// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a postgresql parameterTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Postgresql.NewParameterTemplate(ctx, "parameterTemplate", &Postgresql.ParameterTemplateArgs{
// 			DbEngine:       pulumi.String("postgresql"),
// 			DbMajorVersion: pulumi.String("13"),
// 			DeleteParamSets: pulumi.StringArray{
// 				pulumi.String("lc_time"),
// 			},
// 			ModifyParamEntrySets: postgresql.ParameterTemplateModifyParamEntrySetArray{
// 				&postgresql.ParameterTemplateModifyParamEntrySetArgs{
// 					ExpectedValue: pulumi.String("UTC"),
// 					Name:          pulumi.String("timezone"),
// 				},
// 				&postgresql.ParameterTemplateModifyParamEntrySetArgs{
// 					ExpectedValue: pulumi.String("123"),
// 					Name:          pulumi.String("lock_timeout"),
// 				},
// 			},
// 			TemplateDescription: pulumi.String("For_tf_test"),
// 			TemplateName:        pulumi.String("your_temp_name"),
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
// postgresql parameter_template can be imported using the id, e.g. Notice`modify_param_entry_set` and `delete_param_set` do not support import.
//
// ```sh
//  $ pulumi import tencentcloud:Postgresql/parameterTemplate:ParameterTemplate parameter_template parameter_template_id
// ```
type ParameterTemplate struct {
	pulumi.CustomResourceState

	// Database engine, such as postgresql, mssql_compatible.
	DbEngine pulumi.StringOutput `pulumi:"dbEngine"`
	// The major database version number, such as 11, 12, 13.
	DbMajorVersion pulumi.StringOutput `pulumi:"dbMajorVersion"`
	// The set of parameters that need to be deleted.
	DeleteParamSets pulumi.StringArrayOutput `pulumi:"deleteParamSets"`
	// The set of parameters that need to be modified or added. Note: the same parameter cannot appear in the set of modifying and adding and deleting at the same time.
	ModifyParamEntrySets ParameterTemplateModifyParamEntrySetArrayOutput `pulumi:"modifyParamEntrySets"`
	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateDescription pulumi.StringPtrOutput `pulumi:"templateDescription"`
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewParameterTemplate registers a new resource with the given unique name, arguments, and options.
func NewParameterTemplate(ctx *pulumi.Context,
	name string, args *ParameterTemplateArgs, opts ...pulumi.ResourceOption) (*ParameterTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbEngine == nil {
		return nil, errors.New("invalid value for required argument 'DbEngine'")
	}
	if args.DbMajorVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbMajorVersion'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ParameterTemplate
	err := ctx.RegisterResource("tencentcloud:Postgresql/parameterTemplate:ParameterTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterTemplate gets an existing ParameterTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterTemplateState, opts ...pulumi.ResourceOption) (*ParameterTemplate, error) {
	var resource ParameterTemplate
	err := ctx.ReadResource("tencentcloud:Postgresql/parameterTemplate:ParameterTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterTemplate resources.
type parameterTemplateState struct {
	// Database engine, such as postgresql, mssql_compatible.
	DbEngine *string `pulumi:"dbEngine"`
	// The major database version number, such as 11, 12, 13.
	DbMajorVersion *string `pulumi:"dbMajorVersion"`
	// The set of parameters that need to be deleted.
	DeleteParamSets []string `pulumi:"deleteParamSets"`
	// The set of parameters that need to be modified or added. Note: the same parameter cannot appear in the set of modifying and adding and deleting at the same time.
	ModifyParamEntrySets []ParameterTemplateModifyParamEntrySet `pulumi:"modifyParamEntrySets"`
	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateDescription *string `pulumi:"templateDescription"`
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateName *string `pulumi:"templateName"`
}

type ParameterTemplateState struct {
	// Database engine, such as postgresql, mssql_compatible.
	DbEngine pulumi.StringPtrInput
	// The major database version number, such as 11, 12, 13.
	DbMajorVersion pulumi.StringPtrInput
	// The set of parameters that need to be deleted.
	DeleteParamSets pulumi.StringArrayInput
	// The set of parameters that need to be modified or added. Note: the same parameter cannot appear in the set of modifying and adding and deleting at the same time.
	ModifyParamEntrySets ParameterTemplateModifyParamEntrySetArrayInput
	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateDescription pulumi.StringPtrInput
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateName pulumi.StringPtrInput
}

func (ParameterTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterTemplateState)(nil)).Elem()
}

type parameterTemplateArgs struct {
	// Database engine, such as postgresql, mssql_compatible.
	DbEngine string `pulumi:"dbEngine"`
	// The major database version number, such as 11, 12, 13.
	DbMajorVersion string `pulumi:"dbMajorVersion"`
	// The set of parameters that need to be deleted.
	DeleteParamSets []string `pulumi:"deleteParamSets"`
	// The set of parameters that need to be modified or added. Note: the same parameter cannot appear in the set of modifying and adding and deleting at the same time.
	ModifyParamEntrySets []ParameterTemplateModifyParamEntrySet `pulumi:"modifyParamEntrySets"`
	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateDescription *string `pulumi:"templateDescription"`
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a ParameterTemplate resource.
type ParameterTemplateArgs struct {
	// Database engine, such as postgresql, mssql_compatible.
	DbEngine pulumi.StringInput
	// The major database version number, such as 11, 12, 13.
	DbMajorVersion pulumi.StringInput
	// The set of parameters that need to be deleted.
	DeleteParamSets pulumi.StringArrayInput
	// The set of parameters that need to be modified or added. Note: the same parameter cannot appear in the set of modifying and adding and deleting at the same time.
	ModifyParamEntrySets ParameterTemplateModifyParamEntrySetArrayInput
	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateDescription pulumi.StringPtrInput
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
	TemplateName pulumi.StringInput
}

func (ParameterTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterTemplateArgs)(nil)).Elem()
}

type ParameterTemplateInput interface {
	pulumi.Input

	ToParameterTemplateOutput() ParameterTemplateOutput
	ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput
}

func (*ParameterTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterTemplate)(nil)).Elem()
}

