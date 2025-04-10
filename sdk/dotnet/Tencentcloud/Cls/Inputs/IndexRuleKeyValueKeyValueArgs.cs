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

    public sealed class IndexRuleKeyValueKeyValueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Field index description information.
        /// </summary>
        [Input("value")]
        public Input<Inputs.IndexRuleKeyValueKeyValueValueArgs>? Value { get; set; }

        public IndexRuleKeyValueKeyValueArgs()
        {
        }
    }
}
