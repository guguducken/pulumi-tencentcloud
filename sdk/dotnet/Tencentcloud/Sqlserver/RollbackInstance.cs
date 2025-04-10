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
    /// Provides a resource to create a sqlserver rollback_instance
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
    ///         var rollbackInstance = new Tencentcloud.Sqlserver.RollbackInstance("rollbackInstance", new Tencentcloud.Sqlserver.RollbackInstanceArgs
    ///         {
    ///             InstanceId = "mssql-qelbzgwf",
    ///             RenameRestores = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.RollbackInstanceRenameRestoreArgs
    ///                 {
    ///                     NewName = "rollback_pubsub_db3",
    ///                     OldName = "keep_pubsub_db2",
    ///                 },
    ///             },
    ///             Time = "2023-05-23 01:00:00",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver rollback_instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/rollbackInstance:RollbackInstance rollback_instance rollback_instance_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/rollbackInstance:RollbackInstance")]
    public partial class RollbackInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// TDE encryption, `enable` encrypted, `disable` unencrypted.
        /// </summary>
        [Output("encryptions")]
        public Output<ImmutableArray<Outputs.RollbackInstanceEncryption>> Encryptions { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Rename the databases listed in ReNameRestoreDatabase.
        /// </summary>
        [Output("renameRestores")]
        public Output<ImmutableArray<Outputs.RollbackInstanceRenameRestore>> RenameRestores { get; private set; } = null!;

        /// <summary>
        /// Target time point for rollback.
        /// </summary>
        [Output("time")]
        public Output<string> Time { get; private set; } = null!;


        /// <summary>
        /// Create a RollbackInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RollbackInstance(string name, RollbackInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/rollbackInstance:RollbackInstance", name, args ?? new RollbackInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RollbackInstance(string name, Input<string> id, RollbackInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/rollbackInstance:RollbackInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RollbackInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RollbackInstance Get(string name, Input<string> id, RollbackInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new RollbackInstance(name, id, state, options);
        }
    }

    public sealed class RollbackInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("renameRestores", required: true)]
        private InputList<Inputs.RollbackInstanceRenameRestoreArgs>? _renameRestores;

        /// <summary>
        /// Rename the databases listed in ReNameRestoreDatabase.
        /// </summary>
        public InputList<Inputs.RollbackInstanceRenameRestoreArgs> RenameRestores
        {
            get => _renameRestores ?? (_renameRestores = new InputList<Inputs.RollbackInstanceRenameRestoreArgs>());
            set => _renameRestores = value;
        }

        /// <summary>
        /// Target time point for rollback.
        /// </summary>
        [Input("time", required: true)]
        public Input<string> Time { get; set; } = null!;

        public RollbackInstanceArgs()
        {
        }
    }

    public sealed class RollbackInstanceState : Pulumi.ResourceArgs
    {
        [Input("encryptions")]
        private InputList<Inputs.RollbackInstanceEncryptionGetArgs>? _encryptions;

        /// <summary>
        /// TDE encryption, `enable` encrypted, `disable` unencrypted.
        /// </summary>
        public InputList<Inputs.RollbackInstanceEncryptionGetArgs> Encryptions
        {
            get => _encryptions ?? (_encryptions = new InputList<Inputs.RollbackInstanceEncryptionGetArgs>());
            set => _encryptions = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("renameRestores")]
        private InputList<Inputs.RollbackInstanceRenameRestoreGetArgs>? _renameRestores;

        /// <summary>
        /// Rename the databases listed in ReNameRestoreDatabase.
        /// </summary>
        public InputList<Inputs.RollbackInstanceRenameRestoreGetArgs> RenameRestores
        {
            get => _renameRestores ?? (_renameRestores = new InputList<Inputs.RollbackInstanceRenameRestoreGetArgs>());
            set => _renameRestores = value;
        }

        /// <summary>
        /// Target time point for rollback.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        public RollbackInstanceState()
        {
        }
    }
}
