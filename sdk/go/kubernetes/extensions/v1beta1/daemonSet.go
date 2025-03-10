// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// DaemonSet represents the configuration of a daemon set.
//
// Deprecated: extensions/v1beta1/DaemonSet is deprecated by apps/v1/DaemonSet and not supported by Kubernetes v1.16+ clusters.
type DaemonSet struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecOutput `pulumi:"spec"`
	// The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status DaemonSetStatusPtrOutput `pulumi:"status"`
}

// NewDaemonSet registers a new resource with the given unique name, arguments, and options.
func NewDaemonSet(ctx *pulumi.Context,
	name string, args *DaemonSetArgs, opts ...pulumi.ResourceOption) (*DaemonSet, error) {
	if args == nil {
		args = &DaemonSetArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("DaemonSet")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apps/v1:DaemonSet"),
		},
		{
			Type: pulumi.String("kubernetes:apps/v1beta2:DaemonSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DaemonSet
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:DaemonSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDaemonSet gets an existing DaemonSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDaemonSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DaemonSetState, opts ...pulumi.ResourceOption) (*DaemonSet, error) {
	var resource DaemonSet
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:DaemonSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DaemonSet resources.
type daemonSetState struct {
}

type DaemonSetState struct {
}

func (DaemonSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*daemonSetState)(nil)).Elem()
}

type daemonSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *DaemonSetSpec `pulumi:"spec"`
}

// The set of arguments for constructing a DaemonSet resource.
type DaemonSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecPtrInput
}

func (DaemonSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*daemonSetArgs)(nil)).Elem()
}

type DaemonSetInput interface {
	pulumi.Input

	ToDaemonSetOutput() DaemonSetOutput
	ToDaemonSetOutputWithContext(ctx context.Context) DaemonSetOutput
}

func (*DaemonSet) ElementType() reflect.Type {
	return reflect.TypeOf((**DaemonSet)(nil)).Elem()
}

func (i *DaemonSet) ToDaemonSetOutput() DaemonSetOutput {
	return i.ToDaemonSetOutputWithContext(context.Background())
}

func (i *DaemonSet) ToDaemonSetOutputWithContext(ctx context.Context) DaemonSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetOutput)
}

func (i *DaemonSet) ToOutput(ctx context.Context) pulumix.Output[*DaemonSet] {
	return pulumix.Output[*DaemonSet]{
		OutputState: i.ToDaemonSetOutputWithContext(ctx).OutputState,
	}
}

// DaemonSetArrayInput is an input type that accepts DaemonSetArray and DaemonSetArrayOutput values.
// You can construct a concrete instance of `DaemonSetArrayInput` via:
//
//	DaemonSetArray{ DaemonSetArgs{...} }
type DaemonSetArrayInput interface {
	pulumi.Input

	ToDaemonSetArrayOutput() DaemonSetArrayOutput
	ToDaemonSetArrayOutputWithContext(context.Context) DaemonSetArrayOutput
}

type DaemonSetArray []DaemonSetInput

func (DaemonSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DaemonSet)(nil)).Elem()
}

func (i DaemonSetArray) ToDaemonSetArrayOutput() DaemonSetArrayOutput {
	return i.ToDaemonSetArrayOutputWithContext(context.Background())
}

func (i DaemonSetArray) ToDaemonSetArrayOutputWithContext(ctx context.Context) DaemonSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetArrayOutput)
}

func (i DaemonSetArray) ToOutput(ctx context.Context) pulumix.Output[[]*DaemonSet] {
	return pulumix.Output[[]*DaemonSet]{
		OutputState: i.ToDaemonSetArrayOutputWithContext(ctx).OutputState,
	}
}

// DaemonSetMapInput is an input type that accepts DaemonSetMap and DaemonSetMapOutput values.
// You can construct a concrete instance of `DaemonSetMapInput` via:
//
//	DaemonSetMap{ "key": DaemonSetArgs{...} }
type DaemonSetMapInput interface {
	pulumi.Input

	ToDaemonSetMapOutput() DaemonSetMapOutput
	ToDaemonSetMapOutputWithContext(context.Context) DaemonSetMapOutput
}

