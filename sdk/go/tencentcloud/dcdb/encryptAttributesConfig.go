// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dcdb encryptAttributesConfig
//
// > **NOTE:**  This resource currently only supports the newly created MySQL 8.0.24 version.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		internal, err := Security.GetGroups(ctx, &security.GetGroupsArgs{
// 			Name: pulumi.StringRef("default"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vpc, err := Vpc.GetInstances(ctx, &vpc.GetInstancesArgs{
// 			Name: pulumi.StringRef("Default-VPC"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		subnet, err := Vpc.GetSubnets(ctx, &vpc.GetSubnetsArgs{
// 			VpcId: pulumi.StringRef(vpc.InstanceLists[0].VpcId),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vpcId := subnet.InstanceLists[0].VpcId
// 		subnetId := subnet.InstanceLists[0].SubnetId
// 		sgId := internal.SecurityGroups[0].SecurityGroupId
// 		prepaidInstance, err := Dcdb.NewDbInstance(ctx, "prepaidInstance", &Dcdb.DbInstanceArgs{
// 			InstanceName: pulumi.String("test_dcdb_db_post_instance"),
// 			Zones: pulumi.StringArray{
// 				pulumi.Any(_var.Default_az),
// 			},
// 			Period:         pulumi.Int(1),
// 			ShardMemory:    pulumi.Int(2),
// 			ShardStorage:   pulumi.Int(10),
// 			ShardNodeCount: pulumi.Int(2),
// 			ShardCount:     pulumi.Int(2),
// 			VpcId:          pulumi.String(vpcId),
// 			SubnetId:       pulumi.String(subnetId),
// 			DbVersionId:    pulumi.String("8.0"),
// 			ResourceTags: dcdb.DbInstanceResourceTagArray{
// 				&dcdb.DbInstanceResourceTagArgs{
// 					TagKey:   pulumi.String("aaa"),
// 					TagValue: pulumi.String("bbb"),
// 				},
// 			},
// 			SecurityGroupIds: pulumi.StringArray{
// 				pulumi.String(sgId),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		hourdbInstance, err := Dcdb.NewHourdbInstance(ctx, "hourdbInstance", &Dcdb.HourdbInstanceArgs{
// 			InstanceName: pulumi.String("test_dcdb_db_hourdb_instance"),
// 			Zones: pulumi.StringArray{
// 				pulumi.Any(_var.Default_az),
// 			},
// 			ShardMemory:     pulumi.Int(2),
// 			ShardStorage:    pulumi.Int(10),
// 			ShardNodeCount:  pulumi.Int(2),
// 			ShardCount:      pulumi.Int(2),
// 			VpcId:           pulumi.String(vpcId),
// 			SubnetId:        pulumi.String(subnetId),
// 			SecurityGroupId: pulumi.String(sgId),
// 			DbVersionId:     pulumi.String("8.0"),
// 			ResourceTags: dcdb.HourdbInstanceResourceTagArray{
// 				&dcdb.HourdbInstanceResourceTagArgs{
// 					TagKey:   pulumi.String("aaa"),
// 					TagValue: pulumi.String("bbb"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		prepaidDcdbId := prepaidInstance.ID()
// 		hourdbDcdbId := hourdbInstance.ID()
// 		_, err = Dcdb.NewEncryptAttributesConfig(ctx, "configHourdb", &Dcdb.EncryptAttributesConfigArgs{
// 			InstanceId:     pulumi.String(hourdbDcdbId),
// 			EncryptEnabled: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dcdb.NewEncryptAttributesConfig(ctx, "configPrepaid", &Dcdb.EncryptAttributesConfigArgs{
// 			InstanceId:     pulumi.String(prepaidDcdbId),
// 			EncryptEnabled: pulumi.Int(1),
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
// dcdb encrypt_attributes_config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig encrypt_attributes_config encrypt_attributes_config_id
// ```
type EncryptAttributesConfig struct {
	pulumi.CustomResourceState

	// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
	EncryptEnabled pulumi.IntOutput `pulumi:"encryptEnabled"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewEncryptAttributesConfig registers a new resource with the given unique name, arguments, and options.
func NewEncryptAttributesConfig(ctx *pulumi.Context,
	name string, args *EncryptAttributesConfigArgs, opts ...pulumi.ResourceOption) (*EncryptAttributesConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncryptEnabled == nil {
		return nil, errors.New("invalid value for required argument 'EncryptEnabled'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EncryptAttributesConfig
	err := ctx.RegisterResource("tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptAttributesConfig gets an existing EncryptAttributesConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptAttributesConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptAttributesConfigState, opts ...pulumi.ResourceOption) (*EncryptAttributesConfig, error) {
	var resource EncryptAttributesConfig
	err := ctx.ReadResource("tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptAttributesConfig resources.
type encryptAttributesConfigState struct {
	// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
	EncryptEnabled *int `pulumi:"encryptEnabled"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
}

