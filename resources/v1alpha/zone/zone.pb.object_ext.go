// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha/zone.proto
// DO NOT EDIT!!!

package zone

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha/site"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = fmt.Stringer(nil)
	_ = sort.Interface(nil)

	_ = proto.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
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

func (o *Zone) GotenObjectExt() {}

func (o *Zone) MakeFullFieldMask() *Zone_FieldMask {
	return FullZone_FieldMask()
}

func (o *Zone) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullZone_FieldMask()
}

func (o *Zone) MakeDiffFieldMask(other *Zone) *Zone_FieldMask {
	if o == nil && other == nil {
		return &Zone_FieldMask{}
	}
	if o == nil || other == nil {
		return FullZone_FieldMask()
	}

	res := &Zone_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorDisplayName})
	}
	if o.GetType() != other.GetType() {
		res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorType})
	}
	{
		subMask := o.GetGeometry().MakeDiffFieldMask(other.GetGeometry())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorGeometry})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Zone_FieldSubPath{selector: Zone_FieldPathSelectorGeometry, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetVendorSpec().MakeDiffFieldMask(other.GetVendorSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorVendorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Zone_FieldSubPath{selector: Zone_FieldPathSelectorVendorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetState().MakeDiffFieldMask(other.GetState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Zone_FieldSubPath{selector: Zone_FieldPathSelectorState, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Zone_FieldTerminalPath{selector: Zone_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Zone_FieldSubPath{selector: Zone_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Zone) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Zone))
}

func (o *Zone) Clone() *Zone {
	if o == nil {
		return nil
	}
	result := &Zone{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.DisplayName = o.DisplayName
	result.Type = o.Type
	result.Geometry = o.Geometry.Clone()
	result.VendorSpec = o.VendorSpec.Clone()
	result.State = o.State.Clone()
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *Zone) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Zone) Merge(source *Zone) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	o.DisplayName = source.GetDisplayName()
	o.Type = source.GetType()
	if source.GetGeometry() != nil {
		if o.Geometry == nil {
			o.Geometry = new(workplace_common.Geometry)
		}
		o.Geometry.Merge(source.GetGeometry())
	}
	if source.GetVendorSpec() != nil {
		if o.VendorSpec == nil {
			o.VendorSpec = new(Zone_VendorSpec)
		}
		o.VendorSpec.Merge(source.GetVendorSpec())
	}
	if source.GetState() != nil {
		if o.State == nil {
			o.State = new(Zone_State)
		}
		o.State.Merge(source.GetState())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *Zone) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Zone))
}

func (o *Zone_VendorSpec) GotenObjectExt() {}

func (o *Zone_VendorSpec) MakeFullFieldMask() *Zone_VendorSpec_FieldMask {
	return FullZone_VendorSpec_FieldMask()
}

func (o *Zone_VendorSpec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullZone_VendorSpec_FieldMask()
}

func (o *Zone_VendorSpec) MakeDiffFieldMask(other *Zone_VendorSpec) *Zone_VendorSpec_FieldMask {
	if o == nil && other == nil {
		return &Zone_VendorSpec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullZone_VendorSpec_FieldMask()
	}

	res := &Zone_VendorSpec_FieldMask{}
	{
		subMask := o.GetPointGrab().MakeDiffFieldMask(other.GetPointGrab())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ZoneVendorSpec_FieldTerminalPath{selector: ZoneVendorSpec_FieldPathSelectorPointGrab})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ZoneVendorSpec_FieldSubPath{selector: ZoneVendorSpec_FieldPathSelectorPointGrab, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Zone_VendorSpec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Zone_VendorSpec))
}

func (o *Zone_VendorSpec) Clone() *Zone_VendorSpec {
	if o == nil {
		return nil
	}
	result := &Zone_VendorSpec{}
	result.PointGrab = o.PointGrab.Clone()
	return result
}

func (o *Zone_VendorSpec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Zone_VendorSpec) Merge(source *Zone_VendorSpec) {
	if source.GetPointGrab() != nil {
		if o.PointGrab == nil {
			o.PointGrab = new(Zone_VendorSpec_PointGrab)
		}
		o.PointGrab.Merge(source.GetPointGrab())
	}
}

func (o *Zone_VendorSpec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Zone_VendorSpec))
}

func (o *Zone_State) GotenObjectExt() {}

func (o *Zone_State) MakeFullFieldMask() *Zone_State_FieldMask {
	return FullZone_State_FieldMask()
}

func (o *Zone_State) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullZone_State_FieldMask()
}

func (o *Zone_State) MakeDiffFieldMask(other *Zone_State) *Zone_State_FieldMask {
	if o == nil && other == nil {
		return &Zone_State_FieldMask{}
	}
	if o == nil || other == nil {
		return FullZone_State_FieldMask()
	}

	res := &Zone_State_FieldMask{}
	{
		subMask := o.GetOccupancy().MakeDiffFieldMask(other.GetOccupancy())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ZoneState_FieldTerminalPath{selector: ZoneState_FieldPathSelectorOccupancy})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ZoneState_FieldSubPath{selector: ZoneState_FieldPathSelectorOccupancy, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Zone_State) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Zone_State))
}

