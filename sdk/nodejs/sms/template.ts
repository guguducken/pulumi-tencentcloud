// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sms template
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const template = new tencentcloud.Sms.Template("template", {
 *     international: 0,
 *     remark: "terraform test",
 *     smsType: 0,
 *     templateContent: "Template Content",
 *     templateName: "Template By Terraform",
 * });
 * ```
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateState, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sms/template:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
     */
    public readonly international!: pulumi.Output<number>;
    /**
     * Template remarks, such as reason for application and use case.
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * SMS type. 0: regular SMS, 1: marketing SMS.
     */
    public readonly smsType!: pulumi.Output<number>;
    /**
     * Message Template Content.
     */
    public readonly templateContent!: pulumi.Output<string>;
    /**
     * Message Template name, which must be unique.
     */
    public readonly templateName!: pulumi.Output<string>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateArgs | TemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemplateState | undefined;
            resourceInputs["international"] = state ? state.international : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["smsType"] = state ? state.smsType : undefined;
            resourceInputs["templateContent"] = state ? state.templateContent : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.international === undefined) && !opts.urn) {
                throw new Error("Missing required property 'international'");
            }
            if ((!args || args.remark === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remark'");
            }
            if ((!args || args.smsType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'smsType'");
            }
            if ((!args || args.templateContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateContent'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["international"] = args ? args.international : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["smsType"] = args ? args.smsType : undefined;
            resourceInputs["templateContent"] = args ? args.templateContent : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
     */
    international?: pulumi.Input<number>;
    /**
     * Template remarks, such as reason for application and use case.
     */
    remark?: pulumi.Input<string>;
    /**
     * SMS type. 0: regular SMS, 1: marketing SMS.
     */
    smsType?: pulumi.Input<number>;
    /**
     * Message Template Content.
     */
    templateContent?: pulumi.Input<string>;
    /**
     * Message Template name, which must be unique.
     */
    templateName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
     */
    international: pulumi.Input<number>;
    /**
     * Template remarks, such as reason for application and use case.
     */
    remark: pulumi.Input<string>;
    /**
     * SMS type. 0: regular SMS, 1: marketing SMS.
     */
    smsType: pulumi.Input<number>;
    /**
     * Message Template Content.
     */
    templateContent: pulumi.Input<string>;
    /**
     * Message Template name, which must be unique.
     */
    templateName: pulumi.Input<string>;
}
