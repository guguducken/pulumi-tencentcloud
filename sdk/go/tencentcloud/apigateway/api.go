// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create API of API gateway.
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
// 		service, err := ApiGateway.NewService(ctx, "service", &ApiGateway.ServiceArgs{
// 			ServiceName: pulumi.String("ck"),
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
// 		_, err = ApiGateway.NewApi(ctx, "api", &ApiGateway.ApiArgs{
// 			ServiceId:           service.ID(),
// 			ApiName:             pulumi.String("hello"),
// 			ApiDesc:             pulumi.String("my hello api"),
// 			AuthType:            pulumi.String("NONE"),
// 			Protocol:            pulumi.String("HTTP"),
// 			EnableCors:          pulumi.Bool(true),
// 			RequestConfigPath:   pulumi.String("/user/info"),
// 			RequestConfigMethod: pulumi.String("GET"),
// 			RequestParameters: apigateway.ApiRequestParameterArray{
// 				&apigateway.ApiRequestParameterArgs{
// 					Name:         pulumi.String("name"),
// 					Position:     pulumi.String("QUERY"),
// 					Type:         pulumi.String("string"),
// 					Desc:         pulumi.String("who are you?"),
// 					DefaultValue: pulumi.String("tom"),
// 					Required:     pulumi.Bool(true),
// 				},
// 			},
// 			ServiceConfigType:      pulumi.String("HTTP"),
// 			ServiceConfigTimeout:   pulumi.Int(15),
// 			ServiceConfigUrl:       pulumi.String("http://www.qq.com"),
// 			ServiceConfigPath:      pulumi.String("/user"),
// 			ServiceConfigMethod:    pulumi.String("GET"),
// 			ResponseType:           pulumi.String("HTML"),
// 			ResponseSuccessExample: pulumi.String("success"),
// 			ResponseFailExample:    pulumi.String("fail"),
// 			ResponseErrorCodes: apigateway.ApiResponseErrorCodeArray{
// 				&apigateway.ApiResponseErrorCodeArgs{
// 					Code:          pulumi.Int(100),
// 					Msg:           pulumi.String("system error"),
// 					Desc:          pulumi.String("system error code"),
// 					ConvertedCode: -100,
// 					NeedConvert:   pulumi.Bool(true),
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
type Api struct {
	pulumi.CustomResourceState

	// Custom API description.
	ApiDesc pulumi.StringPtrOutput `pulumi:"apiDesc"`
	// Custom API name.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
	AuthType pulumi.StringPtrOutput `pulumi:"authType"`
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether to enable CORS. Default value: `true`.
	EnableCors pulumi.BoolPtrOutput `pulumi:"enableCors"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	PreLimit pulumi.IntOutput `pulumi:"preLimit"`
	// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	ReleaseLimit pulumi.IntOutput `pulumi:"releaseLimit"`
	// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
	RequestConfigMethod pulumi.StringPtrOutput `pulumi:"requestConfigMethod"`
	// Request frontend path configuration. Like `/user/getinfo`.
	RequestConfigPath pulumi.StringOutput `pulumi:"requestConfigPath"`
	// Frontend request parameters.
	RequestParameters ApiRequestParameterArrayOutput `pulumi:"requestParameters"`
	// Custom error code configuration. Must keep at least one after set.
	ResponseErrorCodes ApiResponseErrorCodeArrayOutput `pulumi:"responseErrorCodes"`
	// Response failure sample of custom response configuration.
	ResponseFailExample pulumi.StringPtrOutput `pulumi:"responseFailExample"`
	// Successful response sample of custom response configuration.
	ResponseSuccessExample pulumi.StringPtrOutput `pulumi:"responseSuccessExample"`
	// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
	ResponseType pulumi.StringOutput `pulumi:"responseType"`
	// API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
	ServiceConfigMethod pulumi.StringPtrOutput `pulumi:"serviceConfigMethod"`
	// Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
	ServiceConfigMockReturnMessage pulumi.StringPtrOutput `pulumi:"serviceConfigMockReturnMessage"`
	// API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
	ServiceConfigPath pulumi.StringPtrOutput `pulumi:"serviceConfigPath"`
	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	ServiceConfigProduct pulumi.StringPtrOutput `pulumi:"serviceConfigProduct"`
	// SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionName pulumi.StringPtrOutput `pulumi:"serviceConfigScfFunctionName"`
	// SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionNamespace pulumi.StringPtrOutput `pulumi:"serviceConfigScfFunctionNamespace"`
	// SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionQualifier pulumi.StringPtrOutput `pulumi:"serviceConfigScfFunctionQualifier"`
	// API backend service timeout period in seconds. Default value: `5`.
	ServiceConfigTimeout pulumi.IntPtrOutput `pulumi:"serviceConfigTimeout"`
	// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
	ServiceConfigType pulumi.StringPtrOutput `pulumi:"serviceConfigType"`
	// API backend service url. This parameter is required when `serviceConfigType` is `HTTP`.
	ServiceConfigUrl pulumi.StringPtrOutput `pulumi:"serviceConfigUrl"`
	// Unique VPC ID.
	ServiceConfigVpcId pulumi.StringPtrOutput `pulumi:"serviceConfigVpcId"`
	// Which service this API belongs. Refer to resource `ApiGateway.Service`.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	TestLimit pulumi.IntOutput `pulumi:"testLimit"`
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiName == nil {
		return nil, errors.New("invalid value for required argument 'ApiName'")
	}
	if args.RequestConfigPath == nil {
		return nil, errors.New("invalid value for required argument 'RequestConfigPath'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("tencentcloud:ApiGateway/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("tencentcloud:ApiGateway/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// Custom API description.
	ApiDesc *string `pulumi:"apiDesc"`
	// Custom API name.
	ApiName *string `pulumi:"apiName"`
	// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
	AuthType *string `pulumi:"authType"`
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime *string `pulumi:"createTime"`
	// Whether to enable CORS. Default value: `true`.
	EnableCors *bool `pulumi:"enableCors"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	PreLimit *int `pulumi:"preLimit"`
	// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	ReleaseLimit *int `pulumi:"releaseLimit"`
	// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
	RequestConfigMethod *string `pulumi:"requestConfigMethod"`
	// Request frontend path configuration. Like `/user/getinfo`.
	RequestConfigPath *string `pulumi:"requestConfigPath"`
	// Frontend request parameters.
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// Custom error code configuration. Must keep at least one after set.
	ResponseErrorCodes []ApiResponseErrorCode `pulumi:"responseErrorCodes"`
	// Response failure sample of custom response configuration.
	ResponseFailExample *string `pulumi:"responseFailExample"`
	// Successful response sample of custom response configuration.
	ResponseSuccessExample *string `pulumi:"responseSuccessExample"`
	// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
	ResponseType *string `pulumi:"responseType"`
	// API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
	ServiceConfigMethod *string `pulumi:"serviceConfigMethod"`
	// Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
	ServiceConfigMockReturnMessage *string `pulumi:"serviceConfigMockReturnMessage"`
	// API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
	ServiceConfigPath *string `pulumi:"serviceConfigPath"`
	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	ServiceConfigProduct *string `pulumi:"serviceConfigProduct"`
	// SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionName *string `pulumi:"serviceConfigScfFunctionName"`
	// SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionNamespace *string `pulumi:"serviceConfigScfFunctionNamespace"`
	// SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionQualifier *string `pulumi:"serviceConfigScfFunctionQualifier"`
	// API backend service timeout period in seconds. Default value: `5`.
	ServiceConfigTimeout *int `pulumi:"serviceConfigTimeout"`
	// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
	ServiceConfigType *string `pulumi:"serviceConfigType"`
	// API backend service url. This parameter is required when `serviceConfigType` is `HTTP`.
	ServiceConfigUrl *string `pulumi:"serviceConfigUrl"`
	// Unique VPC ID.
	ServiceConfigVpcId *string `pulumi:"serviceConfigVpcId"`
	// Which service this API belongs. Refer to resource `ApiGateway.Service`.
	ServiceId *string `pulumi:"serviceId"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	TestLimit *int `pulumi:"testLimit"`
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	UpdateTime *string `pulumi:"updateTime"`
}

type ApiState struct {
	// Custom API description.
	ApiDesc pulumi.StringPtrInput
	// Custom API name.
	ApiName pulumi.StringPtrInput
	// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
	AuthType pulumi.StringPtrInput
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime pulumi.StringPtrInput
	// Whether to enable CORS. Default value: `true`.
	EnableCors pulumi.BoolPtrInput
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	PreLimit pulumi.IntPtrInput
	// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
	Protocol pulumi.StringPtrInput
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	ReleaseLimit pulumi.IntPtrInput
	// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
	RequestConfigMethod pulumi.StringPtrInput
	// Request frontend path configuration. Like `/user/getinfo`.
	RequestConfigPath pulumi.StringPtrInput
	// Frontend request parameters.
	RequestParameters ApiRequestParameterArrayInput
	// Custom error code configuration. Must keep at least one after set.
	ResponseErrorCodes ApiResponseErrorCodeArrayInput
	// Response failure sample of custom response configuration.
	ResponseFailExample pulumi.StringPtrInput
	// Successful response sample of custom response configuration.
	ResponseSuccessExample pulumi.StringPtrInput
	// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
	ResponseType pulumi.StringPtrInput
	// API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
	ServiceConfigMethod pulumi.StringPtrInput
	// Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
	ServiceConfigMockReturnMessage pulumi.StringPtrInput
	// API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
	ServiceConfigPath pulumi.StringPtrInput
	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	ServiceConfigProduct pulumi.StringPtrInput
	// SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionName pulumi.StringPtrInput
	// SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionNamespace pulumi.StringPtrInput
	// SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionQualifier pulumi.StringPtrInput
	// API backend service timeout period in seconds. Default value: `5`.
	ServiceConfigTimeout pulumi.IntPtrInput
	// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
	ServiceConfigType pulumi.StringPtrInput
	// API backend service url. This parameter is required when `serviceConfigType` is `HTTP`.
	ServiceConfigUrl pulumi.StringPtrInput
	// Unique VPC ID.
	ServiceConfigVpcId pulumi.StringPtrInput
	// Which service this API belongs. Refer to resource `ApiGateway.Service`.
	ServiceId pulumi.StringPtrInput
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	TestLimit pulumi.IntPtrInput
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	UpdateTime pulumi.StringPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// Custom API description.
	ApiDesc *string `pulumi:"apiDesc"`
	// Custom API name.
	ApiName string `pulumi:"apiName"`
	// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
	AuthType *string `pulumi:"authType"`
	// Whether to enable CORS. Default value: `true`.
	EnableCors *bool `pulumi:"enableCors"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	PreLimit *int `pulumi:"preLimit"`
	// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	ReleaseLimit *int `pulumi:"releaseLimit"`
	// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
	RequestConfigMethod *string `pulumi:"requestConfigMethod"`
	// Request frontend path configuration. Like `/user/getinfo`.
	RequestConfigPath string `pulumi:"requestConfigPath"`
	// Frontend request parameters.
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// Custom error code configuration. Must keep at least one after set.
	ResponseErrorCodes []ApiResponseErrorCode `pulumi:"responseErrorCodes"`
	// Response failure sample of custom response configuration.
	ResponseFailExample *string `pulumi:"responseFailExample"`
	// Successful response sample of custom response configuration.
	ResponseSuccessExample *string `pulumi:"responseSuccessExample"`
	// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
	ResponseType *string `pulumi:"responseType"`
	// API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
	ServiceConfigMethod *string `pulumi:"serviceConfigMethod"`
	// Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
	ServiceConfigMockReturnMessage *string `pulumi:"serviceConfigMockReturnMessage"`
	// API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
	ServiceConfigPath *string `pulumi:"serviceConfigPath"`
	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	ServiceConfigProduct *string `pulumi:"serviceConfigProduct"`
	// SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionName *string `pulumi:"serviceConfigScfFunctionName"`
	// SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionNamespace *string `pulumi:"serviceConfigScfFunctionNamespace"`
	// SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionQualifier *string `pulumi:"serviceConfigScfFunctionQualifier"`
	// API backend service timeout period in seconds. Default value: `5`.
	ServiceConfigTimeout *int `pulumi:"serviceConfigTimeout"`
	// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
	ServiceConfigType *string `pulumi:"serviceConfigType"`
	// API backend service url. This parameter is required when `serviceConfigType` is `HTTP`.
	ServiceConfigUrl *string `pulumi:"serviceConfigUrl"`
	// Unique VPC ID.
	ServiceConfigVpcId *string `pulumi:"serviceConfigVpcId"`
	// Which service this API belongs. Refer to resource `ApiGateway.Service`.
	ServiceId string `pulumi:"serviceId"`
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	TestLimit *int `pulumi:"testLimit"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// Custom API description.
	ApiDesc pulumi.StringPtrInput
	// Custom API name.
	ApiName pulumi.StringInput
	// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
	AuthType pulumi.StringPtrInput
	// Whether to enable CORS. Default value: `true`.
	EnableCors pulumi.BoolPtrInput
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	PreLimit pulumi.IntPtrInput
	// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
	Protocol pulumi.StringPtrInput
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	ReleaseLimit pulumi.IntPtrInput
	// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
	RequestConfigMethod pulumi.StringPtrInput
	// Request frontend path configuration. Like `/user/getinfo`.
	RequestConfigPath pulumi.StringInput
	// Frontend request parameters.
	RequestParameters ApiRequestParameterArrayInput
	// Custom error code configuration. Must keep at least one after set.
	ResponseErrorCodes ApiResponseErrorCodeArrayInput
	// Response failure sample of custom response configuration.
	ResponseFailExample pulumi.StringPtrInput
	// Successful response sample of custom response configuration.
	ResponseSuccessExample pulumi.StringPtrInput
	// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
	ResponseType pulumi.StringPtrInput
	// API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
	ServiceConfigMethod pulumi.StringPtrInput
	// Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
	ServiceConfigMockReturnMessage pulumi.StringPtrInput
	// API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
	ServiceConfigPath pulumi.StringPtrInput
	// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
	ServiceConfigProduct pulumi.StringPtrInput
	// SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionName pulumi.StringPtrInput
	// SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionNamespace pulumi.StringPtrInput
	// SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
	ServiceConfigScfFunctionQualifier pulumi.StringPtrInput
	// API backend service timeout period in seconds. Default value: `5`.
	ServiceConfigTimeout pulumi.IntPtrInput
	// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
	ServiceConfigType pulumi.StringPtrInput
	// API backend service url. This parameter is required when `serviceConfigType` is `HTTP`.
	ServiceConfigUrl pulumi.StringPtrInput
	// Unique VPC ID.
	ServiceConfigVpcId pulumi.StringPtrInput
	// Which service this API belongs. Refer to resource `ApiGateway.Service`.
	ServiceId pulumi.StringInput
	// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
	TestLimit pulumi.IntPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

// ApiArrayInput is an input type that accepts ApiArray and ApiArrayOutput values.
// You can construct a concrete instance of `ApiArrayInput` via:
//
//          ApiArray{ ApiArgs{...} }
type ApiArrayInput interface {
	pulumi.Input

	ToApiArrayOutput() ApiArrayOutput
	ToApiArrayOutputWithContext(context.Context) ApiArrayOutput
}

type ApiArray []ApiInput

func (ApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (i ApiArray) ToApiArrayOutput() ApiArrayOutput {
	return i.ToApiArrayOutputWithContext(context.Background())
}

func (i ApiArray) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiArrayOutput)
}

// ApiMapInput is an input type that accepts ApiMap and ApiMapOutput values.
// You can construct a concrete instance of `ApiMapInput` via:
//
//          ApiMap{ "key": ApiArgs{...} }
type ApiMapInput interface {
	pulumi.Input

	ToApiMapOutput() ApiMapOutput
	ToApiMapOutputWithContext(context.Context) ApiMapOutput
}

type ApiMap map[string]ApiInput

func (ApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (i ApiMap) ToApiMapOutput() ApiMapOutput {
	return i.ToApiMapOutputWithContext(context.Background())
}

func (i ApiMap) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMapOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// Custom API description.
func (o ApiOutput) ApiDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiDesc }).(pulumi.StringPtrOutput)
}

// Custom API name.
func (o ApiOutput) ApiName() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ApiName }).(pulumi.StringOutput)
}

// API authentication type. Valid values: `SECRET` (key pair authentication),`NONE` (no authentication). Default value: `NONE`.
func (o ApiOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.AuthType }).(pulumi.StringPtrOutput)
}

// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
func (o ApiOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether to enable CORS. Default value: `true`.
func (o ApiOutput) EnableCors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.EnableCors }).(pulumi.BoolPtrOutput)
}

// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
func (o ApiOutput) PreLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Api) pulumi.IntOutput { return v.PreLimit }).(pulumi.IntOutput)
}

// API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
func (o ApiOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
func (o ApiOutput) ReleaseLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Api) pulumi.IntOutput { return v.ReleaseLimit }).(pulumi.IntOutput)
}

// Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
func (o ApiOutput) RequestConfigMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.RequestConfigMethod }).(pulumi.StringPtrOutput)
}

// Request frontend path configuration. Like `/user/getinfo`.
func (o ApiOutput) RequestConfigPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.RequestConfigPath }).(pulumi.StringOutput)
}

// Frontend request parameters.
func (o ApiOutput) RequestParameters() ApiRequestParameterArrayOutput {
	return o.ApplyT(func(v *Api) ApiRequestParameterArrayOutput { return v.RequestParameters }).(ApiRequestParameterArrayOutput)
}

// Custom error code configuration. Must keep at least one after set.
func (o ApiOutput) ResponseErrorCodes() ApiResponseErrorCodeArrayOutput {
	return o.ApplyT(func(v *Api) ApiResponseErrorCodeArrayOutput { return v.ResponseErrorCodes }).(ApiResponseErrorCodeArrayOutput)
}

// Response failure sample of custom response configuration.
func (o ApiOutput) ResponseFailExample() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ResponseFailExample }).(pulumi.StringPtrOutput)
}

// Successful response sample of custom response configuration.
func (o ApiOutput) ResponseSuccessExample() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ResponseSuccessExample }).(pulumi.StringPtrOutput)
}

// Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
func (o ApiOutput) ResponseType() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ResponseType }).(pulumi.StringOutput)
}

// API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
func (o ApiOutput) ServiceConfigMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigMethod }).(pulumi.StringPtrOutput)
}

// Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
func (o ApiOutput) ServiceConfigMockReturnMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigMockReturnMessage }).(pulumi.StringPtrOutput)
}

// API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
func (o ApiOutput) ServiceConfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigPath }).(pulumi.StringPtrOutput)
}

// Backend type. This parameter takes effect when VPC is enabled. Currently, only `clb` is supported.
func (o ApiOutput) ServiceConfigProduct() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigProduct }).(pulumi.StringPtrOutput)
}

// SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
func (o ApiOutput) ServiceConfigScfFunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigScfFunctionName }).(pulumi.StringPtrOutput)
}

// SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
func (o ApiOutput) ServiceConfigScfFunctionNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigScfFunctionNamespace }).(pulumi.StringPtrOutput)
}

// SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
func (o ApiOutput) ServiceConfigScfFunctionQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigScfFunctionQualifier }).(pulumi.StringPtrOutput)
}

// API backend service timeout period in seconds. Default value: `5`.
func (o ApiOutput) ServiceConfigTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.IntPtrOutput { return v.ServiceConfigTimeout }).(pulumi.IntPtrOutput)
}

// API backend service type. Valid values: `WEBSOCKET`, `HTTP`, `SCF`, `MOCK`. Default value: `HTTP`.
func (o ApiOutput) ServiceConfigType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigType }).(pulumi.StringPtrOutput)
}

// API backend service url. This parameter is required when `serviceConfigType` is `HTTP`.
func (o ApiOutput) ServiceConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigUrl }).(pulumi.StringPtrOutput)
}

// Unique VPC ID.
func (o ApiOutput) ServiceConfigVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceConfigVpcId }).(pulumi.StringPtrOutput)
}

// Which service this API belongs. Refer to resource `ApiGateway.Service`.
func (o ApiOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// API QPS value. Enter a positive number to limit the API query rate per second `QPS`.
func (o ApiOutput) TestLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Api) pulumi.IntOutput { return v.TestLimit }).(pulumi.IntOutput)
}

// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
func (o ApiOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ApiArrayOutput struct{ *pulumi.OutputState }

func (ApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (o ApiArrayOutput) ToApiArrayOutput() ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) Index(i pulumi.IntInput) ApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Api {
		return vs[0].([]*Api)[vs[1].(int)]
	}).(ApiOutput)
}

type ApiMapOutput struct{ *pulumi.OutputState }

func (ApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (o ApiMapOutput) ToApiMapOutput() ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return o
}

func (o ApiMapOutput) MapIndex(k pulumi.StringInput) ApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Api {
		return vs[0].(map[string]*Api)[vs[1].(string)]
	}).(ApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiArrayInput)(nil)).Elem(), ApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiMapInput)(nil)).Elem(), ApiMap{})
	pulumi.RegisterOutputType(ApiOutput{})
	pulumi.RegisterOutputType(ApiArrayOutput{})
	pulumi.RegisterOutputType(ApiMapOutput{})
}
