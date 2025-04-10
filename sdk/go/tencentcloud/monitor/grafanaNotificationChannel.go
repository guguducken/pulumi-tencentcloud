// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor grafanaNotificationChannel
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
// 		_, err := Monitor.NewGrafanaNotificationChannel(ctx, "grafanaNotificationChannel", &Monitor.GrafanaNotificationChannelArgs{
// 			ChannelName: pulumi.String("create-channel"),
// 			ExtraOrgIds: pulumi.StringArray{},
// 			InstanceId:  pulumi.String("grafana-50nj6v00"),
// 			OrgId:       pulumi.Int(1),
// 			Receivers: pulumi.StringArray{
// 				pulumi.String("Consumer-6vkna7pevq"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GrafanaNotificationChannel struct {
	pulumi.CustomResourceState

	// plugin id.
	ChannelId pulumi.StringOutput `pulumi:"channelId"`
	// channel name.
	ChannelName pulumi.StringOutput `pulumi:"channelName"`
	// extra grafana organization id list, default to 1 representing Main Org.
	ExtraOrgIds pulumi.StringArrayOutput `pulumi:"extraOrgIds"`
	// grafana instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Grafana organization which channel will be installed, default to 1 representing Main Org.
	OrgId pulumi.IntOutput `pulumi:"orgId"`
	// cloud monitor notification template notice-id list.
	Receivers pulumi.StringArrayOutput `pulumi:"receivers"`
}

// NewGrafanaNotificationChannel registers a new resource with the given unique name, arguments, and options.
func NewGrafanaNotificationChannel(ctx *pulumi.Context,
	name string, args *GrafanaNotificationChannelArgs, opts ...pulumi.ResourceOption) (*GrafanaNotificationChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource GrafanaNotificationChannel
	err := ctx.RegisterResource("tencentcloud:Monitor/grafanaNotificationChannel:GrafanaNotificationChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrafanaNotificationChannel gets an existing GrafanaNotificationChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrafanaNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrafanaNotificationChannelState, opts ...pulumi.ResourceOption) (*GrafanaNotificationChannel, error) {
	var resource GrafanaNotificationChannel
	err := ctx.ReadResource("tencentcloud:Monitor/grafanaNotificationChannel:GrafanaNotificationChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GrafanaNotificationChannel resources.
type grafanaNotificationChannelState struct {
	// plugin id.
	ChannelId *string `pulumi:"channelId"`
	// channel name.
	ChannelName *string `pulumi:"channelName"`
	// extra grafana organization id list, default to 1 representing Main Org.
	ExtraOrgIds []string `pulumi:"extraOrgIds"`
	// grafana instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Grafana organization which channel will be installed, default to 1 representing Main Org.
	OrgId *int `pulumi:"orgId"`
	// cloud monitor notification template notice-id list.
	Receivers []string `pulumi:"receivers"`
}

type GrafanaNotificationChannelState struct {
	// plugin id.
	ChannelId pulumi.StringPtrInput
	// channel name.
	ChannelName pulumi.StringPtrInput
	// extra grafana organization id list, default to 1 representing Main Org.
	ExtraOrgIds pulumi.StringArrayInput
	// grafana instance id.
	InstanceId pulumi.StringPtrInput
	// Grafana organization which channel will be installed, default to 1 representing Main Org.
	OrgId pulumi.IntPtrInput
	// cloud monitor notification template notice-id list.
	Receivers pulumi.StringArrayInput
}

func (GrafanaNotificationChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaNotificationChannelState)(nil)).Elem()
}

type grafanaNotificationChannelArgs struct {
	// channel name.
	ChannelName *string `pulumi:"channelName"`
	// extra grafana organization id list, default to 1 representing Main Org.
	ExtraOrgIds []string `pulumi:"extraOrgIds"`
	// grafana instance id.
	InstanceId string `pulumi:"instanceId"`
	// Grafana organization which channel will be installed, default to 1 representing Main Org.
	OrgId *int `pulumi:"orgId"`
	// cloud monitor notification template notice-id list.
	Receivers []string `pulumi:"receivers"`
}

// The set of arguments for constructing a GrafanaNotificationChannel resource.
type GrafanaNotificationChannelArgs struct {
	// channel name.
	ChannelName pulumi.StringPtrInput
	// extra grafana organization id list, default to 1 representing Main Org.
	ExtraOrgIds pulumi.StringArrayInput
	// grafana instance id.
	InstanceId pulumi.StringInput
	// Grafana organization which channel will be installed, default to 1 representing Main Org.
	OrgId pulumi.IntPtrInput
	// cloud monitor notification template notice-id list.
	Receivers pulumi.StringArrayInput
}

func (GrafanaNotificationChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaNotificationChannelArgs)(nil)).Elem()
}

