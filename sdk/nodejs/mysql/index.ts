// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./auditLogFile";
export * from "./backupDownloadRestriction";
export * from "./backupEncryptionStatus";
export * from "./backupPolicy";
export * from "./dbImportJobOperation";
export * from "./deployGroup";
export * from "./drInstanceToMater";
export * from "./getBackupList";
export * from "./getBackupOverview";
export * from "./getBackupSummaries";
export * from "./getBinLog";
export * from "./getBinlogBackupOverview";
export * from "./getCloneList";
export * from "./getDataBackupOverview";
export * from "./getDatabases";
export * from "./getDbFeatures";
export * from "./getDefaultParams";
export * from "./getErrorLog";
export * from "./getInstTables";
export * from "./getInstance";
export * from "./getInstanceCharset";
export * from "./getInstanceInfo";
export * from "./getInstanceParamRecord";
export * from "./getInstanceRebootTime";
export * from "./getParameterList";
export * from "./getProjectSecurityGroup";
export * from "./getProxyCustom";
export * from "./getRoMinScale";
export * from "./getRollbackRangeTime";
export * from "./getSlowLog";
export * from "./getSlowLogData";
export * from "./getSupportedPrivileges";
export * from "./getSwitchRecord";
export * from "./getUserTask";
export * from "./getZoneConfig";
export * from "./instance";
export * from "./instanceEncryptionOperation";
export * from "./isolateInstance";
export * from "./localBinlogConfig";
export * from "./paramTemplate";
export * from "./passwordComplexity";
export * from "./privilege";
export * from "./proxy";
export * from "./readonlyInstance";
export * from "./reloadBalanceProxyNode";
export * from "./remoteBackupConfig";
export * from "./renewDbInstanceOperation";
export * from "./resetRootAccount";
export * from "./restartDbInstancesOperation";
export * from "./roGroup";
export * from "./roGroupLoadOperation";
export * from "./roInstanceIp";
export * from "./roStartReplication";
export * from "./roStopReplication";
export * from "./rollback";
export * from "./rollbackStop";
export * from "./securityGroupsAttachment";
export * from "./switchForUpgrade";
export * from "./switchMasterSlaveOperation";
export * from "./switchProxy";
export * from "./timeWindow";
export * from "./verifyRootAccount";

