// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a forward rule of layer7 listener.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooProxy, err := Gaap.NewProxy(ctx, "fooProxy", &Gaap.ProxyArgs{
// 			Bandwidth:        pulumi.Int(10),
// 			Concurrent:       pulumi.Int(2),
// 			AccessRegion:     pulumi.String("SouthChina"),
// 			RealserverRegion: pulumi.String("NorthChina"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooLayer7Listener, err := Gaap.NewLayer7Listener(ctx, "fooLayer7Listener", &Gaap.Layer7ListenerArgs{
// 			Protocol: pulumi.String("HTTP"),
// 			Port:     pulumi.Int(80),
// 			ProxyId:  fooProxy.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooRealserver, err := Gaap.NewRealserver(ctx, "fooRealserver", &Gaap.RealserverArgs{
// 			Ip: pulumi.String("1.1.1.1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := Gaap.NewRealserver(ctx, "bar", &Gaap.RealserverArgs{
// 			Ip: pulumi.String("8.8.8.8"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooHttpDomain, err := Gaap.NewHttpDomain(ctx, "fooHttpDomain", &Gaap.HttpDomainArgs{
// 			ListenerId: fooLayer7Listener.ID(),
// 			Domain:     pulumi.String("www.qq.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Gaap.NewHttpRule(ctx, "fooHttpRule", &Gaap.HttpRuleArgs{
// 			ListenerId:        fooLayer7Listener.ID(),
// 			Domain:            fooHttpDomain.Domain,
// 			Path:              pulumi.String("/"),
// 			RealserverType:    pulumi.String("IP"),
// 			HealthCheck:       pulumi.Bool(true),
// 			HealthCheckPath:   pulumi.String("/"),
// 			HealthCheckMethod: pulumi.String("GET"),
// 			HealthCheckStatusCodes: pulumi.IntArray{
// 				pulumi.Int(200),
// 			},
// 			Realservers: gaap.HttpRuleRealserverArray{
// 				&gaap.HttpRuleRealserverArgs{
// 					Id:   fooRealserver.ID(),
// 					Ip:   fooRealserver.Ip,
// 					Port: pulumi.Int(80),
// 				},
// 				&gaap.HttpRuleRealserverArgs{
// 					Id:   bar.ID(),
// 					Ip:   bar.Ip,
// 					Port: pulumi.Int(80),
// 				},
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
// GAAP http rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Gaap/httpRule:HttpRule tencentcloud_gaap_http_rule.foo rule-3bsuu01r
// ```
type HttpRule struct {
	pulumi.CustomResourceState

	// Timeout of the health check response, default value is 2s.
	ConnectTimeout pulumi.IntPtrOutput `pulumi:"connectTimeout"`
	// Forward domain of the forward rule.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost pulumi.StringPtrOutput `pulumi:"forwardHost"`
	// Indicates whether health check is enable.
	HealthCheck pulumi.BoolOutput `pulumi:"healthCheck"`
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod pulumi.StringPtrOutput `pulumi:"healthCheckMethod"`
	// Path of health check. Maximum length is 80.
	HealthCheckPath pulumi.StringPtrOutput `pulumi:"healthCheckPath"`
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	HealthCheckStatusCodes pulumi.IntArrayOutput `pulumi:"healthCheckStatusCodes"`
	// Interval of the health check, default value is 5s.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// ID of the layer7 listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// Path of the forward rule. Maximum length is 80.
	Path pulumi.StringOutput `pulumi:"path"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType pulumi.StringOutput `pulumi:"realserverType"`
	// An information list of GAAP realserver.
	Realservers HttpRuleRealserverArrayOutput `pulumi:"realservers"`
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrOutput `pulumi:"scheduler"`
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni pulumi.StringOutput `pulumi:"sni"`
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch pulumi.StringOutput `pulumi:"sniSwitch"`
}

// NewHttpRule registers a new resource with the given unique name, arguments, and options.
func NewHttpRule(ctx *pulumi.Context,
	name string, args *HttpRuleArgs, opts ...pulumi.ResourceOption) (*HttpRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.HealthCheck == nil {
		return nil, errors.New("invalid value for required argument 'HealthCheck'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.RealserverType == nil {
		return nil, errors.New("invalid value for required argument 'RealserverType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource HttpRule
	err := ctx.RegisterResource("tencentcloud:Gaap/httpRule:HttpRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpRule gets an existing HttpRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpRuleState, opts ...pulumi.ResourceOption) (*HttpRule, error) {
	var resource HttpRule
	err := ctx.ReadResource("tencentcloud:Gaap/httpRule:HttpRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpRule resources.
type httpRuleState struct {
	// Timeout of the health check response, default value is 2s.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// Forward domain of the forward rule.
	Domain *string `pulumi:"domain"`
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost *string `pulumi:"forwardHost"`
	// Indicates whether health check is enable.
	HealthCheck *bool `pulumi:"healthCheck"`
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod *string `pulumi:"healthCheckMethod"`
	// Path of health check. Maximum length is 80.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	HealthCheckStatusCodes []int `pulumi:"healthCheckStatusCodes"`
	// Interval of the health check, default value is 5s.
	Interval *int `pulumi:"interval"`
	// ID of the layer7 listener.
	ListenerId *string `pulumi:"listenerId"`
	// Path of the forward rule. Maximum length is 80.
	Path *string `pulumi:"path"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType *string `pulumi:"realserverType"`
	// An information list of GAAP realserver.
	Realservers []HttpRuleRealserver `pulumi:"realservers"`
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `pulumi:"scheduler"`
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni *string `pulumi:"sni"`
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch *string `pulumi:"sniSwitch"`
}

