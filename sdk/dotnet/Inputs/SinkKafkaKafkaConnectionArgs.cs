// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class SinkKafkaKafkaConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The kafka_connection database name.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The kafka_connection name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The kafka_connection schema name.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        public SinkKafkaKafkaConnectionArgs()
        {
        }
        public static new SinkKafkaKafkaConnectionArgs Empty => new SinkKafkaKafkaConnectionArgs();
    }
}
