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
    'GetLogFilesResult',
    'AwaitableGetLogFilesResult',
    'get_log_files',
    'get_log_files_output',
]

@pulumi.output_type
class GetLogFilesResult:
    """
    A collection of values returned by getLogFiles.
    """
    def __init__(__self__, files=None, id=None, instance_id=None, normal_prefix=None, result_output_file=None, type=None, vpc_prefix=None):
        if files and not isinstance(files, list):
            raise TypeError("Expected argument 'files' to be a list")
        pulumi.set(__self__, "files", files)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if normal_prefix and not isinstance(normal_prefix, str):
            raise TypeError("Expected argument 'normal_prefix' to be a str")
        pulumi.set(__self__, "normal_prefix", normal_prefix)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if type and not isinstance(type, int):
            raise TypeError("Expected argument 'type' to be a int")
        pulumi.set(__self__, "type", type)
        if vpc_prefix and not isinstance(vpc_prefix, str):
            raise TypeError("Expected argument 'vpc_prefix' to be a str")
        pulumi.set(__self__, "vpc_prefix", vpc_prefix)

    @property
    @pulumi.getter
    def files(self) -> Sequence['outputs.GetLogFilesFileResult']:
        """
        Information such as `uri`, `length`, and `mtime` (modification time).
        """
        return pulumi.get(self, "files")

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
    @pulumi.getter(name="normalPrefix")
    def normal_prefix(self) -> str:
        """
        For an instance in a common network, this prefix plus URI can be used as the download address.
        """
        return pulumi.get(self, "normal_prefix")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def type(self) -> int:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcPrefix")
    def vpc_prefix(self) -> str:
        """
        For an instance in a VPC, this prefix plus URI can be used as the download address.
        """
        return pulumi.get(self, "vpc_prefix")


class AwaitableGetLogFilesResult(GetLogFilesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLogFilesResult(
            files=self.files,
            id=self.id,
            instance_id=self.instance_id,
            normal_prefix=self.normal_prefix,
            result_output_file=self.result_output_file,
            type=self.type,
            vpc_prefix=self.vpc_prefix)


def get_log_files(instance_id: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  type: Optional[int] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLogFilesResult:
    """
    Use this data source to query detailed information of mariadb log_files

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    log_files = tencentcloud.Mariadb.get_log_files(instance_id="tdsql-9vqvls95",
        type=1)
    ```


    :param str instance_id: Instance ID in the format of `tdsql-ow728lmc`.
    :param str result_output_file: Used to save results.
    :param int type: Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mariadb/getLogFiles:getLogFiles', __args__, opts=opts, typ=GetLogFilesResult).value

    return AwaitableGetLogFilesResult(
        files=__ret__.files,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        normal_prefix=__ret__.normal_prefix,
        result_output_file=__ret__.result_output_file,
        type=__ret__.type,
        vpc_prefix=__ret__.vpc_prefix)


@_utilities.lift_output_func(get_log_files)
def get_log_files_output(instance_id: Optional[pulumi.Input[str]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         type: Optional[pulumi.Input[int]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLogFilesResult]:
    """
    Use this data source to query detailed information of mariadb log_files

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    log_files = tencentcloud.Mariadb.get_log_files(instance_id="tdsql-9vqvls95",
        type=1)
    ```


    :param str instance_id: Instance ID in the format of `tdsql-ow728lmc`.
    :param str result_output_file: Used to save results.
    :param int type: Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
    """
    ...
