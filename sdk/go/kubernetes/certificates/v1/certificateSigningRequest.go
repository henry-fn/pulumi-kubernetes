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

// CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.
//
// Kubelets use this API to obtain:
//  1. client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).
//  2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving" signerName).
//
// This API can be used to request client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.
type CertificateSigningRequest struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// spec contains the certificate request, and is immutable after creation. Only the request, signerName, expirationSeconds, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
	Spec CertificateSigningRequestSpecOutput `pulumi:"spec"`
	// status contains information about whether the request is approved or denied, and the certificate issued by the signer, or the failure condition indicating signer failure.
	Status CertificateSigningRequestStatusPtrOutput `pulumi:"status"`
}

// NewCertificateSigningRequest registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequest(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CertificateSigningRequest")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CertificateSigningRequest
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1:CertificateSigningRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequest gets an existing CertificateSigningRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestState, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	var resource CertificateSigningRequest
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1:CertificateSigningRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequest resources.
type certificateSigningRequestState struct {
}

type CertificateSigningRequestState struct {
}

func (CertificateSigningRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestState)(nil)).Elem()
}

type certificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec contains the certificate request, and is immutable after creation. Only the request, signerName, expirationSeconds, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
	Spec CertificateSigningRequestSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CertificateSigningRequest resource.
type CertificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// spec contains the certificate request, and is immutable after creation. Only the request, signerName, expirationSeconds, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
	Spec CertificateSigningRequestSpecInput
}

func (CertificateSigningRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestArgs)(nil)).Elem()
}

type CertificateSigningRequestInput interface {
	pulumi.Input

	ToCertificateSigningRequestOutput() CertificateSigningRequestOutput
	ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput
}

func (*CertificateSigningRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequest)(nil)).Elem()
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestOutput() CertificateSigningRequestOutput {
	return i.ToCertificateSigningRequestOutputWithContext(context.Background())
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestOutput)
}

func (i *CertificateSigningRequest) ToOutput(ctx context.Context) pulumix.Output[*CertificateSigningRequest] {
	return pulumix.Output[*CertificateSigningRequest]{
		OutputState: i.ToCertificateSigningRequestOutputWithContext(ctx).OutputState,
	}
}

// CertificateSigningRequestArrayInput is an input type that accepts CertificateSigningRequestArray and CertificateSigningRequestArrayOutput values.
// You can construct a concrete instance of `CertificateSigningRequestArrayInput` via:
//
//	CertificateSigningRequestArray{ CertificateSigningRequestArgs{...} }
type CertificateSigningRequestArrayInput interface {
	pulumi.Input

	ToCertificateSigningRequestArrayOutput() CertificateSigningRequestArrayOutput
	ToCertificateSigningRequestArrayOutputWithContext(context.Context) CertificateSigningRequestArrayOutput
}

type CertificateSigningRequestArray []CertificateSigningRequestInput

func (CertificateSigningRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequest)(nil)).Elem()
}

func (i CertificateSigningRequestArray) ToCertificateSigningRequestArrayOutput() CertificateSigningRequestArrayOutput {
	return i.ToCertificateSigningRequestArrayOutputWithContext(context.Background())
}

func (i CertificateSigningRequestArray) ToCertificateSigningRequestArrayOutputWithContext(ctx context.Context) CertificateSigningRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestArrayOutput)
}

func (i CertificateSigningRequestArray) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateSigningRequest] {
	return pulumix.Output[[]*CertificateSigningRequest]{
		OutputState: i.ToCertificateSigningRequestArrayOutputWithContext(ctx).OutputState,
	}
}

// CertificateSigningRequestMapInput is an input type that accepts CertificateSigningRequestMap and CertificateSigningRequestMapOutput values.
// You can construct a concrete instance of `CertificateSigningRequestMapInput` via:
//
//	CertificateSigningRequestMap{ "key": CertificateSigningRequestArgs{...} }
type CertificateSigningRequestMapInput interface {
	pulumi.Input

	ToCertificateSigningRequestMapOutput() CertificateSigningRequestMapOutput
	ToCertificateSigningRequestMapOutputWithContext(context.Context) CertificateSigningRequestMapOutput
}

