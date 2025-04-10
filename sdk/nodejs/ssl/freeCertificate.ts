// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a Free Certificate.
 *
 * > **NOTE:** Once certificat created, it cannot be removed within 1 hours.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Ssl.FreeCertificate("foo", {
 *     alias: "my_free_cert",
 *     contactEmail: "foo@example.com",
 *     contactPhone: "12345678901",
 *     csrEncryptAlgo: "RSA",
 *     csrKeyParameter: "2048",
 *     csrKeyPassword: "xxxxxxxx",
 *     domain: "example.com",
 *     dvAuthMethod: "DNS_AUTO",
 *     packageType: "2",
 *     validityPeriod: "12",
 * });
 * ```
 *
 * ## Import
 *
 * FreeCertificate instance can be imported, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ssl/freeCertificate:FreeCertificate test free_certificate-id
 * ```
 */
export class FreeCertificate extends pulumi.CustomResource {
    /**
     * Get an existing FreeCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FreeCertificateState, opts?: pulumi.CustomResourceOptions): FreeCertificate {
        return new FreeCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssl/freeCertificate:FreeCertificate';

    /**
     * Returns true if the given object is an instance of FreeCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FreeCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FreeCertificate.__pulumiType;
    }

    /**
     * Specify alias for remark.
     */
    public readonly alias!: pulumi.Output<string | undefined>;
    /**
     * Certificate begin time.
     */
    public /*out*/ readonly certBeginTime!: pulumi.Output<string>;
    /**
     * Certificate end time.
     */
    public /*out*/ readonly certEndTime!: pulumi.Output<string>;
    /**
     * Certificate private key.
     */
    public /*out*/ readonly certificatePrivateKey!: pulumi.Output<string>;
    /**
     * Certificate public key.
     */
    public /*out*/ readonly certificatePublicKey!: pulumi.Output<string>;
    /**
     * Email address.
     */
    public readonly contactEmail!: pulumi.Output<string | undefined>;
    /**
     * Phone number.
     */
    public readonly contactPhone!: pulumi.Output<string | undefined>;
    /**
     * Specify CSR encrypt algorithm, only support `RSA` for now.
     */
    public readonly csrEncryptAlgo!: pulumi.Output<string | undefined>;
    /**
     * Specify CSR key parameter, only support `"2048"` for now.
     */
    public readonly csrKeyParameter!: pulumi.Output<string | undefined>;
    /**
     * Specify CSR key password.
     */
    public readonly csrKeyPassword!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the certificate deployable.
     */
    public /*out*/ readonly deployable!: pulumi.Output<boolean>;
    /**
     * Specify domain name.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
     */
    public readonly dvAuthMethod!: pulumi.Output<string>;
    /**
     * DV certification information.
     */
    public /*out*/ readonly dvAuths!: pulumi.Output<outputs.Ssl.FreeCertificateDvAuth[]>;
    /**
     * Certificate insert time.
     */
    public /*out*/ readonly insertTime!: pulumi.Output<string>;
    /**
     * Specify old certificate ID, used for re-apply.
     */
    public readonly oldCertificateId!: pulumi.Output<string | undefined>;
    /**
     * Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
     */
    public readonly packageType!: pulumi.Output<string | undefined>;
    /**
     * Product zh name.
     */
    public /*out*/ readonly productZhName!: pulumi.Output<string>;
    /**
     * ID of projects which this certification belong to.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether the certificate renewable.
     */
    public /*out*/ readonly renewable!: pulumi.Output<boolean>;
    /**
     * Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Certificate status message.
     */
    public /*out*/ readonly statusMsg!: pulumi.Output<string>;
    /**
     * Certificate status name.
     */
    public /*out*/ readonly statusName!: pulumi.Output<string>;
    /**
     * Specify validity period in month, only support `"12"` months for now.
     */
    public readonly validityPeriod!: pulumi.Output<string | undefined>;
    /**
     * Vulnerability status.
     */
    public /*out*/ readonly vulnerabilityStatus!: pulumi.Output<string>;

