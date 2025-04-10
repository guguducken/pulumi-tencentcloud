// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu.Inputs
{

    public sealed class CcPolicyV2CcBlackWhiteIpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Blacklist and whitelist IP addresses.
        /// </summary>
        [Input("blackWhiteIp", required: true)]
        public Input<string> BlackWhiteIp { get; set; } = null!;

        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Modify time.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        /// <summary>
        /// Protocol.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// IP type, value [black(blacklist IP), white (whitelist IP)].
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CcPolicyV2CcBlackWhiteIpArgs()
        {
        }
    }
}
