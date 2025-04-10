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
    public sealed class GetSecurityPolicyRegionsGeoIpResult
    {
        /// <summary>
        /// Name of the continent.
        /// </summary>
        public readonly string Continent;
        /// <summary>
        /// Name of the country.
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// Province of the region. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string Province;
        /// <summary>
        /// Region ID.
        /// </summary>
        public readonly int RegionId;

        [OutputConstructor]
        private GetSecurityPolicyRegionsGeoIpResult(
            string continent,

            string country,

            string province,

            int regionId)
        {
            Continent = continent;
            Country = country;
            Province = province;
            RegionId = regionId;
        }
    }
}
