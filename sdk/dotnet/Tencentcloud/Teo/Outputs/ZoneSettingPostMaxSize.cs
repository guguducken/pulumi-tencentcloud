// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class ZoneSettingPostMaxSize
    {
        /// <summary>
        /// Maximum size. Value range: 1-500 MB. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? MaxSize;
        /// <summary>
        /// Specifies whether to enable custom setting of the maximum file size.- `on`: Enable. You can set a custom max size.- `off`: Disable. In this case, the max size defaults to 32 MB.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private ZoneSettingPostMaxSize(
            int? maxSize,

            string @switch)
        {
            MaxSize = maxSize;
            Switch = @switch;
        }
    }
}