// Import resources to register:
import { Account } from "./account";
import { AuditLogFile } from "./auditLogFile";
import { BackupDownloadRestriction } from "./backupDownloadRestriction";
import { BackupEncryptionStatus } from "./backupEncryptionStatus";
import { BackupPolicy } from "./backupPolicy";
import { DbImportJobOperation } from "./dbImportJobOperation";
import { DeployGroup } from "./deployGroup";
import { DrInstanceToMater } from "./drInstanceToMater";
import { Instance } from "./instance";
import { InstanceEncryptionOperation } from "./instanceEncryptionOperation";
import { IsolateInstance } from "./isolateInstance";
import { LocalBinlogConfig } from "./localBinlogConfig";
import { ParamTemplate } from "./paramTemplate";
import { PasswordComplexity } from "./passwordComplexity";
import { Privilege } from "./privilege";
import { Proxy } from "./proxy";
import { ReadonlyInstance } from "./readonlyInstance";
import { ReloadBalanceProxyNode } from "./reloadBalanceProxyNode";
import { RemoteBackupConfig } from "./remoteBackupConfig";
import { RenewDbInstanceOperation } from "./renewDbInstanceOperation";
import { ResetRootAccount } from "./resetRootAccount";
import { RestartDbInstancesOperation } from "./restartDbInstancesOperation";
import { RoGroup } from "./roGroup";
import { RoGroupLoadOperation } from "./roGroupLoadOperation";
import { RoInstanceIp } from "./roInstanceIp";
import { RoStartReplication } from "./roStartReplication";
import { RoStopReplication } from "./roStopReplication";
import { Rollback } from "./rollback";
import { RollbackStop } from "./rollbackStop";
import { SecurityGroupsAttachment } from "./securityGroupsAttachment";
import { SwitchForUpgrade } from "./switchForUpgrade";
import { SwitchMasterSlaveOperation } from "./switchMasterSlaveOperation";
import { SwitchProxy } from "./switchProxy";
import { TimeWindow } from "./timeWindow";
import { VerifyRootAccount } from "./verifyRootAccount";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Mysql/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/auditLogFile:AuditLogFile":
                return new AuditLogFile(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction":
                return new BackupDownloadRestriction(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/backupEncryptionStatus:BackupEncryptionStatus":
                return new BackupEncryptionStatus(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/backupPolicy:BackupPolicy":
                return new BackupPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/dbImportJobOperation:DbImportJobOperation":
                return new DbImportJobOperation(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/deployGroup:DeployGroup":
                return new DeployGroup(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/drInstanceToMater:DrInstanceToMater":
                return new DrInstanceToMater(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/instanceEncryptionOperation:InstanceEncryptionOperation":
                return new InstanceEncryptionOperation(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/isolateInstance:IsolateInstance":
                return new IsolateInstance(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig":
                return new LocalBinlogConfig(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/paramTemplate:ParamTemplate":
                return new ParamTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/passwordComplexity:PasswordComplexity":
                return new PasswordComplexity(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/privilege:Privilege":
                return new Privilege(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/proxy:Proxy":
                return new Proxy(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/readonlyInstance:ReadonlyInstance":
                return new ReadonlyInstance(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/reloadBalanceProxyNode:ReloadBalanceProxyNode":
                return new ReloadBalanceProxyNode(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/remoteBackupConfig:RemoteBackupConfig":
                return new RemoteBackupConfig(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/renewDbInstanceOperation:RenewDbInstanceOperation":
                return new RenewDbInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/resetRootAccount:ResetRootAccount":
                return new ResetRootAccount(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation":
                return new RestartDbInstancesOperation(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/roGroup:RoGroup":
                return new RoGroup(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/roGroupLoadOperation:RoGroupLoadOperation":
                return new RoGroupLoadOperation(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/roInstanceIp:RoInstanceIp":
                return new RoInstanceIp(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/roStartReplication:RoStartReplication":
                return new RoStartReplication(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/roStopReplication:RoStopReplication":
                return new RoStopReplication(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/rollback:Rollback":
                return new Rollback(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/rollbackStop:RollbackStop":
                return new RollbackStop(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/securityGroupsAttachment:SecurityGroupsAttachment":
                return new SecurityGroupsAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/switchForUpgrade:SwitchForUpgrade":
                return new SwitchForUpgrade(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/switchMasterSlaveOperation:SwitchMasterSlaveOperation":
                return new SwitchMasterSlaveOperation(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/switchProxy:SwitchProxy":
                return new SwitchProxy(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/timeWindow:TimeWindow":
                return new TimeWindow(name, <any>undefined, { urn })
            case "tencentcloud:Mysql/verifyRootAccount:VerifyRootAccount":
                return new VerifyRootAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/account", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/auditLogFile", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/backupDownloadRestriction", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/backupEncryptionStatus", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/backupPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/dbImportJobOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/deployGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/drInstanceToMater", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/instance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/instanceEncryptionOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/isolateInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/localBinlogConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/paramTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/passwordComplexity", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/privilege", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/proxy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/readonlyInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/reloadBalanceProxyNode", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/remoteBackupConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/renewDbInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/resetRootAccount", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/restartDbInstancesOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/roGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/roGroupLoadOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/roInstanceIp", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/roStartReplication", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/roStopReplication", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/rollback", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/rollbackStop", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/securityGroupsAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/switchForUpgrade", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/switchMasterSlaveOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/switchProxy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/timeWindow", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mysql/verifyRootAccount", _module)
