// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cvm hpcCluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cvm.NewHpcCluster(ctx, "hpcCluster", &Cvm.HpcClusterArgs{
// 			Remark: pulumi.String("create for test"),
// 			Zone:   pulumi.String("ap-beijing-6"),
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
// cvm hpc_cluster can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cvm/hpcCluster:HpcCluster hpc_cluster hpc_cluster_id
// ```
type HpcCluster struct {
	pulumi.CustomResourceState

	// Name of Hpc Cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Remark of Hpc Cluster.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Available zone.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewHpcCluster registers a new resource with the given unique name, arguments, and options.
func NewHpcCluster(ctx *pulumi.Context,
	name string, args *HpcClusterArgs, opts ...pulumi.ResourceOption) (*HpcCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource HpcCluster
	err := ctx.RegisterResource("tencentcloud:Cvm/hpcCluster:HpcCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHpcCluster gets an existing HpcCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHpcCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HpcClusterState, opts ...pulumi.ResourceOption) (*HpcCluster, error) {
	var resource HpcCluster
	err := ctx.ReadResource("tencentcloud:Cvm/hpcCluster:HpcCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HpcCluster resources.
type hpcClusterState struct {
	// Name of Hpc Cluster.
	Name *string `pulumi:"name"`
	// Remark of Hpc Cluster.
	Remark *string `pulumi:"remark"`
	// Available zone.
	Zone *string `pulumi:"zone"`
}

type HpcClusterState struct {
	// Name of Hpc Cluster.
	Name pulumi.StringPtrInput
	// Remark of Hpc Cluster.
	Remark pulumi.StringPtrInput
	// Available zone.
	Zone pulumi.StringPtrInput
}

func (HpcClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hpcClusterState)(nil)).Elem()
}

type hpcClusterArgs struct {
	// Name of Hpc Cluster.
	Name *string `pulumi:"name"`
	// Remark of Hpc Cluster.
	Remark *string `pulumi:"remark"`
	// Available zone.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a HpcCluster resource.
type HpcClusterArgs struct {
	// Name of Hpc Cluster.
	Name pulumi.StringPtrInput
	// Remark of Hpc Cluster.
	Remark pulumi.StringPtrInput
	// Available zone.
	Zone pulumi.StringInput
}

func (HpcClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hpcClusterArgs)(nil)).Elem()
}

type HpcClusterInput interface {
	pulumi.Input

	ToHpcClusterOutput() HpcClusterOutput
	ToHpcClusterOutputWithContext(ctx context.Context) HpcClusterOutput
}

func (*HpcCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**HpcCluster)(nil)).Elem()
}

func (i *HpcCluster) ToHpcClusterOutput() HpcClusterOutput {
	return i.ToHpcClusterOutputWithContext(context.Background())
}

func (i *HpcCluster) ToHpcClusterOutputWithContext(ctx context.Context) HpcClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HpcClusterOutput)
}

// HpcClusterArrayInput is an input type that accepts HpcClusterArray and HpcClusterArrayOutput values.
// You can construct a concrete instance of `HpcClusterArrayInput` via:
//
//          HpcClusterArray{ HpcClusterArgs{...} }
type HpcClusterArrayInput interface {
	pulumi.Input

	ToHpcClusterArrayOutput() HpcClusterArrayOutput
	ToHpcClusterArrayOutputWithContext(context.Context) HpcClusterArrayOutput
}

type HpcClusterArray []HpcClusterInput

func (HpcClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HpcCluster)(nil)).Elem()
}

func (i HpcClusterArray) ToHpcClusterArrayOutput() HpcClusterArrayOutput {
	return i.ToHpcClusterArrayOutputWithContext(context.Background())
}

func (i HpcClusterArray) ToHpcClusterArrayOutputWithContext(ctx context.Context) HpcClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HpcClusterArrayOutput)
}

// HpcClusterMapInput is an input type that accepts HpcClusterMap and HpcClusterMapOutput values.
// You can construct a concrete instance of `HpcClusterMapInput` via:
//
//          HpcClusterMap{ "key": HpcClusterArgs{...} }
type HpcClusterMapInput interface {
	pulumi.Input

	ToHpcClusterMapOutput() HpcClusterMapOutput
	ToHpcClusterMapOutputWithContext(context.Context) HpcClusterMapOutput
}

type HpcClusterMap map[string]HpcClusterInput

func (HpcClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HpcCluster)(nil)).Elem()
}

func (i HpcClusterMap) ToHpcClusterMapOutput() HpcClusterMapOutput {
	return i.ToHpcClusterMapOutputWithContext(context.Background())
}

func (i HpcClusterMap) ToHpcClusterMapOutputWithContext(ctx context.Context) HpcClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HpcClusterMapOutput)
}

type HpcClusterOutput struct{ *pulumi.OutputState }

func (HpcClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HpcCluster)(nil)).Elem()
}

func (o HpcClusterOutput) ToHpcClusterOutput() HpcClusterOutput {
	return o
}

func (o HpcClusterOutput) ToHpcClusterOutputWithContext(ctx context.Context) HpcClusterOutput {
	return o
}

// Name of Hpc Cluster.
func (o HpcClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HpcCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Remark of Hpc Cluster.
func (o HpcClusterOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HpcCluster) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Available zone.
func (o HpcClusterOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *HpcCluster) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type HpcClusterArrayOutput struct{ *pulumi.OutputState }

func (HpcClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HpcCluster)(nil)).Elem()
}

func (o HpcClusterArrayOutput) ToHpcClusterArrayOutput() HpcClusterArrayOutput {
	return o
}

func (o HpcClusterArrayOutput) ToHpcClusterArrayOutputWithContext(ctx context.Context) HpcClusterArrayOutput {
	return o
}

func (o HpcClusterArrayOutput) Index(i pulumi.IntInput) HpcClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HpcCluster {
		return vs[0].([]*HpcCluster)[vs[1].(int)]
	}).(HpcClusterOutput)
}

type HpcClusterMapOutput struct{ *pulumi.OutputState }

func (HpcClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HpcCluster)(nil)).Elem()
}

func (o HpcClusterMapOutput) ToHpcClusterMapOutput() HpcClusterMapOutput {
	return o
}

func (o HpcClusterMapOutput) ToHpcClusterMapOutputWithContext(ctx context.Context) HpcClusterMapOutput {
	return o
}

func (o HpcClusterMapOutput) MapIndex(k pulumi.StringInput) HpcClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HpcCluster {
		return vs[0].(map[string]*HpcCluster)[vs[1].(string)]
	}).(HpcClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HpcClusterInput)(nil)).Elem(), &HpcCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*HpcClusterArrayInput)(nil)).Elem(), HpcClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HpcClusterMapInput)(nil)).Elem(), HpcClusterMap{})
	pulumi.RegisterOutputType(HpcClusterOutput{})
	pulumi.RegisterOutputType(HpcClusterArrayOutput{})
	pulumi.RegisterOutputType(HpcClusterMapOutput{})
}
