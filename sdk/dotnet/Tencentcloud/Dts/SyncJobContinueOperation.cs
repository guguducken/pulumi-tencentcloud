// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts
{
    /// <summary>
    /// Provides a resource to create a dts sync_job_continue_operation
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
    ///         var syncJobContinueOperation = new Tencentcloud.Dts.SyncJobContinueOperation("syncJobContinueOperation", new Tencentcloud.Dts.SyncJobContinueOperationArgs
    ///         {
    ///             JobId = "sync-werwfs23",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dts/syncJobContinueOperation:SyncJobContinueOperation")]
    public partial class SyncJobContinueOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;


        /// <summary>
        /// Create a SyncJobContinueOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncJobContinueOperation(string name, SyncJobContinueOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobContinueOperation:SyncJobContinueOperation", name, args ?? new SyncJobContinueOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncJobContinueOperation(string name, Input<string> id, SyncJobContinueOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobContinueOperation:SyncJobContinueOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyncJobContinueOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncJobContinueOperation Get(string name, Input<string> id, SyncJobContinueOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncJobContinueOperation(name, id, state, options);
        }
    }

    public sealed class SyncJobContinueOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        public SyncJobContinueOperationArgs()
        {
        }
    }

    public sealed class SyncJobContinueOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        public SyncJobContinueOperationState()
        {
        }
    }
}
