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
    /// The connection resource allows you to manage connections in Materialize.
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
    ///     // Create a Postgres Connection
    ///     var examplePostgresConnection = new Materialize.ConnectionPostgres("examplePostgresConnection", new()
    ///     {
    ///         Database = "example",
    ///         Host = "instance.foo000.us-west-1.rds.amazonaws.com",
    ///         Password = new Materialize.Inputs.ConnectionPostgresPasswordArgs
    ///         {
    ///             DatabaseName = "database",
    ///             Name = "example",
    ///             SchemaName = "schema",
    ///         },
    ///         Port = 5432,
    ///         User = new Materialize.Inputs.ConnectionPostgresUserArgs
    ///         {
    ///             Secret = new Materialize.Inputs.ConnectionPostgresUserSecretArgs
    ///             {
    ///                 DatabaseName = "database",
    ///                 Name = "example",
    ///                 SchemaName = "schema",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #Connections can be imported using the connection id
    /// 
    /// ```sh
    ///  $ pulumi import materialize:index/connectionPostgres:ConnectionPostgres example &lt;connection_id&gt;
    /// ```
    /// </summary>
    [MaterializeResourceType("materialize:index/connectionPostgres:ConnectionPostgres")]
    public partial class ConnectionPostgres : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS PrivateLink configuration for the Postgres database.
        /// </summary>
        [Output("awsPrivatelink")]
        public Output<Outputs.ConnectionPostgresAwsPrivatelink?> AwsPrivatelink { get; private set; } = null!;

        /// <summary>
        /// The type of connection.
        /// </summary>
        [Output("connectionType")]
        public Output<string> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// The target Postgres database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The identifier for the connection database.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The Postgres database hostname.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The identifier for the connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Postgres database password.
        /// </summary>
        [Output("password")]
        public Output<Outputs.ConnectionPostgresPassword?> Password { get; private set; } = null!;

        /// <summary>
        /// The Postgres database port.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of the connection.
        /// </summary>
        [Output("qualifiedSqlName")]
        public Output<string> QualifiedSqlName { get; private set; } = null!;

        /// <summary>
        /// The identifier for the connection schema.
        /// </summary>
        [Output("schemaName")]
        public Output<string?> SchemaName { get; private set; } = null!;

        /// <summary>
        /// The SSH tunnel configuration for the Postgres database.
        /// </summary>
        [Output("sshTunnel")]
        public Output<Outputs.ConnectionPostgresSshTunnel?> SshTunnel { get; private set; } = null!;

        /// <summary>
        /// The client certificate for the Postgres database.
        /// </summary>
        [Output("sslCertificate")]
        public Output<Outputs.ConnectionPostgresSslCertificate?> SslCertificate { get; private set; } = null!;

        /// <summary>
        /// The CA certificate for the Postgres database.
        /// </summary>
        [Output("sslCertificateAuthority")]
        public Output<Outputs.ConnectionPostgresSslCertificateAuthority?> SslCertificateAuthority { get; private set; } = null!;

        /// <summary>
        /// The client key for the Postgres database.
        /// </summary>
        [Output("sslKey")]
        public Output<Outputs.ConnectionPostgresSslKey?> SslKey { get; private set; } = null!;

        /// <summary>
        /// The SSL mode for the Postgres database.
        /// </summary>
        [Output("sslMode")]
        public Output<string?> SslMode { get; private set; } = null!;

        /// <summary>
        /// The Postgres database username.
        /// </summary>
        [Output("user")]
        public Output<Outputs.ConnectionPostgresUser> User { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectionPostgres resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionPostgres(string name, ConnectionPostgresArgs args, CustomResourceOptions? options = null)
            : base("materialize:index/connectionPostgres:ConnectionPostgres", name, args ?? new ConnectionPostgresArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionPostgres(string name, Input<string> id, ConnectionPostgresState? state = null, CustomResourceOptions? options = null)
            : base("materialize:index/connectionPostgres:ConnectionPostgres", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectionPostgres resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionPostgres Get(string name, Input<string> id, ConnectionPostgresState? state = null, CustomResourceOptions? options = null)
        {
            return new ConnectionPostgres(name, id, state, options);
        }
    }

    public sealed class ConnectionPostgresArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS PrivateLink configuration for the Postgres database.
        /// </summary>
        [Input("awsPrivatelink")]
        public Input<Inputs.ConnectionPostgresAwsPrivatelinkArgs>? AwsPrivatelink { get; set; }

        /// <summary>
        /// The target Postgres database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The identifier for the connection database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The Postgres database hostname.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The identifier for the connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Postgres database password.
        /// </summary>
        [Input("password")]
        public Input<Inputs.ConnectionPostgresPasswordArgs>? Password { get; set; }

        /// <summary>
        /// The Postgres database port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The identifier for the connection schema.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        /// <summary>
        /// The SSH tunnel configuration for the Postgres database.
        /// </summary>
        [Input("sshTunnel")]
        public Input<Inputs.ConnectionPostgresSshTunnelArgs>? SshTunnel { get; set; }

        /// <summary>
        /// The client certificate for the Postgres database.
        /// </summary>
        [Input("sslCertificate")]
        public Input<Inputs.ConnectionPostgresSslCertificateArgs>? SslCertificate { get; set; }

        /// <summary>
        /// The CA certificate for the Postgres database.
        /// </summary>
        [Input("sslCertificateAuthority")]
        public Input<Inputs.ConnectionPostgresSslCertificateAuthorityArgs>? SslCertificateAuthority { get; set; }

        /// <summary>
        /// The client key for the Postgres database.
        /// </summary>
        [Input("sslKey")]
        public Input<Inputs.ConnectionPostgresSslKeyArgs>? SslKey { get; set; }

        /// <summary>
        /// The SSL mode for the Postgres database.
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        /// <summary>
        /// The Postgres database username.
        /// </summary>
        [Input("user", required: true)]
        public Input<Inputs.ConnectionPostgresUserArgs> User { get; set; } = null!;

        public ConnectionPostgresArgs()
        {
        }
        public static new ConnectionPostgresArgs Empty => new ConnectionPostgresArgs();
    }

    public sealed class ConnectionPostgresState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS PrivateLink configuration for the Postgres database.
        /// </summary>
        [Input("awsPrivatelink")]
        public Input<Inputs.ConnectionPostgresAwsPrivatelinkGetArgs>? AwsPrivatelink { get; set; }

        /// <summary>
        /// The type of connection.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// The target Postgres database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The identifier for the connection database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The Postgres database hostname.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The identifier for the connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Postgres database password.
        /// </summary>
        [Input("password")]
        public Input<Inputs.ConnectionPostgresPasswordGetArgs>? Password { get; set; }

        /// <summary>
        /// The Postgres database port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The fully qualified name of the connection.
        /// </summary>
        [Input("qualifiedSqlName")]
        public Input<string>? QualifiedSqlName { get; set; }

        /// <summary>
        /// The identifier for the connection schema.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        /// <summary>
        /// The SSH tunnel configuration for the Postgres database.
        /// </summary>
        [Input("sshTunnel")]
        public Input<Inputs.ConnectionPostgresSshTunnelGetArgs>? SshTunnel { get; set; }

        /// <summary>
        /// The client certificate for the Postgres database.
        /// </summary>
        [Input("sslCertificate")]
        public Input<Inputs.ConnectionPostgresSslCertificateGetArgs>? SslCertificate { get; set; }

        /// <summary>
        /// The CA certificate for the Postgres database.
        /// </summary>
        [Input("sslCertificateAuthority")]
        public Input<Inputs.ConnectionPostgresSslCertificateAuthorityGetArgs>? SslCertificateAuthority { get; set; }

        /// <summary>
        /// The client key for the Postgres database.
        /// </summary>
        [Input("sslKey")]
        public Input<Inputs.ConnectionPostgresSslKeyGetArgs>? SslKey { get; set; }

        /// <summary>
        /// The SSL mode for the Postgres database.
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        /// <summary>
        /// The Postgres database username.
        /// </summary>
        [Input("user")]
        public Input<Inputs.ConnectionPostgresUserGetArgs>? User { get; set; }

        public ConnectionPostgresState()
        {
        }
        public static new ConnectionPostgresState Empty => new ConnectionPostgresState();
    }
}
