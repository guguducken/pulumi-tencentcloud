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

__all__ = [
    'GetRabbitmqVipInstanceResult',
    'AwaitableGetRabbitmqVipInstanceResult',
    'get_rabbitmq_vip_instance',
    'get_rabbitmq_vip_instance_output',
]

@pulumi.output_type
class GetRabbitmqVipInstanceResult:
    """
    A collection of values returned by getRabbitmqVipInstance.
    """
    def __init__(__self__, filters=None, id=None, instances=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetRabbitmqVipInstanceFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetRabbitmqVipInstanceInstanceResult']:
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetRabbitmqVipInstanceResult(GetRabbitmqVipInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRabbitmqVipInstanceResult(
            filters=self.filters,
            id=self.id,
            instances=self.instances,
            result_output_file=self.result_output_file)


def get_rabbitmq_vip_instance(filters: Optional[Sequence[pulumi.InputType['GetRabbitmqVipInstanceFilterArgs']]] = None,
                              result_output_file: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRabbitmqVipInstanceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getRabbitmqVipInstance:getRabbitmqVipInstance', __args__, opts=opts, typ=GetRabbitmqVipInstanceResult).value

    return AwaitableGetRabbitmqVipInstanceResult(
        filters=__ret__.filters,
        id=__ret__.id,
        instances=__ret__.instances,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_rabbitmq_vip_instance)
def get_rabbitmq_vip_instance_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRabbitmqVipInstanceFilterArgs']]]]] = None,
                                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRabbitmqVipInstanceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
