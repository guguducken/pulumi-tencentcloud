// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    /// <summary>
    /// Provides a resource to create a dcdb db_sync_mode_config
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
    ///         var config = new Tencentcloud.Dcdb.DbSyncModeConfig("config", new Tencentcloud.Dcdb.DbSyncModeConfigArgs
    ///         {
    ///             InstanceId = "%s",
    ///             SyncMode = 2,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dcdb db_sync_mode_config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dcdb/dbSyncModeConfig:DbSyncModeConfig db_sync_mode_config db_sync_mode_config_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dcdb/dbSyncModeConfig:DbSyncModeConfig")]
    public partial class DbSyncModeConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
        /// </summary>
        [Output("syncMode")]
        public Output<int> SyncMode { get; private set; } = null!;


        /// <summary>
        /// Create a DbSyncModeConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbSyncModeConfig(string name, DbSyncModeConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/dbSyncModeConfig:DbSyncModeConfig", name, args ?? new DbSyncModeConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbSyncModeConfig(string name, Input<string> id, DbSyncModeConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/dbSyncModeConfig:DbSyncModeConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DbSyncModeConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbSyncModeConfig Get(string name, Input<string> id, DbSyncModeConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DbSyncModeConfig(name, id, state, options);
        }
    }

    public sealed class DbSyncModeConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
        /// </summary>
        [Input("syncMode", required: true)]
        public Input<int> SyncMode { get; set; } = null!;

        public DbSyncModeConfigArgs()
        {
        }
    }

    public sealed class DbSyncModeConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
        /// </summary>
        [Input("syncMode")]
        public Input<int>? SyncMode { get; set; }

        public DbSyncModeConfigState()
        {
        }
    }
}
