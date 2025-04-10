// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a chdfs fileSystem
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const fileSystem = new tencentcloud.Chdfs.FileSystem("file_system", {
 *     capacityQuota: 10995116277760,
 *     description: "file system for terraform test",
 *     enableRanger: true,
 *     fileSystemName: "terraform-test",
 *     posixAcl: false,
 *     rangerServiceAddresses: [
 *         "127.0.0.1:80",
 *         "127.0.0.1:8000",
 *     ],
 *     superUsers: [
 *         "terraform",
 *         "iac",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * chdfs file_system can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Chdfs/fileSystem:FileSystem file_system file_system_id
 * ```
 */
export class FileSystem extends pulumi.CustomResource {
    /**
     * Get an existing FileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileSystemState, opts?: pulumi.CustomResourceOptions): FileSystem {
        return new FileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Chdfs/fileSystem:FileSystem';

    /**
     * Returns true if the given object is an instance of FileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileSystem.__pulumiType;
    }

    /**
     * file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
     */
    public readonly capacityQuota!: pulumi.Output<number>;
    /**
     * desc of the file system.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * check the ranger address or not.
     */
    public readonly enableRanger!: pulumi.Output<boolean | undefined>;
    /**
     * file system name.
     */
    public readonly fileSystemName!: pulumi.Output<string>;
    /**
     * check POSIX ACL or not.
     */
    public readonly posixAcl!: pulumi.Output<boolean>;
    /**
     * ranger address list, default empty.
     */
    public readonly rangerServiceAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * super users of the file system, default empty.
     */
    public readonly superUsers!: pulumi.Output<string[] | undefined>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileSystemArgs | FileSystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FileSystemState | undefined;
            resourceInputs["capacityQuota"] = state ? state.capacityQuota : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableRanger"] = state ? state.enableRanger : undefined;
            resourceInputs["fileSystemName"] = state ? state.fileSystemName : undefined;
            resourceInputs["posixAcl"] = state ? state.posixAcl : undefined;
            resourceInputs["rangerServiceAddresses"] = state ? state.rangerServiceAddresses : undefined;
            resourceInputs["superUsers"] = state ? state.superUsers : undefined;
        } else {
            const args = argsOrState as FileSystemArgs | undefined;
            if ((!args || args.capacityQuota === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacityQuota'");
            }
            if ((!args || args.fileSystemName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemName'");
            }
            if ((!args || args.posixAcl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'posixAcl'");
            }
            resourceInputs["capacityQuota"] = args ? args.capacityQuota : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableRanger"] = args ? args.enableRanger : undefined;
            resourceInputs["fileSystemName"] = args ? args.fileSystemName : undefined;
            resourceInputs["posixAcl"] = args ? args.posixAcl : undefined;
            resourceInputs["rangerServiceAddresses"] = args ? args.rangerServiceAddresses : undefined;
            resourceInputs["superUsers"] = args ? args.superUsers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FileSystem resources.
 */
export interface FileSystemState {
    /**
     * file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
     */
    capacityQuota?: pulumi.Input<number>;
    /**
     * desc of the file system.
     */
    description?: pulumi.Input<string>;
    /**
     * check the ranger address or not.
     */
    enableRanger?: pulumi.Input<boolean>;
    /**
     * file system name.
     */
    fileSystemName?: pulumi.Input<string>;
    /**
     * check POSIX ACL or not.
     */
    posixAcl?: pulumi.Input<boolean>;
    /**
     * ranger address list, default empty.
     */
    rangerServiceAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * super users of the file system, default empty.
     */
    superUsers?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    /**
     * file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
     */
    capacityQuota: pulumi.Input<number>;
    /**
     * desc of the file system.
     */
    description?: pulumi.Input<string>;
    /**
     * check the ranger address or not.
     */
    enableRanger?: pulumi.Input<boolean>;
    /**
     * file system name.
     */
    fileSystemName: pulumi.Input<string>;
    /**
     * check POSIX ACL or not.
     */
    posixAcl: pulumi.Input<boolean>;
    /**
     * ranger address list, default empty.
     */
    rangerServiceAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * super users of the file system, default empty.
     */
    superUsers?: pulumi.Input<pulumi.Input<string>[]>;
}
