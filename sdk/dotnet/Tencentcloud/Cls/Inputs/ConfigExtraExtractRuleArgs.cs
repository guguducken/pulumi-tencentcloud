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

    public sealed class ConfigExtraExtractRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).
        /// </summary>
        [Input("backtracking")]
        public Input<int>? Backtracking { get; set; }

        /// <summary>
        /// First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.
        /// </summary>
        [Input("beginRegex")]
        public Input<string>? BeginRegex { get; set; }

        /// <summary>
        /// Delimiter for delimited log, which is valid only if log_type is delimiter_log.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("filterKeyRegexes")]
        private InputList<Inputs.ConfigExtraExtractRuleFilterKeyRegexArgs>? _filterKeyRegexes;

        /// <summary>
        /// Log keys to be filtered and the corresponding regex.
        /// </summary>
        public InputList<Inputs.ConfigExtraExtractRuleFilterKeyRegexArgs> FilterKeyRegexes
        {
            get => _filterKeyRegexes ?? (_filterKeyRegexes = new InputList<Inputs.ConfigExtraExtractRuleFilterKeyRegexArgs>());
            set => _filterKeyRegexes = value;
        }

        [Input("keys")]
        private InputList<string>? _keys;

        /// <summary>
        /// Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.
        /// </summary>
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        /// <summary>
        /// Full log matching rule, which is valid only if log_type is fullregex_log.
        /// </summary>
        [Input("logRegex")]
        public Input<string>? LogRegex { get; set; }

        /// <summary>
        /// Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        /// <summary>
        /// Time field key name. time_key and time_format must appear in pair.
        /// </summary>
        [Input("timeKey")]
        public Input<string>? TimeKey { get; set; }

        /// <summary>
        /// Unmatched log key.
        /// </summary>
        [Input("unMatchLogKey")]
        public Input<string>? UnMatchLogKey { get; set; }

        /// <summary>
        /// Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no.
        /// </summary>
        [Input("unMatchUpLoadSwitch")]
        public Input<bool>? UnMatchUpLoadSwitch { get; set; }

        public ConfigExtraExtractRuleArgs()
        {
        }
    }
}
