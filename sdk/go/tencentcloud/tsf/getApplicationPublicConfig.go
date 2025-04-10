// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tsf applicationPublicConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tsf.GetApplicationPublicConfig(ctx, &tsf.GetApplicationPublicConfigArgs{
// 			ConfigId:      pulumi.StringRef("dcfg-p-evjrbgly"),
// 			ConfigName:    pulumi.StringRef("dsadsa"),
// 			ConfigVersion: pulumi.StringRef("123"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupApplicationPublicConfig(ctx *pulumi.Context, args *LookupApplicationPublicConfigArgs, opts ...pulumi.InvokeOption) (*LookupApplicationPublicConfigResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupApplicationPublicConfigResult
	err := ctx.Invoke("tencentcloud:Tsf/getApplicationPublicConfig:getApplicationPublicConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationPublicConfig.
type LookupApplicationPublicConfigArgs struct {
	// Config ID. Query all items if not passed, high priority.
	ConfigId *string `pulumi:"configId"`
	// Config ID list. Query all items if not passed, low priority.
	ConfigIdLists []string `pulumi:"configIdLists"`
	// Config name. Exact query. Query all items if not passed.
	ConfigName *string `pulumi:"configName"`
	// Config version. Exact query. Query all items if not passed.
	ConfigVersion *string `pulumi:"configVersion"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getApplicationPublicConfig.
type LookupApplicationPublicConfigResult struct {
	// Config ID. Note: This field may return null, indicating that no valid value can be obtained.
	ConfigId      *string  `pulumi:"configId"`
	ConfigIdLists []string `pulumi:"configIdLists"`
	// Config name. Note: This field may return null, indicating that no valid value can be obtained.
	ConfigName *string `pulumi:"configName"`
	// Config version. Note: This field may return null, indicating that no valid value can be obtained.
	ConfigVersion *string `pulumi:"configVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Paginated global configuration  list. Note: This field may return null, indicating that no valid value can be obtained.
	Results []GetApplicationPublicConfigResult `pulumi:"results"`
}

func LookupApplicationPublicConfigOutput(ctx *pulumi.Context, args LookupApplicationPublicConfigOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationPublicConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationPublicConfigResult, error) {
			args := v.(LookupApplicationPublicConfigArgs)
			r, err := LookupApplicationPublicConfig(ctx, &args, opts...)
			var s LookupApplicationPublicConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationPublicConfigResultOutput)
}

// A collection of arguments for invoking getApplicationPublicConfig.
type LookupApplicationPublicConfigOutputArgs struct {
	// Config ID. Query all items if not passed, high priority.
	ConfigId pulumi.StringPtrInput `pulumi:"configId"`
	// Config ID list. Query all items if not passed, low priority.
	ConfigIdLists pulumi.StringArrayInput `pulumi:"configIdLists"`
	// Config name. Exact query. Query all items if not passed.
	ConfigName pulumi.StringPtrInput `pulumi:"configName"`
	// Config version. Exact query. Query all items if not passed.
	ConfigVersion pulumi.StringPtrInput `pulumi:"configVersion"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupApplicationPublicConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationPublicConfigArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationPublicConfig.
type LookupApplicationPublicConfigResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationPublicConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationPublicConfigResult)(nil)).Elem()
}

func (o LookupApplicationPublicConfigResultOutput) ToLookupApplicationPublicConfigResultOutput() LookupApplicationPublicConfigResultOutput {
	return o
}

func (o LookupApplicationPublicConfigResultOutput) ToLookupApplicationPublicConfigResultOutputWithContext(ctx context.Context) LookupApplicationPublicConfigResultOutput {
	return o
}

// Config ID. Note: This field may return null, indicating that no valid value can be obtained.
func (o LookupApplicationPublicConfigResultOutput) ConfigId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) *string { return v.ConfigId }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationPublicConfigResultOutput) ConfigIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) []string { return v.ConfigIdLists }).(pulumi.StringArrayOutput)
}

// Config name. Note: This field may return null, indicating that no valid value can be obtained.
func (o LookupApplicationPublicConfigResultOutput) ConfigName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) *string { return v.ConfigName }).(pulumi.StringPtrOutput)
}

// Config version. Note: This field may return null, indicating that no valid value can be obtained.
func (o LookupApplicationPublicConfigResultOutput) ConfigVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) *string { return v.ConfigVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApplicationPublicConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationPublicConfigResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Paginated global configuration  list. Note: This field may return null, indicating that no valid value can be obtained.
func (o LookupApplicationPublicConfigResultOutput) Results() GetApplicationPublicConfigResultArrayOutput {
	return o.ApplyT(func(v LookupApplicationPublicConfigResult) []GetApplicationPublicConfigResult { return v.Results }).(GetApplicationPublicConfigResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationPublicConfigResultOutput{})
}
