// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SQL Server DB resource belongs to SQL Server instance.
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
// 		_, err := Sqlserver.NewDb(ctx, "example", &Sqlserver.DbArgs{
// 			InstanceId: pulumi.Any(tencentcloud_sqlserver_instance.Example.Id),
// 			Charset:    pulumi.String("Chinese_PRC_BIN"),
// 			Remark:     pulumi.String("test-remark"),
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
// SQL Server DB can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/db:Db foo mssql-3cdq7kx5#db_name
// ```
type Db struct {
	pulumi.CustomResourceState

	// Character set DB uses. Valid values: `Chinese_PRC_CI_AS`, `Chinese_PRC_CS_AS`, `Chinese_PRC_BIN`, `Chinese_Taiwan_Stroke_CI_AS`, `SQL_Latin1_General_CP1_CI_AS`, and `SQL_Latin1_General_CP1_CS_AS`. Default value is `Chinese_PRC_CI_AS`.
	Charset pulumi.StringPtrOutput `pulumi:"charset"`
	// Database creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// SQL Server instance ID which DB belongs to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.
	Name pulumi.StringOutput `pulumi:"name"`
	// Remark of the DB.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Database status, could be `creating`, `running`, `modifying` which means changing the remark, and `deleting`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewDb registers a new resource with the given unique name, arguments, and options.
func NewDb(ctx *pulumi.Context,
	name string, args *DbArgs, opts ...pulumi.ResourceOption) (*Db, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Db
	err := ctx.RegisterResource("tencentcloud:Sqlserver/db:Db", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDb gets an existing Db resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbState, opts ...pulumi.ResourceOption) (*Db, error) {
	var resource Db
	err := ctx.ReadResource("tencentcloud:Sqlserver/db:Db", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Db resources.
type dbState struct {
	// Character set DB uses. Valid values: `Chinese_PRC_CI_AS`, `Chinese_PRC_CS_AS`, `Chinese_PRC_BIN`, `Chinese_Taiwan_Stroke_CI_AS`, `SQL_Latin1_General_CP1_CI_AS`, and `SQL_Latin1_General_CP1_CS_AS`. Default value is `Chinese_PRC_CI_AS`.
	Charset *string `pulumi:"charset"`
	// Database creation time.
	CreateTime *string `pulumi:"createTime"`
	// SQL Server instance ID which DB belongs to.
	InstanceId *string `pulumi:"instanceId"`
	// Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.
	Name *string `pulumi:"name"`
	// Remark of the DB.
	Remark *string `pulumi:"remark"`
	// Database status, could be `creating`, `running`, `modifying` which means changing the remark, and `deleting`.
	Status *string `pulumi:"status"`
}

type DbState struct {
	// Character set DB uses. Valid values: `Chinese_PRC_CI_AS`, `Chinese_PRC_CS_AS`, `Chinese_PRC_BIN`, `Chinese_Taiwan_Stroke_CI_AS`, `SQL_Latin1_General_CP1_CI_AS`, and `SQL_Latin1_General_CP1_CS_AS`. Default value is `Chinese_PRC_CI_AS`.
	Charset pulumi.StringPtrInput
	// Database creation time.
	CreateTime pulumi.StringPtrInput
	// SQL Server instance ID which DB belongs to.
	InstanceId pulumi.StringPtrInput
	// Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.
	Name pulumi.StringPtrInput
	// Remark of the DB.
	Remark pulumi.StringPtrInput
	// Database status, could be `creating`, `running`, `modifying` which means changing the remark, and `deleting`.
	Status pulumi.StringPtrInput
}

func (DbState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbState)(nil)).Elem()
}

type dbArgs struct {
	// Character set DB uses. Valid values: `Chinese_PRC_CI_AS`, `Chinese_PRC_CS_AS`, `Chinese_PRC_BIN`, `Chinese_Taiwan_Stroke_CI_AS`, `SQL_Latin1_General_CP1_CI_AS`, and `SQL_Latin1_General_CP1_CS_AS`. Default value is `Chinese_PRC_CI_AS`.
	Charset *string `pulumi:"charset"`
	// SQL Server instance ID which DB belongs to.
	InstanceId string `pulumi:"instanceId"`
	// Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.
	Name *string `pulumi:"name"`
	// Remark of the DB.
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a Db resource.
type DbArgs struct {
	// Character set DB uses. Valid values: `Chinese_PRC_CI_AS`, `Chinese_PRC_CS_AS`, `Chinese_PRC_BIN`, `Chinese_Taiwan_Stroke_CI_AS`, `SQL_Latin1_General_CP1_CI_AS`, and `SQL_Latin1_General_CP1_CS_AS`. Default value is `Chinese_PRC_CI_AS`.
	Charset pulumi.StringPtrInput
	// SQL Server instance ID which DB belongs to.
	InstanceId pulumi.StringInput
	// Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.
	Name pulumi.StringPtrInput
	// Remark of the DB.
	Remark pulumi.StringPtrInput
}

func (DbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbArgs)(nil)).Elem()
}

type DbInput interface {
	pulumi.Input

	ToDbOutput() DbOutput
	ToDbOutputWithContext(ctx context.Context) DbOutput
}

func (*Db) ElementType() reflect.Type {
	return reflect.TypeOf((**Db)(nil)).Elem()
}

func (i *Db) ToDbOutput() DbOutput {
	return i.ToDbOutputWithContext(context.Background())
}

func (i *Db) ToDbOutputWithContext(ctx context.Context) DbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbOutput)
}

