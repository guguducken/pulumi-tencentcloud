// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create security group rule. This resource is similar with tencentcloud_security_group_lite_rule, rules can be ordered and configure descriptions.
 *
 * > **NOTE:** This resource must exclusive in one security group, do not declare additional rule resources of this security group elsewhere.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const sglab1Group = new tencentcloud.security.Group("sglab1Group", {description: "favourite sg_1"});
 * const sglab1GroupRuleSet = new tencentcloud.security.GroupRuleSet("sglab1GroupRuleSet", {
 *     securityGroupId: sglab1Group.id,
 *     ingresses: [
 *         {
 *             cidrBlock: "10.0.0.0/16",
 *             protocol: "TCP",
 *             port: "80",
 *             action: "ACCEPT",
 *             description: "favourite sg rule_1",
 *         },
 *         {
 *             protocol: "TCP",
 *             port: "80",
 *             action: "ACCEPT",
 *             sourceSecurityId: tencentcloud_security_group.sglab_3.id,
 *             description: "favourite sg rule_2",
 *         },
 *     ],
 *     egresses: [
 *         {
 *             action: "ACCEPT",
 *             addressTemplateId: "ipm-xxxxxxxx",
 *             description: "Allow address template",
 *         },
 *         {
 *             action: "ACCEPT",
 *             serviceTemplateGroup: "ppmg-xxxxxxxx",
 *             description: "Allow protocol template",
 *         },
 *         {
 *             cidrBlock: "10.0.0.0/16",
 *             protocol: "TCP",
 *             port: "80",
 *             action: "DROP",
 *             description: "favourite sg egress rule",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Resource tencentcloud_security_group_rule_set can be imported by passing security grou id
 *
 * ```sh
 *  $ pulumi import tencentcloud:Security/groupRuleSet:GroupRuleSet sglab_1 sg-xxxxxxxx
 * ```
 */
export class GroupRuleSet extends pulumi.CustomResource {
    /**
     * Get an existing GroupRuleSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupRuleSetState, opts?: pulumi.CustomResourceOptions): GroupRuleSet {
        return new GroupRuleSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Security/groupRuleSet:GroupRuleSet';

    /**
     * Returns true if the given object is an instance of GroupRuleSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupRuleSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupRuleSet.__pulumiType;
    }

    /**
     * List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
     */
    public readonly egresses!: pulumi.Output<outputs.Security.GroupRuleSetEgress[] | undefined>;
    /**
     * List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
     */
    public readonly ingresses!: pulumi.Output<outputs.Security.GroupRuleSetIngress[] | undefined>;
    /**
     * ID of the security group to be queried.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Security policies version, auto increment for every update.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a GroupRuleSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupRuleSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupRuleSetArgs | GroupRuleSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupRuleSetState | undefined;
            resourceInputs["egresses"] = state ? state.egresses : undefined;
            resourceInputs["ingresses"] = state ? state.ingresses : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as GroupRuleSetArgs | undefined;
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            resourceInputs["egresses"] = args ? args.egresses : undefined;
            resourceInputs["ingresses"] = args ? args.ingresses : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupRuleSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupRuleSet resources.
 */
export interface GroupRuleSetState {
    /**
     * List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
     */
    egresses?: pulumi.Input<pulumi.Input<inputs.Security.GroupRuleSetEgress>[]>;
    /**
     * List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
     */
    ingresses?: pulumi.Input<pulumi.Input<inputs.Security.GroupRuleSetIngress>[]>;
    /**
     * ID of the security group to be queried.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Security policies version, auto increment for every update.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupRuleSet resource.
 */
export interface GroupRuleSetArgs {
    /**
     * List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
     */
    egresses?: pulumi.Input<pulumi.Input<inputs.Security.GroupRuleSetEgress>[]>;
    /**
     * List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
     */
    ingresses?: pulumi.Input<pulumi.Input<inputs.Security.GroupRuleSetIngress>[]>;
    /**
     * ID of the security group to be queried.
     */
    securityGroupId: pulumi.Input<string>;
}
