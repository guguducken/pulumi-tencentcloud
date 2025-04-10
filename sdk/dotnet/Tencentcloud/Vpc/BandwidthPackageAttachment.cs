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
    /// <summary>
    /// Provides a resource to create a vpc bandwidth_package_attachment
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
    ///         var bandwidthPackageAttachment = new Tencentcloud.Vpc.BandwidthPackageAttachment("bandwidthPackageAttachment", new Tencentcloud.Vpc.BandwidthPackageAttachmentArgs
    ///         {
    ///             BandwidthPackageId = "bwp-atmf0p9g",
    ///             NetworkType = "BGP",
    ///             Protocol = "",
    ///             ResourceId = "lb-dv1ai6ma",
    ///             ResourceType = "LoadBalance",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/bandwidthPackageAttachment:BandwidthPackageAttachment")]
    public partial class BandwidthPackageAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Bandwidth package unique ID, in the form of `bwp-xxxx`.
        /// </summary>
        [Output("bandwidthPackageId")]
        public Output<string> BandwidthPackageId { get; private set; } = null!;

        /// <summary>
        /// Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        /// </summary>
        [Output("networkType")]
        public Output<string?> NetworkType { get; private set; } = null!;

        /// <summary>
        /// Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource types, including `Address`, `LoadBalance`.
        /// </summary>
        [Output("resourceType")]
        public Output<string?> ResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackageAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackageAttachment(string name, BandwidthPackageAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/bandwidthPackageAttachment:BandwidthPackageAttachment", name, args ?? new BandwidthPackageAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackageAttachment(string name, Input<string> id, BandwidthPackageAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/bandwidthPackageAttachment:BandwidthPackageAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthPackageAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthPackageAttachment Get(string name, Input<string> id, BandwidthPackageAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthPackageAttachment(name, id, state, options);
        }
    }

    public sealed class BandwidthPackageAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bandwidth package unique ID, in the form of `bwp-xxxx`.
        /// </summary>
        [Input("bandwidthPackageId", required: true)]
        public Input<string> BandwidthPackageId { get; set; } = null!;

        /// <summary>
        /// Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// Resource types, including `Address`, `LoadBalance`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public BandwidthPackageAttachmentArgs()
        {
        }
    }

    public sealed class BandwidthPackageAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bandwidth package unique ID, in the form of `bwp-xxxx`.
        /// </summary>
        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        /// <summary>
        /// Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Resource types, including `Address`, `LoadBalance`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public BandwidthPackageAttachmentState()
        {
        }
    }
}
