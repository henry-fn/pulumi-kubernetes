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

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.
type CSIDriverPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// spec represents the specification of the CSI Driver.
	Spec CSIDriverSpecPatchPtrOutput `pulumi:"spec"`
}

// NewCSIDriverPatch registers a new resource with the given unique name, arguments, and options.
func NewCSIDriverPatch(ctx *pulumi.Context,
	name string, args *CSIDriverPatchArgs, opts ...pulumi.ResourceOption) (*CSIDriverPatch, error) {
	if args == nil {
		args = &CSIDriverPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSIDriver")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:CSIDriverPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CSIDriverPatch
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSIDriverPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIDriverPatch gets an existing CSIDriverPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIDriverPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIDriverPatchState, opts ...pulumi.ResourceOption) (*CSIDriverPatch, error) {
	var resource CSIDriverPatch
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSIDriverPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIDriverPatch resources.
type csidriverPatchState struct {
}

type CSIDriverPatchState struct {
}

func (CSIDriverPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverPatchState)(nil)).Elem()
}

type csidriverPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// spec represents the specification of the CSI Driver.
	Spec *CSIDriverSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a CSIDriverPatch resource.
type CSIDriverPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// spec represents the specification of the CSI Driver.
	Spec CSIDriverSpecPatchPtrInput
}

func (CSIDriverPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverPatchArgs)(nil)).Elem()
}

type CSIDriverPatchInput interface {
	pulumi.Input

	ToCSIDriverPatchOutput() CSIDriverPatchOutput
	ToCSIDriverPatchOutputWithContext(ctx context.Context) CSIDriverPatchOutput
}

func (*CSIDriverPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIDriverPatch)(nil)).Elem()
}

func (i *CSIDriverPatch) ToCSIDriverPatchOutput() CSIDriverPatchOutput {
	return i.ToCSIDriverPatchOutputWithContext(context.Background())
}

func (i *CSIDriverPatch) ToCSIDriverPatchOutputWithContext(ctx context.Context) CSIDriverPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverPatchOutput)
}

func (i *CSIDriverPatch) ToOutput(ctx context.Context) pulumix.Output[*CSIDriverPatch] {
	return pulumix.Output[*CSIDriverPatch]{
		OutputState: i.ToCSIDriverPatchOutputWithContext(ctx).OutputState,
	}
}

// CSIDriverPatchArrayInput is an input type that accepts CSIDriverPatchArray and CSIDriverPatchArrayOutput values.
// You can construct a concrete instance of `CSIDriverPatchArrayInput` via:
//
//	CSIDriverPatchArray{ CSIDriverPatchArgs{...} }
type CSIDriverPatchArrayInput interface {
	pulumi.Input

	ToCSIDriverPatchArrayOutput() CSIDriverPatchArrayOutput
	ToCSIDriverPatchArrayOutputWithContext(context.Context) CSIDriverPatchArrayOutput
}

type CSIDriverPatchArray []CSIDriverPatchInput

func (CSIDriverPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIDriverPatch)(nil)).Elem()
}

func (i CSIDriverPatchArray) ToCSIDriverPatchArrayOutput() CSIDriverPatchArrayOutput {
	return i.ToCSIDriverPatchArrayOutputWithContext(context.Background())
}

func (i CSIDriverPatchArray) ToCSIDriverPatchArrayOutputWithContext(ctx context.Context) CSIDriverPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverPatchArrayOutput)
}

func (i CSIDriverPatchArray) ToOutput(ctx context.Context) pulumix.Output[[]*CSIDriverPatch] {
	return pulumix.Output[[]*CSIDriverPatch]{
		OutputState: i.ToCSIDriverPatchArrayOutputWithContext(ctx).OutputState,
	}
}

// CSIDriverPatchMapInput is an input type that accepts CSIDriverPatchMap and CSIDriverPatchMapOutput values.
// You can construct a concrete instance of `CSIDriverPatchMapInput` via:
//
//	CSIDriverPatchMap{ "key": CSIDriverPatchArgs{...} }
type CSIDriverPatchMapInput interface {
	pulumi.Input

	ToCSIDriverPatchMapOutput() CSIDriverPatchMapOutput
	ToCSIDriverPatchMapOutputWithContext(context.Context) CSIDriverPatchMapOutput
}

