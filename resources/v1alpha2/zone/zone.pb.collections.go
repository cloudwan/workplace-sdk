// Code generated by protoc-gen-goten-resource
// Resource: Zone
// DO NOT EDIT!!!

package zone

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
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
	_ = &area.Area{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
	_ = &floor.Floor{}
	_ = &site.Site{}
)

type ZoneList []*Zone

func (l ZoneList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Zone))
}

func (l ZoneList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(ZoneList)...)
}

func (l ZoneList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ZoneList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l ZoneList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Zone)
}

func (l ZoneList) Length() int {
	return len(l)
}

type ZoneChangeList []*ZoneChange

func (l ZoneChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*ZoneChange))
}

func (l ZoneChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(ZoneChangeList)...)
}

func (l ZoneChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ZoneChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l ZoneChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*ZoneChange)
}

func (l ZoneChangeList) Length() int {
	return len(l)
}

type ZoneNameList []*Name

func (l ZoneNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l ZoneNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(ZoneNameList)...)
}

func (l ZoneNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ZoneNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ZoneNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l ZoneNameList) Length() int {
	return len(l)
}

type ZoneReferenceList []*Reference

func (l ZoneReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l ZoneReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(ZoneReferenceList)...)
}

func (l ZoneReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ZoneReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ZoneReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l ZoneReferenceList) Length() int {
	return len(l)
}

type ZoneParentNameList []*ParentName

func (l ZoneParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l ZoneParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(ZoneParentNameList)...)
}

func (l ZoneParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ZoneParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ZoneParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l ZoneParentNameList) Length() int {
	return len(l)
}

type ZoneParentReferenceList []*ParentReference

func (l ZoneParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l ZoneParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(ZoneParentReferenceList)...)
}

func (l ZoneParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ZoneParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ZoneParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l ZoneParentReferenceList) Length() int {
	return len(l)
}

type ZoneMap map[Name]*Zone

func (m ZoneMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m ZoneMap) Set(res gotenresource.Resource) {
	tRes := res.(*Zone)
	m[*tRes.Name] = tRes
}

func (m ZoneMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ZoneMap) Length() int {
	return len(m)
}

func (m ZoneMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type ZoneChangeMap map[Name]*ZoneChange

func (m ZoneChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m ZoneChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*ZoneChange)
	m[*tChange.GetZoneName()] = tChange
}

func (m ZoneChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ZoneChangeMap) Length() int {
	return len(m)
}

func (m ZoneChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
