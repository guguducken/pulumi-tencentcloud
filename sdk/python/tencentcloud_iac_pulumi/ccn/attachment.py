# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AttachmentArgs', 'Attachment']

@pulumi.input_type
class AttachmentArgs:
    def __init__(__self__, *,
                 ccn_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 instance_region: pulumi.Input[str],
                 instance_type: pulumi.Input[str],
                 ccn_uin: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Attachment resource.
        :param pulumi.Input[str] ccn_id: ID of the CCN.
        :param pulumi.Input[str] instance_id: ID of instance is attached.
        :param pulumi.Input[str] instance_region: The region that the instance locates at.
        :param pulumi.Input[str] instance_type: Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        :param pulumi.Input[str] ccn_uin: Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        :param pulumi.Input[str] description: Remark of attachment.
        """
        pulumi.set(__self__, "ccn_id", ccn_id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        pulumi.set(__self__, "instance_type", instance_type)
        if ccn_uin is not None:
            pulumi.set(__self__, "ccn_uin", ccn_uin)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> pulumi.Input[str]:
        """
        ID of the CCN.
        """
        return pulumi.get(self, "ccn_id")

    @ccn_id.setter
    def ccn_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ccn_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of instance is attached.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> pulumi.Input[str]:
        """
        The region that the instance locates at.
        """
        return pulumi.get(self, "instance_region")

    @instance_region.setter
    def instance_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_region", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="ccnUin")
    def ccn_uin(self) -> Optional[pulumi.Input[str]]:
        """
        Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        """
        return pulumi.get(self, "ccn_uin")

    @ccn_uin.setter
    def ccn_uin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ccn_uin", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Remark of attachment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _AttachmentState:
    def __init__(__self__, *,
                 attached_time: Optional[pulumi.Input[str]] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 ccn_uin: Optional[pulumi.Input[str]] = None,
                 cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Attachment resources.
        :param pulumi.Input[str] attached_time: Time of attaching.
        :param pulumi.Input[str] ccn_id: ID of the CCN.
        :param pulumi.Input[str] ccn_uin: Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_blocks: A network address block of the instance that is attached.
        :param pulumi.Input[str] description: Remark of attachment.
        :param pulumi.Input[str] instance_id: ID of instance is attached.
        :param pulumi.Input[str] instance_region: The region that the instance locates at.
        :param pulumi.Input[str] instance_type: Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        :param pulumi.Input[str] state: States of instance is attached. Valid values: `PENDING`, `ACTIVE`, `EXPIRED`, `REJECTED`, `DELETED`, `FAILED`, `ATTACHING`, `DETACHING` and `DETACHFAILED`. `FAILED` means asynchronous forced disassociation after 2 hours. `DETACHFAILED` means asynchronous forced disassociation after 2 hours.
        """
        if attached_time is not None:
            pulumi.set(__self__, "attached_time", attached_time)
        if ccn_id is not None:
            pulumi.set(__self__, "ccn_id", ccn_id)
        if ccn_uin is not None:
            pulumi.set(__self__, "ccn_uin", ccn_uin)
        if cidr_blocks is not None:
            pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_region is not None:
            pulumi.set(__self__, "instance_region", instance_region)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="attachedTime")
    def attached_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of attaching.
        """
        return pulumi.get(self, "attached_time")

    @attached_time.setter
    def attached_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attached_time", value)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CCN.
        """
        return pulumi.get(self, "ccn_id")

    @ccn_id.setter
    def ccn_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ccn_id", value)

    @property
    @pulumi.getter(name="ccnUin")
    def ccn_uin(self) -> Optional[pulumi.Input[str]]:
        """
        Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        """
        return pulumi.get(self, "ccn_uin")

    @ccn_uin.setter
    def ccn_uin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ccn_uin", value)

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A network address block of the instance that is attached.
        """
        return pulumi.get(self, "cidr_blocks")

    @cidr_blocks.setter
    def cidr_blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_blocks", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Remark of attachment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of instance is attached.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> Optional[pulumi.Input[str]]:
        """
        The region that the instance locates at.
        """
        return pulumi.get(self, "instance_region")

    @instance_region.setter
    def instance_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_region", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        States of instance is attached. Valid values: `PENDING`, `ACTIVE`, `EXPIRED`, `REJECTED`, `DELETED`, `FAILED`, `ATTACHING`, `DETACHING` and `DETACHFAILED`. `FAILED` means asynchronous forced disassociation after 2 hours. `DETACHFAILED` means asynchronous forced disassociation after 2 hours.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class Attachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 ccn_uin: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a CCN attaching resource.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        region = config.get("region")
        if region is None:
            region = "ap-guangzhou"
        otheruin = config.get("otheruin")
        if otheruin is None:
            otheruin = "123353"
        otherccn = config.get("otherccn")
        if otherccn is None:
            otherccn = "ccn-151ssaga"
        vpc = tencentcloud.vpc.Instance("vpc",
            cidr_block="10.0.0.0/16",
            dns_servers=[
                "119.29.29.29",
                "8.8.8.8",
            ],
            is_multicast=False)
        main = tencentcloud.ccn.Instance("main",
            description="ci-temp-test-ccn-des",
            qos="AG")
        attachment = tencentcloud.ccn.Attachment("attachment",
            ccn_id=main.id,
            instance_type="VPC",
            instance_id=vpc.id,
            instance_region=region)
        other_account = tencentcloud.ccn.Attachment("otherAccount",
            ccn_id=otherccn,
            instance_type="VPC",
            instance_id=vpc.id,
            instance_region=region,
            ccn_uin=otheruin)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ccn_id: ID of the CCN.
        :param pulumi.Input[str] ccn_uin: Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        :param pulumi.Input[str] description: Remark of attachment.
        :param pulumi.Input[str] instance_id: ID of instance is attached.
        :param pulumi.Input[str] instance_region: The region that the instance locates at.
        :param pulumi.Input[str] instance_type: Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CCN attaching resource.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        region = config.get("region")
        if region is None:
            region = "ap-guangzhou"
        otheruin = config.get("otheruin")
        if otheruin is None:
            otheruin = "123353"
        otherccn = config.get("otherccn")
        if otherccn is None:
            otherccn = "ccn-151ssaga"
        vpc = tencentcloud.vpc.Instance("vpc",
            cidr_block="10.0.0.0/16",
            dns_servers=[
                "119.29.29.29",
                "8.8.8.8",
            ],
            is_multicast=False)
        main = tencentcloud.ccn.Instance("main",
            description="ci-temp-test-ccn-des",
            qos="AG")
        attachment = tencentcloud.ccn.Attachment("attachment",
            ccn_id=main.id,
            instance_type="VPC",
            instance_id=vpc.id,
            instance_region=region)
        other_account = tencentcloud.ccn.Attachment("otherAccount",
            ccn_id=otherccn,
            instance_type="VPC",
            instance_id=vpc.id,
            instance_region=region,
            ccn_uin=otheruin)
        ```

        :param str resource_name: The name of the resource.
        :param AttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 ccn_uin: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = AttachmentArgs.__new__(AttachmentArgs)

            if ccn_id is None and not opts.urn:
                raise TypeError("Missing required property 'ccn_id'")
            __props__.__dict__["ccn_id"] = ccn_id
            __props__.__dict__["ccn_uin"] = ccn_uin
            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if instance_region is None and not opts.urn:
                raise TypeError("Missing required property 'instance_region'")
            __props__.__dict__["instance_region"] = instance_region
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["attached_time"] = None
            __props__.__dict__["cidr_blocks"] = None
            __props__.__dict__["state"] = None
        super(Attachment, __self__).__init__(
            'tencentcloud:Ccn/attachment:Attachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attached_time: Optional[pulumi.Input[str]] = None,
            ccn_id: Optional[pulumi.Input[str]] = None,
            ccn_uin: Optional[pulumi.Input[str]] = None,
            cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_region: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'Attachment':
        """
        Get an existing Attachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attached_time: Time of attaching.
        :param pulumi.Input[str] ccn_id: ID of the CCN.
        :param pulumi.Input[str] ccn_uin: Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_blocks: A network address block of the instance that is attached.
        :param pulumi.Input[str] description: Remark of attachment.
        :param pulumi.Input[str] instance_id: ID of instance is attached.
        :param pulumi.Input[str] instance_region: The region that the instance locates at.
        :param pulumi.Input[str] instance_type: Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        :param pulumi.Input[str] state: States of instance is attached. Valid values: `PENDING`, `ACTIVE`, `EXPIRED`, `REJECTED`, `DELETED`, `FAILED`, `ATTACHING`, `DETACHING` and `DETACHFAILED`. `FAILED` means asynchronous forced disassociation after 2 hours. `DETACHFAILED` means asynchronous forced disassociation after 2 hours.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AttachmentState.__new__(_AttachmentState)

        __props__.__dict__["attached_time"] = attached_time
        __props__.__dict__["ccn_id"] = ccn_id
        __props__.__dict__["ccn_uin"] = ccn_uin
        __props__.__dict__["cidr_blocks"] = cidr_blocks
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_region"] = instance_region
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["state"] = state
        return Attachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachedTime")
    def attached_time(self) -> pulumi.Output[str]:
        """
        Time of attaching.
        """
        return pulumi.get(self, "attached_time")

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> pulumi.Output[str]:
        """
        ID of the CCN.
        """
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="ccnUin")
    def ccn_uin(self) -> pulumi.Output[str]:
        """
        Uin of the ccn attached. Default is ``, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        """
        return pulumi.get(self, "ccn_uin")

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> pulumi.Output[Sequence[str]]:
        """
        A network address block of the instance that is attached.
        """
        return pulumi.get(self, "cidr_blocks")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Remark of attachment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of instance is attached.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> pulumi.Output[str]:
        """
        The region that the instance locates at.
        """
        return pulumi.get(self, "instance_region")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note: `VPNGW` type is only for whitelist customer now.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        States of instance is attached. Valid values: `PENDING`, `ACTIVE`, `EXPIRED`, `REJECTED`, `DELETED`, `FAILED`, `ATTACHING`, `DETACHING` and `DETACHFAILED`. `FAILED` means asynchronous forced disassociation after 2 hours. `DETACHFAILED` means asynchronous forced disassociation after 2 hours.
        """
        return pulumi.get(self, "state")

