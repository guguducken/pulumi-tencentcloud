// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the list of SQL Server readonly groups.
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
// 		_, err := Sqlserver.GetReadonlyGroups(ctx, &sqlserver.GetReadonlyGroupsArgs{
// 			MasterInstanceId: pulumi.StringRef("mssql-3cdq7kx5"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetReadonlyGroups(ctx *pulumi.Context, args *GetReadonlyGroupsArgs, opts ...pulumi.InvokeOption) (*GetReadonlyGroupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetReadonlyGroupsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getReadonlyGroups:getReadonlyGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReadonlyGroups.
type GetReadonlyGroupsArgs struct {
	// Master SQL Server instance ID.
	MasterInstanceId *string `pulumi:"masterInstanceId"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getReadonlyGroups.
type GetReadonlyGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SQL Server readonly group. Each element contains the following attributes:
	Lists []GetReadonlyGroupsList `pulumi:"lists"`
	// Master instance id.
	MasterInstanceId *string `pulumi:"masterInstanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetReadonlyGroupsOutput(ctx *pulumi.Context, args GetReadonlyGroupsOutputArgs, opts ...pulumi.InvokeOption) GetReadonlyGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReadonlyGroupsResult, error) {
			args := v.(GetReadonlyGroupsArgs)
			r, err := GetReadonlyGroups(ctx, &args, opts...)
			var s GetReadonlyGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReadonlyGroupsResultOutput)
}

// A collection of arguments for invoking getReadonlyGroups.
type GetReadonlyGroupsOutputArgs struct {
	// Master SQL Server instance ID.
	MasterInstanceId pulumi.StringPtrInput `pulumi:"masterInstanceId"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetReadonlyGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReadonlyGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getReadonlyGroups.
type GetReadonlyGroupsResultOutput struct{ *pulumi.OutputState }

func (GetReadonlyGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReadonlyGroupsResult)(nil)).Elem()
}

func (o GetReadonlyGroupsResultOutput) ToGetReadonlyGroupsResultOutput() GetReadonlyGroupsResultOutput {
	return o
}

func (o GetReadonlyGroupsResultOutput) ToGetReadonlyGroupsResultOutputWithContext(ctx context.Context) GetReadonlyGroupsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetReadonlyGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of SQL Server readonly group. Each element contains the following attributes:
func (o GetReadonlyGroupsResultOutput) Lists() GetReadonlyGroupsListArrayOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) []GetReadonlyGroupsList { return v.Lists }).(GetReadonlyGroupsListArrayOutput)
}

// Master instance id.
func (o GetReadonlyGroupsResultOutput) MasterInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) *string { return v.MasterInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetReadonlyGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReadonlyGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReadonlyGroupsResultOutput{})
}
