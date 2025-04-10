// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyConfigRateLimitConfigTemplateDetail
    {
        /// <summary>
        /// Action to take.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Template ID. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Template Name. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// Period.
        /// </summary>
        public readonly int? Period;
        /// <summary>
        /// Punish time.
        /// </summary>
        public readonly int? PunishTime;
        /// <summary>
        /// Threshold.
        /// </summary>
        public readonly int? Threshold;

        [OutputConstructor]
        private SecurityPolicyConfigRateLimitConfigTemplateDetail(
            string? action,

            int? id,

            string? mode,

            int? period,

            int? punishTime,

            int? threshold)
        {
            Action = action;
            Id = id;
            Mode = mode;
            Period = period;
            PunishTime = punishTime;
            Threshold = threshold;
        }
    }
}
