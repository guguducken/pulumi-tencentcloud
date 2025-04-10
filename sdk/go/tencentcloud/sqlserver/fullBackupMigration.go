// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver fullBackupMigration
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
// 		_, err := Sqlserver.NewFullBackupMigration(ctx, "myMigration", &Sqlserver.FullBackupMigrationArgs{
// 			BackupFiles:   pulumi.StringArray{},
// 			InstanceId:    pulumi.String("mssql-qelbzgwf"),
// 			MigrationName: pulumi.String("migration_test"),
// 			RecoveryType:  pulumi.String("FULL"),
// 			UploadType:    pulumi.String("COS_URL"),
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
// sqlserver full_backup_migration can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/fullBackupMigration:FullBackupMigration full_backup_migration full_backup_migration_id
// ```
type FullBackupMigration struct {
	pulumi.CustomResourceState

	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles pulumi.StringArrayOutput `pulumi:"backupFiles"`
	// ID of imported target instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Task name.
	MigrationName pulumi.StringOutput `pulumi:"migrationName"`
	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration.
	RecoveryType pulumi.StringOutput `pulumi:"recoveryType"`
	// Backup upload type. COS_URL: the backup is stored in users Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the applications Cloud Object Storage and needs to be uploaded by the user.
	UploadType pulumi.StringOutput `pulumi:"uploadType"`
}

// NewFullBackupMigration registers a new resource with the given unique name, arguments, and options.
func NewFullBackupMigration(ctx *pulumi.Context,
	name string, args *FullBackupMigrationArgs, opts ...pulumi.ResourceOption) (*FullBackupMigration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.MigrationName == nil {
		return nil, errors.New("invalid value for required argument 'MigrationName'")
	}
	if args.RecoveryType == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryType'")
	}
	if args.UploadType == nil {
		return nil, errors.New("invalid value for required argument 'UploadType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FullBackupMigration
	err := ctx.RegisterResource("tencentcloud:Sqlserver/fullBackupMigration:FullBackupMigration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFullBackupMigration gets an existing FullBackupMigration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFullBackupMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FullBackupMigrationState, opts ...pulumi.ResourceOption) (*FullBackupMigration, error) {
	var resource FullBackupMigration
	err := ctx.ReadResource("tencentcloud:Sqlserver/fullBackupMigration:FullBackupMigration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FullBackupMigration resources.
type fullBackupMigrationState struct {
	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []string `pulumi:"backupFiles"`
	// ID of imported target instance.
	InstanceId *string `pulumi:"instanceId"`
	// Task name.
	MigrationName *string `pulumi:"migrationName"`
	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration.
	RecoveryType *string `pulumi:"recoveryType"`
	// Backup upload type. COS_URL: the backup is stored in users Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the applications Cloud Object Storage and needs to be uploaded by the user.
	UploadType *string `pulumi:"uploadType"`
}

type FullBackupMigrationState struct {
	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles pulumi.StringArrayInput
	// ID of imported target instance.
	InstanceId pulumi.StringPtrInput
	// Task name.
	MigrationName pulumi.StringPtrInput
	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration.
	RecoveryType pulumi.StringPtrInput
	// Backup upload type. COS_URL: the backup is stored in users Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the applications Cloud Object Storage and needs to be uploaded by the user.
	UploadType pulumi.StringPtrInput
}

func (FullBackupMigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fullBackupMigrationState)(nil)).Elem()
}

type fullBackupMigrationArgs struct {
	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles []string `pulumi:"backupFiles"`
	// ID of imported target instance.
	InstanceId string `pulumi:"instanceId"`
	// Task name.
	MigrationName string `pulumi:"migrationName"`
	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration.
	RecoveryType string `pulumi:"recoveryType"`
	// Backup upload type. COS_URL: the backup is stored in users Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the applications Cloud Object Storage and needs to be uploaded by the user.
	UploadType string `pulumi:"uploadType"`
}

// The set of arguments for constructing a FullBackupMigration resource.
type FullBackupMigrationArgs struct {
	// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
	BackupFiles pulumi.StringArrayInput
	// ID of imported target instance.
	InstanceId pulumi.StringInput
	// Task name.
	MigrationName pulumi.StringInput
	// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration.
	RecoveryType pulumi.StringInput
	// Backup upload type. COS_URL: the backup is stored in users Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the applications Cloud Object Storage and needs to be uploaded by the user.
	UploadType pulumi.StringInput
}

func (FullBackupMigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fullBackupMigrationArgs)(nil)).Elem()
}

type FullBackupMigrationInput interface {
	pulumi.Input

	ToFullBackupMigrationOutput() FullBackupMigrationOutput
	ToFullBackupMigrationOutputWithContext(ctx context.Context) FullBackupMigrationOutput
}

func (*FullBackupMigration) ElementType() reflect.Type {
	return reflect.TypeOf((**FullBackupMigration)(nil)).Elem()
}

func (i *FullBackupMigration) ToFullBackupMigrationOutput() FullBackupMigrationOutput {
	return i.ToFullBackupMigrationOutputWithContext(context.Background())
}

func (i *FullBackupMigration) ToFullBackupMigrationOutputWithContext(ctx context.Context) FullBackupMigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullBackupMigrationOutput)
}

// FullBackupMigrationArrayInput is an input type that accepts FullBackupMigrationArray and FullBackupMigrationArrayOutput values.
// You can construct a concrete instance of `FullBackupMigrationArrayInput` via:
//
//          FullBackupMigrationArray{ FullBackupMigrationArgs{...} }
type FullBackupMigrationArrayInput interface {
	pulumi.Input

	ToFullBackupMigrationArrayOutput() FullBackupMigrationArrayOutput
	ToFullBackupMigrationArrayOutputWithContext(context.Context) FullBackupMigrationArrayOutput
}

