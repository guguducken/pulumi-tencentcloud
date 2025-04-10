// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tcr tagRetentionExecutionConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const myNs = new tencentcloud.tcr.Namespace("myNs", {
 *     instanceId: tencentcloud_tcr_instance.mytcr_retention.id,
 *     isPublic: true,
 *     isAutoScan: true,
 *     isPreventVul: true,
 *     severity: "medium",
 *     cveWhitelistItems: [{
 *         cveId: "cve-xxxxx",
 *     }],
 * });
 * const myRule = new tencentcloud.tcr.TagRetentionRule("myRule", {
 *     registryId: tencentcloud_tcr_instance.mytcr_retention.id,
 *     namespaceName: myNs.name,
 *     retentionRule: {
 *         key: "nDaysSinceLastPush",
 *         value: 2,
 *     },
 *     cronSetting: "manual",
 *     disabled: true,
 * });
 * const tagRetentionExecutionConfig = new tencentcloud.tcr.TagRetentionExecutionConfig("tagRetentionExecutionConfig", {
 *     registryId: myRule.registryId,
 *     retentionId: myRule.retentionId,
 *     dryRun: false,
 * });
 * ```
 */
export class TagRetentionExecutionConfig extends pulumi.CustomResource {
    /**
     * Get an existing TagRetentionExecutionConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagRetentionExecutionConfigState, opts?: pulumi.CustomResourceOptions): TagRetentionExecutionConfig {
        return new TagRetentionExecutionConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcr/tagRetentionExecutionConfig:TagRetentionExecutionConfig';

    /**
     * Returns true if the given object is an instance of TagRetentionExecutionConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagRetentionExecutionConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagRetentionExecutionConfig.__pulumiType;
    }

    /**
     * Whether to simulate execution, the default value is false, that is, non-simulation execution.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * execution id.
     */
    public /*out*/ readonly executionId!: pulumi.Output<number>;
    /**
     * instance id.
     */
    public readonly registryId!: pulumi.Output<string>;
    /**
     * retention id.
     */
    public readonly retentionId!: pulumi.Output<number>;

    /**
     * Create a TagRetentionExecutionConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagRetentionExecutionConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagRetentionExecutionConfigArgs | TagRetentionExecutionConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagRetentionExecutionConfigState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["executionId"] = state ? state.executionId : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["retentionId"] = state ? state.retentionId : undefined;
        } else {
            const args = argsOrState as TagRetentionExecutionConfigArgs | undefined;
            if ((!args || args.registryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryId'");
            }
            if ((!args || args.retentionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["registryId"] = args ? args.registryId : undefined;
            resourceInputs["retentionId"] = args ? args.retentionId : undefined;
            resourceInputs["executionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TagRetentionExecutionConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagRetentionExecutionConfig resources.
 */
export interface TagRetentionExecutionConfigState {
    /**
     * Whether to simulate execution, the default value is false, that is, non-simulation execution.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * execution id.
     */
    executionId?: pulumi.Input<number>;
    /**
     * instance id.
     */
    registryId?: pulumi.Input<string>;
    /**
     * retention id.
     */
    retentionId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a TagRetentionExecutionConfig resource.
 */
export interface TagRetentionExecutionConfigArgs {
    /**
     * Whether to simulate execution, the default value is false, that is, non-simulation execution.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * instance id.
     */
    registryId: pulumi.Input<string>;
    /**
     * retention id.
     */
    retentionId: pulumi.Input<number>;
}
