// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mysql slowLogData
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
// 		_, err := Mysql.GetSlowLogData(ctx, &mysql.GetSlowLogDataArgs{
// 			DataBases: []string{
// 				"tf_ci_test",
// 			},
// 			EndTime:    1684392459,
// 			InstType:   pulumi.StringRef("slave"),
// 			InstanceId: "cdb-fitq5t9h",
// 			OrderBy:    pulumi.StringRef("ASC"),
// 			SortBy:     pulumi.StringRef("Timestamp"),
// 			StartTime:  1682664459,
// 			UserHosts: []string{
// 				"169.254.128.158",
// 			},
// 			UserNames: []string{
// 				"keep_dts",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSlowLogData(ctx *pulumi.Context, args *GetSlowLogDataArgs, opts ...pulumi.InvokeOption) (*GetSlowLogDataResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSlowLogDataResult
	err := ctx.Invoke("tencentcloud:Mysql/getSlowLogData:getSlowLogData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlowLogData.
type GetSlowLogDataArgs struct {
	// List of databases accessed.
	DataBases []string `pulumi:"dataBases"`
	// End timestamp. For example 1585142640.
	EndTime int `pulumi:"endTime"`
	// Only valid when the instance is the master instance or disaster recovery instance, the optional value: slave, which means to pull the log of the slave machine.
	InstType *string `pulumi:"instType"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Sort in ascending or descending order. Currently supported: ASC,DESC.
	OrderBy *string `pulumi:"orderBy"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Sort field. Currently supported: Timestamp, QueryTime, LockTime, RowsExamined, RowsSent.
	SortBy *string `pulumi:"sortBy"`
	// Start timestamp. For example 1585142640.
	StartTime int `pulumi:"startTime"`
	// List of client hosts.
	UserHosts []string `pulumi:"userHosts"`
	// A list of client usernames.
	UserNames []string `pulumi:"userNames"`
}

// A collection of values returned by getSlowLogData.
type GetSlowLogDataResult struct {
	DataBases []string `pulumi:"dataBases"`
	EndTime   int      `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstType   *string `pulumi:"instType"`
	InstanceId string  `pulumi:"instanceId"`
	// Query records.
	Items            []GetSlowLogDataItem `pulumi:"items"`
	OrderBy          *string              `pulumi:"orderBy"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
	SortBy           *string              `pulumi:"sortBy"`
	StartTime        int                  `pulumi:"startTime"`
	UserHosts        []string             `pulumi:"userHosts"`
	UserNames        []string             `pulumi:"userNames"`
}

func GetSlowLogDataOutput(ctx *pulumi.Context, args GetSlowLogDataOutputArgs, opts ...pulumi.InvokeOption) GetSlowLogDataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlowLogDataResult, error) {
			args := v.(GetSlowLogDataArgs)
			r, err := GetSlowLogData(ctx, &args, opts...)
			var s GetSlowLogDataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlowLogDataResultOutput)
}

// A collection of arguments for invoking getSlowLogData.
type GetSlowLogDataOutputArgs struct {
	// List of databases accessed.
	DataBases pulumi.StringArrayInput `pulumi:"dataBases"`
	// End timestamp. For example 1585142640.
	EndTime pulumi.IntInput `pulumi:"endTime"`
	// Only valid when the instance is the master instance or disaster recovery instance, the optional value: slave, which means to pull the log of the slave machine.
	InstType pulumi.StringPtrInput `pulumi:"instType"`
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Sort in ascending or descending order. Currently supported: ASC,DESC.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Sort field. Currently supported: Timestamp, QueryTime, LockTime, RowsExamined, RowsSent.
	SortBy pulumi.StringPtrInput `pulumi:"sortBy"`
	// Start timestamp. For example 1585142640.
	StartTime pulumi.IntInput `pulumi:"startTime"`
	// List of client hosts.
	UserHosts pulumi.StringArrayInput `pulumi:"userHosts"`
	// A list of client usernames.
	UserNames pulumi.StringArrayInput `pulumi:"userNames"`
}

func (GetSlowLogDataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogDataArgs)(nil)).Elem()
}

// A collection of values returned by getSlowLogData.
type GetSlowLogDataResultOutput struct{ *pulumi.OutputState }

func (GetSlowLogDataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlowLogDataResult)(nil)).Elem()
}

func (o GetSlowLogDataResultOutput) ToGetSlowLogDataResultOutput() GetSlowLogDataResultOutput {
	return o
}

func (o GetSlowLogDataResultOutput) ToGetSlowLogDataResultOutputWithContext(ctx context.Context) GetSlowLogDataResultOutput {
	return o
}

func (o GetSlowLogDataResultOutput) DataBases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) []string { return v.DataBases }).(pulumi.StringArrayOutput)
}

func (o GetSlowLogDataResultOutput) EndTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) int { return v.EndTime }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlowLogDataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlowLogDataResultOutput) InstType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) *string { return v.InstType }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogDataResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Query records.
func (o GetSlowLogDataResultOutput) Items() GetSlowLogDataItemArrayOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) []GetSlowLogDataItem { return v.Items }).(GetSlowLogDataItemArrayOutput)
}

func (o GetSlowLogDataResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogDataResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogDataResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

func (o GetSlowLogDataResultOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) int { return v.StartTime }).(pulumi.IntOutput)
}

func (o GetSlowLogDataResultOutput) UserHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) []string { return v.UserHosts }).(pulumi.StringArrayOutput)
}

func (o GetSlowLogDataResultOutput) UserNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSlowLogDataResult) []string { return v.UserNames }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlowLogDataResultOutput{})
}
