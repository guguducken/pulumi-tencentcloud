// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getTemplateGroups";
export * from "./getTemplates";
export * from "./template";
export * from "./templateGroup";

// Import resources to register:
import { Template } from "./template";
import { TemplateGroup } from "./templateGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Protocol/template:Template":
                return new Template(name, <any>undefined, { urn })
            case "tencentcloud:Protocol/templateGroup:TemplateGroup":
                return new TemplateGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Protocol/template", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Protocol/templateGroup", _module)
