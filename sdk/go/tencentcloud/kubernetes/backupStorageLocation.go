// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create tke backup storage location.
//
// > **NOTE:** To create this resource, you need to create a cos bucket with prefix "tke-backup" in advance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Kubernetes.NewBackupStorageLocation(ctx, "exampleBackup", &Kubernetes.BackupStorageLocationArgs{
// 			Bucket:        pulumi.String("tke-backup-example-1"),
// 			StorageRegion: pulumi.String("ap-guangzhou"),
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
// tke backup storage location can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Kubernetes/backupStorageLocation:BackupStorageLocation test xxx
// ```
type BackupStorageLocation struct {
	pulumi.CustomResourceState

	// Name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Message of the backup storage location.
	Message pulumi.StringOutput `pulumi:"message"`
	// Name of the backup storage location.
	Name pulumi.StringOutput `pulumi:"name"`
	// Prefix of the bucket.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// State of the backup storage location.
	State pulumi.StringOutput `pulumi:"state"`
	// Region of the storage.
	StorageRegion pulumi.StringOutput `pulumi:"storageRegion"`
}

// NewBackupStorageLocation registers a new resource with the given unique name, arguments, and options.
func NewBackupStorageLocation(ctx *pulumi.Context,
	name string, args *BackupStorageLocationArgs, opts ...pulumi.ResourceOption) (*BackupStorageLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.StorageRegion == nil {
		return nil, errors.New("invalid value for required argument 'StorageRegion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BackupStorageLocation
	err := ctx.RegisterResource("tencentcloud:Kubernetes/backupStorageLocation:BackupStorageLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupStorageLocation gets an existing BackupStorageLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupStorageLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupStorageLocationState, opts ...pulumi.ResourceOption) (*BackupStorageLocation, error) {
	var resource BackupStorageLocation
	err := ctx.ReadResource("tencentcloud:Kubernetes/backupStorageLocation:BackupStorageLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupStorageLocation resources.
type backupStorageLocationState struct {
	// Name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Message of the backup storage location.
	Message *string `pulumi:"message"`
	// Name of the backup storage location.
	Name *string `pulumi:"name"`
	// Prefix of the bucket.
	Path *string `pulumi:"path"`
	// State of the backup storage location.
	State *string `pulumi:"state"`
	// Region of the storage.
	StorageRegion *string `pulumi:"storageRegion"`
}

type BackupStorageLocationState struct {
	// Name of the bucket.
	Bucket pulumi.StringPtrInput
	// Message of the backup storage location.
	Message pulumi.StringPtrInput
	// Name of the backup storage location.
	Name pulumi.StringPtrInput
	// Prefix of the bucket.
	Path pulumi.StringPtrInput
	// State of the backup storage location.
	State pulumi.StringPtrInput
	// Region of the storage.
	StorageRegion pulumi.StringPtrInput
}

func (BackupStorageLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupStorageLocationState)(nil)).Elem()
}

type backupStorageLocationArgs struct {
	// Name of the bucket.
	Bucket string `pulumi:"bucket"`
	// Name of the backup storage location.
	Name *string `pulumi:"name"`
	// Prefix of the bucket.
	Path *string `pulumi:"path"`
	// Region of the storage.
	StorageRegion string `pulumi:"storageRegion"`
}

// The set of arguments for constructing a BackupStorageLocation resource.
type BackupStorageLocationArgs struct {
	// Name of the bucket.
	Bucket pulumi.StringInput
	// Name of the backup storage location.
	Name pulumi.StringPtrInput
	// Prefix of the bucket.
	Path pulumi.StringPtrInput
	// Region of the storage.
	StorageRegion pulumi.StringInput
}

func (BackupStorageLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupStorageLocationArgs)(nil)).Elem()
}

type BackupStorageLocationInput interface {
	pulumi.Input

	ToBackupStorageLocationOutput() BackupStorageLocationOutput
	ToBackupStorageLocationOutputWithContext(ctx context.Context) BackupStorageLocationOutput
}

func (*BackupStorageLocation) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupStorageLocation)(nil)).Elem()
}

func (i *BackupStorageLocation) ToBackupStorageLocationOutput() BackupStorageLocationOutput {
	return i.ToBackupStorageLocationOutputWithContext(context.Background())
}

func (i *BackupStorageLocation) ToBackupStorageLocationOutputWithContext(ctx context.Context) BackupStorageLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupStorageLocationOutput)
}

// BackupStorageLocationArrayInput is an input type that accepts BackupStorageLocationArray and BackupStorageLocationArrayOutput values.
// You can construct a concrete instance of `BackupStorageLocationArrayInput` via:
//
//          BackupStorageLocationArray{ BackupStorageLocationArgs{...} }
type BackupStorageLocationArrayInput interface {
	pulumi.Input

	ToBackupStorageLocationArrayOutput() BackupStorageLocationArrayOutput
	ToBackupStorageLocationArrayOutputWithContext(context.Context) BackupStorageLocationArrayOutput
}

