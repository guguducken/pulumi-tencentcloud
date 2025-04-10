// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a postgresql modifyAccountRemarkOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Postgresql.NewModifyAccountRemarkOperation(ctx, "modifyAccountRemarkOperation", &Postgresql.ModifyAccountRemarkOperationArgs{
// 			DbInstanceId: pulumi.Any(local.Pgsql_id),
// 			UserName:     pulumi.String("root"),
// 			Remark:       pulumi.String("hello_world"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ModifyAccountRemarkOperation struct {
	pulumi.CustomResourceState

	// Instance ID in the format of postgres-4wdeb0zv.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// New remarks corresponding to user `UserName`.
	Remark pulumi.StringOutput `pulumi:"remark"`
	// Instance username.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewModifyAccountRemarkOperation registers a new resource with the given unique name, arguments, and options.
func NewModifyAccountRemarkOperation(ctx *pulumi.Context,
	name string, args *ModifyAccountRemarkOperationArgs, opts ...pulumi.ResourceOption) (*ModifyAccountRemarkOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.Remark == nil {
		return nil, errors.New("invalid value for required argument 'Remark'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ModifyAccountRemarkOperation
	err := ctx.RegisterResource("tencentcloud:Postgresql/modifyAccountRemarkOperation:ModifyAccountRemarkOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModifyAccountRemarkOperation gets an existing ModifyAccountRemarkOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModifyAccountRemarkOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModifyAccountRemarkOperationState, opts ...pulumi.ResourceOption) (*ModifyAccountRemarkOperation, error) {
	var resource ModifyAccountRemarkOperation
	err := ctx.ReadResource("tencentcloud:Postgresql/modifyAccountRemarkOperation:ModifyAccountRemarkOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModifyAccountRemarkOperation resources.
type modifyAccountRemarkOperationState struct {
	// Instance ID in the format of postgres-4wdeb0zv.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// New remarks corresponding to user `UserName`.
	Remark *string `pulumi:"remark"`
	// Instance username.
	UserName *string `pulumi:"userName"`
}

type ModifyAccountRemarkOperationState struct {
	// Instance ID in the format of postgres-4wdeb0zv.
	DbInstanceId pulumi.StringPtrInput
	// New remarks corresponding to user `UserName`.
	Remark pulumi.StringPtrInput
	// Instance username.
	UserName pulumi.StringPtrInput
}

func (ModifyAccountRemarkOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*modifyAccountRemarkOperationState)(nil)).Elem()
}

type modifyAccountRemarkOperationArgs struct {
	// Instance ID in the format of postgres-4wdeb0zv.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// New remarks corresponding to user `UserName`.
	Remark string `pulumi:"remark"`
	// Instance username.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a ModifyAccountRemarkOperation resource.
type ModifyAccountRemarkOperationArgs struct {
	// Instance ID in the format of postgres-4wdeb0zv.
	DbInstanceId pulumi.StringInput
	// New remarks corresponding to user `UserName`.
	Remark pulumi.StringInput
	// Instance username.
	UserName pulumi.StringInput
}

func (ModifyAccountRemarkOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modifyAccountRemarkOperationArgs)(nil)).Elem()
}

type ModifyAccountRemarkOperationInput interface {
	pulumi.Input

	ToModifyAccountRemarkOperationOutput() ModifyAccountRemarkOperationOutput
	ToModifyAccountRemarkOperationOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationOutput
}

func (*ModifyAccountRemarkOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ModifyAccountRemarkOperation)(nil)).Elem()
}

func (i *ModifyAccountRemarkOperation) ToModifyAccountRemarkOperationOutput() ModifyAccountRemarkOperationOutput {
	return i.ToModifyAccountRemarkOperationOutputWithContext(context.Background())
}

func (i *ModifyAccountRemarkOperation) ToModifyAccountRemarkOperationOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyAccountRemarkOperationOutput)
}

// ModifyAccountRemarkOperationArrayInput is an input type that accepts ModifyAccountRemarkOperationArray and ModifyAccountRemarkOperationArrayOutput values.
// You can construct a concrete instance of `ModifyAccountRemarkOperationArrayInput` via:
//
//          ModifyAccountRemarkOperationArray{ ModifyAccountRemarkOperationArgs{...} }
type ModifyAccountRemarkOperationArrayInput interface {
	pulumi.Input

	ToModifyAccountRemarkOperationArrayOutput() ModifyAccountRemarkOperationArrayOutput
	ToModifyAccountRemarkOperationArrayOutputWithContext(context.Context) ModifyAccountRemarkOperationArrayOutput
}

type ModifyAccountRemarkOperationArray []ModifyAccountRemarkOperationInput

func (ModifyAccountRemarkOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ModifyAccountRemarkOperation)(nil)).Elem()
}

