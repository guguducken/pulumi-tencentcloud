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

    public sealed class IndexRuleKeyValueGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Case sensitivity.
        /// </summary>
        [Input("caseSensitive", required: true)]
        public Input<bool> CaseSensitive { get; set; } = null!;

        [Input("keyValues")]
        private InputList<Inputs.IndexRuleKeyValueKeyValueGetArgs>? _keyValues;

        /// <summary>
        /// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
        /// </summary>
        public InputList<Inputs.IndexRuleKeyValueKeyValueGetArgs> KeyValues
        {
            get => _keyValues ?? (_keyValues = new InputList<Inputs.IndexRuleKeyValueKeyValueGetArgs>());
            set => _keyValues = value;
        }

        public IndexRuleKeyValueGetArgs()
        {
        }
    }
}
