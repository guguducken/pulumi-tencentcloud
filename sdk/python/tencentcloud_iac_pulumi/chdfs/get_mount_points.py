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
    'GetMountPointsResult',
    'AwaitableGetMountPointsResult',
    'get_mount_points',
    'get_mount_points_output',
]

@pulumi.output_type
class GetMountPointsResult:
    """
    A collection of values returned by getMountPoints.
    """
    def __init__(__self__, access_group_id=None, file_system_id=None, id=None, mount_points=None, owner_uin=None, result_output_file=None):
        if access_group_id and not isinstance(access_group_id, str):
            raise TypeError("Expected argument 'access_group_id' to be a str")
        pulumi.set(__self__, "access_group_id", access_group_id)
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mount_points and not isinstance(mount_points, list):
            raise TypeError("Expected argument 'mount_points' to be a list")
        pulumi.set(__self__, "mount_points", mount_points)
        if owner_uin and not isinstance(owner_uin, int):
            raise TypeError("Expected argument 'owner_uin' to be a int")
        pulumi.set(__self__, "owner_uin", owner_uin)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> Optional[str]:
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[str]:
        """
        mounted file system id.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mountPoints")
    def mount_points(self) -> Sequence['outputs.GetMountPointsMountPointResult']:
        """
        mount point list.
        """
        return pulumi.get(self, "mount_points")

    @property
    @pulumi.getter(name="ownerUin")
    def owner_uin(self) -> Optional[int]:
        return pulumi.get(self, "owner_uin")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetMountPointsResult(GetMountPointsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMountPointsResult(
            access_group_id=self.access_group_id,
            file_system_id=self.file_system_id,
            id=self.id,
            mount_points=self.mount_points,
            owner_uin=self.owner_uin,
            result_output_file=self.result_output_file)


def get_mount_points(access_group_id: Optional[str] = None,
                     file_system_id: Optional[str] = None,
                     owner_uin: Optional[int] = None,
                     result_output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMountPointsResult:
    """
    Use this data source to query detailed information of chdfs mount_points

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    mount_points = tencentcloud.Chdfs.get_mount_points(file_system_id="f14mpfy5lh4e")
    ```


    :param str access_group_id: get mount points belongs to access group id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
    :param str file_system_id: get mount points belongs to file system id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
    :param int owner_uin: get mount points belongs to owner uin, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['accessGroupId'] = access_group_id
    __args__['fileSystemId'] = file_system_id
    __args__['ownerUin'] = owner_uin
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Chdfs/getMountPoints:getMountPoints', __args__, opts=opts, typ=GetMountPointsResult).value

    return AwaitableGetMountPointsResult(
        access_group_id=__ret__.access_group_id,
        file_system_id=__ret__.file_system_id,
        id=__ret__.id,
        mount_points=__ret__.mount_points,
        owner_uin=__ret__.owner_uin,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_mount_points)
def get_mount_points_output(access_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                            file_system_id: Optional[pulumi.Input[Optional[str]]] = None,
                            owner_uin: Optional[pulumi.Input[Optional[int]]] = None,
                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMountPointsResult]:
    """
    Use this data source to query detailed information of chdfs mount_points

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    mount_points = tencentcloud.Chdfs.get_mount_points(file_system_id="f14mpfy5lh4e")
    ```


    :param str access_group_id: get mount points belongs to access group id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
    :param str file_system_id: get mount points belongs to file system id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
    :param int owner_uin: get mount points belongs to owner uin, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
    :param str result_output_file: Used to save results.
    """
    ...
