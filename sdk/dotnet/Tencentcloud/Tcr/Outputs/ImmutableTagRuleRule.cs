// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr.Outputs
{

    [OutputType]
    public sealed class ImmutableTagRuleRule
    {
        /// <summary>
        /// disable rule.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// ID of the resource.
        /// </summary>
        public readonly int? Id;
        public readonly string? NsName;
        /// <summary>
        /// repository decoration type:repoMatches or repoExcludes.
        /// </summary>
        public readonly string RepositoryDecoration;
        /// <summary>
        /// repository matching rules.
        /// </summary>
        public readonly string RepositoryPattern;
        /// <summary>
        /// tag decoration type: matches or excludes.
        /// </summary>
        public readonly string TagDecoration;
        /// <summary>
        /// tag matching rules.
        /// </summary>
        public readonly string TagPattern;

        [OutputConstructor]
        private ImmutableTagRuleRule(
            bool? disabled,

            int? id,

            string? nsName,

            string repositoryDecoration,

            string repositoryPattern,

            string tagDecoration,

            string tagPattern)
        {
            Disabled = disabled;
            Id = id;
            NsName = nsName;
            RepositoryDecoration = repositoryDecoration;
            RepositoryPattern = repositoryPattern;
            TagDecoration = tagDecoration;
            TagPattern = tagPattern;
        }
    }
}
