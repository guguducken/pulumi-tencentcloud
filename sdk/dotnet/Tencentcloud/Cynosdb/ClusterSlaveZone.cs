// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    /// <summary>
    /// Provides a resource to create a cynosdb cluster slave zone.
    /// 
    /// ## Example Usage
    /// ### Set a new slave zone for a cynosdb cluster.
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
    ///         var gz3 = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             AvailabilityZone = @var.Default_az,
    ///             IsDefault = true,
    ///         }));
    ///         var vpcId = gz3.Apply(gz3 =&gt; gz3.InstanceLists?[0]?.VpcId);
    ///         var subnetId = gz3.Apply(gz3 =&gt; gz3.InstanceLists?[0]?.SubnetId);
    ///         var config = new Config();
    ///         var fixedTags = config.GetObject&lt;dynamic&gt;("fixedTags") ?? 
    ///         {
    ///             { "fixed_resource", "do_not_remove" },
    ///         };
    ///         var @internal = Output.Create(Tencentcloud.Security.GetGroups.InvokeAsync(new Tencentcloud.Security.GetGroupsArgs
    ///         {
    ///             Name = "default",
    ///             Tags = fixedTags,
    ///         }));
    ///         var sgId = @internal.Apply(@internal =&gt; @internal.SecurityGroups?[0]?.SecurityGroupId);
    ///         var exclusive = Output.Create(Tencentcloud.Security.GetGroups.InvokeAsync(new Tencentcloud.Security.GetGroupsArgs
    ///         {
    ///             Name = "test_preset_sg",
    ///         }));
    ///         var sgId2 = exclusive.Apply(exclusive =&gt; exclusive.SecurityGroups?[0]?.SecurityGroupId);
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-4";
    ///         var newAvailabilityZone = config.Get("newAvailabilityZone") ?? "ap-guangzhou-6";
    ///         var myParamTemplate = config.Get("myParamTemplate") ?? "15765";
    ///         var instance = new Tencentcloud.Cynosdb.Cluster("instance", new Tencentcloud.Cynosdb.ClusterArgs
    ///         {
    ///             AvailableZone = availabilityZone,
    ///             VpcId = vpcId,
    ///             SubnetId = subnetId,
    ///             DbType = "MYSQL",
    ///             DbVersion = "5.7",
    ///             StorageLimit = 1000,
    ///             ClusterName = "tf_test_cynosdb_cluster_slave_zone",
    ///             Password = "cynos@123",
    ///             InstanceMaintainDuration = 3600,
    ///             InstanceMaintainStartTime = 10800,
    ///             InstanceMaintainWeekdays = 
    ///             {
    ///                 "Fri",
    ///                 "Mon",
    ///                 "Sat",
    ///                 "Sun",
    ///                 "Thu",
    ///                 "Wed",
    ///                 "Tue",
    ///             },
    ///             InstanceCpuCore = 1,
    ///             InstanceMemorySize = 2,
    ///             ParamItems = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ClusterParamItemArgs
    ///                 {
    ///                     Name = "character_set_server",
    ///                     CurrentValue = "utf8",
    ///                 },
    ///                 new Tencentcloud.Cynosdb.Inputs.ClusterParamItemArgs
    ///                 {
    ///                     Name = "time_zone",
    ///                     CurrentValue = "+09:00",
    ///                 },
    ///             },
    ///             ForceDelete = true,
    ///             RwGroupSgs = 
    ///             {
    ///                 sgId,
    ///             },
    ///             RoGroupSgs = 
    ///             {
    ///                 sgId,
    ///             },
    ///             PrarmTemplateId = myParamTemplate,
    ///         });
    ///         var clusterSlaveZone = new Tencentcloud.Cynosdb.ClusterSlaveZone("clusterSlaveZone", new Tencentcloud.Cynosdb.ClusterSlaveZoneArgs
    ///         {
    ///             ClusterId = instance.Id,
    ///             SlaveZone = newAvailabilityZone,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Update the slave zone with specified value.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var clusterSlaveZone = new Tencentcloud.Cynosdb.ClusterSlaveZone("clusterSlaveZone", new Tencentcloud.Cynosdb.ClusterSlaveZoneArgs
    ///         {
    ///             ClusterId = tencentcloud_cynosdb_cluster.Instance.Id,
    ///             SlaveZone = @var.Availability_zone,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cynosdb cluster_slave_zone can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone cluster_slave_zone cluster_id#slave_zone
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone")]
    public partial class ClusterSlaveZone : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Slave zone.
        /// </summary>
        [Output("slaveZone")]
        public Output<string> SlaveZone { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterSlaveZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterSlaveZone(string name, ClusterSlaveZoneArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone", name, args ?? new ClusterSlaveZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterSlaveZone(string name, Input<string> id, ClusterSlaveZoneState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/clusterSlaveZone:ClusterSlaveZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterSlaveZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterSlaveZone Get(string name, Input<string> id, ClusterSlaveZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterSlaveZone(name, id, state, options);
        }
    }

    public sealed class ClusterSlaveZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Slave zone.
        /// </summary>
        [Input("slaveZone", required: true)]
        public Input<string> SlaveZone { get; set; } = null!;

        public ClusterSlaveZoneArgs()
        {
        }
    }

    public sealed class ClusterSlaveZoneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Slave zone.
        /// </summary>
        [Input("slaveZone")]
        public Input<string>? SlaveZone { get; set; }

        public ClusterSlaveZoneState()
        {
        }
    }
}
