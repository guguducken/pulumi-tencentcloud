// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a eip publicAddressAdjust
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const publicAddressAdjust = new tencentcloud.Eip.PublicAddressAdjust("public_address_adjust", {
 *     addressId: "eip-erft45fu",
 *     instanceId: "ins-cr2rfq78",
 * });
 * ```
 */
export class PublicAddressAdjust extends pulumi.CustomResource {
    /**
     * Get an existing PublicAddressAdjust resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicAddressAdjustState, opts?: pulumi.CustomResourceOptions): PublicAddressAdjust {
        return new PublicAddressAdjust(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Eip/publicAddressAdjust:PublicAddressAdjust';

    /**
     * Returns true if the given object is an instance of PublicAddressAdjust.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicAddressAdjust {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicAddressAdjust.__pulumiType;
    }

    /**
     * A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
     */
    public readonly addressId!: pulumi.Output<string | undefined>;
    /**
     * A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;

    /**
     * Create a PublicAddressAdjust resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PublicAddressAdjustArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicAddressAdjustArgs | PublicAddressAdjustState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicAddressAdjustState | undefined;
            resourceInputs["addressId"] = state ? state.addressId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as PublicAddressAdjustArgs | undefined;
            resourceInputs["addressId"] = args ? args.addressId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicAddressAdjust.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicAddressAdjust resources.
 */
export interface PublicAddressAdjustState {
    /**
     * A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
     */
    addressId?: pulumi.Input<string>;
    /**
     * A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublicAddressAdjust resource.
 */
export interface PublicAddressAdjustArgs {
    /**
     * A unique ID that identifies an EIP instance. The unique ID of EIP is in the form:`eip-erft45fu`.
     */
    addressId?: pulumi.Input<string>;
    /**
     * A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.
     */
    instanceId?: pulumi.Input<string>;
}
