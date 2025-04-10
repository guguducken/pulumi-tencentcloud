// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tdmqRocketmq group
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cluster, err := Tdmq.NewRocketmqCluster(ctx, "cluster", &Tdmq.RocketmqClusterArgs{
// 			ClusterName: pulumi.String("test_rocketmq"),
// 			Remark:      pulumi.String("test recket mq"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		namespace, err := Tdmq.NewRocketmqNamespace(ctx, "namespace", &Tdmq.RocketmqNamespaceArgs{
// 			ClusterId:     cluster.ClusterId,
// 			NamespaceName: pulumi.String("test_namespace"),
// 			Ttl:           pulumi.Int(65000),
// 			RetentionTime: pulumi.Int(65000),
// 			Remark:        pulumi.String("test namespace"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Tdmq.NewRocketmqGroup(ctx, "group", &Tdmq.RocketmqGroupArgs{
// 			GroupName:       pulumi.String("test_rocketmq_group"),
// 			Namespace:       namespace.NamespaceName,
// 			ReadEnable:      pulumi.Bool(true),
// 			BroadcastEnable: pulumi.Bool(true),
// 			ClusterId:       cluster.ClusterId,
// 			Remark:          pulumi.String("test rocketmq group"),
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
// tdmqRocketmq group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup group group_id
// ```
type RocketmqGroup struct {
	pulumi.CustomResourceState

	// Whether to enable broadcast consumption.
	BroadcastEnable pulumi.BoolOutput `pulumi:"broadcastEnable"`
	// Client protocol.
	ClientProtocol pulumi.StringOutput `pulumi:"clientProtocol"`
	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The number of online consumers.
	ConsumerNum pulumi.IntOutput `pulumi:"consumerNum"`
	// Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
	ConsumerType pulumi.StringOutput `pulumi:"consumerType"`
	// `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
	ConsumptionMode pulumi.IntOutput `pulumi:"consumptionMode"`
	// Creation time in milliseconds.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// Group name (8-64 characters).
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Namespace. Currently, only one namespace is supported.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// Whether to enable consumption.
	ReadEnable pulumi.BoolOutput `pulumi:"readEnable"`
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The number of partitions in a retry topic.
	RetryPartitionNum pulumi.IntOutput `pulumi:"retryPartitionNum"`
	// The total number of heaped messages.
	TotalAccumulative pulumi.IntOutput `pulumi:"totalAccumulative"`
	// Consumption TPS.
	Tps pulumi.IntOutput `pulumi:"tps"`
	// Modification time in milliseconds.
	UpdateTime pulumi.IntOutput `pulumi:"updateTime"`
}

// NewRocketmqGroup registers a new resource with the given unique name, arguments, and options.
func NewRocketmqGroup(ctx *pulumi.Context,
	name string, args *RocketmqGroupArgs, opts ...pulumi.ResourceOption) (*RocketmqGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BroadcastEnable == nil {
		return nil, errors.New("invalid value for required argument 'BroadcastEnable'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.ReadEnable == nil {
		return nil, errors.New("invalid value for required argument 'ReadEnable'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RocketmqGroup
	err := ctx.RegisterResource("tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRocketmqGroup gets an existing RocketmqGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRocketmqGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RocketmqGroupState, opts ...pulumi.ResourceOption) (*RocketmqGroup, error) {
	var resource RocketmqGroup
	err := ctx.ReadResource("tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RocketmqGroup resources.
type rocketmqGroupState struct {
	// Whether to enable broadcast consumption.
	BroadcastEnable *bool `pulumi:"broadcastEnable"`
	// Client protocol.
	ClientProtocol *string `pulumi:"clientProtocol"`
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// The number of online consumers.
	ConsumerNum *int `pulumi:"consumerNum"`
	// Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
	ConsumerType *string `pulumi:"consumerType"`
	// `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
	ConsumptionMode *int `pulumi:"consumptionMode"`
	// Creation time in milliseconds.
	CreateTime *int `pulumi:"createTime"`
	// Group name (8-64 characters).
	GroupName *string `pulumi:"groupName"`
	// Namespace. Currently, only one namespace is supported.
	Namespace *string `pulumi:"namespace"`
	// Whether to enable consumption.
	ReadEnable *bool `pulumi:"readEnable"`
	// Remarks (up to 128 characters).
	Remark *string `pulumi:"remark"`
	// The number of partitions in a retry topic.
	RetryPartitionNum *int `pulumi:"retryPartitionNum"`
	// The total number of heaped messages.
	TotalAccumulative *int `pulumi:"totalAccumulative"`
	// Consumption TPS.
	Tps *int `pulumi:"tps"`
	// Modification time in milliseconds.
	UpdateTime *int `pulumi:"updateTime"`
}

type RocketmqGroupState struct {
	// Whether to enable broadcast consumption.
	BroadcastEnable pulumi.BoolPtrInput
	// Client protocol.
	ClientProtocol pulumi.StringPtrInput
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// The number of online consumers.
	ConsumerNum pulumi.IntPtrInput
	// Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
	ConsumerType pulumi.StringPtrInput
	// `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
	ConsumptionMode pulumi.IntPtrInput
	// Creation time in milliseconds.
	CreateTime pulumi.IntPtrInput
	// Group name (8-64 characters).
	GroupName pulumi.StringPtrInput
	// Namespace. Currently, only one namespace is supported.
	Namespace pulumi.StringPtrInput
	// Whether to enable consumption.
	ReadEnable pulumi.BoolPtrInput
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrInput
	// The number of partitions in a retry topic.
	RetryPartitionNum pulumi.IntPtrInput
	// The total number of heaped messages.
	TotalAccumulative pulumi.IntPtrInput
	// Consumption TPS.
	Tps pulumi.IntPtrInput
	// Modification time in milliseconds.
	UpdateTime pulumi.IntPtrInput
}

func (RocketmqGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqGroupState)(nil)).Elem()
}

type rocketmqGroupArgs struct {
	// Whether to enable broadcast consumption.
	BroadcastEnable bool `pulumi:"broadcastEnable"`
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Group name (8-64 characters).
	GroupName string `pulumi:"groupName"`
	// Namespace. Currently, only one namespace is supported.
	Namespace string `pulumi:"namespace"`
	// Whether to enable consumption.
	ReadEnable bool `pulumi:"readEnable"`
	// Remarks (up to 128 characters).
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a RocketmqGroup resource.
type RocketmqGroupArgs struct {
	// Whether to enable broadcast consumption.
	BroadcastEnable pulumi.BoolInput
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Group name (8-64 characters).
	GroupName pulumi.StringInput
	// Namespace. Currently, only one namespace is supported.
	Namespace pulumi.StringInput
	// Whether to enable consumption.
	ReadEnable pulumi.BoolInput
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrInput
}

func (RocketmqGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqGroupArgs)(nil)).Elem()
}

type RocketmqGroupInput interface {
	pulumi.Input

	ToRocketmqGroupOutput() RocketmqGroupOutput
	ToRocketmqGroupOutputWithContext(ctx context.Context) RocketmqGroupOutput
}

func (*RocketmqGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqGroup)(nil)).Elem()
}

func (i *RocketmqGroup) ToRocketmqGroupOutput() RocketmqGroupOutput {
	return i.ToRocketmqGroupOutputWithContext(context.Background())
}

func (i *RocketmqGroup) ToRocketmqGroupOutputWithContext(ctx context.Context) RocketmqGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqGroupOutput)
}

// RocketmqGroupArrayInput is an input type that accepts RocketmqGroupArray and RocketmqGroupArrayOutput values.
// You can construct a concrete instance of `RocketmqGroupArrayInput` via:
//
//          RocketmqGroupArray{ RocketmqGroupArgs{...} }
type RocketmqGroupArrayInput interface {
	pulumi.Input

	ToRocketmqGroupArrayOutput() RocketmqGroupArrayOutput
	ToRocketmqGroupArrayOutputWithContext(context.Context) RocketmqGroupArrayOutput
}

type RocketmqGroupArray []RocketmqGroupInput

func (RocketmqGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqGroup)(nil)).Elem()
}

