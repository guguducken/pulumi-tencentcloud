// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TaskSetBatchTasks struct {
	// Task name.
	Name string `pulumi:"name"`
	// Target address.
	TargetAddress string `pulumi:"targetAddress"`
}

// TaskSetBatchTasksInput is an input type that accepts TaskSetBatchTasksArgs and TaskSetBatchTasksOutput values.
// You can construct a concrete instance of `TaskSetBatchTasksInput` via:
//
//          TaskSetBatchTasksArgs{...}
type TaskSetBatchTasksInput interface {
	pulumi.Input

	ToTaskSetBatchTasksOutput() TaskSetBatchTasksOutput
	ToTaskSetBatchTasksOutputWithContext(context.Context) TaskSetBatchTasksOutput
}

type TaskSetBatchTasksArgs struct {
	// Task name.
	Name pulumi.StringInput `pulumi:"name"`
	// Target address.
	TargetAddress pulumi.StringInput `pulumi:"targetAddress"`
}

func (TaskSetBatchTasksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskSetBatchTasks)(nil)).Elem()
}

func (i TaskSetBatchTasksArgs) ToTaskSetBatchTasksOutput() TaskSetBatchTasksOutput {
	return i.ToTaskSetBatchTasksOutputWithContext(context.Background())
}

func (i TaskSetBatchTasksArgs) ToTaskSetBatchTasksOutputWithContext(ctx context.Context) TaskSetBatchTasksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetBatchTasksOutput)
}

func (i TaskSetBatchTasksArgs) ToTaskSetBatchTasksPtrOutput() TaskSetBatchTasksPtrOutput {
	return i.ToTaskSetBatchTasksPtrOutputWithContext(context.Background())
}

func (i TaskSetBatchTasksArgs) ToTaskSetBatchTasksPtrOutputWithContext(ctx context.Context) TaskSetBatchTasksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetBatchTasksOutput).ToTaskSetBatchTasksPtrOutputWithContext(ctx)
}

// TaskSetBatchTasksPtrInput is an input type that accepts TaskSetBatchTasksArgs, TaskSetBatchTasksPtr and TaskSetBatchTasksPtrOutput values.
// You can construct a concrete instance of `TaskSetBatchTasksPtrInput` via:
//
//          TaskSetBatchTasksArgs{...}
//
//  or:
//
//          nil
type TaskSetBatchTasksPtrInput interface {
	pulumi.Input

	ToTaskSetBatchTasksPtrOutput() TaskSetBatchTasksPtrOutput
	ToTaskSetBatchTasksPtrOutputWithContext(context.Context) TaskSetBatchTasksPtrOutput
}

type taskSetBatchTasksPtrType TaskSetBatchTasksArgs

func TaskSetBatchTasksPtr(v *TaskSetBatchTasksArgs) TaskSetBatchTasksPtrInput {
	return (*taskSetBatchTasksPtrType)(v)
}

func (*taskSetBatchTasksPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSetBatchTasks)(nil)).Elem()
}

func (i *taskSetBatchTasksPtrType) ToTaskSetBatchTasksPtrOutput() TaskSetBatchTasksPtrOutput {
	return i.ToTaskSetBatchTasksPtrOutputWithContext(context.Background())
}

func (i *taskSetBatchTasksPtrType) ToTaskSetBatchTasksPtrOutputWithContext(ctx context.Context) TaskSetBatchTasksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetBatchTasksPtrOutput)
}

type TaskSetBatchTasksOutput struct{ *pulumi.OutputState }

func (TaskSetBatchTasksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskSetBatchTasks)(nil)).Elem()
}

func (o TaskSetBatchTasksOutput) ToTaskSetBatchTasksOutput() TaskSetBatchTasksOutput {
	return o
}

func (o TaskSetBatchTasksOutput) ToTaskSetBatchTasksOutputWithContext(ctx context.Context) TaskSetBatchTasksOutput {
	return o
}

func (o TaskSetBatchTasksOutput) ToTaskSetBatchTasksPtrOutput() TaskSetBatchTasksPtrOutput {
	return o.ToTaskSetBatchTasksPtrOutputWithContext(context.Background())
}

