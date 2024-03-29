// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/floor.proto
// DO NOT EDIT!!!

package floor

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
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
)

func (o *Floor) GotenObjectExt() {}

func (o *Floor) MakeFullFieldMask() *Floor_FieldMask {
	return FullFloor_FieldMask()
}

func (o *Floor) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_FieldMask()
}

func (o *Floor) MakeDiffFieldMask(other *Floor) *Floor_FieldMask {
	if o == nil && other == nil {
		return &Floor_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_FieldMask()
	}

	res := &Floor_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorDisplayName})
	}
	if o.GetOrderingNumber() != other.GetOrderingNumber() {
		res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorOrderingNumber})
	}
	{
		subMask := o.GetGeometry().MakeDiffFieldMask(other.GetGeometry())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorGeometry})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Floor_FieldSubPath{selector: Floor_FieldPathSelectorGeometry, subPath: subpath})
			}
		}
	}

	if len(o.GetVendorMappings()) == len(other.GetVendorMappings()) {
		for i, lValue := range o.GetVendorMappings() {
			rValue := other.GetVendorMappings()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorVendorMappings})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorVendorMappings})
	}
	{
		subMask := o.GetVendorSpec().MakeDiffFieldMask(other.GetVendorSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorVendorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Floor_FieldSubPath{selector: Floor_FieldPathSelectorVendorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetVendorInfo().MakeDiffFieldMask(other.GetVendorInfo())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorVendorInfo})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Floor_FieldSubPath{selector: Floor_FieldPathSelectorVendorInfo, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetVendorState().MakeDiffFieldMask(other.GetVendorState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorVendorState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Floor_FieldSubPath{selector: Floor_FieldPathSelectorVendorState, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Floor_FieldTerminalPath{selector: Floor_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Floor_FieldSubPath{selector: Floor_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Floor) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor))
}

func (o *Floor) Clone() *Floor {
	if o == nil {
		return nil
	}
	result := &Floor{}
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
	result.OrderingNumber = o.OrderingNumber
	result.Geometry = o.Geometry.Clone()
	result.VendorMappings = make([]*workplace_common.VendorMapping, len(o.VendorMappings))
	for i, sourceValue := range o.VendorMappings {
		result.VendorMappings[i] = sourceValue.Clone()
	}
	result.VendorSpec = o.VendorSpec.Clone()
	result.VendorInfo = o.VendorInfo.Clone()
	result.VendorState = o.VendorState.Clone()
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *Floor) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor) Merge(source *Floor) {
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
	o.OrderingNumber = source.GetOrderingNumber()
	if source.GetGeometry() != nil {
		if o.Geometry == nil {
			o.Geometry = new(workplace_common.Geometry)
		}
		o.Geometry.Merge(source.GetGeometry())
	}
	for _, sourceValue := range source.GetVendorMappings() {
		exists := false
		for _, currentValue := range o.VendorMappings {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *workplace_common.VendorMapping
			if sourceValue != nil {
				newDstElement = new(workplace_common.VendorMapping)
				newDstElement.Merge(sourceValue)
			}
			o.VendorMappings = append(o.VendorMappings, newDstElement)
		}
	}

	if source.GetVendorSpec() != nil {
		if o.VendorSpec == nil {
			o.VendorSpec = new(Floor_VendorSpec)
		}
		o.VendorSpec.Merge(source.GetVendorSpec())
	}
	if source.GetVendorInfo() != nil {
		if o.VendorInfo == nil {
			o.VendorInfo = new(Floor_VendorInfo)
		}
		o.VendorInfo.Merge(source.GetVendorInfo())
	}
	if source.GetVendorState() != nil {
		if o.VendorState == nil {
			o.VendorState = new(Floor_VendorState)
		}
		o.VendorState.Merge(source.GetVendorState())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *Floor) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor))
}

func (o *Floor_VendorSpec) GotenObjectExt() {}

func (o *Floor_VendorSpec) MakeFullFieldMask() *Floor_VendorSpec_FieldMask {
	return FullFloor_VendorSpec_FieldMask()
}

