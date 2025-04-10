// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./orgMember";
export * from "./orgNode";
export * from "./policySubAccountAttachment";

// Import resources to register:
import { OrgMember } from "./orgMember";
import { OrgNode } from "./orgNode";
import { PolicySubAccountAttachment } from "./policySubAccountAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Organization/orgMember:OrgMember":
                return new OrgMember(name, <any>undefined, { urn })
            case "tencentcloud:Organization/orgNode:OrgNode":
                return new OrgNode(name, <any>undefined, { urn })
            case "tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment":
                return new PolicySubAccountAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Organization/orgMember", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Organization/orgNode", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Organization/policySubAccountAttachment", _module)