func (o TaskSetBatchTasksOutput) ToTaskSetBatchTasksPtrOutputWithContext(ctx context.Context) TaskSetBatchTasksPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskSetBatchTasks) *TaskSetBatchTasks {
		return &v
	}).(TaskSetBatchTasksPtrOutput)
}

// Task name.
func (o TaskSetBatchTasksOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TaskSetBatchTasks) string { return v.Name }).(pulumi.StringOutput)
}

// Target address.
func (o TaskSetBatchTasksOutput) TargetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v TaskSetBatchTasks) string { return v.TargetAddress }).(pulumi.StringOutput)
}

type TaskSetBatchTasksPtrOutput struct{ *pulumi.OutputState }

func (TaskSetBatchTasksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSetBatchTasks)(nil)).Elem()
}

func (o TaskSetBatchTasksPtrOutput) ToTaskSetBatchTasksPtrOutput() TaskSetBatchTasksPtrOutput {
	return o
}

func (o TaskSetBatchTasksPtrOutput) ToTaskSetBatchTasksPtrOutputWithContext(ctx context.Context) TaskSetBatchTasksPtrOutput {
	return o
}

func (o TaskSetBatchTasksPtrOutput) Elem() TaskSetBatchTasksOutput {
	return o.ApplyT(func(v *TaskSetBatchTasks) TaskSetBatchTasks {
		if v != nil {
			return *v
		}
		var ret TaskSetBatchTasks
		return ret
	}).(TaskSetBatchTasksOutput)
}

// Task name.
func (o TaskSetBatchTasksPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskSetBatchTasks) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Target address.
func (o TaskSetBatchTasksPtrOutput) TargetAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskSetBatchTasks) *string {
		if v == nil {
			return nil
		}
		return &v.TargetAddress
	}).(pulumi.StringPtrOutput)
}

type GetNodeNodeDefine struct {
	// City.
	City string `pulumi:"city"`
	// Node ID.
	Code string `pulumi:"code"`
	// If the node type is base, it is an availability dial test point; if it is blank, it is an advanced dial test point.
	CodeType string `pulumi:"codeType"`
	// District.
	District string `pulumi:"district"`
	// IP type:1 = IPv4,2 = IPv6.
	IpType int `pulumi:"ipType"`
	// Node area:1=Chinese Mainland,2=Hong Kong, Macao and Taiwan,3=Overseas.
	Location int `pulumi:"location"`
	// Node name.
	Name string `pulumi:"name"`
	// Network service provider.
	NetService string `pulumi:"netService"`
	// Node status: 1=running, 2=offline.
	NodeDefineStatus int `pulumi:"nodeDefineStatus"`
	// Node Type;1 = IDC,2 = LastMile,3 = Mobile.
	Type int `pulumi:"type"`
}

// GetNodeNodeDefineInput is an input type that accepts GetNodeNodeDefineArgs and GetNodeNodeDefineOutput values.
// You can construct a concrete instance of `GetNodeNodeDefineInput` via:
//
//          GetNodeNodeDefineArgs{...}
type GetNodeNodeDefineInput interface {
	pulumi.Input

	ToGetNodeNodeDefineOutput() GetNodeNodeDefineOutput
	ToGetNodeNodeDefineOutputWithContext(context.Context) GetNodeNodeDefineOutput
}

type GetNodeNodeDefineArgs struct {
	// City.
	City pulumi.StringInput `pulumi:"city"`
	// Node ID.
	Code pulumi.StringInput `pulumi:"code"`
	// If the node type is base, it is an availability dial test point; if it is blank, it is an advanced dial test point.
	CodeType pulumi.StringInput `pulumi:"codeType"`
	// District.
	District pulumi.StringInput `pulumi:"district"`
	// IP type:1 = IPv4,2 = IPv6.
	IpType pulumi.IntInput `pulumi:"ipType"`
	// Node area:1=Chinese Mainland,2=Hong Kong, Macao and Taiwan,3=Overseas.
	Location pulumi.IntInput `pulumi:"location"`
	// Node name.
	Name pulumi.StringInput `pulumi:"name"`
	// Network service provider.
	NetService pulumi.StringInput `pulumi:"netService"`
	// Node status: 1=running, 2=offline.
	NodeDefineStatus pulumi.IntInput `pulumi:"nodeDefineStatus"`
	// Node Type;1 = IDC,2 = LastMile,3 = Mobile.
	Type pulumi.IntInput `pulumi:"type"`
}

