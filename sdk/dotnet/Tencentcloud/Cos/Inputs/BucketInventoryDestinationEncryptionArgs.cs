// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketInventoryDestinationEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption with COS-managed key. This field can be left empty.
        /// </summary>
        [Input("sseCos")]
        public Input<string>? SseCos { get; set; }

        public BucketInventoryDestinationEncryptionArgs()
        {
        }
    }
}
