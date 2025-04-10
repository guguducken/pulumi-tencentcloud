// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourcePackage struct {
	pulumi.CustomResourceState

	// Validity period of resource package, in days.
	ExpireDay pulumi.IntOutput `pulumi:"expireDay"`
	// Instance Type.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Number of purchased resource packs.
	PackageCount pulumi.IntOutput `pulumi:"packageCount"`
	// Resource Package Name.
	PackageName pulumi.StringPtrOutput `pulumi:"packageName"`
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and
	// overseas.
	PackageRegion pulumi.StringOutput `pulumi:"packageRegion"`
	// Resource package size, calculated in 10000 units; Storage resources: GB.
	PackageSpec pulumi.Float64Output `pulumi:"packageSpec"`
	// Resource package type: CCU computing resource package, DISK storage resource package.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// Resource package version base basic version, common general version, enterprise enterprise version.
	PackageVersion pulumi.StringOutput `pulumi:"packageVersion"`
}

// NewResourcePackage registers a new resource with the given unique name, arguments, and options.
func NewResourcePackage(ctx *pulumi.Context,
	name string, args *ResourcePackageArgs, opts ...pulumi.ResourceOption) (*ResourcePackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpireDay == nil {
		return nil, errors.New("invalid value for required argument 'ExpireDay'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.PackageCount == nil {
		return nil, errors.New("invalid value for required argument 'PackageCount'")
	}
	if args.PackageRegion == nil {
		return nil, errors.New("invalid value for required argument 'PackageRegion'")
	}
	if args.PackageSpec == nil {
		return nil, errors.New("invalid value for required argument 'PackageSpec'")
	}
	if args.PackageType == nil {
		return nil, errors.New("invalid value for required argument 'PackageType'")
	}
	if args.PackageVersion == nil {
		return nil, errors.New("invalid value for required argument 'PackageVersion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ResourcePackage
	err := ctx.RegisterResource("tencentcloud:Cynosdb/resourcePackage:ResourcePackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePackage gets an existing ResourcePackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePackageState, opts ...pulumi.ResourceOption) (*ResourcePackage, error) {
	var resource ResourcePackage
	err := ctx.ReadResource("tencentcloud:Cynosdb/resourcePackage:ResourcePackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePackage resources.
type resourcePackageState struct {
	// Validity period of resource package, in days.
	ExpireDay *int `pulumi:"expireDay"`
	// Instance Type.
	InstanceType *string `pulumi:"instanceType"`
	// Number of purchased resource packs.
	PackageCount *int `pulumi:"packageCount"`
	// Resource Package Name.
	PackageName *string `pulumi:"packageName"`
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and
	// overseas.
	PackageRegion *string `pulumi:"packageRegion"`
	// Resource package size, calculated in 10000 units; Storage resources: GB.
	PackageSpec *float64 `pulumi:"packageSpec"`
	// Resource package type: CCU computing resource package, DISK storage resource package.
	PackageType *string `pulumi:"packageType"`
	// Resource package version base basic version, common general version, enterprise enterprise version.
	PackageVersion *string `pulumi:"packageVersion"`
}

type ResourcePackageState struct {
	// Validity period of resource package, in days.
	ExpireDay pulumi.IntPtrInput
	// Instance Type.
	InstanceType pulumi.StringPtrInput
	// Number of purchased resource packs.
	PackageCount pulumi.IntPtrInput
	// Resource Package Name.
	PackageName pulumi.StringPtrInput
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and
	// overseas.
	PackageRegion pulumi.StringPtrInput
	// Resource package size, calculated in 10000 units; Storage resources: GB.
	PackageSpec pulumi.Float64PtrInput
	// Resource package type: CCU computing resource package, DISK storage resource package.
	PackageType pulumi.StringPtrInput
	// Resource package version base basic version, common general version, enterprise enterprise version.
	PackageVersion pulumi.StringPtrInput
}

func (ResourcePackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePackageState)(nil)).Elem()
}

type resourcePackageArgs struct {
	// Validity period of resource package, in days.
	ExpireDay int `pulumi:"expireDay"`
	// Instance Type.
	InstanceType string `pulumi:"instanceType"`
	// Number of purchased resource packs.
	PackageCount int `pulumi:"packageCount"`
	// Resource Package Name.
	PackageName *string `pulumi:"packageName"`
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and
	// overseas.
	PackageRegion string `pulumi:"packageRegion"`
	// Resource package size, calculated in 10000 units; Storage resources: GB.
	PackageSpec float64 `pulumi:"packageSpec"`
	// Resource package type: CCU computing resource package, DISK storage resource package.
	PackageType string `pulumi:"packageType"`
	// Resource package version base basic version, common general version, enterprise enterprise version.
	PackageVersion string `pulumi:"packageVersion"`
}

// The set of arguments for constructing a ResourcePackage resource.
type ResourcePackageArgs struct {
	// Validity period of resource package, in days.
	ExpireDay pulumi.IntInput
	// Instance Type.
	InstanceType pulumi.StringInput
	// Number of purchased resource packs.
	PackageCount pulumi.IntInput
	// Resource Package Name.
	PackageName pulumi.StringPtrInput
	// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and
	// overseas.
	PackageRegion pulumi.StringInput
	// Resource package size, calculated in 10000 units; Storage resources: GB.
	PackageSpec pulumi.Float64Input
	// Resource package type: CCU computing resource package, DISK storage resource package.
	PackageType pulumi.StringInput
	// Resource package version base basic version, common general version, enterprise enterprise version.
	PackageVersion pulumi.StringInput
}

func (ResourcePackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePackageArgs)(nil)).Elem()
}

type ResourcePackageInput interface {
	pulumi.Input

	ToResourcePackageOutput() ResourcePackageOutput
	ToResourcePackageOutputWithContext(ctx context.Context) ResourcePackageOutput
}

func (*ResourcePackage) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePackage)(nil)).Elem()
}

