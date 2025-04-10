// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class GetZoneDdosPolicyShieldAreaResult
    {
        /// <summary>
        /// DDoS layer 7 application.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneDdosPolicyShieldAreaApplicationResult> Applications;
        /// <summary>
        /// When `Type` is `domain`, this field is `ZoneId`. When `Type` is `application`, this field is `ProxyId`. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string Entity;
        /// <summary>
        /// When `Type` is `domain`, this field is `ZoneName`. When `Type` is `application`, this field is `ProxyName`. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string EntityName;
        /// <summary>
        /// Policy ID.
        /// </summary>
        public readonly int PolicyId;
        /// <summary>
        /// TCP forwarding rule number of layer 4 application.
        /// </summary>
        public readonly int TcpNum;
        /// <summary>
        /// Valid values: `domain`, `application`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// UDP forwarding rule number of layer 4 application.
        /// </summary>
        public readonly int UdpNum;
        /// <summary>
        /// Site ID.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetZoneDdosPolicyShieldAreaResult(
            ImmutableArray<Outputs.GetZoneDdosPolicyShieldAreaApplicationResult> applications,

            string entity,

            string entityName,

            int policyId,

            int tcpNum,

            string type,

            int udpNum,

            string zoneId)
        {
            Applications = applications;
            Entity = entity;
            EntityName = entityName;
            PolicyId = policyId;
            TcpNum = tcpNum;
            Type = type;
            UdpNum = udpNum;
            ZoneId = zoneId;
        }
    }
}
