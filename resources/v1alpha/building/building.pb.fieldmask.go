// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha/building.proto
// DO NOT EDIT!!!

package building

import (
	"encoding/json"
	"strings"

	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha/common"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha/site"
)

// ensure the imports are used
var (
	_ = json.Marshaler(nil)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = proto.Message(nil)
	_ = preflect.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldMask(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &workplace_common.BBox{}
	_ = &site.Site{}
)

type Building_FieldMask struct {
	Paths []Building_FieldPath
}

func FullBuilding_FieldMask() *Building_FieldMask {
	res := &Building_FieldMask{}
	res.Paths = append(res.Paths, &Building_FieldTerminalPath{selector: Building_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Building_FieldTerminalPath{selector: Building_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &Building_FieldTerminalPath{selector: Building_FieldPathSelectorLocation})
	res.Paths = append(res.Paths, &Building_FieldTerminalPath{selector: Building_FieldPathSelectorGeometry})
	res.Paths = append(res.Paths, &Building_FieldTerminalPath{selector: Building_FieldPathSelectorVendorSpec})
	res.Paths = append(res.Paths, &Building_FieldTerminalPath{selector: Building_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *Building_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *Building_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *Building_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseBuilding_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Building_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Building_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *Building_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseBuilding_FieldPath(raw)
	})
}

func (fieldMask *Building_FieldMask) ProtoMessage() {}

func (fieldMask *Building_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Building_FieldMask) Subtract(other *Building_FieldMask) *Building_FieldMask {
	result := &Building_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[Building_FieldPathSelector]gotenobject.FieldMask{
		Building_FieldPathSelectorLocation:   &workplace_common.StreetLocation_FieldMask{},
		Building_FieldPathSelectorGeometry:   &workplace_common.Geometry_FieldMask{},
		Building_FieldPathSelectorVendorSpec: &Building_VendorSpec_FieldMask{},
		Building_FieldPathSelectorMetadata:   &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[Building_FieldPathSelector]gotenobject.FieldMask{
		Building_FieldPathSelectorLocation:   &workplace_common.StreetLocation_FieldMask{},
		Building_FieldPathSelectorGeometry:   &workplace_common.Geometry_FieldMask{},
		Building_FieldPathSelectorVendorSpec: &Building_VendorSpec_FieldMask{},
		Building_FieldPathSelectorMetadata:   &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Building_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Building_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Building_FieldTerminalPath); ok {
					switch tp.selector {
					case Building_FieldPathSelectorLocation:
						mySubMasks[Building_FieldPathSelectorLocation] = workplace_common.FullStreetLocation_FieldMask()
					case Building_FieldPathSelectorGeometry:
						mySubMasks[Building_FieldPathSelectorGeometry] = workplace_common.FullGeometry_FieldMask()
					case Building_FieldPathSelectorVendorSpec:
						mySubMasks[Building_FieldPathSelectorVendorSpec] = FullBuilding_VendorSpec_FieldMask()
					case Building_FieldPathSelectorMetadata:
						mySubMasks[Building_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*Building_FieldSubPath); ok {
					mySubMasks[tp.selector].AppendRawPath(tp.subPath)
				}
			} else {
				result.Paths = append(result.Paths, path)
			}
		}
	}
	for selector, mySubMask := range mySubMasks {
		if mySubMask.PathsCount() > 0 {
			for _, allowedPath := range mySubMask.SubtractRaw(otherSubMasks[selector]).GetRawPaths() {
				result.Paths = append(result.Paths, &Building_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Building_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Building_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Building_FieldMask) FilterInputFields() *Building_FieldMask {
	result := &Building_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Building_FieldPathSelectorMetadata:
			if _, ok := path.(*Building_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Building_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Building_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Building_FieldSubPath{selector: Building_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Building_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Building_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Building_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseBuilding_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Building_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Building_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Building_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Building_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Building_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Building_FieldMask) AppendPath(path Building_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Building_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Building_FieldPath))
}

func (fieldMask *Building_FieldMask) GetPaths() []Building_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Building_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Building_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseBuilding_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Building_FieldMask) Set(target, source *Building) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Building_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Building), source.(*Building))
}

func (fieldMask *Building_FieldMask) Project(source *Building) *Building {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Building{}
	locationMask := &workplace_common.StreetLocation_FieldMask{}
	wholeLocationAccepted := false
	geometryMask := &workplace_common.Geometry_FieldMask{}
	wholeGeometryAccepted := false
	vendorSpecMask := &Building_VendorSpec_FieldMask{}
	wholeVendorSpecAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Building_FieldTerminalPath:
			switch tp.selector {
			case Building_FieldPathSelectorName:
				result.Name = source.Name
			case Building_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case Building_FieldPathSelectorLocation:
				result.Location = source.Location
				wholeLocationAccepted = true
			case Building_FieldPathSelectorGeometry:
				result.Geometry = source.Geometry
				wholeGeometryAccepted = true
			case Building_FieldPathSelectorVendorSpec:
				result.VendorSpec = source.VendorSpec
				wholeVendorSpecAccepted = true
			case Building_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *Building_FieldSubPath:
			switch tp.selector {
			case Building_FieldPathSelectorLocation:
				locationMask.AppendPath(tp.subPath.(workplace_common.StreetLocation_FieldPath))
			case Building_FieldPathSelectorGeometry:
				geometryMask.AppendPath(tp.subPath.(workplace_common.Geometry_FieldPath))
			case Building_FieldPathSelectorVendorSpec:
				vendorSpecMask.AppendPath(tp.subPath.(BuildingVendorSpec_FieldPath))
			case Building_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeLocationAccepted == false && len(locationMask.Paths) > 0 {
		result.Location = locationMask.Project(source.GetLocation())
	}
	if wholeGeometryAccepted == false && len(geometryMask.Paths) > 0 {
		result.Geometry = geometryMask.Project(source.GetGeometry())
	}
	if wholeVendorSpecAccepted == false && len(vendorSpecMask.Paths) > 0 {
		result.VendorSpec = vendorSpecMask.Project(source.GetVendorSpec())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *Building_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Building))
}

func (fieldMask *Building_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Building_VendorSpec_FieldMask struct {
	Paths []BuildingVendorSpec_FieldPath
}

func FullBuilding_VendorSpec_FieldMask() *Building_VendorSpec_FieldMask {
	res := &Building_VendorSpec_FieldMask{}
	res.Paths = append(res.Paths, &BuildingVendorSpec_FieldTerminalPath{selector: BuildingVendorSpec_FieldPathSelectorPointGrab})
	return res
}

func (fieldMask *Building_VendorSpec_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *Building_VendorSpec_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *Building_VendorSpec_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseBuildingVendorSpec_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Building_VendorSpec_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*BuildingVendorSpec_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *Building_VendorSpec_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseBuildingVendorSpec_FieldPath(raw)
	})
}

func (fieldMask *Building_VendorSpec_FieldMask) ProtoMessage() {}

func (fieldMask *Building_VendorSpec_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Building_VendorSpec_FieldMask) Subtract(other *Building_VendorSpec_FieldMask) *Building_VendorSpec_FieldMask {
	result := &Building_VendorSpec_FieldMask{}
	removedSelectors := make([]bool, 1)
	otherSubMasks := map[BuildingVendorSpec_FieldPathSelector]gotenobject.FieldMask{
		BuildingVendorSpec_FieldPathSelectorPointGrab: &Building_VendorSpec_PointGrab_FieldMask{},
	}
	mySubMasks := map[BuildingVendorSpec_FieldPathSelector]gotenobject.FieldMask{
		BuildingVendorSpec_FieldPathSelectorPointGrab: &Building_VendorSpec_PointGrab_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *BuildingVendorSpec_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *BuildingVendorSpec_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*BuildingVendorSpec_FieldTerminalPath); ok {
					switch tp.selector {
					case BuildingVendorSpec_FieldPathSelectorPointGrab:
						mySubMasks[BuildingVendorSpec_FieldPathSelectorPointGrab] = FullBuilding_VendorSpec_PointGrab_FieldMask()
					}
				} else if tp, ok := path.(*BuildingVendorSpec_FieldSubPath); ok {
					mySubMasks[tp.selector].AppendRawPath(tp.subPath)
				}
			} else {
				result.Paths = append(result.Paths, path)
			}
		}
	}
	for selector, mySubMask := range mySubMasks {
		if mySubMask.PathsCount() > 0 {
			for _, allowedPath := range mySubMask.SubtractRaw(otherSubMasks[selector]).GetRawPaths() {
				result.Paths = append(result.Paths, &BuildingVendorSpec_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Building_VendorSpec_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Building_VendorSpec_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Building_VendorSpec_FieldMask) FilterInputFields() *Building_VendorSpec_FieldMask {
	result := &Building_VendorSpec_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Building_VendorSpec_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Building_VendorSpec_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]BuildingVendorSpec_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseBuildingVendorSpec_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Building_VendorSpec_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Building_VendorSpec_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Building_VendorSpec_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Building_VendorSpec_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Building_VendorSpec_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Building_VendorSpec_FieldMask) AppendPath(path BuildingVendorSpec_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Building_VendorSpec_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(BuildingVendorSpec_FieldPath))
}

func (fieldMask *Building_VendorSpec_FieldMask) GetPaths() []BuildingVendorSpec_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Building_VendorSpec_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Building_VendorSpec_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseBuildingVendorSpec_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Building_VendorSpec_FieldMask) Set(target, source *Building_VendorSpec) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Building_VendorSpec_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Building_VendorSpec), source.(*Building_VendorSpec))
}

