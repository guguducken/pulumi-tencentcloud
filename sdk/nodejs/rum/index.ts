// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getOfflineLogConfig";
export * from "./getProject";
export * from "./getTawInstance";
export * from "./getWhitelist";
export * from "./offlineLogConfigAttachment";
export * from "./project";
export * from "./tawInstance";
export * from "./whitelist";

// Import resources to register:
import { OfflineLogConfigAttachment } from "./offlineLogConfigAttachment";
import { Project } from "./project";
import { TawInstance } from "./tawInstance";
import { Whitelist } from "./whitelist";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Rum/offlineLogConfigAttachment:OfflineLogConfigAttachment":
                return new OfflineLogConfigAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Rum/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "tencentcloud:Rum/tawInstance:TawInstance":
                return new TawInstance(name, <any>undefined, { urn })
            case "tencentcloud:Rum/whitelist:Whitelist":
                return new Whitelist(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Rum/offlineLogConfigAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Rum/project", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Rum/tawInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Rum/whitelist", _module)
