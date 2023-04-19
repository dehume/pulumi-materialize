// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Materialize
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

        private static readonly Pulumi.Config __config = new Pulumi.Config("materialize");

        private static readonly __Value<string?> _database = new __Value<string?>(() => __config.Get("database"));
        /// <summary>
        /// The Materialize database
        /// </summary>
        public static string? Database
        {
            get => _database.Get();
            set => _database.Set(value);
        }

        private static readonly __Value<string?> _host = new __Value<string?>(() => __config.Get("host"));
        /// <summary>
        /// Materialize host
        /// </summary>
        public static string? Host
        {
            get => _host.Get();
            set => _host.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password"));
        /// <summary>
        /// Materialize host
        /// </summary>
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<int?> _port = new __Value<int?>(() => __config.GetInt32("port"));
        /// <summary>
        /// The Materialize port number to connect to at the server host
        /// </summary>
        public static int? Port
        {
            get => _port.Get();
            set => _port.Set(value);
        }

        private static readonly __Value<bool?> _testing = new __Value<bool?>(() => __config.GetBoolean("testing"));
        /// <summary>
        /// Enable to test the provider locally
        /// </summary>
        public static bool? Testing
        {
            get => _testing.Get();
            set => _testing.Set(value);
        }

        private static readonly __Value<string?> _username = new __Value<string?>(() => __config.Get("username"));
        /// <summary>
        /// Materialize username
        /// </summary>
        public static string? Username
        {
            get => _username.Get();
            set => _username.Set(value);
        }

    }
}
