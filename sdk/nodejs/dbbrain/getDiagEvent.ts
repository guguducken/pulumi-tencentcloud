// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dbbrain diagEvent
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const diagHistory = tencentcloud.Dbbrain.getDiagHistory({
 *     instanceId: `%s`,
 *     startTime: `%s`,
 *     endTime: `%s`,
 *     product: "mysql",
 * });
 * const diagEvent = diagHistory.then(diagHistory => tencentcloud.Dbbrain.getDiagEvent({
 *     instanceId: `%s`,
 *     eventId: diagHistory.events?[0]?.eventId,
 *     product: "mysql",
 * }));
 * ```
 */
export function getDiagEvent(args: GetDiagEventArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagEventResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getDiagEvent:getDiagEvent", {
        "eventId": args.eventId,
        "instanceId": args.instanceId,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDiagEvent.
 */
export interface GetDiagEventArgs {
    /**
     * Event ID. Obtain it through `Get Instance Diagnosis History DescribeDBDiagHistory`.
     */
    eventId?: number;
    /**
     * isntance id.
     */
    instanceId: string;
    /**
     * Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
     */
    product?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDiagEvent.
 */
export interface GetDiagEventResult {
    /**
     * diagnostic item.
     */
    readonly diagItem: string;
    /**
     * Diagnostic type.
     */
    readonly diagType: string;
    /**
     * End Time.
     */
    readonly endTime: string;
    readonly eventId: number;
    /**
     * Diagnostic event details, output is empty if there is no additional explanatory information.
     */
    readonly explanation: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * reserved text. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly metric: string;
    /**
     * Diagnostic summary.
     */
    readonly outline: string;
    /**
     * Diagnosed problem.
     */
    readonly problem: string;
    readonly product?: string;
    readonly resultOutputFile?: string;
    /**
     * severity. The severity is divided into 5 levels, according to the degree of impact from high to low: 1: Fatal, 2: Serious, 3: Warning, 4: Prompt, 5: Healthy.
     */
    readonly severity: number;
    /**
     * Starting time.
     */
    readonly startTime: string;
    /**
     * A diagnostic suggestion, or empty if there is no suggestion.
     */
    readonly suggestions: string;
}

export function getDiagEventOutput(args: GetDiagEventOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDiagEventResult> {
    return pulumi.output(args).apply(a => getDiagEvent(a, opts))
}

/**
 * A collection of arguments for invoking getDiagEvent.
 */
export interface GetDiagEventOutputArgs {
    /**
     * Event ID. Obtain it through `Get Instance Diagnosis History DescribeDBDiagHistory`.
     */
    eventId?: pulumi.Input<number>;
    /**
     * isntance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
     */
    product?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
