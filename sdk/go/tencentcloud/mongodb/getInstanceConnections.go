// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mongodb instanceConnections
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mongodb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mongodb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mongodb.GetInstanceConnections(ctx, &mongodb.GetInstanceConnectionsArgs{
// 			InstanceId: "cmgo-9d0p6umb",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstanceConnections(ctx *pulumi.Context, args *GetInstanceConnectionsArgs, opts ...pulumi.InvokeOption) (*GetInstanceConnectionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceConnectionsResult
	err := ctx.Invoke("tencentcloud:Mongodb/getInstanceConnections:getInstanceConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceConnections.
type GetInstanceConnectionsArgs struct {
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceConnections.
type GetInstanceConnectionsResult struct {
	// Client connection info.
	Clients []GetInstanceConnectionsClient `pulumi:"clients"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetInstanceConnectionsOutput(ctx *pulumi.Context, args GetInstanceConnectionsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceConnectionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceConnectionsResult, error) {
			args := v.(GetInstanceConnectionsArgs)
			r, err := GetInstanceConnections(ctx, &args, opts...)
			var s GetInstanceConnectionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceConnectionsResultOutput)
}

// A collection of arguments for invoking getInstanceConnections.
type GetInstanceConnectionsOutputArgs struct {
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConnectionsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceConnections.
type GetInstanceConnectionsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConnectionsResult)(nil)).Elem()
}

func (o GetInstanceConnectionsResultOutput) ToGetInstanceConnectionsResultOutput() GetInstanceConnectionsResultOutput {
	return o
}

func (o GetInstanceConnectionsResultOutput) ToGetInstanceConnectionsResultOutputWithContext(ctx context.Context) GetInstanceConnectionsResultOutput {
	return o
}

// Client connection info.
func (o GetInstanceConnectionsResultOutput) Clients() GetInstanceConnectionsClientArrayOutput {
	return o.ApplyT(func(v GetInstanceConnectionsResult) []GetInstanceConnectionsClient { return v.Clients }).(GetInstanceConnectionsClientArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceConnectionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConnectionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceConnectionsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConnectionsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetInstanceConnectionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConnectionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceConnectionsResultOutput{})
}
