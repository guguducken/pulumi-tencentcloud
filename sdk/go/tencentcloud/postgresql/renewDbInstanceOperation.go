// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a postgresql renewDbInstanceOperation
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
// 		_, err := Postgresql.NewRenewDbInstanceOperation(ctx, "renewDbInstanceOperation", &Postgresql.RenewDbInstanceOperationArgs{
// 			DbInstanceId: pulumi.Any(tencentcloud_postgresql_instance.Oper_test_PREPAID.Id),
// 			Period:       pulumi.Int(1),
// 			AutoVoucher:  pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RenewDbInstanceOperation struct {
	pulumi.CustomResourceState

	// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
	AutoVoucher pulumi.IntPtrOutput `pulumi:"autoVoucher"`
	// Instance ID in the format of postgres-6fego161.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// Renewal duration in months.
	Period pulumi.IntOutput `pulumi:"period"`
	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds pulumi.StringArrayOutput `pulumi:"voucherIds"`
}

// NewRenewDbInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewRenewDbInstanceOperation(ctx *pulumi.Context,
	name string, args *RenewDbInstanceOperationArgs, opts ...pulumi.ResourceOption) (*RenewDbInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RenewDbInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRenewDbInstanceOperation gets an existing RenewDbInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRenewDbInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RenewDbInstanceOperationState, opts ...pulumi.ResourceOption) (*RenewDbInstanceOperation, error) {
	var resource RenewDbInstanceOperation
	err := ctx.ReadResource("tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RenewDbInstanceOperation resources.
type renewDbInstanceOperationState struct {
	// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
	AutoVoucher *int `pulumi:"autoVoucher"`
	// Instance ID in the format of postgres-6fego161.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// Renewal duration in months.
	Period *int `pulumi:"period"`
	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds []string `pulumi:"voucherIds"`
}

type RenewDbInstanceOperationState struct {
	// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
	AutoVoucher pulumi.IntPtrInput
	// Instance ID in the format of postgres-6fego161.
	DbInstanceId pulumi.StringPtrInput
	// Renewal duration in months.
	Period pulumi.IntPtrInput
	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds pulumi.StringArrayInput
}

func (RenewDbInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*renewDbInstanceOperationState)(nil)).Elem()
}

type renewDbInstanceOperationArgs struct {
	// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
	AutoVoucher *int `pulumi:"autoVoucher"`
	// Instance ID in the format of postgres-6fego161.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// Renewal duration in months.
	Period int `pulumi:"period"`
	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds []string `pulumi:"voucherIds"`
}

// The set of arguments for constructing a RenewDbInstanceOperation resource.
type RenewDbInstanceOperationArgs struct {
	// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
	AutoVoucher pulumi.IntPtrInput
	// Instance ID in the format of postgres-6fego161.
	DbInstanceId pulumi.StringInput
	// Renewal duration in months.
	Period pulumi.IntInput
	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds pulumi.StringArrayInput
}

func (RenewDbInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*renewDbInstanceOperationArgs)(nil)).Elem()
}

type RenewDbInstanceOperationInput interface {
	pulumi.Input

	ToRenewDbInstanceOperationOutput() RenewDbInstanceOperationOutput
	ToRenewDbInstanceOperationOutputWithContext(ctx context.Context) RenewDbInstanceOperationOutput
}

func (*RenewDbInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewDbInstanceOperation)(nil)).Elem()
}

func (i *RenewDbInstanceOperation) ToRenewDbInstanceOperationOutput() RenewDbInstanceOperationOutput {
	return i.ToRenewDbInstanceOperationOutputWithContext(context.Background())
}

func (i *RenewDbInstanceOperation) ToRenewDbInstanceOperationOutputWithContext(ctx context.Context) RenewDbInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewDbInstanceOperationOutput)
}

// RenewDbInstanceOperationArrayInput is an input type that accepts RenewDbInstanceOperationArray and RenewDbInstanceOperationArrayOutput values.
// You can construct a concrete instance of `RenewDbInstanceOperationArrayInput` via:
//
//          RenewDbInstanceOperationArray{ RenewDbInstanceOperationArgs{...} }
type RenewDbInstanceOperationArrayInput interface {
	pulumi.Input

	ToRenewDbInstanceOperationArrayOutput() RenewDbInstanceOperationArrayOutput
	ToRenewDbInstanceOperationArrayOutputWithContext(context.Context) RenewDbInstanceOperationArrayOutput
}

type RenewDbInstanceOperationArray []RenewDbInstanceOperationInput

func (RenewDbInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewDbInstanceOperation)(nil)).Elem()
}

func (i RenewDbInstanceOperationArray) ToRenewDbInstanceOperationArrayOutput() RenewDbInstanceOperationArrayOutput {
	return i.ToRenewDbInstanceOperationArrayOutputWithContext(context.Background())
}

