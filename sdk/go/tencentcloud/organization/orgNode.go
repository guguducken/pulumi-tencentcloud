// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a organization orgNode
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Organization"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Organization.NewOrgNode(ctx, "orgNode", &Organization.OrgNodeArgs{
// 			ParentNodeId: pulumi.Int(2003721),
// 			Remark:       pulumi.String("for terraform test"),
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
// organization org_node can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Organization/orgNode:OrgNode org_node orgNode_id
// ```
type OrgNode struct {
	pulumi.CustomResourceState

	// Node creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Node name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parent node ID.
	ParentNodeId pulumi.IntOutput `pulumi:"parentNodeId"`
	// Notes.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Node update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrgNode registers a new resource with the given unique name, arguments, and options.
func NewOrgNode(ctx *pulumi.Context,
	name string, args *OrgNodeArgs, opts ...pulumi.ResourceOption) (*OrgNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentNodeId == nil {
		return nil, errors.New("invalid value for required argument 'ParentNodeId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource OrgNode
	err := ctx.RegisterResource("tencentcloud:Organization/orgNode:OrgNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgNode gets an existing OrgNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgNodeState, opts ...pulumi.ResourceOption) (*OrgNode, error) {
	var resource OrgNode
	err := ctx.ReadResource("tencentcloud:Organization/orgNode:OrgNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgNode resources.
type orgNodeState struct {
	// Node creation time.
	CreateTime *string `pulumi:"createTime"`
	// Node name.
	Name *string `pulumi:"name"`
	// Parent node ID.
	ParentNodeId *int `pulumi:"parentNodeId"`
	// Notes.
	Remark *string `pulumi:"remark"`
	// Node update time.
	UpdateTime *string `pulumi:"updateTime"`
}

type OrgNodeState struct {
	// Node creation time.
	CreateTime pulumi.StringPtrInput
	// Node name.
	Name pulumi.StringPtrInput
	// Parent node ID.
	ParentNodeId pulumi.IntPtrInput
	// Notes.
	Remark pulumi.StringPtrInput
	// Node update time.
	UpdateTime pulumi.StringPtrInput
}

func (OrgNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgNodeState)(nil)).Elem()
}

type orgNodeArgs struct {
	// Node name.
	Name *string `pulumi:"name"`
	// Parent node ID.
	ParentNodeId int `pulumi:"parentNodeId"`
	// Notes.
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a OrgNode resource.
type OrgNodeArgs struct {
	// Node name.
	Name pulumi.StringPtrInput
	// Parent node ID.
	ParentNodeId pulumi.IntInput
	// Notes.
	Remark pulumi.StringPtrInput
}

func (OrgNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgNodeArgs)(nil)).Elem()
}

type OrgNodeInput interface {
	pulumi.Input

	ToOrgNodeOutput() OrgNodeOutput
	ToOrgNodeOutputWithContext(ctx context.Context) OrgNodeOutput
}

func (*OrgNode) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgNode)(nil)).Elem()
}

func (i *OrgNode) ToOrgNodeOutput() OrgNodeOutput {
	return i.ToOrgNodeOutputWithContext(context.Background())
}

func (i *OrgNode) ToOrgNodeOutputWithContext(ctx context.Context) OrgNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgNodeOutput)
}

// OrgNodeArrayInput is an input type that accepts OrgNodeArray and OrgNodeArrayOutput values.
// You can construct a concrete instance of `OrgNodeArrayInput` via:
//
//          OrgNodeArray{ OrgNodeArgs{...} }
type OrgNodeArrayInput interface {
	pulumi.Input

	ToOrgNodeArrayOutput() OrgNodeArrayOutput
	ToOrgNodeArrayOutputWithContext(context.Context) OrgNodeArrayOutput
}

type OrgNodeArray []OrgNodeInput

func (OrgNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgNode)(nil)).Elem()
}

func (i OrgNodeArray) ToOrgNodeArrayOutput() OrgNodeArrayOutput {
	return i.ToOrgNodeArrayOutputWithContext(context.Background())
}

func (i OrgNodeArray) ToOrgNodeArrayOutputWithContext(ctx context.Context) OrgNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgNodeArrayOutput)
}

// OrgNodeMapInput is an input type that accepts OrgNodeMap and OrgNodeMapOutput values.
// You can construct a concrete instance of `OrgNodeMapInput` via:
//
//          OrgNodeMap{ "key": OrgNodeArgs{...} }
type OrgNodeMapInput interface {
	pulumi.Input

	ToOrgNodeMapOutput() OrgNodeMapOutput
	ToOrgNodeMapOutputWithContext(context.Context) OrgNodeMapOutput
}

type OrgNodeMap map[string]OrgNodeInput

func (OrgNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgNode)(nil)).Elem()
}

func (i OrgNodeMap) ToOrgNodeMapOutput() OrgNodeMapOutput {
	return i.ToOrgNodeMapOutputWithContext(context.Background())
}

func (i OrgNodeMap) ToOrgNodeMapOutputWithContext(ctx context.Context) OrgNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgNodeMapOutput)
}

type OrgNodeOutput struct{ *pulumi.OutputState }

func (OrgNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgNode)(nil)).Elem()
}

func (o OrgNodeOutput) ToOrgNodeOutput() OrgNodeOutput {
	return o
}

func (o OrgNodeOutput) ToOrgNodeOutputWithContext(ctx context.Context) OrgNodeOutput {
	return o
}

// Node creation time.
func (o OrgNodeOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgNode) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Node name.
func (o OrgNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgNode) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parent node ID.
func (o OrgNodeOutput) ParentNodeId() pulumi.IntOutput {
	return o.ApplyT(func(v *OrgNode) pulumi.IntOutput { return v.ParentNodeId }).(pulumi.IntOutput)
}

// Notes.
func (o OrgNodeOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgNode) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Node update time.
func (o OrgNodeOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgNode) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type OrgNodeArrayOutput struct{ *pulumi.OutputState }

func (OrgNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgNode)(nil)).Elem()
}

func (o OrgNodeArrayOutput) ToOrgNodeArrayOutput() OrgNodeArrayOutput {
	return o
}

func (o OrgNodeArrayOutput) ToOrgNodeArrayOutputWithContext(ctx context.Context) OrgNodeArrayOutput {
	return o
}

func (o OrgNodeArrayOutput) Index(i pulumi.IntInput) OrgNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgNode {
		return vs[0].([]*OrgNode)[vs[1].(int)]
	}).(OrgNodeOutput)
}

type OrgNodeMapOutput struct{ *pulumi.OutputState }

func (OrgNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgNode)(nil)).Elem()
}

func (o OrgNodeMapOutput) ToOrgNodeMapOutput() OrgNodeMapOutput {
	return o
}

func (o OrgNodeMapOutput) ToOrgNodeMapOutputWithContext(ctx context.Context) OrgNodeMapOutput {
	return o
}

func (o OrgNodeMapOutput) MapIndex(k pulumi.StringInput) OrgNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgNode {
		return vs[0].(map[string]*OrgNode)[vs[1].(string)]
	}).(OrgNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgNodeInput)(nil)).Elem(), &OrgNode{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgNodeArrayInput)(nil)).Elem(), OrgNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgNodeMapInput)(nil)).Elem(), OrgNodeMap{})
	pulumi.RegisterOutputType(OrgNodeOutput{})
	pulumi.RegisterOutputType(OrgNodeArrayOutput{})
	pulumi.RegisterOutputType(OrgNodeMapOutput{})
}
