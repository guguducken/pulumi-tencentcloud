// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetSnapshotPoliciesSnapshotPolicyList struct {
	// Storage IDs that the snapshot policy attached.
	AttachedStorageIds []string `pulumi:"attachedStorageIds"`
	// Create time of the snapshot policy.
	CreateTime string `pulumi:"createTime"`
	// Trigger hours of periodic snapshot.
	RepeatHours []int `pulumi:"repeatHours"`
	// Trigger days of periodic snapshot.
	RepeatWeekdays []int `pulumi:"repeatWeekdays"`
	// Retention days of the snapshot.
	RetentionDays int `pulumi:"retentionDays"`
	// ID of the snapshot policy to be queried.
	SnapshotPolicyId string `pulumi:"snapshotPolicyId"`
	// Name of the snapshot policy to be queried.
	SnapshotPolicyName string `pulumi:"snapshotPolicyName"`
	// Status of the snapshot policy.
	Status string `pulumi:"status"`
}

// GetSnapshotPoliciesSnapshotPolicyListInput is an input type that accepts GetSnapshotPoliciesSnapshotPolicyListArgs and GetSnapshotPoliciesSnapshotPolicyListOutput values.
// You can construct a concrete instance of `GetSnapshotPoliciesSnapshotPolicyListInput` via:
//
//          GetSnapshotPoliciesSnapshotPolicyListArgs{...}
type GetSnapshotPoliciesSnapshotPolicyListInput interface {
	pulumi.Input

	ToGetSnapshotPoliciesSnapshotPolicyListOutput() GetSnapshotPoliciesSnapshotPolicyListOutput
	ToGetSnapshotPoliciesSnapshotPolicyListOutputWithContext(context.Context) GetSnapshotPoliciesSnapshotPolicyListOutput
}

type GetSnapshotPoliciesSnapshotPolicyListArgs struct {
	// Storage IDs that the snapshot policy attached.
	AttachedStorageIds pulumi.StringArrayInput `pulumi:"attachedStorageIds"`
	// Create time of the snapshot policy.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Trigger hours of periodic snapshot.
	RepeatHours pulumi.IntArrayInput `pulumi:"repeatHours"`
	// Trigger days of periodic snapshot.
	RepeatWeekdays pulumi.IntArrayInput `pulumi:"repeatWeekdays"`
	// Retention days of the snapshot.
	RetentionDays pulumi.IntInput `pulumi:"retentionDays"`
	// ID of the snapshot policy to be queried.
	SnapshotPolicyId pulumi.StringInput `pulumi:"snapshotPolicyId"`
	// Name of the snapshot policy to be queried.
	SnapshotPolicyName pulumi.StringInput `pulumi:"snapshotPolicyName"`
	// Status of the snapshot policy.
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetSnapshotPoliciesSnapshotPolicyListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotPoliciesSnapshotPolicyList)(nil)).Elem()
}

func (i GetSnapshotPoliciesSnapshotPolicyListArgs) ToGetSnapshotPoliciesSnapshotPolicyListOutput() GetSnapshotPoliciesSnapshotPolicyListOutput {
	return i.ToGetSnapshotPoliciesSnapshotPolicyListOutputWithContext(context.Background())
}

func (i GetSnapshotPoliciesSnapshotPolicyListArgs) ToGetSnapshotPoliciesSnapshotPolicyListOutputWithContext(ctx context.Context) GetSnapshotPoliciesSnapshotPolicyListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSnapshotPoliciesSnapshotPolicyListOutput)
}

// GetSnapshotPoliciesSnapshotPolicyListArrayInput is an input type that accepts GetSnapshotPoliciesSnapshotPolicyListArray and GetSnapshotPoliciesSnapshotPolicyListArrayOutput values.
// You can construct a concrete instance of `GetSnapshotPoliciesSnapshotPolicyListArrayInput` via:
//
//          GetSnapshotPoliciesSnapshotPolicyListArray{ GetSnapshotPoliciesSnapshotPolicyListArgs{...} }
type GetSnapshotPoliciesSnapshotPolicyListArrayInput interface {
	pulumi.Input

	ToGetSnapshotPoliciesSnapshotPolicyListArrayOutput() GetSnapshotPoliciesSnapshotPolicyListArrayOutput
	ToGetSnapshotPoliciesSnapshotPolicyListArrayOutputWithContext(context.Context) GetSnapshotPoliciesSnapshotPolicyListArrayOutput
}

