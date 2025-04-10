// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create custom domain of API gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.ApiGateway.CustomDomain("foo", {
 *     defaultDomain: "service-ohxqslqe-1259649581.gz.apigw.tencentcs.com",
 *     isDefaultMapping: false,
 *     netType: "OUTER",
 *     pathMappings: [
 *         "/good#test",
 *         "/root#release",
 *     ],
 *     protocol: "http",
 *     serviceId: "service-ohxqslqe",
 *     subDomain: "tic-test.dnsv1.com",
 * });
 * ```
 */
export class CustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing CustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomDomainState, opts?: pulumi.CustomResourceOptions): CustomDomain {
        return new CustomDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:ApiGateway/customDomain:CustomDomain';

    /**
     * Returns true if the given object is an instance of CustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDomain.__pulumiType;
    }

    /**
     * Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
     */
    public readonly certificateId!: pulumi.Output<string>;
    /**
     * Default domain name.
     */
    public readonly defaultDomain!: pulumi.Output<string>;
    /**
     * Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
     */
    public readonly isDefaultMapping!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
     */
    public readonly isForcedHttps!: pulumi.Output<boolean | undefined>;
    /**
     * Network type. Valid values: `OUTER`, `INNER`.
     */
    public readonly netType!: pulumi.Output<string>;
    /**
     * Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
     */
    public readonly pathMappings!: pulumi.Output<string[]>;
    /**
     * Protocol supported by service. Valid values: `http`, `https`, `http&https`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Unique service ID.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Custom domain name to be bound.
     */
    public readonly subDomain!: pulumi.Output<string>;

    /**
     * Create a CustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomDomainArgs | CustomDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomDomainState | undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["defaultDomain"] = state ? state.defaultDomain : undefined;
            resourceInputs["isDefaultMapping"] = state ? state.isDefaultMapping : undefined;
            resourceInputs["isForcedHttps"] = state ? state.isForcedHttps : undefined;
            resourceInputs["netType"] = state ? state.netType : undefined;
            resourceInputs["pathMappings"] = state ? state.pathMappings : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subDomain"] = state ? state.subDomain : undefined;
        } else {
            const args = argsOrState as CustomDomainArgs | undefined;
            if ((!args || args.defaultDomain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultDomain'");
            }
            if ((!args || args.netType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'netType'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            if ((!args || args.subDomain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subDomain'");
            }
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["defaultDomain"] = args ? args.defaultDomain : undefined;
            resourceInputs["isDefaultMapping"] = args ? args.isDefaultMapping : undefined;
            resourceInputs["isForcedHttps"] = args ? args.isForcedHttps : undefined;
            resourceInputs["netType"] = args ? args.netType : undefined;
            resourceInputs["pathMappings"] = args ? args.pathMappings : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["subDomain"] = args ? args.subDomain : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomDomain resources.
 */
export interface CustomDomainState {
    /**
     * Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Default domain name.
     */
    defaultDomain?: pulumi.Input<string>;
    /**
     * Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
     */
    isDefaultMapping?: pulumi.Input<boolean>;
    /**
     * Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
     */
    isForcedHttps?: pulumi.Input<boolean>;
    /**
     * Network type. Valid values: `OUTER`, `INNER`.
     */
    netType?: pulumi.Input<string>;
    /**
     * Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
     */
    pathMappings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Protocol supported by service. Valid values: `http`, `https`, `http&https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Unique service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
     */
    status?: pulumi.Input<number>;
    /**
     * Custom domain name to be bound.
     */
    subDomain?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomDomain resource.
 */
export interface CustomDomainArgs {
    /**
     * Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Default domain name.
     */
    defaultDomain: pulumi.Input<string>;
    /**
     * Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `pathMappings` attribute is required.
     */
    isDefaultMapping?: pulumi.Input<boolean>;
    /**
     * Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
     */
    isForcedHttps?: pulumi.Input<boolean>;
    /**
     * Network type. Valid values: `OUTER`, `INNER`.
     */
    netType: pulumi.Input<string>;
    /**
     * Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
     */
    pathMappings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Protocol supported by service. Valid values: `http`, `https`, `http&https`.
     */
    protocol: pulumi.Input<string>;
    /**
     * Unique service ID.
     */
    serviceId: pulumi.Input<string>;
    /**
     * Custom domain name to be bound.
     */
    subDomain: pulumi.Input<string>;
}
