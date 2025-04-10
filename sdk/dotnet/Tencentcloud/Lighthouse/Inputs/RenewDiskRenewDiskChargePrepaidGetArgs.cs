// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse.Inputs
{

    public sealed class RenewDiskRenewDiskChargePrepaidGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current instance expiration time. Such as 2018-01-01 00:00:00. Specifying this parameter can align the expiration time of the instance attached to the disk. One of this parameter and Period must be specified, and cannot be specified at the same time.
        /// </summary>
        [Input("curInstanceDeadline")]
        public Input<string>? CurInstanceDeadline { get; set; }

        /// <summary>
        /// Renewal period.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Automatic renewal falg. Value:NOTIFY_AND_AUTO_RENEW: Notice expires and auto-renews.NOTIFY_AND_MANUAL_RENEW: Notification expires without automatic renewal, users need to manually renew.DISABLE_NOTIFY_AND_AUTO_RENEW: No automatic renewal and no notification.Default: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the disk will be automatically renewed monthly when the account balance is sufficient.
        /// </summary>
        [Input("renewFlag")]
        public Input<string>? RenewFlag { get; set; }

        /// <summary>
        /// newly purchased unit. Default: m.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        public RenewDiskRenewDiskChargePrepaidGetArgs()
        {
        }
    }
}
