// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As.Inputs
{

    public sealed class LoadBalancerForwardLoadBalancerGetArgs : Pulumi.ResourceArgs
    {
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        [Input("locationId")]
        public Input<string>? LocationId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("targetAttributes", required: true)]
        private InputList<Inputs.LoadBalancerForwardLoadBalancerTargetAttributeGetArgs>? _targetAttributes;
        public InputList<Inputs.LoadBalancerForwardLoadBalancerTargetAttributeGetArgs> TargetAttributes
        {
            get => _targetAttributes ?? (_targetAttributes = new InputList<Inputs.LoadBalancerForwardLoadBalancerTargetAttributeGetArgs>());
            set => _targetAttributes = value;
        }

        public LoadBalancerForwardLoadBalancerGetArgs()
        {
        }
    }
}
