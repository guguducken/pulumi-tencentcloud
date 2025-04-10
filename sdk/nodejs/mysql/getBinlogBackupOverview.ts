// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mysql binlogBackupOverview
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const binlogBackupOverview = pulumi.output(tencentcloud.Mysql.getBinlogBackupOverview({
 *     product: "mysql",
 * }));
 * ```
 */
export function getBinlogBackupOverview(args: GetBinlogBackupOverviewArgs, opts?: pulumi.InvokeOptions): Promise<GetBinlogBackupOverviewResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mysql/getBinlogBackupOverview:getBinlogBackupOverview", {
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBinlogBackupOverview.
 */
export interface GetBinlogBackupOverviewArgs {
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
 * A collection of values returned by getBinlogBackupOverview.
 */
export interface GetBinlogBackupOverviewResult {
    /**
     * The number of archived log backups.
     */
    readonly binlogArchiveCount: number;
    /**
     * Archived log backup capacity (in bytes).
     */
    readonly binlogArchiveVolume: number;
    /**
     * The total number of log backups, including remote log backups.
     */
    readonly binlogBackupCount: number;
    /**
     * Total log backup capacity, including off-site log backup (unit is byte).
     */
    readonly binlogBackupVolume: number;
    /**
     * The number of standard storage log backups.
     */
    readonly binlogStandbyCount: number;
    /**
     * Standard storage log backup capacity (in bytes).
     */
    readonly binlogStandbyVolume: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly product: string;
    /**
     * The number of remote log backups.
     */
    readonly remoteBinlogCount: number;
    /**
     * Remote log backup capacity (in bytes).
     */
    readonly remoteBinlogVolume: number;
    readonly resultOutputFile?: string;
}

export function getBinlogBackupOverviewOutput(args: GetBinlogBackupOverviewOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBinlogBackupOverviewResult> {
    return pulumi.output(args).apply(a => getBinlogBackupOverview(a, opts))
}

/**
 * A collection of arguments for invoking getBinlogBackupOverview.
 */
export interface GetBinlogBackupOverviewOutputArgs {
    /**
     * The type of cloud database product to be queried, currently only supports `mysql`.
     */
    product: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
