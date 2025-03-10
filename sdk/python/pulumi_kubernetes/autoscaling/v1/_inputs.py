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
    'CrossVersionObjectReferencePatchArgs',
    'CrossVersionObjectReferenceArgs',
    'HorizontalPodAutoscalerSpecPatchArgs',
    'HorizontalPodAutoscalerSpecArgs',
    'HorizontalPodAutoscalerStatusArgs',
    'HorizontalPodAutoscalerArgs',
]

@pulumi.input_type
class CrossVersionObjectReferencePatchArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        CrossVersionObjectReference contains enough information to let you identify the referred resource.
        :param pulumi.Input[str] api_version: apiVersion is the API version of the referent
        :param pulumi.Input[str] kind: kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[str] name: name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        """
        CrossVersionObjectReferencePatchArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_version=api_version,
            kind=kind,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_version: Optional[pulumi.Input[str]] = None,
             kind: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']

        if api_version is not None:
            _setter("api_version", api_version)
        if kind is not None:
            _setter("kind", kind)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        apiVersion is the API version of the referent
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class CrossVersionObjectReferenceArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 name: pulumi.Input[str],
                 api_version: Optional[pulumi.Input[str]] = None):
        """
        CrossVersionObjectReference contains enough information to let you identify the referred resource.
        :param pulumi.Input[str] kind: kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[str] name: name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        :param pulumi.Input[str] api_version: apiVersion is the API version of the referent
        """
        CrossVersionObjectReferenceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            kind=kind,
            name=name,
            api_version=api_version,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             kind: pulumi.Input[str],
             name: pulumi.Input[str],
             api_version: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']

        _setter("kind", kind)
        _setter("name", name)
        if api_version is not None:
            _setter("api_version", api_version)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        apiVersion is the API version of the referent
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)


@pulumi.input_type
class HorizontalPodAutoscalerSpecPatchArgs:
    def __init__(__self__, *,
                 max_replicas: Optional[pulumi.Input[int]] = None,
                 min_replicas: Optional[pulumi.Input[int]] = None,
                 scale_target_ref: Optional[pulumi.Input['CrossVersionObjectReferencePatchArgs']] = None,
                 target_cpu_utilization_percentage: Optional[pulumi.Input[int]] = None):
        """
        specification of a horizontal pod autoscaler.
        :param pulumi.Input[int] max_replicas: maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        :param pulumi.Input[int] min_replicas: minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        :param pulumi.Input['CrossVersionObjectReferencePatchArgs'] scale_target_ref: reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        :param pulumi.Input[int] target_cpu_utilization_percentage: targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        HorizontalPodAutoscalerSpecPatchArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            max_replicas=max_replicas,
            min_replicas=min_replicas,
            scale_target_ref=scale_target_ref,
            target_cpu_utilization_percentage=target_cpu_utilization_percentage,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             max_replicas: Optional[pulumi.Input[int]] = None,
             min_replicas: Optional[pulumi.Input[int]] = None,
             scale_target_ref: Optional[pulumi.Input['CrossVersionObjectReferencePatchArgs']] = None,
             target_cpu_utilization_percentage: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'maxReplicas' in kwargs:
            max_replicas = kwargs['maxReplicas']
        if 'minReplicas' in kwargs:
            min_replicas = kwargs['minReplicas']
        if 'scaleTargetRef' in kwargs:
            scale_target_ref = kwargs['scaleTargetRef']
        if 'targetCPUUtilizationPercentage' in kwargs:
            target_cpu_utilization_percentage = kwargs['targetCPUUtilizationPercentage']

        if max_replicas is not None:
            _setter("max_replicas", max_replicas)
        if min_replicas is not None:
            _setter("min_replicas", min_replicas)
        if scale_target_ref is not None:
            _setter("scale_target_ref", scale_target_ref)
        if target_cpu_utilization_percentage is not None:
            _setter("target_cpu_utilization_percentage", target_cpu_utilization_percentage)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        """
        return pulumi.get(self, "max_replicas")

    @max_replicas.setter
    def max_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_replicas", value)

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        """
        return pulumi.get(self, "min_replicas")

    @min_replicas.setter
    def min_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_replicas", value)

    @property
    @pulumi.getter(name="scaleTargetRef")
    def scale_target_ref(self) -> Optional[pulumi.Input['CrossVersionObjectReferencePatchArgs']]:
        """
        reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        """
        return pulumi.get(self, "scale_target_ref")

    @scale_target_ref.setter
    def scale_target_ref(self, value: Optional[pulumi.Input['CrossVersionObjectReferencePatchArgs']]):
        pulumi.set(self, "scale_target_ref", value)

    @property
    @pulumi.getter(name="targetCPUUtilizationPercentage")
    def target_cpu_utilization_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        return pulumi.get(self, "target_cpu_utilization_percentage")

    @target_cpu_utilization_percentage.setter
    def target_cpu_utilization_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "target_cpu_utilization_percentage", value)


