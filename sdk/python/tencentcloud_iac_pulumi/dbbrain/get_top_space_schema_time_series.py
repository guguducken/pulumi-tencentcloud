# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetTopSpaceSchemaTimeSeriesResult',
    'AwaitableGetTopSpaceSchemaTimeSeriesResult',
    'get_top_space_schema_time_series',
    'get_top_space_schema_time_series_output',
]

@pulumi.output_type
class GetTopSpaceSchemaTimeSeriesResult:
    """
    A collection of values returned by getTopSpaceSchemaTimeSeries.
    """
    def __init__(__self__, end_date=None, id=None, instance_id=None, limit=None, product=None, result_output_file=None, sort_by=None, start_date=None, top_space_schema_time_series=None):
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sort_by and not isinstance(sort_by, str):
            raise TypeError("Expected argument 'sort_by' to be a str")
        pulumi.set(__self__, "sort_by", sort_by)
        if start_date and not isinstance(start_date, str):
            raise TypeError("Expected argument 'start_date' to be a str")
        pulumi.set(__self__, "start_date", start_date)
        if top_space_schema_time_series and not isinstance(top_space_schema_time_series, list):
            raise TypeError("Expected argument 'top_space_schema_time_series' to be a list")
        pulumi.set(__self__, "top_space_schema_time_series", top_space_schema_time_series)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[str]:
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def product(self) -> Optional[str]:
        return pulumi.get(self, "product")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="sortBy")
    def sort_by(self) -> Optional[str]:
        return pulumi.get(self, "sort_by")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[str]:
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter(name="topSpaceSchemaTimeSeries")
    def top_space_schema_time_series(self) -> Sequence['outputs.GetTopSpaceSchemaTimeSeriesTopSpaceSchemaTimeSeriesResult']:
        """
        The time series data list of the returned top library space statistics.
        """
        return pulumi.get(self, "top_space_schema_time_series")


class AwaitableGetTopSpaceSchemaTimeSeriesResult(GetTopSpaceSchemaTimeSeriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopSpaceSchemaTimeSeriesResult(
            end_date=self.end_date,
            id=self.id,
            instance_id=self.instance_id,
            limit=self.limit,
            product=self.product,
            result_output_file=self.result_output_file,
            sort_by=self.sort_by,
            start_date=self.start_date,
            top_space_schema_time_series=self.top_space_schema_time_series)


def get_top_space_schema_time_series(end_date: Optional[str] = None,
                                     instance_id: Optional[str] = None,
                                     limit: Optional[int] = None,
                                     product: Optional[str] = None,
                                     result_output_file: Optional[str] = None,
                                     sort_by: Optional[str] = None,
                                     start_date: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopSpaceSchemaTimeSeriesResult:
    """
    Use this data source to query detailed information of dbbrain top_space_schema_time_series

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    top_space_schema_time_series = tencentcloud.Dbbrain.get_top_space_schema_time_series(end_date="%s",
        instance_id="%s",
        product="mysql",
        sort_by="DataLength",
        start_date="%s")
    ```


    :param str end_date: The deadline, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the current day.
    :param str instance_id: instance id.
    :param int limit: The number of Top libraries to return, the maximum value is 100, and the default is 20.
    :param str product: Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
    :param str result_output_file: Used to save results.
    :param str sort_by: The sorting field used to filter the Top library. The optional fields include DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, and PhysicalFileSize (only supported by ApsaraDB for MySQL instances). The default for ApsaraDB for MySQL instances is PhysicalFileSize, and the default for other product instances is TotalLength.
    :param str start_date: The start date, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the 6th day before the deadline.
    """
    __args__ = dict()
    __args__['endDate'] = end_date
    __args__['instanceId'] = instance_id
    __args__['limit'] = limit
    __args__['product'] = product
    __args__['resultOutputFile'] = result_output_file
    __args__['sortBy'] = sort_by
    __args__['startDate'] = start_date
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dbbrain/getTopSpaceSchemaTimeSeries:getTopSpaceSchemaTimeSeries', __args__, opts=opts, typ=GetTopSpaceSchemaTimeSeriesResult).value

    return AwaitableGetTopSpaceSchemaTimeSeriesResult(
        end_date=__ret__.end_date,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        limit=__ret__.limit,
        product=__ret__.product,
        result_output_file=__ret__.result_output_file,
        sort_by=__ret__.sort_by,
        start_date=__ret__.start_date,
        top_space_schema_time_series=__ret__.top_space_schema_time_series)


@_utilities.lift_output_func(get_top_space_schema_time_series)
def get_top_space_schema_time_series_output(end_date: Optional[pulumi.Input[Optional[str]]] = None,
                                            instance_id: Optional[pulumi.Input[str]] = None,
                                            limit: Optional[pulumi.Input[Optional[int]]] = None,
                                            product: Optional[pulumi.Input[Optional[str]]] = None,
                                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                            sort_by: Optional[pulumi.Input[Optional[str]]] = None,
                                            start_date: Optional[pulumi.Input[Optional[str]]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopSpaceSchemaTimeSeriesResult]:
    """
    Use this data source to query detailed information of dbbrain top_space_schema_time_series

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    top_space_schema_time_series = tencentcloud.Dbbrain.get_top_space_schema_time_series(end_date="%s",
        instance_id="%s",
        product="mysql",
        sort_by="DataLength",
        start_date="%s")
    ```


    :param str end_date: The deadline, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the current day.
    :param str instance_id: instance id.
    :param int limit: The number of Top libraries to return, the maximum value is 100, and the default is 20.
    :param str product: Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
    :param str result_output_file: Used to save results.
    :param str sort_by: The sorting field used to filter the Top library. The optional fields include DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, and PhysicalFileSize (only supported by ApsaraDB for MySQL instances). The default for ApsaraDB for MySQL instances is PhysicalFileSize, and the default for other product instances is TotalLength.
    :param str start_date: The start date, such as 2021-01-01, the earliest is the 29th day before the current day, and the default is the 6th day before the deadline.
    """
    ...
