# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetSyncJobsResult',
    'AwaitableGetSyncJobsResult',
    'get_sync_jobs',
    'get_sync_jobs_output',
]

@pulumi.output_type
class GetSyncJobsResult:
    """
    A collection of values returned by getSyncJobs.
    """
    def __init__(__self__, id=None, job_id=None, job_name=None, job_type=None, lists=None, order=None, order_seq=None, pay_mode=None, result_output_file=None, run_mode=None, statuses=None, tag_filters=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if job_id and not isinstance(job_id, str):
            raise TypeError("Expected argument 'job_id' to be a str")
        pulumi.set(__self__, "job_id", job_id)
        if job_name and not isinstance(job_name, str):
            raise TypeError("Expected argument 'job_name' to be a str")
        pulumi.set(__self__, "job_name", job_name)
        if job_type and not isinstance(job_type, str):
            raise TypeError("Expected argument 'job_type' to be a str")
        pulumi.set(__self__, "job_type", job_type)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if order and not isinstance(order, str):
            raise TypeError("Expected argument 'order' to be a str")
        pulumi.set(__self__, "order", order)
        if order_seq and not isinstance(order_seq, str):
            raise TypeError("Expected argument 'order_seq' to be a str")
        pulumi.set(__self__, "order_seq", order_seq)
        if pay_mode and not isinstance(pay_mode, str):
            raise TypeError("Expected argument 'pay_mode' to be a str")
        pulumi.set(__self__, "pay_mode", pay_mode)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if run_mode and not isinstance(run_mode, str):
            raise TypeError("Expected argument 'run_mode' to be a str")
        pulumi.set(__self__, "run_mode", run_mode)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if tag_filters and not isinstance(tag_filters, list):
            raise TypeError("Expected argument 'tag_filters' to be a list")
        pulumi.set(__self__, "tag_filters", tag_filters)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[str]:
        """
        job id.
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[str]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> Optional[str]:
        return pulumi.get(self, "job_type")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetSyncJobsListResult']:
        """
        sync job list.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="orderSeq")
    def order_seq(self) -> Optional[str]:
        return pulumi.get(self, "order_seq")

    @property
    @pulumi.getter(name="payMode")
    def pay_mode(self) -> Optional[str]:
        """
        pay mode.
        """
        return pulumi.get(self, "pay_mode")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="runMode")
    def run_mode(self) -> Optional[str]:
        """
        run mode.
        """
        return pulumi.get(self, "run_mode")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        """
        status.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="tagFilters")
    def tag_filters(self) -> Optional[Sequence['outputs.GetSyncJobsTagFilterResult']]:
        return pulumi.get(self, "tag_filters")


class AwaitableGetSyncJobsResult(GetSyncJobsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSyncJobsResult(
            id=self.id,
            job_id=self.job_id,
            job_name=self.job_name,
            job_type=self.job_type,
            lists=self.lists,
            order=self.order,
            order_seq=self.order_seq,
            pay_mode=self.pay_mode,
            result_output_file=self.result_output_file,
            run_mode=self.run_mode,
            statuses=self.statuses,
            tag_filters=self.tag_filters)


def get_sync_jobs(job_id: Optional[str] = None,
                  job_name: Optional[str] = None,
                  job_type: Optional[str] = None,
                  order: Optional[str] = None,
                  order_seq: Optional[str] = None,
                  pay_mode: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  run_mode: Optional[str] = None,
                  statuses: Optional[Sequence[str]] = None,
                  tag_filters: Optional[Sequence[pulumi.InputType['GetSyncJobsTagFilterArgs']]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSyncJobsResult:
    """
    Use this data source to query detailed information of dts syncJobs

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    job = tencentcloud.dts.SyncJob("job",
        job_name="tf_dts_test",
        pay_mode="PostPay",
        src_database_type="mysql",
        src_region="ap-guangzhou",
        dst_database_type="cynosdbmysql",
        dst_region="ap-guangzhou",
        tags=[tencentcloud.dts.SyncJobTagArgs(
            tag_key="aaa",
            tag_value="bbb",
        )],
        auto_renew=0,
        instance_class="micro")
    sync_jobs = tencentcloud.Dts.get_sync_jobs_output(job_id=job.id,
        job_name="tf_dts_test")
    ```


    :param str job_id: job id.
    :param str job_name: job name.
    :param str job_type: job type.
    :param str order: order field.
    :param str order_seq: order way, optional value is DESC or ASC.
    :param str pay_mode: pay mode, optional value is PrePay or PostPay.
    :param str result_output_file: Used to save results.
    :param str run_mode: run mode, optional value is mmediate or Timed.
    :param Sequence[str] statuses: status.
    :param Sequence[pulumi.InputType['GetSyncJobsTagFilterArgs']] tag_filters: tag filters.
    """
    __args__ = dict()
    __args__['jobId'] = job_id
    __args__['jobName'] = job_name
    __args__['jobType'] = job_type
    __args__['order'] = order
    __args__['orderSeq'] = order_seq
    __args__['payMode'] = pay_mode
    __args__['resultOutputFile'] = result_output_file
    __args__['runMode'] = run_mode
    __args__['statuses'] = statuses
    __args__['tagFilters'] = tag_filters
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dts/getSyncJobs:getSyncJobs', __args__, opts=opts, typ=GetSyncJobsResult).value

    return AwaitableGetSyncJobsResult(
        id=__ret__.id,
        job_id=__ret__.job_id,
        job_name=__ret__.job_name,
        job_type=__ret__.job_type,
        lists=__ret__.lists,
        order=__ret__.order,
        order_seq=__ret__.order_seq,
        pay_mode=__ret__.pay_mode,
        result_output_file=__ret__.result_output_file,
        run_mode=__ret__.run_mode,
        statuses=__ret__.statuses,
        tag_filters=__ret__.tag_filters)


@_utilities.lift_output_func(get_sync_jobs)
def get_sync_jobs_output(job_id: Optional[pulumi.Input[Optional[str]]] = None,
                         job_name: Optional[pulumi.Input[Optional[str]]] = None,
                         job_type: Optional[pulumi.Input[Optional[str]]] = None,
                         order: Optional[pulumi.Input[Optional[str]]] = None,
                         order_seq: Optional[pulumi.Input[Optional[str]]] = None,
                         pay_mode: Optional[pulumi.Input[Optional[str]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         run_mode: Optional[pulumi.Input[Optional[str]]] = None,
                         statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         tag_filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetSyncJobsTagFilterArgs']]]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSyncJobsResult]:
    """
    Use this data source to query detailed information of dts syncJobs

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    job = tencentcloud.dts.SyncJob("job",
        job_name="tf_dts_test",
        pay_mode="PostPay",
        src_database_type="mysql",
        src_region="ap-guangzhou",
        dst_database_type="cynosdbmysql",
        dst_region="ap-guangzhou",
        tags=[tencentcloud.dts.SyncJobTagArgs(
            tag_key="aaa",
            tag_value="bbb",
        )],
        auto_renew=0,
        instance_class="micro")
    sync_jobs = tencentcloud.Dts.get_sync_jobs_output(job_id=job.id,
        job_name="tf_dts_test")
    ```


    :param str job_id: job id.
    :param str job_name: job name.
    :param str job_type: job type.
    :param str order: order field.
    :param str order_seq: order way, optional value is DESC or ASC.
    :param str pay_mode: pay mode, optional value is PrePay or PostPay.
    :param str result_output_file: Used to save results.
    :param str run_mode: run mode, optional value is mmediate or Timed.
    :param Sequence[str] statuses: status.
    :param Sequence[pulumi.InputType['GetSyncJobsTagFilterArgs']] tag_filters: tag filters.
    """
    ...
