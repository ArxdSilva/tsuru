// AUTOMATICALLY GENERATED FILE - DO NOT EDIT!
// Please run 'go generate' to update this file.
//
// Copyright 2017 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package permission

var (
	PermAll                              = PermissionRegistry.get("")                                    // [global]
	PermApp                              = PermissionRegistry.get("app")                                 // [global app team pool]
	PermAppAdmin                         = PermissionRegistry.get("app.admin")                           // [global app team pool]
	PermAppAdminQuota                    = PermissionRegistry.get("app.admin.quota")                     // [global app team pool]
	PermAppAdminRoutes                   = PermissionRegistry.get("app.admin.routes")                    // [global app team pool]
	PermAppAdminUnlock                   = PermissionRegistry.get("app.admin.unlock")                    // [global app team pool]
	PermAppBuild                         = PermissionRegistry.get("app.build")                           // [global app team pool]
	PermAppCreate                        = PermissionRegistry.get("app.create")                          // [global team]
	PermAppDelete                        = PermissionRegistry.get("app.delete")                          // [global app team pool]
	PermAppDeploy                        = PermissionRegistry.get("app.deploy")                          // [global app team pool]
	PermAppDeployArchiveUrl              = PermissionRegistry.get("app.deploy.archive-url")              // [global app team pool]
	PermAppDeployBuild                   = PermissionRegistry.get("app.deploy.build")                    // [global app team pool]
	PermAppDeployGit                     = PermissionRegistry.get("app.deploy.git")                      // [global app team pool]
	PermAppDeployImage                   = PermissionRegistry.get("app.deploy.image")                    // [global app team pool]
	PermAppDeployRollback                = PermissionRegistry.get("app.deploy.rollback")                 // [global app team pool]
	PermAppDeployUpload                  = PermissionRegistry.get("app.deploy.upload")                   // [global app team pool]
	PermAppRead                          = PermissionRegistry.get("app.read")                            // [global app team pool]
	PermAppReadCertificate               = PermissionRegistry.get("app.read.certificate")                // [global app team pool]
	PermAppReadDeploy                    = PermissionRegistry.get("app.read.deploy")                     // [global app team pool]
	PermAppReadEnv                       = PermissionRegistry.get("app.read.env")                        // [global app team pool]
	PermAppReadEvents                    = PermissionRegistry.get("app.read.events")                     // [global app team pool]
	PermAppReadLog                       = PermissionRegistry.get("app.read.log")                        // [global app team pool]
	PermAppReadMetric                    = PermissionRegistry.get("app.read.metric")                     // [global app team pool]
	PermAppReadRouter                    = PermissionRegistry.get("app.read.router")                     // [global app team pool]
	PermAppRun                           = PermissionRegistry.get("app.run")                             // [global app team pool]
	PermAppRunShell                      = PermissionRegistry.get("app.run.shell")                       // [global app team pool]
	PermAppUpdate                        = PermissionRegistry.get("app.update")                          // [global app team pool]
	PermAppUpdateBind                    = PermissionRegistry.get("app.update.bind")                     // [global app team pool]
	PermAppUpdateBindVolume              = PermissionRegistry.get("app.update.bind-volume")              // [global app team pool]
	PermAppUpdateCertificate             = PermissionRegistry.get("app.update.certificate")              // [global app team pool]
	PermAppUpdateCertificateSet          = PermissionRegistry.get("app.update.certificate.set")          // [global app team pool]
	PermAppUpdateCertificateUnset        = PermissionRegistry.get("app.update.certificate.unset")        // [global app team pool]
	PermAppUpdateCname                   = PermissionRegistry.get("app.update.cname")                    // [global app team pool]
	PermAppUpdateCnameAdd                = PermissionRegistry.get("app.update.cname.add")                // [global app team pool]
	PermAppUpdateCnameRemove             = PermissionRegistry.get("app.update.cname.remove")             // [global app team pool]
	PermAppUpdateDeploy                  = PermissionRegistry.get("app.update.deploy")                   // [global app team pool]
	PermAppUpdateDeployRollback          = PermissionRegistry.get("app.update.deploy.rollback")          // [global app team pool]
	PermAppUpdateDescription             = PermissionRegistry.get("app.update.description")              // [global app team pool]
	PermAppUpdateEnv                     = PermissionRegistry.get("app.update.env")                      // [global app team pool]
	PermAppUpdateEnvSet                  = PermissionRegistry.get("app.update.env.set")                  // [global app team pool]
	PermAppUpdateEnvUnset                = PermissionRegistry.get("app.update.env.unset")                // [global app team pool]
	PermAppUpdateEvents                  = PermissionRegistry.get("app.update.events")                   // [global app team pool]
	PermAppUpdateGrant                   = PermissionRegistry.get("app.update.grant")                    // [global app team pool]
	PermAppUpdateImageReset              = PermissionRegistry.get("app.update.image-reset")              // [global app team pool]
	PermAppUpdateLog                     = PermissionRegistry.get("app.update.log")                      // [global app team pool]
	PermAppUpdatePlan                    = PermissionRegistry.get("app.update.plan")                     // [global app team pool]
	PermAppUpdatePlatform                = PermissionRegistry.get("app.update.platform")                 // [global app team pool]
	PermAppUpdatePool                    = PermissionRegistry.get("app.update.pool")                     // [global app team pool]
	PermAppUpdateRestart                 = PermissionRegistry.get("app.update.restart")                  // [global app team pool]
	PermAppUpdateRevoke                  = PermissionRegistry.get("app.update.revoke")                   // [global app team pool]
	PermAppUpdateRouter                  = PermissionRegistry.get("app.update.router")                   // [global app team pool]
	PermAppUpdateRouterAdd               = PermissionRegistry.get("app.update.router.add")               // [global app team pool]
	PermAppUpdateRouterRemove            = PermissionRegistry.get("app.update.router.remove")            // [global app team pool]
	PermAppUpdateRouterUpdate            = PermissionRegistry.get("app.update.router.update")            // [global app team pool]
	PermAppUpdateSleep                   = PermissionRegistry.get("app.update.sleep")                    // [global app team pool]
	PermAppUpdateStart                   = PermissionRegistry.get("app.update.start")                    // [global app team pool]
	PermAppUpdateStop                    = PermissionRegistry.get("app.update.stop")                     // [global app team pool]
	PermAppUpdateSwap                    = PermissionRegistry.get("app.update.swap")                     // [global app team pool]
	PermAppUpdateTags                    = PermissionRegistry.get("app.update.tags")                     // [global app team pool]
	PermAppUpdateTeamowner               = PermissionRegistry.get("app.update.teamowner")                // [global app team pool]
	PermAppUpdateUnbind                  = PermissionRegistry.get("app.update.unbind")                   // [global app team pool]
	PermAppUpdateUnbindVolume            = PermissionRegistry.get("app.update.unbind-volume")            // [global app team pool]
	PermAppUpdateUnit                    = PermissionRegistry.get("app.update.unit")                     // [global app team pool]
	PermAppUpdateUnitAdd                 = PermissionRegistry.get("app.update.unit.add")                 // [global app team pool]
	PermAppUpdateUnitRegister            = PermissionRegistry.get("app.update.unit.register")            // [global app team pool]
	PermAppUpdateUnitRemove              = PermissionRegistry.get("app.update.unit.remove")              // [global app team pool]
	PermAppUpdateUnitStatus              = PermissionRegistry.get("app.update.unit.status")              // [global app team pool]
	PermCluster                          = PermissionRegistry.get("cluster")                             // [global]
	PermClusterCreate                    = PermissionRegistry.get("cluster.create")                      // [global]
	PermClusterDelete                    = PermissionRegistry.get("cluster.delete")                      // [global]
	PermClusterRead                      = PermissionRegistry.get("cluster.read")                        // [global]
	PermClusterReadEvents                = PermissionRegistry.get("cluster.read.events")                 // [global]
	PermClusterUpdate                    = PermissionRegistry.get("cluster.update")                      // [global]
	PermDebug                            = PermissionRegistry.get("debug")                               // [global]
	PermEventBlock                       = PermissionRegistry.get("event-block")                         // [global]
	PermEventBlockAdd                    = PermissionRegistry.get("event-block.add")                     // [global]
	PermEventBlockRead                   = PermissionRegistry.get("event-block.read")                    // [global]
	PermEventBlockReadEvents             = PermissionRegistry.get("event-block.read.events")             // [global]
	PermEventBlockRemove                 = PermissionRegistry.get("event-block.remove")                  // [global]
	PermHealing                          = PermissionRegistry.get("healing")                             // [global pool]
	PermHealingDelete                    = PermissionRegistry.get("healing.delete")                      // [global pool]
	PermHealingRead                      = PermissionRegistry.get("healing.read")                        // [global pool]
	PermHealingUpdate                    = PermissionRegistry.get("healing.update")                      // [global pool]
	PermInstall                          = PermissionRegistry.get("install")                             // [global]
	PermInstallManage                    = PermissionRegistry.get("install.manage")                      // [global]
	PermMachine                          = PermissionRegistry.get("machine")                             // [global iaas]
	PermMachineDelete                    = PermissionRegistry.get("machine.delete")                      // [global iaas]
	PermMachineRead                      = PermissionRegistry.get("machine.read")                        // [global iaas]
	PermMachineReadEvents                = PermissionRegistry.get("machine.read.events")                 // [global iaas]
	PermMachineTemplate                  = PermissionRegistry.get("machine.template")                    // [global iaas]
	PermMachineTemplateCreate            = PermissionRegistry.get("machine.template.create")             // [global iaas]
	PermMachineTemplateDelete            = PermissionRegistry.get("machine.template.delete")             // [global iaas]
	PermMachineTemplateRead              = PermissionRegistry.get("machine.template.read")               // [global iaas]
	PermMachineTemplateUpdate            = PermissionRegistry.get("machine.template.update")             // [global iaas]
	PermNode                             = PermissionRegistry.get("node")                                // [global pool]
	PermNodeAutoscale                    = PermissionRegistry.get("node.autoscale")                      // [global]
	PermNodeAutoscaleDelete              = PermissionRegistry.get("node.autoscale.delete")               // [global]
	PermNodeAutoscaleRead                = PermissionRegistry.get("node.autoscale.read")                 // [global]
	PermNodeAutoscaleUpdate              = PermissionRegistry.get("node.autoscale.update")               // [global]
	PermNodeAutoscaleUpdateRun           = PermissionRegistry.get("node.autoscale.update.run")           // [global]
	PermNodeCreate                       = PermissionRegistry.get("node.create")                         // [global pool]
	PermNodeDelete                       = PermissionRegistry.get("node.delete")                         // [global pool]
	PermNodeRead                         = PermissionRegistry.get("node.read")                           // [global pool]
	PermNodeUpdate                       = PermissionRegistry.get("node.update")                         // [global pool]
	PermNodeUpdateMove                   = PermissionRegistry.get("node.update.move")                    // [global pool]
	PermNodeUpdateMoveContainer          = PermissionRegistry.get("node.update.move.container")          // [global pool]
	PermNodeUpdateMoveContainers         = PermissionRegistry.get("node.update.move.containers")         // [global pool]
	PermNodeUpdateRebalance              = PermissionRegistry.get("node.update.rebalance")               // [global pool]
	PermNodecontainer                    = PermissionRegistry.get("nodecontainer")                       // [global pool]
	PermNodecontainerCreate              = PermissionRegistry.get("nodecontainer.create")                // [global pool]
	PermNodecontainerDelete              = PermissionRegistry.get("nodecontainer.delete")                // [global pool]
	PermNodecontainerRead                = PermissionRegistry.get("nodecontainer.read")                  // [global pool]
	PermNodecontainerUpdate              = PermissionRegistry.get("nodecontainer.update")                // [global pool]
	PermNodecontainerUpdateUpgrade       = PermissionRegistry.get("nodecontainer.update.upgrade")        // [global pool]
	PermPlan                             = PermissionRegistry.get("plan")                                // [global]
	PermPlanCreate                       = PermissionRegistry.get("plan.create")                         // [global]
	PermPlanDelete                       = PermissionRegistry.get("plan.delete")                         // [global]
	PermPlanRead                         = PermissionRegistry.get("plan.read")                           // [global]
	PermPlanReadEvents                   = PermissionRegistry.get("plan.read.events")                    // [global]
	PermPlatform                         = PermissionRegistry.get("platform")                            // [global]
	PermPlatformCreate                   = PermissionRegistry.get("platform.create")                     // [global]
	PermPlatformDelete                   = PermissionRegistry.get("platform.delete")                     // [global]
	PermPlatformRead                     = PermissionRegistry.get("platform.read")                       // [global]
	PermPlatformReadEvents               = PermissionRegistry.get("platform.read.events")                // [global]
	PermPlatformUpdate                   = PermissionRegistry.get("platform.update")                     // [global]
	PermPool                             = PermissionRegistry.get("pool")                                // [global pool]
	PermPoolCreate                       = PermissionRegistry.get("pool.create")                         // [global]
	PermPoolDelete                       = PermissionRegistry.get("pool.delete")                         // [global pool]
	PermPoolRead                         = PermissionRegistry.get("pool.read")                           // [global pool]
	PermPoolReadConstraints              = PermissionRegistry.get("pool.read.constraints")               // [global pool]
	PermPoolReadEvents                   = PermissionRegistry.get("pool.read.events")                    // [global pool]
	PermPoolUpdate                       = PermissionRegistry.get("pool.update")                         // [global pool]
	PermPoolUpdateConstraints            = PermissionRegistry.get("pool.update.constraints")             // [global pool]
	PermPoolUpdateConstraintsSet         = PermissionRegistry.get("pool.update.constraints.set")         // [global pool]
	PermPoolUpdateLogs                   = PermissionRegistry.get("pool.update.logs")                    // [global pool]
	PermPoolUpdateTeam                   = PermissionRegistry.get("pool.update.team")                    // [global pool]
	PermPoolUpdateTeamAdd                = PermissionRegistry.get("pool.update.team.add")                // [global pool]
	PermPoolUpdateTeamRemove             = PermissionRegistry.get("pool.update.team.remove")             // [global pool]
	PermRole                             = PermissionRegistry.get("role")                                // [global]
	PermRoleCreate                       = PermissionRegistry.get("role.create")                         // [global]
	PermRoleDefault                      = PermissionRegistry.get("role.default")                        // [global]
	PermRoleDefaultCreate                = PermissionRegistry.get("role.default.create")                 // [global]
	PermRoleDefaultDelete                = PermissionRegistry.get("role.default.delete")                 // [global]
	PermRoleDelete                       = PermissionRegistry.get("role.delete")                         // [global]
	PermRoleRead                         = PermissionRegistry.get("role.read")                           // [global]
	PermRoleReadEvents                   = PermissionRegistry.get("role.read.events")                    // [global]
	PermRoleUpdate                       = PermissionRegistry.get("role.update")                         // [global]
	PermRoleUpdateAssign                 = PermissionRegistry.get("role.update.assign")                  // [global]
	PermRoleUpdateContext                = PermissionRegistry.get("role.update.context")                 // [global]
	PermRoleUpdateContextType            = PermissionRegistry.get("role.update.context.type")            // [global]
	PermRoleUpdateDescription            = PermissionRegistry.get("role.update.description")             // [global]
	PermRoleUpdateDissociate             = PermissionRegistry.get("role.update.dissociate")              // [global]
	PermRoleUpdateName                   = PermissionRegistry.get("role.update.name")                    // [global]
	PermRoleUpdatePermission             = PermissionRegistry.get("role.update.permission")              // [global]
	PermRoleUpdatePermissionAdd          = PermissionRegistry.get("role.update.permission.add")          // [global]
	PermRoleUpdatePermissionRemove       = PermissionRegistry.get("role.update.permission.remove")       // [global]
	PermService                          = PermissionRegistry.get("service")                             // [global service team]
	PermServiceInstance                  = PermissionRegistry.get("service-instance")                    // [global service-instance team]
	PermServiceInstanceCreate            = PermissionRegistry.get("service-instance.create")             // [global team]
	PermServiceInstanceDelete            = PermissionRegistry.get("service-instance.delete")             // [global service-instance team]
	PermServiceInstanceRead              = PermissionRegistry.get("service-instance.read")               // [global service-instance team]
	PermServiceInstanceReadEvents        = PermissionRegistry.get("service-instance.read.events")        // [global service-instance team]
	PermServiceInstanceReadStatus        = PermissionRegistry.get("service-instance.read.status")        // [global service-instance team]
	PermServiceInstanceUpdate            = PermissionRegistry.get("service-instance.update")             // [global service-instance team]
	PermServiceInstanceUpdateBind        = PermissionRegistry.get("service-instance.update.bind")        // [global service-instance team]
	PermServiceInstanceUpdateDescription = PermissionRegistry.get("service-instance.update.description") // [global service-instance team]
	PermServiceInstanceUpdateGrant       = PermissionRegistry.get("service-instance.update.grant")       // [global service-instance team]
	PermServiceInstanceUpdatePlan        = PermissionRegistry.get("service-instance.update.plan")        // [global service-instance team]
	PermServiceInstanceUpdateProxy       = PermissionRegistry.get("service-instance.update.proxy")       // [global service-instance team]
	PermServiceInstanceUpdateRevoke      = PermissionRegistry.get("service-instance.update.revoke")      // [global service-instance team]
	PermServiceInstanceUpdateTags        = PermissionRegistry.get("service-instance.update.tags")        // [global service-instance team]
	PermServiceInstanceUpdateTeamowner   = PermissionRegistry.get("service-instance.update.teamowner")   // [global service-instance team]
	PermServiceInstanceUpdateUnbind      = PermissionRegistry.get("service-instance.update.unbind")      // [global service-instance team]
	PermServiceCreate                    = PermissionRegistry.get("service.create")                      // [global team]
	PermServiceDelete                    = PermissionRegistry.get("service.delete")                      // [global service team]
	PermServiceRead                      = PermissionRegistry.get("service.read")                        // [global service team]
	PermServiceReadDoc                   = PermissionRegistry.get("service.read.doc")                    // [global service team]
	PermServiceReadEvents                = PermissionRegistry.get("service.read.events")                 // [global service team]
	PermServiceReadPlans                 = PermissionRegistry.get("service.read.plans")                  // [global service team]
	PermServiceUpdate                    = PermissionRegistry.get("service.update")                      // [global service team]
	PermServiceUpdateDoc                 = PermissionRegistry.get("service.update.doc")                  // [global service team]
	PermServiceUpdateGrantAccess         = PermissionRegistry.get("service.update.grant-access")         // [global service team]
	PermServiceUpdateProxy               = PermissionRegistry.get("service.update.proxy")                // [global service team]
	PermServiceUpdateRevokeAccess        = PermissionRegistry.get("service.update.revoke-access")        // [global service team]
	PermTeam                             = PermissionRegistry.get("team")                                // [global team]
	PermTeamCreate                       = PermissionRegistry.get("team.create")                         // [global]
	PermTeamDelete                       = PermissionRegistry.get("team.delete")                         // [global team]
	PermTeamInfo                         = PermissionRegistry.get("team.info")                           // [global team]
	PermTeamRead                         = PermissionRegistry.get("team.read")                           // [global team]
	PermTeamReadEvents                   = PermissionRegistry.get("team.read.events")                    // [global team]
	PermTeamUpdate                       = PermissionRegistry.get("team.update")                         // [global team]
	PermUser                             = PermissionRegistry.get("user")                                // [global user]
	PermUserCreate                       = PermissionRegistry.get("user.create")                         // [global]
	PermUserDelete                       = PermissionRegistry.get("user.delete")                         // [global user]
	PermUserRead                         = PermissionRegistry.get("user.read")                           // [global user]
	PermUserReadEvents                   = PermissionRegistry.get("user.read.events")                    // [global user]
	PermUserUpdate                       = PermissionRegistry.get("user.update")                         // [global user]
	PermUserUpdateKey                    = PermissionRegistry.get("user.update.key")                     // [global user]
	PermUserUpdateKeyAdd                 = PermissionRegistry.get("user.update.key.add")                 // [global user]
	PermUserUpdateKeyRemove              = PermissionRegistry.get("user.update.key.remove")              // [global user]
	PermUserUpdatePassword               = PermissionRegistry.get("user.update.password")                // [global user]
	PermUserUpdateQuota                  = PermissionRegistry.get("user.update.quota")                   // [global user]
	PermUserUpdateReset                  = PermissionRegistry.get("user.update.reset")                   // [global user]
	PermUserUpdateToken                  = PermissionRegistry.get("user.update.token")                   // [global user]
	PermVolume                           = PermissionRegistry.get("volume")                              // [global volume team pool]
	PermVolumeCreate                     = PermissionRegistry.get("volume.create")                       // [global team pool]
	PermVolumeDelete                     = PermissionRegistry.get("volume.delete")                       // [global volume team pool]
	PermVolumeRead                       = PermissionRegistry.get("volume.read")                         // [global volume team pool]
	PermVolumeReadEvents                 = PermissionRegistry.get("volume.read.events")                  // [global volume team pool]
	PermVolumeUpdate                     = PermissionRegistry.get("volume.update")                       // [global volume team pool]
	PermVolumeUpdateBind                 = PermissionRegistry.get("volume.update.bind")                  // [global volume team pool]
	PermVolumeUpdateUnbind               = PermissionRegistry.get("volume.update.unbind")                // [global volume team pool]
)
