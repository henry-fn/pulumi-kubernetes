// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// PodTemplate describes a template for creating copies of a predefined pod.
type PodTemplate struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Template defines the pods that will be created from this pod template. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Template PodTemplateSpecOutput `pulumi:"template"`
}

// NewPodTemplate registers a new resource with the given unique name, arguments, and options.
func NewPodTemplate(ctx *pulumi.Context,
	name string, args *PodTemplateArgs, opts ...pulumi.ResourceOption) (*PodTemplate, error) {
	if args == nil {
		args = &PodTemplateArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PodTemplate")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PodTemplate
	err := ctx.RegisterResource("kubernetes:core/v1:PodTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodTemplate gets an existing PodTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodTemplateState, opts ...pulumi.ResourceOption) (*PodTemplate, error) {
	var resource PodTemplate
	err := ctx.ReadResource("kubernetes:core/v1:PodTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodTemplate resources.
type podTemplateState struct {
}

type PodTemplateState struct {
}

func (PodTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*podTemplateState)(nil)).Elem()
}

type podTemplateArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Template defines the pods that will be created from this pod template. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Template *PodTemplateSpec `pulumi:"template"`
}

// The set of arguments for constructing a PodTemplate resource.
type PodTemplateArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Template defines the pods that will be created from this pod template. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Template PodTemplateSpecPtrInput
}

func (PodTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podTemplateArgs)(nil)).Elem()
}

type PodTemplateInput interface {
	pulumi.Input

	ToPodTemplateOutput() PodTemplateOutput
	ToPodTemplateOutputWithContext(ctx context.Context) PodTemplateOutput
}

func (*PodTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**PodTemplate)(nil)).Elem()
}

func (i *PodTemplate) ToPodTemplateOutput() PodTemplateOutput {
	return i.ToPodTemplateOutputWithContext(context.Background())
}

func (i *PodTemplate) ToPodTemplateOutputWithContext(ctx context.Context) PodTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodTemplateOutput)
}

func (i *PodTemplate) ToOutput(ctx context.Context) pulumix.Output[*PodTemplate] {
	return pulumix.Output[*PodTemplate]{
		OutputState: i.ToPodTemplateOutputWithContext(ctx).OutputState,
	}
}

// PodTemplateArrayInput is an input type that accepts PodTemplateArray and PodTemplateArrayOutput values.
// You can construct a concrete instance of `PodTemplateArrayInput` via:
//
//	PodTemplateArray{ PodTemplateArgs{...} }
type PodTemplateArrayInput interface {
	pulumi.Input

	ToPodTemplateArrayOutput() PodTemplateArrayOutput
	ToPodTemplateArrayOutputWithContext(context.Context) PodTemplateArrayOutput
}

type PodTemplateArray []PodTemplateInput

func (PodTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodTemplate)(nil)).Elem()
}

func (i PodTemplateArray) ToPodTemplateArrayOutput() PodTemplateArrayOutput {
	return i.ToPodTemplateArrayOutputWithContext(context.Background())
}

func (i PodTemplateArray) ToPodTemplateArrayOutputWithContext(ctx context.Context) PodTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodTemplateArrayOutput)
}

func (i PodTemplateArray) ToOutput(ctx context.Context) pulumix.Output[[]*PodTemplate] {
	return pulumix.Output[[]*PodTemplate]{
		OutputState: i.ToPodTemplateArrayOutputWithContext(ctx).OutputState,
	}
}

// PodTemplateMapInput is an input type that accepts PodTemplateMap and PodTemplateMapOutput values.
// You can construct a concrete instance of `PodTemplateMapInput` via:
//
//	PodTemplateMap{ "key": PodTemplateArgs{...} }
type PodTemplateMapInput interface {
	pulumi.Input

	ToPodTemplateMapOutput() PodTemplateMapOutput
	ToPodTemplateMapOutputWithContext(context.Context) PodTemplateMapOutput
}

type PodTemplateMap map[string]PodTemplateInput

func (PodTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodTemplate)(nil)).Elem()
}

func (i PodTemplateMap) ToPodTemplateMapOutput() PodTemplateMapOutput {
	return i.ToPodTemplateMapOutputWithContext(context.Background())
}

func (i PodTemplateMap) ToPodTemplateMapOutputWithContext(ctx context.Context) PodTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodTemplateMapOutput)
}

func (i PodTemplateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PodTemplate] {
	return pulumix.Output[map[string]*PodTemplate]{
		OutputState: i.ToPodTemplateMapOutputWithContext(ctx).OutputState,
	}
}

type PodTemplateOutput struct{ *pulumi.OutputState }

func (PodTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodTemplate)(nil)).Elem()
}

func (o PodTemplateOutput) ToPodTemplateOutput() PodTemplateOutput {
	return o
}

func (o PodTemplateOutput) ToPodTemplateOutputWithContext(ctx context.Context) PodTemplateOutput {
	return o
}

func (o PodTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*PodTemplate] {
	return pulumix.Output[*PodTemplate]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodTemplateOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PodTemplate) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodTemplateOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PodTemplate) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PodTemplateOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *PodTemplate) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Template defines the pods that will be created from this pod template. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o PodTemplateOutput) Template() PodTemplateSpecOutput {
	return o.ApplyT(func(v *PodTemplate) PodTemplateSpecOutput { return v.Template }).(PodTemplateSpecOutput)
}

type PodTemplateArrayOutput struct{ *pulumi.OutputState }

func (PodTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodTemplate)(nil)).Elem()
}

func (o PodTemplateArrayOutput) ToPodTemplateArrayOutput() PodTemplateArrayOutput {
	return o
}

func (o PodTemplateArrayOutput) ToPodTemplateArrayOutputWithContext(ctx context.Context) PodTemplateArrayOutput {
	return o
}

func (o PodTemplateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PodTemplate] {
	return pulumix.Output[[]*PodTemplate]{
		OutputState: o.OutputState,
	}
}

func (o PodTemplateArrayOutput) Index(i pulumi.IntInput) PodTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodTemplate {
		return vs[0].([]*PodTemplate)[vs[1].(int)]
	}).(PodTemplateOutput)
}

type PodTemplateMapOutput struct{ *pulumi.OutputState }

func (PodTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodTemplate)(nil)).Elem()
}

func (o PodTemplateMapOutput) ToPodTemplateMapOutput() PodTemplateMapOutput {
	return o
}

func (o PodTemplateMapOutput) ToPodTemplateMapOutputWithContext(ctx context.Context) PodTemplateMapOutput {
	return o
}

func (o PodTemplateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PodTemplate] {
	return pulumix.Output[map[string]*PodTemplate]{
		OutputState: o.OutputState,
	}
}

func (o PodTemplateMapOutput) MapIndex(k pulumi.StringInput) PodTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodTemplate {
		return vs[0].(map[string]*PodTemplate)[vs[1].(string)]
	}).(PodTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodTemplateInput)(nil)).Elem(), &PodTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodTemplateArrayInput)(nil)).Elem(), PodTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodTemplateMapInput)(nil)).Elem(), PodTemplateMap{})
	pulumi.RegisterOutputType(PodTemplateOutput{})
	pulumi.RegisterOutputType(PodTemplateArrayOutput{})
	pulumi.RegisterOutputType(PodTemplateMapOutput{})
}
