// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package reserved

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a reserved instance resource.
//
// > **NOTE:** Reserved instance cannot be deleted and updated. The reserved instance still exist which can be extracted by reservedInstances data source when reserved instance is destroied.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Reserved"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Reserved.NewInstance(ctx, "ri", &Reserved.InstanceArgs{
// 			ConfigId:      pulumi.String("469043dd-28b9-4d89-b557-74f6a8326259"),
// 			InstanceCount: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Reserved instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Reserved/instance:Instance foo 6cc16e7c-47d7-4fae-9b44-ce5c0f59a920
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Configuration ID of the reserved instance.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// Expiry time of the RI.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Number of reserved instances to be purchased.
	InstanceCount pulumi.IntOutput `pulumi:"instanceCount"`
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName pulumi.StringPtrOutput `pulumi:"reservedInstanceName"`
	// Start time of the RI.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Status of the RI at the time of purchase.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.InstanceCount == nil {
		return nil, errors.New("invalid value for required argument 'InstanceCount'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Reserved/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("tencentcloud:Reserved/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Configuration ID of the reserved instance.
	ConfigId *string `pulumi:"configId"`
	// Expiry time of the RI.
	EndTime *string `pulumi:"endTime"`
	// Number of reserved instances to be purchased.
	InstanceCount *int `pulumi:"instanceCount"`
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName *string `pulumi:"reservedInstanceName"`
	// Start time of the RI.
	StartTime *string `pulumi:"startTime"`
	// Status of the RI at the time of purchase.
	Status *string `pulumi:"status"`
}

type InstanceState struct {
	// Configuration ID of the reserved instance.
	ConfigId pulumi.StringPtrInput
	// Expiry time of the RI.
	EndTime pulumi.StringPtrInput
	// Number of reserved instances to be purchased.
	InstanceCount pulumi.IntPtrInput
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName pulumi.StringPtrInput
	// Start time of the RI.
	StartTime pulumi.StringPtrInput
	// Status of the RI at the time of purchase.
	Status pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Configuration ID of the reserved instance.
	ConfigId string `pulumi:"configId"`
	// Number of reserved instances to be purchased.
	InstanceCount int `pulumi:"instanceCount"`
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName *string `pulumi:"reservedInstanceName"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Configuration ID of the reserved instance.
	ConfigId pulumi.StringInput
	// Number of reserved instances to be purchased.
	InstanceCount pulumi.IntInput
	// Reserved Instance display name.
	// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
	// - Up to 60 characters (including pattern strings) are supported.
	ReservedInstanceName pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Configuration ID of the reserved instance.
func (o InstanceOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// Expiry time of the RI.
func (o InstanceOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Number of reserved instances to be purchased.
func (o InstanceOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.InstanceCount }).(pulumi.IntOutput)
}

// Reserved Instance display name.
// - If you do not specify an instance display name, 'Unnamed' is displayed by default.
// - Up to 60 characters (including pattern strings) are supported.
func (o InstanceOutput) ReservedInstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ReservedInstanceName }).(pulumi.StringPtrOutput)
}

// Start time of the RI.
func (o InstanceOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Status of the RI at the time of purchase.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
