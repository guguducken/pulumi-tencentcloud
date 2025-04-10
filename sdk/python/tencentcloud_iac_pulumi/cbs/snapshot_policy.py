# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnapshotPolicyArgs', 'SnapshotPolicy']

@pulumi.input_type
class SnapshotPolicyArgs:
    def __init__(__self__, *,
                 repeat_hours: pulumi.Input[Sequence[pulumi.Input[int]]],
                 repeat_weekdays: pulumi.Input[Sequence[pulumi.Input[int]]],
                 snapshot_policy_name: pulumi.Input[str],
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SnapshotPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_hours: Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_weekdays: Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        :param pulumi.Input[str] snapshot_policy_name: Name of snapshot policy. The maximum length can not exceed 60 bytes.
        :param pulumi.Input[int] retention_days: Retention days of the snapshot, and the default value is 7.
        """
        pulumi.set(__self__, "repeat_hours", repeat_hours)
        pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        pulumi.set(__self__, "snapshot_policy_name", snapshot_policy_name)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="repeatHours")
    def repeat_hours(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        """
        return pulumi.get(self, "repeat_hours")

    @repeat_hours.setter
    def repeat_hours(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "repeat_hours", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> pulumi.Input[str]:
        """
        Name of snapshot policy. The maximum length can not exceed 60 bytes.
        """
        return pulumi.get(self, "snapshot_policy_name")

    @snapshot_policy_name.setter
    def snapshot_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_policy_name", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Retention days of the snapshot, and the default value is 7.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


@pulumi.input_type
class _SnapshotPolicyState:
    def __init__(__self__, *,
                 repeat_hours: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 snapshot_policy_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnapshotPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_hours: Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_weekdays: Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        :param pulumi.Input[int] retention_days: Retention days of the snapshot, and the default value is 7.
        :param pulumi.Input[str] snapshot_policy_name: Name of snapshot policy. The maximum length can not exceed 60 bytes.
        """
        if repeat_hours is not None:
            pulumi.set(__self__, "repeat_hours", repeat_hours)
        if repeat_weekdays is not None:
            pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if snapshot_policy_name is not None:
            pulumi.set(__self__, "snapshot_policy_name", snapshot_policy_name)

    @property
    @pulumi.getter(name="repeatHours")
    def repeat_hours(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        """
        return pulumi.get(self, "repeat_hours")

    @repeat_hours.setter
    def repeat_hours(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "repeat_hours", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Retention days of the snapshot, and the default value is 7.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of snapshot policy. The maximum length can not exceed 60 bytes.
        """
        return pulumi.get(self, "snapshot_policy_name")

    @snapshot_policy_name.setter
    def snapshot_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_policy_name", value)


class SnapshotPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repeat_hours: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a snapshot policy resource.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        snapshot_policy = tencentcloud.cbs.SnapshotPolicy("snapshotPolicy",
            repeat_hours=[1],
            repeat_weekdays=[
                1,
                4,
            ],
            retention_days=7,
            snapshot_policy_name="mysnapshotpolicyname")
        ```

        ## Import

        CBS snapshot policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy snapshot_policy asp-jliex1tn
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_hours: Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_weekdays: Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        :param pulumi.Input[int] retention_days: Retention days of the snapshot, and the default value is 7.
        :param pulumi.Input[str] snapshot_policy_name: Name of snapshot policy. The maximum length can not exceed 60 bytes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a snapshot policy resource.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        snapshot_policy = tencentcloud.cbs.SnapshotPolicy("snapshotPolicy",
            repeat_hours=[1],
            repeat_weekdays=[
                1,
                4,
            ],
            retention_days=7,
            snapshot_policy_name="mysnapshotpolicyname")
        ```

        ## Import

        CBS snapshot policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy snapshot_policy asp-jliex1tn
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repeat_hours: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 snapshot_policy_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = SnapshotPolicyArgs.__new__(SnapshotPolicyArgs)

            if repeat_hours is None and not opts.urn:
                raise TypeError("Missing required property 'repeat_hours'")
            __props__.__dict__["repeat_hours"] = repeat_hours
            if repeat_weekdays is None and not opts.urn:
                raise TypeError("Missing required property 'repeat_weekdays'")
            __props__.__dict__["repeat_weekdays"] = repeat_weekdays
            __props__.__dict__["retention_days"] = retention_days
            if snapshot_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_policy_name'")
            __props__.__dict__["snapshot_policy_name"] = snapshot_policy_name
        super(SnapshotPolicy, __self__).__init__(
            'tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            repeat_hours: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            retention_days: Optional[pulumi.Input[int]] = None,
            snapshot_policy_name: Optional[pulumi.Input[str]] = None) -> 'SnapshotPolicy':
        """
        Get an existing SnapshotPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_hours: Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] repeat_weekdays: Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        :param pulumi.Input[int] retention_days: Retention days of the snapshot, and the default value is 7.
        :param pulumi.Input[str] snapshot_policy_name: Name of snapshot policy. The maximum length can not exceed 60 bytes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotPolicyState.__new__(_SnapshotPolicyState)

        __props__.__dict__["repeat_hours"] = repeat_hours
        __props__.__dict__["repeat_weekdays"] = repeat_weekdays
        __props__.__dict__["retention_days"] = retention_days
        __props__.__dict__["snapshot_policy_name"] = snapshot_policy_name
        return SnapshotPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="repeatHours")
    def repeat_hours(self) -> pulumi.Output[Sequence[int]]:
        """
        Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        """
        return pulumi.get(self, "repeat_hours")

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Output[Sequence[int]]:
        """
        Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        """
        return pulumi.get(self, "repeat_weekdays")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        Retention days of the snapshot, and the default value is 7.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> pulumi.Output[str]:
        """
        Name of snapshot policy. The maximum length can not exceed 60 bytes.
        """
        return pulumi.get(self, "snapshot_policy_name")