func (i RenewDbInstanceOperationArray) ToRenewDbInstanceOperationArrayOutputWithContext(ctx context.Context) RenewDbInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewDbInstanceOperationArrayOutput)
}

// RenewDbInstanceOperationMapInput is an input type that accepts RenewDbInstanceOperationMap and RenewDbInstanceOperationMapOutput values.
// You can construct a concrete instance of `RenewDbInstanceOperationMapInput` via:
//
//          RenewDbInstanceOperationMap{ "key": RenewDbInstanceOperationArgs{...} }
type RenewDbInstanceOperationMapInput interface {
	pulumi.Input

	ToRenewDbInstanceOperationMapOutput() RenewDbInstanceOperationMapOutput
	ToRenewDbInstanceOperationMapOutputWithContext(context.Context) RenewDbInstanceOperationMapOutput
}

type RenewDbInstanceOperationMap map[string]RenewDbInstanceOperationInput

func (RenewDbInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewDbInstanceOperation)(nil)).Elem()
}

func (i RenewDbInstanceOperationMap) ToRenewDbInstanceOperationMapOutput() RenewDbInstanceOperationMapOutput {
	return i.ToRenewDbInstanceOperationMapOutputWithContext(context.Background())
}

func (i RenewDbInstanceOperationMap) ToRenewDbInstanceOperationMapOutputWithContext(ctx context.Context) RenewDbInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewDbInstanceOperationMapOutput)
}

type RenewDbInstanceOperationOutput struct{ *pulumi.OutputState }

func (RenewDbInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewDbInstanceOperation)(nil)).Elem()
}

func (o RenewDbInstanceOperationOutput) ToRenewDbInstanceOperationOutput() RenewDbInstanceOperationOutput {
	return o
}

func (o RenewDbInstanceOperationOutput) ToRenewDbInstanceOperationOutputWithContext(ctx context.Context) RenewDbInstanceOperationOutput {
	return o
}

// Whether to automatically use vouchers. 1:yes, 0:no. Default value:0.
func (o RenewDbInstanceOperationOutput) AutoVoucher() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RenewDbInstanceOperation) pulumi.IntPtrOutput { return v.AutoVoucher }).(pulumi.IntPtrOutput)
}

// Instance ID in the format of postgres-6fego161.
func (o RenewDbInstanceOperationOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RenewDbInstanceOperation) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// Renewal duration in months.
func (o RenewDbInstanceOperationOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *RenewDbInstanceOperation) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// Voucher ID list (only one voucher can be specified currently).
func (o RenewDbInstanceOperationOutput) VoucherIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RenewDbInstanceOperation) pulumi.StringArrayOutput { return v.VoucherIds }).(pulumi.StringArrayOutput)
}

type RenewDbInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (RenewDbInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewDbInstanceOperation)(nil)).Elem()
}

func (o RenewDbInstanceOperationArrayOutput) ToRenewDbInstanceOperationArrayOutput() RenewDbInstanceOperationArrayOutput {
	return o
}

func (o RenewDbInstanceOperationArrayOutput) ToRenewDbInstanceOperationArrayOutputWithContext(ctx context.Context) RenewDbInstanceOperationArrayOutput {
	return o
}

func (o RenewDbInstanceOperationArrayOutput) Index(i pulumi.IntInput) RenewDbInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RenewDbInstanceOperation {
		return vs[0].([]*RenewDbInstanceOperation)[vs[1].(int)]
	}).(RenewDbInstanceOperationOutput)
}

type RenewDbInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (RenewDbInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewDbInstanceOperation)(nil)).Elem()
}

func (o RenewDbInstanceOperationMapOutput) ToRenewDbInstanceOperationMapOutput() RenewDbInstanceOperationMapOutput {
	return o
}

func (o RenewDbInstanceOperationMapOutput) ToRenewDbInstanceOperationMapOutputWithContext(ctx context.Context) RenewDbInstanceOperationMapOutput {
	return o
}

func (o RenewDbInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) RenewDbInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RenewDbInstanceOperation {
		return vs[0].(map[string]*RenewDbInstanceOperation)[vs[1].(string)]
	}).(RenewDbInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RenewDbInstanceOperationInput)(nil)).Elem(), &RenewDbInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewDbInstanceOperationArrayInput)(nil)).Elem(), RenewDbInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewDbInstanceOperationMapInput)(nil)).Elem(), RenewDbInstanceOperationMap{})
	pulumi.RegisterOutputType(RenewDbInstanceOperationOutput{})
	pulumi.RegisterOutputType(RenewDbInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(RenewDbInstanceOperationMapOutput{})
}
