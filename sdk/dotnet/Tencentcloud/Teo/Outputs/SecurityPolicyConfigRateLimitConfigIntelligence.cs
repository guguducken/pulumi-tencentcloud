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
    public sealed class SecurityPolicyConfigRateLimitConfigIntelligence
    {
        /// <summary>
        /// Action to take. Valid values: `monitor`, `alg`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private SecurityPolicyConfigRateLimitConfigIntelligence(
            string? action,

            string? @switch)
        {
            Action = action;
            Switch = @switch;
        }
    }
}
