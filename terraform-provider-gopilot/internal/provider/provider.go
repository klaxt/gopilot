// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	gopilot "terraform-provider-gopilot/internal/provider/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &ScaffoldingProvider{}

// ScaffoldingProvider defines the provider implementation.
type ScaffoldingProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ScaffoldingProviderModel describes the provider data model.
type ScaffoldingProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

// gopilotProviderModel maps provider schema data to a Go type.
type gopilotProviderModel struct {
	Host types.String `tfsdk:"host"`
}

func (p *ScaffoldingProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gopilot"
	resp.Version = p.version
}

func (p *ScaffoldingProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "The host of the provider",
				Required:            true,
			},
		},
	}
}

func (p *ScaffoldingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config gopilotProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown HashiCups API Host",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HASHICUPS_HOST environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("GOPILOT_HOST")
	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing HashiCups API Host",
			"The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API host. "+
				"Set the host value in the configuration or use the HASHICUPS_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Example client configuration for data sources and resources
	// TODO create gopilot client
	// client := http.DefaultClient
	client, err := gopilot.NewClient(&host) // Use the NewClient function to create a client instance
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create HashiCups API Client",
			"An unexpected error occurred when creating the HashiCups API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"HashiCups Client Error: "+err.Error(),
		)
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ScaffoldingProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDeviceResource,
	}
}

func (p *ScaffoldingProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDevicesDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ScaffoldingProvider{
			version: version,
		}
	}
}
