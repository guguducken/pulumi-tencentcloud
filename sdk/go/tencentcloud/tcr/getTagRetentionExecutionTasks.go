// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tcr tagRetentionExecutionTasks
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tcr.GetTagRetentionExecutionTasks(ctx, &tcr.GetTagRetentionExecutionTasksArgs{
// 			ExecutionId: 1,
// 			RegistryId:  "tcr_ins_id",
// 			RetentionId: 1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetTagRetentionExecutionTasks(ctx *pulumi.Context, args *GetTagRetentionExecutionTasksArgs, opts ...pulumi.InvokeOption) (*GetTagRetentionExecutionTasksResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTagRetentionExecutionTasksResult
	err := ctx.Invoke("tencentcloud:Tcr/getTagRetentionExecutionTasks:getTagRetentionExecutionTasks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagRetentionExecutionTasks.
type GetTagRetentionExecutionTasksArgs struct {
	// execution id.
	ExecutionId int `pulumi:"executionId"`
	// instance id.
	RegistryId string `pulumi:"registryId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// retention id.
	RetentionId int `pulumi:"retentionId"`
}

// A collection of values returned by getTagRetentionExecutionTasks.
type GetTagRetentionExecutionTasksResult struct {
	// the rule execution id.
	ExecutionId int `pulumi:"executionId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	RegistryId       string  `pulumi:"registryId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RetentionId      int     `pulumi:"retentionId"`
	// list of version retention tasks.
	RetentionTaskLists []GetTagRetentionExecutionTasksRetentionTaskList `pulumi:"retentionTaskLists"`
}

func GetTagRetentionExecutionTasksOutput(ctx *pulumi.Context, args GetTagRetentionExecutionTasksOutputArgs, opts ...pulumi.InvokeOption) GetTagRetentionExecutionTasksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTagRetentionExecutionTasksResult, error) {
			args := v.(GetTagRetentionExecutionTasksArgs)
			r, err := GetTagRetentionExecutionTasks(ctx, &args, opts...)
			var s GetTagRetentionExecutionTasksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTagRetentionExecutionTasksResultOutput)
}

// A collection of arguments for invoking getTagRetentionExecutionTasks.
type GetTagRetentionExecutionTasksOutputArgs struct {
	// execution id.
	ExecutionId pulumi.IntInput `pulumi:"executionId"`
	// instance id.
	RegistryId pulumi.StringInput `pulumi:"registryId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// retention id.
	RetentionId pulumi.IntInput `pulumi:"retentionId"`
}

func (GetTagRetentionExecutionTasksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagRetentionExecutionTasksArgs)(nil)).Elem()
}

// A collection of values returned by getTagRetentionExecutionTasks.
type GetTagRetentionExecutionTasksResultOutput struct{ *pulumi.OutputState }

func (GetTagRetentionExecutionTasksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagRetentionExecutionTasksResult)(nil)).Elem()
}

func (o GetTagRetentionExecutionTasksResultOutput) ToGetTagRetentionExecutionTasksResultOutput() GetTagRetentionExecutionTasksResultOutput {
	return o
}

func (o GetTagRetentionExecutionTasksResultOutput) ToGetTagRetentionExecutionTasksResultOutputWithContext(ctx context.Context) GetTagRetentionExecutionTasksResultOutput {
	return o
}

// the rule execution id.
func (o GetTagRetentionExecutionTasksResultOutput) ExecutionId() pulumi.IntOutput {
	return o.ApplyT(func(v GetTagRetentionExecutionTasksResult) int { return v.ExecutionId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTagRetentionExecutionTasksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagRetentionExecutionTasksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTagRetentionExecutionTasksResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagRetentionExecutionTasksResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

func (o GetTagRetentionExecutionTasksResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTagRetentionExecutionTasksResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTagRetentionExecutionTasksResultOutput) RetentionId() pulumi.IntOutput {
	return o.ApplyT(func(v GetTagRetentionExecutionTasksResult) int { return v.RetentionId }).(pulumi.IntOutput)
}

// list of version retention tasks.
func (o GetTagRetentionExecutionTasksResultOutput) RetentionTaskLists() GetTagRetentionExecutionTasksRetentionTaskListArrayOutput {
	return o.ApplyT(func(v GetTagRetentionExecutionTasksResult) []GetTagRetentionExecutionTasksRetentionTaskList {
		return v.RetentionTaskLists
	}).(GetTagRetentionExecutionTasksRetentionTaskListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTagRetentionExecutionTasksResultOutput{})
}
