// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tsf microservice
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tsf.NewMicroservice(ctx, "microservice", &Tsf.MicroserviceArgs{
// 			MicroserviceDesc: pulumi.String("desc-microservice"),
// 			MicroserviceName: pulumi.String("test-microservice"),
// 			NamespaceId:      pulumi.String("namespace-vjlkzkgy"),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
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
// tsf microservice can be imported using the namespaceId#microserviceId, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tsf/microservice:Microservice microservice namespace-vjlkzkgy#ms-vjeb43lw
// ```
type Microservice struct {
	pulumi.CustomResourceState

	// Microservice description information.
	MicroserviceDesc pulumi.StringPtrOutput `pulumi:"microserviceDesc"`
	// Microservice name.
	MicroserviceName pulumi.StringOutput `pulumi:"microserviceName"`
	// Namespace ID.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewMicroservice registers a new resource with the given unique name, arguments, and options.
func NewMicroservice(ctx *pulumi.Context,
	name string, args *MicroserviceArgs, opts ...pulumi.ResourceOption) (*Microservice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MicroserviceName == nil {
		return nil, errors.New("invalid value for required argument 'MicroserviceName'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Microservice
	err := ctx.RegisterResource("tencentcloud:Tsf/microservice:Microservice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMicroservice gets an existing Microservice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMicroservice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicroserviceState, opts ...pulumi.ResourceOption) (*Microservice, error) {
	var resource Microservice
	err := ctx.ReadResource("tencentcloud:Tsf/microservice:Microservice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Microservice resources.
type microserviceState struct {
	// Microservice description information.
	MicroserviceDesc *string `pulumi:"microserviceDesc"`
	// Microservice name.
	MicroserviceName *string `pulumi:"microserviceName"`
	// Namespace ID.
	NamespaceId *string `pulumi:"namespaceId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

type MicroserviceState struct {
	// Microservice description information.
	MicroserviceDesc pulumi.StringPtrInput
	// Microservice name.
	MicroserviceName pulumi.StringPtrInput
	// Namespace ID.
	NamespaceId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (MicroserviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*microserviceState)(nil)).Elem()
}

type microserviceArgs struct {
	// Microservice description information.
	MicroserviceDesc *string `pulumi:"microserviceDesc"`
	// Microservice name.
	MicroserviceName string `pulumi:"microserviceName"`
	// Namespace ID.
	NamespaceId string `pulumi:"namespaceId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Microservice resource.
type MicroserviceArgs struct {
	// Microservice description information.
	MicroserviceDesc pulumi.StringPtrInput
	// Microservice name.
	MicroserviceName pulumi.StringInput
	// Namespace ID.
	NamespaceId pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (MicroserviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microserviceArgs)(nil)).Elem()
}

type MicroserviceInput interface {
	pulumi.Input

	ToMicroserviceOutput() MicroserviceOutput
	ToMicroserviceOutputWithContext(ctx context.Context) MicroserviceOutput
}

func (*Microservice) ElementType() reflect.Type {
	return reflect.TypeOf((**Microservice)(nil)).Elem()
}

func (i *Microservice) ToMicroserviceOutput() MicroserviceOutput {
	return i.ToMicroserviceOutputWithContext(context.Background())
}

func (i *Microservice) ToMicroserviceOutputWithContext(ctx context.Context) MicroserviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicroserviceOutput)
}

// MicroserviceArrayInput is an input type that accepts MicroserviceArray and MicroserviceArrayOutput values.
// You can construct a concrete instance of `MicroserviceArrayInput` via:
//
//          MicroserviceArray{ MicroserviceArgs{...} }
type MicroserviceArrayInput interface {
	pulumi.Input

	ToMicroserviceArrayOutput() MicroserviceArrayOutput
	ToMicroserviceArrayOutputWithContext(context.Context) MicroserviceArrayOutput
}

type MicroserviceArray []MicroserviceInput

func (MicroserviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Microservice)(nil)).Elem()
}

func (i MicroserviceArray) ToMicroserviceArrayOutput() MicroserviceArrayOutput {
	return i.ToMicroserviceArrayOutputWithContext(context.Background())
}

func (i MicroserviceArray) ToMicroserviceArrayOutputWithContext(ctx context.Context) MicroserviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicroserviceArrayOutput)
}

// MicroserviceMapInput is an input type that accepts MicroserviceMap and MicroserviceMapOutput values.
// You can construct a concrete instance of `MicroserviceMapInput` via:
//
//          MicroserviceMap{ "key": MicroserviceArgs{...} }
type MicroserviceMapInput interface {
	pulumi.Input

	ToMicroserviceMapOutput() MicroserviceMapOutput
	ToMicroserviceMapOutputWithContext(context.Context) MicroserviceMapOutput
}

type MicroserviceMap map[string]MicroserviceInput

func (MicroserviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Microservice)(nil)).Elem()
}

