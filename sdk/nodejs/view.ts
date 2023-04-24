// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * A non-materialized view, provides an alias for the embedded SELECT statement.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as materialize from "@pulumi/materialize";
 *
 * const simpleViewView = new materialize.View("simpleViewView", {
 *     schemaName: materialize_schema.schema.name,
 *     databaseName: materialize_database.database.name,
 *     statement: `SELECT
 *     *
 * FROM
 *     ${materialize_table.simple_table.qualified_name}
 * `,
 * }, {
 *     dependsOn: [materialize_table.simple_table],
 * });
 * const simpleViewIndex_viewView = new materialize.View("simpleViewIndex/viewView", {
 *     schemaName: materialize_schema.schema.name,
 *     databaseName: materialize_database.database.name,
 *     statement: "SELECT * FROM materialize.public.simple_table",
 * });
 * ```
 *
 * ## Import
 *
 * # Views can be imported using the source id
 *
 * ```sh
 *  $ pulumi import materialize:index/view:View example_view <view_id>
 * ```
 */
export class View extends pulumi.CustomResource {
    /**
     * Get an existing View resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ViewState, opts?: pulumi.CustomResourceOptions): View {
        return new View(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'materialize:index/view:View';

    /**
     * Returns true if the given object is an instance of View.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is View {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === View.__pulumiType;
    }

    /**
     * The identifier for the view database.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * The identifier for the view.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The fully qualified name of the view.
     */
    public /*out*/ readonly qualifiedSqlName!: pulumi.Output<string>;
    /**
     * The identifier for the view schema.
     */
    public readonly schemaName!: pulumi.Output<string | undefined>;
    /**
     * The SQL statement to create the view.
     */
    public readonly statement!: pulumi.Output<string>;

    /**
     * Create a View resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ViewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ViewArgs | ViewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ViewState | undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["qualifiedSqlName"] = state ? state.qualifiedSqlName : undefined;
            resourceInputs["schemaName"] = state ? state.schemaName : undefined;
            resourceInputs["statement"] = state ? state.statement : undefined;
        } else {
            const args = argsOrState as ViewArgs | undefined;
            if ((!args || args.statement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statement'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["schemaName"] = args ? args.schemaName : undefined;
            resourceInputs["statement"] = args ? args.statement : undefined;
            resourceInputs["qualifiedSqlName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(View.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering View resources.
 */
export interface ViewState {
    /**
     * The identifier for the view database.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The identifier for the view.
     */
    name?: pulumi.Input<string>;
    /**
     * The fully qualified name of the view.
     */
    qualifiedSqlName?: pulumi.Input<string>;
    /**
     * The identifier for the view schema.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * The SQL statement to create the view.
     */
    statement?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a View resource.
 */
export interface ViewArgs {
    /**
     * The identifier for the view database.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The identifier for the view.
     */
    name?: pulumi.Input<string>;
    /**
     * The identifier for the view schema.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * The SQL statement to create the view.
     */
    statement: pulumi.Input<string>;
}
