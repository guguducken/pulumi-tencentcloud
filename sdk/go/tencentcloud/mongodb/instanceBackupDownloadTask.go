// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceBackupDownloadTask struct {
	pulumi.CustomResourceState

	// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
	BackupName pulumi.StringOutput `pulumi:"backupName"`
	// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
	// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
	// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
	// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
	// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
	BackupSets InstanceBackupDownloadTaskBackupSetArrayOutput `pulumi:"backupSets"`
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewInstanceBackupDownloadTask registers a new resource with the given unique name, arguments, and options.
func NewInstanceBackupDownloadTask(ctx *pulumi.Context,
	name string, args *InstanceBackupDownloadTaskArgs, opts ...pulumi.ResourceOption) (*InstanceBackupDownloadTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupName == nil {
		return nil, errors.New("invalid value for required argument 'BackupName'")
	}
	if args.BackupSets == nil {
		return nil, errors.New("invalid value for required argument 'BackupSets'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceBackupDownloadTask
	err := ctx.RegisterResource("tencentcloud:Mongodb/instanceBackupDownloadTask:InstanceBackupDownloadTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceBackupDownloadTask gets an existing InstanceBackupDownloadTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceBackupDownloadTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceBackupDownloadTaskState, opts ...pulumi.ResourceOption) (*InstanceBackupDownloadTask, error) {
	var resource InstanceBackupDownloadTask
	err := ctx.ReadResource("tencentcloud:Mongodb/instanceBackupDownloadTask:InstanceBackupDownloadTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceBackupDownloadTask resources.
type instanceBackupDownloadTaskState struct {
	// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
	BackupName *string `pulumi:"backupName"`
	// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
	// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
	// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
	// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
	// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
	BackupSets []InstanceBackupDownloadTaskBackupSet `pulumi:"backupSets"`
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId *string `pulumi:"instanceId"`
}

type InstanceBackupDownloadTaskState struct {
	// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
	BackupName pulumi.StringPtrInput
	// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
	// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
	// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
	// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
	// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
	BackupSets InstanceBackupDownloadTaskBackupSetArrayInput
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId pulumi.StringPtrInput
}

func (InstanceBackupDownloadTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceBackupDownloadTaskState)(nil)).Elem()
}

type instanceBackupDownloadTaskArgs struct {
	// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
	BackupName string `pulumi:"backupName"`
	// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
	// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
	// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
	// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
	// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
	BackupSets []InstanceBackupDownloadTaskBackupSet `pulumi:"backupSets"`
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a InstanceBackupDownloadTask resource.
type InstanceBackupDownloadTaskArgs struct {
	// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
	BackupName pulumi.StringInput
	// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
	// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
	// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
	// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
	// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
	BackupSets InstanceBackupDownloadTaskBackupSetArrayInput
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId pulumi.StringInput
}

func (InstanceBackupDownloadTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceBackupDownloadTaskArgs)(nil)).Elem()
}

type InstanceBackupDownloadTaskInput interface {
	pulumi.Input

	ToInstanceBackupDownloadTaskOutput() InstanceBackupDownloadTaskOutput
	ToInstanceBackupDownloadTaskOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskOutput
}

func (*InstanceBackupDownloadTask) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupDownloadTask)(nil)).Elem()
}

func (i *InstanceBackupDownloadTask) ToInstanceBackupDownloadTaskOutput() InstanceBackupDownloadTaskOutput {
	return i.ToInstanceBackupDownloadTaskOutputWithContext(context.Background())
}

func (i *InstanceBackupDownloadTask) ToInstanceBackupDownloadTaskOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupDownloadTaskOutput)
}

// InstanceBackupDownloadTaskArrayInput is an input type that accepts InstanceBackupDownloadTaskArray and InstanceBackupDownloadTaskArrayOutput values.
// You can construct a concrete instance of `InstanceBackupDownloadTaskArrayInput` via:
//
//          InstanceBackupDownloadTaskArray{ InstanceBackupDownloadTaskArgs{...} }
type InstanceBackupDownloadTaskArrayInput interface {
	pulumi.Input

	ToInstanceBackupDownloadTaskArrayOutput() InstanceBackupDownloadTaskArrayOutput
	ToInstanceBackupDownloadTaskArrayOutputWithContext(context.Context) InstanceBackupDownloadTaskArrayOutput
}

type InstanceBackupDownloadTaskArray []InstanceBackupDownloadTaskInput

func (InstanceBackupDownloadTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceBackupDownloadTask)(nil)).Elem()
}

func (i InstanceBackupDownloadTaskArray) ToInstanceBackupDownloadTaskArrayOutput() InstanceBackupDownloadTaskArrayOutput {
	return i.ToInstanceBackupDownloadTaskArrayOutputWithContext(context.Background())
}

func (i InstanceBackupDownloadTaskArray) ToInstanceBackupDownloadTaskArrayOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupDownloadTaskArrayOutput)
}

