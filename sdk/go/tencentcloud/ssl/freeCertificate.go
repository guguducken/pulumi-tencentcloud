// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a Free Certificate.
//
// > **NOTE:** Once certificat created, it cannot be removed within 1 hours.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ssl.NewFreeCertificate(ctx, "foo", &Ssl.FreeCertificateArgs{
// 			Alias:           pulumi.String("my_free_cert"),
// 			ContactEmail:    pulumi.String("foo@example.com"),
// 			ContactPhone:    pulumi.String("12345678901"),
// 			CsrEncryptAlgo:  pulumi.String("RSA"),
// 			CsrKeyParameter: pulumi.String("2048"),
// 			CsrKeyPassword:  pulumi.String("xxxxxxxx"),
// 			Domain:          pulumi.String("example.com"),
// 			DvAuthMethod:    pulumi.String("DNS_AUTO"),
// 			PackageType:     pulumi.String("2"),
// 			ValidityPeriod:  pulumi.String("12"),
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
// FreeCertificate instance can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ssl/freeCertificate:FreeCertificate test free_certificate-id
// ```
type FreeCertificate struct {
	pulumi.CustomResourceState

	// Specify alias for remark.
	Alias pulumi.StringPtrOutput `pulumi:"alias"`
	// Certificate begin time.
	CertBeginTime pulumi.StringOutput `pulumi:"certBeginTime"`
	// Certificate end time.
	CertEndTime pulumi.StringOutput `pulumi:"certEndTime"`
	// Certificate private key.
	CertificatePrivateKey pulumi.StringOutput `pulumi:"certificatePrivateKey"`
	// Certificate public key.
	CertificatePublicKey pulumi.StringOutput `pulumi:"certificatePublicKey"`
	// Email address.
	ContactEmail pulumi.StringPtrOutput `pulumi:"contactEmail"`
	// Phone number.
	ContactPhone pulumi.StringPtrOutput `pulumi:"contactPhone"`
	// Specify CSR encrypt algorithm, only support `RSA` for now.
	CsrEncryptAlgo pulumi.StringPtrOutput `pulumi:"csrEncryptAlgo"`
	// Specify CSR key parameter, only support `"2048"` for now.
	CsrKeyParameter pulumi.StringPtrOutput `pulumi:"csrKeyParameter"`
	// Specify CSR key password.
	CsrKeyPassword pulumi.StringPtrOutput `pulumi:"csrKeyPassword"`
	// Indicates whether the certificate deployable.
	Deployable pulumi.BoolOutput `pulumi:"deployable"`
	// Specify domain name.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
	DvAuthMethod pulumi.StringOutput `pulumi:"dvAuthMethod"`
	// DV certification information.
	DvAuths FreeCertificateDvAuthArrayOutput `pulumi:"dvAuths"`
	// Certificate insert time.
	InsertTime pulumi.StringOutput `pulumi:"insertTime"`
	// Specify old certificate ID, used for re-apply.
	OldCertificateId pulumi.StringPtrOutput `pulumi:"oldCertificateId"`
	// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
	PackageType pulumi.StringPtrOutput `pulumi:"packageType"`
	// Product zh name.
	ProductZhName pulumi.StringOutput `pulumi:"productZhName"`
	// ID of projects which this certification belong to.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Indicates whether the certificate renewable.
	Renewable pulumi.BoolOutput `pulumi:"renewable"`
	// Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
	Status pulumi.IntOutput `pulumi:"status"`
	// Certificate status message.
	StatusMsg pulumi.StringOutput `pulumi:"statusMsg"`
	// Certificate status name.
	StatusName pulumi.StringOutput `pulumi:"statusName"`
	// Specify validity period in month, only support `"12"` months for now.
	ValidityPeriod pulumi.StringPtrOutput `pulumi:"validityPeriod"`
	// Vulnerability status.
	VulnerabilityStatus pulumi.StringOutput `pulumi:"vulnerabilityStatus"`
}

