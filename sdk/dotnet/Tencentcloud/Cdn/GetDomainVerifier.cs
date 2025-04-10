// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn
{
    public static class GetDomainVerifier
    {
        /// <summary>
        /// Provides a resource to check or create a cdn Domain Verify Record
        /// 
        /// &gt; **NOTE:**
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
        ///         var vr = Output.Create(Tencentcloud.Cdn.GetDomainVerifier.InvokeAsync(new Tencentcloud.Cdn.GetDomainVerifierArgs
        ///         {
        ///             Domain = "www.examplexxx123.com",
        ///             AutoVerify = true,
        ///             FreezeRecord = true,
        ///         }));
        ///         var recordValue = data.Tencentcloud_cdn_domain_verifier.Record;
        ///         var recordType = data.Tencentcloud_cdn_domain_verifier.Record_type;
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainVerifierResult> InvokeAsync(GetDomainVerifierArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainVerifierResult>("tencentcloud:Cdn/getDomainVerifier:getDomainVerifier", args ?? new GetDomainVerifierArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a resource to check or create a cdn Domain Verify Record
        /// 
        /// &gt; **NOTE:**
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
        ///         var vr = Output.Create(Tencentcloud.Cdn.GetDomainVerifier.InvokeAsync(new Tencentcloud.Cdn.GetDomainVerifierArgs
        ///         {
        ///             Domain = "www.examplexxx123.com",
        ///             AutoVerify = true,
        ///             FreezeRecord = true,
        ///         }));
        ///         var recordValue = data.Tencentcloud_cdn_domain_verifier.Record;
        ///         var recordType = data.Tencentcloud_cdn_domain_verifier.Record_type;
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDomainVerifierResult> Invoke(GetDomainVerifierInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainVerifierResult>("tencentcloud:Cdn/getDomainVerifier:getDomainVerifier", args ?? new GetDomainVerifierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainVerifierArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify whether to keep first create result instead of re-create again.
        /// </summary>
        [Input("autoVerify")]
        public bool? AutoVerify { get; set; }

        /// <summary>
        /// Specify domain name, e.g. `www.examplexxx123.com`.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// Indicates failed reason of verification.
        /// </summary>
        [Input("failedReason")]
        public string? FailedReason { get; set; }

        /// <summary>
        /// Specify whether the verification record needs to be freeze instead of refresh every 8 hours, this used for domain verification.
        /// </summary>
        [Input("freezeRecord")]
        public bool? FreezeRecord { get; set; }

        /// <summary>
        /// Used for save result json.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Specify verify type, values: `dns` (default), `file`.
        /// </summary>
        [Input("verifyType")]
        public string? VerifyType { get; set; }

        public GetDomainVerifierArgs()
        {
        }
    }

    public sealed class GetDomainVerifierInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify whether to keep first create result instead of re-create again.
        /// </summary>
        [Input("autoVerify")]
        public Input<bool>? AutoVerify { get; set; }

        /// <summary>
        /// Specify domain name, e.g. `www.examplexxx123.com`.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Indicates failed reason of verification.
        /// </summary>
        [Input("failedReason")]
        public Input<string>? FailedReason { get; set; }

        /// <summary>
        /// Specify whether the verification record needs to be freeze instead of refresh every 8 hours, this used for domain verification.
        /// </summary>
        [Input("freezeRecord")]
        public Input<bool>? FreezeRecord { get; set; }

        /// <summary>
        /// Used for save result json.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Specify verify type, values: `dns` (default), `file`.
        /// </summary>
        [Input("verifyType")]
        public Input<string>? VerifyType { get; set; }

        public GetDomainVerifierInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainVerifierResult
    {
        public readonly bool? AutoVerify;
        public readonly string Domain;
        public readonly string? FailedReason;
        /// <summary>
        /// List of file verified domains.
        /// </summary>
        public readonly ImmutableArray<string> FileVerifyDomains;
        /// <summary>
        /// Name of file verifications.
        /// </summary>
        public readonly string FileVerifyName;
        /// <summary>
        /// File verify URL guidance.
        /// </summary>
        public readonly string FileVerifyUrl;
        public readonly bool? FreezeRecord;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resolution record value.
        /// </summary>
        public readonly string Record;
        /// <summary>
        /// Type of resolution.
        /// </summary>
        public readonly string RecordType;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Sub-domain resolution.
        /// </summary>
        public readonly string SubDomain;
        /// <summary>
        /// Verify result.
        /// </summary>
        public readonly bool VerifyResult;
        public readonly string? VerifyType;

        [OutputConstructor]
        private GetDomainVerifierResult(
            bool? autoVerify,

            string domain,

            string? failedReason,

            ImmutableArray<string> fileVerifyDomains,

            string fileVerifyName,

            string fileVerifyUrl,

            bool? freezeRecord,

            string id,

            string record,

            string recordType,

            string? resultOutputFile,

            string subDomain,

            bool verifyResult,

            string? verifyType)
        {
            AutoVerify = autoVerify;
            Domain = domain;
            FailedReason = failedReason;
            FileVerifyDomains = fileVerifyDomains;
            FileVerifyName = fileVerifyName;
            FileVerifyUrl = fileVerifyUrl;
            FreezeRecord = freezeRecord;
            Id = id;
            Record = record;
            RecordType = recordType;
            ResultOutputFile = resultOutputFile;
            SubDomain = subDomain;
            VerifyResult = verifyResult;
            VerifyType = verifyType;
        }
    }
}
