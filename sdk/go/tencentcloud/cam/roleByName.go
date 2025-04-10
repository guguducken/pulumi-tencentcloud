// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleByName struct {
	pulumi.CustomResourceState

	// Indicates whether the CAM role can login or not.
	ConsoleLogin pulumi.BoolPtrOutput `pulumi:"consoleLogin"`
	// Create time of the CAM role.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the CAM role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document pulumi.StringOutput `pulumi:"document"`
	// Name of CAM role.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of tags used to associate different resources.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The last update time of the CAM role.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRoleByName registers a new resource with the given unique name, arguments, and options.
func NewRoleByName(ctx *pulumi.Context,
	name string, args *RoleByNameArgs, opts ...pulumi.ResourceOption) (*RoleByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Document == nil {
		return nil, errors.New("invalid value for required argument 'Document'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RoleByName
	err := ctx.RegisterResource("tencentcloud:Cam/roleByName:RoleByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleByName gets an existing RoleByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleByNameState, opts ...pulumi.ResourceOption) (*RoleByName, error) {
	var resource RoleByName
	err := ctx.ReadResource("tencentcloud:Cam/roleByName:RoleByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleByName resources.
type roleByNameState struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin *bool `pulumi:"consoleLogin"`
	// Create time of the CAM role.
	CreateTime *string `pulumi:"createTime"`
	// Description of the CAM role.
	Description *string `pulumi:"description"`
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document *string `pulumi:"document"`
	// Name of CAM role.
	Name *string `pulumi:"name"`
	// A list of tags used to associate different resources.
	Tags map[string]interface{} `pulumi:"tags"`
	// The last update time of the CAM role.
	UpdateTime *string `pulumi:"updateTime"`
}

type RoleByNameState struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin pulumi.BoolPtrInput
	// Create time of the CAM role.
	CreateTime pulumi.StringPtrInput
	// Description of the CAM role.
	Description pulumi.StringPtrInput
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document pulumi.StringPtrInput
	// Name of CAM role.
	Name pulumi.StringPtrInput
	// A list of tags used to associate different resources.
	Tags pulumi.MapInput
	// The last update time of the CAM role.
	UpdateTime pulumi.StringPtrInput
}

func (RoleByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleByNameState)(nil)).Elem()
}

type roleByNameArgs struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin *bool `pulumi:"consoleLogin"`
	// Description of the CAM role.
	Description *string `pulumi:"description"`
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document string `pulumi:"document"`
	// Name of CAM role.
	Name *string `pulumi:"name"`
	// A list of tags used to associate different resources.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a RoleByName resource.
type RoleByNameArgs struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin pulumi.BoolPtrInput
	// Description of the CAM role.
	Description pulumi.StringPtrInput
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document pulumi.StringInput
	// Name of CAM role.
	Name pulumi.StringPtrInput
	// A list of tags used to associate different resources.
	Tags pulumi.MapInput
}

func (RoleByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleByNameArgs)(nil)).Elem()
}

type RoleByNameInput interface {
	pulumi.Input

	ToRoleByNameOutput() RoleByNameOutput
	ToRoleByNameOutputWithContext(ctx context.Context) RoleByNameOutput
}

func (*RoleByName) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleByName)(nil)).Elem()
}

func (i *RoleByName) ToRoleByNameOutput() RoleByNameOutput {
	return i.ToRoleByNameOutputWithContext(context.Background())
}

func (i *RoleByName) ToRoleByNameOutputWithContext(ctx context.Context) RoleByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleByNameOutput)
}

// RoleByNameArrayInput is an input type that accepts RoleByNameArray and RoleByNameArrayOutput values.
// You can construct a concrete instance of `RoleByNameArrayInput` via:
//
//          RoleByNameArray{ RoleByNameArgs{...} }
type RoleByNameArrayInput interface {
	pulumi.Input

	ToRoleByNameArrayOutput() RoleByNameArrayOutput
	ToRoleByNameArrayOutputWithContext(context.Context) RoleByNameArrayOutput
}