type BackupStorageLocationArray []BackupStorageLocationInput

func (BackupStorageLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupStorageLocation)(nil)).Elem()
}

func (i BackupStorageLocationArray) ToBackupStorageLocationArrayOutput() BackupStorageLocationArrayOutput {
	return i.ToBackupStorageLocationArrayOutputWithContext(context.Background())
}

func (i BackupStorageLocationArray) ToBackupStorageLocationArrayOutputWithContext(ctx context.Context) BackupStorageLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupStorageLocationArrayOutput)
}

// BackupStorageLocationMapInput is an input type that accepts BackupStorageLocationMap and BackupStorageLocationMapOutput values.
// You can construct a concrete instance of `BackupStorageLocationMapInput` via:
//
//          BackupStorageLocationMap{ "key": BackupStorageLocationArgs{...} }
type BackupStorageLocationMapInput interface {
	pulumi.Input

	ToBackupStorageLocationMapOutput() BackupStorageLocationMapOutput
	ToBackupStorageLocationMapOutputWithContext(context.Context) BackupStorageLocationMapOutput
}

type BackupStorageLocationMap map[string]BackupStorageLocationInput

func (BackupStorageLocationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupStorageLocation)(nil)).Elem()
}

func (i BackupStorageLocationMap) ToBackupStorageLocationMapOutput() BackupStorageLocationMapOutput {
	return i.ToBackupStorageLocationMapOutputWithContext(context.Background())
}

func (i BackupStorageLocationMap) ToBackupStorageLocationMapOutputWithContext(ctx context.Context) BackupStorageLocationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupStorageLocationMapOutput)
}

type BackupStorageLocationOutput struct{ *pulumi.OutputState }

func (BackupStorageLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupStorageLocation)(nil)).Elem()
}

func (o BackupStorageLocationOutput) ToBackupStorageLocationOutput() BackupStorageLocationOutput {
	return o
}

func (o BackupStorageLocationOutput) ToBackupStorageLocationOutputWithContext(ctx context.Context) BackupStorageLocationOutput {
	return o
}

// Name of the bucket.
func (o BackupStorageLocationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStorageLocation) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Message of the backup storage location.
func (o BackupStorageLocationOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStorageLocation) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// Name of the backup storage location.
func (o BackupStorageLocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStorageLocation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Prefix of the bucket.
func (o BackupStorageLocationOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupStorageLocation) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// State of the backup storage location.
func (o BackupStorageLocationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStorageLocation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Region of the storage.
func (o BackupStorageLocationOutput) StorageRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStorageLocation) pulumi.StringOutput { return v.StorageRegion }).(pulumi.StringOutput)
}

type BackupStorageLocationArrayOutput struct{ *pulumi.OutputState }

func (BackupStorageLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupStorageLocation)(nil)).Elem()
}

func (o BackupStorageLocationArrayOutput) ToBackupStorageLocationArrayOutput() BackupStorageLocationArrayOutput {
	return o
}

func (o BackupStorageLocationArrayOutput) ToBackupStorageLocationArrayOutputWithContext(ctx context.Context) BackupStorageLocationArrayOutput {
	return o
}

func (o BackupStorageLocationArrayOutput) Index(i pulumi.IntInput) BackupStorageLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupStorageLocation {
		return vs[0].([]*BackupStorageLocation)[vs[1].(int)]
	}).(BackupStorageLocationOutput)
}

type BackupStorageLocationMapOutput struct{ *pulumi.OutputState }

func (BackupStorageLocationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupStorageLocation)(nil)).Elem()
}

func (o BackupStorageLocationMapOutput) ToBackupStorageLocationMapOutput() BackupStorageLocationMapOutput {
	return o
}

func (o BackupStorageLocationMapOutput) ToBackupStorageLocationMapOutputWithContext(ctx context.Context) BackupStorageLocationMapOutput {
	return o
}

func (o BackupStorageLocationMapOutput) MapIndex(k pulumi.StringInput) BackupStorageLocationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupStorageLocation {
		return vs[0].(map[string]*BackupStorageLocation)[vs[1].(string)]
	}).(BackupStorageLocationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupStorageLocationInput)(nil)).Elem(), &BackupStorageLocation{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupStorageLocationArrayInput)(nil)).Elem(), BackupStorageLocationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupStorageLocationMapInput)(nil)).Elem(), BackupStorageLocationMap{})
	pulumi.RegisterOutputType(BackupStorageLocationOutput{})
	pulumi.RegisterOutputType(BackupStorageLocationArrayOutput{})
	pulumi.RegisterOutputType(BackupStorageLocationMapOutput{})
}
