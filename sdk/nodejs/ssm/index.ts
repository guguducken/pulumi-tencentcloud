// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getSecretVersions";
export * from "./getSecrets";
export * from "./secret";
export * from "./secretVersion";
export * from "./sshKeyPairSecret";

// Import resources to register:
import { Secret } from "./secret";
import { SecretVersion } from "./secretVersion";
import { SshKeyPairSecret } from "./sshKeyPairSecret";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Ssm/secret:Secret":
                return new Secret(name, <any>undefined, { urn })
            case "tencentcloud:Ssm/secretVersion:SecretVersion":
                return new SecretVersion(name, <any>undefined, { urn })
            case "tencentcloud:Ssm/sshKeyPairSecret:SshKeyPairSecret":
                return new SshKeyPairSecret(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/secret", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/secretVersion", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/sshKeyPairSecret", _module)