func (GetNodeNodeDefineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodeNodeDefine)(nil)).Elem()
}

func (i GetNodeNodeDefineArgs) ToGetNodeNodeDefineOutput() GetNodeNodeDefineOutput {
	return i.ToGetNodeNodeDefineOutputWithContext(context.Background())
}

func (i GetNodeNodeDefineArgs) ToGetNodeNodeDefineOutputWithContext(ctx context.Context) GetNodeNodeDefineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNodeNodeDefineOutput)
}

// GetNodeNodeDefineArrayInput is an input type that accepts GetNodeNodeDefineArray and GetNodeNodeDefineArrayOutput values.
// You can construct a concrete instance of `GetNodeNodeDefineArrayInput` via:
//
//          GetNodeNodeDefineArray{ GetNodeNodeDefineArgs{...} }
type GetNodeNodeDefineArrayInput interface {
	pulumi.Input

	ToGetNodeNodeDefineArrayOutput() GetNodeNodeDefineArrayOutput
	ToGetNodeNodeDefineArrayOutputWithContext(context.Context) GetNodeNodeDefineArrayOutput
}

type GetNodeNodeDefineArray []GetNodeNodeDefineInput

func (GetNodeNodeDefineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNodeNodeDefine)(nil)).Elem()
}

func (i GetNodeNodeDefineArray) ToGetNodeNodeDefineArrayOutput() GetNodeNodeDefineArrayOutput {
	return i.ToGetNodeNodeDefineArrayOutputWithContext(context.Background())
}

func (i GetNodeNodeDefineArray) ToGetNodeNodeDefineArrayOutputWithContext(ctx context.Context) GetNodeNodeDefineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNodeNodeDefineArrayOutput)
}

type GetNodeNodeDefineOutput struct{ *pulumi.OutputState }

func (GetNodeNodeDefineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodeNodeDefine)(nil)).Elem()
}

func (o GetNodeNodeDefineOutput) ToGetNodeNodeDefineOutput() GetNodeNodeDefineOutput {
	return o
}

func (o GetNodeNodeDefineOutput) ToGetNodeNodeDefineOutputWithContext(ctx context.Context) GetNodeNodeDefineOutput {
	return o
}

// City.
func (o GetNodeNodeDefineOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) string { return v.City }).(pulumi.StringOutput)
}

// Node ID.
func (o GetNodeNodeDefineOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) string { return v.Code }).(pulumi.StringOutput)
}

// If the node type is base, it is an availability dial test point; if it is blank, it is an advanced dial test point.
func (o GetNodeNodeDefineOutput) CodeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) string { return v.CodeType }).(pulumi.StringOutput)
}

// District.
func (o GetNodeNodeDefineOutput) District() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) string { return v.District }).(pulumi.StringOutput)
}

// IP type:1 = IPv4,2 = IPv6.
func (o GetNodeNodeDefineOutput) IpType() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) int { return v.IpType }).(pulumi.IntOutput)
}

// Node area:1=Chinese Mainland,2=Hong Kong, Macao and Taiwan,3=Overseas.
func (o GetNodeNodeDefineOutput) Location() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) int { return v.Location }).(pulumi.IntOutput)
}

// Node name.
func (o GetNodeNodeDefineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) string { return v.Name }).(pulumi.StringOutput)
}

// Network service provider.
func (o GetNodeNodeDefineOutput) NetService() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) string { return v.NetService }).(pulumi.StringOutput)
}

// Node status: 1=running, 2=offline.
func (o GetNodeNodeDefineOutput) NodeDefineStatus() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) int { return v.NodeDefineStatus }).(pulumi.IntOutput)
}

