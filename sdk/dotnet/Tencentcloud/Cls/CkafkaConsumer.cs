// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls
{
    /// <summary>
    /// Provides a resource to create a cls ckafka_consumer
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
    ///         var ckafkaConsumer = new Tencentcloud.Cls.CkafkaConsumer("ckafkaConsumer", new Tencentcloud.Cls.CkafkaConsumerArgs
    ///         {
    ///             Ckafka = new Tencentcloud.Cls.Inputs.CkafkaConsumerCkafkaArgs
    ///             {
    ///                 InstanceId = "ckafka-qzoeaqx8",
    ///                 InstanceName = "ckafka-instance",
    ///                 TopicId = "topic-c6tm4kpm",
    ///                 TopicName = "name",
    ///                 Vip = "172.16.112.23",
    ///                 Vport = "9092",
    ///             },
    ///             Compression = 1,
    ///             Content = new Tencentcloud.Cls.Inputs.CkafkaConsumerContentArgs
    ///             {
    ///                 EnableTag = true,
    ///                 MetaFields = 
    ///                 {
    ///                     "__FILENAME__",
    ///                     "__HOSTNAME__",
    ///                     "__PKGID__",
    ///                     "__SOURCE__",
    ///                     "__TIMESTAMP__",
    ///                 },
    ///                 TagJsonNotTiled = true,
    ///                 TimestampAccuracy = 2,
    ///             },
    ///             NeedContent = true,
    ///             TopicId = "7e34a3a7-635e-4da8-9005-88106c1fde69",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cls ckafka_consumer can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cls/ckafkaConsumer:CkafkaConsumer ckafka_consumer topic_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cls/ckafkaConsumer:CkafkaConsumer")]
    public partial class CkafkaConsumer : Pulumi.CustomResource
    {
        /// <summary>
        /// ckafka info.
        /// </summary>
        [Output("ckafka")]
        public Output<Outputs.CkafkaConsumerCkafka?> Ckafka { get; private set; } = null!;

        /// <summary>
        /// compression method. 0 for NONE, 2 for SNAPPY, 3 for LZ4.
        /// </summary>
        [Output("compression")]
        public Output<int?> Compression { get; private set; } = null!;

        /// <summary>
        /// metadata information.
        /// </summary>
        [Output("content")]
        public Output<Outputs.CkafkaConsumerContent?> Content { get; private set; } = null!;

        /// <summary>
        /// whether to deliver the metadata information of the log.
        /// </summary>
        [Output("needContent")]
        public Output<bool?> NeedContent { get; private set; } = null!;

        /// <summary>
        /// topic id.
        /// </summary>
        [Output("topicId")]
        public Output<string> TopicId { get; private set; } = null!;


        /// <summary>
        /// Create a CkafkaConsumer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CkafkaConsumer(string name, CkafkaConsumerArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/ckafkaConsumer:CkafkaConsumer", name, args ?? new CkafkaConsumerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CkafkaConsumer(string name, Input<string> id, CkafkaConsumerState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/ckafkaConsumer:CkafkaConsumer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CkafkaConsumer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CkafkaConsumer Get(string name, Input<string> id, CkafkaConsumerState? state = null, CustomResourceOptions? options = null)
        {
            return new CkafkaConsumer(name, id, state, options);
        }
    }

    public sealed class CkafkaConsumerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ckafka info.
        /// </summary>
        [Input("ckafka")]
        public Input<Inputs.CkafkaConsumerCkafkaArgs>? Ckafka { get; set; }

        /// <summary>
        /// compression method. 0 for NONE, 2 for SNAPPY, 3 for LZ4.
        /// </summary>
        [Input("compression")]
        public Input<int>? Compression { get; set; }

        /// <summary>
        /// metadata information.
        /// </summary>
        [Input("content")]
        public Input<Inputs.CkafkaConsumerContentArgs>? Content { get; set; }

        /// <summary>
        /// whether to deliver the metadata information of the log.
        /// </summary>
        [Input("needContent")]
        public Input<bool>? NeedContent { get; set; }

        /// <summary>
        /// topic id.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public CkafkaConsumerArgs()
        {
        }
    }

    public sealed class CkafkaConsumerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ckafka info.
        /// </summary>
        [Input("ckafka")]
        public Input<Inputs.CkafkaConsumerCkafkaGetArgs>? Ckafka { get; set; }

        /// <summary>
        /// compression method. 0 for NONE, 2 for SNAPPY, 3 for LZ4.
        /// </summary>
        [Input("compression")]
        public Input<int>? Compression { get; set; }

        /// <summary>
        /// metadata information.
        /// </summary>
        [Input("content")]
        public Input<Inputs.CkafkaConsumerContentGetArgs>? Content { get; set; }

        /// <summary>
        /// whether to deliver the metadata information of the log.
        /// </summary>
        [Input("needContent")]
        public Input<bool>? NeedContent { get; set; }

        /// <summary>
        /// topic id.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public CkafkaConsumerState()
        {
        }
    }
}
