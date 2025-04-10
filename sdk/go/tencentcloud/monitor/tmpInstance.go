// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor tmpInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Monitor.NewTmpInstance(ctx, "tmpInstance", &Monitor.TmpInstanceArgs{
// 			DataRetentionTime: pulumi.Int(30),
// 			InstanceName:      pulumi.String("demo"),
// 			SubnetId:          pulumi.String("subnet-rdkj0agk"),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 			VpcId: pulumi.String("vpc-2hfyray3"),
// 			Zone:  pulumi.String("ap-guangzhou-3"),
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
// monitor tmpInstance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Monitor/tmpInstance:TmpInstance tmpInstance tmpInstance_id
// ```
type TmpInstance struct {
	pulumi.CustomResourceState

	// Prometheus HTTP API root address.
	ApiRootPath pulumi.StringOutput `pulumi:"apiRootPath"`
	// Data retention time.
	DataRetentionTime pulumi.IntOutput `pulumi:"dataRetentionTime"`
	// Instance name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Instance IPv4 address.
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// Proxy address.
	ProxyAddress pulumi.StringOutput `pulumi:"proxyAddress"`
	// Prometheus remote write address.
	RemoteWrite pulumi.StringOutput `pulumi:"remoteWrite"`
	// Subnet Id.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Vpc Id.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Available zone.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTmpInstance registers a new resource with the given unique name, arguments, and options.
func NewTmpInstance(ctx *pulumi.Context,
	name string, args *TmpInstanceArgs, opts ...pulumi.ResourceOption) (*TmpInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataRetentionTime == nil {
		return nil, errors.New("invalid value for required argument 'DataRetentionTime'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TmpInstance
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpInstance:TmpInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpInstance gets an existing TmpInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpInstanceState, opts ...pulumi.ResourceOption) (*TmpInstance, error) {
	var resource TmpInstance
	err := ctx.ReadResource("tencentcloud:Monitor/tmpInstance:TmpInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpInstance resources.
type tmpInstanceState struct {
	// Prometheus HTTP API root address.
	ApiRootPath *string `pulumi:"apiRootPath"`
	// Data retention time.
	DataRetentionTime *int `pulumi:"dataRetentionTime"`
	// Instance name.
	InstanceName *string `pulumi:"instanceName"`
	// Instance IPv4 address.
	Ipv4Address *string `pulumi:"ipv4Address"`
	// Proxy address.
	ProxyAddress *string `pulumi:"proxyAddress"`
	// Prometheus remote write address.
	RemoteWrite *string `pulumi:"remoteWrite"`
	// Subnet Id.
	SubnetId *string `pulumi:"subnetId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc Id.
	VpcId *string `pulumi:"vpcId"`
	// Available zone.
	Zone *string `pulumi:"zone"`
}

type TmpInstanceState struct {
	// Prometheus HTTP API root address.
	ApiRootPath pulumi.StringPtrInput
	// Data retention time.
	DataRetentionTime pulumi.IntPtrInput
	// Instance name.
	InstanceName pulumi.StringPtrInput
	// Instance IPv4 address.
	Ipv4Address pulumi.StringPtrInput
	// Proxy address.
	ProxyAddress pulumi.StringPtrInput
	// Prometheus remote write address.
	RemoteWrite pulumi.StringPtrInput
	// Subnet Id.
	SubnetId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// Vpc Id.
	VpcId pulumi.StringPtrInput
	// Available zone.
	Zone pulumi.StringPtrInput
}

func (TmpInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpInstanceState)(nil)).Elem()
}

type tmpInstanceArgs struct {
	// Data retention time.
	DataRetentionTime int `pulumi:"dataRetentionTime"`
	// Instance name.
	InstanceName string `pulumi:"instanceName"`
	// Subnet Id.
	SubnetId string `pulumi:"subnetId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc Id.
	VpcId string `pulumi:"vpcId"`
	// Available zone.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a TmpInstance resource.
type TmpInstanceArgs struct {
	// Data retention time.
	DataRetentionTime pulumi.IntInput
	// Instance name.
	InstanceName pulumi.StringInput
	// Subnet Id.
	SubnetId pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
	// Vpc Id.
	VpcId pulumi.StringInput
	// Available zone.
	Zone pulumi.StringInput
}

func (TmpInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpInstanceArgs)(nil)).Elem()
}

type TmpInstanceInput interface {
	pulumi.Input

	ToTmpInstanceOutput() TmpInstanceOutput
	ToTmpInstanceOutputWithContext(ctx context.Context) TmpInstanceOutput
}

func (*TmpInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpInstance)(nil)).Elem()
}

func (i *TmpInstance) ToTmpInstanceOutput() TmpInstanceOutput {
	return i.ToTmpInstanceOutputWithContext(context.Background())
}

func (i *TmpInstance) ToTmpInstanceOutputWithContext(ctx context.Context) TmpInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpInstanceOutput)
}

