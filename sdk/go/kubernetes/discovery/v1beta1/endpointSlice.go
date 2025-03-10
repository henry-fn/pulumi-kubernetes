// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
type EndpointSlice struct {
	pulumi.CustomResourceState

	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	AddressType pulumi.StringOutput `pulumi:"addressType"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	Endpoints EndpointArrayOutput `pulumi:"endpoints"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports EndpointPortArrayOutput `pulumi:"ports"`
}

// NewEndpointSlice registers a new resource with the given unique name, arguments, and options.
func NewEndpointSlice(ctx *pulumi.Context,
	name string, args *EndpointSliceArgs, opts ...pulumi.ResourceOption) (*EndpointSlice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressType == nil {
		return nil, errors.New("invalid value for required argument 'AddressType'")
	}
	if args.Endpoints == nil {
		return nil, errors.New("invalid value for required argument 'Endpoints'")
	}
	args.ApiVersion = pulumi.StringPtr("discovery.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("EndpointSlice")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:discovery.k8s.io/v1:EndpointSlice"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EndpointSlice
	err := ctx.RegisterResource("kubernetes:discovery.k8s.io/v1beta1:EndpointSlice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointSlice gets an existing EndpointSlice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointSlice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointSliceState, opts ...pulumi.ResourceOption) (*EndpointSlice, error) {
	var resource EndpointSlice
	err := ctx.ReadResource("kubernetes:discovery.k8s.io/v1beta1:EndpointSlice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointSlice resources.
type endpointSliceState struct {
}

type EndpointSliceState struct {
}

func (EndpointSliceState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointSliceState)(nil)).Elem()
}

type endpointSliceArgs struct {
	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	AddressType string `pulumi:"addressType"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	Endpoints []Endpoint `pulumi:"endpoints"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports []EndpointPort `pulumi:"ports"`
}

// The set of arguments for constructing a EndpointSlice resource.
type EndpointSliceArgs struct {
	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	AddressType pulumi.StringInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	Endpoints EndpointArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports EndpointPortArrayInput
}

func (EndpointSliceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointSliceArgs)(nil)).Elem()
}

type EndpointSliceInput interface {
	pulumi.Input

	ToEndpointSliceOutput() EndpointSliceOutput
	ToEndpointSliceOutputWithContext(ctx context.Context) EndpointSliceOutput
}

func (*EndpointSlice) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointSlice)(nil)).Elem()
}

func (i *EndpointSlice) ToEndpointSliceOutput() EndpointSliceOutput {
	return i.ToEndpointSliceOutputWithContext(context.Background())
}

func (i *EndpointSlice) ToEndpointSliceOutputWithContext(ctx context.Context) EndpointSliceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointSliceOutput)
}

func (i *EndpointSlice) ToOutput(ctx context.Context) pulumix.Output[*EndpointSlice] {
	return pulumix.Output[*EndpointSlice]{
		OutputState: i.ToEndpointSliceOutputWithContext(ctx).OutputState,
	}
}

// EndpointSliceArrayInput is an input type that accepts EndpointSliceArray and EndpointSliceArrayOutput values.
// You can construct a concrete instance of `EndpointSliceArrayInput` via:
//
//	EndpointSliceArray{ EndpointSliceArgs{...} }
type EndpointSliceArrayInput interface {
	pulumi.Input

	ToEndpointSliceArrayOutput() EndpointSliceArrayOutput
	ToEndpointSliceArrayOutputWithContext(context.Context) EndpointSliceArrayOutput
}

type EndpointSliceArray []EndpointSliceInput

func (EndpointSliceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointSlice)(nil)).Elem()
}

func (i EndpointSliceArray) ToEndpointSliceArrayOutput() EndpointSliceArrayOutput {
	return i.ToEndpointSliceArrayOutputWithContext(context.Background())
}

func (i EndpointSliceArray) ToEndpointSliceArrayOutputWithContext(ctx context.Context) EndpointSliceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointSliceArrayOutput)
}

func (i EndpointSliceArray) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointSlice] {
	return pulumix.Output[[]*EndpointSlice]{
		OutputState: i.ToEndpointSliceArrayOutputWithContext(ctx).OutputState,
	}
}

