# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import meta as _meta

__all__ = [
    'CertificateSigningRequestConditionArgs',
    'CertificateSigningRequestSpecPatchArgs',
    'CertificateSigningRequestSpecArgs',
    'CertificateSigningRequestStatusArgs',
    'CertificateSigningRequestArgs',
]

@pulumi.input_type
class CertificateSigningRequestConditionArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 last_transition_time: Optional[pulumi.Input[str]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: request approval state, currently Approved or Denied.
        :param pulumi.Input[str] last_transition_time: lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
        :param pulumi.Input[str] last_update_time: timestamp for the last update to this condition
        :param pulumi.Input[str] message: human readable message with details about the request state
        :param pulumi.Input[str] reason: brief reason for the request state
        :param pulumi.Input[str] status: Status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown". Defaults to "True". If unset, should be treated as "True".
        """
        CertificateSigningRequestConditionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            type=type,
            last_transition_time=last_transition_time,
            last_update_time=last_update_time,
            message=message,
            reason=reason,
            status=status,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             type: pulumi.Input[str],
             last_transition_time: Optional[pulumi.Input[str]] = None,
             last_update_time: Optional[pulumi.Input[str]] = None,
             message: Optional[pulumi.Input[str]] = None,
             reason: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'lastTransitionTime' in kwargs:
            last_transition_time = kwargs['lastTransitionTime']
        if 'lastUpdateTime' in kwargs:
            last_update_time = kwargs['lastUpdateTime']

        _setter("type", type)
        if last_transition_time is not None:
            _setter("last_transition_time", last_transition_time)
        if last_update_time is not None:
            _setter("last_update_time", last_update_time)
        if message is not None:
            _setter("message", message)
        if reason is not None:
            _setter("reason", reason)
        if status is not None:
            _setter("status", status)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        request approval state, currently Approved or Denied.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="lastTransitionTime")
    def last_transition_time(self) -> Optional[pulumi.Input[str]]:
        """
        lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
        """
        return pulumi.get(self, "last_transition_time")

    @last_transition_time.setter
    def last_transition_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_transition_time", value)

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> Optional[pulumi.Input[str]]:
        """
        timestamp for the last update to this condition
        """
        return pulumi.get(self, "last_update_time")

    @last_update_time.setter
    def last_update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_update_time", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        human readable message with details about the request state
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def reason(self) -> Optional[pulumi.Input[str]]:
        """
        brief reason for the request state
        """
        return pulumi.get(self, "reason")

    @reason.setter
    def reason(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reason", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown". Defaults to "True". If unset, should be treated as "True".
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class CertificateSigningRequestSpecPatchArgs:
    def __init__(__self__, *,
                 extra: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 request: Optional[pulumi.Input[str]] = None,
                 signer_name: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
        :param pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]] extra: Extra information about the requesting user. See user.Info interface for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: Group information about the requesting user. See user.Info interface for details.
        :param pulumi.Input[str] request: Base64-encoded PKCS#10 CSR data
        :param pulumi.Input[str] signer_name: Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
                1. If it's a kubelet client certificate, it is assigned
                   "kubernetes.io/kube-apiserver-client-kubelet".
                2. If it's a kubelet serving certificate, it is assigned
                   "kubernetes.io/kubelet-serving".
                3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
               Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
        :param pulumi.Input[str] uid: UID information about the requesting user. See user.Info interface for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] usages: allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
                    https://tools.ietf.org/html/rfc5280#section-4.2.1.12
        :param pulumi.Input[str] username: Information about the requesting user. See user.Info interface for details.
        """
        CertificateSigningRequestSpecPatchArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            extra=extra,
            groups=groups,
            request=request,
            signer_name=signer_name,
            uid=uid,
            usages=usages,
            username=username,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             extra: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]] = None,
             groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             request: Optional[pulumi.Input[str]] = None,
             signer_name: Optional[pulumi.Input[str]] = None,
             uid: Optional[pulumi.Input[str]] = None,
             usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             username: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'signerName' in kwargs:
            signer_name = kwargs['signerName']

        if extra is not None:
            _setter("extra", extra)
        if groups is not None:
            _setter("groups", groups)
        if request is not None:
            _setter("request", request)
        if signer_name is not None:
            _setter("signer_name", signer_name)
        if uid is not None:
            _setter("uid", uid)
        if usages is not None:
            _setter("usages", usages)
        if username is not None:
            _setter("username", username)

    @property
    @pulumi.getter
    def extra(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]:
        """
        Extra information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "extra")

    @extra.setter
    def extra(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]):
        pulumi.set(self, "extra", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Group information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def request(self) -> Optional[pulumi.Input[str]]:
        """
        Base64-encoded PKCS#10 CSR data
        """
        return pulumi.get(self, "request")

    @request.setter
    def request(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request", value)

    @property
    @pulumi.getter(name="signerName")
    def signer_name(self) -> Optional[pulumi.Input[str]]:
        """
        Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
         1. If it's a kubelet client certificate, it is assigned
            "kubernetes.io/kube-apiserver-client-kubelet".
         2. If it's a kubelet serving certificate, it is assigned
            "kubernetes.io/kubelet-serving".
         3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
        Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
        """
        return pulumi.get(self, "signer_name")

    @signer_name.setter
    def signer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signer_name", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        UID information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter
    def usages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
             https://tools.ietf.org/html/rfc5280#section-4.2.1.12
        """
        return pulumi.get(self, "usages")

    @usages.setter
    def usages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "usages", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class CertificateSigningRequestSpecArgs:
    def __init__(__self__, *,
                 request: pulumi.Input[str],
                 extra: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signer_name: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
        :param pulumi.Input[str] request: Base64-encoded PKCS#10 CSR data
        :param pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]] extra: Extra information about the requesting user. See user.Info interface for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: Group information about the requesting user. See user.Info interface for details.
        :param pulumi.Input[str] signer_name: Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
                1. If it's a kubelet client certificate, it is assigned
                   "kubernetes.io/kube-apiserver-client-kubelet".
                2. If it's a kubelet serving certificate, it is assigned
                   "kubernetes.io/kubelet-serving".
                3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
               Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
        :param pulumi.Input[str] uid: UID information about the requesting user. See user.Info interface for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] usages: allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
                    https://tools.ietf.org/html/rfc5280#section-4.2.1.12
        :param pulumi.Input[str] username: Information about the requesting user. See user.Info interface for details.
        """
        CertificateSigningRequestSpecArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            request=request,
            extra=extra,
            groups=groups,
            signer_name=signer_name,
            uid=uid,
            usages=usages,
            username=username,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             request: pulumi.Input[str],
             extra: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]] = None,
             groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             signer_name: Optional[pulumi.Input[str]] = None,
             uid: Optional[pulumi.Input[str]] = None,
             usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             username: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'signerName' in kwargs:
            signer_name = kwargs['signerName']

        _setter("request", request)
        if extra is not None:
            _setter("extra", extra)
        if groups is not None:
            _setter("groups", groups)
        if signer_name is not None:
            _setter("signer_name", signer_name)
        if uid is not None:
            _setter("uid", uid)
        if usages is not None:
            _setter("usages", usages)
        if username is not None:
            _setter("username", username)

    @property
    @pulumi.getter
    def request(self) -> pulumi.Input[str]:
        """
        Base64-encoded PKCS#10 CSR data
        """
        return pulumi.get(self, "request")

    @request.setter
    def request(self, value: pulumi.Input[str]):
        pulumi.set(self, "request", value)

    @property
    @pulumi.getter
    def extra(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]:
        """
        Extra information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "extra")

    @extra.setter
    def extra(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]):
        pulumi.set(self, "extra", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Group information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="signerName")
    def signer_name(self) -> Optional[pulumi.Input[str]]:
        """
        Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
         1. If it's a kubelet client certificate, it is assigned
            "kubernetes.io/kube-apiserver-client-kubelet".
         2. If it's a kubelet serving certificate, it is assigned
            "kubernetes.io/kubelet-serving".
         3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
        Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
        """
        return pulumi.get(self, "signer_name")

    @signer_name.setter
    def signer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signer_name", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        UID information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter
    def usages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
             https://tools.ietf.org/html/rfc5280#section-4.2.1.12
        """
        return pulumi.get(self, "usages")

    @usages.setter
    def usages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "usages", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Information about the requesting user. See user.Info interface for details.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class CertificateSigningRequestStatusArgs:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateSigningRequestConditionArgs']]]] = None):
        """
        :param pulumi.Input[str] certificate: If request was approved, the controller will place the issued certificate here.
        :param pulumi.Input[Sequence[pulumi.Input['CertificateSigningRequestConditionArgs']]] conditions: Conditions applied to the request, such as approval or denial.
        """
        CertificateSigningRequestStatusArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            certificate=certificate,
            conditions=conditions,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             certificate: Optional[pulumi.Input[str]] = None,
             conditions: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateSigningRequestConditionArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if certificate is not None:
            _setter("certificate", certificate)
        if conditions is not None:
            _setter("conditions", conditions)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        If request was approved, the controller will place the issued certificate here.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CertificateSigningRequestConditionArgs']]]]:
        """
        Conditions applied to the request, such as approval or denial.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateSigningRequestConditionArgs']]]]):
        pulumi.set(self, "conditions", value)


