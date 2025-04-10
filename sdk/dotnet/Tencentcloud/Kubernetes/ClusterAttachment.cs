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
    /// Provide a resource to attach an existing  cvm to kubernetes cluster.
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
    ///         var clusterCidr = config.Get("clusterCidr") ?? "172.16.0.0/16";
    ///         var defaultInstanceType = config.Get("defaultInstanceType") ?? "S1.SMALL1";
    ///         var defaultInstance = Output.Create(Tencentcloud.Images.GetInstance.InvokeAsync(new Tencentcloud.Images.GetInstanceArgs
    ///         {
    ///             ImageTypes = 
    ///             {
    ///                 "PUBLIC_IMAGE",
    ///             },
    ///             OsName = "centos",
    ///         }));
    ///         var vpc = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             IsDefault = true,
    ///             AvailabilityZone = availabilityZone,
    ///         }));
    ///         var defaultTypes = Output.Create(Tencentcloud.Instance.GetTypes.InvokeAsync(new Tencentcloud.Instance.GetTypesArgs
    ///         {
    ///             Filters = 
    ///             {
    ///                 new Tencentcloud.Instance.Inputs.GetTypesFilterArgs
    ///                 {
    ///                     Name = "instance-family",
    ///                     Values = 
    ///                     {
    ///                         "SA2",
    ///                     },
    ///                 },
    ///             },
    ///             CpuCoreCount = 8,
    ///             MemorySize = 16,
    ///         }));
    ///         var foo = new Tencentcloud.Instance.Instance("foo", new Tencentcloud.Instance.InstanceArgs
    ///         {
    ///             InstanceName = "tf-auto-test-1-1",
    ///             AvailabilityZone = availabilityZone,
    ///             ImageId = defaultInstance.Apply(defaultInstance =&gt; defaultInstance.Images?[0]?.ImageId),
    ///             InstanceType = defaultInstanceType,
    ///             SystemDiskType = "CLOUD_PREMIUM",
    ///             SystemDiskSize = 50,
    ///         });
    ///         var managedCluster = new Tencentcloud.Kubernetes.Cluster("managedCluster", new Tencentcloud.Kubernetes.ClusterArgs
    ///         {
    ///             VpcId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.VpcId),
    ///             ClusterCidr = "10.1.0.0/16",
    ///             ClusterMaxPodNum = 32,
    ///             ClusterName = "keep",
    ///             ClusterDesc = "test cluster desc",
    ///             ClusterMaxServiceNum = 32,
    ///             WorkerConfigs = 
    ///             {
    ///                 new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigArgs
    ///                 {
    ///                     Count = 1,
    ///                     AvailabilityZone = availabilityZone,
    ///                     InstanceType = defaultInstanceType,
    ///                     SystemDiskType = "CLOUD_SSD",
    ///                     SystemDiskSize = 60,
    ///                     InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                     InternetMaxBandwidthOut = 100,
    ///                     PublicIpAssigned = true,
    ///                     SubnetId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.SubnetId),
    ///                     DataDisks = 
    ///                     {
    ///                         new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigDataDiskArgs
    ///                         {
    ///                             DiskType = "CLOUD_PREMIUM",
    ///                             DiskSize = 50,
    ///                         },
    ///                     },
    ///                     EnhancedSecurityService = false,
    ///                     EnhancedMonitorService = false,
    ///                     UserData = "dGVzdA==",
    ///                     Password = "ZZXXccvv1212",
    ///                 },
    ///             },
    ///             ClusterDeployType = "MANAGED_CLUSTER",
    ///         });
    ///         var testAttach = new Tencentcloud.Kubernetes.ClusterAttachment("testAttach", new Tencentcloud.Kubernetes.ClusterAttachmentArgs
    ///         {
    ///             ClusterId = managedCluster.Id,
    ///             InstanceId = foo.Id,
    ///             Password = "Lo4wbdit",
    ///             Labels = 
    ///             {
    ///                 { "test1", "test1" },
    ///                 { "test2", "test2" },
    ///             },
    ///             WorkerConfigOverrides = new Tencentcloud.Kubernetes.Inputs.ClusterAttachmentWorkerConfigOverridesArgs
    ///             {
    ///                 DesiredPodNum = 8,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/clusterAttachment:ClusterAttachment")]
    public partial class ClusterAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        /// </summary>
        [Output("hostname")]
        public Output<string?> Hostname { get; private set; } = null!;

        /// <summary>
        /// ID of the CVM instance, this cvm will reinstall the system.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        /// </summary>
        [Output("keyIds")]
        public Output<string?> KeyIds { get; private set; } = null!;

        /// <summary>
        /// Labels of tke attachment exits CVM.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Password to access, should be set if `key_ids` not set.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs after attach to cluster.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// State of the node.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Output("unschedulable")]
        public Output<int?> Unschedulable { get; private set; } = null!;

        /// <summary>
        /// Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        /// </summary>
        [Output("workerConfig")]
        public Output<Outputs.ClusterAttachmentWorkerConfig?> WorkerConfig { get; private set; } = null!;

        /// <summary>
        /// Override variable worker_config, commonly used to attach existing instances.
        /// </summary>
        [Output("workerConfigOverrides")]
        public Output<Outputs.ClusterAttachmentWorkerConfigOverrides?> WorkerConfigOverrides { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterAttachment(string name, ClusterAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/clusterAttachment:ClusterAttachment", name, args ?? new ClusterAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterAttachment(string name, Input<string> id, ClusterAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/clusterAttachment:ClusterAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterAttachment Get(string name, Input<string> id, ClusterAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterAttachment(name, id, state, options);
        }
    }

    public sealed class ClusterAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// ID of the CVM instance, this cvm will reinstall the system.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        /// </summary>
        [Input("keyIds")]
        public Input<string>? KeyIds { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of tke attachment exits CVM.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Password to access, should be set if `key_ids` not set.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        /// <summary>
        /// Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.ClusterAttachmentWorkerConfigArgs>? WorkerConfig { get; set; }

        /// <summary>
        /// Override variable worker_config, commonly used to attach existing instances.
        /// </summary>
        [Input("workerConfigOverrides")]
        public Input<Inputs.ClusterAttachmentWorkerConfigOverridesArgs>? WorkerConfigOverrides { get; set; }

        public ClusterAttachmentArgs()
        {
        }
    }

    public sealed class ClusterAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// ID of the CVM instance, this cvm will reinstall the system.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        /// </summary>
        [Input("keyIds")]
        public Input<string>? KeyIds { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of tke attachment exits CVM.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Password to access, should be set if `key_ids` not set.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group IDs after attach to cluster.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// State of the node.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        /// <summary>
        /// Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.ClusterAttachmentWorkerConfigGetArgs>? WorkerConfig { get; set; }

        /// <summary>
        /// Override variable worker_config, commonly used to attach existing instances.
        /// </summary>
        [Input("workerConfigOverrides")]
        public Input<Inputs.ClusterAttachmentWorkerConfigOverridesGetArgs>? WorkerConfigOverrides { get; set; }

        public ClusterAttachmentState()
        {
        }
    }
}
