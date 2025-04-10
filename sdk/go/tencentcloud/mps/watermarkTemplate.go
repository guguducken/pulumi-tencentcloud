// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mps watermarkTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/base64"
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
// )
//
// func filebase64OrPanic(path string) pulumi.StringPtrInput {
// 	if fileData, err := ioutil.ReadFile(path); err == nil {
// 		return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
// 	} else {
// 		panic(err.Error())
// 	}
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mps.NewWatermarkTemplate(ctx, "watermarkTemplate", &Mps.WatermarkTemplateArgs{
// 			CoordinateOrigin: pulumi.String("TopLeft"),
// 			Type:             pulumi.String("image"),
// 			XPos:             pulumi.String(fmt.Sprintf("%v%v", "12", "%")),
// 			YPos:             pulumi.String(fmt.Sprintf("%v%v", "21", "%")),
// 			ImageTemplate: &mps.WatermarkTemplateImageTemplateArgs{
// 				Height:       pulumi.String("17px"),
// 				ImageContent: filebase64OrPanic("./logo.png"),
// 				RepeatType:   pulumi.String("repeat"),
// 				Width:        pulumi.String("12px"),
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
// mps watermark_template can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mps/watermarkTemplate:WatermarkTemplate watermark_template watermark_template_id
// ```
type WatermarkTemplate struct {
	pulumi.CustomResourceState

	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
	CoordinateOrigin pulumi.StringPtrOutput `pulumi:"coordinateOrigin"`
	// Image watermark template, only when Type is image, this field is required and valid.
	ImageTemplate WatermarkTemplateImageTemplatePtrOutput `pulumi:"imageTemplate"`
	// Watermark template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// SVG watermark template, only when Type is svg, this field is required and valid.
	SvgTemplate WatermarkTemplateSvgTemplatePtrOutput `pulumi:"svgTemplate"`
	// Text watermark template, only when Type is text, this field is required and valid.
	TextTemplate WatermarkTemplateTextTemplatePtrOutput `pulumi:"textTemplate"`
	// Watermark type, optional value:image, text, svg.
	Type pulumi.StringOutput `pulumi:"type"`
	// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
	XPos pulumi.StringPtrOutput `pulumi:"xPos"`
	// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
	YPos pulumi.StringPtrOutput `pulumi:"yPos"`
}

// NewWatermarkTemplate registers a new resource with the given unique name, arguments, and options.
func NewWatermarkTemplate(ctx *pulumi.Context,
	name string, args *WatermarkTemplateArgs, opts ...pulumi.ResourceOption) (*WatermarkTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource WatermarkTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/watermarkTemplate:WatermarkTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatermarkTemplate gets an existing WatermarkTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatermarkTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatermarkTemplateState, opts ...pulumi.ResourceOption) (*WatermarkTemplate, error) {
	var resource WatermarkTemplate
	err := ctx.ReadResource("tencentcloud:Mps/watermarkTemplate:WatermarkTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WatermarkTemplate resources.
type watermarkTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
	CoordinateOrigin *string `pulumi:"coordinateOrigin"`
	// Image watermark template, only when Type is image, this field is required and valid.
	ImageTemplate *WatermarkTemplateImageTemplate `pulumi:"imageTemplate"`
	// Watermark template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// SVG watermark template, only when Type is svg, this field is required and valid.
	SvgTemplate *WatermarkTemplateSvgTemplate `pulumi:"svgTemplate"`
	// Text watermark template, only when Type is text, this field is required and valid.
	TextTemplate *WatermarkTemplateTextTemplate `pulumi:"textTemplate"`
	// Watermark type, optional value:image, text, svg.
	Type *string `pulumi:"type"`
	// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
	XPos *string `pulumi:"xPos"`
	// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
	YPos *string `pulumi:"yPos"`
}

type WatermarkTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
	CoordinateOrigin pulumi.StringPtrInput
	// Image watermark template, only when Type is image, this field is required and valid.
	ImageTemplate WatermarkTemplateImageTemplatePtrInput
	// Watermark template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// SVG watermark template, only when Type is svg, this field is required and valid.
	SvgTemplate WatermarkTemplateSvgTemplatePtrInput
	// Text watermark template, only when Type is text, this field is required and valid.
	TextTemplate WatermarkTemplateTextTemplatePtrInput
	// Watermark type, optional value:image, text, svg.
	Type pulumi.StringPtrInput
	// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
	XPos pulumi.StringPtrInput
	// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
	YPos pulumi.StringPtrInput
}

func (WatermarkTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*watermarkTemplateState)(nil)).Elem()
}

type watermarkTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
	CoordinateOrigin *string `pulumi:"coordinateOrigin"`
	// Image watermark template, only when Type is image, this field is required and valid.
	ImageTemplate *WatermarkTemplateImageTemplate `pulumi:"imageTemplate"`
	// Watermark template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// SVG watermark template, only when Type is svg, this field is required and valid.
	SvgTemplate *WatermarkTemplateSvgTemplate `pulumi:"svgTemplate"`
	// Text watermark template, only when Type is text, this field is required and valid.
	TextTemplate *WatermarkTemplateTextTemplate `pulumi:"textTemplate"`
	// Watermark type, optional value:image, text, svg.
	Type string `pulumi:"type"`
	// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
	XPos *string `pulumi:"xPos"`
	// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
	YPos *string `pulumi:"yPos"`
}

// The set of arguments for constructing a WatermarkTemplate resource.
type WatermarkTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
	CoordinateOrigin pulumi.StringPtrInput
	// Image watermark template, only when Type is image, this field is required and valid.
	ImageTemplate WatermarkTemplateImageTemplatePtrInput
	// Watermark template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// SVG watermark template, only when Type is svg, this field is required and valid.
	SvgTemplate WatermarkTemplateSvgTemplatePtrInput
	// Text watermark template, only when Type is text, this field is required and valid.
	TextTemplate WatermarkTemplateTextTemplatePtrInput
	// Watermark type, optional value:image, text, svg.
	Type pulumi.StringInput
	// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
	XPos pulumi.StringPtrInput
	// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
	YPos pulumi.StringPtrInput
}

func (WatermarkTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watermarkTemplateArgs)(nil)).Elem()
}

type WatermarkTemplateInput interface {
	pulumi.Input

	ToWatermarkTemplateOutput() WatermarkTemplateOutput
	ToWatermarkTemplateOutputWithContext(ctx context.Context) WatermarkTemplateOutput
}

func (*WatermarkTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**WatermarkTemplate)(nil)).Elem()
}

func (i *WatermarkTemplate) ToWatermarkTemplateOutput() WatermarkTemplateOutput {
	return i.ToWatermarkTemplateOutputWithContext(context.Background())
}

func (i *WatermarkTemplate) ToWatermarkTemplateOutputWithContext(ctx context.Context) WatermarkTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatermarkTemplateOutput)
}

// WatermarkTemplateArrayInput is an input type that accepts WatermarkTemplateArray and WatermarkTemplateArrayOutput values.
// You can construct a concrete instance of `WatermarkTemplateArrayInput` via:
//
//          WatermarkTemplateArray{ WatermarkTemplateArgs{...} }
type WatermarkTemplateArrayInput interface {
	pulumi.Input

	ToWatermarkTemplateArrayOutput() WatermarkTemplateArrayOutput
	ToWatermarkTemplateArrayOutputWithContext(context.Context) WatermarkTemplateArrayOutput
}

type WatermarkTemplateArray []WatermarkTemplateInput

func (WatermarkTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WatermarkTemplate)(nil)).Elem()
}

func (i WatermarkTemplateArray) ToWatermarkTemplateArrayOutput() WatermarkTemplateArrayOutput {
	return i.ToWatermarkTemplateArrayOutputWithContext(context.Background())
}

func (i WatermarkTemplateArray) ToWatermarkTemplateArrayOutputWithContext(ctx context.Context) WatermarkTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatermarkTemplateArrayOutput)
}

// WatermarkTemplateMapInput is an input type that accepts WatermarkTemplateMap and WatermarkTemplateMapOutput values.
// You can construct a concrete instance of `WatermarkTemplateMapInput` via:
//
//          WatermarkTemplateMap{ "key": WatermarkTemplateArgs{...} }
type WatermarkTemplateMapInput interface {
	pulumi.Input

	ToWatermarkTemplateMapOutput() WatermarkTemplateMapOutput
	ToWatermarkTemplateMapOutputWithContext(context.Context) WatermarkTemplateMapOutput
}

type WatermarkTemplateMap map[string]WatermarkTemplateInput

func (WatermarkTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WatermarkTemplate)(nil)).Elem()
}