@pulumi.input_type
class HorizontalPodAutoscalerSpecArgs:
    def __init__(__self__, *,
                 max_replicas: pulumi.Input[int],
                 scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'],
                 min_replicas: Optional[pulumi.Input[int]] = None,
                 target_cpu_utilization_percentage: Optional[pulumi.Input[int]] = None):
        """
        specification of a horizontal pod autoscaler.
        :param pulumi.Input[int] max_replicas: maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        :param pulumi.Input['CrossVersionObjectReferenceArgs'] scale_target_ref: reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        :param pulumi.Input[int] min_replicas: minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        :param pulumi.Input[int] target_cpu_utilization_percentage: targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        HorizontalPodAutoscalerSpecArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            max_replicas=max_replicas,
            scale_target_ref=scale_target_ref,
            min_replicas=min_replicas,
            target_cpu_utilization_percentage=target_cpu_utilization_percentage,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             max_replicas: pulumi.Input[int],
             scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'],
             min_replicas: Optional[pulumi.Input[int]] = None,
             target_cpu_utilization_percentage: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'maxReplicas' in kwargs:
            max_replicas = kwargs['maxReplicas']
        if 'scaleTargetRef' in kwargs:
            scale_target_ref = kwargs['scaleTargetRef']
        if 'minReplicas' in kwargs:
            min_replicas = kwargs['minReplicas']
        if 'targetCPUUtilizationPercentage' in kwargs:
            target_cpu_utilization_percentage = kwargs['targetCPUUtilizationPercentage']

        _setter("max_replicas", max_replicas)
        _setter("scale_target_ref", scale_target_ref)
        if min_replicas is not None:
            _setter("min_replicas", min_replicas)
        if target_cpu_utilization_percentage is not None:
            _setter("target_cpu_utilization_percentage", target_cpu_utilization_percentage)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> pulumi.Input[int]:
        """
        maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        """
        return pulumi.get(self, "max_replicas")

    @max_replicas.setter
    def max_replicas(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_replicas", value)

    @property
    @pulumi.getter(name="scaleTargetRef")
    def scale_target_ref(self) -> pulumi.Input['CrossVersionObjectReferenceArgs']:
        """
        reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        """
        return pulumi.get(self, "scale_target_ref")

    @scale_target_ref.setter
    def scale_target_ref(self, value: pulumi.Input['CrossVersionObjectReferenceArgs']):
        pulumi.set(self, "scale_target_ref", value)

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        """
        return pulumi.get(self, "min_replicas")

    @min_replicas.setter
    def min_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_replicas", value)

    @property
    @pulumi.getter(name="targetCPUUtilizationPercentage")
    def target_cpu_utilization_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        return pulumi.get(self, "target_cpu_utilization_percentage")

    @target_cpu_utilization_percentage.setter
    def target_cpu_utilization_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "target_cpu_utilization_percentage", value)


