// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class ConfigExtractRuleFilterKeyRegexArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Log key to be filtered.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Filter rule regex corresponding to key.
        /// </summary>
        [Input("regex")]
        public Input<string>? Regex { get; set; }

        public ConfigExtractRuleFilterKeyRegexArgs()
        {
        }
    }
}
