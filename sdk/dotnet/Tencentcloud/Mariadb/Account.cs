// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    /// <summary>
    /// Provides a resource to create a mariadb account
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
    ///         var account = new Tencentcloud.Mariadb.Account("account", new Tencentcloud.Mariadb.AccountArgs
    ///         {
    ///             Description = "desc",
    ///             Host = "10.101.202.22",
    ///             InstanceId = "tdsql-4pzs5b67",
    ///             Password = "Password123.",
    ///             ReadOnly = 0,
    ///             UserName = "account-test",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mariadb account can be imported using the instance_id#user_name#host, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mariadb/account:Account account tdsql-4pzs5b67#account-test#10.101.202.22
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mariadb/account:Account")]
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// account description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// host.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// account password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// wether account is read only, 0 means not a read only account.
        /// </summary>
        [Output("readOnly")]
        public Output<int?> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// user name.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// host.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// account password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// wether account is read only, 0 means not a read only account.
        /// </summary>
        [Input("readOnly")]
        public Input<int>? ReadOnly { get; set; }

        /// <summary>
        /// user name.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public AccountArgs()
        {
        }
    }

    public sealed class AccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// account password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// wether account is read only, 0 means not a read only account.
        /// </summary>
        [Input("readOnly")]
        public Input<int>? ReadOnly { get; set; }

        /// <summary>
        /// user name.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public AccountState()
        {
        }
    }
}
