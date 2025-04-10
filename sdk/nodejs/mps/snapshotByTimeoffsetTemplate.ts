// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mps snapshotByTimeoffsetTemplate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const snapshotByTimeoffsetTemplate = new tencentcloud.Mps.SnapshotByTimeoffsetTemplate("snapshot_by_timeoffset_template", {
 *     fillType: "stretch",
 *     format: "jpg",
 *     height: 128,
 *     resolutionAdaptive: "open",
 *     width: 140,
 * });
 * ```
 *
 * ## Import
 *
 * mps snapshot_by_timeoffset_template can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate snapshot_by_timeoffset_template snapshot_by_timeoffset_template_id
 * ```
 */
export class SnapshotByTimeoffsetTemplate extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotByTimeoffsetTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotByTimeoffsetTemplateState, opts?: pulumi.CustomResourceOptions): SnapshotByTimeoffsetTemplate {
        return new SnapshotByTimeoffsetTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate';

    /**
     * Returns true if the given object is an instance of SnapshotByTimeoffsetTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotByTimeoffsetTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotByTimeoffsetTemplate.__pulumiType;
    }

    /**
     * Template description information, length limit: 256 characters.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
     */
    public readonly fillType!: pulumi.Output<string | undefined>;
    /**
     * Image format, the value can be jpg, png, webp. Default is jpg.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
     */
    public readonly height!: pulumi.Output<number | undefined>;
    /**
     * Snapshot by timeoffset template name, length limit: 64 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
     */
    public readonly resolutionAdaptive!: pulumi.Output<string | undefined>;
    /**
     * The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
     */
    public readonly width!: pulumi.Output<number | undefined>;

    /**
     * Create a SnapshotByTimeoffsetTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnapshotByTimeoffsetTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotByTimeoffsetTemplateArgs | SnapshotByTimeoffsetTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotByTimeoffsetTemplateState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["fillType"] = state ? state.fillType : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["height"] = state ? state.height : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resolutionAdaptive"] = state ? state.resolutionAdaptive : undefined;
            resourceInputs["width"] = state ? state.width : undefined;
        } else {
            const args = argsOrState as SnapshotByTimeoffsetTemplateArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["fillType"] = args ? args.fillType : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["height"] = args ? args.height : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resolutionAdaptive"] = args ? args.resolutionAdaptive : undefined;
            resourceInputs["width"] = args ? args.width : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnapshotByTimeoffsetTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotByTimeoffsetTemplate resources.
 */
export interface SnapshotByTimeoffsetTemplateState {
    /**
     * Template description information, length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
     */
    fillType?: pulumi.Input<string>;
    /**
     * Image format, the value can be jpg, png, webp. Default is jpg.
     */
    format?: pulumi.Input<string>;
    /**
     * The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
     */
    height?: pulumi.Input<number>;
    /**
     * Snapshot by timeoffset template name, length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
     */
    resolutionAdaptive?: pulumi.Input<string>;
    /**
     * The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
     */
    width?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SnapshotByTimeoffsetTemplate resource.
 */
export interface SnapshotByTimeoffsetTemplateArgs {
    /**
     * Template description information, length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
     */
    fillType?: pulumi.Input<string>;
    /**
     * Image format, the value can be jpg, png, webp. Default is jpg.
     */
    format?: pulumi.Input<string>;
    /**
     * The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
     */
    height?: pulumi.Input<number>;
    /**
     * Snapshot by timeoffset template name, length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
     */
    resolutionAdaptive?: pulumi.Input<string>;
    /**
     * The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
     */
    width?: pulumi.Input<number>;
}
