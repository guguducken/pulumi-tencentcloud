// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mps adaptiveDynamicStreamingTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mps.NewAdaptiveDynamicStreamingTemplate(ctx, "adaptiveDynamicStreamingTemplate", &Mps.AdaptiveDynamicStreamingTemplateArgs{
// 			Comment:                      pulumi.String("terrraform test"),
// 			DisableHigherVideoBitrate:    pulumi.Int(0),
// 			DisableHigherVideoResolution: pulumi.Int(1),
// 			Format:                       pulumi.String("HLS"),
// 			StreamInfos: mps.AdaptiveDynamicStreamingTemplateStreamInfoArray{
// 				&mps.AdaptiveDynamicStreamingTemplateStreamInfoArgs{
// 					Audio: &mps.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs{
// 						AudioChannel: pulumi.Int(1),
// 						Bitrate:      pulumi.Int(55),
// 						Codec:        pulumi.String("libmp3lame"),
// 						SampleRate:   pulumi.Int(32000),
// 					},
// 					RemoveAudio: pulumi.Int(0),
// 					RemoveVideo: pulumi.Int(0),
// 					Video: &mps.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs{
// 						Bitrate:            pulumi.Int(245),
// 						Codec:              pulumi.String("libx264"),
// 						FillType:           pulumi.String("black"),
// 						Fps:                pulumi.Int(30),
// 						Gop:                pulumi.Int(0),
// 						Height:             pulumi.Int(135),
// 						ResolutionAdaptive: pulumi.String("open"),
// 						Vcrf:               pulumi.Int(0),
// 						Width:              pulumi.Int(145),
// 					},
// 				},
// 				&mps.AdaptiveDynamicStreamingTemplateStreamInfoArgs{
// 					Audio: &mps.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs{
// 						AudioChannel: pulumi.Int(2),
// 						Bitrate:      pulumi.Int(60),
// 						Codec:        pulumi.String("libfdk_aac"),
// 						SampleRate:   pulumi.Int(32000),
// 					},
// 					RemoveAudio: pulumi.Int(0),
// 					RemoveVideo: pulumi.Int(0),
// 					Video: &mps.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs{
// 						Bitrate:            pulumi.Int(400),
// 						Codec:              pulumi.String("libx264"),
// 						FillType:           pulumi.String("black"),
// 						Fps:                pulumi.Int(40),
// 						Gop:                pulumi.Int(0),
// 						Height:             pulumi.Int(150),
// 						ResolutionAdaptive: pulumi.String("open"),
// 						Vcrf:               pulumi.Int(0),
// 						Width:              pulumi.Int(160),
// 					},
// 				},
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
// mps adaptive_dynamic_streaming_template can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mps/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate adaptive_dynamic_streaming_template adaptive_dynamic_streaming_template_id
// ```
type AdaptiveDynamicStreamingTemplate struct {
	pulumi.CustomResourceState

	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Whether to prohibit video from low bit rate to high bit rate, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoBitrate pulumi.IntPtrOutput `pulumi:"disableHigherVideoBitrate"`
	// Whether to prohibit the conversion of video resolution to high resolution, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoResolution pulumi.IntPtrOutput `pulumi:"disableHigherVideoResolution"`
	// Adaptive transcoding format, value range:HLS, MPEG-DASH.
	Format pulumi.StringOutput `pulumi:"format"`
	// Template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Convert adaptive code stream to output sub-stream parameter information, and output up to 10 sub-streams.Note: The frame rate of each sub-stream must be consistent; if not, the frame rate of the first sub-stream is used as the output frame rate.
	StreamInfos AdaptiveDynamicStreamingTemplateStreamInfoArrayOutput `pulumi:"streamInfos"`
}