// InstanceBackupDownloadTaskMapInput is an input type that accepts InstanceBackupDownloadTaskMap and InstanceBackupDownloadTaskMapOutput values.
// You can construct a concrete instance of `InstanceBackupDownloadTaskMapInput` via:
//
//          InstanceBackupDownloadTaskMap{ "key": InstanceBackupDownloadTaskArgs{...} }
type InstanceBackupDownloadTaskMapInput interface {
	pulumi.Input

	ToInstanceBackupDownloadTaskMapOutput() InstanceBackupDownloadTaskMapOutput
	ToInstanceBackupDownloadTaskMapOutputWithContext(context.Context) InstanceBackupDownloadTaskMapOutput
}

type InstanceBackupDownloadTaskMap map[string]InstanceBackupDownloadTaskInput

func (InstanceBackupDownloadTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceBackupDownloadTask)(nil)).Elem()
}

func (i InstanceBackupDownloadTaskMap) ToInstanceBackupDownloadTaskMapOutput() InstanceBackupDownloadTaskMapOutput {
	return i.ToInstanceBackupDownloadTaskMapOutputWithContext(context.Background())
}

func (i InstanceBackupDownloadTaskMap) ToInstanceBackupDownloadTaskMapOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupDownloadTaskMapOutput)
}

type InstanceBackupDownloadTaskOutput struct{ *pulumi.OutputState }

func (InstanceBackupDownloadTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupDownloadTask)(nil)).Elem()
}

func (o InstanceBackupDownloadTaskOutput) ToInstanceBackupDownloadTaskOutput() InstanceBackupDownloadTaskOutput {
	return o
}

func (o InstanceBackupDownloadTaskOutput) ToInstanceBackupDownloadTaskOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskOutput {
	return o
}

// The name of the backup file to be downloaded can be obtained through the DescribeDBBackups interface.
func (o InstanceBackupDownloadTaskOutput) BackupName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupDownloadTask) pulumi.StringOutput { return v.BackupName }).(pulumi.StringOutput)
}

// Specifies the node names of replica sets to download or a list of shard names for sharded clusters.For example, the
// replica set cmgo-p8vnipr5, example (fixed value): BackupSets.0=cmgo-p8vnipr5_0, the full amount of data can be
// downloaded.For example, the sharded cluster cmgo-p8vnipr5, for example:
// BackupSets.0=cmgo-p8vnipr5_0&amp;amp;BackupSets.1=cmgo-p8vnipr5_1, that is, to download the data of shard 0 and 1. If
// the sharded cluster needs to be downloaded in full, please pass in the example. Full slice name.
func (o InstanceBackupDownloadTaskOutput) BackupSets() InstanceBackupDownloadTaskBackupSetArrayOutput {
	return o.ApplyT(func(v *InstanceBackupDownloadTask) InstanceBackupDownloadTaskBackupSetArrayOutput {
		return v.BackupSets
	}).(InstanceBackupDownloadTaskBackupSetArrayOutput)
}

// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
func (o InstanceBackupDownloadTaskOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupDownloadTask) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type InstanceBackupDownloadTaskArrayOutput struct{ *pulumi.OutputState }

func (InstanceBackupDownloadTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceBackupDownloadTask)(nil)).Elem()
}

func (o InstanceBackupDownloadTaskArrayOutput) ToInstanceBackupDownloadTaskArrayOutput() InstanceBackupDownloadTaskArrayOutput {
	return o
}

func (o InstanceBackupDownloadTaskArrayOutput) ToInstanceBackupDownloadTaskArrayOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskArrayOutput {
	return o
}

func (o InstanceBackupDownloadTaskArrayOutput) Index(i pulumi.IntInput) InstanceBackupDownloadTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceBackupDownloadTask {
		return vs[0].([]*InstanceBackupDownloadTask)[vs[1].(int)]
	}).(InstanceBackupDownloadTaskOutput)
}

type InstanceBackupDownloadTaskMapOutput struct{ *pulumi.OutputState }

func (InstanceBackupDownloadTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceBackupDownloadTask)(nil)).Elem()
}

func (o InstanceBackupDownloadTaskMapOutput) ToInstanceBackupDownloadTaskMapOutput() InstanceBackupDownloadTaskMapOutput {
	return o
}

func (o InstanceBackupDownloadTaskMapOutput) ToInstanceBackupDownloadTaskMapOutputWithContext(ctx context.Context) InstanceBackupDownloadTaskMapOutput {
	return o
}

func (o InstanceBackupDownloadTaskMapOutput) MapIndex(k pulumi.StringInput) InstanceBackupDownloadTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceBackupDownloadTask {
		return vs[0].(map[string]*InstanceBackupDownloadTask)[vs[1].(string)]
	}).(InstanceBackupDownloadTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupDownloadTaskInput)(nil)).Elem(), &InstanceBackupDownloadTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupDownloadTaskArrayInput)(nil)).Elem(), InstanceBackupDownloadTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupDownloadTaskMapInput)(nil)).Elem(), InstanceBackupDownloadTaskMap{})
	pulumi.RegisterOutputType(InstanceBackupDownloadTaskOutput{})
	pulumi.RegisterOutputType(InstanceBackupDownloadTaskArrayOutput{})
	pulumi.RegisterOutputType(InstanceBackupDownloadTaskMapOutput{})
}