type FullBackupMigrationArray []FullBackupMigrationInput

func (FullBackupMigrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FullBackupMigration)(nil)).Elem()
}

func (i FullBackupMigrationArray) ToFullBackupMigrationArrayOutput() FullBackupMigrationArrayOutput {
	return i.ToFullBackupMigrationArrayOutputWithContext(context.Background())
}

func (i FullBackupMigrationArray) ToFullBackupMigrationArrayOutputWithContext(ctx context.Context) FullBackupMigrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullBackupMigrationArrayOutput)
}

// FullBackupMigrationMapInput is an input type that accepts FullBackupMigrationMap and FullBackupMigrationMapOutput values.
// You can construct a concrete instance of `FullBackupMigrationMapInput` via:
//
//          FullBackupMigrationMap{ "key": FullBackupMigrationArgs{...} }
type FullBackupMigrationMapInput interface {
	pulumi.Input

	ToFullBackupMigrationMapOutput() FullBackupMigrationMapOutput
	ToFullBackupMigrationMapOutputWithContext(context.Context) FullBackupMigrationMapOutput
}

type FullBackupMigrationMap map[string]FullBackupMigrationInput

func (FullBackupMigrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FullBackupMigration)(nil)).Elem()
}

func (i FullBackupMigrationMap) ToFullBackupMigrationMapOutput() FullBackupMigrationMapOutput {
	return i.ToFullBackupMigrationMapOutputWithContext(context.Background())
}

func (i FullBackupMigrationMap) ToFullBackupMigrationMapOutputWithContext(ctx context.Context) FullBackupMigrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullBackupMigrationMapOutput)
}

type FullBackupMigrationOutput struct{ *pulumi.OutputState }

func (FullBackupMigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FullBackupMigration)(nil)).Elem()
}

func (o FullBackupMigrationOutput) ToFullBackupMigrationOutput() FullBackupMigrationOutput {
	return o
}

func (o FullBackupMigrationOutput) ToFullBackupMigrationOutputWithContext(ctx context.Context) FullBackupMigrationOutput {
	return o
}

// If the UploadType is COS_URL, fill in the URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
func (o FullBackupMigrationOutput) BackupFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FullBackupMigration) pulumi.StringArrayOutput { return v.BackupFiles }).(pulumi.StringArrayOutput)
}

// ID of imported target instance.
func (o FullBackupMigrationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FullBackupMigration) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Task name.
func (o FullBackupMigrationOutput) MigrationName() pulumi.StringOutput {
	return o.ApplyT(func(v *FullBackupMigration) pulumi.StringOutput { return v.MigrationName }).(pulumi.StringOutput)
}

// Migration task restoration type. FULL: full backup restoration, FULL_LOG: full backup and transaction log restoration, FULL_DIFF: full backup and differential backup restoration.
func (o FullBackupMigrationOutput) RecoveryType() pulumi.StringOutput {
	return o.ApplyT(func(v *FullBackupMigration) pulumi.StringOutput { return v.RecoveryType }).(pulumi.StringOutput)
}

// Backup upload type. COS_URL: the backup is stored in users Cloud Object Storage, with URL provided. COS_UPLOAD: the backup is stored in the applications Cloud Object Storage and needs to be uploaded by the user.
func (o FullBackupMigrationOutput) UploadType() pulumi.StringOutput {
	return o.ApplyT(func(v *FullBackupMigration) pulumi.StringOutput { return v.UploadType }).(pulumi.StringOutput)
}

type FullBackupMigrationArrayOutput struct{ *pulumi.OutputState }

func (FullBackupMigrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FullBackupMigration)(nil)).Elem()
}

func (o FullBackupMigrationArrayOutput) ToFullBackupMigrationArrayOutput() FullBackupMigrationArrayOutput {
	return o
}

func (o FullBackupMigrationArrayOutput) ToFullBackupMigrationArrayOutputWithContext(ctx context.Context) FullBackupMigrationArrayOutput {
	return o
}

func (o FullBackupMigrationArrayOutput) Index(i pulumi.IntInput) FullBackupMigrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FullBackupMigration {
		return vs[0].([]*FullBackupMigration)[vs[1].(int)]
	}).(FullBackupMigrationOutput)
}

type FullBackupMigrationMapOutput struct{ *pulumi.OutputState }

func (FullBackupMigrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FullBackupMigration)(nil)).Elem()
}

func (o FullBackupMigrationMapOutput) ToFullBackupMigrationMapOutput() FullBackupMigrationMapOutput {
	return o
}

func (o FullBackupMigrationMapOutput) ToFullBackupMigrationMapOutputWithContext(ctx context.Context) FullBackupMigrationMapOutput {
	return o
}

func (o FullBackupMigrationMapOutput) MapIndex(k pulumi.StringInput) FullBackupMigrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FullBackupMigration {
		return vs[0].(map[string]*FullBackupMigration)[vs[1].(string)]
	}).(FullBackupMigrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FullBackupMigrationInput)(nil)).Elem(), &FullBackupMigration{})
	pulumi.RegisterInputType(reflect.TypeOf((*FullBackupMigrationArrayInput)(nil)).Elem(), FullBackupMigrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FullBackupMigrationMapInput)(nil)).Elem(), FullBackupMigrationMap{})
	pulumi.RegisterOutputType(FullBackupMigrationOutput{})
	pulumi.RegisterOutputType(FullBackupMigrationArrayOutput{})
	pulumi.RegisterOutputType(FullBackupMigrationMapOutput{})
}
