// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tsf configTemplate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const configTemplate = new tencentcloud.Tsf.ConfigTemplate("config_template", {
 *     configTemplateDesc: "terraform-test",
 *     configTemplateName: "terraform-template-name",
 *     configTemplateType: "Ribbon",
 *     configTemplateValue: `  ribbon.ReadTimeout: 5000
 *   ribbon.ConnectTimeout: 2000
 *   ribbon.MaxAutoRetries: 0
 *   ribbon.MaxAutoRetriesNextServer: 1
 *   ribbon.OkToRetryOnAllOperations: true
 * `,
 * });
 * ```
 */
export class ConfigTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ConfigTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigTemplateState, opts?: pulumi.CustomResourceOptions): ConfigTemplate {
        return new ConfigTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tsf/configTemplate:ConfigTemplate';

    /**
     * Returns true if the given object is an instance of ConfigTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigTemplate.__pulumiType;
    }

    /**
     * Configuration template description.
     */
    public readonly configTemplateDesc!: pulumi.Output<string | undefined>;
    /**
     * Template Id.
     */
    public /*out*/ readonly configTemplateId!: pulumi.Output<string>;
    /**
     * Configuration template name.
     */
    public readonly configTemplateName!: pulumi.Output<string>;
    /**
     * Configure the microservice framework corresponding to the template.
     */
    public readonly configTemplateType!: pulumi.Output<string>;
    /**
     * Configure template data.
     */
    public readonly configTemplateValue!: pulumi.Output<string>;
    /**
     * creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Program id list.
     */
    public readonly programIdLists!: pulumi.Output<string[] | undefined>;
    /**
     * update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ConfigTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigTemplateArgs | ConfigTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigTemplateState | undefined;
            resourceInputs["configTemplateDesc"] = state ? state.configTemplateDesc : undefined;
            resourceInputs["configTemplateId"] = state ? state.configTemplateId : undefined;
            resourceInputs["configTemplateName"] = state ? state.configTemplateName : undefined;
            resourceInputs["configTemplateType"] = state ? state.configTemplateType : undefined;
            resourceInputs["configTemplateValue"] = state ? state.configTemplateValue : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["programIdLists"] = state ? state.programIdLists : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ConfigTemplateArgs | undefined;
            if ((!args || args.configTemplateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configTemplateName'");
            }
            if ((!args || args.configTemplateType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configTemplateType'");
            }
            if ((!args || args.configTemplateValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configTemplateValue'");
            }
            resourceInputs["configTemplateDesc"] = args ? args.configTemplateDesc : undefined;
            resourceInputs["configTemplateName"] = args ? args.configTemplateName : undefined;
            resourceInputs["configTemplateType"] = args ? args.configTemplateType : undefined;
            resourceInputs["configTemplateValue"] = args ? args.configTemplateValue : undefined;
            resourceInputs["programIdLists"] = args ? args.programIdLists : undefined;
            resourceInputs["configTemplateId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigTemplate resources.
 */
export interface ConfigTemplateState {
    /**
     * Configuration template description.
     */
    configTemplateDesc?: pulumi.Input<string>;
    /**
     * Template Id.
     */
    configTemplateId?: pulumi.Input<string>;
    /**
     * Configuration template name.
     */
    configTemplateName?: pulumi.Input<string>;
    /**
     * Configure the microservice framework corresponding to the template.
     */
    configTemplateType?: pulumi.Input<string>;
    /**
     * Configure template data.
     */
    configTemplateValue?: pulumi.Input<string>;
    /**
     * creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Program id list.
     */
    programIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * update time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigTemplate resource.
 */
export interface ConfigTemplateArgs {
    /**
     * Configuration template description.
     */
    configTemplateDesc?: pulumi.Input<string>;
    /**
     * Configuration template name.
     */
    configTemplateName: pulumi.Input<string>;
    /**
     * Configure the microservice framework corresponding to the template.
     */
    configTemplateType: pulumi.Input<string>;
    /**
     * Configure template data.
     */
    configTemplateValue: pulumi.Input<string>;
    /**
     * Program id list.
     */
    programIdLists?: pulumi.Input<pulumi.Input<string>[]>;
}
