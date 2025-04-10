// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dts syncJobResumeOperation
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
// 		_, err := Dts.NewSyncJobResumeOperation(ctx, "syncJobResumeOperation", &Dts.SyncJobResumeOperationArgs{
// 			JobId: pulumi.String("sync-werwfs23"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SyncJobResumeOperation struct {
	pulumi.CustomResourceState

	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringOutput `pulumi:"jobId"`
}

// NewSyncJobResumeOperation registers a new resource with the given unique name, arguments, and options.
func NewSyncJobResumeOperation(ctx *pulumi.Context,
	name string, args *SyncJobResumeOperationArgs, opts ...pulumi.ResourceOption) (*SyncJobResumeOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SyncJobResumeOperation
	err := ctx.RegisterResource("tencentcloud:Dts/syncJobResumeOperation:SyncJobResumeOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncJobResumeOperation gets an existing SyncJobResumeOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncJobResumeOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncJobResumeOperationState, opts ...pulumi.ResourceOption) (*SyncJobResumeOperation, error) {
	var resource SyncJobResumeOperation
	err := ctx.ReadResource("tencentcloud:Dts/syncJobResumeOperation:SyncJobResumeOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncJobResumeOperation resources.
type syncJobResumeOperationState struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId *string `pulumi:"jobId"`
}

type SyncJobResumeOperationState struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringPtrInput
}

func (SyncJobResumeOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncJobResumeOperationState)(nil)).Elem()
}

type syncJobResumeOperationArgs struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId string `pulumi:"jobId"`
}

// The set of arguments for constructing a SyncJobResumeOperation resource.
type SyncJobResumeOperationArgs struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringInput
}

func (SyncJobResumeOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncJobResumeOperationArgs)(nil)).Elem()
}

type SyncJobResumeOperationInput interface {
	pulumi.Input

	ToSyncJobResumeOperationOutput() SyncJobResumeOperationOutput
	ToSyncJobResumeOperationOutputWithContext(ctx context.Context) SyncJobResumeOperationOutput
}

func (*SyncJobResumeOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncJobResumeOperation)(nil)).Elem()
}

func (i *SyncJobResumeOperation) ToSyncJobResumeOperationOutput() SyncJobResumeOperationOutput {
	return i.ToSyncJobResumeOperationOutputWithContext(context.Background())
}

func (i *SyncJobResumeOperation) ToSyncJobResumeOperationOutputWithContext(ctx context.Context) SyncJobResumeOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobResumeOperationOutput)
}

// SyncJobResumeOperationArrayInput is an input type that accepts SyncJobResumeOperationArray and SyncJobResumeOperationArrayOutput values.
// You can construct a concrete instance of `SyncJobResumeOperationArrayInput` via:
//
//          SyncJobResumeOperationArray{ SyncJobResumeOperationArgs{...} }
type SyncJobResumeOperationArrayInput interface {
	pulumi.Input

	ToSyncJobResumeOperationArrayOutput() SyncJobResumeOperationArrayOutput
	ToSyncJobResumeOperationArrayOutputWithContext(context.Context) SyncJobResumeOperationArrayOutput
}

type SyncJobResumeOperationArray []SyncJobResumeOperationInput

func (SyncJobResumeOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncJobResumeOperation)(nil)).Elem()
}

func (i SyncJobResumeOperationArray) ToSyncJobResumeOperationArrayOutput() SyncJobResumeOperationArrayOutput {
	return i.ToSyncJobResumeOperationArrayOutputWithContext(context.Background())
}

func (i SyncJobResumeOperationArray) ToSyncJobResumeOperationArrayOutputWithContext(ctx context.Context) SyncJobResumeOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobResumeOperationArrayOutput)
}

// SyncJobResumeOperationMapInput is an input type that accepts SyncJobResumeOperationMap and SyncJobResumeOperationMapOutput values.
// You can construct a concrete instance of `SyncJobResumeOperationMapInput` via:
//
//          SyncJobResumeOperationMap{ "key": SyncJobResumeOperationArgs{...} }
type SyncJobResumeOperationMapInput interface {
	pulumi.Input

	ToSyncJobResumeOperationMapOutput() SyncJobResumeOperationMapOutput
	ToSyncJobResumeOperationMapOutputWithContext(context.Context) SyncJobResumeOperationMapOutput
}

type SyncJobResumeOperationMap map[string]SyncJobResumeOperationInput

func (SyncJobResumeOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncJobResumeOperation)(nil)).Elem()
}

func (i SyncJobResumeOperationMap) ToSyncJobResumeOperationMapOutput() SyncJobResumeOperationMapOutput {
	return i.ToSyncJobResumeOperationMapOutputWithContext(context.Background())
}

func (i SyncJobResumeOperationMap) ToSyncJobResumeOperationMapOutputWithContext(ctx context.Context) SyncJobResumeOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobResumeOperationMapOutput)
}

type SyncJobResumeOperationOutput struct{ *pulumi.OutputState }

func (SyncJobResumeOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncJobResumeOperation)(nil)).Elem()
}

func (o SyncJobResumeOperationOutput) ToSyncJobResumeOperationOutput() SyncJobResumeOperationOutput {
	return o
}

func (o SyncJobResumeOperationOutput) ToSyncJobResumeOperationOutputWithContext(ctx context.Context) SyncJobResumeOperationOutput {
	return o
}

// Synchronization instance id (i.e. identifies a synchronization job).
func (o SyncJobResumeOperationOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncJobResumeOperation) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

type SyncJobResumeOperationArrayOutput struct{ *pulumi.OutputState }

func (SyncJobResumeOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncJobResumeOperation)(nil)).Elem()
}

func (o SyncJobResumeOperationArrayOutput) ToSyncJobResumeOperationArrayOutput() SyncJobResumeOperationArrayOutput {
	return o
}

func (o SyncJobResumeOperationArrayOutput) ToSyncJobResumeOperationArrayOutputWithContext(ctx context.Context) SyncJobResumeOperationArrayOutput {
	return o
}

func (o SyncJobResumeOperationArrayOutput) Index(i pulumi.IntInput) SyncJobResumeOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncJobResumeOperation {
		return vs[0].([]*SyncJobResumeOperation)[vs[1].(int)]
	}).(SyncJobResumeOperationOutput)
}

type SyncJobResumeOperationMapOutput struct{ *pulumi.OutputState }

func (SyncJobResumeOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncJobResumeOperation)(nil)).Elem()
}

func (o SyncJobResumeOperationMapOutput) ToSyncJobResumeOperationMapOutput() SyncJobResumeOperationMapOutput {
	return o
}

func (o SyncJobResumeOperationMapOutput) ToSyncJobResumeOperationMapOutputWithContext(ctx context.Context) SyncJobResumeOperationMapOutput {
	return o
}

func (o SyncJobResumeOperationMapOutput) MapIndex(k pulumi.StringInput) SyncJobResumeOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncJobResumeOperation {
		return vs[0].(map[string]*SyncJobResumeOperation)[vs[1].(string)]
	}).(SyncJobResumeOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobResumeOperationInput)(nil)).Elem(), &SyncJobResumeOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobResumeOperationArrayInput)(nil)).Elem(), SyncJobResumeOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobResumeOperationMapInput)(nil)).Elem(), SyncJobResumeOperationMap{})
	pulumi.RegisterOutputType(SyncJobResumeOperationOutput{})
	pulumi.RegisterOutputType(SyncJobResumeOperationArrayOutput{})
	pulumi.RegisterOutputType(SyncJobResumeOperationMapOutput{})
}
