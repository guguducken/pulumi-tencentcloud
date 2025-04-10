// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of scf asyncEventManagement
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Scf.GetAsyncEventManagement(ctx, &scf.GetAsyncEventManagementArgs{
// 			FunctionName: "keep-1676351130",
// 			Namespace:    pulumi.StringRef("default"),
// 			Order:        pulumi.StringRef("ASC"),
// 			Orderby:      pulumi.StringRef("StartTime"),
// 			Qualifier:    pulumi.StringRef(fmt.Sprintf("%v%v", "$", "LATEST")),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAsyncEventManagement(ctx *pulumi.Context, args *GetAsyncEventManagementArgs, opts ...pulumi.InvokeOption) (*GetAsyncEventManagementResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAsyncEventManagementResult
	err := ctx.Invoke("tencentcloud:Scf/getAsyncEventManagement:getAsyncEventManagement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAsyncEventManagement.
type GetAsyncEventManagementArgs struct {
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Filter (event invocation request ID).
	InvokeRequestId *string `pulumi:"invokeRequestId"`
	// Filter (invocation type list), Values: CMQ, CKAFKA_TRIGGER, APIGW, COS, TRIGGER_TIMER, MPS_TRIGGER, CLS_TRIGGER, OTHERS.
	InvokeTypes []string `pulumi:"invokeTypes"`
	// Function namespace.
	Namespace *string `pulumi:"namespace"`
	// Valid values: ASC, DESC. Default value: DESC.
	Order *string `pulumi:"order"`
	// Valid values: StartTime, EndTime. Default value: StartTime.
	Orderby *string `pulumi:"orderby"`
	// Filter (function version).
	Qualifier *string `pulumi:"qualifier"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Filter (event status list), Values: RUNNING, FINISHED, ABORTED, FAILED.
	Statuses []string `pulumi:"statuses"`
}

// A collection of values returned by getAsyncEventManagement.
type GetAsyncEventManagementResult struct {
	// Async event list.
	EventLists   []GetAsyncEventManagementEventList `pulumi:"eventLists"`
	FunctionName string                             `pulumi:"functionName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Invocation request ID.
	InvokeRequestId *string `pulumi:"invokeRequestId"`
	// Invocation type.
	InvokeTypes []string `pulumi:"invokeTypes"`
	Namespace   *string  `pulumi:"namespace"`
	Order       *string  `pulumi:"order"`
	Orderby     *string  `pulumi:"orderby"`
	// Function version.
	Qualifier        *string `pulumi:"qualifier"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Event status. Values: `RUNNING`; `FINISHED` (invoked successfully); `ABORTED` (invocation ended); `FAILED` (invocation failed).
	Statuses []string `pulumi:"statuses"`
}

func GetAsyncEventManagementOutput(ctx *pulumi.Context, args GetAsyncEventManagementOutputArgs, opts ...pulumi.InvokeOption) GetAsyncEventManagementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAsyncEventManagementResult, error) {
			args := v.(GetAsyncEventManagementArgs)
			r, err := GetAsyncEventManagement(ctx, &args, opts...)
			var s GetAsyncEventManagementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAsyncEventManagementResultOutput)
}

// A collection of arguments for invoking getAsyncEventManagement.
type GetAsyncEventManagementOutputArgs struct {
	// Function name.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Filter (event invocation request ID).
	InvokeRequestId pulumi.StringPtrInput `pulumi:"invokeRequestId"`
	// Filter (invocation type list), Values: CMQ, CKAFKA_TRIGGER, APIGW, COS, TRIGGER_TIMER, MPS_TRIGGER, CLS_TRIGGER, OTHERS.
	InvokeTypes pulumi.StringArrayInput `pulumi:"invokeTypes"`
	// Function namespace.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Valid values: ASC, DESC. Default value: DESC.
	Order pulumi.StringPtrInput `pulumi:"order"`
	// Valid values: StartTime, EndTime. Default value: StartTime.
	Orderby pulumi.StringPtrInput `pulumi:"orderby"`
	// Filter (function version).
	Qualifier pulumi.StringPtrInput `pulumi:"qualifier"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Filter (event status list), Values: RUNNING, FINISHED, ABORTED, FAILED.
	Statuses pulumi.StringArrayInput `pulumi:"statuses"`
}

func (GetAsyncEventManagementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAsyncEventManagementArgs)(nil)).Elem()
}

// A collection of values returned by getAsyncEventManagement.
type GetAsyncEventManagementResultOutput struct{ *pulumi.OutputState }

func (GetAsyncEventManagementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAsyncEventManagementResult)(nil)).Elem()
}

func (o GetAsyncEventManagementResultOutput) ToGetAsyncEventManagementResultOutput() GetAsyncEventManagementResultOutput {
	return o
}

func (o GetAsyncEventManagementResultOutput) ToGetAsyncEventManagementResultOutputWithContext(ctx context.Context) GetAsyncEventManagementResultOutput {
	return o
}

// Async event list.
func (o GetAsyncEventManagementResultOutput) EventLists() GetAsyncEventManagementEventListArrayOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) []GetAsyncEventManagementEventList { return v.EventLists }).(GetAsyncEventManagementEventListArrayOutput)
}

func (o GetAsyncEventManagementResultOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) string { return v.FunctionName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAsyncEventManagementResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) string { return v.Id }).(pulumi.StringOutput)
}

// Invocation request ID.
func (o GetAsyncEventManagementResultOutput) InvokeRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) *string { return v.InvokeRequestId }).(pulumi.StringPtrOutput)
}

// Invocation type.
func (o GetAsyncEventManagementResultOutput) InvokeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) []string { return v.InvokeTypes }).(pulumi.StringArrayOutput)
}

func (o GetAsyncEventManagementResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetAsyncEventManagementResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetAsyncEventManagementResultOutput) Orderby() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) *string { return v.Orderby }).(pulumi.StringPtrOutput)
}

// Function version.
func (o GetAsyncEventManagementResultOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) *string { return v.Qualifier }).(pulumi.StringPtrOutput)
}

func (o GetAsyncEventManagementResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Event status. Values: `RUNNING`; `FINISHED` (invoked successfully); `ABORTED` (invocation ended); `FAILED` (invocation failed).
func (o GetAsyncEventManagementResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAsyncEventManagementResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAsyncEventManagementResultOutput{})
}
