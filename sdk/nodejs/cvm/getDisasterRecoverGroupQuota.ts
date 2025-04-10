// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cvm disasterRecoverGroupQuota
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const disasterRecoverGroupQuota = pulumi.output(tencentcloud.Cvm.getDisasterRecoverGroupQuota());
 * ```
 */
export function getDisasterRecoverGroupQuota(args?: GetDisasterRecoverGroupQuotaArgs, opts?: pulumi.InvokeOptions): Promise<GetDisasterRecoverGroupQuotaResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cvm/getDisasterRecoverGroupQuota:getDisasterRecoverGroupQuota", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDisasterRecoverGroupQuota.
 */
export interface GetDisasterRecoverGroupQuotaArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDisasterRecoverGroupQuota.
 */
export interface GetDisasterRecoverGroupQuotaResult {
    /**
     * The number of placement groups that have been created by the current user.
     */
    readonly currentNum: number;
    /**
     * Quota on instances in a physical-machine-type disaster recovery group.
     */
    readonly cvmInHostGroupQuota: number;
    /**
     * Quota on instances in a rack-type disaster recovery group.
     */
    readonly cvmInRackGroupQuota: number;
    /**
     * Quota on instances in a switch-type disaster recovery group.
     */
    readonly cvmInSwGroupQuota: number;
    /**
     * The maximum number of placement groups that can be created.
     */
    readonly groupQuota: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function getDisasterRecoverGroupQuotaOutput(args?: GetDisasterRecoverGroupQuotaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDisasterRecoverGroupQuotaResult> {
    return pulumi.output(args).apply(a => getDisasterRecoverGroupQuota(a, opts))
}

/**
 * A collection of arguments for invoking getDisasterRecoverGroupQuota.
 */
export interface GetDisasterRecoverGroupQuotaOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
