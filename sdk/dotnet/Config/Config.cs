// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("tencentcloud");

        private static readonly __Value<TencentCloudIAC.PulumiPackage.Tencentcloud.Config.Types.AssumeRole?> _assumeRole = new __Value<TencentCloudIAC.PulumiPackage.Tencentcloud.Config.Types.AssumeRole?>(() => __config.GetObject<TencentCloudIAC.PulumiPackage.Tencentcloud.Config.Types.AssumeRole>("assumeRole"));
        /// <summary>
        /// The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.
        /// </summary>
        public static TencentCloudIAC.PulumiPackage.Tencentcloud.Config.Types.AssumeRole? AssumeRole
        {
            get => _assumeRole.Get();
            set => _assumeRole.Set(value);
        }

        private static readonly __Value<string?> _domain = new __Value<string?>(() => __config.Get("domain"));
        /// <summary>
        /// The root domain of the API request, Default is `tencentcloudapi.com`.
        /// </summary>
        public static string? Domain
        {
            get => _domain.Get();
            set => _domain.Set(value);
        }

        private static readonly __Value<string?> _protocol = new __Value<string?>(() => __config.Get("protocol"));
        /// <summary>
        /// The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.
        /// </summary>
        public static string? Protocol
        {
            get => _protocol.Get();
            set => _protocol.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region") ?? Utilities.GetEnv("TENCENTCLOUD_REGION"));
        /// <summary>
        /// This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION`
        /// environment variables. The default input value is ap-guangzhou.
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _secretId = new __Value<string?>(() => __config.Get("secretId") ?? Utilities.GetEnv("TENCENTCLOUD_SECRET_ID"));
        /// <summary>
        /// This is the TencentCloud access key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID`
        /// environment variable.
        /// </summary>
        public static string? SecretId
        {
            get => _secretId.Get();
            set => _secretId.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey") ?? Utilities.GetEnv("TENCENTCLOUD_SECRET_KEY"));
        /// <summary>
        /// This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY`
        /// environment variable.
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

        private static readonly __Value<string?> _securityToken = new __Value<string?>(() => __config.Get("securityToken") ?? Utilities.GetEnv("TENCENTCLOUD_SECURITY_TOKEN"));
        /// <summary>
        /// TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUD_SECURITY_TOKEN`
        /// environment variable. Notice: for supported products, please refer to: [temporary key supported
        /// products](https://intl.cloud.tencent.com/document/product/598/10588).
        /// </summary>
        public static string? SecurityToken
        {
            get => _securityToken.Get();
            set => _securityToken.Set(value);
        }

        public static class Types
        {

             public class AssumeRole
             {
                public string? Policy { get; set; } = null!;
                public string RoleArn { get; set; }
                public int SessionDuration { get; set; }
                public string SessionName { get; set; }
            }
        }
    }
}
