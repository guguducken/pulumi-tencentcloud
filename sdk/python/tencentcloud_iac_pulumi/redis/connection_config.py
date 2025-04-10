# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConnectionConfigArgs', 'ConnectionConfig']

@pulumi.input_type
class ConnectionConfigArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 add_bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 client_limit: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ConnectionConfig resource.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[int] add_bandwidth: Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
               increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
               bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
               the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
               of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
               shards in the standard architecture = 1.
        :param pulumi.Input[int] bandwidth: Additional bandwidth, greater than 0, in MB.
        :param pulumi.Input[int] client_limit: The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
               limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
               number of read replicas +3).
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if add_bandwidth is not None:
            pulumi.set(__self__, "add_bandwidth", add_bandwidth)
        if bandwidth is not None:
            warnings.warn("""Configure `add_bandwidth` instead. This attribute will be removed in the next major version of the provider""", DeprecationWarning)
            pulumi.log.warn("""bandwidth is deprecated: Configure `add_bandwidth` instead. This attribute will be removed in the next major version of the provider""")
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if client_limit is not None:
            pulumi.set(__self__, "client_limit", client_limit)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="addBandwidth")
    def add_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
        increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
        bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
        the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
        of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
        shards in the standard architecture = 1.
        """
        return pulumi.get(self, "add_bandwidth")

    @add_bandwidth.setter
    def add_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "add_bandwidth", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Additional bandwidth, greater than 0, in MB.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="clientLimit")
    def client_limit(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
        limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
        number of read replicas +3).
        """
        return pulumi.get(self, "client_limit")

    @client_limit.setter
    def client_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "client_limit", value)