func (i *ResourcePackage) ToResourcePackageOutput() ResourcePackageOutput {
	return i.ToResourcePackageOutputWithContext(context.Background())
}

func (i *ResourcePackage) ToResourcePackageOutputWithContext(ctx context.Context) ResourcePackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePackageOutput)
}

// ResourcePackageArrayInput is an input type that accepts ResourcePackageArray and ResourcePackageArrayOutput values.
// You can construct a concrete instance of `ResourcePackageArrayInput` via:
//
//          ResourcePackageArray{ ResourcePackageArgs{...} }
type ResourcePackageArrayInput interface {
	pulumi.Input

	ToResourcePackageArrayOutput() ResourcePackageArrayOutput
	ToResourcePackageArrayOutputWithContext(context.Context) ResourcePackageArrayOutput
}

type ResourcePackageArray []ResourcePackageInput

func (ResourcePackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePackage)(nil)).Elem()
}

func (i ResourcePackageArray) ToResourcePackageArrayOutput() ResourcePackageArrayOutput {
	return i.ToResourcePackageArrayOutputWithContext(context.Background())
}

func (i ResourcePackageArray) ToResourcePackageArrayOutputWithContext(ctx context.Context) ResourcePackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePackageArrayOutput)
}

// ResourcePackageMapInput is an input type that accepts ResourcePackageMap and ResourcePackageMapOutput values.
// You can construct a concrete instance of `ResourcePackageMapInput` via:
//
//          ResourcePackageMap{ "key": ResourcePackageArgs{...} }
type ResourcePackageMapInput interface {
	pulumi.Input

	ToResourcePackageMapOutput() ResourcePackageMapOutput
	ToResourcePackageMapOutputWithContext(context.Context) ResourcePackageMapOutput
}

type ResourcePackageMap map[string]ResourcePackageInput

func (ResourcePackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePackage)(nil)).Elem()
}

func (i ResourcePackageMap) ToResourcePackageMapOutput() ResourcePackageMapOutput {
	return i.ToResourcePackageMapOutputWithContext(context.Background())
}

func (i ResourcePackageMap) ToResourcePackageMapOutputWithContext(ctx context.Context) ResourcePackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePackageMapOutput)
}

type ResourcePackageOutput struct{ *pulumi.OutputState }

func (ResourcePackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePackage)(nil)).Elem()
}

func (o ResourcePackageOutput) ToResourcePackageOutput() ResourcePackageOutput {
	return o
}

func (o ResourcePackageOutput) ToResourcePackageOutputWithContext(ctx context.Context) ResourcePackageOutput {
	return o
}

// Validity period of resource package, in days.
func (o ResourcePackageOutput) ExpireDay() pulumi.IntOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.IntOutput { return v.ExpireDay }).(pulumi.IntOutput)
}

// Instance Type.
func (o ResourcePackageOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// Number of purchased resource packs.
func (o ResourcePackageOutput) PackageCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.IntOutput { return v.PackageCount }).(pulumi.IntOutput)
}

// Resource Package Name.
func (o ResourcePackageOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.StringPtrOutput { return v.PackageName }).(pulumi.StringPtrOutput)
}

// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and
// overseas.
func (o ResourcePackageOutput) PackageRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.StringOutput { return v.PackageRegion }).(pulumi.StringOutput)
}

// Resource package size, calculated in 10000 units; Storage resources: GB.
func (o ResourcePackageOutput) PackageSpec() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePackage) pulumi.Float64Output { return v.PackageSpec }).(pulumi.Float64Output)
}

// Resource package type: CCU computing resource package, DISK storage resource package.
func (o ResourcePackageOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Resource package version base basic version, common general version, enterprise enterprise version.
func (o ResourcePackageOutput) PackageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePackage) pulumi.StringOutput { return v.PackageVersion }).(pulumi.StringOutput)
}

type ResourcePackageArrayOutput struct{ *pulumi.OutputState }

func (ResourcePackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePackage)(nil)).Elem()
}

func (o ResourcePackageArrayOutput) ToResourcePackageArrayOutput() ResourcePackageArrayOutput {
	return o
}

func (o ResourcePackageArrayOutput) ToResourcePackageArrayOutputWithContext(ctx context.Context) ResourcePackageArrayOutput {
	return o
}

func (o ResourcePackageArrayOutput) Index(i pulumi.IntInput) ResourcePackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcePackage {
		return vs[0].([]*ResourcePackage)[vs[1].(int)]
	}).(ResourcePackageOutput)
}

type ResourcePackageMapOutput struct{ *pulumi.OutputState }

func (ResourcePackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePackage)(nil)).Elem()
}

func (o ResourcePackageMapOutput) ToResourcePackageMapOutput() ResourcePackageMapOutput {
	return o
}

func (o ResourcePackageMapOutput) ToResourcePackageMapOutputWithContext(ctx context.Context) ResourcePackageMapOutput {
	return o
}

func (o ResourcePackageMapOutput) MapIndex(k pulumi.StringInput) ResourcePackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcePackage {
		return vs[0].(map[string]*ResourcePackage)[vs[1].(string)]
	}).(ResourcePackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePackageInput)(nil)).Elem(), &ResourcePackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePackageArrayInput)(nil)).Elem(), ResourcePackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePackageMapInput)(nil)).Elem(), ResourcePackageMap{})
	pulumi.RegisterOutputType(ResourcePackageOutput{})
	pulumi.RegisterOutputType(ResourcePackageArrayOutput{})
	pulumi.RegisterOutputType(ResourcePackageMapOutput{})
}
