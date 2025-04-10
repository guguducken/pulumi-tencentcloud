// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ckafka topicFlowRanking
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const topicFlowRanking = pulumi.output(tencentcloud.Ckafka.getTopicFlowRanking({
 *     beginDate: "2023-05-29T00:00:00+08:00",
 *     endDate: "2021-05-29T23:59:59+08:00",
 *     instanceId: "ckafka-xxxxxx",
 *     rankingType: "PRO",
 * }));
 * ```
 */
export function getTopicFlowRanking(args: GetTopicFlowRankingArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicFlowRankingResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Ckafka/getTopicFlowRanking:getTopicFlowRanking", {
        "beginDate": args.beginDate,
        "endDate": args.endDate,
        "instanceId": args.instanceId,
        "rankingType": args.rankingType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopicFlowRanking.
 */
export interface GetTopicFlowRankingArgs {
    /**
     * BeginDate.
     */
    beginDate?: string;
    /**
     * EndDate.
     */
    endDate?: string;
    /**
     * InstanceId.
     */
    instanceId: string;
    /**
     * Ranking type. `PRO`: topic production flow, `CON`: topic consumption traffic.
     */
    rankingType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getTopicFlowRanking.
 */
export interface GetTopicFlowRankingResult {
    readonly beginDate?: string;
    readonly endDate?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly rankingType: string;
    readonly resultOutputFile?: string;
    /**
     * result.
     */
    readonly results: outputs.Ckafka.GetTopicFlowRankingResult[];
}

export function getTopicFlowRankingOutput(args: GetTopicFlowRankingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicFlowRankingResult> {
    return pulumi.output(args).apply(a => getTopicFlowRanking(a, opts))
}

/**
 * A collection of arguments for invoking getTopicFlowRanking.
 */
export interface GetTopicFlowRankingOutputArgs {
    /**
     * BeginDate.
     */
    beginDate?: pulumi.Input<string>;
    /**
     * EndDate.
     */
    endDate?: pulumi.Input<string>;
    /**
     * InstanceId.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Ranking type. `PRO`: topic production flow, `CON`: topic consumption traffic.
     */
    rankingType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
