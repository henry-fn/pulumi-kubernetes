# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'ClusterCIDR',
    'ClusterCIDRSpec',
    'ClusterCIDRSpecPatch',
    'IPAddress',
    'IPAddressSpec',
    'IPAddressSpecPatch',
    'ParentReference',
    'ParentReferencePatch',
]

@pulumi.output_type
class ClusterCIDR(dict):
    """
    ClusterCIDR represents a single configuration for per-Node Pod CIDR allocations when the MultiCIDRRangeAllocator is enabled (see the config for kube-controller-manager).  A cluster may have any number of ClusterCIDR resources, all of which will be considered when allocating a CIDR for a Node.  A ClusterCIDR is eligible to be used for a given Node when the node selector matches the node in question and has free CIDRs to allocate.  In case of multiple matching ClusterCIDR resources, the allocator will attempt to break ties using internal heuristics, but any ClusterCIDR whose node selector matches the Node may be used.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterCIDR. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterCIDR.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterCIDR.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 spec: Optional['outputs.ClusterCIDRSpec'] = None):
        """
        ClusterCIDR represents a single configuration for per-Node Pod CIDR allocations when the MultiCIDRRangeAllocator is enabled (see the config for kube-controller-manager).  A cluster may have any number of ClusterCIDR resources, all of which will be considered when allocating a CIDR for a Node.  A ClusterCIDR is eligible to be used for a given Node when the node selector matches the node in question and has free CIDRs to allocate.  In case of multiple matching ClusterCIDR resources, the allocator will attempt to break ties using internal heuristics, but any ClusterCIDR whose node selector matches the Node may be used.
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param '_meta.v1.ObjectMetaArgs' metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param 'ClusterCIDRSpecArgs' spec: spec is the desired state of the ClusterCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        ClusterCIDR._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_version=api_version,
            kind=kind,
            metadata=metadata,
            spec=spec,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_version: Optional[str] = None,
             kind: Optional[str] = None,
             metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
             spec: Optional['outputs.ClusterCIDRSpec'] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']

        if api_version is not None:
            _setter("api_version", 'networking.k8s.io/v1alpha1')
        if kind is not None:
            _setter("kind", 'ClusterCIDR')
        if metadata is not None:
            _setter("metadata", metadata)
        if spec is not None:
            _setter("spec", spec)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> Optional['outputs.ClusterCIDRSpec']:
        """
        spec is the desired state of the ClusterCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "spec")


@pulumi.output_type
class ClusterCIDRSpec(dict):
    """
    ClusterCIDRSpec defines the desired state of ClusterCIDR.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "perNodeHostBits":
            suggest = "per_node_host_bits"
        elif key == "nodeSelector":
            suggest = "node_selector"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterCIDRSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterCIDRSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterCIDRSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 per_node_host_bits: int,
                 ipv4: Optional[str] = None,
                 ipv6: Optional[str] = None,
                 node_selector: Optional['_core.v1.outputs.NodeSelector'] = None):
        """
        ClusterCIDRSpec defines the desired state of ClusterCIDR.
        :param int per_node_host_bits: perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        :param str ipv4: ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param str ipv6: ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param '_core.v1.NodeSelectorArgs' node_selector: nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        """
        ClusterCIDRSpec._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            per_node_host_bits=per_node_host_bits,
            ipv4=ipv4,
            ipv6=ipv6,
            node_selector=node_selector,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             per_node_host_bits: int,
             ipv4: Optional[str] = None,
             ipv6: Optional[str] = None,
             node_selector: Optional['_core.v1.outputs.NodeSelector'] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'perNodeHostBits' in kwargs:
            per_node_host_bits = kwargs['perNodeHostBits']
        if 'nodeSelector' in kwargs:
            node_selector = kwargs['nodeSelector']

        _setter("per_node_host_bits", per_node_host_bits)
        if ipv4 is not None:
            _setter("ipv4", ipv4)
        if ipv6 is not None:
            _setter("ipv6", ipv6)
        if node_selector is not None:
            _setter("node_selector", node_selector)

    @property
    @pulumi.getter(name="perNodeHostBits")
    def per_node_host_bits(self) -> int:
        """
        perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        """
        return pulumi.get(self, "per_node_host_bits")

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[str]:
        """
        ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[str]:
        """
        ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="nodeSelector")
    def node_selector(self) -> Optional['_core.v1.outputs.NodeSelector']:
        """
        nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        """
        return pulumi.get(self, "node_selector")


