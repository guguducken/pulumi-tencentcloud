// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mysql projectSecurityGroup
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mysql.GetProjectSecurityGroup(ctx, &mysql.GetProjectSecurityGroupArgs{
// 			ProjectId: pulumi.IntRef(1250480),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProjectSecurityGroup(ctx *pulumi.Context, args *GetProjectSecurityGroupArgs, opts ...pulumi.InvokeOption) (*GetProjectSecurityGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProjectSecurityGroupResult
	err := ctx.Invoke("tencentcloud:Mysql/getProjectSecurityGroup:getProjectSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectSecurityGroup.
type GetProjectSecurityGroupArgs struct {
	// project id.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getProjectSecurityGroup.
type GetProjectSecurityGroupResult struct {
	// Security group details.
	Groups []GetProjectSecurityGroupGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// project id.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetProjectSecurityGroupOutput(ctx *pulumi.Context, args GetProjectSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) GetProjectSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectSecurityGroupResult, error) {
			args := v.(GetProjectSecurityGroupArgs)
			r, err := GetProjectSecurityGroup(ctx, &args, opts...)
			var s GetProjectSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectSecurityGroupResultOutput)
}

// A collection of arguments for invoking getProjectSecurityGroup.
type GetProjectSecurityGroupOutputArgs struct {
	// project id.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetProjectSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getProjectSecurityGroup.
type GetProjectSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (GetProjectSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectSecurityGroupResult)(nil)).Elem()
}

func (o GetProjectSecurityGroupResultOutput) ToGetProjectSecurityGroupResultOutput() GetProjectSecurityGroupResultOutput {
	return o
}

func (o GetProjectSecurityGroupResultOutput) ToGetProjectSecurityGroupResultOutputWithContext(ctx context.Context) GetProjectSecurityGroupResultOutput {
	return o
}

// Security group details.
func (o GetProjectSecurityGroupResultOutput) Groups() GetProjectSecurityGroupGroupArrayOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupResult) []GetProjectSecurityGroupGroup { return v.Groups }).(GetProjectSecurityGroupGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// project id.
func (o GetProjectSecurityGroupResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetProjectSecurityGroupResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectSecurityGroupResultOutput{})
}
