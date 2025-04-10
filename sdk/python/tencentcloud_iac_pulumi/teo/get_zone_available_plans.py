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
    'GetZoneAvailablePlansResult',
    'AwaitableGetZoneAvailablePlansResult',
    'get_zone_available_plans',
    'get_zone_available_plans_output',
]

@pulumi.output_type
class GetZoneAvailablePlansResult:
    """
    A collection of values returned by getZoneAvailablePlans.
    """
    def __init__(__self__, id=None, plan_info_lists=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if plan_info_lists and not isinstance(plan_info_lists, list):
            raise TypeError("Expected argument 'plan_info_lists' to be a list")
        pulumi.set(__self__, "plan_info_lists", plan_info_lists)
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
    @pulumi.getter(name="planInfoLists")
    def plan_info_lists(self) -> Sequence['outputs.GetZoneAvailablePlansPlanInfoListResult']:
        """
        Zone plans which current account can use.
        """
        return pulumi.get(self, "plan_info_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetZoneAvailablePlansResult(GetZoneAvailablePlansResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneAvailablePlansResult(
            id=self.id,
            plan_info_lists=self.plan_info_lists,
            result_output_file=self.result_output_file)


def get_zone_available_plans(result_output_file: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneAvailablePlansResult:
    """
    Use this data source to query detailed information of teo zoneAvailablePlans

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    zone_available_plans = tencentcloud.Teo.get_zone_available_plans()
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
    __ret__ = pulumi.runtime.invoke('tencentcloud:Teo/getZoneAvailablePlans:getZoneAvailablePlans', __args__, opts=opts, typ=GetZoneAvailablePlansResult).value

    return AwaitableGetZoneAvailablePlansResult(
        id=__ret__.id,
        plan_info_lists=__ret__.plan_info_lists,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_zone_available_plans)
def get_zone_available_plans_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneAvailablePlansResult]:
    """
    Use this data source to query detailed information of teo zoneAvailablePlans

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    zone_available_plans = tencentcloud.Teo.get_zone_available_plans()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
