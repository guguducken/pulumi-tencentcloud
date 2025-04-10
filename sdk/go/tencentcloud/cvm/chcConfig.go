// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cvm chcConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cvm.NewChcConfig(ctx, "chcConfig", &Cvm.ChcConfigArgs{
// 			BmcSecurityGroupIds: pulumi.StringArray{
// 				pulumi.String("sg-xxxxxx"),
// 			},
// 			BmcUser: pulumi.String("admin"),
// 			BmcVirtualPrivateCloud: &cvm.ChcConfigBmcVirtualPrivateCloudArgs{
// 				SubnetId: pulumi.String("subnet-xxxxxx"),
// 				VpcId:    pulumi.String("vpc-xxxxxx"),
// 			},
// 			ChcId: pulumi.String("chc-xxxxxx"),
// 			DeploySecurityGroupIds: pulumi.StringArray{
// 				pulumi.String("sg-xxxxxx"),
// 			},
// 			DeployVirtualPrivateCloud: &cvm.ChcConfigDeployVirtualPrivateCloudArgs{
// 				SubnetId: pulumi.String("subnet-xxxxxx"),
// 				VpcId:    pulumi.String("vpc-xxxxxx"),
// 			},
// 			InstanceName: pulumi.String("xxxxxx"),
// 			Password:     pulumi.String("xxxxxx"),
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
// cvm chc_config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cvm/chcConfig:ChcConfig chc_config chc_config_id
// ```
type ChcConfig struct {
	pulumi.CustomResourceState

	// Out-of-band network security group list.
	BmcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"bmcSecurityGroupIds"`
	// Valid characters: Letters, numbers, hyphens and underscores. Only set when update password.
	BmcUser pulumi.StringPtrOutput `pulumi:"bmcUser"`
	// Out-of-band network information.
	BmcVirtualPrivateCloud ChcConfigBmcVirtualPrivateCloudOutput `pulumi:"bmcVirtualPrivateCloud"`
	// CHC host ID.
	ChcId pulumi.StringOutput `pulumi:"chcId"`
	// Deployment network security group list.
	DeploySecurityGroupIds pulumi.StringArrayOutput `pulumi:"deploySecurityGroupIds"`
	// Deployment network information.
	DeployVirtualPrivateCloud ChcConfigDeployVirtualPrivateCloudOutput `pulumi:"deployVirtualPrivateCloud"`
	// Server type.
	DeviceType pulumi.StringOutput `pulumi:"deviceType"`
	// CHC host name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&amp;amp;*-+=_|{}).
	Password pulumi.StringPtrOutput `pulumi:"password"`
}

// NewChcConfig registers a new resource with the given unique name, arguments, and options.
func NewChcConfig(ctx *pulumi.Context,
	name string, args *ChcConfigArgs, opts ...pulumi.ResourceOption) (*ChcConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChcId == nil {
		return nil, errors.New("invalid value for required argument 'ChcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ChcConfig
	err := ctx.RegisterResource("tencentcloud:Cvm/chcConfig:ChcConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChcConfig gets an existing ChcConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChcConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChcConfigState, opts ...pulumi.ResourceOption) (*ChcConfig, error) {
	var resource ChcConfig
	err := ctx.ReadResource("tencentcloud:Cvm/chcConfig:ChcConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChcConfig resources.
type chcConfigState struct {
	// Out-of-band network security group list.
	BmcSecurityGroupIds []string `pulumi:"bmcSecurityGroupIds"`
	// Valid characters: Letters, numbers, hyphens and underscores. Only set when update password.
	BmcUser *string `pulumi:"bmcUser"`
	// Out-of-band network information.
	BmcVirtualPrivateCloud *ChcConfigBmcVirtualPrivateCloud `pulumi:"bmcVirtualPrivateCloud"`
	// CHC host ID.
	ChcId *string `pulumi:"chcId"`
	// Deployment network security group list.
	DeploySecurityGroupIds []string `pulumi:"deploySecurityGroupIds"`
	// Deployment network information.
	DeployVirtualPrivateCloud *ChcConfigDeployVirtualPrivateCloud `pulumi:"deployVirtualPrivateCloud"`
	// Server type.
	DeviceType *string `pulumi:"deviceType"`
	// CHC host name.
	InstanceName *string `pulumi:"instanceName"`
	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&amp;amp;*-+=_|{}).
	Password *string `pulumi:"password"`
}

type ChcConfigState struct {
	// Out-of-band network security group list.
	BmcSecurityGroupIds pulumi.StringArrayInput
	// Valid characters: Letters, numbers, hyphens and underscores. Only set when update password.
	BmcUser pulumi.StringPtrInput
	// Out-of-band network information.
	BmcVirtualPrivateCloud ChcConfigBmcVirtualPrivateCloudPtrInput
	// CHC host ID.
	ChcId pulumi.StringPtrInput
	// Deployment network security group list.
	DeploySecurityGroupIds pulumi.StringArrayInput
	// Deployment network information.
	DeployVirtualPrivateCloud ChcConfigDeployVirtualPrivateCloudPtrInput
	// Server type.
	DeviceType pulumi.StringPtrInput
	// CHC host name.
	InstanceName pulumi.StringPtrInput
	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&amp;amp;*-+=_|{}).
	Password pulumi.StringPtrInput
}

func (ChcConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*chcConfigState)(nil)).Elem()
}

type chcConfigArgs struct {
	// Out-of-band network security group list.
	BmcSecurityGroupIds []string `pulumi:"bmcSecurityGroupIds"`
	// Valid characters: Letters, numbers, hyphens and underscores. Only set when update password.
	BmcUser *string `pulumi:"bmcUser"`
	// Out-of-band network information.
	BmcVirtualPrivateCloud *ChcConfigBmcVirtualPrivateCloud `pulumi:"bmcVirtualPrivateCloud"`
	// CHC host ID.
	ChcId string `pulumi:"chcId"`
	// Deployment network security group list.
	DeploySecurityGroupIds []string `pulumi:"deploySecurityGroupIds"`
	// Deployment network information.
	DeployVirtualPrivateCloud *ChcConfigDeployVirtualPrivateCloud `pulumi:"deployVirtualPrivateCloud"`
	// Server type.
	DeviceType *string `pulumi:"deviceType"`
	// CHC host name.
	InstanceName *string `pulumi:"instanceName"`
	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&amp;amp;*-+=_|{}).
	Password *string `pulumi:"password"`
}

// The set of arguments for constructing a ChcConfig resource.
type ChcConfigArgs struct {
	// Out-of-band network security group list.
	BmcSecurityGroupIds pulumi.StringArrayInput
	// Valid characters: Letters, numbers, hyphens and underscores. Only set when update password.
	BmcUser pulumi.StringPtrInput
	// Out-of-band network information.
	BmcVirtualPrivateCloud ChcConfigBmcVirtualPrivateCloudPtrInput
	// CHC host ID.
	ChcId pulumi.StringInput
	// Deployment network security group list.
	DeploySecurityGroupIds pulumi.StringArrayInput
	// Deployment network information.
	DeployVirtualPrivateCloud ChcConfigDeployVirtualPrivateCloudPtrInput
	// Server type.
	DeviceType pulumi.StringPtrInput
	// CHC host name.
	InstanceName pulumi.StringPtrInput
	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&amp;amp;*-+=_|{}).
	Password pulumi.StringPtrInput
}

func (ChcConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chcConfigArgs)(nil)).Elem()
}

type ChcConfigInput interface {
	pulumi.Input

	ToChcConfigOutput() ChcConfigOutput
	ToChcConfigOutputWithContext(ctx context.Context) ChcConfigOutput
}

func (*ChcConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ChcConfig)(nil)).Elem()
}

func (i *ChcConfig) ToChcConfigOutput() ChcConfigOutput {
	return i.ToChcConfigOutputWithContext(context.Background())
}

func (i *ChcConfig) ToChcConfigOutputWithContext(ctx context.Context) ChcConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChcConfigOutput)
}

// ChcConfigArrayInput is an input type that accepts ChcConfigArray and ChcConfigArrayOutput values.
// You can construct a concrete instance of `ChcConfigArrayInput` via:
//
//          ChcConfigArray{ ChcConfigArgs{...} }
type ChcConfigArrayInput interface {
	pulumi.Input

	ToChcConfigArrayOutput() ChcConfigArrayOutput
	ToChcConfigArrayOutputWithContext(context.Context) ChcConfigArrayOutput
}

type ChcConfigArray []ChcConfigInput

func (ChcConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ChcConfig)(nil)).Elem()
}

func (i ChcConfigArray) ToChcConfigArrayOutput() ChcConfigArrayOutput {
	return i.ToChcConfigArrayOutputWithContext(context.Background())
}

func (i ChcConfigArray) ToChcConfigArrayOutputWithContext(ctx context.Context) ChcConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChcConfigArrayOutput)
}

// ChcConfigMapInput is an input type that accepts ChcConfigMap and ChcConfigMapOutput values.
// You can construct a concrete instance of `ChcConfigMapInput` via:
//
//          ChcConfigMap{ "key": ChcConfigArgs{...} }
type ChcConfigMapInput interface {
	pulumi.Input

	ToChcConfigMapOutput() ChcConfigMapOutput
	ToChcConfigMapOutputWithContext(context.Context) ChcConfigMapOutput
}

type ChcConfigMap map[string]ChcConfigInput

func (ChcConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ChcConfig)(nil)).Elem()
}

func (i ChcConfigMap) ToChcConfigMapOutput() ChcConfigMapOutput {
	return i.ToChcConfigMapOutputWithContext(context.Background())
}

func (i ChcConfigMap) ToChcConfigMapOutputWithContext(ctx context.Context) ChcConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChcConfigMapOutput)
}

type ChcConfigOutput struct{ *pulumi.OutputState }

func (ChcConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChcConfig)(nil)).Elem()
}

func (o ChcConfigOutput) ToChcConfigOutput() ChcConfigOutput {
	return o
}

func (o ChcConfigOutput) ToChcConfigOutputWithContext(ctx context.Context) ChcConfigOutput {
	return o
}

// Out-of-band network security group list.
func (o ChcConfigOutput) BmcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringArrayOutput { return v.BmcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Valid characters: Letters, numbers, hyphens and underscores. Only set when update password.
func (o ChcConfigOutput) BmcUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringPtrOutput { return v.BmcUser }).(pulumi.StringPtrOutput)
}

// Out-of-band network information.
func (o ChcConfigOutput) BmcVirtualPrivateCloud() ChcConfigBmcVirtualPrivateCloudOutput {
	return o.ApplyT(func(v *ChcConfig) ChcConfigBmcVirtualPrivateCloudOutput { return v.BmcVirtualPrivateCloud }).(ChcConfigBmcVirtualPrivateCloudOutput)
}

// CHC host ID.
func (o ChcConfigOutput) ChcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringOutput { return v.ChcId }).(pulumi.StringOutput)
}

// Deployment network security group list.
func (o ChcConfigOutput) DeploySecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringArrayOutput { return v.DeploySecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Deployment network information.
func (o ChcConfigOutput) DeployVirtualPrivateCloud() ChcConfigDeployVirtualPrivateCloudOutput {
	return o.ApplyT(func(v *ChcConfig) ChcConfigDeployVirtualPrivateCloudOutput { return v.DeployVirtualPrivateCloud }).(ChcConfigDeployVirtualPrivateCloudOutput)
}

// Server type.
func (o ChcConfigOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

// CHC host name.
func (o ChcConfigOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&amp;amp;*-+=_|{}).
func (o ChcConfigOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChcConfig) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

type ChcConfigArrayOutput struct{ *pulumi.OutputState }

func (ChcConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ChcConfig)(nil)).Elem()
}

func (o ChcConfigArrayOutput) ToChcConfigArrayOutput() ChcConfigArrayOutput {
	return o
}

func (o ChcConfigArrayOutput) ToChcConfigArrayOutputWithContext(ctx context.Context) ChcConfigArrayOutput {
	return o
}

func (o ChcConfigArrayOutput) Index(i pulumi.IntInput) ChcConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ChcConfig {
		return vs[0].([]*ChcConfig)[vs[1].(int)]
	}).(ChcConfigOutput)
}

type ChcConfigMapOutput struct{ *pulumi.OutputState }

func (ChcConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ChcConfig)(nil)).Elem()
}

func (o ChcConfigMapOutput) ToChcConfigMapOutput() ChcConfigMapOutput {
	return o
}

func (o ChcConfigMapOutput) ToChcConfigMapOutputWithContext(ctx context.Context) ChcConfigMapOutput {
	return o
}

func (o ChcConfigMapOutput) MapIndex(k pulumi.StringInput) ChcConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ChcConfig {
		return vs[0].(map[string]*ChcConfig)[vs[1].(string)]
	}).(ChcConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChcConfigInput)(nil)).Elem(), &ChcConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChcConfigArrayInput)(nil)).Elem(), ChcConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChcConfigMapInput)(nil)).Elem(), ChcConfigMap{})
	pulumi.RegisterOutputType(ChcConfigOutput{})
	pulumi.RegisterOutputType(ChcConfigArrayOutput{})
	pulumi.RegisterOutputType(ChcConfigMapOutput{})
}
