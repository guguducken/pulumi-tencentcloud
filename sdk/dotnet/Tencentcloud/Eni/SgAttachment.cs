// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eni
{
    /// <summary>
    /// Provides a resource to create a eni_sg_attachment
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
    ///         var eniSgAttachment = new Tencentcloud.Eni.SgAttachment("eniSgAttachment", new Tencentcloud.Eni.SgAttachmentArgs
    ///         {
    ///             NetworkInterfaceIds = "eni-p0hkgx8p",
    ///             SecurityGroupIds = 
    ///             {
    ///                 "sg-902tl7t7",
    ///                 "sg-edmur627",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// vpc eni_sg_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Eni/sgAttachment:SgAttachment eni_sg_attachment eni_sg_attachment_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Eni/sgAttachment:SgAttachment")]
    public partial class SgAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ENI instance ID. Such as:eni-pxir56ns. It Only support set one eni instance now.
        /// </summary>
        [Output("networkInterfaceIds")]
        public Output<string> NetworkInterfaceIds { get; private set; } = null!;

        /// <summary>
        /// Security group instance ID, for example:sg-33ocnj9n, can be obtained through DescribeSecurityGroups. There is a limit of 100 instances per request.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a SgAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SgAttachment(string name, SgAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Eni/sgAttachment:SgAttachment", name, args ?? new SgAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SgAttachment(string name, Input<string> id, SgAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Eni/sgAttachment:SgAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SgAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SgAttachment Get(string name, Input<string> id, SgAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SgAttachment(name, id, state, options);
        }
    }

    public sealed class SgAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ENI instance ID. Such as:eni-pxir56ns. It Only support set one eni instance now.
        /// </summary>
        [Input("networkInterfaceIds", required: true)]
        public Input<string> NetworkInterfaceIds { get; set; } = null!;

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group instance ID, for example:sg-33ocnj9n, can be obtained through DescribeSecurityGroups. There is a limit of 100 instances per request.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        public SgAttachmentArgs()
        {
        }
    }

    public sealed class SgAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ENI instance ID. Such as:eni-pxir56ns. It Only support set one eni instance now.
        /// </summary>
        [Input("networkInterfaceIds")]
        public Input<string>? NetworkInterfaceIds { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group instance ID, for example:sg-33ocnj9n, can be obtained through DescribeSecurityGroups. There is a limit of 100 instances per request.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        public SgAttachmentState()
        {
        }
    }
}
