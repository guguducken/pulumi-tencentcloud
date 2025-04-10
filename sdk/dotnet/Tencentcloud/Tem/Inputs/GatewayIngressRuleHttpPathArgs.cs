// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Inputs
{

    public sealed class GatewayIngressRuleHttpPathArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// backend payload.
        /// </summary>
        [Input("backend", required: true)]
        public Input<Inputs.GatewayIngressRuleHttpPathBackendArgs> Backend { get; set; } = null!;

        /// <summary>
        /// path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GatewayIngressRuleHttpPathArgs()
        {
        }
    }
}
