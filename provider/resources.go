// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package materialize

import (
	"fmt"
	"path/filepath"

	"provider/pkg/version"

	materialize "github.com/MaterializeInc/terraform-provider-materialize/pkg"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "materialize"
	// modules:
	mainMod = "index" // the materialize module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(materialize.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "materialize",
		Description:          "A Pulumi package for creating and managing materialize cloud resources.",
		Keywords:             []string{"pulumi", "materialize", "category/cloud"},
		License:              "Apache-2.0",
		Publisher:            "Materialize Inc",
		LogoURL:              "https://raw.githubusercontent.com/MaterializeInc/pulumi-materialize/main/assets/materialize.svg",
		GitHubOrg:            "MaterializeInc",
		Homepage:             "https://github.com/MaterializeInc/terraform-provider-materialize",
		Repository:           "https://github.com/MaterializeInc/terraform-provider-materialize",
		PluginDownloadURL:    "github://api.github.com/MaterializeInc/pulumi-materialize",
		DisplayName:          "Materialize",
		PreConfigureCallback: preConfigureCallback,
		Config:               map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"materialize_cluster":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cluster")},
			"materialize_cluster_replica":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ClusterReplica")},
			"materialize_connection_aws_privatelink":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectionAwsPrivatelink")},
			"materialize_connection_confluent_schema_registry": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectionConfluentSchemaRegistry")},
			"materialize_connection_kafka":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectionKafka")},
			"materialize_connection_postgres":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectionPostgres")},
			"materialize_connection_ssh_tunnel":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ConnectionSshTunnel")},
			"materialize_database":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Database")},
			"materialize_index":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Index")},
			"materialize_materialized_view":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MaterializedView")},
			"materialize_schema":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Schema")},
			"materialize_secret":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Secret")},
			"materialize_sink_kafka":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SinkKafka")},
			"materialize_source_kafka":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SourceKafka")},
			"materialize_source_load_generator":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SourceLoadgen")},
			"materialize_source_postgres":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SourcePostgres")},
			"materialize_table":                                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Table")},
			"materialize_view":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "View")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"materialize_cluster":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetClusters")},
			"materialize_cluster_replica":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetClusterReplicas")},
			"materialize_connection":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetConnections")},
			"materialize_current_database":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetCurrentDatabase")},
			"materialize_current_cluster":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetCurrentCluster")},
			"materialize_database":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetDatabases")},
			"materialize_index":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetIndexes")},
			"materialize_materialized_view": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMaterializedViews")},
			"materialize_schema":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetSchemas")},
			"materialize_secret":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetSecrets")},
			"materialize_sink":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetSinks")},
			"materialize_source":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetSources")},
			"materialize_table":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetTables")},
			"materialize_view":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetViews")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
