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
    'GetDefaultParametersResult',
    'AwaitableGetDefaultParametersResult',
    'get_default_parameters',
    'get_default_parameters_output',
]

@pulumi.output_type
class GetDefaultParametersResult:
    """
    A collection of values returned by getDefaultParameters.
    """
    def __init__(__self__, db_engine=None, db_major_version=None, id=None, param_info_sets=None, result_output_file=None):
        if db_engine and not isinstance(db_engine, str):
            raise TypeError("Expected argument 'db_engine' to be a str")
        pulumi.set(__self__, "db_engine", db_engine)
        if db_major_version and not isinstance(db_major_version, str):
            raise TypeError("Expected argument 'db_major_version' to be a str")
        pulumi.set(__self__, "db_major_version", db_major_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if param_info_sets and not isinstance(param_info_sets, list):
            raise TypeError("Expected argument 'param_info_sets' to be a list")
        pulumi.set(__self__, "param_info_sets", param_info_sets)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="dbEngine")
    def db_engine(self) -> str:
        return pulumi.get(self, "db_engine")

    @property
    @pulumi.getter(name="dbMajorVersion")
    def db_major_version(self) -> str:
        return pulumi.get(self, "db_major_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="paramInfoSets")
    def param_info_sets(self) -> Sequence['outputs.GetDefaultParametersParamInfoSetResult']:
        """
        Parameter informationNote: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "param_info_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDefaultParametersResult(GetDefaultParametersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultParametersResult(
            db_engine=self.db_engine,
            db_major_version=self.db_major_version,
            id=self.id,
            param_info_sets=self.param_info_sets,
            result_output_file=self.result_output_file)


def get_default_parameters(db_engine: Optional[str] = None,
                           db_major_version: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultParametersResult:
    """
    Use this data source to query detailed information of postgresql default_parameters

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    default_parameters = tencentcloud.Postgresql.get_default_parameters(db_engine="postgresql",
        db_major_version="13")
    ```


    :param str db_engine: Database engine, such as postgresql, mssql_compatible.
    :param str db_major_version: The major database version number, such as 11, 12, 13.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['dbEngine'] = db_engine
    __args__['dbMajorVersion'] = db_major_version
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Postgresql/getDefaultParameters:getDefaultParameters', __args__, opts=opts, typ=GetDefaultParametersResult).value

    return AwaitableGetDefaultParametersResult(
        db_engine=__ret__.db_engine,
        db_major_version=__ret__.db_major_version,
        id=__ret__.id,
        param_info_sets=__ret__.param_info_sets,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_default_parameters)
def get_default_parameters_output(db_engine: Optional[pulumi.Input[str]] = None,
                                  db_major_version: Optional[pulumi.Input[str]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultParametersResult]:
    """
    Use this data source to query detailed information of postgresql default_parameters

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    default_parameters = tencentcloud.Postgresql.get_default_parameters(db_engine="postgresql",
        db_major_version="13")
    ```


    :param str db_engine: Database engine, such as postgresql, mssql_compatible.
    :param str db_major_version: The major database version number, such as 11, 12, 13.
    :param str result_output_file: Used to save results.
    """
    ...
