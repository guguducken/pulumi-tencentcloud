// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to attach API gateway usage plan to service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		plan, err := ApiGateway.NewUsagePlan(ctx, "plan", &ApiGateway.UsagePlanArgs{
// 			UsagePlanName:       pulumi.String("my_plan"),
// 			UsagePlanDesc:       pulumi.String("nice plan"),
// 			MaxRequestNum:       pulumi.Int(100),
// 			MaxRequestNumPreSec: pulumi.Int(10),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		service, err := ApiGateway.NewService(ctx, "service", &ApiGateway.ServiceArgs{
// 			ServiceName: pulumi.String("niceservice"),
// 			Protocol:    pulumi.String("http&https"),
// 			ServiceDesc: pulumi.String("your nice service"),
// 			NetTypes: pulumi.StringArray{
// 				pulumi.String("INNER"),
// 				pulumi.String("OUTER"),
// 			},
// 			IpVersion: pulumi.String("IPv4"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		api, err := ApiGateway.NewApi(ctx, "api", &ApiGateway.ApiArgs{
// 			ServiceId:           service.ID(),
// 			ApiName:             pulumi.String("hello_update"),
// 			ApiDesc:             pulumi.String("my hello api update"),
// 			AuthType:            pulumi.String("SECRET"),
// 			Protocol:            pulumi.String("HTTP"),
// 			EnableCors:          pulumi.Bool(true),
// 			RequestConfigPath:   pulumi.String("/user/info"),
// 			RequestConfigMethod: pulumi.String("POST"),
// 			RequestParameters: apigateway.ApiRequestParameterArray{
// 				&apigateway.ApiRequestParameterArgs{
// 					Name:         pulumi.String("email"),
// 					Position:     pulumi.String("QUERY"),
// 					Type:         pulumi.String("string"),
// 					Desc:         pulumi.String("your email please?"),
// 					DefaultValue: pulumi.String("tom@qq.com"),
// 					Required:     pulumi.Bool(true),
// 				},
// 			},
// 			ServiceConfigType:      pulumi.String("HTTP"),
// 			ServiceConfigTimeout:   pulumi.Int(10),
// 			ServiceConfigUrl:       pulumi.String("http://www.tencent.com"),
// 			ServiceConfigPath:      pulumi.String("/user"),
// 			ServiceConfigMethod:    pulumi.String("POST"),
// 			ResponseType:           pulumi.String("XML"),
// 			ResponseSuccessExample: pulumi.String("<note>success</note>"),
// 			ResponseFailExample:    pulumi.String("<note>fail</note>"),
// 			ResponseErrorCodes: apigateway.ApiResponseErrorCodeArray{
// 				&apigateway.ApiResponseErrorCodeArgs{
// 					Code:          pulumi.Int(10),
// 					Msg:           pulumi.String("system error"),
// 					Desc:          pulumi.String("system error code"),
// 					ConvertedCode: -10,
// 					NeedConvert:   pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ApiGateway.NewUsagePlanAttachment(ctx, "attachService", &ApiGateway.UsagePlanAttachmentArgs{
// 			UsagePlanId: plan.ID(),
// 			ServiceId:   service.ID(),
// 			Environment: pulumi.String("release"),
// 			BindType:    pulumi.String("API"),
// 			ApiId:       api.ID(),
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
// API gateway usage plan attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment attach_service usagePlan-pe7fbdgn#service-kuqd6xqk#release#API#api-p8gtanvy
// ```
type UsagePlanAttachment struct {
	pulumi.CustomResourceState

	// ID of the API. This parameter will be required when `bindType` is `API`.
	ApiId pulumi.StringPtrOutput `pulumi:"apiId"`
	// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
	BindType pulumi.StringPtrOutput `pulumi:"bindType"`
	// The environment to be bound. Valid values: `test`, `prepub`, `release`.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// ID of the service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// ID of the usage plan.
	UsagePlanId pulumi.StringOutput `pulumi:"usagePlanId"`
}

