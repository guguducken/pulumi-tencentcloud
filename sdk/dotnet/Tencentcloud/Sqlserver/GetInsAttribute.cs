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
    public static class GetInsAttribute
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver_ins_attribute
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
        ///         var insAttribute = Output.Create(Tencentcloud.Sqlserver.GetInsAttribute.InvokeAsync(new Tencentcloud.Sqlserver.GetInsAttributeArgs
        ///         {
        ///             InstanceId = "mssql-gyg9xycl",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInsAttributeResult> InvokeAsync(GetInsAttributeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInsAttributeResult>("tencentcloud:Sqlserver/getInsAttribute:getInsAttribute", args ?? new GetInsAttributeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver_ins_attribute
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
        ///         var insAttribute = Output.Create(Tencentcloud.Sqlserver.GetInsAttribute.InvokeAsync(new Tencentcloud.Sqlserver.GetInsAttributeArgs
        ///         {
        ///             InstanceId = "mssql-gyg9xycl",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInsAttributeResult> Invoke(GetInsAttributeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInsAttributeResult>("tencentcloud:Sqlserver/getInsAttribute:getInsAttribute", args ?? new GetInsAttributeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInsAttributeArgs : Pulumi.InvokeArgs
    {
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

        public GetInsAttributeArgs()
        {
        }
    }

    public sealed class GetInsAttributeInvokeArgs : Pulumi.InvokeArgs
    {
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

        public GetInsAttributeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInsAttributeResult
    {
        /// <summary>
        /// Block process threshold in milliseconds.
        /// </summary>
        public readonly int BlockedThreshold;
        /// <summary>
        /// Retention period for the files of slow SQL, blocking, deadlock, and extended events.
        /// </summary>
        public readonly int EventSaveDays;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// The number of retained archive backups.
        /// </summary>
        public readonly int RegularBackupCounts;
        /// <summary>
        /// Archive backup status. Valid values: enable (enabled), disable (disabled).
        /// </summary>
        public readonly string RegularBackupEnable;
        /// <summary>
        /// Archive backup retention period: [90-3650] days.
        /// </summary>
        public readonly int RegularBackupSaveDays;
        /// <summary>
        /// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
        /// </summary>
        public readonly string RegularBackupStartTime;
        /// <summary>
        /// Archive backup policy. Valid values: years (yearly); quarters (quarterly);months` (monthly).
        /// </summary>
        public readonly string RegularBackupStrategy;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// TDE Transparent Data Encryption Configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInsAttributeTdeConfigResult> TdeConfigs;

        [OutputConstructor]
        private GetInsAttributeResult(
            int blockedThreshold,

            int eventSaveDays,

            string id,

            string instanceId,

            int regularBackupCounts,

            string regularBackupEnable,

            int regularBackupSaveDays,

            string regularBackupStartTime,

            string regularBackupStrategy,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetInsAttributeTdeConfigResult> tdeConfigs)
        {
            BlockedThreshold = blockedThreshold;
            EventSaveDays = eventSaveDays;
            Id = id;
            InstanceId = instanceId;
            RegularBackupCounts = regularBackupCounts;
            RegularBackupEnable = regularBackupEnable;
            RegularBackupSaveDays = regularBackupSaveDays;
            RegularBackupStartTime = regularBackupStartTime;
            RegularBackupStrategy = regularBackupStrategy;
            ResultOutputFile = resultOutputFile;
            TdeConfigs = tdeConfigs;
        }
    }
}
