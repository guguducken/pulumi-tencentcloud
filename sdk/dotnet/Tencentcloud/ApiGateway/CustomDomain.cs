// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    /// <summary>
    /// Use this resource to create custom domain of API gateway.
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
    ///         var foo = new Tencentcloud.ApiGateway.CustomDomain("foo", new Tencentcloud.ApiGateway.CustomDomainArgs
    ///         {
    ///             DefaultDomain = "service-ohxqslqe-1259649581.gz.apigw.tencentcs.com",
    ///             IsDefaultMapping = false,
    ///             NetType = "OUTER",
    ///             PathMappings = 
    ///             {
    ///                 "/good#test",
    ///                 "/root#release",
    ///             },
    ///             Protocol = "http",
    ///             ServiceId = "service-ohxqslqe",
    ///             SubDomain = "tic-test.dnsv1.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:ApiGateway/customDomain:CustomDomain")]
    public partial class CustomDomain : Pulumi.CustomResource
    {
        /// <summary>
        /// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&amp;https`.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// Default domain name.
        /// </summary>
        [Output("defaultDomain")]
        public Output<string> DefaultDomain { get; private set; } = null!;

        /// <summary>
        /// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
        /// </summary>
        [Output("isDefaultMapping")]
        public Output<bool?> IsDefaultMapping { get; private set; } = null!;

        /// <summary>
        /// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
        /// </summary>
        [Output("isForcedHttps")]
        public Output<bool?> IsForcedHttps { get; private set; } = null!;

        /// <summary>
        /// Network type. Valid values: `OUTER`, `INNER`.
        /// </summary>
        [Output("netType")]
        public Output<string> NetType { get; private set; } = null!;

        /// <summary>
        /// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
        /// </summary>
        [Output("pathMappings")]
        public Output<ImmutableArray<string>> PathMappings { get; private set; } = null!;

        /// <summary>
        /// Protocol supported by service. Valid values: `http`, `https`, `http&amp;https`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Unique service ID.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Custom domain name to be bound.
        /// </summary>
        [Output("subDomain")]
        public Output<string> SubDomain { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDomain(string name, CustomDomainArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/customDomain:CustomDomain", name, args ?? new CustomDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomDomain(string name, Input<string> id, CustomDomainState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/customDomain:CustomDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomDomain Get(string name, Input<string> id, CustomDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomDomain(name, id, state, options);
        }
    }

    public sealed class CustomDomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&amp;https`.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Default domain name.
        /// </summary>
        [Input("defaultDomain", required: true)]
        public Input<string> DefaultDomain { get; set; } = null!;

        /// <summary>
        /// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
        /// </summary>
        [Input("isDefaultMapping")]
        public Input<bool>? IsDefaultMapping { get; set; }

        /// <summary>
        /// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
        /// </summary>
        [Input("isForcedHttps")]
        public Input<bool>? IsForcedHttps { get; set; }

        /// <summary>
        /// Network type. Valid values: `OUTER`, `INNER`.
        /// </summary>
        [Input("netType", required: true)]
        public Input<string> NetType { get; set; } = null!;

        [Input("pathMappings")]
        private InputList<string>? _pathMappings;

        /// <summary>
        /// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
        /// </summary>
        public InputList<string> PathMappings
        {
            get => _pathMappings ?? (_pathMappings = new InputList<string>());
            set => _pathMappings = value;
        }

        /// <summary>
        /// Protocol supported by service. Valid values: `http`, `https`, `http&amp;https`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Unique service ID.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// Custom domain name to be bound.
        /// </summary>
        [Input("subDomain", required: true)]
        public Input<string> SubDomain { get; set; } = null!;

        public CustomDomainArgs()
        {
        }
    }

    public sealed class CustomDomainState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&amp;https`.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Default domain name.
        /// </summary>
        [Input("defaultDomain")]
        public Input<string>? DefaultDomain { get; set; }

        /// <summary>
        /// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
        /// </summary>
        [Input("isDefaultMapping")]
        public Input<bool>? IsDefaultMapping { get; set; }

        /// <summary>
        /// Whether to force HTTP requests to jump to HTTPS, default to false. When the parameter is true, the API gateway will redirect all HTTP protocol requests using the custom domain name to the HTTPS protocol for forwarding.
        /// </summary>
        [Input("isForcedHttps")]
        public Input<bool>? IsForcedHttps { get; set; }

        /// <summary>
        /// Network type. Valid values: `OUTER`, `INNER`.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        [Input("pathMappings")]
        private InputList<string>? _pathMappings;

        /// <summary>
        /// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
        /// </summary>
        public InputList<string> PathMappings
        {
            get => _pathMappings ?? (_pathMappings = new InputList<string>());
            set => _pathMappings = value;
        }

        /// <summary>
        /// Protocol supported by service. Valid values: `http`, `https`, `http&amp;https`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Unique service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// Domain name resolution status. `1` means normal analysis, `0` means parsing failed.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Custom domain name to be bound.
        /// </summary>
        [Input("subDomain")]
        public Input<string>? SubDomain { get; set; }

        public CustomDomainState()
        {
        }
    }
}
