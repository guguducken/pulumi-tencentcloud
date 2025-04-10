// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm
{
    [TencentcloudResourceType("tencentcloud:Cvm/renewHost:RenewHost")]
    public partial class RenewHost : Pulumi.CustomResource
    {
        /// <summary>
        /// Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can
        /// specify attributes such as the purchase duration of the Subscription instance and whether to set automatic renewal. If
        /// the payment mode of the specified instance is prepaid, this parameter must be passed.
        /// </summary>
        [Output("hostChargePrepaid")]
        public Output<Outputs.RenewHostHostChargePrepaid> HostChargePrepaid { get; private set; } = null!;

        /// <summary>
        /// CDH instance ID.
        /// </summary>
        [Output("hostId")]
        public Output<string> HostId { get; private set; } = null!;


        /// <summary>
        /// Create a RenewHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RenewHost(string name, RenewHostArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/renewHost:RenewHost", name, args ?? new RenewHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RenewHost(string name, Input<string> id, RenewHostState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/renewHost:RenewHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RenewHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RenewHost Get(string name, Input<string> id, RenewHostState? state = null, CustomResourceOptions? options = null)
        {
            return new RenewHost(name, id, state, options);
        }
    }

    public sealed class RenewHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can
        /// specify attributes such as the purchase duration of the Subscription instance and whether to set automatic renewal. If
        /// the payment mode of the specified instance is prepaid, this parameter must be passed.
        /// </summary>
        [Input("hostChargePrepaid", required: true)]
        public Input<Inputs.RenewHostHostChargePrepaidArgs> HostChargePrepaid { get; set; } = null!;

        /// <summary>
        /// CDH instance ID.
        /// </summary>
        [Input("hostId", required: true)]
        public Input<string> HostId { get; set; } = null!;

        public RenewHostArgs()
        {
        }
    }

    public sealed class RenewHostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can
        /// specify attributes such as the purchase duration of the Subscription instance and whether to set automatic renewal. If
        /// the payment mode of the specified instance is prepaid, this parameter must be passed.
        /// </summary>
        [Input("hostChargePrepaid")]
        public Input<Inputs.RenewHostHostChargePrepaidGetArgs>? HostChargePrepaid { get; set; }

        /// <summary>
        /// CDH instance ID.
        /// </summary>
        [Input("hostId")]
        public Input<string>? HostId { get; set; }

        public RenewHostState()
        {
        }
    }
}
