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

    public sealed class InstanceContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command to run.
        /// </summary>
        [Input("command")]
        public Input<string>? Command { get; set; }

        /// <summary>
        /// Container image address.
        /// </summary>
        [Input("containerImage")]
        public Input<string>? ContainerImage { get; set; }

        /// <summary>
        /// Container name.
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        [Input("envs")]
        private InputList<Inputs.InstanceContainerEnvArgs>? _envs;

        /// <summary>
        /// List of environment variables.
        /// </summary>
        public InputList<Inputs.InstanceContainerEnvArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.InstanceContainerEnvArgs>());
            set => _envs = value;
        }

        [Input("publishPorts")]
        private InputList<Inputs.InstanceContainerPublishPortArgs>? _publishPorts;

        /// <summary>
        /// List of mappings of container ports and host ports.
        /// </summary>
        public InputList<Inputs.InstanceContainerPublishPortArgs> PublishPorts
        {
            get => _publishPorts ?? (_publishPorts = new InputList<Inputs.InstanceContainerPublishPortArgs>());
            set => _publishPorts = value;
        }

        [Input("volumes")]
        private InputList<Inputs.InstanceContainerVolumeArgs>? _volumes;

        /// <summary>
        /// List of container mount volumes.
        /// </summary>
        public InputList<Inputs.InstanceContainerVolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.InstanceContainerVolumeArgs>());
            set => _volumes = value;
        }

        public InstanceContainerArgs()
        {
        }
    }
}
