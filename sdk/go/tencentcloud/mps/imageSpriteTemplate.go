// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mps imageSpriteTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mps.NewImageSpriteTemplate(ctx, "imageSpriteTemplate", &Mps.ImageSpriteTemplateArgs{
// 			ColumnCount:        pulumi.Int(10),
// 			FillType:           pulumi.String("stretch"),
// 			Format:             pulumi.String("jpg"),
// 			Height:             pulumi.Int(143),
// 			ResolutionAdaptive: pulumi.String("open"),
// 			RowCount:           pulumi.Int(10),
// 			SampleInterval:     pulumi.Int(10),
// 			SampleType:         pulumi.String("Time"),
// 			Width:              pulumi.Int(182),
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
// mps image_sprite_template can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate image_sprite_template image_sprite_template_id
// ```
type ImageSpriteTemplate struct {
	pulumi.CustomResourceState

	// The number of columns in the small image in the sprite.
	ColumnCount pulumi.IntOutput `pulumi:"columnCount"`
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
	FillType pulumi.StringPtrOutput `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// Image sprite template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrOutput `pulumi:"resolutionAdaptive"`
	// The number of rows in the small image in the sprite.
	RowCount pulumi.IntOutput `pulumi:"rowCount"`
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval pulumi.IntOutput `pulumi:"sampleInterval"`
	// Sampling type, optional value:Percent/Time.
	SampleType pulumi.StringOutput `pulumi:"sampleType"`
	// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewImageSpriteTemplate registers a new resource with the given unique name, arguments, and options.
func NewImageSpriteTemplate(ctx *pulumi.Context,
	name string, args *ImageSpriteTemplateArgs, opts ...pulumi.ResourceOption) (*ImageSpriteTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnCount == nil {
		return nil, errors.New("invalid value for required argument 'ColumnCount'")
	}
	if args.RowCount == nil {
		return nil, errors.New("invalid value for required argument 'RowCount'")
	}
	if args.SampleInterval == nil {
		return nil, errors.New("invalid value for required argument 'SampleInterval'")
	}
	if args.SampleType == nil {
		return nil, errors.New("invalid value for required argument 'SampleType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ImageSpriteTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageSpriteTemplate gets an existing ImageSpriteTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageSpriteTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageSpriteTemplateState, opts ...pulumi.ResourceOption) (*ImageSpriteTemplate, error) {
	var resource ImageSpriteTemplate
	err := ctx.ReadResource("tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageSpriteTemplate resources.
type imageSpriteTemplateState struct {
	// The number of columns in the small image in the sprite.
	ColumnCount *int `pulumi:"columnCount"`
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format *string `pulumi:"format"`
	// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height *int `pulumi:"height"`
	// Image sprite template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// The number of rows in the small image in the sprite.
	RowCount *int `pulumi:"rowCount"`
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval *int `pulumi:"sampleInterval"`
	// Sampling type, optional value:Percent/Time.
	SampleType *string `pulumi:"sampleType"`
	// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width *int `pulumi:"width"`
}

type ImageSpriteTemplateState struct {
	// The number of columns in the small image in the sprite.
	ColumnCount pulumi.IntPtrInput
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrInput
	// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrInput
	// Image sprite template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// The number of rows in the small image in the sprite.
	RowCount pulumi.IntPtrInput
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval pulumi.IntPtrInput
	// Sampling type, optional value:Percent/Time.
	SampleType pulumi.StringPtrInput
	// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrInput
}

func (ImageSpriteTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageSpriteTemplateState)(nil)).Elem()
}

type imageSpriteTemplateArgs struct {
	// The number of columns in the small image in the sprite.
	ColumnCount int `pulumi:"columnCount"`
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format *string `pulumi:"format"`
	// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height *int `pulumi:"height"`
	// Image sprite template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// The number of rows in the small image in the sprite.
	RowCount int `pulumi:"rowCount"`
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval int `pulumi:"sampleInterval"`
	// Sampling type, optional value:Percent/Time.
	SampleType string `pulumi:"sampleType"`
	// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a ImageSpriteTemplate resource.
type ImageSpriteTemplateArgs struct {
	// The number of columns in the small image in the sprite.
	ColumnCount pulumi.IntInput
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrInput
	// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrInput
	// Image sprite template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// The number of rows in the small image in the sprite.
	RowCount pulumi.IntInput
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval pulumi.IntInput
	// Sampling type, optional value:Percent/Time.
	SampleType pulumi.StringInput
	// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrInput
}

func (ImageSpriteTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageSpriteTemplateArgs)(nil)).Elem()
}

type ImageSpriteTemplateInput interface {
	pulumi.Input

	ToImageSpriteTemplateOutput() ImageSpriteTemplateOutput
	ToImageSpriteTemplateOutputWithContext(ctx context.Context) ImageSpriteTemplateOutput
}

func (*ImageSpriteTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageSpriteTemplate)(nil)).Elem()
}

func (i *ImageSpriteTemplate) ToImageSpriteTemplateOutput() ImageSpriteTemplateOutput {
	return i.ToImageSpriteTemplateOutputWithContext(context.Background())
}

func (i *ImageSpriteTemplate) ToImageSpriteTemplateOutputWithContext(ctx context.Context) ImageSpriteTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSpriteTemplateOutput)
}

// ImageSpriteTemplateArrayInput is an input type that accepts ImageSpriteTemplateArray and ImageSpriteTemplateArrayOutput values.
// You can construct a concrete instance of `ImageSpriteTemplateArrayInput` via:
//
//          ImageSpriteTemplateArray{ ImageSpriteTemplateArgs{...} }
type ImageSpriteTemplateArrayInput interface {
	pulumi.Input

	ToImageSpriteTemplateArrayOutput() ImageSpriteTemplateArrayOutput
	ToImageSpriteTemplateArrayOutputWithContext(context.Context) ImageSpriteTemplateArrayOutput
}

type ImageSpriteTemplateArray []ImageSpriteTemplateInput

func (ImageSpriteTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageSpriteTemplate)(nil)).Elem()
}

func (i ImageSpriteTemplateArray) ToImageSpriteTemplateArrayOutput() ImageSpriteTemplateArrayOutput {
	return i.ToImageSpriteTemplateArrayOutputWithContext(context.Background())
}

func (i ImageSpriteTemplateArray) ToImageSpriteTemplateArrayOutputWithContext(ctx context.Context) ImageSpriteTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSpriteTemplateArrayOutput)
}

// ImageSpriteTemplateMapInput is an input type that accepts ImageSpriteTemplateMap and ImageSpriteTemplateMapOutput values.
// You can construct a concrete instance of `ImageSpriteTemplateMapInput` via:
//
//          ImageSpriteTemplateMap{ "key": ImageSpriteTemplateArgs{...} }
type ImageSpriteTemplateMapInput interface {
	pulumi.Input

	ToImageSpriteTemplateMapOutput() ImageSpriteTemplateMapOutput
	ToImageSpriteTemplateMapOutputWithContext(context.Context) ImageSpriteTemplateMapOutput
}

type ImageSpriteTemplateMap map[string]ImageSpriteTemplateInput

func (ImageSpriteTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageSpriteTemplate)(nil)).Elem()
}

func (i ImageSpriteTemplateMap) ToImageSpriteTemplateMapOutput() ImageSpriteTemplateMapOutput {
	return i.ToImageSpriteTemplateMapOutputWithContext(context.Background())
}

func (i ImageSpriteTemplateMap) ToImageSpriteTemplateMapOutputWithContext(ctx context.Context) ImageSpriteTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSpriteTemplateMapOutput)
}

type ImageSpriteTemplateOutput struct{ *pulumi.OutputState }

func (ImageSpriteTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageSpriteTemplate)(nil)).Elem()
}

func (o ImageSpriteTemplateOutput) ToImageSpriteTemplateOutput() ImageSpriteTemplateOutput {
	return o
}

func (o ImageSpriteTemplateOutput) ToImageSpriteTemplateOutputWithContext(ctx context.Context) ImageSpriteTemplateOutput {
	return o
}

// The number of columns in the small image in the sprite.
func (o ImageSpriteTemplateOutput) ColumnCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.IntOutput { return v.ColumnCount }).(pulumi.IntOutput)
}

// Template description information, length limit: 256 characters.
func (o ImageSpriteTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
func (o ImageSpriteTemplateOutput) FillType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.StringPtrOutput { return v.FillType }).(pulumi.StringPtrOutput)
}

// Image format, the value can be jpg, png, webp. Default is jpg.
func (o ImageSpriteTemplateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
func (o ImageSpriteTemplateOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// Image sprite template name, length limit: 64 characters.
func (o ImageSpriteTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
func (o ImageSpriteTemplateOutput) ResolutionAdaptive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.StringPtrOutput { return v.ResolutionAdaptive }).(pulumi.StringPtrOutput)
}

// The number of rows in the small image in the sprite.
func (o ImageSpriteTemplateOutput) RowCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.IntOutput { return v.RowCount }).(pulumi.IntOutput)
}

// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
func (o ImageSpriteTemplateOutput) SampleInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.IntOutput { return v.SampleInterval }).(pulumi.IntOutput)
}

// Sampling type, optional value:Percent/Time.
func (o ImageSpriteTemplateOutput) SampleType() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.StringOutput { return v.SampleType }).(pulumi.StringOutput)
}

// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
func (o ImageSpriteTemplateOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageSpriteTemplate) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type ImageSpriteTemplateArrayOutput struct{ *pulumi.OutputState }

func (ImageSpriteTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageSpriteTemplate)(nil)).Elem()
}

func (o ImageSpriteTemplateArrayOutput) ToImageSpriteTemplateArrayOutput() ImageSpriteTemplateArrayOutput {
	return o
}

func (o ImageSpriteTemplateArrayOutput) ToImageSpriteTemplateArrayOutputWithContext(ctx context.Context) ImageSpriteTemplateArrayOutput {
	return o
}

func (o ImageSpriteTemplateArrayOutput) Index(i pulumi.IntInput) ImageSpriteTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageSpriteTemplate {
		return vs[0].([]*ImageSpriteTemplate)[vs[1].(int)]
	}).(ImageSpriteTemplateOutput)
}

type ImageSpriteTemplateMapOutput struct{ *pulumi.OutputState }

func (ImageSpriteTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageSpriteTemplate)(nil)).Elem()
}

func (o ImageSpriteTemplateMapOutput) ToImageSpriteTemplateMapOutput() ImageSpriteTemplateMapOutput {
	return o
}

func (o ImageSpriteTemplateMapOutput) ToImageSpriteTemplateMapOutputWithContext(ctx context.Context) ImageSpriteTemplateMapOutput {
	return o
}

func (o ImageSpriteTemplateMapOutput) MapIndex(k pulumi.StringInput) ImageSpriteTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageSpriteTemplate {
		return vs[0].(map[string]*ImageSpriteTemplate)[vs[1].(string)]
	}).(ImageSpriteTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageSpriteTemplateInput)(nil)).Elem(), &ImageSpriteTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageSpriteTemplateArrayInput)(nil)).Elem(), ImageSpriteTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageSpriteTemplateMapInput)(nil)).Elem(), ImageSpriteTemplateMap{})
	pulumi.RegisterOutputType(ImageSpriteTemplateOutput{})
	pulumi.RegisterOutputType(ImageSpriteTemplateArrayOutput{})
	pulumi.RegisterOutputType(ImageSpriteTemplateMapOutput{})
}
