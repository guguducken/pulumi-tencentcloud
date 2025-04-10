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
    /// Provides a resource to create a ci media_video_montage_template
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
    ///         var mediaVideoMontageTemplate = new Tencentcloud.Ci.MediaVideoMontageTemplate("mediaVideoMontageTemplate", new Tencentcloud.Ci.MediaVideoMontageTemplateArgs
    ///         {
    ///             Audio = new Tencentcloud.Ci.Inputs.MediaVideoMontageTemplateAudioArgs
    ///             {
    ///                 Bitrate = "128",
    ///                 Channels = "4",
    ///                 Codec = "aac",
    ///                 Remove = "false",
    ///                 Samplerate = "44100",
    ///             },
    ///             AudioMixes = 
    ///             {
    ///                 new Tencentcloud.Ci.Inputs.MediaVideoMontageTemplateAudioMixArgs
    ///                 {
    ///                     AudioSource = "https://terraform-ci-xxxxx.cos.ap-guangzhou.myqcloud.com/mp3%2Fnizhan-test.mp3",
    ///                     MixMode = "Once",
    ///                     Replace = "true",
    ///                 },
    ///             },
    ///             Bucket = "terraform-ci-xxxxx",
    ///             Container = new Tencentcloud.Ci.Inputs.MediaVideoMontageTemplateContainerArgs
    ///             {
    ///                 Format = "mp4",
    ///             },
    ///             Duration = "10.5",
    ///             Video = new Tencentcloud.Ci.Inputs.MediaVideoMontageTemplateVideoArgs
    ///             {
    ///                 Bitrate = "1000",
    ///                 Codec = "H.264",
    ///                 Crf = "",
    ///                 Fps = "25",
    ///                 Height = "",
    ///                 Remove = "",
    ///                 Width = "1280",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ci media_video_montage_template can be imported using the bucket#templateId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate media_video_montage_template terraform-ci-xxxxxx#t193e5ecc1b8154e57a8376b4405ad9c63
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate")]
    public partial class MediaVideoMontageTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
        /// </summary>
        [Output("audio")]
        public Output<Outputs.MediaVideoMontageTemplateAudio?> Audio { get; private set; } = null!;

        /// <summary>
        /// mixing parameters.
        /// </summary>
        [Output("audioMixes")]
        public Output<ImmutableArray<Outputs.MediaVideoMontageTemplateAudioMix>> AudioMixes { get; private set; } = null!;

        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// container format.
        /// </summary>
        [Output("container")]
        public Output<Outputs.MediaVideoMontageTemplateContainer> Container { get; private set; } = null!;

        /// <summary>
        /// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution accuracy is accurate to milliseconds.
        /// </summary>
        [Output("duration")]
        public Output<string?> Duration { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// video information, do not upload Video, which is equivalent to deleting video information.
        /// </summary>
        [Output("video")]
        public Output<Outputs.MediaVideoMontageTemplateVideo?> Video { get; private set; } = null!;


        /// <summary>
        /// Create a MediaVideoMontageTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaVideoMontageTemplate(string name, MediaVideoMontageTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate", name, args ?? new MediaVideoMontageTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaVideoMontageTemplate(string name, Input<string> id, MediaVideoMontageTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaVideoMontageTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaVideoMontageTemplate Get(string name, Input<string> id, MediaVideoMontageTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaVideoMontageTemplate(name, id, state, options);
        }
    }

    public sealed class MediaVideoMontageTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
        /// </summary>
        [Input("audio")]
        public Input<Inputs.MediaVideoMontageTemplateAudioArgs>? Audio { get; set; }

        [Input("audioMixes")]
        private InputList<Inputs.MediaVideoMontageTemplateAudioMixArgs>? _audioMixes;

        /// <summary>
        /// mixing parameters.
        /// </summary>
        public InputList<Inputs.MediaVideoMontageTemplateAudioMixArgs> AudioMixes
        {
            get => _audioMixes ?? (_audioMixes = new InputList<Inputs.MediaVideoMontageTemplateAudioMixArgs>());
            set => _audioMixes = value;
        }

        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// container format.
        /// </summary>
        [Input("container", required: true)]
        public Input<Inputs.MediaVideoMontageTemplateContainerArgs> Container { get; set; } = null!;

        /// <summary>
        /// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution accuracy is accurate to milliseconds.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// video information, do not upload Video, which is equivalent to deleting video information.
        /// </summary>
        [Input("video")]
        public Input<Inputs.MediaVideoMontageTemplateVideoArgs>? Video { get; set; }

        public MediaVideoMontageTemplateArgs()
        {
        }
    }

    public sealed class MediaVideoMontageTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
        /// </summary>
        [Input("audio")]
        public Input<Inputs.MediaVideoMontageTemplateAudioGetArgs>? Audio { get; set; }

        [Input("audioMixes")]
        private InputList<Inputs.MediaVideoMontageTemplateAudioMixGetArgs>? _audioMixes;

        /// <summary>
        /// mixing parameters.
        /// </summary>
        public InputList<Inputs.MediaVideoMontageTemplateAudioMixGetArgs> AudioMixes
        {
            get => _audioMixes ?? (_audioMixes = new InputList<Inputs.MediaVideoMontageTemplateAudioMixGetArgs>());
            set => _audioMixes = value;
        }

        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// container format.
        /// </summary>
        [Input("container")]
        public Input<Inputs.MediaVideoMontageTemplateContainerGetArgs>? Container { get; set; }

        /// <summary>
        /// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution accuracy is accurate to milliseconds.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// video information, do not upload Video, which is equivalent to deleting video information.
        /// </summary>
        [Input("video")]
        public Input<Inputs.MediaVideoMontageTemplateVideoGetArgs>? Video { get; set; }

        public MediaVideoMontageTemplateState()
        {
        }
    }
}
