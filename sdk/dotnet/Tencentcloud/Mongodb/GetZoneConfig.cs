// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mongodb
{
    public static class GetZoneConfig
    {
        /// <summary>
        /// Use this data source to query the available mongodb specifications for different zone.
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
        ///         var mongodb = Output.Create(Tencentcloud.Mongodb.GetZoneConfig.InvokeAsync(new Tencentcloud.Mongodb.GetZoneConfigArgs
        ///         {
        ///             AvailableZone = "ap-guangzhou-2",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZoneConfigResult> InvokeAsync(GetZoneConfigArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZoneConfigResult>("tencentcloud:Mongodb/getZoneConfig:getZoneConfig", args ?? new GetZoneConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query the available mongodb specifications for different zone.
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
        ///         var mongodb = Output.Create(Tencentcloud.Mongodb.GetZoneConfig.InvokeAsync(new Tencentcloud.Mongodb.GetZoneConfigArgs
        ///         {
        ///             AvailableZone = "ap-guangzhou-2",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZoneConfigResult> Invoke(GetZoneConfigInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetZoneConfigResult>("tencentcloud:Mongodb/getZoneConfig:getZoneConfig", args ?? new GetZoneConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The available zone of the Mongodb.
        /// </summary>
        [Input("availableZone")]
        public string? AvailableZone { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetZoneConfigArgs()
        {
        }
    }

    public sealed class GetZoneConfigInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The available zone of the Mongodb.
        /// </summary>
        [Input("availableZone")]
        public Input<string>? AvailableZone { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetZoneConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetZoneConfigResult
    {
        /// <summary>
        /// The available zone of the Mongodb.
        /// </summary>
        public readonly string? AvailableZone;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of zone config. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneConfigListResult> Lists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetZoneConfigResult(
            string? availableZone,

            string id,

            ImmutableArray<Outputs.GetZoneConfigListResult> lists,

            string? resultOutputFile)
        {
            AvailableZone = availableZone;
            Id = id;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
