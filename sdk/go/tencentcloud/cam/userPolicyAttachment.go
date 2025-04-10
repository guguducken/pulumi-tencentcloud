// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CAM user policy attachment.
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
// 		_, err := Cam.NewUserPolicyAttachment(ctx, "foo", &Cam.UserPolicyAttachmentArgs{
// 			UserId:   pulumi.Any(tencentcloud_cam_user.Foo.Id),
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
// CAM user policy attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment foo cam-test#26800353
// ```
type UserPolicyAttachment struct {
	pulumi.CustomResourceState

	// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode pulumi.IntOutput `pulumi:"createMode"`
	// Create time of the CAM user policy attachment.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// ID of the policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// Name of the policy.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
	//
	// Deprecated: It has been deprecated from version 1.59.5. Use `user_name` instead.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
	// Name of the attached CAM user as uniq key.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewUserPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPolicyAttachment(ctx *pulumi.Context,
	name string, args *UserPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserPolicyAttachment
	err := ctx.RegisterResource("tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPolicyAttachment gets an existing UserPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPolicyAttachmentState, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	var resource UserPolicyAttachment
	err := ctx.ReadResource("tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPolicyAttachment resources.
type userPolicyAttachmentState struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode *int `pulumi:"createMode"`
	// Create time of the CAM user policy attachment.
	CreateTime *string `pulumi:"createTime"`
	// ID of the policy.
	PolicyId *string `pulumi:"policyId"`
	// Name of the policy.
	PolicyName *string `pulumi:"policyName"`
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType *string `pulumi:"policyType"`
	// It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
	//
	// Deprecated: It has been deprecated from version 1.59.5. Use `user_name` instead.
	UserId *string `pulumi:"userId"`
	// Name of the attached CAM user as uniq key.
	UserName *string `pulumi:"userName"`
}

type UserPolicyAttachmentState struct {
	// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
	CreateMode pulumi.IntPtrInput
	// Create time of the CAM user policy attachment.
	CreateTime pulumi.StringPtrInput
	// ID of the policy.
	PolicyId pulumi.StringPtrInput
	// Name of the policy.
	PolicyName pulumi.StringPtrInput
	// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
	PolicyType pulumi.StringPtrInput
	// It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
	//
	// Deprecated: It has been deprecated from version 1.59.5. Use `user_name` instead.
	UserId pulumi.StringPtrInput
	// Name of the attached CAM user as uniq key.
	UserName pulumi.StringPtrInput
}

func (UserPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentState)(nil)).Elem()
}

type userPolicyAttachmentArgs struct {
	// ID of the policy.
	PolicyId string `pulumi:"policyId"`
	// It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
	//
	// Deprecated: It has been deprecated from version 1.59.5. Use `user_name` instead.
	UserId *string `pulumi:"userId"`
	// Name of the attached CAM user as uniq key.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a UserPolicyAttachment resource.
type UserPolicyAttachmentArgs struct {
	// ID of the policy.
	PolicyId pulumi.StringInput
	// It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
	//
	// Deprecated: It has been deprecated from version 1.59.5. Use `user_name` instead.
	UserId pulumi.StringPtrInput
	// Name of the attached CAM user as uniq key.
	UserName pulumi.StringPtrInput
}

func (UserPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentArgs)(nil)).Elem()
}

type UserPolicyAttachmentInput interface {
	pulumi.Input

	ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput
	ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput
}

func (*UserPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (i *UserPolicyAttachment) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return i.ToUserPolicyAttachmentOutputWithContext(context.Background())
}

func (i *UserPolicyAttachment) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentOutput)
}

// UserPolicyAttachmentArrayInput is an input type that accepts UserPolicyAttachmentArray and UserPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `UserPolicyAttachmentArrayInput` via:
//
//          UserPolicyAttachmentArray{ UserPolicyAttachmentArgs{...} }
type UserPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput
	ToUserPolicyAttachmentArrayOutputWithContext(context.Context) UserPolicyAttachmentArrayOutput
}

type UserPolicyAttachmentArray []UserPolicyAttachmentInput

func (UserPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachmentArray) ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput {
	return i.ToUserPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i UserPolicyAttachmentArray) ToUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) UserPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentArrayOutput)
}

// UserPolicyAttachmentMapInput is an input type that accepts UserPolicyAttachmentMap and UserPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `UserPolicyAttachmentMapInput` via:
//
//          UserPolicyAttachmentMap{ "key": UserPolicyAttachmentArgs{...} }
type UserPolicyAttachmentMapInput interface {
	pulumi.Input

	ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput
	ToUserPolicyAttachmentMapOutputWithContext(context.Context) UserPolicyAttachmentMapOutput
}

type UserPolicyAttachmentMap map[string]UserPolicyAttachmentInput

func (UserPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachmentMap) ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput {
	return i.ToUserPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i UserPolicyAttachmentMap) ToUserPolicyAttachmentMapOutputWithContext(ctx context.Context) UserPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentMapOutput)
}

type UserPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return o
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return o
}

// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
func (o UserPolicyAttachmentOutput) CreateMode() pulumi.IntOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.IntOutput { return v.CreateMode }).(pulumi.IntOutput)
}

// Create time of the CAM user policy attachment.
func (o UserPolicyAttachmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// ID of the policy.
func (o UserPolicyAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// Name of the policy.
func (o UserPolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
func (o UserPolicyAttachmentOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// It has been deprecated from version 1.59.5. Use `userName` instead. ID of the attached CAM user.
//
// Deprecated: It has been deprecated from version 1.59.5. Use `user_name` instead.
func (o UserPolicyAttachmentOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

// Name of the attached CAM user as uniq key.
func (o UserPolicyAttachmentOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

type UserPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentArrayOutput) ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput {
	return o
}

func (o UserPolicyAttachmentArrayOutput) ToUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) UserPolicyAttachmentArrayOutput {
	return o
}

func (o UserPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) UserPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPolicyAttachment {
		return vs[0].([]*UserPolicyAttachment)[vs[1].(int)]
	}).(UserPolicyAttachmentOutput)
}

type UserPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentMapOutput) ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput {
	return o
}

func (o UserPolicyAttachmentMapOutput) ToUserPolicyAttachmentMapOutputWithContext(ctx context.Context) UserPolicyAttachmentMapOutput {
	return o
}

func (o UserPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) UserPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPolicyAttachment {
		return vs[0].(map[string]*UserPolicyAttachment)[vs[1].(string)]
	}).(UserPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentInput)(nil)).Elem(), &UserPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentArrayInput)(nil)).Elem(), UserPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentMapInput)(nil)).Elem(), UserPolicyAttachmentMap{})
	pulumi.RegisterOutputType(UserPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentMapOutput{})
}
