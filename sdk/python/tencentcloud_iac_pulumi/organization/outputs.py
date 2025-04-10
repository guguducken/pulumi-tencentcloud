# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'OrgMemberOrgPermission',
]

@pulumi.output_type
class OrgMemberOrgPermission(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 name: Optional[str] = None):
        """
        :param int id: Permissions ID.
        :param str name: Member name.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Permissions ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Member name.
        """
        return pulumi.get(self, "name")


