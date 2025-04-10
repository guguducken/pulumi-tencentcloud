// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cynosdb instanceSlowQueries
 *
 * ## Example Usage
 * ### Query slow queries of instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const cynosdbClusterId = config.get("cynosdbClusterId") || "default_cynosdb_cluster";
 * const instanceSlowQueries = tencentcloud.Cynosdb.getInstanceSlowQueries({
 *     instanceId: cynosdbClusterId,
 *     startTime: "2023-06-20 23:19:03",
 *     endTime: "2023-06-30 23:19:03",
 *     username: "keep_dts",
 *     host: `%%`,
 *     database: "tf_ci_test",
 *     orderBy: "QueryTime",
 *     orderByType: "desc",
 * });
 * ```
 * ### Query slow queries by time range
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const cynosdbClusterId = config.get("cynosdbClusterId") || "default_cynosdb_cluster";
 * const instanceSlowQueries = tencentcloud.Cynosdb.getInstanceSlowQueries({
 *     instanceId: cynosdbClusterId,
 *     startTime: "2023-06-20 23:19:03",
 *     endTime: "2023-06-30 23:19:03",
 *     orderBy: "QueryTime",
 *     orderByType: "desc",
 * });
 * ```
 * ### Query slow queries by user and db name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const cynosdbClusterId = config.get("cynosdbClusterId") || "default_cynosdb_cluster";
 * const instanceSlowQueries = tencentcloud.Cynosdb.getInstanceSlowQueries({
 *     instanceId: cynosdbClusterId,
 *     username: "keep_dts",
 *     host: `%%`,
 *     database: "tf_ci_test",
 *     orderBy: "QueryTime",
 *     orderByType: "desc",
 * });
 * ```
 */
export function getInstanceSlowQueries(args: GetInstanceSlowQueriesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceSlowQueriesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cynosdb/getInstanceSlowQueries:getInstanceSlowQueries", {
        "database": args.database,
        "endTime": args.endTime,
        "host": args.host,
        "instanceId": args.instanceId,
        "orderBy": args.orderBy,
        "orderByType": args.orderByType,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceSlowQueries.
 */
export interface GetInstanceSlowQueriesArgs {
    /**
     * Database name.
     */
    database?: string;
    /**
     * Latest transaction start time.
     */
    endTime?: string;
    /**
     * Client host.
     */
    host?: string;
    /**
     * Instance ID.
     */
    instanceId: string;
    /**
     * Sort field, optional values: QueryTime, LockTime, RowsExamined, RowsSent.
     */
    orderBy?: string;
    /**
     * Sort type, optional values: asc, desc.
     */
    orderByType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Earliest transaction start time.
     */
    startTime?: string;
    /**
     * user name.
     */
    username?: string;
}

/**
 * A collection of values returned by getInstanceSlowQueries.
 */
export interface GetInstanceSlowQueriesResult {
    /**
     * Database name.
     */
    readonly database?: string;
    readonly endTime?: string;
    readonly host?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly orderBy?: string;
    readonly orderByType?: string;
    readonly resultOutputFile?: string;
    /**
     * Slow query records.
     */
    readonly slowQueries: outputs.Cynosdb.GetInstanceSlowQueriesSlowQuery[];
    readonly startTime?: string;
    readonly username?: string;
}

export function getInstanceSlowQueriesOutput(args: GetInstanceSlowQueriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceSlowQueriesResult> {
    return pulumi.output(args).apply(a => getInstanceSlowQueries(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceSlowQueries.
 */
export interface GetInstanceSlowQueriesOutputArgs {
    /**
     * Database name.
     */
    database?: pulumi.Input<string>;
    /**
     * Latest transaction start time.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Client host.
     */
    host?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Sort field, optional values: QueryTime, LockTime, RowsExamined, RowsSent.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Sort type, optional values: asc, desc.
     */
    orderByType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Earliest transaction start time.
     */
    startTime?: pulumi.Input<string>;
    /**
     * user name.
     */
    username?: pulumi.Input<string>;
}
