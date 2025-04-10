// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create SQL Server account DB attachment
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
// 		_, err := Sqlserver.NewAccountDbAttachment(ctx, "foo", &Sqlserver.AccountDbAttachmentArgs{
// 			InstanceId:  pulumi.String("mssql-3cdq7kx5"),
// 			AccountName: pulumi.Any(tencentcloud_sqlserver_account.Example.Name),
// 			DbName:      pulumi.Any(tencentcloud_sqlserver_db.Example.Name),
// 			Privilege:   pulumi.String("ReadWrite"),
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
// SQL Server account DB attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/accountDbAttachment:AccountDbAttachment foo mssql-3cdq7kx5#tf_sqlserver_account#test111
// ```
type AccountDbAttachment struct {
	pulumi.CustomResourceState

	// SQL Server account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// SQL Server DB name.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// SQL Server instance ID that the account belongs to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
	Privilege pulumi.StringOutput `pulumi:"privilege"`
}

// NewAccountDbAttachment registers a new resource with the given unique name, arguments, and options.
func NewAccountDbAttachment(ctx *pulumi.Context,
	name string, args *AccountDbAttachmentArgs, opts ...pulumi.ResourceOption) (*AccountDbAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Privilege == nil {
		return nil, errors.New("invalid value for required argument 'Privilege'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AccountDbAttachment
	err := ctx.RegisterResource("tencentcloud:Sqlserver/accountDbAttachment:AccountDbAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountDbAttachment gets an existing AccountDbAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountDbAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountDbAttachmentState, opts ...pulumi.ResourceOption) (*AccountDbAttachment, error) {
	var resource AccountDbAttachment
	err := ctx.ReadResource("tencentcloud:Sqlserver/accountDbAttachment:AccountDbAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountDbAttachment resources.
type accountDbAttachmentState struct {
	// SQL Server account name.
	AccountName *string `pulumi:"accountName"`
	// SQL Server DB name.
	DbName *string `pulumi:"dbName"`
	// SQL Server instance ID that the account belongs to.
	InstanceId *string `pulumi:"instanceId"`
	// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
	Privilege *string `pulumi:"privilege"`
}

type AccountDbAttachmentState struct {
	// SQL Server account name.
	AccountName pulumi.StringPtrInput
	// SQL Server DB name.
	DbName pulumi.StringPtrInput
	// SQL Server instance ID that the account belongs to.
	InstanceId pulumi.StringPtrInput
	// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
	Privilege pulumi.StringPtrInput
}

func (AccountDbAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountDbAttachmentState)(nil)).Elem()
}

type accountDbAttachmentArgs struct {
	// SQL Server account name.
	AccountName string `pulumi:"accountName"`
	// SQL Server DB name.
	DbName string `pulumi:"dbName"`
	// SQL Server instance ID that the account belongs to.
	InstanceId string `pulumi:"instanceId"`
	// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
	Privilege string `pulumi:"privilege"`
}

// The set of arguments for constructing a AccountDbAttachment resource.
type AccountDbAttachmentArgs struct {
	// SQL Server account name.
	AccountName pulumi.StringInput
	// SQL Server DB name.
	DbName pulumi.StringInput
	// SQL Server instance ID that the account belongs to.
	InstanceId pulumi.StringInput
	// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
	Privilege pulumi.StringInput
}

func (AccountDbAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountDbAttachmentArgs)(nil)).Elem()
}

type AccountDbAttachmentInput interface {
	pulumi.Input

	ToAccountDbAttachmentOutput() AccountDbAttachmentOutput
	ToAccountDbAttachmentOutputWithContext(ctx context.Context) AccountDbAttachmentOutput
}

func (*AccountDbAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountDbAttachment)(nil)).Elem()
}

func (i *AccountDbAttachment) ToAccountDbAttachmentOutput() AccountDbAttachmentOutput {
	return i.ToAccountDbAttachmentOutputWithContext(context.Background())
}

func (i *AccountDbAttachment) ToAccountDbAttachmentOutputWithContext(ctx context.Context) AccountDbAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountDbAttachmentOutput)
}

// AccountDbAttachmentArrayInput is an input type that accepts AccountDbAttachmentArray and AccountDbAttachmentArrayOutput values.
// You can construct a concrete instance of `AccountDbAttachmentArrayInput` via:
//
//          AccountDbAttachmentArray{ AccountDbAttachmentArgs{...} }
type AccountDbAttachmentArrayInput interface {
	pulumi.Input

	ToAccountDbAttachmentArrayOutput() AccountDbAttachmentArrayOutput
	ToAccountDbAttachmentArrayOutputWithContext(context.Context) AccountDbAttachmentArrayOutput
}

