// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf unitRules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const unitRules = pulumi.output(tencentcloud.Tsf.getUnitRules({
 *     gatewayInstanceId: "gw-ins-lvdypq5k",
 *     status: "disabled",
 * }));
 * ```
 */
export function getUnitRules(args: GetUnitRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetUnitRulesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tsf/getUnitRules:getUnitRules", {
        "gatewayInstanceId": args.gatewayInstanceId,
        "resultOutputFile": args.resultOutputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getUnitRules.
 */
export interface GetUnitRulesArgs {
    /**
     * gateway instance id.
     */
    gatewayInstanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Enabled state, disabled: unpublished, enabled: published.
     */
    status?: string;
}

/**
 * A collection of values returned by getUnitRules.
 */
export interface GetUnitRulesResult {
    /**
     * Gateway Entity ID.
     */
    readonly gatewayInstanceId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Pagination list information.
     */
    readonly results: outputs.Tsf.GetUnitRulesResult[];
    /**
     * Use status: enabled/disabled.
     */
    readonly status?: string;
}

export function getUnitRulesOutput(args: GetUnitRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUnitRulesResult> {
    return pulumi.output(args).apply(a => getUnitRules(a, opts))
}

/**
 * A collection of arguments for invoking getUnitRules.
 */
export interface GetUnitRulesOutputArgs {
    /**
     * gateway instance id.
     */
    gatewayInstanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Enabled state, disabled: unpublished, enabled: published.
     */
    status?: pulumi.Input<string>;
}
