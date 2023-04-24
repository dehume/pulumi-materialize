// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Materialize
{
    /// <summary>
    /// A cluster replica is the physical resource which maintains dataflow-powered objects.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Materialize = Pulumi.Materialize;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleClusterReplica = new Materialize.ClusterReplica("exampleClusterReplica", new()
    ///     {
    ///         ClusterName = "cluster",
    ///         Size = "2xsmall",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// # Clusters can be imported using the cluster id
    /// 
    /// ```sh
    ///  $ pulumi import materialize:index/clusterReplica:ClusterReplica example_1_cluster_replica &lt;cluster_replica_id&gt;
    /// ```
    /// </summary>
    [MaterializeResourceType("materialize:index/clusterReplica:ClusterReplica")]
    public partial class ClusterReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If you want the replica to reside in a specific availability zone.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The cluster whose resources you want to create an additional computation of.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The amount of effort the replica should exert on compacting arrangements during idle periods. This is an unstable option! It may be changed or removed at any time.
        /// </summary>
        [Output("idleArrangementMergeEffort")]
        public Output<int?> IdleArrangementMergeEffort { get; private set; } = null!;

        /// <summary>
        /// Whether to introspect the gathering of the introspection data.
        /// </summary>
        [Output("introspectionDebugging")]
        public Output<bool?> IntrospectionDebugging { get; private set; } = null!;

        /// <summary>
        /// The interval at which to collect introspection data.
        /// </summary>
        [Output("introspectionInterval")]
        public Output<string?> IntrospectionInterval { get; private set; } = null!;

        /// <summary>
        /// The identifier for the replica.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The size of the replica.
        /// </summary>
        [Output("size")]
        public Output<string> Size { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterReplica(string name, ClusterReplicaArgs args, CustomResourceOptions? options = null)
            : base("materialize:index/clusterReplica:ClusterReplica", name, args ?? new ClusterReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterReplica(string name, Input<string> id, ClusterReplicaState? state = null, CustomResourceOptions? options = null)
            : base("materialize:index/clusterReplica:ClusterReplica", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/MaterializeInc/pulumi-materialize",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterReplica Get(string name, Input<string> id, ClusterReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterReplica(name, id, state, options);
        }
    }

    public sealed class ClusterReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you want the replica to reside in a specific availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The cluster whose resources you want to create an additional computation of.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The amount of effort the replica should exert on compacting arrangements during idle periods. This is an unstable option! It may be changed or removed at any time.
        /// </summary>
        [Input("idleArrangementMergeEffort")]
        public Input<int>? IdleArrangementMergeEffort { get; set; }

        /// <summary>
        /// Whether to introspect the gathering of the introspection data.
        /// </summary>
        [Input("introspectionDebugging")]
        public Input<bool>? IntrospectionDebugging { get; set; }

        /// <summary>
        /// The interval at which to collect introspection data.
        /// </summary>
        [Input("introspectionInterval")]
        public Input<string>? IntrospectionInterval { get; set; }

        /// <summary>
        /// The identifier for the replica.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the replica.
        /// </summary>
        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        public ClusterReplicaArgs()
        {
        }
        public static new ClusterReplicaArgs Empty => new ClusterReplicaArgs();
    }

    public sealed class ClusterReplicaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you want the replica to reside in a specific availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The cluster whose resources you want to create an additional computation of.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The amount of effort the replica should exert on compacting arrangements during idle periods. This is an unstable option! It may be changed or removed at any time.
        /// </summary>
        [Input("idleArrangementMergeEffort")]
        public Input<int>? IdleArrangementMergeEffort { get; set; }

        /// <summary>
        /// Whether to introspect the gathering of the introspection data.
        /// </summary>
        [Input("introspectionDebugging")]
        public Input<bool>? IntrospectionDebugging { get; set; }

        /// <summary>
        /// The interval at which to collect introspection data.
        /// </summary>
        [Input("introspectionInterval")]
        public Input<string>? IntrospectionInterval { get; set; }

        /// <summary>
        /// The identifier for the replica.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the replica.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        public ClusterReplicaState()
        {
        }
        public static new ClusterReplicaState Empty => new ClusterReplicaState();
    }
}
