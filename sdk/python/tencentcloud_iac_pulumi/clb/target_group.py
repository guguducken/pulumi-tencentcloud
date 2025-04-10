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

__all__ = ['TargetGroupArgs', 'TargetGroup']

@pulumi.input_type
class TargetGroupArgs:
    def __init__(__self__, *,
                 port: Optional[pulumi.Input[int]] = None,
                 target_group_instances: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]]] = None,
                 target_group_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TargetGroup resource.
        :param pulumi.Input[int] port: The default port of target group, add server after can use it.
        :param pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]] target_group_instances: It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        :param pulumi.Input[str] target_group_name: Target group name.
        :param pulumi.Input[str] vpc_id: VPC ID, default is based on the network.
        """
        if port is not None:
            pulumi.set(__self__, "port", port)
        if target_group_instances is not None:
            warnings.warn("""It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.""", DeprecationWarning)
            pulumi.log.warn("""target_group_instances is deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.""")
        if target_group_instances is not None:
            pulumi.set(__self__, "target_group_instances", target_group_instances)
        if target_group_name is not None:
            pulumi.set(__self__, "target_group_name", target_group_name)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The default port of target group, add server after can use it.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="targetGroupInstances")
    def target_group_instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]]]:
        """
        It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        """
        return pulumi.get(self, "target_group_instances")

    @target_group_instances.setter
    def target_group_instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]]]):
        pulumi.set(self, "target_group_instances", value)

    @property
    @pulumi.getter(name="targetGroupName")
    def target_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Target group name.
        """
        return pulumi.get(self, "target_group_name")

    @target_group_name.setter
    def target_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_group_name", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID, default is based on the network.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _TargetGroupState:
    def __init__(__self__, *,
                 port: Optional[pulumi.Input[int]] = None,
                 target_group_instances: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]]] = None,
                 target_group_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TargetGroup resources.
        :param pulumi.Input[int] port: The default port of target group, add server after can use it.
        :param pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]] target_group_instances: It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        :param pulumi.Input[str] target_group_name: Target group name.
        :param pulumi.Input[str] vpc_id: VPC ID, default is based on the network.
        """
        if port is not None:
            pulumi.set(__self__, "port", port)
        if target_group_instances is not None:
            warnings.warn("""It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.""", DeprecationWarning)
            pulumi.log.warn("""target_group_instances is deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.""")
        if target_group_instances is not None:
            pulumi.set(__self__, "target_group_instances", target_group_instances)
        if target_group_name is not None:
            pulumi.set(__self__, "target_group_name", target_group_name)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The default port of target group, add server after can use it.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="targetGroupInstances")
    def target_group_instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]]]:
        """
        It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        """
        return pulumi.get(self, "target_group_instances")

    @target_group_instances.setter
    def target_group_instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetGroupInstanceArgs']]]]):
        pulumi.set(self, "target_group_instances", value)

    @property
    @pulumi.getter(name="targetGroupName")
    def target_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Target group name.
        """
        return pulumi.get(self, "target_group_name")

    @target_group_name.setter
    def target_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_group_name", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID, default is based on the network.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class TargetGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 target_group_instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetGroupInstanceArgs']]]]] = None,
                 target_group_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a CLB target group.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        test = tencentcloud.clb.TargetGroup("test",
            port=33,
            target_group_name="test")
        ```

        ## Import

        CLB target group can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clb/targetGroup:TargetGroup test lbtg-3k3io0i0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] port: The default port of target group, add server after can use it.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetGroupInstanceArgs']]]] target_group_instances: It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        :param pulumi.Input[str] target_group_name: Target group name.
        :param pulumi.Input[str] vpc_id: VPC ID, default is based on the network.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TargetGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a CLB target group.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        test = tencentcloud.clb.TargetGroup("test",
            port=33,
            target_group_name="test")
        ```

        ## Import

        CLB target group can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clb/targetGroup:TargetGroup test lbtg-3k3io0i0
        ```

        :param str resource_name: The name of the resource.
        :param TargetGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 target_group_instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetGroupInstanceArgs']]]]] = None,
                 target_group_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = TargetGroupArgs.__new__(TargetGroupArgs)

            __props__.__dict__["port"] = port
            if target_group_instances is not None and not opts.urn:
                warnings.warn("""It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.""", DeprecationWarning)
                pulumi.log.warn("""target_group_instances is deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.""")
            __props__.__dict__["target_group_instances"] = target_group_instances
            __props__.__dict__["target_group_name"] = target_group_name
            __props__.__dict__["vpc_id"] = vpc_id
        super(TargetGroup, __self__).__init__(
            'tencentcloud:Clb/targetGroup:TargetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            port: Optional[pulumi.Input[int]] = None,
            target_group_instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetGroupInstanceArgs']]]]] = None,
            target_group_name: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'TargetGroup':
        """
        Get an existing TargetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] port: The default port of target group, add server after can use it.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetGroupInstanceArgs']]]] target_group_instances: It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        :param pulumi.Input[str] target_group_name: Target group name.
        :param pulumi.Input[str] vpc_id: VPC ID, default is based on the network.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TargetGroupState.__new__(_TargetGroupState)

        __props__.__dict__["port"] = port
        __props__.__dict__["target_group_instances"] = target_group_instances
        __props__.__dict__["target_group_name"] = target_group_name
        __props__.__dict__["vpc_id"] = vpc_id
        return TargetGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The default port of target group, add server after can use it.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="targetGroupInstances")
    def target_group_instances(self) -> pulumi.Output[Optional[Sequence['outputs.TargetGroupTargetGroupInstance']]]:
        """
        It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
        """
        return pulumi.get(self, "target_group_instances")

    @property
    @pulumi.getter(name="targetGroupName")
    def target_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        Target group name.
        """
        return pulumi.get(self, "target_group_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        VPC ID, default is based on the network.
        """
        return pulumi.get(self, "vpc_id")