// NewFreeCertificate registers a new resource with the given unique name, arguments, and options.
func NewFreeCertificate(ctx *pulumi.Context,
	name string, args *FreeCertificateArgs, opts ...pulumi.ResourceOption) (*FreeCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.DvAuthMethod == nil {
		return nil, errors.New("invalid value for required argument 'DvAuthMethod'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FreeCertificate
	err := ctx.RegisterResource("tencentcloud:Ssl/freeCertificate:FreeCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFreeCertificate gets an existing FreeCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFreeCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FreeCertificateState, opts ...pulumi.ResourceOption) (*FreeCertificate, error) {
	var resource FreeCertificate
	err := ctx.ReadResource("tencentcloud:Ssl/freeCertificate:FreeCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FreeCertificate resources.
type freeCertificateState struct {
	// Specify alias for remark.
	Alias *string `pulumi:"alias"`
	// Certificate begin time.
	CertBeginTime *string `pulumi:"certBeginTime"`
	// Certificate end time.
	CertEndTime *string `pulumi:"certEndTime"`
	// Certificate private key.
	CertificatePrivateKey *string `pulumi:"certificatePrivateKey"`
	// Certificate public key.
	CertificatePublicKey *string `pulumi:"certificatePublicKey"`
	// Email address.
	ContactEmail *string `pulumi:"contactEmail"`
	// Phone number.
	ContactPhone *string `pulumi:"contactPhone"`
	// Specify CSR encrypt algorithm, only support `RSA` for now.
	CsrEncryptAlgo *string `pulumi:"csrEncryptAlgo"`
	// Specify CSR key parameter, only support `"2048"` for now.
	CsrKeyParameter *string `pulumi:"csrKeyParameter"`
	// Specify CSR key password.
	CsrKeyPassword *string `pulumi:"csrKeyPassword"`
	// Indicates whether the certificate deployable.
	Deployable *bool `pulumi:"deployable"`
	// Specify domain name.
	Domain *string `pulumi:"domain"`
	// Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
	DvAuthMethod *string `pulumi:"dvAuthMethod"`
	// DV certification information.
	DvAuths []FreeCertificateDvAuth `pulumi:"dvAuths"`
	// Certificate insert time.
	InsertTime *string `pulumi:"insertTime"`
	// Specify old certificate ID, used for re-apply.
	OldCertificateId *string `pulumi:"oldCertificateId"`
	// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
	PackageType *string `pulumi:"packageType"`
	// Product zh name.
	ProductZhName *string `pulumi:"productZhName"`
	// ID of projects which this certification belong to.
	ProjectId *int `pulumi:"projectId"`
	// Indicates whether the certificate renewable.
	Renewable *bool `pulumi:"renewable"`
	// Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
	Status *int `pulumi:"status"`
	// Certificate status message.
	StatusMsg *string `pulumi:"statusMsg"`
	// Certificate status name.
	StatusName *string `pulumi:"statusName"`
	// Specify validity period in month, only support `"12"` months for now.
	ValidityPeriod *string `pulumi:"validityPeriod"`
	// Vulnerability status.
	VulnerabilityStatus *string `pulumi:"vulnerabilityStatus"`
}

type FreeCertificateState struct {
	// Specify alias for remark.
	Alias pulumi.StringPtrInput
	// Certificate begin time.
	CertBeginTime pulumi.StringPtrInput
	// Certificate end time.
	CertEndTime pulumi.StringPtrInput
	// Certificate private key.
	CertificatePrivateKey pulumi.StringPtrInput
	// Certificate public key.
	CertificatePublicKey pulumi.StringPtrInput
	// Email address.
	ContactEmail pulumi.StringPtrInput
	// Phone number.
	ContactPhone pulumi.StringPtrInput
	// Specify CSR encrypt algorithm, only support `RSA` for now.
	CsrEncryptAlgo pulumi.StringPtrInput
	// Specify CSR key parameter, only support `"2048"` for now.
	CsrKeyParameter pulumi.StringPtrInput
	// Specify CSR key password.
	CsrKeyPassword pulumi.StringPtrInput
	// Indicates whether the certificate deployable.
	Deployable pulumi.BoolPtrInput
	// Specify domain name.
	Domain pulumi.StringPtrInput
	// Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
	DvAuthMethod pulumi.StringPtrInput
	// DV certification information.
	DvAuths FreeCertificateDvAuthArrayInput
	// Certificate insert time.
	InsertTime pulumi.StringPtrInput
	// Specify old certificate ID, used for re-apply.
	OldCertificateId pulumi.StringPtrInput
	// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
	PackageType pulumi.StringPtrInput
	// Product zh name.
	ProductZhName pulumi.StringPtrInput
	// ID of projects which this certification belong to.
	ProjectId pulumi.IntPtrInput
	// Indicates whether the certificate renewable.
	Renewable pulumi.BoolPtrInput
	// Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
	Status pulumi.IntPtrInput
	// Certificate status message.
	StatusMsg pulumi.StringPtrInput
	// Certificate status name.
	StatusName pulumi.StringPtrInput
	// Specify validity period in month, only support `"12"` months for now.
	ValidityPeriod pulumi.StringPtrInput
	// Vulnerability status.
	VulnerabilityStatus pulumi.StringPtrInput
}

func (FreeCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*freeCertificateState)(nil)).Elem()
}

type freeCertificateArgs struct {
	// Specify alias for remark.
	Alias *string `pulumi:"alias"`
	// Email address.
	ContactEmail *string `pulumi:"contactEmail"`
	// Phone number.
	ContactPhone *string `pulumi:"contactPhone"`
	// Specify CSR encrypt algorithm, only support `RSA` for now.
	CsrEncryptAlgo *string `pulumi:"csrEncryptAlgo"`
	// Specify CSR key parameter, only support `"2048"` for now.
	CsrKeyParameter *string `pulumi:"csrKeyParameter"`
	// Specify CSR key password.
	CsrKeyPassword *string `pulumi:"csrKeyPassword"`
	// Specify domain name.
	Domain string `pulumi:"domain"`
	// Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
	DvAuthMethod string `pulumi:"dvAuthMethod"`
	// Specify old certificate ID, used for re-apply.
	OldCertificateId *string `pulumi:"oldCertificateId"`
	// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
	PackageType *string `pulumi:"packageType"`
	// ID of projects which this certification belong to.
	ProjectId *int `pulumi:"projectId"`
	// Specify validity period in month, only support `"12"` months for now.
	ValidityPeriod *string `pulumi:"validityPeriod"`
}

// The set of arguments for constructing a FreeCertificate resource.
type FreeCertificateArgs struct {
	// Specify alias for remark.
	Alias pulumi.StringPtrInput
	// Email address.
	ContactEmail pulumi.StringPtrInput
	// Phone number.
	ContactPhone pulumi.StringPtrInput
	// Specify CSR encrypt algorithm, only support `RSA` for now.
	CsrEncryptAlgo pulumi.StringPtrInput
	// Specify CSR key parameter, only support `"2048"` for now.
	CsrKeyParameter pulumi.StringPtrInput
	// Specify CSR key password.
	CsrKeyPassword pulumi.StringPtrInput
	// Specify domain name.
	Domain pulumi.StringInput
	// Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
	DvAuthMethod pulumi.StringInput
	// Specify old certificate ID, used for re-apply.
	OldCertificateId pulumi.StringPtrInput
	// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
	PackageType pulumi.StringPtrInput
	// ID of projects which this certification belong to.
	ProjectId pulumi.IntPtrInput
	// Specify validity period in month, only support `"12"` months for now.
	ValidityPeriod pulumi.StringPtrInput
}

func (FreeCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*freeCertificateArgs)(nil)).Elem()
}

type FreeCertificateInput interface {
	pulumi.Input

	ToFreeCertificateOutput() FreeCertificateOutput
	ToFreeCertificateOutputWithContext(ctx context.Context) FreeCertificateOutput
}

func (*FreeCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**FreeCertificate)(nil)).Elem()
}

