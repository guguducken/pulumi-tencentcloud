// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./extraConfig";
export * from "./getInstances";
export * from "./instance";

// Import resources to register:
import { ExtraConfig } from "./extraConfig";
import { Instance } from "./instance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Dcx/extraConfig:ExtraConfig":
                return new ExtraConfig(name, <any>undefined, { urn })
            case "tencentcloud:Dcx/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Dcx/extraConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dcx/instance", _module)
