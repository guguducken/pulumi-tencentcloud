// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpn.Inputs
{

    public sealed class CustomerGatewayConfigurationDownloadCustomerGatewayVendorGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Platform.
        /// </summary>
        [Input("platform", required: true)]
        public Input<string> Platform { get; set; } = null!;

        /// <summary>
        /// SoftwareVersion.
        /// </summary>
        [Input("softwareVersion", required: true)]
        public Input<string> SoftwareVersion { get; set; } = null!;

        /// <summary>
        /// VendorName.
        /// </summary>
        [Input("vendorName", required: true)]
        public Input<string> VendorName { get; set; } = null!;

        public CustomerGatewayConfigurationDownloadCustomerGatewayVendorGetArgs()
        {
        }
    }
}
