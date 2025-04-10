// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dbbrain sql_filter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		availabilityZone := "ap-guangzhou-3"
// 		if param := cfg.Get("availabilityZone"); param != "" {
// 			availabilityZone = param
// 		}
// 		region := "ap-guangzhou"
// 		if param := cfg.Get("region"); param != "" {
// 			region = param
// 		}
// 		mysql, err := Mysql.GetInstance(ctx, &mysql.GetInstanceArgs{
// 			InstanceName: pulumi.StringRef("instance_name"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		mysqlId := mysql.InstanceLists[0].MysqlId
// 		_, err = Dbbrain.NewSqlFilter(ctx, "sqlFilter", &Dbbrain.SqlFilterArgs{
// 			InstanceId: pulumi.String(mysqlId),
// 			SessionToken: &dbbrain.SqlFilterSessionTokenArgs{
// 				User:     pulumi.String("test"),
// 				Password: pulumi.String("===password==="),
// 			},
// 			SqlType:        pulumi.String("SELECT"),
// 			FilterKey:      pulumi.String("filter_key"),
// 			MaxConcurrency: pulumi.Int(10),
// 			Duration:       pulumi.Int(3600),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SqlFilter struct {
	pulumi.CustomResourceState

	// filter duration.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// filter id.
	FilterId pulumi.IntOutput `pulumi:"filterId"`
	// filter key.
	FilterKey pulumi.StringOutput `pulumi:"filterKey"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// maximum concurreny.
	MaxConcurrency pulumi.IntOutput `pulumi:"maxConcurrency"`
	// product, optional value is &amp;#39;mysql&amp;#39;, &amp;#39;cynosdb&amp;#39;.
	Product pulumi.StringPtrOutput `pulumi:"product"`
	// session token.
	SessionToken SqlFilterSessionTokenOutput `pulumi:"sessionToken"`
	// sql type, optional value is SELECT, UPDATE, DELETE, INSERT, REPLACE.
	SqlType pulumi.StringOutput `pulumi:"sqlType"`
	// filter status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSqlFilter registers a new resource with the given unique name, arguments, and options.
func NewSqlFilter(ctx *pulumi.Context,
	name string, args *SqlFilterArgs, opts ...pulumi.ResourceOption) (*SqlFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.FilterKey == nil {
		return nil, errors.New("invalid value for required argument 'FilterKey'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.MaxConcurrency == nil {
		return nil, errors.New("invalid value for required argument 'MaxConcurrency'")
	}
	if args.SessionToken == nil {
		return nil, errors.New("invalid value for required argument 'SessionToken'")
	}
	if args.SqlType == nil {
		return nil, errors.New("invalid value for required argument 'SqlType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SqlFilter
	err := ctx.RegisterResource("tencentcloud:Dbbrain/sqlFilter:SqlFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlFilter gets an existing SqlFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlFilterState, opts ...pulumi.ResourceOption) (*SqlFilter, error) {
	var resource SqlFilter
	err := ctx.ReadResource("tencentcloud:Dbbrain/sqlFilter:SqlFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlFilter resources.
type sqlFilterState struct {
	// filter duration.
	Duration *int `pulumi:"duration"`
	// filter id.
	FilterId *int `pulumi:"filterId"`
	// filter key.
	FilterKey *string `pulumi:"filterKey"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// maximum concurreny.
	MaxConcurrency *int `pulumi:"maxConcurrency"`
	// product, optional value is &amp;#39;mysql&amp;#39;, &amp;#39;cynosdb&amp;#39;.
	Product *string `pulumi:"product"`
	// session token.
	SessionToken *SqlFilterSessionToken `pulumi:"sessionToken"`
	// sql type, optional value is SELECT, UPDATE, DELETE, INSERT, REPLACE.
	SqlType *string `pulumi:"sqlType"`
	// filter status.
	Status *string `pulumi:"status"`
}

