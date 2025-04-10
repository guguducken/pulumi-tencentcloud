// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a monitor grafana ssoAccount
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
    ///         var ssoAccount = new Tencentcloud.Monitor.GrafanaSsoAccount("ssoAccount", new Tencentcloud.Monitor.GrafanaSsoAccountArgs
    ///         {
    ///             InstanceId = "grafana-50nj6v00",
    ///             Notes = "desc12222",
    ///             Roles = 
    ///             {
    ///                 new Tencentcloud.Monitor.Inputs.GrafanaSsoAccountRoleArgs
    ///                 {
    ///                     Organization = "Main Org.",
    ///                     Role = "Admin",
    ///                 },
    ///             },
    ///             UserId = "111",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// monitor grafana ssoAccount can be imported using the instance_id#user_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount ssoAccount grafana-50nj6v00#111
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount")]
    public partial class GrafanaSsoAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// grafana instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// account related description.
        /// </summary>
        [Output("notes")]
        public Output<string> Notes { get; private set; } = null!;

        /// <summary>
        /// grafana role.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<Outputs.GrafanaSsoAccountRole>> Roles { get; private set; } = null!;

        /// <summary>
        /// sub account uin of specific user.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a GrafanaSsoAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrafanaSsoAccount(string name, GrafanaSsoAccountArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount", name, args ?? new GrafanaSsoAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrafanaSsoAccount(string name, Input<string> id, GrafanaSsoAccountState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrafanaSsoAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrafanaSsoAccount Get(string name, Input<string> id, GrafanaSsoAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new GrafanaSsoAccount(name, id, state, options);
        }
    }

    public sealed class GrafanaSsoAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// grafana instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// account related description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("roles")]
        private InputList<Inputs.GrafanaSsoAccountRoleArgs>? _roles;

        /// <summary>
        /// grafana role.
        /// </summary>
        public InputList<Inputs.GrafanaSsoAccountRoleArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.GrafanaSsoAccountRoleArgs>());
            set => _roles = value;
        }

        /// <summary>
        /// sub account uin of specific user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GrafanaSsoAccountArgs()
        {
        }
    }

    public sealed class GrafanaSsoAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// grafana instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// account related description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("roles")]
        private InputList<Inputs.GrafanaSsoAccountRoleGetArgs>? _roles;

        /// <summary>
        /// grafana role.
        /// </summary>
        public InputList<Inputs.GrafanaSsoAccountRoleGetArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.GrafanaSsoAccountRoleGetArgs>());
            set => _roles = value;
        }

        /// <summary>
        /// sub account uin of specific user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GrafanaSsoAccountState()
        {
        }
    }
}
