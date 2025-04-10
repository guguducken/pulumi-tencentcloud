// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq
{
    public static class GetProInstanceDetail
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmq pro_instance_detail
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
        ///         var proInstanceDetail = Output.Create(Tencentcloud.Tdmq.GetProInstanceDetail.InvokeAsync(new Tencentcloud.Tdmq.GetProInstanceDetailArgs
        ///         {
        ///             ClusterId = "pulsar-9n95ax58b9vn",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProInstanceDetailResult> InvokeAsync(GetProInstanceDetailArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProInstanceDetailResult>("tencentcloud:Tdmq/getProInstanceDetail:getProInstanceDetail", args ?? new GetProInstanceDetailArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmq pro_instance_detail
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
        ///         var proInstanceDetail = Output.Create(Tencentcloud.Tdmq.GetProInstanceDetail.InvokeAsync(new Tencentcloud.Tdmq.GetProInstanceDetailArgs
        ///         {
        ///             ClusterId = "pulsar-9n95ax58b9vn",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProInstanceDetailResult> Invoke(GetProInstanceDetailInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProInstanceDetailResult>("tencentcloud:Tdmq/getProInstanceDetail:getProInstanceDetail", args ?? new GetProInstanceDetailInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProInstanceDetailArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster Id.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetProInstanceDetailArgs()
        {
        }
    }

    public sealed class GetProInstanceDetailInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster Id.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetProInstanceDetailInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProInstanceDetailResult
    {
        /// <summary>
        /// Cluster Id.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Cluster information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProInstanceDetailClusterInfoResult> ClusterInfos;
        /// <summary>
        /// Cluster specification informationNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProInstanceDetailClusterSpecInfoResult> ClusterSpecInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Cluster network access point informationNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProInstanceDetailNetworkAccessPointInfoResult> NetworkAccessPointInfos;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetProInstanceDetailResult(
            string clusterId,

            ImmutableArray<Outputs.GetProInstanceDetailClusterInfoResult> clusterInfos,

            ImmutableArray<Outputs.GetProInstanceDetailClusterSpecInfoResult> clusterSpecInfos,

            string id,

            ImmutableArray<Outputs.GetProInstanceDetailNetworkAccessPointInfoResult> networkAccessPointInfos,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            ClusterInfos = clusterInfos;
            ClusterSpecInfos = clusterSpecInfos;
            Id = id;
            NetworkAccessPointInfos = networkAccessPointInfos;
            ResultOutputFile = resultOutputFile;
        }
    }
}
