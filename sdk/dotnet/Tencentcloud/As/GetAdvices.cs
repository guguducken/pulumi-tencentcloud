// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As
{
    public static class GetAdvices
    {
        /// <summary>
        /// Use this data source to query detailed information of as advices
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
        ///         var advices = Output.Create(Tencentcloud.As.GetAdvices.InvokeAsync(new Tencentcloud.As.GetAdvicesArgs
        ///         {
        ///             AutoScalingGroupIds = 
        ///             {
        ///                 "asc-lo0b94oy",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAdvicesResult> InvokeAsync(GetAdvicesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAdvicesResult>("tencentcloud:As/getAdvices:getAdvices", args ?? new GetAdvicesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of as advices
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
        ///         var advices = Output.Create(Tencentcloud.As.GetAdvices.InvokeAsync(new Tencentcloud.As.GetAdvicesArgs
        ///         {
        ///             AutoScalingGroupIds = 
        ///             {
        ///                 "asc-lo0b94oy",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAdvicesResult> Invoke(GetAdvicesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAdvicesResult>("tencentcloud:As/getAdvices:getAdvices", args ?? new GetAdvicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAdvicesArgs : Pulumi.InvokeArgs
    {
        [Input("autoScalingGroupIds", required: true)]
        private List<string>? _autoScalingGroupIds;

        /// <summary>
        /// List of scaling groups to be queried. Upper limit: 100.
        /// </summary>
        public List<string> AutoScalingGroupIds
        {
            get => _autoScalingGroupIds ?? (_autoScalingGroupIds = new List<string>());
            set => _autoScalingGroupIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAdvicesArgs()
        {
        }
    }

    public sealed class GetAdvicesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("autoScalingGroupIds", required: true)]
        private InputList<string>? _autoScalingGroupIds;

        /// <summary>
        /// List of scaling groups to be queried. Upper limit: 100.
        /// </summary>
        public InputList<string> AutoScalingGroupIds
        {
            get => _autoScalingGroupIds ?? (_autoScalingGroupIds = new InputList<string>());
            set => _autoScalingGroupIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAdvicesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAdvicesResult
    {
        /// <summary>
        /// A collection of suggestions for scaling group configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdvicesAutoScalingAdviceSetResult> AutoScalingAdviceSets;
        public readonly ImmutableArray<string> AutoScalingGroupIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetAdvicesResult(
            ImmutableArray<Outputs.GetAdvicesAutoScalingAdviceSetResult> autoScalingAdviceSets,

            ImmutableArray<string> autoScalingGroupIds,

            string id,

            string? resultOutputFile)
        {
            AutoScalingAdviceSets = autoScalingAdviceSets;
            AutoScalingGroupIds = autoScalingGroupIds;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
