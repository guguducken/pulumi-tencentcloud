// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class InstanceEngineRegionInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine node region.
        /// </summary>
        [Input("engineRegion", required: true)]
        public Input<string> EngineRegion { get; set; } = null!;

        /// <summary>
        /// The number of nodes allocated in this region.
        /// </summary>
        [Input("replica", required: true)]
        public Input<int> Replica { get; set; } = null!;

        [Input("vpcInfos", required: true)]
        private InputList<Inputs.InstanceEngineRegionInfoVpcInfoGetArgs>? _vpcInfos;

        /// <summary>
        /// Cluster network information.
        /// </summary>
        public InputList<Inputs.InstanceEngineRegionInfoVpcInfoGetArgs> VpcInfos
        {
            get => _vpcInfos ?? (_vpcInfos = new InputList<Inputs.InstanceEngineRegionInfoVpcInfoGetArgs>());
            set => _vpcInfos = value;
        }

        public InstanceEngineRegionInfoGetArgs()
        {
        }
    }
}
