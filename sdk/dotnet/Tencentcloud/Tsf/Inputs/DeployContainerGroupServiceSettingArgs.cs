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

    public sealed class DeployContainerGroupServiceSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// 0: Public network, 1: Access within the cluster, 2: NodePort, 3: Access within VPC. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("accessType", required: true)]
        public Input<int> AccessType { get; set; } = null!;

        /// <summary>
        /// When set to true and DisableService is also true, the previously created service will be deleted. Please use with caution. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("allowDeleteService")]
        public Input<bool>? AllowDeleteService { get; set; }

        /// <summary>
        /// Whether to create a Kubernetes service. The default value is false. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("disableService")]
        public Input<bool>? DisableService { get; set; }

        /// <summary>
        /// Whether the service is of headless type. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("headlessService")]
        public Input<bool>? HeadlessService { get; set; }

        /// <summary>
        /// Enable session affinity. true means enabled, false means disabled. The default value is false. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("openSessionAffinity")]
        public Input<bool>? OpenSessionAffinity { get; set; }

        [Input("protocolPorts", required: true)]
        private InputList<Inputs.DeployContainerGroupServiceSettingProtocolPortArgs>? _protocolPorts;

        /// <summary>
        /// Container port mapping. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<Inputs.DeployContainerGroupServiceSettingProtocolPortArgs> ProtocolPorts
        {
            get => _protocolPorts ?? (_protocolPorts = new InputList<Inputs.DeployContainerGroupServiceSettingProtocolPortArgs>());
            set => _protocolPorts = value;
        }

        /// <summary>
        /// Session affinity session time. The default value is 10800. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("sessionAffinityTimeoutSeconds")]
        public Input<int>? SessionAffinityTimeoutSeconds { get; set; }

        /// <summary>
        /// subnet Id.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public DeployContainerGroupServiceSettingArgs()
        {
        }
    }
}
