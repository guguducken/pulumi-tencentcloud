// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mysql backupOverview
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const backupOverview = pulumi.output(tencentcloud.Mysql.getBackupOverview({
 *     product: "mysql",
 * }));
 * ```
 */
export function getBackupOverview(args: GetBackupOverviewArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupOverviewResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mysql/getBackupOverview:getBackupOverview", {
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupOverview.
 */
export interface GetBackupOverviewArgs {
    /**
     * The type of cloud database product to be queried, currently only supports `mysql`.
     */
    product: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBackupOverview.
 */
export interface GetBackupOverviewResult {
    /**
     * Archive backup capacity, including data backup and log backup. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly backupArchiveVolume: number;
    /**
     * The total number of user backups in the current region (including data backups and log backups).
     */
    readonly backupCount: number;
    /**
     * Standard storage backup capacity, including data backup and log backup. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly backupStandbyVolume: number;
    /**
     * The total backup capacity of the user in the current region.
     */
    readonly backupVolume: number;
    /**
     * The billable capacity of the user&amp;#39;s backup in the current region, that is, the part that exceeds the gifted capacity.
     */
    readonly billingVolume: number;
    /**
     * The free backup capacity obtained by the user in the current region.
     */
    readonly freeVolume: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly product: string;
    /**
     * The total capacity of off-site backup of the user in the current region. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly remoteBackupVolume: number;
    readonly resultOutputFile?: string;
}

export function getBackupOverviewOutput(args: GetBackupOverviewOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupOverviewResult> {
    return pulumi.output(args).apply(a => getBackupOverview(a, opts))
}

/**
 * A collection of arguments for invoking getBackupOverview.
 */
export interface GetBackupOverviewOutputArgs {
    /**
     * The type of cloud database product to be queried, currently only supports `mysql`.
     */
    product: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
