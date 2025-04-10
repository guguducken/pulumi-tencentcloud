// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a tdmq namespace.
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
// 		foo, err := Tdmq.NewInstance(ctx, "foo", &Tdmq.InstanceArgs{
// 			ClusterName: pulumi.String("example"),
// 			Remark:      pulumi.String("this is description."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Tdmq.NewNamespace(ctx, "bar", &Tdmq.NamespaceArgs{
// 			ClusterId:   foo.ID(),
// 			EnvironName: pulumi.String("example"),
// 			MsgTtl:      pulumi.Int(300),
// 			Remark:      pulumi.String("this is description."),
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
// Tdmq namespace can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tdmq/namespace:Namespace test namespace_id
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// The Dedicated Cluster Id.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name of namespace to be created.
	EnvironName pulumi.StringOutput `pulumi:"environName"`
	// The expiration time of unconsumed message.
	MsgTtl pulumi.IntOutput `pulumi:"msgTtl"`
	// Description of the namespace.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
	RetentionPolicy pulumi.MapOutput `pulumi:"retentionPolicy"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.EnvironName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironName'")
	}
	if args.MsgTtl == nil {
		return nil, errors.New("invalid value for required argument 'MsgTtl'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("tencentcloud:Tdmq/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("tencentcloud:Tdmq/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// The Dedicated Cluster Id.
	ClusterId *string `pulumi:"clusterId"`
	// The name of namespace to be created.
	EnvironName *string `pulumi:"environName"`
	// The expiration time of unconsumed message.
	MsgTtl *int `pulumi:"msgTtl"`
	// Description of the namespace.
	Remark *string `pulumi:"remark"`
	// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
	RetentionPolicy map[string]interface{} `pulumi:"retentionPolicy"`
}

type NamespaceState struct {
	// The Dedicated Cluster Id.
	ClusterId pulumi.StringPtrInput
	// The name of namespace to be created.
	EnvironName pulumi.StringPtrInput
	// The expiration time of unconsumed message.
	MsgTtl pulumi.IntPtrInput
	// Description of the namespace.
	Remark pulumi.StringPtrInput
	// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
	RetentionPolicy pulumi.MapInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// The Dedicated Cluster Id.
	ClusterId string `pulumi:"clusterId"`
	// The name of namespace to be created.
	EnvironName string `pulumi:"environName"`
	// The expiration time of unconsumed message.
	MsgTtl int `pulumi:"msgTtl"`
	// Description of the namespace.
	Remark *string `pulumi:"remark"`
	// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
	RetentionPolicy map[string]interface{} `pulumi:"retentionPolicy"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// The Dedicated Cluster Id.
	ClusterId pulumi.StringInput
	// The name of namespace to be created.
	EnvironName pulumi.StringInput
	// The expiration time of unconsumed message.
	MsgTtl pulumi.IntInput
	// Description of the namespace.
	Remark pulumi.StringPtrInput
	// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
	RetentionPolicy pulumi.MapInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//          NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//          NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// The Dedicated Cluster Id.
func (o NamespaceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The name of namespace to be created.
func (o NamespaceOutput) EnvironName() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.EnvironName }).(pulumi.StringOutput)
}

// The expiration time of unconsumed message.
func (o NamespaceOutput) MsgTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *Namespace) pulumi.IntOutput { return v.MsgTtl }).(pulumi.IntOutput)
}

// Description of the namespace.
func (o NamespaceOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
func (o NamespaceOutput) RetentionPolicy() pulumi.MapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.MapOutput { return v.RetentionPolicy }).(pulumi.MapOutput)
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].([]*Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].(map[string]*Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceArrayInput)(nil)).Elem(), NamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceMapInput)(nil)).Elem(), NamespaceMap{})
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
