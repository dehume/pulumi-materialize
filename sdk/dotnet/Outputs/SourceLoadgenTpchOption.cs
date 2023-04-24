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
    public sealed class SourceLoadgenTpchOption
    {
        /// <summary>
        /// The scale factor for the generator. Defaults to 0.01 (~ 10MB).
        /// </summary>
        public readonly double? ScaleFactor;
        /// <summary>
        /// Creates subsources for specific tables.
        /// </summary>
        public readonly ImmutableArray<Outputs.SourceLoadgenTpchOptionTable> Tables;
        /// <summary>
        /// The interval at which the next datum should be emitted. Defaults to one second.
        /// </summary>
        public readonly string? TickInterval;

        [OutputConstructor]
        private SourceLoadgenTpchOption(
            double? scaleFactor,

            ImmutableArray<Outputs.SourceLoadgenTpchOptionTable> tables,

            string? tickInterval)
        {
            ScaleFactor = scaleFactor;
            Tables = tables;
            TickInterval = tickInterval;
        }
    }
}
