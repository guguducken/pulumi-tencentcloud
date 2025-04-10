// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CLB listener rule.
//
// > **NOTE:** This resource only be applied to the HTTP or HTTPS listeners.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Clb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Clb.NewListenerRule(ctx, "foo", &Clb.ListenerRuleArgs{
// 			CertificateCaId:         pulumi.String("VfqO4zkB"),
// 			CertificateId:           pulumi.String("VjANRdz8"),
// 			CertificateSslMode:      pulumi.String("MUTUAL"),
// 			ClbId:                   pulumi.String("lb-k2zjp9lv"),
// 			Domain:                  pulumi.String("foo.net"),
// 			HealthCheckHealthNum:    pulumi.Int(3),
// 			HealthCheckHttpCode:     pulumi.Int(2),
// 			HealthCheckHttpDomain:   pulumi.String("Default Domain"),
// 			HealthCheckHttpMethod:   pulumi.String("GET"),
// 			HealthCheckHttpPath:     pulumi.String("Default Path"),
// 			HealthCheckIntervalTime: pulumi.Int(5),
// 			HealthCheckSwitch:       pulumi.Bool(true),
// 			HealthCheckUnhealthNum:  pulumi.Int(3),
// 			ListenerId:              pulumi.String("lbl-hh141sn9"),
// 			Scheduler:               pulumi.String("WRR"),
// 			SessionExpireTime:       pulumi.Int(30),
// 			Url:                     pulumi.String("/bar"),
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
// CLB listener rule can be imported using the id (version >= 1.47.0), e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Clb/listenerRule:ListenerRule foo lb-7a0t6zqb#lbl-hh141sn9#loc-agg236ys
// ```
type ListenerRule struct {
	pulumi.CustomResourceState

	// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateCaId pulumi.StringPtrOutput `pulumi:"certificateCaId"`
	// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateId pulumi.StringPtrOutput `pulumi:"certificateId"`
	// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
	CertificateSslMode pulumi.StringPtrOutput `pulumi:"certificateSslMode"`
	// ID of CLB instance.
	ClbId pulumi.StringOutput `pulumi:"clbId"`
	// Domain name of the listener rule.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
	ForwardType pulumi.StringOutput `pulumi:"forwardType"`
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckHealthNum pulumi.IntOutput `pulumi:"healthCheckHealthNum"`
	// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` 1 means the return value '1xx' is health.  `2`means the return value '2xx' is health.`4`means the return value '3xx' is health.`8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
	HealthCheckHttpCode pulumi.IntOutput `pulumi:"healthCheckHttpCode"`
	// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpDomain pulumi.StringOutput `pulumi:"healthCheckHttpDomain"`
	// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
	HealthCheckHttpMethod pulumi.StringOutput `pulumi:"healthCheckHttpMethod"`
	// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpPath pulumi.StringOutput `pulumi:"healthCheckHttpPath"`
	// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckIntervalTime pulumi.IntOutput `pulumi:"healthCheckIntervalTime"`
	// Indicates whether health check is enabled.
	HealthCheckSwitch pulumi.BoolOutput `pulumi:"healthCheckSwitch"`
	// Time out of health check. The value range is 2-60.
	HealthCheckTimeOut pulumi.IntOutput `pulumi:"healthCheckTimeOut"`
	// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
	HealthCheckType pulumi.StringOutput `pulumi:"healthCheckType"`
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckUnhealthNum pulumi.IntOutput `pulumi:"healthCheckUnhealthNum"`
	// Indicate to apply HTTP2.0 protocol or not.
	Http2Switch pulumi.BoolOutput `pulumi:"http2Switch"`
	// ID of CLB listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// ID of this CLB listener rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	Scheduler pulumi.StringPtrOutput `pulumi:"scheduler"`
	// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	SessionExpireTime pulumi.IntPtrOutput `pulumi:"sessionExpireTime"`
	// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
	TargetType pulumi.StringPtrOutput `pulumi:"targetType"`
	// Url of the listener rule.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewListenerRule registers a new resource with the given unique name, arguments, and options.
func NewListenerRule(ctx *pulumi.Context,
	name string, args *ListenerRuleArgs, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClbId == nil {
		return nil, errors.New("invalid value for required argument 'ClbId'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ListenerRule
	err := ctx.RegisterResource("tencentcloud:Clb/listenerRule:ListenerRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerRule gets an existing ListenerRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerRuleState, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	var resource ListenerRule
	err := ctx.ReadResource("tencentcloud:Clb/listenerRule:ListenerRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerRule resources.
type listenerRuleState struct {
	// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateCaId *string `pulumi:"certificateCaId"`
	// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateId *string `pulumi:"certificateId"`
	// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
	CertificateSslMode *string `pulumi:"certificateSslMode"`
	// ID of CLB instance.
	ClbId *string `pulumi:"clbId"`
	// Domain name of the listener rule.
	Domain *string `pulumi:"domain"`
	// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
	ForwardType *string `pulumi:"forwardType"`
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckHealthNum *int `pulumi:"healthCheckHealthNum"`
	// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` 1 means the return value '1xx' is health.  `2`means the return value '2xx' is health.`4`means the return value '3xx' is health.`8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
	HealthCheckHttpCode *int `pulumi:"healthCheckHttpCode"`
	// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpDomain *string `pulumi:"healthCheckHttpDomain"`
	// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
	HealthCheckHttpMethod *string `pulumi:"healthCheckHttpMethod"`
	// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpPath *string `pulumi:"healthCheckHttpPath"`
	// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckIntervalTime *int `pulumi:"healthCheckIntervalTime"`
	// Indicates whether health check is enabled.
	HealthCheckSwitch *bool `pulumi:"healthCheckSwitch"`
	// Time out of health check. The value range is 2-60.
	HealthCheckTimeOut *int `pulumi:"healthCheckTimeOut"`
	// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckUnhealthNum *int `pulumi:"healthCheckUnhealthNum"`
	// Indicate to apply HTTP2.0 protocol or not.
	Http2Switch *bool `pulumi:"http2Switch"`
	// ID of CLB listener.
	ListenerId *string `pulumi:"listenerId"`
	// ID of this CLB listener rule.
	RuleId *string `pulumi:"ruleId"`
	// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	Scheduler *string `pulumi:"scheduler"`
	// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	SessionExpireTime *int `pulumi:"sessionExpireTime"`
	// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
	TargetType *string `pulumi:"targetType"`
	// Url of the listener rule.
	Url *string `pulumi:"url"`
}

type ListenerRuleState struct {
	// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateCaId pulumi.StringPtrInput
	// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateId pulumi.StringPtrInput
	// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
	CertificateSslMode pulumi.StringPtrInput
	// ID of CLB instance.
	ClbId pulumi.StringPtrInput
	// Domain name of the listener rule.
	Domain pulumi.StringPtrInput
	// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
	ForwardType pulumi.StringPtrInput
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckHealthNum pulumi.IntPtrInput
	// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` 1 means the return value '1xx' is health.  `2`means the return value '2xx' is health.`4`means the return value '3xx' is health.`8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
	HealthCheckHttpCode pulumi.IntPtrInput
	// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpDomain pulumi.StringPtrInput
	// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
	HealthCheckHttpMethod pulumi.StringPtrInput
	// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpPath pulumi.StringPtrInput
	// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckIntervalTime pulumi.IntPtrInput
	// Indicates whether health check is enabled.
	HealthCheckSwitch pulumi.BoolPtrInput
	// Time out of health check. The value range is 2-60.
	HealthCheckTimeOut pulumi.IntPtrInput
	// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
	HealthCheckType pulumi.StringPtrInput
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckUnhealthNum pulumi.IntPtrInput
	// Indicate to apply HTTP2.0 protocol or not.
	Http2Switch pulumi.BoolPtrInput
	// ID of CLB listener.
	ListenerId pulumi.StringPtrInput
	// ID of this CLB listener rule.
	RuleId pulumi.StringPtrInput
	// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	Scheduler pulumi.StringPtrInput
	// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	SessionExpireTime pulumi.IntPtrInput
	// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
	TargetType pulumi.StringPtrInput
	// Url of the listener rule.
	Url pulumi.StringPtrInput
}

func (ListenerRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleState)(nil)).Elem()
}

type listenerRuleArgs struct {
	// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateCaId *string `pulumi:"certificateCaId"`
	// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateId *string `pulumi:"certificateId"`
	// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
	CertificateSslMode *string `pulumi:"certificateSslMode"`
	// ID of CLB instance.
	ClbId string `pulumi:"clbId"`
	// Domain name of the listener rule.
	Domain string `pulumi:"domain"`
	// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
	ForwardType *string `pulumi:"forwardType"`
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckHealthNum *int `pulumi:"healthCheckHealthNum"`
	// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` 1 means the return value '1xx' is health.  `2`means the return value '2xx' is health.`4`means the return value '3xx' is health.`8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
	HealthCheckHttpCode *int `pulumi:"healthCheckHttpCode"`
	// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpDomain *string `pulumi:"healthCheckHttpDomain"`
	// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
	HealthCheckHttpMethod *string `pulumi:"healthCheckHttpMethod"`
	// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpPath *string `pulumi:"healthCheckHttpPath"`
	// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckIntervalTime *int `pulumi:"healthCheckIntervalTime"`
	// Indicates whether health check is enabled.
	HealthCheckSwitch *bool `pulumi:"healthCheckSwitch"`
	// Time out of health check. The value range is 2-60.
	HealthCheckTimeOut *int `pulumi:"healthCheckTimeOut"`
	// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckUnhealthNum *int `pulumi:"healthCheckUnhealthNum"`
	// Indicate to apply HTTP2.0 protocol or not.
	Http2Switch *bool `pulumi:"http2Switch"`
	// ID of CLB listener.
	ListenerId string `pulumi:"listenerId"`
	// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	Scheduler *string `pulumi:"scheduler"`
	// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	SessionExpireTime *int `pulumi:"sessionExpireTime"`
	// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
	TargetType *string `pulumi:"targetType"`
	// Url of the listener rule.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ListenerRule resource.
type ListenerRuleArgs struct {
	// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateCaId pulumi.StringPtrInput
	// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
	CertificateId pulumi.StringPtrInput
	// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
	CertificateSslMode pulumi.StringPtrInput
	// ID of CLB instance.
	ClbId pulumi.StringInput
	// Domain name of the listener rule.
	Domain pulumi.StringInput
	// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
	ForwardType pulumi.StringPtrInput
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckHealthNum pulumi.IntPtrInput
	// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` 1 means the return value '1xx' is health.  `2`means the return value '2xx' is health.`4`means the return value '3xx' is health.`8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
	HealthCheckHttpCode pulumi.IntPtrInput
	// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpDomain pulumi.StringPtrInput
	// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
	HealthCheckHttpMethod pulumi.StringPtrInput
	// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
	HealthCheckHttpPath pulumi.StringPtrInput
	// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckIntervalTime pulumi.IntPtrInput
	// Indicates whether health check is enabled.
	HealthCheckSwitch pulumi.BoolPtrInput
	// Time out of health check. The value range is 2-60.
	HealthCheckTimeOut pulumi.IntPtrInput
	// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
	HealthCheckType pulumi.StringPtrInput
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	HealthCheckUnhealthNum pulumi.IntPtrInput
	// Indicate to apply HTTP2.0 protocol or not.
	Http2Switch pulumi.BoolPtrInput
	// ID of CLB listener.
	ListenerId pulumi.StringInput
	// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	Scheduler pulumi.StringPtrInput
	// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
	SessionExpireTime pulumi.IntPtrInput
	// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
	TargetType pulumi.StringPtrInput
	// Url of the listener rule.
	Url pulumi.StringInput
}

func (ListenerRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleArgs)(nil)).Elem()
}

type ListenerRuleInput interface {
	pulumi.Input

	ToListenerRuleOutput() ListenerRuleOutput
	ToListenerRuleOutputWithContext(ctx context.Context) ListenerRuleOutput
}

func (*ListenerRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRule)(nil)).Elem()
}

func (i *ListenerRule) ToListenerRuleOutput() ListenerRuleOutput {
	return i.ToListenerRuleOutputWithContext(context.Background())
}

func (i *ListenerRule) ToListenerRuleOutputWithContext(ctx context.Context) ListenerRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleOutput)
}

