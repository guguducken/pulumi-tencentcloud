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

    public sealed class AlarmNoticeNoticeReceiverArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// end time allowed to receive messages.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// index.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        [Input("receiverChannels", required: true)]
        private InputList<string>? _receiverChannels;

        /// <summary>
        /// receiver channels, Email,Sms,WeChat or Phone.
        /// </summary>
        public InputList<string> ReceiverChannels
        {
            get => _receiverChannels ?? (_receiverChannels = new InputList<string>());
            set => _receiverChannels = value;
        }

        [Input("receiverIds", required: true)]
        private InputList<int>? _receiverIds;

        /// <summary>
        /// receiver id.
        /// </summary>
        public InputList<int> ReceiverIds
        {
            get => _receiverIds ?? (_receiverIds = new InputList<int>());
            set => _receiverIds = value;
        }

        /// <summary>
        /// receiver type, Uin or Group.
        /// </summary>
        [Input("receiverType", required: true)]
        public Input<string> ReceiverType { get; set; } = null!;

        /// <summary>
        /// start time allowed to receive messages.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public AlarmNoticeNoticeReceiverArgs()
        {
        }
    }
}
