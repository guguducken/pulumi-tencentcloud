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

__all__ = ['FileArgs', 'File']

@pulumi.input_type
class FileArgs:
    def __init__(__self__, *,
                 file_id: pulumi.Input[str],
                 kind: pulumi.Input[int],
                 project_id: pulumi.Input[str],
                 size: pulumi.Input[int],
                 type: pulumi.Input[str],
                 file_infos: Optional[pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]]] = None,
                 head_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_columns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_in_file: Optional[pulumi.Input[bool]] = None,
                 line_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tail_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a File resource.
        :param pulumi.Input[str] file_id: File id.
        :param pulumi.Input[int] kind: File kind, parameter file-1, protocol file-2, request file-3.
        :param pulumi.Input[str] project_id: Project id.
        :param pulumi.Input[int] size: File size.
        :param pulumi.Input[str] type: File type, folder-folder.
        :param pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]] file_infos: Files in a folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] head_lines: The first few lines of data.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] header_columns: Meter head.
        :param pulumi.Input[bool] header_in_file: Whether the header is in the file.
        :param pulumi.Input[int] line_count: Line count.
        :param pulumi.Input[str] name: File name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tail_lines: The last few lines of data.
        """
        pulumi.set(__self__, "file_id", file_id)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "type", type)
        if file_infos is not None:
            pulumi.set(__self__, "file_infos", file_infos)
        if head_lines is not None:
            pulumi.set(__self__, "head_lines", head_lines)
        if header_columns is not None:
            pulumi.set(__self__, "header_columns", header_columns)
        if header_in_file is not None:
            pulumi.set(__self__, "header_in_file", header_in_file)
        if line_count is not None:
            pulumi.set(__self__, "line_count", line_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tail_lines is not None:
            pulumi.set(__self__, "tail_lines", tail_lines)

    @property
    @pulumi.getter(name="fileId")
    def file_id(self) -> pulumi.Input[str]:
        """
        File id.
        """
        return pulumi.get(self, "file_id")

    @file_id.setter
    def file_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_id", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[int]:
        """
        File kind, parameter file-1, protocol file-2, request file-3.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[int]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        Project id.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        File size.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        File type, folder-folder.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="fileInfos")
    def file_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]]]:
        """
        Files in a folder.
        """
        return pulumi.get(self, "file_infos")

    @file_infos.setter
    def file_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]]]):
        pulumi.set(self, "file_infos", value)

    @property
    @pulumi.getter(name="headLines")
    def head_lines(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The first few lines of data.
        """
        return pulumi.get(self, "head_lines")

    @head_lines.setter
    def head_lines(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "head_lines", value)

    @property
    @pulumi.getter(name="headerColumns")
    def header_columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Meter head.
        """
        return pulumi.get(self, "header_columns")

    @header_columns.setter
    def header_columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "header_columns", value)

    @property
    @pulumi.getter(name="headerInFile")
    def header_in_file(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the header is in the file.
        """
        return pulumi.get(self, "header_in_file")

    @header_in_file.setter
    def header_in_file(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "header_in_file", value)

    @property
    @pulumi.getter(name="lineCount")
    def line_count(self) -> Optional[pulumi.Input[int]]:
        """
        Line count.
        """
        return pulumi.get(self, "line_count")

    @line_count.setter
    def line_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "line_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        File name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tailLines")
    def tail_lines(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The last few lines of data.
        """
        return pulumi.get(self, "tail_lines")

    @tail_lines.setter
    def tail_lines(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tail_lines", value)


@pulumi.input_type
class _FileState:
    def __init__(__self__, *,
                 file_id: Optional[pulumi.Input[str]] = None,
                 file_infos: Optional[pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]]] = None,
                 head_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_columns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_in_file: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[int]] = None,
                 line_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 tail_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering File resources.
        :param pulumi.Input[str] file_id: File id.
        :param pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]] file_infos: Files in a folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] head_lines: The first few lines of data.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] header_columns: Meter head.
        :param pulumi.Input[bool] header_in_file: Whether the header is in the file.
        :param pulumi.Input[int] kind: File kind, parameter file-1, protocol file-2, request file-3.
        :param pulumi.Input[int] line_count: Line count.
        :param pulumi.Input[str] name: File name.
        :param pulumi.Input[str] project_id: Project id.
        :param pulumi.Input[int] size: File size.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tail_lines: The last few lines of data.
        :param pulumi.Input[str] type: File type, folder-folder.
        """
        if file_id is not None:
            pulumi.set(__self__, "file_id", file_id)
        if file_infos is not None:
            pulumi.set(__self__, "file_infos", file_infos)
        if head_lines is not None:
            pulumi.set(__self__, "head_lines", head_lines)
        if header_columns is not None:
            pulumi.set(__self__, "header_columns", header_columns)
        if header_in_file is not None:
            pulumi.set(__self__, "header_in_file", header_in_file)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if line_count is not None:
            pulumi.set(__self__, "line_count", line_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if tail_lines is not None:
            pulumi.set(__self__, "tail_lines", tail_lines)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="fileId")
    def file_id(self) -> Optional[pulumi.Input[str]]:
        """
        File id.
        """
        return pulumi.get(self, "file_id")

    @file_id.setter
    def file_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_id", value)

    @property
    @pulumi.getter(name="fileInfos")
    def file_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]]]:
        """
        Files in a folder.
        """
        return pulumi.get(self, "file_infos")

    @file_infos.setter
    def file_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FileFileInfoArgs']]]]):
        pulumi.set(self, "file_infos", value)

    @property
    @pulumi.getter(name="headLines")
    def head_lines(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The first few lines of data.
        """
        return pulumi.get(self, "head_lines")

    @head_lines.setter
    def head_lines(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "head_lines", value)

    @property
    @pulumi.getter(name="headerColumns")
    def header_columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Meter head.
        """
        return pulumi.get(self, "header_columns")

    @header_columns.setter
    def header_columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "header_columns", value)

    @property
    @pulumi.getter(name="headerInFile")
    def header_in_file(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the header is in the file.
        """
        return pulumi.get(self, "header_in_file")

    @header_in_file.setter
    def header_in_file(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "header_in_file", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[int]]:
        """
        File kind, parameter file-1, protocol file-2, request file-3.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="lineCount")
    def line_count(self) -> Optional[pulumi.Input[int]]:
        """
        Line count.
        """
        return pulumi.get(self, "line_count")

    @line_count.setter
    def line_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "line_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        File name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Project id.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        File size.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="tailLines")
    def tail_lines(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The last few lines of data.
        """
        return pulumi.get(self, "tail_lines")

    @tail_lines.setter
    def tail_lines(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tail_lines", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        File type, folder-folder.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class File(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_id: Optional[pulumi.Input[str]] = None,
                 file_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileFileInfoArgs']]]]] = None,
                 head_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_columns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_in_file: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[int]] = None,
                 line_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 tail_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a pts file

        > **NOTE:** Modification is not currently supported, please go to the console to modify.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        file = tencentcloud.pts.File("file",
            file_id="file-de2dbaf8",
            header_in_file=False,
            kind=3,
            line_count=0,
            project_id="project-45vw7v82",
            size=10799,
            type="text/plain")
        ```

        ## Import

        pts file can be imported using the project_id#file_id, e.g.

        ```sh
         $ pulumi import tencentcloud:Pts/file:File file project-45vw7v82#file-de2dbaf8
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_id: File id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileFileInfoArgs']]]] file_infos: Files in a folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] head_lines: The first few lines of data.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] header_columns: Meter head.
        :param pulumi.Input[bool] header_in_file: Whether the header is in the file.
        :param pulumi.Input[int] kind: File kind, parameter file-1, protocol file-2, request file-3.
        :param pulumi.Input[int] line_count: Line count.
        :param pulumi.Input[str] name: File name.
        :param pulumi.Input[str] project_id: Project id.
        :param pulumi.Input[int] size: File size.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tail_lines: The last few lines of data.
        :param pulumi.Input[str] type: File type, folder-folder.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a pts file

        > **NOTE:** Modification is not currently supported, please go to the console to modify.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        file = tencentcloud.pts.File("file",
            file_id="file-de2dbaf8",
            header_in_file=False,
            kind=3,
            line_count=0,
            project_id="project-45vw7v82",
            size=10799,
            type="text/plain")
        ```

        ## Import

        pts file can be imported using the project_id#file_id, e.g.

        ```sh
         $ pulumi import tencentcloud:Pts/file:File file project-45vw7v82#file-de2dbaf8
        ```

        :param str resource_name: The name of the resource.
        :param FileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_id: Optional[pulumi.Input[str]] = None,
                 file_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileFileInfoArgs']]]]] = None,
                 head_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_columns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 header_in_file: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[int]] = None,
                 line_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 tail_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = FileArgs.__new__(FileArgs)

            if file_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_id'")
            __props__.__dict__["file_id"] = file_id
            __props__.__dict__["file_infos"] = file_infos
            __props__.__dict__["head_lines"] = head_lines
            __props__.__dict__["header_columns"] = header_columns
            __props__.__dict__["header_in_file"] = header_in_file
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = kind
            __props__.__dict__["line_count"] = line_count
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            __props__.__dict__["tail_lines"] = tail_lines
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(File, __self__).__init__(
            'tencentcloud:Pts/file:File',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            file_id: Optional[pulumi.Input[str]] = None,
            file_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileFileInfoArgs']]]]] = None,
            head_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            header_columns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            header_in_file: Optional[pulumi.Input[bool]] = None,
            kind: Optional[pulumi.Input[int]] = None,
            line_count: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            tail_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'File':
        """
        Get an existing File resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_id: File id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileFileInfoArgs']]]] file_infos: Files in a folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] head_lines: The first few lines of data.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] header_columns: Meter head.
        :param pulumi.Input[bool] header_in_file: Whether the header is in the file.
        :param pulumi.Input[int] kind: File kind, parameter file-1, protocol file-2, request file-3.
        :param pulumi.Input[int] line_count: Line count.
        :param pulumi.Input[str] name: File name.
        :param pulumi.Input[str] project_id: Project id.
        :param pulumi.Input[int] size: File size.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tail_lines: The last few lines of data.
        :param pulumi.Input[str] type: File type, folder-folder.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileState.__new__(_FileState)

        __props__.__dict__["file_id"] = file_id
        __props__.__dict__["file_infos"] = file_infos
        __props__.__dict__["head_lines"] = head_lines
        __props__.__dict__["header_columns"] = header_columns
        __props__.__dict__["header_in_file"] = header_in_file
        __props__.__dict__["kind"] = kind
        __props__.__dict__["line_count"] = line_count
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["size"] = size
        __props__.__dict__["tail_lines"] = tail_lines
        __props__.__dict__["type"] = type
        return File(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fileId")
    def file_id(self) -> pulumi.Output[str]:
        """
        File id.
        """
        return pulumi.get(self, "file_id")

    @property
    @pulumi.getter(name="fileInfos")
    def file_infos(self) -> pulumi.Output[Optional[Sequence['outputs.FileFileInfo']]]:
        """
        Files in a folder.
        """
        return pulumi.get(self, "file_infos")

    @property
    @pulumi.getter(name="headLines")
    def head_lines(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The first few lines of data.
        """
        return pulumi.get(self, "head_lines")

    @property
    @pulumi.getter(name="headerColumns")
    def header_columns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Meter head.
        """
        return pulumi.get(self, "header_columns")

    @property
    @pulumi.getter(name="headerInFile")
    def header_in_file(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the header is in the file.
        """
        return pulumi.get(self, "header_in_file")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[int]:
        """
        File kind, parameter file-1, protocol file-2, request file-3.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lineCount")
    def line_count(self) -> pulumi.Output[Optional[int]]:
        """
        Line count.
        """
        return pulumi.get(self, "line_count")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        File name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        Project id.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        File size.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="tailLines")
    def tail_lines(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The last few lines of data.
        """
        return pulumi.get(self, "tail_lines")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        File type, folder-folder.
        """
        return pulumi.get(self, "type")

