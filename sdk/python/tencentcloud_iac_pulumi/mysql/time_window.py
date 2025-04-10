# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TimeWindowArgs', 'TimeWindow']

@pulumi.input_type
class TimeWindowArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 time_ranges: pulumi.Input[Sequence[pulumi.Input[str]]],
                 max_delay_time: Optional[pulumi.Input[int]] = None,
                 weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TimeWindow resource.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_ranges: Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        :param pulumi.Input[int] max_delay_time: Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] weekdays: Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "time_ranges", time_ranges)
        if max_delay_time is not None:
            pulumi.set(__self__, "max_delay_time", max_delay_time)
        if weekdays is not None:
            pulumi.set(__self__, "weekdays", weekdays)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="timeRanges")
    def time_ranges(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        """
        return pulumi.get(self, "time_ranges")

    @time_ranges.setter
    def time_ranges(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "time_ranges", value)

    @property
    @pulumi.getter(name="maxDelayTime")
    def max_delay_time(self) -> Optional[pulumi.Input[int]]:
        """
        Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        """
        return pulumi.get(self, "max_delay_time")

    @max_delay_time.setter
    def max_delay_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_delay_time", value)

    @property
    @pulumi.getter
    def weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        return pulumi.get(self, "weekdays")

    @weekdays.setter
    def weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "weekdays", value)


@pulumi.input_type
class _TimeWindowState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_delay_time: Optional[pulumi.Input[int]] = None,
                 time_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering TimeWindow resources.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        :param pulumi.Input[int] max_delay_time: Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_ranges: Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        :param pulumi.Input[Sequence[pulumi.Input[str]]] weekdays: Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_delay_time is not None:
            pulumi.set(__self__, "max_delay_time", max_delay_time)
        if time_ranges is not None:
            pulumi.set(__self__, "time_ranges", time_ranges)
        if weekdays is not None:
            pulumi.set(__self__, "weekdays", weekdays)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxDelayTime")
    def max_delay_time(self) -> Optional[pulumi.Input[int]]:
        """
        Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        """
        return pulumi.get(self, "max_delay_time")

    @max_delay_time.setter
    def max_delay_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_delay_time", value)

    @property
    @pulumi.getter(name="timeRanges")
    def time_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        """
        return pulumi.get(self, "time_ranges")

    @time_ranges.setter
    def time_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "time_ranges", value)

    @property
    @pulumi.getter
    def weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        return pulumi.get(self, "weekdays")

    @weekdays.setter
    def weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "weekdays", value)


class TimeWindow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_delay_time: Optional[pulumi.Input[int]] = None,
                 time_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql time_window

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        time_window = tencentcloud.mysql.TimeWindow("timeWindow",
            instance_id="cdb-lw71b6ar",
            max_delay_time=10,
            time_ranges=["01:00-02:01"],
            weekdays=[
                "friday",
                "monday",
                "saturday",
                "thursday",
                "tuesday",
                "wednesday",
            ])
        ```

        ## Import

        mysql time_window can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Mysql/timeWindow:TimeWindow time_window instanceId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        :param pulumi.Input[int] max_delay_time: Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_ranges: Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        :param pulumi.Input[Sequence[pulumi.Input[str]]] weekdays: Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TimeWindowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql time_window

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        time_window = tencentcloud.mysql.TimeWindow("timeWindow",
            instance_id="cdb-lw71b6ar",
            max_delay_time=10,
            time_ranges=["01:00-02:01"],
            weekdays=[
                "friday",
                "monday",
                "saturday",
                "thursday",
                "tuesday",
                "wednesday",
            ])
        ```

        ## Import

        mysql time_window can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Mysql/timeWindow:TimeWindow time_window instanceId
        ```

        :param str resource_name: The name of the resource.
        :param TimeWindowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TimeWindowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_delay_time: Optional[pulumi.Input[int]] = None,
                 time_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = TimeWindowArgs.__new__(TimeWindowArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["max_delay_time"] = max_delay_time
            if time_ranges is None and not opts.urn:
                raise TypeError("Missing required property 'time_ranges'")
            __props__.__dict__["time_ranges"] = time_ranges
            __props__.__dict__["weekdays"] = weekdays
        super(TimeWindow, __self__).__init__(
            'tencentcloud:Mysql/timeWindow:TimeWindow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            max_delay_time: Optional[pulumi.Input[int]] = None,
            time_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'TimeWindow':
        """
        Get an existing TimeWindow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        :param pulumi.Input[int] max_delay_time: Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_ranges: Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        :param pulumi.Input[Sequence[pulumi.Input[str]]] weekdays: Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TimeWindowState.__new__(_TimeWindowState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_delay_time"] = max_delay_time
        __props__.__dict__["time_ranges"] = time_ranges
        __props__.__dict__["weekdays"] = weekdays
        return TimeWindow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxDelayTime")
    def max_delay_time(self) -> pulumi.Output[Optional[int]]:
        """
        Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        """
        return pulumi.get(self, "max_delay_time")

    @property
    @pulumi.getter(name="timeRanges")
    def time_ranges(self) -> pulumi.Output[Sequence[str]]:
        """
        Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        """
        return pulumi.get(self, "time_ranges")

    @property
    @pulumi.getter
    def weekdays(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        """
        return pulumi.get(self, "weekdays")

