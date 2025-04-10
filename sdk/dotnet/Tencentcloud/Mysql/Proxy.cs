// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a resource to create a mysql proxy
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxy = new Tencentcloud.Mysql.Proxy("proxy", new Tencentcloud.Mysql.ProxyArgs
    ///         {
    ///             ConnectionPoolLimit = 2,
    ///             Desc = "desc1",
    ///             InstanceId = "cdb-fitq5t9h",
    ///             ProxyNodeCustoms = 
    ///             {
    ///                 new Tencentcloud.Mysql.Inputs.ProxyProxyNodeCustomArgs
    ///                 {
    ///                     Cpu = 2,
    ///                     Mem = 4000,
    ///                     NodeCount = 1,
    ///                     Region = "ap-guangzhou",
    ///                     Zone = "ap-guangzhou-3",
    ///                 },
    ///             },
    ///             SecurityGroups = 
    ///             {
    ///                 "sg-edmur627",
    ///             },
    ///             UniqSubnetId = "subnet-ahv6swf2",
    ///             UniqVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.17.101",
    ///             Vport = 3306,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/proxy:Proxy")]
    public partial class Proxy : Pulumi.CustomResource
    {
        /// <summary>
        /// Connection Pool Threshold.
        /// </summary>
        [Output("connectionPoolLimit")]
        public Output<int?> ConnectionPoolLimit { get; private set; } = null!;

        /// <summary>
        /// Describe.
        /// </summary>
        [Output("desc")]
        public Output<string?> Desc { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Proxy address id.
        /// </summary>
        [Output("proxyAddressId")]
        public Output<string> ProxyAddressId { get; private set; } = null!;

        /// <summary>
        /// Proxy group id.
        /// </summary>
        [Output("proxyGroupId")]
        public Output<string> ProxyGroupId { get; private set; } = null!;

        /// <summary>
        /// Node specification configuration.
        /// </summary>
        [Output("proxyNodeCustoms")]
        public Output<ImmutableArray<Outputs.ProxyProxyNodeCustom>> ProxyNodeCustoms { get; private set; } = null!;

        /// <summary>
        /// The current version of the database agent. No need to fill in when creating.
        /// </summary>
        [Output("proxyVersion")]
        public Output<string> ProxyVersion { get; private set; } = null!;

        /// <summary>
        /// Security group.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Output("uniqSubnetId")]
        public Output<string> UniqSubnetId { get; private set; } = null!;

        /// <summary>
        /// Vpc id.
        /// </summary>
        [Output("uniqVpcId")]
        public Output<string> UniqVpcId { get; private set; } = null!;

        /// <summary>
        /// Upgrade time: nowTime (upgrade completed) timeWindow (instance maintenance time), Required when modifying the agent version, No need to fill in when creating.
        /// </summary>
        [Output("upgradeTime")]
        public Output<string?> UpgradeTime { get; private set; } = null!;

        /// <summary>
        /// IP address.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// Port.
        /// </summary>
        [Output("vport")]
        public Output<int> Vport { get; private set; } = null!;


        /// <summary>
        /// Create a Proxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Proxy(string name, ProxyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/proxy:Proxy", name, args ?? new ProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Proxy(string name, Input<string> id, ProxyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/proxy:Proxy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Proxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Proxy Get(string name, Input<string> id, ProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new Proxy(name, id, state, options);
        }
    }

    public sealed class ProxyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection Pool Threshold.
        /// </summary>
        [Input("connectionPoolLimit")]
        public Input<int>? ConnectionPoolLimit { get; set; }

        /// <summary>
        /// Describe.
        /// </summary>
        [Input("desc")]
        public Input<string>? Desc { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("proxyNodeCustoms", required: true)]
        private InputList<Inputs.ProxyProxyNodeCustomArgs>? _proxyNodeCustoms;

        /// <summary>
        /// Node specification configuration.
        /// </summary>
        public InputList<Inputs.ProxyProxyNodeCustomArgs> ProxyNodeCustoms
        {
            get => _proxyNodeCustoms ?? (_proxyNodeCustoms = new InputList<Inputs.ProxyProxyNodeCustomArgs>());
            set => _proxyNodeCustoms = value;
        }

        /// <summary>
        /// The current version of the database agent. No need to fill in when creating.
        /// </summary>
        [Input("proxyVersion")]
        public Input<string>? ProxyVersion { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security group.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Input("uniqSubnetId", required: true)]
        public Input<string> UniqSubnetId { get; set; } = null!;

        /// <summary>
        /// Vpc id.
        /// </summary>
        [Input("uniqVpcId", required: true)]
        public Input<string> UniqVpcId { get; set; } = null!;

        /// <summary>
        /// Upgrade time: nowTime (upgrade completed) timeWindow (instance maintenance time), Required when modifying the agent version, No need to fill in when creating.
        /// </summary>
        [Input("upgradeTime")]
        public Input<string>? UpgradeTime { get; set; }

        /// <summary>
        /// IP address.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Port.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        public ProxyArgs()
        {
        }
    }

    public sealed class ProxyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection Pool Threshold.
        /// </summary>
        [Input("connectionPoolLimit")]
        public Input<int>? ConnectionPoolLimit { get; set; }

        /// <summary>
        /// Describe.
        /// </summary>
        [Input("desc")]
        public Input<string>? Desc { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Proxy address id.
        /// </summary>
        [Input("proxyAddressId")]
        public Input<string>? ProxyAddressId { get; set; }

        /// <summary>
        /// Proxy group id.
        /// </summary>
        [Input("proxyGroupId")]
        public Input<string>? ProxyGroupId { get; set; }

        [Input("proxyNodeCustoms")]
        private InputList<Inputs.ProxyProxyNodeCustomGetArgs>? _proxyNodeCustoms;

        /// <summary>
        /// Node specification configuration.
        /// </summary>
        public InputList<Inputs.ProxyProxyNodeCustomGetArgs> ProxyNodeCustoms
        {
            get => _proxyNodeCustoms ?? (_proxyNodeCustoms = new InputList<Inputs.ProxyProxyNodeCustomGetArgs>());
            set => _proxyNodeCustoms = value;
        }

        /// <summary>
        /// The current version of the database agent. No need to fill in when creating.
        /// </summary>
        [Input("proxyVersion")]
        public Input<string>? ProxyVersion { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security group.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Input("uniqSubnetId")]
        public Input<string>? UniqSubnetId { get; set; }

        /// <summary>
        /// Vpc id.
        /// </summary>
        [Input("uniqVpcId")]
        public Input<string>? UniqVpcId { get; set; }

        /// <summary>
        /// Upgrade time: nowTime (upgrade completed) timeWindow (instance maintenance time), Required when modifying the agent version, No need to fill in when creating.
        /// </summary>
        [Input("upgradeTime")]
        public Input<string>? UpgradeTime { get; set; }

        /// <summary>
        /// IP address.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Port.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        public ProxyState()
        {
        }
    }
}