func (o *Zone_State) Clone() *Zone_State {
	if o == nil {
		return nil
	}
	result := &Zone_State{}
	result.Occupancy = o.Occupancy.Clone()
	return result
}

func (o *Zone_State) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Zone_State) Merge(source *Zone_State) {
	if source.GetOccupancy() != nil {
		if o.Occupancy == nil {
			o.Occupancy = new(Zone_State_Occupancy)
		}
		o.Occupancy.Merge(source.GetOccupancy())
	}
}

func (o *Zone_State) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Zone_State))
}

func (o *Zone_VendorSpec_PointGrab) GotenObjectExt() {}

func (o *Zone_VendorSpec_PointGrab) MakeFullFieldMask() *Zone_VendorSpec_PointGrab_FieldMask {
	return FullZone_VendorSpec_PointGrab_FieldMask()
}

func (o *Zone_VendorSpec_PointGrab) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullZone_VendorSpec_PointGrab_FieldMask()
}

func (o *Zone_VendorSpec_PointGrab) MakeDiffFieldMask(other *Zone_VendorSpec_PointGrab) *Zone_VendorSpec_PointGrab_FieldMask {
	if o == nil && other == nil {
		return &Zone_VendorSpec_PointGrab_FieldMask{}
	}
	if o == nil || other == nil {
		return FullZone_VendorSpec_PointGrab_FieldMask()
	}

	res := &Zone_VendorSpec_PointGrab_FieldMask{}
	if o.GetAreaId() != other.GetAreaId() {
		res.Paths = append(res.Paths, &ZoneVendorSpecPointGrab_FieldTerminalPath{selector: ZoneVendorSpecPointGrab_FieldPathSelectorAreaId})
	}
	{
		subMask := o.GetPolygon().MakeDiffFieldMask(other.GetPolygon())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ZoneVendorSpecPointGrab_FieldTerminalPath{selector: ZoneVendorSpecPointGrab_FieldPathSelectorPolygon})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ZoneVendorSpecPointGrab_FieldSubPath{selector: ZoneVendorSpecPointGrab_FieldPathSelectorPolygon, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Zone_VendorSpec_PointGrab) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Zone_VendorSpec_PointGrab))
}

func (o *Zone_VendorSpec_PointGrab) Clone() *Zone_VendorSpec_PointGrab {
	if o == nil {
		return nil
	}
	result := &Zone_VendorSpec_PointGrab{}
	result.AreaId = o.AreaId
	result.Polygon = o.Polygon.Clone()
	return result
}

func (o *Zone_VendorSpec_PointGrab) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Zone_VendorSpec_PointGrab) Merge(source *Zone_VendorSpec_PointGrab) {
	o.AreaId = source.GetAreaId()
	if source.GetPolygon() != nil {
		if o.Polygon == nil {
			o.Polygon = new(workplace_common.Polygon)
		}
		o.Polygon.Merge(source.GetPolygon())
	}
}

func (o *Zone_VendorSpec_PointGrab) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Zone_VendorSpec_PointGrab))
}

func (o *Zone_State_Occupancy) GotenObjectExt() {}

func (o *Zone_State_Occupancy) MakeFullFieldMask() *Zone_State_Occupancy_FieldMask {
	return FullZone_State_Occupancy_FieldMask()
}

func (o *Zone_State_Occupancy) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullZone_State_Occupancy_FieldMask()
}

func (o *Zone_State_Occupancy) MakeDiffFieldMask(other *Zone_State_Occupancy) *Zone_State_Occupancy_FieldMask {
	if o == nil && other == nil {
		return &Zone_State_Occupancy_FieldMask{}
	}
	if o == nil || other == nil {
		return FullZone_State_Occupancy_FieldMask()
	}

	res := &Zone_State_Occupancy_FieldMask{}
	if o.GetIsOccupied() != other.GetIsOccupied() {
		res.Paths = append(res.Paths, &ZoneStateOccupancy_FieldTerminalPath{selector: ZoneStateOccupancy_FieldPathSelectorIsOccupied})
	}
	if !proto.Equal(o.GetLastOccupiedTime(), other.GetLastOccupiedTime()) {
		res.Paths = append(res.Paths, &ZoneStateOccupancy_FieldTerminalPath{selector: ZoneStateOccupancy_FieldPathSelectorLastOccupiedTime})
	}
	return res
}

func (o *Zone_State_Occupancy) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Zone_State_Occupancy))
}

func (o *Zone_State_Occupancy) Clone() *Zone_State_Occupancy {
	if o == nil {
		return nil
	}
	result := &Zone_State_Occupancy{}
	result.IsOccupied = o.IsOccupied
	result.LastOccupiedTime = proto.Clone(o.LastOccupiedTime).(*timestamp.Timestamp)
	return result
}

func (o *Zone_State_Occupancy) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Zone_State_Occupancy) Merge(source *Zone_State_Occupancy) {
	o.IsOccupied = source.GetIsOccupied()
	if source.GetLastOccupiedTime() != nil {
		if o.LastOccupiedTime == nil {
			o.LastOccupiedTime = new(timestamp.Timestamp)
		}
		proto.Merge(o.LastOccupiedTime, source.GetLastOccupiedTime())
	}
}

func (o *Zone_State_Occupancy) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Zone_State_Occupancy))
}
