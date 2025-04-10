// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chdfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a chdfs fileSystem
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Chdfs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Chdfs.NewFileSystem(ctx, "fileSystem", &Chdfs.FileSystemArgs{
// 			CapacityQuota:  pulumi.Int(10995116277760),
// 			Description:    pulumi.String("file system for terraform test"),
// 			EnableRanger:   pulumi.Bool(true),
// 			FileSystemName: pulumi.String("terraform-test"),
// 			PosixAcl:       pulumi.Bool(false),
// 			RangerServiceAddresses: pulumi.StringArray{
// 				pulumi.String("127.0.0.1:80"),
// 				pulumi.String("127.0.0.1:8000"),
// 			},
// 			SuperUsers: pulumi.StringArray{
// 				pulumi.String("terraform"),
// 				pulumi.String("iac"),
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
// chdfs file_system can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Chdfs/fileSystem:FileSystem file_system file_system_id
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
	CapacityQuota pulumi.IntOutput `pulumi:"capacityQuota"`
	// desc of the file system.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// check the ranger address or not.
	EnableRanger pulumi.BoolPtrOutput `pulumi:"enableRanger"`
	// file system name.
	FileSystemName pulumi.StringOutput `pulumi:"fileSystemName"`
	// check POSIX ACL or not.
	PosixAcl pulumi.BoolOutput `pulumi:"posixAcl"`
	// ranger address list, default empty.
	RangerServiceAddresses pulumi.StringArrayOutput `pulumi:"rangerServiceAddresses"`
	// super users of the file system, default empty.
	SuperUsers pulumi.StringArrayOutput `pulumi:"superUsers"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityQuota == nil {
		return nil, errors.New("invalid value for required argument 'CapacityQuota'")
	}
	if args.FileSystemName == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemName'")
	}
	if args.PosixAcl == nil {
		return nil, errors.New("invalid value for required argument 'PosixAcl'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FileSystem
	err := ctx.RegisterResource("tencentcloud:Chdfs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("tencentcloud:Chdfs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
	CapacityQuota *int `pulumi:"capacityQuota"`
	// desc of the file system.
	Description *string `pulumi:"description"`
	// check the ranger address or not.
	EnableRanger *bool `pulumi:"enableRanger"`
	// file system name.
	FileSystemName *string `pulumi:"fileSystemName"`
	// check POSIX ACL or not.
	PosixAcl *bool `pulumi:"posixAcl"`
	// ranger address list, default empty.
	RangerServiceAddresses []string `pulumi:"rangerServiceAddresses"`
	// super users of the file system, default empty.
	SuperUsers []string `pulumi:"superUsers"`
}

type FileSystemState struct {
	// file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
	CapacityQuota pulumi.IntPtrInput
	// desc of the file system.
	Description pulumi.StringPtrInput
	// check the ranger address or not.
	EnableRanger pulumi.BoolPtrInput
	// file system name.
	FileSystemName pulumi.StringPtrInput
	// check POSIX ACL or not.
	PosixAcl pulumi.BoolPtrInput
	// ranger address list, default empty.
	RangerServiceAddresses pulumi.StringArrayInput
	// super users of the file system, default empty.
	SuperUsers pulumi.StringArrayInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
	CapacityQuota int `pulumi:"capacityQuota"`
	// desc of the file system.
	Description *string `pulumi:"description"`
	// check the ranger address or not.
	EnableRanger *bool `pulumi:"enableRanger"`
	// file system name.
	FileSystemName string `pulumi:"fileSystemName"`
	// check POSIX ACL or not.
	PosixAcl bool `pulumi:"posixAcl"`
	// ranger address list, default empty.
	RangerServiceAddresses []string `pulumi:"rangerServiceAddresses"`
	// super users of the file system, default empty.
	SuperUsers []string `pulumi:"superUsers"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
	CapacityQuota pulumi.IntInput
	// desc of the file system.
	Description pulumi.StringPtrInput
	// check the ranger address or not.
	EnableRanger pulumi.BoolPtrInput
	// file system name.
	FileSystemName pulumi.StringInput
	// check POSIX ACL or not.
	PosixAcl pulumi.BoolInput
	// ranger address list, default empty.
	RangerServiceAddresses pulumi.StringArrayInput
	// super users of the file system, default empty.
	SuperUsers pulumi.StringArrayInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

// FileSystemArrayInput is an input type that accepts FileSystemArray and FileSystemArrayOutput values.
// You can construct a concrete instance of `FileSystemArrayInput` via:
//
//          FileSystemArray{ FileSystemArgs{...} }
type FileSystemArrayInput interface {
	pulumi.Input

	ToFileSystemArrayOutput() FileSystemArrayOutput
	ToFileSystemArrayOutputWithContext(context.Context) FileSystemArrayOutput
}

type FileSystemArray []FileSystemInput

func (FileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (i FileSystemArray) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return i.ToFileSystemArrayOutputWithContext(context.Background())
}

func (i FileSystemArray) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemArrayOutput)
}

// FileSystemMapInput is an input type that accepts FileSystemMap and FileSystemMapOutput values.
// You can construct a concrete instance of `FileSystemMapInput` via:
//
//          FileSystemMap{ "key": FileSystemArgs{...} }
type FileSystemMapInput interface {
	pulumi.Input

	ToFileSystemMapOutput() FileSystemMapOutput
	ToFileSystemMapOutputWithContext(context.Context) FileSystemMapOutput
}

type FileSystemMap map[string]FileSystemInput

func (FileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (i FileSystemMap) ToFileSystemMapOutput() FileSystemMapOutput {
	return i.ToFileSystemMapOutputWithContext(context.Background())
}

func (i FileSystemMap) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemMapOutput)
}

type FileSystemOutput struct{ *pulumi.OutputState }

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

// file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
func (o FileSystemOutput) CapacityQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.IntOutput { return v.CapacityQuota }).(pulumi.IntOutput)
}

// desc of the file system.
func (o FileSystemOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// check the ranger address or not.
func (o FileSystemOutput) EnableRanger() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.BoolPtrOutput { return v.EnableRanger }).(pulumi.BoolPtrOutput)
}

// file system name.
func (o FileSystemOutput) FileSystemName() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.FileSystemName }).(pulumi.StringOutput)
}

// check POSIX ACL or not.
func (o FileSystemOutput) PosixAcl() pulumi.BoolOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.BoolOutput { return v.PosixAcl }).(pulumi.BoolOutput)
}

// ranger address list, default empty.
func (o FileSystemOutput) RangerServiceAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringArrayOutput { return v.RangerServiceAddresses }).(pulumi.StringArrayOutput)
}

// super users of the file system, default empty.
func (o FileSystemOutput) SuperUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringArrayOutput { return v.SuperUsers }).(pulumi.StringArrayOutput)
}

type FileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) Index(i pulumi.IntInput) FileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FileSystem {
		return vs[0].([]*FileSystem)[vs[1].(int)]
	}).(FileSystemOutput)
}

type FileSystemMapOutput struct{ *pulumi.OutputState }

func (FileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (o FileSystemMapOutput) ToFileSystemMapOutput() FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) MapIndex(k pulumi.StringInput) FileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FileSystem {
		return vs[0].(map[string]*FileSystem)[vs[1].(string)]
	}).(FileSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemInput)(nil)).Elem(), &FileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemArrayInput)(nil)).Elem(), FileSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemMapInput)(nil)).Elem(), FileSystemMap{})
	pulumi.RegisterOutputType(FileSystemOutput{})
	pulumi.RegisterOutputType(FileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileSystemMapOutput{})
}
