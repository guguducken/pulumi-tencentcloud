# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InvocationInvokeAttachmentArgs', 'InvocationInvokeAttachment']

@pulumi.input_type
class InvocationInvokeAttachmentArgs:
    def __init__(__self__, *,
                 command_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 output_cos_bucket_url: Optional[pulumi.Input[str]] = None,
                 output_cos_key_prefix: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 working_directory: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InvocationInvokeAttachment resource.
        :param pulumi.Input[str] command_id: Command ID.
        :param pulumi.Input[str] instance_id: ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        :param pulumi.Input[str] output_cos_bucket_url: The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        :param pulumi.Input[str] output_cos_key_prefix: The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        :param pulumi.Input[str] parameters: Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        :param pulumi.Input[int] timeout: Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        :param pulumi.Input[str] username: The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        :param pulumi.Input[str] working_directory: Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        pulumi.set(__self__, "command_id", command_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if output_cos_bucket_url is not None:
            pulumi.set(__self__, "output_cos_bucket_url", output_cos_bucket_url)
        if output_cos_key_prefix is not None:
            pulumi.set(__self__, "output_cos_key_prefix", output_cos_key_prefix)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if working_directory is not None:
            pulumi.set(__self__, "working_directory", working_directory)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> pulumi.Input[str]:
        """
        Command ID.
        """
        return pulumi.get(self, "command_id")

    @command_id.setter
    def command_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "command_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="outputCosBucketUrl")
    def output_cos_bucket_url(self) -> Optional[pulumi.Input[str]]:
        """
        The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        """
        return pulumi.get(self, "output_cos_bucket_url")

    @output_cos_bucket_url.setter
    def output_cos_bucket_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_cos_bucket_url", value)

    @property
    @pulumi.getter(name="outputCosKeyPrefix")
    def output_cos_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        """
        return pulumi.get(self, "output_cos_key_prefix")

    @output_cos_key_prefix.setter
    def output_cos_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_cos_key_prefix", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[str]]:
        """
        Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="workingDirectory")
    def working_directory(self) -> Optional[pulumi.Input[str]]:
        """
        Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        return pulumi.get(self, "working_directory")

    @working_directory.setter
    def working_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "working_directory", value)


@pulumi.input_type
class _InvocationInvokeAttachmentState:
    def __init__(__self__, *,
                 command_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 output_cos_bucket_url: Optional[pulumi.Input[str]] = None,
                 output_cos_key_prefix: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 working_directory: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InvocationInvokeAttachment resources.
        :param pulumi.Input[str] command_id: Command ID.
        :param pulumi.Input[str] instance_id: ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        :param pulumi.Input[str] output_cos_bucket_url: The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        :param pulumi.Input[str] output_cos_key_prefix: The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        :param pulumi.Input[str] parameters: Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        :param pulumi.Input[int] timeout: Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        :param pulumi.Input[str] username: The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        :param pulumi.Input[str] working_directory: Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        if command_id is not None:
            pulumi.set(__self__, "command_id", command_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if output_cos_bucket_url is not None:
            pulumi.set(__self__, "output_cos_bucket_url", output_cos_bucket_url)
        if output_cos_key_prefix is not None:
            pulumi.set(__self__, "output_cos_key_prefix", output_cos_key_prefix)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if working_directory is not None:
            pulumi.set(__self__, "working_directory", working_directory)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> Optional[pulumi.Input[str]]:
        """
        Command ID.
        """
        return pulumi.get(self, "command_id")

    @command_id.setter
    def command_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="outputCosBucketUrl")
    def output_cos_bucket_url(self) -> Optional[pulumi.Input[str]]:
        """
        The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        """
        return pulumi.get(self, "output_cos_bucket_url")

    @output_cos_bucket_url.setter
    def output_cos_bucket_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_cos_bucket_url", value)

    @property
    @pulumi.getter(name="outputCosKeyPrefix")
    def output_cos_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        """
        return pulumi.get(self, "output_cos_key_prefix")

    @output_cos_key_prefix.setter
    def output_cos_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_cos_key_prefix", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[str]]:
        """
        Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="workingDirectory")
    def working_directory(self) -> Optional[pulumi.Input[str]]:
        """
        Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        return pulumi.get(self, "working_directory")

    @working_directory.setter
    def working_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "working_directory", value)


class InvocationInvokeAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 output_cos_bucket_url: Optional[pulumi.Input[str]] = None,
                 output_cos_key_prefix: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 working_directory: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tat invocation_invoke_attachment

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        invocation_invoke_attachment = tencentcloud.tat.InvocationInvokeAttachment("invocationInvokeAttachment",
            command_id="cmd-rxbs7f5z",
            instance_id="ins-881b1c8w",
            output_cos_bucket_url="https://BucketName-123454321.cos.ap-beijing.myqcloud.com",
            output_cos_key_prefix="log",
            timeout=100,
            username="root",
            working_directory="/root")
        ```

        ## Import

        tat invocation can be imported using the invocation_id#instance_id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tat/invocationInvokeAttachment:InvocationInvokeAttachment invocation_invoke_attachment inv-mhs6ca8z#ins-881b1c8w
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command_id: Command ID.
        :param pulumi.Input[str] instance_id: ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        :param pulumi.Input[str] output_cos_bucket_url: The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        :param pulumi.Input[str] output_cos_key_prefix: The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        :param pulumi.Input[str] parameters: Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        :param pulumi.Input[int] timeout: Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        :param pulumi.Input[str] username: The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        :param pulumi.Input[str] working_directory: Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InvocationInvokeAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tat invocation_invoke_attachment

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        invocation_invoke_attachment = tencentcloud.tat.InvocationInvokeAttachment("invocationInvokeAttachment",
            command_id="cmd-rxbs7f5z",
            instance_id="ins-881b1c8w",
            output_cos_bucket_url="https://BucketName-123454321.cos.ap-beijing.myqcloud.com",
            output_cos_key_prefix="log",
            timeout=100,
            username="root",
            working_directory="/root")
        ```

        ## Import

        tat invocation can be imported using the invocation_id#instance_id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tat/invocationInvokeAttachment:InvocationInvokeAttachment invocation_invoke_attachment inv-mhs6ca8z#ins-881b1c8w
        ```

        :param str resource_name: The name of the resource.
        :param InvocationInvokeAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InvocationInvokeAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 output_cos_bucket_url: Optional[pulumi.Input[str]] = None,
                 output_cos_key_prefix: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 working_directory: Optional[pulumi.Input[str]] = None,
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
            __props__ = InvocationInvokeAttachmentArgs.__new__(InvocationInvokeAttachmentArgs)

            if command_id is None and not opts.urn:
                raise TypeError("Missing required property 'command_id'")
            __props__.__dict__["command_id"] = command_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["output_cos_bucket_url"] = output_cos_bucket_url
            __props__.__dict__["output_cos_key_prefix"] = output_cos_key_prefix
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["username"] = username
            __props__.__dict__["working_directory"] = working_directory
        super(InvocationInvokeAttachment, __self__).__init__(
            'tencentcloud:Tat/invocationInvokeAttachment:InvocationInvokeAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            output_cos_bucket_url: Optional[pulumi.Input[str]] = None,
            output_cos_key_prefix: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            username: Optional[pulumi.Input[str]] = None,
            working_directory: Optional[pulumi.Input[str]] = None) -> 'InvocationInvokeAttachment':
        """
        Get an existing InvocationInvokeAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command_id: Command ID.
        :param pulumi.Input[str] instance_id: ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        :param pulumi.Input[str] output_cos_bucket_url: The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        :param pulumi.Input[str] output_cos_key_prefix: The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        :param pulumi.Input[str] parameters: Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        :param pulumi.Input[int] timeout: Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        :param pulumi.Input[str] username: The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        :param pulumi.Input[str] working_directory: Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InvocationInvokeAttachmentState.__new__(_InvocationInvokeAttachmentState)

        __props__.__dict__["command_id"] = command_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["output_cos_bucket_url"] = output_cos_bucket_url
        __props__.__dict__["output_cos_key_prefix"] = output_cos_key_prefix
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["username"] = username
        __props__.__dict__["working_directory"] = working_directory
        return InvocationInvokeAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> pulumi.Output[str]:
        """
        Command ID.
        """
        return pulumi.get(self, "command_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of instances about to execute commands. Supported instance types:  CVM  LIGHTHOUSE.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="outputCosBucketUrl")
    def output_cos_bucket_url(self) -> pulumi.Output[Optional[str]]:
        """
        The COS bucket URL for uploading logs. The URL must start with https, such as https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        """
        return pulumi.get(self, "output_cos_bucket_url")

    @property
    @pulumi.getter(name="outputCosKeyPrefix")
    def output_cos_key_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive slashes.
        """
        return pulumi.get(self, "output_cos_key_prefix")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[str]]:
        """
        Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root is used to execute commands on Linux and the user System is used on Windows.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="workingDirectory")
    def working_directory(self) -> pulumi.Output[Optional[str]]:
        """
        Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for POWERSHELL commands.
        """
        return pulumi.get(self, "working_directory")