func (o *Floor_VendorSpec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorSpec_FieldMask()
}

func (o *Floor_VendorSpec) MakeDiffFieldMask(other *Floor_VendorSpec) *Floor_VendorSpec_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorSpec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorSpec_FieldMask()
	}

	res := &Floor_VendorSpec_FieldMask{}
	{
		subMask := o.GetPointGrab().MakeDiffFieldMask(other.GetPointGrab())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &FloorVendorSpec_FieldTerminalPath{selector: FloorVendorSpec_FieldPathSelectorPointGrab})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &FloorVendorSpec_FieldSubPath{selector: FloorVendorSpec_FieldPathSelectorPointGrab, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMapbox().MakeDiffFieldMask(other.GetMapbox())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &FloorVendorSpec_FieldTerminalPath{selector: FloorVendorSpec_FieldPathSelectorMapbox})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &FloorVendorSpec_FieldSubPath{selector: FloorVendorSpec_FieldPathSelectorMapbox, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Floor_VendorSpec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorSpec))
}

func (o *Floor_VendorSpec) Clone() *Floor_VendorSpec {
	if o == nil {
		return nil
	}
	result := &Floor_VendorSpec{}
	result.PointGrab = o.PointGrab.Clone()
	result.Mapbox = o.Mapbox.Clone()
	return result
}

func (o *Floor_VendorSpec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorSpec) Merge(source *Floor_VendorSpec) {
	if source.GetPointGrab() != nil {
		if o.PointGrab == nil {
			o.PointGrab = new(Floor_VendorSpec_PointGrab)
		}
		o.PointGrab.Merge(source.GetPointGrab())
	}
	if source.GetMapbox() != nil {
		if o.Mapbox == nil {
			o.Mapbox = new(Floor_VendorSpec_Mapbox)
		}
		o.Mapbox.Merge(source.GetMapbox())
	}
}

func (o *Floor_VendorSpec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorSpec))
}

func (o *Floor_VendorInfo) GotenObjectExt() {}

func (o *Floor_VendorInfo) MakeFullFieldMask() *Floor_VendorInfo_FieldMask {
	return FullFloor_VendorInfo_FieldMask()
}

func (o *Floor_VendorInfo) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorInfo_FieldMask()
}

func (o *Floor_VendorInfo) MakeDiffFieldMask(other *Floor_VendorInfo) *Floor_VendorInfo_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorInfo_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorInfo_FieldMask()
	}

	res := &Floor_VendorInfo_FieldMask{}
	{
		subMask := o.GetPointGrab().MakeDiffFieldMask(other.GetPointGrab())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &FloorVendorInfo_FieldTerminalPath{selector: FloorVendorInfo_FieldPathSelectorPointGrab})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &FloorVendorInfo_FieldSubPath{selector: FloorVendorInfo_FieldPathSelectorPointGrab, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Floor_VendorInfo) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorInfo))
}

func (o *Floor_VendorInfo) Clone() *Floor_VendorInfo {
	if o == nil {
		return nil
	}
	result := &Floor_VendorInfo{}
	result.PointGrab = o.PointGrab.Clone()
	return result
}

func (o *Floor_VendorInfo) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorInfo) Merge(source *Floor_VendorInfo) {
	if source.GetPointGrab() != nil {
		if o.PointGrab == nil {
			o.PointGrab = new(Floor_VendorInfo_PointGrab)
		}
		o.PointGrab.Merge(source.GetPointGrab())
	}
}

func (o *Floor_VendorInfo) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorInfo))
}

func (o *Floor_VendorState) GotenObjectExt() {}

func (o *Floor_VendorState) MakeFullFieldMask() *Floor_VendorState_FieldMask {
	return FullFloor_VendorState_FieldMask()
}

func (o *Floor_VendorState) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorState_FieldMask()
}

