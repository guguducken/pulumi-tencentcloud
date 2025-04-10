// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Inputs
{

    public sealed class MediaTranscodeTemplateAudioGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Original audio bit rate, unit: Kbps, Value range: [8, 1000].
        /// </summary>
        [Input("bitrate")]
        public Input<string>? Bitrate { get; set; }

        /// <summary>
        /// number of channels- When Codec is set to aac/flac, support 1, 2, 4, 5, 6, 8- When Codec is set to mp3/opus, support 1, 2- When Codec is set to Vorbis, only 2 is supported- When Codec is set to amr, only 1 is supported- When Codec is set to pcm_s16le, only 1 and 2 are supported- When the encapsulation format is dash, 8 is not supported.
        /// </summary>
        [Input("channels")]
        public Input<string>? Channels { get; set; }

        /// <summary>
        /// Codec format, value aac, mp3, flac, amr, Vorbis, opus, pcm_s16le.
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        /// <summary>
        /// Keep dual audio tracks, the value is true, false. This parameter is invalid when Video.Codec is H.265.
        /// </summary>
        [Input("keepTwoTracks")]
        public Input<string>? KeepTwoTracks { get; set; }

        /// <summary>
        /// Whether to delete the source audio stream, the value is true, false.
        /// </summary>
        [Input("remove")]
        public Input<string>? Remove { get; set; }

        /// <summary>
        /// Sampling bit width- When Codec is set to aac, support fltp- When Codec is set to mp3, fltp, s16p, s32p are supported- When Codec is set to flac, s16, s32, s16p, s32p are supported- When Codec is set to amr, support s16, s16p- When Codec is set to opus, support s16- When Codec is set to pcm_s16le, support s16- When Codec is set to Vorbis, support fltp- This parameter is invalid when Video.Codec is H.265.
        /// </summary>
        [Input("sampleFormat")]
        public Input<string>? SampleFormat { get; set; }

        /// <summary>
        /// Sampling Rate- Unit: Hz- Optional 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 88200, 96000- Different packages, mp3 supports different sampling rates, as shown in the table below- When Codec is set to amr, only 8000 is supported- When Codec is set to opus, it supports 8000, 16000, 24000, 48000.
        /// </summary>
        [Input("samplerate")]
        public Input<string>? Samplerate { get; set; }

        /// <summary>
        /// Convert track, the value is true, false. This parameter is invalid when Video.Codec is H.265.
        /// </summary>
        [Input("switchTrack")]
        public Input<string>? SwitchTrack { get; set; }

        public MediaTranscodeTemplateAudioGetArgs()
        {
        }
    }
}
