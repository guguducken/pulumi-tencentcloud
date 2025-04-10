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
    public sealed class SecurityPolicyConfigIpTableConfigRule
    {
        /// <summary>
        /// Actions to take. Valid values: `drop`, `trans`, `monitor`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Matching content.
        /// </summary>
        public readonly string? MatchContent;
        /// <summary>
        /// Matching type. Valid values: `ip`, `area`.
        /// </summary>
        public readonly string? MatchFrom;
        public readonly int? RuleId;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private SecurityPolicyConfigIpTableConfigRule(
            string? action,

            string? matchContent,

            string? matchFrom,

            int? ruleId,

            string? updateTime)
        {
            Action = action;
            MatchContent = matchContent;
            MatchFrom = matchFrom;
            RuleId = ruleId;
            UpdateTime = updateTime;
        }
    }
}