// NewUsagePlanAttachment registers a new resource with the given unique name, arguments, and options.
func NewUsagePlanAttachment(ctx *pulumi.Context,
	name string, args *UsagePlanAttachmentArgs, opts ...pulumi.ResourceOption) (*UsagePlanAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.UsagePlanId == nil {
		return nil, errors.New("invalid value for required argument 'UsagePlanId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UsagePlanAttachment
	err := ctx.RegisterResource("tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsagePlanAttachment gets an existing UsagePlanAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsagePlanAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsagePlanAttachmentState, opts ...pulumi.ResourceOption) (*UsagePlanAttachment, error) {
	var resource UsagePlanAttachment
	err := ctx.ReadResource("tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UsagePlanAttachment resources.
type usagePlanAttachmentState struct {
	// ID of the API. This parameter will be required when `bindType` is `API`.
	ApiId *string `pulumi:"apiId"`
	// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
	BindType *string `pulumi:"bindType"`
	// The environment to be bound. Valid values: `test`, `prepub`, `release`.
	Environment *string `pulumi:"environment"`
	// ID of the service.
	ServiceId *string `pulumi:"serviceId"`
	// ID of the usage plan.
	UsagePlanId *string `pulumi:"usagePlanId"`
}

type UsagePlanAttachmentState struct {
	// ID of the API. This parameter will be required when `bindType` is `API`.
	ApiId pulumi.StringPtrInput
	// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
	BindType pulumi.StringPtrInput
	// The environment to be bound. Valid values: `test`, `prepub`, `release`.
	Environment pulumi.StringPtrInput
	// ID of the service.
	ServiceId pulumi.StringPtrInput
	// ID of the usage plan.
	UsagePlanId pulumi.StringPtrInput
}

func (UsagePlanAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanAttachmentState)(nil)).Elem()
}

type usagePlanAttachmentArgs struct {
	// ID of the API. This parameter will be required when `bindType` is `API`.
	ApiId *string `pulumi:"apiId"`
	// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
	BindType *string `pulumi:"bindType"`
	// The environment to be bound. Valid values: `test`, `prepub`, `release`.
	Environment string `pulumi:"environment"`
	// ID of the service.
	ServiceId string `pulumi:"serviceId"`
	// ID of the usage plan.
	UsagePlanId string `pulumi:"usagePlanId"`
}

// The set of arguments for constructing a UsagePlanAttachment resource.
type UsagePlanAttachmentArgs struct {
	// ID of the API. This parameter will be required when `bindType` is `API`.
	ApiId pulumi.StringPtrInput
	// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
	BindType pulumi.StringPtrInput
	// The environment to be bound. Valid values: `test`, `prepub`, `release`.
	Environment pulumi.StringInput
	// ID of the service.
	ServiceId pulumi.StringInput
	// ID of the usage plan.
	UsagePlanId pulumi.StringInput
}

func (UsagePlanAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usagePlanAttachmentArgs)(nil)).Elem()
}

type UsagePlanAttachmentInput interface {
	pulumi.Input

	ToUsagePlanAttachmentOutput() UsagePlanAttachmentOutput
	ToUsagePlanAttachmentOutputWithContext(ctx context.Context) UsagePlanAttachmentOutput
}

func (*UsagePlanAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlanAttachment)(nil)).Elem()
}

func (i *UsagePlanAttachment) ToUsagePlanAttachmentOutput() UsagePlanAttachmentOutput {
	return i.ToUsagePlanAttachmentOutputWithContext(context.Background())
}

func (i *UsagePlanAttachment) ToUsagePlanAttachmentOutputWithContext(ctx context.Context) UsagePlanAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanAttachmentOutput)
}

// UsagePlanAttachmentArrayInput is an input type that accepts UsagePlanAttachmentArray and UsagePlanAttachmentArrayOutput values.
// You can construct a concrete instance of `UsagePlanAttachmentArrayInput` via:
//
//          UsagePlanAttachmentArray{ UsagePlanAttachmentArgs{...} }
type UsagePlanAttachmentArrayInput interface {
	pulumi.Input

	ToUsagePlanAttachmentArrayOutput() UsagePlanAttachmentArrayOutput
	ToUsagePlanAttachmentArrayOutputWithContext(context.Context) UsagePlanAttachmentArrayOutput
}

type UsagePlanAttachmentArray []UsagePlanAttachmentInput

func (UsagePlanAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UsagePlanAttachment)(nil)).Elem()
}

func (i UsagePlanAttachmentArray) ToUsagePlanAttachmentArrayOutput() UsagePlanAttachmentArrayOutput {
	return i.ToUsagePlanAttachmentArrayOutputWithContext(context.Background())
}

func (i UsagePlanAttachmentArray) ToUsagePlanAttachmentArrayOutputWithContext(ctx context.Context) UsagePlanAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanAttachmentArrayOutput)
}

