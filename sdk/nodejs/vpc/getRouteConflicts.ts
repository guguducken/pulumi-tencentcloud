// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc routeConflicts
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const routeConflicts = pulumi.output(tencentcloud.Vpc.getRouteConflicts({
 *     destinationCidrBlocks: ["172.18.111.0/24"],
 *     routeTableId: "rtb-6xypllqe",
 * }));
 * ```
 */
export function getRouteConflicts(args: GetRouteConflictsArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteConflictsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Vpc/getRouteConflicts:getRouteConflicts", {
        "destinationCidrBlocks": args.destinationCidrBlocks,
        "resultOutputFile": args.resultOutputFile,
        "routeTableId": args.routeTableId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteConflicts.
 */
export interface GetRouteConflictsArgs {
    /**
     * List of conflicting destinations to check for.
     */
    destinationCidrBlocks: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Routing table instance ID, for example:rtb-azd4dt1c.
     */
    routeTableId: string;
}

/**
 * A collection of values returned by getRouteConflicts.
 */
export interface GetRouteConflictsResult {
    readonly destinationCidrBlocks: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * route conflict list.
     */
    readonly routeConflictSets: outputs.Vpc.GetRouteConflictsRouteConflictSet[];
    /**
     * route table id.
     */
    readonly routeTableId: string;
}

export function getRouteConflictsOutput(args: GetRouteConflictsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteConflictsResult> {
    return pulumi.output(args).apply(a => getRouteConflicts(a, opts))
}

/**
 * A collection of arguments for invoking getRouteConflicts.
 */
export interface GetRouteConflictsOutputArgs {
    /**
     * List of conflicting destinations to check for.
     */
    destinationCidrBlocks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Routing table instance ID, for example:rtb-azd4dt1c.
     */
    routeTableId: pulumi.Input<string>;
}
