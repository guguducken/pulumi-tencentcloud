// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql renewDbInstanceOperation
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const renewDbInstanceOperation = new tencentcloud.Mysql.RenewDbInstanceOperation("renew_db_instance_operation", {
 *     instanceId: "cdb-c1nl9rpv",
 *     modifyPayType: "PREPAID",
 *     timeSpan: 1,
 * });
 * ```
 */
export class RenewDbInstanceOperation extends pulumi.CustomResource {
    /**
     * Get an existing RenewDbInstanceOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RenewDbInstanceOperationState, opts?: pulumi.CustomResourceOptions): RenewDbInstanceOperation {
        return new RenewDbInstanceOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/renewDbInstanceOperation:RenewDbInstanceOperation';

    /**
     * Returns true if the given object is an instance of RenewDbInstanceOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RenewDbInstanceOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RenewDbInstanceOperation.__pulumiType;
    }

    /**
     * Instance expiration time.
     */
    public /*out*/ readonly deadlineTime!: pulumi.Output<string>;
    /**
     * Deal id.
     */
    public /*out*/ readonly dealId!: pulumi.Output<string>;
    /**
     * The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
     */
    public readonly modifyPayType!: pulumi.Output<string | undefined>;
    /**
     * Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
     */
    public readonly timeSpan!: pulumi.Output<number>;

    /**
     * Create a RenewDbInstanceOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RenewDbInstanceOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RenewDbInstanceOperationArgs | RenewDbInstanceOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RenewDbInstanceOperationState | undefined;
            resourceInputs["deadlineTime"] = state ? state.deadlineTime : undefined;
            resourceInputs["dealId"] = state ? state.dealId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["modifyPayType"] = state ? state.modifyPayType : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
        } else {
            const args = argsOrState as RenewDbInstanceOperationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.timeSpan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeSpan'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["modifyPayType"] = args ? args.modifyPayType : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
            resourceInputs["deadlineTime"] = undefined /*out*/;
            resourceInputs["dealId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RenewDbInstanceOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RenewDbInstanceOperation resources.
 */
export interface RenewDbInstanceOperationState {
    /**
     * Instance expiration time.
     */
    deadlineTime?: pulumi.Input<string>;
    /**
     * Deal id.
     */
    dealId?: pulumi.Input<string>;
    /**
     * The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
     */
    instanceId?: pulumi.Input<string>;
    /**
     * If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
     */
    modifyPayType?: pulumi.Input<string>;
    /**
     * Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
     */
    timeSpan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RenewDbInstanceOperation resource.
 */
export interface RenewDbInstanceOperationArgs {
    /**
     * The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
     */
    instanceId: pulumi.Input<string>;
    /**
     * If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
     */
    modifyPayType?: pulumi.Input<string>;
    /**
     * Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
     */
    timeSpan: pulumi.Input<number>;
}