func (i WatermarkTemplateMap) ToWatermarkTemplateMapOutput() WatermarkTemplateMapOutput {
	return i.ToWatermarkTemplateMapOutputWithContext(context.Background())
}

func (i WatermarkTemplateMap) ToWatermarkTemplateMapOutputWithContext(ctx context.Context) WatermarkTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatermarkTemplateMapOutput)
}

type WatermarkTemplateOutput struct{ *pulumi.OutputState }

func (WatermarkTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WatermarkTemplate)(nil)).Elem()
}

func (o WatermarkTemplateOutput) ToWatermarkTemplateOutput() WatermarkTemplateOutput {
	return o
}

func (o WatermarkTemplateOutput) ToWatermarkTemplateOutputWithContext(ctx context.Context) WatermarkTemplateOutput {
	return o
}

// Template description information, length limit: 256 characters.
func (o WatermarkTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
func (o WatermarkTemplateOutput) CoordinateOrigin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) pulumi.StringPtrOutput { return v.CoordinateOrigin }).(pulumi.StringPtrOutput)
}

// Image watermark template, only when Type is image, this field is required and valid.
func (o WatermarkTemplateOutput) ImageTemplate() WatermarkTemplateImageTemplatePtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) WatermarkTemplateImageTemplatePtrOutput { return v.ImageTemplate }).(WatermarkTemplateImageTemplatePtrOutput)
}

// Watermark template name, length limit: 64 characters.
func (o WatermarkTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WatermarkTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SVG watermark template, only when Type is svg, this field is required and valid.
func (o WatermarkTemplateOutput) SvgTemplate() WatermarkTemplateSvgTemplatePtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) WatermarkTemplateSvgTemplatePtrOutput { return v.SvgTemplate }).(WatermarkTemplateSvgTemplatePtrOutput)
}

// Text watermark template, only when Type is text, this field is required and valid.
func (o WatermarkTemplateOutput) TextTemplate() WatermarkTemplateTextTemplatePtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) WatermarkTemplateTextTemplatePtrOutput { return v.TextTemplate }).(WatermarkTemplateTextTemplatePtrOutput)
}

// Watermark type, optional value:image, text, svg.
func (o WatermarkTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WatermarkTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
func (o WatermarkTemplateOutput) XPos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) pulumi.StringPtrOutput { return v.XPos }).(pulumi.StringPtrOutput)
}

// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
func (o WatermarkTemplateOutput) YPos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatermarkTemplate) pulumi.StringPtrOutput { return v.YPos }).(pulumi.StringPtrOutput)
}

type WatermarkTemplateArrayOutput struct{ *pulumi.OutputState }

func (WatermarkTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WatermarkTemplate)(nil)).Elem()
}

func (o WatermarkTemplateArrayOutput) ToWatermarkTemplateArrayOutput() WatermarkTemplateArrayOutput {
	return o
}

func (o WatermarkTemplateArrayOutput) ToWatermarkTemplateArrayOutputWithContext(ctx context.Context) WatermarkTemplateArrayOutput {
	return o
}

func (o WatermarkTemplateArrayOutput) Index(i pulumi.IntInput) WatermarkTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WatermarkTemplate {
		return vs[0].([]*WatermarkTemplate)[vs[1].(int)]
	}).(WatermarkTemplateOutput)
}

type WatermarkTemplateMapOutput struct{ *pulumi.OutputState }

func (WatermarkTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WatermarkTemplate)(nil)).Elem()
}

func (o WatermarkTemplateMapOutput) ToWatermarkTemplateMapOutput() WatermarkTemplateMapOutput {
	return o
}

func (o WatermarkTemplateMapOutput) ToWatermarkTemplateMapOutputWithContext(ctx context.Context) WatermarkTemplateMapOutput {
	return o
}

func (o WatermarkTemplateMapOutput) MapIndex(k pulumi.StringInput) WatermarkTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WatermarkTemplate {
		return vs[0].(map[string]*WatermarkTemplate)[vs[1].(string)]
	}).(WatermarkTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WatermarkTemplateInput)(nil)).Elem(), &WatermarkTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatermarkTemplateArrayInput)(nil)).Elem(), WatermarkTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatermarkTemplateMapInput)(nil)).Elem(), WatermarkTemplateMap{})
	pulumi.RegisterOutputType(WatermarkTemplateOutput{})
	pulumi.RegisterOutputType(WatermarkTemplateArrayOutput{})
	pulumi.RegisterOutputType(WatermarkTemplateMapOutput{})
}
