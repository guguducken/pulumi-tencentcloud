// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    /// <summary>
    /// Provides a resource to create a postgresql renew_db_instance_operation
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
    ///         var renewDbInstanceOperation = new Tencentcloud.Postgresql.RenewDbInstanceOperation("renewDbInstanceOperation", new Tencentcloud.Postgresql.RenewDbInstanceOperationArgs
    ///         {
    ///             DbInstanceId = tencentcloud_postgresql_instance.Oper_test_PREPAID.Id,
    ///             Period = 1,
    ///             AutoVoucher = 0,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation")]
    public partial class RenewDbInstanceOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
        /// </summary>
        [Output("autoVoucher")]
        public Output<int?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Instance ID in the format of postgres-6fego161.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// Renewal duration in months.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;

        /// <summary>
        /// Voucher ID list (only one voucher can be specified currently).
        /// </summary>
        [Output("voucherIds")]
        public Output<ImmutableArray<string>> VoucherIds { get; private set; } = null!;


        /// <summary>
        /// Create a RenewDbInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RenewDbInstanceOperation(string name, RenewDbInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation", name, args ?? new RenewDbInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RenewDbInstanceOperation(string name, Input<string> id, RenewDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RenewDbInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RenewDbInstanceOperation Get(string name, Input<string> id, RenewDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RenewDbInstanceOperation(name, id, state, options);
        }
    }

    public sealed class RenewDbInstanceOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Instance ID in the format of postgres-6fego161.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// Renewal duration in months.
        /// </summary>
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// Voucher ID list (only one voucher can be specified currently).
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        public RenewDbInstanceOperationArgs()
        {
        }
    }

    public sealed class RenewDbInstanceOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Instance ID in the format of postgres-6fego161.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// Renewal duration in months.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// Voucher ID list (only one voucher can be specified currently).
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        public RenewDbInstanceOperationState()
        {
        }
    }
}