// NewAdaptiveDynamicStreamingTemplate registers a new resource with the given unique name, arguments, and options.
func NewAdaptiveDynamicStreamingTemplate(ctx *pulumi.Context,
	name string, args *AdaptiveDynamicStreamingTemplateArgs, opts ...pulumi.ResourceOption) (*AdaptiveDynamicStreamingTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.StreamInfos == nil {
		return nil, errors.New("invalid value for required argument 'StreamInfos'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AdaptiveDynamicStreamingTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdaptiveDynamicStreamingTemplate gets an existing AdaptiveDynamicStreamingTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdaptiveDynamicStreamingTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdaptiveDynamicStreamingTemplateState, opts ...pulumi.ResourceOption) (*AdaptiveDynamicStreamingTemplate, error) {
	var resource AdaptiveDynamicStreamingTemplate
	err := ctx.ReadResource("tencentcloud:Mps/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdaptiveDynamicStreamingTemplate resources.
type adaptiveDynamicStreamingTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Whether to prohibit video from low bit rate to high bit rate, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoBitrate *int `pulumi:"disableHigherVideoBitrate"`
	// Whether to prohibit the conversion of video resolution to high resolution, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoResolution *int `pulumi:"disableHigherVideoResolution"`
	// Adaptive transcoding format, value range:HLS, MPEG-DASH.
	Format *string `pulumi:"format"`
	// Template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Convert adaptive code stream to output sub-stream parameter information, and output up to 10 sub-streams.Note: The frame rate of each sub-stream must be consistent; if not, the frame rate of the first sub-stream is used as the output frame rate.
	StreamInfos []AdaptiveDynamicStreamingTemplateStreamInfo `pulumi:"streamInfos"`
}

type AdaptiveDynamicStreamingTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Whether to prohibit video from low bit rate to high bit rate, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoBitrate pulumi.IntPtrInput
	// Whether to prohibit the conversion of video resolution to high resolution, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoResolution pulumi.IntPtrInput
	// Adaptive transcoding format, value range:HLS, MPEG-DASH.
	Format pulumi.StringPtrInput
	// Template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Convert adaptive code stream to output sub-stream parameter information, and output up to 10 sub-streams.Note: The frame rate of each sub-stream must be consistent; if not, the frame rate of the first sub-stream is used as the output frame rate.
	StreamInfos AdaptiveDynamicStreamingTemplateStreamInfoArrayInput
}

func (AdaptiveDynamicStreamingTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveDynamicStreamingTemplateState)(nil)).Elem()
}

type adaptiveDynamicStreamingTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Whether to prohibit video from low bit rate to high bit rate, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoBitrate *int `pulumi:"disableHigherVideoBitrate"`
	// Whether to prohibit the conversion of video resolution to high resolution, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoResolution *int `pulumi:"disableHigherVideoResolution"`
	// Adaptive transcoding format, value range:HLS, MPEG-DASH.
	Format string `pulumi:"format"`
	// Template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Convert adaptive code stream to output sub-stream parameter information, and output up to 10 sub-streams.Note: The frame rate of each sub-stream must be consistent; if not, the frame rate of the first sub-stream is used as the output frame rate.
	StreamInfos []AdaptiveDynamicStreamingTemplateStreamInfo `pulumi:"streamInfos"`
}

// The set of arguments for constructing a AdaptiveDynamicStreamingTemplate resource.
type AdaptiveDynamicStreamingTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Whether to prohibit video from low bit rate to high bit rate, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoBitrate pulumi.IntPtrInput
	// Whether to prohibit the conversion of video resolution to high resolution, value range:0: no.1: yes.Default value: 0.
	DisableHigherVideoResolution pulumi.IntPtrInput
	// Adaptive transcoding format, value range:HLS, MPEG-DASH.
	Format pulumi.StringInput
	// Template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Convert adaptive code stream to output sub-stream parameter information, and output up to 10 sub-streams.Note: The frame rate of each sub-stream must be consistent; if not, the frame rate of the first sub-stream is used as the output frame rate.
	StreamInfos AdaptiveDynamicStreamingTemplateStreamInfoArrayInput
}

func (AdaptiveDynamicStreamingTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveDynamicStreamingTemplateArgs)(nil)).Elem()
}

type AdaptiveDynamicStreamingTemplateInput interface {
	pulumi.Input

	ToAdaptiveDynamicStreamingTemplateOutput() AdaptiveDynamicStreamingTemplateOutput
	ToAdaptiveDynamicStreamingTemplateOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateOutput
}

func (*AdaptiveDynamicStreamingTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**AdaptiveDynamicStreamingTemplate)(nil)).Elem()
}

