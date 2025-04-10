// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a NAT gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eip"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Nat"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := Vpc.GetInstances(ctx, &vpc.GetInstancesArgs{
// 			Name: pulumi.StringRef("Default-VPC"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		eipDevDnat, err := Eip.NewInstance(ctx, "eipDevDnat", nil)
// 		if err != nil {
// 			return err
// 		}
// 		newEip, err := Eip.NewInstance(ctx, "newEip", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Nat.NewGateway(ctx, "myNat", &Nat.GatewayArgs{
// 			VpcId:         pulumi.String(foo.InstanceLists[0].VpcId),
// 			MaxConcurrent: pulumi.Int(10000000),
// 			Bandwidth:     pulumi.Int(1000),
// 			Zone:          pulumi.String("ap-guangzhou-3"),
// 			AssignedEipSets: pulumi.StringArray{
// 				eipDevDnat.PublicIp,
// 				newEip.PublicIp,
// 			},
// 			Tags: pulumi.AnyMap{
// 				"tf": pulumi.Any("test"),
// 			},
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
// NAT gateway can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Nat/gateway:Gateway foo nat-1asg3t63
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
	AssignedEipSets pulumi.StringArrayOutput `pulumi:"assignedEipSets"`
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// Create time of the NAT gateway.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
	MaxConcurrent pulumi.IntPtrOutput `pulumi:"maxConcurrent"`
	// Name of the NAT gateway.
	Name pulumi.StringOutput `pulumi:"name"`
	// The available tags within this NAT gateway.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// ID of the vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The availability zone, such as `ap-guangzhou-3`.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssignedEipSets == nil {
		return nil, errors.New("invalid value for required argument 'AssignedEipSets'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Gateway
	err := ctx.RegisterResource("tencentcloud:Nat/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("tencentcloud:Nat/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
	AssignedEipSets []string `pulumi:"assignedEipSets"`
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
	Bandwidth *int `pulumi:"bandwidth"`
	// Create time of the NAT gateway.
	CreatedTime *string `pulumi:"createdTime"`
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
	MaxConcurrent *int `pulumi:"maxConcurrent"`
	// Name of the NAT gateway.
	Name *string `pulumi:"name"`
	// The available tags within this NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc.
	VpcId *string `pulumi:"vpcId"`
	// The availability zone, such as `ap-guangzhou-3`.
	Zone *string `pulumi:"zone"`
}

type GatewayState struct {
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
	AssignedEipSets pulumi.StringArrayInput
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
	Bandwidth pulumi.IntPtrInput
	// Create time of the NAT gateway.
	CreatedTime pulumi.StringPtrInput
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
	MaxConcurrent pulumi.IntPtrInput
	// Name of the NAT gateway.
	Name pulumi.StringPtrInput
	// The available tags within this NAT gateway.
	Tags pulumi.MapInput
	// ID of the vpc.
	VpcId pulumi.StringPtrInput
	// The availability zone, such as `ap-guangzhou-3`.
	Zone pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
	AssignedEipSets []string `pulumi:"assignedEipSets"`
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
	Bandwidth *int `pulumi:"bandwidth"`
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
	MaxConcurrent *int `pulumi:"maxConcurrent"`
	// Name of the NAT gateway.
	Name *string `pulumi:"name"`
	// The available tags within this NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the vpc.
	VpcId string `pulumi:"vpcId"`
	// The availability zone, such as `ap-guangzhou-3`.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
	AssignedEipSets pulumi.StringArrayInput
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
	Bandwidth pulumi.IntPtrInput
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
	MaxConcurrent pulumi.IntPtrInput
	// Name of the NAT gateway.
	Name pulumi.StringPtrInput
	// The available tags within this NAT gateway.
	Tags pulumi.MapInput
	// ID of the vpc.
	VpcId pulumi.StringInput
	// The availability zone, such as `ap-guangzhou-3`.
	Zone pulumi.StringPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//          GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//          GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
func (o GatewayOutput) AssignedEipSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringArrayOutput { return v.AssignedEipSets }).(pulumi.StringArrayOutput)
}

// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
func (o GatewayOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntPtrOutput { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// Create time of the NAT gateway.
func (o GatewayOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
func (o GatewayOutput) MaxConcurrent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntPtrOutput { return v.MaxConcurrent }).(pulumi.IntPtrOutput)
}

// Name of the NAT gateway.
func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The available tags within this NAT gateway.
func (o GatewayOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Gateway) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// ID of the vpc.
func (o GatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The availability zone, such as `ap-guangzhou-3`.
func (o GatewayOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].([]*Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].(map[string]*Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
