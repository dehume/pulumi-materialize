// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as materialize from "@pulumi/materialize";
 *
 * const all = pulumi.output(materialize.GetSchemas());
 * const materializeGetSchemas = pulumi.output(materialize.GetSchemas({
 *     databaseName: "materialize",
 * }));
 * ```
 */
export function getSchemas(args?: GetSchemasArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemasResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("materialize:index/getSchemas:GetSchemas", {
        "databaseName": args.databaseName,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSchemas.
 */
export interface GetSchemasArgs {
    /**
     * Limit schemas to a specific database
     */
    databaseName?: string;
}

/**
 * A collection of values returned by GetSchemas.
 */
export interface GetSchemasResult {
    /**
     * Limit schemas to a specific database
     */
    readonly databaseName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The schemas in the account
     */
    readonly schemas: outputs.GetSchemasSchema[];
}

export function getSchemasOutput(args?: GetSchemasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemasResult> {
    return pulumi.output(args).apply(a => getSchemas(a, opts))
}

/**
 * A collection of arguments for invoking GetSchemas.
 */
export interface GetSchemasOutputArgs {
    /**
     * Limit schemas to a specific database
     */
    databaseName?: pulumi.Input<string>;
}
