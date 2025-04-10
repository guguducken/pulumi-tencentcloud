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
    public sealed class ZoneSettingCacheNoCache
    {
        /// <summary>
        /// Whether to cache the configuration.- `on`: Do not cache.- `off`: Cache. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private ZoneSettingCacheNoCache(string? @switch)
        {
            Switch = @switch;
        }
    }
}