func (i *FreeCertificate) ToFreeCertificateOutput() FreeCertificateOutput {
	return i.ToFreeCertificateOutputWithContext(context.Background())
}

func (i *FreeCertificate) ToFreeCertificateOutputWithContext(ctx context.Context) FreeCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FreeCertificateOutput)
}

// FreeCertificateArrayInput is an input type that accepts FreeCertificateArray and FreeCertificateArrayOutput values.
// You can construct a concrete instance of `FreeCertificateArrayInput` via:
//
//          FreeCertificateArray{ FreeCertificateArgs{...} }
type FreeCertificateArrayInput interface {
	pulumi.Input

	ToFreeCertificateArrayOutput() FreeCertificateArrayOutput
	ToFreeCertificateArrayOutputWithContext(context.Context) FreeCertificateArrayOutput
}

type FreeCertificateArray []FreeCertificateInput

func (FreeCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FreeCertificate)(nil)).Elem()
}

func (i FreeCertificateArray) ToFreeCertificateArrayOutput() FreeCertificateArrayOutput {
	return i.ToFreeCertificateArrayOutputWithContext(context.Background())
}

func (i FreeCertificateArray) ToFreeCertificateArrayOutputWithContext(ctx context.Context) FreeCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FreeCertificateArrayOutput)
}

