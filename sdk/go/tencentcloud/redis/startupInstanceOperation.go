// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a redis startupInstanceOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Redis.NewStartupInstanceOperation(ctx, "startupInstanceOperation", &Redis.StartupInstanceOperationArgs{
// 			InstanceId: pulumi.String("crs-c1nl9rpv"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type StartupInstanceOperation struct {
	pulumi.CustomResourceState

	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewStartupInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewStartupInstanceOperation(ctx *pulumi.Context,
	name string, args *StartupInstanceOperationArgs, opts ...pulumi.ResourceOption) (*StartupInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StartupInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Redis/startupInstanceOperation:StartupInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStartupInstanceOperation gets an existing StartupInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStartupInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StartupInstanceOperationState, opts ...pulumi.ResourceOption) (*StartupInstanceOperation, error) {
	var resource StartupInstanceOperation
	err := ctx.ReadResource("tencentcloud:Redis/startupInstanceOperation:StartupInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StartupInstanceOperation resources.
type startupInstanceOperationState struct {
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
}

type StartupInstanceOperationState struct {
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
}

func (StartupInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*startupInstanceOperationState)(nil)).Elem()
}

type startupInstanceOperationArgs struct {
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a StartupInstanceOperation resource.
type StartupInstanceOperationArgs struct {
	// The ID of instance.
	InstanceId pulumi.StringInput
}

func (StartupInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startupInstanceOperationArgs)(nil)).Elem()
}

type StartupInstanceOperationInput interface {
	pulumi.Input

	ToStartupInstanceOperationOutput() StartupInstanceOperationOutput
	ToStartupInstanceOperationOutputWithContext(ctx context.Context) StartupInstanceOperationOutput
}

func (*StartupInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**StartupInstanceOperation)(nil)).Elem()
}

func (i *StartupInstanceOperation) ToStartupInstanceOperationOutput() StartupInstanceOperationOutput {
	return i.ToStartupInstanceOperationOutputWithContext(context.Background())
}

func (i *StartupInstanceOperation) ToStartupInstanceOperationOutputWithContext(ctx context.Context) StartupInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartupInstanceOperationOutput)
}

// StartupInstanceOperationArrayInput is an input type that accepts StartupInstanceOperationArray and StartupInstanceOperationArrayOutput values.
// You can construct a concrete instance of `StartupInstanceOperationArrayInput` via:
//
//          StartupInstanceOperationArray{ StartupInstanceOperationArgs{...} }
type StartupInstanceOperationArrayInput interface {
	pulumi.Input

	ToStartupInstanceOperationArrayOutput() StartupInstanceOperationArrayOutput
	ToStartupInstanceOperationArrayOutputWithContext(context.Context) StartupInstanceOperationArrayOutput
}

type StartupInstanceOperationArray []StartupInstanceOperationInput

func (StartupInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartupInstanceOperation)(nil)).Elem()
}

func (i StartupInstanceOperationArray) ToStartupInstanceOperationArrayOutput() StartupInstanceOperationArrayOutput {
	return i.ToStartupInstanceOperationArrayOutputWithContext(context.Background())
}

func (i StartupInstanceOperationArray) ToStartupInstanceOperationArrayOutputWithContext(ctx context.Context) StartupInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartupInstanceOperationArrayOutput)
}

// StartupInstanceOperationMapInput is an input type that accepts StartupInstanceOperationMap and StartupInstanceOperationMapOutput values.
// You can construct a concrete instance of `StartupInstanceOperationMapInput` via:
//
//          StartupInstanceOperationMap{ "key": StartupInstanceOperationArgs{...} }
type StartupInstanceOperationMapInput interface {
	pulumi.Input

	ToStartupInstanceOperationMapOutput() StartupInstanceOperationMapOutput
	ToStartupInstanceOperationMapOutputWithContext(context.Context) StartupInstanceOperationMapOutput
}

type StartupInstanceOperationMap map[string]StartupInstanceOperationInput

func (StartupInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartupInstanceOperation)(nil)).Elem()
}

func (i StartupInstanceOperationMap) ToStartupInstanceOperationMapOutput() StartupInstanceOperationMapOutput {
	return i.ToStartupInstanceOperationMapOutputWithContext(context.Background())
}

func (i StartupInstanceOperationMap) ToStartupInstanceOperationMapOutputWithContext(ctx context.Context) StartupInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartupInstanceOperationMapOutput)
}

type StartupInstanceOperationOutput struct{ *pulumi.OutputState }

func (StartupInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartupInstanceOperation)(nil)).Elem()
}

func (o StartupInstanceOperationOutput) ToStartupInstanceOperationOutput() StartupInstanceOperationOutput {
	return o
}

func (o StartupInstanceOperationOutput) ToStartupInstanceOperationOutputWithContext(ctx context.Context) StartupInstanceOperationOutput {
	return o
}

// The ID of instance.
func (o StartupInstanceOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *StartupInstanceOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type StartupInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (StartupInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartupInstanceOperation)(nil)).Elem()
}

func (o StartupInstanceOperationArrayOutput) ToStartupInstanceOperationArrayOutput() StartupInstanceOperationArrayOutput {
	return o
}

func (o StartupInstanceOperationArrayOutput) ToStartupInstanceOperationArrayOutputWithContext(ctx context.Context) StartupInstanceOperationArrayOutput {
	return o
}

func (o StartupInstanceOperationArrayOutput) Index(i pulumi.IntInput) StartupInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StartupInstanceOperation {
		return vs[0].([]*StartupInstanceOperation)[vs[1].(int)]
	}).(StartupInstanceOperationOutput)
}

type StartupInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (StartupInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartupInstanceOperation)(nil)).Elem()
}

func (o StartupInstanceOperationMapOutput) ToStartupInstanceOperationMapOutput() StartupInstanceOperationMapOutput {
	return o
}

func (o StartupInstanceOperationMapOutput) ToStartupInstanceOperationMapOutputWithContext(ctx context.Context) StartupInstanceOperationMapOutput {
	return o
}

func (o StartupInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) StartupInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StartupInstanceOperation {
		return vs[0].(map[string]*StartupInstanceOperation)[vs[1].(string)]
	}).(StartupInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StartupInstanceOperationInput)(nil)).Elem(), &StartupInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartupInstanceOperationArrayInput)(nil)).Elem(), StartupInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartupInstanceOperationMapInput)(nil)).Elem(), StartupInstanceOperationMap{})
	pulumi.RegisterOutputType(StartupInstanceOperationOutput{})
	pulumi.RegisterOutputType(StartupInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(StartupInstanceOperationMapOutput{})
}