type CertificateSigningRequestMap map[string]CertificateSigningRequestInput

func (CertificateSigningRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequest)(nil)).Elem()
}

func (i CertificateSigningRequestMap) ToCertificateSigningRequestMapOutput() CertificateSigningRequestMapOutput {
	return i.ToCertificateSigningRequestMapOutputWithContext(context.Background())
}

func (i CertificateSigningRequestMap) ToCertificateSigningRequestMapOutputWithContext(ctx context.Context) CertificateSigningRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestMapOutput)
}

func (i CertificateSigningRequestMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateSigningRequest] {
	return pulumix.Output[map[string]*CertificateSigningRequest]{
		OutputState: i.ToCertificateSigningRequestMapOutputWithContext(ctx).OutputState,
	}
}

type CertificateSigningRequestOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequest)(nil)).Elem()
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestOutput() CertificateSigningRequestOutput {
	return o
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput {
	return o
}

func (o CertificateSigningRequestOutput) ToOutput(ctx context.Context) pulumix.Output[*CertificateSigningRequest] {
	return pulumix.Output[*CertificateSigningRequest]{
		OutputState: o.OutputState,
	}
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CertificateSigningRequestOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateSigningRequest) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CertificateSigningRequestOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateSigningRequest) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o CertificateSigningRequestOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *CertificateSigningRequest) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// spec contains the certificate request, and is immutable after creation. Only the request, signerName, expirationSeconds, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
func (o CertificateSigningRequestOutput) Spec() CertificateSigningRequestSpecOutput {
	return o.ApplyT(func(v *CertificateSigningRequest) CertificateSigningRequestSpecOutput { return v.Spec }).(CertificateSigningRequestSpecOutput)
}

// status contains information about whether the request is approved or denied, and the certificate issued by the signer, or the failure condition indicating signer failure.
func (o CertificateSigningRequestOutput) Status() CertificateSigningRequestStatusPtrOutput {
	return o.ApplyT(func(v *CertificateSigningRequest) CertificateSigningRequestStatusPtrOutput { return v.Status }).(CertificateSigningRequestStatusPtrOutput)
}

type CertificateSigningRequestArrayOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequest)(nil)).Elem()
}

func (o CertificateSigningRequestArrayOutput) ToCertificateSigningRequestArrayOutput() CertificateSigningRequestArrayOutput {
	return o
}

func (o CertificateSigningRequestArrayOutput) ToCertificateSigningRequestArrayOutputWithContext(ctx context.Context) CertificateSigningRequestArrayOutput {
	return o
}

func (o CertificateSigningRequestArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateSigningRequest] {
	return pulumix.Output[[]*CertificateSigningRequest]{
		OutputState: o.OutputState,
	}
}

func (o CertificateSigningRequestArrayOutput) Index(i pulumi.IntInput) CertificateSigningRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateSigningRequest {
		return vs[0].([]*CertificateSigningRequest)[vs[1].(int)]
	}).(CertificateSigningRequestOutput)
}

type CertificateSigningRequestMapOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequest)(nil)).Elem()
}

func (o CertificateSigningRequestMapOutput) ToCertificateSigningRequestMapOutput() CertificateSigningRequestMapOutput {
	return o
}

func (o CertificateSigningRequestMapOutput) ToCertificateSigningRequestMapOutputWithContext(ctx context.Context) CertificateSigningRequestMapOutput {
	return o
}

func (o CertificateSigningRequestMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateSigningRequest] {
	return pulumix.Output[map[string]*CertificateSigningRequest]{
		OutputState: o.OutputState,
	}
}

func (o CertificateSigningRequestMapOutput) MapIndex(k pulumi.StringInput) CertificateSigningRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateSigningRequest {
		return vs[0].(map[string]*CertificateSigningRequest)[vs[1].(string)]
	}).(CertificateSigningRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestInput)(nil)).Elem(), &CertificateSigningRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestArrayInput)(nil)).Elem(), CertificateSigningRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestMapInput)(nil)).Elem(), CertificateSigningRequestMap{})
	pulumi.RegisterOutputType(CertificateSigningRequestOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestArrayOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestMapOutput{})
}
