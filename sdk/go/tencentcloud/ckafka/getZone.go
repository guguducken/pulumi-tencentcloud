// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ckafka zone
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ckafka.GetZone(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupZoneResult
	err := ctx.Invoke("tencentcloud:Ckafka/getZone:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZone.
type LookupZoneArgs struct {
	// cdc professional cluster business parameters.
	CdcId *string `pulumi:"cdcId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getZone.
type LookupZoneResult struct {
	CdcId *string `pulumi:"cdcId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// query result complex object entity.
	Results []GetZoneResult `pulumi:"results"`
}

func LookupZoneOutput(ctx *pulumi.Context, args LookupZoneOutputArgs, opts ...pulumi.InvokeOption) LookupZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZoneResult, error) {
			args := v.(LookupZoneArgs)
			r, err := LookupZone(ctx, &args, opts...)
			var s LookupZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZoneResultOutput)
}

// A collection of arguments for invoking getZone.
type LookupZoneOutputArgs struct {
	// cdc professional cluster business parameters.
	CdcId pulumi.StringPtrInput `pulumi:"cdcId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneArgs)(nil)).Elem()
}

// A collection of values returned by getZone.
type LookupZoneResultOutput struct{ *pulumi.OutputState }

func (LookupZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneResult)(nil)).Elem()
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutput() LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutputWithContext(ctx context.Context) LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) CdcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZoneResult) *string { return v.CdcId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupZoneResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZoneResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// query result complex object entity.
func (o LookupZoneResultOutput) Results() GetZoneResultArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []GetZoneResult { return v.Results }).(GetZoneResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZoneResultOutput{})
}
