// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of VPN customer gateways.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = pulumi.output(tencentcloud.CustomerGateways({
 *     id: "cgw-xfqag",
 *     name: "main",
 *     publicIpAddress: "1.1.1.1",
 *     tags: [{
 *         test: "tf",
 *     }],
 * }));
 * ```
 */
export function getCustomerGateways(args?: GetCustomerGatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomerGatewaysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Vpn/getCustomerGateways:getCustomerGateways", {
        "id": args.id,
        "name": args.name,
        "publicIpAddress": args.publicIpAddress,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomerGateways.
 */
export interface GetCustomerGatewaysArgs {
    /**
     * ID of the VPN customer gateway.
     */
    id?: string;
    /**
     * Name of the customer gateway. The length of character is limited to 1-60.
     */
    name?: string;
    /**
     * Public ip address of the VPN customer gateway.
     */
    publicIpAddress?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Tags of the VPN customer gateway to be queried.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getCustomerGateways.
 */
export interface GetCustomerGatewaysResult {
    /**
     * Information list of the dedicated gateways.
     */
    readonly gatewayLists: outputs.Vpn.GetCustomerGatewaysGatewayList[];
    /**
     * ID of the VPN customer gateway.
     */
    readonly id?: string;
    /**
     * Name of the VPN customer gateway.
     */
    readonly name?: string;
    /**
     * Public ip address of the VPN customer gateway.
     */
    readonly publicIpAddress?: string;
    readonly resultOutputFile?: string;
    /**
     * Tags of the VPN customer gateway.
     */
    readonly tags?: {[key: string]: any};
}

export function getCustomerGatewaysOutput(args?: GetCustomerGatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomerGatewaysResult> {
    return pulumi.output(args).apply(a => getCustomerGateways(a, opts))
}

/**
 * A collection of arguments for invoking getCustomerGateways.
 */
export interface GetCustomerGatewaysOutputArgs {
    /**
     * ID of the VPN customer gateway.
     */
    id?: pulumi.Input<string>;
    /**
     * Name of the customer gateway. The length of character is limited to 1-60.
     */
    name?: pulumi.Input<string>;
    /**
     * Public ip address of the VPN customer gateway.
     */
    publicIpAddress?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Tags of the VPN customer gateway to be queried.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
