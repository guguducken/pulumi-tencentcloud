// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Chdfs
{
    public static class GetAccessGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of chdfs access_groups
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
        ///         var accessGroups = Output.Create(Tencentcloud.Chdfs.GetAccessGroups.InvokeAsync(new Tencentcloud.Chdfs.GetAccessGroupsArgs
        ///         {
        ///             VpcId = "vpc-pewdpc0d",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccessGroupsResult> InvokeAsync(GetAccessGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessGroupsResult>("tencentcloud:Chdfs/getAccessGroups:getAccessGroups", args ?? new GetAccessGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of chdfs access_groups
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
        ///         var accessGroups = Output.Create(Tencentcloud.Chdfs.GetAccessGroups.InvokeAsync(new Tencentcloud.Chdfs.GetAccessGroupsArgs
        ///         {
        ///             VpcId = "vpc-pewdpc0d",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccessGroupsResult> Invoke(GetAccessGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccessGroupsResult>("tencentcloud:Chdfs/getAccessGroups:getAccessGroups", args ?? new GetAccessGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// get groups belongs to the owner uin, must set but only can use one of VpcId and OwnerUin to get the groups.
        /// </summary>
        [Input("ownerUin")]
        public int? OwnerUin { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetAccessGroupsArgs()
        {
        }
    }

    public sealed class GetAccessGroupsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// get groups belongs to the owner uin, must set but only can use one of VpcId and OwnerUin to get the groups.
        /// </summary>
        [Input("ownerUin")]
        public Input<int>? OwnerUin { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetAccessGroupsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessGroupsResult
    {
        /// <summary>
        /// access group list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessGroupsAccessGroupResult> ChdfsAccessGroups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? OwnerUin;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// VPC ID.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private GetAccessGroupsResult(
            ImmutableArray<Outputs.GetAccessGroupsAccessGroupResult> accessGroups,

            string id,

            int? ownerUin,

            string? resultOutputFile,

            string? vpcId)
        {
            ChdfsAccessGroups = accessGroups;
            Id = id;
            OwnerUin = ownerUin;
            ResultOutputFile = resultOutputFile;
            VpcId = vpcId;
        }
    }
}
