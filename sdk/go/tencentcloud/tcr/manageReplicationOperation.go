// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tcr manageReplicationOperation
type ManageReplicationOperation struct {
	pulumi.CustomResourceState

	// rule description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// the region ID of the target instance, such as Guangzhou is 1.
	DestinationRegionId pulumi.IntPtrOutput `pulumi:"destinationRegionId"`
	// copy destination instance Id.
	DestinationRegistryId pulumi.StringOutput `pulumi:"destinationRegistryId"`
	// enable synchronization of configuration items across master account instances.
	PeerReplicationOption ManageReplicationOperationPeerReplicationOptionPtrOutput `pulumi:"peerReplicationOption"`
	// synchronization rules.
	Rule ManageReplicationOperationRuleOutput `pulumi:"rule"`
	// copy source instance Id.
	SourceRegistryId pulumi.StringOutput `pulumi:"sourceRegistryId"`
}

// NewManageReplicationOperation registers a new resource with the given unique name, arguments, and options.
func NewManageReplicationOperation(ctx *pulumi.Context,
	name string, args *ManageReplicationOperationArgs, opts ...pulumi.ResourceOption) (*ManageReplicationOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationRegistryId == nil {
		return nil, errors.New("invalid value for required argument 'DestinationRegistryId'")
	}
	if args.Rule == nil {
		return nil, errors.New("invalid value for required argument 'Rule'")
	}
	if args.SourceRegistryId == nil {
		return nil, errors.New("invalid value for required argument 'SourceRegistryId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ManageReplicationOperation
	err := ctx.RegisterResource("tencentcloud:Tcr/manageReplicationOperation:ManageReplicationOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManageReplicationOperation gets an existing ManageReplicationOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManageReplicationOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManageReplicationOperationState, opts ...pulumi.ResourceOption) (*ManageReplicationOperation, error) {
	var resource ManageReplicationOperation
	err := ctx.ReadResource("tencentcloud:Tcr/manageReplicationOperation:ManageReplicationOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManageReplicationOperation resources.
type manageReplicationOperationState struct {
	// rule description.
	Description *string `pulumi:"description"`
	// the region ID of the target instance, such as Guangzhou is 1.
	DestinationRegionId *int `pulumi:"destinationRegionId"`
	// copy destination instance Id.
	DestinationRegistryId *string `pulumi:"destinationRegistryId"`
	// enable synchronization of configuration items across master account instances.
	PeerReplicationOption *ManageReplicationOperationPeerReplicationOption `pulumi:"peerReplicationOption"`
	// synchronization rules.
	Rule *ManageReplicationOperationRule `pulumi:"rule"`
	// copy source instance Id.
	SourceRegistryId *string `pulumi:"sourceRegistryId"`
}

type ManageReplicationOperationState struct {
	// rule description.
	Description pulumi.StringPtrInput
	// the region ID of the target instance, such as Guangzhou is 1.
	DestinationRegionId pulumi.IntPtrInput
	// copy destination instance Id.
	DestinationRegistryId pulumi.StringPtrInput
	// enable synchronization of configuration items across master account instances.
	PeerReplicationOption ManageReplicationOperationPeerReplicationOptionPtrInput
	// synchronization rules.
	Rule ManageReplicationOperationRulePtrInput
	// copy source instance Id.
	SourceRegistryId pulumi.StringPtrInput
}

func (ManageReplicationOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*manageReplicationOperationState)(nil)).Elem()
}

type manageReplicationOperationArgs struct {
	// rule description.
	Description *string `pulumi:"description"`
	// the region ID of the target instance, such as Guangzhou is 1.
	DestinationRegionId *int `pulumi:"destinationRegionId"`
	// copy destination instance Id.
	DestinationRegistryId string `pulumi:"destinationRegistryId"`
	// enable synchronization of configuration items across master account instances.
	PeerReplicationOption *ManageReplicationOperationPeerReplicationOption `pulumi:"peerReplicationOption"`
	// synchronization rules.
	Rule ManageReplicationOperationRule `pulumi:"rule"`
	// copy source instance Id.
	SourceRegistryId string `pulumi:"sourceRegistryId"`
}

// The set of arguments for constructing a ManageReplicationOperation resource.
type ManageReplicationOperationArgs struct {
	// rule description.
	Description pulumi.StringPtrInput
	// the region ID of the target instance, such as Guangzhou is 1.
	DestinationRegionId pulumi.IntPtrInput
	// copy destination instance Id.
	DestinationRegistryId pulumi.StringInput
	// enable synchronization of configuration items across master account instances.
	PeerReplicationOption ManageReplicationOperationPeerReplicationOptionPtrInput
	// synchronization rules.
	Rule ManageReplicationOperationRuleInput
	// copy source instance Id.
	SourceRegistryId pulumi.StringInput
}

func (ManageReplicationOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*manageReplicationOperationArgs)(nil)).Elem()
}

type ManageReplicationOperationInput interface {
	pulumi.Input

	ToManageReplicationOperationOutput() ManageReplicationOperationOutput
	ToManageReplicationOperationOutputWithContext(ctx context.Context) ManageReplicationOperationOutput
}

func (*ManageReplicationOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ManageReplicationOperation)(nil)).Elem()
}