type GetSnapshotPoliciesSnapshotPolicyListArray []GetSnapshotPoliciesSnapshotPolicyListInput

func (GetSnapshotPoliciesSnapshotPolicyListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSnapshotPoliciesSnapshotPolicyList)(nil)).Elem()
}

func (i GetSnapshotPoliciesSnapshotPolicyListArray) ToGetSnapshotPoliciesSnapshotPolicyListArrayOutput() GetSnapshotPoliciesSnapshotPolicyListArrayOutput {
	return i.ToGetSnapshotPoliciesSnapshotPolicyListArrayOutputWithContext(context.Background())
}

func (i GetSnapshotPoliciesSnapshotPolicyListArray) ToGetSnapshotPoliciesSnapshotPolicyListArrayOutputWithContext(ctx context.Context) GetSnapshotPoliciesSnapshotPolicyListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSnapshotPoliciesSnapshotPolicyListArrayOutput)
}

type GetSnapshotPoliciesSnapshotPolicyListOutput struct{ *pulumi.OutputState }

func (GetSnapshotPoliciesSnapshotPolicyListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotPoliciesSnapshotPolicyList)(nil)).Elem()
}

func (o GetSnapshotPoliciesSnapshotPolicyListOutput) ToGetSnapshotPoliciesSnapshotPolicyListOutput() GetSnapshotPoliciesSnapshotPolicyListOutput {
	return o
}

func (o GetSnapshotPoliciesSnapshotPolicyListOutput) ToGetSnapshotPoliciesSnapshotPolicyListOutputWithContext(ctx context.Context) GetSnapshotPoliciesSnapshotPolicyListOutput {
	return o
}

// Storage IDs that the snapshot policy attached.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) AttachedStorageIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) []string { return v.AttachedStorageIds }).(pulumi.StringArrayOutput)
}

// Create time of the snapshot policy.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Trigger hours of periodic snapshot.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) RepeatHours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) []int { return v.RepeatHours }).(pulumi.IntArrayOutput)
}

// Trigger days of periodic snapshot.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) RepeatWeekdays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) []int { return v.RepeatWeekdays }).(pulumi.IntArrayOutput)
}

// Retention days of the snapshot.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) int { return v.RetentionDays }).(pulumi.IntOutput)
}

// ID of the snapshot policy to be queried.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) SnapshotPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) string { return v.SnapshotPolicyId }).(pulumi.StringOutput)
}

// Name of the snapshot policy to be queried.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) SnapshotPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) string { return v.SnapshotPolicyName }).(pulumi.StringOutput)
}

// Status of the snapshot policy.
func (o GetSnapshotPoliciesSnapshotPolicyListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesSnapshotPolicyList) string { return v.Status }).(pulumi.StringOutput)
}

type GetSnapshotPoliciesSnapshotPolicyListArrayOutput struct{ *pulumi.OutputState }

func (GetSnapshotPoliciesSnapshotPolicyListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSnapshotPoliciesSnapshotPolicyList)(nil)).Elem()
}

func (o GetSnapshotPoliciesSnapshotPolicyListArrayOutput) ToGetSnapshotPoliciesSnapshotPolicyListArrayOutput() GetSnapshotPoliciesSnapshotPolicyListArrayOutput {
	return o
}

func (o GetSnapshotPoliciesSnapshotPolicyListArrayOutput) ToGetSnapshotPoliciesSnapshotPolicyListArrayOutputWithContext(ctx context.Context) GetSnapshotPoliciesSnapshotPolicyListArrayOutput {
	return o
}

func (o GetSnapshotPoliciesSnapshotPolicyListArrayOutput) Index(i pulumi.IntInput) GetSnapshotPoliciesSnapshotPolicyListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSnapshotPoliciesSnapshotPolicyList {
		return vs[0].([]GetSnapshotPoliciesSnapshotPolicyList)[vs[1].(int)]
	}).(GetSnapshotPoliciesSnapshotPolicyListOutput)
}

