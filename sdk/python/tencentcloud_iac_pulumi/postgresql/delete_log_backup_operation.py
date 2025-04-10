# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeleteLogBackupOperationArgs', 'DeleteLogBackupOperation']

@pulumi.input_type
class DeleteLogBackupOperationArgs:
    def __init__(__self__, *,
                 db_instance_id: pulumi.Input[str],
                 log_backup_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DeleteLogBackupOperation resource.
        :param pulumi.Input[str] db_instance_id: Instance ID.
        :param pulumi.Input[str] log_backup_id: Log backup ID.
        """
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "log_backup_id", log_backup_id)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="logBackupId")
    def log_backup_id(self) -> pulumi.Input[str]:
        """
        Log backup ID.
        """
        return pulumi.get(self, "log_backup_id")

    @log_backup_id.setter
    def log_backup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_backup_id", value)


@pulumi.input_type
class _DeleteLogBackupOperationState:
    def __init__(__self__, *,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 log_backup_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DeleteLogBackupOperation resources.
        :param pulumi.Input[str] db_instance_id: Instance ID.
        :param pulumi.Input[str] log_backup_id: Log backup ID.
        """
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if log_backup_id is not None:
            pulumi.set(__self__, "log_backup_id", log_backup_id)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="logBackupId")
    def log_backup_id(self) -> Optional[pulumi.Input[str]]:
        """
        Log backup ID.
        """
        return pulumi.get(self, "log_backup_id")

    @log_backup_id.setter
    def log_backup_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_backup_id", value)


class DeleteLogBackupOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 log_backup_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a postgresql delete_log_backup_operation

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        delete_log_backup_operation = tencentcloud.postgresql.DeleteLogBackupOperation("deleteLogBackupOperation",
            db_instance_id="local.pg_id",
            log_backup_id="local.pg_log_backup_id")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: Instance ID.
        :param pulumi.Input[str] log_backup_id: Log backup ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeleteLogBackupOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a postgresql delete_log_backup_operation

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        delete_log_backup_operation = tencentcloud.postgresql.DeleteLogBackupOperation("deleteLogBackupOperation",
            db_instance_id="local.pg_id",
            log_backup_id="local.pg_log_backup_id")
        ```

        :param str resource_name: The name of the resource.
        :param DeleteLogBackupOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeleteLogBackupOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 log_backup_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DeleteLogBackupOperationArgs.__new__(DeleteLogBackupOperationArgs)

            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            if log_backup_id is None and not opts.urn:
                raise TypeError("Missing required property 'log_backup_id'")
            __props__.__dict__["log_backup_id"] = log_backup_id
        super(DeleteLogBackupOperation, __self__).__init__(
            'tencentcloud:Postgresql/deleteLogBackupOperation:DeleteLogBackupOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            log_backup_id: Optional[pulumi.Input[str]] = None) -> 'DeleteLogBackupOperation':
        """
        Get an existing DeleteLogBackupOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: Instance ID.
        :param pulumi.Input[str] log_backup_id: Log backup ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeleteLogBackupOperationState.__new__(_DeleteLogBackupOperationState)

        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["log_backup_id"] = log_backup_id
        return DeleteLogBackupOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="logBackupId")
    def log_backup_id(self) -> pulumi.Output[str]:
        """
        Log backup ID.
        """
        return pulumi.get(self, "log_backup_id")

