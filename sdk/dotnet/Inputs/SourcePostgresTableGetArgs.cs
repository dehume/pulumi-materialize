// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class SourcePostgresTableGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of the table.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The name of the table.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public SourcePostgresTableGetArgs()
        {
        }
        public static new SourcePostgresTableGetArgs Empty => new SourcePostgresTableGetArgs();
    }
}
