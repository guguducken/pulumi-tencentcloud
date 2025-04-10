// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cam
{
    public static class GetGroupPolicyAttachments
    {
        /// <summary>
        /// Use this data source to query detailed information of CAM group policy attachments
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
        ///         var foo = Output.Create(Tencentcloud.Cam.GetGroupPolicyAttachments.InvokeAsync(new Tencentcloud.Cam.GetGroupPolicyAttachmentsArgs
        ///         {
        ///             GroupId = tencentcloud_cam_group.Foo.Id,
        ///         }));
        ///         var bar = Output.Create(Tencentcloud.Cam.GetGroupPolicyAttachments.InvokeAsync(new Tencentcloud.Cam.GetGroupPolicyAttachmentsArgs
        ///         {
        ///             GroupId = tencentcloud_cam_group.Foo.Id,
        ///             PolicyId = tencentcloud_cam_policy.Foo.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupPolicyAttachmentsResult> InvokeAsync(GetGroupPolicyAttachmentsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupPolicyAttachmentsResult>("tencentcloud:Cam/getGroupPolicyAttachments:getGroupPolicyAttachments", args ?? new GetGroupPolicyAttachmentsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of CAM group policy attachments
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
        ///         var foo = Output.Create(Tencentcloud.Cam.GetGroupPolicyAttachments.InvokeAsync(new Tencentcloud.Cam.GetGroupPolicyAttachmentsArgs
        ///         {
        ///             GroupId = tencentcloud_cam_group.Foo.Id,
        ///         }));
        ///         var bar = Output.Create(Tencentcloud.Cam.GetGroupPolicyAttachments.InvokeAsync(new Tencentcloud.Cam.GetGroupPolicyAttachmentsArgs
        ///         {
        ///             GroupId = tencentcloud_cam_group.Foo.Id,
        ///             PolicyId = tencentcloud_cam_policy.Foo.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupPolicyAttachmentsResult> Invoke(GetGroupPolicyAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupPolicyAttachmentsResult>("tencentcloud:Cam/getGroupPolicyAttachments:getGroupPolicyAttachments", args ?? new GetGroupPolicyAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupPolicyAttachmentsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Mode of creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
        /// </summary>
        [Input("createMode")]
        public int? CreateMode { get; set; }

        /// <summary>
        /// ID of the attached CAM group to be queried.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        /// <summary>
        /// ID of CAM policy to be queried.
        /// </summary>
        [Input("policyId")]
        public string? PolicyId { get; set; }

        /// <summary>
        /// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        /// </summary>
        [Input("policyType")]
        public string? PolicyType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetGroupPolicyAttachmentsArgs()
        {
        }
    }

    public sealed class GetGroupPolicyAttachmentsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Mode of creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
        /// </summary>
        [Input("createMode")]
        public Input<int>? CreateMode { get; set; }

        /// <summary>
        /// ID of the attached CAM group to be queried.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// ID of CAM policy to be queried.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetGroupPolicyAttachmentsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupPolicyAttachmentsResult
    {
        /// <summary>
        /// Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
        /// </summary>
        public readonly int? CreateMode;
        /// <summary>
        /// ID of CAM group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// A list of CAM group policy attachments. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupPolicyAttachmentsGroupPolicyAttachmentListResult> GroupPolicyAttachmentLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of CAM group.
        /// </summary>
        public readonly string? PolicyId;
        /// <summary>
        /// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        /// </summary>
        public readonly string? PolicyType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetGroupPolicyAttachmentsResult(
            int? createMode,

            string groupId,

            ImmutableArray<Outputs.GetGroupPolicyAttachmentsGroupPolicyAttachmentListResult> groupPolicyAttachmentLists,

            string id,

            string? policyId,

            string? policyType,

            string? resultOutputFile)
        {
            CreateMode = createMode;
            GroupId = groupId;
            GroupPolicyAttachmentLists = groupPolicyAttachmentLists;
            Id = id;
            PolicyId = policyId;
            PolicyType = policyType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