type EncryptAttributesConfigState struct {
	// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
	EncryptEnabled pulumi.IntPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
}

func (EncryptAttributesConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptAttributesConfigState)(nil)).Elem()
}

type encryptAttributesConfigArgs struct {
	// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
	EncryptEnabled int `pulumi:"encryptEnabled"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a EncryptAttributesConfig resource.
type EncryptAttributesConfigArgs struct {
	// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
	EncryptEnabled pulumi.IntInput
	// instance id.
	InstanceId pulumi.StringInput
}

func (EncryptAttributesConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptAttributesConfigArgs)(nil)).Elem()
}

type EncryptAttributesConfigInput interface {
	pulumi.Input

	ToEncryptAttributesConfigOutput() EncryptAttributesConfigOutput
	ToEncryptAttributesConfigOutputWithContext(ctx context.Context) EncryptAttributesConfigOutput
}

func (*EncryptAttributesConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptAttributesConfig)(nil)).Elem()
}

func (i *EncryptAttributesConfig) ToEncryptAttributesConfigOutput() EncryptAttributesConfigOutput {
	return i.ToEncryptAttributesConfigOutputWithContext(context.Background())
}

func (i *EncryptAttributesConfig) ToEncryptAttributesConfigOutputWithContext(ctx context.Context) EncryptAttributesConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptAttributesConfigOutput)
}

// EncryptAttributesConfigArrayInput is an input type that accepts EncryptAttributesConfigArray and EncryptAttributesConfigArrayOutput values.
// You can construct a concrete instance of `EncryptAttributesConfigArrayInput` via:
//
//          EncryptAttributesConfigArray{ EncryptAttributesConfigArgs{...} }
type EncryptAttributesConfigArrayInput interface {
	pulumi.Input

	ToEncryptAttributesConfigArrayOutput() EncryptAttributesConfigArrayOutput
	ToEncryptAttributesConfigArrayOutputWithContext(context.Context) EncryptAttributesConfigArrayOutput
}

type EncryptAttributesConfigArray []EncryptAttributesConfigInput

func (EncryptAttributesConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EncryptAttributesConfig)(nil)).Elem()
}

func (i EncryptAttributesConfigArray) ToEncryptAttributesConfigArrayOutput() EncryptAttributesConfigArrayOutput {
	return i.ToEncryptAttributesConfigArrayOutputWithContext(context.Background())
}

func (i EncryptAttributesConfigArray) ToEncryptAttributesConfigArrayOutputWithContext(ctx context.Context) EncryptAttributesConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptAttributesConfigArrayOutput)
}

// EncryptAttributesConfigMapInput is an input type that accepts EncryptAttributesConfigMap and EncryptAttributesConfigMapOutput values.
// You can construct a concrete instance of `EncryptAttributesConfigMapInput` via:
//
//          EncryptAttributesConfigMap{ "key": EncryptAttributesConfigArgs{...} }
type EncryptAttributesConfigMapInput interface {
	pulumi.Input

	ToEncryptAttributesConfigMapOutput() EncryptAttributesConfigMapOutput
	ToEncryptAttributesConfigMapOutputWithContext(context.Context) EncryptAttributesConfigMapOutput
}

type EncryptAttributesConfigMap map[string]EncryptAttributesConfigInput

func (EncryptAttributesConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EncryptAttributesConfig)(nil)).Elem()
}

func (i EncryptAttributesConfigMap) ToEncryptAttributesConfigMapOutput() EncryptAttributesConfigMapOutput {
	return i.ToEncryptAttributesConfigMapOutputWithContext(context.Background())
}

func (i EncryptAttributesConfigMap) ToEncryptAttributesConfigMapOutputWithContext(ctx context.Context) EncryptAttributesConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptAttributesConfigMapOutput)
}

type EncryptAttributesConfigOutput struct{ *pulumi.OutputState }

func (EncryptAttributesConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptAttributesConfig)(nil)).Elem()
}

func (o EncryptAttributesConfigOutput) ToEncryptAttributesConfigOutput() EncryptAttributesConfigOutput {
	return o
}

func (o EncryptAttributesConfigOutput) ToEncryptAttributesConfigOutputWithContext(ctx context.Context) EncryptAttributesConfigOutput {
	return o
}

// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
func (o EncryptAttributesConfigOutput) EncryptEnabled() pulumi.IntOutput {
	return o.ApplyT(func(v *EncryptAttributesConfig) pulumi.IntOutput { return v.EncryptEnabled }).(pulumi.IntOutput)
}

// instance id.
func (o EncryptAttributesConfigOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptAttributesConfig) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type EncryptAttributesConfigArrayOutput struct{ *pulumi.OutputState }

func (EncryptAttributesConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EncryptAttributesConfig)(nil)).Elem()
}

func (o EncryptAttributesConfigArrayOutput) ToEncryptAttributesConfigArrayOutput() EncryptAttributesConfigArrayOutput {
	return o
}

func (o EncryptAttributesConfigArrayOutput) ToEncryptAttributesConfigArrayOutputWithContext(ctx context.Context) EncryptAttributesConfigArrayOutput {
	return o
}

func (o EncryptAttributesConfigArrayOutput) Index(i pulumi.IntInput) EncryptAttributesConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EncryptAttributesConfig {
		return vs[0].([]*EncryptAttributesConfig)[vs[1].(int)]
	}).(EncryptAttributesConfigOutput)
}

type EncryptAttributesConfigMapOutput struct{ *pulumi.OutputState }

func (EncryptAttributesConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EncryptAttributesConfig)(nil)).Elem()
}

func (o EncryptAttributesConfigMapOutput) ToEncryptAttributesConfigMapOutput() EncryptAttributesConfigMapOutput {
	return o
}

func (o EncryptAttributesConfigMapOutput) ToEncryptAttributesConfigMapOutputWithContext(ctx context.Context) EncryptAttributesConfigMapOutput {
	return o
}

func (o EncryptAttributesConfigMapOutput) MapIndex(k pulumi.StringInput) EncryptAttributesConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EncryptAttributesConfig {
		return vs[0].(map[string]*EncryptAttributesConfig)[vs[1].(string)]
	}).(EncryptAttributesConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptAttributesConfigInput)(nil)).Elem(), &EncryptAttributesConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptAttributesConfigArrayInput)(nil)).Elem(), EncryptAttributesConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptAttributesConfigMapInput)(nil)).Elem(), EncryptAttributesConfigMap{})
	pulumi.RegisterOutputType(EncryptAttributesConfigOutput{})
	pulumi.RegisterOutputType(EncryptAttributesConfigArrayOutput{})
	pulumi.RegisterOutputType(EncryptAttributesConfigMapOutput{})
}
