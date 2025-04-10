// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcx.Inputs
{

    public sealed class ExtraConfigNqaInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// detect ip.
        /// </summary>
        [Input("destinationIp")]
        public Input<string>? DestinationIp { get; set; }

        /// <summary>
        /// detect interval.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// detect times.
        /// </summary>
        [Input("probeFailedTimes")]
        public Input<int>? ProbeFailedTimes { get; set; }

        public ExtraConfigNqaInfoGetArgs()
        {
        }
    }
}
