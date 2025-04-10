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
    'GetApplicationConfigResult',
    'AwaitableGetApplicationConfigResult',
    'get_application_config',
    'get_application_config_output',
]

@pulumi.output_type
class GetApplicationConfigResult:
    """
    A collection of values returned by getApplicationConfig.
    """
    def __init__(__self__, application_id=None, config_id=None, config_id_lists=None, config_name=None, config_version=None, id=None, result_output_file=None, results=None):
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if config_id and not isinstance(config_id, str):
            raise TypeError("Expected argument 'config_id' to be a str")
        pulumi.set(__self__, "config_id", config_id)
        if config_id_lists and not isinstance(config_id_lists, list):
            raise TypeError("Expected argument 'config_id_lists' to be a list")
        pulumi.set(__self__, "config_id_lists", config_id_lists)
        if config_name and not isinstance(config_name, str):
            raise TypeError("Expected argument 'config_name' to be a str")
        pulumi.set(__self__, "config_name", config_name)
        if config_version and not isinstance(config_version, str):
            raise TypeError("Expected argument 'config_version' to be a str")
        pulumi.set(__self__, "config_version", config_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[str]:
        """
        application Id. Note: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[str]:
        """
        Configuration ID. Note: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="configIdLists")
    def config_id_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "config_id_lists")

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> Optional[str]:
        """
        Configuration name Note: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> Optional[str]:
        """
        Configuration version. Note: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "config_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetApplicationConfigResultResult']:
        """
        Paginated configuration item list. Note: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "results")


class AwaitableGetApplicationConfigResult(GetApplicationConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationConfigResult(
            application_id=self.application_id,
            config_id=self.config_id,
            config_id_lists=self.config_id_lists,
            config_name=self.config_name,
            config_version=self.config_version,
            id=self.id,
            result_output_file=self.result_output_file,
            results=self.results)


def get_application_config(application_id: Optional[str] = None,
                           config_id: Optional[str] = None,
                           config_id_lists: Optional[Sequence[str]] = None,
                           config_name: Optional[str] = None,
                           config_version: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationConfigResult:
    """
    Use this data source to query detailed information of tsf application_config


    :param str application_id: Application ID, query all when not provided.
    :param str config_id: Configuration ID, query all with higher priority when not provided.
    :param Sequence[str] config_id_lists: Configuration ID list, query all with lower priority when not provided.
    :param str config_name: Configuration name, precise query, query all when not provided.
    :param str config_version: Configuration version, precise query, query all when not provided.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['configId'] = config_id
    __args__['configIdLists'] = config_id_lists
    __args__['configName'] = config_name
    __args__['configVersion'] = config_version
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tsf/getApplicationConfig:getApplicationConfig', __args__, opts=opts, typ=GetApplicationConfigResult).value

    return AwaitableGetApplicationConfigResult(
        application_id=__ret__.application_id,
        config_id=__ret__.config_id,
        config_id_lists=__ret__.config_id_lists,
        config_name=__ret__.config_name,
        config_version=__ret__.config_version,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        results=__ret__.results)


@_utilities.lift_output_func(get_application_config)
def get_application_config_output(application_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  config_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  config_id_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  config_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  config_version: Optional[pulumi.Input[Optional[str]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationConfigResult]:
    """
    Use this data source to query detailed information of tsf application_config


    :param str application_id: Application ID, query all when not provided.
    :param str config_id: Configuration ID, query all with higher priority when not provided.
    :param Sequence[str] config_id_lists: Configuration ID list, query all with lower priority when not provided.
    :param str config_name: Configuration name, precise query, query all when not provided.
    :param str config_version: Configuration version, precise query, query all when not provided.
    :param str result_output_file: Used to save results.
    """
    ...
