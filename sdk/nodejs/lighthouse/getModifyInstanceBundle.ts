// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of lighthouse modifyInstanceBundle
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const modifyInstanceBundle = pulumi.output(tencentcloud.Lighthouse.getModifyInstanceBundle({
 *     filters: [{
 *         name: "bundle-id",
 *         values: ["bundle_gen_mc_med2_02"],
 *     }],
 *     instanceId: "lhins-xxxxxx",
 * }));
 * ```
 */
export function getModifyInstanceBundle(args: GetModifyInstanceBundleArgs, opts?: pulumi.InvokeOptions): Promise<GetModifyInstanceBundleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Lighthouse/getModifyInstanceBundle:getModifyInstanceBundle", {
        "filters": args.filters,
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getModifyInstanceBundle.
 */
export interface GetModifyInstanceBundleArgs {
    /**
     * Filter list.
     * - `bundle-id`: filter by the bundle ID.
     * - `support-platform-type`: filter by system type, valid values: `LINUX_UNIX`, `WINDOWS`.
     * - `bundle-type`: filter according to package type, valid values: `GENERAL_BUNDLE`, `STORAGE_BUNDLE`, `ENTERPRISE_BUNDLE`, `EXCLUSIVE_BUNDLE`, `BEFAST_BUNDLE`.
     * - `bundle-state`: filter according to package status, valid values: `ONLINE`, `OFFLINE`.
     * NOTE: The upper limit of Filters per request is 10. The upper limit of Filter.Values is 5. Parameter does not support specifying both BundleIds and Filters.
     */
    filters?: inputs.Lighthouse.GetModifyInstanceBundleFilter[];
    /**
     * Instance ID.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getModifyInstanceBundle.
 */
export interface GetModifyInstanceBundleResult {
    readonly filters?: outputs.Lighthouse.GetModifyInstanceBundleFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Change package details.
     */
    readonly modifyBundleSets: outputs.Lighthouse.GetModifyInstanceBundleModifyBundleSet[];
    readonly resultOutputFile?: string;
}

export function getModifyInstanceBundleOutput(args: GetModifyInstanceBundleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModifyInstanceBundleResult> {
    return pulumi.output(args).apply(a => getModifyInstanceBundle(a, opts))
}

/**
 * A collection of arguments for invoking getModifyInstanceBundle.
 */
export interface GetModifyInstanceBundleOutputArgs {
    /**
     * Filter list.
     * - `bundle-id`: filter by the bundle ID.
     * - `support-platform-type`: filter by system type, valid values: `LINUX_UNIX`, `WINDOWS`.
     * - `bundle-type`: filter according to package type, valid values: `GENERAL_BUNDLE`, `STORAGE_BUNDLE`, `ENTERPRISE_BUNDLE`, `EXCLUSIVE_BUNDLE`, `BEFAST_BUNDLE`.
     * - `bundle-state`: filter according to package status, valid values: `ONLINE`, `OFFLINE`.
     * NOTE: The upper limit of Filters per request is 10. The upper limit of Filter.Values is 5. Parameter does not support specifying both BundleIds and Filters.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Lighthouse.GetModifyInstanceBundleFilterArgs>[]>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