func (i ModifyAccountRemarkOperationArray) ToModifyAccountRemarkOperationArrayOutput() ModifyAccountRemarkOperationArrayOutput {
	return i.ToModifyAccountRemarkOperationArrayOutputWithContext(context.Background())
}

func (i ModifyAccountRemarkOperationArray) ToModifyAccountRemarkOperationArrayOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyAccountRemarkOperationArrayOutput)
}

// ModifyAccountRemarkOperationMapInput is an input type that accepts ModifyAccountRemarkOperationMap and ModifyAccountRemarkOperationMapOutput values.
// You can construct a concrete instance of `ModifyAccountRemarkOperationMapInput` via:
//
//          ModifyAccountRemarkOperationMap{ "key": ModifyAccountRemarkOperationArgs{...} }
type ModifyAccountRemarkOperationMapInput interface {
	pulumi.Input

	ToModifyAccountRemarkOperationMapOutput() ModifyAccountRemarkOperationMapOutput
	ToModifyAccountRemarkOperationMapOutputWithContext(context.Context) ModifyAccountRemarkOperationMapOutput
}

type ModifyAccountRemarkOperationMap map[string]ModifyAccountRemarkOperationInput

func (ModifyAccountRemarkOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ModifyAccountRemarkOperation)(nil)).Elem()
}

func (i ModifyAccountRemarkOperationMap) ToModifyAccountRemarkOperationMapOutput() ModifyAccountRemarkOperationMapOutput {
	return i.ToModifyAccountRemarkOperationMapOutputWithContext(context.Background())
}

func (i ModifyAccountRemarkOperationMap) ToModifyAccountRemarkOperationMapOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyAccountRemarkOperationMapOutput)
}

type ModifyAccountRemarkOperationOutput struct{ *pulumi.OutputState }

func (ModifyAccountRemarkOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModifyAccountRemarkOperation)(nil)).Elem()
}

func (o ModifyAccountRemarkOperationOutput) ToModifyAccountRemarkOperationOutput() ModifyAccountRemarkOperationOutput {
	return o
}

func (o ModifyAccountRemarkOperationOutput) ToModifyAccountRemarkOperationOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationOutput {
	return o
}

// Instance ID in the format of postgres-4wdeb0zv.
func (o ModifyAccountRemarkOperationOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ModifyAccountRemarkOperation) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// New remarks corresponding to user `UserName`.
func (o ModifyAccountRemarkOperationOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v *ModifyAccountRemarkOperation) pulumi.StringOutput { return v.Remark }).(pulumi.StringOutput)
}

// Instance username.
func (o ModifyAccountRemarkOperationOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ModifyAccountRemarkOperation) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type ModifyAccountRemarkOperationArrayOutput struct{ *pulumi.OutputState }

func (ModifyAccountRemarkOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ModifyAccountRemarkOperation)(nil)).Elem()
}

func (o ModifyAccountRemarkOperationArrayOutput) ToModifyAccountRemarkOperationArrayOutput() ModifyAccountRemarkOperationArrayOutput {
	return o
}

func (o ModifyAccountRemarkOperationArrayOutput) ToModifyAccountRemarkOperationArrayOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationArrayOutput {
	return o
}

func (o ModifyAccountRemarkOperationArrayOutput) Index(i pulumi.IntInput) ModifyAccountRemarkOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ModifyAccountRemarkOperation {
		return vs[0].([]*ModifyAccountRemarkOperation)[vs[1].(int)]
	}).(ModifyAccountRemarkOperationOutput)
}

type ModifyAccountRemarkOperationMapOutput struct{ *pulumi.OutputState }

func (ModifyAccountRemarkOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ModifyAccountRemarkOperation)(nil)).Elem()
}

func (o ModifyAccountRemarkOperationMapOutput) ToModifyAccountRemarkOperationMapOutput() ModifyAccountRemarkOperationMapOutput {
	return o
}

func (o ModifyAccountRemarkOperationMapOutput) ToModifyAccountRemarkOperationMapOutputWithContext(ctx context.Context) ModifyAccountRemarkOperationMapOutput {
	return o
}

func (o ModifyAccountRemarkOperationMapOutput) MapIndex(k pulumi.StringInput) ModifyAccountRemarkOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ModifyAccountRemarkOperation {
		return vs[0].(map[string]*ModifyAccountRemarkOperation)[vs[1].(string)]
	}).(ModifyAccountRemarkOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyAccountRemarkOperationInput)(nil)).Elem(), &ModifyAccountRemarkOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyAccountRemarkOperationArrayInput)(nil)).Elem(), ModifyAccountRemarkOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyAccountRemarkOperationMapInput)(nil)).Elem(), ModifyAccountRemarkOperationMap{})
	pulumi.RegisterOutputType(ModifyAccountRemarkOperationOutput{})
	pulumi.RegisterOutputType(ModifyAccountRemarkOperationArrayOutput{})
	pulumi.RegisterOutputType(ModifyAccountRemarkOperationMapOutput{})
}
