// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    /// <summary>
    /// Provides a resource to create a sqlserver config_database_mdf
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
    ///         var configDatabaseMdf = new Tencentcloud.Sqlserver.ConfigDatabaseMdf("configDatabaseMdf", new Tencentcloud.Sqlserver.ConfigDatabaseMdfArgs
    ///         {
    ///             DbName = "keep_pubsub_db2",
    ///             InstanceId = "mssql-qelbzgwf",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver config_database_mdf can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf config_database_mdf config_database_mdf_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf")]
    public partial class ConfigDatabaseMdf : Pulumi.CustomResource
    {
        /// <summary>
        /// Array of database names.
        /// </summary>
        [Output("dbName")]
        public Output<string> DbName { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigDatabaseMdf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigDatabaseMdf(string name, ConfigDatabaseMdfArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf", name, args ?? new ConfigDatabaseMdfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigDatabaseMdf(string name, Input<string> id, ConfigDatabaseMdfState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigDatabaseMdf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigDatabaseMdf Get(string name, Input<string> id, ConfigDatabaseMdfState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigDatabaseMdf(name, id, state, options);
        }
    }

    public sealed class ConfigDatabaseMdfArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Array of database names.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public ConfigDatabaseMdfArgs()
        {
        }
    }

    public sealed class ConfigDatabaseMdfState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Array of database names.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public ConfigDatabaseMdfState()
        {
        }
    }
}
