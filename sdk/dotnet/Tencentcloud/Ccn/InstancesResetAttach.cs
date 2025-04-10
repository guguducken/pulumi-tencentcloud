// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ccn
{
    /// <summary>
    /// Provides a resource to create a vpc ccn_instances_reset_attach, you can use this resource to reset cross-region attachment.
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
    ///         var ccnInstancesResetAttach = new Tencentcloud.Ccn.InstancesResetAttach("ccnInstancesResetAttach", new Tencentcloud.Ccn.InstancesResetAttachArgs
    ///         {
    ///             CcnId = "ccn-39lqkygf",
    ///             CcnUin = "100022975249",
    ///             Instances = 
    ///             {
    ///                 new Tencentcloud.Ccn.Inputs.InstancesResetAttachInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-j9yhbzpn",
    ///                     InstanceRegion = "ap-guangzhou",
    ///                     InstanceType = "VPC",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ccn/instancesResetAttach:InstancesResetAttach")]
    public partial class InstancesResetAttach : Pulumi.CustomResource
    {
        /// <summary>
        /// CCN Instance ID.
        /// </summary>
        [Output("ccnId")]
        public Output<string> CcnId { get; private set; } = null!;

        /// <summary>
        /// CCN Uin (root account).
        /// </summary>
        [Output("ccnUin")]
        public Output<string> CcnUin { get; private set; } = null!;

        /// <summary>
        /// List Of Attachment Instances.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.InstancesResetAttachInstance>> Instances { get; private set; } = null!;


        /// <summary>
        /// Create a InstancesResetAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstancesResetAttach(string name, InstancesResetAttachArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/instancesResetAttach:InstancesResetAttach", name, args ?? new InstancesResetAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstancesResetAttach(string name, Input<string> id, InstancesResetAttachState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/instancesResetAttach:InstancesResetAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstancesResetAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstancesResetAttach Get(string name, Input<string> id, InstancesResetAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new InstancesResetAttach(name, id, state, options);
        }
    }

    public sealed class InstancesResetAttachArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CCN Instance ID.
        /// </summary>
        [Input("ccnId", required: true)]
        public Input<string> CcnId { get; set; } = null!;

        /// <summary>
        /// CCN Uin (root account).
        /// </summary>
        [Input("ccnUin", required: true)]
        public Input<string> CcnUin { get; set; } = null!;

        [Input("instances", required: true)]
        private InputList<Inputs.InstancesResetAttachInstanceArgs>? _instances;

        /// <summary>
        /// List Of Attachment Instances.
        /// </summary>
        public InputList<Inputs.InstancesResetAttachInstanceArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.InstancesResetAttachInstanceArgs>());
            set => _instances = value;
        }

        public InstancesResetAttachArgs()
        {
        }
    }

    public sealed class InstancesResetAttachState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CCN Instance ID.
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// CCN Uin (root account).
        /// </summary>
        [Input("ccnUin")]
        public Input<string>? CcnUin { get; set; }

        [Input("instances")]
        private InputList<Inputs.InstancesResetAttachInstanceGetArgs>? _instances;

        /// <summary>
        /// List Of Attachment Instances.
        /// </summary>
        public InputList<Inputs.InstancesResetAttachInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.InstancesResetAttachInstanceGetArgs>());
            set => _instances = value;
        }

        public InstancesResetAttachState()
        {
        }
    }
}
