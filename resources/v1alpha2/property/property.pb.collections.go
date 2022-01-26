// Code generated by protoc-gen-goten-resource
// Resource: Property
// DO NOT EDIT!!!

package property

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	device "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	zone "github.com/cloudwan/workplace-sdk/resources/v1alpha2/zone"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &area.Area{}
	_ = &workplace_common.BACNetEntity{}
	_ = &building.Building{}
	_ = &device.Device{}
	_ = &floor.Floor{}
	_ = &site.Site{}
	_ = &zone.Zone{}
)

type PropertyList []*Property

func (l PropertyList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Property))
}

func (l PropertyList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(PropertyList)...)
}

func (l PropertyList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PropertyList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l PropertyList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Property)
}

func (l PropertyList) Length() int {
	return len(l)
}

type PropertyChangeList []*PropertyChange

func (l PropertyChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*PropertyChange))
}

func (l PropertyChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(PropertyChangeList)...)
}

func (l PropertyChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PropertyChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l PropertyChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*PropertyChange)
}

func (l PropertyChangeList) Length() int {
	return len(l)
}

type PropertyNameList []*Name

func (l PropertyNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l PropertyNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(PropertyNameList)...)
}

func (l PropertyNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PropertyNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l PropertyNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l PropertyNameList) Length() int {
	return len(l)
}

type PropertyReferenceList []*Reference

func (l PropertyReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l PropertyReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(PropertyReferenceList)...)
}

func (l PropertyReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PropertyReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l PropertyReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l PropertyReferenceList) Length() int {
	return len(l)
}

type PropertyParentNameList []*ParentName

func (l PropertyParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l PropertyParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(PropertyParentNameList)...)
}

func (l PropertyParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PropertyParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l PropertyParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l PropertyParentNameList) Length() int {
	return len(l)
}

type PropertyParentReferenceList []*ParentReference

func (l PropertyParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l PropertyParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(PropertyParentReferenceList)...)
}

func (l PropertyParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PropertyParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l PropertyParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l PropertyParentReferenceList) Length() int {
	return len(l)
}

type PropertyMap map[Name]*Property

func (m PropertyMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m PropertyMap) Set(res gotenresource.Resource) {
	tRes := res.(*Property)
	m[*tRes.Name] = tRes
}

func (m PropertyMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m PropertyMap) Length() int {
	return len(m)
}

func (m PropertyMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type PropertyChangeMap map[Name]*PropertyChange

func (m PropertyChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m PropertyChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*PropertyChange)
	m[*tChange.GetPropertyName()] = tChange
}

func (m PropertyChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m PropertyChangeMap) Length() int {
	return len(m)
}

func (m PropertyChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}