@pulumi.input_type
class CertificateSigningRequestArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['CertificateSigningRequestSpecArgs']] = None,
                 status: Optional[pulumi.Input['CertificateSigningRequestStatusArgs']] = None):
        """
        Describes a certificate signing request
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['CertificateSigningRequestSpecArgs'] spec: The certificate request itself and any additional information.
        :param pulumi.Input['CertificateSigningRequestStatusArgs'] status: Derived information about the request.
        """
        CertificateSigningRequestArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_version=api_version,
            kind=kind,
            metadata=metadata,
            spec=spec,
            status=status,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_version: Optional[pulumi.Input[str]] = None,
             kind: Optional[pulumi.Input[str]] = None,
             metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
             spec: Optional[pulumi.Input['CertificateSigningRequestSpecArgs']] = None,
             status: Optional[pulumi.Input['CertificateSigningRequestStatusArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']

        if api_version is not None:
            _setter("api_version", 'certificates.k8s.io/v1beta1')
        if kind is not None:
            _setter("kind", 'CertificateSigningRequest')
        if metadata is not None:
            _setter("metadata", metadata)
        if spec is not None:
            _setter("spec", spec)
        if status is not None:
            _setter("status", status)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['CertificateSigningRequestSpecArgs']]:
        """
        The certificate request itself and any additional information.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['CertificateSigningRequestSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['CertificateSigningRequestStatusArgs']]:
        """
        Derived information about the request.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['CertificateSigningRequestStatusArgs']]):
        pulumi.set(self, "status", value)


