// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/property.proto
// DO NOT EDIT!!!

package property

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
	_ = &area.Area{}
	_ = &workplace_common.BACNetEntity{}
	_ = &building.Building{}
	_ = &device.Device{}
	_ = &floor.Floor{}
	_ = &site.Site{}
	_ = &zone.Zone{}
)

type Property_FieldMask struct {
	Paths []Property_FieldPath
}

func FullProperty_FieldMask() *Property_FieldMask {
	res := &Property_FieldMask{}
	res.Paths = append(res.Paths, &Property_FieldTerminalPath{selector: Property_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Property_FieldTerminalPath{selector: Property_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &Property_FieldTerminalPath{selector: Property_FieldPathSelectorBacnet})
	res.Paths = append(res.Paths, &Property_FieldTerminalPath{selector: Property_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &Property_FieldTerminalPath{selector: Property_FieldPathSelectorSitePlacement})
	res.Paths = append(res.Paths, &Property_FieldTerminalPath{selector: Property_FieldPathSelectorMetricOverride})
	return res
}

func (fieldMask *Property_FieldMask) String() string {
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
func (fieldMask *Property_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Property_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProperty_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Property_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Property_FieldTerminalPath); ok {
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

func (fieldMask *Property_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProperty_FieldPath(raw)
	})
}

func (fieldMask *Property_FieldMask) ProtoMessage() {}

func (fieldMask *Property_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Property_FieldMask) Subtract(other *Property_FieldMask) *Property_FieldMask {
	result := &Property_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[Property_FieldPathSelector]gotenobject.FieldMask{
		Property_FieldPathSelectorBacnet:         &workplace_common.BACNetEntity_FieldMask{},
		Property_FieldPathSelectorMetadata:       &ntt_meta.Meta_FieldMask{},
		Property_FieldPathSelectorSitePlacement:  &Property_SitePlacement_FieldMask{},
		Property_FieldPathSelectorMetricOverride: &Property_MetricOverride_FieldMask{},
	}
	mySubMasks := map[Property_FieldPathSelector]gotenobject.FieldMask{
		Property_FieldPathSelectorBacnet:         &workplace_common.BACNetEntity_FieldMask{},
		Property_FieldPathSelectorMetadata:       &ntt_meta.Meta_FieldMask{},
		Property_FieldPathSelectorSitePlacement:  &Property_SitePlacement_FieldMask{},
		Property_FieldPathSelectorMetricOverride: &Property_MetricOverride_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Property_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Property_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Property_FieldTerminalPath); ok {
					switch tp.selector {
					case Property_FieldPathSelectorBacnet:
						mySubMasks[Property_FieldPathSelectorBacnet] = workplace_common.FullBACNetEntity_FieldMask()
					case Property_FieldPathSelectorMetadata:
						mySubMasks[Property_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					case Property_FieldPathSelectorSitePlacement:
						mySubMasks[Property_FieldPathSelectorSitePlacement] = FullProperty_SitePlacement_FieldMask()
					case Property_FieldPathSelectorMetricOverride:
						mySubMasks[Property_FieldPathSelectorMetricOverride] = FullProperty_MetricOverride_FieldMask()
					}
				} else if tp, ok := path.(*Property_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Property_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Property_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Property_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Property_FieldMask) FilterInputFields() *Property_FieldMask {
	result := &Property_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Property_FieldPathSelectorMetadata:
			if _, ok := path.(*Property_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Property_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Property_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Property_FieldSubPath{selector: Property_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Property_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Property_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Property_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProperty_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Property_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Property_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Property_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Property_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Property_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Property_FieldMask) AppendPath(path Property_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Property_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Property_FieldPath))
}

func (fieldMask *Property_FieldMask) GetPaths() []Property_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Property_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Property_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProperty_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Property_FieldMask) Set(target, source *Property) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Property_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Property), source.(*Property))
}

func (fieldMask *Property_FieldMask) Project(source *Property) *Property {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Property{}
	bacnetMask := &workplace_common.BACNetEntity_FieldMask{}
	wholeBacnetAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	sitePlacementMask := &Property_SitePlacement_FieldMask{}
	wholeSitePlacementAccepted := false
	metricOverrideMask := &Property_MetricOverride_FieldMask{}
	wholeMetricOverrideAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Property_FieldTerminalPath:
			switch tp.selector {
			case Property_FieldPathSelectorName:
				result.Name = source.Name
			case Property_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case Property_FieldPathSelectorBacnet:
				result.Bacnet = source.Bacnet
				wholeBacnetAccepted = true
			case Property_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case Property_FieldPathSelectorSitePlacement:
				result.SitePlacement = source.SitePlacement
				wholeSitePlacementAccepted = true
			case Property_FieldPathSelectorMetricOverride:
				result.MetricOverride = source.MetricOverride
				wholeMetricOverrideAccepted = true
			}
		case *Property_FieldSubPath:
			switch tp.selector {
			case Property_FieldPathSelectorBacnet:
				bacnetMask.AppendPath(tp.subPath.(workplace_common.BACNetEntity_FieldPath))
			case Property_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			case Property_FieldPathSelectorSitePlacement:
				sitePlacementMask.AppendPath(tp.subPath.(PropertySitePlacement_FieldPath))
			case Property_FieldPathSelectorMetricOverride:
				metricOverrideMask.AppendPath(tp.subPath.(PropertyMetricOverride_FieldPath))
			}
		}
	}
	if wholeBacnetAccepted == false && len(bacnetMask.Paths) > 0 {
		result.Bacnet = bacnetMask.Project(source.GetBacnet())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeSitePlacementAccepted == false && len(sitePlacementMask.Paths) > 0 {
		result.SitePlacement = sitePlacementMask.Project(source.GetSitePlacement())
	}
	if wholeMetricOverrideAccepted == false && len(metricOverrideMask.Paths) > 0 {
		result.MetricOverride = metricOverrideMask.Project(source.GetMetricOverride())
	}
	return result
}

