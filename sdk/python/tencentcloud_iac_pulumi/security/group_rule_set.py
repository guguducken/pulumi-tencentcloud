# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GroupRuleSetArgs', 'GroupRuleSet']

@pulumi.input_type
class GroupRuleSetArgs:
    def __init__(__self__, *,
                 security_group_id: pulumi.Input[str],
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]]] = None):
        """
        The set of arguments for constructing a GroupRuleSet resource.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]] egresses: List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]] ingresses: List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        pulumi.set(__self__, "security_group_id", security_group_id)
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        ID of the security group to be queried.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]]]:
        """
        List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]]]:
        """
        List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)


@pulumi.input_type
class _GroupRuleSetState:
    def __init__(__self__, *,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupRuleSet resources.
        :param pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]] egresses: List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]] ingresses: List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[str] version: Security policies version, auto increment for every update.
        """
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]]]:
        """
        List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]]]:
        """
        List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupRuleSetIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the security group to be queried.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Security policies version, auto increment for every update.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class GroupRuleSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetEgressArgs']]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetIngressArgs']]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create security group rule. This resource is similar with tencentcloud_security_group_lite_rule, rules can be ordered and configure descriptions.

        > **NOTE:** This resource must exclusive in one security group, do not declare additional rule resources of this security group elsewhere.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sglab1_group = tencentcloud.security.Group("sglab1Group", description="favourite sg_1")
        sglab1_group_rule_set = tencentcloud.security.GroupRuleSet("sglab1GroupRuleSet",
            security_group_id=sglab1_group.id,
            ingresses=[
                tencentcloud.security.GroupRuleSetIngressArgs(
                    cidr_block="10.0.0.0/16",
                    protocol="TCP",
                    port="80",
                    action="ACCEPT",
                    description="favourite sg rule_1",
                ),
                tencentcloud.security.GroupRuleSetIngressArgs(
                    protocol="TCP",
                    port="80",
                    action="ACCEPT",
                    source_security_id=tencentcloud_security_group["sglab_3"]["id"],
                    description="favourite sg rule_2",
                ),
            ],
            egresses=[
                tencentcloud.security.GroupRuleSetEgressArgs(
                    action="ACCEPT",
                    address_template_id="ipm-xxxxxxxx",
                    description="Allow address template",
                ),
                tencentcloud.security.GroupRuleSetEgressArgs(
                    action="ACCEPT",
                    service_template_group="ppmg-xxxxxxxx",
                    description="Allow protocol template",
                ),
                tencentcloud.security.GroupRuleSetEgressArgs(
                    cidr_block="10.0.0.0/16",
                    protocol="TCP",
                    port="80",
                    action="DROP",
                    description="favourite sg egress rule",
                ),
            ])
        ```

        ## Import

        Resource tencentcloud_security_group_rule_set can be imported by passing security grou id

        ```sh
         $ pulumi import tencentcloud:Security/groupRuleSet:GroupRuleSet sglab_1 sg-xxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetEgressArgs']]]] egresses: List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetIngressArgs']]]] ingresses: List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupRuleSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create security group rule. This resource is similar with tencentcloud_security_group_lite_rule, rules can be ordered and configure descriptions.

        > **NOTE:** This resource must exclusive in one security group, do not declare additional rule resources of this security group elsewhere.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sglab1_group = tencentcloud.security.Group("sglab1Group", description="favourite sg_1")
        sglab1_group_rule_set = tencentcloud.security.GroupRuleSet("sglab1GroupRuleSet",
            security_group_id=sglab1_group.id,
            ingresses=[
                tencentcloud.security.GroupRuleSetIngressArgs(
                    cidr_block="10.0.0.0/16",
                    protocol="TCP",
                    port="80",
                    action="ACCEPT",
                    description="favourite sg rule_1",
                ),
                tencentcloud.security.GroupRuleSetIngressArgs(
                    protocol="TCP",
                    port="80",
                    action="ACCEPT",
                    source_security_id=tencentcloud_security_group["sglab_3"]["id"],
                    description="favourite sg rule_2",
                ),
            ],
            egresses=[
                tencentcloud.security.GroupRuleSetEgressArgs(
                    action="ACCEPT",
                    address_template_id="ipm-xxxxxxxx",
                    description="Allow address template",
                ),
                tencentcloud.security.GroupRuleSetEgressArgs(
                    action="ACCEPT",
                    service_template_group="ppmg-xxxxxxxx",
                    description="Allow protocol template",
                ),
                tencentcloud.security.GroupRuleSetEgressArgs(
                    cidr_block="10.0.0.0/16",
                    protocol="TCP",
                    port="80",
                    action="DROP",
                    description="favourite sg egress rule",
                ),
            ])
        ```

        ## Import

        Resource tencentcloud_security_group_rule_set can be imported by passing security grou id

        ```sh
         $ pulumi import tencentcloud:Security/groupRuleSet:GroupRuleSet sglab_1 sg-xxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param GroupRuleSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupRuleSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetEgressArgs']]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetIngressArgs']]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = GroupRuleSetArgs.__new__(GroupRuleSetArgs)

            __props__.__dict__["egresses"] = egresses
            __props__.__dict__["ingresses"] = ingresses
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
            __props__.__dict__["version"] = None
        super(GroupRuleSet, __self__).__init__(
            'tencentcloud:Security/groupRuleSet:GroupRuleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetEgressArgs']]]]] = None,
            ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetIngressArgs']]]]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'GroupRuleSet':
        """
        Get an existing GroupRuleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetEgressArgs']]]] egresses: List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupRuleSetIngressArgs']]]] ingresses: List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[str] version: Security policies version, auto increment for every update.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupRuleSetState.__new__(_GroupRuleSetState)

        __props__.__dict__["egresses"] = egresses
        __props__.__dict__["ingresses"] = ingresses
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["version"] = version
        return GroupRuleSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def egresses(self) -> pulumi.Output[Optional[Sequence['outputs.GroupRuleSetEgress']]]:
        """
        List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        return pulumi.get(self, "egresses")

    @property
    @pulumi.getter
    def ingresses(self) -> pulumi.Output[Optional[Sequence['outputs.GroupRuleSetIngress']]]:
        """
        List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.
        """
        return pulumi.get(self, "ingresses")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        ID of the security group to be queried.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Security policies version, auto increment for every update.
        """
        return pulumi.get(self, "version")

