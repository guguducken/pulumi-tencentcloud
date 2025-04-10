// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of teo zoneDDoSPolicy
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zoneDDoSPolicy = pulumi.output(tencentcloud.Teo.getZoneDdosPolicy({
 *     zoneId: "",
 * }));
 * ```
 */
export function getZoneDdosPolicy(args: GetZoneDdosPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneDdosPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Teo/getZoneDdosPolicy:getZoneDdosPolicy", {
        "resultOutputFile": args.resultOutputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getZoneDdosPolicy.
 */
export interface GetZoneDdosPolicyArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Site ID.
     */
    zoneId: string;
}

/**
 * A collection of values returned by getZoneDdosPolicy.
 */
export interface GetZoneDdosPolicyResult {
    /**
     * All subdomain info. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly domains: outputs.Teo.GetZoneDdosPolicyDomain[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Shielded areas of the zone.
     */
    readonly shieldAreas: outputs.Teo.GetZoneDdosPolicyShieldArea[];
    /**
     * Site ID.
     */
    readonly zoneId: string;
}

export function getZoneDdosPolicyOutput(args: GetZoneDdosPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneDdosPolicyResult> {
    return pulumi.output(args).apply(a => getZoneDdosPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getZoneDdosPolicy.
 */
export interface GetZoneDdosPolicyOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Site ID.
     */
    zoneId: pulumi.Input<string>;
}
