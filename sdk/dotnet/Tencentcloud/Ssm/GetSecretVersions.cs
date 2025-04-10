// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssm
{
    public static class GetSecretVersions
    {
        /// <summary>
        /// Use this data source to query detailed information of SSM secret version
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
        ///         var foo = Output.Create(Tencentcloud.Ssm.GetSecretVersions.InvokeAsync(new Tencentcloud.Ssm.GetSecretVersionsArgs
        ///         {
        ///             SecretName = "test",
        ///             VersionId = "v1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecretVersionsResult> InvokeAsync(GetSecretVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionsResult>("tencentcloud:Ssm/getSecretVersions:getSecretVersions", args ?? new GetSecretVersionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of SSM secret version
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
        ///         var foo = Output.Create(Tencentcloud.Ssm.GetSecretVersions.InvokeAsync(new Tencentcloud.Ssm.GetSecretVersionsArgs
        ///         {
        ///             SecretName = "test",
        ///             VersionId = "v1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecretVersionsResult> Invoke(GetSecretVersionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecretVersionsResult>("tencentcloud:Ssm/getSecretVersions:getSecretVersions", args ?? new GetSecretVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Secret name used to filter result.
        /// </summary>
        [Input("secretName", required: true)]
        public string SecretName { get; set; } = null!;

        /// <summary>
        /// VersionId used to filter result.
        /// </summary>
        [Input("versionId")]
        public string? VersionId { get; set; }

        public GetSecretVersionsArgs()
        {
        }
    }

    public sealed class GetSecretVersionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Secret name used to filter result.
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        /// <summary>
        /// VersionId used to filter result.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public GetSecretVersionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecretVersionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string SecretName;
        /// <summary>
        /// A list of SSM secret versions. When secret status is `Disabled`, this field will not update anymore.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecretVersionsSecretVersionListResult> SecretVersionLists;
        /// <summary>
        /// Version of secret.
        /// </summary>
        public readonly string? VersionId;

        [OutputConstructor]
        private GetSecretVersionsResult(
            string id,

            string? resultOutputFile,

            string secretName,

            ImmutableArray<Outputs.GetSecretVersionsSecretVersionListResult> secretVersionLists,

            string? versionId)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SecretName = secretName;
            SecretVersionLists = secretVersionLists;
            VersionId = versionId;
        }
    }
}
