// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class DeployContainerGroupEnvArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// env param name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// value of env.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Kubernetes ValueFrom configuration. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("valueFrom")]
        public Input<Inputs.DeployContainerGroupEnvValueFromArgs>? ValueFrom { get; set; }

        public DeployContainerGroupEnvArgs()
        {
        }
    }
}
