// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class GetProxyCustomWeightRuleResult
    {
        public readonly int LessThan;
        public readonly int Weight;

        [OutputConstructor]
        private GetProxyCustomWeightRuleResult(
            int lessThan,

            int weight)
        {
            LessThan = lessThan;
            Weight = weight;
        }
    }
}
