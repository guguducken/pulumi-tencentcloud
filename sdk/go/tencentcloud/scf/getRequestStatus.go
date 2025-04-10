// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of scf requestStatus
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Scf.GetRequestStatus(ctx, &scf.GetRequestStatusArgs{
// 			FunctionName:      "keep-1676351130",
// 			FunctionRequestId: "9de9405a-e33a-498d-bb59-e80b7bed1191",
// 			Namespace:         pulumi.StringRef("default"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetRequestStatus(ctx *pulumi.Context, args *GetRequestStatusArgs, opts ...pulumi.InvokeOption) (*GetRequestStatusResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRequestStatusResult
	err := ctx.Invoke("tencentcloud:Scf/getRequestStatus:getRequestStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRequestStatus.
type GetRequestStatusArgs struct {
	// End time of the query. such as `2017-05-16 20:59:59`. If `StartTime` is not specified, `EndTime` defaults to the current time. If `StartTime` is specified, `EndTime` is required, and it need to be later than the `StartTime`.
	EndTime *string `pulumi:"endTime"`
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// ID of the request to be queried.
	FunctionRequestId string `pulumi:"functionRequestId"`
	// Function namespace.
	Namespace *string `pulumi:"namespace"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start time of the query, for example `2017-05-16 20:00:00`. If it's left empty, it defaults to 15 minutes before the current time.
	StartTime *string `pulumi:"startTime"`
}

// A collection of values returned by getRequestStatus.
type GetRequestStatusResult struct {
	// Details of the function running statusNote: this field may return `null`, indicating that no valid values can be obtained.
	Datas   []GetRequestStatusData `pulumi:"datas"`
	EndTime *string                `pulumi:"endTime"`
	// Function name.
	FunctionName      string `pulumi:"functionName"`
	FunctionRequestId string `pulumi:"functionRequestId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Namespace        *string `pulumi:"namespace"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Request start time.
	StartTime *string `pulumi:"startTime"`
}

func GetRequestStatusOutput(ctx *pulumi.Context, args GetRequestStatusOutputArgs, opts ...pulumi.InvokeOption) GetRequestStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRequestStatusResult, error) {
			args := v.(GetRequestStatusArgs)
			r, err := GetRequestStatus(ctx, &args, opts...)
			var s GetRequestStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRequestStatusResultOutput)
}

// A collection of arguments for invoking getRequestStatus.
type GetRequestStatusOutputArgs struct {
	// End time of the query. such as `2017-05-16 20:59:59`. If `StartTime` is not specified, `EndTime` defaults to the current time. If `StartTime` is specified, `EndTime` is required, and it need to be later than the `StartTime`.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// Function name.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// ID of the request to be queried.
	FunctionRequestId pulumi.StringInput `pulumi:"functionRequestId"`
	// Function namespace.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start time of the query, for example `2017-05-16 20:00:00`. If it's left empty, it defaults to 15 minutes before the current time.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (GetRequestStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRequestStatusArgs)(nil)).Elem()
}

// A collection of values returned by getRequestStatus.
type GetRequestStatusResultOutput struct{ *pulumi.OutputState }

func (GetRequestStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRequestStatusResult)(nil)).Elem()
}

func (o GetRequestStatusResultOutput) ToGetRequestStatusResultOutput() GetRequestStatusResultOutput {
	return o
}

func (o GetRequestStatusResultOutput) ToGetRequestStatusResultOutputWithContext(ctx context.Context) GetRequestStatusResultOutput {
	return o
}

// Details of the function running statusNote: this field may return `null`, indicating that no valid values can be obtained.
func (o GetRequestStatusResultOutput) Datas() GetRequestStatusDataArrayOutput {
	return o.ApplyT(func(v GetRequestStatusResult) []GetRequestStatusData { return v.Datas }).(GetRequestStatusDataArrayOutput)
}

func (o GetRequestStatusResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRequestStatusResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Function name.
func (o GetRequestStatusResultOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRequestStatusResult) string { return v.FunctionName }).(pulumi.StringOutput)
}

func (o GetRequestStatusResultOutput) FunctionRequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRequestStatusResult) string { return v.FunctionRequestId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRequestStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRequestStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRequestStatusResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRequestStatusResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetRequestStatusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRequestStatusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Request start time.
func (o GetRequestStatusResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRequestStatusResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRequestStatusResultOutput{})
}
