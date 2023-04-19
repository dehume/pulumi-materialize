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
from ._inputs import *

__all__ = ['SourceLoadgenArgs', 'SourceLoadgen']

@pulumi.input_type
class SourceLoadgenArgs:
    def __init__(__self__, *,
                 load_generator_type: pulumi.Input[str],
                 auction_options: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 counter_options: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tpch_options: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]]] = None):
        """
        The set of arguments for constructing a SourceLoadgen resource.
        :param pulumi.Input[str] load_generator_type: The load generator types: [AUCTION COUNTER TPCH].
        :param pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]] auction_options: Auction Options.
        :param pulumi.Input[str] cluster_name: The cluster to maintain this source. If not specified, the size option must be specified.
        :param pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]] counter_options: Counter Options.
        :param pulumi.Input[str] database_name: The identifier for the source database.
        :param pulumi.Input[str] name: The identifier for the source.
        :param pulumi.Input[str] schema_name: The identifier for the source schema.
        :param pulumi.Input[str] size: The size of the source.
        :param pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]] tpch_options: TPCH Options.
        """
        pulumi.set(__self__, "load_generator_type", load_generator_type)
        if auction_options is not None:
            pulumi.set(__self__, "auction_options", auction_options)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if counter_options is not None:
            pulumi.set(__self__, "counter_options", counter_options)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if schema_name is not None:
            pulumi.set(__self__, "schema_name", schema_name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if tpch_options is not None:
            pulumi.set(__self__, "tpch_options", tpch_options)

    @property
    @pulumi.getter(name="loadGeneratorType")
    def load_generator_type(self) -> pulumi.Input[str]:
        """
        The load generator types: [AUCTION COUNTER TPCH].
        """
        return pulumi.get(self, "load_generator_type")

    @load_generator_type.setter
    def load_generator_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_generator_type", value)

    @property
    @pulumi.getter(name="auctionOptions")
    def auction_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]]]:
        """
        Auction Options.
        """
        return pulumi.get(self, "auction_options")

    @auction_options.setter
    def auction_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]]]):
        pulumi.set(self, "auction_options", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster to maintain this source. If not specified, the size option must be specified.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="counterOptions")
    def counter_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]]]:
        """
        Counter Options.
        """
        return pulumi.get(self, "counter_options")

    @counter_options.setter
    def counter_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]]]):
        pulumi.set(self, "counter_options", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the source database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the source schema.
        """
        return pulumi.get(self, "schema_name")

    @schema_name.setter
    def schema_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the source.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="tpchOptions")
    def tpch_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]]]:
        """
        TPCH Options.
        """
        return pulumi.get(self, "tpch_options")

    @tpch_options.setter
    def tpch_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]]]):
        pulumi.set(self, "tpch_options", value)


@pulumi.input_type
class _SourceLoadgenState:
    def __init__(__self__, *,
                 auction_options: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 counter_options: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 load_generator_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 qualified_sql_name: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 tpch_options: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]]] = None):
        """
        Input properties used for looking up and filtering SourceLoadgen resources.
        :param pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]] auction_options: Auction Options.
        :param pulumi.Input[str] cluster_name: The cluster to maintain this source. If not specified, the size option must be specified.
        :param pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]] counter_options: Counter Options.
        :param pulumi.Input[str] database_name: The identifier for the source database.
        :param pulumi.Input[str] load_generator_type: The load generator types: [AUCTION COUNTER TPCH].
        :param pulumi.Input[str] name: The identifier for the source.
        :param pulumi.Input[str] qualified_sql_name: The fully qualified name of the source.
        :param pulumi.Input[str] schema_name: The identifier for the source schema.
        :param pulumi.Input[str] size: The size of the source.
        :param pulumi.Input[str] source_type: The type of source.
        :param pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]] tpch_options: TPCH Options.
        """
        if auction_options is not None:
            pulumi.set(__self__, "auction_options", auction_options)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if counter_options is not None:
            pulumi.set(__self__, "counter_options", counter_options)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if load_generator_type is not None:
            pulumi.set(__self__, "load_generator_type", load_generator_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if qualified_sql_name is not None:
            pulumi.set(__self__, "qualified_sql_name", qualified_sql_name)
        if schema_name is not None:
            pulumi.set(__self__, "schema_name", schema_name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if tpch_options is not None:
            pulumi.set(__self__, "tpch_options", tpch_options)

    @property
    @pulumi.getter(name="auctionOptions")
    def auction_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]]]:
        """
        Auction Options.
        """
        return pulumi.get(self, "auction_options")

    @auction_options.setter
    def auction_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenAuctionOptionArgs']]]]):
        pulumi.set(self, "auction_options", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster to maintain this source. If not specified, the size option must be specified.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="counterOptions")
    def counter_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]]]:
        """
        Counter Options.
        """
        return pulumi.get(self, "counter_options")

    @counter_options.setter
    def counter_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenCounterOptionArgs']]]]):
        pulumi.set(self, "counter_options", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the source database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="loadGeneratorType")
    def load_generator_type(self) -> Optional[pulumi.Input[str]]:
        """
        The load generator types: [AUCTION COUNTER TPCH].
        """
        return pulumi.get(self, "load_generator_type")

    @load_generator_type.setter
    def load_generator_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_generator_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="qualifiedSqlName")
    def qualified_sql_name(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified name of the source.
        """
        return pulumi.get(self, "qualified_sql_name")

    @qualified_sql_name.setter
    def qualified_sql_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualified_sql_name", value)

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the source schema.
        """
        return pulumi.get(self, "schema_name")

    @schema_name.setter
    def schema_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the source.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of source.
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="tpchOptions")
    def tpch_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]]]:
        """
        TPCH Options.
        """
        return pulumi.get(self, "tpch_options")

    @tpch_options.setter
    def tpch_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SourceLoadgenTpchOptionArgs']]]]):
        pulumi.set(self, "tpch_options", value)


class SourceLoadgen(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auction_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenAuctionOptionArgs']]]]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 counter_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenCounterOptionArgs']]]]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 load_generator_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tpch_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenTpchOptionArgs']]]]] = None,
                 __props__=None):
        """
        A source describes an external system you want Materialize to read data from, and provides details about how to decode and interpret that data.

        ## Import

        # Sources can be imported using the source id

        ```sh
         $ pulumi import materialize:index/sourceLoadgen:SourceLoadgen example_source_load_generator <source_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenAuctionOptionArgs']]]] auction_options: Auction Options.
        :param pulumi.Input[str] cluster_name: The cluster to maintain this source. If not specified, the size option must be specified.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenCounterOptionArgs']]]] counter_options: Counter Options.
        :param pulumi.Input[str] database_name: The identifier for the source database.
        :param pulumi.Input[str] load_generator_type: The load generator types: [AUCTION COUNTER TPCH].
        :param pulumi.Input[str] name: The identifier for the source.
        :param pulumi.Input[str] schema_name: The identifier for the source schema.
        :param pulumi.Input[str] size: The size of the source.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenTpchOptionArgs']]]] tpch_options: TPCH Options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SourceLoadgenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A source describes an external system you want Materialize to read data from, and provides details about how to decode and interpret that data.

        ## Import

        # Sources can be imported using the source id

        ```sh
         $ pulumi import materialize:index/sourceLoadgen:SourceLoadgen example_source_load_generator <source_id>
        ```

        :param str resource_name: The name of the resource.
        :param SourceLoadgenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SourceLoadgenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auction_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenAuctionOptionArgs']]]]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 counter_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenCounterOptionArgs']]]]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 load_generator_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tpch_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenTpchOptionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SourceLoadgenArgs.__new__(SourceLoadgenArgs)

            __props__.__dict__["auction_options"] = auction_options
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["counter_options"] = counter_options
            __props__.__dict__["database_name"] = database_name
            if load_generator_type is None and not opts.urn:
                raise TypeError("Missing required property 'load_generator_type'")
            __props__.__dict__["load_generator_type"] = load_generator_type
            __props__.__dict__["name"] = name
            __props__.__dict__["schema_name"] = schema_name
            __props__.__dict__["size"] = size
            __props__.__dict__["tpch_options"] = tpch_options
            __props__.__dict__["qualified_sql_name"] = None
            __props__.__dict__["source_type"] = None
        super(SourceLoadgen, __self__).__init__(
            'materialize:index/sourceLoadgen:SourceLoadgen',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auction_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenAuctionOptionArgs']]]]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            counter_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenCounterOptionArgs']]]]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            load_generator_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            qualified_sql_name: Optional[pulumi.Input[str]] = None,
            schema_name: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[str]] = None,
            source_type: Optional[pulumi.Input[str]] = None,
            tpch_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenTpchOptionArgs']]]]] = None) -> 'SourceLoadgen':
        """
        Get an existing SourceLoadgen resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenAuctionOptionArgs']]]] auction_options: Auction Options.
        :param pulumi.Input[str] cluster_name: The cluster to maintain this source. If not specified, the size option must be specified.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenCounterOptionArgs']]]] counter_options: Counter Options.
        :param pulumi.Input[str] database_name: The identifier for the source database.
        :param pulumi.Input[str] load_generator_type: The load generator types: [AUCTION COUNTER TPCH].
        :param pulumi.Input[str] name: The identifier for the source.
        :param pulumi.Input[str] qualified_sql_name: The fully qualified name of the source.
        :param pulumi.Input[str] schema_name: The identifier for the source schema.
        :param pulumi.Input[str] size: The size of the source.
        :param pulumi.Input[str] source_type: The type of source.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SourceLoadgenTpchOptionArgs']]]] tpch_options: TPCH Options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SourceLoadgenState.__new__(_SourceLoadgenState)

        __props__.__dict__["auction_options"] = auction_options
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["counter_options"] = counter_options
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["load_generator_type"] = load_generator_type
        __props__.__dict__["name"] = name
        __props__.__dict__["qualified_sql_name"] = qualified_sql_name
        __props__.__dict__["schema_name"] = schema_name
        __props__.__dict__["size"] = size
        __props__.__dict__["source_type"] = source_type
        __props__.__dict__["tpch_options"] = tpch_options
        return SourceLoadgen(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="auctionOptions")
    def auction_options(self) -> pulumi.Output[Optional[Sequence['outputs.SourceLoadgenAuctionOption']]]:
        """
        Auction Options.
        """
        return pulumi.get(self, "auction_options")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The cluster to maintain this source. If not specified, the size option must be specified.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="counterOptions")
    def counter_options(self) -> pulumi.Output[Optional[Sequence['outputs.SourceLoadgenCounterOption']]]:
        """
        Counter Options.
        """
        return pulumi.get(self, "counter_options")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[Optional[str]]:
        """
        The identifier for the source database.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="loadGeneratorType")
    def load_generator_type(self) -> pulumi.Output[str]:
        """
        The load generator types: [AUCTION COUNTER TPCH].
        """
        return pulumi.get(self, "load_generator_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The identifier for the source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="qualifiedSqlName")
    def qualified_sql_name(self) -> pulumi.Output[str]:
        """
        The fully qualified name of the source.
        """
        return pulumi.get(self, "qualified_sql_name")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> pulumi.Output[Optional[str]]:
        """
        The identifier for the source schema.
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        The size of the source.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Output[str]:
        """
        The type of source.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="tpchOptions")
    def tpch_options(self) -> pulumi.Output[Optional[Sequence['outputs.SourceLoadgenTpchOption']]]:
        """
        TPCH Options.
        """
        return pulumi.get(self, "tpch_options")

