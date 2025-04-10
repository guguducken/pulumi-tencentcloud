// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mariadb flow
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mariadb.GetFlow(ctx, &mariadb.GetFlowArgs{
// 			FlowId: 1307,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetFlow(ctx *pulumi.Context, args *GetFlowArgs, opts ...pulumi.InvokeOption) (*GetFlowResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFlowResult
	err := ctx.Invoke("tencentcloud:Mariadb/getFlow:getFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlow.
type GetFlowArgs struct {
	// Flow ID returned by async request API.
	FlowId int `pulumi:"flowId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getFlow.
type GetFlowResult struct {
	FlowId int `pulumi:"flowId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Flow status. 0: succeeded, 1: failed, 2: running.
	Status int `pulumi:"status"`
}

func GetFlowOutput(ctx *pulumi.Context, args GetFlowOutputArgs, opts ...pulumi.InvokeOption) GetFlowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFlowResult, error) {
			args := v.(GetFlowArgs)
			r, err := GetFlow(ctx, &args, opts...)
			var s GetFlowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFlowResultOutput)
}

// A collection of arguments for invoking getFlow.
type GetFlowOutputArgs struct {
	// Flow ID returned by async request API.
	FlowId pulumi.IntInput `pulumi:"flowId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetFlowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowArgs)(nil)).Elem()
}

// A collection of values returned by getFlow.
type GetFlowResultOutput struct{ *pulumi.OutputState }

func (GetFlowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowResult)(nil)).Elem()
}

func (o GetFlowResultOutput) ToGetFlowResultOutput() GetFlowResultOutput {
	return o
}

func (o GetFlowResultOutput) ToGetFlowResultOutputWithContext(ctx context.Context) GetFlowResultOutput {
	return o
}

func (o GetFlowResultOutput) FlowId() pulumi.IntOutput {
	return o.ApplyT(func(v GetFlowResult) int { return v.FlowId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFlowResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Flow status. 0: succeeded, 1: failed, 2: running.
func (o GetFlowResultOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetFlowResult) int { return v.Status }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlowResultOutput{})
}