type GrafanaNotificationChannelInput interface {
	pulumi.Input

	ToGrafanaNotificationChannelOutput() GrafanaNotificationChannelOutput
	ToGrafanaNotificationChannelOutputWithContext(ctx context.Context) GrafanaNotificationChannelOutput
}

func (*GrafanaNotificationChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaNotificationChannel)(nil)).Elem()
}

func (i *GrafanaNotificationChannel) ToGrafanaNotificationChannelOutput() GrafanaNotificationChannelOutput {
	return i.ToGrafanaNotificationChannelOutputWithContext(context.Background())
}

func (i *GrafanaNotificationChannel) ToGrafanaNotificationChannelOutputWithContext(ctx context.Context) GrafanaNotificationChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaNotificationChannelOutput)
}

// GrafanaNotificationChannelArrayInput is an input type that accepts GrafanaNotificationChannelArray and GrafanaNotificationChannelArrayOutput values.
// You can construct a concrete instance of `GrafanaNotificationChannelArrayInput` via:
//
//          GrafanaNotificationChannelArray{ GrafanaNotificationChannelArgs{...} }
type GrafanaNotificationChannelArrayInput interface {
	pulumi.Input

	ToGrafanaNotificationChannelArrayOutput() GrafanaNotificationChannelArrayOutput
	ToGrafanaNotificationChannelArrayOutputWithContext(context.Context) GrafanaNotificationChannelArrayOutput
}

type GrafanaNotificationChannelArray []GrafanaNotificationChannelInput

func (GrafanaNotificationChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrafanaNotificationChannel)(nil)).Elem()
}

func (i GrafanaNotificationChannelArray) ToGrafanaNotificationChannelArrayOutput() GrafanaNotificationChannelArrayOutput {
	return i.ToGrafanaNotificationChannelArrayOutputWithContext(context.Background())
}

func (i GrafanaNotificationChannelArray) ToGrafanaNotificationChannelArrayOutputWithContext(ctx context.Context) GrafanaNotificationChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaNotificationChannelArrayOutput)
}

// GrafanaNotificationChannelMapInput is an input type that accepts GrafanaNotificationChannelMap and GrafanaNotificationChannelMapOutput values.
// You can construct a concrete instance of `GrafanaNotificationChannelMapInput` via:
//
//          GrafanaNotificationChannelMap{ "key": GrafanaNotificationChannelArgs{...} }
type GrafanaNotificationChannelMapInput interface {
	pulumi.Input

	ToGrafanaNotificationChannelMapOutput() GrafanaNotificationChannelMapOutput
	ToGrafanaNotificationChannelMapOutputWithContext(context.Context) GrafanaNotificationChannelMapOutput
}

type GrafanaNotificationChannelMap map[string]GrafanaNotificationChannelInput

func (GrafanaNotificationChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrafanaNotificationChannel)(nil)).Elem()
}

func (i GrafanaNotificationChannelMap) ToGrafanaNotificationChannelMapOutput() GrafanaNotificationChannelMapOutput {
	return i.ToGrafanaNotificationChannelMapOutputWithContext(context.Background())
}

func (i GrafanaNotificationChannelMap) ToGrafanaNotificationChannelMapOutputWithContext(ctx context.Context) GrafanaNotificationChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaNotificationChannelMapOutput)
}