func (fieldMask *Property_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Property))
}

func (fieldMask *Property_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Property_SitePlacement_FieldMask struct {
	Paths []PropertySitePlacement_FieldPath
}

func FullProperty_SitePlacement_FieldMask() *Property_SitePlacement_FieldMask {
	res := &Property_SitePlacement_FieldMask{}
	res.Paths = append(res.Paths, &PropertySitePlacement_FieldTerminalPath{selector: PropertySitePlacement_FieldPathSelectorSite})
	res.Paths = append(res.Paths, &PropertySitePlacement_FieldTerminalPath{selector: PropertySitePlacement_FieldPathSelectorBuilding})
	res.Paths = append(res.Paths, &PropertySitePlacement_FieldTerminalPath{selector: PropertySitePlacement_FieldPathSelectorFloor})
	res.Paths = append(res.Paths, &PropertySitePlacement_FieldTerminalPath{selector: PropertySitePlacement_FieldPathSelectorArea})
	res.Paths = append(res.Paths, &PropertySitePlacement_FieldTerminalPath{selector: PropertySitePlacement_FieldPathSelectorZone})
	return res
}

func (fieldMask *Property_SitePlacement_FieldMask) String() string {
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
func (fieldMask *Property_SitePlacement_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Property_SitePlacement_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParsePropertySitePlacement_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Property_SitePlacement_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 5)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*PropertySitePlacement_FieldTerminalPath); ok {
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

func (fieldMask *Property_SitePlacement_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParsePropertySitePlacement_FieldPath(raw)
	})
}

func (fieldMask *Property_SitePlacement_FieldMask) ProtoMessage() {}

func (fieldMask *Property_SitePlacement_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Property_SitePlacement_FieldMask) Subtract(other *Property_SitePlacement_FieldMask) *Property_SitePlacement_FieldMask {
	result := &Property_SitePlacement_FieldMask{}
	removedSelectors := make([]bool, 5)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *PropertySitePlacement_FieldTerminalPath:
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

func (fieldMask *Property_SitePlacement_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Property_SitePlacement_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Property_SitePlacement_FieldMask) FilterInputFields() *Property_SitePlacement_FieldMask {
	result := &Property_SitePlacement_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Property_SitePlacement_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Property_SitePlacement_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]PropertySitePlacement_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParsePropertySitePlacement_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Property_SitePlacement_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Property_SitePlacement_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Property_SitePlacement_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Property_SitePlacement_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Property_SitePlacement_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Property_SitePlacement_FieldMask) AppendPath(path PropertySitePlacement_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Property_SitePlacement_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(PropertySitePlacement_FieldPath))
}

