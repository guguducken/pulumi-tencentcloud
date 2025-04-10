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

__all__ = ['ExtraConfigArgs', 'ExtraConfig']

@pulumi.input_type
class ExtraConfigArgs:
    def __init__(__self__, *,
                 direct_connect_tunnel_id: pulumi.Input[str],
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bfd_enable: Optional[pulumi.Input[int]] = None,
                 bfd_info: Optional[pulumi.Input['ExtraConfigBfdInfoArgs']] = None,
                 bgp_peer: Optional[pulumi.Input['ExtraConfigBgpPeerArgs']] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 enable_bgp_community: Optional[pulumi.Input[bool]] = None,
                 ipv6_enable: Optional[pulumi.Input[int]] = None,
                 jumbo_enable: Optional[pulumi.Input[int]] = None,
                 nqa_enable: Optional[pulumi.Input[int]] = None,
                 nqa_info: Optional[pulumi.Input['ExtraConfigNqaInfoArgs']] = None,
                 route_filter_prefixes: Optional[pulumi.Input['ExtraConfigRouteFilterPrefixesArgs']] = None,
                 tencent_address: Optional[pulumi.Input[str]] = None,
                 tencent_backup_address: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ExtraConfig resource.
        :param pulumi.Input[str] direct_connect_tunnel_id: direct connect tunnel id.
        :param pulumi.Input[int] bandwidth: direct connect tunnel bandwidth.
        :param pulumi.Input[int] bfd_enable: be enabled BFD.
        :param pulumi.Input['ExtraConfigBfdInfoArgs'] bfd_info: BFD config info.
        :param pulumi.Input['ExtraConfigBgpPeerArgs'] bgp_peer: idc BGP, Asn, AuthKey.
        :param pulumi.Input[str] customer_address: direct connect tunnel user idc connect ip.
        :param pulumi.Input[bool] enable_bgp_community: BGP community attribute.
        :param pulumi.Input[int] ipv6_enable: 0: disable IPv61: enable IPv6.
        :param pulumi.Input[int] jumbo_enable: direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        :param pulumi.Input[int] nqa_enable: be enabled NQA.
        :param pulumi.Input['ExtraConfigNqaInfoArgs'] nqa_info: NQA config info.
        :param pulumi.Input['ExtraConfigRouteFilterPrefixesArgs'] route_filter_prefixes: user filter network prefixes.
        :param pulumi.Input[str] tencent_address: direct connect tunnel tencent cloud connect ip.
        :param pulumi.Input[str] tencent_backup_address: direct connect tunnel tencent cloud backup connect ip.
        :param pulumi.Input[int] vlan: direct connect tunnel vlan id.
        """
        pulumi.set(__self__, "direct_connect_tunnel_id", direct_connect_tunnel_id)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if bfd_enable is not None:
            pulumi.set(__self__, "bfd_enable", bfd_enable)
        if bfd_info is not None:
            pulumi.set(__self__, "bfd_info", bfd_info)
        if bgp_peer is not None:
            pulumi.set(__self__, "bgp_peer", bgp_peer)
        if customer_address is not None:
            pulumi.set(__self__, "customer_address", customer_address)
        if enable_bgp_community is not None:
            pulumi.set(__self__, "enable_bgp_community", enable_bgp_community)
        if ipv6_enable is not None:
            pulumi.set(__self__, "ipv6_enable", ipv6_enable)
        if jumbo_enable is not None:
            pulumi.set(__self__, "jumbo_enable", jumbo_enable)
        if nqa_enable is not None:
            pulumi.set(__self__, "nqa_enable", nqa_enable)
        if nqa_info is not None:
            pulumi.set(__self__, "nqa_info", nqa_info)
        if route_filter_prefixes is not None:
            pulumi.set(__self__, "route_filter_prefixes", route_filter_prefixes)
        if tencent_address is not None:
            pulumi.set(__self__, "tencent_address", tencent_address)
        if tencent_backup_address is not None:
            pulumi.set(__self__, "tencent_backup_address", tencent_backup_address)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter(name="directConnectTunnelId")
    def direct_connect_tunnel_id(self) -> pulumi.Input[str]:
        """
        direct connect tunnel id.
        """
        return pulumi.get(self, "direct_connect_tunnel_id")

    @direct_connect_tunnel_id.setter
    def direct_connect_tunnel_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "direct_connect_tunnel_id", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        direct connect tunnel bandwidth.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="bfdEnable")
    def bfd_enable(self) -> Optional[pulumi.Input[int]]:
        """
        be enabled BFD.
        """
        return pulumi.get(self, "bfd_enable")

    @bfd_enable.setter
    def bfd_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bfd_enable", value)

    @property
    @pulumi.getter(name="bfdInfo")
    def bfd_info(self) -> Optional[pulumi.Input['ExtraConfigBfdInfoArgs']]:
        """
        BFD config info.
        """
        return pulumi.get(self, "bfd_info")

    @bfd_info.setter
    def bfd_info(self, value: Optional[pulumi.Input['ExtraConfigBfdInfoArgs']]):
        pulumi.set(self, "bfd_info", value)

    @property
    @pulumi.getter(name="bgpPeer")
    def bgp_peer(self) -> Optional[pulumi.Input['ExtraConfigBgpPeerArgs']]:
        """
        idc BGP, Asn, AuthKey.
        """
        return pulumi.get(self, "bgp_peer")

    @bgp_peer.setter
    def bgp_peer(self, value: Optional[pulumi.Input['ExtraConfigBgpPeerArgs']]):
        pulumi.set(self, "bgp_peer", value)

    @property
    @pulumi.getter(name="customerAddress")
    def customer_address(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel user idc connect ip.
        """
        return pulumi.get(self, "customer_address")

    @customer_address.setter
    def customer_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_address", value)

    @property
    @pulumi.getter(name="enableBgpCommunity")
    def enable_bgp_community(self) -> Optional[pulumi.Input[bool]]:
        """
        BGP community attribute.
        """
        return pulumi.get(self, "enable_bgp_community")

    @enable_bgp_community.setter
    def enable_bgp_community(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_bgp_community", value)

    @property
    @pulumi.getter(name="ipv6Enable")
    def ipv6_enable(self) -> Optional[pulumi.Input[int]]:
        """
        0: disable IPv61: enable IPv6.
        """
        return pulumi.get(self, "ipv6_enable")

    @ipv6_enable.setter
    def ipv6_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_enable", value)

    @property
    @pulumi.getter(name="jumboEnable")
    def jumbo_enable(self) -> Optional[pulumi.Input[int]]:
        """
        direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        """
        return pulumi.get(self, "jumbo_enable")

    @jumbo_enable.setter
    def jumbo_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "jumbo_enable", value)

    @property
    @pulumi.getter(name="nqaEnable")
    def nqa_enable(self) -> Optional[pulumi.Input[int]]:
        """
        be enabled NQA.
        """
        return pulumi.get(self, "nqa_enable")

    @nqa_enable.setter
    def nqa_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nqa_enable", value)

    @property
    @pulumi.getter(name="nqaInfo")
    def nqa_info(self) -> Optional[pulumi.Input['ExtraConfigNqaInfoArgs']]:
        """
        NQA config info.
        """
        return pulumi.get(self, "nqa_info")

    @nqa_info.setter
    def nqa_info(self, value: Optional[pulumi.Input['ExtraConfigNqaInfoArgs']]):
        pulumi.set(self, "nqa_info", value)

    @property
    @pulumi.getter(name="routeFilterPrefixes")
    def route_filter_prefixes(self) -> Optional[pulumi.Input['ExtraConfigRouteFilterPrefixesArgs']]:
        """
        user filter network prefixes.
        """
        return pulumi.get(self, "route_filter_prefixes")

    @route_filter_prefixes.setter
    def route_filter_prefixes(self, value: Optional[pulumi.Input['ExtraConfigRouteFilterPrefixesArgs']]):
        pulumi.set(self, "route_filter_prefixes", value)

    @property
    @pulumi.getter(name="tencentAddress")
    def tencent_address(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel tencent cloud connect ip.
        """
        return pulumi.get(self, "tencent_address")

    @tencent_address.setter
    def tencent_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tencent_address", value)

    @property
    @pulumi.getter(name="tencentBackupAddress")
    def tencent_backup_address(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel tencent cloud backup connect ip.
        """
        return pulumi.get(self, "tencent_backup_address")

    @tencent_backup_address.setter
    def tencent_backup_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tencent_backup_address", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        direct connect tunnel vlan id.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)


@pulumi.input_type
class _ExtraConfigState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bfd_enable: Optional[pulumi.Input[int]] = None,
                 bfd_info: Optional[pulumi.Input['ExtraConfigBfdInfoArgs']] = None,
                 bgp_peer: Optional[pulumi.Input['ExtraConfigBgpPeerArgs']] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 direct_connect_tunnel_id: Optional[pulumi.Input[str]] = None,
                 enable_bgp_community: Optional[pulumi.Input[bool]] = None,
                 ipv6_enable: Optional[pulumi.Input[int]] = None,
                 jumbo_enable: Optional[pulumi.Input[int]] = None,
                 nqa_enable: Optional[pulumi.Input[int]] = None,
                 nqa_info: Optional[pulumi.Input['ExtraConfigNqaInfoArgs']] = None,
                 route_filter_prefixes: Optional[pulumi.Input['ExtraConfigRouteFilterPrefixesArgs']] = None,
                 tencent_address: Optional[pulumi.Input[str]] = None,
                 tencent_backup_address: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ExtraConfig resources.
        :param pulumi.Input[int] bandwidth: direct connect tunnel bandwidth.
        :param pulumi.Input[int] bfd_enable: be enabled BFD.
        :param pulumi.Input['ExtraConfigBfdInfoArgs'] bfd_info: BFD config info.
        :param pulumi.Input['ExtraConfigBgpPeerArgs'] bgp_peer: idc BGP, Asn, AuthKey.
        :param pulumi.Input[str] customer_address: direct connect tunnel user idc connect ip.
        :param pulumi.Input[str] direct_connect_tunnel_id: direct connect tunnel id.
        :param pulumi.Input[bool] enable_bgp_community: BGP community attribute.
        :param pulumi.Input[int] ipv6_enable: 0: disable IPv61: enable IPv6.
        :param pulumi.Input[int] jumbo_enable: direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        :param pulumi.Input[int] nqa_enable: be enabled NQA.
        :param pulumi.Input['ExtraConfigNqaInfoArgs'] nqa_info: NQA config info.
        :param pulumi.Input['ExtraConfigRouteFilterPrefixesArgs'] route_filter_prefixes: user filter network prefixes.
        :param pulumi.Input[str] tencent_address: direct connect tunnel tencent cloud connect ip.
        :param pulumi.Input[str] tencent_backup_address: direct connect tunnel tencent cloud backup connect ip.
        :param pulumi.Input[int] vlan: direct connect tunnel vlan id.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if bfd_enable is not None:
            pulumi.set(__self__, "bfd_enable", bfd_enable)
        if bfd_info is not None:
            pulumi.set(__self__, "bfd_info", bfd_info)
        if bgp_peer is not None:
            pulumi.set(__self__, "bgp_peer", bgp_peer)
        if customer_address is not None:
            pulumi.set(__self__, "customer_address", customer_address)
        if direct_connect_tunnel_id is not None:
            pulumi.set(__self__, "direct_connect_tunnel_id", direct_connect_tunnel_id)
        if enable_bgp_community is not None:
            pulumi.set(__self__, "enable_bgp_community", enable_bgp_community)
        if ipv6_enable is not None:
            pulumi.set(__self__, "ipv6_enable", ipv6_enable)
        if jumbo_enable is not None:
            pulumi.set(__self__, "jumbo_enable", jumbo_enable)
        if nqa_enable is not None:
            pulumi.set(__self__, "nqa_enable", nqa_enable)
        if nqa_info is not None:
            pulumi.set(__self__, "nqa_info", nqa_info)
        if route_filter_prefixes is not None:
            pulumi.set(__self__, "route_filter_prefixes", route_filter_prefixes)
        if tencent_address is not None:
            pulumi.set(__self__, "tencent_address", tencent_address)
        if tencent_backup_address is not None:
            pulumi.set(__self__, "tencent_backup_address", tencent_backup_address)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        direct connect tunnel bandwidth.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="bfdEnable")
    def bfd_enable(self) -> Optional[pulumi.Input[int]]:
        """
        be enabled BFD.
        """
        return pulumi.get(self, "bfd_enable")

    @bfd_enable.setter
    def bfd_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bfd_enable", value)

    @property
    @pulumi.getter(name="bfdInfo")
    def bfd_info(self) -> Optional[pulumi.Input['ExtraConfigBfdInfoArgs']]:
        """
        BFD config info.
        """
        return pulumi.get(self, "bfd_info")

    @bfd_info.setter
    def bfd_info(self, value: Optional[pulumi.Input['ExtraConfigBfdInfoArgs']]):
        pulumi.set(self, "bfd_info", value)

    @property
    @pulumi.getter(name="bgpPeer")
    def bgp_peer(self) -> Optional[pulumi.Input['ExtraConfigBgpPeerArgs']]:
        """
        idc BGP, Asn, AuthKey.
        """
        return pulumi.get(self, "bgp_peer")

    @bgp_peer.setter
    def bgp_peer(self, value: Optional[pulumi.Input['ExtraConfigBgpPeerArgs']]):
        pulumi.set(self, "bgp_peer", value)

    @property
    @pulumi.getter(name="customerAddress")
    def customer_address(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel user idc connect ip.
        """
        return pulumi.get(self, "customer_address")

    @customer_address.setter
    def customer_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_address", value)

    @property
    @pulumi.getter(name="directConnectTunnelId")
    def direct_connect_tunnel_id(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel id.
        """
        return pulumi.get(self, "direct_connect_tunnel_id")

    @direct_connect_tunnel_id.setter
    def direct_connect_tunnel_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direct_connect_tunnel_id", value)

    @property
    @pulumi.getter(name="enableBgpCommunity")
    def enable_bgp_community(self) -> Optional[pulumi.Input[bool]]:
        """
        BGP community attribute.
        """
        return pulumi.get(self, "enable_bgp_community")

    @enable_bgp_community.setter
    def enable_bgp_community(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_bgp_community", value)

    @property
    @pulumi.getter(name="ipv6Enable")
    def ipv6_enable(self) -> Optional[pulumi.Input[int]]:
        """
        0: disable IPv61: enable IPv6.
        """
        return pulumi.get(self, "ipv6_enable")

    @ipv6_enable.setter
    def ipv6_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_enable", value)

    @property
    @pulumi.getter(name="jumboEnable")
    def jumbo_enable(self) -> Optional[pulumi.Input[int]]:
        """
        direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        """
        return pulumi.get(self, "jumbo_enable")

    @jumbo_enable.setter
    def jumbo_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "jumbo_enable", value)

    @property
    @pulumi.getter(name="nqaEnable")
    def nqa_enable(self) -> Optional[pulumi.Input[int]]:
        """
        be enabled NQA.
        """
        return pulumi.get(self, "nqa_enable")

    @nqa_enable.setter
    def nqa_enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nqa_enable", value)

    @property
    @pulumi.getter(name="nqaInfo")
    def nqa_info(self) -> Optional[pulumi.Input['ExtraConfigNqaInfoArgs']]:
        """
        NQA config info.
        """
        return pulumi.get(self, "nqa_info")

    @nqa_info.setter
    def nqa_info(self, value: Optional[pulumi.Input['ExtraConfigNqaInfoArgs']]):
        pulumi.set(self, "nqa_info", value)

    @property
    @pulumi.getter(name="routeFilterPrefixes")
    def route_filter_prefixes(self) -> Optional[pulumi.Input['ExtraConfigRouteFilterPrefixesArgs']]:
        """
        user filter network prefixes.
        """
        return pulumi.get(self, "route_filter_prefixes")

    @route_filter_prefixes.setter
    def route_filter_prefixes(self, value: Optional[pulumi.Input['ExtraConfigRouteFilterPrefixesArgs']]):
        pulumi.set(self, "route_filter_prefixes", value)

    @property
    @pulumi.getter(name="tencentAddress")
    def tencent_address(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel tencent cloud connect ip.
        """
        return pulumi.get(self, "tencent_address")

    @tencent_address.setter
    def tencent_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tencent_address", value)

    @property
    @pulumi.getter(name="tencentBackupAddress")
    def tencent_backup_address(self) -> Optional[pulumi.Input[str]]:
        """
        direct connect tunnel tencent cloud backup connect ip.
        """
        return pulumi.get(self, "tencent_backup_address")

    @tencent_backup_address.setter
    def tencent_backup_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tencent_backup_address", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        direct connect tunnel vlan id.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)


class ExtraConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bfd_enable: Optional[pulumi.Input[int]] = None,
                 bfd_info: Optional[pulumi.Input[pulumi.InputType['ExtraConfigBfdInfoArgs']]] = None,
                 bgp_peer: Optional[pulumi.Input[pulumi.InputType['ExtraConfigBgpPeerArgs']]] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 direct_connect_tunnel_id: Optional[pulumi.Input[str]] = None,
                 enable_bgp_community: Optional[pulumi.Input[bool]] = None,
                 ipv6_enable: Optional[pulumi.Input[int]] = None,
                 jumbo_enable: Optional[pulumi.Input[int]] = None,
                 nqa_enable: Optional[pulumi.Input[int]] = None,
                 nqa_info: Optional[pulumi.Input[pulumi.InputType['ExtraConfigNqaInfoArgs']]] = None,
                 route_filter_prefixes: Optional[pulumi.Input[pulumi.InputType['ExtraConfigRouteFilterPrefixesArgs']]] = None,
                 tencent_address: Optional[pulumi.Input[str]] = None,
                 tencent_backup_address: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a dc dcx_extra_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        dcx_extra_config = tencentcloud.dcx.ExtraConfig("dcxExtraConfig",
            bandwidth=10,
            bfd_enable=0,
            bfd_info=tencentcloud.dcx.ExtraConfigBfdInfoArgs(
                interval=100,
                probe_failed_times=3,
            ),
            bgp_peer=tencentcloud.dcx.ExtraConfigBgpPeerArgs(
                asn=65101,
                auth_key="test123",
            ),
            customer_address="192.168.1.4",
            direct_connect_tunnel_id="dcx-4z49tnws",
            enable_bgp_community=False,
            ipv6_enable=0,
            jumbo_enable=0,
            nqa_enable=1,
            nqa_info=tencentcloud.dcx.ExtraConfigNqaInfoArgs(
                destination_ip="192.168.2.2",
                interval=100,
                probe_failed_times=3,
            ),
            route_filter_prefixes=tencentcloud.dcx.ExtraConfigRouteFilterPrefixesArgs(
                cidr="192.168.0.0/24",
            ),
            tencent_address="192.168.1.1",
            tencent_backup_address="192.168.1.2",
            vlan=123)
        ```

        ## Import

        dc dcx_extra_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dcx/extraConfig:ExtraConfig dcx_extra_config dcx_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: direct connect tunnel bandwidth.
        :param pulumi.Input[int] bfd_enable: be enabled BFD.
        :param pulumi.Input[pulumi.InputType['ExtraConfigBfdInfoArgs']] bfd_info: BFD config info.
        :param pulumi.Input[pulumi.InputType['ExtraConfigBgpPeerArgs']] bgp_peer: idc BGP, Asn, AuthKey.
        :param pulumi.Input[str] customer_address: direct connect tunnel user idc connect ip.
        :param pulumi.Input[str] direct_connect_tunnel_id: direct connect tunnel id.
        :param pulumi.Input[bool] enable_bgp_community: BGP community attribute.
        :param pulumi.Input[int] ipv6_enable: 0: disable IPv61: enable IPv6.
        :param pulumi.Input[int] jumbo_enable: direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        :param pulumi.Input[int] nqa_enable: be enabled NQA.
        :param pulumi.Input[pulumi.InputType['ExtraConfigNqaInfoArgs']] nqa_info: NQA config info.
        :param pulumi.Input[pulumi.InputType['ExtraConfigRouteFilterPrefixesArgs']] route_filter_prefixes: user filter network prefixes.
        :param pulumi.Input[str] tencent_address: direct connect tunnel tencent cloud connect ip.
        :param pulumi.Input[str] tencent_backup_address: direct connect tunnel tencent cloud backup connect ip.
        :param pulumi.Input[int] vlan: direct connect tunnel vlan id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExtraConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dc dcx_extra_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        dcx_extra_config = tencentcloud.dcx.ExtraConfig("dcxExtraConfig",
            bandwidth=10,
            bfd_enable=0,
            bfd_info=tencentcloud.dcx.ExtraConfigBfdInfoArgs(
                interval=100,
                probe_failed_times=3,
            ),
            bgp_peer=tencentcloud.dcx.ExtraConfigBgpPeerArgs(
                asn=65101,
                auth_key="test123",
            ),
            customer_address="192.168.1.4",
            direct_connect_tunnel_id="dcx-4z49tnws",
            enable_bgp_community=False,
            ipv6_enable=0,
            jumbo_enable=0,
            nqa_enable=1,
            nqa_info=tencentcloud.dcx.ExtraConfigNqaInfoArgs(
                destination_ip="192.168.2.2",
                interval=100,
                probe_failed_times=3,
            ),
            route_filter_prefixes=tencentcloud.dcx.ExtraConfigRouteFilterPrefixesArgs(
                cidr="192.168.0.0/24",
            ),
            tencent_address="192.168.1.1",
            tencent_backup_address="192.168.1.2",
            vlan=123)
        ```

        ## Import

        dc dcx_extra_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dcx/extraConfig:ExtraConfig dcx_extra_config dcx_id
        ```

        :param str resource_name: The name of the resource.
        :param ExtraConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExtraConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bfd_enable: Optional[pulumi.Input[int]] = None,
                 bfd_info: Optional[pulumi.Input[pulumi.InputType['ExtraConfigBfdInfoArgs']]] = None,
                 bgp_peer: Optional[pulumi.Input[pulumi.InputType['ExtraConfigBgpPeerArgs']]] = None,
                 customer_address: Optional[pulumi.Input[str]] = None,
                 direct_connect_tunnel_id: Optional[pulumi.Input[str]] = None,
                 enable_bgp_community: Optional[pulumi.Input[bool]] = None,
                 ipv6_enable: Optional[pulumi.Input[int]] = None,
                 jumbo_enable: Optional[pulumi.Input[int]] = None,
                 nqa_enable: Optional[pulumi.Input[int]] = None,
                 nqa_info: Optional[pulumi.Input[pulumi.InputType['ExtraConfigNqaInfoArgs']]] = None,
                 route_filter_prefixes: Optional[pulumi.Input[pulumi.InputType['ExtraConfigRouteFilterPrefixesArgs']]] = None,
                 tencent_address: Optional[pulumi.Input[str]] = None,
                 tencent_backup_address: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
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
            __props__ = ExtraConfigArgs.__new__(ExtraConfigArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["bfd_enable"] = bfd_enable
            __props__.__dict__["bfd_info"] = bfd_info
            __props__.__dict__["bgp_peer"] = bgp_peer
            __props__.__dict__["customer_address"] = customer_address
            if direct_connect_tunnel_id is None and not opts.urn:
                raise TypeError("Missing required property 'direct_connect_tunnel_id'")
            __props__.__dict__["direct_connect_tunnel_id"] = direct_connect_tunnel_id
            __props__.__dict__["enable_bgp_community"] = enable_bgp_community
            __props__.__dict__["ipv6_enable"] = ipv6_enable
            __props__.__dict__["jumbo_enable"] = jumbo_enable
            __props__.__dict__["nqa_enable"] = nqa_enable
            __props__.__dict__["nqa_info"] = nqa_info
            __props__.__dict__["route_filter_prefixes"] = route_filter_prefixes
            __props__.__dict__["tencent_address"] = tencent_address
            __props__.__dict__["tencent_backup_address"] = tencent_backup_address
            __props__.__dict__["vlan"] = vlan
        super(ExtraConfig, __self__).__init__(
            'tencentcloud:Dcx/extraConfig:ExtraConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            bfd_enable: Optional[pulumi.Input[int]] = None,
            bfd_info: Optional[pulumi.Input[pulumi.InputType['ExtraConfigBfdInfoArgs']]] = None,
            bgp_peer: Optional[pulumi.Input[pulumi.InputType['ExtraConfigBgpPeerArgs']]] = None,
            customer_address: Optional[pulumi.Input[str]] = None,
            direct_connect_tunnel_id: Optional[pulumi.Input[str]] = None,
            enable_bgp_community: Optional[pulumi.Input[bool]] = None,
            ipv6_enable: Optional[pulumi.Input[int]] = None,
            jumbo_enable: Optional[pulumi.Input[int]] = None,
            nqa_enable: Optional[pulumi.Input[int]] = None,
            nqa_info: Optional[pulumi.Input[pulumi.InputType['ExtraConfigNqaInfoArgs']]] = None,
            route_filter_prefixes: Optional[pulumi.Input[pulumi.InputType['ExtraConfigRouteFilterPrefixesArgs']]] = None,
            tencent_address: Optional[pulumi.Input[str]] = None,
            tencent_backup_address: Optional[pulumi.Input[str]] = None,
            vlan: Optional[pulumi.Input[int]] = None) -> 'ExtraConfig':
        """
        Get an existing ExtraConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: direct connect tunnel bandwidth.
        :param pulumi.Input[int] bfd_enable: be enabled BFD.
        :param pulumi.Input[pulumi.InputType['ExtraConfigBfdInfoArgs']] bfd_info: BFD config info.
        :param pulumi.Input[pulumi.InputType['ExtraConfigBgpPeerArgs']] bgp_peer: idc BGP, Asn, AuthKey.
        :param pulumi.Input[str] customer_address: direct connect tunnel user idc connect ip.
        :param pulumi.Input[str] direct_connect_tunnel_id: direct connect tunnel id.
        :param pulumi.Input[bool] enable_bgp_community: BGP community attribute.
        :param pulumi.Input[int] ipv6_enable: 0: disable IPv61: enable IPv6.
        :param pulumi.Input[int] jumbo_enable: direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        :param pulumi.Input[int] nqa_enable: be enabled NQA.
        :param pulumi.Input[pulumi.InputType['ExtraConfigNqaInfoArgs']] nqa_info: NQA config info.
        :param pulumi.Input[pulumi.InputType['ExtraConfigRouteFilterPrefixesArgs']] route_filter_prefixes: user filter network prefixes.
        :param pulumi.Input[str] tencent_address: direct connect tunnel tencent cloud connect ip.
        :param pulumi.Input[str] tencent_backup_address: direct connect tunnel tencent cloud backup connect ip.
        :param pulumi.Input[int] vlan: direct connect tunnel vlan id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExtraConfigState.__new__(_ExtraConfigState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["bfd_enable"] = bfd_enable
        __props__.__dict__["bfd_info"] = bfd_info
        __props__.__dict__["bgp_peer"] = bgp_peer
        __props__.__dict__["customer_address"] = customer_address
        __props__.__dict__["direct_connect_tunnel_id"] = direct_connect_tunnel_id
        __props__.__dict__["enable_bgp_community"] = enable_bgp_community
        __props__.__dict__["ipv6_enable"] = ipv6_enable
        __props__.__dict__["jumbo_enable"] = jumbo_enable
        __props__.__dict__["nqa_enable"] = nqa_enable
        __props__.__dict__["nqa_info"] = nqa_info
        __props__.__dict__["route_filter_prefixes"] = route_filter_prefixes
        __props__.__dict__["tencent_address"] = tencent_address
        __props__.__dict__["tencent_backup_address"] = tencent_backup_address
        __props__.__dict__["vlan"] = vlan
        return ExtraConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        direct connect tunnel bandwidth.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="bfdEnable")
    def bfd_enable(self) -> pulumi.Output[Optional[int]]:
        """
        be enabled BFD.
        """
        return pulumi.get(self, "bfd_enable")

    @property
    @pulumi.getter(name="bfdInfo")
    def bfd_info(self) -> pulumi.Output[Optional['outputs.ExtraConfigBfdInfo']]:
        """
        BFD config info.
        """
        return pulumi.get(self, "bfd_info")

    @property
    @pulumi.getter(name="bgpPeer")
    def bgp_peer(self) -> pulumi.Output[Optional['outputs.ExtraConfigBgpPeer']]:
        """
        idc BGP, Asn, AuthKey.
        """
        return pulumi.get(self, "bgp_peer")

    @property
    @pulumi.getter(name="customerAddress")
    def customer_address(self) -> pulumi.Output[Optional[str]]:
        """
        direct connect tunnel user idc connect ip.
        """
        return pulumi.get(self, "customer_address")

    @property
    @pulumi.getter(name="directConnectTunnelId")
    def direct_connect_tunnel_id(self) -> pulumi.Output[str]:
        """
        direct connect tunnel id.
        """
        return pulumi.get(self, "direct_connect_tunnel_id")

    @property
    @pulumi.getter(name="enableBgpCommunity")
    def enable_bgp_community(self) -> pulumi.Output[Optional[bool]]:
        """
        BGP community attribute.
        """
        return pulumi.get(self, "enable_bgp_community")

    @property
    @pulumi.getter(name="ipv6Enable")
    def ipv6_enable(self) -> pulumi.Output[Optional[int]]:
        """
        0: disable IPv61: enable IPv6.
        """
        return pulumi.get(self, "ipv6_enable")

    @property
    @pulumi.getter(name="jumboEnable")
    def jumbo_enable(self) -> pulumi.Output[Optional[int]]:
        """
        direct connect tunnel support jumbo frame1: enable direct connect tunnel jumbo frame0: disable direct connect tunnel jumbo frame.
        """
        return pulumi.get(self, "jumbo_enable")

    @property
    @pulumi.getter(name="nqaEnable")
    def nqa_enable(self) -> pulumi.Output[Optional[int]]:
        """
        be enabled NQA.
        """
        return pulumi.get(self, "nqa_enable")

    @property
    @pulumi.getter(name="nqaInfo")
    def nqa_info(self) -> pulumi.Output[Optional['outputs.ExtraConfigNqaInfo']]:
        """
        NQA config info.
        """
        return pulumi.get(self, "nqa_info")

    @property
    @pulumi.getter(name="routeFilterPrefixes")
    def route_filter_prefixes(self) -> pulumi.Output[Optional['outputs.ExtraConfigRouteFilterPrefixes']]:
        """
        user filter network prefixes.
        """
        return pulumi.get(self, "route_filter_prefixes")

    @property
    @pulumi.getter(name="tencentAddress")
    def tencent_address(self) -> pulumi.Output[Optional[str]]:
        """
        direct connect tunnel tencent cloud connect ip.
        """
        return pulumi.get(self, "tencent_address")

    @property
    @pulumi.getter(name="tencentBackupAddress")
    def tencent_backup_address(self) -> pulumi.Output[Optional[str]]:
        """
        direct connect tunnel tencent cloud backup connect ip.
        """
        return pulumi.get(self, "tencent_backup_address")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[Optional[int]]:
        """
        direct connect tunnel vlan id.
        """
        return pulumi.get(self, "vlan")

