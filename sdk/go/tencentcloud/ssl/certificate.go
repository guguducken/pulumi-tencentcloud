// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a SSL certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ssl.NewCertificate(ctx, "foo", &Ssl.CertificateArgs{
// 			Cert:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "-----BEGIN CERTIFICATE-----\n", "MIIERzCCAq+gAwIBAgIBAjANBgkqhkiG9w0BAQsFADAoMQ0wCwYDVQQDEwR0ZXN0\n", "MRcwFQYDVQQKEw50ZXJyYWZvcm0gdGVzdDAeFw0xOTA4MTMwMzE5MzlaFw0yOTA4\n", "MTAwMzE5MzlaMC4xEzARBgNVBAMTCnNlcnZlciBzc2wxFzAVBgNVBAoTDnRlcnJh\n", "Zm9ybS10ZXN0MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA1Ryp+DKK\n", "SNFKZsPtwfR+jzOnQ8YFieIKYgakV688d8YgpolenbmeEPrzT87tunFD7G9f6ALG\n", "ND8rj7npj0AowxhOL/h/v1D9u0UsIaj5i2GWJrqNAhGLaxWiEB/hy5WOiwxDrGei\n", "gQqJkFM52Ep7G1Yx7PHJmKFGwN9FhIsFi1cNZfVRopZuCe/RMPNusNVZaIi+qcEf\n", "fsE1cmfmuSlG3Ap0RKOIyR0ajDEzqZn9/0R7VwWCF97qy8TNYk94K/1tq3zyhVzR\n", "Z83xOSfrTqEfb3so3AU2jyKgYdwr/FZS72VCHS8IslgnqJW4izIXZqgIKmHaRZtM\n", "N4jUloi6l/6lktt6Lsgh9xECecxziSJtPMaog88aC8HnMqJJ3kScGCL36GYG+Kaw\n", "5PnDlWXBaeiDe8z/eWK9+Rr2M+rhTNxosAVGfDJyxAXyiX49LQ0v7f9qzwc/0JiD\n", "bvsUv1cm6OgpoEMP9SXqqBdwGqeKbD2/2jlP48xlYP6l1SoJG3GgZ8dbAgMBAAGj\n", "djB0MAwGA1UdEwEB/wQCMAAwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0PAQH/\n", "BAUDAweAADAdBgNVHQ4EFgQULwWKBQNLL9s3cb3tTnyPVg+mpCMwHwYDVR0jBBgw\n", "FoAUKwfrmq791mY831S6UHARHtgYnlgwDQYJKoZIhvcNAQELBQADggGBAMo5RglS\n", "AHdPgaicWJvmvjjexjF/42b7Rz4pPfMjYw6uYO8He/f4UZWv5CZLrbEe7MywaK3y\n", "0OsfH8AhyN29pv2x8g9wbmq7omZIOZ0oCAGduEXs/A/qY/hFaCohdkz/IN8qi6JW\n", "VXreGli3SrpcHFchSwHTyJEXgkutcGAsOvdsOuVSmplOyrkLHc8uUe8SG4j8kGyg\n", "EzaszFjHkR7g1dVyDVUedc588mjkQxYeAamJgfkgIhljWKMa2XzkVMcVfQHfNpM1\n", "n+bu8SmqRt9Wma2bMijKRG/Blm756LoI+skY+WRZmlDnq8zj95TT0vceGP0FUWh5\n", "hKyiocABmpQs9OK9HMi8vgSWISP+fYgkm/bKtKup2NbZBoO5/VL2vCEPInYzUhBO\n", "jCbLMjNjtM5KriCaR7wDARgHiG0gBEPOEW1PIjZ9UOH+LtIxbNZ4eEIIINLHnBHf\n", "L+doVeZtS/gJc4G4Adr5HYuaS9ZxJ0W2uy0eQlOHzjyxR6Mf/rpnilJlcQ==\n", "-----END CERTIFICATE-----\n")),
// 			ProjectId: pulumi.Int(0),
// 			Type:      pulumi.String("CA"),
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
// ssl certificate can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ssl/certificate:Certificate tencentcloud_ssl_certificate.cert GjTNRoK7
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// Beginning time of the SSL certificate.
	BeginTime pulumi.StringOutput `pulumi:"beginTime"`
	// Content of the SSL certificate. Not allowed newline at the start and end.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Creation time of the SSL certificate.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Primary domain of the SSL certificate.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Ending time of the SSL certificate.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// Name of the SSL certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// Certificate authority.
	ProductZhName pulumi.StringOutput `pulumi:"productZhName"`
	// Project ID of the SSL certificate. Default is `0`.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Status of the SSL certificate.
	Status pulumi.IntOutput `pulumi:"status"`
	// ALL domains included in the SSL certificate. Including the primary domain name.
	SubjectNames pulumi.StringArrayOutput `pulumi:"subjectNames"`
	// Tags of the SSL certificate.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Type of the SSL certificate. Valid values: `CA` and `SVR`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cert == nil {
		return nil, errors.New("invalid value for required argument 'Cert'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("tencentcloud:Ssl/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("tencentcloud:Ssl/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Beginning time of the SSL certificate.
	BeginTime *string `pulumi:"beginTime"`
	// Content of the SSL certificate. Not allowed newline at the start and end.
	Cert *string `pulumi:"cert"`
	// Creation time of the SSL certificate.
	CreateTime *string `pulumi:"createTime"`
	// Primary domain of the SSL certificate.
	Domain *string `pulumi:"domain"`
	// Ending time of the SSL certificate.
	EndTime *string `pulumi:"endTime"`
	// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
	Key *string `pulumi:"key"`
	// Name of the SSL certificate.
	Name *string `pulumi:"name"`
	// Certificate authority.
	ProductZhName *string `pulumi:"productZhName"`
	// Project ID of the SSL certificate. Default is `0`.
	ProjectId *int `pulumi:"projectId"`
	// Status of the SSL certificate.
	Status *int `pulumi:"status"`
	// ALL domains included in the SSL certificate. Including the primary domain name.
	SubjectNames []string `pulumi:"subjectNames"`
	// Tags of the SSL certificate.
	Tags map[string]interface{} `pulumi:"tags"`
	// Type of the SSL certificate. Valid values: `CA` and `SVR`.
	Type *string `pulumi:"type"`
}

type CertificateState struct {
	// Beginning time of the SSL certificate.
	BeginTime pulumi.StringPtrInput
	// Content of the SSL certificate. Not allowed newline at the start and end.
	Cert pulumi.StringPtrInput
	// Creation time of the SSL certificate.
	CreateTime pulumi.StringPtrInput
	// Primary domain of the SSL certificate.
	Domain pulumi.StringPtrInput
	// Ending time of the SSL certificate.
	EndTime pulumi.StringPtrInput
	// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
	Key pulumi.StringPtrInput
	// Name of the SSL certificate.
	Name pulumi.StringPtrInput
	// Certificate authority.
	ProductZhName pulumi.StringPtrInput
	// Project ID of the SSL certificate. Default is `0`.
	ProjectId pulumi.IntPtrInput
	// Status of the SSL certificate.
	Status pulumi.IntPtrInput
	// ALL domains included in the SSL certificate. Including the primary domain name.
	SubjectNames pulumi.StringArrayInput
	// Tags of the SSL certificate.
	Tags pulumi.MapInput
	// Type of the SSL certificate. Valid values: `CA` and `SVR`.
	Type pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Content of the SSL certificate. Not allowed newline at the start and end.
	Cert string `pulumi:"cert"`
	// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
	Key *string `pulumi:"key"`
	// Name of the SSL certificate.
	Name *string `pulumi:"name"`
	// Project ID of the SSL certificate. Default is `0`.
	ProjectId *int `pulumi:"projectId"`
	// Tags of the SSL certificate.
	Tags map[string]interface{} `pulumi:"tags"`
	// Type of the SSL certificate. Valid values: `CA` and `SVR`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Content of the SSL certificate. Not allowed newline at the start and end.
	Cert pulumi.StringInput
	// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
	Key pulumi.StringPtrInput
	// Name of the SSL certificate.
	Name pulumi.StringPtrInput
	// Project ID of the SSL certificate. Default is `0`.
	ProjectId pulumi.IntPtrInput
	// Tags of the SSL certificate.
	Tags pulumi.MapInput
	// Type of the SSL certificate. Valid values: `CA` and `SVR`.
	Type pulumi.StringInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//          CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//          CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// Beginning time of the SSL certificate.
func (o CertificateOutput) BeginTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.BeginTime }).(pulumi.StringOutput)
}

// Content of the SSL certificate. Not allowed newline at the start and end.
func (o CertificateOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Cert }).(pulumi.StringOutput)
}

// Creation time of the SSL certificate.
func (o CertificateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Primary domain of the SSL certificate.
func (o CertificateOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Ending time of the SSL certificate.
func (o CertificateOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.
func (o CertificateOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// Name of the SSL certificate.
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Certificate authority.
func (o CertificateOutput) ProductZhName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ProductZhName }).(pulumi.StringOutput)
}

// Project ID of the SSL certificate. Default is `0`.
func (o CertificateOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Status of the SSL certificate.
func (o CertificateOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Certificate) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// ALL domains included in the SSL certificate. Including the primary domain name.
func (o CertificateOutput) SubjectNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.SubjectNames }).(pulumi.StringArrayOutput)
}

// Tags of the SSL certificate.
func (o CertificateOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Type of the SSL certificate. Valid values: `CA` and `SVR`.
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
