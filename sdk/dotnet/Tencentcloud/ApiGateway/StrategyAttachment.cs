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
    /// Use this resource to create IP strategy attachment of API gateway.
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
    ///         var serviceService = new Tencentcloud.ApiGateway.Service("serviceService", new Tencentcloud.ApiGateway.ServiceArgs
    ///         {
    ///             ServiceName = "niceservice",
    ///             Protocol = "http&amp;https",
    ///             ServiceDesc = "your nice service",
    ///             NetTypes = 
    ///             {
    ///                 "INNER",
    ///                 "OUTER",
    ///             },
    ///             IpVersion = "IPv4",
    ///         });
    ///         var testIpStrategy = new Tencentcloud.ApiGateway.IpStrategy("testIpStrategy", new Tencentcloud.ApiGateway.IpStrategyArgs
    ///         {
    ///             ServiceId = serviceService.Id,
    ///             StrategyName = "tf_test",
    ///             StrategyType = "BLACK",
    ///             StrategyData = "9.9.9.9",
    ///         });
    ///         var api = new Tencentcloud.ApiGateway.Api("api", new Tencentcloud.ApiGateway.ApiArgs
    ///         {
    ///             ServiceId = serviceService.Id,
    ///             ApiName = "hello_update",
    ///             ApiDesc = "my hello api update",
    ///             AuthType = "SECRET",
    ///             Protocol = "HTTP",
    ///             EnableCors = true,
    ///             RequestConfigPath = "/user/info",
    ///             RequestConfigMethod = "POST",
    ///             RequestParameters = 
    ///             {
    ///                 new Tencentcloud.ApiGateway.Inputs.ApiRequestParameterArgs
    ///                 {
    ///                     Name = "email",
    ///                     Position = "QUERY",
    ///                     Type = "string",
    ///                     Desc = "your email please?",
    ///                     DefaultValue = "tom@qq.com",
    ///                     Required = true,
    ///                 },
    ///             },
    ///             ServiceConfigType = "HTTP",
    ///             ServiceConfigTimeout = 10,
    ///             ServiceConfigUrl = "http://www.tencent.com",
    ///             ServiceConfigPath = "/user",
    ///             ServiceConfigMethod = "POST",
    ///             ResponseType = "XML",
    ///             ResponseSuccessExample = "&lt;note&gt;success&lt;/note&gt;",
    ///             ResponseFailExample = "&lt;note&gt;fail&lt;/note&gt;",
    ///             ResponseErrorCodes = 
    ///             {
    ///                 new Tencentcloud.ApiGateway.Inputs.ApiResponseErrorCodeArgs
    ///                 {
    ///                     Code = 10,
    ///                     Msg = "system error",
    ///                     Desc = "system error code",
    ///                     ConvertedCode = -10,
    ///                     NeedConvert = true,
    ///                 },
    ///             },
    ///         });
    ///         var serviceServiceRelease = new Tencentcloud.ApiGateway.ServiceRelease("serviceServiceRelease", new Tencentcloud.ApiGateway.ServiceReleaseArgs
    ///         {
    ///             ServiceId = serviceService.Id,
    ///             EnvironmentName = "release",
    ///             ReleaseDesc = "test service release",
    ///         });
    ///         var testStrategyAttachment = new Tencentcloud.ApiGateway.StrategyAttachment("testStrategyAttachment", new Tencentcloud.ApiGateway.StrategyAttachmentArgs
    ///         {
    ///             ServiceId = serviceServiceRelease.ServiceId,
    ///             StrategyId = testIpStrategy.StrategyId,
    ///             EnvironmentName = "release",
    ///             BindApiId = api.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IP strategy attachment of API gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:ApiGateway/strategyAttachment:StrategyAttachment test service-pk2r6bcc#IPStrategy-4kz2ljfi#api-h3wc5r0s#release
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:ApiGateway/strategyAttachment:StrategyAttachment")]
    public partial class StrategyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The API that needs to be bound.
        /// </summary>
        [Output("bindApiId")]
        public Output<string> BindApiId { get; private set; } = null!;

        /// <summary>
        /// The environment of the strategy association. Valid values: `test`, `release`, `prepub`.
        /// </summary>
        [Output("environmentName")]
        public Output<string> EnvironmentName { get; private set; } = null!;

        /// <summary>
        /// The ID of the API gateway service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the API gateway strategy.
        /// </summary>
        [Output("strategyId")]
        public Output<string> StrategyId { get; private set; } = null!;


        /// <summary>
        /// Create a StrategyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StrategyAttachment(string name, StrategyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/strategyAttachment:StrategyAttachment", name, args ?? new StrategyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StrategyAttachment(string name, Input<string> id, StrategyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/strategyAttachment:StrategyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StrategyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StrategyAttachment Get(string name, Input<string> id, StrategyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new StrategyAttachment(name, id, state, options);
        }
    }

    public sealed class StrategyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API that needs to be bound.
        /// </summary>
        [Input("bindApiId", required: true)]
        public Input<string> BindApiId { get; set; } = null!;

        /// <summary>
        /// The environment of the strategy association. Valid values: `test`, `release`, `prepub`.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The ID of the API gateway service.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// The ID of the API gateway strategy.
        /// </summary>
        [Input("strategyId", required: true)]
        public Input<string> StrategyId { get; set; } = null!;

        public StrategyAttachmentArgs()
        {
        }
    }

    public sealed class StrategyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API that needs to be bound.
        /// </summary>
        [Input("bindApiId")]
        public Input<string>? BindApiId { get; set; }

        /// <summary>
        /// The environment of the strategy association. Valid values: `test`, `release`, `prepub`.
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// The ID of the API gateway service.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The ID of the API gateway strategy.
        /// </summary>
        [Input("strategyId")]
        public Input<string>? StrategyId { get; set; }

        public StrategyAttachmentState()
        {
        }
    }
}