func (o *Floor_VendorState) MakeDiffFieldMask(other *Floor_VendorState) *Floor_VendorState_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorState_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorState_FieldMask()
	}

	res := &Floor_VendorState_FieldMask{}
	{
		subMask := o.GetPointGrab().MakeDiffFieldMask(other.GetPointGrab())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &FloorVendorState_FieldTerminalPath{selector: FloorVendorState_FieldPathSelectorPointGrab})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &FloorVendorState_FieldSubPath{selector: FloorVendorState_FieldPathSelectorPointGrab, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Floor_VendorState) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorState))
}

func (o *Floor_VendorState) Clone() *Floor_VendorState {
	if o == nil {
		return nil
	}
	result := &Floor_VendorState{}
	result.PointGrab = o.PointGrab.Clone()
	return result
}

func (o *Floor_VendorState) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorState) Merge(source *Floor_VendorState) {
	if source.GetPointGrab() != nil {
		if o.PointGrab == nil {
			o.PointGrab = new(Floor_VendorState_PointGrab)
		}
		o.PointGrab.Merge(source.GetPointGrab())
	}
}

func (o *Floor_VendorState) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorState))
}

func (o *Floor_VendorSpec_PointGrab) GotenObjectExt() {}

func (o *Floor_VendorSpec_PointGrab) MakeFullFieldMask() *Floor_VendorSpec_PointGrab_FieldMask {
	return FullFloor_VendorSpec_PointGrab_FieldMask()
}

func (o *Floor_VendorSpec_PointGrab) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorSpec_PointGrab_FieldMask()
}

func (o *Floor_VendorSpec_PointGrab) MakeDiffFieldMask(other *Floor_VendorSpec_PointGrab) *Floor_VendorSpec_PointGrab_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorSpec_PointGrab_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorSpec_PointGrab_FieldMask()
	}

	res := &Floor_VendorSpec_PointGrab_FieldMask{}
	if o.GetFloorId() != other.GetFloorId() {
		res.Paths = append(res.Paths, &FloorVendorSpecPointGrab_FieldTerminalPath{selector: FloorVendorSpecPointGrab_FieldPathSelectorFloorId})
	}

	if len(o.GetReferencePoints()) == len(other.GetReferencePoints()) {
		for i, lValue := range o.GetReferencePoints() {
			rValue := other.GetReferencePoints()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &FloorVendorSpecPointGrab_FieldTerminalPath{selector: FloorVendorSpecPointGrab_FieldPathSelectorReferencePoints})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &FloorVendorSpecPointGrab_FieldTerminalPath{selector: FloorVendorSpecPointGrab_FieldPathSelectorReferencePoints})
	}
	return res
}

func (o *Floor_VendorSpec_PointGrab) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorSpec_PointGrab))
}

func (o *Floor_VendorSpec_PointGrab) Clone() *Floor_VendorSpec_PointGrab {
	if o == nil {
		return nil
	}
	result := &Floor_VendorSpec_PointGrab{}
	result.FloorId = o.FloorId
	result.ReferencePoints = make([]*Floor_VendorSpec_PointGrab_ReferencePoint, len(o.ReferencePoints))
	for i, sourceValue := range o.ReferencePoints {
		result.ReferencePoints[i] = sourceValue.Clone()
	}
	return result
}

func (o *Floor_VendorSpec_PointGrab) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorSpec_PointGrab) Merge(source *Floor_VendorSpec_PointGrab) {
	o.FloorId = source.GetFloorId()
	for _, sourceValue := range source.GetReferencePoints() {
		exists := false
		for _, currentValue := range o.ReferencePoints {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Floor_VendorSpec_PointGrab_ReferencePoint
			if sourceValue != nil {
				newDstElement = new(Floor_VendorSpec_PointGrab_ReferencePoint)
				newDstElement.Merge(sourceValue)
			}
			o.ReferencePoints = append(o.ReferencePoints, newDstElement)
		}
	}

}

func (o *Floor_VendorSpec_PointGrab) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorSpec_PointGrab))
}

func (o *Floor_VendorSpec_Mapbox) GotenObjectExt() {}

