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
    'GetProjectSecurityGroupsResult',
    'AwaitableGetProjectSecurityGroupsResult',
    'get_project_security_groups',
    'get_project_security_groups_output',
]

@pulumi.output_type
class GetProjectSecurityGroupsResult:
    """
    A collection of values returned by getProjectSecurityGroups.
    """
    def __init__(__self__, id=None, project_id=None, result_output_file=None, security_group_sets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if security_group_sets and not isinstance(security_group_sets, list):
            raise TypeError("Expected argument 'security_group_sets' to be a list")
        pulumi.set(__self__, "security_group_sets", security_group_sets)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        """
        project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="securityGroupSets")
    def security_group_sets(self) -> Sequence['outputs.GetProjectSecurityGroupsSecurityGroupSetResult']:
        """
        Security group details.
        """
        return pulumi.get(self, "security_group_sets")


class AwaitableGetProjectSecurityGroupsResult(GetProjectSecurityGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectSecurityGroupsResult(
            id=self.id,
            project_id=self.project_id,
            result_output_file=self.result_output_file,
            security_group_sets=self.security_group_sets)


def get_project_security_groups(project_id: Optional[int] = None,
                                result_output_file: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectSecurityGroupsResult:
    """
    Use this data source to query detailed information of sqlserver project_security_groups

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    project_security_groups = tencentcloud.Sqlserver.get_project_security_groups(project_id=0)
    ```


    :param int project_id: Project ID, which can be viewed through the console project management.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getProjectSecurityGroups:getProjectSecurityGroups', __args__, opts=opts, typ=GetProjectSecurityGroupsResult).value

    return AwaitableGetProjectSecurityGroupsResult(
        id=__ret__.id,
        project_id=__ret__.project_id,
        result_output_file=__ret__.result_output_file,
        security_group_sets=__ret__.security_group_sets)


@_utilities.lift_output_func(get_project_security_groups)
def get_project_security_groups_output(project_id: Optional[pulumi.Input[int]] = None,
                                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectSecurityGroupsResult]:
    """
    Use this data source to query detailed information of sqlserver project_security_groups

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    project_security_groups = tencentcloud.Sqlserver.get_project_security_groups(project_id=0)
    ```


    :param int project_id: Project ID, which can be viewed through the console project management.
    :param str result_output_file: Used to save results.
    """
    ...
