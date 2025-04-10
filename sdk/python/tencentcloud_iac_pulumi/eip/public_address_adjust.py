# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PublicAddressAdjustArgs', 'PublicAddressAdjust']

@pulumi.input_type
class PublicAddressAdjustArgs:
    def __init__(__self__, *,
                 address_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PublicAddressAdjust resource.
        :param pulumi.Input[str] address_id: A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        :param pulumi.Input[str] instance_id: A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        if address_id is not None:
            pulumi.set(__self__, "address_id", address_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="addressId")
    def address_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        """
        return pulumi.get(self, "address_id")

    @address_id.setter
    def address_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _PublicAddressAdjustState:
    def __init__(__self__, *,
                 address_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PublicAddressAdjust resources.
        :param pulumi.Input[str] address_id: A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        :param pulumi.Input[str] instance_id: A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        if address_id is not None:
            pulumi.set(__self__, "address_id", address_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="addressId")
    def address_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        """
        return pulumi.get(self, "address_id")

    @address_id.setter
    def address_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class PublicAddressAdjust(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a eip public_address_adjust

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        public_address_adjust = tencentcloud.eip.PublicAddressAdjust("publicAddressAdjust",
            address_id="eip-erft45fu",
            instance_id="ins-cr2rfq78")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_id: A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        :param pulumi.Input[str] instance_id: A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PublicAddressAdjustArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a eip public_address_adjust

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        public_address_adjust = tencentcloud.eip.PublicAddressAdjust("publicAddressAdjust",
            address_id="eip-erft45fu",
            instance_id="ins-cr2rfq78")
        ```

        :param str resource_name: The name of the resource.
        :param PublicAddressAdjustArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PublicAddressAdjustArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PublicAddressAdjustArgs.__new__(PublicAddressAdjustArgs)

            __props__.__dict__["address_id"] = address_id
            __props__.__dict__["instance_id"] = instance_id
        super(PublicAddressAdjust, __self__).__init__(
            'tencentcloud:Eip/publicAddressAdjust:PublicAddressAdjust',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'PublicAddressAdjust':
        """
        Get an existing PublicAddressAdjust resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_id: A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        :param pulumi.Input[str] instance_id: A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PublicAddressAdjustState.__new__(_PublicAddressAdjustState)

        __props__.__dict__["address_id"] = address_id
        __props__.__dict__["instance_id"] = instance_id
        return PublicAddressAdjust(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressId")
    def address_id(self) -> pulumi.Output[Optional[str]]:
        """
        A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
        """
        return pulumi.get(self, "address_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
        """
        return pulumi.get(self, "instance_id")