func (o *Floor_VendorSpec_Mapbox) MakeFullFieldMask() *Floor_VendorSpec_Mapbox_FieldMask {
	return FullFloor_VendorSpec_Mapbox_FieldMask()
}

func (o *Floor_VendorSpec_Mapbox) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorSpec_Mapbox_FieldMask()
}

func (o *Floor_VendorSpec_Mapbox) MakeDiffFieldMask(other *Floor_VendorSpec_Mapbox) *Floor_VendorSpec_Mapbox_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorSpec_Mapbox_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorSpec_Mapbox_FieldMask()
	}

	res := &Floor_VendorSpec_Mapbox_FieldMask{}
	if o.GetFloorPlanTilesetId() != other.GetFloorPlanTilesetId() {
		res.Paths = append(res.Paths, &FloorVendorSpecMapbox_FieldTerminalPath{selector: FloorVendorSpecMapbox_FieldPathSelectorFloorPlanTilesetId})
	}
	if o.GetBearing() != other.GetBearing() {
		res.Paths = append(res.Paths, &FloorVendorSpecMapbox_FieldTerminalPath{selector: FloorVendorSpecMapbox_FieldPathSelectorBearing})
	}
	return res
}

func (o *Floor_VendorSpec_Mapbox) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorSpec_Mapbox))
}

func (o *Floor_VendorSpec_Mapbox) Clone() *Floor_VendorSpec_Mapbox {
	if o == nil {
		return nil
	}
	result := &Floor_VendorSpec_Mapbox{}
	result.FloorPlanTilesetId = o.FloorPlanTilesetId
	result.Bearing = o.Bearing
	return result
}

func (o *Floor_VendorSpec_Mapbox) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorSpec_Mapbox) Merge(source *Floor_VendorSpec_Mapbox) {
	o.FloorPlanTilesetId = source.GetFloorPlanTilesetId()
	o.Bearing = source.GetBearing()
}

func (o *Floor_VendorSpec_Mapbox) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorSpec_Mapbox))
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) GotenObjectExt() {}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) MakeFullFieldMask() *Floor_VendorSpec_PointGrab_ReferencePoint_FieldMask {
	return FullFloor_VendorSpec_PointGrab_ReferencePoint_FieldMask()
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorSpec_PointGrab_ReferencePoint_FieldMask()
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) MakeDiffFieldMask(other *Floor_VendorSpec_PointGrab_ReferencePoint) *Floor_VendorSpec_PointGrab_ReferencePoint_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorSpec_PointGrab_ReferencePoint_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorSpec_PointGrab_ReferencePoint_FieldMask()
	}

	res := &Floor_VendorSpec_PointGrab_ReferencePoint_FieldMask{}
	if !proto.Equal(o.GetLatLng(), other.GetLatLng()) {
		res.Paths = append(res.Paths, &FloorVendorSpecPointGrabReferencePoint_FieldTerminalPath{selector: FloorVendorSpecPointGrabReferencePoint_FieldPathSelectorLatLng})
	}
	{
		subMask := o.GetXy().MakeDiffFieldMask(other.GetXy())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &FloorVendorSpecPointGrabReferencePoint_FieldTerminalPath{selector: FloorVendorSpecPointGrabReferencePoint_FieldPathSelectorXy})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &FloorVendorSpecPointGrabReferencePoint_FieldSubPath{selector: FloorVendorSpecPointGrabReferencePoint_FieldPathSelectorXy, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorSpec_PointGrab_ReferencePoint))
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) Clone() *Floor_VendorSpec_PointGrab_ReferencePoint {
	if o == nil {
		return nil
	}
	result := &Floor_VendorSpec_PointGrab_ReferencePoint{}
	result.LatLng = proto.Clone(o.LatLng).(*latlng.LatLng)
	result.Xy = o.Xy.Clone()
	return result
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) Merge(source *Floor_VendorSpec_PointGrab_ReferencePoint) {
	if source.GetLatLng() != nil {
		if o.LatLng == nil {
			o.LatLng = new(latlng.LatLng)
		}
		proto.Merge(o.LatLng, source.GetLatLng())
	}
	if source.GetXy() != nil {
		if o.Xy == nil {
			o.Xy = new(workplace_common.Point)
		}
		o.Xy.Merge(source.GetXy())
	}
}

