// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class ZoneSettingCacheFollowOriginArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to follow the origin server configuration.- `on`: Enable.- `off`: Disable. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public ZoneSettingCacheFollowOriginArgs()
        {
        }
    }
}
