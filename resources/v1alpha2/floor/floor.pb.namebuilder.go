// Code generated by protoc-gen-goten-resource
// Resource: Floor
// DO NOT EDIT!!!

package floor

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
)

const (
	NamePattern_Project_Region_Site_Building = "projects/{project}/regions/{region}/sites/{site}/buildings/{building}/floors/{floor}"
)

type NamePattern struct {
	Pattern gotenresource.NamePattern `firestore:"pattern"`
}

type NameBuilder struct {
	nameObj Name
}

func NewNameBuilder() *NameBuilder {
	return &NameBuilder{
		nameObj: Name{
			FloorId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project_Region_Site_Building,
				},
			},
		},
	}
}

func (b *NameBuilder) Name() *Name {
	copied := b.nameObj
	return &copied
}

func (b *NameBuilder) Reference() *Reference {
	return b.nameObj.AsReference()
}

func (b *NameBuilder) Parent() *ParentName {
	copied := b.nameObj.ParentName
	return &copied
}

func (b *NameBuilder) ParentReference() *ParentReference {
	return b.nameObj.ParentName.AsReference()
}

func (b *NameBuilder) SetId(id string) *NameBuilder {
	b.nameObj.FloorId = id
	return b
}

func (b *NameBuilder) SetBuilding(parent *building.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case building.NamePattern_Project_Region_Site:
		parentName.Pattern = NamePattern_Project_Region_Site_Building
	}
	parentName.ProjectId = parent.ProjectId
	parentName.RegionId = parent.RegionId
	parentName.SiteId = parent.SiteId
	parentName.BuildingId = parent.BuildingId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.SiteId != "" && parentName.BuildingId != "" {
		parentName.Pattern = NamePattern_Project_Region_Site_Building
	}
	return b
}

func (b *NameBuilder) SetRegionId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.RegionId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.SiteId != "" && parentName.BuildingId != "" {
		parentName.Pattern = NamePattern_Project_Region_Site_Building
	}
	return b
}

func (b *NameBuilder) SetSiteId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.SiteId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.SiteId != "" && parentName.BuildingId != "" {
		parentName.Pattern = NamePattern_Project_Region_Site_Building
	}
	return b
}

func (b *NameBuilder) SetBuildingId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.BuildingId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.SiteId != "" && parentName.BuildingId != "" {
		parentName.Pattern = NamePattern_Project_Region_Site_Building
	}
	return b
}