type GetSnapshotsSnapshotList struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Creation time of snapshot.
	CreateTime string `pulumi:"createTime"`
	// Indicates whether the snapshot is encrypted.
	Encrypt bool `pulumi:"encrypt"`
	// Snapshot creation progress percentage.
	Percent int `pulumi:"percent"`
	// ID of the project within the snapshot.
	ProjectId int `pulumi:"projectId"`
	// ID of the snapshot to be queried.
	SnapshotId string `pulumi:"snapshotId"`
	// Name of the snapshot to be queried.
	SnapshotName string `pulumi:"snapshotName"`
	// ID of the the CBS which this snapshot created from.
	StorageId string `pulumi:"storageId"`
	// Volume of storage which this snapshot created from.
	StorageSize int `pulumi:"storageSize"`
	// Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
	StorageUsage string `pulumi:"storageUsage"`
}

// GetSnapshotsSnapshotListInput is an input type that accepts GetSnapshotsSnapshotListArgs and GetSnapshotsSnapshotListOutput values.
// You can construct a concrete instance of `GetSnapshotsSnapshotListInput` via:
//
//          GetSnapshotsSnapshotListArgs{...}
type GetSnapshotsSnapshotListInput interface {
	pulumi.Input

	ToGetSnapshotsSnapshotListOutput() GetSnapshotsSnapshotListOutput
	ToGetSnapshotsSnapshotListOutputWithContext(context.Context) GetSnapshotsSnapshotListOutput
}

type GetSnapshotsSnapshotListArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// Creation time of snapshot.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Indicates whether the snapshot is encrypted.
	Encrypt pulumi.BoolInput `pulumi:"encrypt"`
	// Snapshot creation progress percentage.
	Percent pulumi.IntInput `pulumi:"percent"`
	// ID of the project within the snapshot.
	ProjectId pulumi.IntInput `pulumi:"projectId"`
	// ID of the snapshot to be queried.
	SnapshotId pulumi.StringInput `pulumi:"snapshotId"`
	// Name of the snapshot to be queried.
	SnapshotName pulumi.StringInput `pulumi:"snapshotName"`
	// ID of the the CBS which this snapshot created from.
	StorageId pulumi.StringInput `pulumi:"storageId"`
	// Volume of storage which this snapshot created from.
	StorageSize pulumi.IntInput `pulumi:"storageSize"`
	// Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
	StorageUsage pulumi.StringInput `pulumi:"storageUsage"`
}

func (GetSnapshotsSnapshotListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsSnapshotList)(nil)).Elem()
}

func (i GetSnapshotsSnapshotListArgs) ToGetSnapshotsSnapshotListOutput() GetSnapshotsSnapshotListOutput {
	return i.ToGetSnapshotsSnapshotListOutputWithContext(context.Background())
}

func (i GetSnapshotsSnapshotListArgs) ToGetSnapshotsSnapshotListOutputWithContext(ctx context.Context) GetSnapshotsSnapshotListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSnapshotsSnapshotListOutput)
}

// GetSnapshotsSnapshotListArrayInput is an input type that accepts GetSnapshotsSnapshotListArray and GetSnapshotsSnapshotListArrayOutput values.
// You can construct a concrete instance of `GetSnapshotsSnapshotListArrayInput` via:
//
//          GetSnapshotsSnapshotListArray{ GetSnapshotsSnapshotListArgs{...} }
type GetSnapshotsSnapshotListArrayInput interface {
	pulumi.Input

	ToGetSnapshotsSnapshotListArrayOutput() GetSnapshotsSnapshotListArrayOutput
	ToGetSnapshotsSnapshotListArrayOutputWithContext(context.Context) GetSnapshotsSnapshotListArrayOutput
}

type GetSnapshotsSnapshotListArray []GetSnapshotsSnapshotListInput

func (GetSnapshotsSnapshotListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSnapshotsSnapshotList)(nil)).Elem()
}

func (i GetSnapshotsSnapshotListArray) ToGetSnapshotsSnapshotListArrayOutput() GetSnapshotsSnapshotListArrayOutput {
	return i.ToGetSnapshotsSnapshotListArrayOutputWithContext(context.Background())
}

func (i GetSnapshotsSnapshotListArray) ToGetSnapshotsSnapshotListArrayOutputWithContext(ctx context.Context) GetSnapshotsSnapshotListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSnapshotsSnapshotListArrayOutput)
}

