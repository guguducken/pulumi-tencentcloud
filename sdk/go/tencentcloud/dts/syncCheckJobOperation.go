// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dts syncCheckJobOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewSyncCheckJobOperation(ctx, "syncCheckJobOperation", &Dts.SyncCheckJobOperationArgs{
// 			JobId: pulumi.String(""),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SyncCheckJobOperation struct {
	pulumi.CustomResourceState

	// Sync job id.
	JobId pulumi.StringOutput `pulumi:"jobId"`
}

// NewSyncCheckJobOperation registers a new resource with the given unique name, arguments, and options.
func NewSyncCheckJobOperation(ctx *pulumi.Context,
	name string, args *SyncCheckJobOperationArgs, opts ...pulumi.ResourceOption) (*SyncCheckJobOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SyncCheckJobOperation
	err := ctx.RegisterResource("tencentcloud:Dts/syncCheckJobOperation:SyncCheckJobOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncCheckJobOperation gets an existing SyncCheckJobOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncCheckJobOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncCheckJobOperationState, opts ...pulumi.ResourceOption) (*SyncCheckJobOperation, error) {
	var resource SyncCheckJobOperation
	err := ctx.ReadResource("tencentcloud:Dts/syncCheckJobOperation:SyncCheckJobOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncCheckJobOperation resources.
type syncCheckJobOperationState struct {
	// Sync job id.
	JobId *string `pulumi:"jobId"`
}

type SyncCheckJobOperationState struct {
	// Sync job id.
	JobId pulumi.StringPtrInput
}

func (SyncCheckJobOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncCheckJobOperationState)(nil)).Elem()
}

type syncCheckJobOperationArgs struct {
	// Sync job id.
	JobId string `pulumi:"jobId"`
}

// The set of arguments for constructing a SyncCheckJobOperation resource.
type SyncCheckJobOperationArgs struct {
	// Sync job id.
	JobId pulumi.StringInput
}

func (SyncCheckJobOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncCheckJobOperationArgs)(nil)).Elem()
}

type SyncCheckJobOperationInput interface {
	pulumi.Input

	ToSyncCheckJobOperationOutput() SyncCheckJobOperationOutput
	ToSyncCheckJobOperationOutputWithContext(ctx context.Context) SyncCheckJobOperationOutput
}

func (*SyncCheckJobOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncCheckJobOperation)(nil)).Elem()
}

func (i *SyncCheckJobOperation) ToSyncCheckJobOperationOutput() SyncCheckJobOperationOutput {
	return i.ToSyncCheckJobOperationOutputWithContext(context.Background())
}

func (i *SyncCheckJobOperation) ToSyncCheckJobOperationOutputWithContext(ctx context.Context) SyncCheckJobOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCheckJobOperationOutput)
}

// SyncCheckJobOperationArrayInput is an input type that accepts SyncCheckJobOperationArray and SyncCheckJobOperationArrayOutput values.
// You can construct a concrete instance of `SyncCheckJobOperationArrayInput` via:
//
//          SyncCheckJobOperationArray{ SyncCheckJobOperationArgs{...} }
type SyncCheckJobOperationArrayInput interface {
	pulumi.Input

	ToSyncCheckJobOperationArrayOutput() SyncCheckJobOperationArrayOutput
	ToSyncCheckJobOperationArrayOutputWithContext(context.Context) SyncCheckJobOperationArrayOutput
}

type SyncCheckJobOperationArray []SyncCheckJobOperationInput

func (SyncCheckJobOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncCheckJobOperation)(nil)).Elem()
}

func (i SyncCheckJobOperationArray) ToSyncCheckJobOperationArrayOutput() SyncCheckJobOperationArrayOutput {
	return i.ToSyncCheckJobOperationArrayOutputWithContext(context.Background())
}

func (i SyncCheckJobOperationArray) ToSyncCheckJobOperationArrayOutputWithContext(ctx context.Context) SyncCheckJobOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCheckJobOperationArrayOutput)
}

