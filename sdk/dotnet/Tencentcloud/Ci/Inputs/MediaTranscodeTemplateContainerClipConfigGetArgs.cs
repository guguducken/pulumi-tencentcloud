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

    public sealed class MediaTranscodeTemplateContainerClipConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fragmentation duration, default 5s.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        public MediaTranscodeTemplateContainerClipConfigGetArgs()
        {
        }
    }
}
