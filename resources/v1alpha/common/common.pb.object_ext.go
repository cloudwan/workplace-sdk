// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha/common.proto
// DO NOT EDIT!!!

package workplace_common

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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
	_ = &latlng.LatLng{}
)

func (o *BBox) GotenObjectExt() {}

func (o *BBox) MakeFullFieldMask() *BBox_FieldMask {
	return FullBBox_FieldMask()
}

func (o *BBox) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullBBox_FieldMask()
}

func (o *BBox) MakeDiffFieldMask(other *BBox) *BBox_FieldMask {
	if o == nil && other == nil {
		return &BBox_FieldMask{}
	}
	if o == nil || other == nil {
		return FullBBox_FieldMask()
	}

	res := &BBox_FieldMask{}
	if !proto.Equal(o.GetSouthWest(), other.GetSouthWest()) {
		res.Paths = append(res.Paths, &BBox_FieldTerminalPath{selector: BBox_FieldPathSelectorSouthWest})
	}
	if !proto.Equal(o.GetNorthEast(), other.GetNorthEast()) {
		res.Paths = append(res.Paths, &BBox_FieldTerminalPath{selector: BBox_FieldPathSelectorNorthEast})
	}
	return res
}

func (o *BBox) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*BBox))
}

func (o *BBox) Clone() *BBox {
	if o == nil {
		return nil
	}
	result := &BBox{}
	result.SouthWest = proto.Clone(o.SouthWest).(*latlng.LatLng)
	result.NorthEast = proto.Clone(o.NorthEast).(*latlng.LatLng)
	return result
}

func (o *BBox) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *BBox) Merge(source *BBox) {
	if source.GetSouthWest() != nil {
		if o.SouthWest == nil {
			o.SouthWest = new(latlng.LatLng)
		}
		proto.Merge(o.SouthWest, source.GetSouthWest())
	}
	if source.GetNorthEast() != nil {
		if o.NorthEast == nil {
			o.NorthEast = new(latlng.LatLng)
		}
		proto.Merge(o.NorthEast, source.GetNorthEast())
	}
}

func (o *BBox) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*BBox))
}

func (o *Polygon) GotenObjectExt() {}

func (o *Polygon) MakeFullFieldMask() *Polygon_FieldMask {
	return FullPolygon_FieldMask()
}

func (o *Polygon) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPolygon_FieldMask()
}

func (o *Polygon) MakeDiffFieldMask(other *Polygon) *Polygon_FieldMask {
	if o == nil && other == nil {
		return &Polygon_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPolygon_FieldMask()
	}

	res := &Polygon_FieldMask{}

	if len(o.GetCoordinates()) == len(other.GetCoordinates()) {
		for i, lValue := range o.GetCoordinates() {
			rValue := other.GetCoordinates()[i]
			if !proto.Equal(lValue, rValue) {
				res.Paths = append(res.Paths, &Polygon_FieldTerminalPath{selector: Polygon_FieldPathSelectorCoordinates})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Polygon_FieldTerminalPath{selector: Polygon_FieldPathSelectorCoordinates})
	}
	return res
}

func (o *Polygon) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Polygon))
}

func (o *Polygon) Clone() *Polygon {
	if o == nil {
		return nil
	}
	result := &Polygon{}
	result.Coordinates = make([]*latlng.LatLng, len(o.Coordinates))
	for i, sourceValue := range o.Coordinates {
		result.Coordinates[i] = proto.Clone(sourceValue).(*latlng.LatLng)
	}
	return result
}

func (o *Polygon) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Polygon) Merge(source *Polygon) {
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

}

func (o *Polygon) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Polygon))
}

func (o *Geometry) GotenObjectExt() {}

func (o *Geometry) MakeFullFieldMask() *Geometry_FieldMask {
	return FullGeometry_FieldMask()
}

func (o *Geometry) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullGeometry_FieldMask()
}