    /**
     * Create a FreeCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FreeCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FreeCertificateArgs | FreeCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FreeCertificateState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["certBeginTime"] = state ? state.certBeginTime : undefined;
            resourceInputs["certEndTime"] = state ? state.certEndTime : undefined;
            resourceInputs["certificatePrivateKey"] = state ? state.certificatePrivateKey : undefined;
            resourceInputs["certificatePublicKey"] = state ? state.certificatePublicKey : undefined;
            resourceInputs["contactEmail"] = state ? state.contactEmail : undefined;
            resourceInputs["contactPhone"] = state ? state.contactPhone : undefined;
            resourceInputs["csrEncryptAlgo"] = state ? state.csrEncryptAlgo : undefined;
            resourceInputs["csrKeyParameter"] = state ? state.csrKeyParameter : undefined;
            resourceInputs["csrKeyPassword"] = state ? state.csrKeyPassword : undefined;
            resourceInputs["deployable"] = state ? state.deployable : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["dvAuthMethod"] = state ? state.dvAuthMethod : undefined;
            resourceInputs["dvAuths"] = state ? state.dvAuths : undefined;
            resourceInputs["insertTime"] = state ? state.insertTime : undefined;
            resourceInputs["oldCertificateId"] = state ? state.oldCertificateId : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["productZhName"] = state ? state.productZhName : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["renewable"] = state ? state.renewable : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusMsg"] = state ? state.statusMsg : undefined;
            resourceInputs["statusName"] = state ? state.statusName : undefined;
            resourceInputs["validityPeriod"] = state ? state.validityPeriod : undefined;
            resourceInputs["vulnerabilityStatus"] = state ? state.vulnerabilityStatus : undefined;
        } else {
            const args = argsOrState as FreeCertificateArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.dvAuthMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dvAuthMethod'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["contactEmail"] = args ? args.contactEmail : undefined;
            resourceInputs["contactPhone"] = args ? args.contactPhone : undefined;
            resourceInputs["csrEncryptAlgo"] = args ? args.csrEncryptAlgo : undefined;
            resourceInputs["csrKeyParameter"] = args ? args.csrKeyParameter : undefined;
            resourceInputs["csrKeyPassword"] = args ? args.csrKeyPassword : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["dvAuthMethod"] = args ? args.dvAuthMethod : undefined;
            resourceInputs["oldCertificateId"] = args ? args.oldCertificateId : undefined;
            resourceInputs["packageType"] = args ? args.packageType : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["validityPeriod"] = args ? args.validityPeriod : undefined;
            resourceInputs["certBeginTime"] = undefined /*out*/;
            resourceInputs["certEndTime"] = undefined /*out*/;
            resourceInputs["certificatePrivateKey"] = undefined /*out*/;
            resourceInputs["certificatePublicKey"] = undefined /*out*/;
            resourceInputs["deployable"] = undefined /*out*/;
            resourceInputs["dvAuths"] = undefined /*out*/;
            resourceInputs["insertTime"] = undefined /*out*/;
            resourceInputs["productZhName"] = undefined /*out*/;
            resourceInputs["renewable"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMsg"] = undefined /*out*/;
            resourceInputs["statusName"] = undefined /*out*/;
            resourceInputs["vulnerabilityStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FreeCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FreeCertificate resources.
 */
export interface FreeCertificateState {
    /**
     * Specify alias for remark.
     */
    alias?: pulumi.Input<string>;
    /**
     * Certificate begin time.
     */
    certBeginTime?: pulumi.Input<string>;
    /**
     * Certificate end time.
     */
    certEndTime?: pulumi.Input<string>;
    /**
     * Certificate private key.
     */
    certificatePrivateKey?: pulumi.Input<string>;
    /**
     * Certificate public key.
     */
    certificatePublicKey?: pulumi.Input<string>;
    /**
     * Email address.
     */
    contactEmail?: pulumi.Input<string>;
    /**
     * Phone number.
     */
    contactPhone?: pulumi.Input<string>;
    /**
     * Specify CSR encrypt algorithm, only support `RSA` for now.
     */
    csrEncryptAlgo?: pulumi.Input<string>;
    /**
     * Specify CSR key parameter, only support `"2048"` for now.
     */
    csrKeyParameter?: pulumi.Input<string>;
    /**
     * Specify CSR key password.
     */
    csrKeyPassword?: pulumi.Input<string>;
    /**
     * Indicates whether the certificate deployable.
     */
    deployable?: pulumi.Input<boolean>;
    /**
     * Specify domain name.
     */
    domain?: pulumi.Input<string>;
    /**
     * Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
     */
    dvAuthMethod?: pulumi.Input<string>;
    /**
     * DV certification information.
     */
    dvAuths?: pulumi.Input<pulumi.Input<inputs.Ssl.FreeCertificateDvAuth>[]>;
    /**
     * Certificate insert time.
     */
    insertTime?: pulumi.Input<string>;
    /**
     * Specify old certificate ID, used for re-apply.
     */
    oldCertificateId?: pulumi.Input<string>;
    /**
     * Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
     */
    packageType?: pulumi.Input<string>;
    /**
     * Product zh name.
     */
    productZhName?: pulumi.Input<string>;
    /**
     * ID of projects which this certification belong to.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Indicates whether the certificate renewable.
     */
    renewable?: pulumi.Input<boolean>;
    /**
     * Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
     */
    status?: pulumi.Input<number>;
    /**
     * Certificate status message.
     */
    statusMsg?: pulumi.Input<string>;
    /**
     * Certificate status name.
     */
    statusName?: pulumi.Input<string>;
    /**
     * Specify validity period in month, only support `"12"` months for now.
     */
    validityPeriod?: pulumi.Input<string>;
    /**
     * Vulnerability status.
     */
    vulnerabilityStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FreeCertificate resource.
 */
export interface FreeCertificateArgs {
    /**
     * Specify alias for remark.
     */
    alias?: pulumi.Input<string>;
    /**
     * Email address.
     */
    contactEmail?: pulumi.Input<string>;
    /**
     * Phone number.
     */
    contactPhone?: pulumi.Input<string>;
    /**
     * Specify CSR encrypt algorithm, only support `RSA` for now.
     */
    csrEncryptAlgo?: pulumi.Input<string>;
    /**
     * Specify CSR key parameter, only support `"2048"` for now.
     */
    csrKeyParameter?: pulumi.Input<string>;
    /**
     * Specify CSR key password.
     */
    csrKeyPassword?: pulumi.Input<string>;
    /**
     * Specify domain name.
     */
    domain: pulumi.Input<string>;
    /**
     * Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
     */
    dvAuthMethod: pulumi.Input<string>;
    /**
     * Specify old certificate ID, used for re-apply.
     */
    oldCertificateId?: pulumi.Input<string>;
    /**
     * Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
     */
    packageType?: pulumi.Input<string>;
    /**
     * ID of projects which this certification belong to.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Specify validity period in month, only support `"12"` months for now.
     */
    validityPeriod?: pulumi.Input<string>;
}
