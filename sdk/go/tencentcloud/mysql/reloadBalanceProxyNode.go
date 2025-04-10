// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql reloadBalanceProxyNode
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mysql.NewReloadBalanceProxyNode(ctx, "reloadBalanceProxyNode", &Mysql.ReloadBalanceProxyNodeArgs{
// 			ProxyAddressId: pulumi.String("proxyaddr-4wc4y1pq"),
// 			ProxyGroupId:   pulumi.String("proxy-gmi1f78l"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ReloadBalanceProxyNode struct {
	pulumi.CustomResourceState

	// Proxy address id.
	ProxyAddressId pulumi.StringPtrOutput `pulumi:"proxyAddressId"`
	// Proxy id.
	ProxyGroupId pulumi.StringOutput `pulumi:"proxyGroupId"`
}

// NewReloadBalanceProxyNode registers a new resource with the given unique name, arguments, and options.
func NewReloadBalanceProxyNode(ctx *pulumi.Context,
	name string, args *ReloadBalanceProxyNodeArgs, opts ...pulumi.ResourceOption) (*ReloadBalanceProxyNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProxyGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ProxyGroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReloadBalanceProxyNode
	err := ctx.RegisterResource("tencentcloud:Mysql/reloadBalanceProxyNode:ReloadBalanceProxyNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReloadBalanceProxyNode gets an existing ReloadBalanceProxyNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReloadBalanceProxyNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReloadBalanceProxyNodeState, opts ...pulumi.ResourceOption) (*ReloadBalanceProxyNode, error) {
	var resource ReloadBalanceProxyNode
	err := ctx.ReadResource("tencentcloud:Mysql/reloadBalanceProxyNode:ReloadBalanceProxyNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReloadBalanceProxyNode resources.
type reloadBalanceProxyNodeState struct {
	// Proxy address id.
	ProxyAddressId *string `pulumi:"proxyAddressId"`
	// Proxy id.
	ProxyGroupId *string `pulumi:"proxyGroupId"`
}

type ReloadBalanceProxyNodeState struct {
	// Proxy address id.
	ProxyAddressId pulumi.StringPtrInput
	// Proxy id.
	ProxyGroupId pulumi.StringPtrInput
}

func (ReloadBalanceProxyNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*reloadBalanceProxyNodeState)(nil)).Elem()
}

type reloadBalanceProxyNodeArgs struct {
	// Proxy address id.
	ProxyAddressId *string `pulumi:"proxyAddressId"`
	// Proxy id.
	ProxyGroupId string `pulumi:"proxyGroupId"`
}

// The set of arguments for constructing a ReloadBalanceProxyNode resource.
type ReloadBalanceProxyNodeArgs struct {
	// Proxy address id.
	ProxyAddressId pulumi.StringPtrInput
	// Proxy id.
	ProxyGroupId pulumi.StringInput
}

func (ReloadBalanceProxyNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reloadBalanceProxyNodeArgs)(nil)).Elem()
}

type ReloadBalanceProxyNodeInput interface {
	pulumi.Input

	ToReloadBalanceProxyNodeOutput() ReloadBalanceProxyNodeOutput
	ToReloadBalanceProxyNodeOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeOutput
}

func (*ReloadBalanceProxyNode) ElementType() reflect.Type {
	return reflect.TypeOf((**ReloadBalanceProxyNode)(nil)).Elem()
}

func (i *ReloadBalanceProxyNode) ToReloadBalanceProxyNodeOutput() ReloadBalanceProxyNodeOutput {
	return i.ToReloadBalanceProxyNodeOutputWithContext(context.Background())
}

func (i *ReloadBalanceProxyNode) ToReloadBalanceProxyNodeOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReloadBalanceProxyNodeOutput)
}

// ReloadBalanceProxyNodeArrayInput is an input type that accepts ReloadBalanceProxyNodeArray and ReloadBalanceProxyNodeArrayOutput values.
// You can construct a concrete instance of `ReloadBalanceProxyNodeArrayInput` via:
//
//          ReloadBalanceProxyNodeArray{ ReloadBalanceProxyNodeArgs{...} }
type ReloadBalanceProxyNodeArrayInput interface {
	pulumi.Input

	ToReloadBalanceProxyNodeArrayOutput() ReloadBalanceProxyNodeArrayOutput
	ToReloadBalanceProxyNodeArrayOutputWithContext(context.Context) ReloadBalanceProxyNodeArrayOutput
}

type ReloadBalanceProxyNodeArray []ReloadBalanceProxyNodeInput

func (ReloadBalanceProxyNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReloadBalanceProxyNode)(nil)).Elem()
}

func (i ReloadBalanceProxyNodeArray) ToReloadBalanceProxyNodeArrayOutput() ReloadBalanceProxyNodeArrayOutput {
	return i.ToReloadBalanceProxyNodeArrayOutputWithContext(context.Background())
}

