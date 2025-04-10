# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TableEntryArgs', 'TableEntry']

@pulumi.input_type
class TableEntryArgs:
    def __init__(__self__, *,
                 destination_cidr_block: pulumi.Input[str],
                 next_hub: pulumi.Input[str],
                 next_type: pulumi.Input[str],
                 route_table_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TableEntry resource.
        :param pulumi.Input[str] destination_cidr_block: Destination address block.
        :param pulumi.Input[str] next_hub: ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        :param pulumi.Input[str] next_type: Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        :param pulumi.Input[str] route_table_id: ID of routing table to which this entry belongs.
        :param pulumi.Input[str] description: Description of the routing table entry.
        :param pulumi.Input[bool] disabled: Whether the entry is disabled, default is `false`.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "next_hub", next_hub)
        pulumi.set(__self__, "next_type", next_type)
        pulumi.set(__self__, "route_table_id", route_table_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        Destination address block.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="nextHub")
    def next_hub(self) -> pulumi.Input[str]:
        """
        ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        """
        return pulumi.get(self, "next_hub")

    @next_hub.setter
    def next_hub(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_hub", value)

    @property
    @pulumi.getter(name="nextType")
    def next_type(self) -> pulumi.Input[str]:
        """
        Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        """
        return pulumi.get(self, "next_type")

    @next_type.setter
    def next_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_type", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Input[str]:
        """
        ID of routing table to which this entry belongs.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the routing table entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the entry is disabled, default is `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)


@pulumi.input_type
class _TableEntryState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 next_hub: Optional[pulumi.Input[str]] = None,
                 next_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TableEntry resources.
        :param pulumi.Input[str] description: Description of the routing table entry.
        :param pulumi.Input[str] destination_cidr_block: Destination address block.
        :param pulumi.Input[bool] disabled: Whether the entry is disabled, default is `false`.
        :param pulumi.Input[str] next_hub: ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        :param pulumi.Input[str] next_type: Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        :param pulumi.Input[str] route_table_id: ID of routing table to which this entry belongs.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if next_hub is not None:
            pulumi.set(__self__, "next_hub", next_hub)
        if next_type is not None:
            pulumi.set(__self__, "next_type", next_type)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the routing table entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        Destination address block.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the entry is disabled, default is `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="nextHub")
    def next_hub(self) -> Optional[pulumi.Input[str]]:
        """
        ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        """
        return pulumi.get(self, "next_hub")

    @next_hub.setter
    def next_hub(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hub", value)

    @property
    @pulumi.getter(name="nextType")
    def next_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        """
        return pulumi.get(self, "next_type")

    @next_type.setter
    def next_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_type", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of routing table to which this entry belongs.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)


class TableEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 next_hub: Optional[pulumi.Input[str]] = None,
                 next_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create an entry of a routing table.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "na-siliconvalley-1"
        foo_instance = tencentcloud.vpc.Instance("fooInstance", cidr_block="10.0.0.0/16")
        foo_table = tencentcloud.route.Table("fooTable", vpc_id=foo_instance.id)
        foo_subnet_instance_instance = tencentcloud.subnet.Instance("fooSubnet/instanceInstance",
            vpc_id=foo_instance.id,
            cidr_block="10.0.12.0/24",
            availability_zone=availability_zone,
            route_table_id=foo_table.id)
        instance = tencentcloud.route.TableEntry("instance",
            route_table_id=foo_table.id,
            destination_cidr_block="10.4.4.0/24",
            next_type="EIP",
            next_hub="0",
            description="ci-test-route-table-entry")
        ```

        ## Import

        Route table entry can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Route/tableEntry:TableEntry foo 83517.rtb-mlhpg09u
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the routing table entry.
        :param pulumi.Input[str] destination_cidr_block: Destination address block.
        :param pulumi.Input[bool] disabled: Whether the entry is disabled, default is `false`.
        :param pulumi.Input[str] next_hub: ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        :param pulumi.Input[str] next_type: Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        :param pulumi.Input[str] route_table_id: ID of routing table to which this entry belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TableEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create an entry of a routing table.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "na-siliconvalley-1"
        foo_instance = tencentcloud.vpc.Instance("fooInstance", cidr_block="10.0.0.0/16")
        foo_table = tencentcloud.route.Table("fooTable", vpc_id=foo_instance.id)
        foo_subnet_instance_instance = tencentcloud.subnet.Instance("fooSubnet/instanceInstance",
            vpc_id=foo_instance.id,
            cidr_block="10.0.12.0/24",
            availability_zone=availability_zone,
            route_table_id=foo_table.id)
        instance = tencentcloud.route.TableEntry("instance",
            route_table_id=foo_table.id,
            destination_cidr_block="10.4.4.0/24",
            next_type="EIP",
            next_hub="0",
            description="ci-test-route-table-entry")
        ```

        ## Import

        Route table entry can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Route/tableEntry:TableEntry foo 83517.rtb-mlhpg09u
        ```

        :param str resource_name: The name of the resource.
        :param TableEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TableEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 next_hub: Optional[pulumi.Input[str]] = None,
                 next_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TableEntryArgs.__new__(TableEntryArgs)

            __props__.__dict__["description"] = description
            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["disabled"] = disabled
            if next_hub is None and not opts.urn:
                raise TypeError("Missing required property 'next_hub'")
            __props__.__dict__["next_hub"] = next_hub
            if next_type is None and not opts.urn:
                raise TypeError("Missing required property 'next_type'")
            __props__.__dict__["next_type"] = next_type
            if route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_id'")
            __props__.__dict__["route_table_id"] = route_table_id
        super(TableEntry, __self__).__init__(
            'tencentcloud:Route/tableEntry:TableEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            next_hub: Optional[pulumi.Input[str]] = None,
            next_type: Optional[pulumi.Input[str]] = None,
            route_table_id: Optional[pulumi.Input[str]] = None) -> 'TableEntry':
        """
        Get an existing TableEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the routing table entry.
        :param pulumi.Input[str] destination_cidr_block: Destination address block.
        :param pulumi.Input[bool] disabled: Whether the entry is disabled, default is `false`.
        :param pulumi.Input[str] next_hub: ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        :param pulumi.Input[str] next_type: Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        :param pulumi.Input[str] route_table_id: ID of routing table to which this entry belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TableEntryState.__new__(_TableEntryState)

        __props__.__dict__["description"] = description
        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["next_hub"] = next_hub
        __props__.__dict__["next_type"] = next_type
        __props__.__dict__["route_table_id"] = route_table_id
        return TableEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the routing table entry.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        Destination address block.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the entry is disabled, default is `false`.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="nextHub")
    def next_hub(self) -> pulumi.Output[str]:
        """
        ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        """
        return pulumi.get(self, "next_hub")

    @property
    @pulumi.getter(name="nextType")
    def next_type(self) -> pulumi.Output[str]:
        """
        Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
        """
        return pulumi.get(self, "next_type")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[str]:
        """
        ID of routing table to which this entry belongs.
        """
        return pulumi.get(self, "route_table_id")