type CSIDriverPatchMap map[string]CSIDriverPatchInput

func (CSIDriverPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIDriverPatch)(nil)).Elem()
}

func (i CSIDriverPatchMap) ToCSIDriverPatchMapOutput() CSIDriverPatchMapOutput {
	return i.ToCSIDriverPatchMapOutputWithContext(context.Background())
}

func (i CSIDriverPatchMap) ToCSIDriverPatchMapOutputWithContext(ctx context.Context) CSIDriverPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverPatchMapOutput)
}

func (i CSIDriverPatchMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CSIDriverPatch] {
	return pulumix.Output[map[string]*CSIDriverPatch]{
		OutputState: i.ToCSIDriverPatchMapOutputWithContext(ctx).OutputState,
	}
}

type CSIDriverPatchOutput struct{ *pulumi.OutputState }

func (CSIDriverPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIDriverPatch)(nil)).Elem()
}

func (o CSIDriverPatchOutput) ToCSIDriverPatchOutput() CSIDriverPatchOutput {
	return o
}

func (o CSIDriverPatchOutput) ToCSIDriverPatchOutputWithContext(ctx context.Context) CSIDriverPatchOutput {
	return o
}

func (o CSIDriverPatchOutput) ToOutput(ctx context.Context) pulumix.Output[*CSIDriverPatch] {
	return pulumix.Output[*CSIDriverPatch]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSIDriverPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIDriverPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSIDriverPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSIDriverPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CSIDriverPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *CSIDriverPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// spec represents the specification of the CSI Driver.
func (o CSIDriverPatchOutput) Spec() CSIDriverSpecPatchPtrOutput {
	return o.ApplyT(func(v *CSIDriverPatch) CSIDriverSpecPatchPtrOutput { return v.Spec }).(CSIDriverSpecPatchPtrOutput)
}

type CSIDriverPatchArrayOutput struct{ *pulumi.OutputState }

func (CSIDriverPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIDriverPatch)(nil)).Elem()
}

func (o CSIDriverPatchArrayOutput) ToCSIDriverPatchArrayOutput() CSIDriverPatchArrayOutput {
	return o
}

func (o CSIDriverPatchArrayOutput) ToCSIDriverPatchArrayOutputWithContext(ctx context.Context) CSIDriverPatchArrayOutput {
	return o
}

func (o CSIDriverPatchArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CSIDriverPatch] {
	return pulumix.Output[[]*CSIDriverPatch]{
		OutputState: o.OutputState,
	}
}

func (o CSIDriverPatchArrayOutput) Index(i pulumi.IntInput) CSIDriverPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSIDriverPatch {
		return vs[0].([]*CSIDriverPatch)[vs[1].(int)]
	}).(CSIDriverPatchOutput)
}

type CSIDriverPatchMapOutput struct{ *pulumi.OutputState }

func (CSIDriverPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIDriverPatch)(nil)).Elem()
}

func (o CSIDriverPatchMapOutput) ToCSIDriverPatchMapOutput() CSIDriverPatchMapOutput {
	return o
}

func (o CSIDriverPatchMapOutput) ToCSIDriverPatchMapOutputWithContext(ctx context.Context) CSIDriverPatchMapOutput {
	return o
}

func (o CSIDriverPatchMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CSIDriverPatch] {
	return pulumix.Output[map[string]*CSIDriverPatch]{
		OutputState: o.OutputState,
	}
}

func (o CSIDriverPatchMapOutput) MapIndex(k pulumi.StringInput) CSIDriverPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSIDriverPatch {
		return vs[0].(map[string]*CSIDriverPatch)[vs[1].(string)]
	}).(CSIDriverPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverPatchInput)(nil)).Elem(), &CSIDriverPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverPatchArrayInput)(nil)).Elem(), CSIDriverPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverPatchMapInput)(nil)).Elem(), CSIDriverPatchMap{})
	pulumi.RegisterOutputType(CSIDriverPatchOutput{})
	pulumi.RegisterOutputType(CSIDriverPatchArrayOutput{})
	pulumi.RegisterOutputType(CSIDriverPatchMapOutput{})
}