type HttpRuleState struct {
	// Timeout of the health check response, default value is 2s.
	ConnectTimeout pulumi.IntPtrInput
	// Forward domain of the forward rule.
	Domain pulumi.StringPtrInput
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost pulumi.StringPtrInput
	// Indicates whether health check is enable.
	HealthCheck pulumi.BoolPtrInput
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod pulumi.StringPtrInput
	// Path of health check. Maximum length is 80.
	HealthCheckPath pulumi.StringPtrInput
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	HealthCheckStatusCodes pulumi.IntArrayInput
	// Interval of the health check, default value is 5s.
	Interval pulumi.IntPtrInput
	// ID of the layer7 listener.
	ListenerId pulumi.StringPtrInput
	// Path of the forward rule. Maximum length is 80.
	Path pulumi.StringPtrInput
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType pulumi.StringPtrInput
	// An information list of GAAP realserver.
	Realservers HttpRuleRealserverArrayInput
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrInput
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni pulumi.StringPtrInput
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch pulumi.StringPtrInput
}

func (HttpRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpRuleState)(nil)).Elem()
}

type httpRuleArgs struct {
	// Timeout of the health check response, default value is 2s.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// Forward domain of the forward rule.
	Domain string `pulumi:"domain"`
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost *string `pulumi:"forwardHost"`
	// Indicates whether health check is enable.
	HealthCheck bool `pulumi:"healthCheck"`
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod *string `pulumi:"healthCheckMethod"`
	// Path of health check. Maximum length is 80.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	HealthCheckStatusCodes []int `pulumi:"healthCheckStatusCodes"`
	// Interval of the health check, default value is 5s.
	Interval *int `pulumi:"interval"`
	// ID of the layer7 listener.
	ListenerId string `pulumi:"listenerId"`
	// Path of the forward rule. Maximum length is 80.
	Path string `pulumi:"path"`
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType string `pulumi:"realserverType"`
	// An information list of GAAP realserver.
	Realservers []HttpRuleRealserver `pulumi:"realservers"`
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler *string `pulumi:"scheduler"`
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni *string `pulumi:"sni"`
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch *string `pulumi:"sniSwitch"`
}

