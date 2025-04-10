// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As
{
    /// <summary>
    /// Provides a resource to create a configuration for an AS (Auto scaling) instance.
    /// 
    /// &gt; **NOTE:**  In order to ensure the integrity of customer data, if the cvm instance was destroyed due to shrinking, it will keep the cbs associate with cvm by default. If you want to destroy together, please set `delete_with_instance` to `true`.
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
    ///         var launchConfiguration = new Tencentcloud.As.ScalingConfig("launchConfiguration", new Tencentcloud.As.ScalingConfigArgs
    ///         {
    ///             ConfigurationName = "launch-configuration",
    ///             DataDisks = 
    ///             {
    ///                 new Tencentcloud.As.Inputs.ScalingConfigDataDiskArgs
    ///                 {
    ///                     DiskSize = 50,
    ///                     DiskType = "CLOUD_PREMIUM",
    ///                 },
    ///             },
    ///             EnhancedMonitorService = false,
    ///             EnhancedSecurityService = false,
    ///             ImageId = "img-9qabwvbn",
    ///             InstanceTags = 
    ///             {
    ///                 { "tag", "as" },
    ///             },
    ///             InstanceTypes = 
    ///             {
    ///                 "SA1.SMALL1",
    ///             },
    ///             InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///             InternetMaxBandwidthOut = 10,
    ///             Password = "test123#",
    ///             ProjectId = 0,
    ///             PublicIpAssigned = true,
    ///             SystemDiskSize = 50,
    ///             SystemDiskType = "CLOUD_PREMIUM",
    ///             UserData = "dGVzdA==",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Using SPOT charge type
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var launchConfiguration = new Tencentcloud.As.ScalingConfig("launchConfiguration", new Tencentcloud.As.ScalingConfigArgs
    ///         {
    ///             ConfigurationName = "launch-configuration",
    ///             ImageId = "img-9qabwvbn",
    ///             InstanceChargeType = "SPOTPAID",
    ///             InstanceTypes = 
    ///             {
    ///                 "SA1.SMALL1",
    ///             },
    ///             SpotInstanceType = "one-time",
    ///             SpotMaxPrice = "1000",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AutoScaling Configuration can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:As/scalingConfig:ScalingConfig scaling_config asc-n32ymck2
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:As/scalingConfig:ScalingConfig")]
    public partial class ScalingConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// CAM role name authorized to access.
        /// </summary>
        [Output("camRoleName")]
        public Output<string?> CamRoleName { get; private set; } = null!;

        /// <summary>
        /// Name of a launch configuration.
        /// </summary>
        [Output("configurationName")]
        public Output<string> ConfigurationName { get; private set; } = null!;

        /// <summary>
        /// The time when the launch configuration was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        [Output("dataDisks")]
        public Output<ImmutableArray<Outputs.ScalingConfigDataDisk>> DataDisks { get; private set; } = null!;

        /// <summary>
        /// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
        /// </summary>
        [Output("diskTypePolicy")]
        public Output<string?> DiskTypePolicy { get; private set; } = null!;

        /// <summary>
        /// To specify whether to enable cloud monitor service. Default is `TRUE`.
        /// </summary>
        [Output("enhancedMonitorService")]
        public Output<bool?> EnhancedMonitorService { get; private set; } = null!;

        /// <summary>
        /// To specify whether to enable cloud security service. Default is `TRUE`.
        /// </summary>
        [Output("enhancedSecurityService")]
        public Output<bool?> EnhancedSecurityService { get; private set; } = null!;

        /// <summary>
        /// An available image ID for a cvm instance.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        /// </summary>
        [Output("instanceChargeTypePrepaidPeriod")]
        public Output<int?> InstanceChargeTypePrepaidPeriod { get; private set; } = null!;

        /// <summary>
        /// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
        /// </summary>
        [Output("instanceChargeTypePrepaidRenewFlag")]
        public Output<string> InstanceChargeTypePrepaidRenewFlag { get; private set; } = null!;

        /// <summary>
        /// Settings of CVM instance names.
        /// </summary>
        [Output("instanceNameSettings")]
        public Output<Outputs.ScalingConfigInstanceNameSettings?> InstanceNameSettings { get; private set; } = null!;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        [Output("instanceTags")]
        public Output<ImmutableDictionary<string, object>?> InstanceTags { get; private set; } = null!;

        /// <summary>
        /// Specified types of CVM instances.
        /// </summary>
        [Output("instanceTypes")]
        public Output<ImmutableArray<string>> InstanceTypes { get; private set; } = null!;

        /// <summary>
        /// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
        /// </summary>
        [Output("internetChargeType")]
        public Output<string?> InternetChargeType { get; private set; } = null!;

        /// <summary>
        /// Max bandwidth of Internet access in Mbps. Default is `0`.
        /// </summary>
        [Output("internetMaxBandwidthOut")]
        public Output<int?> InternetMaxBandwidthOut { get; private set; } = null!;

        /// <summary>
        /// Specify whether to keep original settings of a CVM image. And it can't be used with password or key_ids together.
        /// </summary>
        [Output("keepImageLogin")]
        public Output<bool?> KeepImageLogin { get; private set; } = null!;

        /// <summary>
        /// ID list of keys.
        /// </summary>
        [Output("keyIds")]
        public Output<ImmutableArray<string>> KeyIds { get; private set; } = null!;

        /// <summary>
        /// Password to access.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Specifys to which project the configuration belongs.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Specify whether to assign an Internet IP address.
        /// </summary>
        [Output("publicIpAssigned")]
        public Output<bool?> PublicIpAssigned { get; private set; } = null!;

        /// <summary>
        /// Security groups to which a CVM instance belongs.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Type of spot instance, only support `one-time` now. Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Output("spotInstanceType")]
        public Output<string?> SpotInstanceType { get; private set; } = null!;

        /// <summary>
        /// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Output("spotMaxPrice")]
        public Output<string?> SpotMaxPrice { get; private set; } = null!;

        /// <summary>
        /// Current statues of a launch configuration.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Volume of system disk in GB. Default is `50`.
        /// </summary>
        [Output("systemDiskSize")]
        public Output<int?> SystemDiskSize { get; private set; } = null!;

        /// <summary>
        /// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when disk_type_policy is ORIGINAL.
        /// </summary>
        [Output("systemDiskType")]
        public Output<string?> SystemDiskType { get; private set; } = null!;

        /// <summary>
        /// ase64-encoded User Data text, the length limit is 16KB.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;


        /// <summary>
        /// Create a ScalingConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScalingConfig(string name, ScalingConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:As/scalingConfig:ScalingConfig", name, args ?? new ScalingConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScalingConfig(string name, Input<string> id, ScalingConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:As/scalingConfig:ScalingConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScalingConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScalingConfig Get(string name, Input<string> id, ScalingConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ScalingConfig(name, id, state, options);
        }
    }

    public sealed class ScalingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CAM role name authorized to access.
        /// </summary>
        [Input("camRoleName")]
        public Input<string>? CamRoleName { get; set; }

        /// <summary>
        /// Name of a launch configuration.
        /// </summary>
        [Input("configurationName", required: true)]
        public Input<string> ConfigurationName { get; set; } = null!;

        [Input("dataDisks")]
        private InputList<Inputs.ScalingConfigDataDiskArgs>? _dataDisks;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public InputList<Inputs.ScalingConfigDataDiskArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.ScalingConfigDataDiskArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
        /// </summary>
        [Input("diskTypePolicy")]
        public Input<string>? DiskTypePolicy { get; set; }

        /// <summary>
        /// To specify whether to enable cloud monitor service. Default is `TRUE`.
        /// </summary>
        [Input("enhancedMonitorService")]
        public Input<bool>? EnhancedMonitorService { get; set; }

        /// <summary>
        /// To specify whether to enable cloud security service. Default is `TRUE`.
        /// </summary>
        [Input("enhancedSecurityService")]
        public Input<bool>? EnhancedSecurityService { get; set; }

        /// <summary>
        /// An available image ID for a cvm instance.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        /// </summary>
        [Input("instanceChargeTypePrepaidPeriod")]
        public Input<int>? InstanceChargeTypePrepaidPeriod { get; set; }

        /// <summary>
        /// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
        /// </summary>
        [Input("instanceChargeTypePrepaidRenewFlag")]
        public Input<string>? InstanceChargeTypePrepaidRenewFlag { get; set; }

        /// <summary>
        /// Settings of CVM instance names.
        /// </summary>
        [Input("instanceNameSettings")]
        public Input<Inputs.ScalingConfigInstanceNameSettingsArgs>? InstanceNameSettings { get; set; }

        [Input("instanceTags")]
        private InputMap<object>? _instanceTags;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        public InputMap<object> InstanceTags
        {
            get => _instanceTags ?? (_instanceTags = new InputMap<object>());
            set => _instanceTags = value;
        }

        [Input("instanceTypes", required: true)]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// Specified types of CVM instances.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        /// <summary>
        /// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// Max bandwidth of Internet access in Mbps. Default is `0`.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// Specify whether to keep original settings of a CVM image. And it can't be used with password or key_ids together.
        /// </summary>
        [Input("keepImageLogin")]
        public Input<bool>? KeepImageLogin { get; set; }

        [Input("keyIds")]
        private InputList<string>? _keyIds;

        /// <summary>
        /// ID list of keys.
        /// </summary>
        public InputList<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new InputList<string>());
            set => _keyIds = value;
        }

        /// <summary>
        /// Password to access.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Specifys to which project the configuration belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Specify whether to assign an Internet IP address.
        /// </summary>
        [Input("publicIpAssigned")]
        public Input<bool>? PublicIpAssigned { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security groups to which a CVM instance belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Type of spot instance, only support `one-time` now. Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        /// <summary>
        /// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Input("spotMaxPrice")]
        public Input<string>? SpotMaxPrice { get; set; }

        /// <summary>
        /// Volume of system disk in GB. Default is `50`.
        /// </summary>
        [Input("systemDiskSize")]
        public Input<int>? SystemDiskSize { get; set; }

        /// <summary>
        /// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when disk_type_policy is ORIGINAL.
        /// </summary>
        [Input("systemDiskType")]
        public Input<string>? SystemDiskType { get; set; }

        /// <summary>
        /// ase64-encoded User Data text, the length limit is 16KB.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public ScalingConfigArgs()
        {
        }
    }

    public sealed class ScalingConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CAM role name authorized to access.
        /// </summary>
        [Input("camRoleName")]
        public Input<string>? CamRoleName { get; set; }

        /// <summary>
        /// Name of a launch configuration.
        /// </summary>
        [Input("configurationName")]
        public Input<string>? ConfigurationName { get; set; }

        /// <summary>
        /// The time when the launch configuration was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.ScalingConfigDataDiskGetArgs>? _dataDisks;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public InputList<Inputs.ScalingConfigDataDiskGetArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.ScalingConfigDataDiskGetArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
        /// </summary>
        [Input("diskTypePolicy")]
        public Input<string>? DiskTypePolicy { get; set; }

        /// <summary>
        /// To specify whether to enable cloud monitor service. Default is `TRUE`.
        /// </summary>
        [Input("enhancedMonitorService")]
        public Input<bool>? EnhancedMonitorService { get; set; }

        /// <summary>
        /// To specify whether to enable cloud security service. Default is `TRUE`.
        /// </summary>
        [Input("enhancedSecurityService")]
        public Input<bool>? EnhancedSecurityService { get; set; }

        /// <summary>
        /// An available image ID for a cvm instance.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        /// </summary>
        [Input("instanceChargeTypePrepaidPeriod")]
        public Input<int>? InstanceChargeTypePrepaidPeriod { get; set; }

        /// <summary>
        /// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
        /// </summary>
        [Input("instanceChargeTypePrepaidRenewFlag")]
        public Input<string>? InstanceChargeTypePrepaidRenewFlag { get; set; }

        /// <summary>
        /// Settings of CVM instance names.
        /// </summary>
        [Input("instanceNameSettings")]
        public Input<Inputs.ScalingConfigInstanceNameSettingsGetArgs>? InstanceNameSettings { get; set; }

        [Input("instanceTags")]
        private InputMap<object>? _instanceTags;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        public InputMap<object> InstanceTags
        {
            get => _instanceTags ?? (_instanceTags = new InputMap<object>());
            set => _instanceTags = value;
        }

        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// Specified types of CVM instances.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        /// <summary>
        /// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// Max bandwidth of Internet access in Mbps. Default is `0`.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// Specify whether to keep original settings of a CVM image. And it can't be used with password or key_ids together.
        /// </summary>
        [Input("keepImageLogin")]
        public Input<bool>? KeepImageLogin { get; set; }

        [Input("keyIds")]
        private InputList<string>? _keyIds;

        /// <summary>
        /// ID list of keys.
        /// </summary>
        public InputList<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new InputList<string>());
            set => _keyIds = value;
        }

        /// <summary>
        /// Password to access.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Specifys to which project the configuration belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Specify whether to assign an Internet IP address.
        /// </summary>
        [Input("publicIpAssigned")]
        public Input<bool>? PublicIpAssigned { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security groups to which a CVM instance belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Type of spot instance, only support `one-time` now. Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        /// <summary>
        /// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Input("spotMaxPrice")]
        public Input<string>? SpotMaxPrice { get; set; }

        /// <summary>
        /// Current statues of a launch configuration.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Volume of system disk in GB. Default is `50`.
        /// </summary>
        [Input("systemDiskSize")]
        public Input<int>? SystemDiskSize { get; set; }

        /// <summary>
        /// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when disk_type_policy is ORIGINAL.
        /// </summary>
        [Input("systemDiskType")]
        public Input<string>? SystemDiskType { get; set; }

        /// <summary>
        /// ase64-encoded User Data text, the length limit is 16KB.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public ScalingConfigState()
        {
        }
    }
}
