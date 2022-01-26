// Code generated by protoc-gen-goten-access
// Service: Workplace
// DO NOT EDIT!!!

package workplace_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	agent_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/agent"
	area_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/area"
	building_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/building"
	device_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/device"
	device_group_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/device_group"
	floor_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/floor"
	property_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/property"
	site_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/site"
	vendor_connection_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/vendor_connection"
	zone_access "github.com/cloudwan/workplace-sdk/access/v1alpha2/zone"
	workplace_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/workplace"
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

type WorkplaceApiAccess interface {
	gotenresource.Access

	agent.AgentAccess
	area.AreaAccess
	building.BuildingAccess
	device.DeviceAccess
	device_group.DeviceGroupAccess
	floor.FloorAccess
	property.PropertyAccess
	site.SiteAccess
	vendor_connection.VendorConnectionAccess
	zone.ZoneAccess
}

type apiWorkplaceAccess struct {
	gotenresource.Access

	agent.AgentAccess
	area.AreaAccess
	building.BuildingAccess
	device.DeviceAccess
	device_group.DeviceGroupAccess
	floor.FloorAccess
	property.PropertyAccess
	site.SiteAccess
	vendor_connection.VendorConnectionAccess
	zone.ZoneAccess
}

func NewApiAccess(client workplace_client.WorkplaceClient) WorkplaceApiAccess {

	agentAccess := agent_access.NewApiAgentAccess(client)
	areaAccess := area_access.NewApiAreaAccess(client)
	buildingAccess := building_access.NewApiBuildingAccess(client)
	deviceAccess := device_access.NewApiDeviceAccess(client)
	deviceGroupAccess := device_group_access.NewApiDeviceGroupAccess(client)
	floorAccess := floor_access.NewApiFloorAccess(client)
	propertyAccess := property_access.NewApiPropertyAccess(client)
	siteAccess := site_access.NewApiSiteAccess(client)
	vendorConnectionAccess := vendor_connection_access.NewApiVendorConnectionAccess(client)
	zoneAccess := zone_access.NewApiZoneAccess(client)

	return &apiWorkplaceAccess{
		Access: gotenresource.NewCompositeAccess(

			agent.AsAnyCastAccess(agentAccess),
			area.AsAnyCastAccess(areaAccess),
			building.AsAnyCastAccess(buildingAccess),
			device.AsAnyCastAccess(deviceAccess),
			device_group.AsAnyCastAccess(deviceGroupAccess),
			floor.AsAnyCastAccess(floorAccess),
			property.AsAnyCastAccess(propertyAccess),
			site.AsAnyCastAccess(siteAccess),
			vendor_connection.AsAnyCastAccess(vendorConnectionAccess),
			zone.AsAnyCastAccess(zoneAccess),
		),

		AgentAccess:            agentAccess,
		AreaAccess:             areaAccess,
		BuildingAccess:         buildingAccess,
		DeviceAccess:           deviceAccess,
		DeviceGroupAccess:      deviceGroupAccess,
		FloorAccess:            floorAccess,
		PropertyAccess:         propertyAccess,
		SiteAccess:             siteAccess,
		VendorConnectionAccess: vendorConnectionAccess,
		ZoneAccess:             zoneAccess,
	}
}