func (i RocketmqGroupArray) ToRocketmqGroupArrayOutput() RocketmqGroupArrayOutput {
	return i.ToRocketmqGroupArrayOutputWithContext(context.Background())
}

func (i RocketmqGroupArray) ToRocketmqGroupArrayOutputWithContext(ctx context.Context) RocketmqGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqGroupArrayOutput)
}

// RocketmqGroupMapInput is an input type that accepts RocketmqGroupMap and RocketmqGroupMapOutput values.
// You can construct a concrete instance of `RocketmqGroupMapInput` via:
//
//          RocketmqGroupMap{ "key": RocketmqGroupArgs{...} }
type RocketmqGroupMapInput interface {
	pulumi.Input

	ToRocketmqGroupMapOutput() RocketmqGroupMapOutput
	ToRocketmqGroupMapOutputWithContext(context.Context) RocketmqGroupMapOutput
}

type RocketmqGroupMap map[string]RocketmqGroupInput

func (RocketmqGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqGroup)(nil)).Elem()
}

func (i RocketmqGroupMap) ToRocketmqGroupMapOutput() RocketmqGroupMapOutput {
	return i.ToRocketmqGroupMapOutputWithContext(context.Background())
}

func (i RocketmqGroupMap) ToRocketmqGroupMapOutputWithContext(ctx context.Context) RocketmqGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqGroupMapOutput)
}

