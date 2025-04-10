// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cvm chcHosts
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const chcHosts = pulumi.output(tencentcloud.Cvm.getChcHosts({
 *     chcIds: ["chc-xxxxxx"],
 *     filters: [{
 *         name: "zone",
 *         values: ["ap-guangzhou-7"],
 *     }],
 * }));
 * ```
 */
export function getChcHosts(args?: GetChcHostsArgs, opts?: pulumi.InvokeOptions): Promise<GetChcHostsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cvm/getChcHosts:getChcHosts", {
        "chcIds": args.chcIds,
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getChcHosts.
 */
export interface GetChcHostsArgs {
    /**
     * CHC host ID. Up to 100 instances per request is allowed. ChcIds and Filters cannot be specified at the same time.
     */
    chcIds?: string[];
    /**
     * - `zone` Filter by the availability zone, such as ap-guangzhou-1. Valid values: See [Regions and Availability Zones](https://www.tencentcloud.com/document/product/213/6091?from_cn_redirect=1).
     * - `instance-name` Filter by the instance name.
     * - `instance-state` Filter by the instance status. For status details, see [InstanceStatus](https://www.tencentcloud.com/document/api/213/15753?from_cn_redirect=1#InstanceStatus).
     * - `device-type` Filter by the device type.
     * - `vpc-id` Filter by the unique VPC ID.
     * - `subnet-id` Filter by the unique VPC subnet ID.
     */
    filters?: inputs.Cvm.GetChcHostsFilter[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getChcHosts.
 */
export interface GetChcHostsResult {
    /**
     * List of returned instances.
     */
    readonly chcHostSets: outputs.Cvm.GetChcHostsChcHostSet[];
    readonly chcIds?: string[];
    readonly filters?: outputs.Cvm.GetChcHostsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function getChcHostsOutput(args?: GetChcHostsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChcHostsResult> {
    return pulumi.output(args).apply(a => getChcHosts(a, opts))
}

/**
 * A collection of arguments for invoking getChcHosts.
 */
export interface GetChcHostsOutputArgs {
    /**
     * CHC host ID. Up to 100 instances per request is allowed. ChcIds and Filters cannot be specified at the same time.
     */
    chcIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * - `zone` Filter by the availability zone, such as ap-guangzhou-1. Valid values: See [Regions and Availability Zones](https://www.tencentcloud.com/document/product/213/6091?from_cn_redirect=1).
     * - `instance-name` Filter by the instance name.
     * - `instance-state` Filter by the instance status. For status details, see [InstanceStatus](https://www.tencentcloud.com/document/api/213/15753?from_cn_redirect=1#InstanceStatus).
     * - `device-type` Filter by the device type.
     * - `vpc-id` Filter by the unique VPC ID.
     * - `subnet-id` Filter by the unique VPC subnet ID.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Cvm.GetChcHostsFilterArgs>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
