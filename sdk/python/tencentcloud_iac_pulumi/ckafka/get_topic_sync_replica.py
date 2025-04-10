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
    'GetTopicSyncReplicaResult',
    'AwaitableGetTopicSyncReplicaResult',
    'get_topic_sync_replica',
    'get_topic_sync_replica_output',
]

@pulumi.output_type
class GetTopicSyncReplicaResult:
    """
    A collection of values returned by getTopicSyncReplica.
    """
    def __init__(__self__, id=None, instance_id=None, out_of_sync_replica_only=None, result_output_file=None, topic_in_sync_replica_lists=None, topic_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if out_of_sync_replica_only and not isinstance(out_of_sync_replica_only, bool):
            raise TypeError("Expected argument 'out_of_sync_replica_only' to be a bool")
        pulumi.set(__self__, "out_of_sync_replica_only", out_of_sync_replica_only)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if topic_in_sync_replica_lists and not isinstance(topic_in_sync_replica_lists, list):
            raise TypeError("Expected argument 'topic_in_sync_replica_lists' to be a list")
        pulumi.set(__self__, "topic_in_sync_replica_lists", topic_in_sync_replica_lists)
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        pulumi.set(__self__, "topic_name", topic_name)

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
    @pulumi.getter(name="outOfSyncReplicaOnly")
    def out_of_sync_replica_only(self) -> Optional[bool]:
        return pulumi.get(self, "out_of_sync_replica_only")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="topicInSyncReplicaLists")
    def topic_in_sync_replica_lists(self) -> Sequence['outputs.GetTopicSyncReplicaTopicInSyncReplicaListResult']:
        """
        Topic details and copy collection.
        """
        return pulumi.get(self, "topic_in_sync_replica_lists")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> str:
        return pulumi.get(self, "topic_name")


class AwaitableGetTopicSyncReplicaResult(GetTopicSyncReplicaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicSyncReplicaResult(
            id=self.id,
            instance_id=self.instance_id,
            out_of_sync_replica_only=self.out_of_sync_replica_only,
            result_output_file=self.result_output_file,
            topic_in_sync_replica_lists=self.topic_in_sync_replica_lists,
            topic_name=self.topic_name)


def get_topic_sync_replica(instance_id: Optional[str] = None,
                           out_of_sync_replica_only: Optional[bool] = None,
                           result_output_file: Optional[str] = None,
                           topic_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicSyncReplicaResult:
    """
    Use this data source to query detailed information of ckafka topic_sync_replica

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    topic_sync_replica = tencentcloud.Ckafka.get_topic_sync_replica(instance_id="ckafka-xxxxxx",
        topic_name="xxxxxx")
    ```


    :param str instance_id: InstanceId.
    :param bool out_of_sync_replica_only: Filter only unsynced replicas.
    :param str result_output_file: Used to save results.
    :param str topic_name: TopicName.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['outOfSyncReplicaOnly'] = out_of_sync_replica_only
    __args__['resultOutputFile'] = result_output_file
    __args__['topicName'] = topic_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getTopicSyncReplica:getTopicSyncReplica', __args__, opts=opts, typ=GetTopicSyncReplicaResult).value

    return AwaitableGetTopicSyncReplicaResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        out_of_sync_replica_only=__ret__.out_of_sync_replica_only,
        result_output_file=__ret__.result_output_file,
        topic_in_sync_replica_lists=__ret__.topic_in_sync_replica_lists,
        topic_name=__ret__.topic_name)


@_utilities.lift_output_func(get_topic_sync_replica)
def get_topic_sync_replica_output(instance_id: Optional[pulumi.Input[str]] = None,
                                  out_of_sync_replica_only: Optional[pulumi.Input[Optional[bool]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  topic_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicSyncReplicaResult]:
    """
    Use this data source to query detailed information of ckafka topic_sync_replica

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    topic_sync_replica = tencentcloud.Ckafka.get_topic_sync_replica(instance_id="ckafka-xxxxxx",
        topic_name="xxxxxx")
    ```


    :param str instance_id: InstanceId.
    :param bool out_of_sync_replica_only: Filter only unsynced replicas.
    :param str result_output_file: Used to save results.
    :param str topic_name: TopicName.
    """
    ...
