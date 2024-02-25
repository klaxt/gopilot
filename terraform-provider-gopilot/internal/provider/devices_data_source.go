package provider

import (
	"context"
	"fmt"
	"strconv"

	gopilot "terraform-provider-gopilot/internal/provider/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &devicesDataSource{}
	_ datasource.DataSourceWithConfigure = &devicesDataSource{}
)

func NewDevicesDataSource() datasource.DataSource {
	return &devicesDataSource{}
}

type devicesDataSource struct {
	client *gopilot.Client
}

func (d *devicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *devicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"model": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"color": schema.StringAttribute{ // Added color attribute
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *devicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state devicesDataSourceModel

	devices, err := d.client.GetDevices()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read HashiCups Coffees",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, device := range devices {
		tflog.Info(ctx, "--Add device--"+device.Name)

		deviceState := deviceModel{
			ID:     types.Int64Value(int64(device.ID)),
			Name:   types.StringValue(device.Name),
			Model:  types.StringValue(device.Model),
			Color:  types.StringValue(device.Color),
			Status: types.StringValue(device.Status),
		}

		state.Devices = append(state.Devices, deviceState)
	}

	deviceMap := make(map[string]interface{})
	for i, device := range state.Devices {
		deviceMap[strconv.Itoa(i)] = device
	}

	tflog.Info(ctx, "--devices--", deviceMap)

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *devicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*gopilot.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

type devicesDataSourceModel struct {
	Devices []deviceModel `tfsdk:"devices"`
}
type deviceModel struct {
	ID     types.Int64  `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Model  types.String `tfsdk:"model"`
	Status types.String `tfsdk:"status"`
	Color  types.String `tfsdk:"color"`
}