func (i MicroserviceMap) ToMicroserviceMapOutput() MicroserviceMapOutput {
	return i.ToMicroserviceMapOutputWithContext(context.Background())
}

func (i MicroserviceMap) ToMicroserviceMapOutputWithContext(ctx context.Context) MicroserviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicroserviceMapOutput)
}

type MicroserviceOutput struct{ *pulumi.OutputState }

func (MicroserviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Microservice)(nil)).Elem()
}

func (o MicroserviceOutput) ToMicroserviceOutput() MicroserviceOutput {
	return o
}

func (o MicroserviceOutput) ToMicroserviceOutputWithContext(ctx context.Context) MicroserviceOutput {
	return o
}

// Microservice description information.
func (o MicroserviceOutput) MicroserviceDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Microservice) pulumi.StringPtrOutput { return v.MicroserviceDesc }).(pulumi.StringPtrOutput)
}

// Microservice name.
func (o MicroserviceOutput) MicroserviceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Microservice) pulumi.StringOutput { return v.MicroserviceName }).(pulumi.StringOutput)
}

// Namespace ID.
func (o MicroserviceOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Microservice) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// Tag description list.
func (o MicroserviceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Microservice) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type MicroserviceArrayOutput struct{ *pulumi.OutputState }

func (MicroserviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Microservice)(nil)).Elem()
}

func (o MicroserviceArrayOutput) ToMicroserviceArrayOutput() MicroserviceArrayOutput {
	return o
}

func (o MicroserviceArrayOutput) ToMicroserviceArrayOutputWithContext(ctx context.Context) MicroserviceArrayOutput {
	return o
}

func (o MicroserviceArrayOutput) Index(i pulumi.IntInput) MicroserviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Microservice {
		return vs[0].([]*Microservice)[vs[1].(int)]
	}).(MicroserviceOutput)
}

type MicroserviceMapOutput struct{ *pulumi.OutputState }

func (MicroserviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Microservice)(nil)).Elem()
}

func (o MicroserviceMapOutput) ToMicroserviceMapOutput() MicroserviceMapOutput {
	return o
}

func (o MicroserviceMapOutput) ToMicroserviceMapOutputWithContext(ctx context.Context) MicroserviceMapOutput {
	return o
}

func (o MicroserviceMapOutput) MapIndex(k pulumi.StringInput) MicroserviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Microservice {
		return vs[0].(map[string]*Microservice)[vs[1].(string)]
	}).(MicroserviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MicroserviceInput)(nil)).Elem(), &Microservice{})
	pulumi.RegisterInputType(reflect.TypeOf((*MicroserviceArrayInput)(nil)).Elem(), MicroserviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MicroserviceMapInput)(nil)).Elem(), MicroserviceMap{})
	pulumi.RegisterOutputType(MicroserviceOutput{})
	pulumi.RegisterOutputType(MicroserviceArrayOutput{})
	pulumi.RegisterOutputType(MicroserviceMapOutput{})
}
