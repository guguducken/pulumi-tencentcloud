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
    public static class GetUploadBackupInfo
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver upload_backup_info
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
        ///         var uploadBackupInfo = Output.Create(Tencentcloud.Sqlserver.GetUploadBackupInfo.InvokeAsync(new Tencentcloud.Sqlserver.GetUploadBackupInfoArgs
        ///         {
        ///             BackupMigrationId = "mssql-backup-migration-8a0f3eht",
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUploadBackupInfoResult> InvokeAsync(GetUploadBackupInfoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUploadBackupInfoResult>("tencentcloud:Sqlserver/getUploadBackupInfo:getUploadBackupInfo", args ?? new GetUploadBackupInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver upload_backup_info
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
        ///         var uploadBackupInfo = Output.Create(Tencentcloud.Sqlserver.GetUploadBackupInfo.InvokeAsync(new Tencentcloud.Sqlserver.GetUploadBackupInfoArgs
        ///         {
        ///             BackupMigrationId = "mssql-backup-migration-8a0f3eht",
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUploadBackupInfoResult> Invoke(GetUploadBackupInfoInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUploadBackupInfoResult>("tencentcloud:Sqlserver/getUploadBackupInfo:getUploadBackupInfo", args ?? new GetUploadBackupInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUploadBackupInfoArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Input("backupMigrationId", required: true)]
        public string BackupMigrationId { get; set; } = null!;

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

        public GetUploadBackupInfoArgs()
        {
        }
    }

    public sealed class GetUploadBackupInfoInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Input("backupMigrationId", required: true)]
        public Input<string> BackupMigrationId { get; set; } = null!;

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

        public GetUploadBackupInfoInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUploadBackupInfoResult
    {
        public readonly string BackupMigrationId;
        /// <summary>
        /// Bucket name.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// Temporary key expiration time.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Storage path.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Bucket location information.
        /// </summary>
        public readonly string Region;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Temporary key start time.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Temporary key ID.
        /// </summary>
        public readonly string TmpSecretId;
        /// <summary>
        /// Temporary key (Key).
        /// </summary>
        public readonly string TmpSecretKey;
        /// <summary>
        /// Temporary key (Token).
        /// </summary>
        public readonly string XCosSecurityToken;

        [OutputConstructor]
        private GetUploadBackupInfoResult(
            string backupMigrationId,

            string bucketName,

            string expiredTime,

            string id,

            string instanceId,

            string path,

            string region,

            string? resultOutputFile,

            string startTime,

            string tmpSecretId,

            string tmpSecretKey,

            string xCosSecurityToken)
        {
            BackupMigrationId = backupMigrationId;
            BucketName = bucketName;
            ExpiredTime = expiredTime;
            Id = id;
            InstanceId = instanceId;
            Path = path;
            Region = region;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            TmpSecretId = tmpSecretId;
            TmpSecretKey = tmpSecretKey;
            XCosSecurityToken = xCosSecurityToken;
        }
    }
}
