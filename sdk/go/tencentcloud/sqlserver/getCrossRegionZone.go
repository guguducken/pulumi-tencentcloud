// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of sqlserver datasourceCrossRegionZone
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
// 		_, err := Sqlserver.GetCrossRegionZone(ctx, &sqlserver.GetCrossRegionZoneArgs{
// 			InstanceId: "mssql-qelbzgwf",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCrossRegionZone(ctx *pulumi.Context, args *GetCrossRegionZoneArgs, opts ...pulumi.InvokeOption) (*GetCrossRegionZoneResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCrossRegionZoneResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getCrossRegionZone:getCrossRegionZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCrossRegionZone.
type GetCrossRegionZoneArgs struct {
	// Instance ID in the format of mssql-j8kv137v.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCrossRegionZone.
type GetCrossRegionZoneResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The string ID of the region where the standby machine is located, such as: ap-guangzhou.
	Region           string  `pulumi:"region"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The string ID of the availability zone where the standby machine is located, such as: ap-guangzhou-1.
	Zone string `pulumi:"zone"`
}

func GetCrossRegionZoneOutput(ctx *pulumi.Context, args GetCrossRegionZoneOutputArgs, opts ...pulumi.InvokeOption) GetCrossRegionZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCrossRegionZoneResult, error) {
			args := v.(GetCrossRegionZoneArgs)
			r, err := GetCrossRegionZone(ctx, &args, opts...)
			var s GetCrossRegionZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCrossRegionZoneResultOutput)
}

// A collection of arguments for invoking getCrossRegionZone.
type GetCrossRegionZoneOutputArgs struct {
	// Instance ID in the format of mssql-j8kv137v.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCrossRegionZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCrossRegionZoneArgs)(nil)).Elem()
}

// A collection of values returned by getCrossRegionZone.
type GetCrossRegionZoneResultOutput struct{ *pulumi.OutputState }

func (GetCrossRegionZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCrossRegionZoneResult)(nil)).Elem()
}

func (o GetCrossRegionZoneResultOutput) ToGetCrossRegionZoneResultOutput() GetCrossRegionZoneResultOutput {
	return o
}

func (o GetCrossRegionZoneResultOutput) ToGetCrossRegionZoneResultOutputWithContext(ctx context.Context) GetCrossRegionZoneResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCrossRegionZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCrossRegionZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCrossRegionZoneResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCrossRegionZoneResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The string ID of the region where the standby machine is located, such as: ap-guangzhou.
func (o GetCrossRegionZoneResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCrossRegionZoneResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetCrossRegionZoneResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionZoneResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// The string ID of the availability zone where the standby machine is located, such as: ap-guangzhou-1.
func (o GetCrossRegionZoneResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetCrossRegionZoneResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCrossRegionZoneResultOutput{})
}
