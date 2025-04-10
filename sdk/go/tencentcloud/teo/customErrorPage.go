// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a teo customErrorPage
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Teo.NewCustomErrorPage(ctx, "errorPage0", &Teo.CustomErrorPageArgs{
// 			ZoneId:  pulumi.Any(data.Tencentcloud_teo_zone_ddos_policy.Zone_policy.Zone_id),
// 			Entity:  pulumi.Any(data.Tencentcloud_teo_zone_ddos_policy.Zone_policy.Shield_areas[0].Application[0].Host),
// 			Content: pulumi.String("<html lang='en'><body><div><p>test content</p></div></body></html>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CustomErrorPage struct {
	pulumi.CustomResourceState

	// Page content.
	Content pulumi.StringOutput `pulumi:"content"`
	// Subdomain.
	Entity pulumi.StringOutput `pulumi:"entity"`
	// Page name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Page ID.
	PageId pulumi.StringOutput `pulumi:"pageId"`
	// Site ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCustomErrorPage registers a new resource with the given unique name, arguments, and options.
func NewCustomErrorPage(ctx *pulumi.Context,
	name string, args *CustomErrorPageArgs, opts ...pulumi.ResourceOption) (*CustomErrorPage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CustomErrorPage
	err := ctx.RegisterResource("tencentcloud:Teo/customErrorPage:CustomErrorPage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomErrorPage gets an existing CustomErrorPage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomErrorPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomErrorPageState, opts ...pulumi.ResourceOption) (*CustomErrorPage, error) {
	var resource CustomErrorPage
	err := ctx.ReadResource("tencentcloud:Teo/customErrorPage:CustomErrorPage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomErrorPage resources.
type customErrorPageState struct {
	// Page content.
	Content *string `pulumi:"content"`
	// Subdomain.
	Entity *string `pulumi:"entity"`
	// Page name.
	Name *string `pulumi:"name"`
	// Page ID.
	PageId *string `pulumi:"pageId"`
	// Site ID.
	ZoneId *string `pulumi:"zoneId"`
}

type CustomErrorPageState struct {
	// Page content.
	Content pulumi.StringPtrInput
	// Subdomain.
	Entity pulumi.StringPtrInput
	// Page name.
	Name pulumi.StringPtrInput
	// Page ID.
	PageId pulumi.StringPtrInput
	// Site ID.
	ZoneId pulumi.StringPtrInput
}

func (CustomErrorPageState) ElementType() reflect.Type {
	return reflect.TypeOf((*customErrorPageState)(nil)).Elem()
}

type customErrorPageArgs struct {
	// Page content.
	Content string `pulumi:"content"`
	// Subdomain.
	Entity string `pulumi:"entity"`
	// Page name.
	Name *string `pulumi:"name"`
	// Site ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a CustomErrorPage resource.
type CustomErrorPageArgs struct {
	// Page content.
	Content pulumi.StringInput
	// Subdomain.
	Entity pulumi.StringInput
	// Page name.
	Name pulumi.StringPtrInput
	// Site ID.
	ZoneId pulumi.StringInput
}

func (CustomErrorPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customErrorPageArgs)(nil)).Elem()
}

type CustomErrorPageInput interface {
	pulumi.Input

	ToCustomErrorPageOutput() CustomErrorPageOutput
	ToCustomErrorPageOutputWithContext(ctx context.Context) CustomErrorPageOutput
}

func (*CustomErrorPage) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomErrorPage)(nil)).Elem()
}

func (i *CustomErrorPage) ToCustomErrorPageOutput() CustomErrorPageOutput {
	return i.ToCustomErrorPageOutputWithContext(context.Background())
}

func (i *CustomErrorPage) ToCustomErrorPageOutputWithContext(ctx context.Context) CustomErrorPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomErrorPageOutput)
}

// CustomErrorPageArrayInput is an input type that accepts CustomErrorPageArray and CustomErrorPageArrayOutput values.
// You can construct a concrete instance of `CustomErrorPageArrayInput` via:
//
//          CustomErrorPageArray{ CustomErrorPageArgs{...} }
type CustomErrorPageArrayInput interface {
	pulumi.Input

	ToCustomErrorPageArrayOutput() CustomErrorPageArrayOutput
	ToCustomErrorPageArrayOutputWithContext(context.Context) CustomErrorPageArrayOutput
}

