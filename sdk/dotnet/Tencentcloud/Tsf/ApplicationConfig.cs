// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    /// <summary>
    /// Provides a resource to create a tsf application_config
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
    ///         var applicationConfig = new Tencentcloud.Tsf.ApplicationConfig("applicationConfig", new Tencentcloud.Tsf.ApplicationConfigArgs
    ///         {
    ///             ApplicationId = "application-ym9mxmza",
    ///             ConfigName = "test-2",
    ///             ConfigValue = "name: \"name\"",
    ///             ConfigVersion = "1.0",
    ///             ConfigVersionDesc = "test2",
    ///             EncodeWithBase64 = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/applicationConfig:ApplicationConfig")]
    public partial class ApplicationConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// configuration item name.
        /// </summary>
        [Output("configName")]
        public Output<string> ConfigName { get; private set; } = null!;

        /// <summary>
        /// configuration item value type.
        /// </summary>
        [Output("configType")]
        public Output<string?> ConfigType { get; private set; } = null!;

        /// <summary>
        /// configuration item value.
        /// </summary>
        [Output("configValue")]
        public Output<string> ConfigValue { get; private set; } = null!;

        /// <summary>
        /// configuration item version.
        /// </summary>
        [Output("configVersion")]
        public Output<string> ConfigVersion { get; private set; } = null!;

        /// <summary>
        /// configuration item version description.
        /// </summary>
        [Output("configVersionDesc")]
        public Output<string?> ConfigVersionDesc { get; private set; } = null!;

        /// <summary>
        /// Base64 encoded configuration items.
        /// </summary>
        [Output("encodeWithBase64")]
        public Output<bool?> EncodeWithBase64 { get; private set; } = null!;

        /// <summary>
        /// Program id list.
        /// </summary>
        [Output("programIdLists")]
        public Output<ImmutableArray<string>> ProgramIdLists { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationConfig(string name, ApplicationConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/applicationConfig:ApplicationConfig", name, args ?? new ApplicationConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationConfig(string name, Input<string> id, ApplicationConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/applicationConfig:ApplicationConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationConfig Get(string name, Input<string> id, ApplicationConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationConfig(name, id, state, options);
        }
    }

    public sealed class ApplicationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// configuration item name.
        /// </summary>
        [Input("configName", required: true)]
        public Input<string> ConfigName { get; set; } = null!;

        /// <summary>
        /// configuration item value type.
        /// </summary>
        [Input("configType")]
        public Input<string>? ConfigType { get; set; }

        /// <summary>
        /// configuration item value.
        /// </summary>
        [Input("configValue", required: true)]
        public Input<string> ConfigValue { get; set; } = null!;

        /// <summary>
        /// configuration item version.
        /// </summary>
        [Input("configVersion", required: true)]
        public Input<string> ConfigVersion { get; set; } = null!;

        /// <summary>
        /// configuration item version description.
        /// </summary>
        [Input("configVersionDesc")]
        public Input<string>? ConfigVersionDesc { get; set; }

        /// <summary>
        /// Base64 encoded configuration items.
        /// </summary>
        [Input("encodeWithBase64")]
        public Input<bool>? EncodeWithBase64 { get; set; }

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// Program id list.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        public ApplicationConfigArgs()
        {
        }
    }

    public sealed class ApplicationConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// configuration item name.
        /// </summary>
        [Input("configName")]
        public Input<string>? ConfigName { get; set; }

        /// <summary>
        /// configuration item value type.
        /// </summary>
        [Input("configType")]
        public Input<string>? ConfigType { get; set; }

        /// <summary>
        /// configuration item value.
        /// </summary>
        [Input("configValue")]
        public Input<string>? ConfigValue { get; set; }

        /// <summary>
        /// configuration item version.
        /// </summary>
        [Input("configVersion")]
        public Input<string>? ConfigVersion { get; set; }

        /// <summary>
        /// configuration item version description.
        /// </summary>
        [Input("configVersionDesc")]
        public Input<string>? ConfigVersionDesc { get; set; }

        /// <summary>
        /// Base64 encoded configuration items.
        /// </summary>
        [Input("encodeWithBase64")]
        public Input<bool>? EncodeWithBase64 { get; set; }

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// Program id list.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        public ApplicationConfigState()
        {
        }
    }
}
