// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CAM role policy attachment.
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
// 		_, err := Cam.NewRolePolicyAttachment(ctx, "foo", &Cam.RolePolicyAttachmentArgs{
// 			RoleId:   pulumi.Any(tencentcloud_cam_role.Foo.Id),
// 			PolicyId: pulumi.Any(tencentcloud_cam_policy.Foo.Id),
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
// CAM role policy attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment foo 4611686018427922725#26800353
// ```
type RolePolicyAttachment struct {
	pulumi.CustomResourceState

	// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode pulumi.IntOutput `pulumi:"createMode"`
	// The create time of the CAM role policy attachment.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// ID of the policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The name of the policy.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// ID of the attached CAM role.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewRolePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewRolePolicyAttachment(ctx *pulumi.Context,
	name string, args *RolePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*RolePolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RolePolicyAttachment
	err := ctx.RegisterResource("tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRolePolicyAttachment gets an existing RolePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRolePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolePolicyAttachmentState, opts ...pulumi.ResourceOption) (*RolePolicyAttachment, error) {
	var resource RolePolicyAttachment
	err := ctx.ReadResource("tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RolePolicyAttachment resources.
type rolePolicyAttachmentState struct {
	// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *int `pulumi:"createMode"`
	// The create time of the CAM role policy attachment.
	CreateTime *string `pulumi:"createTime"`
	// ID of the policy.
	PolicyId *string `pulumi:"policyId"`
	// The name of the policy.
	PolicyName *string `pulumi:"policyName"`
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType *string `pulumi:"policyType"`
	// ID of the attached CAM role.
	RoleId *string `pulumi:"roleId"`
}

type RolePolicyAttachmentState struct {
	// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode pulumi.IntPtrInput
	// The create time of the CAM role policy attachment.
	CreateTime pulumi.StringPtrInput
	// ID of the policy.
	PolicyId pulumi.StringPtrInput
	// The name of the policy.
	PolicyName pulumi.StringPtrInput
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType pulumi.StringPtrInput
	// ID of the attached CAM role.
	RoleId pulumi.StringPtrInput
}

func (RolePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePolicyAttachmentState)(nil)).Elem()
}

type rolePolicyAttachmentArgs struct {
	// ID of the policy.
	PolicyId string `pulumi:"policyId"`
	// ID of the attached CAM role.
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a RolePolicyAttachment resource.
type RolePolicyAttachmentArgs struct {
	// ID of the policy.
	PolicyId pulumi.StringInput
	// ID of the attached CAM role.
	RoleId pulumi.StringInput
}

func (RolePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePolicyAttachmentArgs)(nil)).Elem()
}

type RolePolicyAttachmentInput interface {
	pulumi.Input

	ToRolePolicyAttachmentOutput() RolePolicyAttachmentOutput
	ToRolePolicyAttachmentOutputWithContext(ctx context.Context) RolePolicyAttachmentOutput
}

func (*RolePolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePolicyAttachment)(nil)).Elem()
}

func (i *RolePolicyAttachment) ToRolePolicyAttachmentOutput() RolePolicyAttachmentOutput {
	return i.ToRolePolicyAttachmentOutputWithContext(context.Background())
}

func (i *RolePolicyAttachment) ToRolePolicyAttachmentOutputWithContext(ctx context.Context) RolePolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyAttachmentOutput)
}

// RolePolicyAttachmentArrayInput is an input type that accepts RolePolicyAttachmentArray and RolePolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `RolePolicyAttachmentArrayInput` via:
//
//          RolePolicyAttachmentArray{ RolePolicyAttachmentArgs{...} }
type RolePolicyAttachmentArrayInput interface {
	pulumi.Input

	ToRolePolicyAttachmentArrayOutput() RolePolicyAttachmentArrayOutput
	ToRolePolicyAttachmentArrayOutputWithContext(context.Context) RolePolicyAttachmentArrayOutput
}

type RolePolicyAttachmentArray []RolePolicyAttachmentInput

func (RolePolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RolePolicyAttachment)(nil)).Elem()
}

