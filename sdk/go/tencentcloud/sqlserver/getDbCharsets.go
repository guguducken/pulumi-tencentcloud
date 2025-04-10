// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of sqlserver datasourceDBCharsets
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
// 		_, err := Sqlserver.GetDbCharsets(ctx, &sqlserver.GetDbCharsetsArgs{
// 			InstanceId: "mssql-qelbzgwf",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDbCharsets(ctx *pulumi.Context, args *GetDbCharsetsArgs, opts ...pulumi.InvokeOption) (*GetDbCharsetsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDbCharsetsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getDbCharsets:getDbCharsets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDbCharsets.
type GetDbCharsetsArgs struct {
	// Instance ID in the format of mssql-j8kv137v.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDbCharsets.
type GetDbCharsetsResult struct {
	// Database character set list.
	DatabaseCharsets []string `pulumi:"databaseCharsets"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDbCharsetsOutput(ctx *pulumi.Context, args GetDbCharsetsOutputArgs, opts ...pulumi.InvokeOption) GetDbCharsetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDbCharsetsResult, error) {
			args := v.(GetDbCharsetsArgs)
			r, err := GetDbCharsets(ctx, &args, opts...)
			var s GetDbCharsetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDbCharsetsResultOutput)
}

// A collection of arguments for invoking getDbCharsets.
type GetDbCharsetsOutputArgs struct {
	// Instance ID in the format of mssql-j8kv137v.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDbCharsetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDbCharsetsArgs)(nil)).Elem()
}

// A collection of values returned by getDbCharsets.
type GetDbCharsetsResultOutput struct{ *pulumi.OutputState }

func (GetDbCharsetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDbCharsetsResult)(nil)).Elem()
}

func (o GetDbCharsetsResultOutput) ToGetDbCharsetsResultOutput() GetDbCharsetsResultOutput {
	return o
}

func (o GetDbCharsetsResultOutput) ToGetDbCharsetsResultOutputWithContext(ctx context.Context) GetDbCharsetsResultOutput {
	return o
}

// Database character set list.
func (o GetDbCharsetsResultOutput) DatabaseCharsets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDbCharsetsResult) []string { return v.DatabaseCharsets }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDbCharsetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbCharsetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDbCharsetsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbCharsetsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDbCharsetsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDbCharsetsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDbCharsetsResultOutput{})
}