func (i *ParameterTemplate) ToParameterTemplateOutput() ParameterTemplateOutput {
	return i.ToParameterTemplateOutputWithContext(context.Background())
}

func (i *ParameterTemplate) ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateOutput)
}

// ParameterTemplateArrayInput is an input type that accepts ParameterTemplateArray and ParameterTemplateArrayOutput values.
// You can construct a concrete instance of `ParameterTemplateArrayInput` via:
//
//          ParameterTemplateArray{ ParameterTemplateArgs{...} }
type ParameterTemplateArrayInput interface {
	pulumi.Input

	ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput
	ToParameterTemplateArrayOutputWithContext(context.Context) ParameterTemplateArrayOutput
}

type ParameterTemplateArray []ParameterTemplateInput

func (ParameterTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterTemplate)(nil)).Elem()
}

func (i ParameterTemplateArray) ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput {
	return i.ToParameterTemplateArrayOutputWithContext(context.Background())
}

func (i ParameterTemplateArray) ToParameterTemplateArrayOutputWithContext(ctx context.Context) ParameterTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateArrayOutput)
}

// ParameterTemplateMapInput is an input type that accepts ParameterTemplateMap and ParameterTemplateMapOutput values.
// You can construct a concrete instance of `ParameterTemplateMapInput` via:
//
//          ParameterTemplateMap{ "key": ParameterTemplateArgs{...} }
type ParameterTemplateMapInput interface {
	pulumi.Input

	ToParameterTemplateMapOutput() ParameterTemplateMapOutput
	ToParameterTemplateMapOutputWithContext(context.Context) ParameterTemplateMapOutput
}

type ParameterTemplateMap map[string]ParameterTemplateInput

func (ParameterTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterTemplate)(nil)).Elem()
}

func (i ParameterTemplateMap) ToParameterTemplateMapOutput() ParameterTemplateMapOutput {
	return i.ToParameterTemplateMapOutputWithContext(context.Background())
}

func (i ParameterTemplateMap) ToParameterTemplateMapOutputWithContext(ctx context.Context) ParameterTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateMapOutput)
}

type ParameterTemplateOutput struct{ *pulumi.OutputState }

func (ParameterTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateOutput) ToParameterTemplateOutput() ParameterTemplateOutput {
	return o
}

func (o ParameterTemplateOutput) ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput {
	return o
}

// Database engine, such as postgresql, mssql_compatible.
func (o ParameterTemplateOutput) DbEngine() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.DbEngine }).(pulumi.StringOutput)
}

// The major database version number, such as 11, 12, 13.
func (o ParameterTemplateOutput) DbMajorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.DbMajorVersion }).(pulumi.StringOutput)
}

// The set of parameters that need to be deleted.
func (o ParameterTemplateOutput) DeleteParamSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringArrayOutput { return v.DeleteParamSets }).(pulumi.StringArrayOutput)
}

// The set of parameters that need to be modified or added. Note: the same parameter cannot appear in the set of modifying and adding and deleting at the same time.
func (o ParameterTemplateOutput) ModifyParamEntrySets() ParameterTemplateModifyParamEntrySetArrayOutput {
	return o.ApplyT(func(v *ParameterTemplate) ParameterTemplateModifyParamEntrySetArrayOutput {
		return v.ModifyParamEntrySets
	}).(ParameterTemplateModifyParamEntrySetArrayOutput)
}

// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
func (o ParameterTemplateOutput) TemplateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringPtrOutput { return v.TemplateDescription }).(pulumi.StringPtrOutput)
}

// Template name, which can contain 1-60 letters, digits, and symbols (-_./()+=:@).
func (o ParameterTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type ParameterTemplateArrayOutput struct{ *pulumi.OutputState }

func (ParameterTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateArrayOutput) ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput {
	return o
}

func (o ParameterTemplateArrayOutput) ToParameterTemplateArrayOutputWithContext(ctx context.Context) ParameterTemplateArrayOutput {
	return o
}

func (o ParameterTemplateArrayOutput) Index(i pulumi.IntInput) ParameterTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ParameterTemplate {
		return vs[0].([]*ParameterTemplate)[vs[1].(int)]
	}).(ParameterTemplateOutput)
}

type ParameterTemplateMapOutput struct{ *pulumi.OutputState }

func (ParameterTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateMapOutput) ToParameterTemplateMapOutput() ParameterTemplateMapOutput {
	return o
}

func (o ParameterTemplateMapOutput) ToParameterTemplateMapOutputWithContext(ctx context.Context) ParameterTemplateMapOutput {
	return o
}

func (o ParameterTemplateMapOutput) MapIndex(k pulumi.StringInput) ParameterTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ParameterTemplate {
		return vs[0].(map[string]*ParameterTemplate)[vs[1].(string)]
	}).(ParameterTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateInput)(nil)).Elem(), &ParameterTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateArrayInput)(nil)).Elem(), ParameterTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateMapInput)(nil)).Elem(), ParameterTemplateMap{})
	pulumi.RegisterOutputType(ParameterTemplateOutput{})
	pulumi.RegisterOutputType(ParameterTemplateArrayOutput{})
	pulumi.RegisterOutputType(ParameterTemplateMapOutput{})
}
