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
    /// <summary>
    /// Provides a resource to create a sqlserver config_backup_strategy
    /// 
    /// ## Example Usage
    /// ### Daily backup
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Tencentcloud.Sqlserver.ConfigBackupStrategy("config", new Tencentcloud.Sqlserver.ConfigBackupStrategyArgs
    ///         {
    ///             InstanceId = local.Sqlserver_id,
    ///             BackupType = "daily",
    ///             BackupTime = 0,
    ///             BackupDay = 1,
    ///             BackupModel = "master_no_pkg",
    ///             BackupCycles = 
    ///             {
    ///                 1,
    ///             },
    ///             BackupSaveDays = 7,
    ///             RegularBackupEnable = "disable",
    ///             RegularBackupSaveDays = 90,
    ///             RegularBackupStrategy = "months",
    ///             RegularBackupCounts = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Weekly backup
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Tencentcloud.Sqlserver.ConfigBackupStrategy("config", new Tencentcloud.Sqlserver.ConfigBackupStrategyArgs
    ///         {
    ///             InstanceId = local.Sqlserver_id,
    ///             BackupType = "weekly",
    ///             BackupTime = 0,
    ///             BackupModel = "master_no_pkg",
    ///             BackupCycles = 
    ///             {
    ///                 1,
    ///                 3,
    ///                 5,
    ///             },
    ///             BackupSaveDays = 7,
    ///             RegularBackupEnable = "disable",
    ///             RegularBackupSaveDays = 90,
    ///             RegularBackupStrategy = "months",
    ///             RegularBackupCounts = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Regular backup
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Tencentcloud.Sqlserver.ConfigBackupStrategy("config", new Tencentcloud.Sqlserver.ConfigBackupStrategyArgs
    ///         {
    ///             InstanceId = local.Sqlserver_id,
    ///             BackupTime = 0,
    ///             BackupModel = "master_no_pkg",
    ///             BackupCycles = 
    ///             {
    ///                 1,
    ///                 3,
    ///             },
    ///             BackupSaveDays = 7,
    ///             RegularBackupEnable = "enable",
    ///             RegularBackupSaveDays = 120,
    ///             RegularBackupStrategy = "months",
    ///             RegularBackupCounts = 1,
    ///             RegularBackupStartTime = "%s",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver config_backup_strategy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/configBackupStrategy:ConfigBackupStrategy config_backup_strategy config_backup_strategy_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/configBackupStrategy:ConfigBackupStrategy")]
    public partial class ConfigBackupStrategy : Pulumi.CustomResource
    {
        /// <summary>
        /// The days of the week on which backup will be performed when `BackupType` is weekly. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
        /// </summary>
        [Output("backupCycles")]
        public Output<ImmutableArray<int>> BackupCycles { get; private set; } = null!;

        /// <summary>
        /// Backup interval in days when the BackupType is daily. The current value can only be 1.
        /// </summary>
        [Output("backupDay")]
        public Output<int?> BackupDay { get; private set; } = null!;

        /// <summary>
        /// Backup mode. Valid values: master_pkg (archive the backup files of the primary node), master_no_pkg (do not archive the backup files of the primary node), slave_pkg (archive the backup files of the replica node), slave_no_pkg (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
        /// </summary>
        [Output("backupModel")]
        public Output<string?> BackupModel { get; private set; } = null!;

        /// <summary>
        /// Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
        /// </summary>
        [Output("backupSaveDays")]
        public Output<int?> BackupSaveDays { get; private set; } = null!;

        /// <summary>
        /// Backup time. Value range: an integer from 0 to 23.
        /// </summary>
        [Output("backupTime")]
        public Output<int?> BackupTime { get; private set; } = null!;

        /// <summary>
        /// Backup type. Valid values: weekly (when length(BackupDay) &lt;=7 &amp;&amp; length(BackupDay) &gt;=2), daily (when length(BackupDay)=1). Default value: daily.
        /// </summary>
        [Output("backupType")]
        public Output<string?> BackupType { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The number of retained archive backups. Default value: 1.
        /// </summary>
        [Output("regularBackupCounts")]
        public Output<int?> RegularBackupCounts { get; private set; } = null!;

        /// <summary>
        /// Archive backup status. Valid values: enable (enabled); disable (disabled). Default value: disable.
        /// </summary>
        [Output("regularBackupEnable")]
        public Output<string?> RegularBackupEnable { get; private set; } = null!;

        /// <summary>
        /// Archive backup retention days. Value range: 90-3650 days. Default value: 365 days.
        /// </summary>
        [Output("regularBackupSaveDays")]
        public Output<int?> RegularBackupSaveDays { get; private set; } = null!;

        /// <summary>
        /// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
        /// </summary>
        [Output("regularBackupStartTime")]
        public Output<string?> RegularBackupStartTime { get; private set; } = null!;

        /// <summary>
        /// Archive backup policy. Valid values: years (yearly); quarters (quarterly); months(monthly); Default value: `months`.
        /// </summary>
        [Output("regularBackupStrategy")]
        public Output<string?> RegularBackupStrategy { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigBackupStrategy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigBackupStrategy(string name, ConfigBackupStrategyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configBackupStrategy:ConfigBackupStrategy", name, args ?? new ConfigBackupStrategyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigBackupStrategy(string name, Input<string> id, ConfigBackupStrategyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configBackupStrategy:ConfigBackupStrategy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigBackupStrategy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigBackupStrategy Get(string name, Input<string> id, ConfigBackupStrategyState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigBackupStrategy(name, id, state, options);
        }
    }

    public sealed class ConfigBackupStrategyArgs : Pulumi.ResourceArgs
    {
        [Input("backupCycles")]
        private InputList<int>? _backupCycles;

        /// <summary>
        /// The days of the week on which backup will be performed when `BackupType` is weekly. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
        /// </summary>
        public InputList<int> BackupCycles
        {
            get => _backupCycles ?? (_backupCycles = new InputList<int>());
            set => _backupCycles = value;
        }

        /// <summary>
        /// Backup interval in days when the BackupType is daily. The current value can only be 1.
        /// </summary>
        [Input("backupDay")]
        public Input<int>? BackupDay { get; set; }

        /// <summary>
        /// Backup mode. Valid values: master_pkg (archive the backup files of the primary node), master_no_pkg (do not archive the backup files of the primary node), slave_pkg (archive the backup files of the replica node), slave_no_pkg (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
        /// </summary>
        [Input("backupModel")]
        public Input<string>? BackupModel { get; set; }

        /// <summary>
        /// Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
        /// </summary>
        [Input("backupSaveDays")]
        public Input<int>? BackupSaveDays { get; set; }

        /// <summary>
        /// Backup time. Value range: an integer from 0 to 23.
        /// </summary>
        [Input("backupTime")]
        public Input<int>? BackupTime { get; set; }

        /// <summary>
        /// Backup type. Valid values: weekly (when length(BackupDay) &lt;=7 &amp;&amp; length(BackupDay) &gt;=2), daily (when length(BackupDay)=1). Default value: daily.
        /// </summary>
        [Input("backupType")]
        public Input<string>? BackupType { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The number of retained archive backups. Default value: 1.
        /// </summary>
        [Input("regularBackupCounts")]
        public Input<int>? RegularBackupCounts { get; set; }

        /// <summary>
        /// Archive backup status. Valid values: enable (enabled); disable (disabled). Default value: disable.
        /// </summary>
        [Input("regularBackupEnable")]
        public Input<string>? RegularBackupEnable { get; set; }

        /// <summary>
        /// Archive backup retention days. Value range: 90-3650 days. Default value: 365 days.
        /// </summary>
        [Input("regularBackupSaveDays")]
        public Input<int>? RegularBackupSaveDays { get; set; }

        /// <summary>
        /// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
        /// </summary>
        [Input("regularBackupStartTime")]
        public Input<string>? RegularBackupStartTime { get; set; }

        /// <summary>
        /// Archive backup policy. Valid values: years (yearly); quarters (quarterly); months(monthly); Default value: `months`.
        /// </summary>
        [Input("regularBackupStrategy")]
        public Input<string>? RegularBackupStrategy { get; set; }

        public ConfigBackupStrategyArgs()
        {
        }
    }

    public sealed class ConfigBackupStrategyState : Pulumi.ResourceArgs
    {
        [Input("backupCycles")]
        private InputList<int>? _backupCycles;

        /// <summary>
        /// The days of the week on which backup will be performed when `BackupType` is weekly. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
        /// </summary>
        public InputList<int> BackupCycles
        {
            get => _backupCycles ?? (_backupCycles = new InputList<int>());
            set => _backupCycles = value;
        }

        /// <summary>
        /// Backup interval in days when the BackupType is daily. The current value can only be 1.
        /// </summary>
        [Input("backupDay")]
        public Input<int>? BackupDay { get; set; }

        /// <summary>
        /// Backup mode. Valid values: master_pkg (archive the backup files of the primary node), master_no_pkg (do not archive the backup files of the primary node), slave_pkg (archive the backup files of the replica node), slave_no_pkg (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
        /// </summary>
        [Input("backupModel")]
        public Input<string>? BackupModel { get; set; }

        /// <summary>
        /// Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
        /// </summary>
        [Input("backupSaveDays")]
        public Input<int>? BackupSaveDays { get; set; }

        /// <summary>
        /// Backup time. Value range: an integer from 0 to 23.
        /// </summary>
        [Input("backupTime")]
        public Input<int>? BackupTime { get; set; }

        /// <summary>
        /// Backup type. Valid values: weekly (when length(BackupDay) &lt;=7 &amp;&amp; length(BackupDay) &gt;=2), daily (when length(BackupDay)=1). Default value: daily.
        /// </summary>
        [Input("backupType")]
        public Input<string>? BackupType { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The number of retained archive backups. Default value: 1.
        /// </summary>
        [Input("regularBackupCounts")]
        public Input<int>? RegularBackupCounts { get; set; }

        /// <summary>
        /// Archive backup status. Valid values: enable (enabled); disable (disabled). Default value: disable.
        /// </summary>
        [Input("regularBackupEnable")]
        public Input<string>? RegularBackupEnable { get; set; }

        /// <summary>
        /// Archive backup retention days. Value range: 90-3650 days. Default value: 365 days.
        /// </summary>
        [Input("regularBackupSaveDays")]
        public Input<int>? RegularBackupSaveDays { get; set; }

        /// <summary>
        /// Archive backup start date in YYYY-MM-DD format, which is the current time by default.
        /// </summary>
        [Input("regularBackupStartTime")]
        public Input<string>? RegularBackupStartTime { get; set; }

        /// <summary>
        /// Archive backup policy. Valid values: years (yearly); quarters (quarterly); months(monthly); Default value: `months`.
        /// </summary>
        [Input("regularBackupStrategy")]
        public Input<string>? RegularBackupStrategy { get; set; }

        public ConfigBackupStrategyState()
        {
        }
    }
}
