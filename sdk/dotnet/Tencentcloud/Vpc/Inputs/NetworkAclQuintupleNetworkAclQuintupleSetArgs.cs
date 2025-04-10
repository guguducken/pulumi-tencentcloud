// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Inputs
{

    public sealed class NetworkAclQuintupleNetworkAclQuintupleSetArgs : Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<Inputs.NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs>? _egresses;
        public InputList<Inputs.NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<Inputs.NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<Inputs.NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs>? _ingresses;
        public InputList<Inputs.NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<Inputs.NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs>());
            set => _ingresses = value;
        }

        public NetworkAclQuintupleNetworkAclQuintupleSetArgs()
        {
        }
    }
}
