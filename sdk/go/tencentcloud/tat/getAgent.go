// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tat agent
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tat"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tat"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tat.GetAgent(ctx, &tat.GetAgentArgs{
// 			Filters: []tat.GetAgentFilter{
// 				tat.GetAgentFilter{
// 					Name: "environment",
// 					Values: []string{
// 						"Linux",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAgent(ctx *pulumi.Context, args *GetAgentArgs, opts ...pulumi.InvokeOption) (*GetAgentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAgentResult
	err := ctx.Invoke("tencentcloud:Tat/getAgent:getAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAgent.
type GetAgentArgs struct {
	// Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
	Filters []GetAgentFilter `pulumi:"filters"`
	// List of instance IDs for the query.
	InstanceIds []string `pulumi:"instanceIds"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAgent.
type GetAgentResult struct {
	// List of agent message.
	AutomationAgentSets []GetAgentAutomationAgentSet `pulumi:"automationAgentSets"`
	Filters             []GetAgentFilter             `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	InstanceIds      []string `pulumi:"instanceIds"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

func GetAgentOutput(ctx *pulumi.Context, args GetAgentOutputArgs, opts ...pulumi.InvokeOption) GetAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAgentResult, error) {
			args := v.(GetAgentArgs)
			r, err := GetAgent(ctx, &args, opts...)
			var s GetAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAgentResultOutput)
}

// A collection of arguments for invoking getAgent.
type GetAgentOutputArgs struct {
	// Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
	Filters GetAgentFilterArrayInput `pulumi:"filters"`
	// List of instance IDs for the query.
	InstanceIds pulumi.StringArrayInput `pulumi:"instanceIds"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentArgs)(nil)).Elem()
}

// A collection of values returned by getAgent.
type GetAgentResultOutput struct{ *pulumi.OutputState }

func (GetAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentResult)(nil)).Elem()
}

func (o GetAgentResultOutput) ToGetAgentResultOutput() GetAgentResultOutput {
	return o
}

func (o GetAgentResultOutput) ToGetAgentResultOutputWithContext(ctx context.Context) GetAgentResultOutput {
	return o
}

// List of agent message.
func (o GetAgentResultOutput) AutomationAgentSets() GetAgentAutomationAgentSetArrayOutput {
	return o.ApplyT(func(v GetAgentResult) []GetAgentAutomationAgentSet { return v.AutomationAgentSets }).(GetAgentAutomationAgentSetArrayOutput)
}

func (o GetAgentResultOutput) Filters() GetAgentFilterArrayOutput {
	return o.ApplyT(func(v GetAgentResult) []GetAgentFilter { return v.Filters }).(GetAgentFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAgentResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAgentResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

func (o GetAgentResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAgentResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAgentResultOutput{})
}
