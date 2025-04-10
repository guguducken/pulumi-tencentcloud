// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTargetResourceScfParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of messages sent in each batch, the default is 1000.
        /// </summary>
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        /// <summary>
        /// SCF function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// The number of retries after the SCF call fails, the default is 5.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// SCF cloud function namespace, the default is default.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// SCF cloud function version and alias, the default is DEFAULT.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        public DatahubTaskTargetResourceScfParamGetArgs()
        {
        }
    }
}
