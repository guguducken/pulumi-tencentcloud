// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf publicConfigSummary
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describePublicConfigSummary = pulumi.output(tencentcloud.TsfDescribePublicConfigSummary({
 *     configIdList: ["dcfg-p-ygbdw5mv"],
 *     // config_tag_list = [""]
 *     disableProgramAuthCheck: true,
 *     orderBy: "last_update_time",
 *     orderType: 0,
 *     searchWord: "test",
 * }));
 * ```
 */
export function getPublicConfigSummary(args?: GetPublicConfigSummaryArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicConfigSummaryResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tsf/getPublicConfigSummary:getPublicConfigSummary", {
        "configIdLists": args.configIdLists,
        "configTagLists": args.configTagLists,
        "disableProgramAuthCheck": args.disableProgramAuthCheck,
        "orderBy": args.orderBy,
        "orderType": args.orderType,
        "resultOutputFile": args.resultOutputFile,
        "searchWord": args.searchWord,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublicConfigSummary.
 */
export interface GetPublicConfigSummaryArgs {
    /**
     * Config Id List.
     */
    configIdLists?: string[];
    /**
     * config tag list.
     */
    configTagLists?: string[];
    /**
     * Whether to disable dataset authentication.
     */
    disableProgramAuthCheck?: boolean;
    /**
     * Sort by time: creation_time; Sort by name: config_name.
     */
    orderBy?: string;
    /**
     * Pass 0 for ascending order and 1 for descending order.
     */
    orderType?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Query keyword for fuzzy search: configuration item name. If not passed in, the full set will be queried.
     */
    searchWord?: string;
}

/**
 * A collection of values returned by getPublicConfigSummary.
 */
export interface GetPublicConfigSummaryResult {
    readonly configIdLists?: string[];
    readonly configTagLists?: string[];
    readonly disableProgramAuthCheck?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly orderBy?: string;
    readonly orderType?: number;
    readonly resultOutputFile?: string;
    /**
     * Public config Page Item.
     */
    readonly results: outputs.Tsf.GetPublicConfigSummaryResult[];
    readonly searchWord?: string;
}

export function getPublicConfigSummaryOutput(args?: GetPublicConfigSummaryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicConfigSummaryResult> {
    return pulumi.output(args).apply(a => getPublicConfigSummary(a, opts))
}

/**
 * A collection of arguments for invoking getPublicConfigSummary.
 */
export interface GetPublicConfigSummaryOutputArgs {
    /**
     * Config Id List.
     */
    configIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * config tag list.
     */
    configTagLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to disable dataset authentication.
     */
    disableProgramAuthCheck?: pulumi.Input<boolean>;
    /**
     * Sort by time: creation_time; Sort by name: config_name.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * Pass 0 for ascending order and 1 for descending order.
     */
    orderType?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Query keyword for fuzzy search: configuration item name. If not passed in, the full set will be queried.
     */
    searchWord?: pulumi.Input<string>;
}
