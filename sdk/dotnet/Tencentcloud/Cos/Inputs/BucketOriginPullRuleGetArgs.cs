// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketOriginPullRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("customHttpHeaders")]
        private InputMap<object>? _customHttpHeaders;

        /// <summary>
        /// Specifies the custom headers that you can add for COS to access your origin server.
        /// </summary>
        public InputMap<object> CustomHttpHeaders
        {
            get => _customHttpHeaders ?? (_customHttpHeaders = new InputMap<object>());
            set => _customHttpHeaders = value;
        }

        [Input("followHttpHeaders")]
        private InputList<string>? _followHttpHeaders;

        /// <summary>
        /// Specifies the pass through headers when accessing the origin server.
        /// </summary>
        public InputList<string> FollowHttpHeaders
        {
            get => _followHttpHeaders ?? (_followHttpHeaders = new InputList<string>());
            set => _followHttpHeaders = value;
        }

        /// <summary>
        /// Specifies whether to pass through COS request query string when accessing the origin server.
        /// </summary>
        [Input("followQueryString")]
        public Input<bool>? FollowQueryString { get; set; }

        /// <summary>
        /// Specifies whether to follow 3XX redirect to another origin server to pull data from.
        /// </summary>
        [Input("followRedirection")]
        public Input<bool>? FollowRedirection { get; set; }

        /// <summary>
        /// Allows only a domain name or IP address. You can optionally append a port number to the address.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Triggers the origin-pull rule when the requested file name matches this prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Priority of origin-pull rules, do not set the same value for multiple rules.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// the protocol used for COS to access the specified origin server. The available value include `HTTP`, `HTTPS` and `FOLLOW`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// If `true`, COS will not return 3XX status code when pulling data from an origin server. Current available zone: ap-beijing, ap-shanghai, ap-singapore, ap-mumbai.
        /// </summary>
        [Input("syncBackToSource")]
        public Input<bool>? SyncBackToSource { get; set; }

        public BucketOriginPullRuleGetArgs()
        {
        }
    }
}
