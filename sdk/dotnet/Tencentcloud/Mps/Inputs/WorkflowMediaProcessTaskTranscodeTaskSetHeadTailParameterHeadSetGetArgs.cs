// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid when Type is COS, this item is required, indicating media processing COS object information.
        /// </summary>
        [Input("cosInputInfo")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetCosInputInfoGetArgs>? CosInputInfo { get; set; }

        /// <summary>
        /// Enter the type of source object, which supports COS and URL.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Valid when Type is URL, this item is required, indicating media processing URL object information.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("urlInputInfo")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetUrlInputInfoGetArgs>? UrlInputInfo { get; set; }

        public WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs()
        {
        }
    }
}
