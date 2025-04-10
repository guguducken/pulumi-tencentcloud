// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Availability
{
    public static class GetRegions
    {
        /// <summary>
        /// Use this data source to get the available regions. By default only `AVAILABLE` regions will be returned, but `UNAVAILABLE` regions can also be fetched when `include_unavailable` is specified.
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
        ///         var myFavouriteRegion = Output.Create(Tencentcloud.Availability.GetRegions.InvokeAsync(new Tencentcloud.Availability.GetRegionsArgs
        ///         {
        ///             Name = "ap-guangzhou",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegionsResult> InvokeAsync(GetRegionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionsResult>("tencentcloud:Availability/getRegions:getRegions", args ?? new GetRegionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the available regions. By default only `AVAILABLE` regions will be returned, but `UNAVAILABLE` regions can also be fetched when `include_unavailable` is specified.
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
        ///         var myFavouriteRegion = Output.Create(Tencentcloud.Availability.GetRegions.InvokeAsync(new Tencentcloud.Availability.GetRegionsArgs
        ///         {
        ///             Name = "ap-guangzhou",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegionsResult> Invoke(GetRegionsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionsResult>("tencentcloud:Availability/getRegions:getRegions", args ?? new GetRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A bool variable indicates that the query will include `UNAVAILABLE` regions.
        /// </summary>
        [Input("includeUnavailable")]
        public bool? IncludeUnavailable { get; set; }

        /// <summary>
        /// When specified, only the region with the exactly name match will be returned. `default` value means it consistent with the provider region.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRegionsArgs()
        {
        }
    }

    public sealed class GetRegionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A bool variable indicates that the query will include `UNAVAILABLE` regions.
        /// </summary>
        [Input("includeUnavailable")]
        public Input<bool>? IncludeUnavailable { get; set; }

        /// <summary>
        /// When specified, only the region with the exactly name match will be returned. `default` value means it consistent with the provider region.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRegionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeUnavailable;
        /// <summary>
        /// The name of the region, like `ap-guangzhou`.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A list of regions will be exported and its every element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegionsRegionResult> AvailabilityRegions;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRegionsResult(
            string id,

            bool? includeUnavailable,

            string? name,

            ImmutableArray<Outputs.GetRegionsRegionResult> regions,

            string? resultOutputFile)
        {
            Id = id;
            IncludeUnavailable = includeUnavailable;
            Name = name;
            AvailabilityRegions = regions;
            ResultOutputFile = resultOutputFile;
        }
    }
}
