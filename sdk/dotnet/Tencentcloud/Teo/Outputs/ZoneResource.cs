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
    public sealed class ZoneResource
    {
        /// <summary>
        /// Valid values: `mainland`, `overseas`.
        /// </summary>
        public readonly string? Area;
        /// <summary>
        /// Whether to automatically renew. Valid values:- `0`: Default.- `1`: Enable automatic renewal.- `2`: Disable automatic renewal.
        /// </summary>
        public readonly int? AutoRenewFlag;
        /// <summary>
        /// Resource creation date.
        /// </summary>
        public readonly string? CreateTime;
        /// <summary>
        /// Enable time of the resource.
        /// </summary>
        public readonly string? EnableTime;
        /// <summary>
        /// Expire time of the resource.
        /// </summary>
        public readonly string? ExpireTime;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource pay mode. Valid values:- `0`: post pay mode.
        /// </summary>
        public readonly int? PayMode;
        /// <summary>
        /// Associated plan ID.
        /// </summary>
        public readonly string? PlanId;
        /// <summary>
        /// Site status. Valid values:- `active`: NS is switched.- `pending`: NS is not switched.- `moved`: NS is moved.- `deactivated`: this site is blocked.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Price inquiry parameters.
        /// </summary>
        public readonly ImmutableArray<Outputs.ZoneResourceSv> Svs;

        [OutputConstructor]
        private ZoneResource(
            string? area,

            int? autoRenewFlag,

            string? createTime,

            string? enableTime,

            string? expireTime,

            string? id,

            int? payMode,

            string? planId,

            string? status,

            ImmutableArray<Outputs.ZoneResourceSv> svs)
        {
            Area = area;
            AutoRenewFlag = autoRenewFlag;
            CreateTime = createTime;
            EnableTime = enableTime;
            ExpireTime = expireTime;
            Id = id;
            PayMode = payMode;
            PlanId = planId;
            Status = status;
            Svs = svs;
        }
    }
}
