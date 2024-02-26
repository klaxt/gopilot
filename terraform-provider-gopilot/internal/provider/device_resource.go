package provider

import (
	"context"
	"fmt"
	"strconv"
	gopilot "terraform-provider-gopilot/internal/provider/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &deviceResource{}
	_ resource.ResourceWithConfigure = &deviceResource{}
)

// Configure adds the provider configured client to the resource.
func (r *deviceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

// NewDeviceResource is a helper function to simplify the provider implementation.
func NewDeviceResource() resource.Resource {
	return &deviceResource{}
}

// deviceResource is the resource implementation.
type deviceResource struct {
	client *gopilot.Client
}

// Metadata returns the resource type name.
func (r *deviceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

// Schema defines the schema for the resource.
func (r *deviceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"model": schema.StringAttribute{
				Required: true,
			},
			"status": schema.StringAttribute{
				Required: true,
			},
			"color": schema.StringAttribute{ // Added color attribute
				Required: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *deviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan deviceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var planDevice = gopilot.Device{
		Name:   plan.Name.ValueString(),
		Model:  plan.Model.ValueString(),
		Status: plan.Status.ValueString(),
		Color:  plan.Color.ValueString(),
	}

	// Create new order
	device, err := r.client.CreateDevice(planDevice)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	// plan.ID = types.StringValue(strconv.Itoa(order.ID))
	plan.ID = types.Int64Value(int64(device.ID))
	plan.Name = types.StringValue(device.Name)
	plan.Model = types.StringValue(device.Model)
	plan.Color = types.StringValue(device.Color)
	plan.Status = types.StringValue(device.Status)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *deviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state deviceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from HashiCups
	device, err := r.client.GetDevice(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading HashiCups Order",
			"Could not read HashiCups order ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state = deviceModel{
		ID:     types.Int64Value(int64(device.ID)),
		Name:   types.StringValue(device.Name),
		Model:  types.StringValue(device.Model),
		Color:  types.StringValue(device.Color),
		Status: types.StringValue(device.Status),
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	tflog.Info(ctx, fmt.Sprintf("--Storing Read State--%d %s", state.ID, state.Name))
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *deviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan deviceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var currentState deviceModel
	diags2 := req.State.Get(ctx, &currentState)
	resp.Diagnostics.Append(diags2...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, fmt.Sprintf("--Update Current Context--%d %s", currentState.ID.ValueInt64(), currentState.Name))
	tflog.Info(ctx, fmt.Sprintf("--Update Plan Context--%d %s", plan.ID, plan.Name))

	// Generate API request body from plan
	var planDevice = gopilot.Device{
		ID:     int(currentState.ID.ValueInt64()),
		Name:   plan.Name.ValueString(),
		Model:  plan.Model.ValueString(),
		Status: plan.Status.ValueString(),
		Color:  plan.Color.ValueString(),
	}

	tflog.Info(ctx, fmt.Sprintf("--Update Device--%d", currentState.ID.ValueInt64()))
	// Update existing order
	_, err := r.client.UpdateDevice(currentState.ID.ValueInt64(), planDevice)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating HashiCups Order",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not
	// populated.
	tflog.Info(ctx, "--Get updated device--")
	device, err := r.client.GetDevice(int64(planDevice.ID))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading HashiCups Order",
			"Could not read HashiCups order ID "+strconv.FormatInt(int64(planDevice.ID), 10)+": "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("--Updated Device--%d %s", device.ID, device.Name))

	// Update resource state with updated items and timestamp
	plan.ID = types.Int64Value(int64(device.ID))
	plan.Name = types.StringValue(device.Name)
	plan.Model = types.StringValue(device.Model)
	plan.Color = types.StringValue(device.Color)
	plan.Status = types.StringValue(device.Status)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *deviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, fmt.Sprintf("--Deleting Device--"))
	// Retrieve values from state
	var state deviceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("--Delete Device--%d %s", state.ID.ValueInt64(), state.Name))
	// Delete existing order
	err := r.client.DeleteDevice(state.ID.ValueInt64())
	tflog.Info(ctx, "--Deleted Device--")
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}
