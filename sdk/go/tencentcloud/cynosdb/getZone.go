// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cynosdb zone
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cynosdb.GetZone(ctx, &cynosdb.GetZoneArgs{
// 			IncludeVirtualZones: pulumi.BoolRef(true),
// 			ShowPermission:      pulumi.BoolRef(true),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetZone(ctx *pulumi.Context, args *GetZoneArgs, opts ...pulumi.InvokeOption) (*GetZoneResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetZoneResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getZone:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZone.
type GetZoneArgs struct {
	// Is virtual zone included.
	IncludeVirtualZones *bool `pulumi:"includeVirtualZones"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Whether to display all available zones under the region and display the permissions of each available zone of the user.
	ShowPermission *bool `pulumi:"showPermission"`
}

// A collection of values returned by getZone.
type GetZoneResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                  string `pulumi:"id"`
	IncludeVirtualZones *bool  `pulumi:"includeVirtualZones"`
	// Information of region.
	RegionSets       []GetZoneRegionSet `pulumi:"regionSets"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	ShowPermission   *bool              `pulumi:"showPermission"`
}

func GetZoneOutput(ctx *pulumi.Context, args GetZoneOutputArgs, opts ...pulumi.InvokeOption) GetZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZoneResult, error) {
			args := v.(GetZoneArgs)
			r, err := GetZone(ctx, &args, opts...)
			var s GetZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZoneResultOutput)
}

// A collection of arguments for invoking getZone.
type GetZoneOutputArgs struct {
	// Is virtual zone included.
	IncludeVirtualZones pulumi.BoolPtrInput `pulumi:"includeVirtualZones"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Whether to display all available zones under the region and display the permissions of each available zone of the user.
	ShowPermission pulumi.BoolPtrInput `pulumi:"showPermission"`
}

func (GetZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneArgs)(nil)).Elem()
}

// A collection of values returned by getZone.
type GetZoneResultOutput struct{ *pulumi.OutputState }

func (GetZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneResult)(nil)).Elem()
}

func (o GetZoneResultOutput) ToGetZoneResultOutput() GetZoneResultOutput {
	return o
}

func (o GetZoneResultOutput) ToGetZoneResultOutputWithContext(ctx context.Context) GetZoneResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetZoneResultOutput) IncludeVirtualZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetZoneResult) *bool { return v.IncludeVirtualZones }).(pulumi.BoolPtrOutput)
}

// Information of region.
func (o GetZoneResultOutput) RegionSets() GetZoneRegionSetArrayOutput {
	return o.ApplyT(func(v GetZoneResult) []GetZoneRegionSet { return v.RegionSets }).(GetZoneRegionSetArrayOutput)
}

func (o GetZoneResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetZoneResultOutput) ShowPermission() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetZoneResult) *bool { return v.ShowPermission }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZoneResultOutput{})
}
