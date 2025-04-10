# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'StreamLiveInputInputSettingArgs',
]

@pulumi.input_type
class StreamLiveInputInputSettingArgs:
    def __init__(__self__, *,
                 app_name: Optional[pulumi.Input[str]] = None,
                 delay_time: Optional[pulumi.Input[int]] = None,
                 input_address: Optional[pulumi.Input[str]] = None,
                 input_domain: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 source_url: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] app_name: Application name, which is valid if `Type` is `RTMP_PUSH` and can contain 1-32 letters and digitsNote: This field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[int] delay_time: Delayed time (ms) for playback, which is valid if `Type` is `RTMP_PUSH`Value range: 0 (default) or 10000-600000The value must be a multiple of 1,000.Note: This field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[str] input_address: RTP/UDP input address, which does not need to be entered for the input parameter.Note: this field may return null, indicating that no valid values can be obtained.
        :param pulumi.Input[str] input_domain: The domain of an SRT_PUSH address. If this is a request parameter, you do not need to specify it.Note: This field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[str] password: The password, which is used for authentication.Note: This field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[str] source_type: Source type for stream pulling and relaying. To pull content from private-read COS buckets under the current account, set this parameter to `TencentCOS`; otherwise, leave it empty.Note: this field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[str] source_url: Source URL, which is valid if `Type` is `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL` and can contain 1-512 charactersNote: This field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[str] stream_name: Stream name, which is valid if `Type` is `RTMP_PUSH` and can contain 1-32 letters and digitsNote: This field may return `null`, indicating that no valid value was found.
        :param pulumi.Input[str] user_name: The username, which is used for authentication.Note: This field may return `null`, indicating that no valid value was found.
        """
        if app_name is not None:
            pulumi.set(__self__, "app_name", app_name)
        if delay_time is not None:
            pulumi.set(__self__, "delay_time", delay_time)
        if input_address is not None:
            pulumi.set(__self__, "input_address", input_address)
        if input_domain is not None:
            pulumi.set(__self__, "input_domain", input_domain)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if source_url is not None:
            pulumi.set(__self__, "source_url", source_url)
        if stream_name is not None:
            pulumi.set(__self__, "stream_name", stream_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> Optional[pulumi.Input[str]]:
        """
        Application name, which is valid if `Type` is `RTMP_PUSH` and can contain 1-32 letters and digitsNote: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="delayTime")
    def delay_time(self) -> Optional[pulumi.Input[int]]:
        """
        Delayed time (ms) for playback, which is valid if `Type` is `RTMP_PUSH`Value range: 0 (default) or 10000-600000The value must be a multiple of 1,000.Note: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "delay_time")

    @delay_time.setter
    def delay_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delay_time", value)

    @property
    @pulumi.getter(name="inputAddress")
    def input_address(self) -> Optional[pulumi.Input[str]]:
        """
        RTP/UDP input address, which does not need to be entered for the input parameter.Note: this field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "input_address")

    @input_address.setter
    def input_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_address", value)

    @property
    @pulumi.getter(name="inputDomain")
    def input_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain of an SRT_PUSH address. If this is a request parameter, you do not need to specify it.Note: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "input_domain")

    @input_domain.setter
    def input_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_domain", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password, which is used for authentication.Note: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[str]]:
        """
        Source type for stream pulling and relaying. To pull content from private-read COS buckets under the current account, set this parameter to `TencentCOS`; otherwise, leave it empty.Note: this field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="sourceUrl")
    def source_url(self) -> Optional[pulumi.Input[str]]:
        """
        Source URL, which is valid if `Type` is `RTMP_PULL`, `HLS_PULL`, or `MP4_PULL` and can contain 1-512 charactersNote: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "source_url")

    @source_url.setter
    def source_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_url", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[pulumi.Input[str]]:
        """
        Stream name, which is valid if `Type` is `RTMP_PUSH` and can contain 1-32 letters and digitsNote: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The username, which is used for authentication.Note: This field may return `null`, indicating that no valid value was found.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