@pulumi.input_type
class _ConnectionConfigState:
    def __init__(__self__, *,
                 add_bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 base_bandwidth: Optional[pulumi.Input[int]] = None,
                 client_limit: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_add_bandwidth: Optional[pulumi.Input[int]] = None,
                 min_add_bandwidth: Optional[pulumi.Input[int]] = None,
                 total_bandwidth: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ConnectionConfig resources.
        :param pulumi.Input[int] add_bandwidth: Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
               increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
               bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
               the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
               of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
               shards in the standard architecture = 1.
        :param pulumi.Input[int] bandwidth: Additional bandwidth, greater than 0, in MB.
        :param pulumi.Input[int] base_bandwidth: standard bandwidth. Refers to the bandwidth allocated by the system to each node when an instance is purchased.
        :param pulumi.Input[int] client_limit: The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
               limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
               number of read replicas +3).
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[int] max_add_bandwidth: Additional bandwidth is capped.
        :param pulumi.Input[int] min_add_bandwidth: Additional bandwidth sets the lower limit.
        :param pulumi.Input[int] total_bandwidth: Total bandwidth of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards *
               (number of primary nodes + number of read-only replica nodes), the number of shards of the standard architecture = 1, in
               Mb/s.
        """
        if add_bandwidth is not None:
            pulumi.set(__self__, "add_bandwidth", add_bandwidth)
        if bandwidth is not None:
            warnings.warn("""Configure `add_bandwidth` instead. This attribute will be removed in the next major version of the provider""", DeprecationWarning)
            pulumi.log.warn("""bandwidth is deprecated: Configure `add_bandwidth` instead. This attribute will be removed in the next major version of the provider""")
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if base_bandwidth is not None:
            pulumi.set(__self__, "base_bandwidth", base_bandwidth)
        if client_limit is not None:
            pulumi.set(__self__, "client_limit", client_limit)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_add_bandwidth is not None:
            pulumi.set(__self__, "max_add_bandwidth", max_add_bandwidth)
        if min_add_bandwidth is not None:
            pulumi.set(__self__, "min_add_bandwidth", min_add_bandwidth)
        if total_bandwidth is not None:
            pulumi.set(__self__, "total_bandwidth", total_bandwidth)

    @property
    @pulumi.getter(name="addBandwidth")
    def add_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
        increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
        bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
        the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
        of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
        shards in the standard architecture = 1.
        """
        return pulumi.get(self, "add_bandwidth")

    @add_bandwidth.setter
    def add_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "add_bandwidth", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Additional bandwidth, greater than 0, in MB.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="baseBandwidth")
    def base_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        standard bandwidth. Refers to the bandwidth allocated by the system to each node when an instance is purchased.
        """
        return pulumi.get(self, "base_bandwidth")

    @base_bandwidth.setter
    def base_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "base_bandwidth", value)

    @property
    @pulumi.getter(name="clientLimit")
    def client_limit(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
        limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
        number of read replicas +3).
        """
        return pulumi.get(self, "client_limit")

    @client_limit.setter
    def client_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "client_limit", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxAddBandwidth")
    def max_add_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Additional bandwidth is capped.
        """
        return pulumi.get(self, "max_add_bandwidth")

    @max_add_bandwidth.setter
    def max_add_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_add_bandwidth", value)

    @property
    @pulumi.getter(name="minAddBandwidth")
    def min_add_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Additional bandwidth sets the lower limit.
        """
        return pulumi.get(self, "min_add_bandwidth")

    @min_add_bandwidth.setter
    def min_add_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_add_bandwidth", value)

    @property
    @pulumi.getter(name="totalBandwidth")
    def total_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Total bandwidth of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards *
        (number of primary nodes + number of read-only replica nodes), the number of shards of the standard architecture = 1, in
        Mb/s.
        """
        return pulumi.get(self, "total_bandwidth")

    @total_bandwidth.setter
    def total_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "total_bandwidth", value)


class ConnectionConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 client_limit: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ConnectionConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] add_bandwidth: Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
               increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
               bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
               the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
               of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
               shards in the standard architecture = 1.
        :param pulumi.Input[int] bandwidth: Additional bandwidth, greater than 0, in MB.
        :param pulumi.Input[int] client_limit: The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
               limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
               number of read replicas +3).
        :param pulumi.Input[str] instance_id: The ID of instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ConnectionConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ConnectionConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 client_limit: Optional[pulumi.Input[int]] = None,
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
            __props__ = ConnectionConfigArgs.__new__(ConnectionConfigArgs)

            __props__.__dict__["add_bandwidth"] = add_bandwidth
            if bandwidth is not None and not opts.urn:
                warnings.warn("""Configure `add_bandwidth` instead. This attribute will be removed in the next major version of the provider""", DeprecationWarning)
                pulumi.log.warn("""bandwidth is deprecated: Configure `add_bandwidth` instead. This attribute will be removed in the next major version of the provider""")
            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["client_limit"] = client_limit
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["base_bandwidth"] = None
            __props__.__dict__["max_add_bandwidth"] = None
            __props__.__dict__["min_add_bandwidth"] = None
            __props__.__dict__["total_bandwidth"] = None
        super(ConnectionConfig, __self__).__init__(
            'tencentcloud:Redis/connectionConfig:ConnectionConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_bandwidth: Optional[pulumi.Input[int]] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            base_bandwidth: Optional[pulumi.Input[int]] = None,
            client_limit: Optional[pulumi.Input[int]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            max_add_bandwidth: Optional[pulumi.Input[int]] = None,
            min_add_bandwidth: Optional[pulumi.Input[int]] = None,
            total_bandwidth: Optional[pulumi.Input[int]] = None) -> 'ConnectionConfig':
        """
        Get an existing ConnectionConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] add_bandwidth: Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
               increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
               bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
               the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
               of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
               shards in the standard architecture = 1.
        :param pulumi.Input[int] bandwidth: Additional bandwidth, greater than 0, in MB.
        :param pulumi.Input[int] base_bandwidth: standard bandwidth. Refers to the bandwidth allocated by the system to each node when an instance is purchased.
        :param pulumi.Input[int] client_limit: The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
               limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
               number of read replicas +3).
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[int] max_add_bandwidth: Additional bandwidth is capped.
        :param pulumi.Input[int] min_add_bandwidth: Additional bandwidth sets the lower limit.
        :param pulumi.Input[int] total_bandwidth: Total bandwidth of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards *
               (number of primary nodes + number of read-only replica nodes), the number of shards of the standard architecture = 1, in
               Mb/s.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConnectionConfigState.__new__(_ConnectionConfigState)

        __props__.__dict__["add_bandwidth"] = add_bandwidth
        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["base_bandwidth"] = base_bandwidth
        __props__.__dict__["client_limit"] = client_limit
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_add_bandwidth"] = max_add_bandwidth
        __props__.__dict__["min_add_bandwidth"] = min_add_bandwidth
        __props__.__dict__["total_bandwidth"] = total_bandwidth
        return ConnectionConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addBandwidth")
    def add_bandwidth(self) -> pulumi.Output[int]:
        """
        Refers to the additional bandwidth of the instance. When the standard bandwidth does not meet the demand, the user can
        increase the bandwidth by himself. When the read-only copy is enabled, the total bandwidth of the instance = additional
        bandwidth * number of fragments + standard bandwidth * number of fragments * Max ([number of read-only replicas, 1] ),
        the number of shards in the standard architecture = 1, and when read-only replicas are not enabled, the total bandwidth
        of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards, and the number of
        shards in the standard architecture = 1.
        """
        return pulumi.get(self, "add_bandwidth")

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        Additional bandwidth, greater than 0, in MB.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="baseBandwidth")
    def base_bandwidth(self) -> pulumi.Output[int]:
        """
        standard bandwidth. Refers to the bandwidth allocated by the system to each node when an instance is purchased.
        """
        return pulumi.get(self, "base_bandwidth")

    @property
    @pulumi.getter(name="clientLimit")
    def client_limit(self) -> pulumi.Output[Optional[int]]:
        """
        The total number of connections per shard.If read-only replicas are not enabled, the lower limit is 10,000 and the upper
        limit is 40,000.When you enable read-only replicas, the minimum limit is 10,000 and the upper limit is 10,000 × (the
        number of read replicas +3).
        """
        return pulumi.get(self, "client_limit")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxAddBandwidth")
    def max_add_bandwidth(self) -> pulumi.Output[int]:
        """
        Additional bandwidth is capped.
        """
        return pulumi.get(self, "max_add_bandwidth")

    @property
    @pulumi.getter(name="minAddBandwidth")
    def min_add_bandwidth(self) -> pulumi.Output[int]:
        """
        Additional bandwidth sets the lower limit.
        """
        return pulumi.get(self, "min_add_bandwidth")

    @property
    @pulumi.getter(name="totalBandwidth")
    def total_bandwidth(self) -> pulumi.Output[int]:
        """
        Total bandwidth of the instance = additional bandwidth * number of shards + standard bandwidth * number of shards *
        (number of primary nodes + number of read-only replica nodes), the number of shards of the standard architecture = 1, in
        Mb/s.
        """
        return pulumi.get(self, "total_bandwidth")