// FreeCertificateMapInput is an input type that accepts FreeCertificateMap and FreeCertificateMapOutput values.
// You can construct a concrete instance of `FreeCertificateMapInput` via:
//
//          FreeCertificateMap{ "key": FreeCertificateArgs{...} }
type FreeCertificateMapInput interface {
	pulumi.Input

	ToFreeCertificateMapOutput() FreeCertificateMapOutput
	ToFreeCertificateMapOutputWithContext(context.Context) FreeCertificateMapOutput
}

type FreeCertificateMap map[string]FreeCertificateInput

func (FreeCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FreeCertificate)(nil)).Elem()
}

func (i FreeCertificateMap) ToFreeCertificateMapOutput() FreeCertificateMapOutput {
	return i.ToFreeCertificateMapOutputWithContext(context.Background())
}

func (i FreeCertificateMap) ToFreeCertificateMapOutputWithContext(ctx context.Context) FreeCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FreeCertificateMapOutput)
}

type FreeCertificateOutput struct{ *pulumi.OutputState }

func (FreeCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FreeCertificate)(nil)).Elem()
}

func (o FreeCertificateOutput) ToFreeCertificateOutput() FreeCertificateOutput {
	return o
}

func (o FreeCertificateOutput) ToFreeCertificateOutputWithContext(ctx context.Context) FreeCertificateOutput {
	return o
}

// Specify alias for remark.
func (o FreeCertificateOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.Alias }).(pulumi.StringPtrOutput)
}

// Certificate begin time.
func (o FreeCertificateOutput) CertBeginTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.CertBeginTime }).(pulumi.StringOutput)
}

// Certificate end time.
func (o FreeCertificateOutput) CertEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.CertEndTime }).(pulumi.StringOutput)
}

// Certificate private key.
func (o FreeCertificateOutput) CertificatePrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.CertificatePrivateKey }).(pulumi.StringOutput)
}

// Certificate public key.
func (o FreeCertificateOutput) CertificatePublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.CertificatePublicKey }).(pulumi.StringOutput)
}

// Email address.
func (o FreeCertificateOutput) ContactEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.ContactEmail }).(pulumi.StringPtrOutput)
}

// Phone number.
func (o FreeCertificateOutput) ContactPhone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.ContactPhone }).(pulumi.StringPtrOutput)
}

// Specify CSR encrypt algorithm, only support `RSA` for now.
func (o FreeCertificateOutput) CsrEncryptAlgo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.CsrEncryptAlgo }).(pulumi.StringPtrOutput)
}

// Specify CSR key parameter, only support `"2048"` for now.
func (o FreeCertificateOutput) CsrKeyParameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.CsrKeyParameter }).(pulumi.StringPtrOutput)
}

// Specify CSR key password.
func (o FreeCertificateOutput) CsrKeyPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.CsrKeyPassword }).(pulumi.StringPtrOutput)
}

