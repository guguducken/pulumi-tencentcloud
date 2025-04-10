// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cvm launchTemplate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const demo = new tencentcloud.Cvm.LaunchTemplate("demo", {
 *     imageId: "img-xxxxxxxxx",
 *     launchTemplateName: "test",
 *     placement: {
 *         hostIds: [],
 *         hostIps: [],
 *         projectId: 0,
 *         zone: "ap-guangzhou-6",
 *     },
 * });
 * ```
 */
export class LaunchTemplate extends pulumi.CustomResource {
    /**
     * Get an existing LaunchTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaunchTemplateState, opts?: pulumi.CustomResourceOptions): LaunchTemplate {
        return new LaunchTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cvm/launchTemplate:LaunchTemplate';

    /**
     * Returns true if the given object is an instance of LaunchTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchTemplate.__pulumiType;
    }

    /**
     * Timed task.
     */
    public readonly actionTimer!: pulumi.Output<outputs.Cvm.LaunchTemplateActionTimer>;
    /**
     * The role name of CAM.
     */
    public readonly camRoleName!: pulumi.Output<string | undefined>;
    /**
     * A string to used guarantee request idempotency.
     */
    public readonly clientToken!: pulumi.Output<string | undefined>;
    /**
     * Data disk configuration information of the instance.
     */
    public readonly dataDisks!: pulumi.Output<outputs.Cvm.LaunchTemplateDataDisk[] | undefined>;
    /**
     * Instance destruction protection flag.
     */
    public readonly disableApiTermination!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of disaster recover group.
     */
    public readonly disasterRecoverGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to preflight only this request, true or false.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
     */
    public readonly enhancedService!: pulumi.Output<outputs.Cvm.LaunchTemplateEnhancedService>;
    /**
     * The host name of CVM.
     */
    public readonly hostName!: pulumi.Output<string | undefined>;
    /**
     * The ID of HPC cluster.
     */
    public readonly hpcClusterId!: pulumi.Output<string | undefined>;
    /**
     * Image ID.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The configuration of charge prepaid.
     */
    public readonly instanceChargePrepaid!: pulumi.Output<outputs.Cvm.LaunchTemplateInstanceChargePrepaid | undefined>;
    /**
     * The charge type of instance. Default value: POSTPAID_BY_HOUR.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * The number of instances purchased.
     */
    public readonly instanceCount!: pulumi.Output<number | undefined>;
    /**
     * The marketplace options of instance.
     */
    public readonly instanceMarketOptions!: pulumi.Output<outputs.Cvm.LaunchTemplateInstanceMarketOptions | undefined>;
    /**
     * The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
     */
    public readonly internetAccessible!: pulumi.Output<outputs.Cvm.LaunchTemplateInternetAccessible>;
    /**
     * The name of launch template.
     */
    public readonly launchTemplateName!: pulumi.Output<string>;
    /**
     * Instance launch template version description.
     */
    public readonly launchTemplateVersionDescription!: pulumi.Output<string | undefined>;
    /**
     * The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
     */
    public readonly loginSettings!: pulumi.Output<outputs.Cvm.LaunchTemplateLoginSettings>;
    /**
     * The location of instance.
     */
    public readonly placement!: pulumi.Output<outputs.Cvm.LaunchTemplatePlacement>;
    /**
     * The security group ID of instance. If this parameter is not specified, the default security group is bound.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
     */
    public readonly systemDisk!: pulumi.Output<outputs.Cvm.LaunchTemplateSystemDisk>;
    /**
     * Tag description list.
     */
    public readonly tagSpecifications!: pulumi.Output<outputs.Cvm.LaunchTemplateTagSpecification[] | undefined>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The data of users.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
     */
    public readonly virtualPrivateCloud!: pulumi.Output<outputs.Cvm.LaunchTemplateVirtualPrivateCloud>;

    /**
     * Create a LaunchTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LaunchTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaunchTemplateArgs | LaunchTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LaunchTemplateState | undefined;
            resourceInputs["actionTimer"] = state ? state.actionTimer : undefined;
            resourceInputs["camRoleName"] = state ? state.camRoleName : undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["disableApiTermination"] = state ? state.disableApiTermination : undefined;
            resourceInputs["disasterRecoverGroupIds"] = state ? state.disasterRecoverGroupIds : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["enhancedService"] = state ? state.enhancedService : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["hpcClusterId"] = state ? state.hpcClusterId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["instanceChargePrepaid"] = state ? state.instanceChargePrepaid : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceCount"] = state ? state.instanceCount : undefined;
            resourceInputs["instanceMarketOptions"] = state ? state.instanceMarketOptions : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetAccessible"] = state ? state.internetAccessible : undefined;
            resourceInputs["launchTemplateName"] = state ? state.launchTemplateName : undefined;
            resourceInputs["launchTemplateVersionDescription"] = state ? state.launchTemplateVersionDescription : undefined;
            resourceInputs["loginSettings"] = state ? state.loginSettings : undefined;
            resourceInputs["placement"] = state ? state.placement : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["systemDisk"] = state ? state.systemDisk : undefined;
            resourceInputs["tagSpecifications"] = state ? state.tagSpecifications : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["virtualPrivateCloud"] = state ? state.virtualPrivateCloud : undefined;
        } else {
            const args = argsOrState as LaunchTemplateArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.launchTemplateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'launchTemplateName'");
            }
            if ((!args || args.placement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placement'");
            }
            resourceInputs["actionTimer"] = args ? args.actionTimer : undefined;
            resourceInputs["camRoleName"] = args ? args.camRoleName : undefined;
            resourceInputs["clientToken"] = args ? args.clientToken : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["disableApiTermination"] = args ? args.disableApiTermination : undefined;
            resourceInputs["disasterRecoverGroupIds"] = args ? args.disasterRecoverGroupIds : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["enhancedService"] = args ? args.enhancedService : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["hpcClusterId"] = args ? args.hpcClusterId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["instanceChargePrepaid"] = args ? args.instanceChargePrepaid : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceCount"] = args ? args.instanceCount : undefined;
            resourceInputs["instanceMarketOptions"] = args ? args.instanceMarketOptions : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetAccessible"] = args ? args.internetAccessible : undefined;
            resourceInputs["launchTemplateName"] = args ? args.launchTemplateName : undefined;
            resourceInputs["launchTemplateVersionDescription"] = args ? args.launchTemplateVersionDescription : undefined;
            resourceInputs["loginSettings"] = args ? args.loginSettings : undefined;
            resourceInputs["placement"] = args ? args.placement : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["systemDisk"] = args ? args.systemDisk : undefined;
            resourceInputs["tagSpecifications"] = args ? args.tagSpecifications : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["virtualPrivateCloud"] = args ? args.virtualPrivateCloud : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LaunchTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LaunchTemplate resources.
 */
export interface LaunchTemplateState {
    /**
     * Timed task.
     */
    actionTimer?: pulumi.Input<inputs.Cvm.LaunchTemplateActionTimer>;
    /**
     * The role name of CAM.
     */
    camRoleName?: pulumi.Input<string>;
    /**
     * A string to used guarantee request idempotency.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * Data disk configuration information of the instance.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateDataDisk>[]>;
    /**
     * Instance destruction protection flag.
     */
    disableApiTermination?: pulumi.Input<boolean>;
    /**
     * The ID of disaster recover group.
     */
    disasterRecoverGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to preflight only this request, true or false.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
     */
    enhancedService?: pulumi.Input<inputs.Cvm.LaunchTemplateEnhancedService>;
    /**
     * The host name of CVM.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The ID of HPC cluster.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * Image ID.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The configuration of charge prepaid.
     */
    instanceChargePrepaid?: pulumi.Input<inputs.Cvm.LaunchTemplateInstanceChargePrepaid>;
    /**
     * The charge type of instance. Default value: POSTPAID_BY_HOUR.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The number of instances purchased.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * The marketplace options of instance.
     */
    instanceMarketOptions?: pulumi.Input<inputs.Cvm.LaunchTemplateInstanceMarketOptions>;
    /**
     * The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
     */
    internetAccessible?: pulumi.Input<inputs.Cvm.LaunchTemplateInternetAccessible>;
    /**
     * The name of launch template.
     */
    launchTemplateName?: pulumi.Input<string>;
    /**
     * Instance launch template version description.
     */
    launchTemplateVersionDescription?: pulumi.Input<string>;
    /**
     * The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
     */
    loginSettings?: pulumi.Input<inputs.Cvm.LaunchTemplateLoginSettings>;
    /**
     * The location of instance.
     */
    placement?: pulumi.Input<inputs.Cvm.LaunchTemplatePlacement>;
    /**
     * The security group ID of instance. If this parameter is not specified, the default security group is bound.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
     */
    systemDisk?: pulumi.Input<inputs.Cvm.LaunchTemplateSystemDisk>;
    /**
     * Tag description list.
     */
    tagSpecifications?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateTagSpecification>[]>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The data of users.
     */
    userData?: pulumi.Input<string>;
    /**
     * The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
     */
    virtualPrivateCloud?: pulumi.Input<inputs.Cvm.LaunchTemplateVirtualPrivateCloud>;
}

/**
 * The set of arguments for constructing a LaunchTemplate resource.
 */
export interface LaunchTemplateArgs {
    /**
     * Timed task.
     */
    actionTimer?: pulumi.Input<inputs.Cvm.LaunchTemplateActionTimer>;
    /**
     * The role name of CAM.
     */
    camRoleName?: pulumi.Input<string>;
    /**
     * A string to used guarantee request idempotency.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * Data disk configuration information of the instance.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateDataDisk>[]>;
    /**
     * Instance destruction protection flag.
     */
    disableApiTermination?: pulumi.Input<boolean>;
    /**
     * The ID of disaster recover group.
     */
    disasterRecoverGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to preflight only this request, true or false.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
     */
    enhancedService?: pulumi.Input<inputs.Cvm.LaunchTemplateEnhancedService>;
    /**
     * The host name of CVM.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The ID of HPC cluster.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * Image ID.
     */
    imageId: pulumi.Input<string>;
    /**
     * The configuration of charge prepaid.
     */
    instanceChargePrepaid?: pulumi.Input<inputs.Cvm.LaunchTemplateInstanceChargePrepaid>;
    /**
     * The charge type of instance. Default value: POSTPAID_BY_HOUR.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The number of instances purchased.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * The marketplace options of instance.
     */
    instanceMarketOptions?: pulumi.Input<inputs.Cvm.LaunchTemplateInstanceMarketOptions>;
    /**
     * The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
     */
    internetAccessible?: pulumi.Input<inputs.Cvm.LaunchTemplateInternetAccessible>;
    /**
     * The name of launch template.
     */
    launchTemplateName: pulumi.Input<string>;
    /**
     * Instance launch template version description.
     */
    launchTemplateVersionDescription?: pulumi.Input<string>;
    /**
     * The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
     */
    loginSettings?: pulumi.Input<inputs.Cvm.LaunchTemplateLoginSettings>;
    /**
     * The location of instance.
     */
    placement: pulumi.Input<inputs.Cvm.LaunchTemplatePlacement>;
    /**
     * The security group ID of instance. If this parameter is not specified, the default security group is bound.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
     */
    systemDisk?: pulumi.Input<inputs.Cvm.LaunchTemplateSystemDisk>;
    /**
     * Tag description list.
     */
    tagSpecifications?: pulumi.Input<pulumi.Input<inputs.Cvm.LaunchTemplateTagSpecification>[]>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The data of users.
     */
    userData?: pulumi.Input<string>;
    /**
     * The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
     */
    virtualPrivateCloud?: pulumi.Input<inputs.Cvm.LaunchTemplateVirtualPrivateCloud>;
}
