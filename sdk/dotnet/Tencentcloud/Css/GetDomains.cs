// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css
{
    public static class GetDomains
    {
        /// <summary>
        /// Use this data source to query detailed information of css domains
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var domains = Output.Create(Tencentcloud.Css.GetDomains.InvokeAsync(new Tencentcloud.Css.GetDomainsArgs
        ///         {
        ///             DomainType = 0,
        ///             IsDelayLive = 0,
        ///             PlayType = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("tencentcloud:Css/getDomains:getDomains", args ?? new GetDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of css domains
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var domains = Output.Create(Tencentcloud.Css.GetDomains.InvokeAsync(new Tencentcloud.Css.GetDomainsArgs
        ///         {
        ///             DomainType = 0,
        ///             IsDelayLive = 0,
        ///             PlayType = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDomainsResult> Invoke(GetDomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainsResult>("tencentcloud:Css/getDomains:getDomains", args ?? new GetDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// domain name prefix.
        /// </summary>
        [Input("domainPrefix")]
        public string? DomainPrefix { get; set; }

        /// <summary>
        /// domain name status filter. 0-disable, 1-enable.
        /// </summary>
        [Input("domainStatus")]
        public int? DomainStatus { get; set; }

        /// <summary>
        /// Domain name type filtering. 0-push, 1-play.
        /// </summary>
        [Input("domainType")]
        public int? DomainType { get; set; }

        /// <summary>
        /// 0 normal live broadcast 1 slow live broadcast default 0.
        /// </summary>
        [Input("isDelayLive")]
        public int? IsDelayLive { get; set; }

        /// <summary>
        /// Playing area, this parameter is meaningful only when DomainType=1. 1: Domestic.2: Global.3: Overseas.
        /// </summary>
        [Input("playType")]
        public int? PlayType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDomainsArgs()
        {
        }
    }

    public sealed class GetDomainsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// domain name prefix.
        /// </summary>
        [Input("domainPrefix")]
        public Input<string>? DomainPrefix { get; set; }

        /// <summary>
        /// domain name status filter. 0-disable, 1-enable.
        /// </summary>
        [Input("domainStatus")]
        public Input<int>? DomainStatus { get; set; }

        /// <summary>
        /// Domain name type filtering. 0-push, 1-play.
        /// </summary>
        [Input("domainType")]
        public Input<int>? DomainType { get; set; }

        /// <summary>
        /// 0 normal live broadcast 1 slow live broadcast default 0.
        /// </summary>
        [Input("isDelayLive")]
        public Input<int>? IsDelayLive { get; set; }

        /// <summary>
        /// Playing area, this parameter is meaningful only when DomainType=1. 1: Domestic.2: Global.3: Overseas.
        /// </summary>
        [Input("playType")]
        public Input<int>? PlayType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDomainsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        /// <summary>
        /// A list of domain name details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainListResult> DomainLists;
        public readonly string? DomainPrefix;
        public readonly int? DomainStatus;
        public readonly int? DomainType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether to slow live broadcast: 0: normal live broadcast. 1: Slow live broadcast.
        /// </summary>
        public readonly int? IsDelayLive;
        /// <summary>
        /// Playing area, this parameter is meaningful only when Type=1. 1: Domestic. 2: Global. 3: Overseas.
        /// </summary>
        public readonly int? PlayType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDomainsResult(
            ImmutableArray<Outputs.GetDomainsDomainListResult> domainLists,

            string? domainPrefix,

            int? domainStatus,

            int? domainType,

            string id,

            int? isDelayLive,

            int? playType,

            string? resultOutputFile)
        {
            DomainLists = domainLists;
            DomainPrefix = domainPrefix;
            DomainStatus = domainStatus;
            DomainType = domainType;
            Id = id;
            IsDelayLive = isDelayLive;
            PlayType = playType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
