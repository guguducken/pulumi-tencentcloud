// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dbbrain noPrimaryKeyTables
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dbbrain.GetNoPrimaryKeyTables(ctx, &dbbrain.GetNoPrimaryKeyTablesArgs{
// 			Date:       "",
// 			InstanceId: "",
// 			Product:    pulumi.StringRef(""),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetNoPrimaryKeyTables(ctx *pulumi.Context, args *GetNoPrimaryKeyTablesArgs, opts ...pulumi.InvokeOption) (*GetNoPrimaryKeyTablesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetNoPrimaryKeyTablesResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getNoPrimaryKeyTables:getNoPrimaryKeyTables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNoPrimaryKeyTables.
type GetNoPrimaryKeyTablesArgs struct {
	// Query date, such as 2021-05-27, the earliest date is 30 days ago.
	Date string `pulumi:"date"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Service product type, supported values: `mysql` - ApsaraDB for MySQL, the default is `mysql`.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getNoPrimaryKeyTables.
type GetNoPrimaryKeyTablesResult struct {
	Date string `pulumi:"date"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The difference with yesterday&amp;#39;s scan of the table without a primary key. A positive number means an increase, a negative number means a decrease, and 0 means no change.
	NoPrimaryKeyTableCountDiff int `pulumi:"noPrimaryKeyTableCountDiff"`
	// A list of tables without primary keys.
	NoPrimaryKeyTables []GetNoPrimaryKeyTablesNoPrimaryKeyTable `pulumi:"noPrimaryKeyTables"`
	Product            *string                                  `pulumi:"product"`
	ResultOutputFile   *string                                  `pulumi:"resultOutputFile"`
	// Collection timestamp (seconds).
	Timestamp int `pulumi:"timestamp"`
}

func GetNoPrimaryKeyTablesOutput(ctx *pulumi.Context, args GetNoPrimaryKeyTablesOutputArgs, opts ...pulumi.InvokeOption) GetNoPrimaryKeyTablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNoPrimaryKeyTablesResult, error) {
			args := v.(GetNoPrimaryKeyTablesArgs)
			r, err := GetNoPrimaryKeyTables(ctx, &args, opts...)
			var s GetNoPrimaryKeyTablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNoPrimaryKeyTablesResultOutput)
}

// A collection of arguments for invoking getNoPrimaryKeyTables.
type GetNoPrimaryKeyTablesOutputArgs struct {
	// Query date, such as 2021-05-27, the earliest date is 30 days ago.
	Date pulumi.StringInput `pulumi:"date"`
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Service product type, supported values: `mysql` - ApsaraDB for MySQL, the default is `mysql`.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetNoPrimaryKeyTablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNoPrimaryKeyTablesArgs)(nil)).Elem()
}

// A collection of values returned by getNoPrimaryKeyTables.
type GetNoPrimaryKeyTablesResultOutput struct{ *pulumi.OutputState }

func (GetNoPrimaryKeyTablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNoPrimaryKeyTablesResult)(nil)).Elem()
}

func (o GetNoPrimaryKeyTablesResultOutput) ToGetNoPrimaryKeyTablesResultOutput() GetNoPrimaryKeyTablesResultOutput {
	return o
}

func (o GetNoPrimaryKeyTablesResultOutput) ToGetNoPrimaryKeyTablesResultOutputWithContext(ctx context.Context) GetNoPrimaryKeyTablesResultOutput {
	return o
}

func (o GetNoPrimaryKeyTablesResultOutput) Date() pulumi.StringOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) string { return v.Date }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNoPrimaryKeyTablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNoPrimaryKeyTablesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The difference with yesterday&amp;#39;s scan of the table without a primary key. A positive number means an increase, a negative number means a decrease, and 0 means no change.
func (o GetNoPrimaryKeyTablesResultOutput) NoPrimaryKeyTableCountDiff() pulumi.IntOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) int { return v.NoPrimaryKeyTableCountDiff }).(pulumi.IntOutput)
}

// A list of tables without primary keys.
func (o GetNoPrimaryKeyTablesResultOutput) NoPrimaryKeyTables() GetNoPrimaryKeyTablesNoPrimaryKeyTableArrayOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) []GetNoPrimaryKeyTablesNoPrimaryKeyTable {
		return v.NoPrimaryKeyTables
	}).(GetNoPrimaryKeyTablesNoPrimaryKeyTableArrayOutput)
}

func (o GetNoPrimaryKeyTablesResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetNoPrimaryKeyTablesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Collection timestamp (seconds).
func (o GetNoPrimaryKeyTablesResultOutput) Timestamp() pulumi.IntOutput {
	return o.ApplyT(func(v GetNoPrimaryKeyTablesResult) int { return v.Timestamp }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNoPrimaryKeyTablesResultOutput{})
}
