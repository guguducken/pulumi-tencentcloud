// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of as instances
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const scalingGroup = new tencentcloud.as.ScalingGroup("scalingGroup", {
 *     scalingGroupName: "tf-as-group-ds-ins-basic",
 *     configurationId: "your_launch_configuration_id",
 *     maxSize: 1,
 *     minSize: 1,
 *     vpcId: "your_vpc_id",
 *     subnetIds: ["your_subnet_id"],
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const instances = tencentcloud.As.getInstancesOutput({
 *     filters: [{
 *         name: "auto-scaling-group-id",
 *         values: [scalingGroup.id],
 *     }],
 * });
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:As/getInstances:getInstances", {
        "filters": args.filters,
        "instanceIds": args.instanceIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Filter conditions. If there are multiple Filters, the relationship between Filters is a logical AND (AND) relationship. If there are multiple Values in the same Filter, the relationship between Values under the same Filter is a logical OR (OR) relationship.
     */
    filters?: inputs.As.GetInstancesFilter[];
    /**
     * Instance ID of the cloud server (CVM) to be queried. The limit is 100 per request.
     */
    instanceIds?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly filters?: outputs.As.GetInstancesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    /**
     * List of instance details.
     */
    readonly instanceLists: outputs.As.GetInstancesInstanceList[];
    readonly resultOutputFile?: string;
}

export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply(a => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * Filter conditions. If there are multiple Filters, the relationship between Filters is a logical AND (AND) relationship. If there are multiple Values in the same Filter, the relationship between Values under the same Filter is a logical OR (OR) relationship.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.As.GetInstancesFilterArgs>[]>;
    /**
     * Instance ID of the cloud server (CVM) to be queried. The limit is 100 per request.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