@pulumi.input_type
class HorizontalPodAutoscalerStatusArgs:
    def __init__(__self__, *,
                 current_replicas: pulumi.Input[int],
                 desired_replicas: pulumi.Input[int],
                 current_cpu_utilization_percentage: Optional[pulumi.Input[int]] = None,
                 last_scale_time: Optional[pulumi.Input[str]] = None,
                 observed_generation: Optional[pulumi.Input[int]] = None):
        """
        current status of a horizontal pod autoscaler
        :param pulumi.Input[int] current_replicas: currentReplicas is the current number of replicas of pods managed by this autoscaler.
        :param pulumi.Input[int] desired_replicas: desiredReplicas is the  desired number of replicas of pods managed by this autoscaler.
        :param pulumi.Input[int] current_cpu_utilization_percentage: currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        :param pulumi.Input[str] last_scale_time: lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        :param pulumi.Input[int] observed_generation: observedGeneration is the most recent generation observed by this autoscaler.
        """
        HorizontalPodAutoscalerStatusArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            current_replicas=current_replicas,
            desired_replicas=desired_replicas,
            current_cpu_utilization_percentage=current_cpu_utilization_percentage,
            last_scale_time=last_scale_time,
            observed_generation=observed_generation,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             current_replicas: pulumi.Input[int],
             desired_replicas: pulumi.Input[int],
             current_cpu_utilization_percentage: Optional[pulumi.Input[int]] = None,
             last_scale_time: Optional[pulumi.Input[str]] = None,
             observed_generation: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'currentReplicas' in kwargs:
            current_replicas = kwargs['currentReplicas']
        if 'desiredReplicas' in kwargs:
            desired_replicas = kwargs['desiredReplicas']
        if 'currentCPUUtilizationPercentage' in kwargs:
            current_cpu_utilization_percentage = kwargs['currentCPUUtilizationPercentage']
        if 'lastScaleTime' in kwargs:
            last_scale_time = kwargs['lastScaleTime']
        if 'observedGeneration' in kwargs:
            observed_generation = kwargs['observedGeneration']

        _setter("current_replicas", current_replicas)
        _setter("desired_replicas", desired_replicas)
        if current_cpu_utilization_percentage is not None:
            _setter("current_cpu_utilization_percentage", current_cpu_utilization_percentage)
        if last_scale_time is not None:
            _setter("last_scale_time", last_scale_time)
        if observed_generation is not None:
            _setter("observed_generation", observed_generation)

    @property
    @pulumi.getter(name="currentReplicas")
    def current_replicas(self) -> pulumi.Input[int]:
        """
        currentReplicas is the current number of replicas of pods managed by this autoscaler.
        """
        return pulumi.get(self, "current_replicas")

    @current_replicas.setter
    def current_replicas(self, value: pulumi.Input[int]):
        pulumi.set(self, "current_replicas", value)

    @property
    @pulumi.getter(name="desiredReplicas")
    def desired_replicas(self) -> pulumi.Input[int]:
        """
        desiredReplicas is the  desired number of replicas of pods managed by this autoscaler.
        """
        return pulumi.get(self, "desired_replicas")

    @desired_replicas.setter
    def desired_replicas(self, value: pulumi.Input[int]):
        pulumi.set(self, "desired_replicas", value)

    @property
    @pulumi.getter(name="currentCPUUtilizationPercentage")
    def current_cpu_utilization_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        """
        return pulumi.get(self, "current_cpu_utilization_percentage")

    @current_cpu_utilization_percentage.setter
    def current_cpu_utilization_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "current_cpu_utilization_percentage", value)

    @property
    @pulumi.getter(name="lastScaleTime")
    def last_scale_time(self) -> Optional[pulumi.Input[str]]:
        """
        lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        """
        return pulumi.get(self, "last_scale_time")

    @last_scale_time.setter
    def last_scale_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_scale_time", value)

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> Optional[pulumi.Input[int]]:
        """
        observedGeneration is the most recent generation observed by this autoscaler.
        """
        return pulumi.get(self, "observed_generation")

    @observed_generation.setter
    def observed_generation(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "observed_generation", value)


@pulumi.input_type
class HorizontalPodAutoscalerArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = None,
                 status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = None):
        """
        configuration of a horizontal pod autoscaler.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['HorizontalPodAutoscalerSpecArgs'] spec: spec defines the behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        :param pulumi.Input['HorizontalPodAutoscalerStatusArgs'] status: status is the current information about the autoscaler.
        """
        HorizontalPodAutoscalerArgs._configure(
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
             spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = None,
             status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']

        if api_version is not None:
            _setter("api_version", 'autoscaling/v1')
        if kind is not None:
            _setter("kind", 'HorizontalPodAutoscaler')
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
        """
        Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']]:
        """
        spec defines the behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']]:
        """
        status is the current information about the autoscaler.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']]):
        pulumi.set(self, "status", value)


