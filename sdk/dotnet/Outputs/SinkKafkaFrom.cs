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
    public sealed class SinkKafkaFrom
    {
        /// <summary>
        /// The from database name.
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// The from name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The from schema name.
        /// </summary>
        public readonly string? SchemaName;

        [OutputConstructor]
        private SinkKafkaFrom(
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
