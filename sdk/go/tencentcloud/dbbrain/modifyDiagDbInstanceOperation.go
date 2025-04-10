// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dbbrain modifyDiagDbInstanceConf
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
// 		_, err := Dbbrain.NewModifyDiagDbInstanceOperation(ctx, "on", &Dbbrain.ModifyDiagDbInstanceOperationArgs{
// 			InstanceConfs: &dbbrain.ModifyDiagDbInstanceOperationInstanceConfsArgs{
// 				DailyInspection: pulumi.String("Yes"),
// 				OverviewDisplay: pulumi.String("Yes"),
// 			},
// 			InstanceIds: pulumi.StringArray{
// 				pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 			},
// 			Product: pulumi.String("mysql"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
// 		_, err := Dbbrain.NewModifyDiagDbInstanceOperation(ctx, "off", &Dbbrain.ModifyDiagDbInstanceOperationArgs{
// 			InstanceConfs: &dbbrain.ModifyDiagDbInstanceOperationInstanceConfsArgs{
// 				DailyInspection: pulumi.String("No"),
// 				OverviewDisplay: pulumi.String("No"),
// 			},
// 			InstanceIds: pulumi.StringArray{
// 				pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 			},
// 			Product: pulumi.String("mysql"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ModifyDiagDbInstanceOperation struct {
	pulumi.CustomResourceState

	// Instance configuration, including inspection, overview switch, etc.
	InstanceConfs ModifyDiagDbInstanceOperationInstanceConfsOutput `pulumi:"instanceConfs"`
	// Specifies the ID of the instance whose inspection status is changed.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
	Product pulumi.StringOutput `pulumi:"product"`
	// Effective instance region, the value is All, which means all regions.
	Regions pulumi.StringPtrOutput `pulumi:"regions"`
}

// NewModifyDiagDbInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewModifyDiagDbInstanceOperation(ctx *pulumi.Context,
	name string, args *ModifyDiagDbInstanceOperationArgs, opts ...pulumi.ResourceOption) (*ModifyDiagDbInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceConfs == nil {
		return nil, errors.New("invalid value for required argument 'InstanceConfs'")
	}
	if args.Product == nil {
		return nil, errors.New("invalid value for required argument 'Product'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ModifyDiagDbInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Dbbrain/modifyDiagDbInstanceOperation:ModifyDiagDbInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModifyDiagDbInstanceOperation gets an existing ModifyDiagDbInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModifyDiagDbInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModifyDiagDbInstanceOperationState, opts ...pulumi.ResourceOption) (*ModifyDiagDbInstanceOperation, error) {
	var resource ModifyDiagDbInstanceOperation
	err := ctx.ReadResource("tencentcloud:Dbbrain/modifyDiagDbInstanceOperation:ModifyDiagDbInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModifyDiagDbInstanceOperation resources.
type modifyDiagDbInstanceOperationState struct {
	// Instance configuration, including inspection, overview switch, etc.
	InstanceConfs *ModifyDiagDbInstanceOperationInstanceConfs `pulumi:"instanceConfs"`
	// Specifies the ID of the instance whose inspection status is changed.
	InstanceIds []string `pulumi:"instanceIds"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
	Product *string `pulumi:"product"`
	// Effective instance region, the value is All, which means all regions.
	Regions *string `pulumi:"regions"`
}

type ModifyDiagDbInstanceOperationState struct {
	// Instance configuration, including inspection, overview switch, etc.
	InstanceConfs ModifyDiagDbInstanceOperationInstanceConfsPtrInput
	// Specifies the ID of the instance whose inspection status is changed.
	InstanceIds pulumi.StringArrayInput
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
	Product pulumi.StringPtrInput
	// Effective instance region, the value is All, which means all regions.
	Regions pulumi.StringPtrInput
}

func (ModifyDiagDbInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*modifyDiagDbInstanceOperationState)(nil)).Elem()
}

type modifyDiagDbInstanceOperationArgs struct {
	// Instance configuration, including inspection, overview switch, etc.
	InstanceConfs ModifyDiagDbInstanceOperationInstanceConfs `pulumi:"instanceConfs"`
	// Specifies the ID of the instance whose inspection status is changed.
	InstanceIds []string `pulumi:"instanceIds"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
	Product string `pulumi:"product"`
	// Effective instance region, the value is All, which means all regions.
	Regions *string `pulumi:"regions"`
}

// The set of arguments for constructing a ModifyDiagDbInstanceOperation resource.
type ModifyDiagDbInstanceOperationArgs struct {
	// Instance configuration, including inspection, overview switch, etc.
	InstanceConfs ModifyDiagDbInstanceOperationInstanceConfsInput
	// Specifies the ID of the instance whose inspection status is changed.
	InstanceIds pulumi.StringArrayInput
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
	Product pulumi.StringInput
	// Effective instance region, the value is All, which means all regions.
	Regions pulumi.StringPtrInput
}

func (ModifyDiagDbInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modifyDiagDbInstanceOperationArgs)(nil)).Elem()
}

type ModifyDiagDbInstanceOperationInput interface {
	pulumi.Input

	ToModifyDiagDbInstanceOperationOutput() ModifyDiagDbInstanceOperationOutput
	ToModifyDiagDbInstanceOperationOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationOutput
}

func (*ModifyDiagDbInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ModifyDiagDbInstanceOperation)(nil)).Elem()
}

