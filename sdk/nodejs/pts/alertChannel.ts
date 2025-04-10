// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a pts alertChannel
 *
 * > **NOTE:** Modification is not currently supported, please go to the console to modify.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.monitor.AlarmNotice("example", {
 *     noticeType: "ALL",
 *     noticeLanguage: "zh-CN",
 *     userNotices: [{
 *         receiverType: "USER",
 *         startTime: 0,
 *         endTime: 1,
 *         noticeWays: [
 *             "EMAIL",
 *             "SMS",
 *             "WECHAT",
 *         ],
 *         userIds: [10001],
 *         groupIds: [],
 *         phoneOrders: [10001],
 *         phoneCircleTimes: 2,
 *         phoneCircleInterval: 50,
 *         phoneInnerInterval: 60,
 *         needPhoneArriveNotice: 1,
 *         phoneCallType: "CIRCLE",
 *         weekdays: [
 *             1,
 *             2,
 *             3,
 *             4,
 *             5,
 *             6,
 *             7,
 *         ],
 *     }],
 *     urlNotices: [{
 *         url: "https://www.mytest.com/validate",
 *         endTime: 0,
 *         startTime: 1,
 *         weekdays: [
 *             1,
 *             2,
 *             3,
 *             4,
 *             5,
 *             6,
 *             7,
 *         ],
 *     }],
 * });
 * const project = new tencentcloud.pts.Project("project", {
 *     description: "desc",
 *     tags: [{
 *         tagKey: "createdBy",
 *         tagValue: "terraform",
 *     }],
 * });
 * const alertChannel = new tencentcloud.pts.AlertChannel("alertChannel", {
 *     noticeId: example.id,
 *     projectId: project.id,
 *     ampConsumerId: "Consumer-vvy1xxxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * pts alert_channel can be imported using the project_id#notice_id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Pts/alertChannel:AlertChannel alert_channel project-kww5v8se#notice-kl66t6y9
 * ```
 */
export class AlertChannel extends pulumi.CustomResource {
    /**
     * Get an existing AlertChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertChannelState, opts?: pulumi.CustomResourceOptions): AlertChannel {
        return new AlertChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Pts/alertChannel:AlertChannel';

    /**
     * Returns true if the given object is an instance of AlertChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertChannel.__pulumiType;
    }

    /**
     * AMP Consumer ID.
     */
    public readonly ampConsumerId!: pulumi.Output<string | undefined>;
    /**
     * App ID Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly appId!: pulumi.Output<number>;
    /**
     * Creation time Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Notice ID.
     */
    public readonly noticeId!: pulumi.Output<string>;
    /**
     * Project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Status Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Sub-user ID Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly subAccountUin!: pulumi.Output<string>;
    /**
     * User ID Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly uin!: pulumi.Output<string>;
    /**
     * Update time Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a AlertChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertChannelArgs | AlertChannelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertChannelState | undefined;
            resourceInputs["ampConsumerId"] = state ? state.ampConsumerId : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["noticeId"] = state ? state.noticeId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subAccountUin"] = state ? state.subAccountUin : undefined;
            resourceInputs["uin"] = state ? state.uin : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as AlertChannelArgs | undefined;
            if ((!args || args.noticeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'noticeId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["ampConsumerId"] = args ? args.ampConsumerId : undefined;
            resourceInputs["noticeId"] = args ? args.noticeId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subAccountUin"] = undefined /*out*/;
            resourceInputs["uin"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AlertChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertChannel resources.
 */
export interface AlertChannelState {
    /**
     * AMP Consumer ID.
     */
    ampConsumerId?: pulumi.Input<string>;
    /**
     * App ID Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    appId?: pulumi.Input<number>;
    /**
     * Creation time Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Notice ID.
     */
    noticeId?: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Status Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    status?: pulumi.Input<number>;
    /**
     * Sub-user ID Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    subAccountUin?: pulumi.Input<string>;
    /**
     * User ID Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    uin?: pulumi.Input<string>;
    /**
     * Update time Note: this field may return null, indicating that a valid value cannot be obtained.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlertChannel resource.
 */
export interface AlertChannelArgs {
    /**
     * AMP Consumer ID.
     */
    ampConsumerId?: pulumi.Input<string>;
    /**
     * Notice ID.
     */
    noticeId: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId: pulumi.Input<string>;
}
