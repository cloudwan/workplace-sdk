// Code generated by protoc-gen-goten-client
// Service: Workplace
// DO NOT EDIT!!!

package workplace_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	agent_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/agent"
	area_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/area"
	building_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/building"
	device_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/device"
	device_group_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/device_group"
	floor_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/floor"
	property_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/property"
	site_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/site"
	vendor_connection_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/vendor_connection"
	zone_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/zone"
	agent "github.com/cloudwan/workplace-sdk/resources/v1alpha2/agent"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	device "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device"
	device_group "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device_group"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	property "github.com/cloudwan/workplace-sdk/resources/v1alpha2/property"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	vendor_connection "github.com/cloudwan/workplace-sdk/resources/v1alpha2/vendor_connection"
	zone "github.com/cloudwan/workplace-sdk/resources/v1alpha2/zone"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = gotenclient.MethodDescriptor(nil)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &agent.Agent{}
	_ = &agent_client.GetAgentRequest{}
	_ = &area.Area{}
	_ = &area_client.GetAreaRequest{}
	_ = &building.Building{}
	_ = &building_client.GetBuildingRequest{}
	_ = &device.Device{}
	_ = &device_group.DeviceGroup{}
	_ = &device_group_client.GetDeviceGroupRequest{}
	_ = &device_client.GetDeviceRequest{}
	_ = &floor.Floor{}
	_ = &floor_client.GetFloorRequest{}
	_ = &property.Property{}
	_ = &property_client.GetPropertyRequest{}
	_ = &site.Site{}
	_ = &site_client.GetSiteRequest{}
	_ = &vendor_connection.PointGrab{}
	_ = &vendor_connection_client.GetVendorConnectionRequest{}
	_ = &zone.Zone{}
	_ = &zone_client.GetZoneRequest{}
)

var (
	descriptorInitialized bool
	workplaceDescriptor   *WorkplaceDescriptor
)

type WorkplaceDescriptor struct{}

func (d *WorkplaceDescriptor) GetServiceName() string {
	return "workplace"
}

func (d *WorkplaceDescriptor) GetServiceDomain() string {
	return "workplace.edgelq.com"
}

func (d *WorkplaceDescriptor) GetVersion() string {
	return "v1alpha2"
}

func (d *WorkplaceDescriptor) GetNextVersion() string {

	return ""
}

func (d *WorkplaceDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		agent.GetDescriptor(),
		area.GetDescriptor(),
		building.GetDescriptor(),
		device.GetDescriptor(),
		device_group.GetDescriptor(),
		floor.GetDescriptor(),
		property.GetDescriptor(),
		site.GetDescriptor(),
		vendor_connection.GetDescriptor(),
		zone.GetDescriptor(),
	}
}

func (d *WorkplaceDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		agent_client.GetAgentServiceDescriptor(),
		area_client.GetAreaServiceDescriptor(),
		building_client.GetBuildingServiceDescriptor(),
		device_group_client.GetDeviceGroupServiceDescriptor(),
		device_client.GetDeviceServiceDescriptor(),
		floor_client.GetFloorServiceDescriptor(),
		property_client.GetPropertyServiceDescriptor(),
		site_client.GetSiteServiceDescriptor(),
		vendor_connection_client.GetVendorConnectionServiceDescriptor(),
		zone_client.GetZoneServiceDescriptor(),
	}
}

func GetWorkplaceDescriptor() *WorkplaceDescriptor {
	return workplaceDescriptor
}

func initDescriptor() {
	workplaceDescriptor = &WorkplaceDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(workplaceDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
