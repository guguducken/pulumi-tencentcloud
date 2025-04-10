// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query monitor events(There is a lot of data and it is recommended to output to a file)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Monitor.GetProductEvent(ctx, &monitor.GetProductEventArgs{
// 			IsAlarmConfig: pulumi.IntRef(0),
// 			ProductNames: []string{
// 				"cvm",
// 			},
// 			StartTime: pulumi.IntRef(1588700283),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProductEvent(ctx *pulumi.Context, args *GetProductEventArgs, opts ...pulumi.InvokeOption) (*GetProductEventResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProductEventResult
	err := ctx.Invoke("tencentcloud:Monitor/getProductEvent:getProductEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProductEvent.
type GetProductEventArgs struct {
	// Dimensional composition of instance objects.
	Dimensions []GetProductEventDimension `pulumi:"dimensions"`
	// End timestamp for this query, eg:`1588232111`. Default start time is `now-3000`.
	EndTime *int `pulumi:"endTime"`
	// Event name filtering, such as `guestReboot` indicates that the machine restart.
	EventNames []string `pulumi:"eventNames"`
	// Affect objects, such as `ins-19708ino`.
	InstanceIds []string `pulumi:"instanceIds"`
	// Alarm status configuration filter, 1means configured, 0(default) means not configured.
	IsAlarmConfig *int `pulumi:"isAlarmConfig"`
	// Product type filtering, such as `cvm` for cloud server.
	ProductNames []string `pulumi:"productNames"`
	// Project ID filter.
	ProjectIds []string `pulumi:"projectIds"`
	// Region filter, such as `gz`.
	RegionLists []string `pulumi:"regionLists"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start timestamp for this query, eg:`1588230000`. Default start time is `now-3600`.
	StartTime *int `pulumi:"startTime"`
	// Event status filter, value range `-`,`alarm`,`recover`, indicating recovered, unrecovered and stateless.
	Statuses []string `pulumi:"statuses"`
	// Event type filtering, with value range `abnormal`,`statusChange`, indicating state change and abnormal events.
	Types []string `pulumi:"types"`
}

// A collection of values returned by getProductEvent.
type GetProductEventResult struct {
	// A list of event dimensions. Each element contains the following attributes:
	Dimensions []GetProductEventDimension `pulumi:"dimensions"`
	EndTime    *int                       `pulumi:"endTime"`
	// Event short name.
	EventNames []string `pulumi:"eventNames"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instance ID of this event.
	InstanceIds []string `pulumi:"instanceIds"`
	// Whether to configure alarm.
	IsAlarmConfig *int `pulumi:"isAlarmConfig"`
	// A list events. Each element contains the following attributes:
	Lists []GetProductEventList `pulumi:"lists"`
	// Product short name.
	ProductNames []string `pulumi:"productNames"`
	// Project ID of this instance.
	ProjectIds       []string `pulumi:"projectIds"`
	RegionLists      []string `pulumi:"regionLists"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	// The start timestamp of this event.
	StartTime *int `pulumi:"startTime"`
	// The status of this event.
	Statuses []string `pulumi:"statuses"`
	// The type of this event.
	Types []string `pulumi:"types"`
}

func GetProductEventOutput(ctx *pulumi.Context, args GetProductEventOutputArgs, opts ...pulumi.InvokeOption) GetProductEventResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductEventResult, error) {
			args := v.(GetProductEventArgs)
			r, err := GetProductEvent(ctx, &args, opts...)
			var s GetProductEventResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductEventResultOutput)
}

// A collection of arguments for invoking getProductEvent.
type GetProductEventOutputArgs struct {
	// Dimensional composition of instance objects.
	Dimensions GetProductEventDimensionArrayInput `pulumi:"dimensions"`
	// End timestamp for this query, eg:`1588232111`. Default start time is `now-3000`.
	EndTime pulumi.IntPtrInput `pulumi:"endTime"`
	// Event name filtering, such as `guestReboot` indicates that the machine restart.
	EventNames pulumi.StringArrayInput `pulumi:"eventNames"`
	// Affect objects, such as `ins-19708ino`.
	InstanceIds pulumi.StringArrayInput `pulumi:"instanceIds"`
	// Alarm status configuration filter, 1means configured, 0(default) means not configured.
	IsAlarmConfig pulumi.IntPtrInput `pulumi:"isAlarmConfig"`
	// Product type filtering, such as `cvm` for cloud server.
	ProductNames pulumi.StringArrayInput `pulumi:"productNames"`
	// Project ID filter.
	ProjectIds pulumi.StringArrayInput `pulumi:"projectIds"`
	// Region filter, such as `gz`.
	RegionLists pulumi.StringArrayInput `pulumi:"regionLists"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start timestamp for this query, eg:`1588230000`. Default start time is `now-3600`.
	StartTime pulumi.IntPtrInput `pulumi:"startTime"`
	// Event status filter, value range `-`,`alarm`,`recover`, indicating recovered, unrecovered and stateless.
	Statuses pulumi.StringArrayInput `pulumi:"statuses"`
	// Event type filtering, with value range `abnormal`,`statusChange`, indicating state change and abnormal events.
	Types pulumi.StringArrayInput `pulumi:"types"`
}

func (GetProductEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductEventArgs)(nil)).Elem()
}

// A collection of values returned by getProductEvent.
type GetProductEventResultOutput struct{ *pulumi.OutputState }

func (GetProductEventResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductEventResult)(nil)).Elem()
}

func (o GetProductEventResultOutput) ToGetProductEventResultOutput() GetProductEventResultOutput {
	return o
}

func (o GetProductEventResultOutput) ToGetProductEventResultOutputWithContext(ctx context.Context) GetProductEventResultOutput {
	return o
}

// A list of event dimensions. Each element contains the following attributes:
func (o GetProductEventResultOutput) Dimensions() GetProductEventDimensionArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []GetProductEventDimension { return v.Dimensions }).(GetProductEventDimensionArrayOutput)
}

func (o GetProductEventResultOutput) EndTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProductEventResult) *int { return v.EndTime }).(pulumi.IntPtrOutput)
}

// Event short name.
func (o GetProductEventResultOutput) EventNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.EventNames }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProductEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductEventResult) string { return v.Id }).(pulumi.StringOutput)
}

// The instance ID of this event.
func (o GetProductEventResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// Whether to configure alarm.
func (o GetProductEventResultOutput) IsAlarmConfig() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProductEventResult) *int { return v.IsAlarmConfig }).(pulumi.IntPtrOutput)
}

// A list events. Each element contains the following attributes:
func (o GetProductEventResultOutput) Lists() GetProductEventListArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []GetProductEventList { return v.Lists }).(GetProductEventListArrayOutput)
}

// Product short name.
func (o GetProductEventResultOutput) ProductNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.ProductNames }).(pulumi.StringArrayOutput)
}

// Project ID of this instance.
func (o GetProductEventResultOutput) ProjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.ProjectIds }).(pulumi.StringArrayOutput)
}

func (o GetProductEventResultOutput) RegionLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.RegionLists }).(pulumi.StringArrayOutput)
}

func (o GetProductEventResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductEventResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// The start timestamp of this event.
func (o GetProductEventResultOutput) StartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetProductEventResult) *int { return v.StartTime }).(pulumi.IntPtrOutput)
}

// The status of this event.
func (o GetProductEventResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

// The type of this event.
func (o GetProductEventResultOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductEventResult) []string { return v.Types }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductEventResultOutput{})
}
