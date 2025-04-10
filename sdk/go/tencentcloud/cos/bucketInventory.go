// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cos bucketInventory
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cos.NewBucketInventory(ctx, "bucketInventory", &Cos.BucketInventoryArgs{
// 			Bucket: pulumi.String("keep-test-xxxxxx"),
// 			Destination: &cos.BucketInventoryDestinationArgs{
// 				AccountId: pulumi.String(""),
// 				Bucket:    pulumi.String("qcs::cos:ap-guangzhou::keep-test-xxxxxx"),
// 				Format:    pulumi.String("CSV"),
// 				Prefix:    pulumi.String("cos_bucket_inventory"),
// 			},
// 			Filter: &cos.BucketInventoryFilterArgs{
// 				Period: &cos.BucketInventoryFilterPeriodArgs{
// 					StartTime: pulumi.String("1687276800"),
// 				},
// 			},
// 			IncludedObjectVersions: pulumi.String("Current"),
// 			IsEnabled:              pulumi.String("true"),
// 			OptionalFields: &cos.BucketInventoryOptionalFieldsArgs{
// 				Fields: pulumi.StringArray{
// 					pulumi.String("Size"),
// 					pulumi.String("ETag"),
// 				},
// 			},
// 			Schedule: &cos.BucketInventoryScheduleArgs{
// 				Frequency: pulumi.String("Weekly"),
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
// cos bucket_inventory can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cos/bucketInventory:BucketInventory bucket_inventory bucket_inventory_id
// ```
type BucketInventory struct {
	pulumi.CustomResourceState

	// Bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Information about the inventory result destination.
	Destination BucketInventoryDestinationOutput `pulumi:"destination"`
	// Filters objects prefixed with the specified value to analyze.
	Filter BucketInventoryFilterPtrOutput `pulumi:"filter"`
	// Whether to include object versions in the inventory. All or No.
	IncludedObjectVersions pulumi.StringOutput `pulumi:"includedObjectVersions"`
	// Whether to enable the inventory. true or false.
	IsEnabled pulumi.StringOutput `pulumi:"isEnabled"`
	// Inventory Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Analysis items to include in the inventory result	.
	OptionalFields BucketInventoryOptionalFieldsPtrOutput `pulumi:"optionalFields"`
	// Inventory job cycle.
	Schedule BucketInventoryScheduleOutput `pulumi:"schedule"`
}

// NewBucketInventory registers a new resource with the given unique name, arguments, and options.
func NewBucketInventory(ctx *pulumi.Context,
	name string, args *BucketInventoryArgs, opts ...pulumi.ResourceOption) (*BucketInventory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.IncludedObjectVersions == nil {
		return nil, errors.New("invalid value for required argument 'IncludedObjectVersions'")
	}
	if args.IsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsEnabled'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BucketInventory
	err := ctx.RegisterResource("tencentcloud:Cos/bucketInventory:BucketInventory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketInventory gets an existing BucketInventory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketInventory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketInventoryState, opts ...pulumi.ResourceOption) (*BucketInventory, error) {
	var resource BucketInventory
	err := ctx.ReadResource("tencentcloud:Cos/bucketInventory:BucketInventory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketInventory resources.
type bucketInventoryState struct {
	// Bucket name.
	Bucket *string `pulumi:"bucket"`
	// Information about the inventory result destination.
	Destination *BucketInventoryDestination `pulumi:"destination"`
	// Filters objects prefixed with the specified value to analyze.
	Filter *BucketInventoryFilter `pulumi:"filter"`
	// Whether to include object versions in the inventory. All or No.
	IncludedObjectVersions *string `pulumi:"includedObjectVersions"`
	// Whether to enable the inventory. true or false.
	IsEnabled *string `pulumi:"isEnabled"`
	// Inventory Name.
	Name *string `pulumi:"name"`
	// Analysis items to include in the inventory result	.
	OptionalFields *BucketInventoryOptionalFields `pulumi:"optionalFields"`
	// Inventory job cycle.
	Schedule *BucketInventorySchedule `pulumi:"schedule"`
}

type BucketInventoryState struct {
	// Bucket name.
	Bucket pulumi.StringPtrInput
	// Information about the inventory result destination.
	Destination BucketInventoryDestinationPtrInput
	// Filters objects prefixed with the specified value to analyze.
	Filter BucketInventoryFilterPtrInput
	// Whether to include object versions in the inventory. All or No.
	IncludedObjectVersions pulumi.StringPtrInput
	// Whether to enable the inventory. true or false.
	IsEnabled pulumi.StringPtrInput
	// Inventory Name.
	Name pulumi.StringPtrInput
	// Analysis items to include in the inventory result	.
	OptionalFields BucketInventoryOptionalFieldsPtrInput
	// Inventory job cycle.
	Schedule BucketInventorySchedulePtrInput
}

func (BucketInventoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketInventoryState)(nil)).Elem()
}

type bucketInventoryArgs struct {
	// Bucket name.
	Bucket string `pulumi:"bucket"`
	// Information about the inventory result destination.
	Destination BucketInventoryDestination `pulumi:"destination"`
	// Filters objects prefixed with the specified value to analyze.
	Filter *BucketInventoryFilter `pulumi:"filter"`
	// Whether to include object versions in the inventory. All or No.
	IncludedObjectVersions string `pulumi:"includedObjectVersions"`
	// Whether to enable the inventory. true or false.
	IsEnabled string `pulumi:"isEnabled"`
	// Inventory Name.
	Name *string `pulumi:"name"`
	// Analysis items to include in the inventory result	.
	OptionalFields *BucketInventoryOptionalFields `pulumi:"optionalFields"`
	// Inventory job cycle.
	Schedule BucketInventorySchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a BucketInventory resource.
type BucketInventoryArgs struct {
	// Bucket name.
	Bucket pulumi.StringInput
	// Information about the inventory result destination.
	Destination BucketInventoryDestinationInput
	// Filters objects prefixed with the specified value to analyze.
	Filter BucketInventoryFilterPtrInput
	// Whether to include object versions in the inventory. All or No.
	IncludedObjectVersions pulumi.StringInput
	// Whether to enable the inventory. true or false.
	IsEnabled pulumi.StringInput
	// Inventory Name.
	Name pulumi.StringPtrInput
	// Analysis items to include in the inventory result	.
	OptionalFields BucketInventoryOptionalFieldsPtrInput
	// Inventory job cycle.
	Schedule BucketInventoryScheduleInput
}

func (BucketInventoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketInventoryArgs)(nil)).Elem()
}

type BucketInventoryInput interface {
	pulumi.Input

	ToBucketInventoryOutput() BucketInventoryOutput
	ToBucketInventoryOutputWithContext(ctx context.Context) BucketInventoryOutput
}

func (*BucketInventory) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketInventory)(nil)).Elem()
}

