// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The connection resource allows you to manage connections in Materialize.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as materialize from "@pulumi/materialize";
 *
 * // Create a Postgres Connection
 * const examplePostgresConnection = new materialize.ConnectionPostgres("example_postgres_connection", {
 *     database: "example",
 *     host: "instance.foo000.us-west-1.rds.amazonaws.com",
 *     password: {
 *         databaseName: "database",
 *         name: "example",
 *         schemaName: "schema",
 *     },
 *     port: 5432,
 *     user: {
 *         secret: {
 *             databaseName: "database",
 *             name: "example",
 *             schemaName: "schema",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * #Connections can be imported using the connection id
 *
 * ```sh
 *  $ pulumi import materialize:index/connectionPostgres:ConnectionPostgres example <connection_id>
 * ```
 */
export class ConnectionPostgres extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionPostgres resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionPostgresState, opts?: pulumi.CustomResourceOptions): ConnectionPostgres {
        return new ConnectionPostgres(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'materialize:index/connectionPostgres:ConnectionPostgres';

    /**
     * Returns true if the given object is an instance of ConnectionPostgres.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionPostgres {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionPostgres.__pulumiType;
    }

    /**
     * The AWS PrivateLink configuration for the Postgres database.
     */
    public readonly awsPrivatelink!: pulumi.Output<outputs.ConnectionPostgresAwsPrivatelink | undefined>;
    /**
     * The type of connection.
     */
    public /*out*/ readonly connectionType!: pulumi.Output<string>;
    /**
     * The target Postgres database.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The identifier for the connection database.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * The Postgres database hostname.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The identifier for the connection.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Postgres database password.
     */
    public readonly password!: pulumi.Output<outputs.ConnectionPostgresPassword | undefined>;
    /**
     * The Postgres database port.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The fully qualified name of the connection.
     */
    public /*out*/ readonly qualifiedSqlName!: pulumi.Output<string>;
    /**
     * The identifier for the connection schema.
     */
    public readonly schemaName!: pulumi.Output<string | undefined>;
    /**
     * The SSH tunnel configuration for the Postgres database.
     */
    public readonly sshTunnel!: pulumi.Output<outputs.ConnectionPostgresSshTunnel | undefined>;
    /**
     * The client certificate for the Postgres database.
     */
    public readonly sslCertificate!: pulumi.Output<outputs.ConnectionPostgresSslCertificate | undefined>;
    /**
     * The CA certificate for the Postgres database.
     */
    public readonly sslCertificateAuthority!: pulumi.Output<outputs.ConnectionPostgresSslCertificateAuthority | undefined>;
    /**
     * The client key for the Postgres database.
     */
    public readonly sslKey!: pulumi.Output<outputs.ConnectionPostgresSslKey | undefined>;
    /**
     * The SSL mode for the Postgres database.
     */
    public readonly sslMode!: pulumi.Output<string | undefined>;
    /**
     * The Postgres database username.
     */
    public readonly user!: pulumi.Output<outputs.ConnectionPostgresUser>;

    /**
     * Create a ConnectionPostgres resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionPostgresArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionPostgresArgs | ConnectionPostgresState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionPostgresState | undefined;
            resourceInputs["awsPrivatelink"] = state ? state.awsPrivatelink : undefined;
            resourceInputs["connectionType"] = state ? state.connectionType : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["qualifiedSqlName"] = state ? state.qualifiedSqlName : undefined;
            resourceInputs["schemaName"] = state ? state.schemaName : undefined;
            resourceInputs["sshTunnel"] = state ? state.sshTunnel : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["sslCertificateAuthority"] = state ? state.sslCertificateAuthority : undefined;
            resourceInputs["sslKey"] = state ? state.sslKey : undefined;
            resourceInputs["sslMode"] = state ? state.sslMode : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as ConnectionPostgresArgs | undefined;
            if ((!args || args.database === undefined) && !opts.urn) {
                throw new Error("Missing required property 'database'");
            }
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["awsPrivatelink"] = args ? args.awsPrivatelink : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["schemaName"] = args ? args.schemaName : undefined;
            resourceInputs["sshTunnel"] = args ? args.sshTunnel : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["sslCertificateAuthority"] = args ? args.sslCertificateAuthority : undefined;
            resourceInputs["sslKey"] = args ? args.sslKey : undefined;
            resourceInputs["sslMode"] = args ? args.sslMode : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["connectionType"] = undefined /*out*/;
            resourceInputs["qualifiedSqlName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConnectionPostgres.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectionPostgres resources.
 */
export interface ConnectionPostgresState {
    /**
     * The AWS PrivateLink configuration for the Postgres database.
     */
    awsPrivatelink?: pulumi.Input<inputs.ConnectionPostgresAwsPrivatelink>;
    /**
     * The type of connection.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * The target Postgres database.
     */
    database?: pulumi.Input<string>;
    /**
     * The identifier for the connection database.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The Postgres database hostname.
     */
    host?: pulumi.Input<string>;
    /**
     * The identifier for the connection.
     */
    name?: pulumi.Input<string>;
    /**
     * The Postgres database password.
     */
    password?: pulumi.Input<inputs.ConnectionPostgresPassword>;
    /**
     * The Postgres database port.
     */
    port?: pulumi.Input<number>;
    /**
     * The fully qualified name of the connection.
     */
    qualifiedSqlName?: pulumi.Input<string>;
    /**
     * The identifier for the connection schema.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * The SSH tunnel configuration for the Postgres database.
     */
    sshTunnel?: pulumi.Input<inputs.ConnectionPostgresSshTunnel>;
    /**
     * The client certificate for the Postgres database.
     */
    sslCertificate?: pulumi.Input<inputs.ConnectionPostgresSslCertificate>;
    /**
     * The CA certificate for the Postgres database.
     */
    sslCertificateAuthority?: pulumi.Input<inputs.ConnectionPostgresSslCertificateAuthority>;
    /**
     * The client key for the Postgres database.
     */
    sslKey?: pulumi.Input<inputs.ConnectionPostgresSslKey>;
    /**
     * The SSL mode for the Postgres database.
     */
    sslMode?: pulumi.Input<string>;
    /**
     * The Postgres database username.
     */
    user?: pulumi.Input<inputs.ConnectionPostgresUser>;
}

/**
 * The set of arguments for constructing a ConnectionPostgres resource.
 */
export interface ConnectionPostgresArgs {
    /**
     * The AWS PrivateLink configuration for the Postgres database.
     */
    awsPrivatelink?: pulumi.Input<inputs.ConnectionPostgresAwsPrivatelink>;
    /**
     * The target Postgres database.
     */
    database: pulumi.Input<string>;
    /**
     * The identifier for the connection database.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The Postgres database hostname.
     */
    host: pulumi.Input<string>;
    /**
     * The identifier for the connection.
     */
    name?: pulumi.Input<string>;
    /**
     * The Postgres database password.
     */
    password?: pulumi.Input<inputs.ConnectionPostgresPassword>;
    /**
     * The Postgres database port.
     */
    port?: pulumi.Input<number>;
    /**
     * The identifier for the connection schema.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * The SSH tunnel configuration for the Postgres database.
     */
    sshTunnel?: pulumi.Input<inputs.ConnectionPostgresSshTunnel>;
    /**
     * The client certificate for the Postgres database.
     */
    sslCertificate?: pulumi.Input<inputs.ConnectionPostgresSslCertificate>;
    /**
     * The CA certificate for the Postgres database.
     */
    sslCertificateAuthority?: pulumi.Input<inputs.ConnectionPostgresSslCertificateAuthority>;
    /**
     * The client key for the Postgres database.
     */
    sslKey?: pulumi.Input<inputs.ConnectionPostgresSslKey>;
    /**
     * The SSL mode for the Postgres database.
     */
    sslMode?: pulumi.Input<string>;
    /**
     * The Postgres database username.
     */
    user: pulumi.Input<inputs.ConnectionPostgresUser>;
}
