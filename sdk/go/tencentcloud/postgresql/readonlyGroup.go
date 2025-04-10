// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create postgresql readonly group.
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
// 		_, err := Postgresql.NewReadonlyGroup(ctx, "group", &Postgresql.ReadonlyGroupArgs{
// 			MasterDbInstanceId:       pulumi.String("postgres-gzg9jb2n"),
// 			MaxReplayLag:             pulumi.Int(100),
// 			MaxReplayLatency:         pulumi.Int(512),
// 			MinDelayEliminateReserve: pulumi.Int(1),
// 			ProjectId:                pulumi.Int(0),
// 			ReplayLagEliminate:       pulumi.Int(1),
// 			ReplayLatencyEliminate:   pulumi.Int(1),
// 			SubnetId:                 pulumi.String("subnet-enm92y0m"),
// 			VpcId:                    pulumi.String("vpc-86v957zb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ReadonlyGroup struct {
	pulumi.CustomResourceState

	// Create time of the postgresql instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Primary instance ID.
	MasterDbInstanceId pulumi.StringOutput `pulumi:"masterDbInstanceId"`
	// Delay threshold in ms.
	MaxReplayLag pulumi.IntOutput `pulumi:"maxReplayLag"`
	// Delayed log size threshold in MB.
	MaxReplayLatency pulumi.IntOutput `pulumi:"maxReplayLatency"`
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve pulumi.IntOutput `pulumi:"minDelayEliminateReserve"`
	// RO group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Project ID.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate pulumi.IntOutput `pulumi:"replayLagEliminate"`
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate pulumi.IntOutput `pulumi:"replayLatencyEliminate"`
	// ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
	SecurityGroupsIds pulumi.StringArrayOutput `pulumi:"securityGroupsIds"`
	// VPC subnet ID.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewReadonlyGroup registers a new resource with the given unique name, arguments, and options.
func NewReadonlyGroup(ctx *pulumi.Context,
	name string, args *ReadonlyGroupArgs, opts ...pulumi.ResourceOption) (*ReadonlyGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MasterDbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'MasterDbInstanceId'")
	}
	if args.MaxReplayLag == nil {
		return nil, errors.New("invalid value for required argument 'MaxReplayLag'")
	}
	if args.MaxReplayLatency == nil {
		return nil, errors.New("invalid value for required argument 'MaxReplayLatency'")
	}
	if args.MinDelayEliminateReserve == nil {
		return nil, errors.New("invalid value for required argument 'MinDelayEliminateReserve'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ReplayLagEliminate == nil {
		return nil, errors.New("invalid value for required argument 'ReplayLagEliminate'")
	}
	if args.ReplayLatencyEliminate == nil {
		return nil, errors.New("invalid value for required argument 'ReplayLatencyEliminate'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReadonlyGroup
	err := ctx.RegisterResource("tencentcloud:Postgresql/readonlyGroup:ReadonlyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadonlyGroup gets an existing ReadonlyGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadonlyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadonlyGroupState, opts ...pulumi.ResourceOption) (*ReadonlyGroup, error) {
	var resource ReadonlyGroup
	err := ctx.ReadResource("tencentcloud:Postgresql/readonlyGroup:ReadonlyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadonlyGroup resources.
type readonlyGroupState struct {
	// Create time of the postgresql instance.
	CreateTime *string `pulumi:"createTime"`
	// Primary instance ID.
	MasterDbInstanceId *string `pulumi:"masterDbInstanceId"`
	// Delay threshold in ms.
	MaxReplayLag *int `pulumi:"maxReplayLag"`
	// Delayed log size threshold in MB.
	MaxReplayLatency *int `pulumi:"maxReplayLatency"`
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve *int `pulumi:"minDelayEliminateReserve"`
	// RO group name.
	Name *string `pulumi:"name"`
	// Project ID.
	ProjectId *int `pulumi:"projectId"`
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate *int `pulumi:"replayLagEliminate"`
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate *int `pulumi:"replayLatencyEliminate"`
	// ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
	SecurityGroupsIds []string `pulumi:"securityGroupsIds"`
	// VPC subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type ReadonlyGroupState struct {
	// Create time of the postgresql instance.
	CreateTime pulumi.StringPtrInput
	// Primary instance ID.
	MasterDbInstanceId pulumi.StringPtrInput
	// Delay threshold in ms.
	MaxReplayLag pulumi.IntPtrInput
	// Delayed log size threshold in MB.
	MaxReplayLatency pulumi.IntPtrInput
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve pulumi.IntPtrInput
	// RO group name.
	Name pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.IntPtrInput
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate pulumi.IntPtrInput
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate pulumi.IntPtrInput
	// ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
	SecurityGroupsIds pulumi.StringArrayInput
	// VPC subnet ID.
	SubnetId pulumi.StringPtrInput
	// VPC ID.
	VpcId pulumi.StringPtrInput
}

func (ReadonlyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*readonlyGroupState)(nil)).Elem()
}

type readonlyGroupArgs struct {
	// Primary instance ID.
	MasterDbInstanceId string `pulumi:"masterDbInstanceId"`
	// Delay threshold in ms.
	MaxReplayLag int `pulumi:"maxReplayLag"`
	// Delayed log size threshold in MB.
	MaxReplayLatency int `pulumi:"maxReplayLatency"`
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve int `pulumi:"minDelayEliminateReserve"`
	// RO group name.
	Name *string `pulumi:"name"`
	// Project ID.
	ProjectId int `pulumi:"projectId"`
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate int `pulumi:"replayLagEliminate"`
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate int `pulumi:"replayLatencyEliminate"`
	// ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
	SecurityGroupsIds []string `pulumi:"securityGroupsIds"`
	// VPC subnet ID.
	SubnetId string `pulumi:"subnetId"`
	// VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ReadonlyGroup resource.
type ReadonlyGroupArgs struct {
	// Primary instance ID.
	MasterDbInstanceId pulumi.StringInput
	// Delay threshold in ms.
	MaxReplayLag pulumi.IntInput
	// Delayed log size threshold in MB.
	MaxReplayLatency pulumi.IntInput
	// The minimum number of read-only replicas that must be retained in an RO group.
	MinDelayEliminateReserve pulumi.IntInput
	// RO group name.
	Name pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.IntInput
	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLagEliminate pulumi.IntInput
	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
	ReplayLatencyEliminate pulumi.IntInput
	// ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
	SecurityGroupsIds pulumi.StringArrayInput
	// VPC subnet ID.
	SubnetId pulumi.StringInput
	// VPC ID.
	VpcId pulumi.StringInput
}

func (ReadonlyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readonlyGroupArgs)(nil)).Elem()
}

type ReadonlyGroupInput interface {
	pulumi.Input

	ToReadonlyGroupOutput() ReadonlyGroupOutput
	ToReadonlyGroupOutputWithContext(ctx context.Context) ReadonlyGroupOutput
}

func (*ReadonlyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadonlyGroup)(nil)).Elem()
}

