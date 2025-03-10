# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'HelmReleaseSettingsArgs',
    'KubeClientSettingsArgs',
]

@pulumi.input_type
class HelmReleaseSettingsArgs:
    def __init__(__self__, *,
                 driver: Optional[pulumi.Input[str]] = None,
                 plugins_path: Optional[pulumi.Input[str]] = None,
                 registry_config_path: Optional[pulumi.Input[str]] = None,
                 repository_cache: Optional[pulumi.Input[str]] = None,
                 repository_config_path: Optional[pulumi.Input[str]] = None):
        """
        Options to configure the Helm Release resource.
        :param pulumi.Input[str] driver: The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
        :param pulumi.Input[str] plugins_path: The path to the helm plugins directory.
        :param pulumi.Input[str] registry_config_path: The path to the registry config file.
        :param pulumi.Input[str] repository_cache: The path to the file containing cached repository indexes.
        :param pulumi.Input[str] repository_config_path: The path to the file containing repository names and URLs.
        """
        HelmReleaseSettingsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            driver=driver,
            plugins_path=plugins_path,
            registry_config_path=registry_config_path,
            repository_cache=repository_cache,
            repository_config_path=repository_config_path,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             driver: Optional[pulumi.Input[str]] = None,
             plugins_path: Optional[pulumi.Input[str]] = None,
             registry_config_path: Optional[pulumi.Input[str]] = None,
             repository_cache: Optional[pulumi.Input[str]] = None,
             repository_config_path: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'pluginsPath' in kwargs:
            plugins_path = kwargs['pluginsPath']
        if 'registryConfigPath' in kwargs:
            registry_config_path = kwargs['registryConfigPath']
        if 'repositoryCache' in kwargs:
            repository_cache = kwargs['repositoryCache']
        if 'repositoryConfigPath' in kwargs:
            repository_config_path = kwargs['repositoryConfigPath']

        if driver is None:
            driver = _utilities.get_env('PULUMI_K8S_HELM_DRIVER')
        if driver is not None:
            _setter("driver", driver)
        if plugins_path is None:
            plugins_path = _utilities.get_env('PULUMI_K8S_HELM_PLUGINS_PATH')
        if plugins_path is not None:
            _setter("plugins_path", plugins_path)
        if registry_config_path is None:
            registry_config_path = _utilities.get_env('PULUMI_K8S_HELM_REGISTRY_CONFIG_PATH')
        if registry_config_path is not None:
            _setter("registry_config_path", registry_config_path)
        if repository_cache is None:
            repository_cache = _utilities.get_env('PULUMI_K8S_HELM_REPOSITORY_CACHE')
        if repository_cache is not None:
            _setter("repository_cache", repository_cache)
        if repository_config_path is None:
            repository_config_path = _utilities.get_env('PULUMI_K8S_HELM_REPOSITORY_CONFIG_PATH')
        if repository_config_path is not None:
            _setter("repository_config_path", repository_config_path)

    @property
    @pulumi.getter
    def driver(self) -> Optional[pulumi.Input[str]]:
        """
        The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
        """
        return pulumi.get(self, "driver")

    @driver.setter
    def driver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "driver", value)

    @property
    @pulumi.getter(name="pluginsPath")
    def plugins_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the helm plugins directory.
        """
        return pulumi.get(self, "plugins_path")

    @plugins_path.setter
    def plugins_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugins_path", value)

    @property
    @pulumi.getter(name="registryConfigPath")
    def registry_config_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the registry config file.
        """
        return pulumi.get(self, "registry_config_path")

    @registry_config_path.setter
    def registry_config_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_config_path", value)

    @property
    @pulumi.getter(name="repositoryCache")
    def repository_cache(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the file containing cached repository indexes.
        """
        return pulumi.get(self, "repository_cache")

    @repository_cache.setter
    def repository_cache(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_cache", value)

    @property
    @pulumi.getter(name="repositoryConfigPath")
    def repository_config_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the file containing repository names and URLs.
        """
        return pulumi.get(self, "repository_config_path")

    @repository_config_path.setter
    def repository_config_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_config_path", value)


@pulumi.input_type
class KubeClientSettingsArgs:
    def __init__(__self__, *,
                 burst: Optional[pulumi.Input[int]] = None,
                 qps: Optional[pulumi.Input[float]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Options for tuning the Kubernetes client used by a Provider.
        :param pulumi.Input[int] burst: Maximum burst for throttle. Default value is 10.
        :param pulumi.Input[float] qps: Maximum queries per second (QPS) to the API server from this client. Default value is 5.
        :param pulumi.Input[int] timeout: Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
        """
        KubeClientSettingsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            burst=burst,
            qps=qps,
            timeout=timeout,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             burst: Optional[pulumi.Input[int]] = None,
             qps: Optional[pulumi.Input[float]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if burst is None:
            burst = _utilities.get_env_int('PULUMI_K8S_CLIENT_BURST')
        if burst is not None:
            _setter("burst", burst)
        if qps is None:
            qps = _utilities.get_env_float('PULUMI_K8S_CLIENT_QPS')
        if qps is not None:
            _setter("qps", qps)
        if timeout is None:
            timeout = _utilities.get_env_int('PULUMI_K8S_CLIENT_TIMEOUT')
        if timeout is not None:
            _setter("timeout", timeout)

    @property
    @pulumi.getter
    def burst(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum burst for throttle. Default value is 10.
        """
        return pulumi.get(self, "burst")

    @burst.setter
    def burst(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "burst", value)

    @property
    @pulumi.getter
    def qps(self) -> Optional[pulumi.Input[float]]:
        """
        Maximum queries per second (QPS) to the API server from this client. Default value is 5.
        """
        return pulumi.get(self, "qps")

    @qps.setter
    def qps(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "qps", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