func (i *BucketInventory) ToBucketInventoryOutput() BucketInventoryOutput {
	return i.ToBucketInventoryOutputWithContext(context.Background())
}

func (i *BucketInventory) ToBucketInventoryOutputWithContext(ctx context.Context) BucketInventoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketInventoryOutput)
}

// BucketInventoryArrayInput is an input type that accepts BucketInventoryArray and BucketInventoryArrayOutput values.
// You can construct a concrete instance of `BucketInventoryArrayInput` via:
//
//          BucketInventoryArray{ BucketInventoryArgs{...} }
type BucketInventoryArrayInput interface {
	pulumi.Input

	ToBucketInventoryArrayOutput() BucketInventoryArrayOutput
	ToBucketInventoryArrayOutputWithContext(context.Context) BucketInventoryArrayOutput
}

type BucketInventoryArray []BucketInventoryInput

func (BucketInventoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketInventory)(nil)).Elem()
}

func (i BucketInventoryArray) ToBucketInventoryArrayOutput() BucketInventoryArrayOutput {
	return i.ToBucketInventoryArrayOutputWithContext(context.Background())
}

func (i BucketInventoryArray) ToBucketInventoryArrayOutputWithContext(ctx context.Context) BucketInventoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketInventoryArrayOutput)
}

// BucketInventoryMapInput is an input type that accepts BucketInventoryMap and BucketInventoryMapOutput values.
// You can construct a concrete instance of `BucketInventoryMapInput` via:
//
//          BucketInventoryMap{ "key": BucketInventoryArgs{...} }
type BucketInventoryMapInput interface {
	pulumi.Input

	ToBucketInventoryMapOutput() BucketInventoryMapOutput
	ToBucketInventoryMapOutputWithContext(context.Context) BucketInventoryMapOutput
}