// SyncCheckJobOperationMapInput is an input type that accepts SyncCheckJobOperationMap and SyncCheckJobOperationMapOutput values.
// You can construct a concrete instance of `SyncCheckJobOperationMapInput` via:
//
//          SyncCheckJobOperationMap{ "key": SyncCheckJobOperationArgs{...} }
type SyncCheckJobOperationMapInput interface {
	pulumi.Input

	ToSyncCheckJobOperationMapOutput() SyncCheckJobOperationMapOutput
	ToSyncCheckJobOperationMapOutputWithContext(context.Context) SyncCheckJobOperationMapOutput
}

type SyncCheckJobOperationMap map[string]SyncCheckJobOperationInput

func (SyncCheckJobOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncCheckJobOperation)(nil)).Elem()
}

func (i SyncCheckJobOperationMap) ToSyncCheckJobOperationMapOutput() SyncCheckJobOperationMapOutput {
	return i.ToSyncCheckJobOperationMapOutputWithContext(context.Background())
}

func (i SyncCheckJobOperationMap) ToSyncCheckJobOperationMapOutputWithContext(ctx context.Context) SyncCheckJobOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCheckJobOperationMapOutput)
}

type SyncCheckJobOperationOutput struct{ *pulumi.OutputState }

func (SyncCheckJobOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncCheckJobOperation)(nil)).Elem()
}

func (o SyncCheckJobOperationOutput) ToSyncCheckJobOperationOutput() SyncCheckJobOperationOutput {
	return o
}

func (o SyncCheckJobOperationOutput) ToSyncCheckJobOperationOutputWithContext(ctx context.Context) SyncCheckJobOperationOutput {
	return o
}

// Sync job id.
func (o SyncCheckJobOperationOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncCheckJobOperation) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

type SyncCheckJobOperationArrayOutput struct{ *pulumi.OutputState }

func (SyncCheckJobOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncCheckJobOperation)(nil)).Elem()
}

func (o SyncCheckJobOperationArrayOutput) ToSyncCheckJobOperationArrayOutput() SyncCheckJobOperationArrayOutput {
	return o
}

func (o SyncCheckJobOperationArrayOutput) ToSyncCheckJobOperationArrayOutputWithContext(ctx context.Context) SyncCheckJobOperationArrayOutput {
	return o
}

func (o SyncCheckJobOperationArrayOutput) Index(i pulumi.IntInput) SyncCheckJobOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncCheckJobOperation {
		return vs[0].([]*SyncCheckJobOperation)[vs[1].(int)]
	}).(SyncCheckJobOperationOutput)
}

type SyncCheckJobOperationMapOutput struct{ *pulumi.OutputState }

func (SyncCheckJobOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncCheckJobOperation)(nil)).Elem()
}

func (o SyncCheckJobOperationMapOutput) ToSyncCheckJobOperationMapOutput() SyncCheckJobOperationMapOutput {
	return o
}

func (o SyncCheckJobOperationMapOutput) ToSyncCheckJobOperationMapOutputWithContext(ctx context.Context) SyncCheckJobOperationMapOutput {
	return o
}

func (o SyncCheckJobOperationMapOutput) MapIndex(k pulumi.StringInput) SyncCheckJobOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncCheckJobOperation {
		return vs[0].(map[string]*SyncCheckJobOperation)[vs[1].(string)]
	}).(SyncCheckJobOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCheckJobOperationInput)(nil)).Elem(), &SyncCheckJobOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCheckJobOperationArrayInput)(nil)).Elem(), SyncCheckJobOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCheckJobOperationMapInput)(nil)).Elem(), SyncCheckJobOperationMap{})
	pulumi.RegisterOutputType(SyncCheckJobOperationOutput{})
	pulumi.RegisterOutputType(SyncCheckJobOperationArrayOutput{})
	pulumi.RegisterOutputType(SyncCheckJobOperationMapOutput{})
}