// ListenerRuleArrayInput is an input type that accepts ListenerRuleArray and ListenerRuleArrayOutput values.
// You can construct a concrete instance of `ListenerRuleArrayInput` via:
//
//          ListenerRuleArray{ ListenerRuleArgs{...} }
type ListenerRuleArrayInput interface {
	pulumi.Input

	ToListenerRuleArrayOutput() ListenerRuleArrayOutput
	ToListenerRuleArrayOutputWithContext(context.Context) ListenerRuleArrayOutput
}

type ListenerRuleArray []ListenerRuleInput

func (ListenerRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListenerRule)(nil)).Elem()
}

func (i ListenerRuleArray) ToListenerRuleArrayOutput() ListenerRuleArrayOutput {
	return i.ToListenerRuleArrayOutputWithContext(context.Background())
}

func (i ListenerRuleArray) ToListenerRuleArrayOutputWithContext(ctx context.Context) ListenerRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleArrayOutput)
}

// ListenerRuleMapInput is an input type that accepts ListenerRuleMap and ListenerRuleMapOutput values.
// You can construct a concrete instance of `ListenerRuleMapInput` via:
//
//          ListenerRuleMap{ "key": ListenerRuleArgs{...} }
type ListenerRuleMapInput interface {
	pulumi.Input

	ToListenerRuleMapOutput() ListenerRuleMapOutput
	ToListenerRuleMapOutputWithContext(context.Context) ListenerRuleMapOutput
}