// UsagePlanAttachmentMapInput is an input type that accepts UsagePlanAttachmentMap and UsagePlanAttachmentMapOutput values.
// You can construct a concrete instance of `UsagePlanAttachmentMapInput` via:
//
//          UsagePlanAttachmentMap{ "key": UsagePlanAttachmentArgs{...} }
type UsagePlanAttachmentMapInput interface {
	pulumi.Input

	ToUsagePlanAttachmentMapOutput() UsagePlanAttachmentMapOutput
	ToUsagePlanAttachmentMapOutputWithContext(context.Context) UsagePlanAttachmentMapOutput
}

type UsagePlanAttachmentMap map[string]UsagePlanAttachmentInput

func (UsagePlanAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UsagePlanAttachment)(nil)).Elem()
}

func (i UsagePlanAttachmentMap) ToUsagePlanAttachmentMapOutput() UsagePlanAttachmentMapOutput {
	return i.ToUsagePlanAttachmentMapOutputWithContext(context.Background())
}

func (i UsagePlanAttachmentMap) ToUsagePlanAttachmentMapOutputWithContext(ctx context.Context) UsagePlanAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsagePlanAttachmentMapOutput)
}

type UsagePlanAttachmentOutput struct{ *pulumi.OutputState }

func (UsagePlanAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsagePlanAttachment)(nil)).Elem()
}

func (o UsagePlanAttachmentOutput) ToUsagePlanAttachmentOutput() UsagePlanAttachmentOutput {
	return o
}

func (o UsagePlanAttachmentOutput) ToUsagePlanAttachmentOutputWithContext(ctx context.Context) UsagePlanAttachmentOutput {
	return o
}

// ID of the API. This parameter will be required when `bindType` is `API`.
func (o UsagePlanAttachmentOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsagePlanAttachment) pulumi.StringPtrOutput { return v.ApiId }).(pulumi.StringPtrOutput)
}

// Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
func (o UsagePlanAttachmentOutput) BindType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsagePlanAttachment) pulumi.StringPtrOutput { return v.BindType }).(pulumi.StringPtrOutput)
}

// The environment to be bound. Valid values: `test`, `prepub`, `release`.
func (o UsagePlanAttachmentOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanAttachment) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// ID of the service.
func (o UsagePlanAttachmentOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanAttachment) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// ID of the usage plan.
func (o UsagePlanAttachmentOutput) UsagePlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *UsagePlanAttachment) pulumi.StringOutput { return v.UsagePlanId }).(pulumi.StringOutput)
}

type UsagePlanAttachmentArrayOutput struct{ *pulumi.OutputState }

func (UsagePlanAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UsagePlanAttachment)(nil)).Elem()
}

func (o UsagePlanAttachmentArrayOutput) ToUsagePlanAttachmentArrayOutput() UsagePlanAttachmentArrayOutput {
	return o
}

func (o UsagePlanAttachmentArrayOutput) ToUsagePlanAttachmentArrayOutputWithContext(ctx context.Context) UsagePlanAttachmentArrayOutput {
	return o
}

func (o UsagePlanAttachmentArrayOutput) Index(i pulumi.IntInput) UsagePlanAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UsagePlanAttachment {
		return vs[0].([]*UsagePlanAttachment)[vs[1].(int)]
	}).(UsagePlanAttachmentOutput)
}

type UsagePlanAttachmentMapOutput struct{ *pulumi.OutputState }

func (UsagePlanAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UsagePlanAttachment)(nil)).Elem()
}

func (o UsagePlanAttachmentMapOutput) ToUsagePlanAttachmentMapOutput() UsagePlanAttachmentMapOutput {
	return o
}

func (o UsagePlanAttachmentMapOutput) ToUsagePlanAttachmentMapOutputWithContext(ctx context.Context) UsagePlanAttachmentMapOutput {
	return o
}

func (o UsagePlanAttachmentMapOutput) MapIndex(k pulumi.StringInput) UsagePlanAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UsagePlanAttachment {
		return vs[0].(map[string]*UsagePlanAttachment)[vs[1].(string)]
	}).(UsagePlanAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanAttachmentInput)(nil)).Elem(), &UsagePlanAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanAttachmentArrayInput)(nil)).Elem(), UsagePlanAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsagePlanAttachmentMapInput)(nil)).Elem(), UsagePlanAttachmentMap{})
	pulumi.RegisterOutputType(UsagePlanAttachmentOutput{})
	pulumi.RegisterOutputType(UsagePlanAttachmentArrayOutput{})
	pulumi.RegisterOutputType(UsagePlanAttachmentMapOutput{})
}
