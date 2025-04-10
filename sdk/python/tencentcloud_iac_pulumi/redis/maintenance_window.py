# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MaintenanceWindowArgs', 'MaintenanceWindow']

@pulumi.input_type
class MaintenanceWindowArgs:
    def __init__(__self__, *,
                 end_time: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 start_time: pulumi.Input[str]):
        """
        The set of arguments for constructing a MaintenanceWindow resource.
        :param pulumi.Input[str] end_time: The end time of the maintenance window, e.g. 19:00.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] start_time: Maintenance window start time, e.g. 17:00.
        """
        pulumi.set(__self__, "end_time", end_time)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Input[str]:
        """
        The end time of the maintenance window, e.g. 19:00.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input[str]:
        """
        Maintenance window start time, e.g. 17:00.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_time", value)


@pulumi.input_type
class _MaintenanceWindowState:
    def __init__(__self__, *,
                 end_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MaintenanceWindow resources.
        :param pulumi.Input[str] end_time: The end time of the maintenance window, e.g. 19:00.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] start_time: Maintenance window start time, e.g. 17:00.
        """
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The end time of the maintenance window, e.g. 19:00.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Maintenance window start time, e.g. 17:00.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)


class MaintenanceWindow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a redis maintenance_window

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        maintenance_window = tencentcloud.redis.MaintenanceWindow("maintenanceWindow",
            end_time="19:00",
            instance_id="crs-c1nl9rpv",
            start_time="17:00")
        ```

        ## Import

        redis maintenance_window can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Redis/maintenanceWindow:MaintenanceWindow maintenance_window maintenance_window_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_time: The end time of the maintenance window, e.g. 19:00.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] start_time: Maintenance window start time, e.g. 17:00.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MaintenanceWindowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a redis maintenance_window

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        maintenance_window = tencentcloud.redis.MaintenanceWindow("maintenanceWindow",
            end_time="19:00",
            instance_id="crs-c1nl9rpv",
            start_time="17:00")
        ```

        ## Import

        redis maintenance_window can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Redis/maintenanceWindow:MaintenanceWindow maintenance_window maintenance_window_id
        ```

        :param str resource_name: The name of the resource.
        :param MaintenanceWindowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MaintenanceWindowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
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
            __props__ = MaintenanceWindowArgs.__new__(MaintenanceWindowArgs)

            if end_time is None and not opts.urn:
                raise TypeError("Missing required property 'end_time'")
            __props__.__dict__["end_time"] = end_time
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if start_time is None and not opts.urn:
                raise TypeError("Missing required property 'start_time'")
            __props__.__dict__["start_time"] = start_time
        super(MaintenanceWindow, __self__).__init__(
            'tencentcloud:Redis/maintenanceWindow:MaintenanceWindow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[str]] = None) -> 'MaintenanceWindow':
        """
        Get an existing MaintenanceWindow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_time: The end time of the maintenance window, e.g. 19:00.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] start_time: Maintenance window start time, e.g. 17:00.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MaintenanceWindowState.__new__(_MaintenanceWindowState)

        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["start_time"] = start_time
        return MaintenanceWindow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        The end time of the maintenance window, e.g. 19:00.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        Maintenance window start time, e.g. 17:00.
        """
        return pulumi.get(self, "start_time")