// TmpInstanceArrayInput is an input type that accepts TmpInstanceArray and TmpInstanceArrayOutput values.
// You can construct a concrete instance of `TmpInstanceArrayInput` via:
//
//          TmpInstanceArray{ TmpInstanceArgs{...} }
type TmpInstanceArrayInput interface {
	pulumi.Input

	ToTmpInstanceArrayOutput() TmpInstanceArrayOutput
	ToTmpInstanceArrayOutputWithContext(context.Context) TmpInstanceArrayOutput
}

type TmpInstanceArray []TmpInstanceInput

func (TmpInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpInstance)(nil)).Elem()
}

func (i TmpInstanceArray) ToTmpInstanceArrayOutput() TmpInstanceArrayOutput {
	return i.ToTmpInstanceArrayOutputWithContext(context.Background())
}

func (i TmpInstanceArray) ToTmpInstanceArrayOutputWithContext(ctx context.Context) TmpInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpInstanceArrayOutput)
}

// TmpInstanceMapInput is an input type that accepts TmpInstanceMap and TmpInstanceMapOutput values.
// You can construct a concrete instance of `TmpInstanceMapInput` via:
//
//          TmpInstanceMap{ "key": TmpInstanceArgs{...} }
type TmpInstanceMapInput interface {
	pulumi.Input

	ToTmpInstanceMapOutput() TmpInstanceMapOutput
	ToTmpInstanceMapOutputWithContext(context.Context) TmpInstanceMapOutput
}

type TmpInstanceMap map[string]TmpInstanceInput

func (TmpInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpInstance)(nil)).Elem()
}

func (i TmpInstanceMap) ToTmpInstanceMapOutput() TmpInstanceMapOutput {
	return i.ToTmpInstanceMapOutputWithContext(context.Background())
}

func (i TmpInstanceMap) ToTmpInstanceMapOutputWithContext(ctx context.Context) TmpInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpInstanceMapOutput)
}

type TmpInstanceOutput struct{ *pulumi.OutputState }

func (TmpInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpInstance)(nil)).Elem()
}

func (o TmpInstanceOutput) ToTmpInstanceOutput() TmpInstanceOutput {
	return o
}

func (o TmpInstanceOutput) ToTmpInstanceOutputWithContext(ctx context.Context) TmpInstanceOutput {
	return o
}

// Prometheus HTTP API root address.
func (o TmpInstanceOutput) ApiRootPath() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.ApiRootPath }).(pulumi.StringOutput)
}

// Data retention time.
func (o TmpInstanceOutput) DataRetentionTime() pulumi.IntOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.IntOutput { return v.DataRetentionTime }).(pulumi.IntOutput)
}

// Instance name.
func (o TmpInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Instance IPv4 address.
func (o TmpInstanceOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.Ipv4Address }).(pulumi.StringOutput)
}

// Proxy address.
func (o TmpInstanceOutput) ProxyAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.ProxyAddress }).(pulumi.StringOutput)
}

// Prometheus remote write address.
func (o TmpInstanceOutput) RemoteWrite() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.RemoteWrite }).(pulumi.StringOutput)
}

// Subnet Id.
func (o TmpInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Tag description list.
func (o TmpInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Vpc Id.
func (o TmpInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Available zone.
func (o TmpInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type TmpInstanceArrayOutput struct{ *pulumi.OutputState }

func (TmpInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpInstance)(nil)).Elem()
}

func (o TmpInstanceArrayOutput) ToTmpInstanceArrayOutput() TmpInstanceArrayOutput {
	return o
}

func (o TmpInstanceArrayOutput) ToTmpInstanceArrayOutputWithContext(ctx context.Context) TmpInstanceArrayOutput {
	return o
}

func (o TmpInstanceArrayOutput) Index(i pulumi.IntInput) TmpInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpInstance {
		return vs[0].([]*TmpInstance)[vs[1].(int)]
	}).(TmpInstanceOutput)
}

type TmpInstanceMapOutput struct{ *pulumi.OutputState }

func (TmpInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpInstance)(nil)).Elem()
}

func (o TmpInstanceMapOutput) ToTmpInstanceMapOutput() TmpInstanceMapOutput {
	return o
}

func (o TmpInstanceMapOutput) ToTmpInstanceMapOutputWithContext(ctx context.Context) TmpInstanceMapOutput {
	return o
}

func (o TmpInstanceMapOutput) MapIndex(k pulumi.StringInput) TmpInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpInstance {
		return vs[0].(map[string]*TmpInstance)[vs[1].(string)]
	}).(TmpInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpInstanceInput)(nil)).Elem(), &TmpInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpInstanceArrayInput)(nil)).Elem(), TmpInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpInstanceMapInput)(nil)).Elem(), TmpInstanceMap{})
	pulumi.RegisterOutputType(TmpInstanceOutput{})
	pulumi.RegisterOutputType(TmpInstanceArrayOutput{})
	pulumi.RegisterOutputType(TmpInstanceMapOutput{})
}
