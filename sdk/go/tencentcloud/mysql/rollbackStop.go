// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql rollbackStop
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mysql.NewRollbackStop(ctx, "rollbackStop", &Mysql.RollbackStopArgs{
// 			InstanceId: pulumi.String("cdb-fitq5t9h"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RollbackStop struct {
	pulumi.CustomResourceState

	// Cloud database instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewRollbackStop registers a new resource with the given unique name, arguments, and options.
func NewRollbackStop(ctx *pulumi.Context,
	name string, args *RollbackStopArgs, opts ...pulumi.ResourceOption) (*RollbackStop, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RollbackStop
	err := ctx.RegisterResource("tencentcloud:Mysql/rollbackStop:RollbackStop", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRollbackStop gets an existing RollbackStop resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRollbackStop(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RollbackStopState, opts ...pulumi.ResourceOption) (*RollbackStop, error) {
	var resource RollbackStop
	err := ctx.ReadResource("tencentcloud:Mysql/rollbackStop:RollbackStop", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RollbackStop resources.
type rollbackStopState struct {
	// Cloud database instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type RollbackStopState struct {
	// Cloud database instance ID.
	InstanceId pulumi.StringPtrInput
}

func (RollbackStopState) ElementType() reflect.Type {
	return reflect.TypeOf((*rollbackStopState)(nil)).Elem()
}

type rollbackStopArgs struct {
	// Cloud database instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a RollbackStop resource.
type RollbackStopArgs struct {
	// Cloud database instance ID.
	InstanceId pulumi.StringInput
}

func (RollbackStopArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rollbackStopArgs)(nil)).Elem()
}

type RollbackStopInput interface {
	pulumi.Input

	ToRollbackStopOutput() RollbackStopOutput
	ToRollbackStopOutputWithContext(ctx context.Context) RollbackStopOutput
}

func (*RollbackStop) ElementType() reflect.Type {
	return reflect.TypeOf((**RollbackStop)(nil)).Elem()
}

func (i *RollbackStop) ToRollbackStopOutput() RollbackStopOutput {
	return i.ToRollbackStopOutputWithContext(context.Background())
}

func (i *RollbackStop) ToRollbackStopOutputWithContext(ctx context.Context) RollbackStopOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollbackStopOutput)
}

// RollbackStopArrayInput is an input type that accepts RollbackStopArray and RollbackStopArrayOutput values.
// You can construct a concrete instance of `RollbackStopArrayInput` via:
//
//          RollbackStopArray{ RollbackStopArgs{...} }
type RollbackStopArrayInput interface {
	pulumi.Input

	ToRollbackStopArrayOutput() RollbackStopArrayOutput
	ToRollbackStopArrayOutputWithContext(context.Context) RollbackStopArrayOutput
}

type RollbackStopArray []RollbackStopInput

func (RollbackStopArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RollbackStop)(nil)).Elem()
}

func (i RollbackStopArray) ToRollbackStopArrayOutput() RollbackStopArrayOutput {
	return i.ToRollbackStopArrayOutputWithContext(context.Background())
}

func (i RollbackStopArray) ToRollbackStopArrayOutputWithContext(ctx context.Context) RollbackStopArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollbackStopArrayOutput)
}

// RollbackStopMapInput is an input type that accepts RollbackStopMap and RollbackStopMapOutput values.
// You can construct a concrete instance of `RollbackStopMapInput` via:
//
//          RollbackStopMap{ "key": RollbackStopArgs{...} }
type RollbackStopMapInput interface {
	pulumi.Input

	ToRollbackStopMapOutput() RollbackStopMapOutput
	ToRollbackStopMapOutputWithContext(context.Context) RollbackStopMapOutput
}

type RollbackStopMap map[string]RollbackStopInput

func (RollbackStopMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RollbackStop)(nil)).Elem()
}

func (i RollbackStopMap) ToRollbackStopMapOutput() RollbackStopMapOutput {
	return i.ToRollbackStopMapOutputWithContext(context.Background())
}

func (i RollbackStopMap) ToRollbackStopMapOutputWithContext(ctx context.Context) RollbackStopMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollbackStopMapOutput)
}

type RollbackStopOutput struct{ *pulumi.OutputState }

func (RollbackStopOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollbackStop)(nil)).Elem()
}

func (o RollbackStopOutput) ToRollbackStopOutput() RollbackStopOutput {
	return o
}

func (o RollbackStopOutput) ToRollbackStopOutputWithContext(ctx context.Context) RollbackStopOutput {
	return o
}

// Cloud database instance ID.
func (o RollbackStopOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RollbackStop) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type RollbackStopArrayOutput struct{ *pulumi.OutputState }

func (RollbackStopArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RollbackStop)(nil)).Elem()
}

func (o RollbackStopArrayOutput) ToRollbackStopArrayOutput() RollbackStopArrayOutput {
	return o
}

func (o RollbackStopArrayOutput) ToRollbackStopArrayOutputWithContext(ctx context.Context) RollbackStopArrayOutput {
	return o
}

func (o RollbackStopArrayOutput) Index(i pulumi.IntInput) RollbackStopOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RollbackStop {
		return vs[0].([]*RollbackStop)[vs[1].(int)]
	}).(RollbackStopOutput)
}

type RollbackStopMapOutput struct{ *pulumi.OutputState }

func (RollbackStopMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RollbackStop)(nil)).Elem()
}

func (o RollbackStopMapOutput) ToRollbackStopMapOutput() RollbackStopMapOutput {
	return o
}

func (o RollbackStopMapOutput) ToRollbackStopMapOutputWithContext(ctx context.Context) RollbackStopMapOutput {
	return o
}

func (o RollbackStopMapOutput) MapIndex(k pulumi.StringInput) RollbackStopOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RollbackStop {
		return vs[0].(map[string]*RollbackStop)[vs[1].(string)]
	}).(RollbackStopOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RollbackStopInput)(nil)).Elem(), &RollbackStop{})
	pulumi.RegisterInputType(reflect.TypeOf((*RollbackStopArrayInput)(nil)).Elem(), RollbackStopArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RollbackStopMapInput)(nil)).Elem(), RollbackStopMap{})
	pulumi.RegisterOutputType(RollbackStopOutput{})
	pulumi.RegisterOutputType(RollbackStopArrayOutput{})
	pulumi.RegisterOutputType(RollbackStopMapOutput{})
}
