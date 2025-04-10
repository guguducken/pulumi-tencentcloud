// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts
{
    /// <summary>
    /// Provides a resource to create a pts file
    /// 
    /// &gt; **NOTE:** Modification is not currently supported, please go to the console to modify.
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
    ///         var file = new Tencentcloud.Pts.File("file", new Tencentcloud.Pts.FileArgs
    ///         {
    ///             FileId = "file-de2dbaf8",
    ///             HeaderInFile = false,
    ///             Kind = 3,
    ///             LineCount = 0,
    ///             ProjectId = "project-45vw7v82",
    ///             Size = 10799,
    ///             Type = "text/plain",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// pts file can be imported using the project_id#file_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Pts/file:File file project-45vw7v82#file-de2dbaf8
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Pts/file:File")]
    public partial class File : Pulumi.CustomResource
    {
        /// <summary>
        /// File id.
        /// </summary>
        [Output("fileId")]
        public Output<string> FileId { get; private set; } = null!;

        /// <summary>
        /// Files in a folder.
        /// </summary>
        [Output("fileInfos")]
        public Output<ImmutableArray<Outputs.FileFileInfo>> FileInfos { get; private set; } = null!;

        /// <summary>
        /// The first few lines of data.
        /// </summary>
        [Output("headLines")]
        public Output<ImmutableArray<string>> HeadLines { get; private set; } = null!;

        /// <summary>
        /// Meter head.
        /// </summary>
        [Output("headerColumns")]
        public Output<ImmutableArray<string>> HeaderColumns { get; private set; } = null!;

        /// <summary>
        /// Whether the header is in the file.
        /// </summary>
        [Output("headerInFile")]
        public Output<bool?> HeaderInFile { get; private set; } = null!;

        /// <summary>
        /// File kind, parameter file-1, protocol file-2, request file-3.
        /// </summary>
        [Output("kind")]
        public Output<int> Kind { get; private set; } = null!;

        /// <summary>
        /// Line count.
        /// </summary>
        [Output("lineCount")]
        public Output<int?> LineCount { get; private set; } = null!;

        /// <summary>
        /// File name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Project id.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// File size.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The last few lines of data.
        /// </summary>
        [Output("tailLines")]
        public Output<ImmutableArray<string>> TailLines { get; private set; } = null!;

        /// <summary>
        /// File type, folder-folder.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a File resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public File(string name, FileArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/file:File", name, args ?? new FileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private File(string name, Input<string> id, FileState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/file:File", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing File resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static File Get(string name, Input<string> id, FileState? state = null, CustomResourceOptions? options = null)
        {
            return new File(name, id, state, options);
        }
    }

    public sealed class FileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File id.
        /// </summary>
        [Input("fileId", required: true)]
        public Input<string> FileId { get; set; } = null!;

        [Input("fileInfos")]
        private InputList<Inputs.FileFileInfoArgs>? _fileInfos;

        /// <summary>
        /// Files in a folder.
        /// </summary>
        public InputList<Inputs.FileFileInfoArgs> FileInfos
        {
            get => _fileInfos ?? (_fileInfos = new InputList<Inputs.FileFileInfoArgs>());
            set => _fileInfos = value;
        }

        [Input("headLines")]
        private InputList<string>? _headLines;

        /// <summary>
        /// The first few lines of data.
        /// </summary>
        public InputList<string> HeadLines
        {
            get => _headLines ?? (_headLines = new InputList<string>());
            set => _headLines = value;
        }

        [Input("headerColumns")]
        private InputList<string>? _headerColumns;

        /// <summary>
        /// Meter head.
        /// </summary>
        public InputList<string> HeaderColumns
        {
            get => _headerColumns ?? (_headerColumns = new InputList<string>());
            set => _headerColumns = value;
        }

        /// <summary>
        /// Whether the header is in the file.
        /// </summary>
        [Input("headerInFile")]
        public Input<bool>? HeaderInFile { get; set; }

        /// <summary>
        /// File kind, parameter file-1, protocol file-2, request file-3.
        /// </summary>
        [Input("kind", required: true)]
        public Input<int> Kind { get; set; } = null!;

        /// <summary>
        /// Line count.
        /// </summary>
        [Input("lineCount")]
        public Input<int>? LineCount { get; set; }

        /// <summary>
        /// File name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project id.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// File size.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("tailLines")]
        private InputList<string>? _tailLines;

        /// <summary>
        /// The last few lines of data.
        /// </summary>
        public InputList<string> TailLines
        {
            get => _tailLines ?? (_tailLines = new InputList<string>());
            set => _tailLines = value;
        }

        /// <summary>
        /// File type, folder-folder.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FileArgs()
        {
        }
    }

    public sealed class FileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File id.
        /// </summary>
        [Input("fileId")]
        public Input<string>? FileId { get; set; }

        [Input("fileInfos")]
        private InputList<Inputs.FileFileInfoGetArgs>? _fileInfos;

        /// <summary>
        /// Files in a folder.
        /// </summary>
        public InputList<Inputs.FileFileInfoGetArgs> FileInfos
        {
            get => _fileInfos ?? (_fileInfos = new InputList<Inputs.FileFileInfoGetArgs>());
            set => _fileInfos = value;
        }

        [Input("headLines")]
        private InputList<string>? _headLines;

        /// <summary>
        /// The first few lines of data.
        /// </summary>
        public InputList<string> HeadLines
        {
            get => _headLines ?? (_headLines = new InputList<string>());
            set => _headLines = value;
        }

        [Input("headerColumns")]
        private InputList<string>? _headerColumns;

        /// <summary>
        /// Meter head.
        /// </summary>
        public InputList<string> HeaderColumns
        {
            get => _headerColumns ?? (_headerColumns = new InputList<string>());
            set => _headerColumns = value;
        }

        /// <summary>
        /// Whether the header is in the file.
        /// </summary>
        [Input("headerInFile")]
        public Input<bool>? HeaderInFile { get; set; }

        /// <summary>
        /// File kind, parameter file-1, protocol file-2, request file-3.
        /// </summary>
        [Input("kind")]
        public Input<int>? Kind { get; set; }

        /// <summary>
        /// Line count.
        /// </summary>
        [Input("lineCount")]
        public Input<int>? LineCount { get; set; }

        /// <summary>
        /// File name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project id.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// File size.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("tailLines")]
        private InputList<string>? _tailLines;

        /// <summary>
        /// The last few lines of data.
        /// </summary>
        public InputList<string> TailLines
        {
            get => _tailLines ?? (_tailLines = new InputList<string>());
            set => _tailLines = value;
        }

        /// <summary>
        /// File type, folder-folder.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FileState()
        {
        }
    }
}