type GrafanaNotificationChannelOutput struct{ *pulumi.OutputState }

func (GrafanaNotificationChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaNotificationChannel)(nil)).Elem()
}

func (o GrafanaNotificationChannelOutput) ToGrafanaNotificationChannelOutput() GrafanaNotificationChannelOutput {
	return o
}

func (o GrafanaNotificationChannelOutput) ToGrafanaNotificationChannelOutputWithContext(ctx context.Context) GrafanaNotificationChannelOutput {
	return o
}

// plugin id.
func (o GrafanaNotificationChannelOutput) ChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaNotificationChannel) pulumi.StringOutput { return v.ChannelId }).(pulumi.StringOutput)
}

// channel name.
func (o GrafanaNotificationChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaNotificationChannel) pulumi.StringOutput { return v.ChannelName }).(pulumi.StringOutput)
}

// extra grafana organization id list, default to 1 representing Main Org.
func (o GrafanaNotificationChannelOutput) ExtraOrgIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GrafanaNotificationChannel) pulumi.StringArrayOutput { return v.ExtraOrgIds }).(pulumi.StringArrayOutput)
}

// grafana instance id.
func (o GrafanaNotificationChannelOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaNotificationChannel) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Grafana organization which channel will be installed, default to 1 representing Main Org.
func (o GrafanaNotificationChannelOutput) OrgId() pulumi.IntOutput {
	return o.ApplyT(func(v *GrafanaNotificationChannel) pulumi.IntOutput { return v.OrgId }).(pulumi.IntOutput)
}

// cloud monitor notification template notice-id list.
func (o GrafanaNotificationChannelOutput) Receivers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GrafanaNotificationChannel) pulumi.StringArrayOutput { return v.Receivers }).(pulumi.StringArrayOutput)
}

type GrafanaNotificationChannelArrayOutput struct{ *pulumi.OutputState }

func (GrafanaNotificationChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrafanaNotificationChannel)(nil)).Elem()
}

func (o GrafanaNotificationChannelArrayOutput) ToGrafanaNotificationChannelArrayOutput() GrafanaNotificationChannelArrayOutput {
	return o
}

func (o GrafanaNotificationChannelArrayOutput) ToGrafanaNotificationChannelArrayOutputWithContext(ctx context.Context) GrafanaNotificationChannelArrayOutput {
	return o
}

func (o GrafanaNotificationChannelArrayOutput) Index(i pulumi.IntInput) GrafanaNotificationChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GrafanaNotificationChannel {
		return vs[0].([]*GrafanaNotificationChannel)[vs[1].(int)]
	}).(GrafanaNotificationChannelOutput)
}

type GrafanaNotificationChannelMapOutput struct{ *pulumi.OutputState }

func (GrafanaNotificationChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrafanaNotificationChannel)(nil)).Elem()
}

func (o GrafanaNotificationChannelMapOutput) ToGrafanaNotificationChannelMapOutput() GrafanaNotificationChannelMapOutput {
	return o
}

func (o GrafanaNotificationChannelMapOutput) ToGrafanaNotificationChannelMapOutputWithContext(ctx context.Context) GrafanaNotificationChannelMapOutput {
	return o
}

func (o GrafanaNotificationChannelMapOutput) MapIndex(k pulumi.StringInput) GrafanaNotificationChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GrafanaNotificationChannel {
		return vs[0].(map[string]*GrafanaNotificationChannel)[vs[1].(string)]
	}).(GrafanaNotificationChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrafanaNotificationChannelInput)(nil)).Elem(), &GrafanaNotificationChannel{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrafanaNotificationChannelArrayInput)(nil)).Elem(), GrafanaNotificationChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrafanaNotificationChannelMapInput)(nil)).Elem(), GrafanaNotificationChannelMap{})
	pulumi.RegisterOutputType(GrafanaNotificationChannelOutput{})
	pulumi.RegisterOutputType(GrafanaNotificationChannelArrayOutput{})
	pulumi.RegisterOutputType(GrafanaNotificationChannelMapOutput{})
}