// The set of arguments for constructing a HttpRule resource.
type HttpRuleArgs struct {
	// Timeout of the health check response, default value is 2s.
	ConnectTimeout pulumi.IntPtrInput
	// Forward domain of the forward rule.
	Domain pulumi.StringInput
	// The default value of requested host which is forwarded to the realserver by the listener is `default`.
	ForwardHost pulumi.StringPtrInput
	// Indicates whether health check is enable.
	HealthCheck pulumi.BoolInput
	// Method of the health check. Valid value: `GET` and `HEAD`.
	HealthCheckMethod pulumi.StringPtrInput
	// Path of health check. Maximum length is 80.
	HealthCheckPath pulumi.StringPtrInput
	// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
	HealthCheckStatusCodes pulumi.IntArrayInput
	// Interval of the health check, default value is 5s.
	Interval pulumi.IntPtrInput
	// ID of the layer7 listener.
	ListenerId pulumi.StringInput
	// Path of the forward rule. Maximum length is 80.
	Path pulumi.StringInput
	// Type of the realserver. Valid value: `IP` and `DOMAIN`.
	RealserverType pulumi.StringInput
	// An information list of GAAP realserver.
	Realservers HttpRuleRealserverArrayInput
	// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
	Scheduler pulumi.StringPtrInput
	// ServerNameIndication (SNI) is required when the SNI switch is turned on.
	Sni pulumi.StringPtrInput
	// ServerNameIndication (SNI) switch. ON means on and OFF means off.
	SniSwitch pulumi.StringPtrInput
}

func (HttpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpRuleArgs)(nil)).Elem()
}

type HttpRuleInput interface {
	pulumi.Input

	ToHttpRuleOutput() HttpRuleOutput
	ToHttpRuleOutputWithContext(ctx context.Context) HttpRuleOutput
}

func (*HttpRule) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRule)(nil)).Elem()
}

func (i *HttpRule) ToHttpRuleOutput() HttpRuleOutput {
	return i.ToHttpRuleOutputWithContext(context.Background())
}

func (i *HttpRule) ToHttpRuleOutputWithContext(ctx context.Context) HttpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRuleOutput)
}

// HttpRuleArrayInput is an input type that accepts HttpRuleArray and HttpRuleArrayOutput values.
// You can construct a concrete instance of `HttpRuleArrayInput` via:
//
//          HttpRuleArray{ HttpRuleArgs{...} }
type HttpRuleArrayInput interface {
	pulumi.Input

	ToHttpRuleArrayOutput() HttpRuleArrayOutput
	ToHttpRuleArrayOutputWithContext(context.Context) HttpRuleArrayOutput
}

type HttpRuleArray []HttpRuleInput

func (HttpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpRule)(nil)).Elem()
}

func (i HttpRuleArray) ToHttpRuleArrayOutput() HttpRuleArrayOutput {
	return i.ToHttpRuleArrayOutputWithContext(context.Background())
}

func (i HttpRuleArray) ToHttpRuleArrayOutputWithContext(ctx context.Context) HttpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRuleArrayOutput)
}

// HttpRuleMapInput is an input type that accepts HttpRuleMap and HttpRuleMapOutput values.
// You can construct a concrete instance of `HttpRuleMapInput` via:
//
//          HttpRuleMap{ "key": HttpRuleArgs{...} }
type HttpRuleMapInput interface {
	pulumi.Input

	ToHttpRuleMapOutput() HttpRuleMapOutput
	ToHttpRuleMapOutputWithContext(context.Context) HttpRuleMapOutput
}

type HttpRuleMap map[string]HttpRuleInput

func (HttpRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpRule)(nil)).Elem()
}

func (i HttpRuleMap) ToHttpRuleMapOutput() HttpRuleMapOutput {
	return i.ToHttpRuleMapOutputWithContext(context.Background())
}

func (i HttpRuleMap) ToHttpRuleMapOutputWithContext(ctx context.Context) HttpRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRuleMapOutput)
}

type HttpRuleOutput struct{ *pulumi.OutputState }

func (HttpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRule)(nil)).Elem()
}

func (o HttpRuleOutput) ToHttpRuleOutput() HttpRuleOutput {
	return o
}

func (o HttpRuleOutput) ToHttpRuleOutputWithContext(ctx context.Context) HttpRuleOutput {
	return o
}

// Timeout of the health check response, default value is 2s.
func (o HttpRuleOutput) ConnectTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.IntPtrOutput { return v.ConnectTimeout }).(pulumi.IntPtrOutput)
}