type RoleByNameArray []RoleByNameInput

func (RoleByNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleByName)(nil)).Elem()
}

func (i RoleByNameArray) ToRoleByNameArrayOutput() RoleByNameArrayOutput {
	return i.ToRoleByNameArrayOutputWithContext(context.Background())
}

func (i RoleByNameArray) ToRoleByNameArrayOutputWithContext(ctx context.Context) RoleByNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleByNameArrayOutput)
}

// RoleByNameMapInput is an input type that accepts RoleByNameMap and RoleByNameMapOutput values.
// You can construct a concrete instance of `RoleByNameMapInput` via:
//
//          RoleByNameMap{ "key": RoleByNameArgs{...} }
type RoleByNameMapInput interface {
	pulumi.Input

	ToRoleByNameMapOutput() RoleByNameMapOutput
	ToRoleByNameMapOutputWithContext(context.Context) RoleByNameMapOutput
}

type RoleByNameMap map[string]RoleByNameInput

func (RoleByNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleByName)(nil)).Elem()
}

func (i RoleByNameMap) ToRoleByNameMapOutput() RoleByNameMapOutput {
	return i.ToRoleByNameMapOutputWithContext(context.Background())
}

func (i RoleByNameMap) ToRoleByNameMapOutputWithContext(ctx context.Context) RoleByNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleByNameMapOutput)
}

type RoleByNameOutput struct{ *pulumi.OutputState }

func (RoleByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleByName)(nil)).Elem()
}

func (o RoleByNameOutput) ToRoleByNameOutput() RoleByNameOutput {
	return o
}

func (o RoleByNameOutput) ToRoleByNameOutputWithContext(ctx context.Context) RoleByNameOutput {
	return o
}

// Indicates whether the CAM role can login or not.
func (o RoleByNameOutput) ConsoleLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.BoolPtrOutput { return v.ConsoleLogin }).(pulumi.BoolPtrOutput)
}

// Create time of the CAM role.
func (o RoleByNameOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the CAM role.
func (o RoleByNameOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
// replaced with the uin it stands for.
func (o RoleByNameOutput) Document() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.StringOutput { return v.Document }).(pulumi.StringOutput)
}

// Name of CAM role.
func (o RoleByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of tags used to associate different resources.
func (o RoleByNameOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The last update time of the CAM role.
func (o RoleByNameOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleByName) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type RoleByNameArrayOutput struct{ *pulumi.OutputState }

func (RoleByNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleByName)(nil)).Elem()
}

func (o RoleByNameArrayOutput) ToRoleByNameArrayOutput() RoleByNameArrayOutput {
	return o
}

func (o RoleByNameArrayOutput) ToRoleByNameArrayOutputWithContext(ctx context.Context) RoleByNameArrayOutput {
	return o
}

func (o RoleByNameArrayOutput) Index(i pulumi.IntInput) RoleByNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleByName {
		return vs[0].([]*RoleByName)[vs[1].(int)]
	}).(RoleByNameOutput)
}

type RoleByNameMapOutput struct{ *pulumi.OutputState }

func (RoleByNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleByName)(nil)).Elem()
}

func (o RoleByNameMapOutput) ToRoleByNameMapOutput() RoleByNameMapOutput {
	return o
}

func (o RoleByNameMapOutput) ToRoleByNameMapOutputWithContext(ctx context.Context) RoleByNameMapOutput {
	return o
}

func (o RoleByNameMapOutput) MapIndex(k pulumi.StringInput) RoleByNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleByName {
		return vs[0].(map[string]*RoleByName)[vs[1].(string)]
	}).(RoleByNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleByNameInput)(nil)).Elem(), &RoleByName{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleByNameArrayInput)(nil)).Elem(), RoleByNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleByNameMapInput)(nil)).Elem(), RoleByNameMap{})
	pulumi.RegisterOutputType(RoleByNameOutput{})
	pulumi.RegisterOutputType(RoleByNameArrayOutput{})
	pulumi.RegisterOutputType(RoleByNameMapOutput{})
}
