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
    public sealed class SourceKafkaValueFormatProtobuf
    {
        public readonly string Message;
        public readonly Outputs.SourceKafkaValueFormatProtobufSchemaRegistryConnection SchemaRegistryConnection;

        [OutputConstructor]
        private SourceKafkaValueFormatProtobuf(
            string message,

            Outputs.SourceKafkaValueFormatProtobufSchemaRegistryConnection schemaRegistryConnection)
        {
            Message = message;
            SchemaRegistryConnection = schemaRegistryConnection;
        }
    }
}
