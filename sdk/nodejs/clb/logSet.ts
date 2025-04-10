// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an exclusive CLB Logset.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Clb.LogSet("foo", {
 *     period: 7,
 * });
 * ```
 *
 * ## Import
 *
 * CLB log set can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Clb/logSet:LogSet foo 4eb9e3a8-9c42-4b32-9ddf-e215e9c92764
 * ```
 */
export class LogSet extends pulumi.CustomResource {
    /**
     * Get an existing LogSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogSetState, opts?: pulumi.CustomResourceOptions): LogSet {
        return new LogSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Clb/logSet:LogSet';

    /**
     * Returns true if the given object is an instance of LogSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogSet.__pulumiType;
    }

    /**
     * Logset creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Logset name, which unique and fixed `clbLogset` among all CLS logsets.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Logset retention period in days. Maximun value is `90`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Number of log topics in logset.
     */
    public /*out*/ readonly topicCount!: pulumi.Output<string>;

    /**
     * Create a LogSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogSetArgs | LogSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogSetState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["topicCount"] = state ? state.topicCount : undefined;
        } else {
            const args = argsOrState as LogSetArgs | undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["topicCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogSet resources.
 */
export interface LogSetState {
    /**
     * Logset creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Logset name, which unique and fixed `clbLogset` among all CLS logsets.
     */
    name?: pulumi.Input<string>;
    /**
     * Logset retention period in days. Maximun value is `90`.
     */
    period?: pulumi.Input<number>;
    /**
     * Number of log topics in logset.
     */
    topicCount?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogSet resource.
 */
export interface LogSetArgs {
    /**
     * Logset retention period in days. Maximun value is `90`.
     */
    period?: pulumi.Input<number>;
}