// Node Type;1 = IDC,2 = LastMile,3 = Mobile.
func (o GetNodeNodeDefineOutput) Type() pulumi.IntOutput {
	return o.ApplyT(func(v GetNodeNodeDefine) int { return v.Type }).(pulumi.IntOutput)
}

type GetNodeNodeDefineArrayOutput struct{ *pulumi.OutputState }

func (GetNodeNodeDefineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNodeNodeDefine)(nil)).Elem()
}

func (o GetNodeNodeDefineArrayOutput) ToGetNodeNodeDefineArrayOutput() GetNodeNodeDefineArrayOutput {
	return o
}

func (o GetNodeNodeDefineArrayOutput) ToGetNodeNodeDefineArrayOutputWithContext(ctx context.Context) GetNodeNodeDefineArrayOutput {
	return o
}

func (o GetNodeNodeDefineArrayOutput) Index(i pulumi.IntInput) GetNodeNodeDefineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNodeNodeDefine {
		return vs[0].([]GetNodeNodeDefine)[vs[1].(int)]
	}).(GetNodeNodeDefineOutput)
}

type GetProbeDataDetailedSingleDataDefine struct {
	// Fields.
	Fields []GetProbeDataDetailedSingleDataDefineField `pulumi:"fields"`
	// Labels.
	Labels []GetProbeDataDetailedSingleDataDefineLabel `pulumi:"labels"`
	// Probe time.
	ProbeTime int `pulumi:"probeTime"`
}

// GetProbeDataDetailedSingleDataDefineInput is an input type that accepts GetProbeDataDetailedSingleDataDefineArgs and GetProbeDataDetailedSingleDataDefineOutput values.
// You can construct a concrete instance of `GetProbeDataDetailedSingleDataDefineInput` via:
//
//          GetProbeDataDetailedSingleDataDefineArgs{...}
type GetProbeDataDetailedSingleDataDefineInput interface {
	pulumi.Input

	ToGetProbeDataDetailedSingleDataDefineOutput() GetProbeDataDetailedSingleDataDefineOutput
	ToGetProbeDataDetailedSingleDataDefineOutputWithContext(context.Context) GetProbeDataDetailedSingleDataDefineOutput
}

type GetProbeDataDetailedSingleDataDefineArgs struct {
	// Fields.
	Fields GetProbeDataDetailedSingleDataDefineFieldArrayInput `pulumi:"fields"`
	// Labels.
	Labels GetProbeDataDetailedSingleDataDefineLabelArrayInput `pulumi:"labels"`
	// Probe time.
	ProbeTime pulumi.IntInput `pulumi:"probeTime"`
}

func (GetProbeDataDetailedSingleDataDefineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataDetailedSingleDataDefine)(nil)).Elem()
}

func (i GetProbeDataDetailedSingleDataDefineArgs) ToGetProbeDataDetailedSingleDataDefineOutput() GetProbeDataDetailedSingleDataDefineOutput {
	return i.ToGetProbeDataDetailedSingleDataDefineOutputWithContext(context.Background())
}

func (i GetProbeDataDetailedSingleDataDefineArgs) ToGetProbeDataDetailedSingleDataDefineOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProbeDataDetailedSingleDataDefineOutput)
}

// GetProbeDataDetailedSingleDataDefineArrayInput is an input type that accepts GetProbeDataDetailedSingleDataDefineArray and GetProbeDataDetailedSingleDataDefineArrayOutput values.
// You can construct a concrete instance of `GetProbeDataDetailedSingleDataDefineArrayInput` via:
//
//          GetProbeDataDetailedSingleDataDefineArray{ GetProbeDataDetailedSingleDataDefineArgs{...} }
type GetProbeDataDetailedSingleDataDefineArrayInput interface {
	pulumi.Input

	ToGetProbeDataDetailedSingleDataDefineArrayOutput() GetProbeDataDetailedSingleDataDefineArrayOutput
	ToGetProbeDataDetailedSingleDataDefineArrayOutputWithContext(context.Context) GetProbeDataDetailedSingleDataDefineArrayOutput
}

type GetProbeDataDetailedSingleDataDefineArray []GetProbeDataDetailedSingleDataDefineInput

func (GetProbeDataDetailedSingleDataDefineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProbeDataDetailedSingleDataDefine)(nil)).Elem()
}

