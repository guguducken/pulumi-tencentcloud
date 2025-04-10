// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver configDatabaseCt
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Sqlserver.NewConfigDatabaseCt(ctx, "configDatabaseCt", &Sqlserver.ConfigDatabaseCtArgs{
// 			DbName:     pulumi.String("keep_pubsub_db2"),
// 			InstanceId: pulumi.String("mssql-qelbzgwf"),
// 			ModifyType: pulumi.String("disable"),
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
// sqlserver tencentcloud_sqlserver_config_database_ct can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt config_database_ct config_database_ct_id
// ```
type ConfigDatabaseCt struct {
	pulumi.CustomResourceState

	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
	ChangeRetentionDay pulumi.IntOutput `pulumi:"changeRetentionDay"`
	// database name.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Enable or disable CT. Valid values: enable, disable.
	ModifyType pulumi.StringOutput `pulumi:"modifyType"`
}

// NewConfigDatabaseCt registers a new resource with the given unique name, arguments, and options.
func NewConfigDatabaseCt(ctx *pulumi.Context,
	name string, args *ConfigDatabaseCtArgs, opts ...pulumi.ResourceOption) (*ConfigDatabaseCt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ModifyType == nil {
		return nil, errors.New("invalid value for required argument 'ModifyType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ConfigDatabaseCt
	err := ctx.RegisterResource("tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigDatabaseCt gets an existing ConfigDatabaseCt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigDatabaseCt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigDatabaseCtState, opts ...pulumi.ResourceOption) (*ConfigDatabaseCt, error) {
	var resource ConfigDatabaseCt
	err := ctx.ReadResource("tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigDatabaseCt resources.
type configDatabaseCtState struct {
	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
	ChangeRetentionDay *int `pulumi:"changeRetentionDay"`
	// database name.
	DbName *string `pulumi:"dbName"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Enable or disable CT. Valid values: enable, disable.
	ModifyType *string `pulumi:"modifyType"`
}

type ConfigDatabaseCtState struct {
	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
	ChangeRetentionDay pulumi.IntPtrInput
	// database name.
	DbName pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Enable or disable CT. Valid values: enable, disable.
	ModifyType pulumi.StringPtrInput
}

func (ConfigDatabaseCtState) ElementType() reflect.Type {
	return reflect.TypeOf((*configDatabaseCtState)(nil)).Elem()
}

type configDatabaseCtArgs struct {
	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
	ChangeRetentionDay *int `pulumi:"changeRetentionDay"`
	// database name.
	DbName string `pulumi:"dbName"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Enable or disable CT. Valid values: enable, disable.
	ModifyType string `pulumi:"modifyType"`
}

// The set of arguments for constructing a ConfigDatabaseCt resource.
type ConfigDatabaseCtArgs struct {
	// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
	ChangeRetentionDay pulumi.IntPtrInput
	// database name.
	DbName pulumi.StringInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// Enable or disable CT. Valid values: enable, disable.
	ModifyType pulumi.StringInput
}

func (ConfigDatabaseCtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configDatabaseCtArgs)(nil)).Elem()
}

type ConfigDatabaseCtInput interface {
	pulumi.Input

	ToConfigDatabaseCtOutput() ConfigDatabaseCtOutput
	ToConfigDatabaseCtOutputWithContext(ctx context.Context) ConfigDatabaseCtOutput
}

func (*ConfigDatabaseCt) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigDatabaseCt)(nil)).Elem()
}

func (i *ConfigDatabaseCt) ToConfigDatabaseCtOutput() ConfigDatabaseCtOutput {
	return i.ToConfigDatabaseCtOutputWithContext(context.Background())
}

func (i *ConfigDatabaseCt) ToConfigDatabaseCtOutputWithContext(ctx context.Context) ConfigDatabaseCtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigDatabaseCtOutput)
}

// ConfigDatabaseCtArrayInput is an input type that accepts ConfigDatabaseCtArray and ConfigDatabaseCtArrayOutput values.
// You can construct a concrete instance of `ConfigDatabaseCtArrayInput` via:
//
//          ConfigDatabaseCtArray{ ConfigDatabaseCtArgs{...} }
type ConfigDatabaseCtArrayInput interface {
	pulumi.Input

	ToConfigDatabaseCtArrayOutput() ConfigDatabaseCtArrayOutput
	ToConfigDatabaseCtArrayOutputWithContext(context.Context) ConfigDatabaseCtArrayOutput
}

type ConfigDatabaseCtArray []ConfigDatabaseCtInput

func (ConfigDatabaseCtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigDatabaseCt)(nil)).Elem()
}

func (i ConfigDatabaseCtArray) ToConfigDatabaseCtArrayOutput() ConfigDatabaseCtArrayOutput {
	return i.ToConfigDatabaseCtArrayOutputWithContext(context.Background())
}