func (i *ModifyDiagDbInstanceOperation) ToModifyDiagDbInstanceOperationOutput() ModifyDiagDbInstanceOperationOutput {
	return i.ToModifyDiagDbInstanceOperationOutputWithContext(context.Background())
}

func (i *ModifyDiagDbInstanceOperation) ToModifyDiagDbInstanceOperationOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyDiagDbInstanceOperationOutput)
}

// ModifyDiagDbInstanceOperationArrayInput is an input type that accepts ModifyDiagDbInstanceOperationArray and ModifyDiagDbInstanceOperationArrayOutput values.
// You can construct a concrete instance of `ModifyDiagDbInstanceOperationArrayInput` via:
//
//          ModifyDiagDbInstanceOperationArray{ ModifyDiagDbInstanceOperationArgs{...} }
type ModifyDiagDbInstanceOperationArrayInput interface {
	pulumi.Input

	ToModifyDiagDbInstanceOperationArrayOutput() ModifyDiagDbInstanceOperationArrayOutput
	ToModifyDiagDbInstanceOperationArrayOutputWithContext(context.Context) ModifyDiagDbInstanceOperationArrayOutput
}

type ModifyDiagDbInstanceOperationArray []ModifyDiagDbInstanceOperationInput

func (ModifyDiagDbInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ModifyDiagDbInstanceOperation)(nil)).Elem()
}

func (i ModifyDiagDbInstanceOperationArray) ToModifyDiagDbInstanceOperationArrayOutput() ModifyDiagDbInstanceOperationArrayOutput {
	return i.ToModifyDiagDbInstanceOperationArrayOutputWithContext(context.Background())
}

func (i ModifyDiagDbInstanceOperationArray) ToModifyDiagDbInstanceOperationArrayOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyDiagDbInstanceOperationArrayOutput)
}

// ModifyDiagDbInstanceOperationMapInput is an input type that accepts ModifyDiagDbInstanceOperationMap and ModifyDiagDbInstanceOperationMapOutput values.
// You can construct a concrete instance of `ModifyDiagDbInstanceOperationMapInput` via:
//
//          ModifyDiagDbInstanceOperationMap{ "key": ModifyDiagDbInstanceOperationArgs{...} }
type ModifyDiagDbInstanceOperationMapInput interface {
	pulumi.Input

	ToModifyDiagDbInstanceOperationMapOutput() ModifyDiagDbInstanceOperationMapOutput
	ToModifyDiagDbInstanceOperationMapOutputWithContext(context.Context) ModifyDiagDbInstanceOperationMapOutput
}

type ModifyDiagDbInstanceOperationMap map[string]ModifyDiagDbInstanceOperationInput

func (ModifyDiagDbInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ModifyDiagDbInstanceOperation)(nil)).Elem()
}