type AccountDbAttachmentArray []AccountDbAttachmentInput

func (AccountDbAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountDbAttachment)(nil)).Elem()
}

func (i AccountDbAttachmentArray) ToAccountDbAttachmentArrayOutput() AccountDbAttachmentArrayOutput {
	return i.ToAccountDbAttachmentArrayOutputWithContext(context.Background())
}

func (i AccountDbAttachmentArray) ToAccountDbAttachmentArrayOutputWithContext(ctx context.Context) AccountDbAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountDbAttachmentArrayOutput)
}

// AccountDbAttachmentMapInput is an input type that accepts AccountDbAttachmentMap and AccountDbAttachmentMapOutput values.
// You can construct a concrete instance of `AccountDbAttachmentMapInput` via:
//
//          AccountDbAttachmentMap{ "key": AccountDbAttachmentArgs{...} }
type AccountDbAttachmentMapInput interface {
	pulumi.Input

	ToAccountDbAttachmentMapOutput() AccountDbAttachmentMapOutput
	ToAccountDbAttachmentMapOutputWithContext(context.Context) AccountDbAttachmentMapOutput
}

type AccountDbAttachmentMap map[string]AccountDbAttachmentInput

func (AccountDbAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountDbAttachment)(nil)).Elem()
}

func (i AccountDbAttachmentMap) ToAccountDbAttachmentMapOutput() AccountDbAttachmentMapOutput {
	return i.ToAccountDbAttachmentMapOutputWithContext(context.Background())
}

func (i AccountDbAttachmentMap) ToAccountDbAttachmentMapOutputWithContext(ctx context.Context) AccountDbAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountDbAttachmentMapOutput)
}

type AccountDbAttachmentOutput struct{ *pulumi.OutputState }

func (AccountDbAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountDbAttachment)(nil)).Elem()
}

func (o AccountDbAttachmentOutput) ToAccountDbAttachmentOutput() AccountDbAttachmentOutput {
	return o
}

func (o AccountDbAttachmentOutput) ToAccountDbAttachmentOutputWithContext(ctx context.Context) AccountDbAttachmentOutput {
	return o
}

// SQL Server account name.
func (o AccountDbAttachmentOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountDbAttachment) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// SQL Server DB name.
func (o AccountDbAttachmentOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountDbAttachment) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// SQL Server instance ID that the account belongs to.
func (o AccountDbAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountDbAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Privilege of the account on DB. Valid values: `ReadOnly`, `ReadWrite`.
func (o AccountDbAttachmentOutput) Privilege() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountDbAttachment) pulumi.StringOutput { return v.Privilege }).(pulumi.StringOutput)
}

type AccountDbAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AccountDbAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountDbAttachment)(nil)).Elem()
}

func (o AccountDbAttachmentArrayOutput) ToAccountDbAttachmentArrayOutput() AccountDbAttachmentArrayOutput {
	return o
}

func (o AccountDbAttachmentArrayOutput) ToAccountDbAttachmentArrayOutputWithContext(ctx context.Context) AccountDbAttachmentArrayOutput {
	return o
}

func (o AccountDbAttachmentArrayOutput) Index(i pulumi.IntInput) AccountDbAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountDbAttachment {
		return vs[0].([]*AccountDbAttachment)[vs[1].(int)]
	}).(AccountDbAttachmentOutput)
}

type AccountDbAttachmentMapOutput struct{ *pulumi.OutputState }

func (AccountDbAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountDbAttachment)(nil)).Elem()
}

func (o AccountDbAttachmentMapOutput) ToAccountDbAttachmentMapOutput() AccountDbAttachmentMapOutput {
	return o
}

func (o AccountDbAttachmentMapOutput) ToAccountDbAttachmentMapOutputWithContext(ctx context.Context) AccountDbAttachmentMapOutput {
	return o
}

func (o AccountDbAttachmentMapOutput) MapIndex(k pulumi.StringInput) AccountDbAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountDbAttachment {
		return vs[0].(map[string]*AccountDbAttachment)[vs[1].(string)]
	}).(AccountDbAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountDbAttachmentInput)(nil)).Elem(), &AccountDbAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountDbAttachmentArrayInput)(nil)).Elem(), AccountDbAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountDbAttachmentMapInput)(nil)).Elem(), AccountDbAttachmentMap{})
	pulumi.RegisterOutputType(AccountDbAttachmentOutput{})
	pulumi.RegisterOutputType(AccountDbAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AccountDbAttachmentMapOutput{})
}
