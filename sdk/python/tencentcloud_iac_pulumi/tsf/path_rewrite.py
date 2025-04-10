# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PathRewriteArgs', 'PathRewrite']

@pulumi.input_type
class PathRewriteArgs:
    def __init__(__self__, *,
                 blocked: pulumi.Input[str],
                 gateway_group_id: pulumi.Input[str],
                 order: pulumi.Input[int],
                 regex: pulumi.Input[str],
                 replacement: pulumi.Input[str]):
        """
        The set of arguments for constructing a PathRewrite resource.
        :param pulumi.Input[str] blocked: Whether to shield the mapped path, Y: Yes N: No.
        :param pulumi.Input[str] gateway_group_id: gateway deployment group ID.
        :param pulumi.Input[int] order: rule order, the smaller the higher the priority.
        :param pulumi.Input[str] regex: regular expression.
        :param pulumi.Input[str] replacement: content to replace.
        """
        pulumi.set(__self__, "blocked", blocked)
        pulumi.set(__self__, "gateway_group_id", gateway_group_id)
        pulumi.set(__self__, "order", order)
        pulumi.set(__self__, "regex", regex)
        pulumi.set(__self__, "replacement", replacement)

    @property
    @pulumi.getter
    def blocked(self) -> pulumi.Input[str]:
        """
        Whether to shield the mapped path, Y: Yes N: No.
        """
        return pulumi.get(self, "blocked")

    @blocked.setter
    def blocked(self, value: pulumi.Input[str]):
        pulumi.set(self, "blocked", value)

    @property
    @pulumi.getter(name="gatewayGroupId")
    def gateway_group_id(self) -> pulumi.Input[str]:
        """
        gateway deployment group ID.
        """
        return pulumi.get(self, "gateway_group_id")

    @gateway_group_id.setter
    def gateway_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_group_id", value)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[int]:
        """
        rule order, the smaller the higher the priority.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def regex(self) -> pulumi.Input[str]:
        """
        regular expression.
        """
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: pulumi.Input[str]):
        pulumi.set(self, "regex", value)

    @property
    @pulumi.getter
    def replacement(self) -> pulumi.Input[str]:
        """
        content to replace.
        """
        return pulumi.get(self, "replacement")

    @replacement.setter
    def replacement(self, value: pulumi.Input[str]):
        pulumi.set(self, "replacement", value)


@pulumi.input_type
class _PathRewriteState:
    def __init__(__self__, *,
                 blocked: Optional[pulumi.Input[str]] = None,
                 gateway_group_id: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 path_rewrite_id: Optional[pulumi.Input[str]] = None,
                 regex: Optional[pulumi.Input[str]] = None,
                 replacement: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PathRewrite resources.
        :param pulumi.Input[str] blocked: Whether to shield the mapped path, Y: Yes N: No.
        :param pulumi.Input[str] gateway_group_id: gateway deployment group ID.
        :param pulumi.Input[int] order: rule order, the smaller the higher the priority.
        :param pulumi.Input[str] path_rewrite_id: path rewrite rule ID.
        :param pulumi.Input[str] regex: regular expression.
        :param pulumi.Input[str] replacement: content to replace.
        """
        if blocked is not None:
            pulumi.set(__self__, "blocked", blocked)
        if gateway_group_id is not None:
            pulumi.set(__self__, "gateway_group_id", gateway_group_id)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if path_rewrite_id is not None:
            pulumi.set(__self__, "path_rewrite_id", path_rewrite_id)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)
        if replacement is not None:
            pulumi.set(__self__, "replacement", replacement)

    @property
    @pulumi.getter
    def blocked(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to shield the mapped path, Y: Yes N: No.
        """
        return pulumi.get(self, "blocked")

    @blocked.setter
    def blocked(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blocked", value)

    @property
    @pulumi.getter(name="gatewayGroupId")
    def gateway_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        gateway deployment group ID.
        """
        return pulumi.get(self, "gateway_group_id")

    @gateway_group_id.setter
    def gateway_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_group_id", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        rule order, the smaller the higher the priority.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="pathRewriteId")
    def path_rewrite_id(self) -> Optional[pulumi.Input[str]]:
        """
        path rewrite rule ID.
        """
        return pulumi.get(self, "path_rewrite_id")

    @path_rewrite_id.setter
    def path_rewrite_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path_rewrite_id", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[pulumi.Input[str]]:
        """
        regular expression.
        """
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "regex", value)

    @property
    @pulumi.getter
    def replacement(self) -> Optional[pulumi.Input[str]]:
        """
        content to replace.
        """
        return pulumi.get(self, "replacement")

    @replacement.setter
    def replacement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replacement", value)


