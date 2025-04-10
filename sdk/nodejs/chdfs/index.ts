// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessGroup";
export * from "./accessRule";
export * from "./fileSystem";
export * from "./getAccessGroups";
export * from "./getFileSystems";
export * from "./getMountPoints";
export * from "./lifeCycleRule";
export * from "./mountPoint";
export * from "./mountPointAttachment";

// Import resources to register:
import { AccessGroup } from "./accessGroup";
import { AccessRule } from "./accessRule";
import { FileSystem } from "./fileSystem";
import { LifeCycleRule } from "./lifeCycleRule";
import { MountPoint } from "./mountPoint";
import { MountPointAttachment } from "./mountPointAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Chdfs/accessGroup:AccessGroup":
                return new AccessGroup(name, <any>undefined, { urn })
            case "tencentcloud:Chdfs/accessRule:AccessRule":
                return new AccessRule(name, <any>undefined, { urn })
            case "tencentcloud:Chdfs/fileSystem:FileSystem":
                return new FileSystem(name, <any>undefined, { urn })
            case "tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule":
                return new LifeCycleRule(name, <any>undefined, { urn })
            case "tencentcloud:Chdfs/mountPoint:MountPoint":
                return new MountPoint(name, <any>undefined, { urn })
            case "tencentcloud:Chdfs/mountPointAttachment:MountPointAttachment":
                return new MountPointAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Chdfs/accessGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Chdfs/accessRule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Chdfs/fileSystem", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Chdfs/lifeCycleRule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Chdfs/mountPoint", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Chdfs/mountPointAttachment", _module)
