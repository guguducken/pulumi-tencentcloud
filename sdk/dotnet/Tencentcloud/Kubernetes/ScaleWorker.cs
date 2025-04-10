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
    /// Provide a resource to increase instance to cluster
    /// 
    /// &gt; **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-3";
    ///         var subnet = config.Get("subnet") ?? "subnet-pqfek0t8";
    ///         var scaleInstanceType = config.Get("scaleInstanceType") ?? "S2.LARGE16";
    ///         var testScale = new Tencentcloud.Kubernetes.ScaleWorker("testScale", new Tencentcloud.Kubernetes.ScaleWorkerArgs
    ///         {
    ///             ClusterId = "cls-godovr32",
    ///             DesiredPodNum = 16,
    ///             Labels = 
    ///             {
    ///                 { "test1", "test1" },
    ///                 { "test2", "test2" },
    ///             },
    ///             WorkerConfig = new Tencentcloud.Kubernetes.Inputs.ScaleWorkerWorkerConfigArgs
    ///             {
    ///                 Count = 3,
    ///                 AvailabilityZone = availabilityZone,
    ///                 InstanceType = scaleInstanceType,
    ///                 SubnetId = subnet,
    ///                 SystemDiskType = "CLOUD_SSD",
    ///                 SystemDiskSize = 50,
    ///                 InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                 InternetMaxBandwidthOut = 100,
    ///                 PublicIpAssigned = true,
    ///                 DataDisks = 
    ///                 {
    ///                     new Tencentcloud.Kubernetes.Inputs.ScaleWorkerWorkerConfigDataDiskArgs
    ///                     {
    ///                         DiskType = "CLOUD_PREMIUM",
    ///                         DiskSize = 50,
    ///                     },
    ///                 },
    ///                 EnhancedSecurityService = false,
    ///                 EnhancedMonitorService = false,
    ///                 UserData = "dGVzdA==",
    ///                 Password = "AABBccdd1122",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Use Kubelet
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-3";
    ///         var subnet = config.Get("subnet") ?? "subnet-pqfek0t8";
    ///         var scaleInstanceType = config.Get("scaleInstanceType") ?? "S2.LARGE16";
    ///         var testScale = new Tencentcloud.Kubernetes.ScaleWorker("testScale", new Tencentcloud.Kubernetes.ScaleWorkerArgs
    ///         {
    ///             ClusterId = "cls-godovr32",
    ///             ExtraArgs = 
    ///             {
    ///                 "root-dir=/var/lib/kubelet",
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "test1", "test1" },
    ///                 { "test2", "test2" },
    ///             },
    ///             WorkerConfig = new Tencentcloud.Kubernetes.Inputs.ScaleWorkerWorkerConfigArgs
    ///             {
    ///                 Count = 3,
    ///                 AvailabilityZone = availabilityZone,
    ///                 InstanceType = scaleInstanceType,
    ///                 SubnetId = subnet,
    ///                 SystemDiskType = "CLOUD_SSD",
    ///                 SystemDiskSize = 50,
    ///                 InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                 InternetMaxBandwidthOut = 100,
    ///                 PublicIpAssigned = true,
    ///                 DataDisks = 
    ///                 {
    ///                     new Tencentcloud.Kubernetes.Inputs.ScaleWorkerWorkerConfigDataDiskArgs
    ///                     {
    ///                         DiskType = "CLOUD_PREMIUM",
    ///                         DiskSize = 50,
    ///                     },
    ///                 },
    ///                 EnhancedSecurityService = false,
    ///                 EnhancedMonitorService = false,
    ///                 UserData = "dGVzdA==",
    ///                 Password = "AABBccdd1122",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/scaleWorker:ScaleWorker")]
    public partial class ScaleWorker : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        [Output("dataDisks")]
        public Output<ImmutableArray<Outputs.ScaleWorkerDataDisk>> DataDisks { get; private set; } = null!;

        /// <summary>
        /// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
        /// </summary>
        [Output("desiredPodNum")]
        public Output<int?> DesiredPodNum { get; private set; } = null!;

        /// <summary>
        /// Docker graph path. Default is `/var/lib/docker`.
        /// </summary>
        [Output("dockerGraphPath")]
        public Output<string?> DockerGraphPath { get; private set; } = null!;

        /// <summary>
        /// Custom parameter information related to the node.
        /// </summary>
        [Output("extraArgs")]
        public Output<ImmutableArray<string>> ExtraArgs { get; private set; } = null!;

        /// <summary>
        /// GPU driver parameters.
        /// </summary>
        [Output("gpuArgs")]
        public Output<Outputs.ScaleWorkerGpuArgs?> GpuArgs { get; private set; } = null!;

        /// <summary>
        /// Labels of kubernetes scale worker created nodes.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Mount target. Default is not mounting.
        /// </summary>
        [Output("mountTarget")]
        public Output<string?> MountTarget { get; private set; } = null!;

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Output("unschedulable")]
        public Output<int?> Unschedulable { get; private set; } = null!;

        /// <summary>
        /// Deploy the machine configuration information of the 'WORK' service, and create &lt;=20 units for common users.
        /// </summary>
        [Output("workerConfig")]
        public Output<Outputs.ScaleWorkerWorkerConfig> WorkerConfig { get; private set; } = null!;

        /// <summary>
        /// An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
        /// </summary>
        [Output("workerInstancesLists")]
        public Output<ImmutableArray<Outputs.ScaleWorkerWorkerInstancesList>> WorkerInstancesLists { get; private set; } = null!;


        /// <summary>
        /// Create a ScaleWorker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScaleWorker(string name, ScaleWorkerArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/scaleWorker:ScaleWorker", name, args ?? new ScaleWorkerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScaleWorker(string name, Input<string> id, ScaleWorkerState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/scaleWorker:ScaleWorker", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScaleWorker resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScaleWorker Get(string name, Input<string> id, ScaleWorkerState? state = null, CustomResourceOptions? options = null)
        {
            return new ScaleWorker(name, id, state, options);
        }
    }

    public sealed class ScaleWorkerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("dataDisks")]
        private InputList<Inputs.ScaleWorkerDataDiskArgs>? _dataDisks;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public InputList<Inputs.ScaleWorkerDataDiskArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.ScaleWorkerDataDiskArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
        /// </summary>
        [Input("desiredPodNum")]
        public Input<int>? DesiredPodNum { get; set; }

        /// <summary>
        /// Docker graph path. Default is `/var/lib/docker`.
        /// </summary>
        [Input("dockerGraphPath")]
        public Input<string>? DockerGraphPath { get; set; }

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;

        /// <summary>
        /// Custom parameter information related to the node.
        /// </summary>
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        /// <summary>
        /// GPU driver parameters.
        /// </summary>
        [Input("gpuArgs")]
        public Input<Inputs.ScaleWorkerGpuArgsArgs>? GpuArgs { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of kubernetes scale worker created nodes.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Mount target. Default is not mounting.
        /// </summary>
        [Input("mountTarget")]
        public Input<string>? MountTarget { get; set; }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        /// <summary>
        /// Deploy the machine configuration information of the 'WORK' service, and create &lt;=20 units for common users.
        /// </summary>
        [Input("workerConfig", required: true)]
        public Input<Inputs.ScaleWorkerWorkerConfigArgs> WorkerConfig { get; set; } = null!;

        public ScaleWorkerArgs()
        {
        }
    }

    public sealed class ScaleWorkerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.ScaleWorkerDataDiskGetArgs>? _dataDisks;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public InputList<Inputs.ScaleWorkerDataDiskGetArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.ScaleWorkerDataDiskGetArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
        /// </summary>
        [Input("desiredPodNum")]
        public Input<int>? DesiredPodNum { get; set; }

        /// <summary>
        /// Docker graph path. Default is `/var/lib/docker`.
        /// </summary>
        [Input("dockerGraphPath")]
        public Input<string>? DockerGraphPath { get; set; }

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;

        /// <summary>
        /// Custom parameter information related to the node.
        /// </summary>
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        /// <summary>
        /// GPU driver parameters.
        /// </summary>
        [Input("gpuArgs")]
        public Input<Inputs.ScaleWorkerGpuArgsGetArgs>? GpuArgs { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of kubernetes scale worker created nodes.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Mount target. Default is not mounting.
        /// </summary>
        [Input("mountTarget")]
        public Input<string>? MountTarget { get; set; }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        /// <summary>
        /// Deploy the machine configuration information of the 'WORK' service, and create &lt;=20 units for common users.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.ScaleWorkerWorkerConfigGetArgs>? WorkerConfig { get; set; }

        [Input("workerInstancesLists")]
        private InputList<Inputs.ScaleWorkerWorkerInstancesListGetArgs>? _workerInstancesLists;

        /// <summary>
        /// An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.ScaleWorkerWorkerInstancesListGetArgs> WorkerInstancesLists
        {
            get => _workerInstancesLists ?? (_workerInstancesLists = new InputList<Inputs.ScaleWorkerWorkerInstancesListGetArgs>());
            set => _workerInstancesLists = value;
        }

        public ScaleWorkerState()
        {
        }
    }
}
