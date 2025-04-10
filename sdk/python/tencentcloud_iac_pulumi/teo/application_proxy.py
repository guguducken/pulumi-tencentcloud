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

__all__ = ['ApplicationProxyArgs', 'ApplicationProxy']

@pulumi.input_type
class ApplicationProxyArgs:
    def __init__(__self__, *,
                 accelerate_type: pulumi.Input[int],
                 plat_type: pulumi.Input[str],
                 proxy_name: pulumi.Input[str],
                 security_type: pulumi.Input[int],
                 zone_id: pulumi.Input[str],
                 ipv6: Optional[pulumi.Input['ApplicationProxyIpv6Args']] = None,
                 proxy_type: Optional[pulumi.Input[str]] = None,
                 session_persist_time: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationProxy resource.
        :param pulumi.Input[int] accelerate_type: - `0`: Disable acceleration.- `1`: Enable acceleration.
        :param pulumi.Input[str] plat_type: Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        :param pulumi.Input[str] proxy_name: When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        :param pulumi.Input[int] security_type: - `0`: Disable security protection.- `1`: Enable security protection.
        :param pulumi.Input[str] zone_id: Site ID.
        :param pulumi.Input['ApplicationProxyIpv6Args'] ipv6: IPv6 access configuration.
        :param pulumi.Input[str] proxy_type: Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        :param pulumi.Input[int] session_persist_time: Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        :param pulumi.Input[str] status: Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        """
        pulumi.set(__self__, "accelerate_type", accelerate_type)
        pulumi.set(__self__, "plat_type", plat_type)
        pulumi.set(__self__, "proxy_name", proxy_name)
        pulumi.set(__self__, "security_type", security_type)
        pulumi.set(__self__, "zone_id", zone_id)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if proxy_type is not None:
            pulumi.set(__self__, "proxy_type", proxy_type)
        if session_persist_time is not None:
            pulumi.set(__self__, "session_persist_time", session_persist_time)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accelerateType")
    def accelerate_type(self) -> pulumi.Input[int]:
        """
        - `0`: Disable acceleration.- `1`: Enable acceleration.
        """
        return pulumi.get(self, "accelerate_type")

    @accelerate_type.setter
    def accelerate_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "accelerate_type", value)

    @property
    @pulumi.getter(name="platType")
    def plat_type(self) -> pulumi.Input[str]:
        """
        Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        """
        return pulumi.get(self, "plat_type")

    @plat_type.setter
    def plat_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "plat_type", value)

    @property
    @pulumi.getter(name="proxyName")
    def proxy_name(self) -> pulumi.Input[str]:
        """
        When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        """
        return pulumi.get(self, "proxy_name")

    @proxy_name.setter
    def proxy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "proxy_name", value)

    @property
    @pulumi.getter(name="securityType")
    def security_type(self) -> pulumi.Input[int]:
        """
        - `0`: Disable security protection.- `1`: Enable security protection.
        """
        return pulumi.get(self, "security_type")

    @security_type.setter
    def security_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "security_type", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input['ApplicationProxyIpv6Args']]:
        """
        IPv6 access configuration.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input['ApplicationProxyIpv6Args']]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter(name="proxyType")
    def proxy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        """
        return pulumi.get(self, "proxy_type")

    @proxy_type.setter
    def proxy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_type", value)

    @property
    @pulumi.getter(name="sessionPersistTime")
    def session_persist_time(self) -> Optional[pulumi.Input[int]]:
        """
        Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        """
        return pulumi.get(self, "session_persist_time")

    @session_persist_time.setter
    def session_persist_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_persist_time", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _ApplicationProxyState:
    def __init__(__self__, *,
                 accelerate_type: Optional[pulumi.Input[int]] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 ban_status: Optional[pulumi.Input[str]] = None,
                 host_id: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input['ApplicationProxyIpv6Args']] = None,
                 plat_type: Optional[pulumi.Input[str]] = None,
                 proxy_id: Optional[pulumi.Input[str]] = None,
                 proxy_name: Optional[pulumi.Input[str]] = None,
                 proxy_type: Optional[pulumi.Input[str]] = None,
                 schedule_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_type: Optional[pulumi.Input[int]] = None,
                 session_persist_time: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationProxy resources.
        :param pulumi.Input[int] accelerate_type: - `0`: Disable acceleration.- `1`: Enable acceleration.
        :param pulumi.Input[str] area: Acceleration area. Valid values: `mainland`, `overseas`.
        :param pulumi.Input[str] ban_status: Application proxy block status. Valid values: `banned`, `banning`, `recover`, `recovering`.
        :param pulumi.Input[str] host_id: When `ProxyType` is hostname, this field is the ID of the subdomain.
        :param pulumi.Input['ApplicationProxyIpv6Args'] ipv6: IPv6 access configuration.
        :param pulumi.Input[str] plat_type: Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        :param pulumi.Input[str] proxy_id: Proxy ID.
        :param pulumi.Input[str] proxy_name: When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        :param pulumi.Input[str] proxy_type: Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] schedule_values: Scheduling information.
        :param pulumi.Input[int] security_type: - `0`: Disable security protection.- `1`: Enable security protection.
        :param pulumi.Input[int] session_persist_time: Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        :param pulumi.Input[str] status: Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        :param pulumi.Input[str] update_time: Last modification date.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        if accelerate_type is not None:
            pulumi.set(__self__, "accelerate_type", accelerate_type)
        if area is not None:
            pulumi.set(__self__, "area", area)
        if ban_status is not None:
            pulumi.set(__self__, "ban_status", ban_status)
        if host_id is not None:
            pulumi.set(__self__, "host_id", host_id)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if plat_type is not None:
            pulumi.set(__self__, "plat_type", plat_type)
        if proxy_id is not None:
            pulumi.set(__self__, "proxy_id", proxy_id)
        if proxy_name is not None:
            pulumi.set(__self__, "proxy_name", proxy_name)
        if proxy_type is not None:
            pulumi.set(__self__, "proxy_type", proxy_type)
        if schedule_values is not None:
            pulumi.set(__self__, "schedule_values", schedule_values)
        if security_type is not None:
            pulumi.set(__self__, "security_type", security_type)
        if session_persist_time is not None:
            pulumi.set(__self__, "session_persist_time", session_persist_time)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="accelerateType")
    def accelerate_type(self) -> Optional[pulumi.Input[int]]:
        """
        - `0`: Disable acceleration.- `1`: Enable acceleration.
        """
        return pulumi.get(self, "accelerate_type")

    @accelerate_type.setter
    def accelerate_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "accelerate_type", value)

    @property
    @pulumi.getter
    def area(self) -> Optional[pulumi.Input[str]]:
        """
        Acceleration area. Valid values: `mainland`, `overseas`.
        """
        return pulumi.get(self, "area")

    @area.setter
    def area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "area", value)

    @property
    @pulumi.getter(name="banStatus")
    def ban_status(self) -> Optional[pulumi.Input[str]]:
        """
        Application proxy block status. Valid values: `banned`, `banning`, `recover`, `recovering`.
        """
        return pulumi.get(self, "ban_status")

    @ban_status.setter
    def ban_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ban_status", value)

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> Optional[pulumi.Input[str]]:
        """
        When `ProxyType` is hostname, this field is the ID of the subdomain.
        """
        return pulumi.get(self, "host_id")

    @host_id.setter
    def host_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_id", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input['ApplicationProxyIpv6Args']]:
        """
        IPv6 access configuration.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input['ApplicationProxyIpv6Args']]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter(name="platType")
    def plat_type(self) -> Optional[pulumi.Input[str]]:
        """
        Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        """
        return pulumi.get(self, "plat_type")

    @plat_type.setter
    def plat_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plat_type", value)

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Proxy ID.
        """
        return pulumi.get(self, "proxy_id")

    @proxy_id.setter
    def proxy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_id", value)

    @property
    @pulumi.getter(name="proxyName")
    def proxy_name(self) -> Optional[pulumi.Input[str]]:
        """
        When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        """
        return pulumi.get(self, "proxy_name")

    @proxy_name.setter
    def proxy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_name", value)

    @property
    @pulumi.getter(name="proxyType")
    def proxy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        """
        return pulumi.get(self, "proxy_type")

    @proxy_type.setter
    def proxy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_type", value)

    @property
    @pulumi.getter(name="scheduleValues")
    def schedule_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Scheduling information.
        """
        return pulumi.get(self, "schedule_values")

    @schedule_values.setter
    def schedule_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "schedule_values", value)

    @property
    @pulumi.getter(name="securityType")
    def security_type(self) -> Optional[pulumi.Input[int]]:
        """
        - `0`: Disable security protection.- `1`: Enable security protection.
        """
        return pulumi.get(self, "security_type")

    @security_type.setter
    def security_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "security_type", value)

    @property
    @pulumi.getter(name="sessionPersistTime")
    def session_persist_time(self) -> Optional[pulumi.Input[int]]:
        """
        Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        """
        return pulumi.get(self, "session_persist_time")

    @session_persist_time.setter
    def session_persist_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_persist_time", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last modification date.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class ApplicationProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerate_type: Optional[pulumi.Input[int]] = None,
                 ipv6: Optional[pulumi.Input[pulumi.InputType['ApplicationProxyIpv6Args']]] = None,
                 plat_type: Optional[pulumi.Input[str]] = None,
                 proxy_name: Optional[pulumi.Input[str]] = None,
                 proxy_type: Optional[pulumi.Input[str]] = None,
                 security_type: Optional[pulumi.Input[int]] = None,
                 session_persist_time: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a teo application_proxy

        ## Import

        teo application_proxy can be imported using the zoneId#proxyId, e.g.

        ```sh
         $ pulumi import tencentcloud:Teo/applicationProxy:ApplicationProxy application_proxy zone-2983wizgxqvm#proxy-6972528a-373a-11ed-afca-52540044a456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] accelerate_type: - `0`: Disable acceleration.- `1`: Enable acceleration.
        :param pulumi.Input[pulumi.InputType['ApplicationProxyIpv6Args']] ipv6: IPv6 access configuration.
        :param pulumi.Input[str] plat_type: Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        :param pulumi.Input[str] proxy_name: When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        :param pulumi.Input[str] proxy_type: Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        :param pulumi.Input[int] security_type: - `0`: Disable security protection.- `1`: Enable security protection.
        :param pulumi.Input[int] session_persist_time: Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        :param pulumi.Input[str] status: Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a teo application_proxy

        ## Import

        teo application_proxy can be imported using the zoneId#proxyId, e.g.

        ```sh
         $ pulumi import tencentcloud:Teo/applicationProxy:ApplicationProxy application_proxy zone-2983wizgxqvm#proxy-6972528a-373a-11ed-afca-52540044a456
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerate_type: Optional[pulumi.Input[int]] = None,
                 ipv6: Optional[pulumi.Input[pulumi.InputType['ApplicationProxyIpv6Args']]] = None,
                 plat_type: Optional[pulumi.Input[str]] = None,
                 proxy_name: Optional[pulumi.Input[str]] = None,
                 proxy_type: Optional[pulumi.Input[str]] = None,
                 security_type: Optional[pulumi.Input[int]] = None,
                 session_persist_time: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApplicationProxyArgs.__new__(ApplicationProxyArgs)

            if accelerate_type is None and not opts.urn:
                raise TypeError("Missing required property 'accelerate_type'")
            __props__.__dict__["accelerate_type"] = accelerate_type
            __props__.__dict__["ipv6"] = ipv6
            if plat_type is None and not opts.urn:
                raise TypeError("Missing required property 'plat_type'")
            __props__.__dict__["plat_type"] = plat_type
            if proxy_name is None and not opts.urn:
                raise TypeError("Missing required property 'proxy_name'")
            __props__.__dict__["proxy_name"] = proxy_name
            __props__.__dict__["proxy_type"] = proxy_type
            if security_type is None and not opts.urn:
                raise TypeError("Missing required property 'security_type'")
            __props__.__dict__["security_type"] = security_type
            __props__.__dict__["session_persist_time"] = session_persist_time
            __props__.__dict__["status"] = status
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["area"] = None
            __props__.__dict__["ban_status"] = None
            __props__.__dict__["host_id"] = None
            __props__.__dict__["proxy_id"] = None
            __props__.__dict__["schedule_values"] = None
            __props__.__dict__["update_time"] = None
        super(ApplicationProxy, __self__).__init__(
            'tencentcloud:Teo/applicationProxy:ApplicationProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accelerate_type: Optional[pulumi.Input[int]] = None,
            area: Optional[pulumi.Input[str]] = None,
            ban_status: Optional[pulumi.Input[str]] = None,
            host_id: Optional[pulumi.Input[str]] = None,
            ipv6: Optional[pulumi.Input[pulumi.InputType['ApplicationProxyIpv6Args']]] = None,
            plat_type: Optional[pulumi.Input[str]] = None,
            proxy_id: Optional[pulumi.Input[str]] = None,
            proxy_name: Optional[pulumi.Input[str]] = None,
            proxy_type: Optional[pulumi.Input[str]] = None,
            schedule_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            security_type: Optional[pulumi.Input[int]] = None,
            session_persist_time: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'ApplicationProxy':
        """
        Get an existing ApplicationProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] accelerate_type: - `0`: Disable acceleration.- `1`: Enable acceleration.
        :param pulumi.Input[str] area: Acceleration area. Valid values: `mainland`, `overseas`.
        :param pulumi.Input[str] ban_status: Application proxy block status. Valid values: `banned`, `banning`, `recover`, `recovering`.
        :param pulumi.Input[str] host_id: When `ProxyType` is hostname, this field is the ID of the subdomain.
        :param pulumi.Input[pulumi.InputType['ApplicationProxyIpv6Args']] ipv6: IPv6 access configuration.
        :param pulumi.Input[str] plat_type: Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        :param pulumi.Input[str] proxy_id: Proxy ID.
        :param pulumi.Input[str] proxy_name: When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        :param pulumi.Input[str] proxy_type: Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] schedule_values: Scheduling information.
        :param pulumi.Input[int] security_type: - `0`: Disable security protection.- `1`: Enable security protection.
        :param pulumi.Input[int] session_persist_time: Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        :param pulumi.Input[str] status: Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        :param pulumi.Input[str] update_time: Last modification date.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationProxyState.__new__(_ApplicationProxyState)

        __props__.__dict__["accelerate_type"] = accelerate_type
        __props__.__dict__["area"] = area
        __props__.__dict__["ban_status"] = ban_status
        __props__.__dict__["host_id"] = host_id
        __props__.__dict__["ipv6"] = ipv6
        __props__.__dict__["plat_type"] = plat_type
        __props__.__dict__["proxy_id"] = proxy_id
        __props__.__dict__["proxy_name"] = proxy_name
        __props__.__dict__["proxy_type"] = proxy_type
        __props__.__dict__["schedule_values"] = schedule_values
        __props__.__dict__["security_type"] = security_type
        __props__.__dict__["session_persist_time"] = session_persist_time
        __props__.__dict__["status"] = status
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["zone_id"] = zone_id
        return ApplicationProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accelerateType")
    def accelerate_type(self) -> pulumi.Output[int]:
        """
        - `0`: Disable acceleration.- `1`: Enable acceleration.
        """
        return pulumi.get(self, "accelerate_type")

    @property
    @pulumi.getter
    def area(self) -> pulumi.Output[str]:
        """
        Acceleration area. Valid values: `mainland`, `overseas`.
        """
        return pulumi.get(self, "area")

    @property
    @pulumi.getter(name="banStatus")
    def ban_status(self) -> pulumi.Output[str]:
        """
        Application proxy block status. Valid values: `banned`, `banning`, `recover`, `recovering`.
        """
        return pulumi.get(self, "ban_status")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> pulumi.Output[str]:
        """
        When `ProxyType` is hostname, this field is the ID of the subdomain.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter
    def ipv6(self) -> pulumi.Output['outputs.ApplicationProxyIpv6']:
        """
        IPv6 access configuration.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="platType")
    def plat_type(self) -> pulumi.Output[str]:
        """
        Scheduling mode.- `ip`: Anycast IP.- `domain`: CNAME.
        """
        return pulumi.get(self, "plat_type")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> pulumi.Output[str]:
        """
        Proxy ID.
        """
        return pulumi.get(self, "proxy_id")

    @property
    @pulumi.getter(name="proxyName")
    def proxy_name(self) -> pulumi.Output[str]:
        """
        When `ProxyType` is hostname, `ProxyName` is the domain or subdomain name.When `ProxyType` is instance, `ProxyName` is the name of proxy application.
        """
        return pulumi.get(self, "proxy_name")

    @property
    @pulumi.getter(name="proxyType")
    def proxy_type(self) -> pulumi.Output[str]:
        """
        Layer 4 proxy mode. Valid values:- `hostname`: subdomain mode.- `instance`: instance mode.
        """
        return pulumi.get(self, "proxy_type")

    @property
    @pulumi.getter(name="scheduleValues")
    def schedule_values(self) -> pulumi.Output[Sequence[str]]:
        """
        Scheduling information.
        """
        return pulumi.get(self, "schedule_values")

    @property
    @pulumi.getter(name="securityType")
    def security_type(self) -> pulumi.Output[int]:
        """
        - `0`: Disable security protection.- `1`: Enable security protection.
        """
        return pulumi.get(self, "security_type")

    @property
    @pulumi.getter(name="sessionPersistTime")
    def session_persist_time(self) -> pulumi.Output[int]:
        """
        Session persistence duration. Value range: 30-3600 (in seconds), default value is 600.
        """
        return pulumi.get(self, "session_persist_time")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of this application proxy. Valid values to set is `online` and `offline`.- `online`: Enable.- `offline`: Disable.- `progress`: Deploying.- `stopping`: Deactivating.- `fail`: Deploy or deactivate failed.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last modification date.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