func (i *ReadonlyGroup) ToReadonlyGroupOutput() ReadonlyGroupOutput {
	return i.ToReadonlyGroupOutputWithContext(context.Background())
}

func (i *ReadonlyGroup) ToReadonlyGroupOutputWithContext(ctx context.Context) ReadonlyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyGroupOutput)
}

// ReadonlyGroupArrayInput is an input type that accepts ReadonlyGroupArray and ReadonlyGroupArrayOutput values.
// You can construct a concrete instance of `ReadonlyGroupArrayInput` via:
//
//          ReadonlyGroupArray{ ReadonlyGroupArgs{...} }
type ReadonlyGroupArrayInput interface {
	pulumi.Input

	ToReadonlyGroupArrayOutput() ReadonlyGroupArrayOutput
	ToReadonlyGroupArrayOutputWithContext(context.Context) ReadonlyGroupArrayOutput
}

type ReadonlyGroupArray []ReadonlyGroupInput

func (ReadonlyGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadonlyGroup)(nil)).Elem()
}

func (i ReadonlyGroupArray) ToReadonlyGroupArrayOutput() ReadonlyGroupArrayOutput {
	return i.ToReadonlyGroupArrayOutputWithContext(context.Background())
}

func (i ReadonlyGroupArray) ToReadonlyGroupArrayOutputWithContext(ctx context.Context) ReadonlyGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyGroupArrayOutput)
}

// ReadonlyGroupMapInput is an input type that accepts ReadonlyGroupMap and ReadonlyGroupMapOutput values.
// You can construct a concrete instance of `ReadonlyGroupMapInput` via:
//
//          ReadonlyGroupMap{ "key": ReadonlyGroupArgs{...} }
type ReadonlyGroupMapInput interface {
	pulumi.Input

	ToReadonlyGroupMapOutput() ReadonlyGroupMapOutput
	ToReadonlyGroupMapOutputWithContext(context.Context) ReadonlyGroupMapOutput
}