func (i ReloadBalanceProxyNodeArray) ToReloadBalanceProxyNodeArrayOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReloadBalanceProxyNodeArrayOutput)
}

// ReloadBalanceProxyNodeMapInput is an input type that accepts ReloadBalanceProxyNodeMap and ReloadBalanceProxyNodeMapOutput values.
// You can construct a concrete instance of `ReloadBalanceProxyNodeMapInput` via:
//
//          ReloadBalanceProxyNodeMap{ "key": ReloadBalanceProxyNodeArgs{...} }
type ReloadBalanceProxyNodeMapInput interface {
	pulumi.Input

	ToReloadBalanceProxyNodeMapOutput() ReloadBalanceProxyNodeMapOutput
	ToReloadBalanceProxyNodeMapOutputWithContext(context.Context) ReloadBalanceProxyNodeMapOutput
}

type ReloadBalanceProxyNodeMap map[string]ReloadBalanceProxyNodeInput

func (ReloadBalanceProxyNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReloadBalanceProxyNode)(nil)).Elem()
}

func (i ReloadBalanceProxyNodeMap) ToReloadBalanceProxyNodeMapOutput() ReloadBalanceProxyNodeMapOutput {
	return i.ToReloadBalanceProxyNodeMapOutputWithContext(context.Background())
}

func (i ReloadBalanceProxyNodeMap) ToReloadBalanceProxyNodeMapOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReloadBalanceProxyNodeMapOutput)
}

type ReloadBalanceProxyNodeOutput struct{ *pulumi.OutputState }

func (ReloadBalanceProxyNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReloadBalanceProxyNode)(nil)).Elem()
}

func (o ReloadBalanceProxyNodeOutput) ToReloadBalanceProxyNodeOutput() ReloadBalanceProxyNodeOutput {
	return o
}

func (o ReloadBalanceProxyNodeOutput) ToReloadBalanceProxyNodeOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeOutput {
	return o
}

// Proxy address id.
func (o ReloadBalanceProxyNodeOutput) ProxyAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReloadBalanceProxyNode) pulumi.StringPtrOutput { return v.ProxyAddressId }).(pulumi.StringPtrOutput)
}

// Proxy id.
func (o ReloadBalanceProxyNodeOutput) ProxyGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReloadBalanceProxyNode) pulumi.StringOutput { return v.ProxyGroupId }).(pulumi.StringOutput)
}

type ReloadBalanceProxyNodeArrayOutput struct{ *pulumi.OutputState }

func (ReloadBalanceProxyNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReloadBalanceProxyNode)(nil)).Elem()
}

func (o ReloadBalanceProxyNodeArrayOutput) ToReloadBalanceProxyNodeArrayOutput() ReloadBalanceProxyNodeArrayOutput {
	return o
}

func (o ReloadBalanceProxyNodeArrayOutput) ToReloadBalanceProxyNodeArrayOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeArrayOutput {
	return o
}

func (o ReloadBalanceProxyNodeArrayOutput) Index(i pulumi.IntInput) ReloadBalanceProxyNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReloadBalanceProxyNode {
		return vs[0].([]*ReloadBalanceProxyNode)[vs[1].(int)]
	}).(ReloadBalanceProxyNodeOutput)
}

type ReloadBalanceProxyNodeMapOutput struct{ *pulumi.OutputState }

func (ReloadBalanceProxyNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReloadBalanceProxyNode)(nil)).Elem()
}

func (o ReloadBalanceProxyNodeMapOutput) ToReloadBalanceProxyNodeMapOutput() ReloadBalanceProxyNodeMapOutput {
	return o
}

func (o ReloadBalanceProxyNodeMapOutput) ToReloadBalanceProxyNodeMapOutputWithContext(ctx context.Context) ReloadBalanceProxyNodeMapOutput {
	return o
}

func (o ReloadBalanceProxyNodeMapOutput) MapIndex(k pulumi.StringInput) ReloadBalanceProxyNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReloadBalanceProxyNode {
		return vs[0].(map[string]*ReloadBalanceProxyNode)[vs[1].(string)]
	}).(ReloadBalanceProxyNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReloadBalanceProxyNodeInput)(nil)).Elem(), &ReloadBalanceProxyNode{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReloadBalanceProxyNodeArrayInput)(nil)).Elem(), ReloadBalanceProxyNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReloadBalanceProxyNodeMapInput)(nil)).Elem(), ReloadBalanceProxyNodeMap{})
	pulumi.RegisterOutputType(ReloadBalanceProxyNodeOutput{})
	pulumi.RegisterOutputType(ReloadBalanceProxyNodeArrayOutput{})
	pulumi.RegisterOutputType(ReloadBalanceProxyNodeMapOutput{})
}
