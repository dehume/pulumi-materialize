# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSourcesResult',
    'AwaitableGetSourcesResult',
    'get_sources',
    'get_sources_output',
]

@pulumi.output_type
class GetSourcesResult:
    """
    A collection of values returned by GetSources.
    """
    def __init__(__self__, database_name=None, id=None, schema_name=None, sources=None):
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        pulumi.set(__self__, "database_name", database_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if schema_name and not isinstance(schema_name, str):
            raise TypeError("Expected argument 'schema_name' to be a str")
        pulumi.set(__self__, "schema_name", schema_name)
        if sources and not isinstance(sources, list):
            raise TypeError("Expected argument 'sources' to be a list")
        pulumi.set(__self__, "sources", sources)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[str]:
        """
        Limit sources to a specific database
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> Optional[str]:
        """
        Limit sources to a specific schema within a specific database
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter
    def sources(self) -> Sequence['outputs.GetSourcesSourceResult']:
        """
        The sources in the account
        """
        return pulumi.get(self, "sources")


class AwaitableGetSourcesResult(GetSourcesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSourcesResult(
            database_name=self.database_name,
            id=self.id,
            schema_name=self.schema_name,
            sources=self.sources)


def get_sources(database_name: Optional[str] = None,
                schema_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSourcesResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_materialize as materialize

    all = materialize.get_sources()
    materialize = materialize.get_sources(database_name="materialize")
    materialize_schema = materialize.get_sources(database_name="materialize",
        schema_name="schema")
    ```


    :param str database_name: Limit sources to a specific database
    :param str schema_name: Limit sources to a specific schema within a specific database
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['schemaName'] = schema_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('materialize:index/getSources:GetSources', __args__, opts=opts, typ=GetSourcesResult).value

    return AwaitableGetSourcesResult(
        database_name=__ret__.database_name,
        id=__ret__.id,
        schema_name=__ret__.schema_name,
        sources=__ret__.sources)


@_utilities.lift_output_func(get_sources)
def get_sources_output(database_name: Optional[pulumi.Input[Optional[str]]] = None,
                       schema_name: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSourcesResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_materialize as materialize

    all = materialize.get_sources()
    materialize = materialize.get_sources(database_name="materialize")
    materialize_schema = materialize.get_sources(database_name="materialize",
        schema_name="schema")
    ```


    :param str database_name: Limit sources to a specific database
    :param str schema_name: Limit sources to a specific schema within a specific database
    """
    ...
