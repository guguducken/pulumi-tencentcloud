// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vpc sgSnapshotFileContent
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Vpc.GetSgSnapshotFileContent(ctx, &vpc.GetSgSnapshotFileContentArgs{
// 			SecurityGroupId:  "sg-ntrgm89v",
// 			SnapshotFileId:   "ssfile-017gepjxpr",
// 			SnapshotPolicyId: "sspolicy-ebjofe71",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSgSnapshotFileContent(ctx *pulumi.Context, args *GetSgSnapshotFileContentArgs, opts ...pulumi.InvokeOption) (*GetSgSnapshotFileContentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSgSnapshotFileContentResult
	err := ctx.Invoke("tencentcloud:Vpc/getSgSnapshotFileContent:getSgSnapshotFileContent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSgSnapshotFileContent.
type GetSgSnapshotFileContentArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Security group ID.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Snapshot file ID.
	SnapshotFileId string `pulumi:"snapshotFileId"`
	// Snapshot policy IDs.
	SnapshotPolicyId string `pulumi:"snapshotPolicyId"`
}

// A collection of values returned by getSgSnapshotFileContent.
type GetSgSnapshotFileContentResult struct {
	// Backup data.
	BackupDatas []GetSgSnapshotFileContentBackupData `pulumi:"backupDatas"`
	// Backup time.
	BackupTime string `pulumi:"backupTime"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Security group ID.
	InstanceId string `pulumi:"instanceId"`
	// Operator.
	Operator string `pulumi:"operator"`
	// Original data.
	OriginalDatas    []GetSgSnapshotFileContentOriginalData `pulumi:"originalDatas"`
	ResultOutputFile *string                                `pulumi:"resultOutputFile"`
	// The security group instance ID, such as `sg-ohuuioma`.
	SecurityGroupId  string `pulumi:"securityGroupId"`
	SnapshotFileId   string `pulumi:"snapshotFileId"`
	SnapshotPolicyId string `pulumi:"snapshotPolicyId"`
}

func GetSgSnapshotFileContentOutput(ctx *pulumi.Context, args GetSgSnapshotFileContentOutputArgs, opts ...pulumi.InvokeOption) GetSgSnapshotFileContentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSgSnapshotFileContentResult, error) {
			args := v.(GetSgSnapshotFileContentArgs)
			r, err := GetSgSnapshotFileContent(ctx, &args, opts...)
			var s GetSgSnapshotFileContentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSgSnapshotFileContentResultOutput)
}

// A collection of arguments for invoking getSgSnapshotFileContent.
type GetSgSnapshotFileContentOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Security group ID.
	SecurityGroupId pulumi.StringInput `pulumi:"securityGroupId"`
	// Snapshot file ID.
	SnapshotFileId pulumi.StringInput `pulumi:"snapshotFileId"`
	// Snapshot policy IDs.
	SnapshotPolicyId pulumi.StringInput `pulumi:"snapshotPolicyId"`
}

func (GetSgSnapshotFileContentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSgSnapshotFileContentArgs)(nil)).Elem()
}

// A collection of values returned by getSgSnapshotFileContent.
type GetSgSnapshotFileContentResultOutput struct{ *pulumi.OutputState }

func (GetSgSnapshotFileContentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSgSnapshotFileContentResult)(nil)).Elem()
}

func (o GetSgSnapshotFileContentResultOutput) ToGetSgSnapshotFileContentResultOutput() GetSgSnapshotFileContentResultOutput {
	return o
}

func (o GetSgSnapshotFileContentResultOutput) ToGetSgSnapshotFileContentResultOutputWithContext(ctx context.Context) GetSgSnapshotFileContentResultOutput {
	return o
}

// Backup data.
func (o GetSgSnapshotFileContentResultOutput) BackupDatas() GetSgSnapshotFileContentBackupDataArrayOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) []GetSgSnapshotFileContentBackupData { return v.BackupDatas }).(GetSgSnapshotFileContentBackupDataArrayOutput)
}

// Backup time.
func (o GetSgSnapshotFileContentResultOutput) BackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.BackupTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSgSnapshotFileContentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Security group ID.
func (o GetSgSnapshotFileContentResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Operator.
func (o GetSgSnapshotFileContentResultOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.Operator }).(pulumi.StringOutput)
}

// Original data.
func (o GetSgSnapshotFileContentResultOutput) OriginalDatas() GetSgSnapshotFileContentOriginalDataArrayOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) []GetSgSnapshotFileContentOriginalData { return v.OriginalDatas }).(GetSgSnapshotFileContentOriginalDataArrayOutput)
}

func (o GetSgSnapshotFileContentResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// The security group instance ID, such as `sg-ohuuioma`.
func (o GetSgSnapshotFileContentResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o GetSgSnapshotFileContentResultOutput) SnapshotFileId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.SnapshotFileId }).(pulumi.StringOutput)
}

func (o GetSgSnapshotFileContentResultOutput) SnapshotPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSgSnapshotFileContentResult) string { return v.SnapshotPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSgSnapshotFileContentResultOutput{})
}
