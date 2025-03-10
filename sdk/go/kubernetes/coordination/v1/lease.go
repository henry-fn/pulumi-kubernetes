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

// Lease defines a lease concept.
type Lease struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// spec contains the specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecOutput `pulumi:"spec"`
}

// NewLease registers a new resource with the given unique name, arguments, and options.
func NewLease(ctx *pulumi.Context,
	name string, args *LeaseArgs, opts ...pulumi.ResourceOption) (*Lease, error) {
	if args == nil {
		args = &LeaseArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("coordination.k8s.io/v1")
	args.Kind = pulumi.StringPtr("Lease")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:coordination.k8s.io/v1beta1:Lease"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Lease
	err := ctx.RegisterResource("kubernetes:coordination.k8s.io/v1:Lease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLease gets an existing Lease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LeaseState, opts ...pulumi.ResourceOption) (*Lease, error) {
	var resource Lease
	err := ctx.ReadResource("kubernetes:coordination.k8s.io/v1:Lease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lease resources.
type leaseState struct {
}

type LeaseState struct {
}

func (LeaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*leaseState)(nil)).Elem()
}

type leaseArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec contains the specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *LeaseSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Lease resource.
type LeaseArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// spec contains the specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecPtrInput
}

func (LeaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*leaseArgs)(nil)).Elem()
}

type LeaseInput interface {
	pulumi.Input

	ToLeaseOutput() LeaseOutput
	ToLeaseOutputWithContext(ctx context.Context) LeaseOutput
}

func (*Lease) ElementType() reflect.Type {
	return reflect.TypeOf((**Lease)(nil)).Elem()
}

func (i *Lease) ToLeaseOutput() LeaseOutput {
	return i.ToLeaseOutputWithContext(context.Background())
}

func (i *Lease) ToLeaseOutputWithContext(ctx context.Context) LeaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseOutput)
}

func (i *Lease) ToOutput(ctx context.Context) pulumix.Output[*Lease] {
	return pulumix.Output[*Lease]{
		OutputState: i.ToLeaseOutputWithContext(ctx).OutputState,
	}
}

// LeaseArrayInput is an input type that accepts LeaseArray and LeaseArrayOutput values.
// You can construct a concrete instance of `LeaseArrayInput` via:
//
//	LeaseArray{ LeaseArgs{...} }
type LeaseArrayInput interface {
	pulumi.Input

	ToLeaseArrayOutput() LeaseArrayOutput
	ToLeaseArrayOutputWithContext(context.Context) LeaseArrayOutput
}

type LeaseArray []LeaseInput

func (LeaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Lease)(nil)).Elem()
}

func (i LeaseArray) ToLeaseArrayOutput() LeaseArrayOutput {
	return i.ToLeaseArrayOutputWithContext(context.Background())
}

func (i LeaseArray) ToLeaseArrayOutputWithContext(ctx context.Context) LeaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseArrayOutput)
}

func (i LeaseArray) ToOutput(ctx context.Context) pulumix.Output[[]*Lease] {
	return pulumix.Output[[]*Lease]{
		OutputState: i.ToLeaseArrayOutputWithContext(ctx).OutputState,
	}
}

// LeaseMapInput is an input type that accepts LeaseMap and LeaseMapOutput values.
// You can construct a concrete instance of `LeaseMapInput` via:
//
//	LeaseMap{ "key": LeaseArgs{...} }
type LeaseMapInput interface {
	pulumi.Input

	ToLeaseMapOutput() LeaseMapOutput
	ToLeaseMapOutputWithContext(context.Context) LeaseMapOutput
}

type LeaseMap map[string]LeaseInput

func (LeaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Lease)(nil)).Elem()
}

func (i LeaseMap) ToLeaseMapOutput() LeaseMapOutput {
	return i.ToLeaseMapOutputWithContext(context.Background())
}

func (i LeaseMap) ToLeaseMapOutputWithContext(ctx context.Context) LeaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseMapOutput)
}

func (i LeaseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Lease] {
	return pulumix.Output[map[string]*Lease]{
		OutputState: i.ToLeaseMapOutputWithContext(ctx).OutputState,
	}
}

type LeaseOutput struct{ *pulumi.OutputState }

func (LeaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lease)(nil)).Elem()
}

func (o LeaseOutput) ToLeaseOutput() LeaseOutput {
	return o
}

func (o LeaseOutput) ToLeaseOutputWithContext(ctx context.Context) LeaseOutput {
	return o
}

func (o LeaseOutput) ToOutput(ctx context.Context) pulumix.Output[*Lease] {
	return pulumix.Output[*Lease]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LeaseOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Lease) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LeaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Lease) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o LeaseOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *Lease) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// spec contains the specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o LeaseOutput) Spec() LeaseSpecOutput {
	return o.ApplyT(func(v *Lease) LeaseSpecOutput { return v.Spec }).(LeaseSpecOutput)
}

type LeaseArrayOutput struct{ *pulumi.OutputState }

func (LeaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Lease)(nil)).Elem()
}

func (o LeaseArrayOutput) ToLeaseArrayOutput() LeaseArrayOutput {
	return o
}

func (o LeaseArrayOutput) ToLeaseArrayOutputWithContext(ctx context.Context) LeaseArrayOutput {
	return o
}

func (o LeaseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Lease] {
	return pulumix.Output[[]*Lease]{
		OutputState: o.OutputState,
	}
}

func (o LeaseArrayOutput) Index(i pulumi.IntInput) LeaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Lease {
		return vs[0].([]*Lease)[vs[1].(int)]
	}).(LeaseOutput)
}

type LeaseMapOutput struct{ *pulumi.OutputState }

func (LeaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Lease)(nil)).Elem()
}

func (o LeaseMapOutput) ToLeaseMapOutput() LeaseMapOutput {
	return o
}

func (o LeaseMapOutput) ToLeaseMapOutputWithContext(ctx context.Context) LeaseMapOutput {
	return o
}

func (o LeaseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Lease] {
	return pulumix.Output[map[string]*Lease]{
		OutputState: o.OutputState,
	}
}

func (o LeaseMapOutput) MapIndex(k pulumi.StringInput) LeaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Lease {
		return vs[0].(map[string]*Lease)[vs[1].(string)]
	}).(LeaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseInput)(nil)).Elem(), &Lease{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseArrayInput)(nil)).Elem(), LeaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseMapInput)(nil)).Elem(), LeaseMap{})
	pulumi.RegisterOutputType(LeaseOutput{})
	pulumi.RegisterOutputType(LeaseArrayOutput{})
	pulumi.RegisterOutputType(LeaseMapOutput{})
}
