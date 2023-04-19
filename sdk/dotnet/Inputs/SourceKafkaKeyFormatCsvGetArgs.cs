// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class SourceKafkaKeyFormatCsvGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("column")]
        public Input<int>? Column { get; set; }

        [Input("delimitedBy")]
        public Input<string>? DelimitedBy { get; set; }

        [Input("headers")]
        private InputList<string>? _headers;
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        public SourceKafkaKeyFormatCsvGetArgs()
        {
        }
        public static new SourceKafkaKeyFormatCsvGetArgs Empty => new SourceKafkaKeyFormatCsvGetArgs();
    }
}
