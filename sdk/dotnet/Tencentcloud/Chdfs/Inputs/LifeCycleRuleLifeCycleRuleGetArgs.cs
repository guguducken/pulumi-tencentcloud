// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Chdfs.Inputs
{

    public sealed class LifeCycleRuleLifeCycleRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("lifeCycleRuleId")]
        public Input<int>? LifeCycleRuleId { get; set; }

        /// <summary>
        /// rule name.
        /// </summary>
        [Input("lifeCycleRuleName")]
        public Input<string>? LifeCycleRuleName { get; set; }

        /// <summary>
        /// rule op path.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// rule status, 1:open, 2:close.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("transitions")]
        private InputList<Inputs.LifeCycleRuleLifeCycleRuleTransitionGetArgs>? _transitions;

        /// <summary>
        /// life cycle rule transition list.
        /// </summary>
        public InputList<Inputs.LifeCycleRuleLifeCycleRuleTransitionGetArgs> Transitions
        {
            get => _transitions ?? (_transitions = new InputList<Inputs.LifeCycleRuleLifeCycleRuleTransitionGetArgs>());
            set => _transitions = value;
        }

        public LifeCycleRuleLifeCycleRuleGetArgs()
        {
        }
    }
}
