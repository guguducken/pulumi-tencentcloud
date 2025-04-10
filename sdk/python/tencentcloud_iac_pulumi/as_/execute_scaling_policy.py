# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ExecuteScalingPolicyArgs', 'ExecuteScalingPolicy']

@pulumi.input_type
class ExecuteScalingPolicyArgs:
    def __init__(__self__, *,
                 auto_scaling_policy_id: pulumi.Input[str],
                 honor_cooldown: Optional[pulumi.Input[bool]] = None,
                 trigger_source: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ExecuteScalingPolicy resource.
        :param pulumi.Input[str] auto_scaling_policy_id: Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        :param pulumi.Input[bool] honor_cooldown: Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        :param pulumi.Input[str] trigger_source: Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        pulumi.set(__self__, "auto_scaling_policy_id", auto_scaling_policy_id)
        if honor_cooldown is not None:
            pulumi.set(__self__, "honor_cooldown", honor_cooldown)
        if trigger_source is not None:
            pulumi.set(__self__, "trigger_source", trigger_source)

    @property
    @pulumi.getter(name="autoScalingPolicyId")
    def auto_scaling_policy_id(self) -> pulumi.Input[str]:
        """
        Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        """
        return pulumi.get(self, "auto_scaling_policy_id")

    @auto_scaling_policy_id.setter
    def auto_scaling_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_scaling_policy_id", value)

    @property
    @pulumi.getter(name="honorCooldown")
    def honor_cooldown(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        """
        return pulumi.get(self, "honor_cooldown")

    @honor_cooldown.setter
    def honor_cooldown(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "honor_cooldown", value)

    @property
    @pulumi.getter(name="triggerSource")
    def trigger_source(self) -> Optional[pulumi.Input[str]]:
        """
        Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        return pulumi.get(self, "trigger_source")

    @trigger_source.setter
    def trigger_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_source", value)


@pulumi.input_type
class _ExecuteScalingPolicyState:
    def __init__(__self__, *,
                 auto_scaling_policy_id: Optional[pulumi.Input[str]] = None,
                 honor_cooldown: Optional[pulumi.Input[bool]] = None,
                 trigger_source: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExecuteScalingPolicy resources.
        :param pulumi.Input[str] auto_scaling_policy_id: Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        :param pulumi.Input[bool] honor_cooldown: Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        :param pulumi.Input[str] trigger_source: Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        if auto_scaling_policy_id is not None:
            pulumi.set(__self__, "auto_scaling_policy_id", auto_scaling_policy_id)
        if honor_cooldown is not None:
            pulumi.set(__self__, "honor_cooldown", honor_cooldown)
        if trigger_source is not None:
            pulumi.set(__self__, "trigger_source", trigger_source)

    @property
    @pulumi.getter(name="autoScalingPolicyId")
    def auto_scaling_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        """
        return pulumi.get(self, "auto_scaling_policy_id")

    @auto_scaling_policy_id.setter
    def auto_scaling_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_scaling_policy_id", value)

    @property
    @pulumi.getter(name="honorCooldown")
    def honor_cooldown(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        """
        return pulumi.get(self, "honor_cooldown")

    @honor_cooldown.setter
    def honor_cooldown(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "honor_cooldown", value)

    @property
    @pulumi.getter(name="triggerSource")
    def trigger_source(self) -> Optional[pulumi.Input[str]]:
        """
        Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        return pulumi.get(self, "trigger_source")

    @trigger_source.setter
    def trigger_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_source", value)


class ExecuteScalingPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_policy_id: Optional[pulumi.Input[str]] = None,
                 honor_cooldown: Optional[pulumi.Input[bool]] = None,
                 trigger_source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a as execute_scaling_policy

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        execute_scaling_policy = tencentcloud.as_.ExecuteScalingPolicy("executeScalingPolicy",
            auto_scaling_policy_id="asp-519acdug",
            honor_cooldown=False,
            trigger_source="API")
        ```

        ## Import

        as execute_scaling_policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy execute_scaling_policy execute_scaling_policy_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_policy_id: Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        :param pulumi.Input[bool] honor_cooldown: Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        :param pulumi.Input[str] trigger_source: Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExecuteScalingPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a as execute_scaling_policy

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        execute_scaling_policy = tencentcloud.as_.ExecuteScalingPolicy("executeScalingPolicy",
            auto_scaling_policy_id="asp-519acdug",
            honor_cooldown=False,
            trigger_source="API")
        ```

        ## Import

        as execute_scaling_policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy execute_scaling_policy execute_scaling_policy_id
        ```

        :param str resource_name: The name of the resource.
        :param ExecuteScalingPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExecuteScalingPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_policy_id: Optional[pulumi.Input[str]] = None,
                 honor_cooldown: Optional[pulumi.Input[bool]] = None,
                 trigger_source: Optional[pulumi.Input[str]] = None,
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
            __props__ = ExecuteScalingPolicyArgs.__new__(ExecuteScalingPolicyArgs)

            if auto_scaling_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'auto_scaling_policy_id'")
            __props__.__dict__["auto_scaling_policy_id"] = auto_scaling_policy_id
            __props__.__dict__["honor_cooldown"] = honor_cooldown
            __props__.__dict__["trigger_source"] = trigger_source
        super(ExecuteScalingPolicy, __self__).__init__(
            'tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_scaling_policy_id: Optional[pulumi.Input[str]] = None,
            honor_cooldown: Optional[pulumi.Input[bool]] = None,
            trigger_source: Optional[pulumi.Input[str]] = None) -> 'ExecuteScalingPolicy':
        """
        Get an existing ExecuteScalingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_policy_id: Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        :param pulumi.Input[bool] honor_cooldown: Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        :param pulumi.Input[str] trigger_source: Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExecuteScalingPolicyState.__new__(_ExecuteScalingPolicyState)

        __props__.__dict__["auto_scaling_policy_id"] = auto_scaling_policy_id
        __props__.__dict__["honor_cooldown"] = honor_cooldown
        __props__.__dict__["trigger_source"] = trigger_source
        return ExecuteScalingPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingPolicyId")
    def auto_scaling_policy_id(self) -> pulumi.Output[str]:
        """
        Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        """
        return pulumi.get(self, "auto_scaling_policy_id")

    @property
    @pulumi.getter(name="honorCooldown")
    def honor_cooldown(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        """
        return pulumi.get(self, "honor_cooldown")

    @property
    @pulumi.getter(name="triggerSource")
    def trigger_source(self) -> pulumi.Output[Optional[str]]:
        """
        Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        """
        return pulumi.get(self, "trigger_source")