type GetSnapshotsSnapshotListOutput struct{ *pulumi.OutputState }

func (GetSnapshotsSnapshotListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsSnapshotList)(nil)).Elem()
}

func (o GetSnapshotsSnapshotListOutput) ToGetSnapshotsSnapshotListOutput() GetSnapshotsSnapshotListOutput {
	return o
}

func (o GetSnapshotsSnapshotListOutput) ToGetSnapshotsSnapshotListOutputWithContext(ctx context.Context) GetSnapshotsSnapshotListOutput {
	return o
}

// The available zone that the CBS instance locates at.
func (o GetSnapshotsSnapshotListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Creation time of snapshot.
func (o GetSnapshotsSnapshotListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether the snapshot is encrypted.
func (o GetSnapshotsSnapshotListOutput) Encrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) bool { return v.Encrypt }).(pulumi.BoolOutput)
}

// Snapshot creation progress percentage.
func (o GetSnapshotsSnapshotListOutput) Percent() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) int { return v.Percent }).(pulumi.IntOutput)
}

// ID of the project within the snapshot.
func (o GetSnapshotsSnapshotListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) int { return v.ProjectId }).(pulumi.IntOutput)
}

// ID of the snapshot to be queried.
func (o GetSnapshotsSnapshotListOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) string { return v.SnapshotId }).(pulumi.StringOutput)
}

// Name of the snapshot to be queried.
func (o GetSnapshotsSnapshotListOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) string { return v.SnapshotName }).(pulumi.StringOutput)
}

// ID of the the CBS which this snapshot created from.
func (o GetSnapshotsSnapshotListOutput) StorageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) string { return v.StorageId }).(pulumi.StringOutput)
}

// Volume of storage which this snapshot created from.
func (o GetSnapshotsSnapshotListOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) int { return v.StorageSize }).(pulumi.IntOutput)
}

// Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
func (o GetSnapshotsSnapshotListOutput) StorageUsage() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsSnapshotList) string { return v.StorageUsage }).(pulumi.StringOutput)
}

type GetSnapshotsSnapshotListArrayOutput struct{ *pulumi.OutputState }

func (GetSnapshotsSnapshotListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSnapshotsSnapshotList)(nil)).Elem()
}

func (o GetSnapshotsSnapshotListArrayOutput) ToGetSnapshotsSnapshotListArrayOutput() GetSnapshotsSnapshotListArrayOutput {
	return o
}

func (o GetSnapshotsSnapshotListArrayOutput) ToGetSnapshotsSnapshotListArrayOutputWithContext(ctx context.Context) GetSnapshotsSnapshotListArrayOutput {
	return o
}

func (o GetSnapshotsSnapshotListArrayOutput) Index(i pulumi.IntInput) GetSnapshotsSnapshotListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSnapshotsSnapshotList {
		return vs[0].([]GetSnapshotsSnapshotList)[vs[1].(int)]
	}).(GetSnapshotsSnapshotListOutput)
}

