// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class NetworkAclQuintupleNetworkAclQuintupleSetIngress
    {
        public readonly string? Action;
        public readonly string? CreateTime;
        public readonly string? Description;
        public readonly string? DestinationCidr;
        public readonly string? DestinationPort;
        public readonly string? NetworkAclDirection;
        public readonly string? NetworkAclQuintupleEntryId;
        public readonly int? Priority;
        public readonly string? Protocol;
        public readonly string? SourceCidr;
        public readonly string? SourcePort;

        [OutputConstructor]
        private NetworkAclQuintupleNetworkAclQuintupleSetIngress(
            string? action,

            string? createTime,

            string? description,

            string? destinationCidr,

            string? destinationPort,

            string? networkAclDirection,

            string? networkAclQuintupleEntryId,

            int? priority,

            string? protocol,

            string? sourceCidr,

            string? sourcePort)
        {
            Action = action;
            CreateTime = createTime;
            Description = description;
            DestinationCidr = destinationCidr;
            DestinationPort = destinationPort;
            NetworkAclDirection = networkAclDirection;
            NetworkAclQuintupleEntryId = networkAclQuintupleEntryId;
            Priority = priority;
            Protocol = protocol;
            SourceCidr = sourceCidr;
            SourcePort = sourcePort;
        }
    }
}
