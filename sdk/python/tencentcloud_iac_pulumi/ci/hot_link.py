# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HotLinkArgs', 'HotLink']

@pulumi.input_type
class HotLinkArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 type: pulumi.Input[str],
                 urls: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a HotLink resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] type: Anti-leech type, `white` is whitelist, `black` is blacklist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] urls: domain address.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "urls", urls)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Anti-leech type, `white` is whitelist, `black` is blacklist.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def urls(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        domain address.
        """
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "urls", value)


@pulumi.input_type
class _HotLinkState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering HotLink resources.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] type: Anti-leech type, `white` is whitelist, `black` is blacklist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] urls: domain address.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Anti-leech type, `white` is whitelist, `black` is blacklist.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        domain address.
        """
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "urls", value)


class HotLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a ci hot_link

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        hot_link = tencentcloud.ci.HotLink("hotLink",
            bucket="terraform-ci-xxxxxx",
            type="white",
            urls=[
                "10.0.0.1",
                "10.0.0.2",
            ])
        ```

        ## Import

        ci hot_link can be imported using the bucket, e.g.

        ```sh
         $ pulumi import tencentcloud:Ci/hotLink:HotLink hot_link terraform-ci-xxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] type: Anti-leech type, `white` is whitelist, `black` is blacklist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] urls: domain address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HotLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a ci hot_link

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        hot_link = tencentcloud.ci.HotLink("hotLink",
            bucket="terraform-ci-xxxxxx",
            type="white",
            urls=[
                "10.0.0.1",
                "10.0.0.2",
            ])
        ```

        ## Import

        ci hot_link can be imported using the bucket, e.g.

        ```sh
         $ pulumi import tencentcloud:Ci/hotLink:HotLink hot_link terraform-ci-xxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param HotLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HotLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = HotLinkArgs.__new__(HotLinkArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if urls is None and not opts.urn:
                raise TypeError("Missing required property 'urls'")
            __props__.__dict__["urls"] = urls
        super(HotLink, __self__).__init__(
            'tencentcloud:Ci/hotLink:HotLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'HotLink':
        """
        Get an existing HotLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] type: Anti-leech type, `white` is whitelist, `black` is blacklist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] urls: domain address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HotLinkState.__new__(_HotLinkState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["type"] = type
        __props__.__dict__["urls"] = urls
        return HotLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Anti-leech type, `white` is whitelist, `black` is blacklist.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def urls(self) -> pulumi.Output[Sequence[str]]:
        """
        domain address.
        """
        return pulumi.get(self, "urls")

