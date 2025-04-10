// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc networkInterfaceLimit
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const networkInterfaceLimit = pulumi.output(tencentcloud.Vpc.getNetworkInterfaceLimit({
 *     instanceId: "ins-cr2rfq78",
 * }));
 * ```
 */
export function getNetworkInterfaceLimit(args: GetNetworkInterfaceLimitArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInterfaceLimitResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Vpc/getNetworkInterfaceLimit:getNetworkInterfaceLimit", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkInterfaceLimit.
 */
export interface GetNetworkInterfaceLimitArgs {
    /**
     * ID of a CVM instance or ENI to query.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getNetworkInterfaceLimit.
 */
export interface GetNetworkInterfaceLimitResult {
    /**
     * Quota of IP addresses that can be allocated to each standard-mounted ENI.
     */
    readonly eniPrivateIpAddressQuantity: number;
    /**
     * Quota of ENIs mounted to a CVM instance in a standard way.
     */
    readonly eniQuantity: number;
    /**
     * Quota of IP addresses that can be allocated to each extension-mounted ENI.Note: this field may return `null`, indicating that no valid values can be obtained.
     */
    readonly extendEniPrivateIpAddressQuantity: number;
    /**
     * Quota of ENIs mounted to a CVM instance as an extensionNote: this field may return `null`, indicating that no valid values can be obtained.
     */
    readonly extendEniQuantity: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly resultOutputFile?: string;
    /**
     * The quota of IPs that can be assigned to each relayed ENI.Note: This field may return `null`, indicating that no valid values can be obtained.
     */
    readonly subEniPrivateIpAddressQuantity: number;
    /**
     * The quota of relayed ENIsNote: This field may return `null`, indicating that no valid values can be obtained.
     */
    readonly subEniQuantity: number;
}

export function getNetworkInterfaceLimitOutput(args: GetNetworkInterfaceLimitOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkInterfaceLimitResult> {
    return pulumi.output(args).apply(a => getNetworkInterfaceLimit(a, opts))
}

/**
 * A collection of arguments for invoking getNetworkInterfaceLimit.
 */
export interface GetNetworkInterfaceLimitOutputArgs {
    /**
     * ID of a CVM instance or ENI to query.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
