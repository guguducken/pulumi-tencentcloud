// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tat
{
    public static class GetCommand
    {
        /// <summary>
        /// Use this data source to query detailed information of tat command
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
        ///         var command = Output.Create(Tencentcloud.Tat.GetCommand.InvokeAsync(new Tencentcloud.Tat.GetCommandArgs
        ///         {
        ///             CommandType = "SHELL",
        ///             CreatedBy = "TAT",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCommandResult> InvokeAsync(GetCommandArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCommandResult>("tencentcloud:Tat/getCommand:getCommand", args ?? new GetCommandArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tat command
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
        ///         var command = Output.Create(Tencentcloud.Tat.GetCommand.InvokeAsync(new Tencentcloud.Tat.GetCommandArgs
        ///         {
        ///             CommandType = "SHELL",
        ///             CreatedBy = "TAT",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCommandResult> Invoke(GetCommandInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCommandResult>("tencentcloud:Tat/getCommand:getCommand", args ?? new GetCommandInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCommandArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        [Input("commandId")]
        public string? CommandId { get; set; }

        /// <summary>
        /// Command name.
        /// </summary>
        [Input("commandName")]
        public string? CommandName { get; set; }

        /// <summary>
        /// Command type, Value is `SHELL` or `POWERSHELL`.
        /// </summary>
        [Input("commandType")]
        public string? CommandType { get; set; }

        /// <summary>
        /// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
        /// </summary>
        [Input("createdBy")]
        public string? CreatedBy { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetCommandArgs()
        {
        }
    }

    public sealed class GetCommandInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        [Input("commandId")]
        public Input<string>? CommandId { get; set; }

        /// <summary>
        /// Command name.
        /// </summary>
        [Input("commandName")]
        public Input<string>? CommandName { get; set; }

        /// <summary>
        /// Command type, Value is `SHELL` or `POWERSHELL`.
        /// </summary>
        [Input("commandType")]
        public Input<string>? CommandType { get; set; }

        /// <summary>
        /// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetCommandInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCommandResult
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        public readonly string? CommandId;
        /// <summary>
        /// Command name.
        /// </summary>
        public readonly string? CommandName;
        /// <summary>
        /// List of command details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCommandCommandSetResult> CommandSets;
        /// <summary>
        /// Command type.
        /// </summary>
        public readonly string? CommandType;
        /// <summary>
        /// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
        /// </summary>
        public readonly string? CreatedBy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetCommandResult(
            string? commandId,

            string? commandName,

            ImmutableArray<Outputs.GetCommandCommandSetResult> commandSets,

            string? commandType,

            string? createdBy,

            string id,

            string? resultOutputFile)
        {
            CommandId = commandId;
            CommandName = commandName;
            CommandSets = commandSets;
            CommandType = commandType;
            CreatedBy = createdBy;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