// Indicates whether the certificate deployable.
func (o FreeCertificateOutput) Deployable() pulumi.BoolOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.BoolOutput { return v.Deployable }).(pulumi.BoolOutput)
}

// Specify domain name.
func (o FreeCertificateOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Specify DV authorize method. Available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
func (o FreeCertificateOutput) DvAuthMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.DvAuthMethod }).(pulumi.StringOutput)
}

// DV certification information.
func (o FreeCertificateOutput) DvAuths() FreeCertificateDvAuthArrayOutput {
	return o.ApplyT(func(v *FreeCertificate) FreeCertificateDvAuthArrayOutput { return v.DvAuths }).(FreeCertificateDvAuthArrayOutput)
}

// Certificate insert time.
func (o FreeCertificateOutput) InsertTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.InsertTime }).(pulumi.StringOutput)
}

// Specify old certificate ID, used for re-apply.
func (o FreeCertificateOutput) OldCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.OldCertificateId }).(pulumi.StringPtrOutput)
}

// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
func (o FreeCertificateOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.PackageType }).(pulumi.StringPtrOutput)
}

// Product zh name.
func (o FreeCertificateOutput) ProductZhName() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.ProductZhName }).(pulumi.StringOutput)
}

// ID of projects which this certification belong to.
func (o FreeCertificateOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Indicates whether the certificate renewable.
func (o FreeCertificateOutput) Renewable() pulumi.BoolOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.BoolOutput { return v.Renewable }).(pulumi.BoolOutput)
}

// Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
func (o FreeCertificateOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Certificate status message.
func (o FreeCertificateOutput) StatusMsg() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.StatusMsg }).(pulumi.StringOutput)
}

// Certificate status name.
func (o FreeCertificateOutput) StatusName() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.StatusName }).(pulumi.StringOutput)
}

// Specify validity period in month, only support `"12"` months for now.
func (o FreeCertificateOutput) ValidityPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringPtrOutput { return v.ValidityPeriod }).(pulumi.StringPtrOutput)
}

// Vulnerability status.
func (o FreeCertificateOutput) VulnerabilityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FreeCertificate) pulumi.StringOutput { return v.VulnerabilityStatus }).(pulumi.StringOutput)
}

type FreeCertificateArrayOutput struct{ *pulumi.OutputState }

func (FreeCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FreeCertificate)(nil)).Elem()
}

func (o FreeCertificateArrayOutput) ToFreeCertificateArrayOutput() FreeCertificateArrayOutput {
	return o
}

func (o FreeCertificateArrayOutput) ToFreeCertificateArrayOutputWithContext(ctx context.Context) FreeCertificateArrayOutput {
	return o
}

func (o FreeCertificateArrayOutput) Index(i pulumi.IntInput) FreeCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FreeCertificate {
		return vs[0].([]*FreeCertificate)[vs[1].(int)]
	}).(FreeCertificateOutput)
}

type FreeCertificateMapOutput struct{ *pulumi.OutputState }

func (FreeCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FreeCertificate)(nil)).Elem()
}

func (o FreeCertificateMapOutput) ToFreeCertificateMapOutput() FreeCertificateMapOutput {
	return o
}

func (o FreeCertificateMapOutput) ToFreeCertificateMapOutputWithContext(ctx context.Context) FreeCertificateMapOutput {
	return o
}

func (o FreeCertificateMapOutput) MapIndex(k pulumi.StringInput) FreeCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FreeCertificate {
		return vs[0].(map[string]*FreeCertificate)[vs[1].(string)]
	}).(FreeCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FreeCertificateInput)(nil)).Elem(), &FreeCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*FreeCertificateArrayInput)(nil)).Elem(), FreeCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FreeCertificateMapInput)(nil)).Elem(), FreeCertificateMap{})
	pulumi.RegisterOutputType(FreeCertificateOutput{})
	pulumi.RegisterOutputType(FreeCertificateArrayOutput{})
	pulumi.RegisterOutputType(FreeCertificateMapOutput{})
}
