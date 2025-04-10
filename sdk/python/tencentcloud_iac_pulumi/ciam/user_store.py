# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserStoreArgs', 'UserStore']

@pulumi.input_type
class UserStoreArgs:
    def __init__(__self__, *,
                 user_pool_name: pulumi.Input[str],
                 user_pool_desc: Optional[pulumi.Input[str]] = None,
                 user_pool_logo: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserStore resource.
        :param pulumi.Input[str] user_pool_name: User Store Name.
        :param pulumi.Input[str] user_pool_desc: User Store Description.
        :param pulumi.Input[str] user_pool_logo: User Store Logo.
        """
        pulumi.set(__self__, "user_pool_name", user_pool_name)
        if user_pool_desc is not None:
            pulumi.set(__self__, "user_pool_desc", user_pool_desc)
        if user_pool_logo is not None:
            pulumi.set(__self__, "user_pool_logo", user_pool_logo)

    @property
    @pulumi.getter(name="userPoolName")
    def user_pool_name(self) -> pulumi.Input[str]:
        """
        User Store Name.
        """
        return pulumi.get(self, "user_pool_name")

    @user_pool_name.setter
    def user_pool_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_pool_name", value)

    @property
    @pulumi.getter(name="userPoolDesc")
    def user_pool_desc(self) -> Optional[pulumi.Input[str]]:
        """
        User Store Description.
        """
        return pulumi.get(self, "user_pool_desc")

    @user_pool_desc.setter
    def user_pool_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_desc", value)

    @property
    @pulumi.getter(name="userPoolLogo")
    def user_pool_logo(self) -> Optional[pulumi.Input[str]]:
        """
        User Store Logo.
        """
        return pulumi.get(self, "user_pool_logo")

    @user_pool_logo.setter
    def user_pool_logo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_logo", value)


@pulumi.input_type
class _UserStoreState:
    def __init__(__self__, *,
                 user_pool_desc: Optional[pulumi.Input[str]] = None,
                 user_pool_logo: Optional[pulumi.Input[str]] = None,
                 user_pool_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserStore resources.
        :param pulumi.Input[str] user_pool_desc: User Store Description.
        :param pulumi.Input[str] user_pool_logo: User Store Logo.
        :param pulumi.Input[str] user_pool_name: User Store Name.
        """
        if user_pool_desc is not None:
            pulumi.set(__self__, "user_pool_desc", user_pool_desc)
        if user_pool_logo is not None:
            pulumi.set(__self__, "user_pool_logo", user_pool_logo)
        if user_pool_name is not None:
            pulumi.set(__self__, "user_pool_name", user_pool_name)

    @property
    @pulumi.getter(name="userPoolDesc")
    def user_pool_desc(self) -> Optional[pulumi.Input[str]]:
        """
        User Store Description.
        """
        return pulumi.get(self, "user_pool_desc")

    @user_pool_desc.setter
    def user_pool_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_desc", value)

    @property
    @pulumi.getter(name="userPoolLogo")
    def user_pool_logo(self) -> Optional[pulumi.Input[str]]:
        """
        User Store Logo.
        """
        return pulumi.get(self, "user_pool_logo")

    @user_pool_logo.setter
    def user_pool_logo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_logo", value)

    @property
    @pulumi.getter(name="userPoolName")
    def user_pool_name(self) -> Optional[pulumi.Input[str]]:
        """
        User Store Name.
        """
        return pulumi.get(self, "user_pool_name")

    @user_pool_name.setter
    def user_pool_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_name", value)


class UserStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 user_pool_desc: Optional[pulumi.Input[str]] = None,
                 user_pool_logo: Optional[pulumi.Input[str]] = None,
                 user_pool_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a ciam user store

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        user_store = tencentcloud.ciam.UserStore("userStore",
            user_pool_desc="for terraform test 123",
            user_pool_logo="https://ciam-prd-1302490086.cos.ap-guangzhou.myqcloud.com/temporary/92630252a2c5422d9663db5feafd619b.png",
            user_pool_name="tf_user_store")
        ```

        ## Import

        ciam user_store can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Ciam/userStore:UserStore user_store userStoreId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] user_pool_desc: User Store Description.
        :param pulumi.Input[str] user_pool_logo: User Store Logo.
        :param pulumi.Input[str] user_pool_name: User Store Name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserStoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a ciam user store

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        user_store = tencentcloud.ciam.UserStore("userStore",
            user_pool_desc="for terraform test 123",
            user_pool_logo="https://ciam-prd-1302490086.cos.ap-guangzhou.myqcloud.com/temporary/92630252a2c5422d9663db5feafd619b.png",
            user_pool_name="tf_user_store")
        ```

        ## Import

        ciam user_store can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Ciam/userStore:UserStore user_store userStoreId
        ```

        :param str resource_name: The name of the resource.
        :param UserStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 user_pool_desc: Optional[pulumi.Input[str]] = None,
                 user_pool_logo: Optional[pulumi.Input[str]] = None,
                 user_pool_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = UserStoreArgs.__new__(UserStoreArgs)

            __props__.__dict__["user_pool_desc"] = user_pool_desc
            __props__.__dict__["user_pool_logo"] = user_pool_logo
            if user_pool_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_pool_name'")
            __props__.__dict__["user_pool_name"] = user_pool_name
        super(UserStore, __self__).__init__(
            'tencentcloud:Ciam/userStore:UserStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            user_pool_desc: Optional[pulumi.Input[str]] = None,
            user_pool_logo: Optional[pulumi.Input[str]] = None,
            user_pool_name: Optional[pulumi.Input[str]] = None) -> 'UserStore':
        """
        Get an existing UserStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] user_pool_desc: User Store Description.
        :param pulumi.Input[str] user_pool_logo: User Store Logo.
        :param pulumi.Input[str] user_pool_name: User Store Name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserStoreState.__new__(_UserStoreState)

        __props__.__dict__["user_pool_desc"] = user_pool_desc
        __props__.__dict__["user_pool_logo"] = user_pool_logo
        __props__.__dict__["user_pool_name"] = user_pool_name
        return UserStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="userPoolDesc")
    def user_pool_desc(self) -> pulumi.Output[Optional[str]]:
        """
        User Store Description.
        """
        return pulumi.get(self, "user_pool_desc")

    @property
    @pulumi.getter(name="userPoolLogo")
    def user_pool_logo(self) -> pulumi.Output[Optional[str]]:
        """
        User Store Logo.
        """
        return pulumi.get(self, "user_pool_logo")

    @property
    @pulumi.getter(name="userPoolName")
    def user_pool_name(self) -> pulumi.Output[str]:
        """
        User Store Name.
        """
        return pulumi.get(self, "user_pool_name")

