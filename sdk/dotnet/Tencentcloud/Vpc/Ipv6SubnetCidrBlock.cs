// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    /// <summary>
    /// Provides a resource to create a vpc ipv6_subnet_cidr_block
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
    ///         var ipv6SubnetCidrBlock = new Tencentcloud.Vpc.Ipv6SubnetCidrBlock("ipv6SubnetCidrBlock", new Tencentcloud.Vpc.Ipv6SubnetCidrBlockArgs
    ///         {
    ///             Ipv6SubnetCidrBlocks = new Tencentcloud.Vpc.Inputs.Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksArgs
    ///             {
    ///                 Ipv6CidrBlock = "2402:4e00:1019:6a7b::/64",
    ///                 SubnetId = "subnet-plg028y8",
    ///             },
    ///             VpcId = "vpc-7w3kgnpl",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// vpc ipv6_subnet_cidr_block can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpc/ipv6SubnetCidrBlock:Ipv6SubnetCidrBlock ipv6_subnet_cidr_block ipv6_subnet_cidr_block_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/ipv6SubnetCidrBlock:Ipv6SubnetCidrBlock")]
    public partial class Ipv6SubnetCidrBlock : Pulumi.CustomResource
    {
        /// <summary>
        /// Allocate a list of `IPv6` subnets.
        /// </summary>
        [Output("ipv6SubnetCidrBlocks")]
        public Output<Outputs.Ipv6SubnetCidrBlockIpv6SubnetCidrBlocks> Ipv6SubnetCidrBlocks { get; private set; } = null!;

        /// <summary>
        /// The private network `ID` where the subnet is located. Such as:`vpc-f49l6u0z`.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv6SubnetCidrBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv6SubnetCidrBlock(string name, Ipv6SubnetCidrBlockArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/ipv6SubnetCidrBlock:Ipv6SubnetCidrBlock", name, args ?? new Ipv6SubnetCidrBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv6SubnetCidrBlock(string name, Input<string> id, Ipv6SubnetCidrBlockState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/ipv6SubnetCidrBlock:Ipv6SubnetCidrBlock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv6SubnetCidrBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv6SubnetCidrBlock Get(string name, Input<string> id, Ipv6SubnetCidrBlockState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv6SubnetCidrBlock(name, id, state, options);
        }
    }

    public sealed class Ipv6SubnetCidrBlockArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allocate a list of `IPv6` subnets.
        /// </summary>
        [Input("ipv6SubnetCidrBlocks", required: true)]
        public Input<Inputs.Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksArgs> Ipv6SubnetCidrBlocks { get; set; } = null!;

        /// <summary>
        /// The private network `ID` where the subnet is located. Such as:`vpc-f49l6u0z`.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public Ipv6SubnetCidrBlockArgs()
        {
        }
    }

    public sealed class Ipv6SubnetCidrBlockState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allocate a list of `IPv6` subnets.
        /// </summary>
        [Input("ipv6SubnetCidrBlocks")]
        public Input<Inputs.Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksGetArgs>? Ipv6SubnetCidrBlocks { get; set; }

        /// <summary>
        /// The private network `ID` where the subnet is located. Such as:`vpc-f49l6u0z`.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public Ipv6SubnetCidrBlockState()
        {
        }
    }
}
