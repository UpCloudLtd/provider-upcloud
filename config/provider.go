/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/UpCloudLtd/provider-upcloud/config/database"
	"github.com/UpCloudLtd/provider-upcloud/config/kubernetes"
	"github.com/UpCloudLtd/provider-upcloud/config/network"
	"github.com/UpCloudLtd/provider-upcloud/config/objectstorage"
	"github.com/UpCloudLtd/provider-upcloud/config/server"
	"github.com/UpCloudLtd/provider-upcloud/config/storage"

	"github.com/UpCloudLtd/terraform-provider-upcloud/upcloud"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "upcloud"
	modulePath     = "github.com/UpCloudLtd/provider-upcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upcloud.io"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(sdkResourcesList()),
		ujconfig.WithTerraformProvider(upcloud.Provider()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(pluginFrameworkResourcesList()),
		ujconfig.WithTerraformPluginFrameworkProvider(upcloud.New()),
		ujconfig.WithFeaturesPackage("internal/features"),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		database.Configure,
		kubernetes.Configure,
		network.Configure,
		objectstorage.Configure,
		server.Configure,
		storage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func pluginFrameworkResourcesList() []string {
	return formatResourceNames([][]string{
		database.PluginFrameworkResources,
		kubernetes.PluginFrameworkResources,
		network.PluginFrameworkResources,
		objectstorage.PluginFrameworkResources,
		server.PluginFrameworkResources,
		storage.PluginFrameworkResources,
	})
}

func sdkResourcesList() []string {
	return formatResourceNames([][]string{
		database.SDKResources,
		kubernetes.SDKResources,
		network.SDKResources,
		objectstorage.SDKResources,
		server.SDKResources,
		storage.SDKResources,
	})
}

func formatResourceNames(resources [][]string) []string {
	var formatted []string
	for _, r := range resources {
		for _, name := range r {
			// $ is added to match the exact string since the format is regex.
			formatted = append(formatted, name+"$")
		}
	}

	return formatted
}
