// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a pts job
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const job = new tencentcloud.Pts.Job("job", {
 *     jobOwner: "username",
 *     // debug = ""
 *     note: "desc",
 *     projectId: "project-45vw7v82",
 *     scenarioId: "scenario-22q19f3k",
 * });
 * ```
 *
 * ## Import
 *
 * pts job can be imported using the projectId#scenarioId#jobId, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Pts/job:Job job project-45vw7v82#scenario-22q19f3k#job-dtm93vx0
 * ```
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Pts/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Cause of interruption.
     */
    public /*out*/ readonly abortReason!: pulumi.Output<number>;
    /**
     * Creation time of the job.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Scheduled job ID.
     */
    public /*out*/ readonly cronId!: pulumi.Output<string>;
    /**
     * Dataset file for the job.
     */
    public /*out*/ readonly datasets!: pulumi.Output<outputs.Pts.JobDataset[]>;
    /**
     * Whether to debug.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * Domain name binding configuration.
     */
    public /*out*/ readonly domainNameConfigs!: pulumi.Output<outputs.Pts.JobDomainNameConfig[]>;
    /**
     * Job duration.
     */
    public /*out*/ readonly duration!: pulumi.Output<number>;
    /**
     * End time of the job.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * Percentage of error rate.
     */
    public /*out*/ readonly errorRate!: pulumi.Output<number>;
    /**
     * Job owner.
     */
    public readonly jobOwner!: pulumi.Output<string>;
    /**
     * Pressure configuration of job.
     */
    public /*out*/ readonly loads!: pulumi.Output<outputs.Pts.JobLoad[]>;
    /**
     * Maximum requests per second.
     */
    public /*out*/ readonly maxRequestsPerSecond!: pulumi.Output<number>;
    /**
     * Maximum number of VU for the job.
     */
    public /*out*/ readonly maxVirtualUserCount!: pulumi.Output<number>;
    /**
     * Note.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * Expansion package file information.
     */
    public /*out*/ readonly plugins!: pulumi.Output<outputs.Pts.JobPlugin[]>;
    /**
     * Project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Protocol script information.
     */
    public /*out*/ readonly protocols!: pulumi.Output<outputs.Pts.JobProtocol[]>;
    /**
     * Request file information.
     */
    public /*out*/ readonly requestFiles!: pulumi.Output<outputs.Pts.JobRequestFile[]>;
    /**
     * Total number of requests.
     */
    public /*out*/ readonly requestTotal!: pulumi.Output<number>;
    /**
     * Average number of requests per second.
     */
    public /*out*/ readonly requestsPerSecond!: pulumi.Output<number>;
    /**
     * Average response time.
     */
    public /*out*/ readonly responseTimeAverage!: pulumi.Output<number>;
    /**
     * Maximum response time.
     */
    public /*out*/ readonly responseTimeMax!: pulumi.Output<number>;
    /**
     * Minimum response time.
     */
    public /*out*/ readonly responseTimeMin!: pulumi.Output<number>;
    /**
     * 90th percentile response time.
     */
    public /*out*/ readonly responseTimeP90!: pulumi.Output<number>;
    /**
     * 95th percentile response time.
     */
    public /*out*/ readonly responseTimeP95!: pulumi.Output<number>;
    /**
     * 99th percentile response time.
     */
    public /*out*/ readonly responseTimeP99!: pulumi.Output<number>;
    /**
     * Pts scenario id.
     */
    public readonly scenarioId!: pulumi.Output<string>;
    /**
     * Start time of the job.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * The running status of the task; `0`: JobUnknown, `1`: JobCreated, `2`: JobPending, `3`: JobPreparing, `4`: JobSelectClustering, `5`: JobCreateTasking, `6`: JobSyncTasking, `11`: JobRunning, `12`: JobFinished, `13`: JobPrepareException, `14`: JobFinishException, `15`: JobAborting, `16`: JobAborted, `17`: JobAbortException, `18`: JobDeleted, `19`: JobSelectClusterException, `20`: JobCreateTaskException, `21`: JobSyncTaskException.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Test script information.
     */
    public /*out*/ readonly testScripts!: pulumi.Output<outputs.Pts.JobTestScript[]>;
    /**
     * Scene Type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobState | undefined;
            resourceInputs["abortReason"] = state ? state.abortReason : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["cronId"] = state ? state.cronId : undefined;
            resourceInputs["datasets"] = state ? state.datasets : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["domainNameConfigs"] = state ? state.domainNameConfigs : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["errorRate"] = state ? state.errorRate : undefined;
            resourceInputs["jobOwner"] = state ? state.jobOwner : undefined;
            resourceInputs["loads"] = state ? state.loads : undefined;
            resourceInputs["maxRequestsPerSecond"] = state ? state.maxRequestsPerSecond : undefined;
            resourceInputs["maxVirtualUserCount"] = state ? state.maxVirtualUserCount : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
            resourceInputs["plugins"] = state ? state.plugins : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["requestFiles"] = state ? state.requestFiles : undefined;
            resourceInputs["requestTotal"] = state ? state.requestTotal : undefined;
            resourceInputs["requestsPerSecond"] = state ? state.requestsPerSecond : undefined;
            resourceInputs["responseTimeAverage"] = state ? state.responseTimeAverage : undefined;
            resourceInputs["responseTimeMax"] = state ? state.responseTimeMax : undefined;
            resourceInputs["responseTimeMin"] = state ? state.responseTimeMin : undefined;
            resourceInputs["responseTimeP90"] = state ? state.responseTimeP90 : undefined;
            resourceInputs["responseTimeP95"] = state ? state.responseTimeP95 : undefined;
            resourceInputs["responseTimeP99"] = state ? state.responseTimeP99 : undefined;
            resourceInputs["scenarioId"] = state ? state.scenarioId : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["testScripts"] = state ? state.testScripts : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.jobOwner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobOwner'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.scenarioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scenarioId'");
            }
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["jobOwner"] = args ? args.jobOwner : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["scenarioId"] = args ? args.scenarioId : undefined;
            resourceInputs["abortReason"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["cronId"] = undefined /*out*/;
            resourceInputs["datasets"] = undefined /*out*/;
            resourceInputs["domainNameConfigs"] = undefined /*out*/;
            resourceInputs["duration"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["errorRate"] = undefined /*out*/;
            resourceInputs["loads"] = undefined /*out*/;
            resourceInputs["maxRequestsPerSecond"] = undefined /*out*/;
            resourceInputs["maxVirtualUserCount"] = undefined /*out*/;
            resourceInputs["plugins"] = undefined /*out*/;
            resourceInputs["protocols"] = undefined /*out*/;
            resourceInputs["requestFiles"] = undefined /*out*/;
            resourceInputs["requestTotal"] = undefined /*out*/;
            resourceInputs["requestsPerSecond"] = undefined /*out*/;
            resourceInputs["responseTimeAverage"] = undefined /*out*/;
            resourceInputs["responseTimeMax"] = undefined /*out*/;
            resourceInputs["responseTimeMin"] = undefined /*out*/;
            resourceInputs["responseTimeP90"] = undefined /*out*/;
            resourceInputs["responseTimeP95"] = undefined /*out*/;
            resourceInputs["responseTimeP99"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["testScripts"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * Cause of interruption.
     */
    abortReason?: pulumi.Input<number>;
    /**
     * Creation time of the job.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Scheduled job ID.
     */
    cronId?: pulumi.Input<string>;
    /**
     * Dataset file for the job.
     */
    datasets?: pulumi.Input<pulumi.Input<inputs.Pts.JobDataset>[]>;
    /**
     * Whether to debug.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Domain name binding configuration.
     */
    domainNameConfigs?: pulumi.Input<pulumi.Input<inputs.Pts.JobDomainNameConfig>[]>;
    /**
     * Job duration.
     */
    duration?: pulumi.Input<number>;
    /**
     * End time of the job.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Percentage of error rate.
     */
    errorRate?: pulumi.Input<number>;
    /**
     * Job owner.
     */
    jobOwner?: pulumi.Input<string>;
    /**
     * Pressure configuration of job.
     */
    loads?: pulumi.Input<pulumi.Input<inputs.Pts.JobLoad>[]>;
    /**
     * Maximum requests per second.
     */
    maxRequestsPerSecond?: pulumi.Input<number>;
    /**
     * Maximum number of VU for the job.
     */
    maxVirtualUserCount?: pulumi.Input<number>;
    /**
     * Note.
     */
    note?: pulumi.Input<string>;
    /**
     * Expansion package file information.
     */
    plugins?: pulumi.Input<pulumi.Input<inputs.Pts.JobPlugin>[]>;
    /**
     * Project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Protocol script information.
     */
    protocols?: pulumi.Input<pulumi.Input<inputs.Pts.JobProtocol>[]>;
    /**
     * Request file information.
     */
    requestFiles?: pulumi.Input<pulumi.Input<inputs.Pts.JobRequestFile>[]>;
    /**
     * Total number of requests.
     */
    requestTotal?: pulumi.Input<number>;
    /**
     * Average number of requests per second.
     */
    requestsPerSecond?: pulumi.Input<number>;
    /**
     * Average response time.
     */
    responseTimeAverage?: pulumi.Input<number>;
    /**
     * Maximum response time.
     */
    responseTimeMax?: pulumi.Input<number>;
    /**
     * Minimum response time.
     */
    responseTimeMin?: pulumi.Input<number>;
    /**
     * 90th percentile response time.
     */
    responseTimeP90?: pulumi.Input<number>;
    /**
     * 95th percentile response time.
     */
    responseTimeP95?: pulumi.Input<number>;
    /**
     * 99th percentile response time.
     */
    responseTimeP99?: pulumi.Input<number>;
    /**
     * Pts scenario id.
     */
    scenarioId?: pulumi.Input<string>;
    /**
     * Start time of the job.
     */
    startTime?: pulumi.Input<string>;
    /**
     * The running status of the task; `0`: JobUnknown, `1`: JobCreated, `2`: JobPending, `3`: JobPreparing, `4`: JobSelectClustering, `5`: JobCreateTasking, `6`: JobSyncTasking, `11`: JobRunning, `12`: JobFinished, `13`: JobPrepareException, `14`: JobFinishException, `15`: JobAborting, `16`: JobAborted, `17`: JobAbortException, `18`: JobDeleted, `19`: JobSelectClusterException, `20`: JobCreateTaskException, `21`: JobSyncTaskException.
     */
    status?: pulumi.Input<number>;
    /**
     * Test script information.
     */
    testScripts?: pulumi.Input<pulumi.Input<inputs.Pts.JobTestScript>[]>;
    /**
     * Scene Type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Whether to debug.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Job owner.
     */
    jobOwner: pulumi.Input<string>;
    /**
     * Note.
     */
    note?: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId: pulumi.Input<string>;
    /**
     * Pts scenario id.
     */
    scenarioId: pulumi.Input<string>;
}
