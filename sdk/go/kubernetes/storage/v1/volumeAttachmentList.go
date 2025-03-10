// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// VolumeAttachmentList is a collection of VolumeAttachment objects.
type VolumeAttachmentList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is the list of VolumeAttachments
	Items VolumeAttachmentTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewVolumeAttachmentList registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachmentList(ctx *pulumi.Context,
	name string, args *VolumeAttachmentListArgs, opts ...pulumi.ResourceOption) (*VolumeAttachmentList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("VolumeAttachmentList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VolumeAttachmentList
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:VolumeAttachmentList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachmentList gets an existing VolumeAttachmentList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachmentList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachmentListState, opts ...pulumi.ResourceOption) (*VolumeAttachmentList, error) {
	var resource VolumeAttachmentList
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:VolumeAttachmentList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachmentList resources.
type volumeAttachmentListState struct {
}

type VolumeAttachmentListState struct {
}

func (VolumeAttachmentListState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentListState)(nil)).Elem()
}

type volumeAttachmentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of VolumeAttachments
	Items []VolumeAttachmentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a VolumeAttachmentList resource.
type VolumeAttachmentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of VolumeAttachments
	Items VolumeAttachmentTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (VolumeAttachmentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentListArgs)(nil)).Elem()
}

type VolumeAttachmentListInput interface {
	pulumi.Input

	ToVolumeAttachmentListOutput() VolumeAttachmentListOutput
	ToVolumeAttachmentListOutputWithContext(ctx context.Context) VolumeAttachmentListOutput
}

func (*VolumeAttachmentList) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachmentList)(nil)).Elem()
}

func (i *VolumeAttachmentList) ToVolumeAttachmentListOutput() VolumeAttachmentListOutput {
	return i.ToVolumeAttachmentListOutputWithContext(context.Background())
}

func (i *VolumeAttachmentList) ToVolumeAttachmentListOutputWithContext(ctx context.Context) VolumeAttachmentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentListOutput)
}

func (i *VolumeAttachmentList) ToOutput(ctx context.Context) pulumix.Output[*VolumeAttachmentList] {
	return pulumix.Output[*VolumeAttachmentList]{
		OutputState: i.ToVolumeAttachmentListOutputWithContext(ctx).OutputState,
	}
}

// VolumeAttachmentListArrayInput is an input type that accepts VolumeAttachmentListArray and VolumeAttachmentListArrayOutput values.
// You can construct a concrete instance of `VolumeAttachmentListArrayInput` via:
//
//	VolumeAttachmentListArray{ VolumeAttachmentListArgs{...} }
type VolumeAttachmentListArrayInput interface {
	pulumi.Input

	ToVolumeAttachmentListArrayOutput() VolumeAttachmentListArrayOutput
	ToVolumeAttachmentListArrayOutputWithContext(context.Context) VolumeAttachmentListArrayOutput
}

type VolumeAttachmentListArray []VolumeAttachmentListInput

func (VolumeAttachmentListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttachmentList)(nil)).Elem()
}

func (i VolumeAttachmentListArray) ToVolumeAttachmentListArrayOutput() VolumeAttachmentListArrayOutput {
	return i.ToVolumeAttachmentListArrayOutputWithContext(context.Background())
}

func (i VolumeAttachmentListArray) ToVolumeAttachmentListArrayOutputWithContext(ctx context.Context) VolumeAttachmentListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentListArrayOutput)
}

func (i VolumeAttachmentListArray) ToOutput(ctx context.Context) pulumix.Output[[]*VolumeAttachmentList] {
	return pulumix.Output[[]*VolumeAttachmentList]{
		OutputState: i.ToVolumeAttachmentListArrayOutputWithContext(ctx).OutputState,
	}
}

// VolumeAttachmentListMapInput is an input type that accepts VolumeAttachmentListMap and VolumeAttachmentListMapOutput values.
// You can construct a concrete instance of `VolumeAttachmentListMapInput` via:
//
//	VolumeAttachmentListMap{ "key": VolumeAttachmentListArgs{...} }
type VolumeAttachmentListMapInput interface {
	pulumi.Input

	ToVolumeAttachmentListMapOutput() VolumeAttachmentListMapOutput
	ToVolumeAttachmentListMapOutputWithContext(context.Context) VolumeAttachmentListMapOutput
}

