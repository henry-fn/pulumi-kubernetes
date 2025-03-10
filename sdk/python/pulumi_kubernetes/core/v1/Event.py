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
from ... import meta as _meta
from ._inputs import *

__all__ = ['EventInitArgs', 'Event']

@pulumi.input_type
class EventInitArgs:
    def __init__(__self__, *,
                 involved_object: pulumi.Input['ObjectReferenceArgs'],
                 metadata: pulumi.Input['_meta.v1.ObjectMetaArgs'],
                 action: Optional[pulumi.Input[str]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[int]] = None,
                 event_time: Optional[pulumi.Input[str]] = None,
                 first_timestamp: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 last_timestamp: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 related: Optional[pulumi.Input['ObjectReferenceArgs']] = None,
                 reporting_component: Optional[pulumi.Input[str]] = None,
                 reporting_instance: Optional[pulumi.Input[str]] = None,
                 series: Optional[pulumi.Input['EventSeriesArgs']] = None,
                 source: Optional[pulumi.Input['EventSourceArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Event resource.
        :param pulumi.Input['ObjectReferenceArgs'] involved_object: The object that this event is about.
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[str] action: What action was taken/failed regarding to the Regarding object.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[int] count: The number of times this event has occurred.
        :param pulumi.Input[str] event_time: Time when this Event was first observed.
        :param pulumi.Input[str] first_timestamp: The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[str] last_timestamp: The time at which the most recent occurrence of this event was recorded.
        :param pulumi.Input[str] message: A human-readable description of the status of this operation.
        :param pulumi.Input[str] reason: This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
        :param pulumi.Input['ObjectReferenceArgs'] related: Optional secondary object for more complex actions.
        :param pulumi.Input[str] reporting_component: Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        :param pulumi.Input[str] reporting_instance: ID of the controller instance, e.g. `kubelet-xyzf`.
        :param pulumi.Input['EventSeriesArgs'] series: Data about the Event series this event represents or nil if it's a singleton Event.
        :param pulumi.Input['EventSourceArgs'] source: The component reporting this event. Should be a short machine understandable string.
        :param pulumi.Input[str] type: Type of this event (Normal, Warning), new types could be added in the future
        """
        EventInitArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            involved_object=involved_object,
            metadata=metadata,
            action=action,
            api_version=api_version,
            count=count,
            event_time=event_time,
            first_timestamp=first_timestamp,
            kind=kind,
            last_timestamp=last_timestamp,
            message=message,
            reason=reason,
            related=related,
            reporting_component=reporting_component,
            reporting_instance=reporting_instance,
            series=series,
            source=source,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             involved_object: pulumi.Input['ObjectReferenceArgs'],
             metadata: pulumi.Input['_meta.v1.ObjectMetaArgs'],
             action: Optional[pulumi.Input[str]] = None,
             api_version: Optional[pulumi.Input[str]] = None,
             count: Optional[pulumi.Input[int]] = None,
             event_time: Optional[pulumi.Input[str]] = None,
             first_timestamp: Optional[pulumi.Input[str]] = None,
             kind: Optional[pulumi.Input[str]] = None,
             last_timestamp: Optional[pulumi.Input[str]] = None,
             message: Optional[pulumi.Input[str]] = None,
             reason: Optional[pulumi.Input[str]] = None,
             related: Optional[pulumi.Input['ObjectReferenceArgs']] = None,
             reporting_component: Optional[pulumi.Input[str]] = None,
             reporting_instance: Optional[pulumi.Input[str]] = None,
             series: Optional[pulumi.Input['EventSeriesArgs']] = None,
             source: Optional[pulumi.Input['EventSourceArgs']] = None,
             type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'involvedObject' in kwargs:
            involved_object = kwargs['involvedObject']
        if 'apiVersion' in kwargs:
            api_version = kwargs['apiVersion']
        if 'eventTime' in kwargs:
            event_time = kwargs['eventTime']
        if 'firstTimestamp' in kwargs:
            first_timestamp = kwargs['firstTimestamp']
        if 'lastTimestamp' in kwargs:
            last_timestamp = kwargs['lastTimestamp']
        if 'reportingComponent' in kwargs:
            reporting_component = kwargs['reportingComponent']
        if 'reportingInstance' in kwargs:
            reporting_instance = kwargs['reportingInstance']

        _setter("involved_object", involved_object)
        _setter("metadata", metadata)
        if action is not None:
            _setter("action", action)
        if api_version is not None:
            _setter("api_version", 'v1')
        if count is not None:
            _setter("count", count)
        if event_time is not None:
            _setter("event_time", event_time)
        if first_timestamp is not None:
            _setter("first_timestamp", first_timestamp)
        if kind is not None:
            _setter("kind", 'Event')
        if last_timestamp is not None:
            _setter("last_timestamp", last_timestamp)
        if message is not None:
            _setter("message", message)
        if reason is not None:
            _setter("reason", reason)
        if related is not None:
            _setter("related", related)
        if reporting_component is not None:
            _setter("reporting_component", reporting_component)
        if reporting_instance is not None:
            _setter("reporting_instance", reporting_instance)
        if series is not None:
            _setter("series", series)
        if source is not None:
            _setter("source", source)
        if type is not None:
            _setter("type", type)

    @property
    @pulumi.getter(name="involvedObject")
    def involved_object(self) -> pulumi.Input['ObjectReferenceArgs']:
        """
        The object that this event is about.
        """
        return pulumi.get(self, "involved_object")

    @involved_object.setter
    def involved_object(self, value: pulumi.Input['ObjectReferenceArgs']):
        pulumi.set(self, "involved_object", value)

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Input['_meta.v1.ObjectMetaArgs']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: pulumi.Input['_meta.v1.ObjectMetaArgs']):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        What action was taken/failed regarding to the Regarding object.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

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
    def count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of times this event has occurred.
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter(name="eventTime")
    def event_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time when this Event was first observed.
        """
        return pulumi.get(self, "event_time")

    @event_time.setter
    def event_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_time", value)

    @property
    @pulumi.getter(name="firstTimestamp")
    def first_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
        """
        return pulumi.get(self, "first_timestamp")

    @first_timestamp.setter
    def first_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_timestamp", value)

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
    @pulumi.getter(name="lastTimestamp")
    def last_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the most recent occurrence of this event was recorded.
        """
        return pulumi.get(self, "last_timestamp")

    @last_timestamp.setter
    def last_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_timestamp", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable description of the status of this operation.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def reason(self) -> Optional[pulumi.Input[str]]:
        """
        This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
        """
        return pulumi.get(self, "reason")

    @reason.setter
    def reason(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reason", value)

    @property
    @pulumi.getter
    def related(self) -> Optional[pulumi.Input['ObjectReferenceArgs']]:
        """
        Optional secondary object for more complex actions.
        """
        return pulumi.get(self, "related")

    @related.setter
    def related(self, value: Optional[pulumi.Input['ObjectReferenceArgs']]):
        pulumi.set(self, "related", value)

    @property
    @pulumi.getter(name="reportingComponent")
    def reporting_component(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        """
        return pulumi.get(self, "reporting_component")

    @reporting_component.setter
    def reporting_component(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reporting_component", value)

    @property
    @pulumi.getter(name="reportingInstance")
    def reporting_instance(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the controller instance, e.g. `kubelet-xyzf`.
        """
        return pulumi.get(self, "reporting_instance")

    @reporting_instance.setter
    def reporting_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reporting_instance", value)

    @property
    @pulumi.getter
    def series(self) -> Optional[pulumi.Input['EventSeriesArgs']]:
        """
        Data about the Event series this event represents or nil if it's a singleton Event.
        """
        return pulumi.get(self, "series")

    @series.setter
    def series(self, value: Optional[pulumi.Input['EventSeriesArgs']]):
        pulumi.set(self, "series", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input['EventSourceArgs']]:
        """
        The component reporting this event. Should be a short machine understandable string.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input['EventSourceArgs']]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of this event (Normal, Warning), new types could be added in the future
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Event(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[int]] = None,
                 event_time: Optional[pulumi.Input[str]] = None,
                 first_timestamp: Optional[pulumi.Input[str]] = None,
                 involved_object: Optional[pulumi.Input[pulumi.InputType['ObjectReferenceArgs']]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 last_timestamp: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 related: Optional[pulumi.Input[pulumi.InputType['ObjectReferenceArgs']]] = None,
                 reporting_component: Optional[pulumi.Input[str]] = None,
                 reporting_instance: Optional[pulumi.Input[str]] = None,
                 series: Optional[pulumi.Input[pulumi.InputType['EventSeriesArgs']]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['EventSourceArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: What action was taken/failed regarding to the Regarding object.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[int] count: The number of times this event has occurred.
        :param pulumi.Input[str] event_time: Time when this Event was first observed.
        :param pulumi.Input[str] first_timestamp: The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
        :param pulumi.Input[pulumi.InputType['ObjectReferenceArgs']] involved_object: The object that this event is about.
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[str] last_timestamp: The time at which the most recent occurrence of this event was recorded.
        :param pulumi.Input[str] message: A human-readable description of the status of this operation.
        :param pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[str] reason: This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
        :param pulumi.Input[pulumi.InputType['ObjectReferenceArgs']] related: Optional secondary object for more complex actions.
        :param pulumi.Input[str] reporting_component: Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        :param pulumi.Input[str] reporting_instance: ID of the controller instance, e.g. `kubelet-xyzf`.
        :param pulumi.Input[pulumi.InputType['EventSeriesArgs']] series: Data about the Event series this event represents or nil if it's a singleton Event.
        :param pulumi.Input[pulumi.InputType['EventSourceArgs']] source: The component reporting this event. Should be a short machine understandable string.
        :param pulumi.Input[str] type: Type of this event (Normal, Warning), new types could be added in the future
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.

        :param str resource_name: The name of the resource.
        :param EventInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            EventInitArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[int]] = None,
                 event_time: Optional[pulumi.Input[str]] = None,
                 first_timestamp: Optional[pulumi.Input[str]] = None,
                 involved_object: Optional[pulumi.Input[pulumi.InputType['ObjectReferenceArgs']]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 last_timestamp: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 related: Optional[pulumi.Input[pulumi.InputType['ObjectReferenceArgs']]] = None,
                 reporting_component: Optional[pulumi.Input[str]] = None,
                 reporting_instance: Optional[pulumi.Input[str]] = None,
                 series: Optional[pulumi.Input[pulumi.InputType['EventSeriesArgs']]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['EventSourceArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventInitArgs.__new__(EventInitArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["api_version"] = 'v1'
            __props__.__dict__["count"] = count
            __props__.__dict__["event_time"] = event_time
            __props__.__dict__["first_timestamp"] = first_timestamp
            if involved_object is not None and not isinstance(involved_object, ObjectReferenceArgs):
                involved_object = involved_object or {}
                def _setter(key, value):
                    involved_object[key] = value
                ObjectReferenceArgs._configure(_setter, **involved_object)
            if involved_object is None and not opts.urn:
                raise TypeError("Missing required property 'involved_object'")
            __props__.__dict__["involved_object"] = involved_object
            __props__.__dict__["kind"] = 'Event'
            __props__.__dict__["last_timestamp"] = last_timestamp
            __props__.__dict__["message"] = message
            if metadata is not None and not isinstance(metadata, _meta.v1.ObjectMetaArgs):
                metadata = metadata or {}
                def _setter(key, value):
                    metadata[key] = value
                _meta.v1.ObjectMetaArgs._configure(_setter, **metadata)
            if metadata is None and not opts.urn:
                raise TypeError("Missing required property 'metadata'")
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["reason"] = reason
            if related is not None and not isinstance(related, ObjectReferenceArgs):
                related = related or {}
                def _setter(key, value):
                    related[key] = value
                ObjectReferenceArgs._configure(_setter, **related)
            __props__.__dict__["related"] = related
            __props__.__dict__["reporting_component"] = reporting_component
            __props__.__dict__["reporting_instance"] = reporting_instance
            if series is not None and not isinstance(series, EventSeriesArgs):
                series = series or {}
                def _setter(key, value):
                    series[key] = value
                EventSeriesArgs._configure(_setter, **series)
            __props__.__dict__["series"] = series
            if source is not None and not isinstance(source, EventSourceArgs):
                source = source or {}
                def _setter(key, value):
                    source[key] = value
                EventSourceArgs._configure(_setter, **source)
            __props__.__dict__["source"] = source
            __props__.__dict__["type"] = type
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="kubernetes:events.k8s.io/v1:Event"), pulumi.Alias(type_="kubernetes:events.k8s.io/v1beta1:Event")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Event, __self__).__init__(
            'kubernetes:core/v1:Event',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Event':
        """
        Get an existing Event resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EventInitArgs.__new__(EventInitArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["api_version"] = None
        __props__.__dict__["count"] = None
        __props__.__dict__["event_time"] = None
        __props__.__dict__["first_timestamp"] = None
        __props__.__dict__["involved_object"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["last_timestamp"] = None
        __props__.__dict__["message"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["reason"] = None
        __props__.__dict__["related"] = None
        __props__.__dict__["reporting_component"] = None
        __props__.__dict__["reporting_instance"] = None
        __props__.__dict__["series"] = None
        __props__.__dict__["source"] = None
        __props__.__dict__["type"] = None
        return Event(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        What action was taken/failed regarding to the Regarding object.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def count(self) -> pulumi.Output[int]:
        """
        The number of times this event has occurred.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="eventTime")
    def event_time(self) -> pulumi.Output[str]:
        """
        Time when this Event was first observed.
        """
        return pulumi.get(self, "event_time")

    @property
    @pulumi.getter(name="firstTimestamp")
    def first_timestamp(self) -> pulumi.Output[str]:
        """
        The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
        """
        return pulumi.get(self, "first_timestamp")

    @property
    @pulumi.getter(name="involvedObject")
    def involved_object(self) -> pulumi.Output['outputs.ObjectReference']:
        """
        The object that this event is about.
        """
        return pulumi.get(self, "involved_object")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastTimestamp")
    def last_timestamp(self) -> pulumi.Output[str]:
        """
        The time at which the most recent occurrence of this event was recorded.
        """
        return pulumi.get(self, "last_timestamp")

    @property
    @pulumi.getter
    def message(self) -> pulumi.Output[str]:
        """
        A human-readable description of the status of this operation.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def reason(self) -> pulumi.Output[str]:
        """
        This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
        """
        return pulumi.get(self, "reason")

    @property
    @pulumi.getter
    def related(self) -> pulumi.Output['outputs.ObjectReference']:
        """
        Optional secondary object for more complex actions.
        """
        return pulumi.get(self, "related")

    @property
    @pulumi.getter(name="reportingComponent")
    def reporting_component(self) -> pulumi.Output[str]:
        """
        Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        """
        return pulumi.get(self, "reporting_component")

    @property
    @pulumi.getter(name="reportingInstance")
    def reporting_instance(self) -> pulumi.Output[str]:
        """
        ID of the controller instance, e.g. `kubelet-xyzf`.
        """
        return pulumi.get(self, "reporting_instance")

    @property
    @pulumi.getter
    def series(self) -> pulumi.Output['outputs.EventSeries']:
        """
        Data about the Event series this event represents or nil if it's a singleton Event.
        """
        return pulumi.get(self, "series")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output['outputs.EventSource']:
        """
        The component reporting this event. Should be a short machine understandable string.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of this event (Normal, Warning), new types could be added in the future
        """
        return pulumi.get(self, "type")