func (i RolePolicyAttachmentArray) ToRolePolicyAttachmentArrayOutput() RolePolicyAttachmentArrayOutput {
	return i.ToRolePolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i RolePolicyAttachmentArray) ToRolePolicyAttachmentArrayOutputWithContext(ctx context.Context) RolePolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyAttachmentArrayOutput)
}

// RolePolicyAttachmentMapInput is an input type that accepts RolePolicyAttachmentMap and RolePolicyAttachmentMapOutput values.
// You can construct a concrete instance of `RolePolicyAttachmentMapInput` via:
//
//          RolePolicyAttachmentMap{ "key": RolePolicyAttachmentArgs{...} }
type RolePolicyAttachmentMapInput interface {
	pulumi.Input

	ToRolePolicyAttachmentMapOutput() RolePolicyAttachmentMapOutput
	ToRolePolicyAttachmentMapOutputWithContext(context.Context) RolePolicyAttachmentMapOutput
}

type RolePolicyAttachmentMap map[string]RolePolicyAttachmentInput

func (RolePolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RolePolicyAttachment)(nil)).Elem()
}

func (i RolePolicyAttachmentMap) ToRolePolicyAttachmentMapOutput() RolePolicyAttachmentMapOutput {
	return i.ToRolePolicyAttachmentMapOutputWithContext(context.Background())
}

func (i RolePolicyAttachmentMap) ToRolePolicyAttachmentMapOutputWithContext(ctx context.Context) RolePolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyAttachmentMapOutput)
}

type RolePolicyAttachmentOutput struct{ *pulumi.OutputState }

func (RolePolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePolicyAttachment)(nil)).Elem()
}

func (o RolePolicyAttachmentOutput) ToRolePolicyAttachmentOutput() RolePolicyAttachmentOutput {
	return o
}

func (o RolePolicyAttachmentOutput) ToRolePolicyAttachmentOutputWithContext(ctx context.Context) RolePolicyAttachmentOutput {
	return o
}

// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
func (o RolePolicyAttachmentOutput) CreateMode() pulumi.IntOutput {
	return o.ApplyT(func(v *RolePolicyAttachment) pulumi.IntOutput { return v.CreateMode }).(pulumi.IntOutput)
}

// The create time of the CAM role policy attachment.
func (o RolePolicyAttachmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RolePolicyAttachment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// ID of the policy.
func (o RolePolicyAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *RolePolicyAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The name of the policy.
func (o RolePolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *RolePolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
func (o RolePolicyAttachmentOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *RolePolicyAttachment) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// ID of the attached CAM role.
func (o RolePolicyAttachmentOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RolePolicyAttachment) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type RolePolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (RolePolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RolePolicyAttachment)(nil)).Elem()
}

func (o RolePolicyAttachmentArrayOutput) ToRolePolicyAttachmentArrayOutput() RolePolicyAttachmentArrayOutput {
	return o
}

func (o RolePolicyAttachmentArrayOutput) ToRolePolicyAttachmentArrayOutputWithContext(ctx context.Context) RolePolicyAttachmentArrayOutput {
	return o
}

func (o RolePolicyAttachmentArrayOutput) Index(i pulumi.IntInput) RolePolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RolePolicyAttachment {
		return vs[0].([]*RolePolicyAttachment)[vs[1].(int)]
	}).(RolePolicyAttachmentOutput)
}

type RolePolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (RolePolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RolePolicyAttachment)(nil)).Elem()
}

func (o RolePolicyAttachmentMapOutput) ToRolePolicyAttachmentMapOutput() RolePolicyAttachmentMapOutput {
	return o
}

func (o RolePolicyAttachmentMapOutput) ToRolePolicyAttachmentMapOutputWithContext(ctx context.Context) RolePolicyAttachmentMapOutput {
	return o
}

func (o RolePolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) RolePolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RolePolicyAttachment {
		return vs[0].(map[string]*RolePolicyAttachment)[vs[1].(string)]
	}).(RolePolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RolePolicyAttachmentInput)(nil)).Elem(), &RolePolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolePolicyAttachmentArrayInput)(nil)).Elem(), RolePolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolePolicyAttachmentMapInput)(nil)).Elem(), RolePolicyAttachmentMap{})
	pulumi.RegisterOutputType(RolePolicyAttachmentOutput{})
	pulumi.RegisterOutputType(RolePolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(RolePolicyAttachmentMapOutput{})
}