type ReadonlyGroupMap map[string]ReadonlyGroupInput

func (ReadonlyGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadonlyGroup)(nil)).Elem()
}

func (i ReadonlyGroupMap) ToReadonlyGroupMapOutput() ReadonlyGroupMapOutput {
	return i.ToReadonlyGroupMapOutputWithContext(context.Background())
}

func (i ReadonlyGroupMap) ToReadonlyGroupMapOutputWithContext(ctx context.Context) ReadonlyGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyGroupMapOutput)
}

type ReadonlyGroupOutput struct{ *pulumi.OutputState }

func (ReadonlyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadonlyGroup)(nil)).Elem()
}

func (o ReadonlyGroupOutput) ToReadonlyGroupOutput() ReadonlyGroupOutput {
	return o
}

func (o ReadonlyGroupOutput) ToReadonlyGroupOutputWithContext(ctx context.Context) ReadonlyGroupOutput {
	return o
}

// Create time of the postgresql instance.
func (o ReadonlyGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Primary instance ID.
func (o ReadonlyGroupOutput) MasterDbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.StringOutput { return v.MasterDbInstanceId }).(pulumi.StringOutput)
}

// Delay threshold in ms.
func (o ReadonlyGroupOutput) MaxReplayLag() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.IntOutput { return v.MaxReplayLag }).(pulumi.IntOutput)
}

// Delayed log size threshold in MB.
func (o ReadonlyGroupOutput) MaxReplayLatency() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.IntOutput { return v.MaxReplayLatency }).(pulumi.IntOutput)
}

// The minimum number of read-only replicas that must be retained in an RO group.
func (o ReadonlyGroupOutput) MinDelayEliminateReserve() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.IntOutput { return v.MinDelayEliminateReserve }).(pulumi.IntOutput)
}

// RO group name.
func (o ReadonlyGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project ID.
func (o ReadonlyGroupOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
func (o ReadonlyGroupOutput) ReplayLagEliminate() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.IntOutput { return v.ReplayLagEliminate }).(pulumi.IntOutput)
}

// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
func (o ReadonlyGroupOutput) ReplayLatencyEliminate() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.IntOutput { return v.ReplayLatencyEliminate }).(pulumi.IntOutput)
}

// ID of security group. If both vpcId and subnetId are not set, this argument should not be set either.
func (o ReadonlyGroupOutput) SecurityGroupsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.StringArrayOutput { return v.SecurityGroupsIds }).(pulumi.StringArrayOutput)
}

// VPC subnet ID.
func (o ReadonlyGroupOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// VPC ID.
func (o ReadonlyGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ReadonlyGroupArrayOutput struct{ *pulumi.OutputState }

func (ReadonlyGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadonlyGroup)(nil)).Elem()
}

func (o ReadonlyGroupArrayOutput) ToReadonlyGroupArrayOutput() ReadonlyGroupArrayOutput {
	return o
}

func (o ReadonlyGroupArrayOutput) ToReadonlyGroupArrayOutputWithContext(ctx context.Context) ReadonlyGroupArrayOutput {
	return o
}

func (o ReadonlyGroupArrayOutput) Index(i pulumi.IntInput) ReadonlyGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReadonlyGroup {
		return vs[0].([]*ReadonlyGroup)[vs[1].(int)]
	}).(ReadonlyGroupOutput)
}

type ReadonlyGroupMapOutput struct{ *pulumi.OutputState }

func (ReadonlyGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadonlyGroup)(nil)).Elem()
}

func (o ReadonlyGroupMapOutput) ToReadonlyGroupMapOutput() ReadonlyGroupMapOutput {
	return o
}

func (o ReadonlyGroupMapOutput) ToReadonlyGroupMapOutputWithContext(ctx context.Context) ReadonlyGroupMapOutput {
	return o
}

func (o ReadonlyGroupMapOutput) MapIndex(k pulumi.StringInput) ReadonlyGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReadonlyGroup {
		return vs[0].(map[string]*ReadonlyGroup)[vs[1].(string)]
	}).(ReadonlyGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyGroupInput)(nil)).Elem(), &ReadonlyGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyGroupArrayInput)(nil)).Elem(), ReadonlyGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyGroupMapInput)(nil)).Elem(), ReadonlyGroupMap{})
	pulumi.RegisterOutputType(ReadonlyGroupOutput{})
	pulumi.RegisterOutputType(ReadonlyGroupArrayOutput{})
	pulumi.RegisterOutputType(ReadonlyGroupMapOutput{})
}
