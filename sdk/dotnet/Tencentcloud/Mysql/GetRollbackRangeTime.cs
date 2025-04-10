// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    public static class GetRollbackRangeTime
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql rollback_range_time
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
        ///         var rollbackRangeTime = Output.Create(Tencentcloud.Mysql.GetRollbackRangeTime.InvokeAsync(new Tencentcloud.Mysql.GetRollbackRangeTimeArgs
        ///         {
        ///             InstanceIds = 
        ///             {
        ///                 "cdb-fitq5t9h",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRollbackRangeTimeResult> InvokeAsync(GetRollbackRangeTimeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRollbackRangeTimeResult>("tencentcloud:Mysql/getRollbackRangeTime:getRollbackRangeTime", args ?? new GetRollbackRangeTimeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql rollback_range_time
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
        ///         var rollbackRangeTime = Output.Create(Tencentcloud.Mysql.GetRollbackRangeTime.InvokeAsync(new Tencentcloud.Mysql.GetRollbackRangeTimeArgs
        ///         {
        ///             InstanceIds = 
        ///             {
        ///                 "cdb-fitq5t9h",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRollbackRangeTimeResult> Invoke(GetRollbackRangeTimeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRollbackRangeTimeResult>("tencentcloud:Mysql/getRollbackRangeTime:getRollbackRangeTime", args ?? new GetRollbackRangeTimeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRollbackRangeTimeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// If the clone instance is not in the same region as the source instance, fill in the region where the clone instance is located, for example: ap-guangzhou.
        /// </summary>
        [Input("backupRegion")]
        public string? BackupRegion { get; set; }

        [Input("instanceIds", required: true)]
        private List<string>? _instanceIds;

        /// <summary>
        /// A list of instance IDs, the format of a single instance ID is: cdb-c1nl9rpv. Same instance ID as displayed in the ApsaraDB for Console page.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Whether the clone instance is in the same zone as the source instance, yes: `false`, no: `true`.
        /// </summary>
        [Input("isRemoteZone")]
        public string? IsRemoteZone { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRollbackRangeTimeArgs()
        {
        }
    }

    public sealed class GetRollbackRangeTimeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// If the clone instance is not in the same region as the source instance, fill in the region where the clone instance is located, for example: ap-guangzhou.
        /// </summary>
        [Input("backupRegion")]
        public Input<string>? BackupRegion { get; set; }

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// A list of instance IDs, the format of a single instance ID is: cdb-c1nl9rpv. Same instance ID as displayed in the ApsaraDB for Console page.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Whether the clone instance is in the same zone as the source instance, yes: `false`, no: `true`.
        /// </summary>
        [Input("isRemoteZone")]
        public Input<string>? IsRemoteZone { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRollbackRangeTimeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRollbackRangeTimeResult
    {
        public readonly string? BackupRegion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string? IsRemoteZone;
        /// <summary>
        /// Returned parameter information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRollbackRangeTimeItemResult> Items;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRollbackRangeTimeResult(
            string? backupRegion,

            string id,

            ImmutableArray<string> instanceIds,

            string? isRemoteZone,

            ImmutableArray<Outputs.GetRollbackRangeTimeItemResult> items,

            string? resultOutputFile)
        {
            BackupRegion = backupRegion;
            Id = id;
            InstanceIds = instanceIds;
            IsRemoteZone = isRemoteZone;
            Items = items;
            ResultOutputFile = resultOutputFile;
        }
    }
}