type GetStoragesSetStorageList struct {
	// Indicates whether the CBS is mounted the CVM.
	Attached bool `pulumi:"attached"`
	// The available zone that the CBS instance locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
	ChargeType string `pulumi:"chargeType"`
	// Creation time of CBS.
	CreateTime string `pulumi:"createTime"`
	// Indicates whether CBS is encrypted.
	Encrypt bool `pulumi:"encrypt"`
	// ID of the CVM instance that be mounted by this CBS.
	InstanceId string `pulumi:"instanceId"`
	// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
	PrepaidRenewFlag string `pulumi:"prepaidRenewFlag"`
	// ID of the project with which the CBS is associated.
	ProjectId int `pulumi:"projectId"`
	// Status of CBS.
	Status string `pulumi:"status"`
	// ID of the CBS to be queried.
	StorageId string `pulumi:"storageId"`
	// Name of the CBS to be queried.
	StorageName string `pulumi:"storageName"`
	// Volume of CBS.
	StorageSize int `pulumi:"storageSize"`
	// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
	StorageType string `pulumi:"storageType"`
	// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
	StorageUsage string `pulumi:"storageUsage"`
	// The available tags within this CBS.
	Tags map[string]interface{} `pulumi:"tags"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance int `pulumi:"throughputPerformance"`
}

// GetStoragesSetStorageListInput is an input type that accepts GetStoragesSetStorageListArgs and GetStoragesSetStorageListOutput values.
// You can construct a concrete instance of `GetStoragesSetStorageListInput` via:
//
//          GetStoragesSetStorageListArgs{...}
type GetStoragesSetStorageListInput interface {
	pulumi.Input

	ToGetStoragesSetStorageListOutput() GetStoragesSetStorageListOutput
	ToGetStoragesSetStorageListOutputWithContext(context.Context) GetStoragesSetStorageListOutput
}

type GetStoragesSetStorageListArgs struct {
	// Indicates whether the CBS is mounted the CVM.
	Attached pulumi.BoolInput `pulumi:"attached"`
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
	ChargeType pulumi.StringInput `pulumi:"chargeType"`
	// Creation time of CBS.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Indicates whether CBS is encrypted.
	Encrypt pulumi.BoolInput `pulumi:"encrypt"`
	// ID of the CVM instance that be mounted by this CBS.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
	PrepaidRenewFlag pulumi.StringInput `pulumi:"prepaidRenewFlag"`
	// ID of the project with which the CBS is associated.
	ProjectId pulumi.IntInput `pulumi:"projectId"`
	// Status of CBS.
	Status pulumi.StringInput `pulumi:"status"`
	// ID of the CBS to be queried.
	StorageId pulumi.StringInput `pulumi:"storageId"`
	// Name of the CBS to be queried.
	StorageName pulumi.StringInput `pulumi:"storageName"`
	// Volume of CBS.
	StorageSize pulumi.IntInput `pulumi:"storageSize"`
	// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
	StorageType pulumi.StringInput `pulumi:"storageType"`
	// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
	StorageUsage pulumi.StringInput `pulumi:"storageUsage"`
	// The available tags within this CBS.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance pulumi.IntInput `pulumi:"throughputPerformance"`
}

func (GetStoragesSetStorageListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesSetStorageList)(nil)).Elem()
}

func (i GetStoragesSetStorageListArgs) ToGetStoragesSetStorageListOutput() GetStoragesSetStorageListOutput {
	return i.ToGetStoragesSetStorageListOutputWithContext(context.Background())
}

func (i GetStoragesSetStorageListArgs) ToGetStoragesSetStorageListOutputWithContext(ctx context.Context) GetStoragesSetStorageListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStoragesSetStorageListOutput)
}

// GetStoragesSetStorageListArrayInput is an input type that accepts GetStoragesSetStorageListArray and GetStoragesSetStorageListArrayOutput values.
// You can construct a concrete instance of `GetStoragesSetStorageListArrayInput` via:
//
//          GetStoragesSetStorageListArray{ GetStoragesSetStorageListArgs{...} }
type GetStoragesSetStorageListArrayInput interface {
	pulumi.Input

	ToGetStoragesSetStorageListArrayOutput() GetStoragesSetStorageListArrayOutput
	ToGetStoragesSetStorageListArrayOutputWithContext(context.Context) GetStoragesSetStorageListArrayOutput
}

type GetStoragesSetStorageListArray []GetStoragesSetStorageListInput

func (GetStoragesSetStorageListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStoragesSetStorageList)(nil)).Elem()
}

func (i GetStoragesSetStorageListArray) ToGetStoragesSetStorageListArrayOutput() GetStoragesSetStorageListArrayOutput {
	return i.ToGetStoragesSetStorageListArrayOutputWithContext(context.Background())
}

func (i GetStoragesSetStorageListArray) ToGetStoragesSetStorageListArrayOutputWithContext(ctx context.Context) GetStoragesSetStorageListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStoragesSetStorageListArrayOutput)
}

type GetStoragesSetStorageListOutput struct{ *pulumi.OutputState }

func (GetStoragesSetStorageListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesSetStorageList)(nil)).Elem()
}

func (o GetStoragesSetStorageListOutput) ToGetStoragesSetStorageListOutput() GetStoragesSetStorageListOutput {
	return o
}

func (o GetStoragesSetStorageListOutput) ToGetStoragesSetStorageListOutputWithContext(ctx context.Context) GetStoragesSetStorageListOutput {
	return o
}

// Indicates whether the CBS is mounted the CVM.
func (o GetStoragesSetStorageListOutput) Attached() pulumi.BoolOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) bool { return v.Attached }).(pulumi.BoolOutput)
}

// The available zone that the CBS instance locates at.
func (o GetStoragesSetStorageListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
func (o GetStoragesSetStorageListOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.ChargeType }).(pulumi.StringOutput)
}

// Creation time of CBS.
func (o GetStoragesSetStorageListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether CBS is encrypted.
func (o GetStoragesSetStorageListOutput) Encrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) bool { return v.Encrypt }).(pulumi.BoolOutput)
}

// ID of the CVM instance that be mounted by this CBS.
func (o GetStoragesSetStorageListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
func (o GetStoragesSetStorageListOutput) PrepaidRenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.PrepaidRenewFlag }).(pulumi.StringOutput)
}

// ID of the project with which the CBS is associated.
func (o GetStoragesSetStorageListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) int { return v.ProjectId }).(pulumi.IntOutput)
}

// Status of CBS.
func (o GetStoragesSetStorageListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.Status }).(pulumi.StringOutput)
}

// ID of the CBS to be queried.
func (o GetStoragesSetStorageListOutput) StorageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.StorageId }).(pulumi.StringOutput)
}

// Name of the CBS to be queried.
func (o GetStoragesSetStorageListOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.StorageName }).(pulumi.StringOutput)
}

// Volume of CBS.
func (o GetStoragesSetStorageListOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) int { return v.StorageSize }).(pulumi.IntOutput)
}

// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
func (o GetStoragesSetStorageListOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.StorageType }).(pulumi.StringOutput)
}

// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
func (o GetStoragesSetStorageListOutput) StorageUsage() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) string { return v.StorageUsage }).(pulumi.StringOutput)
}

// The available tags within this CBS.
func (o GetStoragesSetStorageListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
func (o GetStoragesSetStorageListOutput) ThroughputPerformance() pulumi.IntOutput {
	return o.ApplyT(func(v GetStoragesSetStorageList) int { return v.ThroughputPerformance }).(pulumi.IntOutput)
}

type GetStoragesSetStorageListArrayOutput struct{ *pulumi.OutputState }

func (GetStoragesSetStorageListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStoragesSetStorageList)(nil)).Elem()
}

func (o GetStoragesSetStorageListArrayOutput) ToGetStoragesSetStorageListArrayOutput() GetStoragesSetStorageListArrayOutput {
	return o
}

func (o GetStoragesSetStorageListArrayOutput) ToGetStoragesSetStorageListArrayOutputWithContext(ctx context.Context) GetStoragesSetStorageListArrayOutput {
	return o
}

func (o GetStoragesSetStorageListArrayOutput) Index(i pulumi.IntInput) GetStoragesSetStorageListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetStoragesSetStorageList {
		return vs[0].([]GetStoragesSetStorageList)[vs[1].(int)]
	}).(GetStoragesSetStorageListOutput)
}

type GetStoragesStorageList struct {
	// Indicates whether the CBS is mounted the CVM.
	Attached bool `pulumi:"attached"`
	// The available zone that the CBS instance locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
	ChargeType string `pulumi:"chargeType"`
	// Creation time of CBS.
	CreateTime string `pulumi:"createTime"`
	// Indicates whether CBS is encrypted.
	Encrypt bool `pulumi:"encrypt"`
	// ID of the CVM instance that be mounted by this CBS.
	InstanceId string `pulumi:"instanceId"`
	// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
	PrepaidRenewFlag string `pulumi:"prepaidRenewFlag"`
	// ID of the project with which the CBS is associated.
	ProjectId int `pulumi:"projectId"`
	// Status of CBS.
	Status string `pulumi:"status"`
	// ID of the CBS to be queried.
	StorageId string `pulumi:"storageId"`
	// Name of the CBS to be queried.
	StorageName string `pulumi:"storageName"`
	// Volume of CBS.
	StorageSize int `pulumi:"storageSize"`
	// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
	StorageType string `pulumi:"storageType"`
	// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
	StorageUsage string `pulumi:"storageUsage"`
	// The available tags within this CBS.
	Tags map[string]interface{} `pulumi:"tags"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance int `pulumi:"throughputPerformance"`
}

