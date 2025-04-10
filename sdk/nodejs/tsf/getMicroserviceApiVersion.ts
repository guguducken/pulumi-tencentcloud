// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf microserviceApiVersion
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const microserviceApiVersion = pulumi.output(tencentcloud.Tsf.getMicroserviceApiVersion({
 *     method: "get",
 *     microserviceId: "ms-yq3jo6jd",
 *     path: "",
 * }));
 * ```
 */
export function getMicroserviceApiVersion(args: GetMicroserviceApiVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetMicroserviceApiVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tsf/getMicroserviceApiVersion:getMicroserviceApiVersion", {
        "method": args.method,
        "microserviceId": args.microserviceId,
        "path": args.path,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMicroserviceApiVersion.
 */
export interface GetMicroserviceApiVersionArgs {
    /**
     * request method.
     */
    method?: string;
    /**
     * Microservice ID.
     */
    microserviceId: string;
    /**
     * api path.
     */
    path?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getMicroserviceApiVersion.
 */
export interface GetMicroserviceApiVersionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly method?: string;
    readonly microserviceId: string;
    readonly path?: string;
    readonly resultOutputFile?: string;
    /**
     * api version list.
     */
    readonly results: outputs.Tsf.GetMicroserviceApiVersionResult[];
}

export function getMicroserviceApiVersionOutput(args: GetMicroserviceApiVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMicroserviceApiVersionResult> {
    return pulumi.output(args).apply(a => getMicroserviceApiVersion(a, opts))
}

/**
 * A collection of arguments for invoking getMicroserviceApiVersion.
 */
export interface GetMicroserviceApiVersionOutputArgs {
    /**
     * request method.
     */
    method?: pulumi.Input<string>;
    /**
     * Microservice ID.
     */
    microserviceId: pulumi.Input<string>;
    /**
     * api path.
     */
    path?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
