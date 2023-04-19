// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class SinkKafkaEnvelopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The generated schemas have a Debezium-style diff envelope to capture changes in the input view or source.
        /// </summary>
        [Input("debezium")]
        public Input<bool>? Debezium { get; set; }

        /// <summary>
        /// The sink emits data with upsert semantics: updates and inserts for the given key are expressed as a value, and deletes are expressed as a null value payload in Kafka.
        /// </summary>
        [Input("upsert")]
        public Input<bool>? Upsert { get; set; }

        public SinkKafkaEnvelopeArgs()
        {
        }
        public static new SinkKafkaEnvelopeArgs Empty => new SinkKafkaEnvelopeArgs();
    }
}