func (i *ManageReplicationOperation) ToManageReplicationOperationOutput() ManageReplicationOperationOutput {
	return i.ToManageReplicationOperationOutputWithContext(context.Background())
}

func (i *ManageReplicationOperation) ToManageReplicationOperationOutputWithContext(ctx context.Context) ManageReplicationOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManageReplicationOperationOutput)
}

// ManageReplicationOperationArrayInput is an input type that accepts ManageReplicationOperationArray and ManageReplicationOperationArrayOutput values.
// You can construct a concrete instance of `ManageReplicationOperationArrayInput` via:
//
//          ManageReplicationOperationArray{ ManageReplicationOperationArgs{...} }
type ManageReplicationOperationArrayInput interface {
	pulumi.Input

	ToManageReplicationOperationArrayOutput() ManageReplicationOperationArrayOutput
	ToManageReplicationOperationArrayOutputWithContext(context.Context) ManageReplicationOperationArrayOutput
}

type ManageReplicationOperationArray []ManageReplicationOperationInput

func (ManageReplicationOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManageReplicationOperation)(nil)).Elem()
}

func (i ManageReplicationOperationArray) ToManageReplicationOperationArrayOutput() ManageReplicationOperationArrayOutput {
	return i.ToManageReplicationOperationArrayOutputWithContext(context.Background())
}

func (i ManageReplicationOperationArray) ToManageReplicationOperationArrayOutputWithContext(ctx context.Context) ManageReplicationOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManageReplicationOperationArrayOutput)
}

// ManageReplicationOperationMapInput is an input type that accepts ManageReplicationOperationMap and ManageReplicationOperationMapOutput values.
// You can construct a concrete instance of `ManageReplicationOperationMapInput` via:
//
//          ManageReplicationOperationMap{ "key": ManageReplicationOperationArgs{...} }
type ManageReplicationOperationMapInput interface {
	pulumi.Input

	ToManageReplicationOperationMapOutput() ManageReplicationOperationMapOutput
	ToManageReplicationOperationMapOutputWithContext(context.Context) ManageReplicationOperationMapOutput
}

type ManageReplicationOperationMap map[string]ManageReplicationOperationInput

func (ManageReplicationOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManageReplicationOperation)(nil)).Elem()
}

func (i ManageReplicationOperationMap) ToManageReplicationOperationMapOutput() ManageReplicationOperationMapOutput {
	return i.ToManageReplicationOperationMapOutputWithContext(context.Background())
}

func (i ManageReplicationOperationMap) ToManageReplicationOperationMapOutputWithContext(ctx context.Context) ManageReplicationOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManageReplicationOperationMapOutput)
}

type ManageReplicationOperationOutput struct{ *pulumi.OutputState }