func (o *Geometry) MakeDiffFieldMask(other *Geometry) *Geometry_FieldMask {
	if o == nil && other == nil {
		return &Geometry_FieldMask{}
	}
	if o == nil || other == nil {
		return FullGeometry_FieldMask()
	}

	res := &Geometry_FieldMask{}
	if !proto.Equal(o.GetCenter(), other.GetCenter()) {
		res.Paths = append(res.Paths, &Geometry_FieldTerminalPath{selector: Geometry_FieldPathSelectorCenter})
	}
	{
		subMask := o.GetBbox().MakeDiffFieldMask(other.GetBbox())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Geometry_FieldTerminalPath{selector: Geometry_FieldPathSelectorBbox})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Geometry_FieldSubPath{selector: Geometry_FieldPathSelectorBbox, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetPolygon().MakeDiffFieldMask(other.GetPolygon())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Geometry_FieldTerminalPath{selector: Geometry_FieldPathSelectorPolygon})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Geometry_FieldSubPath{selector: Geometry_FieldPathSelectorPolygon, subPath: subpath})
			}
		}
	}
	if o.GetPanning() != other.GetPanning() {
		res.Paths = append(res.Paths, &Geometry_FieldTerminalPath{selector: Geometry_FieldPathSelectorPanning})
	}
	return res
}

func (o *Geometry) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Geometry))
}

func (o *Geometry) Clone() *Geometry {
	if o == nil {
		return nil
	}
	result := &Geometry{}
	result.Center = proto.Clone(o.Center).(*latlng.LatLng)
	result.Bbox = o.Bbox.Clone()
	result.Polygon = o.Polygon.Clone()
	result.Panning = o.Panning
	return result
}

func (o *Geometry) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Geometry) Merge(source *Geometry) {
	if source.GetCenter() != nil {
		if o.Center == nil {
			o.Center = new(latlng.LatLng)
		}
		proto.Merge(o.Center, source.GetCenter())
	}
	if source.GetBbox() != nil {
		if o.Bbox == nil {
			o.Bbox = new(BBox)
		}
		o.Bbox.Merge(source.GetBbox())
	}
	if source.GetPolygon() != nil {
		if o.Polygon == nil {
			o.Polygon = new(Polygon)
		}
		o.Polygon.Merge(source.GetPolygon())
	}
	o.Panning = source.GetPanning()
}

func (o *Geometry) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Geometry))
}

func (o *StreetLocation) GotenObjectExt() {}

func (o *StreetLocation) MakeFullFieldMask() *StreetLocation_FieldMask {
	return FullStreetLocation_FieldMask()
}

func (o *StreetLocation) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullStreetLocation_FieldMask()
}

func (o *StreetLocation) MakeDiffFieldMask(other *StreetLocation) *StreetLocation_FieldMask {
	if o == nil && other == nil {
		return &StreetLocation_FieldMask{}
	}
	if o == nil || other == nil {
		return FullStreetLocation_FieldMask()
	}

	res := &StreetLocation_FieldMask{}
	if o.GetStreetAddress() != other.GetStreetAddress() {
		res.Paths = append(res.Paths, &StreetLocation_FieldTerminalPath{selector: StreetLocation_FieldPathSelectorStreetAddress})
	}
	if !proto.Equal(o.GetStreetCoordinates(), other.GetStreetCoordinates()) {
		res.Paths = append(res.Paths, &StreetLocation_FieldTerminalPath{selector: StreetLocation_FieldPathSelectorStreetCoordinates})
	}
	return res
}

func (o *StreetLocation) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*StreetLocation))
}

func (o *StreetLocation) Clone() *StreetLocation {
	if o == nil {
		return nil
	}
	result := &StreetLocation{}
	result.StreetAddress = o.StreetAddress
	result.StreetCoordinates = proto.Clone(o.StreetCoordinates).(*latlng.LatLng)
	return result
}

func (o *StreetLocation) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *StreetLocation) Merge(source *StreetLocation) {
	o.StreetAddress = source.GetStreetAddress()
	if source.GetStreetCoordinates() != nil {
		if o.StreetCoordinates == nil {
			o.StreetCoordinates = new(latlng.LatLng)
		}
		proto.Merge(o.StreetCoordinates, source.GetStreetCoordinates())
	}
}

func (o *StreetLocation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*StreetLocation))
}

func (o *VendorMapping) GotenObjectExt() {}

func (o *VendorMapping) MakeFullFieldMask() *VendorMapping_FieldMask {
	return FullVendorMapping_FieldMask()
}

func (o *VendorMapping) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullVendorMapping_FieldMask()
}

