// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Inputs
{

    public sealed class SyncConfigOptionsConflictHandleOptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Columns covered by the condition. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("conditionColumn")]
        public Input<string>? ConditionColumn { get; set; }

        /// <summary>
        /// Conditional Override Operation. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("conditionOperator")]
        public Input<string>? ConditionOperator { get; set; }

        /// <summary>
        /// Conditional Override Priority Processing. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("conditionOrderInSrcAndDst")]
        public Input<string>? ConditionOrderInSrcAndDst { get; set; }

        public SyncConfigOptionsConflictHandleOptionGetArgs()
        {
        }
    }
}
