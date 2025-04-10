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
    /// Provides a resource to create a monitor tmpCvmAgent
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
    ///         var tmpCvmAgent = new Tencentcloud.Monitor.TmpCvmAgent("tmpCvmAgent", new Tencentcloud.Monitor.TmpCvmAgentArgs
    ///         {
    ///             InstanceId = "prom-dko9d0nu",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// monitor tmpCvmAgent can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent tmpCvmAgent instance_id#agent_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent")]
    public partial class TmpCvmAgent : Pulumi.CustomResource
    {
        /// <summary>
        /// Agent id.
        /// </summary>
        [Output("agentId")]
        public Output<string> AgentId { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Agent name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a TmpCvmAgent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpCvmAgent(string name, TmpCvmAgentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent", name, args ?? new TmpCvmAgentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpCvmAgent(string name, Input<string> id, TmpCvmAgentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpCvmAgent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpCvmAgent Get(string name, Input<string> id, TmpCvmAgentState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpCvmAgent(name, id, state, options);
        }
    }

    public sealed class TmpCvmAgentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Agent name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TmpCvmAgentArgs()
        {
        }
    }

    public sealed class TmpCvmAgentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Agent id.
        /// </summary>
        [Input("agentId")]
        public Input<string>? AgentId { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Agent name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TmpCvmAgentState()
        {
        }
    }
}
