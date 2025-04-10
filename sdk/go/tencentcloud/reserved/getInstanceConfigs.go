// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package reserved

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query reserved instances configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Reserved"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Reserved"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Reserved.GetInstanceConfigs(ctx, &reserved.GetInstanceConfigsArgs{
// 			AvailabilityZone: pulumi.StringRef("na-siliconvalley-1"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstanceConfigs(ctx *pulumi.Context, args *GetInstanceConfigsArgs, opts ...pulumi.InvokeOption) (*GetInstanceConfigsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceConfigsResult
	err := ctx.Invoke("tencentcloud:Reserved/getInstanceConfigs:getInstanceConfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceConfigs.
type GetInstanceConfigsArgs struct {
	// The available zone that the reserved instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
	Duration *int `pulumi:"duration"`
	// The type of reserved instance.
	InstanceType *string `pulumi:"instanceType"`
	// Filter by Payment Type. Such as All Upfront.
	OfferingType *string `pulumi:"offeringType"`
	// Filter by the Platform Description (that is, operating system) for Reserved Instance billing. Shaped like: linux.
	ProductDescription *string `pulumi:"productDescription"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceConfigs.
type GetInstanceConfigsResult struct {
	// Availability zone of the purchasable reserved instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// An information list of reserved instance configuration. Each element contains the following attributes:
	ConfigLists []GetInstanceConfigsConfigList `pulumi:"configLists"`
	// Validity period of the reserved instance.
	Duration *int `pulumi:"duration"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance type of the reserved instance.
	InstanceType *string `pulumi:"instanceType"`
	// OfferingType of the reserved instance.
	OfferingType       *string `pulumi:"offeringType"`
	ProductDescription *string `pulumi:"productDescription"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

func GetInstanceConfigsOutput(ctx *pulumi.Context, args GetInstanceConfigsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceConfigsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceConfigsResult, error) {
			args := v.(GetInstanceConfigsArgs)
			r, err := GetInstanceConfigs(ctx, &args, opts...)
			var s GetInstanceConfigsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceConfigsResultOutput)
}

// A collection of arguments for invoking getInstanceConfigs.
type GetInstanceConfigsOutputArgs struct {
	// The available zone that the reserved instance locates at.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Validity period of the reserved instance. Valid values are `31536000`(1 year) and `94608000`(3 years).
	Duration pulumi.IntPtrInput `pulumi:"duration"`
	// The type of reserved instance.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Filter by Payment Type. Such as All Upfront.
	OfferingType pulumi.StringPtrInput `pulumi:"offeringType"`
	// Filter by the Platform Description (that is, operating system) for Reserved Instance billing. Shaped like: linux.
	ProductDescription pulumi.StringPtrInput `pulumi:"productDescription"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceConfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConfigsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceConfigs.
type GetInstanceConfigsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceConfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConfigsResult)(nil)).Elem()
}

func (o GetInstanceConfigsResultOutput) ToGetInstanceConfigsResultOutput() GetInstanceConfigsResultOutput {
	return o
}

func (o GetInstanceConfigsResultOutput) ToGetInstanceConfigsResultOutputWithContext(ctx context.Context) GetInstanceConfigsResultOutput {
	return o
}

// Availability zone of the purchasable reserved instance.
func (o GetInstanceConfigsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// An information list of reserved instance configuration. Each element contains the following attributes:
func (o GetInstanceConfigsResultOutput) ConfigLists() GetInstanceConfigsConfigListArrayOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) []GetInstanceConfigsConfigList { return v.ConfigLists }).(GetInstanceConfigsConfigListArrayOutput)
}

// Validity period of the reserved instance.
func (o GetInstanceConfigsResultOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *int { return v.Duration }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceConfigsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance type of the reserved instance.
func (o GetInstanceConfigsResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// OfferingType of the reserved instance.
func (o GetInstanceConfigsResultOutput) OfferingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.OfferingType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceConfigsResultOutput) ProductDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.ProductDescription }).(pulumi.StringPtrOutput)
}

func (o GetInstanceConfigsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceConfigsResultOutput{})
}
