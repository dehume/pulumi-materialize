// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class ConnectionKafkaKafkaBrokerPrivatelinkConnectionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for the connection database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The identifier for the connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The identifier for the connection schema.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        public ConnectionKafkaKafkaBrokerPrivatelinkConnectionGetArgs()
        {
        }
        public static new ConnectionKafkaKafkaBrokerPrivatelinkConnectionGetArgs Empty => new ConnectionKafkaKafkaBrokerPrivatelinkConnectionGetArgs();
    }
}
