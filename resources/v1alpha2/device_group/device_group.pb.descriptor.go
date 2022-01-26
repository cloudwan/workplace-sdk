// Code generated by protoc-gen-goten-resource
// Resource: DeviceGroup
// DO NOT EDIT!!!

package device_group

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"DeviceGroup", "DeviceGroups", "workplace.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorName},
			"pattern", "deviceGroupId",
			[]string{"projectId"},
			[]gotenresource.NamePattern{NamePattern_Project}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewDeviceGroup() *DeviceGroup {
	return &DeviceGroup{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewDeviceGroup()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewDeviceGroupName() *Name {
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
func (d *Descriptor) NewDeviceGroupCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewDeviceGroupCursor()
}
func (d *Descriptor) NewDeviceGroupChange() *DeviceGroupChange {
	return &DeviceGroupChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewDeviceGroupChange()
}

func (d *Descriptor) NewDeviceGroupQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewDeviceGroupQueryResultSnapshot()
}
func (d *Descriptor) NewDeviceGroupQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewDeviceGroupQueryResultChange()
}

func (d *Descriptor) NewDeviceGroupList(size, reserved int) DeviceGroupList {
	return make(DeviceGroupList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(DeviceGroupList, size, reserved)
}
func (d *Descriptor) NewDeviceGroupChangeList(size, reserved int) DeviceGroupChangeList {
	return make(DeviceGroupChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(DeviceGroupChangeList, size, reserved)
}

func (d *Descriptor) NewDeviceGroupNameList(size, reserved int) DeviceGroupNameList {
	return make(DeviceGroupNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(DeviceGroupNameList, size, reserved)
}

func (d *Descriptor) NewDeviceGroupReferenceList(size, reserved int) DeviceGroupReferenceList {
	return make(DeviceGroupReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(DeviceGroupReferenceList, size, reserved)
}
func (d *Descriptor) NewDeviceGroupParentNameList(size, reserved int) DeviceGroupParentNameList {
	return make(DeviceGroupParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(DeviceGroupParentNameList, size, reserved)
}
func (d *Descriptor) NewDeviceGroupParentReferenceList(size, reserved int) DeviceGroupParentReferenceList {
	return make(DeviceGroupParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(DeviceGroupParentReferenceList, size, reserved)
}

func (d *Descriptor) NewDeviceGroupMap(reserved int) DeviceGroupMap {
	return make(DeviceGroupMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(DeviceGroupMap, reserved)
}
func (d *Descriptor) NewDeviceGroupChangeMap(reserved int) DeviceGroupChangeMap {
	return make(DeviceGroupChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(DeviceGroupChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseDeviceGroup_FieldPath(raw)
}

func (d *Descriptor) ParseDeviceGroupName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}