class PathRewrite(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked: Optional[pulumi.Input[str]] = None,
                 gateway_group_id: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 regex: Optional[pulumi.Input[str]] = None,
                 replacement: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tsf path_rewrite

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        path_rewrite = tencentcloud.tsf.PathRewrite("pathRewrite",
            blocked="N",
            gateway_group_id="group-a2j9zxpv",
            order=2,
            regex="/test",
            replacement="/tt")
        ```

        ## Import

        tsf path_rewrite can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tsf/pathRewrite:PathRewrite path_rewrite rewrite-nygq33v2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blocked: Whether to shield the mapped path, Y: Yes N: No.
        :param pulumi.Input[str] gateway_group_id: gateway deployment group ID.
        :param pulumi.Input[int] order: rule order, the smaller the higher the priority.
        :param pulumi.Input[str] regex: regular expression.
        :param pulumi.Input[str] replacement: content to replace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PathRewriteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tsf path_rewrite

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        path_rewrite = tencentcloud.tsf.PathRewrite("pathRewrite",
            blocked="N",
            gateway_group_id="group-a2j9zxpv",
            order=2,
            regex="/test",
            replacement="/tt")
        ```

        ## Import

        tsf path_rewrite can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tsf/pathRewrite:PathRewrite path_rewrite rewrite-nygq33v2
        ```

        :param str resource_name: The name of the resource.
        :param PathRewriteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PathRewriteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked: Optional[pulumi.Input[str]] = None,
                 gateway_group_id: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 regex: Optional[pulumi.Input[str]] = None,
                 replacement: Optional[pulumi.Input[str]] = None,
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
            __props__ = PathRewriteArgs.__new__(PathRewriteArgs)

            if blocked is None and not opts.urn:
                raise TypeError("Missing required property 'blocked'")
            __props__.__dict__["blocked"] = blocked
            if gateway_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_group_id'")
            __props__.__dict__["gateway_group_id"] = gateway_group_id
            if order is None and not opts.urn:
                raise TypeError("Missing required property 'order'")
            __props__.__dict__["order"] = order
            if regex is None and not opts.urn:
                raise TypeError("Missing required property 'regex'")
            __props__.__dict__["regex"] = regex
            if replacement is None and not opts.urn:
                raise TypeError("Missing required property 'replacement'")
            __props__.__dict__["replacement"] = replacement
            __props__.__dict__["path_rewrite_id"] = None
        super(PathRewrite, __self__).__init__(
            'tencentcloud:Tsf/pathRewrite:PathRewrite',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blocked: Optional[pulumi.Input[str]] = None,
            gateway_group_id: Optional[pulumi.Input[str]] = None,
            order: Optional[pulumi.Input[int]] = None,
            path_rewrite_id: Optional[pulumi.Input[str]] = None,
            regex: Optional[pulumi.Input[str]] = None,
            replacement: Optional[pulumi.Input[str]] = None) -> 'PathRewrite':
        """
        Get an existing PathRewrite resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blocked: Whether to shield the mapped path, Y: Yes N: No.
        :param pulumi.Input[str] gateway_group_id: gateway deployment group ID.
        :param pulumi.Input[int] order: rule order, the smaller the higher the priority.
        :param pulumi.Input[str] path_rewrite_id: path rewrite rule ID.
        :param pulumi.Input[str] regex: regular expression.
        :param pulumi.Input[str] replacement: content to replace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PathRewriteState.__new__(_PathRewriteState)

        __props__.__dict__["blocked"] = blocked
        __props__.__dict__["gateway_group_id"] = gateway_group_id
        __props__.__dict__["order"] = order
        __props__.__dict__["path_rewrite_id"] = path_rewrite_id
        __props__.__dict__["regex"] = regex
        __props__.__dict__["replacement"] = replacement
        return PathRewrite(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def blocked(self) -> pulumi.Output[str]:
        """
        Whether to shield the mapped path, Y: Yes N: No.
        """
        return pulumi.get(self, "blocked")

    @property
    @pulumi.getter(name="gatewayGroupId")
    def gateway_group_id(self) -> pulumi.Output[str]:
        """
        gateway deployment group ID.
        """
        return pulumi.get(self, "gateway_group_id")

    @property
    @pulumi.getter
    def order(self) -> pulumi.Output[int]:
        """
        rule order, the smaller the higher the priority.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="pathRewriteId")
    def path_rewrite_id(self) -> pulumi.Output[str]:
        """
        path rewrite rule ID.
        """
        return pulumi.get(self, "path_rewrite_id")

    @property
    @pulumi.getter
    def regex(self) -> pulumi.Output[str]:
        """
        regular expression.
        """
        return pulumi.get(self, "regex")

    @property
    @pulumi.getter
    def replacement(self) -> pulumi.Output[str]:
        """
        content to replace.
        """
        return pulumi.get(self, "replacement")

