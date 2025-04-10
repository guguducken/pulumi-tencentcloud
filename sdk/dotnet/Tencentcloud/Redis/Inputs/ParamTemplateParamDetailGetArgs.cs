// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis.Inputs
{

    public sealed class ParamTemplateParamDetailGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current value.
        /// </summary>
        [Input("currentValue")]
        public Input<string>? CurrentValue { get; set; }

        /// <summary>
        /// Default value.
        /// </summary>
        [Input("default")]
        public Input<string>? Default { get; set; }

        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enumValues")]
        private InputList<string>? _enumValues;

        /// <summary>
        /// Enum values.
        /// </summary>
        public InputList<string> EnumValues
        {
            get => _enumValues ?? (_enumValues = new InputList<string>());
            set => _enumValues = value;
        }

        /// <summary>
        /// Maximum value.
        /// </summary>
        [Input("max")]
        public Input<string>? Max { get; set; }

        /// <summary>
        /// Minimum value.
        /// </summary>
        [Input("min")]
        public Input<string>? Min { get; set; }

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Indicates whether to reboot redis instance if modified.
        /// </summary>
        [Input("needReboot")]
        public Input<int>? NeedReboot { get; set; }

        /// <summary>
        /// Parameter type.
        /// </summary>
        [Input("paramType")]
        public Input<string>? ParamType { get; set; }

        public ParamTemplateParamDetailGetArgs()
        {
        }
    }
}
