// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    public static class GetProxyVersion
    {
        /// <summary>
        /// Use this data source to query detailed information of cynosdb proxy_version
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var proxyVersion = Output.Create(Tencentcloud.Cynosdb.GetProxyVersion.InvokeAsync(new Tencentcloud.Cynosdb.GetProxyVersionArgs
        ///         {
        ///             ClusterId = "cynosdbmysql-bws8h88b",
        ///             ProxyGroupId = "cynosdbmysql-proxy-l6zf9t30",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProxyVersionResult> InvokeAsync(GetProxyVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProxyVersionResult>("tencentcloud:Cynosdb/getProxyVersion:getProxyVersion", args ?? new GetProxyVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cynosdb proxy_version
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var proxyVersion = Output.Create(Tencentcloud.Cynosdb.GetProxyVersion.InvokeAsync(new Tencentcloud.Cynosdb.GetProxyVersionArgs
        ///         {
        ///             ClusterId = "cynosdbmysql-bws8h88b",
        ///             ProxyGroupId = "cynosdbmysql-proxy-l6zf9t30",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProxyVersionResult> Invoke(GetProxyVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProxyVersionResult>("tencentcloud:Cynosdb/getProxyVersion:getProxyVersion", args ?? new GetProxyVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProxyVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Database Agent Group ID.
        /// </summary>
        [Input("proxyGroupId")]
        public string? ProxyGroupId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetProxyVersionArgs()
        {
        }
    }

    public sealed class GetProxyVersionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Database Agent Group ID.
        /// </summary>
        [Input("proxyGroupId")]
        public Input<string>? ProxyGroupId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetProxyVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProxyVersionResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// Current proxy version number note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string CurrentProxyVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ProxyGroupId;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Supported Database Agent Version Collection Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> SupportProxyVersions;

        [OutputConstructor]
        private GetProxyVersionResult(
            string clusterId,

            string currentProxyVersion,

            string id,

            string? proxyGroupId,

            string? resultOutputFile,

            ImmutableArray<string> supportProxyVersions)
        {
            ClusterId = clusterId;
            CurrentProxyVersion = currentProxyVersion;
            Id = id;
            ProxyGroupId = proxyGroupId;
            ResultOutputFile = resultOutputFile;
            SupportProxyVersions = supportProxyVersions;
        }
    }
}
