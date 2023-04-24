// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize
{
    public static class GetCurrentDatabase
    {
        public static Task<GetCurrentDatabaseResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCurrentDatabaseResult>("materialize:index/getCurrentDatabase:GetCurrentDatabase", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetCurrentDatabaseResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetCurrentDatabaseResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