func (i *AdaptiveDynamicStreamingTemplate) ToAdaptiveDynamicStreamingTemplateOutput() AdaptiveDynamicStreamingTemplateOutput {
	return i.ToAdaptiveDynamicStreamingTemplateOutputWithContext(context.Background())
}

func (i *AdaptiveDynamicStreamingTemplate) ToAdaptiveDynamicStreamingTemplateOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveDynamicStreamingTemplateOutput)
}

// AdaptiveDynamicStreamingTemplateArrayInput is an input type that accepts AdaptiveDynamicStreamingTemplateArray and AdaptiveDynamicStreamingTemplateArrayOutput values.
// You can construct a concrete instance of `AdaptiveDynamicStreamingTemplateArrayInput` via:
//
//          AdaptiveDynamicStreamingTemplateArray{ AdaptiveDynamicStreamingTemplateArgs{...} }
type AdaptiveDynamicStreamingTemplateArrayInput interface {
	pulumi.Input

	ToAdaptiveDynamicStreamingTemplateArrayOutput() AdaptiveDynamicStreamingTemplateArrayOutput
	ToAdaptiveDynamicStreamingTemplateArrayOutputWithContext(context.Context) AdaptiveDynamicStreamingTemplateArrayOutput
}

type AdaptiveDynamicStreamingTemplateArray []AdaptiveDynamicStreamingTemplateInput

func (AdaptiveDynamicStreamingTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdaptiveDynamicStreamingTemplate)(nil)).Elem()
}

func (i AdaptiveDynamicStreamingTemplateArray) ToAdaptiveDynamicStreamingTemplateArrayOutput() AdaptiveDynamicStreamingTemplateArrayOutput {
	return i.ToAdaptiveDynamicStreamingTemplateArrayOutputWithContext(context.Background())
}

func (i AdaptiveDynamicStreamingTemplateArray) ToAdaptiveDynamicStreamingTemplateArrayOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveDynamicStreamingTemplateArrayOutput)
}

// AdaptiveDynamicStreamingTemplateMapInput is an input type that accepts AdaptiveDynamicStreamingTemplateMap and AdaptiveDynamicStreamingTemplateMapOutput values.
// You can construct a concrete instance of `AdaptiveDynamicStreamingTemplateMapInput` via:
//
//          AdaptiveDynamicStreamingTemplateMap{ "key": AdaptiveDynamicStreamingTemplateArgs{...} }
type AdaptiveDynamicStreamingTemplateMapInput interface {
	pulumi.Input

	ToAdaptiveDynamicStreamingTemplateMapOutput() AdaptiveDynamicStreamingTemplateMapOutput
	ToAdaptiveDynamicStreamingTemplateMapOutputWithContext(context.Context) AdaptiveDynamicStreamingTemplateMapOutput
}

type AdaptiveDynamicStreamingTemplateMap map[string]AdaptiveDynamicStreamingTemplateInput

func (AdaptiveDynamicStreamingTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdaptiveDynamicStreamingTemplate)(nil)).Elem()
}

func (i AdaptiveDynamicStreamingTemplateMap) ToAdaptiveDynamicStreamingTemplateMapOutput() AdaptiveDynamicStreamingTemplateMapOutput {
	return i.ToAdaptiveDynamicStreamingTemplateMapOutputWithContext(context.Background())
}

func (i AdaptiveDynamicStreamingTemplateMap) ToAdaptiveDynamicStreamingTemplateMapOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveDynamicStreamingTemplateMapOutput)
}

type AdaptiveDynamicStreamingTemplateOutput struct{ *pulumi.OutputState }

func (AdaptiveDynamicStreamingTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdaptiveDynamicStreamingTemplate)(nil)).Elem()
}

func (o AdaptiveDynamicStreamingTemplateOutput) ToAdaptiveDynamicStreamingTemplateOutput() AdaptiveDynamicStreamingTemplateOutput {
	return o
}

func (o AdaptiveDynamicStreamingTemplateOutput) ToAdaptiveDynamicStreamingTemplateOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateOutput {
	return o
}

// Template description information, length limit: 256 characters.
func (o AdaptiveDynamicStreamingTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdaptiveDynamicStreamingTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Whether to prohibit video from low bit rate to high bit rate, value range:0: no.1: yes.Default value: 0.
func (o AdaptiveDynamicStreamingTemplateOutput) DisableHigherVideoBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AdaptiveDynamicStreamingTemplate) pulumi.IntPtrOutput { return v.DisableHigherVideoBitrate }).(pulumi.IntPtrOutput)
}

