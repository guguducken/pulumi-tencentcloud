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
    public sealed class SecurityPolicyConfigRateLimitConfig
    {
        /// <summary>
        /// Intelligent client filter.
        /// </summary>
        public readonly Outputs.SecurityPolicyConfigRateLimitConfigIntelligence? Intelligence;
        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string? Switch;
        /// <summary>
        /// Default Template. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.SecurityPolicyConfigRateLimitConfigTemplate? Template;
        /// <summary>
        /// Custom configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityPolicyConfigRateLimitConfigUserRule> UserRules;

        [OutputConstructor]
        private SecurityPolicyConfigRateLimitConfig(
            Outputs.SecurityPolicyConfigRateLimitConfigIntelligence? intelligence,

            string? @switch,

            Outputs.SecurityPolicyConfigRateLimitConfigTemplate? template,

            ImmutableArray<Outputs.SecurityPolicyConfigRateLimitConfigUserRule> userRules)
        {
            Intelligence = intelligence;
            Switch = @switch;
            Template = template;
            UserRules = userRules;
        }
    }
}
