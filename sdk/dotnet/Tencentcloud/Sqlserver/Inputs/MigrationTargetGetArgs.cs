// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver.Inputs
{

    public sealed class MigrationTargetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the migration target instance, in the format mssql-si2823jyl.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Password of the migration target instance.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// User name of the migration target instance.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public MigrationTargetGetArgs()
        {
        }
    }
}
