// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a rum offlineLogConfigAttachment
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const offlineLogConfigAttachment = new tencentcloud.Rum.OfflineLogConfigAttachment("offline_log_config_attachment", {
 *     projectKey: "ZEYrYfvaYQ30jRdmPx",
 *     uniqueId: "100027012454",
 * });
 * ```
 *
 * ## Import
 *
 * rum offline_log_config_attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Rum/offlineLogConfigAttachment:OfflineLogConfigAttachment offline_log_config_attachment ZEYrYfvaYQ30jRdmPx#100027012454
 * ```
 */
export class OfflineLogConfigAttachment extends pulumi.CustomResource {
    /**
     * Get an existing OfflineLogConfigAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OfflineLogConfigAttachmentState, opts?: pulumi.CustomResourceOptions): OfflineLogConfigAttachment {
        return new OfflineLogConfigAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Rum/offlineLogConfigAttachment:OfflineLogConfigAttachment';

    /**
     * Returns true if the given object is an instance of OfflineLogConfigAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OfflineLogConfigAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OfflineLogConfigAttachment.__pulumiType;
    }

    /**
     * Interface call information.
     */
    public /*out*/ readonly msg!: pulumi.Output<string>;
    /**
     * Unique project key for reporting.
     */
    public readonly projectKey!: pulumi.Output<string>;
    /**
     * Unique identifier of the user to be listened on(aid or uin).
     */
    public readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a OfflineLogConfigAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OfflineLogConfigAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OfflineLogConfigAttachmentArgs | OfflineLogConfigAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OfflineLogConfigAttachmentState | undefined;
            resourceInputs["msg"] = state ? state.msg : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["uniqueId"] = state ? state.uniqueId : undefined;
        } else {
            const args = argsOrState as OfflineLogConfigAttachmentArgs | undefined;
            if ((!args || args.projectKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectKey'");
            }
            if ((!args || args.uniqueId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uniqueId'");
            }
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["uniqueId"] = args ? args.uniqueId : undefined;
            resourceInputs["msg"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OfflineLogConfigAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OfflineLogConfigAttachment resources.
 */
export interface OfflineLogConfigAttachmentState {
    /**
     * Interface call information.
     */
    msg?: pulumi.Input<string>;
    /**
     * Unique project key for reporting.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Unique identifier of the user to be listened on(aid or uin).
     */
    uniqueId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OfflineLogConfigAttachment resource.
 */
export interface OfflineLogConfigAttachmentArgs {
    /**
     * Unique project key for reporting.
     */
    projectKey: pulumi.Input<string>;
    /**
     * Unique identifier of the user to be listened on(aid or uin).
     */
    uniqueId: pulumi.Input<string>;
}
