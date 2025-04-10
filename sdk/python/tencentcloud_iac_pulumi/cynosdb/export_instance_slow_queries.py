# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ExportInstanceSlowQueriesArgs', 'ExportInstanceSlowQueries']

@pulumi.input_type
class ExportInstanceSlowQueriesArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 database: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ExportInstanceSlowQueries resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] database: Database name.
        :param pulumi.Input[str] end_time: Latest transaction start time.
        :param pulumi.Input[str] file_type: File type, optional values: csv, original.
        :param pulumi.Input[str] host: Client host.
        :param pulumi.Input[str] start_time: Earliest transaction start time.
        :param pulumi.Input[str] username: user name.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if file_type is not None:
            pulumi.set(__self__, "file_type", file_type)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Database name.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        Latest transaction start time.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> Optional[pulumi.Input[str]]:
        """
        File type, optional values: csv, original.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Client host.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Earliest transaction start time.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        user name.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _ExportInstanceSlowQueriesState:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExportInstanceSlowQueries resources.
        :param pulumi.Input[str] database: Database name.
        :param pulumi.Input[str] end_time: Latest transaction start time.
        :param pulumi.Input[str] file_content: Slow query export content.
        :param pulumi.Input[str] file_type: File type, optional values: csv, original.
        :param pulumi.Input[str] host: Client host.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] start_time: Earliest transaction start time.
        :param pulumi.Input[str] username: user name.
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if file_content is not None:
            pulumi.set(__self__, "file_content", file_content)
        if file_type is not None:
            pulumi.set(__self__, "file_type", file_type)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Database name.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        Latest transaction start time.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> Optional[pulumi.Input[str]]:
        """
        Slow query export content.
        """
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_content", value)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> Optional[pulumi.Input[str]]:
        """
        File type, optional values: csv, original.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Client host.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Earliest transaction start time.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        user name.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class ExportInstanceSlowQueries(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a cynosdb export_instance_slow_queries

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        export_instance_slow_queries = tencentcloud.cynosdb.ExportInstanceSlowQueries("exportInstanceSlowQueries",
            database="db1",
            end_time="2022-01-01 14:00:00",
            file_type="csv",
            host="10.10.10.10",
            instance_id="cynosdbmysql-ins-123",
            start_time="2022-01-01 12:00:00",
            username="root")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: Database name.
        :param pulumi.Input[str] end_time: Latest transaction start time.
        :param pulumi.Input[str] file_type: File type, optional values: csv, original.
        :param pulumi.Input[str] host: Client host.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] start_time: Earliest transaction start time.
        :param pulumi.Input[str] username: user name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExportInstanceSlowQueriesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cynosdb export_instance_slow_queries

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        export_instance_slow_queries = tencentcloud.cynosdb.ExportInstanceSlowQueries("exportInstanceSlowQueries",
            database="db1",
            end_time="2022-01-01 14:00:00",
            file_type="csv",
            host="10.10.10.10",
            instance_id="cynosdbmysql-ins-123",
            start_time="2022-01-01 12:00:00",
            username="root")
        ```

        :param str resource_name: The name of the resource.
        :param ExportInstanceSlowQueriesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExportInstanceSlowQueriesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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
            __props__ = ExportInstanceSlowQueriesArgs.__new__(ExportInstanceSlowQueriesArgs)

            __props__.__dict__["database"] = database
            __props__.__dict__["end_time"] = end_time
            __props__.__dict__["file_type"] = file_type
            __props__.__dict__["host"] = host
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["username"] = username
            __props__.__dict__["file_content"] = None
        super(ExportInstanceSlowQueries, __self__).__init__(
            'tencentcloud:Cynosdb/exportInstanceSlowQueries:ExportInstanceSlowQueries',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            file_content: Optional[pulumi.Input[str]] = None,
            file_type: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'ExportInstanceSlowQueries':
        """
        Get an existing ExportInstanceSlowQueries resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: Database name.
        :param pulumi.Input[str] end_time: Latest transaction start time.
        :param pulumi.Input[str] file_content: Slow query export content.
        :param pulumi.Input[str] file_type: File type, optional values: csv, original.
        :param pulumi.Input[str] host: Client host.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] start_time: Earliest transaction start time.
        :param pulumi.Input[str] username: user name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExportInstanceSlowQueriesState.__new__(_ExportInstanceSlowQueriesState)

        __props__.__dict__["database"] = database
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["file_content"] = file_content
        __props__.__dict__["file_type"] = file_type
        __props__.__dict__["host"] = host
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["username"] = username
        return ExportInstanceSlowQueries(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        """
        Database name.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[str]]:
        """
        Latest transaction start time.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Output[str]:
        """
        Slow query export content.
        """
        return pulumi.get(self, "file_content")

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> pulumi.Output[Optional[str]]:
        """
        File type, optional values: csv, original.
        """
        return pulumi.get(self, "file_type")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        Client host.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[str]]:
        """
        Earliest transaction start time.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        user name.
        """
        return pulumi.get(self, "username")

