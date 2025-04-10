// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci
{
    /// <summary>
    /// Provides a resource to create a ci media_snapshot_template
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
    ///         var mediaSnapshotTemplate = new Tencentcloud.Ci.MediaSnapshotTemplate("mediaSnapshotTemplate", new Tencentcloud.Ci.MediaSnapshotTemplateArgs
    ///         {
    ///             Bucket = "terraform-ci-xxxxxx",
    ///             Snapshot = new Tencentcloud.Ci.Inputs.MediaSnapshotTemplateSnapshotArgs
    ///             {
    ///                 Count = "10",
    ///                 SnapshotOutMode = "SnapshotAndSprite",
    ///                 SpriteSnapshotConfig = new Tencentcloud.Ci.Inputs.MediaSnapshotTemplateSnapshotSpriteSnapshotConfigArgs
    ///                 {
    ///                     Color = "White",
    ///                     Columns = "10",
    ///                     Lines = "10",
    ///                     Margin = "10",
    ///                     Padding = "10",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ci media_snapshot_template can be imported using the bucket#templateId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate media_snapshot_template terraform-ci-xxxxxx#t18210645f96564eaf80e86b1f58c20152
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate")]
    public partial class MediaSnapshotTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// screenshot.
        /// </summary>
        [Output("snapshot")]
        public Output<Outputs.MediaSnapshotTemplateSnapshot> Snapshot { get; private set; } = null!;

        /// <summary>
        /// Template ID.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// update time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a MediaSnapshotTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaSnapshotTemplate(string name, MediaSnapshotTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate", name, args ?? new MediaSnapshotTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaSnapshotTemplate(string name, Input<string> id, MediaSnapshotTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaSnapshotTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaSnapshotTemplate Get(string name, Input<string> id, MediaSnapshotTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaSnapshotTemplate(name, id, state, options);
        }
    }

    public sealed class MediaSnapshotTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// screenshot.
        /// </summary>
        [Input("snapshot", required: true)]
        public Input<Inputs.MediaSnapshotTemplateSnapshotArgs> Snapshot { get; set; } = null!;

        public MediaSnapshotTemplateArgs()
        {
        }
    }

    public sealed class MediaSnapshotTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// screenshot.
        /// </summary>
        [Input("snapshot")]
        public Input<Inputs.MediaSnapshotTemplateSnapshotGetArgs>? Snapshot { get; set; }

        /// <summary>
        /// Template ID.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// update time.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public MediaSnapshotTemplateState()
        {
        }
    }
}
