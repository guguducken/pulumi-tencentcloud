// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver renewDbInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Sqlserver.NewRenewDbInstance(ctx, "renewDbInstance", &Sqlserver.RenewDbInstanceArgs{
// 			InstanceId: pulumi.String("mssql-i1z41iwd"),
// 			Period:     pulumi.Int(1),
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
// sqlserver renew_db_instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance renew_db_instance renew_db_instance_id
// ```
type RenewDbInstance struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// How many months to renew, the value range is 1-48, the default is 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
}

// NewRenewDbInstance registers a new resource with the given unique name, arguments, and options.
func NewRenewDbInstance(ctx *pulumi.Context,
	name string, args *RenewDbInstanceArgs, opts ...pulumi.ResourceOption) (*RenewDbInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RenewDbInstance
	err := ctx.RegisterResource("tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRenewDbInstance gets an existing RenewDbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRenewDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RenewDbInstanceState, opts ...pulumi.ResourceOption) (*RenewDbInstance, error) {
	var resource RenewDbInstance
	err := ctx.ReadResource("tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RenewDbInstance resources.
type renewDbInstanceState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// How many months to renew, the value range is 1-48, the default is 1.
	Period *int `pulumi:"period"`
}

type RenewDbInstanceState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// How many months to renew, the value range is 1-48, the default is 1.
	Period pulumi.IntPtrInput
}

func (RenewDbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*renewDbInstanceState)(nil)).Elem()
}

type renewDbInstanceArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// How many months to renew, the value range is 1-48, the default is 1.
	Period *int `pulumi:"period"`
}

// The set of arguments for constructing a RenewDbInstance resource.
type RenewDbInstanceArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
	// How many months to renew, the value range is 1-48, the default is 1.
	Period pulumi.IntPtrInput
}

func (RenewDbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*renewDbInstanceArgs)(nil)).Elem()
}

type RenewDbInstanceInput interface {
	pulumi.Input

	ToRenewDbInstanceOutput() RenewDbInstanceOutput
	ToRenewDbInstanceOutputWithContext(ctx context.Context) RenewDbInstanceOutput
}

func (*RenewDbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewDbInstance)(nil)).Elem()
}

func (i *RenewDbInstance) ToRenewDbInstanceOutput() RenewDbInstanceOutput {
	return i.ToRenewDbInstanceOutputWithContext(context.Background())
}

func (i *RenewDbInstance) ToRenewDbInstanceOutputWithContext(ctx context.Context) RenewDbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewDbInstanceOutput)
}

// RenewDbInstanceArrayInput is an input type that accepts RenewDbInstanceArray and RenewDbInstanceArrayOutput values.
// You can construct a concrete instance of `RenewDbInstanceArrayInput` via:
//
//          RenewDbInstanceArray{ RenewDbInstanceArgs{...} }
type RenewDbInstanceArrayInput interface {
	pulumi.Input

	ToRenewDbInstanceArrayOutput() RenewDbInstanceArrayOutput
	ToRenewDbInstanceArrayOutputWithContext(context.Context) RenewDbInstanceArrayOutput
}

type RenewDbInstanceArray []RenewDbInstanceInput

func (RenewDbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewDbInstance)(nil)).Elem()
}

func (i RenewDbInstanceArray) ToRenewDbInstanceArrayOutput() RenewDbInstanceArrayOutput {
	return i.ToRenewDbInstanceArrayOutputWithContext(context.Background())
}

func (i RenewDbInstanceArray) ToRenewDbInstanceArrayOutputWithContext(ctx context.Context) RenewDbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewDbInstanceArrayOutput)
}

// RenewDbInstanceMapInput is an input type that accepts RenewDbInstanceMap and RenewDbInstanceMapOutput values.
// You can construct a concrete instance of `RenewDbInstanceMapInput` via:
//
//          RenewDbInstanceMap{ "key": RenewDbInstanceArgs{...} }
type RenewDbInstanceMapInput interface {
	pulumi.Input

	ToRenewDbInstanceMapOutput() RenewDbInstanceMapOutput
	ToRenewDbInstanceMapOutputWithContext(context.Context) RenewDbInstanceMapOutput
}

type RenewDbInstanceMap map[string]RenewDbInstanceInput

func (RenewDbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewDbInstance)(nil)).Elem()
}

func (i RenewDbInstanceMap) ToRenewDbInstanceMapOutput() RenewDbInstanceMapOutput {
	return i.ToRenewDbInstanceMapOutputWithContext(context.Background())
}

func (i RenewDbInstanceMap) ToRenewDbInstanceMapOutputWithContext(ctx context.Context) RenewDbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewDbInstanceMapOutput)
}

type RenewDbInstanceOutput struct{ *pulumi.OutputState }

func (RenewDbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewDbInstance)(nil)).Elem()
}

func (o RenewDbInstanceOutput) ToRenewDbInstanceOutput() RenewDbInstanceOutput {
	return o
}

func (o RenewDbInstanceOutput) ToRenewDbInstanceOutputWithContext(ctx context.Context) RenewDbInstanceOutput {
	return o
}

// Instance ID.
func (o RenewDbInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RenewDbInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// How many months to renew, the value range is 1-48, the default is 1.
func (o RenewDbInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RenewDbInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

type RenewDbInstanceArrayOutput struct{ *pulumi.OutputState }

func (RenewDbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewDbInstance)(nil)).Elem()
}

func (o RenewDbInstanceArrayOutput) ToRenewDbInstanceArrayOutput() RenewDbInstanceArrayOutput {
	return o
}

func (o RenewDbInstanceArrayOutput) ToRenewDbInstanceArrayOutputWithContext(ctx context.Context) RenewDbInstanceArrayOutput {
	return o
}

func (o RenewDbInstanceArrayOutput) Index(i pulumi.IntInput) RenewDbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RenewDbInstance {
		return vs[0].([]*RenewDbInstance)[vs[1].(int)]
	}).(RenewDbInstanceOutput)
}

type RenewDbInstanceMapOutput struct{ *pulumi.OutputState }

func (RenewDbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewDbInstance)(nil)).Elem()
}

func (o RenewDbInstanceMapOutput) ToRenewDbInstanceMapOutput() RenewDbInstanceMapOutput {
	return o
}

func (o RenewDbInstanceMapOutput) ToRenewDbInstanceMapOutputWithContext(ctx context.Context) RenewDbInstanceMapOutput {
	return o
}

func (o RenewDbInstanceMapOutput) MapIndex(k pulumi.StringInput) RenewDbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RenewDbInstance {
		return vs[0].(map[string]*RenewDbInstance)[vs[1].(string)]
	}).(RenewDbInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RenewDbInstanceInput)(nil)).Elem(), &RenewDbInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewDbInstanceArrayInput)(nil)).Elem(), RenewDbInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewDbInstanceMapInput)(nil)).Elem(), RenewDbInstanceMap{})
	pulumi.RegisterOutputType(RenewDbInstanceOutput{})
	pulumi.RegisterOutputType(RenewDbInstanceArrayOutput{})
	pulumi.RegisterOutputType(RenewDbInstanceMapOutput{})
}