func (o *Floor_VendorSpec_PointGrab_ReferencePoint) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorSpec_PointGrab_ReferencePoint))
}

func (o *Floor_VendorInfo_PointGrab) GotenObjectExt() {}

func (o *Floor_VendorInfo_PointGrab) MakeFullFieldMask() *Floor_VendorInfo_PointGrab_FieldMask {
	return FullFloor_VendorInfo_PointGrab_FieldMask()
}

func (o *Floor_VendorInfo_PointGrab) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorInfo_PointGrab_FieldMask()
}

func (o *Floor_VendorInfo_PointGrab) MakeDiffFieldMask(other *Floor_VendorInfo_PointGrab) *Floor_VendorInfo_PointGrab_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorInfo_PointGrab_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorInfo_PointGrab_FieldMask()
	}

	res := &Floor_VendorInfo_PointGrab_FieldMask{}

	if len(o.GetReferencePoints()) == len(other.GetReferencePoints()) {
		for i, lValue := range o.GetReferencePoints() {
			rValue := other.GetReferencePoints()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &FloorVendorInfoPointGrab_FieldTerminalPath{selector: FloorVendorInfoPointGrab_FieldPathSelectorReferencePoints})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &FloorVendorInfoPointGrab_FieldTerminalPath{selector: FloorVendorInfoPointGrab_FieldPathSelectorReferencePoints})
	}
	return res
}

func (o *Floor_VendorInfo_PointGrab) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorInfo_PointGrab))
}

func (o *Floor_VendorInfo_PointGrab) Clone() *Floor_VendorInfo_PointGrab {
	if o == nil {
		return nil
	}
	result := &Floor_VendorInfo_PointGrab{}
	result.ReferencePoints = make([]*workplace_common.ReferencePoint, len(o.ReferencePoints))
	for i, sourceValue := range o.ReferencePoints {
		result.ReferencePoints[i] = sourceValue.Clone()
	}
	return result
}

func (o *Floor_VendorInfo_PointGrab) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorInfo_PointGrab) Merge(source *Floor_VendorInfo_PointGrab) {
	for _, sourceValue := range source.GetReferencePoints() {
		exists := false
		for _, currentValue := range o.ReferencePoints {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *workplace_common.ReferencePoint
			if sourceValue != nil {
				newDstElement = new(workplace_common.ReferencePoint)
				newDstElement.Merge(sourceValue)
			}
			o.ReferencePoints = append(o.ReferencePoints, newDstElement)
		}
	}

}

func (o *Floor_VendorInfo_PointGrab) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorInfo_PointGrab))
}

func (o *Floor_VendorState_PointGrab) GotenObjectExt() {}

func (o *Floor_VendorState_PointGrab) MakeFullFieldMask() *Floor_VendorState_PointGrab_FieldMask {
	return FullFloor_VendorState_PointGrab_FieldMask()
}

func (o *Floor_VendorState_PointGrab) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorState_PointGrab_FieldMask()
}

func (o *Floor_VendorState_PointGrab) MakeDiffFieldMask(other *Floor_VendorState_PointGrab) *Floor_VendorState_PointGrab_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorState_PointGrab_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorState_PointGrab_FieldMask()
	}

	res := &Floor_VendorState_PointGrab_FieldMask{}
	{
		subMask := o.GetPeoplePositions().MakeDiffFieldMask(other.GetPeoplePositions())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &FloorVendorStatePointGrab_FieldTerminalPath{selector: FloorVendorStatePointGrab_FieldPathSelectorPeoplePositions})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &FloorVendorStatePointGrab_FieldSubPath{selector: FloorVendorStatePointGrab_FieldPathSelectorPeoplePositions, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Floor_VendorState_PointGrab) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorState_PointGrab))
}

