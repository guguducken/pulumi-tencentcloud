# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetAvailableZoneResult',
    'AwaitableGetAvailableZoneResult',
    'get_available_zone',
    'get_available_zone_output',
]

@pulumi.output_type
class GetAvailableZoneResult:
    """
    A collection of values returned by getAvailableZone.
    """
    def __init__(__self__, id=None, region_zones=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region_zones and not isinstance(region_zones, list):
            raise TypeError("Expected argument 'region_zones' to be a list")
        pulumi.set(__self__, "region_zones", region_zones)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="regionZones")
    def region_zones(self) -> Sequence['outputs.GetAvailableZoneRegionZoneResult']:
        """
        Information such as resource availability in each AZ and the supported storage classes and protocols.
        """
        return pulumi.get(self, "region_zones")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetAvailableZoneResult(GetAvailableZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAvailableZoneResult(
            id=self.id,
            region_zones=self.region_zones,
            result_output_file=self.result_output_file)


def get_available_zone(result_output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAvailableZoneResult:
    """
    Use this data source to query detailed information of cfs available_zone

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    available_zone = tencentcloud.Cfs.get_available_zone()
    ```


    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cfs/getAvailableZone:getAvailableZone', __args__, opts=opts, typ=GetAvailableZoneResult).value

    return AwaitableGetAvailableZoneResult(
        id=__ret__.id,
        region_zones=__ret__.region_zones,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_available_zone)
def get_available_zone_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAvailableZoneResult]:
    """
    Use this data source to query detailed information of cfs available_zone

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    available_zone = tencentcloud.Cfs.get_available_zone()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