type CustomErrorPageArray []CustomErrorPageInput

func (CustomErrorPageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomErrorPage)(nil)).Elem()
}

func (i CustomErrorPageArray) ToCustomErrorPageArrayOutput() CustomErrorPageArrayOutput {
	return i.ToCustomErrorPageArrayOutputWithContext(context.Background())
}

func (i CustomErrorPageArray) ToCustomErrorPageArrayOutputWithContext(ctx context.Context) CustomErrorPageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomErrorPageArrayOutput)
}

// CustomErrorPageMapInput is an input type that accepts CustomErrorPageMap and CustomErrorPageMapOutput values.
// You can construct a concrete instance of `CustomErrorPageMapInput` via:
//
//          CustomErrorPageMap{ "key": CustomErrorPageArgs{...} }
type CustomErrorPageMapInput interface {
	pulumi.Input

	ToCustomErrorPageMapOutput() CustomErrorPageMapOutput
	ToCustomErrorPageMapOutputWithContext(context.Context) CustomErrorPageMapOutput
}

type CustomErrorPageMap map[string]CustomErrorPageInput

func (CustomErrorPageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomErrorPage)(nil)).Elem()
}

func (i CustomErrorPageMap) ToCustomErrorPageMapOutput() CustomErrorPageMapOutput {
	return i.ToCustomErrorPageMapOutputWithContext(context.Background())
}

func (i CustomErrorPageMap) ToCustomErrorPageMapOutputWithContext(ctx context.Context) CustomErrorPageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomErrorPageMapOutput)
}

type CustomErrorPageOutput struct{ *pulumi.OutputState }

func (CustomErrorPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomErrorPage)(nil)).Elem()
}

func (o CustomErrorPageOutput) ToCustomErrorPageOutput() CustomErrorPageOutput {
	return o
}

func (o CustomErrorPageOutput) ToCustomErrorPageOutputWithContext(ctx context.Context) CustomErrorPageOutput {
	return o
}

// Page content.
func (o CustomErrorPageOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomErrorPage) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Subdomain.
func (o CustomErrorPageOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomErrorPage) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// Page name.
func (o CustomErrorPageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomErrorPage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Page ID.
func (o CustomErrorPageOutput) PageId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomErrorPage) pulumi.StringOutput { return v.PageId }).(pulumi.StringOutput)
}

// Site ID.
func (o CustomErrorPageOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomErrorPage) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type CustomErrorPageArrayOutput struct{ *pulumi.OutputState }

func (CustomErrorPageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomErrorPage)(nil)).Elem()
}

func (o CustomErrorPageArrayOutput) ToCustomErrorPageArrayOutput() CustomErrorPageArrayOutput {
	return o
}

func (o CustomErrorPageArrayOutput) ToCustomErrorPageArrayOutputWithContext(ctx context.Context) CustomErrorPageArrayOutput {
	return o
}

func (o CustomErrorPageArrayOutput) Index(i pulumi.IntInput) CustomErrorPageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomErrorPage {
		return vs[0].([]*CustomErrorPage)[vs[1].(int)]
	}).(CustomErrorPageOutput)
}

type CustomErrorPageMapOutput struct{ *pulumi.OutputState }

func (CustomErrorPageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomErrorPage)(nil)).Elem()
}

func (o CustomErrorPageMapOutput) ToCustomErrorPageMapOutput() CustomErrorPageMapOutput {
	return o
}

func (o CustomErrorPageMapOutput) ToCustomErrorPageMapOutputWithContext(ctx context.Context) CustomErrorPageMapOutput {
	return o
}

func (o CustomErrorPageMapOutput) MapIndex(k pulumi.StringInput) CustomErrorPageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomErrorPage {
		return vs[0].(map[string]*CustomErrorPage)[vs[1].(string)]
	}).(CustomErrorPageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomErrorPageInput)(nil)).Elem(), &CustomErrorPage{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomErrorPageArrayInput)(nil)).Elem(), CustomErrorPageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomErrorPageMapInput)(nil)).Elem(), CustomErrorPageMap{})
	pulumi.RegisterOutputType(CustomErrorPageOutput{})
	pulumi.RegisterOutputType(CustomErrorPageArrayOutput{})
	pulumi.RegisterOutputType(CustomErrorPageMapOutput{})
}
