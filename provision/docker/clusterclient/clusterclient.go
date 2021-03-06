// Copyright 2017 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clusterclient

import (
	"io"
	"time"

	"github.com/fsouza/go-dockerclient"
	"github.com/pkg/errors"
	"github.com/tsuru/docker-cluster/cluster"
	"github.com/tsuru/tsuru/db/storage"
	tsuruIo "github.com/tsuru/tsuru/io"
	"github.com/tsuru/tsuru/log"
	"github.com/tsuru/tsuru/net"
	"github.com/tsuru/tsuru/provision"
	"github.com/tsuru/tsuru/provision/docker/container"
	"github.com/tsuru/tsuru/provision/dockercommon"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ClusterClient struct {
	*cluster.Cluster
	Collection    func() *storage.Collection
	Limiter       provision.ActionLimiter
	PossibleNodes []string
}

var (
	_ provision.BuilderDockerClient  = &ClusterClient{}
	_ provision.ExecDockerClient     = &ClusterClient{}
	_ container.ContainerStateClient = &ClusterClient{}
)

func (c *ClusterClient) SetTimeout(time.Duration) {
	// noop, cluster already handles timeouts per operation correctly
}

func (c *ClusterClient) PullAndCreateContainer(opts docker.CreateContainerOptions, w io.Writer) (cont *docker.Container, hostAddr string, err error) {
	var dbCont *container.Container
	if opts.Context != nil {
		dbCont, _ = opts.Context.Value(container.ContainerCtxKey{}).(*container.Container)
	}
	if dbCont == nil {
		// No need to register in db as BS won't associate this container with
		// tsuru.
		schedulerOpts := &container.SchedulerOpts{
			FilterNodes: c.PossibleNodes,
		}
		hostAddr, cont, err = c.Cluster.CreateContainerSchedulerOpts(opts, schedulerOpts, net.StreamInactivityTimeout)
		hostAddr = net.URLToHost(hostAddr)
		return cont, hostAddr, err
	}
	coll := c.Collection()
	defer coll.Close()
	defer func() {
		if err == nil {
			return
		}
		dbErr := coll.Remove(bson.M{"name": dbCont.Name})
		if dbErr != nil && dbErr != mgo.ErrNotFound {
			log.Errorf("error trying to remove container in db after failure %#v: %v", cont, dbErr)
		}
		if cont != nil {
			removeErr := c.Cluster.RemoveContainer(docker.RemoveContainerOptions{
				ID:            cont.ID,
				RemoveVolumes: true,
				Force:         true,
			})
			if removeErr != nil {
				log.Errorf("error trying to remove container in docker after update failure %#v: %v", cont, removeErr)
			}
		}
	}()
	err = coll.Insert(dbCont)
	if err != nil {
		return nil, "", err
	}
	schedulerOpts := &container.SchedulerOpts{
		AppName:       dbCont.AppName,
		ProcessName:   dbCont.ProcessName,
		UpdateName:    true,
		ActionLimiter: c.Limiter,
		FilterNodes:   c.PossibleNodes,
	}
	var addr string
	pullOpts := docker.PullImageOptions{
		Repository:        opts.Config.Image,
		InactivityTimeout: net.StreamInactivityTimeout,
	}
	if w != nil {
		pullOpts.OutputStream = &tsuruIo.DockerErrorCheckWriter{W: w}
		pullOpts.RawJSONStream = true
	}
	var nodes []string
	if dbCont.HostAddr != "" {
		var node cluster.Node
		node, err = dockercommon.GetNodeByHost(c.Cluster, dbCont.HostAddr)
		if err != nil {
			return nil, "", err
		}
		nodes = []string{node.Address}
	}
	addr, cont, err = c.Cluster.CreateContainerPullOptsSchedulerOpts(
		opts,
		pullOpts,
		dockercommon.RegistryAuthConfig(),
		schedulerOpts,
		nodes...,
	)
	if schedulerOpts.LimiterDone != nil {
		schedulerOpts.LimiterDone()
	}
	if err != nil {
		return nil, "", err
	}
	hostAddr = net.URLToHost(addr)
	err = coll.Update(bson.M{"name": cont.Name}, bson.M{"$set": bson.M{
		"id":       cont.ID,
		"hostaddr": hostAddr,
	}})
	if err != nil {
		return nil, "", errors.Wrap(err, "unable to update container ID in db")
	}
	return cont, hostAddr, nil
}

func (c *ClusterClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	err := c.Cluster.RemoveContainer(opts)
	if err != nil {
		log.Errorf("error trying to remove container %q: %v", opts.ID, err)
	}
	coll := c.Collection()
	defer coll.Close()
	dbErr := coll.Remove(bson.M{"id": opts.ID})
	if dbErr != nil && dbErr != mgo.ErrNotFound {
		log.Errorf("error trying to remove container in db %q: %v", opts.ID, dbErr)
	}
	return err
}

func (sc *ClusterClient) SetContainerState(c *container.Container, state container.ContainerState) error {
	coll := sc.Collection()
	defer coll.Close()
	switch state {
	case container.ContainerStateNewStatus:
		return coll.Update(
			bson.M{"id": c.ID, "status": bson.M{"$ne": provision.StatusBuilding.String()}},
			bson.M{"$set": bson.M{
				"status":                  c.Status,
				"statusbeforeerror":       c.StatusBeforeError,
				"laststatusupdate":        c.LastStatusUpdate,
				"lastsuccessstatusupdate": c.LastSuccessStatusUpdate,
			}},
		)
	case container.ContainerStateRemoved:
		return coll.Remove(bson.M{"id": c.ID})
	default:
		return coll.Update(bson.M{"id": c.ID}, c)
	}
}
