// Code generated by protoc-gen-goten-resource
// Resource: DeviceGroup
// DO NOT EDIT!!!

package device_group

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
)

type DeviceGroupList []*DeviceGroup

func (l DeviceGroupList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*DeviceGroup))
}

func (l DeviceGroupList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(DeviceGroupList)...)
}

func (l DeviceGroupList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceGroupList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l DeviceGroupList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*DeviceGroup)
}

func (l DeviceGroupList) Length() int {
	return len(l)
}

type DeviceGroupChangeList []*DeviceGroupChange

func (l DeviceGroupChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*DeviceGroupChange))
}

func (l DeviceGroupChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(DeviceGroupChangeList)...)
}

func (l DeviceGroupChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceGroupChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l DeviceGroupChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*DeviceGroupChange)
}

func (l DeviceGroupChangeList) Length() int {
	return len(l)
}

type DeviceGroupNameList []*Name

func (l DeviceGroupNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l DeviceGroupNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(DeviceGroupNameList)...)
}

func (l DeviceGroupNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceGroupNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceGroupNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l DeviceGroupNameList) Length() int {
	return len(l)
}

type DeviceGroupReferenceList []*Reference

func (l DeviceGroupReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l DeviceGroupReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(DeviceGroupReferenceList)...)
}

func (l DeviceGroupReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceGroupReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceGroupReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l DeviceGroupReferenceList) Length() int {
	return len(l)
}

type DeviceGroupParentNameList []*ParentName

func (l DeviceGroupParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l DeviceGroupParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(DeviceGroupParentNameList)...)
}

func (l DeviceGroupParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceGroupParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceGroupParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l DeviceGroupParentNameList) Length() int {
	return len(l)
}

type DeviceGroupParentReferenceList []*ParentReference

func (l DeviceGroupParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l DeviceGroupParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(DeviceGroupParentReferenceList)...)
}

func (l DeviceGroupParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceGroupParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceGroupParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l DeviceGroupParentReferenceList) Length() int {
	return len(l)
}

type DeviceGroupMap map[Name]*DeviceGroup

func (m DeviceGroupMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m DeviceGroupMap) Set(res gotenresource.Resource) {
	tRes := res.(*DeviceGroup)
	m[*tRes.Name] = tRes
}

func (m DeviceGroupMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceGroupMap) Length() int {
	return len(m)
}

func (m DeviceGroupMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type DeviceGroupChangeMap map[Name]*DeviceGroupChange

func (m DeviceGroupChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m DeviceGroupChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*DeviceGroupChange)
	m[*tChange.GetDeviceGroupName()] = tChange
}

func (m DeviceGroupChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceGroupChangeMap) Length() int {
	return len(m)
}

func (m DeviceGroupChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
