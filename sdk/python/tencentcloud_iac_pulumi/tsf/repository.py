# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 bucket_region: pulumi.Input[str],
                 repository_name: pulumi.Input[str],
                 repository_type: pulumi.Input[str],
                 directory: Optional[pulumi.Input[str]] = None,
                 repository_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[str] bucket_name: the name of the bucket where the warehouse is located.
        :param pulumi.Input[str] bucket_region: Bucket region where the warehouse is located.
        :param pulumi.Input[str] repository_name: warehouse name.
        :param pulumi.Input[str] repository_type: warehouse type (default warehouse: default, private warehouse: private).
        :param pulumi.Input[str] directory: directory.
        :param pulumi.Input[str] repository_desc: warehouse description.
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "bucket_region", bucket_region)
        pulumi.set(__self__, "repository_name", repository_name)
        pulumi.set(__self__, "repository_type", repository_type)
        if directory is not None:
            pulumi.set(__self__, "directory", directory)
        if repository_desc is not None:
            pulumi.set(__self__, "repository_desc", repository_desc)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        the name of the bucket where the warehouse is located.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="bucketRegion")
    def bucket_region(self) -> pulumi.Input[str]:
        """
        Bucket region where the warehouse is located.
        """
        return pulumi.get(self, "bucket_region")

    @bucket_region.setter
    def bucket_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_region", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Input[str]:
        """
        warehouse name.
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_name", value)

    @property
    @pulumi.getter(name="repositoryType")
    def repository_type(self) -> pulumi.Input[str]:
        """
        warehouse type (default warehouse: default, private warehouse: private).
        """
        return pulumi.get(self, "repository_type")

    @repository_type.setter
    def repository_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_type", value)

    @property
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[str]]:
        """
        directory.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory", value)

    @property
    @pulumi.getter(name="repositoryDesc")
    def repository_desc(self) -> Optional[pulumi.Input[str]]:
        """
        warehouse description.
        """
        return pulumi.get(self, "repository_desc")

    @repository_desc.setter
    def repository_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_desc", value)


