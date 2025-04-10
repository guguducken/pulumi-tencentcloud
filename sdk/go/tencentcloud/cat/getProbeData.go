// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cat probe data
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cat"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cat"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cat.GetProbeData(ctx, &cat.GetProbeDataArgs{
// 			Ascending: true,
// 			BeginTime: 1667923200000,
// 			EndTime:   1667996208428,
// 			Limit:     20,
// 			Offset:    0,
// 			SelectedFields: []string{
// 				"terraform",
// 			},
// 			SortField: "ProbeTime",
// 			TaskIds: []string{
// 				"task-knare1mk",
// 			},
// 			TaskType: "AnalyzeTaskType_Network",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProbeData(ctx *pulumi.Context, args *GetProbeDataArgs, opts ...pulumi.InvokeOption) (*GetProbeDataResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProbeDataResult
	err := ctx.Invoke("tencentcloud:Cat/getProbeData:getProbeData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProbeData.
type GetProbeDataArgs struct {
	// true is Ascending.
	Ascending bool `pulumi:"ascending"`
	// Start timestamp (in milliseconds).
	BeginTime int `pulumi:"beginTime"`
	// City list.
	Cities []string `pulumi:"cities"`
	// Code list.
	Codes []string `pulumi:"codes"`
	// Districts list.
	Districts []string `pulumi:"districts"`
	// End timestamp (in milliseconds).
	EndTime int `pulumi:"endTime"`
	// ErrorTypes list.
	ErrorTypes []string `pulumi:"errorTypes"`
	// Limit.
	Limit int `pulumi:"limit"`
	// Offset.
	Offset int `pulumi:"offset"`
	// Operators list.
	Operators []string `pulumi:"operators"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Selected Fields.
	SelectedFields []string `pulumi:"selectedFields"`
	// Fields to be sorted ProbeTime dial test time sorting can be filled in You can also fill in the selected fields in SelectedFields.
	SortField string `pulumi:"sortField"`
	// TaskID list.
	TaskIds []string `pulumi:"taskIds"`
	// Task Type in AnalyzeTaskType_Network,AnalyzeTaskType_Browse,AnalyzeTaskType_UploadDownload,AnalyzeTaskType_Transport,AnalyzeTaskType_MediaStream.
	TaskType string `pulumi:"taskType"`
}

// A collection of values returned by getProbeData.
type GetProbeDataResult struct {
	Ascending bool     `pulumi:"ascending"`
	BeginTime int      `pulumi:"beginTime"`
	Cities    []string `pulumi:"cities"`
	Codes     []string `pulumi:"codes"`
	// Probe node list.
	DetailedSingleDataDefines []GetProbeDataDetailedSingleDataDefine `pulumi:"detailedSingleDataDefines"`
	Districts                 []string                               `pulumi:"districts"`
	EndTime                   int                                    `pulumi:"endTime"`
	ErrorTypes                []string                               `pulumi:"errorTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	Limit            int      `pulumi:"limit"`
	Offset           int      `pulumi:"offset"`
	Operators        []string `pulumi:"operators"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	SelectedFields   []string `pulumi:"selectedFields"`
	SortField        string   `pulumi:"sortField"`
	TaskIds          []string `pulumi:"taskIds"`
	TaskType         string   `pulumi:"taskType"`
}

func GetProbeDataOutput(ctx *pulumi.Context, args GetProbeDataOutputArgs, opts ...pulumi.InvokeOption) GetProbeDataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProbeDataResult, error) {
			args := v.(GetProbeDataArgs)
			r, err := GetProbeData(ctx, &args, opts...)
			var s GetProbeDataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProbeDataResultOutput)
}

// A collection of arguments for invoking getProbeData.
type GetProbeDataOutputArgs struct {
	// true is Ascending.
	Ascending pulumi.BoolInput `pulumi:"ascending"`
	// Start timestamp (in milliseconds).
	BeginTime pulumi.IntInput `pulumi:"beginTime"`
	// City list.
	Cities pulumi.StringArrayInput `pulumi:"cities"`
	// Code list.
	Codes pulumi.StringArrayInput `pulumi:"codes"`
	// Districts list.
	Districts pulumi.StringArrayInput `pulumi:"districts"`
	// End timestamp (in milliseconds).
	EndTime pulumi.IntInput `pulumi:"endTime"`
	// ErrorTypes list.
	ErrorTypes pulumi.StringArrayInput `pulumi:"errorTypes"`
	// Limit.
	Limit pulumi.IntInput `pulumi:"limit"`
	// Offset.
	Offset pulumi.IntInput `pulumi:"offset"`
	// Operators list.
	Operators pulumi.StringArrayInput `pulumi:"operators"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Selected Fields.
	SelectedFields pulumi.StringArrayInput `pulumi:"selectedFields"`
	// Fields to be sorted ProbeTime dial test time sorting can be filled in You can also fill in the selected fields in SelectedFields.
	SortField pulumi.StringInput `pulumi:"sortField"`
	// TaskID list.
	TaskIds pulumi.StringArrayInput `pulumi:"taskIds"`
	// Task Type in AnalyzeTaskType_Network,AnalyzeTaskType_Browse,AnalyzeTaskType_UploadDownload,AnalyzeTaskType_Transport,AnalyzeTaskType_MediaStream.
	TaskType pulumi.StringInput `pulumi:"taskType"`
}

func (GetProbeDataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataArgs)(nil)).Elem()
}

// A collection of values returned by getProbeData.
type GetProbeDataResultOutput struct{ *pulumi.OutputState }

func (GetProbeDataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProbeDataResult)(nil)).Elem()
}

