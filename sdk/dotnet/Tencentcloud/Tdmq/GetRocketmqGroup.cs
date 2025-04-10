// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq
{
    public static class GetRocketmqGroup
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq group
        /// </summary>
        public static Task<GetRocketmqGroupResult> InvokeAsync(GetRocketmqGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRocketmqGroupResult>("tencentcloud:Tdmq/getRocketmqGroup:getRocketmqGroup", args ?? new GetRocketmqGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq group
        /// </summary>
        public static Output<GetRocketmqGroupResult> Invoke(GetRocketmqGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRocketmqGroupResult>("tencentcloud:Tdmq/getRocketmqGroup:getRocketmqGroup", args ?? new GetRocketmqGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRocketmqGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Consumer group query by consumer group name. Fuzzy query is supported.
        /// </summary>
        [Input("filterGroup")]
        public string? FilterGroup { get; set; }

        /// <summary>
        /// Subscription group name. After it is specified, the information of only this subscription group will be returned.
        /// </summary>
        [Input("filterOneGroup")]
        public string? FilterOneGroup { get; set; }

        /// <summary>
        /// Topic name, which can be used to query all subscription groups under the topic.
        /// </summary>
        [Input("filterTopic")]
        public string? FilterTopic { get; set; }

        /// <summary>
        /// Namespace.
        /// </summary>
        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRocketmqGroupArgs()
        {
        }
    }

    public sealed class GetRocketmqGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Consumer group query by consumer group name. Fuzzy query is supported.
        /// </summary>
        [Input("filterGroup")]
        public Input<string>? FilterGroup { get; set; }

        /// <summary>
        /// Subscription group name. After it is specified, the information of only this subscription group will be returned.
        /// </summary>
        [Input("filterOneGroup")]
        public Input<string>? FilterOneGroup { get; set; }

        /// <summary>
        /// Topic name, which can be used to query all subscription groups under the topic.
        /// </summary>
        [Input("filterTopic")]
        public Input<string>? FilterTopic { get; set; }

        /// <summary>
        /// Namespace.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRocketmqGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRocketmqGroupResult
    {
        public readonly string ClusterId;
        public readonly string? FilterGroup;
        public readonly string? FilterOneGroup;
        public readonly string? FilterTopic;
        /// <summary>
        /// List of subscription groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqGroupGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string NamespaceId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRocketmqGroupResult(
            string clusterId,

            string? filterGroup,

            string? filterOneGroup,

            string? filterTopic,

            ImmutableArray<Outputs.GetRocketmqGroupGroupResult> groups,

            string id,

            string namespaceId,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            FilterGroup = filterGroup;
            FilterOneGroup = filterOneGroup;
            FilterTopic = filterTopic;
            Groups = groups;
            Id = id;
            NamespaceId = namespaceId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
