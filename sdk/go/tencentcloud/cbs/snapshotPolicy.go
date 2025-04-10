// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a snapshot policy resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cbs.NewSnapshotPolicy(ctx, "snapshotPolicy", &Cbs.SnapshotPolicyArgs{
// 			RepeatHours: pulumi.IntArray{
// 				pulumi.Int(1),
// 			},
// 			RepeatWeekdays: pulumi.IntArray{
// 				pulumi.Int(1),
// 				pulumi.Int(4),
// 			},
// 			RetentionDays:      pulumi.Int(7),
// 			SnapshotPolicyName: pulumi.String("mysnapshotpolicyname"),
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
// CBS snapshot policy can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy snapshot_policy asp-jliex1tn
// ```
type SnapshotPolicy struct {
	pulumi.CustomResourceState

	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours pulumi.IntArrayOutput `pulumi:"repeatHours"`
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays pulumi.IntArrayOutput `pulumi:"repeatWeekdays"`
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName pulumi.StringOutput `pulumi:"snapshotPolicyName"`
}

// NewSnapshotPolicy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotPolicy(ctx *pulumi.Context,
	name string, args *SnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepeatHours == nil {
		return nil, errors.New("invalid value for required argument 'RepeatHours'")
	}
	if args.RepeatWeekdays == nil {
		return nil, errors.New("invalid value for required argument 'RepeatWeekdays'")
	}
	if args.SnapshotPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotPolicyName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SnapshotPolicy
	err := ctx.RegisterResource("tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotPolicy gets an existing SnapshotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotPolicyState, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	var resource SnapshotPolicy
	err := ctx.ReadResource("tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotPolicy resources.
type snapshotPolicyState struct {
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours []int `pulumi:"repeatHours"`
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays []int `pulumi:"repeatWeekdays"`
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays *int `pulumi:"retentionDays"`
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName *string `pulumi:"snapshotPolicyName"`
}

type SnapshotPolicyState struct {
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours pulumi.IntArrayInput
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays pulumi.IntArrayInput
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays pulumi.IntPtrInput
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName pulumi.StringPtrInput
}

func (SnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyState)(nil)).Elem()
}

type snapshotPolicyArgs struct {
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours []int `pulumi:"repeatHours"`
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays []int `pulumi:"repeatWeekdays"`
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays *int `pulumi:"retentionDays"`
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName string `pulumi:"snapshotPolicyName"`
}

// The set of arguments for constructing a SnapshotPolicy resource.
type SnapshotPolicyArgs struct {
	// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
	RepeatHours pulumi.IntArrayInput
	// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
	RepeatWeekdays pulumi.IntArrayInput
	// Retention days of the snapshot, and the default value is 7.
	RetentionDays pulumi.IntPtrInput
	// Name of snapshot policy. The maximum length can not exceed 60 bytes.
	SnapshotPolicyName pulumi.StringInput
}

func (SnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyArgs)(nil)).Elem()
}

type SnapshotPolicyInput interface {
	pulumi.Input

	ToSnapshotPolicyOutput() SnapshotPolicyOutput
	ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput
}

func (*SnapshotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicy)(nil)).Elem()
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return i.ToSnapshotPolicyOutputWithContext(context.Background())
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyOutput)
}

// SnapshotPolicyArrayInput is an input type that accepts SnapshotPolicyArray and SnapshotPolicyArrayOutput values.
// You can construct a concrete instance of `SnapshotPolicyArrayInput` via:
//
//          SnapshotPolicyArray{ SnapshotPolicyArgs{...} }
type SnapshotPolicyArrayInput interface {
	pulumi.Input

	ToSnapshotPolicyArrayOutput() SnapshotPolicyArrayOutput
	ToSnapshotPolicyArrayOutputWithContext(context.Context) SnapshotPolicyArrayOutput
}

type SnapshotPolicyArray []SnapshotPolicyInput

func (SnapshotPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotPolicy)(nil)).Elem()
}

func (i SnapshotPolicyArray) ToSnapshotPolicyArrayOutput() SnapshotPolicyArrayOutput {
	return i.ToSnapshotPolicyArrayOutputWithContext(context.Background())
}

func (i SnapshotPolicyArray) ToSnapshotPolicyArrayOutputWithContext(ctx context.Context) SnapshotPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyArrayOutput)
}

// SnapshotPolicyMapInput is an input type that accepts SnapshotPolicyMap and SnapshotPolicyMapOutput values.
// You can construct a concrete instance of `SnapshotPolicyMapInput` via:
//
//          SnapshotPolicyMap{ "key": SnapshotPolicyArgs{...} }
type SnapshotPolicyMapInput interface {
	pulumi.Input

	ToSnapshotPolicyMapOutput() SnapshotPolicyMapOutput
	ToSnapshotPolicyMapOutputWithContext(context.Context) SnapshotPolicyMapOutput
}

type SnapshotPolicyMap map[string]SnapshotPolicyInput

func (SnapshotPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotPolicy)(nil)).Elem()
}

func (i SnapshotPolicyMap) ToSnapshotPolicyMapOutput() SnapshotPolicyMapOutput {
	return i.ToSnapshotPolicyMapOutputWithContext(context.Background())
}

func (i SnapshotPolicyMap) ToSnapshotPolicyMapOutputWithContext(ctx context.Context) SnapshotPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyMapOutput)
}

type SnapshotPolicyOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return o
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return o
}

// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
func (o SnapshotPolicyOutput) RepeatHours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.IntArrayOutput { return v.RepeatHours }).(pulumi.IntArrayOutput)
}

// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
func (o SnapshotPolicyOutput) RepeatWeekdays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.IntArrayOutput { return v.RepeatWeekdays }).(pulumi.IntArrayOutput)
}

// Retention days of the snapshot, and the default value is 7.
func (o SnapshotPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Name of snapshot policy. The maximum length can not exceed 60 bytes.
func (o SnapshotPolicyOutput) SnapshotPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.SnapshotPolicyName }).(pulumi.StringOutput)
}

type SnapshotPolicyArrayOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyArrayOutput) ToSnapshotPolicyArrayOutput() SnapshotPolicyArrayOutput {
	return o
}

func (o SnapshotPolicyArrayOutput) ToSnapshotPolicyArrayOutputWithContext(ctx context.Context) SnapshotPolicyArrayOutput {
	return o
}

func (o SnapshotPolicyArrayOutput) Index(i pulumi.IntInput) SnapshotPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotPolicy {
		return vs[0].([]*SnapshotPolicy)[vs[1].(int)]
	}).(SnapshotPolicyOutput)
}

type SnapshotPolicyMapOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyMapOutput) ToSnapshotPolicyMapOutput() SnapshotPolicyMapOutput {
	return o
}

func (o SnapshotPolicyMapOutput) ToSnapshotPolicyMapOutputWithContext(ctx context.Context) SnapshotPolicyMapOutput {
	return o
}

func (o SnapshotPolicyMapOutput) MapIndex(k pulumi.StringInput) SnapshotPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotPolicy {
		return vs[0].(map[string]*SnapshotPolicy)[vs[1].(string)]
	}).(SnapshotPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyInput)(nil)).Elem(), &SnapshotPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyArrayInput)(nil)).Elem(), SnapshotPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyMapInput)(nil)).Elem(), SnapshotPolicyMap{})
	pulumi.RegisterOutputType(SnapshotPolicyOutput{})
	pulumi.RegisterOutputType(SnapshotPolicyArrayOutput{})
	pulumi.RegisterOutputType(SnapshotPolicyMapOutput{})
}
