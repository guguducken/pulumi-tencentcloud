# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InternetAddressConfigArgs', 'InternetAddressConfig']

@pulumi.input_type
class InternetAddressConfigArgs:
    def __init__(__self__, *,
                 enable: pulumi.Input[bool],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a InternetAddressConfig resource.
        :param pulumi.Input[bool] enable: whether enable internet address.
        :param pulumi.Input[str] instance_id: internet public address id.
        """
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[bool]:
        """
        whether enable internet address.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        internet public address id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _InternetAddressConfigState:
    def __init__(__self__, *,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InternetAddressConfig resources.
        :param pulumi.Input[bool] enable: whether enable internet address.
        :param pulumi.Input[str] instance_id: internet public address id.
        """
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        whether enable internet address.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        internet public address id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class InternetAddressConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dc internet_address_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        internet_address = tencentcloud.dc.InternetAddress("internetAddress",
            mask_len=30,
            addr_type=2,
            addr_proto=0)
        internet_address_config = tencentcloud.dc.InternetAddressConfig("internetAddressConfig",
            instance_id=internet_address.id,
            enable=False)
        ```

        ## Import

        dc internet_address_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dc/internetAddressConfig:InternetAddressConfig internet_address_config internet_address_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable: whether enable internet address.
        :param pulumi.Input[str] instance_id: internet public address id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InternetAddressConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dc internet_address_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        internet_address = tencentcloud.dc.InternetAddress("internetAddress",
            mask_len=30,
            addr_type=2,
            addr_proto=0)
        internet_address_config = tencentcloud.dc.InternetAddressConfig("internetAddressConfig",
            instance_id=internet_address.id,
            enable=False)
        ```

        ## Import

        dc internet_address_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dc/internetAddressConfig:InternetAddressConfig internet_address_config internet_address_id
        ```

        :param str resource_name: The name of the resource.
        :param InternetAddressConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InternetAddressConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = InternetAddressConfigArgs.__new__(InternetAddressConfigArgs)

            if enable is None and not opts.urn:
                raise TypeError("Missing required property 'enable'")
            __props__.__dict__["enable"] = enable
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(InternetAddressConfig, __self__).__init__(
            'tencentcloud:Dc/internetAddressConfig:InternetAddressConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'InternetAddressConfig':
        """
        Get an existing InternetAddressConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable: whether enable internet address.
        :param pulumi.Input[str] instance_id: internet public address id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InternetAddressConfigState.__new__(_InternetAddressConfigState)

        __props__.__dict__["enable"] = enable
        __props__.__dict__["instance_id"] = instance_id
        return InternetAddressConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[bool]:
        """
        whether enable internet address.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        internet public address id.
        """
        return pulumi.get(self, "instance_id")

