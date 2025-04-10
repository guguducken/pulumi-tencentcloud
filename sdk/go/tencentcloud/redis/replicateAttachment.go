// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a redis replicateAttachment
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Redis.NewReplicateAttachment(ctx, "replicateAttachment", &Redis.ReplicateAttachmentArgs{
// 			GroupId:          pulumi.String("crs-rpl-c1nl9rpv"),
// 			InstanceIds:      pulumi.StringArray{},
// 			MasterInstanceId: pulumi.String("crs-c1nl9rpv"),
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
// redis replicate_attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Redis/replicateAttachment:ReplicateAttachment replicate_attachment replicate_attachment_id
// ```
type ReplicateAttachment struct {
	pulumi.CustomResourceState

	// The ID of group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// All instance ids of the replication group.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// The ID of master instance.
	MasterInstanceId pulumi.StringOutput `pulumi:"masterInstanceId"`
}

// NewReplicateAttachment registers a new resource with the given unique name, arguments, and options.
func NewReplicateAttachment(ctx *pulumi.Context,
	name string, args *ReplicateAttachmentArgs, opts ...pulumi.ResourceOption) (*ReplicateAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.InstanceIds == nil {
		return nil, errors.New("invalid value for required argument 'InstanceIds'")
	}
	if args.MasterInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'MasterInstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReplicateAttachment
	err := ctx.RegisterResource("tencentcloud:Redis/replicateAttachment:ReplicateAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicateAttachment gets an existing ReplicateAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicateAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicateAttachmentState, opts ...pulumi.ResourceOption) (*ReplicateAttachment, error) {
	var resource ReplicateAttachment
	err := ctx.ReadResource("tencentcloud:Redis/replicateAttachment:ReplicateAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicateAttachment resources.
type replicateAttachmentState struct {
	// The ID of group.
	GroupId *string `pulumi:"groupId"`
	// All instance ids of the replication group.
	InstanceIds []string `pulumi:"instanceIds"`
	// The ID of master instance.
	MasterInstanceId *string `pulumi:"masterInstanceId"`
}

type ReplicateAttachmentState struct {
	// The ID of group.
	GroupId pulumi.StringPtrInput
	// All instance ids of the replication group.
	InstanceIds pulumi.StringArrayInput
	// The ID of master instance.
	MasterInstanceId pulumi.StringPtrInput
}

func (ReplicateAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicateAttachmentState)(nil)).Elem()
}

type replicateAttachmentArgs struct {
	// The ID of group.
	GroupId string `pulumi:"groupId"`
	// All instance ids of the replication group.
	InstanceIds []string `pulumi:"instanceIds"`
	// The ID of master instance.
	MasterInstanceId string `pulumi:"masterInstanceId"`
}

// The set of arguments for constructing a ReplicateAttachment resource.
type ReplicateAttachmentArgs struct {
	// The ID of group.
	GroupId pulumi.StringInput
	// All instance ids of the replication group.
	InstanceIds pulumi.StringArrayInput
	// The ID of master instance.
	MasterInstanceId pulumi.StringInput
}

func (ReplicateAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicateAttachmentArgs)(nil)).Elem()
}

type ReplicateAttachmentInput interface {
	pulumi.Input

	ToReplicateAttachmentOutput() ReplicateAttachmentOutput
	ToReplicateAttachmentOutputWithContext(ctx context.Context) ReplicateAttachmentOutput
}

func (*ReplicateAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicateAttachment)(nil)).Elem()
}

func (i *ReplicateAttachment) ToReplicateAttachmentOutput() ReplicateAttachmentOutput {
	return i.ToReplicateAttachmentOutputWithContext(context.Background())
}

func (i *ReplicateAttachment) ToReplicateAttachmentOutputWithContext(ctx context.Context) ReplicateAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicateAttachmentOutput)
}

// ReplicateAttachmentArrayInput is an input type that accepts ReplicateAttachmentArray and ReplicateAttachmentArrayOutput values.
// You can construct a concrete instance of `ReplicateAttachmentArrayInput` via:
//
//          ReplicateAttachmentArray{ ReplicateAttachmentArgs{...} }
type ReplicateAttachmentArrayInput interface {
	pulumi.Input

	ToReplicateAttachmentArrayOutput() ReplicateAttachmentArrayOutput
	ToReplicateAttachmentArrayOutputWithContext(context.Context) ReplicateAttachmentArrayOutput
}

