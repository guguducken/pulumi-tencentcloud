// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver businessIntelligenceFile
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
// 		businessIntelligenceInstance, err := Sqlserver.NewBusinessIntelligenceInstance(ctx, "businessIntelligenceInstance", &Sqlserver.BusinessIntelligenceInstanceArgs{
// 			Zone:               pulumi.String("ap-guangzhou-6"),
// 			Memory:             pulumi.Int(4),
// 			Storage:            pulumi.Int(20),
// 			Cpu:                pulumi.Int(2),
// 			MachineType:        pulumi.String("CLOUD_PREMIUM"),
// 			ProjectId:          pulumi.Int(0),
// 			SubnetId:           pulumi.String("subnet-dwj7ipnc"),
// 			VpcId:              pulumi.String("vpc-4owdpnwr"),
// 			DbVersion:          pulumi.String("201603"),
// 			SecurityGroupLists: pulumi.StringArray{},
// 			Weeklies: pulumi.IntArray{
// 				pulumi.Int(1),
// 				pulumi.Int(2),
// 				pulumi.Int(3),
// 				pulumi.Int(4),
// 				pulumi.Int(5),
// 				pulumi.Int(6),
// 				pulumi.Int(7),
// 			},
// 			StartTime:    pulumi.String("00:00"),
// 			Span:         pulumi.Int(6),
// 			InstanceName: pulumi.String("create_db_name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Sqlserver.NewBusinessIntelligenceFile(ctx, "businessIntelligenceFile", &Sqlserver.BusinessIntelligenceFileArgs{
// 			InstanceId: businessIntelligenceInstance.ID(),
// 			FileUrl:    pulumi.String("https://keep-sqlserver-1308919341.cos.ap-guangzhou.myqcloud.com/test.xlsx"),
// 			FileType:   pulumi.String("FLAT"),
// 			Remark:     pulumi.String("test case."),
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
// sqlserver business_intelligence_file can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/businessIntelligenceFile:BusinessIntelligenceFile business_intelligence_file business_intelligence_file_id
// ```
type BusinessIntelligenceFile struct {
	pulumi.CustomResourceState

	// File Type FLAT - Flat File as Data Source, SSIS - ssis project package.
	FileType pulumi.StringOutput `pulumi:"fileType"`
	// Cos Url.
	FileUrl pulumi.StringOutput `pulumi:"fileUrl"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// remark.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
}

// NewBusinessIntelligenceFile registers a new resource with the given unique name, arguments, and options.
func NewBusinessIntelligenceFile(ctx *pulumi.Context,
	name string, args *BusinessIntelligenceFileArgs, opts ...pulumi.ResourceOption) (*BusinessIntelligenceFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileType == nil {
		return nil, errors.New("invalid value for required argument 'FileType'")
	}
	if args.FileUrl == nil {
		return nil, errors.New("invalid value for required argument 'FileUrl'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BusinessIntelligenceFile
	err := ctx.RegisterResource("tencentcloud:Sqlserver/businessIntelligenceFile:BusinessIntelligenceFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBusinessIntelligenceFile gets an existing BusinessIntelligenceFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBusinessIntelligenceFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BusinessIntelligenceFileState, opts ...pulumi.ResourceOption) (*BusinessIntelligenceFile, error) {
	var resource BusinessIntelligenceFile
	err := ctx.ReadResource("tencentcloud:Sqlserver/businessIntelligenceFile:BusinessIntelligenceFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BusinessIntelligenceFile resources.
type businessIntelligenceFileState struct {
	// File Type FLAT - Flat File as Data Source, SSIS - ssis project package.
	FileType *string `pulumi:"fileType"`
	// Cos Url.
	FileUrl *string `pulumi:"fileUrl"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// remark.
	Remark *string `pulumi:"remark"`
}

