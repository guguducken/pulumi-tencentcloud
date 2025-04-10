// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tcr replicationInstanceCreateTasks
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const createTasks = tencentcloud.Tcr.getReplicationInstanceCreateTasks({
 *     replicationRegistryId: local.dst_registry_id,
 *     replicationRegionId: local.dst_region_id,
 * });
 * ```
 */
export function getReplicationInstanceCreateTasks(args: GetReplicationInstanceCreateTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationInstanceCreateTasksResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tcr/getReplicationInstanceCreateTasks:getReplicationInstanceCreateTasks", {
        "replicationRegionId": args.replicationRegionId,
        "replicationRegistryId": args.replicationRegistryId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getReplicationInstanceCreateTasks.
 */
export interface GetReplicationInstanceCreateTasksArgs {
    /**
     * synchronization instance region Id, see ReplicationRegionId in DescribeReplicationInstances.
     */
    replicationRegionId: number;
    /**
     * synchronization instance Id, see RegistryId in DescribeReplicationInstances.
     */
    replicationRegistryId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getReplicationInstanceCreateTasks.
 */
export interface GetReplicationInstanceCreateTasksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly replicationRegionId: number;
    readonly replicationRegistryId: string;
    readonly resultOutputFile?: string;
    /**
     * overall task status.
     */
    readonly status: string;
    /**
     * task details.
     */
    readonly taskDetails: outputs.Tcr.GetReplicationInstanceCreateTasksTaskDetail[];
}

export function getReplicationInstanceCreateTasksOutput(args: GetReplicationInstanceCreateTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationInstanceCreateTasksResult> {
    return pulumi.output(args).apply(a => getReplicationInstanceCreateTasks(a, opts))
}

/**
 * A collection of arguments for invoking getReplicationInstanceCreateTasks.
 */
export interface GetReplicationInstanceCreateTasksOutputArgs {
    /**
     * synchronization instance region Id, see ReplicationRegionId in DescribeReplicationInstances.
     */
    replicationRegionId: pulumi.Input<number>;
    /**
     * synchronization instance Id, see RegistryId in DescribeReplicationInstances.
     */
    replicationRegistryId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
