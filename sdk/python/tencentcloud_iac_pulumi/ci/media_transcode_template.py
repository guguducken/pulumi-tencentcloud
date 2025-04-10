# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MediaTranscodeTemplateArgs', 'MediaTranscodeTemplate']

@pulumi.input_type
class MediaTranscodeTemplateArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 container: pulumi.Input['MediaTranscodeTemplateContainerArgs'],
                 audio: Optional[pulumi.Input['MediaTranscodeTemplateAudioArgs']] = None,
                 audio_mixes: Optional[pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 time_interval: Optional[pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs']] = None,
                 trans_config: Optional[pulumi.Input['MediaTranscodeTemplateTransConfigArgs']] = None,
                 video: Optional[pulumi.Input['MediaTranscodeTemplateVideoArgs']] = None):
        """
        The set of arguments for constructing a MediaTranscodeTemplate resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input['MediaTranscodeTemplateContainerArgs'] container: container format.
        :param pulumi.Input['MediaTranscodeTemplateAudioArgs'] audio: Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        :param pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]] audio_mixes: mixing parameters.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs'] time_interval: time interval.
        :param pulumi.Input['MediaTranscodeTemplateTransConfigArgs'] trans_config: transcoding configuration.
        :param pulumi.Input['MediaTranscodeTemplateVideoArgs'] video: video information, do not upload Video, which is equivalent to deleting video information.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "container", container)
        if audio is not None:
            pulumi.set(__self__, "audio", audio)
        if audio_mixes is not None:
            pulumi.set(__self__, "audio_mixes", audio_mixes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if time_interval is not None:
            pulumi.set(__self__, "time_interval", time_interval)
        if trans_config is not None:
            pulumi.set(__self__, "trans_config", trans_config)
        if video is not None:
            pulumi.set(__self__, "video", video)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def container(self) -> pulumi.Input['MediaTranscodeTemplateContainerArgs']:
        """
        container format.
        """
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: pulumi.Input['MediaTranscodeTemplateContainerArgs']):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def audio(self) -> Optional[pulumi.Input['MediaTranscodeTemplateAudioArgs']]:
        """
        Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        """
        return pulumi.get(self, "audio")

    @audio.setter
    def audio(self, value: Optional[pulumi.Input['MediaTranscodeTemplateAudioArgs']]):
        pulumi.set(self, "audio", value)

    @property
    @pulumi.getter(name="audioMixes")
    def audio_mixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]]]:
        """
        mixing parameters.
        """
        return pulumi.get(self, "audio_mixes")

    @audio_mixes.setter
    def audio_mixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]]]):
        pulumi.set(self, "audio_mixes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="timeInterval")
    def time_interval(self) -> Optional[pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs']]:
        """
        time interval.
        """
        return pulumi.get(self, "time_interval")

    @time_interval.setter
    def time_interval(self, value: Optional[pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs']]):
        pulumi.set(self, "time_interval", value)

    @property
    @pulumi.getter(name="transConfig")
    def trans_config(self) -> Optional[pulumi.Input['MediaTranscodeTemplateTransConfigArgs']]:
        """
        transcoding configuration.
        """
        return pulumi.get(self, "trans_config")

    @trans_config.setter
    def trans_config(self, value: Optional[pulumi.Input['MediaTranscodeTemplateTransConfigArgs']]):
        pulumi.set(self, "trans_config", value)

    @property
    @pulumi.getter
    def video(self) -> Optional[pulumi.Input['MediaTranscodeTemplateVideoArgs']]:
        """
        video information, do not upload Video, which is equivalent to deleting video information.
        """
        return pulumi.get(self, "video")

    @video.setter
    def video(self, value: Optional[pulumi.Input['MediaTranscodeTemplateVideoArgs']]):
        pulumi.set(self, "video", value)


@pulumi.input_type
class _MediaTranscodeTemplateState:
    def __init__(__self__, *,
                 audio: Optional[pulumi.Input['MediaTranscodeTemplateAudioArgs']] = None,
                 audio_mixes: Optional[pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 container: Optional[pulumi.Input['MediaTranscodeTemplateContainerArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 time_interval: Optional[pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs']] = None,
                 trans_config: Optional[pulumi.Input['MediaTranscodeTemplateTransConfigArgs']] = None,
                 video: Optional[pulumi.Input['MediaTranscodeTemplateVideoArgs']] = None):
        """
        Input properties used for looking up and filtering MediaTranscodeTemplate resources.
        :param pulumi.Input['MediaTranscodeTemplateAudioArgs'] audio: Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        :param pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]] audio_mixes: mixing parameters.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input['MediaTranscodeTemplateContainerArgs'] container: container format.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs'] time_interval: time interval.
        :param pulumi.Input['MediaTranscodeTemplateTransConfigArgs'] trans_config: transcoding configuration.
        :param pulumi.Input['MediaTranscodeTemplateVideoArgs'] video: video information, do not upload Video, which is equivalent to deleting video information.
        """
        if audio is not None:
            pulumi.set(__self__, "audio", audio)
        if audio_mixes is not None:
            pulumi.set(__self__, "audio_mixes", audio_mixes)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if container is not None:
            pulumi.set(__self__, "container", container)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if time_interval is not None:
            pulumi.set(__self__, "time_interval", time_interval)
        if trans_config is not None:
            pulumi.set(__self__, "trans_config", trans_config)
        if video is not None:
            pulumi.set(__self__, "video", video)

    @property
    @pulumi.getter
    def audio(self) -> Optional[pulumi.Input['MediaTranscodeTemplateAudioArgs']]:
        """
        Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        """
        return pulumi.get(self, "audio")

    @audio.setter
    def audio(self, value: Optional[pulumi.Input['MediaTranscodeTemplateAudioArgs']]):
        pulumi.set(self, "audio", value)

    @property
    @pulumi.getter(name="audioMixes")
    def audio_mixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]]]:
        """
        mixing parameters.
        """
        return pulumi.get(self, "audio_mixes")

    @audio_mixes.setter
    def audio_mixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MediaTranscodeTemplateAudioMixArgs']]]]):
        pulumi.set(self, "audio_mixes", value)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def container(self) -> Optional[pulumi.Input['MediaTranscodeTemplateContainerArgs']]:
        """
        container format.
        """
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: Optional[pulumi.Input['MediaTranscodeTemplateContainerArgs']]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="timeInterval")
    def time_interval(self) -> Optional[pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs']]:
        """
        time interval.
        """
        return pulumi.get(self, "time_interval")

    @time_interval.setter
    def time_interval(self, value: Optional[pulumi.Input['MediaTranscodeTemplateTimeIntervalArgs']]):
        pulumi.set(self, "time_interval", value)

    @property
    @pulumi.getter(name="transConfig")
    def trans_config(self) -> Optional[pulumi.Input['MediaTranscodeTemplateTransConfigArgs']]:
        """
        transcoding configuration.
        """
        return pulumi.get(self, "trans_config")

    @trans_config.setter
    def trans_config(self, value: Optional[pulumi.Input['MediaTranscodeTemplateTransConfigArgs']]):
        pulumi.set(self, "trans_config", value)

    @property
    @pulumi.getter
    def video(self) -> Optional[pulumi.Input['MediaTranscodeTemplateVideoArgs']]:
        """
        video information, do not upload Video, which is equivalent to deleting video information.
        """
        return pulumi.get(self, "video")

    @video.setter
    def video(self, value: Optional[pulumi.Input['MediaTranscodeTemplateVideoArgs']]):
        pulumi.set(self, "video", value)


class MediaTranscodeTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audio: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioArgs']]] = None,
                 audio_mixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioMixArgs']]]]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 container: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateContainerArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 time_interval: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTimeIntervalArgs']]] = None,
                 trans_config: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTransConfigArgs']]] = None,
                 video: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateVideoArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a ci media_transcode_template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        media_transcode_template = tencentcloud.ci.MediaTranscodeTemplate("mediaTranscodeTemplate",
            audio=tencentcloud.ci.MediaTranscodeTemplateAudioArgs(
                bitrate="128",
                channels="4",
                codec="aac",
                keep_two_tracks="false",
                remove="false",
                sample_format="",
                samplerate="44100",
                switch_track="false",
            ),
            audio_mixes=[tencentcloud.ci.MediaTranscodeTemplateAudioMixArgs(
                audio_source="https://terraform-ci-1308919341.cos.ap-guangzhou.myqcloud.com/mp3%2Fnizhan-test.mp3",
                effect_config=tencentcloud.ci.MediaTranscodeTemplateAudioMixEffectConfigArgs(
                    bgm_fade_time="1.7",
                    enable_bgm_fade="true",
                    enable_end_fadeout="false",
                    enable_start_fadein="true",
                    end_fadeout_time="0",
                    start_fadein_time="3",
                ),
                mix_mode="Once",
                replace="true",
            )],
            bucket="terraform-ci-1308919341",
            container=tencentcloud.ci.MediaTranscodeTemplateContainerArgs(
                format="mp4",
            ),
            time_interval=tencentcloud.ci.MediaTranscodeTemplateTimeIntervalArgs(
                duration="60",
                start="0",
            ),
            trans_config=tencentcloud.ci.MediaTranscodeTemplateTransConfigArgs(
                adj_dar_method="scale",
                audio_bitrate_adj_method="0",
                delete_metadata="false",
                is_check_audio_bitrate="false",
                is_check_reso="false",
                is_check_video_bitrate="false",
                is_hdr2_sdr="false",
                reso_adj_method="1",
                video_bitrate_adj_method="0",
            ),
            video=tencentcloud.ci.MediaTranscodeTemplateVideoArgs(
                bitrate="1000",
                codec="H.264",
                fps="30",
                long_short_mode="false",
                preset="medium",
                profile="high",
                remove="false",
                width="1280",
            ))
        ```

        ## Import

        ci media_transcode_template can be imported using the bucket#templateId, e.g.

        ```sh
         $ pulumi import tencentcloud:Ci/mediaTranscodeTemplate:MediaTranscodeTemplate media_transcode_template media_transcode_template_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioArgs']] audio: Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioMixArgs']]]] audio_mixes: mixing parameters.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateContainerArgs']] container: container format.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTimeIntervalArgs']] time_interval: time interval.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTransConfigArgs']] trans_config: transcoding configuration.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateVideoArgs']] video: video information, do not upload Video, which is equivalent to deleting video information.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MediaTranscodeTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a ci media_transcode_template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        media_transcode_template = tencentcloud.ci.MediaTranscodeTemplate("mediaTranscodeTemplate",
            audio=tencentcloud.ci.MediaTranscodeTemplateAudioArgs(
                bitrate="128",
                channels="4",
                codec="aac",
                keep_two_tracks="false",
                remove="false",
                sample_format="",
                samplerate="44100",
                switch_track="false",
            ),
            audio_mixes=[tencentcloud.ci.MediaTranscodeTemplateAudioMixArgs(
                audio_source="https://terraform-ci-1308919341.cos.ap-guangzhou.myqcloud.com/mp3%2Fnizhan-test.mp3",
                effect_config=tencentcloud.ci.MediaTranscodeTemplateAudioMixEffectConfigArgs(
                    bgm_fade_time="1.7",
                    enable_bgm_fade="true",
                    enable_end_fadeout="false",
                    enable_start_fadein="true",
                    end_fadeout_time="0",
                    start_fadein_time="3",
                ),
                mix_mode="Once",
                replace="true",
            )],
            bucket="terraform-ci-1308919341",
            container=tencentcloud.ci.MediaTranscodeTemplateContainerArgs(
                format="mp4",
            ),
            time_interval=tencentcloud.ci.MediaTranscodeTemplateTimeIntervalArgs(
                duration="60",
                start="0",
            ),
            trans_config=tencentcloud.ci.MediaTranscodeTemplateTransConfigArgs(
                adj_dar_method="scale",
                audio_bitrate_adj_method="0",
                delete_metadata="false",
                is_check_audio_bitrate="false",
                is_check_reso="false",
                is_check_video_bitrate="false",
                is_hdr2_sdr="false",
                reso_adj_method="1",
                video_bitrate_adj_method="0",
            ),
            video=tencentcloud.ci.MediaTranscodeTemplateVideoArgs(
                bitrate="1000",
                codec="H.264",
                fps="30",
                long_short_mode="false",
                preset="medium",
                profile="high",
                remove="false",
                width="1280",
            ))
        ```

        ## Import

        ci media_transcode_template can be imported using the bucket#templateId, e.g.

        ```sh
         $ pulumi import tencentcloud:Ci/mediaTranscodeTemplate:MediaTranscodeTemplate media_transcode_template media_transcode_template_id
        ```

        :param str resource_name: The name of the resource.
        :param MediaTranscodeTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MediaTranscodeTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audio: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioArgs']]] = None,
                 audio_mixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioMixArgs']]]]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 container: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateContainerArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 time_interval: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTimeIntervalArgs']]] = None,
                 trans_config: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTransConfigArgs']]] = None,
                 video: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateVideoArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MediaTranscodeTemplateArgs.__new__(MediaTranscodeTemplateArgs)

            __props__.__dict__["audio"] = audio
            __props__.__dict__["audio_mixes"] = audio_mixes
            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if container is None and not opts.urn:
                raise TypeError("Missing required property 'container'")
            __props__.__dict__["container"] = container
            __props__.__dict__["name"] = name
            __props__.__dict__["time_interval"] = time_interval
            __props__.__dict__["trans_config"] = trans_config
            __props__.__dict__["video"] = video
        super(MediaTranscodeTemplate, __self__).__init__(
            'tencentcloud:Ci/mediaTranscodeTemplate:MediaTranscodeTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            audio: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioArgs']]] = None,
            audio_mixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioMixArgs']]]]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            container: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateContainerArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            time_interval: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTimeIntervalArgs']]] = None,
            trans_config: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTransConfigArgs']]] = None,
            video: Optional[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateVideoArgs']]] = None) -> 'MediaTranscodeTemplate':
        """
        Get an existing MediaTranscodeTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioArgs']] audio: Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaTranscodeTemplateAudioMixArgs']]]] audio_mixes: mixing parameters.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateContainerArgs']] container: container format.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTimeIntervalArgs']] time_interval: time interval.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateTransConfigArgs']] trans_config: transcoding configuration.
        :param pulumi.Input[pulumi.InputType['MediaTranscodeTemplateVideoArgs']] video: video information, do not upload Video, which is equivalent to deleting video information.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MediaTranscodeTemplateState.__new__(_MediaTranscodeTemplateState)

        __props__.__dict__["audio"] = audio
        __props__.__dict__["audio_mixes"] = audio_mixes
        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["container"] = container
        __props__.__dict__["name"] = name
        __props__.__dict__["time_interval"] = time_interval
        __props__.__dict__["trans_config"] = trans_config
        __props__.__dict__["video"] = video
        return MediaTranscodeTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def audio(self) -> pulumi.Output[Optional['outputs.MediaTranscodeTemplateAudio']]:
        """
        Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        """
        return pulumi.get(self, "audio")

    @property
    @pulumi.getter(name="audioMixes")
    def audio_mixes(self) -> pulumi.Output[Optional[Sequence['outputs.MediaTranscodeTemplateAudioMix']]]:
        """
        mixing parameters.
        """
        return pulumi.get(self, "audio_mixes")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def container(self) -> pulumi.Output['outputs.MediaTranscodeTemplateContainer']:
        """
        container format.
        """
        return pulumi.get(self, "container")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="timeInterval")
    def time_interval(self) -> pulumi.Output[Optional['outputs.MediaTranscodeTemplateTimeInterval']]:
        """
        time interval.
        """
        return pulumi.get(self, "time_interval")

    @property
    @pulumi.getter(name="transConfig")
    def trans_config(self) -> pulumi.Output[Optional['outputs.MediaTranscodeTemplateTransConfig']]:
        """
        transcoding configuration.
        """
        return pulumi.get(self, "trans_config")

    @property
    @pulumi.getter
    def video(self) -> pulumi.Output[Optional['outputs.MediaTranscodeTemplateVideo']]:
        """
        video information, do not upload Video, which is equivalent to deleting video information.
        """
        return pulumi.get(self, "video")

