// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes
{
    /// <summary>
    /// Provide a resource to create an auto scaling group for kubernetes cluster.
    /// 
    /// &gt; **NOTE:**  We recommend the usage of one cluster with essential worker config + node pool to manage cluster and nodes. Its a more flexible way than manage worker config with tencentcloud_kubernetes_cluster, tencentcloud.Kubernetes.ScaleWorker or exist node management of `tencentcloud_kubernetes_attachment`. Cause some unchangeable parameters of `worker_config` may cause the whole cluster resource `force new`.
    /// 
    /// &gt; **NOTE:**  In order to ensure the integrity of customer data, if you destroy nodepool instance, it will keep the cvm instance associate with nodepool by default. If you want to destroy together, please set `delete_keep_instance` to `false`.
    /// 
    /// &gt; **NOTE:**  In order to ensure the integrity of customer data, if the cvm instance was destroyed due to shrinking, it will keep the cbs associate with cvm by default. If you want to destroy together, please set `delete_with_instance` to `true`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-3";
    ///         var clusterCidr = config.Get("clusterCidr") ?? "172.31.0.0/16";
    ///         var vpc = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             IsDefault = true,
    ///             AvailabilityZone = availabilityZone,
    ///         }));
    ///         var defaultInstanceType = config.Get("defaultInstanceType") ?? "S1.SMALL1";
    ///         //this is the cluster with empty worker config
    ///         var managedCluster = new Tencentcloud.Kubernetes.Cluster("managedCluster", new Tencentcloud.Kubernetes.ClusterArgs
    ///         {
    ///             VpcId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.VpcId),
    ///             ClusterCidr = clusterCidr,
    ///             ClusterMaxPodNum = 32,
    ///             ClusterName = "tf-tke-unit-test",
    ///             ClusterDesc = "test cluster desc",
    ///             ClusterMaxServiceNum = 32,
    ///             ClusterVersion = "1.18.4",
    ///             ClusterDeployType = "MANAGED_CLUSTER",
    ///         });
    ///         //this is one example of managing node using node pool
    ///         var mynodepool = new Tencentcloud.Kubernetes.NodePool("mynodepool", new Tencentcloud.Kubernetes.NodePoolArgs
    ///         {
    ///             ClusterId = managedCluster.Id,
    ///             MaxSize = 6,
    ///             MinSize = 1,
    ///             VpcId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.VpcId),
    ///             SubnetIds = 
    ///             {
    ///                 vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.SubnetId),
    ///             },
    ///             RetryPolicy = "INCREMENTAL_INTERVALS",
    ///             DesiredCapacity = 4,
    ///             EnableAutoScale = true,
    ///             MultiZoneSubnetPolicy = "EQUALITY",
    ///             AutoScalingConfig = new Tencentcloud.Kubernetes.Inputs.NodePoolAutoScalingConfigArgs
    ///             {
    ///                 InstanceType = defaultInstanceType,
    ///                 SystemDiskType = "CLOUD_PREMIUM",
    ///                 SystemDiskSize = 50,
    ///                 SecurityGroupIds = 
    ///                 {
    ///                     "sg-24vswocp",
    ///                 },
    ///                 DataDisks = 
    ///                 {
    ///                     new Tencentcloud.Kubernetes.Inputs.NodePoolAutoScalingConfigDataDiskArgs
    ///                     {
    ///                         DiskType = "CLOUD_PREMIUM",
    ///                         DiskSize = 50,
    ///                     },
    ///                 },
    ///                 InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                 InternetMaxBandwidthOut = 10,
    ///                 PublicIpAssigned = true,
    ///                 Password = "test123#",
    ///                 EnhancedSecurityService = false,
    ///                 EnhancedMonitorService = false,
    ///                 HostName = "12.123.0.0",
    ///                 HostNameStyle = "ORIGINAL",
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "test1", "test1" },
    ///                 { "test2", "test2" },
    ///             },
    ///             Taints = 
    ///             {
    ///                 new Tencentcloud.Kubernetes.Inputs.NodePoolTaintArgs
    ///                 {
    ///                     Key = "test_taint",
    ///                     Value = "taint_value",
    ///                     Effect = "PreferNoSchedule",
    ///                 },
    ///                 new Tencentcloud.Kubernetes.Inputs.NodePoolTaintArgs
    ///                 {
    ///                     Key = "test_taint2",
    ///                     Value = "taint_value2",
    ///                     Effect = "PreferNoSchedule",
    ///                 },
    ///             },
    ///             NodeConfig = new Tencentcloud.Kubernetes.Inputs.NodePoolNodeConfigArgs
    ///             {
    ///                 ExtraArgs = 
    ///                 {
    ///                     "root-dir=/var/lib/kubelet",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Using Spot CVM Instance
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mynodepool = new Tencentcloud.Kubernetes.NodePool("mynodepool", new Tencentcloud.Kubernetes.NodePoolArgs
    ///         {
    ///             ClusterId = tencentcloud_kubernetes_cluster.Managed_cluster.Id,
    ///             MaxSize = 6,
    ///             MinSize = 1,
    ///             VpcId = data.Tencentcloud_vpc_subnets.Vpc.Instance_list[0].Vpc_id,
    ///             SubnetIds = 
    ///             {
    ///                 data.Tencentcloud_vpc_subnets.Vpc.Instance_list[0].Subnet_id,
    ///             },
    ///             RetryPolicy = "INCREMENTAL_INTERVALS",
    ///             DesiredCapacity = 4,
    ///             EnableAutoScale = true,
    ///             MultiZoneSubnetPolicy = "EQUALITY",
    ///             AutoScalingConfig = new Tencentcloud.Kubernetes.Inputs.NodePoolAutoScalingConfigArgs
    ///             {
    ///                 InstanceType = @var.Default_instance_type,
    ///                 SystemDiskType = "CLOUD_PREMIUM",
    ///                 SystemDiskSize = 50,
    ///                 SecurityGroupIds = 
    ///                 {
    ///                     "sg-24vswocp",
    ///                 },
    ///                 InstanceChargeType = "SPOTPAID",
    ///                 SpotInstanceType = "one-time",
    ///                 SpotMaxPrice = "1000",
    ///                 DataDisks = 
    ///                 {
    ///                     new Tencentcloud.Kubernetes.Inputs.NodePoolAutoScalingConfigDataDiskArgs
    ///                     {
    ///                         DiskType = "CLOUD_PREMIUM",
    ///                         DiskSize = 50,
    ///                     },
    ///                 },
    ///                 InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                 InternetMaxBandwidthOut = 10,
    ///                 PublicIpAssigned = true,
    ///                 Password = "test123#",
    ///                 EnhancedSecurityService = false,
    ///                 EnhancedMonitorService = false,
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "test1", "test1" },
    ///                 { "test2", "test2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/nodePool:NodePool")]
    public partial class NodePool : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto scaling config parameters.
        /// </summary>
        [Output("autoScalingConfig")]
        public Output<Outputs.NodePoolAutoScalingConfig> AutoScalingConfig { get; private set; } = null!;

        /// <summary>
        /// The auto scaling group ID.
        /// </summary>
        [Output("autoScalingGroupId")]
        public Output<string> AutoScalingGroupId { get; private set; } = null!;

        /// <summary>
        /// The total of autoscaling added node.
        /// </summary>
        [Output("autoscalingAddedTotal")]
        public Output<int> AutoscalingAddedTotal { get; private set; } = null!;

        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Seconds of scaling group cool down. Default value is `300`.
        /// </summary>
        [Output("defaultCooldown")]
        public Output<int> DefaultCooldown { get; private set; } = null!;

        /// <summary>
        /// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
        /// </summary>
        [Output("deleteKeepInstance")]
        public Output<bool?> DeleteKeepInstance { get; private set; } = null!;

        /// <summary>
        /// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
        /// </summary>
        [Output("desiredCapacity")]
        public Output<int> DesiredCapacity { get; private set; } = null!;

        /// <summary>
        /// Indicate whether to enable auto scaling or not.
        /// </summary>
        [Output("enableAutoScale")]
        public Output<bool?> EnableAutoScale { get; private set; } = null!;

        /// <summary>
        /// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The launch config ID.
        /// </summary>
        [Output("launchConfigId")]
        public Output<string> LaunchConfigId { get; private set; } = null!;

        /// <summary>
        /// The total of manually added node.
        /// </summary>
        [Output("manuallyAddedTotal")]
        public Output<int> ManuallyAddedTotal { get; private set; } = null!;

        /// <summary>
        /// Maximum number of node.
        /// </summary>
        [Output("maxSize")]
        public Output<int> MaxSize { get; private set; } = null!;

        /// <summary>
        /// Minimum number of node.
        /// </summary>
        [Output("minSize")]
        public Output<int> MinSize { get; private set; } = null!;

        /// <summary>
        /// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
        /// </summary>
        [Output("multiZoneSubnetPolicy")]
        public Output<string?> MultiZoneSubnetPolicy { get; private set; } = null!;

        /// <summary>
        /// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (`-`) and decimal points.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Node config.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.NodePoolNodeConfig?> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// The total node count.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        /// <summary>
        /// Operating system of the cluster. Please refer to [TencentCloud Documentation](https://www.tencentcloud.com/document/product/457/46750?lang=en&amp;pg=#list-of-public-images-supported-by-tke) for available values. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.
        /// </summary>
        [Output("nodeOs")]
        public Output<string?> NodeOs { get; private set; } = null!;

        /// <summary>
        /// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter will only affect new nodes, not including the existing nodes.
        /// </summary>
        [Output("nodeOsType")]
        public Output<string?> NodeOsType { get; private set; } = null!;

        /// <summary>
        /// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
        /// </summary>
        [Output("retryPolicy")]
        public Output<string?> RetryPolicy { get; private set; } = null!;

        /// <summary>
        /// Name of relative scaling group.
        /// </summary>
        [Output("scalingGroupName")]
        public Output<string> ScalingGroupName { get; private set; } = null!;

        /// <summary>
        /// Project ID the scaling group belongs to.
        /// </summary>
        [Output("scalingGroupProjectId")]
        public Output<int?> ScalingGroupProjectId { get; private set; } = null!;

        /// <summary>
        /// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances), `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).
        /// </summary>
        [Output("scalingMode")]
        public Output<string?> ScalingMode { get; private set; } = null!;

        /// <summary>
        /// Status of the node pool.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// ID list of subnet, and for VPC it is required.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// Node pool tag specifications, will passthroughs to the scaling instances.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Taints of kubernetes node pool created nodes.
        /// </summary>
        [Output("taints")]
        public Output<ImmutableArray<Outputs.NodePoolTaint>> Taints { get; private set; } = null!;

        /// <summary>
        /// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
        /// </summary>
        [Output("terminationPolicies")]
        public Output<string> TerminationPolicies { get; private set; } = null!;

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Output("unschedulable")]
        public Output<int?> Unschedulable { get; private set; } = null!;

        /// <summary>
        /// ID of VPC network.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// List of auto scaling group available zones, for Basic network it is required.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a NodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodePool(string name, NodePoolArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/nodePool:NodePool", name, args ?? new NodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodePool(string name, Input<string> id, NodePoolState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/nodePool:NodePool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodePool Get(string name, Input<string> id, NodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new NodePool(name, id, state, options);
        }
    }

    public sealed class NodePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto scaling config parameters.
        /// </summary>
        [Input("autoScalingConfig", required: true)]
        public Input<Inputs.NodePoolAutoScalingConfigArgs> AutoScalingConfig { get; set; } = null!;

        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Seconds of scaling group cool down. Default value is `300`.
        /// </summary>
        [Input("defaultCooldown")]
        public Input<int>? DefaultCooldown { get; set; }

        /// <summary>
        /// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
        /// </summary>
        [Input("deleteKeepInstance")]
        public Input<bool>? DeleteKeepInstance { get; set; }

        /// <summary>
        /// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
        /// </summary>
        [Input("desiredCapacity")]
        public Input<int>? DesiredCapacity { get; set; }

        /// <summary>
        /// Indicate whether to enable auto scaling or not.
        /// </summary>
        [Input("enableAutoScale")]
        public Input<bool>? EnableAutoScale { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Maximum number of node.
        /// </summary>
        [Input("maxSize", required: true)]
        public Input<int> MaxSize { get; set; } = null!;

        /// <summary>
        /// Minimum number of node.
        /// </summary>
        [Input("minSize", required: true)]
        public Input<int> MinSize { get; set; } = null!;

        /// <summary>
        /// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
        /// </summary>
        [Input("multiZoneSubnetPolicy")]
        public Input<string>? MultiZoneSubnetPolicy { get; set; }

        /// <summary>
        /// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (`-`) and decimal points.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Node config.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.NodePoolNodeConfigArgs>? NodeConfig { get; set; }

        /// <summary>
        /// Operating system of the cluster. Please refer to [TencentCloud Documentation](https://www.tencentcloud.com/document/product/457/46750?lang=en&amp;pg=#list-of-public-images-supported-by-tke) for available values. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.
        /// </summary>
        [Input("nodeOs")]
        public Input<string>? NodeOs { get; set; }

        /// <summary>
        /// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter will only affect new nodes, not including the existing nodes.
        /// </summary>
        [Input("nodeOsType")]
        public Input<string>? NodeOsType { get; set; }

        /// <summary>
        /// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
        /// </summary>
        [Input("retryPolicy")]
        public Input<string>? RetryPolicy { get; set; }

        /// <summary>
        /// Name of relative scaling group.
        /// </summary>
        [Input("scalingGroupName")]
        public Input<string>? ScalingGroupName { get; set; }

        /// <summary>
        /// Project ID the scaling group belongs to.
        /// </summary>
        [Input("scalingGroupProjectId")]
        public Input<int>? ScalingGroupProjectId { get; set; }

        /// <summary>
        /// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances), `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).
        /// </summary>
        [Input("scalingMode")]
        public Input<string>? ScalingMode { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// ID list of subnet, and for VPC it is required.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Node pool tag specifications, will passthroughs to the scaling instances.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("taints")]
        private InputList<Inputs.NodePoolTaintArgs>? _taints;

        /// <summary>
        /// Taints of kubernetes node pool created nodes.
        /// </summary>
        public InputList<Inputs.NodePoolTaintArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.NodePoolTaintArgs>());
            set => _taints = value;
        }

        /// <summary>
        /// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
        /// </summary>
        [Input("terminationPolicies")]
        public Input<string>? TerminationPolicies { get; set; }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        /// <summary>
        /// ID of VPC network.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// List of auto scaling group available zones, for Basic network it is required.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public NodePoolArgs()
        {
        }
    }

    public sealed class NodePoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto scaling config parameters.
        /// </summary>
        [Input("autoScalingConfig")]
        public Input<Inputs.NodePoolAutoScalingConfigGetArgs>? AutoScalingConfig { get; set; }

        /// <summary>
        /// The auto scaling group ID.
        /// </summary>
        [Input("autoScalingGroupId")]
        public Input<string>? AutoScalingGroupId { get; set; }

        /// <summary>
        /// The total of autoscaling added node.
        /// </summary>
        [Input("autoscalingAddedTotal")]
        public Input<int>? AutoscalingAddedTotal { get; set; }

        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Seconds of scaling group cool down. Default value is `300`.
        /// </summary>
        [Input("defaultCooldown")]
        public Input<int>? DefaultCooldown { get; set; }

        /// <summary>
        /// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
        /// </summary>
        [Input("deleteKeepInstance")]
        public Input<bool>? DeleteKeepInstance { get; set; }

        /// <summary>
        /// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
        /// </summary>
        [Input("desiredCapacity")]
        public Input<int>? DesiredCapacity { get; set; }

        /// <summary>
        /// Indicate whether to enable auto scaling or not.
        /// </summary>
        [Input("enableAutoScale")]
        public Input<bool>? EnableAutoScale { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The launch config ID.
        /// </summary>
        [Input("launchConfigId")]
        public Input<string>? LaunchConfigId { get; set; }

        /// <summary>
        /// The total of manually added node.
        /// </summary>
        [Input("manuallyAddedTotal")]
        public Input<int>? ManuallyAddedTotal { get; set; }

        /// <summary>
        /// Maximum number of node.
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// Minimum number of node.
        /// </summary>
        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        /// <summary>
        /// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
        /// </summary>
        [Input("multiZoneSubnetPolicy")]
        public Input<string>? MultiZoneSubnetPolicy { get; set; }

        /// <summary>
        /// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (`-`) and decimal points.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Node config.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.NodePoolNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The total node count.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// Operating system of the cluster. Please refer to [TencentCloud Documentation](https://www.tencentcloud.com/document/product/457/46750?lang=en&amp;pg=#list-of-public-images-supported-by-tke) for available values. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.
        /// </summary>
        [Input("nodeOs")]
        public Input<string>? NodeOs { get; set; }

        /// <summary>
        /// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter will only affect new nodes, not including the existing nodes.
        /// </summary>
        [Input("nodeOsType")]
        public Input<string>? NodeOsType { get; set; }

        /// <summary>
        /// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
        /// </summary>
        [Input("retryPolicy")]
        public Input<string>? RetryPolicy { get; set; }

        /// <summary>
        /// Name of relative scaling group.
        /// </summary>
        [Input("scalingGroupName")]
        public Input<string>? ScalingGroupName { get; set; }

        /// <summary>
        /// Project ID the scaling group belongs to.
        /// </summary>
        [Input("scalingGroupProjectId")]
        public Input<int>? ScalingGroupProjectId { get; set; }

        /// <summary>
        /// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances), `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).
        /// </summary>
        [Input("scalingMode")]
        public Input<string>? ScalingMode { get; set; }

        /// <summary>
        /// Status of the node pool.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// ID list of subnet, and for VPC it is required.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Node pool tag specifications, will passthroughs to the scaling instances.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("taints")]
        private InputList<Inputs.NodePoolTaintGetArgs>? _taints;

        /// <summary>
        /// Taints of kubernetes node pool created nodes.
        /// </summary>
        public InputList<Inputs.NodePoolTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.NodePoolTaintGetArgs>());
            set => _taints = value;
        }

        /// <summary>
        /// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
        /// </summary>
        [Input("terminationPolicies")]
        public Input<string>? TerminationPolicies { get; set; }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        /// <summary>
        /// ID of VPC network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// List of auto scaling group available zones, for Basic network it is required.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public NodePoolState()
        {
        }
    }
}