func (i ModifyDiagDbInstanceOperationMap) ToModifyDiagDbInstanceOperationMapOutput() ModifyDiagDbInstanceOperationMapOutput {
	return i.ToModifyDiagDbInstanceOperationMapOutputWithContext(context.Background())
}

func (i ModifyDiagDbInstanceOperationMap) ToModifyDiagDbInstanceOperationMapOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyDiagDbInstanceOperationMapOutput)
}

type ModifyDiagDbInstanceOperationOutput struct{ *pulumi.OutputState }

func (ModifyDiagDbInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModifyDiagDbInstanceOperation)(nil)).Elem()
}

func (o ModifyDiagDbInstanceOperationOutput) ToModifyDiagDbInstanceOperationOutput() ModifyDiagDbInstanceOperationOutput {
	return o
}

func (o ModifyDiagDbInstanceOperationOutput) ToModifyDiagDbInstanceOperationOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationOutput {
	return o
}

// Instance configuration, including inspection, overview switch, etc.
func (o ModifyDiagDbInstanceOperationOutput) InstanceConfs() ModifyDiagDbInstanceOperationInstanceConfsOutput {
	return o.ApplyT(func(v *ModifyDiagDbInstanceOperation) ModifyDiagDbInstanceOperationInstanceConfsOutput {
		return v.InstanceConfs
	}).(ModifyDiagDbInstanceOperationInstanceConfsOutput)
}

// Specifies the ID of the instance whose inspection status is changed.
func (o ModifyDiagDbInstanceOperationOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModifyDiagDbInstanceOperation) pulumi.StringArrayOutput { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL.
func (o ModifyDiagDbInstanceOperationOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v *ModifyDiagDbInstanceOperation) pulumi.StringOutput { return v.Product }).(pulumi.StringOutput)
}

// Effective instance region, the value is All, which means all regions.
func (o ModifyDiagDbInstanceOperationOutput) Regions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModifyDiagDbInstanceOperation) pulumi.StringPtrOutput { return v.Regions }).(pulumi.StringPtrOutput)
}

type ModifyDiagDbInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (ModifyDiagDbInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ModifyDiagDbInstanceOperation)(nil)).Elem()
}

func (o ModifyDiagDbInstanceOperationArrayOutput) ToModifyDiagDbInstanceOperationArrayOutput() ModifyDiagDbInstanceOperationArrayOutput {
	return o
}

func (o ModifyDiagDbInstanceOperationArrayOutput) ToModifyDiagDbInstanceOperationArrayOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationArrayOutput {
	return o
}

func (o ModifyDiagDbInstanceOperationArrayOutput) Index(i pulumi.IntInput) ModifyDiagDbInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ModifyDiagDbInstanceOperation {
		return vs[0].([]*ModifyDiagDbInstanceOperation)[vs[1].(int)]
	}).(ModifyDiagDbInstanceOperationOutput)
}

type ModifyDiagDbInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (ModifyDiagDbInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ModifyDiagDbInstanceOperation)(nil)).Elem()
}

func (o ModifyDiagDbInstanceOperationMapOutput) ToModifyDiagDbInstanceOperationMapOutput() ModifyDiagDbInstanceOperationMapOutput {
	return o
}

func (o ModifyDiagDbInstanceOperationMapOutput) ToModifyDiagDbInstanceOperationMapOutputWithContext(ctx context.Context) ModifyDiagDbInstanceOperationMapOutput {
	return o
}

func (o ModifyDiagDbInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) ModifyDiagDbInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ModifyDiagDbInstanceOperation {
		return vs[0].(map[string]*ModifyDiagDbInstanceOperation)[vs[1].(string)]
	}).(ModifyDiagDbInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyDiagDbInstanceOperationInput)(nil)).Elem(), &ModifyDiagDbInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyDiagDbInstanceOperationArrayInput)(nil)).Elem(), ModifyDiagDbInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyDiagDbInstanceOperationMapInput)(nil)).Elem(), ModifyDiagDbInstanceOperationMap{})
	pulumi.RegisterOutputType(ModifyDiagDbInstanceOperationOutput{})
	pulumi.RegisterOutputType(ModifyDiagDbInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(ModifyDiagDbInstanceOperationMapOutput{})
}