func (i GetProbeDataDetailedSingleDataDefineArray) ToGetProbeDataDetailedSingleDataDefineArrayOutput() GetProbeDataDetailedSingleDataDefineArrayOutput {
	return i.ToGetProbeDataDetailedSingleDataDefineArrayOutputWithContext(context.Background())
}

func (i GetProbeDataDetailedSingleDataDefineArray) ToGetProbeDataDetailedSingleDataDefineArrayOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProbeDataDetailedSingleDataDefineArrayOutput)
}

type GetProbeDataDetailedSingleDataDefineOutput struct{ *pulumi.OutputState }

func (GetProbeDataDetailedSingleDataDefineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataDetailedSingleDataDefine)(nil)).Elem()
}

func (o GetProbeDataDetailedSingleDataDefineOutput) ToGetProbeDataDetailedSingleDataDefineOutput() GetProbeDataDetailedSingleDataDefineOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineOutput) ToGetProbeDataDetailedSingleDataDefineOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineOutput {
	return o
}

// Fields.
func (o GetProbeDataDetailedSingleDataDefineOutput) Fields() GetProbeDataDetailedSingleDataDefineFieldArrayOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefine) []GetProbeDataDetailedSingleDataDefineField {
		return v.Fields
	}).(GetProbeDataDetailedSingleDataDefineFieldArrayOutput)
}

// Labels.
func (o GetProbeDataDetailedSingleDataDefineOutput) Labels() GetProbeDataDetailedSingleDataDefineLabelArrayOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefine) []GetProbeDataDetailedSingleDataDefineLabel {
		return v.Labels
	}).(GetProbeDataDetailedSingleDataDefineLabelArrayOutput)
}

// Probe time.
func (o GetProbeDataDetailedSingleDataDefineOutput) ProbeTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefine) int { return v.ProbeTime }).(pulumi.IntOutput)
}

type GetProbeDataDetailedSingleDataDefineArrayOutput struct{ *pulumi.OutputState }

func (GetProbeDataDetailedSingleDataDefineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProbeDataDetailedSingleDataDefine)(nil)).Elem()
}

func (o GetProbeDataDetailedSingleDataDefineArrayOutput) ToGetProbeDataDetailedSingleDataDefineArrayOutput() GetProbeDataDetailedSingleDataDefineArrayOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineArrayOutput) ToGetProbeDataDetailedSingleDataDefineArrayOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineArrayOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineArrayOutput) Index(i pulumi.IntInput) GetProbeDataDetailedSingleDataDefineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProbeDataDetailedSingleDataDefine {
		return vs[0].([]GetProbeDataDetailedSingleDataDefine)[vs[1].(int)]
	}).(GetProbeDataDetailedSingleDataDefineOutput)
}

type GetProbeDataDetailedSingleDataDefineField struct {
	// ID.
	Id int `pulumi:"id"`
	// Custom Field Name/Description.
	Name string `pulumi:"name"`
	// Value.
	Value float64 `pulumi:"value"`
}

// GetProbeDataDetailedSingleDataDefineFieldInput is an input type that accepts GetProbeDataDetailedSingleDataDefineFieldArgs and GetProbeDataDetailedSingleDataDefineFieldOutput values.
// You can construct a concrete instance of `GetProbeDataDetailedSingleDataDefineFieldInput` via:
//
//          GetProbeDataDetailedSingleDataDefineFieldArgs{...}
type GetProbeDataDetailedSingleDataDefineFieldInput interface {
	pulumi.Input

	ToGetProbeDataDetailedSingleDataDefineFieldOutput() GetProbeDataDetailedSingleDataDefineFieldOutput
	ToGetProbeDataDetailedSingleDataDefineFieldOutputWithContext(context.Context) GetProbeDataDetailedSingleDataDefineFieldOutput
}

type GetProbeDataDetailedSingleDataDefineFieldArgs struct {
	// ID.
	Id pulumi.IntInput `pulumi:"id"`
	// Custom Field Name/Description.
	Name pulumi.StringInput `pulumi:"name"`
	// Value.
	Value pulumi.Float64Input `pulumi:"value"`
}

