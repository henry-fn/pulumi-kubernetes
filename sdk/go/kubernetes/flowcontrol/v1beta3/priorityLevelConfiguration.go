// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta3

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// PriorityLevelConfiguration represents the configuration of a priority level.
type PriorityLevelConfiguration struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// `spec` is the specification of the desired behavior of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PriorityLevelConfigurationSpecOutput `pulumi:"spec"`
	// `status` is the current status of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status PriorityLevelConfigurationStatusPtrOutput `pulumi:"status"`
}

// NewPriorityLevelConfiguration registers a new resource with the given unique name, arguments, and options.
func NewPriorityLevelConfiguration(ctx *pulumi.Context,
	name string, args *PriorityLevelConfigurationArgs, opts ...pulumi.ResourceOption) (*PriorityLevelConfiguration, error) {
	if args == nil {
		args = &PriorityLevelConfigurationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("flowcontrol.apiserver.k8s.io/v1beta3")
	args.Kind = pulumi.StringPtr("PriorityLevelConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfiguration"),
		},
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1beta1:PriorityLevelConfiguration"),
		},
		{
			Type: pulumi.String("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:PriorityLevelConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PriorityLevelConfiguration
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:PriorityLevelConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityLevelConfiguration gets an existing PriorityLevelConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPriorityLevelConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriorityLevelConfigurationState, opts ...pulumi.ResourceOption) (*PriorityLevelConfiguration, error) {
	var resource PriorityLevelConfiguration
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:PriorityLevelConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PriorityLevelConfiguration resources.
type priorityLevelConfigurationState struct {
}

type PriorityLevelConfigurationState struct {
}

func (PriorityLevelConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityLevelConfigurationState)(nil)).Elem()
}

type priorityLevelConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// `spec` is the specification of the desired behavior of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PriorityLevelConfigurationSpec `pulumi:"spec"`
}

// The set of arguments for constructing a PriorityLevelConfiguration resource.
type PriorityLevelConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// `spec` is the specification of the desired behavior of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PriorityLevelConfigurationSpecPtrInput
}

func (PriorityLevelConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityLevelConfigurationArgs)(nil)).Elem()
}

type PriorityLevelConfigurationInput interface {
	pulumi.Input

	ToPriorityLevelConfigurationOutput() PriorityLevelConfigurationOutput
	ToPriorityLevelConfigurationOutputWithContext(ctx context.Context) PriorityLevelConfigurationOutput
}

func (*PriorityLevelConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityLevelConfiguration)(nil)).Elem()
}

func (i *PriorityLevelConfiguration) ToPriorityLevelConfigurationOutput() PriorityLevelConfigurationOutput {
	return i.ToPriorityLevelConfigurationOutputWithContext(context.Background())
}

func (i *PriorityLevelConfiguration) ToPriorityLevelConfigurationOutputWithContext(ctx context.Context) PriorityLevelConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityLevelConfigurationOutput)
}

func (i *PriorityLevelConfiguration) ToOutput(ctx context.Context) pulumix.Output[*PriorityLevelConfiguration] {
	return pulumix.Output[*PriorityLevelConfiguration]{
		OutputState: i.ToPriorityLevelConfigurationOutputWithContext(ctx).OutputState,
	}
}

// PriorityLevelConfigurationArrayInput is an input type that accepts PriorityLevelConfigurationArray and PriorityLevelConfigurationArrayOutput values.
// You can construct a concrete instance of `PriorityLevelConfigurationArrayInput` via:
//
//	PriorityLevelConfigurationArray{ PriorityLevelConfigurationArgs{...} }
type PriorityLevelConfigurationArrayInput interface {
	pulumi.Input

	ToPriorityLevelConfigurationArrayOutput() PriorityLevelConfigurationArrayOutput
	ToPriorityLevelConfigurationArrayOutputWithContext(context.Context) PriorityLevelConfigurationArrayOutput
}

type PriorityLevelConfigurationArray []PriorityLevelConfigurationInput

func (PriorityLevelConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityLevelConfiguration)(nil)).Elem()
}

func (i PriorityLevelConfigurationArray) ToPriorityLevelConfigurationArrayOutput() PriorityLevelConfigurationArrayOutput {
	return i.ToPriorityLevelConfigurationArrayOutputWithContext(context.Background())
}

func (i PriorityLevelConfigurationArray) ToPriorityLevelConfigurationArrayOutputWithContext(ctx context.Context) PriorityLevelConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityLevelConfigurationArrayOutput)
}

func (i PriorityLevelConfigurationArray) ToOutput(ctx context.Context) pulumix.Output[[]*PriorityLevelConfiguration] {
	return pulumix.Output[[]*PriorityLevelConfiguration]{
		OutputState: i.ToPriorityLevelConfigurationArrayOutputWithContext(ctx).OutputState,
	}
}

// PriorityLevelConfigurationMapInput is an input type that accepts PriorityLevelConfigurationMap and PriorityLevelConfigurationMapOutput values.
// You can construct a concrete instance of `PriorityLevelConfigurationMapInput` via:
//
//	PriorityLevelConfigurationMap{ "key": PriorityLevelConfigurationArgs{...} }
type PriorityLevelConfigurationMapInput interface {
	pulumi.Input

	ToPriorityLevelConfigurationMapOutput() PriorityLevelConfigurationMapOutput
	ToPriorityLevelConfigurationMapOutputWithContext(context.Context) PriorityLevelConfigurationMapOutput
}

