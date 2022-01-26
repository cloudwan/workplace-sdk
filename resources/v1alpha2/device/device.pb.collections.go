// Code generated by protoc-gen-goten-resource
// Resource: Device
// DO NOT EDIT!!!

package device

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	agent "github.com/cloudwan/workplace-sdk/resources/v1alpha2/agent"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	zone "github.com/cloudwan/workplace-sdk/resources/v1alpha2/zone"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &agent.Agent{}
	_ = &area.Area{}
	_ = &workplace_common.BACNetEntity{}
	_ = &building.Building{}
	_ = &floor.Floor{}
	_ = &site.Site{}
	_ = &zone.Zone{}
)

type DeviceList []*Device

func (l DeviceList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Device))
}

func (l DeviceList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(DeviceList)...)
}

func (l DeviceList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l DeviceList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Device)
}

func (l DeviceList) Length() int {
	return len(l)
}

type DeviceChangeList []*DeviceChange

func (l DeviceChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*DeviceChange))
}

func (l DeviceChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(DeviceChangeList)...)
}

func (l DeviceChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l DeviceChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*DeviceChange)
}

func (l DeviceChangeList) Length() int {
	return len(l)
}

type DeviceNameList []*Name

func (l DeviceNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l DeviceNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(DeviceNameList)...)
}

func (l DeviceNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l DeviceNameList) Length() int {
	return len(l)
}

type DeviceReferenceList []*Reference

func (l DeviceReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l DeviceReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(DeviceReferenceList)...)
}

func (l DeviceReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l DeviceReferenceList) Length() int {
	return len(l)
}

type DeviceParentNameList []*ParentName

func (l DeviceParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l DeviceParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(DeviceParentNameList)...)
}

func (l DeviceParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l DeviceParentNameList) Length() int {
	return len(l)
}

type DeviceParentReferenceList []*ParentReference

func (l DeviceParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l DeviceParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(DeviceParentReferenceList)...)
}

func (l DeviceParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l DeviceParentReferenceList) Length() int {
	return len(l)
}

type DeviceMap map[Name]*Device

func (m DeviceMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m DeviceMap) Set(res gotenresource.Resource) {
	tRes := res.(*Device)
	m[*tRes.Name] = tRes
}

func (m DeviceMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceMap) Length() int {
	return len(m)
}

func (m DeviceMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type DeviceChangeMap map[Name]*DeviceChange

func (m DeviceChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m DeviceChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*DeviceChange)
	m[*tChange.GetDeviceName()] = tChange
}

func (m DeviceChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceChangeMap) Length() int {
	return len(m)
}

func (m DeviceChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