type BusinessIntelligenceFileState struct {
	// File Type FLAT - Flat File as Data Source, SSIS - ssis project package.
	FileType pulumi.StringPtrInput
	// Cos Url.
	FileUrl pulumi.StringPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
	// remark.
	Remark pulumi.StringPtrInput
}

func (BusinessIntelligenceFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*businessIntelligenceFileState)(nil)).Elem()
}

type businessIntelligenceFileArgs struct {
	// File Type FLAT - Flat File as Data Source, SSIS - ssis project package.
	FileType string `pulumi:"fileType"`
	// Cos Url.
	FileUrl string `pulumi:"fileUrl"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// remark.
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a BusinessIntelligenceFile resource.
type BusinessIntelligenceFileArgs struct {
	// File Type FLAT - Flat File as Data Source, SSIS - ssis project package.
	FileType pulumi.StringInput
	// Cos Url.
	FileUrl pulumi.StringInput
	// instance id.
	InstanceId pulumi.StringInput
	// remark.
	Remark pulumi.StringPtrInput
}

func (BusinessIntelligenceFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*businessIntelligenceFileArgs)(nil)).Elem()
}

type BusinessIntelligenceFileInput interface {
	pulumi.Input

	ToBusinessIntelligenceFileOutput() BusinessIntelligenceFileOutput
	ToBusinessIntelligenceFileOutputWithContext(ctx context.Context) BusinessIntelligenceFileOutput
}

func (*BusinessIntelligenceFile) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIntelligenceFile)(nil)).Elem()
}

func (i *BusinessIntelligenceFile) ToBusinessIntelligenceFileOutput() BusinessIntelligenceFileOutput {
	return i.ToBusinessIntelligenceFileOutputWithContext(context.Background())
}

func (i *BusinessIntelligenceFile) ToBusinessIntelligenceFileOutputWithContext(ctx context.Context) BusinessIntelligenceFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIntelligenceFileOutput)
}

// BusinessIntelligenceFileArrayInput is an input type that accepts BusinessIntelligenceFileArray and BusinessIntelligenceFileArrayOutput values.
// You can construct a concrete instance of `BusinessIntelligenceFileArrayInput` via:
//
//          BusinessIntelligenceFileArray{ BusinessIntelligenceFileArgs{...} }
type BusinessIntelligenceFileArrayInput interface {
	pulumi.Input

	ToBusinessIntelligenceFileArrayOutput() BusinessIntelligenceFileArrayOutput
	ToBusinessIntelligenceFileArrayOutputWithContext(context.Context) BusinessIntelligenceFileArrayOutput
}

type BusinessIntelligenceFileArray []BusinessIntelligenceFileInput

func (BusinessIntelligenceFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BusinessIntelligenceFile)(nil)).Elem()
}

func (i BusinessIntelligenceFileArray) ToBusinessIntelligenceFileArrayOutput() BusinessIntelligenceFileArrayOutput {
	return i.ToBusinessIntelligenceFileArrayOutputWithContext(context.Background())
}

func (i BusinessIntelligenceFileArray) ToBusinessIntelligenceFileArrayOutputWithContext(ctx context.Context) BusinessIntelligenceFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIntelligenceFileArrayOutput)
}

// BusinessIntelligenceFileMapInput is an input type that accepts BusinessIntelligenceFileMap and BusinessIntelligenceFileMapOutput values.
// You can construct a concrete instance of `BusinessIntelligenceFileMapInput` via:
//
//          BusinessIntelligenceFileMap{ "key": BusinessIntelligenceFileArgs{...} }
type BusinessIntelligenceFileMapInput interface {
	pulumi.Input

	ToBusinessIntelligenceFileMapOutput() BusinessIntelligenceFileMapOutput
	ToBusinessIntelligenceFileMapOutputWithContext(context.Context) BusinessIntelligenceFileMapOutput
}

type BusinessIntelligenceFileMap map[string]BusinessIntelligenceFileInput

func (BusinessIntelligenceFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BusinessIntelligenceFile)(nil)).Elem()
}

func (i BusinessIntelligenceFileMap) ToBusinessIntelligenceFileMapOutput() BusinessIntelligenceFileMapOutput {
	return i.ToBusinessIntelligenceFileMapOutputWithContext(context.Background())
}

func (i BusinessIntelligenceFileMap) ToBusinessIntelligenceFileMapOutputWithContext(ctx context.Context) BusinessIntelligenceFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIntelligenceFileMapOutput)
}

type BusinessIntelligenceFileOutput struct{ *pulumi.OutputState }

func (BusinessIntelligenceFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIntelligenceFile)(nil)).Elem()
}

func (o BusinessIntelligenceFileOutput) ToBusinessIntelligenceFileOutput() BusinessIntelligenceFileOutput {
	return o
}

func (o BusinessIntelligenceFileOutput) ToBusinessIntelligenceFileOutputWithContext(ctx context.Context) BusinessIntelligenceFileOutput {
	return o
}

// File Type FLAT - Flat File as Data Source, SSIS - ssis project package.
func (o BusinessIntelligenceFileOutput) FileType() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessIntelligenceFile) pulumi.StringOutput { return v.FileType }).(pulumi.StringOutput)
}

// Cos Url.
func (o BusinessIntelligenceFileOutput) FileUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessIntelligenceFile) pulumi.StringOutput { return v.FileUrl }).(pulumi.StringOutput)
}

// instance id.
func (o BusinessIntelligenceFileOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessIntelligenceFile) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// remark.
func (o BusinessIntelligenceFileOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIntelligenceFile) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

type BusinessIntelligenceFileArrayOutput struct{ *pulumi.OutputState }

func (BusinessIntelligenceFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BusinessIntelligenceFile)(nil)).Elem()
}

func (o BusinessIntelligenceFileArrayOutput) ToBusinessIntelligenceFileArrayOutput() BusinessIntelligenceFileArrayOutput {
	return o
}

func (o BusinessIntelligenceFileArrayOutput) ToBusinessIntelligenceFileArrayOutputWithContext(ctx context.Context) BusinessIntelligenceFileArrayOutput {
	return o
}

func (o BusinessIntelligenceFileArrayOutput) Index(i pulumi.IntInput) BusinessIntelligenceFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BusinessIntelligenceFile {
		return vs[0].([]*BusinessIntelligenceFile)[vs[1].(int)]
	}).(BusinessIntelligenceFileOutput)
}

type BusinessIntelligenceFileMapOutput struct{ *pulumi.OutputState }

func (BusinessIntelligenceFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BusinessIntelligenceFile)(nil)).Elem()
}

func (o BusinessIntelligenceFileMapOutput) ToBusinessIntelligenceFileMapOutput() BusinessIntelligenceFileMapOutput {
	return o
}

func (o BusinessIntelligenceFileMapOutput) ToBusinessIntelligenceFileMapOutputWithContext(ctx context.Context) BusinessIntelligenceFileMapOutput {
	return o
}

func (o BusinessIntelligenceFileMapOutput) MapIndex(k pulumi.StringInput) BusinessIntelligenceFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BusinessIntelligenceFile {
		return vs[0].(map[string]*BusinessIntelligenceFile)[vs[1].(string)]
	}).(BusinessIntelligenceFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BusinessIntelligenceFileInput)(nil)).Elem(), &BusinessIntelligenceFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*BusinessIntelligenceFileArrayInput)(nil)).Elem(), BusinessIntelligenceFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BusinessIntelligenceFileMapInput)(nil)).Elem(), BusinessIntelligenceFileMap{})
	pulumi.RegisterOutputType(BusinessIntelligenceFileOutput{})
	pulumi.RegisterOutputType(BusinessIntelligenceFileArrayOutput{})
	pulumi.RegisterOutputType(BusinessIntelligenceFileMapOutput{})
}
