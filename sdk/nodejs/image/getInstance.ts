// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an available image for the user.
 *
 * The Images data source fetch proper image, which could be one of the private images of the user and images of system resources provided by TencentCloud, as well as other public images and those available on the image market.
 *
 * > **NOTE:** This data source will be deprecated, please use `tencentcloud.Images.getInstance` instead.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const myFavorateImage = pulumi.output(tencentcloud.Image.getInstance({
 *     filters: [{
 *         name: "image-type",
 *         values: ["PUBLIC_IMAGE"],
 *     }],
 *     osName: "centos",
 * }));
 * ```
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Image/getInstance:getInstance", {
        "filters": args.filters,
        "imageNameRegex": args.imageNameRegex,
        "osName": args.osName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * One or more name/value pairs to filter.
     */
    filters?: inputs.Image.GetInstanceFilter[];
    /**
     * A regex string to apply to the image list returned by TencentCloud. **NOTE**: it is not wildcard, should look like `imageNameRegex = "^CentOS\s+6\.8\s+64\w*"`.
     */
    imageNameRegex?: string;
    /**
     * A string to apply with fuzzy match to the osName attribute on the image list returned by TencentCloud. **NOTE**: when osName is provided, highest priority is applied in this field instead of `imageNameRegex`.
     */
    osName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly filters?: outputs.Image.GetInstanceFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * An image id indicate the uniqueness of a certain image,  which can be used for instance creation or resetting.
     */
    readonly imageId: string;
    /**
     * Name of this image.
     */
    readonly imageName: string;
    readonly imageNameRegex?: string;
    readonly osName?: string;
    readonly resultOutputFile?: string;
}

export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * One or more name/value pairs to filter.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Image.GetInstanceFilterArgs>[]>;
    /**
     * A regex string to apply to the image list returned by TencentCloud. **NOTE**: it is not wildcard, should look like `imageNameRegex = "^CentOS\s+6\.8\s+64\w*"`.
     */
    imageNameRegex?: pulumi.Input<string>;
    /**
     * A string to apply with fuzzy match to the osName attribute on the image list returned by TencentCloud. **NOTE**: when osName is provided, highest priority is applied in this field instead of `imageNameRegex`.
     */
    osName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
