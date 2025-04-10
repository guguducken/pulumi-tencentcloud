// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskSourceResourcePostgreSqlParamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Upstream data format (JSON|Debezium), required when the database synchronization mode matches the default field.
        /// </summary>
        [Input("dataFormat")]
        public Input<string>? DataFormat { get; set; }

        /// <summary>
        /// INSERT means insert using Insert mode, UPSERT means insert using Upsert mode.
        /// </summary>
        [Input("dataTargetInsertMode")]
        public Input<string>? DataTargetInsertMode { get; set; }

        /// <summary>
        /// When DataInsertMode=UPSERT, pass in the primary key that the current upsert depends on.
        /// </summary>
        [Input("dataTargetPrimaryKeyField")]
        public Input<string>? DataTargetPrimaryKeyField { get; set; }

        [Input("dataTargetRecordMappings")]
        private InputList<Inputs.DatahubTaskSourceResourcePostgreSqlParamDataTargetRecordMappingArgs>? _dataTargetRecordMappings;

        /// <summary>
        /// Mapping relationship between tables and messages.
        /// </summary>
        public InputList<Inputs.DatahubTaskSourceResourcePostgreSqlParamDataTargetRecordMappingArgs> DataTargetRecordMappings
        {
            get => _dataTargetRecordMappings ?? (_dataTargetRecordMappings = new InputList<Inputs.DatahubTaskSourceResourcePostgreSqlParamDataTargetRecordMappingArgs>());
            set => _dataTargetRecordMappings = value;
        }

        /// <summary>
        /// PostgreSQL database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Whether to discard messages that fail to parse, the default is true.
        /// </summary>
        [Input("dropInvalidMessage")]
        public Input<bool>? DropInvalidMessage { get; set; }

        /// <summary>
        /// Whether the input table is a regular expression.
        /// </summary>
        [Input("isTableRegular")]
        public Input<bool>? IsTableRegular { get; set; }

        /// <summary>
        /// Format  library1.table1:field 1,field2;library2.table2:field2, between tables; (semicolon) separated, between fields, (comma) separated. The table that is not specified defaults to the primary key of the table.
        /// </summary>
        [Input("keyColumns")]
        public Input<string>? KeyColumns { get; set; }

        /// <summary>
        /// (decoderbufs/pgoutput), default decoderbufs.
        /// </summary>
        [Input("pluginName", required: true)]
        public Input<string> PluginName { get; set; } = null!;

        /// <summary>
        /// If the value is true, the message will carry the schema corresponding to the message structure, if the value is false, it will not carry.
        /// </summary>
        [Input("recordWithSchema")]
        public Input<bool>? RecordWithSchema { get; set; }

        /// <summary>
        /// PostgreSQL connection Id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// never|initial, default initial.
        /// </summary>
        [Input("snapshotMode")]
        public Input<string>? SnapshotMode { get; set; }

        /// <summary>
        /// PostgreSQL tableName, * is the non-system table in all the monitored databases, you can use, to monitor multiple data tables, but the data table needs to be filled in the format of Schema name.Data table name, and you need to fill in a regular expression When, the format is Schema name.data table name.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public DatahubTaskSourceResourcePostgreSqlParamArgs()
        {
        }
    }
}