// GetStoragesStorageListInput is an input type that accepts GetStoragesStorageListArgs and GetStoragesStorageListOutput values.
// You can construct a concrete instance of `GetStoragesStorageListInput` via:
//
//          GetStoragesStorageListArgs{...}
type GetStoragesStorageListInput interface {
	pulumi.Input

	ToGetStoragesStorageListOutput() GetStoragesStorageListOutput
	ToGetStoragesStorageListOutputWithContext(context.Context) GetStoragesStorageListOutput
}

type GetStoragesStorageListArgs struct {
	// Indicates whether the CBS is mounted the CVM.
	Attached pulumi.BoolInput `pulumi:"attached"`
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
	ChargeType pulumi.StringInput `pulumi:"chargeType"`
	// Creation time of CBS.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Indicates whether CBS is encrypted.
	Encrypt pulumi.BoolInput `pulumi:"encrypt"`
	// ID of the CVM instance that be mounted by this CBS.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
	PrepaidRenewFlag pulumi.StringInput `pulumi:"prepaidRenewFlag"`
	// ID of the project with which the CBS is associated.
	ProjectId pulumi.IntInput `pulumi:"projectId"`
	// Status of CBS.
	Status pulumi.StringInput `pulumi:"status"`
	// ID of the CBS to be queried.
	StorageId pulumi.StringInput `pulumi:"storageId"`
	// Name of the CBS to be queried.
	StorageName pulumi.StringInput `pulumi:"storageName"`
	// Volume of CBS.
	StorageSize pulumi.IntInput `pulumi:"storageSize"`
	// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
	StorageType pulumi.StringInput `pulumi:"storageType"`
	// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
	StorageUsage pulumi.StringInput `pulumi:"storageUsage"`
	// The available tags within this CBS.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance pulumi.IntInput `pulumi:"throughputPerformance"`
}

