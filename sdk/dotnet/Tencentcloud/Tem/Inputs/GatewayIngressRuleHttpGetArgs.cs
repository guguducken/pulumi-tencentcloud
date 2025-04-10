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

    public sealed class GatewayIngressRuleHttpGetArgs : Pulumi.ResourceArgs
    {
        [Input("paths", required: true)]
        private InputList<Inputs.GatewayIngressRuleHttpPathGetArgs>? _paths;

        /// <summary>
        /// path payload.
        /// </summary>
        public InputList<Inputs.GatewayIngressRuleHttpPathGetArgs> Paths
        {
            get => _paths ?? (_paths = new InputList<Inputs.GatewayIngressRuleHttpPathGetArgs>());
            set => _paths = value;
        }

        public GatewayIngressRuleHttpGetArgs()
        {
        }
    }
}
