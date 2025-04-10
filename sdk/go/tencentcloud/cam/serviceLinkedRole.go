// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cam serviceLinkedRole
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cam.NewServiceLinkedRole(ctx, "serviceLinkedRole", &Cam.ServiceLinkedRoleArgs{
// 			CustomSuffix: pulumi.String("x-1"),
// 			Description:  pulumi.String("desc cam"),
// 			QcsServiceNames: pulumi.StringArray{
// 				pulumi.String("cvm.qcloud.com"),
// 				pulumi.String("ekslog.tke.cloud.tencent.com"),
// 			},
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
type ServiceLinkedRole struct {
	pulumi.CustomResourceState

	// The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name.
	CustomSuffix pulumi.StringPtrOutput `pulumi:"customSuffix"`
	// role description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Authorization service, the Tencent Cloud service principal with this role attached.
	QcsServiceNames pulumi.StringArrayOutput `pulumi:"qcsServiceNames"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewServiceLinkedRole registers a new resource with the given unique name, arguments, and options.
func NewServiceLinkedRole(ctx *pulumi.Context,
	name string, args *ServiceLinkedRoleArgs, opts ...pulumi.ResourceOption) (*ServiceLinkedRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QcsServiceNames == nil {
		return nil, errors.New("invalid value for required argument 'QcsServiceNames'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServiceLinkedRole
	err := ctx.RegisterResource("tencentcloud:Cam/serviceLinkedRole:ServiceLinkedRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceLinkedRole gets an existing ServiceLinkedRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceLinkedRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceLinkedRoleState, opts ...pulumi.ResourceOption) (*ServiceLinkedRole, error) {
	var resource ServiceLinkedRole
	err := ctx.ReadResource("tencentcloud:Cam/serviceLinkedRole:ServiceLinkedRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceLinkedRole resources.
type serviceLinkedRoleState struct {
	// The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name.
	CustomSuffix *string `pulumi:"customSuffix"`
	// role description.
	Description *string `pulumi:"description"`
	// Authorization service, the Tencent Cloud service principal with this role attached.
	QcsServiceNames []string `pulumi:"qcsServiceNames"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ServiceLinkedRoleState struct {
	// The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name.
	CustomSuffix pulumi.StringPtrInput
	// role description.
	Description pulumi.StringPtrInput
	// Authorization service, the Tencent Cloud service principal with this role attached.
	QcsServiceNames pulumi.StringArrayInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (ServiceLinkedRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceLinkedRoleState)(nil)).Elem()
}

type serviceLinkedRoleArgs struct {
	// The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name.
	CustomSuffix *string `pulumi:"customSuffix"`
	// role description.
	Description *string `pulumi:"description"`
	// Authorization service, the Tencent Cloud service principal with this role attached.
	QcsServiceNames []string `pulumi:"qcsServiceNames"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceLinkedRole resource.
type ServiceLinkedRoleArgs struct {
	// The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name.
	CustomSuffix pulumi.StringPtrInput
	// role description.
	Description pulumi.StringPtrInput
	// Authorization service, the Tencent Cloud service principal with this role attached.
	QcsServiceNames pulumi.StringArrayInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (ServiceLinkedRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceLinkedRoleArgs)(nil)).Elem()
}

type ServiceLinkedRoleInput interface {
	pulumi.Input

	ToServiceLinkedRoleOutput() ServiceLinkedRoleOutput
	ToServiceLinkedRoleOutputWithContext(ctx context.Context) ServiceLinkedRoleOutput
}

func (*ServiceLinkedRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLinkedRole)(nil)).Elem()
}

func (i *ServiceLinkedRole) ToServiceLinkedRoleOutput() ServiceLinkedRoleOutput {
	return i.ToServiceLinkedRoleOutputWithContext(context.Background())
}

func (i *ServiceLinkedRole) ToServiceLinkedRoleOutputWithContext(ctx context.Context) ServiceLinkedRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLinkedRoleOutput)
}

// ServiceLinkedRoleArrayInput is an input type that accepts ServiceLinkedRoleArray and ServiceLinkedRoleArrayOutput values.
// You can construct a concrete instance of `ServiceLinkedRoleArrayInput` via:
//
//          ServiceLinkedRoleArray{ ServiceLinkedRoleArgs{...} }
type ServiceLinkedRoleArrayInput interface {
	pulumi.Input

	ToServiceLinkedRoleArrayOutput() ServiceLinkedRoleArrayOutput
	ToServiceLinkedRoleArrayOutputWithContext(context.Context) ServiceLinkedRoleArrayOutput
}

type ServiceLinkedRoleArray []ServiceLinkedRoleInput

func (ServiceLinkedRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceLinkedRole)(nil)).Elem()
}

func (i ServiceLinkedRoleArray) ToServiceLinkedRoleArrayOutput() ServiceLinkedRoleArrayOutput {
	return i.ToServiceLinkedRoleArrayOutputWithContext(context.Background())
}

func (i ServiceLinkedRoleArray) ToServiceLinkedRoleArrayOutputWithContext(ctx context.Context) ServiceLinkedRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLinkedRoleArrayOutput)
}

// ServiceLinkedRoleMapInput is an input type that accepts ServiceLinkedRoleMap and ServiceLinkedRoleMapOutput values.
// You can construct a concrete instance of `ServiceLinkedRoleMapInput` via:
//
//          ServiceLinkedRoleMap{ "key": ServiceLinkedRoleArgs{...} }
type ServiceLinkedRoleMapInput interface {
	pulumi.Input

	ToServiceLinkedRoleMapOutput() ServiceLinkedRoleMapOutput
	ToServiceLinkedRoleMapOutputWithContext(context.Context) ServiceLinkedRoleMapOutput
}

type ServiceLinkedRoleMap map[string]ServiceLinkedRoleInput

func (ServiceLinkedRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceLinkedRole)(nil)).Elem()
}

func (i ServiceLinkedRoleMap) ToServiceLinkedRoleMapOutput() ServiceLinkedRoleMapOutput {
	return i.ToServiceLinkedRoleMapOutputWithContext(context.Background())
}

func (i ServiceLinkedRoleMap) ToServiceLinkedRoleMapOutputWithContext(ctx context.Context) ServiceLinkedRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLinkedRoleMapOutput)
}

type ServiceLinkedRoleOutput struct{ *pulumi.OutputState }

func (ServiceLinkedRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLinkedRole)(nil)).Elem()
}

func (o ServiceLinkedRoleOutput) ToServiceLinkedRoleOutput() ServiceLinkedRoleOutput {
	return o
}

func (o ServiceLinkedRoleOutput) ToServiceLinkedRoleOutputWithContext(ctx context.Context) ServiceLinkedRoleOutput {
	return o
}

// The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name.
func (o ServiceLinkedRoleOutput) CustomSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceLinkedRole) pulumi.StringPtrOutput { return v.CustomSuffix }).(pulumi.StringPtrOutput)
}

// role description.
func (o ServiceLinkedRoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceLinkedRole) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Authorization service, the Tencent Cloud service principal with this role attached.
func (o ServiceLinkedRoleOutput) QcsServiceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceLinkedRole) pulumi.StringArrayOutput { return v.QcsServiceNames }).(pulumi.StringArrayOutput)
}

// Tag description list.
func (o ServiceLinkedRoleOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServiceLinkedRole) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type ServiceLinkedRoleArrayOutput struct{ *pulumi.OutputState }

func (ServiceLinkedRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceLinkedRole)(nil)).Elem()
}

func (o ServiceLinkedRoleArrayOutput) ToServiceLinkedRoleArrayOutput() ServiceLinkedRoleArrayOutput {
	return o
}

func (o ServiceLinkedRoleArrayOutput) ToServiceLinkedRoleArrayOutputWithContext(ctx context.Context) ServiceLinkedRoleArrayOutput {
	return o
}

func (o ServiceLinkedRoleArrayOutput) Index(i pulumi.IntInput) ServiceLinkedRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceLinkedRole {
		return vs[0].([]*ServiceLinkedRole)[vs[1].(int)]
	}).(ServiceLinkedRoleOutput)
}

type ServiceLinkedRoleMapOutput struct{ *pulumi.OutputState }

func (ServiceLinkedRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceLinkedRole)(nil)).Elem()
}

func (o ServiceLinkedRoleMapOutput) ToServiceLinkedRoleMapOutput() ServiceLinkedRoleMapOutput {
	return o
}

func (o ServiceLinkedRoleMapOutput) ToServiceLinkedRoleMapOutputWithContext(ctx context.Context) ServiceLinkedRoleMapOutput {
	return o
}

func (o ServiceLinkedRoleMapOutput) MapIndex(k pulumi.StringInput) ServiceLinkedRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceLinkedRole {
		return vs[0].(map[string]*ServiceLinkedRole)[vs[1].(string)]
	}).(ServiceLinkedRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLinkedRoleInput)(nil)).Elem(), &ServiceLinkedRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLinkedRoleArrayInput)(nil)).Elem(), ServiceLinkedRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLinkedRoleMapInput)(nil)).Elem(), ServiceLinkedRoleMap{})
	pulumi.RegisterOutputType(ServiceLinkedRoleOutput{})
	pulumi.RegisterOutputType(ServiceLinkedRoleArrayOutput{})
	pulumi.RegisterOutputType(ServiceLinkedRoleMapOutput{})
}
