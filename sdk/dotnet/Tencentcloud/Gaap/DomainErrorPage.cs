// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Gaap
{
    /// <summary>
    /// Provide a resource to custom error page info for a GAAP HTTP domain.
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
    ///         var fooProxy = new Tencentcloud.Gaap.Proxy("fooProxy", new Tencentcloud.Gaap.ProxyArgs
    ///         {
    ///             Bandwidth = 10,
    ///             Concurrent = 2,
    ///             AccessRegion = "SouthChina",
    ///             RealserverRegion = "NorthChina",
    ///         });
    ///         var fooLayer7Listener = new Tencentcloud.Gaap.Layer7Listener("fooLayer7Listener", new Tencentcloud.Gaap.Layer7ListenerArgs
    ///         {
    ///             Protocol = "HTTP",
    ///             Port = 80,
    ///             ProxyId = fooProxy.Id,
    ///         });
    ///         var fooHttpDomain = new Tencentcloud.Gaap.HttpDomain("fooHttpDomain", new Tencentcloud.Gaap.HttpDomainArgs
    ///         {
    ///             ListenerId = fooLayer7Listener.Id,
    ///             Domain = "www.qq.com",
    ///         });
    ///         var fooDomainErrorPage = new Tencentcloud.Gaap.DomainErrorPage("fooDomainErrorPage", new Tencentcloud.Gaap.DomainErrorPageArgs
    ///         {
    ///             ListenerId = fooLayer7Listener.Id,
    ///             Domain = fooHttpDomain.Domain,
    ///             ErrorCodes = 
    ///             {
    ///                 404,
    ///                 503,
    ///             },
    ///             Body = "bad request",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Gaap/domainErrorPage:DomainErrorPage")]
    public partial class DomainErrorPage : Pulumi.CustomResource
    {
        /// <summary>
        /// New response body.
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        /// <summary>
        /// Response headers to be removed.
        /// </summary>
        [Output("clearHeaders")]
        public Output<ImmutableArray<string>> ClearHeaders { get; private set; } = null!;

        /// <summary>
        /// HTTP domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Original error codes.
        /// </summary>
        [Output("errorCodes")]
        public Output<ImmutableArray<int>> ErrorCodes { get; private set; } = null!;

        /// <summary>
        /// ID of the layer7 listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// New error code.
        /// </summary>
        [Output("newErrorCode")]
        public Output<int?> NewErrorCode { get; private set; } = null!;

        /// <summary>
        /// Response headers to be set.
        /// </summary>
        [Output("setHeaders")]
        public Output<ImmutableDictionary<string, object>?> SetHeaders { get; private set; } = null!;


        /// <summary>
        /// Create a DomainErrorPage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainErrorPage(string name, DomainErrorPageArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/domainErrorPage:DomainErrorPage", name, args ?? new DomainErrorPageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainErrorPage(string name, Input<string> id, DomainErrorPageState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/domainErrorPage:DomainErrorPage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainErrorPage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainErrorPage Get(string name, Input<string> id, DomainErrorPageState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainErrorPage(name, id, state, options);
        }
    }

    public sealed class DomainErrorPageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// New response body.
        /// </summary>
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        [Input("clearHeaders")]
        private InputList<string>? _clearHeaders;

        /// <summary>
        /// Response headers to be removed.
        /// </summary>
        public InputList<string> ClearHeaders
        {
            get => _clearHeaders ?? (_clearHeaders = new InputList<string>());
            set => _clearHeaders = value;
        }

        /// <summary>
        /// HTTP domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("errorCodes", required: true)]
        private InputList<int>? _errorCodes;

        /// <summary>
        /// Original error codes.
        /// </summary>
        public InputList<int> ErrorCodes
        {
            get => _errorCodes ?? (_errorCodes = new InputList<int>());
            set => _errorCodes = value;
        }

        /// <summary>
        /// ID of the layer7 listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// New error code.
        /// </summary>
        [Input("newErrorCode")]
        public Input<int>? NewErrorCode { get; set; }

        [Input("setHeaders")]
        private InputMap<object>? _setHeaders;

        /// <summary>
        /// Response headers to be set.
        /// </summary>
        public InputMap<object> SetHeaders
        {
            get => _setHeaders ?? (_setHeaders = new InputMap<object>());
            set => _setHeaders = value;
        }

        public DomainErrorPageArgs()
        {
        }
    }

    public sealed class DomainErrorPageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// New response body.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("clearHeaders")]
        private InputList<string>? _clearHeaders;

        /// <summary>
        /// Response headers to be removed.
        /// </summary>
        public InputList<string> ClearHeaders
        {
            get => _clearHeaders ?? (_clearHeaders = new InputList<string>());
            set => _clearHeaders = value;
        }

        /// <summary>
        /// HTTP domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("errorCodes")]
        private InputList<int>? _errorCodes;

        /// <summary>
        /// Original error codes.
        /// </summary>
        public InputList<int> ErrorCodes
        {
            get => _errorCodes ?? (_errorCodes = new InputList<int>());
            set => _errorCodes = value;
        }

        /// <summary>
        /// ID of the layer7 listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// New error code.
        /// </summary>
        [Input("newErrorCode")]
        public Input<int>? NewErrorCode { get; set; }

        [Input("setHeaders")]
        private InputMap<object>? _setHeaders;

        /// <summary>
        /// Response headers to be set.
        /// </summary>
        public InputMap<object> SetHeaders
        {
            get => _setHeaders ?? (_setHeaders = new InputMap<object>());
            set => _setHeaders = value;
        }

        public DomainErrorPageState()
        {
        }
    }
}
