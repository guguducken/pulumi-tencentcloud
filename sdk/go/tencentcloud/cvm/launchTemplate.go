// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cvm launchTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cvm.NewLaunchTemplate(ctx, "demo", &Cvm.LaunchTemplateArgs{
// 			ImageId:            pulumi.String("img-xxxxxxxxx"),
// 			LaunchTemplateName: pulumi.String("test"),
// 			Placement: &cvm.LaunchTemplatePlacementArgs{
// 				HostIds:   pulumi.StringArray{},
// 				HostIps:   pulumi.StringArray{},
// 				ProjectId: pulumi.Int(0),
// 				Zone:      pulumi.String("ap-guangzhou-6"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LaunchTemplate struct {
	pulumi.CustomResourceState

	// Timed task.
	ActionTimer LaunchTemplateActionTimerOutput `pulumi:"actionTimer"`
	// The role name of CAM.
	CamRoleName pulumi.StringPtrOutput `pulumi:"camRoleName"`
	// A string to used guarantee request idempotency.
	ClientToken pulumi.StringPtrOutput `pulumi:"clientToken"`
	// Data disk configuration information of the instance.
	DataDisks LaunchTemplateDataDiskArrayOutput `pulumi:"dataDisks"`
	// Instance destruction protection flag.
	DisableApiTermination pulumi.BoolPtrOutput `pulumi:"disableApiTermination"`
	// The ID of disaster recover group.
	DisasterRecoverGroupIds pulumi.StringArrayOutput `pulumi:"disasterRecoverGroupIds"`
	// Whether to preflight only this request, true or false.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
	EnhancedService LaunchTemplateEnhancedServiceOutput `pulumi:"enhancedService"`
	// The host name of CVM.
	HostName pulumi.StringPtrOutput `pulumi:"hostName"`
	// The ID of HPC cluster.
	HpcClusterId pulumi.StringPtrOutput `pulumi:"hpcClusterId"`
	// Image ID.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The configuration of charge prepaid.
	InstanceChargePrepaid LaunchTemplateInstanceChargePrepaidPtrOutput `pulumi:"instanceChargePrepaid"`
	// The charge type of instance. Default value: POSTPAID_BY_HOUR.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// The number of instances purchased.
	InstanceCount pulumi.IntPtrOutput `pulumi:"instanceCount"`
	// The marketplace options of instance.
	InstanceMarketOptions LaunchTemplateInstanceMarketOptionsPtrOutput `pulumi:"instanceMarketOptions"`
	// The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
	InternetAccessible LaunchTemplateInternetAccessibleOutput `pulumi:"internetAccessible"`
	// The name of launch template.
	LaunchTemplateName pulumi.StringOutput `pulumi:"launchTemplateName"`
	// Instance launch template version description.
	LaunchTemplateVersionDescription pulumi.StringPtrOutput `pulumi:"launchTemplateVersionDescription"`
	// The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
	LoginSettings LaunchTemplateLoginSettingsOutput `pulumi:"loginSettings"`
	// The location of instance.
	Placement LaunchTemplatePlacementOutput `pulumi:"placement"`
	// The security group ID of instance. If this parameter is not specified, the default security group is bound.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
	SystemDisk LaunchTemplateSystemDiskOutput `pulumi:"systemDisk"`
	// Tag description list.
	TagSpecifications LaunchTemplateTagSpecificationArrayOutput `pulumi:"tagSpecifications"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The data of users.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
	VirtualPrivateCloud LaunchTemplateVirtualPrivateCloudOutput `pulumi:"virtualPrivateCloud"`
}

// NewLaunchTemplate registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplate(ctx *pulumi.Context,
	name string, args *LaunchTemplateArgs, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.LaunchTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'LaunchTemplateName'")
	}
	if args.Placement == nil {
		return nil, errors.New("invalid value for required argument 'Placement'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LaunchTemplate
	err := ctx.RegisterResource("tencentcloud:Cvm/launchTemplate:LaunchTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplate gets an existing LaunchTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateState, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	var resource LaunchTemplate
	err := ctx.ReadResource("tencentcloud:Cvm/launchTemplate:LaunchTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplate resources.
type launchTemplateState struct {
	// Timed task.
	ActionTimer *LaunchTemplateActionTimer `pulumi:"actionTimer"`
	// The role name of CAM.
	CamRoleName *string `pulumi:"camRoleName"`
	// A string to used guarantee request idempotency.
	ClientToken *string `pulumi:"clientToken"`
	// Data disk configuration information of the instance.
	DataDisks []LaunchTemplateDataDisk `pulumi:"dataDisks"`
	// Instance destruction protection flag.
	DisableApiTermination *bool `pulumi:"disableApiTermination"`
	// The ID of disaster recover group.
	DisasterRecoverGroupIds []string `pulumi:"disasterRecoverGroupIds"`
	// Whether to preflight only this request, true or false.
	DryRun *bool `pulumi:"dryRun"`
	// Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
	EnhancedService *LaunchTemplateEnhancedService `pulumi:"enhancedService"`
	// The host name of CVM.
	HostName *string `pulumi:"hostName"`
	// The ID of HPC cluster.
	HpcClusterId *string `pulumi:"hpcClusterId"`
	// Image ID.
	ImageId *string `pulumi:"imageId"`
	// The configuration of charge prepaid.
	InstanceChargePrepaid *LaunchTemplateInstanceChargePrepaid `pulumi:"instanceChargePrepaid"`
	// The charge type of instance. Default value: POSTPAID_BY_HOUR.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The number of instances purchased.
	InstanceCount *int `pulumi:"instanceCount"`
	// The marketplace options of instance.
	InstanceMarketOptions *LaunchTemplateInstanceMarketOptions `pulumi:"instanceMarketOptions"`
	// The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
	InstanceName *string `pulumi:"instanceName"`
	// The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
	InstanceType *string `pulumi:"instanceType"`
	// The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
	InternetAccessible *LaunchTemplateInternetAccessible `pulumi:"internetAccessible"`
	// The name of launch template.
	LaunchTemplateName *string `pulumi:"launchTemplateName"`
	// Instance launch template version description.
	LaunchTemplateVersionDescription *string `pulumi:"launchTemplateVersionDescription"`
	// The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
	LoginSettings *LaunchTemplateLoginSettings `pulumi:"loginSettings"`
	// The location of instance.
	Placement *LaunchTemplatePlacement `pulumi:"placement"`
	// The security group ID of instance. If this parameter is not specified, the default security group is bound.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
	SystemDisk *LaunchTemplateSystemDisk `pulumi:"systemDisk"`
	// Tag description list.
	TagSpecifications []LaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// The data of users.
	UserData *string `pulumi:"userData"`
	// The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
	VirtualPrivateCloud *LaunchTemplateVirtualPrivateCloud `pulumi:"virtualPrivateCloud"`
}

type LaunchTemplateState struct {
	// Timed task.
	ActionTimer LaunchTemplateActionTimerPtrInput
	// The role name of CAM.
	CamRoleName pulumi.StringPtrInput
	// A string to used guarantee request idempotency.
	ClientToken pulumi.StringPtrInput
	// Data disk configuration information of the instance.
	DataDisks LaunchTemplateDataDiskArrayInput
	// Instance destruction protection flag.
	DisableApiTermination pulumi.BoolPtrInput
	// The ID of disaster recover group.
	DisasterRecoverGroupIds pulumi.StringArrayInput
	// Whether to preflight only this request, true or false.
	DryRun pulumi.BoolPtrInput
	// Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
	EnhancedService LaunchTemplateEnhancedServicePtrInput
	// The host name of CVM.
	HostName pulumi.StringPtrInput
	// The ID of HPC cluster.
	HpcClusterId pulumi.StringPtrInput
	// Image ID.
	ImageId pulumi.StringPtrInput
	// The configuration of charge prepaid.
	InstanceChargePrepaid LaunchTemplateInstanceChargePrepaidPtrInput
	// The charge type of instance. Default value: POSTPAID_BY_HOUR.
	InstanceChargeType pulumi.StringPtrInput
	// The number of instances purchased.
	InstanceCount pulumi.IntPtrInput
	// The marketplace options of instance.
	InstanceMarketOptions LaunchTemplateInstanceMarketOptionsPtrInput
	// The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
	InstanceName pulumi.StringPtrInput
	// The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
	InstanceType pulumi.StringPtrInput
	// The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
	InternetAccessible LaunchTemplateInternetAccessiblePtrInput
	// The name of launch template.
	LaunchTemplateName pulumi.StringPtrInput
	// Instance launch template version description.
	LaunchTemplateVersionDescription pulumi.StringPtrInput
	// The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
	LoginSettings LaunchTemplateLoginSettingsPtrInput
	// The location of instance.
	Placement LaunchTemplatePlacementPtrInput
	// The security group ID of instance. If this parameter is not specified, the default security group is bound.
	SecurityGroupIds pulumi.StringArrayInput
	// System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
	SystemDisk LaunchTemplateSystemDiskPtrInput
	// Tag description list.
	TagSpecifications LaunchTemplateTagSpecificationArrayInput
	// Tag description list.
	Tags pulumi.MapInput
	// The data of users.
	UserData pulumi.StringPtrInput
	// The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
	VirtualPrivateCloud LaunchTemplateVirtualPrivateCloudPtrInput
}

func (LaunchTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateState)(nil)).Elem()
}

type launchTemplateArgs struct {
	// Timed task.
	ActionTimer *LaunchTemplateActionTimer `pulumi:"actionTimer"`
	// The role name of CAM.
	CamRoleName *string `pulumi:"camRoleName"`
	// A string to used guarantee request idempotency.
	ClientToken *string `pulumi:"clientToken"`
	// Data disk configuration information of the instance.
	DataDisks []LaunchTemplateDataDisk `pulumi:"dataDisks"`
	// Instance destruction protection flag.
	DisableApiTermination *bool `pulumi:"disableApiTermination"`
	// The ID of disaster recover group.
	DisasterRecoverGroupIds []string `pulumi:"disasterRecoverGroupIds"`
	// Whether to preflight only this request, true or false.
	DryRun *bool `pulumi:"dryRun"`
	// Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
	EnhancedService *LaunchTemplateEnhancedService `pulumi:"enhancedService"`
	// The host name of CVM.
	HostName *string `pulumi:"hostName"`
	// The ID of HPC cluster.
	HpcClusterId *string `pulumi:"hpcClusterId"`
	// Image ID.
	ImageId string `pulumi:"imageId"`
	// The configuration of charge prepaid.
	InstanceChargePrepaid *LaunchTemplateInstanceChargePrepaid `pulumi:"instanceChargePrepaid"`
	// The charge type of instance. Default value: POSTPAID_BY_HOUR.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The number of instances purchased.
	InstanceCount *int `pulumi:"instanceCount"`
	// The marketplace options of instance.
	InstanceMarketOptions *LaunchTemplateInstanceMarketOptions `pulumi:"instanceMarketOptions"`
	// The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
	InstanceName *string `pulumi:"instanceName"`
	// The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
	InstanceType *string `pulumi:"instanceType"`
	// The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
	InternetAccessible *LaunchTemplateInternetAccessible `pulumi:"internetAccessible"`
	// The name of launch template.
	LaunchTemplateName string `pulumi:"launchTemplateName"`
	// Instance launch template version description.
	LaunchTemplateVersionDescription *string `pulumi:"launchTemplateVersionDescription"`
	// The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
	LoginSettings *LaunchTemplateLoginSettings `pulumi:"loginSettings"`
	// The location of instance.
	Placement LaunchTemplatePlacement `pulumi:"placement"`
	// The security group ID of instance. If this parameter is not specified, the default security group is bound.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
	SystemDisk *LaunchTemplateSystemDisk `pulumi:"systemDisk"`
	// Tag description list.
	TagSpecifications []LaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// The data of users.
	UserData *string `pulumi:"userData"`
	// The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
	VirtualPrivateCloud *LaunchTemplateVirtualPrivateCloud `pulumi:"virtualPrivateCloud"`
}

// The set of arguments for constructing a LaunchTemplate resource.
type LaunchTemplateArgs struct {
	// Timed task.
	ActionTimer LaunchTemplateActionTimerPtrInput
	// The role name of CAM.
	CamRoleName pulumi.StringPtrInput
	// A string to used guarantee request idempotency.
	ClientToken pulumi.StringPtrInput
	// Data disk configuration information of the instance.
	DataDisks LaunchTemplateDataDiskArrayInput
	// Instance destruction protection flag.
	DisableApiTermination pulumi.BoolPtrInput
	// The ID of disaster recover group.
	DisasterRecoverGroupIds pulumi.StringArrayInput
	// Whether to preflight only this request, true or false.
	DryRun pulumi.BoolPtrInput
	// Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
	EnhancedService LaunchTemplateEnhancedServicePtrInput
	// The host name of CVM.
	HostName pulumi.StringPtrInput
	// The ID of HPC cluster.
	HpcClusterId pulumi.StringPtrInput
	// Image ID.
	ImageId pulumi.StringInput
	// The configuration of charge prepaid.
	InstanceChargePrepaid LaunchTemplateInstanceChargePrepaidPtrInput
	// The charge type of instance. Default value: POSTPAID_BY_HOUR.
	InstanceChargeType pulumi.StringPtrInput
	// The number of instances purchased.
	InstanceCount pulumi.IntPtrInput
	// The marketplace options of instance.
	InstanceMarketOptions LaunchTemplateInstanceMarketOptionsPtrInput
	// The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
	InstanceName pulumi.StringPtrInput
	// The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
	InstanceType pulumi.StringPtrInput
	// The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
	InternetAccessible LaunchTemplateInternetAccessiblePtrInput
	// The name of launch template.
	LaunchTemplateName pulumi.StringInput
	// Instance launch template version description.
	LaunchTemplateVersionDescription pulumi.StringPtrInput
	// The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
	LoginSettings LaunchTemplateLoginSettingsPtrInput
	// The location of instance.
	Placement LaunchTemplatePlacementInput
	// The security group ID of instance. If this parameter is not specified, the default security group is bound.
	SecurityGroupIds pulumi.StringArrayInput
	// System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
	SystemDisk LaunchTemplateSystemDiskPtrInput
	// Tag description list.
	TagSpecifications LaunchTemplateTagSpecificationArrayInput
	// Tag description list.
	Tags pulumi.MapInput
	// The data of users.
	UserData pulumi.StringPtrInput
	// The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
	VirtualPrivateCloud LaunchTemplateVirtualPrivateCloudPtrInput
}

func (LaunchTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateArgs)(nil)).Elem()
}

type LaunchTemplateInput interface {
	pulumi.Input

	ToLaunchTemplateOutput() LaunchTemplateOutput
	ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput
}

func (*LaunchTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplate)(nil)).Elem()
}

func (i *LaunchTemplate) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return i.ToLaunchTemplateOutputWithContext(context.Background())
}

func (i *LaunchTemplate) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateOutput)
}

// LaunchTemplateArrayInput is an input type that accepts LaunchTemplateArray and LaunchTemplateArrayOutput values.
// You can construct a concrete instance of `LaunchTemplateArrayInput` via:
//
//          LaunchTemplateArray{ LaunchTemplateArgs{...} }
type LaunchTemplateArrayInput interface {
	pulumi.Input

	ToLaunchTemplateArrayOutput() LaunchTemplateArrayOutput
	ToLaunchTemplateArrayOutputWithContext(context.Context) LaunchTemplateArrayOutput
}

type LaunchTemplateArray []LaunchTemplateInput

func (LaunchTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LaunchTemplate)(nil)).Elem()
}

func (i LaunchTemplateArray) ToLaunchTemplateArrayOutput() LaunchTemplateArrayOutput {
	return i.ToLaunchTemplateArrayOutputWithContext(context.Background())
}

func (i LaunchTemplateArray) ToLaunchTemplateArrayOutputWithContext(ctx context.Context) LaunchTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateArrayOutput)
}

// LaunchTemplateMapInput is an input type that accepts LaunchTemplateMap and LaunchTemplateMapOutput values.
// You can construct a concrete instance of `LaunchTemplateMapInput` via:
//
//          LaunchTemplateMap{ "key": LaunchTemplateArgs{...} }
type LaunchTemplateMapInput interface {
	pulumi.Input

	ToLaunchTemplateMapOutput() LaunchTemplateMapOutput
	ToLaunchTemplateMapOutputWithContext(context.Context) LaunchTemplateMapOutput
}

type LaunchTemplateMap map[string]LaunchTemplateInput

func (LaunchTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LaunchTemplate)(nil)).Elem()
}

func (i LaunchTemplateMap) ToLaunchTemplateMapOutput() LaunchTemplateMapOutput {
	return i.ToLaunchTemplateMapOutputWithContext(context.Background())
}

func (i LaunchTemplateMap) ToLaunchTemplateMapOutputWithContext(ctx context.Context) LaunchTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateMapOutput)
}

type LaunchTemplateOutput struct{ *pulumi.OutputState }

func (LaunchTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplate)(nil)).Elem()
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return o
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return o
}

// Timed task.
func (o LaunchTemplateOutput) ActionTimer() LaunchTemplateActionTimerOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateActionTimerOutput { return v.ActionTimer }).(LaunchTemplateActionTimerOutput)
}

// The role name of CAM.
func (o LaunchTemplateOutput) CamRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.CamRoleName }).(pulumi.StringPtrOutput)
}

// A string to used guarantee request idempotency.
func (o LaunchTemplateOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.ClientToken }).(pulumi.StringPtrOutput)
}

// Data disk configuration information of the instance.
func (o LaunchTemplateOutput) DataDisks() LaunchTemplateDataDiskArrayOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateDataDiskArrayOutput { return v.DataDisks }).(LaunchTemplateDataDiskArrayOutput)
}

// Instance destruction protection flag.
func (o LaunchTemplateOutput) DisableApiTermination() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.BoolPtrOutput { return v.DisableApiTermination }).(pulumi.BoolPtrOutput)
}

// The ID of disaster recover group.
func (o LaunchTemplateOutput) DisasterRecoverGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringArrayOutput { return v.DisasterRecoverGroupIds }).(pulumi.StringArrayOutput)
}

// Whether to preflight only this request, true or false.
func (o LaunchTemplateOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.
func (o LaunchTemplateOutput) EnhancedService() LaunchTemplateEnhancedServiceOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateEnhancedServiceOutput { return v.EnhancedService }).(LaunchTemplateEnhancedServiceOutput)
}

// The host name of CVM.
func (o LaunchTemplateOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.HostName }).(pulumi.StringPtrOutput)
}

// The ID of HPC cluster.
func (o LaunchTemplateOutput) HpcClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.HpcClusterId }).(pulumi.StringPtrOutput)
}

// Image ID.
func (o LaunchTemplateOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The configuration of charge prepaid.
func (o LaunchTemplateOutput) InstanceChargePrepaid() LaunchTemplateInstanceChargePrepaidPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateInstanceChargePrepaidPtrOutput { return v.InstanceChargePrepaid }).(LaunchTemplateInstanceChargePrepaidPtrOutput)
}

// The charge type of instance. Default value: POSTPAID_BY_HOUR.
func (o LaunchTemplateOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.InstanceChargeType }).(pulumi.StringOutput)
}

// The number of instances purchased.
func (o LaunchTemplateOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.IntPtrOutput { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

// The marketplace options of instance.
func (o LaunchTemplateOutput) InstanceMarketOptions() LaunchTemplateInstanceMarketOptionsPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateInstanceMarketOptionsPtrOutput { return v.InstanceMarketOptions }).(LaunchTemplateInstanceMarketOptionsPtrOutput)
}

// The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.
func (o LaunchTemplateOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.
func (o LaunchTemplateOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.
func (o LaunchTemplateOutput) InternetAccessible() LaunchTemplateInternetAccessibleOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateInternetAccessibleOutput { return v.InternetAccessible }).(LaunchTemplateInternetAccessibleOutput)
}

// The name of launch template.
func (o LaunchTemplateOutput) LaunchTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringOutput { return v.LaunchTemplateName }).(pulumi.StringOutput)
}

// Instance launch template version description.
func (o LaunchTemplateOutput) LaunchTemplateVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.LaunchTemplateVersionDescription }).(pulumi.StringPtrOutput)
}

// The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.
func (o LaunchTemplateOutput) LoginSettings() LaunchTemplateLoginSettingsOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateLoginSettingsOutput { return v.LoginSettings }).(LaunchTemplateLoginSettingsOutput)
}

// The location of instance.
func (o LaunchTemplateOutput) Placement() LaunchTemplatePlacementOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplatePlacementOutput { return v.Placement }).(LaunchTemplatePlacementOutput)
}

// The security group ID of instance. If this parameter is not specified, the default security group is bound.
func (o LaunchTemplateOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.
func (o LaunchTemplateOutput) SystemDisk() LaunchTemplateSystemDiskOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateSystemDiskOutput { return v.SystemDisk }).(LaunchTemplateSystemDiskOutput)
}

// Tag description list.
func (o LaunchTemplateOutput) TagSpecifications() LaunchTemplateTagSpecificationArrayOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateTagSpecificationArrayOutput { return v.TagSpecifications }).(LaunchTemplateTagSpecificationArrayOutput)
}

// Tag description list.
func (o LaunchTemplateOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The data of users.
func (o LaunchTemplateOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplate) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// The configuration information of VPC. If this parameter is not specified, the basic network is used by default.
func (o LaunchTemplateOutput) VirtualPrivateCloud() LaunchTemplateVirtualPrivateCloudOutput {
	return o.ApplyT(func(v *LaunchTemplate) LaunchTemplateVirtualPrivateCloudOutput { return v.VirtualPrivateCloud }).(LaunchTemplateVirtualPrivateCloudOutput)
}

type LaunchTemplateArrayOutput struct{ *pulumi.OutputState }

func (LaunchTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LaunchTemplate)(nil)).Elem()
}

func (o LaunchTemplateArrayOutput) ToLaunchTemplateArrayOutput() LaunchTemplateArrayOutput {
	return o
}

func (o LaunchTemplateArrayOutput) ToLaunchTemplateArrayOutputWithContext(ctx context.Context) LaunchTemplateArrayOutput {
	return o
}

func (o LaunchTemplateArrayOutput) Index(i pulumi.IntInput) LaunchTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LaunchTemplate {
		return vs[0].([]*LaunchTemplate)[vs[1].(int)]
	}).(LaunchTemplateOutput)
}

type LaunchTemplateMapOutput struct{ *pulumi.OutputState }

func (LaunchTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LaunchTemplate)(nil)).Elem()
}

func (o LaunchTemplateMapOutput) ToLaunchTemplateMapOutput() LaunchTemplateMapOutput {
	return o
}

func (o LaunchTemplateMapOutput) ToLaunchTemplateMapOutputWithContext(ctx context.Context) LaunchTemplateMapOutput {
	return o
}

func (o LaunchTemplateMapOutput) MapIndex(k pulumi.StringInput) LaunchTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LaunchTemplate {
		return vs[0].(map[string]*LaunchTemplate)[vs[1].(string)]
	}).(LaunchTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateInput)(nil)).Elem(), &LaunchTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateArrayInput)(nil)).Elem(), LaunchTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateMapInput)(nil)).Elem(), LaunchTemplateMap{})
	pulumi.RegisterOutputType(LaunchTemplateOutput{})
	pulumi.RegisterOutputType(LaunchTemplateArrayOutput{})
	pulumi.RegisterOutputType(LaunchTemplateMapOutput{})
}
