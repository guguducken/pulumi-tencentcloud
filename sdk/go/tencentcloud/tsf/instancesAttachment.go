// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tsf instancesAttachment
type InstancesAttachment struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Image feature ID list.
	FeatureIdLists pulumi.StringArrayOutput `pulumi:"featureIdLists"`
	// Operating system image ID.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Additional instance parameter information.
	InstanceAdvancedSettings InstancesAttachmentInstanceAdvancedSettingsPtrOutput `pulumi:"instanceAdvancedSettings"`
	// Cloud server ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF system image, M: Manual installation of agent.
	InstanceImportMode pulumi.StringPtrOutput `pulumi:"instanceImportMode"`
	// Associated key for system reinstallation.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// Image customization type.
	OsCustomizeType pulumi.StringPtrOutput `pulumi:"osCustomizeType"`
	// Operating system name.
	OsName pulumi.StringPtrOutput `pulumi:"osName"`
	// Reset system password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Security group.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Security group setting.
	SgId pulumi.StringPtrOutput `pulumi:"sgId"`
}

// NewInstancesAttachment registers a new resource with the given unique name, arguments, and options.
func NewInstancesAttachment(ctx *pulumi.Context,
	name string, args *InstancesAttachmentArgs, opts ...pulumi.ResourceOption) (*InstancesAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstancesAttachment
	err := ctx.RegisterResource("tencentcloud:Tsf/instancesAttachment:InstancesAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstancesAttachment gets an existing InstancesAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstancesAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancesAttachmentState, opts ...pulumi.ResourceOption) (*InstancesAttachment, error) {
	var resource InstancesAttachment
	err := ctx.ReadResource("tencentcloud:Tsf/instancesAttachment:InstancesAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstancesAttachment resources.
type instancesAttachmentState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Image feature ID list.
	FeatureIdLists []string `pulumi:"featureIdLists"`
	// Operating system image ID.
	ImageId *string `pulumi:"imageId"`
	// Additional instance parameter information.
	InstanceAdvancedSettings *InstancesAttachmentInstanceAdvancedSettings `pulumi:"instanceAdvancedSettings"`
	// Cloud server ID.
	InstanceId *string `pulumi:"instanceId"`
	// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF system image, M: Manual installation of agent.
	InstanceImportMode *string `pulumi:"instanceImportMode"`
	// Associated key for system reinstallation.
	KeyId *string `pulumi:"keyId"`
	// Image customization type.
	OsCustomizeType *string `pulumi:"osCustomizeType"`
	// Operating system name.
	OsName *string `pulumi:"osName"`
	// Reset system password.
	Password *string `pulumi:"password"`
	// Security group.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Security group setting.
	SgId *string `pulumi:"sgId"`
}

type InstancesAttachmentState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Image feature ID list.
	FeatureIdLists pulumi.StringArrayInput
	// Operating system image ID.
	ImageId pulumi.StringPtrInput
	// Additional instance parameter information.
	InstanceAdvancedSettings InstancesAttachmentInstanceAdvancedSettingsPtrInput
	// Cloud server ID.
	InstanceId pulumi.StringPtrInput
	// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF system image, M: Manual installation of agent.
	InstanceImportMode pulumi.StringPtrInput
	// Associated key for system reinstallation.
	KeyId pulumi.StringPtrInput
	// Image customization type.
	OsCustomizeType pulumi.StringPtrInput
	// Operating system name.
	OsName pulumi.StringPtrInput
	// Reset system password.
	Password pulumi.StringPtrInput
	// Security group.
	SecurityGroupIds pulumi.StringArrayInput
	// Security group setting.
	SgId pulumi.StringPtrInput
}

func (InstancesAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancesAttachmentState)(nil)).Elem()
}

type instancesAttachmentArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Image feature ID list.
	FeatureIdLists []string `pulumi:"featureIdLists"`
	// Operating system image ID.
	ImageId *string `pulumi:"imageId"`
	// Additional instance parameter information.
	InstanceAdvancedSettings *InstancesAttachmentInstanceAdvancedSettings `pulumi:"instanceAdvancedSettings"`
	// Cloud server ID.
	InstanceId string `pulumi:"instanceId"`
	// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF system image, M: Manual installation of agent.
	InstanceImportMode *string `pulumi:"instanceImportMode"`
	// Associated key for system reinstallation.
	KeyId *string `pulumi:"keyId"`
	// Image customization type.
	OsCustomizeType *string `pulumi:"osCustomizeType"`
	// Operating system name.
	OsName *string `pulumi:"osName"`
	// Reset system password.
	Password *string `pulumi:"password"`
	// Security group.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Security group setting.
	SgId *string `pulumi:"sgId"`
}

// The set of arguments for constructing a InstancesAttachment resource.
type InstancesAttachmentArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Image feature ID list.
	FeatureIdLists pulumi.StringArrayInput
	// Operating system image ID.
	ImageId pulumi.StringPtrInput
	// Additional instance parameter information.
	InstanceAdvancedSettings InstancesAttachmentInstanceAdvancedSettingsPtrInput
	// Cloud server ID.
	InstanceId pulumi.StringInput
	// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF system image, M: Manual installation of agent.
	InstanceImportMode pulumi.StringPtrInput
	// Associated key for system reinstallation.
	KeyId pulumi.StringPtrInput
	// Image customization type.
	OsCustomizeType pulumi.StringPtrInput
	// Operating system name.
	OsName pulumi.StringPtrInput
	// Reset system password.
	Password pulumi.StringPtrInput
	// Security group.
	SecurityGroupIds pulumi.StringArrayInput
	// Security group setting.
	SgId pulumi.StringPtrInput
}

func (InstancesAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancesAttachmentArgs)(nil)).Elem()
}

type InstancesAttachmentInput interface {
	pulumi.Input

	ToInstancesAttachmentOutput() InstancesAttachmentOutput
	ToInstancesAttachmentOutputWithContext(ctx context.Context) InstancesAttachmentOutput
}

func (*InstancesAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancesAttachment)(nil)).Elem()
}

func (i *InstancesAttachment) ToInstancesAttachmentOutput() InstancesAttachmentOutput {
	return i.ToInstancesAttachmentOutputWithContext(context.Background())
}

func (i *InstancesAttachment) ToInstancesAttachmentOutputWithContext(ctx context.Context) InstancesAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesAttachmentOutput)
}

// InstancesAttachmentArrayInput is an input type that accepts InstancesAttachmentArray and InstancesAttachmentArrayOutput values.
// You can construct a concrete instance of `InstancesAttachmentArrayInput` via:
//
//          InstancesAttachmentArray{ InstancesAttachmentArgs{...} }
type InstancesAttachmentArrayInput interface {
	pulumi.Input

	ToInstancesAttachmentArrayOutput() InstancesAttachmentArrayOutput
	ToInstancesAttachmentArrayOutputWithContext(context.Context) InstancesAttachmentArrayOutput
}

type InstancesAttachmentArray []InstancesAttachmentInput

func (InstancesAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancesAttachment)(nil)).Elem()
}

func (i InstancesAttachmentArray) ToInstancesAttachmentArrayOutput() InstancesAttachmentArrayOutput {
	return i.ToInstancesAttachmentArrayOutputWithContext(context.Background())
}

func (i InstancesAttachmentArray) ToInstancesAttachmentArrayOutputWithContext(ctx context.Context) InstancesAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesAttachmentArrayOutput)
}

// InstancesAttachmentMapInput is an input type that accepts InstancesAttachmentMap and InstancesAttachmentMapOutput values.
// You can construct a concrete instance of `InstancesAttachmentMapInput` via:
//
//          InstancesAttachmentMap{ "key": InstancesAttachmentArgs{...} }
type InstancesAttachmentMapInput interface {
	pulumi.Input

	ToInstancesAttachmentMapOutput() InstancesAttachmentMapOutput
	ToInstancesAttachmentMapOutputWithContext(context.Context) InstancesAttachmentMapOutput
}

type InstancesAttachmentMap map[string]InstancesAttachmentInput

func (InstancesAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancesAttachment)(nil)).Elem()
}

func (i InstancesAttachmentMap) ToInstancesAttachmentMapOutput() InstancesAttachmentMapOutput {
	return i.ToInstancesAttachmentMapOutputWithContext(context.Background())
}

