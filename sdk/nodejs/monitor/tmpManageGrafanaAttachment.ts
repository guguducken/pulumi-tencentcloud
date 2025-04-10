// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a monitor tmpManageGrafanaAttachment
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const manageGrafanaAttachment = new tencentcloud.Monitor.TmpManageGrafanaAttachment("manage_grafana_attachment", {
 *     grafanaId: "grafana-xxxxxx",
 *     instanceId: "prom-xxxxxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * monitor tmp_manage_grafana_attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Monitor/tmpManageGrafanaAttachment:TmpManageGrafanaAttachment manage_grafana_attachment prom-xxxxxxxx
 * ```
 */
export class TmpManageGrafanaAttachment extends pulumi.CustomResource {
    /**
     * Get an existing TmpManageGrafanaAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TmpManageGrafanaAttachmentState, opts?: pulumi.CustomResourceOptions): TmpManageGrafanaAttachment {
        return new TmpManageGrafanaAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/tmpManageGrafanaAttachment:TmpManageGrafanaAttachment';

    /**
     * Returns true if the given object is an instance of TmpManageGrafanaAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TmpManageGrafanaAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TmpManageGrafanaAttachment.__pulumiType;
    }

    /**
     * Grafana instance ID.
     */
    public readonly grafanaId!: pulumi.Output<string>;
    /**
     * Prometheus instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a TmpManageGrafanaAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TmpManageGrafanaAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TmpManageGrafanaAttachmentArgs | TmpManageGrafanaAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TmpManageGrafanaAttachmentState | undefined;
            resourceInputs["grafanaId"] = state ? state.grafanaId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as TmpManageGrafanaAttachmentArgs | undefined;
            if ((!args || args.grafanaId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'grafanaId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["grafanaId"] = args ? args.grafanaId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TmpManageGrafanaAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TmpManageGrafanaAttachment resources.
 */
export interface TmpManageGrafanaAttachmentState {
    /**
     * Grafana instance ID.
     */
    grafanaId?: pulumi.Input<string>;
    /**
     * Prometheus instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TmpManageGrafanaAttachment resource.
 */
export interface TmpManageGrafanaAttachmentArgs {
    /**
     * Grafana instance ID.
     */
    grafanaId: pulumi.Input<string>;
    /**
     * Prometheus instance ID.
     */
    instanceId: pulumi.Input<string>;
}