// DbArrayInput is an input type that accepts DbArray and DbArrayOutput values.
// You can construct a concrete instance of `DbArrayInput` via:
//
//          DbArray{ DbArgs{...} }
type DbArrayInput interface {
	pulumi.Input

	ToDbArrayOutput() DbArrayOutput
	ToDbArrayOutputWithContext(context.Context) DbArrayOutput
}

type DbArray []DbInput

func (DbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Db)(nil)).Elem()
}

func (i DbArray) ToDbArrayOutput() DbArrayOutput {
	return i.ToDbArrayOutputWithContext(context.Background())
}

func (i DbArray) ToDbArrayOutputWithContext(ctx context.Context) DbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbArrayOutput)
}

// DbMapInput is an input type that accepts DbMap and DbMapOutput values.
// You can construct a concrete instance of `DbMapInput` via:
//
//          DbMap{ "key": DbArgs{...} }
type DbMapInput interface {
	pulumi.Input

	ToDbMapOutput() DbMapOutput
	ToDbMapOutputWithContext(context.Context) DbMapOutput
}

type DbMap map[string]DbInput

func (DbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Db)(nil)).Elem()
}

func (i DbMap) ToDbMapOutput() DbMapOutput {
	return i.ToDbMapOutputWithContext(context.Background())
}

func (i DbMap) ToDbMapOutputWithContext(ctx context.Context) DbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbMapOutput)
}

type DbOutput struct{ *pulumi.OutputState }

func (DbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Db)(nil)).Elem()
}

func (o DbOutput) ToDbOutput() DbOutput {
	return o
}

func (o DbOutput) ToDbOutputWithContext(ctx context.Context) DbOutput {
	return o
}

// Character set DB uses. Valid values: `Chinese_PRC_CI_AS`, `Chinese_PRC_CS_AS`, `Chinese_PRC_BIN`, `Chinese_Taiwan_Stroke_CI_AS`, `SQL_Latin1_General_CP1_CI_AS`, and `SQL_Latin1_General_CP1_CS_AS`. Default value is `Chinese_PRC_CI_AS`.
func (o DbOutput) Charset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Db) pulumi.StringPtrOutput { return v.Charset }).(pulumi.StringPtrOutput)
}

// Database creation time.
func (o DbOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Db) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// SQL Server instance ID which DB belongs to.
func (o DbOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Db) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.
func (o DbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Db) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Remark of the DB.
func (o DbOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Db) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Database status, could be `creating`, `running`, `modifying` which means changing the remark, and `deleting`.
func (o DbOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Db) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type DbArrayOutput struct{ *pulumi.OutputState }

func (DbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Db)(nil)).Elem()
}

func (o DbArrayOutput) ToDbArrayOutput() DbArrayOutput {
	return o
}

func (o DbArrayOutput) ToDbArrayOutputWithContext(ctx context.Context) DbArrayOutput {
	return o
}

func (o DbArrayOutput) Index(i pulumi.IntInput) DbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Db {
		return vs[0].([]*Db)[vs[1].(int)]
	}).(DbOutput)
}

type DbMapOutput struct{ *pulumi.OutputState }

func (DbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Db)(nil)).Elem()
}

func (o DbMapOutput) ToDbMapOutput() DbMapOutput {
	return o
}

func (o DbMapOutput) ToDbMapOutputWithContext(ctx context.Context) DbMapOutput {
	return o
}

func (o DbMapOutput) MapIndex(k pulumi.StringInput) DbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Db {
		return vs[0].(map[string]*Db)[vs[1].(string)]
	}).(DbOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbInput)(nil)).Elem(), &Db{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbArrayInput)(nil)).Elem(), DbArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbMapInput)(nil)).Elem(), DbMap{})
	pulumi.RegisterOutputType(DbOutput{})
	pulumi.RegisterOutputType(DbArrayOutput{})
	pulumi.RegisterOutputType(DbMapOutput{})
}
