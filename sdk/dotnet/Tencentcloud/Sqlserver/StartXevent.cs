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
    /// Provides a resource to create a sqlserver start_xevent
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
    ///         var startXevent = new Tencentcloud.Sqlserver.StartXevent("startXevent", new Tencentcloud.Sqlserver.StartXeventArgs
    ///         {
    ///             EventConfigs = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.StartXeventEventConfigArgs
    ///                 {
    ///                     EventType = "slow",
    ///                     Threshold = 0,
    ///                 },
    ///             },
    ///             InstanceId = "mssql-gyg9xycl",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/startXevent:StartXevent")]
    public partial class StartXevent : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to start or stop an extended event.
        /// </summary>
        [Output("eventConfigs")]
        public Output<ImmutableArray<Outputs.StartXeventEventConfig>> EventConfigs { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a StartXevent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StartXevent(string name, StartXeventArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/startXevent:StartXevent", name, args ?? new StartXeventArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StartXevent(string name, Input<string> id, StartXeventState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/startXevent:StartXevent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StartXevent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StartXevent Get(string name, Input<string> id, StartXeventState? state = null, CustomResourceOptions? options = null)
        {
            return new StartXevent(name, id, state, options);
        }
    }

    public sealed class StartXeventArgs : Pulumi.ResourceArgs
    {
        [Input("eventConfigs", required: true)]
        private InputList<Inputs.StartXeventEventConfigArgs>? _eventConfigs;

        /// <summary>
        /// Whether to start or stop an extended event.
        /// </summary>
        public InputList<Inputs.StartXeventEventConfigArgs> EventConfigs
        {
            get => _eventConfigs ?? (_eventConfigs = new InputList<Inputs.StartXeventEventConfigArgs>());
            set => _eventConfigs = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public StartXeventArgs()
        {
        }
    }

    public sealed class StartXeventState : Pulumi.ResourceArgs
    {
        [Input("eventConfigs")]
        private InputList<Inputs.StartXeventEventConfigGetArgs>? _eventConfigs;

        /// <summary>
        /// Whether to start or stop an extended event.
        /// </summary>
        public InputList<Inputs.StartXeventEventConfigGetArgs> EventConfigs
        {
            get => _eventConfigs ?? (_eventConfigs = new InputList<Inputs.StartXeventEventConfigGetArgs>());
            set => _eventConfigs = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public StartXeventState()
        {
        }
    }
}