func (GetStoragesStorageListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesStorageList)(nil)).Elem()
}

func (i GetStoragesStorageListArgs) ToGetStoragesStorageListOutput() GetStoragesStorageListOutput {
	return i.ToGetStoragesStorageListOutputWithContext(context.Background())
}

func (i GetStoragesStorageListArgs) ToGetStoragesStorageListOutputWithContext(ctx context.Context) GetStoragesStorageListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStoragesStorageListOutput)
}

// GetStoragesStorageListArrayInput is an input type that accepts GetStoragesStorageListArray and GetStoragesStorageListArrayOutput values.
// You can construct a concrete instance of `GetStoragesStorageListArrayInput` via:
//
//          GetStoragesStorageListArray{ GetStoragesStorageListArgs{...} }
type GetStoragesStorageListArrayInput interface {
	pulumi.Input

	ToGetStoragesStorageListArrayOutput() GetStoragesStorageListArrayOutput
	ToGetStoragesStorageListArrayOutputWithContext(context.Context) GetStoragesStorageListArrayOutput
}

type GetStoragesStorageListArray []GetStoragesStorageListInput

func (GetStoragesStorageListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStoragesStorageList)(nil)).Elem()
}

func (i GetStoragesStorageListArray) ToGetStoragesStorageListArrayOutput() GetStoragesStorageListArrayOutput {
	return i.ToGetStoragesStorageListArrayOutputWithContext(context.Background())
}

func (i GetStoragesStorageListArray) ToGetStoragesStorageListArrayOutputWithContext(ctx context.Context) GetStoragesStorageListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStoragesStorageListArrayOutput)
}

type GetStoragesStorageListOutput struct{ *pulumi.OutputState }

func (GetStoragesStorageListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesStorageList)(nil)).Elem()
}

func (o GetStoragesStorageListOutput) ToGetStoragesStorageListOutput() GetStoragesStorageListOutput {
	return o
}

func (o GetStoragesStorageListOutput) ToGetStoragesStorageListOutputWithContext(ctx context.Context) GetStoragesStorageListOutput {
	return o
}

// Indicates whether the CBS is mounted the CVM.
func (o GetStoragesStorageListOutput) Attached() pulumi.BoolOutput {
	return o.ApplyT(func(v GetStoragesStorageList) bool { return v.Attached }).(pulumi.BoolOutput)
}

// The available zone that the CBS instance locates at.
func (o GetStoragesStorageListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
func (o GetStoragesStorageListOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.ChargeType }).(pulumi.StringOutput)
}