// Forward domain of the forward rule.
func (o HttpRuleOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The default value of requested host which is forwarded to the realserver by the listener is `default`.
func (o HttpRuleOutput) ForwardHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringPtrOutput { return v.ForwardHost }).(pulumi.StringPtrOutput)
}

// Indicates whether health check is enable.
func (o HttpRuleOutput) HealthCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.BoolOutput { return v.HealthCheck }).(pulumi.BoolOutput)
}

// Method of the health check. Valid value: `GET` and `HEAD`.
func (o HttpRuleOutput) HealthCheckMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringPtrOutput { return v.HealthCheckMethod }).(pulumi.StringPtrOutput)
}

// Path of health check. Maximum length is 80.
func (o HttpRuleOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringPtrOutput { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

// Return code of confirmed normal. Valid value: `100`, `200`, `300`, `400` and `500`.
func (o HttpRuleOutput) HealthCheckStatusCodes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.IntArrayOutput { return v.HealthCheckStatusCodes }).(pulumi.IntArrayOutput)
}

// Interval of the health check, default value is 5s.
func (o HttpRuleOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// ID of the layer7 listener.
func (o HttpRuleOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// Path of the forward rule. Maximum length is 80.
func (o HttpRuleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Type of the realserver. Valid value: `IP` and `DOMAIN`.
func (o HttpRuleOutput) RealserverType() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringOutput { return v.RealserverType }).(pulumi.StringOutput)
}

// An information list of GAAP realserver.
func (o HttpRuleOutput) Realservers() HttpRuleRealserverArrayOutput {
	return o.ApplyT(func(v *HttpRule) HttpRuleRealserverArrayOutput { return v.Realservers }).(HttpRuleRealserverArrayOutput)
}

// Scheduling policy of the forward rule, default value is `rr`. Valid value: `rr`, `wrr` and `lc`.
func (o HttpRuleOutput) Scheduler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringPtrOutput { return v.Scheduler }).(pulumi.StringPtrOutput)
}

// ServerNameIndication (SNI) is required when the SNI switch is turned on.
func (o HttpRuleOutput) Sni() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringOutput { return v.Sni }).(pulumi.StringOutput)
}

// ServerNameIndication (SNI) switch. ON means on and OFF means off.
func (o HttpRuleOutput) SniSwitch() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRule) pulumi.StringOutput { return v.SniSwitch }).(pulumi.StringOutput)
}

type HttpRuleArrayOutput struct{ *pulumi.OutputState }

func (HttpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpRule)(nil)).Elem()
}

func (o HttpRuleArrayOutput) ToHttpRuleArrayOutput() HttpRuleArrayOutput {
	return o
}

func (o HttpRuleArrayOutput) ToHttpRuleArrayOutputWithContext(ctx context.Context) HttpRuleArrayOutput {
	return o
}

func (o HttpRuleArrayOutput) Index(i pulumi.IntInput) HttpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpRule {
		return vs[0].([]*HttpRule)[vs[1].(int)]
	}).(HttpRuleOutput)
}

type HttpRuleMapOutput struct{ *pulumi.OutputState }

func (HttpRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpRule)(nil)).Elem()
}

func (o HttpRuleMapOutput) ToHttpRuleMapOutput() HttpRuleMapOutput {
	return o
}

func (o HttpRuleMapOutput) ToHttpRuleMapOutputWithContext(ctx context.Context) HttpRuleMapOutput {
	return o
}

func (o HttpRuleMapOutput) MapIndex(k pulumi.StringInput) HttpRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpRule {
		return vs[0].(map[string]*HttpRule)[vs[1].(string)]
	}).(HttpRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRuleInput)(nil)).Elem(), &HttpRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRuleArrayInput)(nil)).Elem(), HttpRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRuleMapInput)(nil)).Elem(), HttpRuleMap{})
	pulumi.RegisterOutputType(HttpRuleOutput{})
	pulumi.RegisterOutputType(HttpRuleArrayOutput{})
	pulumi.RegisterOutputType(HttpRuleMapOutput{})
}