type SqlFilterState struct {
	// filter duration.
	Duration pulumi.IntPtrInput
	// filter id.
	FilterId pulumi.IntPtrInput
	// filter key.
	FilterKey pulumi.StringPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
	// maximum concurreny.
	MaxConcurrency pulumi.IntPtrInput
	// product, optional value is &amp;#39;mysql&amp;#39;, &amp;#39;cynosdb&amp;#39;.
	Product pulumi.StringPtrInput
	// session token.
	SessionToken SqlFilterSessionTokenPtrInput
	// sql type, optional value is SELECT, UPDATE, DELETE, INSERT, REPLACE.
	SqlType pulumi.StringPtrInput
	// filter status.
	Status pulumi.StringPtrInput
}

func (SqlFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlFilterState)(nil)).Elem()
}

type sqlFilterArgs struct {
	// filter duration.
	Duration int `pulumi:"duration"`
	// filter key.
	FilterKey string `pulumi:"filterKey"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// maximum concurreny.
	MaxConcurrency int `pulumi:"maxConcurrency"`
	// product, optional value is &amp;#39;mysql&amp;#39;, &amp;#39;cynosdb&amp;#39;.
	Product *string `pulumi:"product"`
	// session token.
	SessionToken SqlFilterSessionToken `pulumi:"sessionToken"`
	// sql type, optional value is SELECT, UPDATE, DELETE, INSERT, REPLACE.
	SqlType string `pulumi:"sqlType"`
	// filter status.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a SqlFilter resource.
type SqlFilterArgs struct {
	// filter duration.
	Duration pulumi.IntInput
	// filter key.
	FilterKey pulumi.StringInput
	// instance id.
	InstanceId pulumi.StringInput
	// maximum concurreny.
	MaxConcurrency pulumi.IntInput
	// product, optional value is &amp;#39;mysql&amp;#39;, &amp;#39;cynosdb&amp;#39;.
	Product pulumi.StringPtrInput
	// session token.
	SessionToken SqlFilterSessionTokenInput
	// sql type, optional value is SELECT, UPDATE, DELETE, INSERT, REPLACE.
	SqlType pulumi.StringInput
	// filter status.
	Status pulumi.StringPtrInput
}

func (SqlFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlFilterArgs)(nil)).Elem()
}

type SqlFilterInput interface {
	pulumi.Input

	ToSqlFilterOutput() SqlFilterOutput
	ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput
}

func (*SqlFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (i *SqlFilter) ToSqlFilterOutput() SqlFilterOutput {
	return i.ToSqlFilterOutputWithContext(context.Background())
}

func (i *SqlFilter) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput)
}

// SqlFilterArrayInput is an input type that accepts SqlFilterArray and SqlFilterArrayOutput values.
// You can construct a concrete instance of `SqlFilterArrayInput` via:
//
//          SqlFilterArray{ SqlFilterArgs{...} }
type SqlFilterArrayInput interface {
	pulumi.Input

	ToSqlFilterArrayOutput() SqlFilterArrayOutput
	ToSqlFilterArrayOutputWithContext(context.Context) SqlFilterArrayOutput
}

type SqlFilterArray []SqlFilterInput

func (SqlFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlFilter)(nil)).Elem()
}

func (i SqlFilterArray) ToSqlFilterArrayOutput() SqlFilterArrayOutput {
	return i.ToSqlFilterArrayOutputWithContext(context.Background())
}

func (i SqlFilterArray) ToSqlFilterArrayOutputWithContext(ctx context.Context) SqlFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterArrayOutput)
}

// SqlFilterMapInput is an input type that accepts SqlFilterMap and SqlFilterMapOutput values.
// You can construct a concrete instance of `SqlFilterMapInput` via:
//
//          SqlFilterMap{ "key": SqlFilterArgs{...} }
type SqlFilterMapInput interface {
	pulumi.Input

	ToSqlFilterMapOutput() SqlFilterMapOutput
	ToSqlFilterMapOutputWithContext(context.Context) SqlFilterMapOutput
}

type SqlFilterMap map[string]SqlFilterInput

func (SqlFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlFilter)(nil)).Elem()
}

func (i SqlFilterMap) ToSqlFilterMapOutput() SqlFilterMapOutput {
	return i.ToSqlFilterMapOutputWithContext(context.Background())
}

func (i SqlFilterMap) ToSqlFilterMapOutputWithContext(ctx context.Context) SqlFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterMapOutput)
}

type SqlFilterOutput struct{ *pulumi.OutputState }

func (SqlFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (o SqlFilterOutput) ToSqlFilterOutput() SqlFilterOutput {
	return o
}

func (o SqlFilterOutput) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return o
}

// filter duration.
func (o SqlFilterOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.IntOutput { return v.Duration }).(pulumi.IntOutput)
}

// filter id.
func (o SqlFilterOutput) FilterId() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.IntOutput { return v.FilterId }).(pulumi.IntOutput)
}

// filter key.
func (o SqlFilterOutput) FilterKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.StringOutput { return v.FilterKey }).(pulumi.StringOutput)
}

// instance id.
func (o SqlFilterOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// maximum concurreny.
func (o SqlFilterOutput) MaxConcurrency() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.IntOutput { return v.MaxConcurrency }).(pulumi.IntOutput)
}

// product, optional value is &amp;#39;mysql&amp;#39;, &amp;#39;cynosdb&amp;#39;.
func (o SqlFilterOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.StringPtrOutput { return v.Product }).(pulumi.StringPtrOutput)
}

// session token.
func (o SqlFilterOutput) SessionToken() SqlFilterSessionTokenOutput {
	return o.ApplyT(func(v *SqlFilter) SqlFilterSessionTokenOutput { return v.SessionToken }).(SqlFilterSessionTokenOutput)
}

// sql type, optional value is SELECT, UPDATE, DELETE, INSERT, REPLACE.
func (o SqlFilterOutput) SqlType() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.StringOutput { return v.SqlType }).(pulumi.StringOutput)
}

// filter status.
func (o SqlFilterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlFilter) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SqlFilterArrayOutput struct{ *pulumi.OutputState }

func (SqlFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlFilter)(nil)).Elem()
}

func (o SqlFilterArrayOutput) ToSqlFilterArrayOutput() SqlFilterArrayOutput {
	return o
}

func (o SqlFilterArrayOutput) ToSqlFilterArrayOutputWithContext(ctx context.Context) SqlFilterArrayOutput {
	return o
}

func (o SqlFilterArrayOutput) Index(i pulumi.IntInput) SqlFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqlFilter {
		return vs[0].([]*SqlFilter)[vs[1].(int)]
	}).(SqlFilterOutput)
}

type SqlFilterMapOutput struct{ *pulumi.OutputState }

func (SqlFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlFilter)(nil)).Elem()
}

func (o SqlFilterMapOutput) ToSqlFilterMapOutput() SqlFilterMapOutput {
	return o
}

func (o SqlFilterMapOutput) ToSqlFilterMapOutputWithContext(ctx context.Context) SqlFilterMapOutput {
	return o
}

func (o SqlFilterMapOutput) MapIndex(k pulumi.StringInput) SqlFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqlFilter {
		return vs[0].(map[string]*SqlFilter)[vs[1].(string)]
	}).(SqlFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlFilterInput)(nil)).Elem(), &SqlFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlFilterArrayInput)(nil)).Elem(), SqlFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlFilterMapInput)(nil)).Elem(), SqlFilterMap{})
	pulumi.RegisterOutputType(SqlFilterOutput{})
	pulumi.RegisterOutputType(SqlFilterArrayOutput{})
	pulumi.RegisterOutputType(SqlFilterMapOutput{})
}
