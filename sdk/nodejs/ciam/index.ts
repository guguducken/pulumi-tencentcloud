// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./userGroup";
export * from "./userStore";

// Import resources to register:
import { UserGroup } from "./userGroup";
import { UserStore } from "./userStore";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Ciam/userGroup:UserGroup":
                return new UserGroup(name, <any>undefined, { urn })
            case "tencentcloud:Ciam/userStore:UserStore":
                return new UserStore(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Ciam/userGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ciam/userStore", _module)
