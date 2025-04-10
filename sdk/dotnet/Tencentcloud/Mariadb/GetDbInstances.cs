// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    public static class GetDbInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of mariadb dbInstances
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
        ///         var dbInstances = Output.Create(Tencentcloud.Mariadb.GetDbInstances.InvokeAsync(new Tencentcloud.Mariadb.GetDbInstancesArgs
        ///         {
        ///             InstanceIds = 
        ///             {
        ///                 "tdsql-ijxtqk5p",
        ///             },
        ///             ProjectIds = 
        ///             {
        ///                 0,
        ///             },
        ///             SubnetId = "3454730",
        ///             VpcId = "5556791",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDbInstancesResult> InvokeAsync(GetDbInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbInstancesResult>("tencentcloud:Mariadb/getDbInstances:getDbInstances", args ?? new GetDbInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mariadb dbInstances
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
        ///         var dbInstances = Output.Create(Tencentcloud.Mariadb.GetDbInstances.InvokeAsync(new Tencentcloud.Mariadb.GetDbInstancesArgs
        ///         {
        ///             InstanceIds = 
        ///             {
        ///                 "tdsql-ijxtqk5p",
        ///             },
        ///             ProjectIds = 
        ///             {
        ///                 0,
        ///             },
        ///             SubnetId = "3454730",
        ///             VpcId = "5556791",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDbInstancesResult> Invoke(GetDbInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDbInstancesResult>("tencentcloud:Mariadb/getDbInstances:getDbInstances", args ?? new GetDbInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDbInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// instance ids.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        [Input("projectIds")]
        private List<int>? _projectIds;

        /// <summary>
        /// project ids.
        /// </summary>
        public List<int> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new List<int>());
            set => _projectIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// instance name or vip.
        /// </summary>
        [Input("searchName")]
        public string? SearchName { get; set; }

        /// <summary>
        /// subnet id.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        /// <summary>
        /// vpc id.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetDbInstancesArgs()
        {
        }
    }

    public sealed class GetDbInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// instance ids.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        [Input("projectIds")]
        private InputList<int>? _projectIds;

        /// <summary>
        /// project ids.
        /// </summary>
        public InputList<int> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<int>());
            set => _projectIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// instance name or vip.
        /// </summary>
        [Input("searchName")]
        public Input<string>? SearchName { get; set; }

        /// <summary>
        /// subnet id.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// vpc id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetDbInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbInstancesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        /// <summary>
        /// instances info.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDbInstancesInstanceResult> Instances;
        public readonly ImmutableArray<int> ProjectIds;
        public readonly string? ResultOutputFile;
        public readonly string? SearchName;
        /// <summary>
        /// subnet id.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// vpc id.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private GetDbInstancesResult(
            string id,

            ImmutableArray<string> instanceIds,

            ImmutableArray<Outputs.GetDbInstancesInstanceResult> instances,

            ImmutableArray<int> projectIds,

            string? resultOutputFile,

            string? searchName,

            string? subnetId,

            string? vpcId)
        {
            Id = id;
            InstanceIds = instanceIds;
            Instances = instances;
            ProjectIds = projectIds;
            ResultOutputFile = resultOutputFile;
            SearchName = searchName;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
