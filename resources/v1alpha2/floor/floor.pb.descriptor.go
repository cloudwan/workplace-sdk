// Code generated by protoc-gen-goten-resource
// Resource: Floor
// DO NOT EDIT!!!

package floor

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
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

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Floor", "Floors", "workplace.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorName},
			"pattern", "floorId",
			[]string{"projectId", "regionId", "siteId", "buildingId"},
			[]gotenresource.NamePattern{NamePattern_Project_Region_Site_Building}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewFloor() *Floor {
	return &Floor{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewFloor()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewFloorName() *Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewGetQuery() gotenresource.GetQuery {
	return &GetQuery{}
}

func (d *Descriptor) NewListQuery() gotenresource.ListQuery {
	return &ListQuery{}
}

func (d *Descriptor) NewSearchQuery() gotenresource.SearchQuery {
	return nil
}

func (d *Descriptor) NewWatchQuery() gotenresource.WatchQuery {
	return &WatchQuery{}
}
func (d *Descriptor) NewFloorCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewFloorCursor()
}
func (d *Descriptor) NewFloorChange() *FloorChange {
	return &FloorChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewFloorChange()
}

func (d *Descriptor) NewFloorQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewFloorQueryResultSnapshot()
}
func (d *Descriptor) NewFloorQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewFloorQueryResultChange()
}

func (d *Descriptor) NewFloorList(size, reserved int) FloorList {
	return make(FloorList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(FloorList, size, reserved)
}
func (d *Descriptor) NewFloorChangeList(size, reserved int) FloorChangeList {
	return make(FloorChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(FloorChangeList, size, reserved)
}

func (d *Descriptor) NewFloorNameList(size, reserved int) FloorNameList {
	return make(FloorNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(FloorNameList, size, reserved)
}

func (d *Descriptor) NewFloorReferenceList(size, reserved int) FloorReferenceList {
	return make(FloorReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(FloorReferenceList, size, reserved)
}
func (d *Descriptor) NewFloorParentNameList(size, reserved int) FloorParentNameList {
	return make(FloorParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(FloorParentNameList, size, reserved)
}
func (d *Descriptor) NewFloorParentReferenceList(size, reserved int) FloorParentReferenceList {
	return make(FloorParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(FloorParentReferenceList, size, reserved)
}

func (d *Descriptor) NewFloorMap(reserved int) FloorMap {
	return make(FloorMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(FloorMap, reserved)
}
func (d *Descriptor) NewFloorChangeMap(reserved int) FloorChangeMap {
	return make(FloorChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(FloorChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseFloor_FieldPath(raw)
}

func (d *Descriptor) ParseFloorName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}