func (GetProbeDataDetailedSingleDataDefineFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineField)(nil)).Elem()
}

func (i GetProbeDataDetailedSingleDataDefineFieldArgs) ToGetProbeDataDetailedSingleDataDefineFieldOutput() GetProbeDataDetailedSingleDataDefineFieldOutput {
	return i.ToGetProbeDataDetailedSingleDataDefineFieldOutputWithContext(context.Background())
}

func (i GetProbeDataDetailedSingleDataDefineFieldArgs) ToGetProbeDataDetailedSingleDataDefineFieldOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProbeDataDetailedSingleDataDefineFieldOutput)
}

// GetProbeDataDetailedSingleDataDefineFieldArrayInput is an input type that accepts GetProbeDataDetailedSingleDataDefineFieldArray and GetProbeDataDetailedSingleDataDefineFieldArrayOutput values.
// You can construct a concrete instance of `GetProbeDataDetailedSingleDataDefineFieldArrayInput` via:
//
//          GetProbeDataDetailedSingleDataDefineFieldArray{ GetProbeDataDetailedSingleDataDefineFieldArgs{...} }
type GetProbeDataDetailedSingleDataDefineFieldArrayInput interface {
	pulumi.Input

	ToGetProbeDataDetailedSingleDataDefineFieldArrayOutput() GetProbeDataDetailedSingleDataDefineFieldArrayOutput
	ToGetProbeDataDetailedSingleDataDefineFieldArrayOutputWithContext(context.Context) GetProbeDataDetailedSingleDataDefineFieldArrayOutput
}

type GetProbeDataDetailedSingleDataDefineFieldArray []GetProbeDataDetailedSingleDataDefineFieldInput

func (GetProbeDataDetailedSingleDataDefineFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProbeDataDetailedSingleDataDefineField)(nil)).Elem()
}

func (i GetProbeDataDetailedSingleDataDefineFieldArray) ToGetProbeDataDetailedSingleDataDefineFieldArrayOutput() GetProbeDataDetailedSingleDataDefineFieldArrayOutput {
	return i.ToGetProbeDataDetailedSingleDataDefineFieldArrayOutputWithContext(context.Background())
}

func (i GetProbeDataDetailedSingleDataDefineFieldArray) ToGetProbeDataDetailedSingleDataDefineFieldArrayOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProbeDataDetailedSingleDataDefineFieldArrayOutput)
}

type GetProbeDataDetailedSingleDataDefineFieldOutput struct{ *pulumi.OutputState }

func (GetProbeDataDetailedSingleDataDefineFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineField)(nil)).Elem()
}

func (o GetProbeDataDetailedSingleDataDefineFieldOutput) ToGetProbeDataDetailedSingleDataDefineFieldOutput() GetProbeDataDetailedSingleDataDefineFieldOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineFieldOutput) ToGetProbeDataDetailedSingleDataDefineFieldOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineFieldOutput {
	return o
}

// ID.
func (o GetProbeDataDetailedSingleDataDefineFieldOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefineField) int { return v.Id }).(pulumi.IntOutput)
}

// Custom Field Name/Description.
func (o GetProbeDataDetailedSingleDataDefineFieldOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefineField) string { return v.Name }).(pulumi.StringOutput)
}

// Value.
func (o GetProbeDataDetailedSingleDataDefineFieldOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefineField) float64 { return v.Value }).(pulumi.Float64Output)
}

type GetProbeDataDetailedSingleDataDefineFieldArrayOutput struct{ *pulumi.OutputState }

func (GetProbeDataDetailedSingleDataDefineFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProbeDataDetailedSingleDataDefineField)(nil)).Elem()
}

func (o GetProbeDataDetailedSingleDataDefineFieldArrayOutput) ToGetProbeDataDetailedSingleDataDefineFieldArrayOutput() GetProbeDataDetailedSingleDataDefineFieldArrayOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineFieldArrayOutput) ToGetProbeDataDetailedSingleDataDefineFieldArrayOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineFieldArrayOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineFieldArrayOutput) Index(i pulumi.IntInput) GetProbeDataDetailedSingleDataDefineFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProbeDataDetailedSingleDataDefineField {
		return vs[0].([]GetProbeDataDetailedSingleDataDefineField)[vs[1].(int)]
	}).(GetProbeDataDetailedSingleDataDefineFieldOutput)
}

