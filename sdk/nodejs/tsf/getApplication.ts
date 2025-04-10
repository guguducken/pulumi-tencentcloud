// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tsf application
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const application = pulumi.output(tencentcloud.Tsf.getApplication({
 *     // application_resource_type_list = [""]
 *     applicationIdLists: ["application-a24x29xv"],
 *     applicationType: "V",
 *     microserviceType: "N",
 * }));
 * ```
 */
export function getApplication(args?: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tsf/getApplication:getApplication", {
        "applicationIdLists": args.applicationIdLists,
        "applicationResourceTypeLists": args.applicationResourceTypeLists,
        "applicationType": args.applicationType,
        "microserviceType": args.microserviceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationArgs {
    /**
     * Id list.
     */
    applicationIdLists?: string[];
    /**
     * An array of application resource types.
     */
    applicationResourceTypeLists?: string[];
    /**
     * The application type. V OR C, V means VM, C means container.
     */
    applicationType?: string;
    /**
     * The microservice type of the application.
     */
    microserviceType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getApplication.
 */
export interface GetApplicationResult {
    readonly applicationIdLists?: string[];
    readonly applicationResourceTypeLists?: string[];
    /**
     * The type of the application.
     */
    readonly applicationType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The microservice type of the application.
     */
    readonly microserviceType?: string;
    readonly resultOutputFile?: string;
    /**
     * The application paging list information.
     */
    readonly results: outputs.Tsf.GetApplicationResult[];
}

export function getApplicationOutput(args?: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply(a => getApplication(a, opts))
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationOutputArgs {
    /**
     * Id list.
     */
    applicationIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of application resource types.
     */
    applicationResourceTypeLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The application type. V OR C, V means VM, C means container.
     */
    applicationType?: pulumi.Input<string>;
    /**
     * The microservice type of the application.
     */
    microserviceType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
