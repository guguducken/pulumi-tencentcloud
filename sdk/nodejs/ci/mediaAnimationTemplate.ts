// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ci mediaAnimationTemplate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mediaAnimationTemplate = new tencentcloud.Ci.MediaAnimationTemplate("media_animation_template", {
 *     bucket: "terraform-ci-1308919341",
 *     container: {
 *         format: "gif",
 *     },
 *     timeInterval: {
 *         duration: "60",
 *         start: "0",
 *     },
 *     video: {
 *         animateFramesPerSecond: "",
 *         animateOnlyKeepKeyFrame: "true",
 *         animateTimeIntervalOfFrame: "",
 *         codec: "gif",
 *         fps: "20",
 *         height: "",
 *         quality: "",
 *         width: "1280",
 *     },
 * });
 * ```
 */
export class MediaAnimationTemplate extends pulumi.CustomResource {
    /**
     * Get an existing MediaAnimationTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MediaAnimationTemplateState, opts?: pulumi.CustomResourceOptions): MediaAnimationTemplate {
        return new MediaAnimationTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ci/mediaAnimationTemplate:MediaAnimationTemplate';

    /**
     * Returns true if the given object is an instance of MediaAnimationTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MediaAnimationTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MediaAnimationTemplate.__pulumiType;
    }

    /**
     * bucket name.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * container format.
     */
    public readonly container!: pulumi.Output<outputs.Ci.MediaAnimationTemplateContainer>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * time interval.
     */
    public readonly timeInterval!: pulumi.Output<outputs.Ci.MediaAnimationTemplateTimeInterval | undefined>;
    /**
     * video information, do not upload Video, which is equivalent to deleting video information.
     */
    public readonly video!: pulumi.Output<outputs.Ci.MediaAnimationTemplateVideo | undefined>;

    /**
     * Create a MediaAnimationTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MediaAnimationTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MediaAnimationTemplateArgs | MediaAnimationTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MediaAnimationTemplateState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["container"] = state ? state.container : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["timeInterval"] = state ? state.timeInterval : undefined;
            resourceInputs["video"] = state ? state.video : undefined;
        } else {
            const args = argsOrState as MediaAnimationTemplateArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.container === undefined) && !opts.urn) {
                throw new Error("Missing required property 'container'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["container"] = args ? args.container : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["timeInterval"] = args ? args.timeInterval : undefined;
            resourceInputs["video"] = args ? args.video : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MediaAnimationTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MediaAnimationTemplate resources.
 */
export interface MediaAnimationTemplateState {
    /**
     * bucket name.
     */
    bucket?: pulumi.Input<string>;
    /**
     * container format.
     */
    container?: pulumi.Input<inputs.Ci.MediaAnimationTemplateContainer>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    name?: pulumi.Input<string>;
    /**
     * time interval.
     */
    timeInterval?: pulumi.Input<inputs.Ci.MediaAnimationTemplateTimeInterval>;
    /**
     * video information, do not upload Video, which is equivalent to deleting video information.
     */
    video?: pulumi.Input<inputs.Ci.MediaAnimationTemplateVideo>;
}

/**
 * The set of arguments for constructing a MediaAnimationTemplate resource.
 */
export interface MediaAnimationTemplateArgs {
    /**
     * bucket name.
     */
    bucket: pulumi.Input<string>;
    /**
     * container format.
     */
    container: pulumi.Input<inputs.Ci.MediaAnimationTemplateContainer>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    name?: pulumi.Input<string>;
    /**
     * time interval.
     */
    timeInterval?: pulumi.Input<inputs.Ci.MediaAnimationTemplateTimeInterval>;
    /**
     * video information, do not upload Video, which is equivalent to deleting video information.
     */
    video?: pulumi.Input<inputs.Ci.MediaAnimationTemplateVideo>;
}