func (fieldMask *Building_VendorSpec_FieldMask) Project(source *Building_VendorSpec) *Building_VendorSpec {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Building_VendorSpec{}
	pointGrabMask := &Building_VendorSpec_PointGrab_FieldMask{}
	wholePointGrabAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *BuildingVendorSpec_FieldTerminalPath:
			switch tp.selector {
			case BuildingVendorSpec_FieldPathSelectorPointGrab:
				result.PointGrab = source.PointGrab
				wholePointGrabAccepted = true
			}
		case *BuildingVendorSpec_FieldSubPath:
			switch tp.selector {
			case BuildingVendorSpec_FieldPathSelectorPointGrab:
				pointGrabMask.AppendPath(tp.subPath.(BuildingVendorSpecPointGrab_FieldPath))
			}
		}
	}
	if wholePointGrabAccepted == false && len(pointGrabMask.Paths) > 0 {
		result.PointGrab = pointGrabMask.Project(source.GetPointGrab())
	}
	return result
}

func (fieldMask *Building_VendorSpec_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Building_VendorSpec))
}

func (fieldMask *Building_VendorSpec_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Building_VendorSpec_PointGrab_FieldMask struct {
	Paths []BuildingVendorSpecPointGrab_FieldPath
}

func FullBuilding_VendorSpec_PointGrab_FieldMask() *Building_VendorSpec_PointGrab_FieldMask {
	res := &Building_VendorSpec_PointGrab_FieldMask{}
	res.Paths = append(res.Paths, &BuildingVendorSpecPointGrab_FieldTerminalPath{selector: BuildingVendorSpecPointGrab_FieldPathSelectorBuildingId})
	return res
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseBuildingVendorSpecPointGrab_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*BuildingVendorSpecPointGrab_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseBuildingVendorSpecPointGrab_FieldPath(raw)
	})
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) ProtoMessage() {}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) Subtract(other *Building_VendorSpec_PointGrab_FieldMask) *Building_VendorSpec_PointGrab_FieldMask {
	result := &Building_VendorSpec_PointGrab_FieldMask{}
	removedSelectors := make([]bool, 1)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *BuildingVendorSpecPointGrab_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			result.Paths = append(result.Paths, path)
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Building_VendorSpec_PointGrab_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) FilterInputFields() *Building_VendorSpec_PointGrab_FieldMask {
	result := &Building_VendorSpec_PointGrab_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]BuildingVendorSpecPointGrab_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseBuildingVendorSpecPointGrab_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Building_VendorSpec_PointGrab_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Building_VendorSpec_PointGrab_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) AppendPath(path BuildingVendorSpecPointGrab_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(BuildingVendorSpecPointGrab_FieldPath))
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) GetPaths() []BuildingVendorSpecPointGrab_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseBuildingVendorSpecPointGrab_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) Set(target, source *Building_VendorSpec_PointGrab) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Building_VendorSpec_PointGrab), source.(*Building_VendorSpec_PointGrab))
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) Project(source *Building_VendorSpec_PointGrab) *Building_VendorSpec_PointGrab {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Building_VendorSpec_PointGrab{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *BuildingVendorSpecPointGrab_FieldTerminalPath:
			switch tp.selector {
			case BuildingVendorSpecPointGrab_FieldPathSelectorBuildingId:
				result.BuildingId = source.BuildingId
			}
		}
	}
	return result
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Building_VendorSpec_PointGrab))
}

func (fieldMask *Building_VendorSpec_PointGrab_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