@pulumi.output_type
class ClusterCIDRSpecPatch(dict):
    """
    ClusterCIDRSpec defines the desired state of ClusterCIDR.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nodeSelector":
            suggest = "node_selector"
        elif key == "perNodeHostBits":
            suggest = "per_node_host_bits"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterCIDRSpecPatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterCIDRSpecPatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterCIDRSpecPatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ipv4: Optional[str] = None,
                 ipv6: Optional[str] = None,
                 node_selector: Optional['_core.v1.outputs.NodeSelectorPatch'] = None,
                 per_node_host_bits: Optional[int] = None):
        """
        ClusterCIDRSpec defines the desired state of ClusterCIDR.
        :param str ipv4: ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param str ipv6: ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        :param '_core.v1.NodeSelectorPatchArgs' node_selector: nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        :param int per_node_host_bits: perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        """
        ClusterCIDRSpecPatch._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ipv4=ipv4,
            ipv6=ipv6,
            node_selector=node_selector,
            per_node_host_bits=per_node_host_bits,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ipv4: Optional[str] = None,
             ipv6: Optional[str] = None,
             node_selector: Optional['_core.v1.outputs.NodeSelectorPatch'] = None,
             per_node_host_bits: Optional[int] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'nodeSelector' in kwargs:
            node_selector = kwargs['nodeSelector']
        if 'perNodeHostBits' in kwargs:
            per_node_host_bits = kwargs['perNodeHostBits']

        if ipv4 is not None:
            _setter("ipv4", ipv4)
        if ipv6 is not None:
            _setter("ipv6", ipv6)
        if node_selector is not None:
            _setter("node_selector", node_selector)
        if per_node_host_bits is not None:
            _setter("per_node_host_bits", per_node_host_bits)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[str]:
        """
        ipv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[str]:
        """
        ipv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of ipv4 and ipv6 must be specified. This field is immutable.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="nodeSelector")
    def node_selector(self) -> Optional['_core.v1.outputs.NodeSelectorPatch']:
        """
        nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.
        """
        return pulumi.get(self, "node_selector")

    @property
    @pulumi.getter(name="perNodeHostBits")
    def per_node_host_bits(self) -> Optional[int]:
        """
        perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
        """
        return pulumi.get(self, "per_node_host_bits")


@pulumi.output_type
class IPAddress(dict):
    """
    IPAddress represents a single IP of a single IP Family. The object is designed to be used by APIs that operate on IP addresses. The object is used by the Service core API for allocation of IP addresses. An IP address can be represented in different formats, to guarantee the uniqueness of the IP, the name of the object is the IP address in canonical format, four decimal digits separated by dots suppressing leading zeros for IPv4 and the representation defined by RFC 5952 for IPv6. Valid: 192.168.1.5 or 2001:db8::1 or 2001:db8:aaaa:bbbb:cccc:dddd:eeee:1 Invalid: 10.01.2.3 or 2001:db8:0:0:0::1
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IPAddress. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IPAddress.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IPAddress.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 spec: Optional['outputs.IPAddressSpec'] = None):
        """
        IPAddress represents a single IP of a single IP Family. The object is designed to be used by APIs that operate on IP addresses. The object is used by the Service core API for allocation of IP addresses. An IP address can be represented in different formats, to guarantee the uniqueness of the IP, the name of the object is the IP address in canonical format, four decimal digits separated by dots suppressing leading zeros for IPv4 and the representation defined by RFC 5952 for IPv6. Valid: 192.168.1.5 or 2001:db8::1 or 2001:db8:aaaa:bbbb:cccc:dddd:eeee:1 Invalid: 10.01.2.3 or 2001:db8:0:0:0::1
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param '_meta.v1.ObjectMetaArgs' metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param 'IPAddressSpecArgs' spec: spec is the desired state of the IPAddress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        IPAddress._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_version=api_version,
            kind=kind,
            metadata=metadata,
            spec=spec,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_version: Optional[str] = None,
             kind: Optional[str] = None,
             metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
             spec: Optional['outputs.IPAddressSpec'] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']

        if api_version is not None:
            _setter("api_version", 'networking.k8s.io/v1alpha1')
        if kind is not None:
            _setter("kind", 'IPAddress')
        if metadata is not None:
            _setter("metadata", metadata)
        if spec is not None:
            _setter("spec", spec)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> Optional['outputs.IPAddressSpec']:
        """
        spec is the desired state of the IPAddress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "spec")


