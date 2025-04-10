// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class GetProcedureTemplatesTemplateListMediaProcessTaskCoverBySnapshotTaskListWatermarkListResult
    {
        /// <summary>
        /// Video transcoding template ID.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// End time offset of blur in seconds. If this parameter is left empty or `0` is entered, the blur will exist till the last video frame; If this value is greater than `0` (e.g., n), the blur will exist till second n; If this value is smaller than `0` (e.g., -n), the blur will exist till second n before the last video frame.
        /// </summary>
        public readonly double? EndTimeOffset;
        /// <summary>
        /// Start time offset of blur in seconds. If this parameter is left empty or `0` is entered, the blur will appear upon the first video frame. If this parameter is left empty or `0` is entered, the blur will appear upon the first video frame; If this value is greater than `0` (e.g., n), the blur will appear at second n after the first video frame; If this value is smaller than `0` (e.g., -n), the blur will appear at second n before the last video frame.
        /// </summary>
        public readonly double? StartTimeOffset;
        public readonly string? SvgContent;
        public readonly string? TextContent;

        [OutputConstructor]
        private GetProcedureTemplatesTemplateListMediaProcessTaskCoverBySnapshotTaskListWatermarkListResult(
            string definition,

            double? endTimeOffset,

            double? startTimeOffset,

            string? svgContent,

            string? textContent)
        {
            Definition = definition;
            EndTimeOffset = endTimeOffset;
            StartTimeOffset = startTimeOffset;
            SvgContent = svgContent;
            TextContent = textContent;
        }
    }
}
