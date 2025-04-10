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

__all__ = ['TagRetentionRuleArgs', 'TagRetentionRule']

@pulumi.input_type
class TagRetentionRuleArgs:
    def __init__(__self__, *,
                 cron_setting: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 registry_id: pulumi.Input[str],
                 retention_rule: pulumi.Input['TagRetentionRuleRetentionRuleArgs'],
                 disabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TagRetentionRule resource.
        :param pulumi.Input[str] cron_setting: Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        :param pulumi.Input[str] namespace_name: The Name of the namespace.
        :param pulumi.Input[str] registry_id: The main instance ID.
        :param pulumi.Input['TagRetentionRuleRetentionRuleArgs'] retention_rule: Retention Policy.
        :param pulumi.Input[bool] disabled: Whether to disable the rule, with the default value of false.
        """
        pulumi.set(__self__, "cron_setting", cron_setting)
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "registry_id", registry_id)
        pulumi.set(__self__, "retention_rule", retention_rule)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter(name="cronSetting")
    def cron_setting(self) -> pulumi.Input[str]:
        """
        Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        """
        return pulumi.get(self, "cron_setting")

    @cron_setting.setter
    def cron_setting(self, value: pulumi.Input[str]):
        pulumi.set(self, "cron_setting", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        The Name of the namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Input[str]:
        """
        The main instance ID.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="retentionRule")
    def retention_rule(self) -> pulumi.Input['TagRetentionRuleRetentionRuleArgs']:
        """
        Retention Policy.
        """
        return pulumi.get(self, "retention_rule")

    @retention_rule.setter
    def retention_rule(self, value: pulumi.Input['TagRetentionRuleRetentionRuleArgs']):
        pulumi.set(self, "retention_rule", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to disable the rule, with the default value of false.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)


@pulumi.input_type
class _TagRetentionRuleState:
    def __init__(__self__, *,
                 cron_setting: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 retention_id: Optional[pulumi.Input[int]] = None,
                 retention_rule: Optional[pulumi.Input['TagRetentionRuleRetentionRuleArgs']] = None):
        """
        Input properties used for looking up and filtering TagRetentionRule resources.
        :param pulumi.Input[str] cron_setting: Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        :param pulumi.Input[bool] disabled: Whether to disable the rule, with the default value of false.
        :param pulumi.Input[str] namespace_name: The Name of the namespace.
        :param pulumi.Input[str] registry_id: The main instance ID.
        :param pulumi.Input[int] retention_id: The ID of the retention task.
        :param pulumi.Input['TagRetentionRuleRetentionRuleArgs'] retention_rule: Retention Policy.
        """
        if cron_setting is not None:
            pulumi.set(__self__, "cron_setting", cron_setting)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if registry_id is not None:
            pulumi.set(__self__, "registry_id", registry_id)
        if retention_id is not None:
            pulumi.set(__self__, "retention_id", retention_id)
        if retention_rule is not None:
            pulumi.set(__self__, "retention_rule", retention_rule)

    @property
    @pulumi.getter(name="cronSetting")
    def cron_setting(self) -> Optional[pulumi.Input[str]]:
        """
        Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        """
        return pulumi.get(self, "cron_setting")

    @cron_setting.setter
    def cron_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_setting", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to disable the rule, with the default value of false.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[str]]:
        """
        The main instance ID.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="retentionId")
    def retention_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the retention task.
        """
        return pulumi.get(self, "retention_id")

    @retention_id.setter
    def retention_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_id", value)

    @property
    @pulumi.getter(name="retentionRule")
    def retention_rule(self) -> Optional[pulumi.Input['TagRetentionRuleRetentionRuleArgs']]:
        """
        Retention Policy.
        """
        return pulumi.get(self, "retention_rule")

    @retention_rule.setter
    def retention_rule(self, value: Optional[pulumi.Input['TagRetentionRuleRetentionRuleArgs']]):
        pulumi.set(self, "retention_rule", value)


class TagRetentionRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_setting: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 retention_rule: Optional[pulumi.Input[pulumi.InputType['TagRetentionRuleRetentionRuleArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a tcr tag_retention_rule

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_ns = tencentcloud.tcr.Namespace("myNs",
            instance_id=local["tcr_id"],
            is_public=True,
            is_auto_scan=True,
            is_prevent_vul=True,
            severity="medium",
            cve_whitelist_items=[tencentcloud.tcr.NamespaceCveWhitelistItemArgs(
                cve_id="cve-xxxxx",
            )])
        my_rule = tencentcloud.tcr.TagRetentionRule("myRule",
            registry_id=local["tcr_id"],
            namespace_name=my_ns.name,
            retention_rule=tencentcloud.tcr.TagRetentionRuleRetentionRuleArgs(
                key="nDaysSinceLastPush",
                value=2,
            ),
            cron_setting="daily",
            disabled=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_setting: Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        :param pulumi.Input[bool] disabled: Whether to disable the rule, with the default value of false.
        :param pulumi.Input[str] namespace_name: The Name of the namespace.
        :param pulumi.Input[str] registry_id: The main instance ID.
        :param pulumi.Input[pulumi.InputType['TagRetentionRuleRetentionRuleArgs']] retention_rule: Retention Policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagRetentionRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tcr tag_retention_rule

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_ns = tencentcloud.tcr.Namespace("myNs",
            instance_id=local["tcr_id"],
            is_public=True,
            is_auto_scan=True,
            is_prevent_vul=True,
            severity="medium",
            cve_whitelist_items=[tencentcloud.tcr.NamespaceCveWhitelistItemArgs(
                cve_id="cve-xxxxx",
            )])
        my_rule = tencentcloud.tcr.TagRetentionRule("myRule",
            registry_id=local["tcr_id"],
            namespace_name=my_ns.name,
            retention_rule=tencentcloud.tcr.TagRetentionRuleRetentionRuleArgs(
                key="nDaysSinceLastPush",
                value=2,
            ),
            cron_setting="daily",
            disabled=True)
        ```

        :param str resource_name: The name of the resource.
        :param TagRetentionRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagRetentionRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_setting: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 retention_rule: Optional[pulumi.Input[pulumi.InputType['TagRetentionRuleRetentionRuleArgs']]] = None,
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
            __props__ = TagRetentionRuleArgs.__new__(TagRetentionRuleArgs)

            if cron_setting is None and not opts.urn:
                raise TypeError("Missing required property 'cron_setting'")
            __props__.__dict__["cron_setting"] = cron_setting
            __props__.__dict__["disabled"] = disabled
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            if registry_id is None and not opts.urn:
                raise TypeError("Missing required property 'registry_id'")
            __props__.__dict__["registry_id"] = registry_id
            if retention_rule is None and not opts.urn:
                raise TypeError("Missing required property 'retention_rule'")
            __props__.__dict__["retention_rule"] = retention_rule
            __props__.__dict__["retention_id"] = None
        super(TagRetentionRule, __self__).__init__(
            'tencentcloud:Tcr/tagRetentionRule:TagRetentionRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_setting: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            registry_id: Optional[pulumi.Input[str]] = None,
            retention_id: Optional[pulumi.Input[int]] = None,
            retention_rule: Optional[pulumi.Input[pulumi.InputType['TagRetentionRuleRetentionRuleArgs']]] = None) -> 'TagRetentionRule':
        """
        Get an existing TagRetentionRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_setting: Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        :param pulumi.Input[bool] disabled: Whether to disable the rule, with the default value of false.
        :param pulumi.Input[str] namespace_name: The Name of the namespace.
        :param pulumi.Input[str] registry_id: The main instance ID.
        :param pulumi.Input[int] retention_id: The ID of the retention task.
        :param pulumi.Input[pulumi.InputType['TagRetentionRuleRetentionRuleArgs']] retention_rule: Retention Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagRetentionRuleState.__new__(_TagRetentionRuleState)

        __props__.__dict__["cron_setting"] = cron_setting
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["retention_id"] = retention_id
        __props__.__dict__["retention_rule"] = retention_rule
        return TagRetentionRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronSetting")
    def cron_setting(self) -> pulumi.Output[str]:
        """
        Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        """
        return pulumi.get(self, "cron_setting")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to disable the rule, with the default value of false.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The Name of the namespace.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[str]:
        """
        The main instance ID.
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="retentionId")
    def retention_id(self) -> pulumi.Output[int]:
        """
        The ID of the retention task.
        """
        return pulumi.get(self, "retention_id")

    @property
    @pulumi.getter(name="retentionRule")
    def retention_rule(self) -> pulumi.Output['outputs.TagRetentionRuleRetentionRule']:
        """
        Retention Policy.
        """
        return pulumi.get(self, "retention_rule")