func (i ConfigDatabaseCtArray) ToConfigDatabaseCtArrayOutputWithContext(ctx context.Context) ConfigDatabaseCtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigDatabaseCtArrayOutput)
}

// ConfigDatabaseCtMapInput is an input type that accepts ConfigDatabaseCtMap and ConfigDatabaseCtMapOutput values.
// You can construct a concrete instance of `ConfigDatabaseCtMapInput` via:
//
//          ConfigDatabaseCtMap{ "key": ConfigDatabaseCtArgs{...} }
type ConfigDatabaseCtMapInput interface {
	pulumi.Input

	ToConfigDatabaseCtMapOutput() ConfigDatabaseCtMapOutput
	ToConfigDatabaseCtMapOutputWithContext(context.Context) ConfigDatabaseCtMapOutput
}

type ConfigDatabaseCtMap map[string]ConfigDatabaseCtInput

func (ConfigDatabaseCtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigDatabaseCt)(nil)).Elem()
}

func (i ConfigDatabaseCtMap) ToConfigDatabaseCtMapOutput() ConfigDatabaseCtMapOutput {
	return i.ToConfigDatabaseCtMapOutputWithContext(context.Background())
}

func (i ConfigDatabaseCtMap) ToConfigDatabaseCtMapOutputWithContext(ctx context.Context) ConfigDatabaseCtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigDatabaseCtMapOutput)
}

type ConfigDatabaseCtOutput struct{ *pulumi.OutputState }

func (ConfigDatabaseCtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigDatabaseCt)(nil)).Elem()
}

func (o ConfigDatabaseCtOutput) ToConfigDatabaseCtOutput() ConfigDatabaseCtOutput {
	return o
}

func (o ConfigDatabaseCtOutput) ToConfigDatabaseCtOutputWithContext(ctx context.Context) ConfigDatabaseCtOutput {
	return o
}

// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
func (o ConfigDatabaseCtOutput) ChangeRetentionDay() pulumi.IntOutput {
	return o.ApplyT(func(v *ConfigDatabaseCt) pulumi.IntOutput { return v.ChangeRetentionDay }).(pulumi.IntOutput)
}

// database name.
func (o ConfigDatabaseCtOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigDatabaseCt) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// Instance ID.
func (o ConfigDatabaseCtOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigDatabaseCt) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Enable or disable CT. Valid values: enable, disable.
func (o ConfigDatabaseCtOutput) ModifyType() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigDatabaseCt) pulumi.StringOutput { return v.ModifyType }).(pulumi.StringOutput)
}

type ConfigDatabaseCtArrayOutput struct{ *pulumi.OutputState }

func (ConfigDatabaseCtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigDatabaseCt)(nil)).Elem()
}

func (o ConfigDatabaseCtArrayOutput) ToConfigDatabaseCtArrayOutput() ConfigDatabaseCtArrayOutput {
	return o
}

func (o ConfigDatabaseCtArrayOutput) ToConfigDatabaseCtArrayOutputWithContext(ctx context.Context) ConfigDatabaseCtArrayOutput {
	return o
}

func (o ConfigDatabaseCtArrayOutput) Index(i pulumi.IntInput) ConfigDatabaseCtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigDatabaseCt {
		return vs[0].([]*ConfigDatabaseCt)[vs[1].(int)]
	}).(ConfigDatabaseCtOutput)
}

type ConfigDatabaseCtMapOutput struct{ *pulumi.OutputState }

func (ConfigDatabaseCtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigDatabaseCt)(nil)).Elem()
}

func (o ConfigDatabaseCtMapOutput) ToConfigDatabaseCtMapOutput() ConfigDatabaseCtMapOutput {
	return o
}

func (o ConfigDatabaseCtMapOutput) ToConfigDatabaseCtMapOutputWithContext(ctx context.Context) ConfigDatabaseCtMapOutput {
	return o
}

func (o ConfigDatabaseCtMapOutput) MapIndex(k pulumi.StringInput) ConfigDatabaseCtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigDatabaseCt {
		return vs[0].(map[string]*ConfigDatabaseCt)[vs[1].(string)]
	}).(ConfigDatabaseCtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigDatabaseCtInput)(nil)).Elem(), &ConfigDatabaseCt{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigDatabaseCtArrayInput)(nil)).Elem(), ConfigDatabaseCtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigDatabaseCtMapInput)(nil)).Elem(), ConfigDatabaseCtMap{})
	pulumi.RegisterOutputType(ConfigDatabaseCtOutput{})
	pulumi.RegisterOutputType(ConfigDatabaseCtArrayOutput{})
	pulumi.RegisterOutputType(ConfigDatabaseCtMapOutput{})
}
