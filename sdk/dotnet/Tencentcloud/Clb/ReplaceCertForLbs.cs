// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb
{
    /// <summary>
    /// Provides a resource to create a clb replace_cert_for_lbs
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var replaceCertForLbs = new Tencentcloud.Clb.ReplaceCertForLbs("replaceCertForLbs", new Tencentcloud.Clb.ReplaceCertForLbsArgs
    ///         {
    ///             Certificate = new Tencentcloud.Clb.Inputs.ReplaceCertForLbsCertificateArgs
    ///             {
    ///                 CertCaContent = "XXXXX",
    ///                 CertCaName = "test",
    ///             },
    ///             OldCertificateId = "zjUMifFK",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Clb/replaceCertForLbs:ReplaceCertForLbs")]
    public partial class ReplaceCertForLbs : Pulumi.CustomResource
    {
        /// <summary>
        /// Information such as the content of the new certificate.
        /// </summary>
        [Output("certificate")]
        public Output<Outputs.ReplaceCertForLbsCertificate> Certificate { get; private set; } = null!;

        /// <summary>
        /// ID of the certificate to be replaced, which can be a server certificate or a client certificate.
        /// </summary>
        [Output("oldCertificateId")]
        public Output<string> OldCertificateId { get; private set; } = null!;


        /// <summary>
        /// Create a ReplaceCertForLbs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplaceCertForLbs(string name, ReplaceCertForLbsArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/replaceCertForLbs:ReplaceCertForLbs", name, args ?? new ReplaceCertForLbsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplaceCertForLbs(string name, Input<string> id, ReplaceCertForLbsState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/replaceCertForLbs:ReplaceCertForLbs", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReplaceCertForLbs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplaceCertForLbs Get(string name, Input<string> id, ReplaceCertForLbsState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplaceCertForLbs(name, id, state, options);
        }
    }

    public sealed class ReplaceCertForLbsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information such as the content of the new certificate.
        /// </summary>
        [Input("certificate", required: true)]
        public Input<Inputs.ReplaceCertForLbsCertificateArgs> Certificate { get; set; } = null!;

        /// <summary>
        /// ID of the certificate to be replaced, which can be a server certificate or a client certificate.
        /// </summary>
        [Input("oldCertificateId", required: true)]
        public Input<string> OldCertificateId { get; set; } = null!;

        public ReplaceCertForLbsArgs()
        {
        }
    }

    public sealed class ReplaceCertForLbsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information such as the content of the new certificate.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.ReplaceCertForLbsCertificateGetArgs>? Certificate { get; set; }

        /// <summary>
        /// ID of the certificate to be replaced, which can be a server certificate or a client certificate.
        /// </summary>
        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        public ReplaceCertForLbsState()
        {
        }
    }
}
