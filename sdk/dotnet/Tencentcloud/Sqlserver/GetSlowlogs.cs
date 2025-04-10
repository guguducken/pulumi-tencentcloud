// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    public static class GetSlowlogs
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver slowlogs
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
        ///         var slowlogs = Output.Create(Tencentcloud.Sqlserver.GetSlowlogs.InvokeAsync(new Tencentcloud.Sqlserver.GetSlowlogsArgs
        ///         {
        ///             EndTime = "2023-05-18 00:00:00",
        ///             InstanceId = "mssql-qelbzgwf",
        ///             StartTime = "2020-05-01 00:00:00",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSlowlogsResult> InvokeAsync(GetSlowlogsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSlowlogsResult>("tencentcloud:Sqlserver/getSlowlogs:getSlowlogs", args ?? new GetSlowlogsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver slowlogs
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
        ///         var slowlogs = Output.Create(Tencentcloud.Sqlserver.GetSlowlogs.InvokeAsync(new Tencentcloud.Sqlserver.GetSlowlogsArgs
        ///         {
        ///             EndTime = "2023-05-18 00:00:00",
        ///             InstanceId = "mssql-qelbzgwf",
        ///             StartTime = "2020-05-01 00:00:00",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSlowlogsResult> Invoke(GetSlowlogsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSlowlogsResult>("tencentcloud:Sqlserver/getSlowlogs:getSlowlogs", args ?? new GetSlowlogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowlogsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Query end time.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Query start time.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetSlowlogsArgs()
        {
        }
    }

    public sealed class GetSlowlogsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Query end time.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Query start time.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetSlowlogsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSlowlogsResult
    {
        /// <summary>
        /// File generation end time.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Information list of slow query logs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlowlogsSlowlogResult> SqlserverSlowlogs;
        /// <summary>
        /// File generation start time.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GetSlowlogsResult(
            string endTime,

            string id,

            string instanceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSlowlogsSlowlogResult> slowlogs,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            SqlserverSlowlogs = slowlogs;
            StartTime = startTime;
        }
    }
}