type ListenerRuleMap map[string]ListenerRuleInput

func (ListenerRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListenerRule)(nil)).Elem()
}

func (i ListenerRuleMap) ToListenerRuleMapOutput() ListenerRuleMapOutput {
	return i.ToListenerRuleMapOutputWithContext(context.Background())
}

func (i ListenerRuleMap) ToListenerRuleMapOutputWithContext(ctx context.Context) ListenerRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleMapOutput)
}

type ListenerRuleOutput struct{ *pulumi.OutputState }

func (ListenerRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRule)(nil)).Elem()
}

func (o ListenerRuleOutput) ToListenerRuleOutput() ListenerRuleOutput {
	return o
}

func (o ListenerRuleOutput) ToListenerRuleOutputWithContext(ctx context.Context) ListenerRuleOutput {
	return o
}

// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
func (o ListenerRuleOutput) CertificateCaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringPtrOutput { return v.CertificateCaId }).(pulumi.StringPtrOutput)
}

// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
func (o ListenerRuleOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringPtrOutput { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
func (o ListenerRuleOutput) CertificateSslMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringPtrOutput { return v.CertificateSslMode }).(pulumi.StringPtrOutput)
}

// ID of CLB instance.
func (o ListenerRuleOutput) ClbId() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.ClbId }).(pulumi.StringOutput)
}

