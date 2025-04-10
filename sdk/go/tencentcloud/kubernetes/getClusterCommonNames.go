// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a datasource to query cluster CommonNames.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Kubernetes.GetClusterCommonNames(ctx, &kubernetes.GetClusterCommonNamesArgs{
// 			ClusterId: pulumi.StringRef("cls-12345678"),
// 			SubaccountUins: []string{
// 				"1234567890",
// 				"0987654321",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetClusterCommonNames(ctx *pulumi.Context, args *GetClusterCommonNamesArgs, opts ...pulumi.InvokeOption) (*GetClusterCommonNamesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetClusterCommonNamesResult
	err := ctx.Invoke("tencentcloud:Kubernetes/getClusterCommonNames:getClusterCommonNames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterCommonNames.
type GetClusterCommonNamesArgs struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Used for save result.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// List of Role ID. Up to 50 sub-accounts can be passed in at a time.
	RoleIds []string `pulumi:"roleIds"`
	// List of sub-account. Up to 50 sub-accounts can be passed in at a time.
	SubaccountUins []string `pulumi:"subaccountUins"`
}

// A collection of values returned by getClusterCommonNames.
type GetClusterCommonNamesResult struct {
	ClusterId *string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of the CommonName in the certificate of the client corresponding to the sub-account UIN.
	Lists            []GetClusterCommonNamesList `pulumi:"lists"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
	RoleIds          []string                    `pulumi:"roleIds"`
	SubaccountUins   []string                    `pulumi:"subaccountUins"`
}

func GetClusterCommonNamesOutput(ctx *pulumi.Context, args GetClusterCommonNamesOutputArgs, opts ...pulumi.InvokeOption) GetClusterCommonNamesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterCommonNamesResult, error) {
			args := v.(GetClusterCommonNamesArgs)
			r, err := GetClusterCommonNames(ctx, &args, opts...)
			var s GetClusterCommonNamesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterCommonNamesResultOutput)
}

// A collection of arguments for invoking getClusterCommonNames.
type GetClusterCommonNamesOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Used for save result.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// List of Role ID. Up to 50 sub-accounts can be passed in at a time.
	RoleIds pulumi.StringArrayInput `pulumi:"roleIds"`
	// List of sub-account. Up to 50 sub-accounts can be passed in at a time.
	SubaccountUins pulumi.StringArrayInput `pulumi:"subaccountUins"`
}

func (GetClusterCommonNamesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCommonNamesArgs)(nil)).Elem()
}

// A collection of values returned by getClusterCommonNames.
type GetClusterCommonNamesResultOutput struct{ *pulumi.OutputState }

func (GetClusterCommonNamesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCommonNamesResult)(nil)).Elem()
}

func (o GetClusterCommonNamesResultOutput) ToGetClusterCommonNamesResultOutput() GetClusterCommonNamesResultOutput {
	return o
}

func (o GetClusterCommonNamesResultOutput) ToGetClusterCommonNamesResultOutputWithContext(ctx context.Context) GetClusterCommonNamesResultOutput {
	return o
}

func (o GetClusterCommonNamesResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterCommonNamesResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterCommonNamesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCommonNamesResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of the CommonName in the certificate of the client corresponding to the sub-account UIN.
func (o GetClusterCommonNamesResultOutput) Lists() GetClusterCommonNamesListArrayOutput {
	return o.ApplyT(func(v GetClusterCommonNamesResult) []GetClusterCommonNamesList { return v.Lists }).(GetClusterCommonNamesListArrayOutput)
}

func (o GetClusterCommonNamesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterCommonNamesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetClusterCommonNamesResultOutput) RoleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClusterCommonNamesResult) []string { return v.RoleIds }).(pulumi.StringArrayOutput)
}

func (o GetClusterCommonNamesResultOutput) SubaccountUins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClusterCommonNamesResult) []string { return v.SubaccountUins }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterCommonNamesResultOutput{})
}
