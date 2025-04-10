// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Inputs
{

    public sealed class AccountPrivilegesAccountsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// user host.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// user name.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public AccountPrivilegesAccountsArgs()
        {
        }
    }
}
