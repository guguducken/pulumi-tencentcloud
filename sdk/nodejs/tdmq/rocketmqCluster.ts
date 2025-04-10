// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tdmqRocketmq cluster
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const cluster = new tencentcloud.Tdmq.RocketmqCluster("cluster", {
 *     clusterName: "test_rocketmq",
 *     remark: "test rocket mq",
 * });
 * ```
 *
 * ## Import
 *
 * tdmqRocketmq cluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster cluster cluster_id
 * ```
 */
export class RocketmqCluster extends pulumi.CustomResource {
    /**
     * Get an existing RocketmqCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RocketmqClusterState, opts?: pulumi.CustomResourceOptions): RocketmqCluster {
        return new RocketmqCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster';

    /**
     * Returns true if the given object is an instance of RocketmqCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RocketmqCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RocketmqCluster.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Creation time in milliseconds.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * Whether it is an exclusive instance.
     */
    public /*out*/ readonly isVip!: pulumi.Output<boolean>;
    /**
     * Public network access address.
     */
    public /*out*/ readonly publicEndPoint!: pulumi.Output<string>;
    /**
     * Region information.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Cluster description (up to 128 characters).
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Rocketmq cluster identification.
     */
    public /*out*/ readonly rocketMQFlag!: pulumi.Output<boolean>;
    /**
     * Whether the namespace access point is supported.
     */
    public /*out*/ readonly supportNamespaceEndpoint!: pulumi.Output<boolean>;
    /**
     * VPC access address.
     */
    public /*out*/ readonly vpcEndPoint!: pulumi.Output<string>;
    /**
     * Vpc list.
     */
    public /*out*/ readonly vpcs!: pulumi.Output<outputs.Tdmq.RocketmqClusterVpc[]>;

    /**
     * Create a RocketmqCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RocketmqClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RocketmqClusterArgs | RocketmqClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RocketmqClusterState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["isVip"] = state ? state.isVip : undefined;
            resourceInputs["publicEndPoint"] = state ? state.publicEndPoint : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["rocketMQFlag"] = state ? state.rocketMQFlag : undefined;
            resourceInputs["supportNamespaceEndpoint"] = state ? state.supportNamespaceEndpoint : undefined;
            resourceInputs["vpcEndPoint"] = state ? state.vpcEndPoint : undefined;
            resourceInputs["vpcs"] = state ? state.vpcs : undefined;
        } else {
            const args = argsOrState as RocketmqClusterArgs | undefined;
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["clusterId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["isVip"] = undefined /*out*/;
            resourceInputs["publicEndPoint"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["rocketMQFlag"] = undefined /*out*/;
            resourceInputs["supportNamespaceEndpoint"] = undefined /*out*/;
            resourceInputs["vpcEndPoint"] = undefined /*out*/;
            resourceInputs["vpcs"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RocketmqCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RocketmqCluster resources.
 */
export interface RocketmqClusterState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Creation time in milliseconds.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Whether it is an exclusive instance.
     */
    isVip?: pulumi.Input<boolean>;
    /**
     * Public network access address.
     */
    publicEndPoint?: pulumi.Input<string>;
    /**
     * Region information.
     */
    region?: pulumi.Input<string>;
    /**
     * Cluster description (up to 128 characters).
     */
    remark?: pulumi.Input<string>;
    /**
     * Rocketmq cluster identification.
     */
    rocketMQFlag?: pulumi.Input<boolean>;
    /**
     * Whether the namespace access point is supported.
     */
    supportNamespaceEndpoint?: pulumi.Input<boolean>;
    /**
     * VPC access address.
     */
    vpcEndPoint?: pulumi.Input<string>;
    /**
     * Vpc list.
     */
    vpcs?: pulumi.Input<pulumi.Input<inputs.Tdmq.RocketmqClusterVpc>[]>;
}

/**
 * The set of arguments for constructing a RocketmqCluster resource.
 */
export interface RocketmqClusterArgs {
    /**
     * Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Cluster description (up to 128 characters).
     */
    remark?: pulumi.Input<string>;
}