// Domain name of the listener rule.
func (o ListenerRuleOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`.
func (o ListenerRuleOutput) ForwardType() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.ForwardType }).(pulumi.StringOutput)
}

// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
func (o ListenerRuleOutput) HealthCheckHealthNum() pulumi.IntOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.IntOutput { return v.HealthCheckHealthNum }).(pulumi.IntOutput)
}

// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` 1 means the return value '1xx' is health.  `2`means the return value '2xx' is health.`4`means the return value '3xx' is health.`8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
func (o ListenerRuleOutput) HealthCheckHttpCode() pulumi.IntOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.IntOutput { return v.HealthCheckHttpCode }).(pulumi.IntOutput)
}

// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
func (o ListenerRuleOutput) HealthCheckHttpDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.HealthCheckHttpDomain }).(pulumi.StringOutput)
}

// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
func (o ListenerRuleOutput) HealthCheckHttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.HealthCheckHttpMethod }).(pulumi.StringOutput)
}

// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
func (o ListenerRuleOutput) HealthCheckHttpPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.HealthCheckHttpPath }).(pulumi.StringOutput)
}

// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
func (o ListenerRuleOutput) HealthCheckIntervalTime() pulumi.IntOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.IntOutput { return v.HealthCheckIntervalTime }).(pulumi.IntOutput)
}

// Indicates whether health check is enabled.
func (o ListenerRuleOutput) HealthCheckSwitch() pulumi.BoolOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.BoolOutput { return v.HealthCheckSwitch }).(pulumi.BoolOutput)
}

// Time out of health check. The value range is 2-60.
func (o ListenerRuleOutput) HealthCheckTimeOut() pulumi.IntOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.IntOutput { return v.HealthCheckTimeOut }).(pulumi.IntOutput)
}

// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
func (o ListenerRuleOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.HealthCheckType }).(pulumi.StringOutput)
}

// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
func (o ListenerRuleOutput) HealthCheckUnhealthNum() pulumi.IntOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.IntOutput { return v.HealthCheckUnhealthNum }).(pulumi.IntOutput)
}

// Indicate to apply HTTP2.0 protocol or not.
func (o ListenerRuleOutput) Http2Switch() pulumi.BoolOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.BoolOutput { return v.Http2Switch }).(pulumi.BoolOutput)
}

// ID of CLB listener.
func (o ListenerRuleOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// ID of this CLB listener rule.
func (o ListenerRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
func (o ListenerRuleOutput) Scheduler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringPtrOutput { return v.Scheduler }).(pulumi.StringPtrOutput)
}

// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `Clb.ListenerRule`.
func (o ListenerRuleOutput) SessionExpireTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.IntPtrOutput { return v.SessionExpireTime }).(pulumi.IntPtrOutput)
}

// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
func (o ListenerRuleOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringPtrOutput { return v.TargetType }).(pulumi.StringPtrOutput)
}

// Url of the listener rule.
func (o ListenerRuleOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ListenerRule) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ListenerRuleArrayOutput struct{ *pulumi.OutputState }

func (ListenerRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListenerRule)(nil)).Elem()
}

func (o ListenerRuleArrayOutput) ToListenerRuleArrayOutput() ListenerRuleArrayOutput {
	return o
}

func (o ListenerRuleArrayOutput) ToListenerRuleArrayOutputWithContext(ctx context.Context) ListenerRuleArrayOutput {
	return o
}

func (o ListenerRuleArrayOutput) Index(i pulumi.IntInput) ListenerRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ListenerRule {
		return vs[0].([]*ListenerRule)[vs[1].(int)]
	}).(ListenerRuleOutput)
}

type ListenerRuleMapOutput struct{ *pulumi.OutputState }

func (ListenerRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListenerRule)(nil)).Elem()
}

func (o ListenerRuleMapOutput) ToListenerRuleMapOutput() ListenerRuleMapOutput {
	return o
}

func (o ListenerRuleMapOutput) ToListenerRuleMapOutputWithContext(ctx context.Context) ListenerRuleMapOutput {
	return o
}

func (o ListenerRuleMapOutput) MapIndex(k pulumi.StringInput) ListenerRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ListenerRule {
		return vs[0].(map[string]*ListenerRule)[vs[1].(string)]
	}).(ListenerRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRuleInput)(nil)).Elem(), &ListenerRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRuleArrayInput)(nil)).Elem(), ListenerRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRuleMapInput)(nil)).Elem(), ListenerRuleMap{})
	pulumi.RegisterOutputType(ListenerRuleOutput{})
	pulumi.RegisterOutputType(ListenerRuleArrayOutput{})
	pulumi.RegisterOutputType(ListenerRuleMapOutput{})
}