type RocketmqGroupOutput struct{ *pulumi.OutputState }

func (RocketmqGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqGroup)(nil)).Elem()
}

func (o RocketmqGroupOutput) ToRocketmqGroupOutput() RocketmqGroupOutput {
	return o
}

func (o RocketmqGroupOutput) ToRocketmqGroupOutputWithContext(ctx context.Context) RocketmqGroupOutput {
	return o
}

// Whether to enable broadcast consumption.
func (o RocketmqGroupOutput) BroadcastEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.BoolOutput { return v.BroadcastEnable }).(pulumi.BoolOutput)
}

// Client protocol.
func (o RocketmqGroupOutput) ClientProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.StringOutput { return v.ClientProtocol }).(pulumi.StringOutput)
}

// Cluster ID.
func (o RocketmqGroupOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The number of online consumers.
func (o RocketmqGroupOutput) ConsumerNum() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.ConsumerNum }).(pulumi.IntOutput)
}

// Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
func (o RocketmqGroupOutput) ConsumerType() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.StringOutput { return v.ConsumerType }).(pulumi.StringOutput)
}

// `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
func (o RocketmqGroupOutput) ConsumptionMode() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.ConsumptionMode }).(pulumi.IntOutput)
}

// Creation time in milliseconds.
func (o RocketmqGroupOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// Group name (8-64 characters).
func (o RocketmqGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Namespace. Currently, only one namespace is supported.
func (o RocketmqGroupOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// Whether to enable consumption.
func (o RocketmqGroupOutput) ReadEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.BoolOutput { return v.ReadEnable }).(pulumi.BoolOutput)
}

// Remarks (up to 128 characters).
func (o RocketmqGroupOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// The number of partitions in a retry topic.
func (o RocketmqGroupOutput) RetryPartitionNum() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.RetryPartitionNum }).(pulumi.IntOutput)
}

// The total number of heaped messages.
func (o RocketmqGroupOutput) TotalAccumulative() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.TotalAccumulative }).(pulumi.IntOutput)
}

// Consumption TPS.
func (o RocketmqGroupOutput) Tps() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.Tps }).(pulumi.IntOutput)
}

// Modification time in milliseconds.
func (o RocketmqGroupOutput) UpdateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqGroup) pulumi.IntOutput { return v.UpdateTime }).(pulumi.IntOutput)
}

type RocketmqGroupArrayOutput struct{ *pulumi.OutputState }

func (RocketmqGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqGroup)(nil)).Elem()
}

func (o RocketmqGroupArrayOutput) ToRocketmqGroupArrayOutput() RocketmqGroupArrayOutput {
	return o
}

func (o RocketmqGroupArrayOutput) ToRocketmqGroupArrayOutputWithContext(ctx context.Context) RocketmqGroupArrayOutput {
	return o
}

func (o RocketmqGroupArrayOutput) Index(i pulumi.IntInput) RocketmqGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RocketmqGroup {
		return vs[0].([]*RocketmqGroup)[vs[1].(int)]
	}).(RocketmqGroupOutput)
}

type RocketmqGroupMapOutput struct{ *pulumi.OutputState }

func (RocketmqGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqGroup)(nil)).Elem()
}

func (o RocketmqGroupMapOutput) ToRocketmqGroupMapOutput() RocketmqGroupMapOutput {
	return o
}

func (o RocketmqGroupMapOutput) ToRocketmqGroupMapOutputWithContext(ctx context.Context) RocketmqGroupMapOutput {
	return o
}

func (o RocketmqGroupMapOutput) MapIndex(k pulumi.StringInput) RocketmqGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RocketmqGroup {
		return vs[0].(map[string]*RocketmqGroup)[vs[1].(string)]
	}).(RocketmqGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqGroupInput)(nil)).Elem(), &RocketmqGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqGroupArrayInput)(nil)).Elem(), RocketmqGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqGroupMapInput)(nil)).Elem(), RocketmqGroupMap{})
	pulumi.RegisterOutputType(RocketmqGroupOutput{})
	pulumi.RegisterOutputType(RocketmqGroupArrayOutput{})
	pulumi.RegisterOutputType(RocketmqGroupMapOutput{})
}
