// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize.Inputs
{

    public sealed class ConnectionKafkaSslCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ssl*certificate secret value.
        /// </summary>
        [Input("secret")]
        public Input<Inputs.ConnectionKafkaSslCertificateSecretArgs>? Secret { get; set; }

        /// <summary>
        /// The ssl_certificate text value.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public ConnectionKafkaSslCertificateArgs()
        {
        }
        public static new ConnectionKafkaSslCertificateArgs Empty => new ConnectionKafkaSslCertificateArgs();
    }
}