type DaemonSetMap map[string]DaemonSetInput

func (DaemonSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DaemonSet)(nil)).Elem()
}

func (i DaemonSetMap) ToDaemonSetMapOutput() DaemonSetMapOutput {
	return i.ToDaemonSetMapOutputWithContext(context.Background())
}

func (i DaemonSetMap) ToDaemonSetMapOutputWithContext(ctx context.Context) DaemonSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetMapOutput)
}

func (i DaemonSetMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DaemonSet] {
	return pulumix.Output[map[string]*DaemonSet]{
		OutputState: i.ToDaemonSetMapOutputWithContext(ctx).OutputState,
	}
}

type DaemonSetOutput struct{ *pulumi.OutputState }

func (DaemonSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaemonSet)(nil)).Elem()
}

func (o DaemonSetOutput) ToDaemonSetOutput() DaemonSetOutput {
	return o
}

func (o DaemonSetOutput) ToDaemonSetOutputWithContext(ctx context.Context) DaemonSetOutput {
	return o
}

func (o DaemonSetOutput) ToOutput(ctx context.Context) pulumix.Output[*DaemonSet] {
	return pulumix.Output[*DaemonSet]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o DaemonSetOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DaemonSet) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o DaemonSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DaemonSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o DaemonSetOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *DaemonSet) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o DaemonSetOutput) Spec() DaemonSetSpecOutput {
	return o.ApplyT(func(v *DaemonSet) DaemonSetSpecOutput { return v.Spec }).(DaemonSetSpecOutput)
}

// The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o DaemonSetOutput) Status() DaemonSetStatusPtrOutput {
	return o.ApplyT(func(v *DaemonSet) DaemonSetStatusPtrOutput { return v.Status }).(DaemonSetStatusPtrOutput)
}

type DaemonSetArrayOutput struct{ *pulumi.OutputState }

func (DaemonSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DaemonSet)(nil)).Elem()
}

func (o DaemonSetArrayOutput) ToDaemonSetArrayOutput() DaemonSetArrayOutput {
	return o
}

func (o DaemonSetArrayOutput) ToDaemonSetArrayOutputWithContext(ctx context.Context) DaemonSetArrayOutput {
	return o
}

func (o DaemonSetArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DaemonSet] {
	return pulumix.Output[[]*DaemonSet]{
		OutputState: o.OutputState,
	}
}

func (o DaemonSetArrayOutput) Index(i pulumi.IntInput) DaemonSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DaemonSet {
		return vs[0].([]*DaemonSet)[vs[1].(int)]
	}).(DaemonSetOutput)
}

type DaemonSetMapOutput struct{ *pulumi.OutputState }

func (DaemonSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DaemonSet)(nil)).Elem()
}

func (o DaemonSetMapOutput) ToDaemonSetMapOutput() DaemonSetMapOutput {
	return o
}

func (o DaemonSetMapOutput) ToDaemonSetMapOutputWithContext(ctx context.Context) DaemonSetMapOutput {
	return o
}

func (o DaemonSetMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DaemonSet] {
	return pulumix.Output[map[string]*DaemonSet]{
		OutputState: o.OutputState,
	}
}

func (o DaemonSetMapOutput) MapIndex(k pulumi.StringInput) DaemonSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DaemonSet {
		return vs[0].(map[string]*DaemonSet)[vs[1].(string)]
	}).(DaemonSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetInput)(nil)).Elem(), &DaemonSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetArrayInput)(nil)).Elem(), DaemonSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetMapInput)(nil)).Elem(), DaemonSetMap{})
	pulumi.RegisterOutputType(DaemonSetOutput{})
	pulumi.RegisterOutputType(DaemonSetArrayOutput{})
	pulumi.RegisterOutputType(DaemonSetMapOutput{})
}
