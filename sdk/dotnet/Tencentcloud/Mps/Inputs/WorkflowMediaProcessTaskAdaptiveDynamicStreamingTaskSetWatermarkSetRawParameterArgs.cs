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

    public sealed class WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Origin position, currently only supports:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.Default: TopLeft.
        /// </summary>
        [Input("coordinateOrigin")]
        public Input<string>? CoordinateOrigin { get; set; }

        /// <summary>
        /// Image watermark template, when Type is image, this field is required. When Type is text, this field is invalid.
        /// </summary>
        [Input("imageTemplate")]
        public Input<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateArgs>? ImageTemplate { get; set; }

        /// <summary>
        /// Watermark type, optional value:image: image watermark.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats:When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default: 0px.
        /// </summary>
        [Input("xPos")]
        public Input<string>? XPos { get; set; }

        /// <summary>
        /// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats:When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default: 0px.
        /// </summary>
        [Input("yPos")]
        public Input<string>? YPos { get; set; }

        public WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterArgs()
        {
        }
    }
}
