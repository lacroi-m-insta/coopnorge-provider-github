/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/coopnorge/provider-github/config/actionssecret"
	"github.com/coopnorge/provider-github/config/actionsvariable"
	"github.com/coopnorge/provider-github/config/branch"
	"github.com/coopnorge/provider-github/config/branchprotection"
	"github.com/coopnorge/provider-github/config/defaultbranch"
	"github.com/coopnorge/provider-github/config/deploykey"
	"github.com/coopnorge/provider-github/config/emugroupmapping"
	"github.com/coopnorge/provider-github/config/membership"
	"github.com/coopnorge/provider-github/config/organization"
	"github.com/coopnorge/provider-github/config/organizationruleset"
	"github.com/coopnorge/provider-github/config/pullrequest"
	"github.com/coopnorge/provider-github/config/repository"
	"github.com/coopnorge/provider-github/config/repositoryfile"
	"github.com/coopnorge/provider-github/config/repositorywebhook"
	"github.com/coopnorge/provider-github/config/team"
	"github.com/coopnorge/provider-github/config/teammembership"
	"github.com/coopnorge/provider-github/config/teamrepository"
	"github.com/coopnorge/provider-github/config/teamsettings"
	"github.com/coopnorge/provider-github/config/teamsyncgroupmapping"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "github"
	modulePath     = "github.com/coopnorge/provider-github"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("github.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		repository.Configure,
		branch.Configure,
		deploykey.Configure,
		repositoryfile.Configure,
		pullrequest.Configure,
		team.Configure,
		emugroupmapping.Configure,
		teammembership.Configure,
		teamrepository.Configure,
		defaultbranch.Configure,
		branchprotection.Configure,
		repositorywebhook.Configure,
		actionssecret.Configure,
		actionsvariable.Configure,
		organization.Configure,
		organizationruleset.Configure,
		membership.Configure,
		teamsettings.Configure,
		teamsyncgroupmapping.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
