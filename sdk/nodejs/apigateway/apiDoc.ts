// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a APIGateway ApiDoc
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const myApiDoc = new tencentcloud.ApiGateway.ApiDoc("my_api_doc", {
 *     apiDocName: "doc_test1",
 *     apiIds: [
 *         "api-test1",
 *         "api-test2",
 *     ],
 *     environment: "release",
 *     serviceId: "service_test1",
 * });
 * ```
 */
export class ApiDoc extends pulumi.CustomResource {
    /**
     * Get an existing ApiDoc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiDocState, opts?: pulumi.CustomResourceOptions): ApiDoc {
        return new ApiDoc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:ApiGateway/apiDoc:ApiDoc';

    /**
     * Returns true if the given object is an instance of ApiDoc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiDoc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiDoc.__pulumiType;
    }

    /**
     * Api Document count.
     */
    public /*out*/ readonly apiCount!: pulumi.Output<number>;
    /**
     * Api Document ID.
     */
    public /*out*/ readonly apiDocId!: pulumi.Output<string>;
    /**
     * Api Document name.
     */
    public readonly apiDocName!: pulumi.Output<string>;
    /**
     * API Document Build Status.
     */
    public /*out*/ readonly apiDocStatus!: pulumi.Output<string>;
    /**
     * API Document Access URI.
     */
    public /*out*/ readonly apiDocUri!: pulumi.Output<string>;
    /**
     * List of APIs for generating documents.
     */
    public readonly apiIds!: pulumi.Output<string[]>;
    /**
     * List of names for generating documents.
     */
    public /*out*/ readonly apiNames!: pulumi.Output<string[]>;
    /**
     * Env name.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Number of API document releases.
     */
    public /*out*/ readonly releaseCount!: pulumi.Output<number>;
    /**
     * Service name.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * API Document service name.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * API Document Sharing Password.
     */
    public /*out*/ readonly sharePassword!: pulumi.Output<string>;
    /**
     * API Document update time.
     */
    public /*out*/ readonly updatedTime!: pulumi.Output<string>;
    /**
     * API Document Viewing Times.
     */
    public /*out*/ readonly viewCount!: pulumi.Output<number>;

    /**
     * Create a ApiDoc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiDocArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiDocArgs | ApiDocState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiDocState | undefined;
            resourceInputs["apiCount"] = state ? state.apiCount : undefined;
            resourceInputs["apiDocId"] = state ? state.apiDocId : undefined;
            resourceInputs["apiDocName"] = state ? state.apiDocName : undefined;
            resourceInputs["apiDocStatus"] = state ? state.apiDocStatus : undefined;
            resourceInputs["apiDocUri"] = state ? state.apiDocUri : undefined;
            resourceInputs["apiIds"] = state ? state.apiIds : undefined;
            resourceInputs["apiNames"] = state ? state.apiNames : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["releaseCount"] = state ? state.releaseCount : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["sharePassword"] = state ? state.sharePassword : undefined;
            resourceInputs["updatedTime"] = state ? state.updatedTime : undefined;
            resourceInputs["viewCount"] = state ? state.viewCount : undefined;
        } else {
            const args = argsOrState as ApiDocArgs | undefined;
            if ((!args || args.apiDocName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiDocName'");
            }
            if ((!args || args.apiIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiIds'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["apiDocName"] = args ? args.apiDocName : undefined;
            resourceInputs["apiIds"] = args ? args.apiIds : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["apiCount"] = undefined /*out*/;
            resourceInputs["apiDocId"] = undefined /*out*/;
            resourceInputs["apiDocStatus"] = undefined /*out*/;
            resourceInputs["apiDocUri"] = undefined /*out*/;
            resourceInputs["apiNames"] = undefined /*out*/;
            resourceInputs["releaseCount"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["sharePassword"] = undefined /*out*/;
            resourceInputs["updatedTime"] = undefined /*out*/;
            resourceInputs["viewCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiDoc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiDoc resources.
 */
export interface ApiDocState {
    /**
     * Api Document count.
     */
    apiCount?: pulumi.Input<number>;
    /**
     * Api Document ID.
     */
    apiDocId?: pulumi.Input<string>;
    /**
     * Api Document name.
     */
    apiDocName?: pulumi.Input<string>;
    /**
     * API Document Build Status.
     */
    apiDocStatus?: pulumi.Input<string>;
    /**
     * API Document Access URI.
     */
    apiDocUri?: pulumi.Input<string>;
    /**
     * List of APIs for generating documents.
     */
    apiIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of names for generating documents.
     */
    apiNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Env name.
     */
    environment?: pulumi.Input<string>;
    /**
     * Number of API document releases.
     */
    releaseCount?: pulumi.Input<number>;
    /**
     * Service name.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * API Document service name.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * API Document Sharing Password.
     */
    sharePassword?: pulumi.Input<string>;
    /**
     * API Document update time.
     */
    updatedTime?: pulumi.Input<string>;
    /**
     * API Document Viewing Times.
     */
    viewCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ApiDoc resource.
 */
export interface ApiDocArgs {
    /**
     * Api Document name.
     */
    apiDocName: pulumi.Input<string>;
    /**
     * List of APIs for generating documents.
     */
    apiIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Env name.
     */
    environment: pulumi.Input<string>;
    /**
     * Service name.
     */
    serviceId: pulumi.Input<string>;
}
