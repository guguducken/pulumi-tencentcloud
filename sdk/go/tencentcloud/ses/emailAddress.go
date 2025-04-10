// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ses emailAddress
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ses"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ses.NewEmailAddress(ctx, "emailAddress", &Ses.EmailAddressArgs{
// 			EmailAddress:    pulumi.String("aaa@iac-tf.cloud"),
// 			EmailSenderName: pulumi.String("aaa"),
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
// ses email_address can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ses/emailAddress:EmailAddress email_address aaa@iac-tf.cloud
// ```
type EmailAddress struct {
	pulumi.CustomResourceState

	// Your sender address. (You can create up to 10 sender addresses for each domain.).
	EmailAddress pulumi.StringOutput `pulumi:"emailAddress"`
	// Sender name.
	EmailSenderName pulumi.StringPtrOutput `pulumi:"emailSenderName"`
}

// NewEmailAddress registers a new resource with the given unique name, arguments, and options.
func NewEmailAddress(ctx *pulumi.Context,
	name string, args *EmailAddressArgs, opts ...pulumi.ResourceOption) (*EmailAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EmailAddress == nil {
		return nil, errors.New("invalid value for required argument 'EmailAddress'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EmailAddress
	err := ctx.RegisterResource("tencentcloud:Ses/emailAddress:EmailAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailAddress gets an existing EmailAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailAddressState, opts ...pulumi.ResourceOption) (*EmailAddress, error) {
	var resource EmailAddress
	err := ctx.ReadResource("tencentcloud:Ses/emailAddress:EmailAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailAddress resources.
type emailAddressState struct {
	// Your sender address. (You can create up to 10 sender addresses for each domain.).
	EmailAddress *string `pulumi:"emailAddress"`
	// Sender name.
	EmailSenderName *string `pulumi:"emailSenderName"`
}

type EmailAddressState struct {
	// Your sender address. (You can create up to 10 sender addresses for each domain.).
	EmailAddress pulumi.StringPtrInput
	// Sender name.
	EmailSenderName pulumi.StringPtrInput
}

func (EmailAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailAddressState)(nil)).Elem()
}

type emailAddressArgs struct {
	// Your sender address. (You can create up to 10 sender addresses for each domain.).
	EmailAddress string `pulumi:"emailAddress"`
	// Sender name.
	EmailSenderName *string `pulumi:"emailSenderName"`
}

// The set of arguments for constructing a EmailAddress resource.
type EmailAddressArgs struct {
	// Your sender address. (You can create up to 10 sender addresses for each domain.).
	EmailAddress pulumi.StringInput
	// Sender name.
	EmailSenderName pulumi.StringPtrInput
}

func (EmailAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailAddressArgs)(nil)).Elem()
}

type EmailAddressInput interface {
	pulumi.Input

	ToEmailAddressOutput() EmailAddressOutput
	ToEmailAddressOutputWithContext(ctx context.Context) EmailAddressOutput
}

func (*EmailAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailAddress)(nil)).Elem()
}

func (i *EmailAddress) ToEmailAddressOutput() EmailAddressOutput {
	return i.ToEmailAddressOutputWithContext(context.Background())
}

func (i *EmailAddress) ToEmailAddressOutputWithContext(ctx context.Context) EmailAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailAddressOutput)
}

// EmailAddressArrayInput is an input type that accepts EmailAddressArray and EmailAddressArrayOutput values.
// You can construct a concrete instance of `EmailAddressArrayInput` via:
//
//          EmailAddressArray{ EmailAddressArgs{...} }
type EmailAddressArrayInput interface {
	pulumi.Input

	ToEmailAddressArrayOutput() EmailAddressArrayOutput
	ToEmailAddressArrayOutputWithContext(context.Context) EmailAddressArrayOutput
}

type EmailAddressArray []EmailAddressInput

func (EmailAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailAddress)(nil)).Elem()
}

func (i EmailAddressArray) ToEmailAddressArrayOutput() EmailAddressArrayOutput {
	return i.ToEmailAddressArrayOutputWithContext(context.Background())
}

func (i EmailAddressArray) ToEmailAddressArrayOutputWithContext(ctx context.Context) EmailAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailAddressArrayOutput)
}

// EmailAddressMapInput is an input type that accepts EmailAddressMap and EmailAddressMapOutput values.
// You can construct a concrete instance of `EmailAddressMapInput` via:
//
//          EmailAddressMap{ "key": EmailAddressArgs{...} }
type EmailAddressMapInput interface {
	pulumi.Input

	ToEmailAddressMapOutput() EmailAddressMapOutput
	ToEmailAddressMapOutputWithContext(context.Context) EmailAddressMapOutput
}

type EmailAddressMap map[string]EmailAddressInput

func (EmailAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailAddress)(nil)).Elem()
}

func (i EmailAddressMap) ToEmailAddressMapOutput() EmailAddressMapOutput {
	return i.ToEmailAddressMapOutputWithContext(context.Background())
}

func (i EmailAddressMap) ToEmailAddressMapOutputWithContext(ctx context.Context) EmailAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailAddressMapOutput)
}

type EmailAddressOutput struct{ *pulumi.OutputState }

func (EmailAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailAddress)(nil)).Elem()
}

func (o EmailAddressOutput) ToEmailAddressOutput() EmailAddressOutput {
	return o
}

func (o EmailAddressOutput) ToEmailAddressOutputWithContext(ctx context.Context) EmailAddressOutput {
	return o
}

// Your sender address. (You can create up to 10 sender addresses for each domain.).
func (o EmailAddressOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailAddress) pulumi.StringOutput { return v.EmailAddress }).(pulumi.StringOutput)
}

// Sender name.
func (o EmailAddressOutput) EmailSenderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailAddress) pulumi.StringPtrOutput { return v.EmailSenderName }).(pulumi.StringPtrOutput)
}

type EmailAddressArrayOutput struct{ *pulumi.OutputState }

func (EmailAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailAddress)(nil)).Elem()
}

func (o EmailAddressArrayOutput) ToEmailAddressArrayOutput() EmailAddressArrayOutput {
	return o
}

func (o EmailAddressArrayOutput) ToEmailAddressArrayOutputWithContext(ctx context.Context) EmailAddressArrayOutput {
	return o
}

func (o EmailAddressArrayOutput) Index(i pulumi.IntInput) EmailAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailAddress {
		return vs[0].([]*EmailAddress)[vs[1].(int)]
	}).(EmailAddressOutput)
}

type EmailAddressMapOutput struct{ *pulumi.OutputState }

func (EmailAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailAddress)(nil)).Elem()
}

func (o EmailAddressMapOutput) ToEmailAddressMapOutput() EmailAddressMapOutput {
	return o
}

func (o EmailAddressMapOutput) ToEmailAddressMapOutputWithContext(ctx context.Context) EmailAddressMapOutput {
	return o
}

func (o EmailAddressMapOutput) MapIndex(k pulumi.StringInput) EmailAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailAddress {
		return vs[0].(map[string]*EmailAddress)[vs[1].(string)]
	}).(EmailAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailAddressInput)(nil)).Elem(), &EmailAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailAddressArrayInput)(nil)).Elem(), EmailAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailAddressMapInput)(nil)).Elem(), EmailAddressMap{})
	pulumi.RegisterOutputType(EmailAddressOutput{})
	pulumi.RegisterOutputType(EmailAddressArrayOutput{})
	pulumi.RegisterOutputType(EmailAddressMapOutput{})
}