type PriorityLevelConfigurationMap map[string]PriorityLevelConfigurationInput

func (PriorityLevelConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityLevelConfiguration)(nil)).Elem()
}

func (i PriorityLevelConfigurationMap) ToPriorityLevelConfigurationMapOutput() PriorityLevelConfigurationMapOutput {
	return i.ToPriorityLevelConfigurationMapOutputWithContext(context.Background())
}

func (i PriorityLevelConfigurationMap) ToPriorityLevelConfigurationMapOutputWithContext(ctx context.Context) PriorityLevelConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityLevelConfigurationMapOutput)
}

func (i PriorityLevelConfigurationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PriorityLevelConfiguration] {
	return pulumix.Output[map[string]*PriorityLevelConfiguration]{
		OutputState: i.ToPriorityLevelConfigurationMapOutputWithContext(ctx).OutputState,
	}
}

type PriorityLevelConfigurationOutput struct{ *pulumi.OutputState }

func (PriorityLevelConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityLevelConfiguration)(nil)).Elem()
}

func (o PriorityLevelConfigurationOutput) ToPriorityLevelConfigurationOutput() PriorityLevelConfigurationOutput {
	return o
}

func (o PriorityLevelConfigurationOutput) ToPriorityLevelConfigurationOutputWithContext(ctx context.Context) PriorityLevelConfigurationOutput {
	return o
}

func (o PriorityLevelConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*PriorityLevelConfiguration] {
	return pulumix.Output[*PriorityLevelConfiguration]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PriorityLevelConfigurationOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PriorityLevelConfiguration) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PriorityLevelConfigurationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PriorityLevelConfiguration) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PriorityLevelConfigurationOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *PriorityLevelConfiguration) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// `spec` is the specification of the desired behavior of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o PriorityLevelConfigurationOutput) Spec() PriorityLevelConfigurationSpecOutput {
	return o.ApplyT(func(v *PriorityLevelConfiguration) PriorityLevelConfigurationSpecOutput { return v.Spec }).(PriorityLevelConfigurationSpecOutput)
}

// `status` is the current status of a "request-priority". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o PriorityLevelConfigurationOutput) Status() PriorityLevelConfigurationStatusPtrOutput {
	return o.ApplyT(func(v *PriorityLevelConfiguration) PriorityLevelConfigurationStatusPtrOutput { return v.Status }).(PriorityLevelConfigurationStatusPtrOutput)
}

type PriorityLevelConfigurationArrayOutput struct{ *pulumi.OutputState }

func (PriorityLevelConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityLevelConfiguration)(nil)).Elem()
}

func (o PriorityLevelConfigurationArrayOutput) ToPriorityLevelConfigurationArrayOutput() PriorityLevelConfigurationArrayOutput {
	return o
}

func (o PriorityLevelConfigurationArrayOutput) ToPriorityLevelConfigurationArrayOutputWithContext(ctx context.Context) PriorityLevelConfigurationArrayOutput {
	return o
}

func (o PriorityLevelConfigurationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PriorityLevelConfiguration] {
	return pulumix.Output[[]*PriorityLevelConfiguration]{
		OutputState: o.OutputState,
	}
}

func (o PriorityLevelConfigurationArrayOutput) Index(i pulumi.IntInput) PriorityLevelConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PriorityLevelConfiguration {
		return vs[0].([]*PriorityLevelConfiguration)[vs[1].(int)]
	}).(PriorityLevelConfigurationOutput)
}

type PriorityLevelConfigurationMapOutput struct{ *pulumi.OutputState }

func (PriorityLevelConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityLevelConfiguration)(nil)).Elem()
}

func (o PriorityLevelConfigurationMapOutput) ToPriorityLevelConfigurationMapOutput() PriorityLevelConfigurationMapOutput {
	return o
}

func (o PriorityLevelConfigurationMapOutput) ToPriorityLevelConfigurationMapOutputWithContext(ctx context.Context) PriorityLevelConfigurationMapOutput {
	return o
}

func (o PriorityLevelConfigurationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PriorityLevelConfiguration] {
	return pulumix.Output[map[string]*PriorityLevelConfiguration]{
		OutputState: o.OutputState,
	}
}

func (o PriorityLevelConfigurationMapOutput) MapIndex(k pulumi.StringInput) PriorityLevelConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PriorityLevelConfiguration {
		return vs[0].(map[string]*PriorityLevelConfiguration)[vs[1].(string)]
	}).(PriorityLevelConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityLevelConfigurationInput)(nil)).Elem(), &PriorityLevelConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityLevelConfigurationArrayInput)(nil)).Elem(), PriorityLevelConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityLevelConfigurationMapInput)(nil)).Elem(), PriorityLevelConfigurationMap{})
	pulumi.RegisterOutputType(PriorityLevelConfigurationOutput{})
	pulumi.RegisterOutputType(PriorityLevelConfigurationArrayOutput{})
	pulumi.RegisterOutputType(PriorityLevelConfigurationMapOutput{})
}