func (o *VendorMapping) MakeDiffFieldMask(other *VendorMapping) *VendorMapping_FieldMask {
	if o == nil && other == nil {
		return &VendorMapping_FieldMask{}
	}
	if o == nil || other == nil {
		return FullVendorMapping_FieldMask()
	}

	res := &VendorMapping_FieldMask{}
	if o.GetVendor() != other.GetVendor() {
		res.Paths = append(res.Paths, &VendorMapping_FieldTerminalPath{selector: VendorMapping_FieldPathSelectorVendor})
	}
	if o.GetId() != other.GetId() {
		res.Paths = append(res.Paths, &VendorMapping_FieldTerminalPath{selector: VendorMapping_FieldPathSelectorId})
	}
	return res
}

func (o *VendorMapping) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*VendorMapping))
}

func (o *VendorMapping) Clone() *VendorMapping {
	if o == nil {
		return nil
	}
	result := &VendorMapping{}
	result.Vendor = o.Vendor
	result.Id = o.Id
	return result
}

func (o *VendorMapping) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *VendorMapping) Merge(source *VendorMapping) {
	o.Vendor = source.GetVendor()
	o.Id = source.GetId()
}

func (o *VendorMapping) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*VendorMapping))
}

func (o *Point) GotenObjectExt() {}

func (o *Point) MakeFullFieldMask() *Point_FieldMask {
	return FullPoint_FieldMask()
}

func (o *Point) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPoint_FieldMask()
}

func (o *Point) MakeDiffFieldMask(other *Point) *Point_FieldMask {
	if o == nil && other == nil {
		return &Point_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPoint_FieldMask()
	}

	res := &Point_FieldMask{}
	if o.GetX() != other.GetX() {
		res.Paths = append(res.Paths, &Point_FieldTerminalPath{selector: Point_FieldPathSelectorX})
	}
	if o.GetY() != other.GetY() {
		res.Paths = append(res.Paths, &Point_FieldTerminalPath{selector: Point_FieldPathSelectorY})
	}
	return res
}

func (o *Point) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Point))
}

func (o *Point) Clone() *Point {
	if o == nil {
		return nil
	}
	result := &Point{}
	result.X = o.X
	result.Y = o.Y
	return result
}

func (o *Point) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Point) Merge(source *Point) {
	o.X = source.GetX()
	o.Y = source.GetY()
}

func (o *Point) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Point))
}

func (o *ReferencePoint) GotenObjectExt() {}

func (o *ReferencePoint) MakeFullFieldMask() *ReferencePoint_FieldMask {
	return FullReferencePoint_FieldMask()
}

func (o *ReferencePoint) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullReferencePoint_FieldMask()
}

func (o *ReferencePoint) MakeDiffFieldMask(other *ReferencePoint) *ReferencePoint_FieldMask {
	if o == nil && other == nil {
		return &ReferencePoint_FieldMask{}
	}
	if o == nil || other == nil {
		return FullReferencePoint_FieldMask()
	}

	res := &ReferencePoint_FieldMask{}
	{
		subMask := o.GetPoint().MakeDiffFieldMask(other.GetPoint())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ReferencePoint_FieldTerminalPath{selector: ReferencePoint_FieldPathSelectorPoint})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ReferencePoint_FieldSubPath{selector: ReferencePoint_FieldPathSelectorPoint, subPath: subpath})
			}
		}
	}
	if !proto.Equal(o.GetLatLng(), other.GetLatLng()) {
		res.Paths = append(res.Paths, &ReferencePoint_FieldTerminalPath{selector: ReferencePoint_FieldPathSelectorLatLng})
	}
	return res
}

func (o *ReferencePoint) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ReferencePoint))
}

func (o *ReferencePoint) Clone() *ReferencePoint {
	if o == nil {
		return nil
	}
	result := &ReferencePoint{}
	result.Point = o.Point.Clone()
	result.LatLng = proto.Clone(o.LatLng).(*latlng.LatLng)
	return result
}

func (o *ReferencePoint) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ReferencePoint) Merge(source *ReferencePoint) {
	if source.GetPoint() != nil {
		if o.Point == nil {
			o.Point = new(Point)
		}
		o.Point.Merge(source.GetPoint())
	}
	if source.GetLatLng() != nil {
		if o.LatLng == nil {
			o.LatLng = new(latlng.LatLng)
		}
		proto.Merge(o.LatLng, source.GetLatLng())
	}
}

func (o *ReferencePoint) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ReferencePoint))
}