type BucketInventoryMap map[string]BucketInventoryInput

func (BucketInventoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketInventory)(nil)).Elem()
}

func (i BucketInventoryMap) ToBucketInventoryMapOutput() BucketInventoryMapOutput {
	return i.ToBucketInventoryMapOutputWithContext(context.Background())
}

func (i BucketInventoryMap) ToBucketInventoryMapOutputWithContext(ctx context.Context) BucketInventoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketInventoryMapOutput)
}

type BucketInventoryOutput struct{ *pulumi.OutputState }

func (BucketInventoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketInventory)(nil)).Elem()
}

func (o BucketInventoryOutput) ToBucketInventoryOutput() BucketInventoryOutput {
	return o
}

func (o BucketInventoryOutput) ToBucketInventoryOutputWithContext(ctx context.Context) BucketInventoryOutput {
	return o
}

// Bucket name.
func (o BucketInventoryOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Information about the inventory result destination.
func (o BucketInventoryOutput) Destination() BucketInventoryDestinationOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryDestinationOutput { return v.Destination }).(BucketInventoryDestinationOutput)
}

// Filters objects prefixed with the specified value to analyze.
func (o BucketInventoryOutput) Filter() BucketInventoryFilterPtrOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryFilterPtrOutput { return v.Filter }).(BucketInventoryFilterPtrOutput)
}

// Whether to include object versions in the inventory. All or No.
func (o BucketInventoryOutput) IncludedObjectVersions() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.IncludedObjectVersions }).(pulumi.StringOutput)
}

// Whether to enable the inventory. true or false.
func (o BucketInventoryOutput) IsEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.IsEnabled }).(pulumi.StringOutput)
}

// Inventory Name.
func (o BucketInventoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Analysis items to include in the inventory result	.
func (o BucketInventoryOutput) OptionalFields() BucketInventoryOptionalFieldsPtrOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryOptionalFieldsPtrOutput { return v.OptionalFields }).(BucketInventoryOptionalFieldsPtrOutput)
}

// Inventory job cycle.
func (o BucketInventoryOutput) Schedule() BucketInventoryScheduleOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryScheduleOutput { return v.Schedule }).(BucketInventoryScheduleOutput)
}

type BucketInventoryArrayOutput struct{ *pulumi.OutputState }

func (BucketInventoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketInventory)(nil)).Elem()
}

func (o BucketInventoryArrayOutput) ToBucketInventoryArrayOutput() BucketInventoryArrayOutput {
	return o
}

func (o BucketInventoryArrayOutput) ToBucketInventoryArrayOutputWithContext(ctx context.Context) BucketInventoryArrayOutput {
	return o
}

func (o BucketInventoryArrayOutput) Index(i pulumi.IntInput) BucketInventoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketInventory {
		return vs[0].([]*BucketInventory)[vs[1].(int)]
	}).(BucketInventoryOutput)
}

type BucketInventoryMapOutput struct{ *pulumi.OutputState }

func (BucketInventoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketInventory)(nil)).Elem()
}

func (o BucketInventoryMapOutput) ToBucketInventoryMapOutput() BucketInventoryMapOutput {
	return o
}

func (o BucketInventoryMapOutput) ToBucketInventoryMapOutputWithContext(ctx context.Context) BucketInventoryMapOutput {
	return o
}

func (o BucketInventoryMapOutput) MapIndex(k pulumi.StringInput) BucketInventoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketInventory {
		return vs[0].(map[string]*BucketInventory)[vs[1].(string)]
	}).(BucketInventoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInventoryInput)(nil)).Elem(), &BucketInventory{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInventoryArrayInput)(nil)).Elem(), BucketInventoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInventoryMapInput)(nil)).Elem(), BucketInventoryMap{})
	pulumi.RegisterOutputType(BucketInventoryOutput{})
	pulumi.RegisterOutputType(BucketInventoryArrayOutput{})
	pulumi.RegisterOutputType(BucketInventoryMapOutput{})
}