func (fieldMask *Property_SitePlacement_FieldMask) GetPaths() []PropertySitePlacement_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Property_SitePlacement_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Property_SitePlacement_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParsePropertySitePlacement_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Property_SitePlacement_FieldMask) Set(target, source *Property_SitePlacement) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Property_SitePlacement_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Property_SitePlacement), source.(*Property_SitePlacement))
}

func (fieldMask *Property_SitePlacement_FieldMask) Project(source *Property_SitePlacement) *Property_SitePlacement {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Property_SitePlacement{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *PropertySitePlacement_FieldTerminalPath:
			switch tp.selector {
			case PropertySitePlacement_FieldPathSelectorSite:
				result.Site = source.Site
			case PropertySitePlacement_FieldPathSelectorBuilding:
				result.Building = source.Building
			case PropertySitePlacement_FieldPathSelectorFloor:
				result.Floor = source.Floor
			case PropertySitePlacement_FieldPathSelectorArea:
				result.Area = source.Area
			case PropertySitePlacement_FieldPathSelectorZone:
				result.Zone = source.Zone
			}
		}
	}
	return result
}

func (fieldMask *Property_SitePlacement_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Property_SitePlacement))
}

func (fieldMask *Property_SitePlacement_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Property_MetricOverride_FieldMask struct {
	Paths []PropertyMetricOverride_FieldPath
}

func FullProperty_MetricOverride_FieldMask() *Property_MetricOverride_FieldMask {
	res := &Property_MetricOverride_FieldMask{}
	res.Paths = append(res.Paths, &PropertyMetricOverride_FieldTerminalPath{selector: PropertyMetricOverride_FieldPathSelectorPropertyName})
	res.Paths = append(res.Paths, &PropertyMetricOverride_FieldTerminalPath{selector: PropertyMetricOverride_FieldPathSelectorUnit})
	return res
}

func (fieldMask *Property_MetricOverride_FieldMask) String() string {
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
func (fieldMask *Property_MetricOverride_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Property_MetricOverride_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParsePropertyMetricOverride_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Property_MetricOverride_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*PropertyMetricOverride_FieldTerminalPath); ok {
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

func (fieldMask *Property_MetricOverride_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParsePropertyMetricOverride_FieldPath(raw)
	})
}

func (fieldMask *Property_MetricOverride_FieldMask) ProtoMessage() {}

func (fieldMask *Property_MetricOverride_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Property_MetricOverride_FieldMask) Subtract(other *Property_MetricOverride_FieldMask) *Property_MetricOverride_FieldMask {
	result := &Property_MetricOverride_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *PropertyMetricOverride_FieldTerminalPath:
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

func (fieldMask *Property_MetricOverride_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Property_MetricOverride_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Property_MetricOverride_FieldMask) FilterInputFields() *Property_MetricOverride_FieldMask {
	result := &Property_MetricOverride_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Property_MetricOverride_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Property_MetricOverride_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]PropertyMetricOverride_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParsePropertyMetricOverride_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Property_MetricOverride_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Property_MetricOverride_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Property_MetricOverride_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Property_MetricOverride_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Property_MetricOverride_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Property_MetricOverride_FieldMask) AppendPath(path PropertyMetricOverride_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Property_MetricOverride_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(PropertyMetricOverride_FieldPath))
}

func (fieldMask *Property_MetricOverride_FieldMask) GetPaths() []PropertyMetricOverride_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Property_MetricOverride_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Property_MetricOverride_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParsePropertyMetricOverride_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Property_MetricOverride_FieldMask) Set(target, source *Property_MetricOverride) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Property_MetricOverride_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Property_MetricOverride), source.(*Property_MetricOverride))
}

func (fieldMask *Property_MetricOverride_FieldMask) Project(source *Property_MetricOverride) *Property_MetricOverride {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Property_MetricOverride{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *PropertyMetricOverride_FieldTerminalPath:
			switch tp.selector {
			case PropertyMetricOverride_FieldPathSelectorPropertyName:
				result.PropertyName = source.PropertyName
			case PropertyMetricOverride_FieldPathSelectorUnit:
				result.Unit = source.Unit
			}
		}
	}
	return result
}

func (fieldMask *Property_MetricOverride_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Property_MetricOverride))
}

func (fieldMask *Property_MetricOverride_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
