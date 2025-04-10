// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an available EMR for the user.
//
// The EMR data source fetch proper EMR from user's EMR pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Emr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Emr"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Emr.GetInstance(ctx, &emr.GetInstanceArgs{
// 			DisplayStrategy: "clusterList",
// 			InstanceIds: []string{
// 				"emr-rnzqrleq",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstance(ctx *pulumi.Context, args *GetInstanceArgs, opts ...pulumi.InvokeOption) (*GetInstanceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceResult
	err := ctx.Invoke("tencentcloud:Emr/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// Display strategy(e.g.:clusterList, monitorManage).
	DisplayStrategy string `pulumi:"displayStrategy"`
	// fetch all instances with same prefix(e.g.:emr-xxxxxx).
	InstanceIds []string `pulumi:"instanceIds"`
	// Fetch all instances which owner same project. Default 0 meaning use default project id.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	// A list of clusters will be exported and its every element contains the following attributes:
	Clusters        []GetInstanceCluster `pulumi:"clusters"`
	DisplayStrategy string               `pulumi:"displayStrategy"`
	// The provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	InstanceIds []string `pulumi:"instanceIds"`
	// Project id of instance.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetInstanceOutput(ctx *pulumi.Context, args GetInstanceOutputArgs, opts ...pulumi.InvokeOption) GetInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceResult, error) {
			args := v.(GetInstanceArgs)
			r, err := GetInstance(ctx, &args, opts...)
			var s GetInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type GetInstanceOutputArgs struct {
	// Display strategy(e.g.:clusterList, monitorManage).
	DisplayStrategy pulumi.StringInput `pulumi:"displayStrategy"`
	// fetch all instances with same prefix(e.g.:emr-xxxxxx).
	InstanceIds pulumi.StringArrayInput `pulumi:"instanceIds"`
	// Fetch all instances which owner same project. Default 0 meaning use default project id.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type GetInstanceResultOutput struct{ *pulumi.OutputState }

func (GetInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceResult)(nil)).Elem()
}

func (o GetInstanceResultOutput) ToGetInstanceResultOutput() GetInstanceResultOutput {
	return o
}

func (o GetInstanceResultOutput) ToGetInstanceResultOutputWithContext(ctx context.Context) GetInstanceResultOutput {
	return o
}

// A list of clusters will be exported and its every element contains the following attributes:
func (o GetInstanceResultOutput) Clusters() GetInstanceClusterArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []GetInstanceCluster { return v.Clusters }).(GetInstanceClusterArrayOutput)
}

func (o GetInstanceResultOutput) DisplayStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.DisplayStrategy }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// Project id of instance.
func (o GetInstanceResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetInstanceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceResultOutput{})
}
