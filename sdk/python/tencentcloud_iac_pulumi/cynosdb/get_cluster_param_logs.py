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
    'GetClusterParamLogsResult',
    'AwaitableGetClusterParamLogsResult',
    'get_cluster_param_logs',
    'get_cluster_param_logs_output',
]

@pulumi.output_type
class GetClusterParamLogsResult:
    """
    A collection of values returned by getClusterParamLogs.
    """
    def __init__(__self__, cluster_id=None, cluster_param_logs=None, id=None, instance_ids=None, order_by=None, order_by_type=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_param_logs and not isinstance(cluster_param_logs, list):
            raise TypeError("Expected argument 'cluster_param_logs' to be a list")
        pulumi.set(__self__, "cluster_param_logs", cluster_param_logs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if order_by_type and not isinstance(order_by_type, str):
            raise TypeError("Expected argument 'order_by_type' to be a str")
        pulumi.set(__self__, "order_by_type", order_by_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterParamLogs")
    def cluster_param_logs(self) -> Sequence['outputs.GetClusterParamLogsClusterParamLogResult']:
        """
        Parameter modification record note: This field may return null, indicating that a valid value cannot be obtained.
        """
        return pulumi.get(self, "cluster_param_logs")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter(name="orderByType")
    def order_by_type(self) -> Optional[str]:
        return pulumi.get(self, "order_by_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetClusterParamLogsResult(GetClusterParamLogsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterParamLogsResult(
            cluster_id=self.cluster_id,
            cluster_param_logs=self.cluster_param_logs,
            id=self.id,
            instance_ids=self.instance_ids,
            order_by=self.order_by,
            order_by_type=self.order_by_type,
            result_output_file=self.result_output_file)


def get_cluster_param_logs(cluster_id: Optional[str] = None,
                           instance_ids: Optional[Sequence[str]] = None,
                           order_by: Optional[str] = None,
                           order_by_type: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterParamLogsResult:
    """
    Use this data source to query detailed information of cynosdb cluster_param_logs

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cluster_param_logs = tencentcloud.Cynosdb.get_cluster_param_logs(cluster_id="cynosdbmysql-bws8h88b",
        instance_ids=["cynosdbmysql-ins-afqx1hy0"],
        order_by="CreateTime",
        order_by_type="DESC")
    ```


    :param str cluster_id: Cluster ID.
    :param Sequence[str] instance_ids: Instance ID list, used to record specific instances of operations.
    :param str order_by: Sort field, defining which field to sort based on when returning results.
    :param str order_by_type: Define specific sorting rules, limited to one of desc, asc, DESC, or ASC.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['instanceIds'] = instance_ids
    __args__['orderBy'] = order_by
    __args__['orderByType'] = order_by_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cynosdb/getClusterParamLogs:getClusterParamLogs', __args__, opts=opts, typ=GetClusterParamLogsResult).value

    return AwaitableGetClusterParamLogsResult(
        cluster_id=__ret__.cluster_id,
        cluster_param_logs=__ret__.cluster_param_logs,
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        order_by=__ret__.order_by,
        order_by_type=__ret__.order_by_type,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_cluster_param_logs)
def get_cluster_param_logs_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                  instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  order_by: Optional[pulumi.Input[Optional[str]]] = None,
                                  order_by_type: Optional[pulumi.Input[Optional[str]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterParamLogsResult]:
    """
    Use this data source to query detailed information of cynosdb cluster_param_logs

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cluster_param_logs = tencentcloud.Cynosdb.get_cluster_param_logs(cluster_id="cynosdbmysql-bws8h88b",
        instance_ids=["cynosdbmysql-ins-afqx1hy0"],
        order_by="CreateTime",
        order_by_type="DESC")
    ```


    :param str cluster_id: Cluster ID.
    :param Sequence[str] instance_ids: Instance ID list, used to record specific instances of operations.
    :param str order_by: Sort field, defining which field to sort based on when returning results.
    :param str order_by_type: Define specific sorting rules, limited to one of desc, asc, DESC, or ASC.
    :param str result_output_file: Used to save results.
    """
    ...
