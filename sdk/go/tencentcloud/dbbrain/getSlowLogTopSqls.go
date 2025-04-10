// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dbbrain slowLogTopSqls
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dbbrain.GetSlowLogTopSqls(ctx, &dbbrain.GetSlowLogTopSqlsArgs{
// 			EndTime:    fmt.Sprintf("%v%v", "%", "s"),
// 			InstanceId: fmt.Sprintf("%v%v", "%", "s"),
// 			OrderBy:    pulumi.StringRef("ASC"),
// 			Product:    pulumi.StringRef("mysql"),
// 			SortBy:     pulumi.StringRef("QueryTimeMax"),
// 			StartTime:  fmt.Sprintf("%v%v", "%", "s"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSlowLogTopSqls(ctx *pulumi.Context, args *GetSlowLogTopSqlsArgs, opts ...pulumi.InvokeOption) (*GetSlowLogTopSqlsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSlowLogTopSqlsResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getSlowLogTopSqls:getSlowLogTopSqls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowLogTopSqls.
type GetSlowLogTopSqlsArgs struct {
	// The deadline, such as `2019-09-11 10:13:14`, the interval between the deadline and the start time is less than 7 days.
	EndTime string `pulumi:"endTime"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// The sorting method supports ASC (ascending) and DESC (descending). The default is DESC.
	OrderBy *string `pulumi:"orderBy"`
	// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Array of database names.
	SchemaLists []GetSlowLogTopSqlsSchemaList `pulumi:"schemaLists"`
	// Sort key, currently supports sort keys such as QueryTime, ExecTimes, RowsSent, LockTime and RowsExamined, the default is QueryTime.
	SortBy *string `pulumi:"sortBy"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getSlowLogTopSqls.
type GetSlowLogTopSqlsResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	OrderBy          *string `pulumi:"orderBy"`
	Product          *string `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Slow log top sql list.
	Rows        []GetSlowLogTopSqlsRow        `pulumi:"rows"`
	SchemaLists []GetSlowLogTopSqlsSchemaList `pulumi:"schemaLists"`
	SortBy      *string                       `pulumi:"sortBy"`
	StartTime   string                        `pulumi:"startTime"`
}

func GetSlowLogTopSqlsOutput(ctx *pulumi.Context, args GetSlowLogTopSqlsOutputArgs, opts ...pulumi.InvokeOption) GetSlowLogTopSqlsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowLogTopSqlsResult, error) {
			args := v.(GetSlowLogTopSqlsArgs)
			r, err := GetSlowLogTopSqls(ctx, &args, opts...)
			var s GetSlowLogTopSqlsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowLogTopSqlsResultOutput)
}

// A collection of arguments for invoking getSlowLogTopSqls.
type GetSlowLogTopSqlsOutputArgs struct {
	// The deadline, such as `2019-09-11 10:13:14`, the interval between the deadline and the start time is less than 7 days.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The sorting method supports ASC (ascending) and DESC (descending). The default is DESC.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Array of database names.
	SchemaLists GetSlowLogTopSqlsSchemaListArrayInput `pulumi:"schemaLists"`
	// Sort key, currently supports sort keys such as QueryTime, ExecTimes, RowsSent, LockTime and RowsExamined, the default is QueryTime.
	SortBy pulumi.StringPtrInput `pulumi:"sortBy"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetSlowLogTopSqlsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogTopSqlsArgs)(nil)).Elem()
}

// A collection of values returned by getSlowLogTopSqls.
type GetSlowLogTopSqlsResultOutput struct{ *pulumi.OutputState }

func (GetSlowLogTopSqlsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogTopSqlsResult)(nil)).Elem()
}

func (o GetSlowLogTopSqlsResultOutput) ToGetSlowLogTopSqlsResultOutput() GetSlowLogTopSqlsResultOutput {
	return o
}

func (o GetSlowLogTopSqlsResultOutput) ToGetSlowLogTopSqlsResultOutputWithContext(ctx context.Context) GetSlowLogTopSqlsResultOutput {
	return o
}

func (o GetSlowLogTopSqlsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowLogTopSqlsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowLogTopSqlsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSlowLogTopSqlsResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Slow log top sql list.
func (o GetSlowLogTopSqlsResultOutput) Rows() GetSlowLogTopSqlsRowArrayOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) []GetSlowLogTopSqlsRow { return v.Rows }).(GetSlowLogTopSqlsRowArrayOutput)
}

func (o GetSlowLogTopSqlsResultOutput) SchemaLists() GetSlowLogTopSqlsSchemaListArrayOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) []GetSlowLogTopSqlsSchemaList { return v.SchemaLists }).(GetSlowLogTopSqlsSchemaListArrayOutput)
}

func (o GetSlowLogTopSqlsResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogTopSqlsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogTopSqlsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowLogTopSqlsResultOutput{})
}
