// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dcdb slowLogs
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const slowLogs = tencentcloud.Dcdb.getSlowLogs({
 *     instanceId: local.dcdb_id,
 *     startTime: `%s`,
 *     endTime: `%s`,
 *     shardId: "shard-1b5r04az",
 *     db: "tf_test_db",
 *     orderBy: "query_time_sum",
 *     orderByType: "desc",
 *     slave: 0,
 * });
 * ```
 */
export function getSlowLogs(args: GetSlowLogsArgs, opts?: pulumi.InvokeOptions): Promise<GetSlowLogsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dcdb/getSlowLogs:getSlowLogs", {
        "db": args.db,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "orderBy": args.orderBy,
        "orderByType": args.orderByType,
        "resultOutputFile": args.resultOutputFile,
        "shardId": args.shardId,
        "slave": args.slave,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlowLogs.
 */
export interface GetSlowLogsArgs {
    /**
     * Specific name of the database to be queried.
     */
    db?: string;
    /**
     * Query end time in the format of 2016-08-22 14:55:20.
     */
    endTime?: string;
    /**
     * Instance ID in the format of `tdsqlshard-ow728lmc`.
     */
    instanceId: string;
    /**
     * Sorting metric. Valid values: query_time_sum, query_count.
     */
    orderBy?: string;
    /**
     * Sorting order. Valid values: desc, asc.
     */
    orderByType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Instance shard ID in the format of `shard-rc754ljk`.
     */
    shardId: string;
    /**
     * Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica).
     */
    slave?: number;
    /**
     * Query start time in the format of 2016-07-23 14:55:20.
     */
    startTime: string;
}

/**
 * A collection of values returned by getSlowLogs.
 */
export interface GetSlowLogsResult {
    /**
     * Slow query log data.
     */
    readonly datas: outputs.Dcdb.GetSlowLogsData[];
    /**
     * Database name.
     */
    readonly db?: string;
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Total statement lock time.
     */
    readonly lockTimeSum: number;
    readonly orderBy?: string;
    readonly orderByType?: string;
    /**
     * Total number of statement queries.
     */
    readonly queryCount: number;
    /**
     * Total statement query time.
     */
    readonly queryTimeSum: number;
    readonly resultOutputFile?: string;
    readonly shardId: string;
    readonly slave?: number;
    readonly startTime: string;
}

export function getSlowLogsOutput(args: GetSlowLogsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSlowLogsResult> {
    return pulumi.output(args).apply(a => getSlowLogs(a, opts))
}

/**
 * A collection of arguments for invoking getSlowLogs.
 */
export interface GetSlowLogsOutputArgs {
    /**
     * Specific name of the database to be queried.
     */
    db?: pulumi.Input<string>;
    /**
     * Query end time in the format of 2016-08-22 14:55:20.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Instance ID in the format of `tdsqlshard-ow728lmc`.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Sorting metric. Valid values: query_time_sum, query_count.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Sorting order. Valid values: desc, asc.
     */
    orderByType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Instance shard ID in the format of `shard-rc754ljk`.
     */
    shardId: pulumi.Input<string>;
    /**
     * Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica).
     */
    slave?: pulumi.Input<number>;
    /**
     * Query start time in the format of 2016-07-23 14:55:20.
     */
    startTime: pulumi.Input<string>;
}
