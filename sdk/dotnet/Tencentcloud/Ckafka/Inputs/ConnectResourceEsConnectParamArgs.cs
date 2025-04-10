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

    public sealed class ConnectResourceEsConnectParamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to update to the associated Datahub task, default: false.
        /// </summary>
        [Input("isUpdate")]
        public Input<bool>? IsUpdate { get; set; }

        /// <summary>
        /// Es The password of the connection source.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Es port.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Instance resource of Es connection source.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Whether the Es connection source is a self-built cluster.
        /// </summary>
        [Input("selfBuilt", required: true)]
        public Input<bool> SelfBuilt { get; set; } = null!;

        /// <summary>
        /// The instance vip of the Es connection source, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        [Input("serviceVip")]
        public Input<string>? ServiceVip { get; set; }

        /// <summary>
        /// The vpc Id of the Es connection source, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        [Input("uniqVpcId")]
        public Input<string>? UniqVpcId { get; set; }

        /// <summary>
        /// Es The username of the connection source.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public ConnectResourceEsConnectParamArgs()
        {
        }
    }
}
