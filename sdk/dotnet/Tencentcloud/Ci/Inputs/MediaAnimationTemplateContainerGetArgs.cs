// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Inputs
{

    public sealed class MediaAnimationTemplateContainerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Package format.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public MediaAnimationTemplateContainerGetArgs()
        {
        }
    }
}
