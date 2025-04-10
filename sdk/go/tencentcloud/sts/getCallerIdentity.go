// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of sts callerIdentity
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Sts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Sts.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCallerIdentity(ctx *pulumi.Context, args *GetCallerIdentityArgs, opts ...pulumi.InvokeOption) (*GetCallerIdentityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCallerIdentityResult
	err := ctx.Invoke("tencentcloud:Sts/getCallerIdentity:getCallerIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCallerIdentity.
type GetCallerIdentityArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCallerIdentity.
type GetCallerIdentityResult struct {
	// The primary account Uin to which the current caller belongs.
	AccountId string `pulumi:"accountId"`
	// Current caller ARN.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Account Uin to which the key belongs:- The caller is a cloud account, and the returned current account Uin- The caller is a role, and the returned account Uin that applies for the role key.
	PrincipalId      string  `pulumi:"principalId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Identity type.
	Type string `pulumi:"type"`
	// Identity:- When the caller is a cloud account, the current account `Uin` is returned.- When the caller is a role, it returns `roleId:roleSessionName`- When the caller is a federated identity, it returns `uin:federatedUserName`.
	UserId string `pulumi:"userId"`
}

func GetCallerIdentityOutput(ctx *pulumi.Context, args GetCallerIdentityOutputArgs, opts ...pulumi.InvokeOption) GetCallerIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCallerIdentityResult, error) {
			args := v.(GetCallerIdentityArgs)
			r, err := GetCallerIdentity(ctx, &args, opts...)
			var s GetCallerIdentityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCallerIdentityResultOutput)
}

// A collection of arguments for invoking getCallerIdentity.
type GetCallerIdentityOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCallerIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCallerIdentityArgs)(nil)).Elem()
}

// A collection of values returned by getCallerIdentity.
type GetCallerIdentityResultOutput struct{ *pulumi.OutputState }

func (GetCallerIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCallerIdentityResult)(nil)).Elem()
}

func (o GetCallerIdentityResultOutput) ToGetCallerIdentityResultOutput() GetCallerIdentityResultOutput {
	return o
}

func (o GetCallerIdentityResultOutput) ToGetCallerIdentityResultOutputWithContext(ctx context.Context) GetCallerIdentityResultOutput {
	return o
}

// The primary account Uin to which the current caller belongs.
func (o GetCallerIdentityResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// Current caller ARN.
func (o GetCallerIdentityResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCallerIdentityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) string { return v.Id }).(pulumi.StringOutput)
}

// Account Uin to which the key belongs:- The caller is a cloud account, and the returned current account Uin- The caller is a role, and the returned account Uin that applies for the role key.
func (o GetCallerIdentityResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o GetCallerIdentityResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Identity type.
func (o GetCallerIdentityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) string { return v.Type }).(pulumi.StringOutput)
}

// Identity:- When the caller is a cloud account, the current account `Uin` is returned.- When the caller is a role, it returns `roleId:roleSessionName`- When the caller is a federated identity, it returns `uin:federatedUserName`.
func (o GetCallerIdentityResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCallerIdentityResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCallerIdentityResultOutput{})
}
