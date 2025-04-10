// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cam userSamlConfig
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
// 		_, err := Cam.NewUserSamlConfig(ctx, "userSamlConfig", &Cam.UserSamlConfigArgs{
// 			SamlMetadataDocument: pulumi.String("./metadataDocument.xml"),
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
// cam user_saml_config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cam/userSamlConfig:UserSamlConfig user_saml_config user_id
// ```
type UserSamlConfig struct {
	pulumi.CustomResourceState

	// The path used to save the samlMetadat file.
	MetadataDocumentFile pulumi.StringPtrOutput `pulumi:"metadataDocumentFile"`
	// SAML metadata document, xml format, support string content or file path.
	SamlMetadataDocument pulumi.StringOutput `pulumi:"samlMetadataDocument"`
	// Status: `0`: not set, `11`: enabled, `2`: disabled.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewUserSamlConfig registers a new resource with the given unique name, arguments, and options.
func NewUserSamlConfig(ctx *pulumi.Context,
	name string, args *UserSamlConfigArgs, opts ...pulumi.ResourceOption) (*UserSamlConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SamlMetadataDocument == nil {
		return nil, errors.New("invalid value for required argument 'SamlMetadataDocument'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserSamlConfig
	err := ctx.RegisterResource("tencentcloud:Cam/userSamlConfig:UserSamlConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSamlConfig gets an existing UserSamlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSamlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSamlConfigState, opts ...pulumi.ResourceOption) (*UserSamlConfig, error) {
	var resource UserSamlConfig
	err := ctx.ReadResource("tencentcloud:Cam/userSamlConfig:UserSamlConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSamlConfig resources.
type userSamlConfigState struct {
	// The path used to save the samlMetadat file.
	MetadataDocumentFile *string `pulumi:"metadataDocumentFile"`
	// SAML metadata document, xml format, support string content or file path.
	SamlMetadataDocument *string `pulumi:"samlMetadataDocument"`
	// Status: `0`: not set, `11`: enabled, `2`: disabled.
	Status *int `pulumi:"status"`
}

type UserSamlConfigState struct {
	// The path used to save the samlMetadat file.
	MetadataDocumentFile pulumi.StringPtrInput
	// SAML metadata document, xml format, support string content or file path.
	SamlMetadataDocument pulumi.StringPtrInput
	// Status: `0`: not set, `11`: enabled, `2`: disabled.
	Status pulumi.IntPtrInput
}

func (UserSamlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSamlConfigState)(nil)).Elem()
}

type userSamlConfigArgs struct {
	// The path used to save the samlMetadat file.
	MetadataDocumentFile *string `pulumi:"metadataDocumentFile"`
	// SAML metadata document, xml format, support string content or file path.
	SamlMetadataDocument string `pulumi:"samlMetadataDocument"`
}

// The set of arguments for constructing a UserSamlConfig resource.
type UserSamlConfigArgs struct {
	// The path used to save the samlMetadat file.
	MetadataDocumentFile pulumi.StringPtrInput
	// SAML metadata document, xml format, support string content or file path.
	SamlMetadataDocument pulumi.StringInput
}

func (UserSamlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSamlConfigArgs)(nil)).Elem()
}

type UserSamlConfigInput interface {
	pulumi.Input

	ToUserSamlConfigOutput() UserSamlConfigOutput
	ToUserSamlConfigOutputWithContext(ctx context.Context) UserSamlConfigOutput
}

func (*UserSamlConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSamlConfig)(nil)).Elem()
}

func (i *UserSamlConfig) ToUserSamlConfigOutput() UserSamlConfigOutput {
	return i.ToUserSamlConfigOutputWithContext(context.Background())
}

func (i *UserSamlConfig) ToUserSamlConfigOutputWithContext(ctx context.Context) UserSamlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSamlConfigOutput)
}

// UserSamlConfigArrayInput is an input type that accepts UserSamlConfigArray and UserSamlConfigArrayOutput values.
// You can construct a concrete instance of `UserSamlConfigArrayInput` via:
//
//          UserSamlConfigArray{ UserSamlConfigArgs{...} }
type UserSamlConfigArrayInput interface {
	pulumi.Input

	ToUserSamlConfigArrayOutput() UserSamlConfigArrayOutput
	ToUserSamlConfigArrayOutputWithContext(context.Context) UserSamlConfigArrayOutput
}

type UserSamlConfigArray []UserSamlConfigInput

