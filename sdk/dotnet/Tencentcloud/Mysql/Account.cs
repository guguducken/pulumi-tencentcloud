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
    /// Provides a MySQL account resource for database management. A MySQL instance supports multiple database account.
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
    ///         var @default = new Tencentcloud.Mysql.Account("default", new Tencentcloud.Mysql.AccountArgs
    ///         {
    ///             Description = "My test account",
    ///             MaxUserConnections = 10,
    ///             MysqlId = "terraform-test-local-database",
    ///             Password = "********",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mysql account can be imported using the mysqlId#accountName, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mysql/account:Account default cdb-gqg6j82x#tf_account
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/account:Account")]
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// Database description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Account host, default is `%`.
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// The maximum number of available connections for a new account, the default value is 10240, and the maximum value that can be set is 10240.
        /// </summary>
        [Output("maxUserConnections")]
        public Output<int> MaxUserConnections { get; private set; } = null!;

        /// <summary>
        /// Instance ID to which the account belongs.
        /// </summary>
        [Output("mysqlId")]
        public Output<string> MysqlId { get; private set; } = null!;

        /// <summary>
        /// Account name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Operation password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Database description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Account host, default is `%`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The maximum number of available connections for a new account, the default value is 10240, and the maximum value that can be set is 10240.
        /// </summary>
        [Input("maxUserConnections")]
        public Input<int>? MaxUserConnections { get; set; }

        /// <summary>
        /// Instance ID to which the account belongs.
        /// </summary>
        [Input("mysqlId", required: true)]
        public Input<string> MysqlId { get; set; } = null!;

        /// <summary>
        /// Account name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Operation password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        public AccountArgs()
        {
        }
    }

    public sealed class AccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Account host, default is `%`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The maximum number of available connections for a new account, the default value is 10240, and the maximum value that can be set is 10240.
        /// </summary>
        [Input("maxUserConnections")]
        public Input<int>? MaxUserConnections { get; set; }

        /// <summary>
        /// Instance ID to which the account belongs.
        /// </summary>
        [Input("mysqlId")]
        public Input<string>? MysqlId { get; set; }

        /// <summary>
        /// Account name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Operation password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public AccountState()
        {
        }
    }
}