// EndpointSliceMapInput is an input type that accepts EndpointSliceMap and EndpointSliceMapOutput values.
// You can construct a concrete instance of `EndpointSliceMapInput` via:
//
//	EndpointSliceMap{ "key": EndpointSliceArgs{...} }
type EndpointSliceMapInput interface {
	pulumi.Input

	ToEndpointSliceMapOutput() EndpointSliceMapOutput
	ToEndpointSliceMapOutputWithContext(context.Context) EndpointSliceMapOutput
}

type EndpointSliceMap map[string]EndpointSliceInput

func (EndpointSliceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointSlice)(nil)).Elem()
}

func (i EndpointSliceMap) ToEndpointSliceMapOutput() EndpointSliceMapOutput {
	return i.ToEndpointSliceMapOutputWithContext(context.Background())
}

func (i EndpointSliceMap) ToEndpointSliceMapOutputWithContext(ctx context.Context) EndpointSliceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointSliceMapOutput)
}

func (i EndpointSliceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointSlice] {
	return pulumix.Output[map[string]*EndpointSlice]{
		OutputState: i.ToEndpointSliceMapOutputWithContext(ctx).OutputState,
	}
}

type EndpointSliceOutput struct{ *pulumi.OutputState }

func (EndpointSliceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointSlice)(nil)).Elem()
}

func (o EndpointSliceOutput) ToEndpointSliceOutput() EndpointSliceOutput {
	return o
}

func (o EndpointSliceOutput) ToEndpointSliceOutputWithContext(ctx context.Context) EndpointSliceOutput {
	return o
}

func (o EndpointSliceOutput) ToOutput(ctx context.Context) pulumix.Output[*EndpointSlice] {
	return pulumix.Output[*EndpointSlice]{
		OutputState: o.OutputState,
	}
}

// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
func (o EndpointSliceOutput) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointSlice) pulumi.StringOutput { return v.AddressType }).(pulumi.StringOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EndpointSliceOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointSlice) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
func (o EndpointSliceOutput) Endpoints() EndpointArrayOutput {
	return o.ApplyT(func(v *EndpointSlice) EndpointArrayOutput { return v.Endpoints }).(EndpointArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EndpointSliceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointSlice) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata.
func (o EndpointSliceOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *EndpointSlice) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
func (o EndpointSliceOutput) Ports() EndpointPortArrayOutput {
	return o.ApplyT(func(v *EndpointSlice) EndpointPortArrayOutput { return v.Ports }).(EndpointPortArrayOutput)
}

type EndpointSliceArrayOutput struct{ *pulumi.OutputState }

func (EndpointSliceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointSlice)(nil)).Elem()
}

func (o EndpointSliceArrayOutput) ToEndpointSliceArrayOutput() EndpointSliceArrayOutput {
	return o
}

func (o EndpointSliceArrayOutput) ToEndpointSliceArrayOutputWithContext(ctx context.Context) EndpointSliceArrayOutput {
	return o
}

func (o EndpointSliceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointSlice] {
	return pulumix.Output[[]*EndpointSlice]{
		OutputState: o.OutputState,
	}
}

func (o EndpointSliceArrayOutput) Index(i pulumi.IntInput) EndpointSliceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointSlice {
		return vs[0].([]*EndpointSlice)[vs[1].(int)]
	}).(EndpointSliceOutput)
}

type EndpointSliceMapOutput struct{ *pulumi.OutputState }

func (EndpointSliceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointSlice)(nil)).Elem()
}

func (o EndpointSliceMapOutput) ToEndpointSliceMapOutput() EndpointSliceMapOutput {
	return o
}

func (o EndpointSliceMapOutput) ToEndpointSliceMapOutputWithContext(ctx context.Context) EndpointSliceMapOutput {
	return o
}

func (o EndpointSliceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointSlice] {
	return pulumix.Output[map[string]*EndpointSlice]{
		OutputState: o.OutputState,
	}
}

func (o EndpointSliceMapOutput) MapIndex(k pulumi.StringInput) EndpointSliceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointSlice {
		return vs[0].(map[string]*EndpointSlice)[vs[1].(string)]
	}).(EndpointSliceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointSliceInput)(nil)).Elem(), &EndpointSlice{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointSliceArrayInput)(nil)).Elem(), EndpointSliceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointSliceMapInput)(nil)).Elem(), EndpointSliceMap{})
	pulumi.RegisterOutputType(EndpointSliceOutput{})
	pulumi.RegisterOutputType(EndpointSliceArrayOutput{})
	pulumi.RegisterOutputType(EndpointSliceMapOutput{})
}
