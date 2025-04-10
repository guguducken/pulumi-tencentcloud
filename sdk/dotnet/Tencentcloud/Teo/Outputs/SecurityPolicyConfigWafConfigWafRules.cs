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
    public sealed class SecurityPolicyConfigWafConfigWafRules
    {
        /// <summary>
        /// Block mode rules list. See details in data source `waf_managed_rules`.
        /// </summary>
        public readonly ImmutableArray<int> BlockRuleIds;
        /// <summary>
        /// Observe rules list. See details in data source `waf_managed_rules`.
        /// </summary>
        public readonly ImmutableArray<int> ObserveRuleIds;
        /// <summary>
        /// Whether to host the rules&amp;#39; configuration.- `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private SecurityPolicyConfigWafConfigWafRules(
            ImmutableArray<int> blockRuleIds,

            ImmutableArray<int> observeRuleIds,

            string @switch)
        {
            BlockRuleIds = blockRuleIds;
            ObserveRuleIds = observeRuleIds;
            Switch = @switch;
        }
    }
}
