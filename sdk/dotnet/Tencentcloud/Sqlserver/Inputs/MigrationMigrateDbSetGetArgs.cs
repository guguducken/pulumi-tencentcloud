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

    public sealed class MigrationMigrateDbSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the migration database.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        public MigrationMigrateDbSetGetArgs()
        {
        }
    }
}
