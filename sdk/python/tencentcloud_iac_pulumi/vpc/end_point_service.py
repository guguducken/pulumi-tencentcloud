# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EndPointServiceArgs', 'EndPointService']

@pulumi.input_type
class EndPointServiceArgs:
    def __init__(__self__, *,
                 auto_accept_flag: pulumi.Input[bool],
                 end_point_service_name: pulumi.Input[str],
                 service_instance_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 service_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndPointService resource.
        :param pulumi.Input[bool] auto_accept_flag: Whether to automatically accept.
        :param pulumi.Input[str] end_point_service_name: Name of end point service.
        :param pulumi.Input[str] service_instance_id: Id of service instance, like lb-xxx.
        :param pulumi.Input[str] vpc_id: ID of vpc instance.
        :param pulumi.Input[str] service_type: Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        """
        pulumi.set(__self__, "auto_accept_flag", auto_accept_flag)
        pulumi.set(__self__, "end_point_service_name", end_point_service_name)
        pulumi.set(__self__, "service_instance_id", service_instance_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)

    @property
    @pulumi.getter(name="autoAcceptFlag")
    def auto_accept_flag(self) -> pulumi.Input[bool]:
        """
        Whether to automatically accept.
        """
        return pulumi.get(self, "auto_accept_flag")

    @auto_accept_flag.setter
    def auto_accept_flag(self, value: pulumi.Input[bool]):
        pulumi.set(self, "auto_accept_flag", value)

    @property
    @pulumi.getter(name="endPointServiceName")
    def end_point_service_name(self) -> pulumi.Input[str]:
        """
        Name of end point service.
        """
        return pulumi.get(self, "end_point_service_name")

    @end_point_service_name.setter
    def end_point_service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "end_point_service_name", value)

    @property
    @pulumi.getter(name="serviceInstanceId")
    def service_instance_id(self) -> pulumi.Input[str]:
        """
        Id of service instance, like lb-xxx.
        """
        return pulumi.get(self, "service_instance_id")

    @service_instance_id.setter
    def service_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_instance_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        ID of vpc instance.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_type", value)


@pulumi.input_type
class _EndPointServiceState:
    def __init__(__self__, *,
                 auto_accept_flag: Optional[pulumi.Input[bool]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 end_point_count: Optional[pulumi.Input[int]] = None,
                 end_point_service_name: Optional[pulumi.Input[str]] = None,
                 service_instance_id: Optional[pulumi.Input[str]] = None,
                 service_owner: Optional[pulumi.Input[str]] = None,
                 service_type: Optional[pulumi.Input[str]] = None,
                 service_vip: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndPointService resources.
        :param pulumi.Input[bool] auto_accept_flag: Whether to automatically accept.
        :param pulumi.Input[str] create_time: Create Time.
        :param pulumi.Input[int] end_point_count: Count of end point.
        :param pulumi.Input[str] end_point_service_name: Name of end point service.
        :param pulumi.Input[str] service_instance_id: Id of service instance, like lb-xxx.
        :param pulumi.Input[str] service_owner: APPID.
        :param pulumi.Input[str] service_type: Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        :param pulumi.Input[str] service_vip: VIP of backend service.
        :param pulumi.Input[str] vpc_id: ID of vpc instance.
        """
        if auto_accept_flag is not None:
            pulumi.set(__self__, "auto_accept_flag", auto_accept_flag)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if end_point_count is not None:
            pulumi.set(__self__, "end_point_count", end_point_count)
        if end_point_service_name is not None:
            pulumi.set(__self__, "end_point_service_name", end_point_service_name)
        if service_instance_id is not None:
            pulumi.set(__self__, "service_instance_id", service_instance_id)
        if service_owner is not None:
            pulumi.set(__self__, "service_owner", service_owner)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)
        if service_vip is not None:
            pulumi.set(__self__, "service_vip", service_vip)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="autoAcceptFlag")
    def auto_accept_flag(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically accept.
        """
        return pulumi.get(self, "auto_accept_flag")

    @auto_accept_flag.setter
    def auto_accept_flag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_accept_flag", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create Time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="endPointCount")
    def end_point_count(self) -> Optional[pulumi.Input[int]]:
        """
        Count of end point.
        """
        return pulumi.get(self, "end_point_count")

    @end_point_count.setter
    def end_point_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_point_count", value)

    @property
    @pulumi.getter(name="endPointServiceName")
    def end_point_service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of end point service.
        """
        return pulumi.get(self, "end_point_service_name")

    @end_point_service_name.setter
    def end_point_service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_point_service_name", value)

    @property
    @pulumi.getter(name="serviceInstanceId")
    def service_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of service instance, like lb-xxx.
        """
        return pulumi.get(self, "service_instance_id")

    @service_instance_id.setter
    def service_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_instance_id", value)

    @property
    @pulumi.getter(name="serviceOwner")
    def service_owner(self) -> Optional[pulumi.Input[str]]:
        """
        APPID.
        """
        return pulumi.get(self, "service_owner")

    @service_owner.setter
    def service_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_owner", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_type", value)

    @property
    @pulumi.getter(name="serviceVip")
    def service_vip(self) -> Optional[pulumi.Input[str]]:
        """
        VIP of backend service.
        """
        return pulumi.get(self, "service_vip")

    @service_vip.setter
    def service_vip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_vip", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of vpc instance.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class EndPointService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_accept_flag: Optional[pulumi.Input[bool]] = None,
                 end_point_service_name: Optional[pulumi.Input[str]] = None,
                 service_instance_id: Optional[pulumi.Input[str]] = None,
                 service_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc end_point_service

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        end_point_service = tencentcloud.vpc.EndPointService("endPointService",
            auto_accept_flag=False,
            end_point_service_name="terraform-endpoint-service",
            service_instance_id="lb-o5f6x7ke",
            service_type="CLB",
            vpc_id="vpc-391sv4w3")
        ```

        ## Import

        vpc end_point_service can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/endPointService:EndPointService end_point_service end_point_service_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_accept_flag: Whether to automatically accept.
        :param pulumi.Input[str] end_point_service_name: Name of end point service.
        :param pulumi.Input[str] service_instance_id: Id of service instance, like lb-xxx.
        :param pulumi.Input[str] service_type: Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        :param pulumi.Input[str] vpc_id: ID of vpc instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndPointServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc end_point_service

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        end_point_service = tencentcloud.vpc.EndPointService("endPointService",
            auto_accept_flag=False,
            end_point_service_name="terraform-endpoint-service",
            service_instance_id="lb-o5f6x7ke",
            service_type="CLB",
            vpc_id="vpc-391sv4w3")
        ```

        ## Import

        vpc end_point_service can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/endPointService:EndPointService end_point_service end_point_service_id
        ```

        :param str resource_name: The name of the resource.
        :param EndPointServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndPointServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_accept_flag: Optional[pulumi.Input[bool]] = None,
                 end_point_service_name: Optional[pulumi.Input[str]] = None,
                 service_instance_id: Optional[pulumi.Input[str]] = None,
                 service_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = EndPointServiceArgs.__new__(EndPointServiceArgs)

            if auto_accept_flag is None and not opts.urn:
                raise TypeError("Missing required property 'auto_accept_flag'")
            __props__.__dict__["auto_accept_flag"] = auto_accept_flag
            if end_point_service_name is None and not opts.urn:
                raise TypeError("Missing required property 'end_point_service_name'")
            __props__.__dict__["end_point_service_name"] = end_point_service_name
            if service_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_instance_id'")
            __props__.__dict__["service_instance_id"] = service_instance_id
            __props__.__dict__["service_type"] = service_type
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["end_point_count"] = None
            __props__.__dict__["service_owner"] = None
            __props__.__dict__["service_vip"] = None
        super(EndPointService, __self__).__init__(
            'tencentcloud:Vpc/endPointService:EndPointService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_accept_flag: Optional[pulumi.Input[bool]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            end_point_count: Optional[pulumi.Input[int]] = None,
            end_point_service_name: Optional[pulumi.Input[str]] = None,
            service_instance_id: Optional[pulumi.Input[str]] = None,
            service_owner: Optional[pulumi.Input[str]] = None,
            service_type: Optional[pulumi.Input[str]] = None,
            service_vip: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'EndPointService':
        """
        Get an existing EndPointService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_accept_flag: Whether to automatically accept.
        :param pulumi.Input[str] create_time: Create Time.
        :param pulumi.Input[int] end_point_count: Count of end point.
        :param pulumi.Input[str] end_point_service_name: Name of end point service.
        :param pulumi.Input[str] service_instance_id: Id of service instance, like lb-xxx.
        :param pulumi.Input[str] service_owner: APPID.
        :param pulumi.Input[str] service_type: Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        :param pulumi.Input[str] service_vip: VIP of backend service.
        :param pulumi.Input[str] vpc_id: ID of vpc instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndPointServiceState.__new__(_EndPointServiceState)

        __props__.__dict__["auto_accept_flag"] = auto_accept_flag
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["end_point_count"] = end_point_count
        __props__.__dict__["end_point_service_name"] = end_point_service_name
        __props__.__dict__["service_instance_id"] = service_instance_id
        __props__.__dict__["service_owner"] = service_owner
        __props__.__dict__["service_type"] = service_type
        __props__.__dict__["service_vip"] = service_vip
        __props__.__dict__["vpc_id"] = vpc_id
        return EndPointService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoAcceptFlag")
    def auto_accept_flag(self) -> pulumi.Output[bool]:
        """
        Whether to automatically accept.
        """
        return pulumi.get(self, "auto_accept_flag")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create Time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="endPointCount")
    def end_point_count(self) -> pulumi.Output[int]:
        """
        Count of end point.
        """
        return pulumi.get(self, "end_point_count")

    @property
    @pulumi.getter(name="endPointServiceName")
    def end_point_service_name(self) -> pulumi.Output[str]:
        """
        Name of end point service.
        """
        return pulumi.get(self, "end_point_service_name")

    @property
    @pulumi.getter(name="serviceInstanceId")
    def service_instance_id(self) -> pulumi.Output[str]:
        """
        Id of service instance, like lb-xxx.
        """
        return pulumi.get(self, "service_instance_id")

    @property
    @pulumi.getter(name="serviceOwner")
    def service_owner(self) -> pulumi.Output[str]:
        """
        APPID.
        """
        return pulumi.get(self, "service_owner")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> pulumi.Output[str]:
        """
        Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter(name="serviceVip")
    def service_vip(self) -> pulumi.Output[str]:
        """
        VIP of backend service.
        """
        return pulumi.get(self, "service_vip")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        ID of vpc instance.
        """
        return pulumi.get(self, "vpc_id")

