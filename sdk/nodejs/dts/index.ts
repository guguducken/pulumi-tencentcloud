// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./compareTask";
export * from "./compareTaskStopOperation";
export * from "./getCompareTasks";
export * from "./getMigrateDbInstances";
export * from "./getMigrateJobs";
export * from "./getSyncJobs";
export * from "./migrateJob";
export * from "./migrateJobConfig";
export * from "./migrateJobResumeOperation";
export * from "./migrateJobStartOperation";
export * from "./migrateService";
export * from "./syncCheckJobOperation";
export * from "./syncConfig";
export * from "./syncJob";
export * from "./syncJobContinueOperation";
export * from "./syncJobIsolateOperation";
export * from "./syncJobPauseOperation";
export * from "./syncJobRecoverOperation";
export * from "./syncJobResizeOperation";
export * from "./syncJobResumeOperation";
export * from "./syncJobStartOperation";
export * from "./syncJobStopOperation";

// Import resources to register:
import { CompareTask } from "./compareTask";
import { CompareTaskStopOperation } from "./compareTaskStopOperation";
import { MigrateJob } from "./migrateJob";
import { MigrateJobConfig } from "./migrateJobConfig";
import { MigrateJobResumeOperation } from "./migrateJobResumeOperation";
import { MigrateJobStartOperation } from "./migrateJobStartOperation";
import { MigrateService } from "./migrateService";
import { SyncCheckJobOperation } from "./syncCheckJobOperation";
import { SyncConfig } from "./syncConfig";
import { SyncJob } from "./syncJob";
import { SyncJobContinueOperation } from "./syncJobContinueOperation";
import { SyncJobIsolateOperation } from "./syncJobIsolateOperation";
import { SyncJobPauseOperation } from "./syncJobPauseOperation";
import { SyncJobRecoverOperation } from "./syncJobRecoverOperation";
import { SyncJobResizeOperation } from "./syncJobResizeOperation";
import { SyncJobResumeOperation } from "./syncJobResumeOperation";
import { SyncJobStartOperation } from "./syncJobStartOperation";
import { SyncJobStopOperation } from "./syncJobStopOperation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Dts/compareTask:CompareTask":
                return new CompareTask(name, <any>undefined, { urn })
            case "tencentcloud:Dts/compareTaskStopOperation:CompareTaskStopOperation":
                return new CompareTaskStopOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/migrateJob:MigrateJob":
                return new MigrateJob(name, <any>undefined, { urn })
            case "tencentcloud:Dts/migrateJobConfig:MigrateJobConfig":
                return new MigrateJobConfig(name, <any>undefined, { urn })
            case "tencentcloud:Dts/migrateJobResumeOperation:MigrateJobResumeOperation":
                return new MigrateJobResumeOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/migrateJobStartOperation:MigrateJobStartOperation":
                return new MigrateJobStartOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/migrateService:MigrateService":
                return new MigrateService(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncCheckJobOperation:SyncCheckJobOperation":
                return new SyncCheckJobOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncConfig:SyncConfig":
                return new SyncConfig(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJob:SyncJob":
                return new SyncJob(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobContinueOperation:SyncJobContinueOperation":
                return new SyncJobContinueOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobIsolateOperation:SyncJobIsolateOperation":
                return new SyncJobIsolateOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobPauseOperation:SyncJobPauseOperation":
                return new SyncJobPauseOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobRecoverOperation:SyncJobRecoverOperation":
                return new SyncJobRecoverOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobResizeOperation:SyncJobResizeOperation":
                return new SyncJobResizeOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobResumeOperation:SyncJobResumeOperation":
                return new SyncJobResumeOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation":
                return new SyncJobStartOperation(name, <any>undefined, { urn })
            case "tencentcloud:Dts/syncJobStopOperation:SyncJobStopOperation":
                return new SyncJobStopOperation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/compareTask", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/compareTaskStopOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/migrateJob", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/migrateJobConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/migrateJobResumeOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/migrateJobStartOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/migrateService", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncCheckJobOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJob", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobContinueOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobIsolateOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobPauseOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobRecoverOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobResizeOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobResumeOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobStartOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dts/syncJobStopOperation", _module)
