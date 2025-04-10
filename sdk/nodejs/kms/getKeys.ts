// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of KMS key
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = pulumi.output(tencentcloud.Kms.getKeys({
 *     keyState: 0,
 *     keyUsage: "ALL",
 *     origin: "TENCENT_KMS",
 *     searchKeyAlias: "test",
 * }));
 * ```
 */
export function getKeys(args?: GetKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetKeysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Kms/getKeys:getKeys", {
        "keyState": args.keyState,
        "keyUsage": args.keyUsage,
        "orderType": args.orderType,
        "origin": args.origin,
        "resultOutputFile": args.resultOutputFile,
        "role": args.role,
        "searchKeyAlias": args.searchKeyAlias,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeys.
 */
export interface GetKeysArgs {
    /**
     * Filter by state of CMK. `0` - all CMKs are queried, `1` - only Enabled CMKs are queried, `2` - only Disabled CMKs are queried, `3` - only PendingDelete CMKs are queried, `4` - only PendingImport CMKs are queried, `5` - only Archived CMKs are queried.
     */
    keyState?: number;
    /**
     * Filter by usage of CMK. Available values include `ALL`, `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
     */
    keyUsage?: string;
    /**
     * Order to sort the CMK create time. `0` - desc, `1` - asc. Default value is `0`.
     */
    orderType?: number;
    /**
     * Filter by origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user, `ALL` - all CMKs. Default value is `ALL`.
     */
    origin?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Filter by role of the CMK creator. `0` - created by user, `1` - created by cloud product. Default value is `0`.
     */
    role?: number;
    /**
     * Words used to match the results, and the words can be: keyId and alias.
     */
    searchKeyAlias?: string;
    /**
     * Tags to filter CMK.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getKeys.
 */
export interface GetKeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of KMS keys.
     */
    readonly keyLists: outputs.Kms.GetKeysKeyList[];
    /**
     * State of CMK.
     */
    readonly keyState?: number;
    /**
     * Usage of CMK.
     */
    readonly keyUsage?: string;
    readonly orderType?: number;
    /**
     * Origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user.
     */
    readonly origin?: string;
    readonly resultOutputFile?: string;
    readonly role?: number;
    readonly searchKeyAlias?: string;
    readonly tags?: {[key: string]: any};
}

export function getKeysOutput(args?: GetKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeysResult> {
    return pulumi.output(args).apply(a => getKeys(a, opts))
}

/**
 * A collection of arguments for invoking getKeys.
 */
export interface GetKeysOutputArgs {
    /**
     * Filter by state of CMK. `0` - all CMKs are queried, `1` - only Enabled CMKs are queried, `2` - only Disabled CMKs are queried, `3` - only PendingDelete CMKs are queried, `4` - only PendingImport CMKs are queried, `5` - only Archived CMKs are queried.
     */
    keyState?: pulumi.Input<number>;
    /**
     * Filter by usage of CMK. Available values include `ALL`, `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
     */
    keyUsage?: pulumi.Input<string>;
    /**
     * Order to sort the CMK create time. `0` - desc, `1` - asc. Default value is `0`.
     */
    orderType?: pulumi.Input<number>;
    /**
     * Filter by origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user, `ALL` - all CMKs. Default value is `ALL`.
     */
    origin?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Filter by role of the CMK creator. `0` - created by user, `1` - created by cloud product. Default value is `0`.
     */
    role?: pulumi.Input<number>;
    /**
     * Words used to match the results, and the words can be: keyId and alias.
     */
    searchKeyAlias?: pulumi.Input<string>;
    /**
     * Tags to filter CMK.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
