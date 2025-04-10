// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of clb idleLoadbalancers
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const idleInstance = pulumi.output(tencentcloud.Clb.getIdleInstances({
 *     loadBalancerRegion: "ap-guangzhou",
 * }));
 * ```
 */
export function getIdleInstances(args?: GetIdleInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetIdleInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Clb/getIdleInstances:getIdleInstances", {
        "loadBalancerRegion": args.loadBalancerRegion,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdleInstances.
 */
export interface GetIdleInstancesArgs {
    /**
     * CLB instance region.
     */
    loadBalancerRegion?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getIdleInstances.
 */
export interface GetIdleInstancesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of idle CLBs. Note: This field may return null, indicating that no valid values can be obtained.
     */
    readonly idleLoadBalancers: outputs.Clb.GetIdleInstancesIdleLoadBalancer[];
    readonly loadBalancerRegion?: string;
    readonly resultOutputFile?: string;
}

export function getIdleInstancesOutput(args?: GetIdleInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdleInstancesResult> {
    return pulumi.output(args).apply(a => getIdleInstances(a, opts))
}

/**
 * A collection of arguments for invoking getIdleInstances.
 */
export interface GetIdleInstancesOutputArgs {
    /**
     * CLB instance region.
     */
    loadBalancerRegion?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
