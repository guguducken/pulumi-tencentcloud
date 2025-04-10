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
    public sealed class SecurityPolicyConfigAclConfigUserRule
    {
        /// <summary>
        /// Valid values: `monitor`, `drop`.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Conditions of the rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityPolicyConfigAclConfigUserRuleCondition> Conditions;
        /// <summary>
        /// Name of the custom response page.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// ID of the custom response page.
        /// </summary>
        public readonly int? PageId;
        /// <summary>
        /// Punish time, Valid value range: 0-2 days.
        /// </summary>
        public readonly int? PunishTime;
        /// <summary>
        /// Time unit of the punish time. Valid values: `second`, `minutes`, `hour`.
        /// </summary>
        public readonly string? PunishTimeUnit;
        /// <summary>
        /// Redirect target URL, must be an sub-domain from one of the account&amp;#39;s site.
        /// </summary>
        public readonly string? RedirectUrl;
        /// <summary>
        /// Response code to use when redirecting.
        /// </summary>
        public readonly int? ResponseCode;
        public readonly int? RuleId;
        /// <summary>
        /// Rule Name.
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// Priority of the rule. Valid value range: 1-100.
        /// </summary>
        public readonly int RulePriority;
        /// <summary>
        /// Status of the rule. Valid values: `on`, `off`, `hour`.
        /// </summary>
        public readonly string RuleStatus;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private SecurityPolicyConfigAclConfigUserRule(
            string action,

            ImmutableArray<Outputs.SecurityPolicyConfigAclConfigUserRuleCondition> conditions,

            string? name,

            int? pageId,

            int? punishTime,

            string? punishTimeUnit,

            string? redirectUrl,

            int? responseCode,

            int? ruleId,

            string ruleName,

            int rulePriority,

            string ruleStatus,

            string? updateTime)
        {
            Action = action;
            Conditions = conditions;
            Name = name;
            PageId = pageId;
            PunishTime = punishTime;
            PunishTimeUnit = punishTimeUnit;
            RedirectUrl = redirectUrl;
            ResponseCode = responseCode;
            RuleId = ruleId;
            RuleName = ruleName;
            RulePriority = rulePriority;
            RuleStatus = ruleStatus;
            UpdateTime = updateTime;
        }
    }
}