func (i InstancesAttachmentMap) ToInstancesAttachmentMapOutputWithContext(ctx context.Context) InstancesAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesAttachmentMapOutput)
}

type InstancesAttachmentOutput struct{ *pulumi.OutputState }

func (InstancesAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancesAttachment)(nil)).Elem()
}

func (o InstancesAttachmentOutput) ToInstancesAttachmentOutput() InstancesAttachmentOutput {
	return o
}

func (o InstancesAttachmentOutput) ToInstancesAttachmentOutputWithContext(ctx context.Context) InstancesAttachmentOutput {
	return o
}

// Cluster ID.
func (o InstancesAttachmentOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Image feature ID list.
func (o InstancesAttachmentOutput) FeatureIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringArrayOutput { return v.FeatureIdLists }).(pulumi.StringArrayOutput)
}

// Operating system image ID.
func (o InstancesAttachmentOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.ImageId }).(pulumi.StringPtrOutput)
}

// Additional instance parameter information.
func (o InstancesAttachmentOutput) InstanceAdvancedSettings() InstancesAttachmentInstanceAdvancedSettingsPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) InstancesAttachmentInstanceAdvancedSettingsPtrOutput {
		return v.InstanceAdvancedSettings
	}).(InstancesAttachmentInstanceAdvancedSettingsPtrOutput)
}

// Cloud server ID.
func (o InstancesAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Cloud server import mode, required for virtual machine clusters, not required for container clusters. R: Reinstall TSF system image, M: Manual installation of agent.
func (o InstancesAttachmentOutput) InstanceImportMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.InstanceImportMode }).(pulumi.StringPtrOutput)
}

// Associated key for system reinstallation.
func (o InstancesAttachmentOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// Image customization type.
func (o InstancesAttachmentOutput) OsCustomizeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.OsCustomizeType }).(pulumi.StringPtrOutput)
}

// Operating system name.
func (o InstancesAttachmentOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.OsName }).(pulumi.StringPtrOutput)
}

// Reset system password.
func (o InstancesAttachmentOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Security group.
func (o InstancesAttachmentOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Security group setting.
func (o InstancesAttachmentOutput) SgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstancesAttachment) pulumi.StringPtrOutput { return v.SgId }).(pulumi.StringPtrOutput)
}

type InstancesAttachmentArrayOutput struct{ *pulumi.OutputState }

func (InstancesAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancesAttachment)(nil)).Elem()
}

func (o InstancesAttachmentArrayOutput) ToInstancesAttachmentArrayOutput() InstancesAttachmentArrayOutput {
	return o
}

func (o InstancesAttachmentArrayOutput) ToInstancesAttachmentArrayOutputWithContext(ctx context.Context) InstancesAttachmentArrayOutput {
	return o
}

func (o InstancesAttachmentArrayOutput) Index(i pulumi.IntInput) InstancesAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstancesAttachment {
		return vs[0].([]*InstancesAttachment)[vs[1].(int)]
	}).(InstancesAttachmentOutput)
}

type InstancesAttachmentMapOutput struct{ *pulumi.OutputState }

func (InstancesAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancesAttachment)(nil)).Elem()
}

func (o InstancesAttachmentMapOutput) ToInstancesAttachmentMapOutput() InstancesAttachmentMapOutput {
	return o
}

func (o InstancesAttachmentMapOutput) ToInstancesAttachmentMapOutputWithContext(ctx context.Context) InstancesAttachmentMapOutput {
	return o
}

func (o InstancesAttachmentMapOutput) MapIndex(k pulumi.StringInput) InstancesAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstancesAttachment {
		return vs[0].(map[string]*InstancesAttachment)[vs[1].(string)]
	}).(InstancesAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesAttachmentInput)(nil)).Elem(), &InstancesAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesAttachmentArrayInput)(nil)).Elem(), InstancesAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesAttachmentMapInput)(nil)).Elem(), InstancesAttachmentMap{})
	pulumi.RegisterOutputType(InstancesAttachmentOutput{})
	pulumi.RegisterOutputType(InstancesAttachmentArrayOutput{})
	pulumi.RegisterOutputType(InstancesAttachmentMapOutput{})
}
