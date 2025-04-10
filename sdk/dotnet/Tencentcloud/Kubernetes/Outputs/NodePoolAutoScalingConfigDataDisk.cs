// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class NodePoolAutoScalingConfigDataDisk
    {
        /// <summary>
        /// Indicates whether the disk remove after instance terminated. Default is `false`.
        /// </summary>
        public readonly bool? DeleteWithInstance;
        /// <summary>
        /// Volume of disk in GB. Default is `0`.
        /// </summary>
        public readonly int? DiskSize;
        /// <summary>
        /// Types of disk. Valid value: `CLOUD_PREMIUM` and `CLOUD_SSD`.
        /// </summary>
        public readonly string? DiskType;
        /// <summary>
        /// Specify whether to encrypt data disk, default: false. NOTE: Make sure the instance type is offering and the cam role `QcloudKMSAccessForCVMRole` was provided.
        /// </summary>
        public readonly bool? Encrypt;
        /// <summary>
        /// Data disk snapshot ID.
        /// </summary>
        public readonly string? SnapshotId;
        /// <summary>
        /// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD` and `data_size` &gt; 460GB.
        /// </summary>
        public readonly int? ThroughputPerformance;

        [OutputConstructor]
        private NodePoolAutoScalingConfigDataDisk(
            bool? deleteWithInstance,

            int? diskSize,

            string? diskType,

            bool? encrypt,

            string? snapshotId,

            int? throughputPerformance)
        {
            DeleteWithInstance = deleteWithInstance;
            DiskSize = diskSize;
            DiskType = diskType;
            Encrypt = encrypt;
            SnapshotId = snapshotId;
            ThroughputPerformance = throughputPerformance;
        }
    }
}
