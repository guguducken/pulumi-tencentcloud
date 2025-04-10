// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls
{
    /// <summary>
    /// Provides a resource to create a cls alarm
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var alarm = new Tencentcloud.Cls.Alarm("alarm", new Tencentcloud.Cls.AlarmArgs
    ///         {
    ///             AlarmNoticeIds = 
    ///             {
    ///                 "notice-0850756b-245d-4bc7-bb27-2a58fffc780b",
    ///             },
    ///             AlarmPeriod = 15,
    ///             AlarmTargets = 
    ///             {
    ///                 new Tencentcloud.Cls.Inputs.AlarmAlarmTargetArgs
    ///                 {
    ///                     EndTimeOffset = 0,
    ///                     LogsetId = "33aaf0ae-6163-411b-a415-9f27450f68db",
    ///                     Number = 1,
    ///                     Query = "status:&gt;500 | select count(*) as errorCounts",
    ///                     StartTimeOffset = -15,
    ///                     TopicId = "88735a07-bea4-4985-8763-e9deb6da4fad",
    ///                 },
    ///             },
    ///             Analyses = 
    ///             {
    ///                 new Tencentcloud.Cls.Inputs.AlarmAnalysisArgs
    ///                 {
    ///                     ConfigInfos = 
    ///                     {
    ///                         new Tencentcloud.Cls.Inputs.AlarmAnalysisConfigInfoArgs
    ///                         {
    ///                             Key = "QueryIndex",
    ///                             Value = "1",
    ///                         },
    ///                     },
    ///                     Content = "__FILENAME__",
    ///                     Name = "terraform",
    ///                     Type = "field",
    ///                 },
    ///             },
    ///             Condition = "test",
    ///             MessageTemplate = "{{.Label}}",
    ///             MonitorTime = new Tencentcloud.Cls.Inputs.AlarmMonitorTimeArgs
    ///             {
    ///                 Time = 1,
    ///                 Type = "Period",
    ///             },
    ///             Status = true,
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///             TriggerCount = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cls alarm can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cls/alarm:Alarm alarm alarm_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cls/alarm:Alarm")]
    public partial class Alarm : Pulumi.CustomResource
    {
        /// <summary>
        /// list of alarm notice id.
        /// </summary>
        [Output("alarmNoticeIds")]
        public Output<ImmutableArray<string>> AlarmNoticeIds { get; private set; } = null!;

        /// <summary>
        /// alarm repeat cycle.
        /// </summary>
        [Output("alarmPeriod")]
        public Output<int> AlarmPeriod { get; private set; } = null!;

        /// <summary>
        /// list of alarm target.
        /// </summary>
        [Output("alarmTargets")]
        public Output<ImmutableArray<Outputs.AlarmAlarmTarget>> AlarmTargets { get; private set; } = null!;

        /// <summary>
        /// multidimensional analysis.
        /// </summary>
        [Output("analyses")]
        public Output<ImmutableArray<Outputs.AlarmAnalysis>> Analyses { get; private set; } = null!;

        /// <summary>
        /// user define callback.
        /// </summary>
        [Output("callBack")]
        public Output<Outputs.AlarmCallBack?> CallBack { get; private set; } = null!;

        /// <summary>
        /// triggering conditions.
        /// </summary>
        [Output("condition")]
        public Output<string> Condition { get; private set; } = null!;

        /// <summary>
        /// user define alarm notice.
        /// </summary>
        [Output("messageTemplate")]
        public Output<string?> MessageTemplate { get; private set; } = null!;

        /// <summary>
        /// monitor task execution time.
        /// </summary>
        [Output("monitorTime")]
        public Output<Outputs.AlarmMonitorTime> MonitorTime { get; private set; } = null!;

        /// <summary>
        /// log alarm name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// whether to enable the alarm policy.
        /// </summary>
        [Output("status")]
        public Output<bool?> Status { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// continuous cycle.
        /// </summary>
        [Output("triggerCount")]
        public Output<int> TriggerCount { get; private set; } = null!;


        /// <summary>
        /// Create a Alarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alarm(string name, AlarmArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/alarm:Alarm", name, args ?? new AlarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alarm(string name, Input<string> id, AlarmState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/alarm:Alarm", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Alarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alarm Get(string name, Input<string> id, AlarmState? state = null, CustomResourceOptions? options = null)
        {
            return new Alarm(name, id, state, options);
        }
    }

    public sealed class AlarmArgs : Pulumi.ResourceArgs
    {
        [Input("alarmNoticeIds", required: true)]
        private InputList<string>? _alarmNoticeIds;

        /// <summary>
        /// list of alarm notice id.
        /// </summary>
        public InputList<string> AlarmNoticeIds
        {
            get => _alarmNoticeIds ?? (_alarmNoticeIds = new InputList<string>());
            set => _alarmNoticeIds = value;
        }

        /// <summary>
        /// alarm repeat cycle.
        /// </summary>
        [Input("alarmPeriod", required: true)]
        public Input<int> AlarmPeriod { get; set; } = null!;

        [Input("alarmTargets", required: true)]
        private InputList<Inputs.AlarmAlarmTargetArgs>? _alarmTargets;

        /// <summary>
        /// list of alarm target.
        /// </summary>
        public InputList<Inputs.AlarmAlarmTargetArgs> AlarmTargets
        {
            get => _alarmTargets ?? (_alarmTargets = new InputList<Inputs.AlarmAlarmTargetArgs>());
            set => _alarmTargets = value;
        }

        [Input("analyses")]
        private InputList<Inputs.AlarmAnalysisArgs>? _analyses;

        /// <summary>
        /// multidimensional analysis.
        /// </summary>
        public InputList<Inputs.AlarmAnalysisArgs> Analyses
        {
            get => _analyses ?? (_analyses = new InputList<Inputs.AlarmAnalysisArgs>());
            set => _analyses = value;
        }

        /// <summary>
        /// user define callback.
        /// </summary>
        [Input("callBack")]
        public Input<Inputs.AlarmCallBackArgs>? CallBack { get; set; }

        /// <summary>
        /// triggering conditions.
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        /// <summary>
        /// user define alarm notice.
        /// </summary>
        [Input("messageTemplate")]
        public Input<string>? MessageTemplate { get; set; }

        /// <summary>
        /// monitor task execution time.
        /// </summary>
        [Input("monitorTime", required: true)]
        public Input<Inputs.AlarmMonitorTimeArgs> MonitorTime { get; set; } = null!;

        /// <summary>
        /// log alarm name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// whether to enable the alarm policy.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// continuous cycle.
        /// </summary>
        [Input("triggerCount", required: true)]
        public Input<int> TriggerCount { get; set; } = null!;

        public AlarmArgs()
        {
        }
    }

    public sealed class AlarmState : Pulumi.ResourceArgs
    {
        [Input("alarmNoticeIds")]
        private InputList<string>? _alarmNoticeIds;

        /// <summary>
        /// list of alarm notice id.
        /// </summary>
        public InputList<string> AlarmNoticeIds
        {
            get => _alarmNoticeIds ?? (_alarmNoticeIds = new InputList<string>());
            set => _alarmNoticeIds = value;
        }

        /// <summary>
        /// alarm repeat cycle.
        /// </summary>
        [Input("alarmPeriod")]
        public Input<int>? AlarmPeriod { get; set; }

        [Input("alarmTargets")]
        private InputList<Inputs.AlarmAlarmTargetGetArgs>? _alarmTargets;

        /// <summary>
        /// list of alarm target.
        /// </summary>
        public InputList<Inputs.AlarmAlarmTargetGetArgs> AlarmTargets
        {
            get => _alarmTargets ?? (_alarmTargets = new InputList<Inputs.AlarmAlarmTargetGetArgs>());
            set => _alarmTargets = value;
        }

        [Input("analyses")]
        private InputList<Inputs.AlarmAnalysisGetArgs>? _analyses;

        /// <summary>
        /// multidimensional analysis.
        /// </summary>
        public InputList<Inputs.AlarmAnalysisGetArgs> Analyses
        {
            get => _analyses ?? (_analyses = new InputList<Inputs.AlarmAnalysisGetArgs>());
            set => _analyses = value;
        }

        /// <summary>
        /// user define callback.
        /// </summary>
        [Input("callBack")]
        public Input<Inputs.AlarmCallBackGetArgs>? CallBack { get; set; }

        /// <summary>
        /// triggering conditions.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// user define alarm notice.
        /// </summary>
        [Input("messageTemplate")]
        public Input<string>? MessageTemplate { get; set; }

        /// <summary>
        /// monitor task execution time.
        /// </summary>
        [Input("monitorTime")]
        public Input<Inputs.AlarmMonitorTimeGetArgs>? MonitorTime { get; set; }

        /// <summary>
        /// log alarm name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// whether to enable the alarm policy.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// continuous cycle.
        /// </summary>
        [Input("triggerCount")]
        public Input<int>? TriggerCount { get; set; }

        public AlarmState()
        {
        }
    }
}
