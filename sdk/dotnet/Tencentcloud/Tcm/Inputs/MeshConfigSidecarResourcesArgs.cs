// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Inputs
{

    public sealed class MeshConfigSidecarResourcesArgs : Pulumi.ResourceArgs
    {
        [Input("limits")]
        private InputList<Inputs.MeshConfigSidecarResourcesLimitArgs>? _limits;

        /// <summary>
        /// Sidecar limits.
        /// </summary>
        public InputList<Inputs.MeshConfigSidecarResourcesLimitArgs> Limits
        {
            get => _limits ?? (_limits = new InputList<Inputs.MeshConfigSidecarResourcesLimitArgs>());
            set => _limits = value;
        }

        [Input("requests")]
        private InputList<Inputs.MeshConfigSidecarResourcesRequestArgs>? _requests;

        /// <summary>
        /// Sidecar requests.
        /// </summary>
        public InputList<Inputs.MeshConfigSidecarResourcesRequestArgs> Requests
        {
            get => _requests ?? (_requests = new InputList<Inputs.MeshConfigSidecarResourcesRequestArgs>());
            set => _requests = value;
        }

        public MeshConfigSidecarResourcesArgs()
        {
        }
    }
}
