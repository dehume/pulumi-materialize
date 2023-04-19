# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('materialize')


class _ExportableConfig(types.ModuleType):
    @property
    def database(self) -> Optional[str]:
        """
        The Materialize database
        """
        return __config__.get('database')

    @property
    def host(self) -> Optional[str]:
        """
        Materialize host
        """
        return __config__.get('host')

    @property
    def password(self) -> Optional[str]:
        """
        Materialize host
        """
        return __config__.get('password')

    @property
    def port(self) -> Optional[int]:
        """
        The Materialize port number to connect to at the server host
        """
        return __config__.get_int('port')

    @property
    def testing(self) -> Optional[bool]:
        """
        Enable to test the provider locally
        """
        return __config__.get_bool('testing')

    @property
    def username(self) -> Optional[str]:
        """
        Materialize username
        """
        return __config__.get('username')