func (o *Floor_VendorState_PointGrab) Clone() *Floor_VendorState_PointGrab {
	if o == nil {
		return nil
	}
	result := &Floor_VendorState_PointGrab{}
	result.PeoplePositions = o.PeoplePositions.Clone()
	return result
}

func (o *Floor_VendorState_PointGrab) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorState_PointGrab) Merge(source *Floor_VendorState_PointGrab) {
	if source.GetPeoplePositions() != nil {
		if o.PeoplePositions == nil {
			o.PeoplePositions = new(Floor_VendorState_PointGrab_PeoplePositions)
		}
		o.PeoplePositions.Merge(source.GetPeoplePositions())
	}
}

func (o *Floor_VendorState_PointGrab) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorState_PointGrab))
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) GotenObjectExt() {}

func (o *Floor_VendorState_PointGrab_PeoplePositions) MakeFullFieldMask() *Floor_VendorState_PointGrab_PeoplePositions_FieldMask {
	return FullFloor_VendorState_PointGrab_PeoplePositions_FieldMask()
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullFloor_VendorState_PointGrab_PeoplePositions_FieldMask()
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) MakeDiffFieldMask(other *Floor_VendorState_PointGrab_PeoplePositions) *Floor_VendorState_PointGrab_PeoplePositions_FieldMask {
	if o == nil && other == nil {
		return &Floor_VendorState_PointGrab_PeoplePositions_FieldMask{}
	}
	if o == nil || other == nil {
		return FullFloor_VendorState_PointGrab_PeoplePositions_FieldMask()
	}

	res := &Floor_VendorState_PointGrab_PeoplePositions_FieldMask{}

	if len(o.GetCoordinates()) == len(other.GetCoordinates()) {
		for i, lValue := range o.GetCoordinates() {
			rValue := other.GetCoordinates()[i]
			if !proto.Equal(lValue, rValue) {
				res.Paths = append(res.Paths, &FloorVendorStatePointGrabPeoplePositions_FieldTerminalPath{selector: FloorVendorStatePointGrabPeoplePositions_FieldPathSelectorCoordinates})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &FloorVendorStatePointGrabPeoplePositions_FieldTerminalPath{selector: FloorVendorStatePointGrabPeoplePositions_FieldPathSelectorCoordinates})
	}
	if !proto.Equal(o.GetLastReportTimestamp(), other.GetLastReportTimestamp()) {
		res.Paths = append(res.Paths, &FloorVendorStatePointGrabPeoplePositions_FieldTerminalPath{selector: FloorVendorStatePointGrabPeoplePositions_FieldPathSelectorLastReportTimestamp})
	}
	return res
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Floor_VendorState_PointGrab_PeoplePositions))
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) Clone() *Floor_VendorState_PointGrab_PeoplePositions {
	if o == nil {
		return nil
	}
	result := &Floor_VendorState_PointGrab_PeoplePositions{}
	result.Coordinates = make([]*latlng.LatLng, len(o.Coordinates))
	for i, sourceValue := range o.Coordinates {
		result.Coordinates[i] = proto.Clone(sourceValue).(*latlng.LatLng)
	}
	result.LastReportTimestamp = proto.Clone(o.LastReportTimestamp).(*timestamp.Timestamp)
	return result
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) Merge(source *Floor_VendorState_PointGrab_PeoplePositions) {
	for _, sourceValue := range source.GetCoordinates() {
		exists := false
		for _, currentValue := range o.Coordinates {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *latlng.LatLng
			if sourceValue != nil {
				newDstElement = new(latlng.LatLng)
				proto.Merge(newDstElement, sourceValue)
			}
			o.Coordinates = append(o.Coordinates, newDstElement)
		}
	}

	if source.GetLastReportTimestamp() != nil {
		if o.LastReportTimestamp == nil {
			o.LastReportTimestamp = new(timestamp.Timestamp)
		}
		proto.Merge(o.LastReportTimestamp, source.GetLastReportTimestamp())
	}
}

func (o *Floor_VendorState_PointGrab_PeoplePositions) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Floor_VendorState_PointGrab_PeoplePositions))
}
