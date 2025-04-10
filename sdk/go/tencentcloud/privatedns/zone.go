// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a Private Dns Zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/PrivateDns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/PrivateDns"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := PrivateDns.NewZone(ctx, "foo", &PrivateDns.ZoneArgs{
// 			AccountVpcSets: privatedns.ZoneAccountVpcSetArray{
// 				&privatedns.ZoneAccountVpcSetArgs{
// 					Region:    pulumi.String("ap-guangzhou"),
// 					Uin:       pulumi.String("454xxxxxxx"),
// 					UniqVpcId: pulumi.String("vpc-xxxxx"),
// 					VpcName:   pulumi.String("test-redis"),
// 				},
// 			},
// 			DnsForwardStatus: pulumi.String("DISABLED"),
// 			Domain:           pulumi.String("domain.com"),
// 			Remark:           pulumi.String("test"),
// 			Tags: pulumi.AnyMap{
// 				"created_by": pulumi.Any{
// 					nil,
// 				},
// 				"terraform": pulumi.Any{
// 					nil,
// 				},
// 			},
// 			VpcSets: privatedns.ZoneVpcSetArray{
// 				&privatedns.ZoneVpcSetArgs{
// 					Region:    pulumi.String("ap-guangzhou"),
// 					UniqVpcId: pulumi.String("vpc-xxxxx"),
// 				},
// 			},
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
// Private Dns Zone can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:PrivateDns/zone:Zone foo zone_id
// ```
type Zone struct {
	pulumi.CustomResourceState

	// List of authorized accounts' VPCs to associate with the private domain.
	AccountVpcSets ZoneAccountVpcSetArrayOutput `pulumi:"accountVpcSets"`
	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
	DnsForwardStatus pulumi.StringPtrOutput `pulumi:"dnsForwardStatus"`
	// Domain name, which must be in the format of standard TLD.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Remarks.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
	//
	// Deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.
	TagSets ZoneTagSetArrayOutput `pulumi:"tagSets"`
	// Tags of the private dns zone.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Associates the private domain to a VPC when it is created.
	VpcSets ZoneVpcSetArrayOutput `pulumi:"vpcSets"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Zone
	err := ctx.RegisterResource("tencentcloud:PrivateDns/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("tencentcloud:PrivateDns/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// List of authorized accounts' VPCs to associate with the private domain.
	AccountVpcSets []ZoneAccountVpcSet `pulumi:"accountVpcSets"`
	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
	DnsForwardStatus *string `pulumi:"dnsForwardStatus"`
	// Domain name, which must be in the format of standard TLD.
	Domain *string `pulumi:"domain"`
	// Remarks.
	Remark *string `pulumi:"remark"`
	// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
	//
	// Deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.
	TagSets []ZoneTagSet `pulumi:"tagSets"`
	// Tags of the private dns zone.
	Tags map[string]interface{} `pulumi:"tags"`
	// Associates the private domain to a VPC when it is created.
	VpcSets []ZoneVpcSet `pulumi:"vpcSets"`
}

type ZoneState struct {
	// List of authorized accounts' VPCs to associate with the private domain.
	AccountVpcSets ZoneAccountVpcSetArrayInput
	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
	DnsForwardStatus pulumi.StringPtrInput
	// Domain name, which must be in the format of standard TLD.
	Domain pulumi.StringPtrInput
	// Remarks.
	Remark pulumi.StringPtrInput
	// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
	//
	// Deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.
	TagSets ZoneTagSetArrayInput
	// Tags of the private dns zone.
	Tags pulumi.MapInput
	// Associates the private domain to a VPC when it is created.
	VpcSets ZoneVpcSetArrayInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// List of authorized accounts' VPCs to associate with the private domain.
	AccountVpcSets []ZoneAccountVpcSet `pulumi:"accountVpcSets"`
	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
	DnsForwardStatus *string `pulumi:"dnsForwardStatus"`
	// Domain name, which must be in the format of standard TLD.
	Domain string `pulumi:"domain"`
	// Remarks.
	Remark *string `pulumi:"remark"`
	// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
	//
	// Deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.
	TagSets []ZoneTagSet `pulumi:"tagSets"`
	// Tags of the private dns zone.
	Tags map[string]interface{} `pulumi:"tags"`
	// Associates the private domain to a VPC when it is created.
	VpcSets []ZoneVpcSet `pulumi:"vpcSets"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// List of authorized accounts' VPCs to associate with the private domain.
	AccountVpcSets ZoneAccountVpcSetArrayInput
	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
	DnsForwardStatus pulumi.StringPtrInput
	// Domain name, which must be in the format of standard TLD.
	Domain pulumi.StringInput
	// Remarks.
	Remark pulumi.StringPtrInput
	// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
	//
	// Deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.
	TagSets ZoneTagSetArrayInput
	// Tags of the private dns zone.
	Tags pulumi.MapInput
	// Associates the private domain to a VPC when it is created.
	VpcSets ZoneVpcSetArrayInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//          ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//          ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

// List of authorized accounts' VPCs to associate with the private domain.
func (o ZoneOutput) AccountVpcSets() ZoneAccountVpcSetArrayOutput {
	return o.ApplyT(func(v *Zone) ZoneAccountVpcSetArrayOutput { return v.AccountVpcSets }).(ZoneAccountVpcSetArrayOutput)
}

// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
func (o ZoneOutput) DnsForwardStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.DnsForwardStatus }).(pulumi.StringPtrOutput)
}

// Domain name, which must be in the format of standard TLD.
func (o ZoneOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Remarks.
func (o ZoneOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
//
// Deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.
func (o ZoneOutput) TagSets() ZoneTagSetArrayOutput {
	return o.ApplyT(func(v *Zone) ZoneTagSetArrayOutput { return v.TagSets }).(ZoneTagSetArrayOutput)
}

// Tags of the private dns zone.
func (o ZoneOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Zone) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Associates the private domain to a VPC when it is created.
func (o ZoneOutput) VpcSets() ZoneVpcSetArrayOutput {
	return o.ApplyT(func(v *Zone) ZoneVpcSetArrayOutput { return v.VpcSets }).(ZoneVpcSetArrayOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].([]*Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].(map[string]*Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}
