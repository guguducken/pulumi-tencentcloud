// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cynosdb clusterParamLogs
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cynosdb.GetClusterParamLogs(ctx, &cynosdb.GetClusterParamLogsArgs{
// 			ClusterId: "cynosdbmysql-bws8h88b",
// 			InstanceIds: []string{
// 				"cynosdbmysql-ins-afqx1hy0",
// 			},
// 			OrderBy:     pulumi.StringRef("CreateTime"),
// 			OrderByType: pulumi.StringRef("DESC"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetClusterParamLogs(ctx *pulumi.Context, args *GetClusterParamLogsArgs, opts ...pulumi.InvokeOption) (*GetClusterParamLogsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetClusterParamLogsResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getClusterParamLogs:getClusterParamLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterParamLogs.
type GetClusterParamLogsArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Instance ID list, used to record specific instances of operations.
	InstanceIds []string `pulumi:"instanceIds"`
	// Sort field, defining which field to sort based on when returning results.
	OrderBy *string `pulumi:"orderBy"`
	// Define specific sorting rules, limited to one of desc, asc, DESC, or ASC.
	OrderByType *string `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getClusterParamLogs.
type GetClusterParamLogsResult struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Parameter modification record note: This field may return null, indicating that a valid value cannot be obtained.
	ClusterParamLogs []GetClusterParamLogsClusterParamLog `pulumi:"clusterParamLogs"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	InstanceIds      []string `pulumi:"instanceIds"`
	OrderBy          *string  `pulumi:"orderBy"`
	OrderByType      *string  `pulumi:"orderByType"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

func GetClusterParamLogsOutput(ctx *pulumi.Context, args GetClusterParamLogsOutputArgs, opts ...pulumi.InvokeOption) GetClusterParamLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterParamLogsResult, error) {
			args := v.(GetClusterParamLogsArgs)
			r, err := GetClusterParamLogs(ctx, &args, opts...)
			var s GetClusterParamLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterParamLogsResultOutput)
}

// A collection of arguments for invoking getClusterParamLogs.
type GetClusterParamLogsOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Instance ID list, used to record specific instances of operations.
	InstanceIds pulumi.StringArrayInput `pulumi:"instanceIds"`
	// Sort field, defining which field to sort based on when returning results.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Define specific sorting rules, limited to one of desc, asc, DESC, or ASC.
	OrderByType pulumi.StringPtrInput `pulumi:"orderByType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetClusterParamLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterParamLogsArgs)(nil)).Elem()
}

// A collection of values returned by getClusterParamLogs.
type GetClusterParamLogsResultOutput struct{ *pulumi.OutputState }

func (GetClusterParamLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterParamLogsResult)(nil)).Elem()
}

func (o GetClusterParamLogsResultOutput) ToGetClusterParamLogsResultOutput() GetClusterParamLogsResultOutput {
	return o
}

func (o GetClusterParamLogsResultOutput) ToGetClusterParamLogsResultOutputWithContext(ctx context.Context) GetClusterParamLogsResultOutput {
	return o
}

// Cluster ID.
func (o GetClusterParamLogsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Parameter modification record note: This field may return null, indicating that a valid value cannot be obtained.
func (o GetClusterParamLogsResultOutput) ClusterParamLogs() GetClusterParamLogsClusterParamLogArrayOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) []GetClusterParamLogsClusterParamLog { return v.ClusterParamLogs }).(GetClusterParamLogsClusterParamLogArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterParamLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClusterParamLogsResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

func (o GetClusterParamLogsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetClusterParamLogsResultOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) *string { return v.OrderByType }).(pulumi.StringPtrOutput)
}

func (o GetClusterParamLogsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterParamLogsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterParamLogsResultOutput{})
}
