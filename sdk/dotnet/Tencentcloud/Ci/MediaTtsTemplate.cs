// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci
{
    /// <summary>
    /// Provides a resource to create a ci media_tts_template
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
    ///         var mediaTtsTemplate = new Tencentcloud.Ci.MediaTtsTemplate("mediaTtsTemplate", new Tencentcloud.Ci.MediaTtsTemplateArgs
    ///         {
    ///             Bucket = "terraform-ci-xxxxxx",
    ///             Codec = "pcm",
    ///             Mode = "Asyc",
    ///             Speed = "100",
    ///             VoiceType = "ruxue",
    ///             Volume = "0",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ci media_tts_template can be imported using the bucket#templateId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate media_tts_template terraform-ci-xxxxxx#t1ed421df8bd2140b6b73474f70f99b0f8
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate")]
    public partial class MediaTtsTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
        /// </summary>
        [Output("codec")]
        public Output<string?> Codec { get; private set; } = null!;

        /// <summary>
        /// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Speech rate, the default value is 100, [50,200].
        /// </summary>
        [Output("speed")]
        public Output<string?> Speed { get; private set; } = null!;

        /// <summary>
        /// Timbre, the default value is ruxue.
        /// </summary>
        [Output("voiceType")]
        public Output<string?> VoiceType { get; private set; } = null!;

        /// <summary>
        /// Volume, default value 0, [-10,10].
        /// </summary>
        [Output("volume")]
        public Output<string?> Volume { get; private set; } = null!;


        /// <summary>
        /// Create a MediaTtsTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaTtsTemplate(string name, MediaTtsTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate", name, args ?? new MediaTtsTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaTtsTemplate(string name, Input<string> id, MediaTtsTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaTtsTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaTtsTemplate Get(string name, Input<string> id, MediaTtsTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaTtsTemplate(name, id, state, options);
        }
    }

    public sealed class MediaTtsTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        /// <summary>
        /// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Speech rate, the default value is 100, [50,200].
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Timbre, the default value is ruxue.
        /// </summary>
        [Input("voiceType")]
        public Input<string>? VoiceType { get; set; }

        /// <summary>
        /// Volume, default value 0, [-10,10].
        /// </summary>
        [Input("volume")]
        public Input<string>? Volume { get; set; }

        public MediaTtsTemplateArgs()
        {
        }
    }

    public sealed class MediaTtsTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        /// <summary>
        /// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Speech rate, the default value is 100, [50,200].
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Timbre, the default value is ruxue.
        /// </summary>
        [Input("voiceType")]
        public Input<string>? VoiceType { get; set; }

        /// <summary>
        /// Volume, default value 0, [-10,10].
        /// </summary>
        [Input("volume")]
        public Input<string>? Volume { get; set; }

        public MediaTtsTemplateState()
        {
        }
    }
}