type GetProbeDataDetailedSingleDataDefineLabel struct {
	// ID.
	Id int `pulumi:"id"`
	// Custom Field Name/Description.
	Name string `pulumi:"name"`
	// Value.
	Value string `pulumi:"value"`
}

// GetProbeDataDetailedSingleDataDefineLabelInput is an input type that accepts GetProbeDataDetailedSingleDataDefineLabelArgs and GetProbeDataDetailedSingleDataDefineLabelOutput values.
// You can construct a concrete instance of `GetProbeDataDetailedSingleDataDefineLabelInput` via:
//
//          GetProbeDataDetailedSingleDataDefineLabelArgs{...}
type GetProbeDataDetailedSingleDataDefineLabelInput interface {
	pulumi.Input

	ToGetProbeDataDetailedSingleDataDefineLabelOutput() GetProbeDataDetailedSingleDataDefineLabelOutput
	ToGetProbeDataDetailedSingleDataDefineLabelOutputWithContext(context.Context) GetProbeDataDetailedSingleDataDefineLabelOutput
}

type GetProbeDataDetailedSingleDataDefineLabelArgs struct {
	// ID.
	Id pulumi.IntInput `pulumi:"id"`
	// Custom Field Name/Description.
	Name pulumi.StringInput `pulumi:"name"`
	// Value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetProbeDataDetailedSingleDataDefineLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineLabel)(nil)).Elem()
}

func (i GetProbeDataDetailedSingleDataDefineLabelArgs) ToGetProbeDataDetailedSingleDataDefineLabelOutput() GetProbeDataDetailedSingleDataDefineLabelOutput {
	return i.ToGetProbeDataDetailedSingleDataDefineLabelOutputWithContext(context.Background())
}

func (i GetProbeDataDetailedSingleDataDefineLabelArgs) ToGetProbeDataDetailedSingleDataDefineLabelOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProbeDataDetailedSingleDataDefineLabelOutput)
}

// GetProbeDataDetailedSingleDataDefineLabelArrayInput is an input type that accepts GetProbeDataDetailedSingleDataDefineLabelArray and GetProbeDataDetailedSingleDataDefineLabelArrayOutput values.
// You can construct a concrete instance of `GetProbeDataDetailedSingleDataDefineLabelArrayInput` via:
//
//          GetProbeDataDetailedSingleDataDefineLabelArray{ GetProbeDataDetailedSingleDataDefineLabelArgs{...} }
type GetProbeDataDetailedSingleDataDefineLabelArrayInput interface {
	pulumi.Input

	ToGetProbeDataDetailedSingleDataDefineLabelArrayOutput() GetProbeDataDetailedSingleDataDefineLabelArrayOutput
	ToGetProbeDataDetailedSingleDataDefineLabelArrayOutputWithContext(context.Context) GetProbeDataDetailedSingleDataDefineLabelArrayOutput
}

type GetProbeDataDetailedSingleDataDefineLabelArray []GetProbeDataDetailedSingleDataDefineLabelInput

func (GetProbeDataDetailedSingleDataDefineLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProbeDataDetailedSingleDataDefineLabel)(nil)).Elem()
}

func (i GetProbeDataDetailedSingleDataDefineLabelArray) ToGetProbeDataDetailedSingleDataDefineLabelArrayOutput() GetProbeDataDetailedSingleDataDefineLabelArrayOutput {
	return i.ToGetProbeDataDetailedSingleDataDefineLabelArrayOutputWithContext(context.Background())
}

func (i GetProbeDataDetailedSingleDataDefineLabelArray) ToGetProbeDataDetailedSingleDataDefineLabelArrayOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProbeDataDetailedSingleDataDefineLabelArrayOutput)
}

type GetProbeDataDetailedSingleDataDefineLabelOutput struct{ *pulumi.OutputState }

