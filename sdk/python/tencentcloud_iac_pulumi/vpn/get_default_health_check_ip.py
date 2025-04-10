# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDefaultHealthCheckIpResult',
    'AwaitableGetDefaultHealthCheckIpResult',
    'get_default_health_check_ip',
    'get_default_health_check_ip_output',
]

@pulumi.output_type
class GetDefaultHealthCheckIpResult:
    """
    A collection of values returned by getDefaultHealthCheckIp.
    """
    def __init__(__self__, health_check_local_ip=None, health_check_remote_ip=None, id=None, result_output_file=None, vpn_gateway_id=None):
        if health_check_local_ip and not isinstance(health_check_local_ip, str):
            raise TypeError("Expected argument 'health_check_local_ip' to be a str")
        pulumi.set(__self__, "health_check_local_ip", health_check_local_ip)
        if health_check_remote_ip and not isinstance(health_check_remote_ip, str):
            raise TypeError("Expected argument 'health_check_remote_ip' to be a str")
        pulumi.set(__self__, "health_check_remote_ip", health_check_remote_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if vpn_gateway_id and not isinstance(vpn_gateway_id, str):
            raise TypeError("Expected argument 'vpn_gateway_id' to be a str")
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="healthCheckLocalIp")
    def health_check_local_ip(self) -> str:
        """
        local ip of health check.
        """
        return pulumi.get(self, "health_check_local_ip")

    @property
    @pulumi.getter(name="healthCheckRemoteIp")
    def health_check_remote_ip(self) -> str:
        """
        remote ip for health check.
        """
        return pulumi.get(self, "health_check_remote_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> str:
        return pulumi.get(self, "vpn_gateway_id")


class AwaitableGetDefaultHealthCheckIpResult(GetDefaultHealthCheckIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultHealthCheckIpResult(
            health_check_local_ip=self.health_check_local_ip,
            health_check_remote_ip=self.health_check_remote_ip,
            id=self.id,
            result_output_file=self.result_output_file,
            vpn_gateway_id=self.vpn_gateway_id)


def get_default_health_check_ip(result_output_file: Optional[str] = None,
                                vpn_gateway_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultHealthCheckIpResult:
    """
    Use this data source to query detailed information of vpn default_health_check_ip

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    default_health_check_ip = tencentcloud.Vpn.get_default_health_check_ip(vpn_gateway_id="vpngw-gt8bianl")
    ```


    :param str result_output_file: Used to save results.
    :param str vpn_gateway_id: vpn gateway id.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    __args__['vpnGatewayId'] = vpn_gateway_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpn/getDefaultHealthCheckIp:getDefaultHealthCheckIp', __args__, opts=opts, typ=GetDefaultHealthCheckIpResult).value

    return AwaitableGetDefaultHealthCheckIpResult(
        health_check_local_ip=__ret__.health_check_local_ip,
        health_check_remote_ip=__ret__.health_check_remote_ip,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        vpn_gateway_id=__ret__.vpn_gateway_id)


@_utilities.lift_output_func(get_default_health_check_ip)
def get_default_health_check_ip_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultHealthCheckIpResult]:
    """
    Use this data source to query detailed information of vpn default_health_check_ip

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    default_health_check_ip = tencentcloud.Vpn.get_default_health_check_ip(vpn_gateway_id="vpngw-gt8bianl")
    ```


    :param str result_output_file: Used to save results.
    :param str vpn_gateway_id: vpn gateway id.
    """
    ...