func (UserSamlConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSamlConfig)(nil)).Elem()
}

func (i UserSamlConfigArray) ToUserSamlConfigArrayOutput() UserSamlConfigArrayOutput {
	return i.ToUserSamlConfigArrayOutputWithContext(context.Background())
}

func (i UserSamlConfigArray) ToUserSamlConfigArrayOutputWithContext(ctx context.Context) UserSamlConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSamlConfigArrayOutput)
}

// UserSamlConfigMapInput is an input type that accepts UserSamlConfigMap and UserSamlConfigMapOutput values.
// You can construct a concrete instance of `UserSamlConfigMapInput` via:
//
//          UserSamlConfigMap{ "key": UserSamlConfigArgs{...} }
type UserSamlConfigMapInput interface {
	pulumi.Input

	ToUserSamlConfigMapOutput() UserSamlConfigMapOutput
	ToUserSamlConfigMapOutputWithContext(context.Context) UserSamlConfigMapOutput
}

type UserSamlConfigMap map[string]UserSamlConfigInput

func (UserSamlConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSamlConfig)(nil)).Elem()
}

func (i UserSamlConfigMap) ToUserSamlConfigMapOutput() UserSamlConfigMapOutput {
	return i.ToUserSamlConfigMapOutputWithContext(context.Background())
}

func (i UserSamlConfigMap) ToUserSamlConfigMapOutputWithContext(ctx context.Context) UserSamlConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSamlConfigMapOutput)
}

type UserSamlConfigOutput struct{ *pulumi.OutputState }

func (UserSamlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSamlConfig)(nil)).Elem()
}

func (o UserSamlConfigOutput) ToUserSamlConfigOutput() UserSamlConfigOutput {
	return o
}

func (o UserSamlConfigOutput) ToUserSamlConfigOutputWithContext(ctx context.Context) UserSamlConfigOutput {
	return o
}

// The path used to save the samlMetadat file.
func (o UserSamlConfigOutput) MetadataDocumentFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSamlConfig) pulumi.StringPtrOutput { return v.MetadataDocumentFile }).(pulumi.StringPtrOutput)
}

// SAML metadata document, xml format, support string content or file path.
func (o UserSamlConfigOutput) SamlMetadataDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSamlConfig) pulumi.StringOutput { return v.SamlMetadataDocument }).(pulumi.StringOutput)
}

// Status: `0`: not set, `11`: enabled, `2`: disabled.
func (o UserSamlConfigOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *UserSamlConfig) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

type UserSamlConfigArrayOutput struct{ *pulumi.OutputState }

func (UserSamlConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSamlConfig)(nil)).Elem()
}

func (o UserSamlConfigArrayOutput) ToUserSamlConfigArrayOutput() UserSamlConfigArrayOutput {
	return o
}

func (o UserSamlConfigArrayOutput) ToUserSamlConfigArrayOutputWithContext(ctx context.Context) UserSamlConfigArrayOutput {
	return o
}

func (o UserSamlConfigArrayOutput) Index(i pulumi.IntInput) UserSamlConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserSamlConfig {
		return vs[0].([]*UserSamlConfig)[vs[1].(int)]
	}).(UserSamlConfigOutput)
}

type UserSamlConfigMapOutput struct{ *pulumi.OutputState }

func (UserSamlConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSamlConfig)(nil)).Elem()
}

func (o UserSamlConfigMapOutput) ToUserSamlConfigMapOutput() UserSamlConfigMapOutput {
	return o
}

func (o UserSamlConfigMapOutput) ToUserSamlConfigMapOutputWithContext(ctx context.Context) UserSamlConfigMapOutput {
	return o
}

func (o UserSamlConfigMapOutput) MapIndex(k pulumi.StringInput) UserSamlConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserSamlConfig {
		return vs[0].(map[string]*UserSamlConfig)[vs[1].(string)]
	}).(UserSamlConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserSamlConfigInput)(nil)).Elem(), &UserSamlConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSamlConfigArrayInput)(nil)).Elem(), UserSamlConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSamlConfigMapInput)(nil)).Elem(), UserSamlConfigMap{})
	pulumi.RegisterOutputType(UserSamlConfigOutput{})
	pulumi.RegisterOutputType(UserSamlConfigArrayOutput{})
	pulumi.RegisterOutputType(UserSamlConfigMapOutput{})
}
