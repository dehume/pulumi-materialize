// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class SourceKafkaValueFormatProtobufGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("message", required: true)]
        public Input<string> Message { get; set; } = null!;

        [Input("schemaRegistryConnection", required: true)]
        public Input<Inputs.SourceKafkaValueFormatProtobufSchemaRegistryConnectionGetArgs> SchemaRegistryConnection { get; set; } = null!;

        public SourceKafkaValueFormatProtobufGetArgs()
        {
        }
        public static new SourceKafkaValueFormatProtobufGetArgs Empty => new SourceKafkaValueFormatProtobufGetArgs();
    }
}
