// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a monitor tmpTkeBasicConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const tmpTkeBasicConfig = new tencentcloud.Monitor.TmpTkeBasicConfig("tmp_tke_basic_config", {
 *     clusterId: "cls-xxxxxx",
 *     clusterType: "eks",
 *     instanceId: "prom-xxxxxx",
 *     metricsNames: ["kube_job_status_succeeded"],
 * });
 * ```
 */
export class TmpTkeBasicConfig extends pulumi.CustomResource {
    /**
     * Get an existing TmpTkeBasicConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TmpTkeBasicConfigState, opts?: pulumi.CustomResourceOptions): TmpTkeBasicConfig {
        return new TmpTkeBasicConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/tmpTkeBasicConfig:TmpTkeBasicConfig';

    /**
     * Returns true if the given object is an instance of TmpTkeBasicConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TmpTkeBasicConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TmpTkeBasicConfig.__pulumiType;
    }

    /**
     * ID of cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Type of cluster.
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * Full configuration in yaml format.
     */
    public /*out*/ readonly config!: pulumi.Output<string>;
    /**
     * config type, `serviceMonitors`, `podMonitors`, `rawJobs`.
     */
    public /*out*/ readonly configType!: pulumi.Output<string>;
    /**
     * ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Configure the name of the metric to keep on.
     */
    public readonly metricsNames!: pulumi.Output<string[]>;
    /**
     * Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a TmpTkeBasicConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TmpTkeBasicConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TmpTkeBasicConfigArgs | TmpTkeBasicConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TmpTkeBasicConfigState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["configType"] = state ? state.configType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["metricsNames"] = state ? state.metricsNames : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as TmpTkeBasicConfigArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.metricsNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricsNames'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["metricsNames"] = args ? args.metricsNames : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["configType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TmpTkeBasicConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TmpTkeBasicConfig resources.
 */
export interface TmpTkeBasicConfigState {
    /**
     * ID of cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Type of cluster.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * Full configuration in yaml format.
     */
    config?: pulumi.Input<string>;
    /**
     * config type, `serviceMonitors`, `podMonitors`, `rawJobs`.
     */
    configType?: pulumi.Input<string>;
    /**
     * ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Configure the name of the metric to keep on.
     */
    metricsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TmpTkeBasicConfig resource.
 */
export interface TmpTkeBasicConfigArgs {
    /**
     * ID of cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Type of cluster.
     */
    clusterType: pulumi.Input<string>;
    /**
     * ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Configure the name of the metric to keep on.
     */
    metricsNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name. The naming rule is: namespace/name. If you don&#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
     */
    name?: pulumi.Input<string>;
}
