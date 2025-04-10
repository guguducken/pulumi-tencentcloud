# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstancesAcceptAttachInstanceArgs',
    'InstancesRejectAttachInstanceArgs',
    'InstancesResetAttachInstanceArgs',
    'GetCrossBorderRegionBandwidthLimitsFilterArgs',
]

@pulumi.input_type
class InstancesAcceptAttachInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 instance_region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance_id: Attachment Instance ID.
        :param pulumi.Input[str] instance_region: Instance Region.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] instance_type: InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        :param pulumi.Input[str] route_table_id: ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Attachment Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> pulumi.Input[str]:
        """
        Instance Region.
        """
        return pulumi.get(self, "instance_region")

    @instance_region.setter
    def instance_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_region", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)


@pulumi.input_type
class InstancesRejectAttachInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 instance_region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance_id: Attachment Instance ID.
        :param pulumi.Input[str] instance_region: Instance Region.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] instance_type: InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        :param pulumi.Input[str] route_table_id: ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Attachment Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> pulumi.Input[str]:
        """
        Instance Region.
        """
        return pulumi.get(self, "instance_region")

    @instance_region.setter
    def instance_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_region", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)


@pulumi.input_type
class InstancesResetAttachInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 instance_region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance_id: Attachment Instance ID.
        :param pulumi.Input[str] instance_region: Instance Region.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] instance_type: InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        :param pulumi.Input[str] route_table_id: ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Attachment Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> pulumi.Input[str]:
        """
        Instance Region.
        """
        return pulumi.get(self, "instance_region")

    @instance_region.setter
    def instance_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_region", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)


@pulumi.input_type
class GetCrossBorderRegionBandwidthLimitsFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: attribute name.
        :param Sequence[str] values: Value of the field.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        attribute name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Value of the field.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


