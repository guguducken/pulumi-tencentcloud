// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of sqlserver slowlogs
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Sqlserver.GetSlowlogs(ctx, &sqlserver.GetSlowlogsArgs{
// 			EndTime:    "2023-05-18 00:00:00",
// 			InstanceId: "mssql-qelbzgwf",
// 			StartTime:  "2020-05-01 00:00:00",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSlowlogs(ctx *pulumi.Context, args *GetSlowlogsArgs, opts ...pulumi.InvokeOption) (*GetSlowlogsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSlowlogsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getSlowlogs:getSlowlogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowlogs.
type GetSlowlogsArgs struct {
	// Query end time.
	EndTime string `pulumi:"endTime"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Query start time.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getSlowlogs.
type GetSlowlogsResult struct {
	// File generation end time.
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Information list of slow query logs.
	Slowlogs []GetSlowlogsSlowlog `pulumi:"slowlogs"`
	// File generation start time.
	StartTime string `pulumi:"startTime"`
}

func GetSlowlogsOutput(ctx *pulumi.Context, args GetSlowlogsOutputArgs, opts ...pulumi.InvokeOption) GetSlowlogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowlogsResult, error) {
			args := v.(GetSlowlogsArgs)
			r, err := GetSlowlogs(ctx, &args, opts...)
			var s GetSlowlogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowlogsResultOutput)
}

// A collection of arguments for invoking getSlowlogs.
type GetSlowlogsOutputArgs struct {
	// Query end time.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Query start time.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetSlowlogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowlogsArgs)(nil)).Elem()
}

// A collection of values returned by getSlowlogs.
type GetSlowlogsResultOutput struct{ *pulumi.OutputState }

func (GetSlowlogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowlogsResult)(nil)).Elem()
}

func (o GetSlowlogsResultOutput) ToGetSlowlogsResultOutput() GetSlowlogsResultOutput {
	return o
}

func (o GetSlowlogsResultOutput) ToGetSlowlogsResultOutputWithContext(ctx context.Context) GetSlowlogsResultOutput {
	return o
}

// File generation end time.
func (o GetSlowlogsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowlogsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowlogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowlogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowlogsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowlogsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSlowlogsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowlogsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Information list of slow query logs.
func (o GetSlowlogsResultOutput) Slowlogs() GetSlowlogsSlowlogArrayOutput {
	return o.ApplyT(func(v GetSlowlogsResult) []GetSlowlogsSlowlog { return v.Slowlogs }).(GetSlowlogsSlowlogArrayOutput)
}

// File generation start time.
func (o GetSlowlogsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowlogsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowlogsResultOutput{})
}