// Whether to prohibit the conversion of video resolution to high resolution, value range:0: no.1: yes.Default value: 0.
func (o AdaptiveDynamicStreamingTemplateOutput) DisableHigherVideoResolution() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AdaptiveDynamicStreamingTemplate) pulumi.IntPtrOutput { return v.DisableHigherVideoResolution }).(pulumi.IntPtrOutput)
}

// Adaptive transcoding format, value range:HLS, MPEG-DASH.
func (o AdaptiveDynamicStreamingTemplateOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveDynamicStreamingTemplate) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Template name, length limit: 64 characters.
func (o AdaptiveDynamicStreamingTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveDynamicStreamingTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Convert adaptive code stream to output sub-stream parameter information, and output up to 10 sub-streams.Note: The frame rate of each sub-stream must be consistent; if not, the frame rate of the first sub-stream is used as the output frame rate.
func (o AdaptiveDynamicStreamingTemplateOutput) StreamInfos() AdaptiveDynamicStreamingTemplateStreamInfoArrayOutput {
	return o.ApplyT(func(v *AdaptiveDynamicStreamingTemplate) AdaptiveDynamicStreamingTemplateStreamInfoArrayOutput {
		return v.StreamInfos
	}).(AdaptiveDynamicStreamingTemplateStreamInfoArrayOutput)
}

type AdaptiveDynamicStreamingTemplateArrayOutput struct{ *pulumi.OutputState }

func (AdaptiveDynamicStreamingTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdaptiveDynamicStreamingTemplate)(nil)).Elem()
}

func (o AdaptiveDynamicStreamingTemplateArrayOutput) ToAdaptiveDynamicStreamingTemplateArrayOutput() AdaptiveDynamicStreamingTemplateArrayOutput {
	return o
}

func (o AdaptiveDynamicStreamingTemplateArrayOutput) ToAdaptiveDynamicStreamingTemplateArrayOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateArrayOutput {
	return o
}

func (o AdaptiveDynamicStreamingTemplateArrayOutput) Index(i pulumi.IntInput) AdaptiveDynamicStreamingTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdaptiveDynamicStreamingTemplate {
		return vs[0].([]*AdaptiveDynamicStreamingTemplate)[vs[1].(int)]
	}).(AdaptiveDynamicStreamingTemplateOutput)
}

type AdaptiveDynamicStreamingTemplateMapOutput struct{ *pulumi.OutputState }

func (AdaptiveDynamicStreamingTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdaptiveDynamicStreamingTemplate)(nil)).Elem()
}

func (o AdaptiveDynamicStreamingTemplateMapOutput) ToAdaptiveDynamicStreamingTemplateMapOutput() AdaptiveDynamicStreamingTemplateMapOutput {
	return o
}

func (o AdaptiveDynamicStreamingTemplateMapOutput) ToAdaptiveDynamicStreamingTemplateMapOutputWithContext(ctx context.Context) AdaptiveDynamicStreamingTemplateMapOutput {
	return o
}

func (o AdaptiveDynamicStreamingTemplateMapOutput) MapIndex(k pulumi.StringInput) AdaptiveDynamicStreamingTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdaptiveDynamicStreamingTemplate {
		return vs[0].(map[string]*AdaptiveDynamicStreamingTemplate)[vs[1].(string)]
	}).(AdaptiveDynamicStreamingTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdaptiveDynamicStreamingTemplateInput)(nil)).Elem(), &AdaptiveDynamicStreamingTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdaptiveDynamicStreamingTemplateArrayInput)(nil)).Elem(), AdaptiveDynamicStreamingTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdaptiveDynamicStreamingTemplateMapInput)(nil)).Elem(), AdaptiveDynamicStreamingTemplateMap{})
	pulumi.RegisterOutputType(AdaptiveDynamicStreamingTemplateOutput{})
	pulumi.RegisterOutputType(AdaptiveDynamicStreamingTemplateArrayOutput{})
	pulumi.RegisterOutputType(AdaptiveDynamicStreamingTemplateMapOutput{})
}
