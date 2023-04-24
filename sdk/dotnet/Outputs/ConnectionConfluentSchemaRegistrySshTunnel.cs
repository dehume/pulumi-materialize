// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Outputs
{

    [OutputType]
    public sealed class ConnectionConfluentSchemaRegistrySshTunnel
    {
        /// <summary>
        /// The ssh_tunnel database name.
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// The ssh_tunnel name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ssh_tunnel schema name.
        /// </summary>
        public readonly string? SchemaName;

        [OutputConstructor]
        private ConnectionConfluentSchemaRegistrySshTunnel(
            string? databaseName,

            string name,

            string? schemaName)
        {
            DatabaseName = databaseName;
            Name = name;
            SchemaName = schemaName;
        }
    }
}
