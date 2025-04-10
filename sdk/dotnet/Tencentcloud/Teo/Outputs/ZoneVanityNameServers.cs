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
    public sealed class ZoneVanityNameServers
    {
        /// <summary>
        /// List of custom name servers.
        /// </summary>
        public readonly ImmutableArray<string> Servers;
        /// <summary>
        /// Whether to enable the custom name server.- `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private ZoneVanityNameServers(
            ImmutableArray<string> servers,

            string @switch)
        {
            Servers = servers;
            Switch = @switch;
        }
    }
}