func (ManageReplicationOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManageReplicationOperation)(nil)).Elem()
}

func (o ManageReplicationOperationOutput) ToManageReplicationOperationOutput() ManageReplicationOperationOutput {
	return o
}

func (o ManageReplicationOperationOutput) ToManageReplicationOperationOutputWithContext(ctx context.Context) ManageReplicationOperationOutput {
	return o
}

// rule description.
func (o ManageReplicationOperationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManageReplicationOperation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// the region ID of the target instance, such as Guangzhou is 1.
func (o ManageReplicationOperationOutput) DestinationRegionId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManageReplicationOperation) pulumi.IntPtrOutput { return v.DestinationRegionId }).(pulumi.IntPtrOutput)
}

// copy destination instance Id.
func (o ManageReplicationOperationOutput) DestinationRegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManageReplicationOperation) pulumi.StringOutput { return v.DestinationRegistryId }).(pulumi.StringOutput)
}

// enable synchronization of configuration items across master account instances.
func (o ManageReplicationOperationOutput) PeerReplicationOption() ManageReplicationOperationPeerReplicationOptionPtrOutput {
	return o.ApplyT(func(v *ManageReplicationOperation) ManageReplicationOperationPeerReplicationOptionPtrOutput {
		return v.PeerReplicationOption
	}).(ManageReplicationOperationPeerReplicationOptionPtrOutput)
}

// synchronization rules.
func (o ManageReplicationOperationOutput) Rule() ManageReplicationOperationRuleOutput {
	return o.ApplyT(func(v *ManageReplicationOperation) ManageReplicationOperationRuleOutput { return v.Rule }).(ManageReplicationOperationRuleOutput)
}

// copy source instance Id.
func (o ManageReplicationOperationOutput) SourceRegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManageReplicationOperation) pulumi.StringOutput { return v.SourceRegistryId }).(pulumi.StringOutput)
}

type ManageReplicationOperationArrayOutput struct{ *pulumi.OutputState }

func (ManageReplicationOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManageReplicationOperation)(nil)).Elem()
}

func (o ManageReplicationOperationArrayOutput) ToManageReplicationOperationArrayOutput() ManageReplicationOperationArrayOutput {
	return o
}

func (o ManageReplicationOperationArrayOutput) ToManageReplicationOperationArrayOutputWithContext(ctx context.Context) ManageReplicationOperationArrayOutput {
	return o
}

func (o ManageReplicationOperationArrayOutput) Index(i pulumi.IntInput) ManageReplicationOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManageReplicationOperation {
		return vs[0].([]*ManageReplicationOperation)[vs[1].(int)]
	}).(ManageReplicationOperationOutput)
}

type ManageReplicationOperationMapOutput struct{ *pulumi.OutputState }

func (ManageReplicationOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManageReplicationOperation)(nil)).Elem()
}

func (o ManageReplicationOperationMapOutput) ToManageReplicationOperationMapOutput() ManageReplicationOperationMapOutput {
	return o
}

func (o ManageReplicationOperationMapOutput) ToManageReplicationOperationMapOutputWithContext(ctx context.Context) ManageReplicationOperationMapOutput {
	return o
}

func (o ManageReplicationOperationMapOutput) MapIndex(k pulumi.StringInput) ManageReplicationOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManageReplicationOperation {
		return vs[0].(map[string]*ManageReplicationOperation)[vs[1].(string)]
	}).(ManageReplicationOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManageReplicationOperationInput)(nil)).Elem(), &ManageReplicationOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManageReplicationOperationArrayInput)(nil)).Elem(), ManageReplicationOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManageReplicationOperationMapInput)(nil)).Elem(), ManageReplicationOperationMap{})
	pulumi.RegisterOutputType(ManageReplicationOperationOutput{})
	pulumi.RegisterOutputType(ManageReplicationOperationArrayOutput{})
	pulumi.RegisterOutputType(ManageReplicationOperationMapOutput{})
}
