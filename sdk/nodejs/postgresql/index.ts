// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./backupDownloadRestrictionConfig";
export * from "./backupPlanConfig";
export * from "./baseBackup";
export * from "./deleteLogBackupOperation";
export * from "./disisolateDbInstanceOperation";
export * from "./getBackupDownloadUrls";
export * from "./getBaseBackups";
export * from "./getDbInstanceClasses";
export * from "./getDbInstanceVersions";
export * from "./getDefaultParameters";
export * from "./getInstances";
export * from "./getLogBackups";
export * from "./getParameterTemplates";
export * from "./getReadonlyGroups";
export * from "./getRecoveryTime";
export * from "./getRegions";
export * from "./getSpecinfos";
export * from "./getXlogs";
export * from "./getZones";
export * from "./instance";
export * from "./isolateDbInstanceOperation";
export * from "./modifyAccountRemarkOperation";
export * from "./modifySwitchTimePeriodOperation";
export * from "./parameterTemplate";
export * from "./readonlyAttachment";
export * from "./readonlyGroup";
export * from "./readonlyInstance";
export * from "./rebalanceReadonlyGroupOperation";
export * from "./renewDbInstanceOperation";
export * from "./restartDbInstanceOperation";
export * from "./securityGroupConfig";

// Import resources to register:
import { BackupDownloadRestrictionConfig } from "./backupDownloadRestrictionConfig";
import { BackupPlanConfig } from "./backupPlanConfig";
import { BaseBackup } from "./baseBackup";
import { DeleteLogBackupOperation } from "./deleteLogBackupOperation";
import { DisisolateDbInstanceOperation } from "./disisolateDbInstanceOperation";
import { Instance } from "./instance";
import { IsolateDbInstanceOperation } from "./isolateDbInstanceOperation";
import { ModifyAccountRemarkOperation } from "./modifyAccountRemarkOperation";
import { ModifySwitchTimePeriodOperation } from "./modifySwitchTimePeriodOperation";
import { ParameterTemplate } from "./parameterTemplate";
import { ReadonlyAttachment } from "./readonlyAttachment";
import { ReadonlyGroup } from "./readonlyGroup";
import { ReadonlyInstance } from "./readonlyInstance";
import { RebalanceReadonlyGroupOperation } from "./rebalanceReadonlyGroupOperation";
import { RenewDbInstanceOperation } from "./renewDbInstanceOperation";
import { RestartDbInstanceOperation } from "./restartDbInstanceOperation";
import { SecurityGroupConfig } from "./securityGroupConfig";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig":
                return new BackupDownloadRestrictionConfig(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/backupPlanConfig:BackupPlanConfig":
                return new BackupPlanConfig(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/baseBackup:BaseBackup":
                return new BaseBackup(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/deleteLogBackupOperation:DeleteLogBackupOperation":
                return new DeleteLogBackupOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation":
                return new DisisolateDbInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/isolateDbInstanceOperation:IsolateDbInstanceOperation":
                return new IsolateDbInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/modifyAccountRemarkOperation:ModifyAccountRemarkOperation":
                return new ModifyAccountRemarkOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/modifySwitchTimePeriodOperation:ModifySwitchTimePeriodOperation":
                return new ModifySwitchTimePeriodOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/parameterTemplate:ParameterTemplate":
                return new ParameterTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/readonlyAttachment:ReadonlyAttachment":
                return new ReadonlyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/readonlyGroup:ReadonlyGroup":
                return new ReadonlyGroup(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance":
                return new ReadonlyInstance(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/rebalanceReadonlyGroupOperation:RebalanceReadonlyGroupOperation":
                return new RebalanceReadonlyGroupOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation":
                return new RenewDbInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation":
                return new RestartDbInstanceOperation(name, <any>undefined, { urn })
            case "tencentcloud:Postgresql/securityGroupConfig:SecurityGroupConfig":
                return new SecurityGroupConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/backupDownloadRestrictionConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/backupPlanConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/baseBackup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/deleteLogBackupOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/disisolateDbInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/instance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/isolateDbInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/modifyAccountRemarkOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/modifySwitchTimePeriodOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/parameterTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/readonlyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/readonlyGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/readonlyInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/rebalanceReadonlyGroupOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/renewDbInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/restartDbInstanceOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Postgresql/securityGroupConfig", _module)
