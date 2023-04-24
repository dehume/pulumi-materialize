// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class IndexObjNameGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The obj_name database name.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The obj_name name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The obj_name schema name.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        public IndexObjNameGetArgs()
        {
        }
        public static new IndexObjNameGetArgs Empty => new IndexObjNameGetArgs();
    }
}
