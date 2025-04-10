// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    [TencentcloudResourceType("tencentcloud:Vpc/classicLinkAttachment:ClassicLinkAttachment")]
    public partial class ClassicLinkAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// CVM instance ID. It only support set one instance now.
        /// </summary>
        [Output("instanceIds")]
        public Output<string> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a ClassicLinkAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClassicLinkAttachment(string name, ClassicLinkAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/classicLinkAttachment:ClassicLinkAttachment", name, args ?? new ClassicLinkAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClassicLinkAttachment(string name, Input<string> id, ClassicLinkAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/classicLinkAttachment:ClassicLinkAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClassicLinkAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClassicLinkAttachment Get(string name, Input<string> id, ClassicLinkAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ClassicLinkAttachment(name, id, state, options);
        }
    }

    public sealed class ClassicLinkAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CVM instance ID. It only support set one instance now.
        /// </summary>
        [Input("instanceIds", required: true)]
        public Input<string> InstanceIds { get; set; } = null!;

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ClassicLinkAttachmentArgs()
        {
        }
    }

    public sealed class ClassicLinkAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CVM instance ID. It only support set one instance now.
        /// </summary>
        [Input("instanceIds")]
        public Input<string>? InstanceIds { get; set; }

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClassicLinkAttachmentState()
        {
        }
    }
}
