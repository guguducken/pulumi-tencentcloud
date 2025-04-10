# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApiDocArgs', 'ApiDoc']

@pulumi.input_type
class ApiDocArgs:
    def __init__(__self__, *,
                 api_doc_name: pulumi.Input[str],
                 api_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 environment: pulumi.Input[str],
                 service_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApiDoc resource.
        :param pulumi.Input[str] api_doc_name: Api Document name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_ids: List of APIs for generating documents.
        :param pulumi.Input[str] environment: Env name.
        :param pulumi.Input[str] service_id: Service name.
        """
        pulumi.set(__self__, "api_doc_name", api_doc_name)
        pulumi.set(__self__, "api_ids", api_ids)
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="apiDocName")
    def api_doc_name(self) -> pulumi.Input[str]:
        """
        Api Document name.
        """
        return pulumi.get(self, "api_doc_name")

    @api_doc_name.setter
    def api_doc_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_doc_name", value)

    @property
    @pulumi.getter(name="apiIds")
    def api_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of APIs for generating documents.
        """
        return pulumi.get(self, "api_ids")

    @api_ids.setter
    def api_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "api_ids", value)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        Env name.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        Service name.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)


@pulumi.input_type
class _ApiDocState:
    def __init__(__self__, *,
                 api_count: Optional[pulumi.Input[int]] = None,
                 api_doc_id: Optional[pulumi.Input[str]] = None,
                 api_doc_name: Optional[pulumi.Input[str]] = None,
                 api_doc_status: Optional[pulumi.Input[str]] = None,
                 api_doc_uri: Optional[pulumi.Input[str]] = None,
                 api_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 release_count: Optional[pulumi.Input[int]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 share_password: Optional[pulumi.Input[str]] = None,
                 updated_time: Optional[pulumi.Input[str]] = None,
                 view_count: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ApiDoc resources.
        :param pulumi.Input[int] api_count: Api Document count.
        :param pulumi.Input[str] api_doc_id: Api Document ID.
        :param pulumi.Input[str] api_doc_name: Api Document name.
        :param pulumi.Input[str] api_doc_status: API Document Build Status.
        :param pulumi.Input[str] api_doc_uri: API Document Access URI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_ids: List of APIs for generating documents.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_names: List of names for generating documents.
        :param pulumi.Input[str] environment: Env name.
        :param pulumi.Input[int] release_count: Number of API document releases.
        :param pulumi.Input[str] service_id: Service name.
        :param pulumi.Input[str] service_name: API Document service name.
        :param pulumi.Input[str] share_password: API Document Sharing Password.
        :param pulumi.Input[str] updated_time: API Document update time.
        :param pulumi.Input[int] view_count: API Document Viewing Times.
        """
        if api_count is not None:
            pulumi.set(__self__, "api_count", api_count)
        if api_doc_id is not None:
            pulumi.set(__self__, "api_doc_id", api_doc_id)
        if api_doc_name is not None:
            pulumi.set(__self__, "api_doc_name", api_doc_name)
        if api_doc_status is not None:
            pulumi.set(__self__, "api_doc_status", api_doc_status)
        if api_doc_uri is not None:
            pulumi.set(__self__, "api_doc_uri", api_doc_uri)
        if api_ids is not None:
            pulumi.set(__self__, "api_ids", api_ids)
        if api_names is not None:
            pulumi.set(__self__, "api_names", api_names)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if release_count is not None:
            pulumi.set(__self__, "release_count", release_count)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if share_password is not None:
            pulumi.set(__self__, "share_password", share_password)
        if updated_time is not None:
            pulumi.set(__self__, "updated_time", updated_time)
        if view_count is not None:
            pulumi.set(__self__, "view_count", view_count)

    @property
    @pulumi.getter(name="apiCount")
    def api_count(self) -> Optional[pulumi.Input[int]]:
        """
        Api Document count.
        """
        return pulumi.get(self, "api_count")

    @api_count.setter
    def api_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "api_count", value)

    @property
    @pulumi.getter(name="apiDocId")
    def api_doc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Api Document ID.
        """
        return pulumi.get(self, "api_doc_id")

    @api_doc_id.setter
    def api_doc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_doc_id", value)

    @property
    @pulumi.getter(name="apiDocName")
    def api_doc_name(self) -> Optional[pulumi.Input[str]]:
        """
        Api Document name.
        """
        return pulumi.get(self, "api_doc_name")

    @api_doc_name.setter
    def api_doc_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_doc_name", value)

    @property
    @pulumi.getter(name="apiDocStatus")
    def api_doc_status(self) -> Optional[pulumi.Input[str]]:
        """
        API Document Build Status.
        """
        return pulumi.get(self, "api_doc_status")

    @api_doc_status.setter
    def api_doc_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_doc_status", value)

    @property
    @pulumi.getter(name="apiDocUri")
    def api_doc_uri(self) -> Optional[pulumi.Input[str]]:
        """
        API Document Access URI.
        """
        return pulumi.get(self, "api_doc_uri")

    @api_doc_uri.setter
    def api_doc_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_doc_uri", value)

    @property
    @pulumi.getter(name="apiIds")
    def api_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of APIs for generating documents.
        """
        return pulumi.get(self, "api_ids")

    @api_ids.setter
    def api_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "api_ids", value)

    @property
    @pulumi.getter(name="apiNames")
    def api_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of names for generating documents.
        """
        return pulumi.get(self, "api_names")

    @api_names.setter
    def api_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "api_names", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        Env name.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="releaseCount")
    def release_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of API document releases.
        """
        return pulumi.get(self, "release_count")

    @release_count.setter
    def release_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "release_count", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        Service name.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        API Document service name.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="sharePassword")
    def share_password(self) -> Optional[pulumi.Input[str]]:
        """
        API Document Sharing Password.
        """
        return pulumi.get(self, "share_password")

    @share_password.setter
    def share_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_password", value)

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> Optional[pulumi.Input[str]]:
        """
        API Document update time.
        """
        return pulumi.get(self, "updated_time")

    @updated_time.setter
    def updated_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_time", value)

    @property
    @pulumi.getter(name="viewCount")
    def view_count(self) -> Optional[pulumi.Input[int]]:
        """
        API Document Viewing Times.
        """
        return pulumi.get(self, "view_count")

    @view_count.setter
    def view_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "view_count", value)


class ApiDoc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_doc_name: Optional[pulumi.Input[str]] = None,
                 api_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a APIGateway ApiDoc

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_api_doc = tencentcloud.api_gateway.ApiDoc("myApiDoc",
            api_doc_name="doc_test1",
            api_ids=[
                "api-test1",
                "api-test2",
            ],
            environment="release",
            service_id="service_test1")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_doc_name: Api Document name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_ids: List of APIs for generating documents.
        :param pulumi.Input[str] environment: Env name.
        :param pulumi.Input[str] service_id: Service name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiDocArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a APIGateway ApiDoc

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_api_doc = tencentcloud.api_gateway.ApiDoc("myApiDoc",
            api_doc_name="doc_test1",
            api_ids=[
                "api-test1",
                "api-test2",
            ],
            environment="release",
            service_id="service_test1")
        ```

        :param str resource_name: The name of the resource.
        :param ApiDocArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiDocArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_doc_name: Optional[pulumi.Input[str]] = None,
                 api_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiDocArgs.__new__(ApiDocArgs)

            if api_doc_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_doc_name'")
            __props__.__dict__["api_doc_name"] = api_doc_name
            if api_ids is None and not opts.urn:
                raise TypeError("Missing required property 'api_ids'")
            __props__.__dict__["api_ids"] = api_ids
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["api_count"] = None
            __props__.__dict__["api_doc_id"] = None
            __props__.__dict__["api_doc_status"] = None
            __props__.__dict__["api_doc_uri"] = None
            __props__.__dict__["api_names"] = None
            __props__.__dict__["release_count"] = None
            __props__.__dict__["service_name"] = None
            __props__.__dict__["share_password"] = None
            __props__.__dict__["updated_time"] = None
            __props__.__dict__["view_count"] = None
        super(ApiDoc, __self__).__init__(
            'tencentcloud:ApiGateway/apiDoc:ApiDoc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_count: Optional[pulumi.Input[int]] = None,
            api_doc_id: Optional[pulumi.Input[str]] = None,
            api_doc_name: Optional[pulumi.Input[str]] = None,
            api_doc_status: Optional[pulumi.Input[str]] = None,
            api_doc_uri: Optional[pulumi.Input[str]] = None,
            api_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            api_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            release_count: Optional[pulumi.Input[int]] = None,
            service_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            share_password: Optional[pulumi.Input[str]] = None,
            updated_time: Optional[pulumi.Input[str]] = None,
            view_count: Optional[pulumi.Input[int]] = None) -> 'ApiDoc':
        """
        Get an existing ApiDoc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] api_count: Api Document count.
        :param pulumi.Input[str] api_doc_id: Api Document ID.
        :param pulumi.Input[str] api_doc_name: Api Document name.
        :param pulumi.Input[str] api_doc_status: API Document Build Status.
        :param pulumi.Input[str] api_doc_uri: API Document Access URI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_ids: List of APIs for generating documents.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_names: List of names for generating documents.
        :param pulumi.Input[str] environment: Env name.
        :param pulumi.Input[int] release_count: Number of API document releases.
        :param pulumi.Input[str] service_id: Service name.
        :param pulumi.Input[str] service_name: API Document service name.
        :param pulumi.Input[str] share_password: API Document Sharing Password.
        :param pulumi.Input[str] updated_time: API Document update time.
        :param pulumi.Input[int] view_count: API Document Viewing Times.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiDocState.__new__(_ApiDocState)

        __props__.__dict__["api_count"] = api_count
        __props__.__dict__["api_doc_id"] = api_doc_id
        __props__.__dict__["api_doc_name"] = api_doc_name
        __props__.__dict__["api_doc_status"] = api_doc_status
        __props__.__dict__["api_doc_uri"] = api_doc_uri
        __props__.__dict__["api_ids"] = api_ids
        __props__.__dict__["api_names"] = api_names
        __props__.__dict__["environment"] = environment
        __props__.__dict__["release_count"] = release_count
        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["share_password"] = share_password
        __props__.__dict__["updated_time"] = updated_time
        __props__.__dict__["view_count"] = view_count
        return ApiDoc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiCount")
    def api_count(self) -> pulumi.Output[int]:
        """
        Api Document count.
        """
        return pulumi.get(self, "api_count")

    @property
    @pulumi.getter(name="apiDocId")
    def api_doc_id(self) -> pulumi.Output[str]:
        """
        Api Document ID.
        """
        return pulumi.get(self, "api_doc_id")

    @property
    @pulumi.getter(name="apiDocName")
    def api_doc_name(self) -> pulumi.Output[str]:
        """
        Api Document name.
        """
        return pulumi.get(self, "api_doc_name")

    @property
    @pulumi.getter(name="apiDocStatus")
    def api_doc_status(self) -> pulumi.Output[str]:
        """
        API Document Build Status.
        """
        return pulumi.get(self, "api_doc_status")

    @property
    @pulumi.getter(name="apiDocUri")
    def api_doc_uri(self) -> pulumi.Output[str]:
        """
        API Document Access URI.
        """
        return pulumi.get(self, "api_doc_uri")

    @property
    @pulumi.getter(name="apiIds")
    def api_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of APIs for generating documents.
        """
        return pulumi.get(self, "api_ids")

    @property
    @pulumi.getter(name="apiNames")
    def api_names(self) -> pulumi.Output[Sequence[str]]:
        """
        List of names for generating documents.
        """
        return pulumi.get(self, "api_names")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        Env name.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="releaseCount")
    def release_count(self) -> pulumi.Output[int]:
        """
        Number of API document releases.
        """
        return pulumi.get(self, "release_count")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        Service name.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        API Document service name.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="sharePassword")
    def share_password(self) -> pulumi.Output[str]:
        """
        API Document Sharing Password.
        """
        return pulumi.get(self, "share_password")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> pulumi.Output[str]:
        """
        API Document update time.
        """
        return pulumi.get(self, "updated_time")

    @property
    @pulumi.getter(name="viewCount")
    def view_count(self) -> pulumi.Output[int]:
        """
        API Document Viewing Times.
        """
        return pulumi.get(self, "view_count")

