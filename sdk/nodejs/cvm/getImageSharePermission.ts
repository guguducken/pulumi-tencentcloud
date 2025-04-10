// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cvm imageSharePermission
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const imageSharePermission = pulumi.output(tencentcloud.Cvm.getImageSharePermission({
 *     imageId: "img-xxxxxx",
 * }));
 * ```
 */
export function getImageSharePermission(args: GetImageSharePermissionArgs, opts?: pulumi.InvokeOptions): Promise<GetImageSharePermissionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cvm/getImageSharePermission:getImageSharePermission", {
        "imageId": args.imageId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageSharePermission.
 */
export interface GetImageSharePermissionArgs {
    /**
     * The ID of the image to be shared.
     */
    imageId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getImageSharePermission.
 */
export interface GetImageSharePermissionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId: string;
    readonly resultOutputFile?: string;
    /**
     * Information on image sharing.
     */
    readonly sharePermissionSets: outputs.Cvm.GetImageSharePermissionSharePermissionSet[];
}

export function getImageSharePermissionOutput(args: GetImageSharePermissionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageSharePermissionResult> {
    return pulumi.output(args).apply(a => getImageSharePermission(a, opts))
}

/**
 * A collection of arguments for invoking getImageSharePermission.
 */
export interface GetImageSharePermissionOutputArgs {
    /**
     * The ID of the image to be shared.
     */
    imageId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
