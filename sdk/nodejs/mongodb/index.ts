// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getInstanceBackups";
export * from "./getInstanceConnections";
export * from "./getInstanceCurrentOp";
export * from "./getInstanceParams";
export * from "./getInstanceSlowLog";
export * from "./getInstances";
export * from "./getZoneConfig";
export * from "./instance";
export * from "./instanceAccount";
export * from "./instanceBackup";
export * from "./instanceBackupDownloadTask";
export * from "./shardingInstance";
export * from "./standbyInstance";

// Import resources to register:
import { Instance } from "./instance";
import { InstanceAccount } from "./instanceAccount";
import { InstanceBackup } from "./instanceBackup";
import { InstanceBackupDownloadTask } from "./instanceBackupDownloadTask";
import { ShardingInstance } from "./shardingInstance";
import { StandbyInstance } from "./standbyInstance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Mongodb/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tencentcloud:Mongodb/instanceAccount:InstanceAccount":
                return new InstanceAccount(name, <any>undefined, { urn })
            case "tencentcloud:Mongodb/instanceBackup:InstanceBackup":
                return new InstanceBackup(name, <any>undefined, { urn })
            case "tencentcloud:Mongodb/instanceBackupDownloadTask:InstanceBackupDownloadTask":
                return new InstanceBackupDownloadTask(name, <any>undefined, { urn })
            case "tencentcloud:Mongodb/shardingInstance:ShardingInstance":
                return new ShardingInstance(name, <any>undefined, { urn })
            case "tencentcloud:Mongodb/standbyInstance:StandbyInstance":
                return new StandbyInstance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Mongodb/instance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mongodb/instanceAccount", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mongodb/instanceBackup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mongodb/instanceBackupDownloadTask", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mongodb/shardingInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mongodb/standbyInstance", _module)