func (GetProbeDataDetailedSingleDataDefineLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineLabel)(nil)).Elem()
}

func (o GetProbeDataDetailedSingleDataDefineLabelOutput) ToGetProbeDataDetailedSingleDataDefineLabelOutput() GetProbeDataDetailedSingleDataDefineLabelOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineLabelOutput) ToGetProbeDataDetailedSingleDataDefineLabelOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineLabelOutput {
	return o
}

// ID.
func (o GetProbeDataDetailedSingleDataDefineLabelOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefineLabel) int { return v.Id }).(pulumi.IntOutput)
}

// Custom Field Name/Description.
func (o GetProbeDataDetailedSingleDataDefineLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefineLabel) string { return v.Name }).(pulumi.StringOutput)
}

// Value.
func (o GetProbeDataDetailedSingleDataDefineLabelOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetProbeDataDetailedSingleDataDefineLabel) string { return v.Value }).(pulumi.StringOutput)
}

type GetProbeDataDetailedSingleDataDefineLabelArrayOutput struct{ *pulumi.OutputState }

func (GetProbeDataDetailedSingleDataDefineLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProbeDataDetailedSingleDataDefineLabel)(nil)).Elem()
}

func (o GetProbeDataDetailedSingleDataDefineLabelArrayOutput) ToGetProbeDataDetailedSingleDataDefineLabelArrayOutput() GetProbeDataDetailedSingleDataDefineLabelArrayOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineLabelArrayOutput) ToGetProbeDataDetailedSingleDataDefineLabelArrayOutputWithContext(ctx context.Context) GetProbeDataDetailedSingleDataDefineLabelArrayOutput {
	return o
}

func (o GetProbeDataDetailedSingleDataDefineLabelArrayOutput) Index(i pulumi.IntInput) GetProbeDataDetailedSingleDataDefineLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProbeDataDetailedSingleDataDefineLabel {
		return vs[0].([]GetProbeDataDetailedSingleDataDefineLabel)[vs[1].(int)]
	}).(GetProbeDataDetailedSingleDataDefineLabelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetBatchTasksInput)(nil)).Elem(), TaskSetBatchTasksArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetBatchTasksPtrInput)(nil)).Elem(), TaskSetBatchTasksArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNodeNodeDefineInput)(nil)).Elem(), GetNodeNodeDefineArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNodeNodeDefineArrayInput)(nil)).Elem(), GetNodeNodeDefineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineInput)(nil)).Elem(), GetProbeDataDetailedSingleDataDefineArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineArrayInput)(nil)).Elem(), GetProbeDataDetailedSingleDataDefineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineFieldInput)(nil)).Elem(), GetProbeDataDetailedSingleDataDefineFieldArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineFieldArrayInput)(nil)).Elem(), GetProbeDataDetailedSingleDataDefineFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineLabelInput)(nil)).Elem(), GetProbeDataDetailedSingleDataDefineLabelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProbeDataDetailedSingleDataDefineLabelArrayInput)(nil)).Elem(), GetProbeDataDetailedSingleDataDefineLabelArray{})
	pulumi.RegisterOutputType(TaskSetBatchTasksOutput{})
	pulumi.RegisterOutputType(TaskSetBatchTasksPtrOutput{})
	pulumi.RegisterOutputType(GetNodeNodeDefineOutput{})
	pulumi.RegisterOutputType(GetNodeNodeDefineArrayOutput{})
	pulumi.RegisterOutputType(GetProbeDataDetailedSingleDataDefineOutput{})
	pulumi.RegisterOutputType(GetProbeDataDetailedSingleDataDefineArrayOutput{})
	pulumi.RegisterOutputType(GetProbeDataDetailedSingleDataDefineFieldOutput{})
	pulumi.RegisterOutputType(GetProbeDataDetailedSingleDataDefineFieldArrayOutput{})
	pulumi.RegisterOutputType(GetProbeDataDetailedSingleDataDefineLabelOutput{})
	pulumi.RegisterOutputType(GetProbeDataDetailedSingleDataDefineLabelArrayOutput{})
}
