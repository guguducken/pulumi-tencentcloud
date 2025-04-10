# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceConfigArgs', 'InstanceConfig']

@pulumi.input_type
class InstanceConfigArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 rs_access_strategy: pulumi.Input[int]):
        """
        The set of arguments for constructing a InstanceConfig resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] rs_access_strategy: RS nearest access mode, 0-no policy, 1-nearest access.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "rs_access_strategy", rs_access_strategy)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="rsAccessStrategy")
    def rs_access_strategy(self) -> pulumi.Input[int]:
        """
        RS nearest access mode, 0-no policy, 1-nearest access.
        """
        return pulumi.get(self, "rs_access_strategy")

    @rs_access_strategy.setter
    def rs_access_strategy(self, value: pulumi.Input[int]):
        pulumi.set(self, "rs_access_strategy", value)


@pulumi.input_type
class _InstanceConfigState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 rs_access_strategy: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering InstanceConfig resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] rs_access_strategy: RS nearest access mode, 0-no policy, 1-nearest access.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if rs_access_strategy is not None:
            pulumi.set(__self__, "rs_access_strategy", rs_access_strategy)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="rsAccessStrategy")
    def rs_access_strategy(self) -> Optional[pulumi.Input[int]]:
        """
        RS nearest access mode, 0-no policy, 1-nearest access.
        """
        return pulumi.get(self, "rs_access_strategy")

    @rs_access_strategy.setter
    def rs_access_strategy(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rs_access_strategy", value)


class InstanceConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 rs_access_strategy: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a dcdb instance_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        instance_config = tencentcloud.dcdb.InstanceConfig("instanceConfig",
            instance_id=local["dcdb_id"],
            rs_access_strategy=0)
        ```

        ## Import

        dcdb instance_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dcdb/instanceConfig:InstanceConfig instance_config instance_config_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] rs_access_strategy: RS nearest access mode, 0-no policy, 1-nearest access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dcdb instance_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        instance_config = tencentcloud.dcdb.InstanceConfig("instanceConfig",
            instance_id=local["dcdb_id"],
            rs_access_strategy=0)
        ```

        ## Import

        dcdb instance_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dcdb/instanceConfig:InstanceConfig instance_config instance_config_id
        ```

        :param str resource_name: The name of the resource.
        :param InstanceConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 rs_access_strategy: Optional[pulumi.Input[int]] = None,
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
            __props__ = InstanceConfigArgs.__new__(InstanceConfigArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if rs_access_strategy is None and not opts.urn:
                raise TypeError("Missing required property 'rs_access_strategy'")
            __props__.__dict__["rs_access_strategy"] = rs_access_strategy
        super(InstanceConfig, __self__).__init__(
            'tencentcloud:Dcdb/instanceConfig:InstanceConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            rs_access_strategy: Optional[pulumi.Input[int]] = None) -> 'InstanceConfig':
        """
        Get an existing InstanceConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] rs_access_strategy: RS nearest access mode, 0-no policy, 1-nearest access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceConfigState.__new__(_InstanceConfigState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["rs_access_strategy"] = rs_access_strategy
        return InstanceConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="rsAccessStrategy")
    def rs_access_strategy(self) -> pulumi.Output[int]:
        """
        RS nearest access mode, 0-no policy, 1-nearest access.
        """
        return pulumi.get(self, "rs_access_strategy")

