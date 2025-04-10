// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cynosdb resourcePackageSaleSpecs
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
// 		_, err := Cynosdb.GetResourcePackageSaleSpecs(ctx, &cynosdb.GetResourcePackageSaleSpecsArgs{
// 			InstanceType:  "cynosdb-serverless",
// 			PackageRegion: "china",
// 			PackageType:   "CCU",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetResourcePackageSaleSpecs(ctx *pulumi.Context, args *GetResourcePackageSaleSpecsArgs, opts ...pulumi.InvokeOption) (*GetResourcePackageSaleSpecsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetResourcePackageSaleSpecsResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getResourcePackageSaleSpecs:getResourcePackageSaleSpecs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsArgs struct {
	// Instance Type. Value range: cynosdb-serverless, cynosdb, cdb.
	InstanceType string `pulumi:"instanceType"`
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
	PackageRegion string `pulumi:"packageRegion"`
	// Resource package type CCU - Computing resource package DISK - Storage resource package.
	PackageType string `pulumi:"packageType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsResult struct {
	// Resource package details note: This field may return null, indicating that a valid value cannot be obtained.
	Details []GetResourcePackageSaleSpecsDetail `pulumi:"details"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	InstanceType string `pulumi:"instanceType"`
	// Note: This field may return null, indicating that a valid value cannot be obtained.
	PackageRegion string `pulumi:"packageRegion"`
	// Resource package type CCU - Compute resource package DISK - Store resource package Note: This field may return null, indicating that a valid value cannot be obtained.
	PackageType      string  `pulumi:"packageType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetResourcePackageSaleSpecsOutput(ctx *pulumi.Context, args GetResourcePackageSaleSpecsOutputArgs, opts ...pulumi.InvokeOption) GetResourcePackageSaleSpecsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResourcePackageSaleSpecsResult, error) {
			args := v.(GetResourcePackageSaleSpecsArgs)
			r, err := GetResourcePackageSaleSpecs(ctx, &args, opts...)
			var s GetResourcePackageSaleSpecsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetResourcePackageSaleSpecsResultOutput)
}

// A collection of arguments for invoking getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsOutputArgs struct {
	// Instance Type. Value range: cynosdb-serverless, cynosdb, cdb.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
	PackageRegion pulumi.StringInput `pulumi:"packageRegion"`
	// Resource package type CCU - Computing resource package DISK - Storage resource package.
	PackageType pulumi.StringInput `pulumi:"packageType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetResourcePackageSaleSpecsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcePackageSaleSpecsArgs)(nil)).Elem()
}

// A collection of values returned by getResourcePackageSaleSpecs.
type GetResourcePackageSaleSpecsResultOutput struct{ *pulumi.OutputState }

func (GetResourcePackageSaleSpecsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcePackageSaleSpecsResult)(nil)).Elem()
}

func (o GetResourcePackageSaleSpecsResultOutput) ToGetResourcePackageSaleSpecsResultOutput() GetResourcePackageSaleSpecsResultOutput {
	return o
}

func (o GetResourcePackageSaleSpecsResultOutput) ToGetResourcePackageSaleSpecsResultOutputWithContext(ctx context.Context) GetResourcePackageSaleSpecsResultOutput {
	return o
}

// Resource package details note: This field may return null, indicating that a valid value cannot be obtained.
func (o GetResourcePackageSaleSpecsResultOutput) Details() GetResourcePackageSaleSpecsDetailArrayOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) []GetResourcePackageSaleSpecsDetail { return v.Details }).(GetResourcePackageSaleSpecsDetailArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetResourcePackageSaleSpecsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetResourcePackageSaleSpecsResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

// Note: This field may return null, indicating that a valid value cannot be obtained.
func (o GetResourcePackageSaleSpecsResultOutput) PackageRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.PackageRegion }).(pulumi.StringOutput)
}

// Resource package type CCU - Compute resource package DISK - Store resource package Note: This field may return null, indicating that a valid value cannot be obtained.
func (o GetResourcePackageSaleSpecsResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetResourcePackageSaleSpecsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResourcePackageSaleSpecsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourcePackageSaleSpecsResultOutput{})
}
