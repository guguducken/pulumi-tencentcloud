// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql param template
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const paramTemplate = new tencentcloud.Mysql.ParamTemplate("param_template", {
 *     description: "terraform-test",
 *     engineType: "InnoDB",
 *     engineVersion: "8.0",
 *     paramLists: [
 *         {
 *             currentValue: "1",
 *             name: "auto_increment_increment",
 *         },
 *         {
 *             currentValue: "1",
 *             name: "auto_increment_offset",
 *         },
 *         {
 *             currentValue: "ON",
 *             name: "automatic_sp_privileges",
 *         },
 *     ],
 *     templateType: "HIGH_STABILITY",
 * });
 * ```
 *
 * ## Import
 *
 * mysql param template can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mysql/paramTemplate:ParamTemplate param_template template_id
 * ```
 */
export class ParamTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ParamTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParamTemplateState, opts?: pulumi.CustomResourceOptions): ParamTemplate {
        return new ParamTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/paramTemplate:ParamTemplate';

    /**
     * Returns true if the given object is an instance of ParamTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ParamTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ParamTemplate.__pulumiType;
    }

    /**
     * The description of parameter template.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
     */
    public readonly engineType!: pulumi.Output<string | undefined>;
    /**
     * The version of MySQL.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of parameter template.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * parameter list.
     */
    public readonly paramLists!: pulumi.Output<outputs.Mysql.ParamTemplateParamList[]>;
    /**
     * The ID of source parameter template.
     */
    public readonly templateId!: pulumi.Output<number>;
    /**
     * The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
     */
    public readonly templateType!: pulumi.Output<string | undefined>;

    /**
     * Create a ParamTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ParamTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParamTemplateArgs | ParamTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ParamTemplateState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["engineType"] = state ? state.engineType : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["paramLists"] = state ? state.paramLists : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
            resourceInputs["templateType"] = state ? state.templateType : undefined;
        } else {
            const args = argsOrState as ParamTemplateArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineType"] = args ? args.engineType : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["paramLists"] = args ? args.paramLists : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["templateType"] = args ? args.templateType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ParamTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ParamTemplate resources.
 */
export interface ParamTemplateState {
    /**
     * The description of parameter template.
     */
    description?: pulumi.Input<string>;
    /**
     * The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
     */
    engineType?: pulumi.Input<string>;
    /**
     * The version of MySQL.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The name of parameter template.
     */
    name?: pulumi.Input<string>;
    /**
     * parameter list.
     */
    paramLists?: pulumi.Input<pulumi.Input<inputs.Mysql.ParamTemplateParamList>[]>;
    /**
     * The ID of source parameter template.
     */
    templateId?: pulumi.Input<number>;
    /**
     * The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
     */
    templateType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ParamTemplate resource.
 */
export interface ParamTemplateArgs {
    /**
     * The description of parameter template.
     */
    description?: pulumi.Input<string>;
    /**
     * The engine type of instance, optional value is InnoDB or RocksDB, default to InnoDB.
     */
    engineType?: pulumi.Input<string>;
    /**
     * The version of MySQL.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The name of parameter template.
     */
    name?: pulumi.Input<string>;
    /**
     * parameter list.
     */
    paramLists?: pulumi.Input<pulumi.Input<inputs.Mysql.ParamTemplateParamList>[]>;
    /**
     * The ID of source parameter template.
     */
    templateId?: pulumi.Input<number>;
    /**
     * The default type of parameter template, supported value is HIGH_STABILITY or HIGH_PERFORMANCE.
     */
    templateType?: pulumi.Input<string>;
}
