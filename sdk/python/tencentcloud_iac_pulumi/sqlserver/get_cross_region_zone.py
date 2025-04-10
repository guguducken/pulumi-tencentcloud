# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetCrossRegionZoneResult',
    'AwaitableGetCrossRegionZoneResult',
    'get_cross_region_zone',
    'get_cross_region_zone_output',
]

@pulumi.output_type
class GetCrossRegionZoneResult:
    """
    A collection of values returned by getCrossRegionZone.
    """
    def __init__(__self__, id=None, instance_id=None, region=None, result_output_file=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The string ID of the region where the standby machine is located, such as: ap-guangzhou.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The string ID of the availability zone where the standby machine is located, such as: ap-guangzhou-1.
        """
        return pulumi.get(self, "zone")


class AwaitableGetCrossRegionZoneResult(GetCrossRegionZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCrossRegionZoneResult(
            id=self.id,
            instance_id=self.instance_id,
            region=self.region,
            result_output_file=self.result_output_file,
            zone=self.zone)


def get_cross_region_zone(instance_id: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCrossRegionZoneResult:
    """
    Use this data source to query detailed information of sqlserver datasource_cross_region_zone

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cross_region_zone = tencentcloud.Sqlserver.get_cross_region_zone(instance_id="mssql-qelbzgwf")
    ```


    :param str instance_id: Instance ID in the format of mssql-j8kv137v.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getCrossRegionZone:getCrossRegionZone', __args__, opts=opts, typ=GetCrossRegionZoneResult).value

    return AwaitableGetCrossRegionZoneResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        region=__ret__.region,
        result_output_file=__ret__.result_output_file,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_cross_region_zone)
def get_cross_region_zone_output(instance_id: Optional[pulumi.Input[str]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCrossRegionZoneResult]:
    """
    Use this data source to query detailed information of sqlserver datasource_cross_region_zone

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cross_region_zone = tencentcloud.Sqlserver.get_cross_region_zone(instance_id="mssql-qelbzgwf")
    ```


    :param str instance_id: Instance ID in the format of mssql-j8kv137v.
    :param str result_output_file: Used to save results.
    """
    ...