type ReplicateAttachmentArray []ReplicateAttachmentInput

func (ReplicateAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicateAttachment)(nil)).Elem()
}

func (i ReplicateAttachmentArray) ToReplicateAttachmentArrayOutput() ReplicateAttachmentArrayOutput {
	return i.ToReplicateAttachmentArrayOutputWithContext(context.Background())
}

func (i ReplicateAttachmentArray) ToReplicateAttachmentArrayOutputWithContext(ctx context.Context) ReplicateAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicateAttachmentArrayOutput)
}

// ReplicateAttachmentMapInput is an input type that accepts ReplicateAttachmentMap and ReplicateAttachmentMapOutput values.
// You can construct a concrete instance of `ReplicateAttachmentMapInput` via:
//
//          ReplicateAttachmentMap{ "key": ReplicateAttachmentArgs{...} }
type ReplicateAttachmentMapInput interface {
	pulumi.Input

	ToReplicateAttachmentMapOutput() ReplicateAttachmentMapOutput
	ToReplicateAttachmentMapOutputWithContext(context.Context) ReplicateAttachmentMapOutput
}

type ReplicateAttachmentMap map[string]ReplicateAttachmentInput

func (ReplicateAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicateAttachment)(nil)).Elem()
}

func (i ReplicateAttachmentMap) ToReplicateAttachmentMapOutput() ReplicateAttachmentMapOutput {
	return i.ToReplicateAttachmentMapOutputWithContext(context.Background())
}

func (i ReplicateAttachmentMap) ToReplicateAttachmentMapOutputWithContext(ctx context.Context) ReplicateAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicateAttachmentMapOutput)
}

type ReplicateAttachmentOutput struct{ *pulumi.OutputState }

func (ReplicateAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicateAttachment)(nil)).Elem()
}

func (o ReplicateAttachmentOutput) ToReplicateAttachmentOutput() ReplicateAttachmentOutput {
	return o
}

func (o ReplicateAttachmentOutput) ToReplicateAttachmentOutputWithContext(ctx context.Context) ReplicateAttachmentOutput {
	return o
}

// The ID of group.
func (o ReplicateAttachmentOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicateAttachment) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// All instance ids of the replication group.
func (o ReplicateAttachmentOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReplicateAttachment) pulumi.StringArrayOutput { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// The ID of master instance.
func (o ReplicateAttachmentOutput) MasterInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicateAttachment) pulumi.StringOutput { return v.MasterInstanceId }).(pulumi.StringOutput)
}

type ReplicateAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ReplicateAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicateAttachment)(nil)).Elem()
}

func (o ReplicateAttachmentArrayOutput) ToReplicateAttachmentArrayOutput() ReplicateAttachmentArrayOutput {
	return o
}

func (o ReplicateAttachmentArrayOutput) ToReplicateAttachmentArrayOutputWithContext(ctx context.Context) ReplicateAttachmentArrayOutput {
	return o
}

func (o ReplicateAttachmentArrayOutput) Index(i pulumi.IntInput) ReplicateAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicateAttachment {
		return vs[0].([]*ReplicateAttachment)[vs[1].(int)]
	}).(ReplicateAttachmentOutput)
}

type ReplicateAttachmentMapOutput struct{ *pulumi.OutputState }

func (ReplicateAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicateAttachment)(nil)).Elem()
}

func (o ReplicateAttachmentMapOutput) ToReplicateAttachmentMapOutput() ReplicateAttachmentMapOutput {
	return o
}

func (o ReplicateAttachmentMapOutput) ToReplicateAttachmentMapOutputWithContext(ctx context.Context) ReplicateAttachmentMapOutput {
	return o
}

func (o ReplicateAttachmentMapOutput) MapIndex(k pulumi.StringInput) ReplicateAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicateAttachment {
		return vs[0].(map[string]*ReplicateAttachment)[vs[1].(string)]
	}).(ReplicateAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicateAttachmentInput)(nil)).Elem(), &ReplicateAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicateAttachmentArrayInput)(nil)).Elem(), ReplicateAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicateAttachmentMapInput)(nil)).Elem(), ReplicateAttachmentMap{})
	pulumi.RegisterOutputType(ReplicateAttachmentOutput{})
	pulumi.RegisterOutputType(ReplicateAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ReplicateAttachmentMapOutput{})
}
