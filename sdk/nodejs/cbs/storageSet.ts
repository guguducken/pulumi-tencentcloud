// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create CBS set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const storage = new tencentcloud.Cbs.StorageSet("storage", {
 *     availabilityZone: "ap-guangzhou-3",
 *     diskCount: 10,
 *     encrypt: false,
 *     projectId: 0,
 *     storageName: "mystorage",
 *     storageSize: 100,
 *     storageType: "CLOUD_SSD",
 * });
 * ```
 */
export class StorageSet extends pulumi.CustomResource {
    /**
     * Get an existing StorageSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageSetState, opts?: pulumi.CustomResourceOptions): StorageSet {
        return new StorageSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cbs/storageSet:StorageSet';

    /**
     * Returns true if the given object is an instance of StorageSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageSet.__pulumiType;
    }

    /**
     * Indicates whether the CBS is mounted the CVM.
     */
    public /*out*/ readonly attached!: pulumi.Output<boolean>;
    /**
     * The available zone that the CBS instance locates at.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * The number of disks to be purchased. Default 1.
     */
    public readonly diskCount!: pulumi.Output<number | undefined>;
    /**
     * disk id list.
     */
    public /*out*/ readonly diskIds!: pulumi.Output<string[]>;
    /**
     * Indicates whether CBS is encrypted.
     */
    public readonly encrypt!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the project to which the instance belongs.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * ID of the snapshot. If specified, created the CBS by this snapshot.
     */
    public readonly snapshotId!: pulumi.Output<string>;
    /**
     * Name of CBS. The maximum length can not exceed 60 bytes.
     */
    public readonly storageName!: pulumi.Output<string>;
    /**
     * Volume of CBS, and unit is GB.
     */
    public readonly storageSize!: pulumi.Output<number>;
    /**
     * Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
     */
    public /*out*/ readonly storageStatus!: pulumi.Output<string>;
    /**
     * Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
     */
    public readonly throughputPerformance!: pulumi.Output<number | undefined>;

    /**
     * Create a StorageSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageSetArgs | StorageSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StorageSetState | undefined;
            resourceInputs["attached"] = state ? state.attached : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["diskCount"] = state ? state.diskCount : undefined;
            resourceInputs["diskIds"] = state ? state.diskIds : undefined;
            resourceInputs["encrypt"] = state ? state.encrypt : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["storageName"] = state ? state.storageName : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["storageStatus"] = state ? state.storageStatus : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["throughputPerformance"] = state ? state.throughputPerformance : undefined;
        } else {
            const args = argsOrState as StorageSetArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.storageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageName'");
            }
            if ((!args || args.storageSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSize'");
            }
            if ((!args || args.storageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageType'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["diskCount"] = args ? args.diskCount : undefined;
            resourceInputs["encrypt"] = args ? args.encrypt : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["storageName"] = args ? args.storageName : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["throughputPerformance"] = args ? args.throughputPerformance : undefined;
            resourceInputs["attached"] = undefined /*out*/;
            resourceInputs["diskIds"] = undefined /*out*/;
            resourceInputs["storageStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StorageSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StorageSet resources.
 */
export interface StorageSetState {
    /**
     * Indicates whether the CBS is mounted the CVM.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * The available zone that the CBS instance locates at.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The number of disks to be purchased. Default 1.
     */
    diskCount?: pulumi.Input<number>;
    /**
     * disk id list.
     */
    diskIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether CBS is encrypted.
     */
    encrypt?: pulumi.Input<boolean>;
    /**
     * ID of the project to which the instance belongs.
     */
    projectId?: pulumi.Input<number>;
    /**
     * ID of the snapshot. If specified, created the CBS by this snapshot.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Name of CBS. The maximum length can not exceed 60 bytes.
     */
    storageName?: pulumi.Input<string>;
    /**
     * Volume of CBS, and unit is GB.
     */
    storageSize?: pulumi.Input<number>;
    /**
     * Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
     */
    storageStatus?: pulumi.Input<string>;
    /**
     * Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
     */
    storageType?: pulumi.Input<string>;
    /**
     * Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
     */
    throughputPerformance?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a StorageSet resource.
 */
export interface StorageSetArgs {
    /**
     * The available zone that the CBS instance locates at.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The number of disks to be purchased. Default 1.
     */
    diskCount?: pulumi.Input<number>;
    /**
     * Indicates whether CBS is encrypted.
     */
    encrypt?: pulumi.Input<boolean>;
    /**
     * ID of the project to which the instance belongs.
     */
    projectId?: pulumi.Input<number>;
    /**
     * ID of the snapshot. If specified, created the CBS by this snapshot.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Name of CBS. The maximum length can not exceed 60 bytes.
     */
    storageName: pulumi.Input<string>;
    /**
     * Volume of CBS, and unit is GB.
     */
    storageSize: pulumi.Input<number>;
    /**
     * Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
     */
    storageType: pulumi.Input<string>;
    /**
     * Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
     */
    throughputPerformance?: pulumi.Input<number>;
}
