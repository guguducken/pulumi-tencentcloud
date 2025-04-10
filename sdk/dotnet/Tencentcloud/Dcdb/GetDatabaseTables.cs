// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    public static class GetDatabaseTables
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb database_tables
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
        ///         var databaseTables = Output.Create(Tencentcloud.Dcdb.GetDatabaseTables.InvokeAsync(new Tencentcloud.Dcdb.GetDatabaseTablesArgs
        ///         {
        ///             Cols = 
        ///             {
        ///                 
        ///                 {
        ///                     { "col", "" },
        ///                     { "gt", 
        ///                     {
        ///                         ,
        ///                         ,
        ///                     } },
        ///                     { "lt", 
        ///                     {
        ///                         ,
        ///                         ,
        ///                     } },
        ///                     { "nil", 
        ///                     {
        ///                         ,
        ///                         ,
        ///                     } },
        ///                     { "type", "" },
        ///                 },
        ///             },
        ///             DbName = "",
        ///             Gt = 
        ///             {
        ///                 ,
        ///                 ,
        ///                 ,
        ///             },
        ///             InstanceId = "dcdbt-ow7t8lmc",
        ///             Lt = 
        ///             {
        ///                 ,
        ///                 ,
        ///                 ,
        ///             },
        ///             Nil = 
        ///             {
        ///                 ,
        ///                 ,
        ///                 ,
        ///             },
        ///             Table = "",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseTablesResult> InvokeAsync(GetDatabaseTablesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseTablesResult>("tencentcloud:Dcdb/getDatabaseTables:getDatabaseTables", args ?? new GetDatabaseTablesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb database_tables
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
        ///         var databaseTables = Output.Create(Tencentcloud.Dcdb.GetDatabaseTables.InvokeAsync(new Tencentcloud.Dcdb.GetDatabaseTablesArgs
        ///         {
        ///             Cols = 
        ///             {
        ///                 
        ///                 {
        ///                     { "col", "" },
        ///                     { "gt", 
        ///                     {
        ///                         ,
        ///                         ,
        ///                     } },
        ///                     { "lt", 
        ///                     {
        ///                         ,
        ///                         ,
        ///                     } },
        ///                     { "nil", 
        ///                     {
        ///                         ,
        ///                         ,
        ///                     } },
        ///                     { "type", "" },
        ///                 },
        ///             },
        ///             DbName = "",
        ///             Gt = 
        ///             {
        ///                 ,
        ///                 ,
        ///                 ,
        ///             },
        ///             InstanceId = "dcdbt-ow7t8lmc",
        ///             Lt = 
        ///             {
        ///                 ,
        ///                 ,
        ///                 ,
        ///             },
        ///             Nil = 
        ///             {
        ///                 ,
        ///                 ,
        ///                 ,
        ///             },
        ///             Table = "",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabaseTablesResult> Invoke(GetDatabaseTablesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatabaseTablesResult>("tencentcloud:Dcdb/getDatabaseTables:getDatabaseTables", args ?? new GetDatabaseTablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseTablesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database name, obtained through the DescribeDatabases api.
        /// </summary>
        [Input("dbName", required: true)]
        public string DbName { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Table name, obtained through the DescribeDatabaseObjects api.
        /// </summary>
        [Input("table", required: true)]
        public string Table { get; set; } = null!;

        public GetDatabaseTablesArgs()
        {
        }
    }

    public sealed class GetDatabaseTablesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database name, obtained through the DescribeDatabases api.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Table name, obtained through the DescribeDatabaseObjects api.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public GetDatabaseTablesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseTablesResult
    {
        /// <summary>
        /// Column information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseTablesColResult> Cols;
        public readonly string DbName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly string Table;

        [OutputConstructor]
        private GetDatabaseTablesResult(
            ImmutableArray<Outputs.GetDatabaseTablesColResult> cols,

            string dbName,

            string id,

            string instanceId,

            string? resultOutputFile,

            string table)
        {
            Cols = cols;
            DbName = dbName;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            Table = table;
        }
    }
}
