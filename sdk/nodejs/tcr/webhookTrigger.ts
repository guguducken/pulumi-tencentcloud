// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tcr webhookTrigger
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mytcrWebhooktrigger = new tencentcloud.tcr.Instance("mytcrWebhooktrigger", {
 *     instanceType: "basic",
 *     deleteBucket: true,
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const myNs = new tencentcloud.tcr.Namespace("myNs", {
 *     instanceId: mytcrWebhooktrigger.id,
 *     isPublic: true,
 *     isAutoScan: true,
 *     isPreventVul: true,
 *     severity: "medium",
 *     cveWhitelistItems: [{
 *         cveId: "cve-xxxxx",
 *     }],
 * });
 * const idTest = tencentcloud.Tcr.getNamespacesOutput({
 *     instanceId: myNs.instanceId,
 * });
 * const nsId = idTest.apply(idTest => idTest.namespaceLists?[0]?.id);
 * const myTrigger = new tencentcloud.tcr.WebhookTrigger("myTrigger", {
 *     registryId: mytcrWebhooktrigger.id,
 *     namespace: myNs.name,
 *     trigger: {
 *         name: `trigger-%s`,
 *         targets: [{
 *             address: "http://example.org/post",
 *             headers: [{
 *                 key: "X-Custom-Header",
 *                 values: ["a"],
 *             }],
 *         }],
 *         eventTypes: ["pushImage"],
 *         condition: ".*",
 *         enabled: true,
 *         description: "this is trigger description",
 *         namespaceId: nsId,
 *     },
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * tcr webhook_trigger can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tcr/webhookTrigger:WebhookTrigger webhook_trigger webhook_trigger_id
 * ```
 */
export class WebhookTrigger extends pulumi.CustomResource {
    /**
     * Get an existing WebhookTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookTriggerState, opts?: pulumi.CustomResourceOptions): WebhookTrigger {
        return new WebhookTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcr/webhookTrigger:WebhookTrigger';

    /**
     * Returns true if the given object is an instance of WebhookTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebhookTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebhookTrigger.__pulumiType;
    }

    /**
     * namespace name.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * instance Id.
     */
    public readonly registryId!: pulumi.Output<string>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * trigger parameters.
     */
    public readonly trigger!: pulumi.Output<outputs.Tcr.WebhookTriggerTrigger>;

    /**
     * Create a WebhookTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookTriggerArgs | WebhookTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebhookTriggerState | undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["trigger"] = state ? state.trigger : undefined;
        } else {
            const args = argsOrState as WebhookTriggerArgs | undefined;
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.registryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryId'");
            }
            if ((!args || args.trigger === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trigger'");
            }
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["registryId"] = args ? args.registryId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trigger"] = args ? args.trigger : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebhookTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebhookTrigger resources.
 */
export interface WebhookTriggerState {
    /**
     * namespace name.
     */
    namespace?: pulumi.Input<string>;
    /**
     * instance Id.
     */
    registryId?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * trigger parameters.
     */
    trigger?: pulumi.Input<inputs.Tcr.WebhookTriggerTrigger>;
}

/**
 * The set of arguments for constructing a WebhookTrigger resource.
 */
export interface WebhookTriggerArgs {
    /**
     * namespace name.
     */
    namespace: pulumi.Input<string>;
    /**
     * instance Id.
     */
    registryId: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * trigger parameters.
     */
    trigger: pulumi.Input<inputs.Tcr.WebhookTriggerTrigger>;
}