func (o GetProbeDataResultOutput) ToGetProbeDataResultOutput() GetProbeDataResultOutput {
	return o
}

func (o GetProbeDataResultOutput) ToGetProbeDataResultOutputWithContext(ctx context.Context) GetProbeDataResultOutput {
	return o
}

func (o GetProbeDataResultOutput) Ascending() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProbeDataResult) bool { return v.Ascending }).(pulumi.BoolOutput)
}

func (o GetProbeDataResultOutput) BeginTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataResult) int { return v.BeginTime }).(pulumi.IntOutput)
}

func (o GetProbeDataResultOutput) Cities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.Cities }).(pulumi.StringArrayOutput)
}

func (o GetProbeDataResultOutput) Codes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.Codes }).(pulumi.StringArrayOutput)
}

// Probe node list.
func (o GetProbeDataResultOutput) DetailedSingleDataDefines() GetProbeDataDetailedSingleDataDefineArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []GetProbeDataDetailedSingleDataDefine { return v.DetailedSingleDataDefines }).(GetProbeDataDetailedSingleDataDefineArrayOutput)
}

func (o GetProbeDataResultOutput) Districts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.Districts }).(pulumi.StringArrayOutput)
}

func (o GetProbeDataResultOutput) EndTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataResult) int { return v.EndTime }).(pulumi.IntOutput)
}

func (o GetProbeDataResultOutput) ErrorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.ErrorTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProbeDataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProbeDataResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProbeDataResultOutput) Limit() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataResult) int { return v.Limit }).(pulumi.IntOutput)
}

func (o GetProbeDataResultOutput) Offset() pulumi.IntOutput {
	return o.ApplyT(func(v GetProbeDataResult) int { return v.Offset }).(pulumi.IntOutput)
}

func (o GetProbeDataResultOutput) Operators() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.Operators }).(pulumi.StringArrayOutput)
}

func (o GetProbeDataResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProbeDataResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetProbeDataResultOutput) SelectedFields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.SelectedFields }).(pulumi.StringArrayOutput)
}

func (o GetProbeDataResultOutput) SortField() pulumi.StringOutput {
	return o.ApplyT(func(v GetProbeDataResult) string { return v.SortField }).(pulumi.StringOutput)
}

func (o GetProbeDataResultOutput) TaskIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProbeDataResult) []string { return v.TaskIds }).(pulumi.StringArrayOutput)
}

func (o GetProbeDataResultOutput) TaskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetProbeDataResult) string { return v.TaskType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProbeDataResultOutput{})
}
