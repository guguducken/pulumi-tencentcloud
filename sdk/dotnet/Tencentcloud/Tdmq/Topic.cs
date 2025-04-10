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
    /// <summary>
    /// Provide a resource to create a TDMQ topic.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Tdmq.Instance("foo", new Tencentcloud.Tdmq.InstanceArgs
    ///         {
    ///             ClusterName = "example",
    ///             Remark = "this is description.",
    ///         });
    ///         var barNamespace = new Tencentcloud.Tdmq.Namespace("barNamespace", new Tencentcloud.Tdmq.NamespaceArgs
    ///         {
    ///             ClusterId = foo.Id,
    ///             EnvironName = "example",
    ///             MsgTtl = 300,
    ///             Remark = "this is description.",
    ///         });
    ///         var barTopic = new Tencentcloud.Tdmq.Topic("barTopic", new Tencentcloud.Tdmq.TopicArgs
    ///         {
    ///             ClusterId = foo.Id,
    ///             EnvironId = barNamespace.Id,
    ///             Partitions = 6,
    ///             Remark = "this is description.",
    ///             TopicName = "example",
    ///             TopicType = 0,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Tdmq Topic can be imported, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tdmq/topic:Topic test topic_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tdmq/topic:Topic")]
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// The Dedicated Cluster Id.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Creation time of resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The name of tdmq namespace.
        /// </summary>
        [Output("environId")]
        public Output<string> EnvironId { get; private set; } = null!;

        /// <summary>
        /// The partitions of topic.
        /// </summary>
        [Output("partitions")]
        public Output<int> Partitions { get; private set; } = null!;

        /// <summary>
        /// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
        /// </summary>
        [Output("pulsarTopicType")]
        public Output<int> PulsarTopicType { get; private set; } = null!;

        /// <summary>
        /// Description of the namespace.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// The name of topic to be created.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;

        /// <summary>
        /// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
        /// </summary>
        [Output("topicType")]
        public Output<int> TopicType { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tdmq/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tdmq/topic:Topic", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Dedicated Cluster Id.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of tdmq namespace.
        /// </summary>
        [Input("environId", required: true)]
        public Input<string> EnvironId { get; set; } = null!;

        /// <summary>
        /// The partitions of topic.
        /// </summary>
        [Input("partitions", required: true)]
        public Input<int> Partitions { get; set; } = null!;

        /// <summary>
        /// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
        /// </summary>
        [Input("pulsarTopicType")]
        public Input<int>? PulsarTopicType { get; set; }

        /// <summary>
        /// Description of the namespace.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The name of topic to be created.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        /// <summary>
        /// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
        /// </summary>
        [Input("topicType")]
        public Input<int>? TopicType { get; set; }

        public TopicArgs()
        {
        }
    }

    public sealed class TopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Dedicated Cluster Id.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Creation time of resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The name of tdmq namespace.
        /// </summary>
        [Input("environId")]
        public Input<string>? EnvironId { get; set; }

        /// <summary>
        /// The partitions of topic.
        /// </summary>
        [Input("partitions")]
        public Input<int>? Partitions { get; set; }

        /// <summary>
        /// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
        /// </summary>
        [Input("pulsarTopicType")]
        public Input<int>? PulsarTopicType { get; set; }

        /// <summary>
        /// Description of the namespace.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The name of topic to be created.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        /// <summary>
        /// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
        /// </summary>
        [Input("topicType")]
        public Input<int>? TopicType { get; set; }

        public TopicState()
        {
        }
    }
}
