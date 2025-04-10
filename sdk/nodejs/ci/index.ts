// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./bucketAttachment";
export * from "./bucketPicStyle";
export * from "./guetzli";
export * from "./hotLink";
export * from "./mediaAnimationTemplate";
export * from "./mediaConcatTemplate";
export * from "./mediaPicProcessTemplate";
export * from "./mediaSmartCoverTemplate";
export * from "./mediaSnapshotTemplate";
export * from "./mediaSpeechRecognitionTemplate";
export * from "./mediaSuperResolutionTemplate";
export * from "./mediaTranscodeProTemplate";
export * from "./mediaTranscodeTemplate";
export * from "./mediaTtsTemplate";
export * from "./mediaVideoMontageTemplate";
export * from "./mediaVideoProcessTemplate";
export * from "./mediaVoiceSeparateTemplate";
export * from "./mediaWatermarkTemplate";
export * from "./originalImageProtection";

// Import resources to register:
import { BucketAttachment } from "./bucketAttachment";
import { BucketPicStyle } from "./bucketPicStyle";
import { Guetzli } from "./guetzli";
import { HotLink } from "./hotLink";
import { MediaAnimationTemplate } from "./mediaAnimationTemplate";
import { MediaConcatTemplate } from "./mediaConcatTemplate";
import { MediaPicProcessTemplate } from "./mediaPicProcessTemplate";
import { MediaSmartCoverTemplate } from "./mediaSmartCoverTemplate";
import { MediaSnapshotTemplate } from "./mediaSnapshotTemplate";
import { MediaSpeechRecognitionTemplate } from "./mediaSpeechRecognitionTemplate";
import { MediaSuperResolutionTemplate } from "./mediaSuperResolutionTemplate";
import { MediaTranscodeProTemplate } from "./mediaTranscodeProTemplate";
import { MediaTranscodeTemplate } from "./mediaTranscodeTemplate";
import { MediaTtsTemplate } from "./mediaTtsTemplate";
import { MediaVideoMontageTemplate } from "./mediaVideoMontageTemplate";
import { MediaVideoProcessTemplate } from "./mediaVideoProcessTemplate";
import { MediaVoiceSeparateTemplate } from "./mediaVoiceSeparateTemplate";
import { MediaWatermarkTemplate } from "./mediaWatermarkTemplate";
import { OriginalImageProtection } from "./originalImageProtection";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Ci/bucketAttachment:BucketAttachment":
                return new BucketAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Ci/bucketPicStyle:BucketPicStyle":
                return new BucketPicStyle(name, <any>undefined, { urn })
            case "tencentcloud:Ci/guetzli:Guetzli":
                return new Guetzli(name, <any>undefined, { urn })
            case "tencentcloud:Ci/hotLink:HotLink":
                return new HotLink(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaAnimationTemplate:MediaAnimationTemplate":
                return new MediaAnimationTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaConcatTemplate:MediaConcatTemplate":
                return new MediaConcatTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaPicProcessTemplate:MediaPicProcessTemplate":
                return new MediaPicProcessTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaSmartCoverTemplate:MediaSmartCoverTemplate":
                return new MediaSmartCoverTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate":
                return new MediaSnapshotTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate":
                return new MediaSpeechRecognitionTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate":
                return new MediaSuperResolutionTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaTranscodeProTemplate:MediaTranscodeProTemplate":
                return new MediaTranscodeProTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaTranscodeTemplate:MediaTranscodeTemplate":
                return new MediaTranscodeTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate":
                return new MediaTtsTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate":
                return new MediaVideoMontageTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaVideoProcessTemplate:MediaVideoProcessTemplate":
                return new MediaVideoProcessTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaVoiceSeparateTemplate:MediaVoiceSeparateTemplate":
                return new MediaVoiceSeparateTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate":
                return new MediaWatermarkTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Ci/originalImageProtection:OriginalImageProtection":
                return new OriginalImageProtection(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/bucketAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/bucketPicStyle", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/guetzli", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/hotLink", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaAnimationTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaConcatTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaPicProcessTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaSmartCoverTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaSnapshotTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaSpeechRecognitionTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaSuperResolutionTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaTranscodeProTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaTranscodeTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaTtsTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaVideoMontageTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaVideoProcessTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaVoiceSeparateTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/mediaWatermarkTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ci/originalImageProtection", _module)
