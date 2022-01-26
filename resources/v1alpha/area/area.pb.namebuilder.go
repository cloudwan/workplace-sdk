// Code generated by protoc-gen-goten-resource
// Resource: Area
// DO NOT EDIT!!!

package area

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha/floor"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &workplace_common.BBox{}
	_ = &floor.Floor{}
)

const (
	NamePattern_Project_Site_Building_Floor = "projects/{project}/sites/{site}/buildings/{building}/floors/{floor}/areas/{area}"
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
			AreaId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project_Site_Building_Floor,
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
	b.nameObj.AreaId = id
	return b
}

func (b *NameBuilder) SetFloor(parent *floor.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case floor.NamePattern_Project_Site_Building:
		parentName.Pattern = NamePattern_Project_Site_Building_Floor
	}
	parentName.ProjectId = parent.ProjectId
	parentName.SiteId = parent.SiteId
	parentName.BuildingId = parent.BuildingId
	parentName.FloorId = parent.FloorId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.SiteId != "" && parentName.BuildingId != "" && parentName.FloorId != "" {
		parentName.Pattern = NamePattern_Project_Site_Building_Floor
	}
	return b
}

func (b *NameBuilder) SetSiteId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.SiteId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.SiteId != "" && parentName.BuildingId != "" && parentName.FloorId != "" {
		parentName.Pattern = NamePattern_Project_Site_Building_Floor
	}
	return b
}

func (b *NameBuilder) SetBuildingId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.BuildingId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.SiteId != "" && parentName.BuildingId != "" && parentName.FloorId != "" {
		parentName.Pattern = NamePattern_Project_Site_Building_Floor
	}
	return b
}

func (b *NameBuilder) SetFloorId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.FloorId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.SiteId != "" && parentName.BuildingId != "" && parentName.FloorId != "" {
		parentName.Pattern = NamePattern_Project_Site_Building_Floor
	}
	return b
}