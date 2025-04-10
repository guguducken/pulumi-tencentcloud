// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this resource to create postgresql instance.
 *
 * > **Note:** To update the charge type, please update the `chargeType` and specify the `period` for the charging period. It only supports updating from `POSTPAID_BY_HOUR` to `PREPAID`, and the `period` field only valid in that upgrading case.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-1";
 * // create vpc
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * // create vpc subnet
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: availabilityZone,
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.20.0/28",
 *     isMulticast: false,
 * });
 * // create postgresql
 * const foo = new tencentcloud.postgresql.Instance("foo", {
 *     availabilityZone: availabilityZone,
 *     chargeType: "POSTPAID_BY_HOUR",
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     engineVersion: "10.4",
 *     rootUser: "root123",
 *     rootPassword: `Root123$`,
 *     charset: "UTF8",
 *     projectId: 0,
 *     memory: 2,
 *     storage: 10,
 *     tags: {
 *         test: "tf",
 *     },
 * });
 * ```
 * ### Create a multi available zone bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-6";
 * const standbyAvailabilityZone = config.get("standbyAvailabilityZone") || "ap-guangzhou-7";
 * // create vpc
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * // create vpc subnet
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: availabilityZone,
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.20.0/28",
 *     isMulticast: false,
 * });
 * // create postgresql
 * const foo = new tencentcloud.postgresql.Instance("foo", {
 *     availabilityZone: availabilityZone,
 *     chargeType: "POSTPAID_BY_HOUR",
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     engineVersion: "10.4",
 *     rootUser: "root123",
 *     rootPassword: `Root123$`,
 *     charset: "UTF8",
 *     projectId: 0,
 *     memory: 2,
 *     storage: 10,
 *     dbNodeSets: [
 *         {
 *             role: "Primary",
 *             zone: availabilityZone,
 *         },
 *         {
 *             zone: standbyAvailabilityZone,
 *         },
 *     ],
 *     tags: {
 *         test: "tf",
 *     },
 * });
 * ```
 * ### create pgsql with kms key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const pg = new tencentcloud.Postgresql.Instance("pg", {
 *     availabilityZone: "ap-guangzhou-6",
 *     backupPlan: {
 *         backupPeriods: [
 *             "tuesday",
 *             "wednesday",
 *         ],
 *         baseBackupRetentionPeriod: 7,
 *         maxBackupStartTime: "01:10:11",
 *         minBackupStartTime: "00:10:11",
 *     },
 *     chargeType: "POSTPAID_BY_HOUR",
 *     charset: "LATIN1",
 *     //  db_major_vesion   = "11"
 *     dbKernelVersion: "v11.12_r1.3",
 *     engineVersion: "11.12",
 *     kmsKeyId: "788c606a-c7b7-11ec-82d1-5254001e5c4e",
 *     kmsRegion: "ap-guangzhou",
 *     memory: 4,
 *     needSupportTde: 1,
 *     projectId: 0,
 *     rootPassword: "xxxxxxxxxx",
 *     storage: 100,
 *     subnetId: "subnet-enm92y0m",
 *     tags: {
 *         tf: "test",
 *     },
 *     vpcId: "vpc-86v957zb",
 * });
 * ```
 * ### upgrade kernel version
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const test = new tencentcloud.postgresql.Instance("test", {
 *     availabilityZone: data.tencentcloud_availability_zones_by_product.zone.zones[5].name,
 *     chargeType: "POSTPAID_BY_HOUR",
 *     vpcId: local.vpc_id,
 *     subnetId: local.subnet_id,
 *     engineVersion: "13.3",
 *     rootPassword: "*",
 *     charset: "LATIN1",
 *     projectId: 0,
 *     publicAccessSwitch: false,
 *     securityGroups: [local.sg_id],
 *     memory: 4,
 *     storage: 250,
 *     backupPlan: {
 *         minBackupStartTime: "01:10:11",
 *         maxBackupStartTime: "02:10:11",
 *         baseBackupRetentionPeriod: 5,
 *         backupPeriods: [
 *             "monday",
 *             "thursday",
 *             "sunday",
 *         ],
 *     },
 *     dbKernelVersion: "v13.3_r1.4",
 *     tags: {
 *         tf: "teest",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * postgresql instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Postgresql/instance:Instance foo postgres-cda1iex1
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Postgresql/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * Whether to use voucher, `1` for enabled.
     */
    public readonly autoVoucher!: pulumi.Output<number | undefined>;
    /**
     * Availability zone. NOTE: This field could not be modified, please use `dbNodeSet` instead of modification. The changes on this field will be suppressed when using the `dbNodeSet`.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Specify DB backup plan.
     */
    public readonly backupPlan!: pulumi.Output<outputs.Postgresql.InstanceBackupPlan | undefined>;
    /**
     * Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Charset of the root account. Valid values are `UTF8`,`LATIN1`.
     */
    public readonly charset!: pulumi.Output<string | undefined>;
    /**
     * Create time of the postgresql instance.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
     */
    public readonly dbKernelVersion!: pulumi.Output<string>;
    /**
     * PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
     */
    public readonly dbMajorVersion!: pulumi.Output<string>;
    /**
     * `dbMajorVesion` will be deprecated, use `dbMajorVersion` instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
     *
     * @deprecated `db_major_vesion` will be deprecated, use `db_major_version` instead.
     */
    public readonly dbMajorVesion!: pulumi.Output<string>;
    /**
     * Specify instance node info for disaster migration.
     */
    public readonly dbNodeSets!: pulumi.Output<outputs.Postgresql.InstanceDbNodeSet[] | undefined>;
    /**
     * Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * KeyId of the custom key.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Region of the custom key.
     */
    public readonly kmsRegion!: pulumi.Output<string>;
    /**
     * max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
     */
    public readonly maxStandbyArchiveDelay!: pulumi.Output<number>;
    /**
     * max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
     */
    public readonly maxStandbyStreamingDelay!: pulumi.Output<number>;
    /**
     * Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Name of the postgresql instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether to support data transparent encryption, 1: yes, 0: no (default).
     */
    public readonly needSupportTde!: pulumi.Output<number>;
    /**
     * Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * IP for private access.
     */
    public /*out*/ readonly privateAccessIp!: pulumi.Output<string>;
    /**
     * Port for private access.
     */
    public /*out*/ readonly privateAccessPort!: pulumi.Output<number>;
    /**
     * Project id, default value is `0`.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Host for public access.
     */
    public /*out*/ readonly publicAccessHost!: pulumi.Output<string>;
    /**
     * Port for public access.
     */
    public /*out*/ readonly publicAccessPort!: pulumi.Output<number>;
    /**
     * Indicates whether to enable the access to an instance from public network or not.
     */
    public readonly publicAccessSwitch!: pulumi.Output<boolean | undefined>;
    /**
     * Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
     */
    public readonly rootPassword!: pulumi.Output<string>;
    /**
     * Instance root account name. This parameter is optional, Default value is `root`.
     */
    public readonly rootUser!: pulumi.Output<string | undefined>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloud.Postgresql.getSpecinfos` provides.
     */
    public readonly storage!: pulumi.Output<number>;
    /**
     * ID of subnet.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The available tags within this postgresql.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Uid of the postgresql instance.
     */
    public /*out*/ readonly uid!: pulumi.Output<number>;
    /**
     * Specify Voucher Ids if `autoVoucher` was `1`, only support using 1 vouchers for now.
     */
    public readonly voucherIds!: pulumi.Output<string[] | undefined>;
    /**
     * ID of VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["backupPlan"] = state ? state.backupPlan : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["charset"] = state ? state.charset : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dbKernelVersion"] = state ? state.dbKernelVersion : undefined;
            resourceInputs["dbMajorVersion"] = state ? state.dbMajorVersion : undefined;
            resourceInputs["dbMajorVesion"] = state ? state.dbMajorVesion : undefined;
            resourceInputs["dbNodeSets"] = state ? state.dbNodeSets : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["kmsRegion"] = state ? state.kmsRegion : undefined;
            resourceInputs["maxStandbyArchiveDelay"] = state ? state.maxStandbyArchiveDelay : undefined;
            resourceInputs["maxStandbyStreamingDelay"] = state ? state.maxStandbyStreamingDelay : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["needSupportTde"] = state ? state.needSupportTde : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["privateAccessIp"] = state ? state.privateAccessIp : undefined;
            resourceInputs["privateAccessPort"] = state ? state.privateAccessPort : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["publicAccessHost"] = state ? state.publicAccessHost : undefined;
            resourceInputs["publicAccessPort"] = state ? state.publicAccessPort : undefined;
            resourceInputs["publicAccessSwitch"] = state ? state.publicAccessSwitch : undefined;
            resourceInputs["rootPassword"] = state ? state.rootPassword : undefined;
            resourceInputs["rootUser"] = state ? state.rootUser : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["voucherIds"] = state ? state.voucherIds : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.rootPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rootPassword'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["backupPlan"] = args ? args.backupPlan : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["charset"] = args ? args.charset : undefined;
            resourceInputs["dbKernelVersion"] = args ? args.dbKernelVersion : undefined;
            resourceInputs["dbMajorVersion"] = args ? args.dbMajorVersion : undefined;
            resourceInputs["dbMajorVesion"] = args ? args.dbMajorVesion : undefined;
            resourceInputs["dbNodeSets"] = args ? args.dbNodeSets : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["kmsRegion"] = args ? args.kmsRegion : undefined;
            resourceInputs["maxStandbyArchiveDelay"] = args ? args.maxStandbyArchiveDelay : undefined;
            resourceInputs["maxStandbyStreamingDelay"] = args ? args.maxStandbyStreamingDelay : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["needSupportTde"] = args ? args.needSupportTde : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["publicAccessSwitch"] = args ? args.publicAccessSwitch : undefined;
            resourceInputs["rootPassword"] = args ? args.rootPassword : undefined;
            resourceInputs["rootUser"] = args ? args.rootUser : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["voucherIds"] = args ? args.voucherIds : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["privateAccessIp"] = undefined /*out*/;
            resourceInputs["privateAccessPort"] = undefined /*out*/;
            resourceInputs["publicAccessHost"] = undefined /*out*/;
            resourceInputs["publicAccessPort"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * Whether to use voucher, `1` for enabled.
     */
    autoVoucher?: pulumi.Input<number>;
    /**
     * Availability zone. NOTE: This field could not be modified, please use `dbNodeSet` instead of modification. The changes on this field will be suppressed when using the `dbNodeSet`.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specify DB backup plan.
     */
    backupPlan?: pulumi.Input<inputs.Postgresql.InstanceBackupPlan>;
    /**
     * Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Charset of the root account. Valid values are `UTF8`,`LATIN1`.
     */
    charset?: pulumi.Input<string>;
    /**
     * Create time of the postgresql instance.
     */
    createTime?: pulumi.Input<string>;
    /**
     * PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
     */
    dbKernelVersion?: pulumi.Input<string>;
    /**
     * PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
     */
    dbMajorVersion?: pulumi.Input<string>;
    /**
     * `dbMajorVesion` will be deprecated, use `dbMajorVersion` instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
     *
     * @deprecated `db_major_vesion` will be deprecated, use `db_major_version` instead.
     */
    dbMajorVesion?: pulumi.Input<string>;
    /**
     * Specify instance node info for disaster migration.
     */
    dbNodeSets?: pulumi.Input<pulumi.Input<inputs.Postgresql.InstanceDbNodeSet>[]>;
    /**
     * Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * KeyId of the custom key.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Region of the custom key.
     */
    kmsRegion?: pulumi.Input<string>;
    /**
     * max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
     */
    maxStandbyArchiveDelay?: pulumi.Input<number>;
    /**
     * max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
     */
    maxStandbyStreamingDelay?: pulumi.Input<number>;
    /**
     * Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
     */
    memory?: pulumi.Input<number>;
    /**
     * Name of the postgresql instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to support data transparent encryption, 1: yes, 0: no (default).
     */
    needSupportTde?: pulumi.Input<number>;
    /**
     * Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
     */
    period?: pulumi.Input<number>;
    /**
     * IP for private access.
     */
    privateAccessIp?: pulumi.Input<string>;
    /**
     * Port for private access.
     */
    privateAccessPort?: pulumi.Input<number>;
    /**
     * Project id, default value is `0`.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Host for public access.
     */
    publicAccessHost?: pulumi.Input<string>;
    /**
     * Port for public access.
     */
    publicAccessPort?: pulumi.Input<number>;
    /**
     * Indicates whether to enable the access to an instance from public network or not.
     */
    publicAccessSwitch?: pulumi.Input<boolean>;
    /**
     * Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
     */
    rootPassword?: pulumi.Input<string>;
    /**
     * Instance root account name. This parameter is optional, Default value is `root`.
     */
    rootUser?: pulumi.Input<string>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloud.Postgresql.getSpecinfos` provides.
     */
    storage?: pulumi.Input<number>;
    /**
     * ID of subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The available tags within this postgresql.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Uid of the postgresql instance.
     */
    uid?: pulumi.Input<number>;
    /**
     * Specify Voucher Ids if `autoVoucher` was `1`, only support using 1 vouchers for now.
     */
    voucherIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * Whether to use voucher, `1` for enabled.
     */
    autoVoucher?: pulumi.Input<number>;
    /**
     * Availability zone. NOTE: This field could not be modified, please use `dbNodeSet` instead of modification. The changes on this field will be suppressed when using the `dbNodeSet`.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Specify DB backup plan.
     */
    backupPlan?: pulumi.Input<inputs.Postgresql.InstanceBackupPlan>;
    /**
     * Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Charset of the root account. Valid values are `UTF8`,`LATIN1`.
     */
    charset?: pulumi.Input<string>;
    /**
     * PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
     */
    dbKernelVersion?: pulumi.Input<string>;
    /**
     * PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
     */
    dbMajorVersion?: pulumi.Input<string>;
    /**
     * `dbMajorVesion` will be deprecated, use `dbMajorVersion` instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
     *
     * @deprecated `db_major_vesion` will be deprecated, use `db_major_version` instead.
     */
    dbMajorVesion?: pulumi.Input<string>;
    /**
     * Specify instance node info for disaster migration.
     */
    dbNodeSets?: pulumi.Input<pulumi.Input<inputs.Postgresql.InstanceDbNodeSet>[]>;
    /**
     * Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * KeyId of the custom key.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Region of the custom key.
     */
    kmsRegion?: pulumi.Input<string>;
    /**
     * max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
     */
    maxStandbyArchiveDelay?: pulumi.Input<number>;
    /**
     * max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
     */
    maxStandbyStreamingDelay?: pulumi.Input<number>;
    /**
     * Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
     */
    memory: pulumi.Input<number>;
    /**
     * Name of the postgresql instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to support data transparent encryption, 1: yes, 0: no (default).
     */
    needSupportTde?: pulumi.Input<number>;
    /**
     * Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
     */
    period?: pulumi.Input<number>;
    /**
     * Project id, default value is `0`.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Indicates whether to enable the access to an instance from public network or not.
     */
    publicAccessSwitch?: pulumi.Input<boolean>;
    /**
     * Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
     */
    rootPassword: pulumi.Input<string>;
    /**
     * Instance root account name. This parameter is optional, Default value is `root`.
     */
    rootUser?: pulumi.Input<string>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloud.Postgresql.getSpecinfos` provides.
     */
    storage: pulumi.Input<number>;
    /**
     * ID of subnet.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The available tags within this postgresql.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specify Voucher Ids if `autoVoucher` was `1`, only support using 1 vouchers for now.
     */
    voucherIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of VPC.
     */
    vpcId: pulumi.Input<string>;
}
