// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ckafka region
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const region = pulumi.output(tencentcloud.Ckafka.getRegion());
 * ```
 */
export function getRegion(args?: GetRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Ckafka/getRegion:getRegion", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getRegion.
 */
export interface GetRegionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Return a list of region enumeration results.
     */
    readonly results: outputs.Ckafka.GetRegionResult[];
}

export function getRegionOutput(args?: GetRegionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionResult> {
    return pulumi.output(args).apply(a => getRegion(a, opts))
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
