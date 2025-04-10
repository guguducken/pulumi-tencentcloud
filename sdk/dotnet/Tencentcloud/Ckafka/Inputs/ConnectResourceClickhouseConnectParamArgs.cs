// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class ConnectResourceClickhouseConnectParamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to update to the associated Datahub task, default: false.
        /// </summary>
        [Input("isUpdate")]
        public Input<bool>? IsUpdate { get; set; }

        /// <summary>
        /// Password for Clickhouse connection source.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Clickhouse connection port.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Instance resources for Click House connection sources.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Whether the Clickhouse connection source is a self-built cluster.
        /// </summary>
        [Input("selfBuilt", required: true)]
        public Input<bool> SelfBuilt { get; set; } = null!;

        /// <summary>
        /// Instance VIP of the ClickHouse connection source, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        [Input("serviceVip")]
        public Input<string>? ServiceVip { get; set; }

        /// <summary>
        /// The vpc Id of the source of the ClickHouse connection, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        [Input("uniqVpcId")]
        public Input<string>? UniqVpcId { get; set; }

        /// <summary>
        /// The username of the clickhouse connection source.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public ConnectResourceClickhouseConnectParamArgs()
        {
        }
    }
}
