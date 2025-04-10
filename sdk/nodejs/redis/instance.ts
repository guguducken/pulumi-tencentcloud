// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a Redis instance and set its attributes.
 *
 * > **NOTE:** The argument vpcId and subnetId is now required because Basic Network Instance is no longer supported.
 *
 * > **NOTE:** Both adding and removing replications in one change is supported but not recommend.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zone = tencentcloud.Redis.getZoneConfig({});
 * const redisInstanceTest2 = new tencentcloud.redis.Instance("redisInstanceTest2", {
 *     availabilityZone: zone.then(zone => zone.lists?[0]?.zone),
 *     typeId: zone.then(zone => zone.lists?[0]?.typeId),
 *     password: "test12345789",
 *     memSize: 8192,
 *     redisShardNum: zone.then(zone => zone.lists?[0]?.redisShardNums?[0]),
 *     redisReplicasNum: zone.then(zone => zone.lists?[0]?.redisReplicasNums?[0]),
 *     port: 6379,
 * });
 * ```
 *
 * ## Import
 *
 * Redis instance can be imported, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Redis/instance:Instance redislab redis-id
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
    public static readonly __pulumiType = 'tencentcloud:Redis/instance:Instance';

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
     * Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * The time when the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * IP address of an instance. When the `operationNetwork` is `changeVip`, this parameter needs to be configured.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.
     */
    public readonly memSize!: pulumi.Output<number>;
    /**
     * Instance name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
     */
    public readonly noAuth!: pulumi.Output<boolean | undefined>;
    /**
     * Readonly Primary/Replica nodes.
     */
    public /*out*/ readonly nodeInfos!: pulumi.Output<outputs.Redis.InstanceNodeInfo[]>;
    /**
     * Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
     */
    public readonly operationNetwork!: pulumi.Output<string | undefined>;
    /**
     * Specify params template id. If not set, will use default template.
     */
    public readonly paramsTemplateId!: pulumi.Output<string | undefined>;
    /**
     * Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password empty.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The port used to access a redis instance. The default value is 6379. When the `operationNetwork` is `changeVPort` or `changeVip`, this parameter needs to be configured.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when chargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
     */
    public readonly prepaidPeriod!: pulumi.Output<number | undefined>;
    /**
     * Specifies which project the instance should belong to.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Original intranet IPv4 address retention time: unit: day, value range: `0`, `1`, `2`, `3`, `7`, `15`.
     */
    public readonly recycle!: pulumi.Output<number | undefined>;
    /**
     * The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replicaZoneIds`.
     */
    public readonly redisReplicasNum!: pulumi.Output<number | undefined>;
    /**
     * The number of instance shard, default is 1. This is not required for standalone and master slave versions.
     */
    public readonly redisShardNum!: pulumi.Output<number>;
    /**
     * ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
     */
    public readonly replicaZoneIds!: pulumi.Output<number[]>;
    /**
     * Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
     */
    public readonly replicasReadOnly!: pulumi.Output<boolean>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Current status of an instance, maybe: init, processing, online, isolate and todelete.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies which subnet the instance should belong to. When the `operationNetwork` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Instance tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: `clusterCkv`,`cluster_redis5.0`,`clusterRedis`,`masterSlaveCkv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`masterSlaveRedis`,`standaloneRedis`, specific region support specific types, need to refer data `tencentcloud.Redis.getZoneConfig`.
     *
     * @deprecated It has been deprecated from version 1.33.1. Please use 'type_id' instead.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Instance type. Available values reference data source `tencentcloud.Redis.getZoneConfig` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
     */
    public readonly typeId!: pulumi.Output<number | undefined>;
    /**
     * ID of the vpc with which the instance is to be associated. When the `operationNetwork` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
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
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["memSize"] = state ? state.memSize : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["noAuth"] = state ? state.noAuth : undefined;
            resourceInputs["nodeInfos"] = state ? state.nodeInfos : undefined;
            resourceInputs["operationNetwork"] = state ? state.operationNetwork : undefined;
            resourceInputs["paramsTemplateId"] = state ? state.paramsTemplateId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["prepaidPeriod"] = state ? state.prepaidPeriod : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["recycle"] = state ? state.recycle : undefined;
            resourceInputs["redisReplicasNum"] = state ? state.redisReplicasNum : undefined;
            resourceInputs["redisShardNum"] = state ? state.redisShardNum : undefined;
            resourceInputs["replicaZoneIds"] = state ? state.replicaZoneIds : undefined;
            resourceInputs["replicasReadOnly"] = state ? state.replicasReadOnly : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["typeId"] = state ? state.typeId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.memSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memSize'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["memSize"] = args ? args.memSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["noAuth"] = args ? args.noAuth : undefined;
            resourceInputs["operationNetwork"] = args ? args.operationNetwork : undefined;
            resourceInputs["paramsTemplateId"] = args ? args.paramsTemplateId : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["prepaidPeriod"] = args ? args.prepaidPeriod : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["recycle"] = args ? args.recycle : undefined;
            resourceInputs["redisReplicasNum"] = args ? args.redisReplicasNum : undefined;
            resourceInputs["redisShardNum"] = args ? args.redisShardNum : undefined;
            resourceInputs["replicaZoneIds"] = args ? args.replicaZoneIds : undefined;
            resourceInputs["replicasReadOnly"] = args ? args.replicasReadOnly : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["typeId"] = args ? args.typeId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["nodeInfos"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
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
     * Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The time when the instance was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * IP address of an instance. When the `operationNetwork` is `changeVip`, this parameter needs to be configured.
     */
    ip?: pulumi.Input<string>;
    /**
     * The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.
     */
    memSize?: pulumi.Input<number>;
    /**
     * Instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
     */
    noAuth?: pulumi.Input<boolean>;
    /**
     * Readonly Primary/Replica nodes.
     */
    nodeInfos?: pulumi.Input<pulumi.Input<inputs.Redis.InstanceNodeInfo>[]>;
    /**
     * Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
     */
    operationNetwork?: pulumi.Input<string>;
    /**
     * Specify params template id. If not set, will use default template.
     */
    paramsTemplateId?: pulumi.Input<string>;
    /**
     * Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password empty.
     */
    password?: pulumi.Input<string>;
    /**
     * The port used to access a redis instance. The default value is 6379. When the `operationNetwork` is `changeVPort` or `changeVip`, this parameter needs to be configured.
     */
    port?: pulumi.Input<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when chargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * Specifies which project the instance should belong to.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Original intranet IPv4 address retention time: unit: day, value range: `0`, `1`, `2`, `3`, `7`, `15`.
     */
    recycle?: pulumi.Input<number>;
    /**
     * The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replicaZoneIds`.
     */
    redisReplicasNum?: pulumi.Input<number>;
    /**
     * The number of instance shard, default is 1. This is not required for standalone and master slave versions.
     */
    redisShardNum?: pulumi.Input<number>;
    /**
     * ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
     */
    replicaZoneIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
     */
    replicasReadOnly?: pulumi.Input<boolean>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Current status of an instance, maybe: init, processing, online, isolate and todelete.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies which subnet the instance should belong to. When the `operationNetwork` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Instance tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: `clusterCkv`,`cluster_redis5.0`,`clusterRedis`,`masterSlaveCkv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`masterSlaveRedis`,`standaloneRedis`, specific region support specific types, need to refer data `tencentcloud.Redis.getZoneConfig`.
     *
     * @deprecated It has been deprecated from version 1.33.1. Please use 'type_id' instead.
     */
    type?: pulumi.Input<string>;
    /**
     * Instance type. Available values reference data source `tencentcloud.Redis.getZoneConfig` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
     */
    typeId?: pulumi.Input<number>;
    /**
     * ID of the vpc with which the instance is to be associated. When the `operationNetwork` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * IP address of an instance. When the `operationNetwork` is `changeVip`, this parameter needs to be configured.
     */
    ip?: pulumi.Input<string>;
    /**
     * The memory volume of an available instance(in MB), please refer to `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.
     */
    memSize: pulumi.Input<number>;
    /**
     * Instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
     */
    noAuth?: pulumi.Input<boolean>;
    /**
     * Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
     */
    operationNetwork?: pulumi.Input<string>;
    /**
     * Specify params template id. If not set, will use default template.
     */
    paramsTemplateId?: pulumi.Input<string>;
    /**
     * Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password empty.
     */
    password?: pulumi.Input<string>;
    /**
     * The port used to access a redis instance. The default value is 6379. When the `operationNetwork` is `changeVPort` or `changeVip`, this parameter needs to be configured.
     */
    port?: pulumi.Input<number>;
    /**
     * The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when chargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * Specifies which project the instance should belong to.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Original intranet IPv4 address retention time: unit: day, value range: `0`, `1`, `2`, `3`, `7`, `15`.
     */
    recycle?: pulumi.Input<number>;
    /**
     * The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replicaZoneIds`.
     */
    redisReplicasNum?: pulumi.Input<number>;
    /**
     * The number of instance shard, default is 1. This is not required for standalone and master slave versions.
     */
    redisShardNum?: pulumi.Input<number>;
    /**
     * ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
     */
    replicaZoneIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
     */
    replicasReadOnly?: pulumi.Input<boolean>;
    /**
     * ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies which subnet the instance should belong to. When the `operationNetwork` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Instance tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: `clusterCkv`,`cluster_redis5.0`,`clusterRedis`,`masterSlaveCkv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`masterSlaveRedis`,`standaloneRedis`, specific region support specific types, need to refer data `tencentcloud.Redis.getZoneConfig`.
     *
     * @deprecated It has been deprecated from version 1.33.1. Please use 'type_id' instead.
     */
    type?: pulumi.Input<string>;
    /**
     * Instance type. Available values reference data source `tencentcloud.Redis.getZoneConfig` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
     */
    typeId?: pulumi.Input<number>;
    /**
     * ID of the vpc with which the instance is to be associated. When the `operationNetwork` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
     */
    vpcId?: pulumi.Input<string>;
}
