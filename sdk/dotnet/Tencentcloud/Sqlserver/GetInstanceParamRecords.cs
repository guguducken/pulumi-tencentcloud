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
    public static class GetInstanceParamRecords
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver instance_param_records
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
        ///         var instanceParamRecords = Output.Create(Tencentcloud.Sqlserver.GetInstanceParamRecords.InvokeAsync(new Tencentcloud.Sqlserver.GetInstanceParamRecordsArgs
        ///         {
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceParamRecordsResult> InvokeAsync(GetInstanceParamRecordsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceParamRecordsResult>("tencentcloud:Sqlserver/getInstanceParamRecords:getInstanceParamRecords", args ?? new GetInstanceParamRecordsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver instance_param_records
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
        ///         var instanceParamRecords = Output.Create(Tencentcloud.Sqlserver.GetInstanceParamRecords.InvokeAsync(new Tencentcloud.Sqlserver.GetInstanceParamRecordsArgs
        ///         {
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceParamRecordsResult> Invoke(GetInstanceParamRecordsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceParamRecordsResult>("tencentcloud:Sqlserver/getInstanceParamRecords:getInstanceParamRecords", args ?? new GetInstanceParamRecordsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceParamRecordsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter InstanceId of the DescribeDBInstances API.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceParamRecordsArgs()
        {
        }
    }

    public sealed class GetInstanceParamRecordsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter InstanceId of the DescribeDBInstances API.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceParamRecordsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceParamRecordsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Parameter modification records.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceParamRecordsItemResult> Items;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceParamRecordsResult(
            string id,

            string instanceId,

            ImmutableArray<Outputs.GetInstanceParamRecordsItemResult> items,

            string? resultOutputFile)
        {
            Id = id;
            InstanceId = instanceId;
            Items = items;
            ResultOutputFile = resultOutputFile;
        }
    }
}