// Creation time of CBS.
func (o GetStoragesStorageListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether CBS is encrypted.
func (o GetStoragesStorageListOutput) Encrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v GetStoragesStorageList) bool { return v.Encrypt }).(pulumi.BoolOutput)
}

// ID of the CVM instance that be mounted by this CBS.
func (o GetStoragesStorageListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
func (o GetStoragesStorageListOutput) PrepaidRenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.PrepaidRenewFlag }).(pulumi.StringOutput)
}

// ID of the project with which the CBS is associated.
func (o GetStoragesStorageListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetStoragesStorageList) int { return v.ProjectId }).(pulumi.IntOutput)
}

// Status of CBS.
func (o GetStoragesStorageListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.Status }).(pulumi.StringOutput)
}

// ID of the CBS to be queried.
func (o GetStoragesStorageListOutput) StorageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.StorageId }).(pulumi.StringOutput)
}

// Name of the CBS to be queried.
func (o GetStoragesStorageListOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.StorageName }).(pulumi.StringOutput)
}

// Volume of CBS.
func (o GetStoragesStorageListOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetStoragesStorageList) int { return v.StorageSize }).(pulumi.IntOutput)
}

// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
func (o GetStoragesStorageListOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.StorageType }).(pulumi.StringOutput)
}

// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
func (o GetStoragesStorageListOutput) StorageUsage() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesStorageList) string { return v.StorageUsage }).(pulumi.StringOutput)
}

// The available tags within this CBS.
func (o GetStoragesStorageListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetStoragesStorageList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
func (o GetStoragesStorageListOutput) ThroughputPerformance() pulumi.IntOutput {
	return o.ApplyT(func(v GetStoragesStorageList) int { return v.ThroughputPerformance }).(pulumi.IntOutput)
}

type GetStoragesStorageListArrayOutput struct{ *pulumi.OutputState }

func (GetStoragesStorageListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStoragesStorageList)(nil)).Elem()
}

func (o GetStoragesStorageListArrayOutput) ToGetStoragesStorageListArrayOutput() GetStoragesStorageListArrayOutput {
	return o
}

func (o GetStoragesStorageListArrayOutput) ToGetStoragesStorageListArrayOutputWithContext(ctx context.Context) GetStoragesStorageListArrayOutput {
	return o
}

func (o GetStoragesStorageListArrayOutput) Index(i pulumi.IntInput) GetStoragesStorageListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetStoragesStorageList {
		return vs[0].([]GetStoragesStorageList)[vs[1].(int)]
	}).(GetStoragesStorageListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetSnapshotPoliciesSnapshotPolicyListInput)(nil)).Elem(), GetSnapshotPoliciesSnapshotPolicyListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSnapshotPoliciesSnapshotPolicyListArrayInput)(nil)).Elem(), GetSnapshotPoliciesSnapshotPolicyListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSnapshotsSnapshotListInput)(nil)).Elem(), GetSnapshotsSnapshotListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSnapshotsSnapshotListArrayInput)(nil)).Elem(), GetSnapshotsSnapshotListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetStoragesSetStorageListInput)(nil)).Elem(), GetStoragesSetStorageListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetStoragesSetStorageListArrayInput)(nil)).Elem(), GetStoragesSetStorageListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetStoragesStorageListInput)(nil)).Elem(), GetStoragesStorageListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetStoragesStorageListArrayInput)(nil)).Elem(), GetStoragesStorageListArray{})
	pulumi.RegisterOutputType(GetSnapshotPoliciesSnapshotPolicyListOutput{})
	pulumi.RegisterOutputType(GetSnapshotPoliciesSnapshotPolicyListArrayOutput{})
	pulumi.RegisterOutputType(GetSnapshotsSnapshotListOutput{})
	pulumi.RegisterOutputType(GetSnapshotsSnapshotListArrayOutput{})
	pulumi.RegisterOutputType(GetStoragesSetStorageListOutput{})
	pulumi.RegisterOutputType(GetStoragesSetStorageListArrayOutput{})
	pulumi.RegisterOutputType(GetStoragesStorageListOutput{})
	pulumi.RegisterOutputType(GetStoragesStorageListArrayOutput{})
}