@pulumi.output_type
class IPAddressSpec(dict):
    """
    IPAddressSpec describe the attributes in an IP Address.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "parentRef":
            suggest = "parent_ref"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IPAddressSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IPAddressSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IPAddressSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 parent_ref: Optional['outputs.ParentReference'] = None):
        """
        IPAddressSpec describe the attributes in an IP Address.
        :param 'ParentReferenceArgs' parent_ref: ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        IPAddressSpec._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            parent_ref=parent_ref,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             parent_ref: Optional['outputs.ParentReference'] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'parentRef' in kwargs:
            parent_ref = kwargs['parentRef']

        if parent_ref is not None:
            _setter("parent_ref", parent_ref)

    @property
    @pulumi.getter(name="parentRef")
    def parent_ref(self) -> Optional['outputs.ParentReference']:
        """
        ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        return pulumi.get(self, "parent_ref")


@pulumi.output_type
class IPAddressSpecPatch(dict):
    """
    IPAddressSpec describe the attributes in an IP Address.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "parentRef":
            suggest = "parent_ref"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IPAddressSpecPatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IPAddressSpecPatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IPAddressSpecPatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 parent_ref: Optional['outputs.ParentReferencePatch'] = None):
        """
        IPAddressSpec describe the attributes in an IP Address.
        :param 'ParentReferencePatchArgs' parent_ref: ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        IPAddressSpecPatch._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            parent_ref=parent_ref,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             parent_ref: Optional['outputs.ParentReferencePatch'] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'parentRef' in kwargs:
            parent_ref = kwargs['parentRef']

        if parent_ref is not None:
            _setter("parent_ref", parent_ref)

    @property
    @pulumi.getter(name="parentRef")
    def parent_ref(self) -> Optional['outputs.ParentReferencePatch']:
        """
        ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
        """
        return pulumi.get(self, "parent_ref")


@pulumi.output_type
class ParentReference(dict):
    """
    ParentReference describes a reference to a parent object.
    """
    def __init__(__self__, *,
                 group: Optional[str] = None,
                 name: Optional[str] = None,
                 namespace: Optional[str] = None,
                 resource: Optional[str] = None,
                 uid: Optional[str] = None):
        """
        ParentReference describes a reference to a parent object.
        :param str group: Group is the group of the object being referenced.
        :param str name: Name is the name of the object being referenced.
        :param str namespace: Namespace is the namespace of the object being referenced.
        :param str resource: Resource is the resource of the object being referenced.
        :param str uid: UID is the uid of the object being referenced.
        """
        ParentReference._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            name=name,
            namespace=namespace,
            resource=resource,
            uid=uid,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[str] = None,
             name: Optional[str] = None,
             namespace: Optional[str] = None,
             resource: Optional[str] = None,
             uid: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if group is not None:
            _setter("group", group)
        if name is not None:
            _setter("name", name)
        if namespace is not None:
            _setter("namespace", namespace)
        if resource is not None:
            _setter("resource", resource)
        if uid is not None:
            _setter("uid", uid)

    @property
    @pulumi.getter
    def group(self) -> Optional[str]:
        """
        Group is the group of the object being referenced.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name is the name of the object being referenced.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        Namespace is the namespace of the object being referenced.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def resource(self) -> Optional[str]:
        """
        Resource is the resource of the object being referenced.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        UID is the uid of the object being referenced.
        """
        return pulumi.get(self, "uid")


@pulumi.output_type
class ParentReferencePatch(dict):
    """
    ParentReference describes a reference to a parent object.
    """
    def __init__(__self__, *,
                 group: Optional[str] = None,
                 name: Optional[str] = None,
                 namespace: Optional[str] = None,
                 resource: Optional[str] = None,
                 uid: Optional[str] = None):
        """
        ParentReference describes a reference to a parent object.
        :param str group: Group is the group of the object being referenced.
        :param str name: Name is the name of the object being referenced.
        :param str namespace: Namespace is the namespace of the object being referenced.
        :param str resource: Resource is the resource of the object being referenced.
        :param str uid: UID is the uid of the object being referenced.
        """
        ParentReferencePatch._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            name=name,
            namespace=namespace,
            resource=resource,
            uid=uid,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[str] = None,
             name: Optional[str] = None,
             namespace: Optional[str] = None,
             resource: Optional[str] = None,
             uid: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if group is not None:
            _setter("group", group)
        if name is not None:
            _setter("name", name)
        if namespace is not None:
            _setter("namespace", namespace)
        if resource is not None:
            _setter("resource", resource)
        if uid is not None:
            _setter("uid", uid)

    @property
    @pulumi.getter
    def group(self) -> Optional[str]:
        """
        Group is the group of the object being referenced.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name is the name of the object being referenced.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        Namespace is the namespace of the object being referenced.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def resource(self) -> Optional[str]:
        """
        Resource is the resource of the object being referenced.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        UID is the uid of the object being referenced.
        """
        return pulumi.get(self, "uid")


