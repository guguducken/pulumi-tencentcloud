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

    public sealed class WorkflowAiRecognitionTaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Video Intelligent Recognition Template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        public WorkflowAiRecognitionTaskArgs()
        {
        }
    }
}