type VolumeAttachmentListMap map[string]VolumeAttachmentListInput

func (VolumeAttachmentListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttachmentList)(nil)).Elem()
}

func (i VolumeAttachmentListMap) ToVolumeAttachmentListMapOutput() VolumeAttachmentListMapOutput {
	return i.ToVolumeAttachmentListMapOutputWithContext(context.Background())
}

func (i VolumeAttachmentListMap) ToVolumeAttachmentListMapOutputWithContext(ctx context.Context) VolumeAttachmentListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentListMapOutput)
}

func (i VolumeAttachmentListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VolumeAttachmentList] {
	return pulumix.Output[map[string]*VolumeAttachmentList]{
		OutputState: i.ToVolumeAttachmentListMapOutputWithContext(ctx).OutputState,
	}
}

type VolumeAttachmentListOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachmentList)(nil)).Elem()
}

func (o VolumeAttachmentListOutput) ToVolumeAttachmentListOutput() VolumeAttachmentListOutput {
	return o
}

func (o VolumeAttachmentListOutput) ToVolumeAttachmentListOutputWithContext(ctx context.Context) VolumeAttachmentListOutput {
	return o
}

func (o VolumeAttachmentListOutput) ToOutput(ctx context.Context) pulumix.Output[*VolumeAttachmentList] {
	return pulumix.Output[*VolumeAttachmentList]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VolumeAttachmentListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachmentList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of VolumeAttachments
func (o VolumeAttachmentListOutput) Items() VolumeAttachmentTypeArrayOutput {
	return o.ApplyT(func(v *VolumeAttachmentList) VolumeAttachmentTypeArrayOutput { return v.Items }).(VolumeAttachmentTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VolumeAttachmentListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachmentList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VolumeAttachmentListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *VolumeAttachmentList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type VolumeAttachmentListArrayOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttachmentList)(nil)).Elem()
}

func (o VolumeAttachmentListArrayOutput) ToVolumeAttachmentListArrayOutput() VolumeAttachmentListArrayOutput {
	return o
}

func (o VolumeAttachmentListArrayOutput) ToVolumeAttachmentListArrayOutputWithContext(ctx context.Context) VolumeAttachmentListArrayOutput {
	return o
}

func (o VolumeAttachmentListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VolumeAttachmentList] {
	return pulumix.Output[[]*VolumeAttachmentList]{
		OutputState: o.OutputState,
	}
}

func (o VolumeAttachmentListArrayOutput) Index(i pulumi.IntInput) VolumeAttachmentListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeAttachmentList {
		return vs[0].([]*VolumeAttachmentList)[vs[1].(int)]
	}).(VolumeAttachmentListOutput)
}

type VolumeAttachmentListMapOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttachmentList)(nil)).Elem()
}

func (o VolumeAttachmentListMapOutput) ToVolumeAttachmentListMapOutput() VolumeAttachmentListMapOutput {
	return o
}

func (o VolumeAttachmentListMapOutput) ToVolumeAttachmentListMapOutputWithContext(ctx context.Context) VolumeAttachmentListMapOutput {
	return o
}

func (o VolumeAttachmentListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VolumeAttachmentList] {
	return pulumix.Output[map[string]*VolumeAttachmentList]{
		OutputState: o.OutputState,
	}
}

func (o VolumeAttachmentListMapOutput) MapIndex(k pulumi.StringInput) VolumeAttachmentListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeAttachmentList {
		return vs[0].(map[string]*VolumeAttachmentList)[vs[1].(string)]
	}).(VolumeAttachmentListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentListInput)(nil)).Elem(), &VolumeAttachmentList{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentListArrayInput)(nil)).Elem(), VolumeAttachmentListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentListMapInput)(nil)).Elem(), VolumeAttachmentListMap{})
	pulumi.RegisterOutputType(VolumeAttachmentListOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentListArrayOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentListMapOutput{})
}
