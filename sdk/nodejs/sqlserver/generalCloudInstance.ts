// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver generalCloudInstance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const generalCloudInstance = new tencentcloud.sqlserver.GeneralCloudInstance("generalCloudInstance", {
 *     zone: "ap-guangzhou-6",
 *     memory: 4,
 *     storage: 20,
 *     cpu: 2,
 *     machineType: "CLOUD_HSSD",
 *     instanceChargeType: "POSTPAID",
 *     projectId: 0,
 *     subnetId: local.subnet_id,
 *     vpcId: local.vpc_id,
 *     dbVersion: "2008R2",
 *     securityGroupLists: [local.sg_id],
 *     weeklies: [
 *         1,
 *         2,
 *         3,
 *         5,
 *         6,
 *         7,
 *     ],
 *     startTime: "00:00",
 *     span: 6,
 *     resourceTags: [{
 *         tagKey: "test",
 *         tagValue: "test",
 *     }],
 *     collation: "Chinese_PRC_CI_AS",
 *     timeZone: "China Standard Time",
 * });
 * ```
 *
 * ## Import
 *
 * sqlserver general_cloud_instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/generalCloudInstance:GeneralCloudInstance general_cloud_instance general_cloud_instance_id
 * ```
 */
export class GeneralCloudInstance extends pulumi.CustomResource {
    /**
     * Get an existing GeneralCloudInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GeneralCloudInstanceState, opts?: pulumi.CustomResourceOptions): GeneralCloudInstance {
        return new GeneralCloudInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/generalCloudInstance:GeneralCloudInstance';

    /**
     * Returns true if the given object is an instance of GeneralCloudInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GeneralCloudInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GeneralCloudInstance.__pulumiType;
    }

    /**
     * Automatic renewal flag: 0-normal renewal 1-automatic renewal, the default is 1 automatic renewal. Valid only when purchasing a prepaid instance. Valid only when the 'instance_charge_type' parameter value is 'PREPAID'.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * System character set collation, default: Chinese_PRC_CI_AS.
     */
    public readonly collation!: pulumi.Output<string | undefined>;
    /**
     * Cpu, unit: CORE.
     */
    public readonly cpu!: pulumi.Output<number>;
    /**
     * sqlserver version, currently all supported versions are: 2008R2 (SQL Server 2008 R2 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 201202 (SQL Server 2012 Standard), 2014SP2 (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise), 201702 (SQL Server 2017 Standard), 2019 (SQL Server 2019 Enterprise), 201902 (SQL Server 2019 Standard). Each region supports different versions for sale, and the version information that can be sold in each region can be pulled through the DescribeProductConfig interface. If left blank, the default version is 2008R2.
     */
    public readonly dbVersion!: pulumi.Output<string | undefined>;
    /**
     * It has been deprecated from version 1.81.2. Upgrade the high-availability architecture of sqlserver, upgrade from mirror disaster recovery to always on cluster disaster recovery, only support 2017 and above and support always on high-availability instances, do not support downgrading to mirror disaster recovery, CLUSTER-upgrade to always on capacity Disaster, if not filled, the high-availability architecture will not be modified.
     *
     * @deprecated It has been deprecated from version 1.81.2.
     */
    public readonly haType!: pulumi.Output<string | undefined>;
    /**
     * Payment mode, the value supports PREPAID (prepaid), POSTPAID (postpaid).
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * The host disk type of the purchased instance, CLOUD_HSSD-enhanced SSD cloud disk for virtual machines, CLOUD_TSSD-extremely fast SSD cloud disk for virtual machines, CLOUD_BSSD-universal SSD cloud disk for virtual machines.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Memory, unit: GB.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Name of the SQL Server instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Purchase instance period, the default value is 1, which means one month. The value cannot exceed 48. Valid only when the 'instance_charge_type' parameter value is 'PREPAID'.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * project ID.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * A collection of tags bound to the new instance.
     */
    public readonly resourceTags!: pulumi.Output<outputs.Sqlserver.GeneralCloudInstanceResourceTag[] | undefined>;
    /**
     * Security group list, fill in the security group ID in the form of sg-xxx.
     */
    public readonly securityGroupLists!: pulumi.Output<string[] | undefined>;
    /**
     * Maintainable time window configuration, duration, unit: hour.
     */
    public readonly span!: pulumi.Output<number | undefined>;
    /**
     * Maintainable time window configuration, daily maintainable start time.
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * instance disk storage, unit: GB.
     */
    public readonly storage!: pulumi.Output<number>;
    /**
     * VPC subnet ID, in the form of subnet-bdoe83fa; SubnetId and VpcId need to be set at the same time or not set at the same time.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * System time zone, default: China Standard Time.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;
    /**
     * VPC network ID, in the form of vpc-dsp338hz; SubnetId and VpcId need to be set at the same time or not set at the same time.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * Maintainable time window configuration, in weeks, indicates the days of the week that allow maintenance, 1-7 represent Monday to weekend respectively.
     */
    public readonly weeklies!: pulumi.Output<number[] | undefined>;
    /**
     * Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the DescribeZones API.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a GeneralCloudInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GeneralCloudInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GeneralCloudInstanceArgs | GeneralCloudInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GeneralCloudInstanceState | undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["collation"] = state ? state.collation : undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["dbVersion"] = state ? state.dbVersion : undefined;
            resourceInputs["haType"] = state ? state.haType : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["machineType"] = state ? state.machineType : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["resourceTags"] = state ? state.resourceTags : undefined;
            resourceInputs["securityGroupLists"] = state ? state.securityGroupLists : undefined;
            resourceInputs["span"] = state ? state.span : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["weeklies"] = state ? state.weeklies : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as GeneralCloudInstanceArgs | undefined;
            if ((!args || args.cpu === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cpu'");
            }
            if ((!args || args.machineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineType'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["collation"] = args ? args.collation : undefined;
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["dbVersion"] = args ? args.dbVersion : undefined;
            resourceInputs["haType"] = args ? args.haType : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["machineType"] = args ? args.machineType : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resourceTags"] = args ? args.resourceTags : undefined;
            resourceInputs["securityGroupLists"] = args ? args.securityGroupLists : undefined;
            resourceInputs["span"] = args ? args.span : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["weeklies"] = args ? args.weeklies : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GeneralCloudInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GeneralCloudInstance resources.
 */
export interface GeneralCloudInstanceState {
    /**
     * Automatic renewal flag: 0-normal renewal 1-automatic renewal, the default is 1 automatic renewal. Valid only when purchasing a prepaid instance. Valid only when the 'instance_charge_type' parameter value is 'PREPAID'.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * System character set collation, default: Chinese_PRC_CI_AS.
     */
    collation?: pulumi.Input<string>;
    /**
     * Cpu, unit: CORE.
     */
    cpu?: pulumi.Input<number>;
    /**
     * sqlserver version, currently all supported versions are: 2008R2 (SQL Server 2008 R2 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 201202 (SQL Server 2012 Standard), 2014SP2 (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise), 201702 (SQL Server 2017 Standard), 2019 (SQL Server 2019 Enterprise), 201902 (SQL Server 2019 Standard). Each region supports different versions for sale, and the version information that can be sold in each region can be pulled through the DescribeProductConfig interface. If left blank, the default version is 2008R2.
     */
    dbVersion?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.81.2. Upgrade the high-availability architecture of sqlserver, upgrade from mirror disaster recovery to always on cluster disaster recovery, only support 2017 and above and support always on high-availability instances, do not support downgrading to mirror disaster recovery, CLUSTER-upgrade to always on capacity Disaster, if not filled, the high-availability architecture will not be modified.
     *
     * @deprecated It has been deprecated from version 1.81.2.
     */
    haType?: pulumi.Input<string>;
    /**
     * Payment mode, the value supports PREPAID (prepaid), POSTPAID (postpaid).
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The host disk type of the purchased instance, CLOUD_HSSD-enhanced SSD cloud disk for virtual machines, CLOUD_TSSD-extremely fast SSD cloud disk for virtual machines, CLOUD_BSSD-universal SSD cloud disk for virtual machines.
     */
    machineType?: pulumi.Input<string>;
    /**
     * Memory, unit: GB.
     */
    memory?: pulumi.Input<number>;
    /**
     * Name of the SQL Server instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Purchase instance period, the default value is 1, which means one month. The value cannot exceed 48. Valid only when the 'instance_charge_type' parameter value is 'PREPAID'.
     */
    period?: pulumi.Input<number>;
    /**
     * project ID.
     */
    projectId?: pulumi.Input<number>;
    /**
     * A collection of tags bound to the new instance.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.Sqlserver.GeneralCloudInstanceResourceTag>[]>;
    /**
     * Security group list, fill in the security group ID in the form of sg-xxx.
     */
    securityGroupLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maintainable time window configuration, duration, unit: hour.
     */
    span?: pulumi.Input<number>;
    /**
     * Maintainable time window configuration, daily maintainable start time.
     */
    startTime?: pulumi.Input<string>;
    /**
     * instance disk storage, unit: GB.
     */
    storage?: pulumi.Input<number>;
    /**
     * VPC subnet ID, in the form of subnet-bdoe83fa; SubnetId and VpcId need to be set at the same time or not set at the same time.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * System time zone, default: China Standard Time.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * VPC network ID, in the form of vpc-dsp338hz; SubnetId and VpcId need to be set at the same time or not set at the same time.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Maintainable time window configuration, in weeks, indicates the days of the week that allow maintenance, 1-7 represent Monday to weekend respectively.
     */
    weeklies?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the DescribeZones API.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GeneralCloudInstance resource.
 */
export interface GeneralCloudInstanceArgs {
    /**
     * Automatic renewal flag: 0-normal renewal 1-automatic renewal, the default is 1 automatic renewal. Valid only when purchasing a prepaid instance. Valid only when the 'instance_charge_type' parameter value is 'PREPAID'.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * System character set collation, default: Chinese_PRC_CI_AS.
     */
    collation?: pulumi.Input<string>;
    /**
     * Cpu, unit: CORE.
     */
    cpu: pulumi.Input<number>;
    /**
     * sqlserver version, currently all supported versions are: 2008R2 (SQL Server 2008 R2 Enterprise), 2012SP3 (SQL Server 2012 Enterprise), 201202 (SQL Server 2012 Standard), 2014SP2 (SQL Server 2014 Enterprise), 201402 (SQL Server 2014 Standard), 2016SP1 (SQL Server 2016 Enterprise), 201602 (SQL Server 2016 Standard), 2017 (SQL Server 2017 Enterprise), 201702 (SQL Server 2017 Standard), 2019 (SQL Server 2019 Enterprise), 201902 (SQL Server 2019 Standard). Each region supports different versions for sale, and the version information that can be sold in each region can be pulled through the DescribeProductConfig interface. If left blank, the default version is 2008R2.
     */
    dbVersion?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.81.2. Upgrade the high-availability architecture of sqlserver, upgrade from mirror disaster recovery to always on cluster disaster recovery, only support 2017 and above and support always on high-availability instances, do not support downgrading to mirror disaster recovery, CLUSTER-upgrade to always on capacity Disaster, if not filled, the high-availability architecture will not be modified.
     *
     * @deprecated It has been deprecated from version 1.81.2.
     */
    haType?: pulumi.Input<string>;
    /**
     * Payment mode, the value supports PREPAID (prepaid), POSTPAID (postpaid).
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The host disk type of the purchased instance, CLOUD_HSSD-enhanced SSD cloud disk for virtual machines, CLOUD_TSSD-extremely fast SSD cloud disk for virtual machines, CLOUD_BSSD-universal SSD cloud disk for virtual machines.
     */
    machineType: pulumi.Input<string>;
    /**
     * Memory, unit: GB.
     */
    memory: pulumi.Input<number>;
    /**
     * Name of the SQL Server instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Purchase instance period, the default value is 1, which means one month. The value cannot exceed 48. Valid only when the 'instance_charge_type' parameter value is 'PREPAID'.
     */
    period?: pulumi.Input<number>;
    /**
     * project ID.
     */
    projectId?: pulumi.Input<number>;
    /**
     * A collection of tags bound to the new instance.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.Sqlserver.GeneralCloudInstanceResourceTag>[]>;
    /**
     * Security group list, fill in the security group ID in the form of sg-xxx.
     */
    securityGroupLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maintainable time window configuration, duration, unit: hour.
     */
    span?: pulumi.Input<number>;
    /**
     * Maintainable time window configuration, daily maintainable start time.
     */
    startTime?: pulumi.Input<string>;
    /**
     * instance disk storage, unit: GB.
     */
    storage: pulumi.Input<number>;
    /**
     * VPC subnet ID, in the form of subnet-bdoe83fa; SubnetId and VpcId need to be set at the same time or not set at the same time.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * System time zone, default: China Standard Time.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * VPC network ID, in the form of vpc-dsp338hz; SubnetId and VpcId need to be set at the same time or not set at the same time.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Maintainable time window configuration, in weeks, indicates the days of the week that allow maintenance, 1-7 represent Monday to weekend respectively.
     */
    weeklies?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Instance AZ, such as ap-guangzhou-1 (Guangzhou Zone 1). Purchasable AZs for an instance can be obtained through the DescribeZones API.
     */
    zone: pulumi.Input<string>;
}
