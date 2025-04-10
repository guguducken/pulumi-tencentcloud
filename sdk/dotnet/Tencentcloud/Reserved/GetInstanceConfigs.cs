// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Reserved
{
    public static class GetInstanceConfigs
    {
        /// <summary>
        /// Use this data source to query reserved instances configuration.
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
        ///         var config = Output.Create(Tencentcloud.Reserved.GetInstanceConfigs.InvokeAsync(new Tencentcloud.Reserved.GetInstanceConfigsArgs
        ///         {
        ///             AvailabilityZone = "na-siliconvalley-1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceConfigsResult> InvokeAsync(GetInstanceConfigsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceConfigsResult>("tencentcloud:Reserved/getInstanceConfigs:getInstanceConfigs", args ?? new GetInstanceConfigsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query reserved instances configuration.
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
        ///         var config = Output.Create(Tencentcloud.Reserved.GetInstanceConfigs.InvokeAsync(new Tencentcloud.Reserved.GetInstanceConfigsArgs
        ///         {
        ///             AvailabilityZone = "na-siliconvalley-1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceConfigsResult> Invoke(GetInstanceConfigsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceConfigsResult>("tencentcloud:Reserved/getInstanceConfigs:getInstanceConfigs", args ?? new GetInstanceConfigsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceConfigsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The available zone that the reserved instance locates at.
        /// </summary>
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        /// <summary>
        /// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
        /// </summary>
        [Input("duration")]
        public int? Duration { get; set; }

        /// <summary>
        /// The type of reserved instance.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// Filter by Payment Type. Such as All Upfront.
        /// </summary>
        [Input("offeringType")]
        public string? OfferingType { get; set; }

        /// <summary>
        /// Filter by the Platform Description (that is, operating system) for Reserved Instance billing. Shaped like: linux.
        /// </summary>
        [Input("productDescription")]
        public string? ProductDescription { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceConfigsArgs()
        {
        }
    }

    public sealed class GetInstanceConfigsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The available zone that the reserved instance locates at.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The type of reserved instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Filter by Payment Type. Such as All Upfront.
        /// </summary>
        [Input("offeringType")]
        public Input<string>? OfferingType { get; set; }

        /// <summary>
        /// Filter by the Platform Description (that is, operating system) for Reserved Instance billing. Shaped like: linux.
        /// </summary>
        [Input("productDescription")]
        public Input<string>? ProductDescription { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceConfigsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceConfigsResult
    {
        /// <summary>
        /// Availability zone of the purchasable reserved instance.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// An information list of reserved instance configuration. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceConfigsConfigListResult> ConfigLists;
        /// <summary>
        /// Validity period of the reserved instance.
        /// </summary>
        public readonly int? Duration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Instance type of the reserved instance.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// OfferingType of the reserved instance.
        /// </summary>
        public readonly string? OfferingType;
        public readonly string? ProductDescription;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceConfigsResult(
            string? availabilityZone,

            ImmutableArray<Outputs.GetInstanceConfigsConfigListResult> configLists,

            int? duration,

            string id,

            string? instanceType,

            string? offeringType,

            string? productDescription,

            string? resultOutputFile)
        {
            AvailabilityZone = availabilityZone;
            ConfigLists = configLists;
            Duration = duration;
            Id = id;
            InstanceType = instanceType;
            OfferingType = offeringType;
            ProductDescription = productDescription;
            ResultOutputFile = resultOutputFile;
        }
    }
}
