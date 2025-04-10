// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    public static class GetDeliveryConfigByGroupId
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf delivery_config_by_group_id
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
        ///         var deliveryConfigByGroupId = Output.Create(Tencentcloud.Tsf.GetDeliveryConfigByGroupId.InvokeAsync(new Tencentcloud.Tsf.GetDeliveryConfigByGroupIdArgs
        ///         {
        ///             GroupId = "group-yrjkln9v",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDeliveryConfigByGroupIdResult> InvokeAsync(GetDeliveryConfigByGroupIdArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeliveryConfigByGroupIdResult>("tencentcloud:Tsf/getDeliveryConfigByGroupId:getDeliveryConfigByGroupId", args ?? new GetDeliveryConfigByGroupIdArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf delivery_config_by_group_id
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
        ///         var deliveryConfigByGroupId = Output.Create(Tencentcloud.Tsf.GetDeliveryConfigByGroupId.InvokeAsync(new Tencentcloud.Tsf.GetDeliveryConfigByGroupIdArgs
        ///         {
        ///             GroupId = "group-yrjkln9v",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDeliveryConfigByGroupIdResult> Invoke(GetDeliveryConfigByGroupIdInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDeliveryConfigByGroupIdResult>("tencentcloud:Tsf/getDeliveryConfigByGroupId:getDeliveryConfigByGroupId", args ?? new GetDeliveryConfigByGroupIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeliveryConfigByGroupIdArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// groupId.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDeliveryConfigByGroupIdArgs()
        {
        }
    }

    public sealed class GetDeliveryConfigByGroupIdInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// groupId.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDeliveryConfigByGroupIdInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeliveryConfigByGroupIdResult
    {
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// configuration item for deliver to a Kafka.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDeliveryConfigByGroupIdResultResult> Results;

        [OutputConstructor]
        private GetDeliveryConfigByGroupIdResult(
            string groupId,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetDeliveryConfigByGroupIdResultResult> results)
        {
            GroupId = groupId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
        }
    }
}
