# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PolicyByNameArgs', 'PolicyByName']

@pulumi.input_type
class PolicyByNameArgs:
    def __init__(__self__, *,
                 document: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PolicyByName resource.
        :param pulumi.Input[str] document: Document of the CAM policy. The syntax refers to [CAM
               POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
               terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
               Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        :param pulumi.Input[str] description: Description of the CAM policy.
        :param pulumi.Input[str] name: Name of CAM policy.
        """
        pulumi.set(__self__, "document", document)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def document(self) -> pulumi.Input[str]:
        """
        Document of the CAM policy. The syntax refers to [CAM
        POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
        terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
        Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        """
        return pulumi.get(self, "document")

    @document.setter
    def document(self, value: pulumi.Input[str]):
        pulumi.set(self, "document", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the CAM policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of CAM policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PolicyByNameState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[int]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyByName resources.
        :param pulumi.Input[str] create_time: Create time of the CAM policy.
        :param pulumi.Input[str] description: Description of the CAM policy.
        :param pulumi.Input[str] document: Document of the CAM policy. The syntax refers to [CAM
               POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
               terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
               Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        :param pulumi.Input[str] name: Name of CAM policy.
        :param pulumi.Input[int] type: Type of the policy strategy. Valid values: `1`, `2`. `1` means customer strategy and `2` means preset strategy.
        :param pulumi.Input[str] update_time: The last update time of the CAM policy.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if document is not None:
            pulumi.set(__self__, "document", document)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the CAM policy.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the CAM policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def document(self) -> Optional[pulumi.Input[str]]:
        """
        Document of the CAM policy. The syntax refers to [CAM
        POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
        terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
        Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        """
        return pulumi.get(self, "document")

    @document.setter
    def document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of CAM policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[int]]:
        """
        Type of the policy strategy. Valid values: `1`, `2`. `1` means customer strategy and `2` means preset strategy.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The last update time of the CAM policy.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class PolicyByName(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PolicyByName resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the CAM policy.
        :param pulumi.Input[str] document: Document of the CAM policy. The syntax refers to [CAM
               POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
               terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
               Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        :param pulumi.Input[str] name: Name of CAM policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyByNameArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PolicyByName resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PolicyByNameArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyByNameArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = PolicyByNameArgs.__new__(PolicyByNameArgs)

            __props__.__dict__["description"] = description
            if document is None and not opts.urn:
                raise TypeError("Missing required property 'document'")
            __props__.__dict__["document"] = document
            __props__.__dict__["name"] = name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["update_time"] = None
        super(PolicyByName, __self__).__init__(
            'tencentcloud:Cam/policyByName:PolicyByName',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            document: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[int]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'PolicyByName':
        """
        Get an existing PolicyByName resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time of the CAM policy.
        :param pulumi.Input[str] description: Description of the CAM policy.
        :param pulumi.Input[str] document: Document of the CAM policy. The syntax refers to [CAM
               POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
               terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
               Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        :param pulumi.Input[str] name: Name of CAM policy.
        :param pulumi.Input[int] type: Type of the policy strategy. Valid values: `1`, `2`. `1` means customer strategy and `2` means preset strategy.
        :param pulumi.Input[str] update_time: The last update time of the CAM policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyByNameState.__new__(_PolicyByNameState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["document"] = document
        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        __props__.__dict__["update_time"] = update_time
        return PolicyByName(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the CAM policy.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the CAM policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def document(self) -> pulumi.Output[str]:
        """
        Document of the CAM policy. The syntax refers to [CAM
        POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
        terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
        Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        """
        return pulumi.get(self, "document")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of CAM policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[int]:
        """
        Type of the policy strategy. Valid values: `1`, `2`. `1` means customer strategy and `2` means preset strategy.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update time of the CAM policy.
        """
        return pulumi.get(self, "update_time")

