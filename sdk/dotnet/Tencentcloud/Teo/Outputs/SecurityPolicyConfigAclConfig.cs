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
    public sealed class SecurityPolicyConfigAclConfig
    {
        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string Switch;
        /// <summary>
        /// Custom configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityPolicyConfigAclConfigUserRule> UserRules;

        [OutputConstructor]
        private SecurityPolicyConfigAclConfig(
            string @switch,

            ImmutableArray<Outputs.SecurityPolicyConfigAclConfigUserRule> userRules)
        {
            Switch = @switch;
            UserRules = userRules;
        }
    }
}
