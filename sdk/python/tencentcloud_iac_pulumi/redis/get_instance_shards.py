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
    'GetInstanceShardsResult',
    'AwaitableGetInstanceShardsResult',
    'get_instance_shards',
    'get_instance_shards_output',
]

@pulumi.output_type
class GetInstanceShardsResult:
    """
    A collection of values returned by getInstanceShards.
    """
    def __init__(__self__, filter_slave=None, id=None, instance_id=None, instance_shards=None, result_output_file=None):
        if filter_slave and not isinstance(filter_slave, bool):
            raise TypeError("Expected argument 'filter_slave' to be a bool")
        pulumi.set(__self__, "filter_slave", filter_slave)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_shards and not isinstance(instance_shards, list):
            raise TypeError("Expected argument 'instance_shards' to be a list")
        pulumi.set(__self__, "instance_shards", instance_shards)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="filterSlave")
    def filter_slave(self) -> Optional[bool]:
        return pulumi.get(self, "filter_slave")

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
    @pulumi.getter(name="instanceShards")
    def instance_shards(self) -> Sequence['outputs.GetInstanceShardsInstanceShardResult']:
        """
        Instance shard list information.
        """
        return pulumi.get(self, "instance_shards")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceShardsResult(GetInstanceShardsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceShardsResult(
            filter_slave=self.filter_slave,
            id=self.id,
            instance_id=self.instance_id,
            instance_shards=self.instance_shards,
            result_output_file=self.result_output_file)


def get_instance_shards(filter_slave: Optional[bool] = None,
                        instance_id: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceShardsResult:
    """
    Use this data source to query detailed information of redis instance_shards

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_shards = tencentcloud.Redis.get_instance_shards(filter_slave=False,
        instance_id="crs-c1nl9rpv")
    ```


    :param bool filter_slave: Whether to filter out slave information.
    :param str instance_id: The ID of instance.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['filterSlave'] = filter_slave
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Redis/getInstanceShards:getInstanceShards', __args__, opts=opts, typ=GetInstanceShardsResult).value

    return AwaitableGetInstanceShardsResult(
        filter_slave=__ret__.filter_slave,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_shards=__ret__.instance_shards,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_instance_shards)
def get_instance_shards_output(filter_slave: Optional[pulumi.Input[Optional[bool]]] = None,
                               instance_id: Optional[pulumi.Input[str]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceShardsResult]:
    """
    Use this data source to query detailed information of redis instance_shards

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_shards = tencentcloud.Redis.get_instance_shards(filter_slave=False,
        instance_id="crs-c1nl9rpv")
    ```


    :param bool filter_slave: Whether to filter out slave information.
    :param str instance_id: The ID of instance.
    :param str result_output_file: Used to save results.
    """
    ...
