// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod
{
    /// <summary>
    /// Provide a resource to create a DnsPod record.
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
    ///         var demo = new Tencentcloud.Dnspod.Record("demo", new Tencentcloud.Dnspod.RecordArgs
    ///         {
    ///             Domain = "mikatong.com",
    ///             RecordLine = "默认",
    ///             RecordType = "A",
    ///             SubDomain = "demo",
    ///             Value = "1.2.3.9",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// DnsPod Domain record can be imported using the Domain#RecordId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dnspod/record:Record demo arunma.com#1194109872
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dnspod/record:Record")]
    public partial class Record : Pulumi.CustomResource
    {
        /// <summary>
        /// The Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The D monitoring status of the record.
        /// </summary>
        [Output("monitorStatus")]
        public Output<string> MonitorStatus { get; private set; } = null!;

        /// <summary>
        /// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
        /// </summary>
        [Output("mx")]
        public Output<int?> Mx { get; private set; } = null!;

        /// <summary>
        /// The record line.
        /// </summary>
        [Output("recordLine")]
        public Output<string> RecordLine { get; private set; } = null!;

        /// <summary>
        /// The record type.
        /// </summary>
        [Output("recordType")]
        public Output<string> RecordType { get; private set; } = null!;

        /// <summary>
        /// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The host records, default value is `@`.
        /// </summary>
        [Output("subDomain")]
        public Output<string?> SubDomain { get; private set; } = null!;

        /// <summary>
        /// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The record value.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a Record resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Record(string name, RecordArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/record:Record", name, args ?? new RecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Record(string name, Input<string> id, RecordState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/record:Record", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Record resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Record Get(string name, Input<string> id, RecordState? state = null, CustomResourceOptions? options = null)
        {
            return new Record(name, id, state, options);
        }
    }

    public sealed class RecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
        /// </summary>
        [Input("mx")]
        public Input<int>? Mx { get; set; }

        /// <summary>
        /// The record line.
        /// </summary>
        [Input("recordLine", required: true)]
        public Input<string> RecordLine { get; set; } = null!;

        /// <summary>
        /// The record type.
        /// </summary>
        [Input("recordType", required: true)]
        public Input<string> RecordType { get; set; } = null!;

        /// <summary>
        /// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The host records, default value is `@`.
        /// </summary>
        [Input("subDomain")]
        public Input<string>? SubDomain { get; set; }

        /// <summary>
        /// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The record value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public RecordArgs()
        {
        }
    }

    public sealed class RecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The D monitoring status of the record.
        /// </summary>
        [Input("monitorStatus")]
        public Input<string>? MonitorStatus { get; set; }

        /// <summary>
        /// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
        /// </summary>
        [Input("mx")]
        public Input<int>? Mx { get; set; }

        /// <summary>
        /// The record line.
        /// </summary>
        [Input("recordLine")]
        public Input<string>? RecordLine { get; set; }

        /// <summary>
        /// The record type.
        /// </summary>
        [Input("recordType")]
        public Input<string>? RecordType { get; set; }

        /// <summary>
        /// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The host records, default value is `@`.
        /// </summary>
        [Input("subDomain")]
        public Input<string>? SubDomain { get; set; }

        /// <summary>
        /// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The record value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public RecordState()
        {
        }
    }
}
