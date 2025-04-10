// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a css domain
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Css"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Css.NewDomain(ctx, "domain", &Css.DomainArgs{
// 			DomainName:        pulumi.String("iac-tf.cloud"),
// 			DomainType:        pulumi.Int(0),
// 			IsDelayLive:       pulumi.Int(0),
// 			IsMiniProgramLive: pulumi.Int(0),
// 			PlayType:          pulumi.Int(1),
// 			VerifyOwnerType:   pulumi.String("dbCheck"),
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
// css domain can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Css/domain:Domain domain domain_name
// ```
type Domain struct {
	pulumi.CustomResourceState

	// Domain Name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Domain type: `0`: push stream. `1`: playback.
	DomainType pulumi.IntOutput `pulumi:"domainType"`
	// Switch. true: enable the specified domain, false: disable the specified domain.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// Whether it is LCB: `0`: LVB. `1`: LCB. Default value is 0.
	IsDelayLive pulumi.IntPtrOutput `pulumi:"isDelayLive"`
	// `0`: LVB. `1`: LVB on Mini Program. Note: this field may return null, indicating that no valid values can be obtained. Default value is 0.
	IsMiniProgramLive pulumi.IntPtrOutput `pulumi:"isMiniProgramLive"`
	// Play Type. This parameter is valid only if `DomainType` is 1. Available values: `1`: in Mainland China. `2`: global. `3`: outside Mainland China. Default value is 1.
	PlayType pulumi.IntPtrOutput `pulumi:"playType"`
	// Domain name attribution verification type. `dnsCheck`, `fileCheck`, `dbCheck`. The default is `dbCheck`.
	VerifyOwnerType pulumi.StringPtrOutput `pulumi:"verifyOwnerType"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.DomainType == nil {
		return nil, errors.New("invalid value for required argument 'DomainType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("tencentcloud:Css/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("tencentcloud:Css/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Domain Name.
	DomainName *string `pulumi:"domainName"`
	// Domain type: `0`: push stream. `1`: playback.
	DomainType *int `pulumi:"domainType"`
	// Switch. true: enable the specified domain, false: disable the specified domain.
	Enable *bool `pulumi:"enable"`
	// Whether it is LCB: `0`: LVB. `1`: LCB. Default value is 0.
	IsDelayLive *int `pulumi:"isDelayLive"`
	// `0`: LVB. `1`: LVB on Mini Program. Note: this field may return null, indicating that no valid values can be obtained. Default value is 0.
	IsMiniProgramLive *int `pulumi:"isMiniProgramLive"`
	// Play Type. This parameter is valid only if `DomainType` is 1. Available values: `1`: in Mainland China. `2`: global. `3`: outside Mainland China. Default value is 1.
	PlayType *int `pulumi:"playType"`
	// Domain name attribution verification type. `dnsCheck`, `fileCheck`, `dbCheck`. The default is `dbCheck`.
	VerifyOwnerType *string `pulumi:"verifyOwnerType"`
}

type DomainState struct {
	// Domain Name.
	DomainName pulumi.StringPtrInput
	// Domain type: `0`: push stream. `1`: playback.
	DomainType pulumi.IntPtrInput
	// Switch. true: enable the specified domain, false: disable the specified domain.
	Enable pulumi.BoolPtrInput
	// Whether it is LCB: `0`: LVB. `1`: LCB. Default value is 0.
	IsDelayLive pulumi.IntPtrInput
	// `0`: LVB. `1`: LVB on Mini Program. Note: this field may return null, indicating that no valid values can be obtained. Default value is 0.
	IsMiniProgramLive pulumi.IntPtrInput
	// Play Type. This parameter is valid only if `DomainType` is 1. Available values: `1`: in Mainland China. `2`: global. `3`: outside Mainland China. Default value is 1.
	PlayType pulumi.IntPtrInput
	// Domain name attribution verification type. `dnsCheck`, `fileCheck`, `dbCheck`. The default is `dbCheck`.
	VerifyOwnerType pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Domain Name.
	DomainName string `pulumi:"domainName"`
	// Domain type: `0`: push stream. `1`: playback.
	DomainType int `pulumi:"domainType"`
	// Switch. true: enable the specified domain, false: disable the specified domain.
	Enable *bool `pulumi:"enable"`
	// Whether it is LCB: `0`: LVB. `1`: LCB. Default value is 0.
	IsDelayLive *int `pulumi:"isDelayLive"`
	// `0`: LVB. `1`: LVB on Mini Program. Note: this field may return null, indicating that no valid values can be obtained. Default value is 0.
	IsMiniProgramLive *int `pulumi:"isMiniProgramLive"`
	// Play Type. This parameter is valid only if `DomainType` is 1. Available values: `1`: in Mainland China. `2`: global. `3`: outside Mainland China. Default value is 1.
	PlayType *int `pulumi:"playType"`
	// Domain name attribution verification type. `dnsCheck`, `fileCheck`, `dbCheck`. The default is `dbCheck`.
	VerifyOwnerType *string `pulumi:"verifyOwnerType"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Domain Name.
	DomainName pulumi.StringInput
	// Domain type: `0`: push stream. `1`: playback.
	DomainType pulumi.IntInput
	// Switch. true: enable the specified domain, false: disable the specified domain.
	Enable pulumi.BoolPtrInput
	// Whether it is LCB: `0`: LVB. `1`: LCB. Default value is 0.
	IsDelayLive pulumi.IntPtrInput
	// `0`: LVB. `1`: LVB on Mini Program. Note: this field may return null, indicating that no valid values can be obtained. Default value is 0.
	IsMiniProgramLive pulumi.IntPtrInput
	// Play Type. This parameter is valid only if `DomainType` is 1. Available values: `1`: in Mainland China. `2`: global. `3`: outside Mainland China. Default value is 1.
	PlayType pulumi.IntPtrInput
	// Domain name attribution verification type. `dnsCheck`, `fileCheck`, `dbCheck`. The default is `dbCheck`.
	VerifyOwnerType pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//          DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//          DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Domain Name.
func (o DomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Domain type: `0`: push stream. `1`: playback.
func (o DomainOutput) DomainType() pulumi.IntOutput {
	return o.ApplyT(func(v *Domain) pulumi.IntOutput { return v.DomainType }).(pulumi.IntOutput)
}

// Switch. true: enable the specified domain, false: disable the specified domain.
func (o DomainOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.Enable }).(pulumi.BoolPtrOutput)
}

// Whether it is LCB: `0`: LVB. `1`: LCB. Default value is 0.
func (o DomainOutput) IsDelayLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.IntPtrOutput { return v.IsDelayLive }).(pulumi.IntPtrOutput)
}

// `0`: LVB. `1`: LVB on Mini Program. Note: this field may return null, indicating that no valid values can be obtained. Default value is 0.
func (o DomainOutput) IsMiniProgramLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.IntPtrOutput { return v.IsMiniProgramLive }).(pulumi.IntPtrOutput)
}

// Play Type. This parameter is valid only if `DomainType` is 1. Available values: `1`: in Mainland China. `2`: global. `3`: outside Mainland China. Default value is 1.
func (o DomainOutput) PlayType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.IntPtrOutput { return v.PlayType }).(pulumi.IntPtrOutput)
}

// Domain name attribution verification type. `dnsCheck`, `fileCheck`, `dbCheck`. The default is `dbCheck`.
func (o DomainOutput) VerifyOwnerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.VerifyOwnerType }).(pulumi.StringPtrOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].([]*Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].(map[string]*Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
