// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver configDatabaseMdf
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const configDatabaseMdf = new tencentcloud.Sqlserver.ConfigDatabaseMdf("config_database_mdf", {
 *     dbName: "keep_pubsub_db2",
 *     instanceId: "mssql-qelbzgwf",
 * });
 * ```
 *
 * ## Import
 *
 * sqlserver config_database_mdf can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf config_database_mdf config_database_mdf_id
 * ```
 */
export class ConfigDatabaseMdf extends pulumi.CustomResource {
    /**
     * Get an existing ConfigDatabaseMdf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigDatabaseMdfState, opts?: pulumi.CustomResourceOptions): ConfigDatabaseMdf {
        return new ConfigDatabaseMdf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf';

    /**
     * Returns true if the given object is an instance of ConfigDatabaseMdf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigDatabaseMdf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigDatabaseMdf.__pulumiType;
    }

    /**
     * Array of database names.
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a ConfigDatabaseMdf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigDatabaseMdfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigDatabaseMdfArgs | ConfigDatabaseMdfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigDatabaseMdfState | undefined;
            resourceInputs["dbName"] = state ? state.dbName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as ConfigDatabaseMdfArgs | undefined;
            if ((!args || args.dbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["dbName"] = args ? args.dbName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigDatabaseMdf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigDatabaseMdf resources.
 */
export interface ConfigDatabaseMdfState {
    /**
     * Array of database names.
     */
    dbName?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigDatabaseMdf resource.
 */
export interface ConfigDatabaseMdfArgs {
    /**
     * Array of database names.
     */
    dbName: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
}
