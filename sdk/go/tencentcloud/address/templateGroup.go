// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package address

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage address template group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Address"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Address.NewTemplateGroup(ctx, "foo", &Address.TemplateGroupArgs{
// 			TemplateIds: pulumi.StringArray{
// 				pulumi.String("ipl-axaf24151"),
// 				pulumi.String("ipl-axaf24152"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Address template group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Address/templateGroup:TemplateGroup foo ipmg-0np3u974
// ```
type TemplateGroup struct {
	pulumi.CustomResourceState

	// Name of the address template group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Template ID list.
	TemplateIds pulumi.StringArrayOutput `pulumi:"templateIds"`
}

// NewTemplateGroup registers a new resource with the given unique name, arguments, and options.
func NewTemplateGroup(ctx *pulumi.Context,
	name string, args *TemplateGroupArgs, opts ...pulumi.ResourceOption) (*TemplateGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateIds == nil {
		return nil, errors.New("invalid value for required argument 'TemplateIds'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TemplateGroup
	err := ctx.RegisterResource("tencentcloud:Address/templateGroup:TemplateGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateGroup gets an existing TemplateGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateGroupState, opts ...pulumi.ResourceOption) (*TemplateGroup, error) {
	var resource TemplateGroup
	err := ctx.ReadResource("tencentcloud:Address/templateGroup:TemplateGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateGroup resources.
type templateGroupState struct {
	// Name of the address template group.
	Name *string `pulumi:"name"`
	// Template ID list.
	TemplateIds []string `pulumi:"templateIds"`
}

type TemplateGroupState struct {
	// Name of the address template group.
	Name pulumi.StringPtrInput
	// Template ID list.
	TemplateIds pulumi.StringArrayInput
}

func (TemplateGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateGroupState)(nil)).Elem()
}

type templateGroupArgs struct {
	// Name of the address template group.
	Name *string `pulumi:"name"`
	// Template ID list.
	TemplateIds []string `pulumi:"templateIds"`
}

// The set of arguments for constructing a TemplateGroup resource.
type TemplateGroupArgs struct {
	// Name of the address template group.
	Name pulumi.StringPtrInput
	// Template ID list.
	TemplateIds pulumi.StringArrayInput
}

func (TemplateGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateGroupArgs)(nil)).Elem()
}

type TemplateGroupInput interface {
	pulumi.Input

	ToTemplateGroupOutput() TemplateGroupOutput
	ToTemplateGroupOutputWithContext(ctx context.Context) TemplateGroupOutput
}

func (*TemplateGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateGroup)(nil)).Elem()
}

func (i *TemplateGroup) ToTemplateGroupOutput() TemplateGroupOutput {
	return i.ToTemplateGroupOutputWithContext(context.Background())
}

func (i *TemplateGroup) ToTemplateGroupOutputWithContext(ctx context.Context) TemplateGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateGroupOutput)
}

// TemplateGroupArrayInput is an input type that accepts TemplateGroupArray and TemplateGroupArrayOutput values.
// You can construct a concrete instance of `TemplateGroupArrayInput` via:
//
//          TemplateGroupArray{ TemplateGroupArgs{...} }
type TemplateGroupArrayInput interface {
	pulumi.Input

	ToTemplateGroupArrayOutput() TemplateGroupArrayOutput
	ToTemplateGroupArrayOutputWithContext(context.Context) TemplateGroupArrayOutput
}

type TemplateGroupArray []TemplateGroupInput

func (TemplateGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateGroup)(nil)).Elem()
}

func (i TemplateGroupArray) ToTemplateGroupArrayOutput() TemplateGroupArrayOutput {
	return i.ToTemplateGroupArrayOutputWithContext(context.Background())
}

func (i TemplateGroupArray) ToTemplateGroupArrayOutputWithContext(ctx context.Context) TemplateGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateGroupArrayOutput)
}

// TemplateGroupMapInput is an input type that accepts TemplateGroupMap and TemplateGroupMapOutput values.
// You can construct a concrete instance of `TemplateGroupMapInput` via:
//
//          TemplateGroupMap{ "key": TemplateGroupArgs{...} }
type TemplateGroupMapInput interface {
	pulumi.Input

	ToTemplateGroupMapOutput() TemplateGroupMapOutput
	ToTemplateGroupMapOutputWithContext(context.Context) TemplateGroupMapOutput
}

type TemplateGroupMap map[string]TemplateGroupInput

func (TemplateGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateGroup)(nil)).Elem()
}

func (i TemplateGroupMap) ToTemplateGroupMapOutput() TemplateGroupMapOutput {
	return i.ToTemplateGroupMapOutputWithContext(context.Background())
}

func (i TemplateGroupMap) ToTemplateGroupMapOutputWithContext(ctx context.Context) TemplateGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateGroupMapOutput)
}

type TemplateGroupOutput struct{ *pulumi.OutputState }

func (TemplateGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateGroup)(nil)).Elem()
}

func (o TemplateGroupOutput) ToTemplateGroupOutput() TemplateGroupOutput {
	return o
}

func (o TemplateGroupOutput) ToTemplateGroupOutputWithContext(ctx context.Context) TemplateGroupOutput {
	return o
}

// Name of the address template group.
func (o TemplateGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Template ID list.
func (o TemplateGroupOutput) TemplateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TemplateGroup) pulumi.StringArrayOutput { return v.TemplateIds }).(pulumi.StringArrayOutput)
}

type TemplateGroupArrayOutput struct{ *pulumi.OutputState }

func (TemplateGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateGroup)(nil)).Elem()
}

func (o TemplateGroupArrayOutput) ToTemplateGroupArrayOutput() TemplateGroupArrayOutput {
	return o
}

func (o TemplateGroupArrayOutput) ToTemplateGroupArrayOutputWithContext(ctx context.Context) TemplateGroupArrayOutput {
	return o
}

func (o TemplateGroupArrayOutput) Index(i pulumi.IntInput) TemplateGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemplateGroup {
		return vs[0].([]*TemplateGroup)[vs[1].(int)]
	}).(TemplateGroupOutput)
}

type TemplateGroupMapOutput struct{ *pulumi.OutputState }

func (TemplateGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateGroup)(nil)).Elem()
}

func (o TemplateGroupMapOutput) ToTemplateGroupMapOutput() TemplateGroupMapOutput {
	return o
}

func (o TemplateGroupMapOutput) ToTemplateGroupMapOutputWithContext(ctx context.Context) TemplateGroupMapOutput {
	return o
}

func (o TemplateGroupMapOutput) MapIndex(k pulumi.StringInput) TemplateGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemplateGroup {
		return vs[0].(map[string]*TemplateGroup)[vs[1].(string)]
	}).(TemplateGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateGroupInput)(nil)).Elem(), &TemplateGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateGroupArrayInput)(nil)).Elem(), TemplateGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateGroupMapInput)(nil)).Elem(), TemplateGroupMap{})
	pulumi.RegisterOutputType(TemplateGroupOutput{})
	pulumi.RegisterOutputType(TemplateGroupArrayOutput{})
	pulumi.RegisterOutputType(TemplateGroupMapOutput{})
}