@pulumi.input_type
class _RepositoryState:
    def __init__(__self__, *,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 bucket_region: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 is_used: Optional[pulumi.Input[bool]] = None,
                 repository_desc: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Repository resources.
        :param pulumi.Input[str] bucket_name: the name of the bucket where the warehouse is located.
        :param pulumi.Input[str] bucket_region: Bucket region where the warehouse is located.
        :param pulumi.Input[str] create_time: warehouse creation time.
        :param pulumi.Input[str] directory: directory.
        :param pulumi.Input[bool] is_used: whether the repository is in use.
        :param pulumi.Input[str] repository_desc: warehouse description.
        :param pulumi.Input[str] repository_id: Warehouse ID.
        :param pulumi.Input[str] repository_name: warehouse name.
        :param pulumi.Input[str] repository_type: warehouse type (default warehouse: default, private warehouse: private).
        """
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if bucket_region is not None:
            pulumi.set(__self__, "bucket_region", bucket_region)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if directory is not None:
            pulumi.set(__self__, "directory", directory)
        if is_used is not None:
            pulumi.set(__self__, "is_used", is_used)
        if repository_desc is not None:
            pulumi.set(__self__, "repository_desc", repository_desc)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if repository_name is not None:
            pulumi.set(__self__, "repository_name", repository_name)
        if repository_type is not None:
            pulumi.set(__self__, "repository_type", repository_type)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        the name of the bucket where the warehouse is located.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="bucketRegion")
    def bucket_region(self) -> Optional[pulumi.Input[str]]:
        """
        Bucket region where the warehouse is located.
        """
        return pulumi.get(self, "bucket_region")

    @bucket_region.setter
    def bucket_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_region", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        warehouse creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[str]]:
        """
        directory.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory", value)

    @property
    @pulumi.getter(name="isUsed")
    def is_used(self) -> Optional[pulumi.Input[bool]]:
        """
        whether the repository is in use.
        """
        return pulumi.get(self, "is_used")

    @is_used.setter
    def is_used(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_used", value)

    @property
    @pulumi.getter(name="repositoryDesc")
    def repository_desc(self) -> Optional[pulumi.Input[str]]:
        """
        warehouse description.
        """
        return pulumi.get(self, "repository_desc")

    @repository_desc.setter
    def repository_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_desc", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        Warehouse ID.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> Optional[pulumi.Input[str]]:
        """
        warehouse name.
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_name", value)

    @property
    @pulumi.getter(name="repositoryType")
    def repository_type(self) -> Optional[pulumi.Input[str]]:
        """
        warehouse type (default warehouse: default, private warehouse: private).
        """
        return pulumi.get(self, "repository_type")

    @repository_type.setter
    def repository_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_type", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 bucket_region: Optional[pulumi.Input[str]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 repository_desc: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tsf repository

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        repository = tencentcloud.tsf.Repository("repository",
            bucket_name="",
            bucket_region="",
            directory="",
            repository_desc="",
            repository_name="",
            repository_type="")
        ```

        ## Import

        tsf repository can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tsf/repository:Repository repository repository_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: the name of the bucket where the warehouse is located.
        :param pulumi.Input[str] bucket_region: Bucket region where the warehouse is located.
        :param pulumi.Input[str] directory: directory.
        :param pulumi.Input[str] repository_desc: warehouse description.
        :param pulumi.Input[str] repository_name: warehouse name.
        :param pulumi.Input[str] repository_type: warehouse type (default warehouse: default, private warehouse: private).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tsf repository

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        repository = tencentcloud.tsf.Repository("repository",
            bucket_name="",
            bucket_region="",
            directory="",
            repository_desc="",
            repository_name="",
            repository_type="")
        ```

        ## Import

        tsf repository can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tsf/repository:Repository repository repository_id
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 bucket_region: Optional[pulumi.Input[str]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 repository_desc: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            if bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_name'")
            __props__.__dict__["bucket_name"] = bucket_name
            if bucket_region is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_region'")
            __props__.__dict__["bucket_region"] = bucket_region
            __props__.__dict__["directory"] = directory
            __props__.__dict__["repository_desc"] = repository_desc
            if repository_name is None and not opts.urn:
                raise TypeError("Missing required property 'repository_name'")
            __props__.__dict__["repository_name"] = repository_name
            if repository_type is None and not opts.urn:
                raise TypeError("Missing required property 'repository_type'")
            __props__.__dict__["repository_type"] = repository_type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["is_used"] = None
            __props__.__dict__["repository_id"] = None
        super(Repository, __self__).__init__(
            'tencentcloud:Tsf/repository:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_name: Optional[pulumi.Input[str]] = None,
            bucket_region: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            directory: Optional[pulumi.Input[str]] = None,
            is_used: Optional[pulumi.Input[bool]] = None,
            repository_desc: Optional[pulumi.Input[str]] = None,
            repository_id: Optional[pulumi.Input[str]] = None,
            repository_name: Optional[pulumi.Input[str]] = None,
            repository_type: Optional[pulumi.Input[str]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: the name of the bucket where the warehouse is located.
        :param pulumi.Input[str] bucket_region: Bucket region where the warehouse is located.
        :param pulumi.Input[str] create_time: warehouse creation time.
        :param pulumi.Input[str] directory: directory.
        :param pulumi.Input[bool] is_used: whether the repository is in use.
        :param pulumi.Input[str] repository_desc: warehouse description.
        :param pulumi.Input[str] repository_id: Warehouse ID.
        :param pulumi.Input[str] repository_name: warehouse name.
        :param pulumi.Input[str] repository_type: warehouse type (default warehouse: default, private warehouse: private).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryState.__new__(_RepositoryState)

        __props__.__dict__["bucket_name"] = bucket_name
        __props__.__dict__["bucket_region"] = bucket_region
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["directory"] = directory
        __props__.__dict__["is_used"] = is_used
        __props__.__dict__["repository_desc"] = repository_desc
        __props__.__dict__["repository_id"] = repository_id
        __props__.__dict__["repository_name"] = repository_name
        __props__.__dict__["repository_type"] = repository_type
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[str]:
        """
        the name of the bucket where the warehouse is located.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="bucketRegion")
    def bucket_region(self) -> pulumi.Output[str]:
        """
        Bucket region where the warehouse is located.
        """
        return pulumi.get(self, "bucket_region")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        warehouse creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Output[str]:
        """
        directory.
        """
        return pulumi.get(self, "directory")

    @property
    @pulumi.getter(name="isUsed")
    def is_used(self) -> pulumi.Output[bool]:
        """
        whether the repository is in use.
        """
        return pulumi.get(self, "is_used")

    @property
    @pulumi.getter(name="repositoryDesc")
    def repository_desc(self) -> pulumi.Output[str]:
        """
        warehouse description.
        """
        return pulumi.get(self, "repository_desc")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[str]:
        """
        Warehouse ID.
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Output[str]:
        """
        warehouse name.
        """
        return pulumi.get(self, "repository_name")

    @property
    @pulumi.getter(name="repositoryType")
    def repository_type(self) -> pulumi.Output[str]:
        """
        warehouse type (default warehouse: default, private warehouse: private).
        """
        return pulumi.get(self, "repository_type")

