// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/bacnet.proto
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
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &duration.Duration{}
)

func (o *BACNetEntity) GotenObjectExt() {}

func (o *BACNetEntity) MakeFullFieldMask() *BACNetEntity_FieldMask {
	return FullBACNetEntity_FieldMask()
}

func (o *BACNetEntity) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullBACNetEntity_FieldMask()
}

func (o *BACNetEntity) MakeDiffFieldMask(other *BACNetEntity) *BACNetEntity_FieldMask {
	if o == nil && other == nil {
		return &BACNetEntity_FieldMask{}
	}
	if o == nil || other == nil {
		return FullBACNetEntity_FieldMask()
	}

	res := &BACNetEntity_FieldMask{}
	if o.GetObjectType() != other.GetObjectType() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorObjectType})
	}
	if o.GetPropertyType() != other.GetPropertyType() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorPropertyType})
	}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorName})
	}
	if o.GetUuid() != other.GetUuid() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorUuid})
	}
	if o.GetUri() != other.GetUri() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorUri})
	}
	if o.GetObjectId() != other.GetObjectId() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorObjectId})
	}
	if !proto.Equal(o.GetUpdateInterval(), other.GetUpdateInterval()) {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorUpdateInterval})
	}
	if o.GetProtoUrl() != other.GetProtoUrl() {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorProtoUrl})
	}

	if len(o.GetEnumValues()) == len(other.GetEnumValues()) {
		for i, lValue := range o.GetEnumValues() {
			rValue := other.GetEnumValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorEnumValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorEnumValues})
	}

	if len(o.GetStringTags()) == len(other.GetStringTags()) {
		for i, lValue := range o.GetStringTags() {
			rValue := other.GetStringTags()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorStringTags})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorStringTags})
	}

	if len(o.GetMarkerTags()) == len(other.GetMarkerTags()) {
		for i, lValue := range o.GetMarkerTags() {
			rValue := other.GetMarkerTags()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorMarkerTags})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &BACNetEntity_FieldTerminalPath{selector: BACNetEntity_FieldPathSelectorMarkerTags})
	}
	return res
}

func (o *BACNetEntity) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*BACNetEntity))
}

func (o *BACNetEntity) Clone() *BACNetEntity {
	if o == nil {
		return nil
	}
	result := &BACNetEntity{}
	result.ObjectType = o.ObjectType
	result.PropertyType = o.PropertyType
	result.Name = o.Name
	result.Uuid = o.Uuid
	result.Uri = o.Uri
	result.ObjectId = o.ObjectId
	result.UpdateInterval = proto.Clone(o.UpdateInterval).(*duration.Duration)
	result.ProtoUrl = o.ProtoUrl
	result.EnumValues = map[string]string{}
	for key, sourceValue := range o.EnumValues {
		result.EnumValues[key] = sourceValue
	}
	result.StringTags = map[string]string{}
	for key, sourceValue := range o.StringTags {
		result.StringTags[key] = sourceValue
	}
	result.MarkerTags = make([]string, len(o.MarkerTags))
	for i, sourceValue := range o.MarkerTags {
		result.MarkerTags[i] = sourceValue
	}
	return result
}

func (o *BACNetEntity) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *BACNetEntity) Merge(source *BACNetEntity) {
	o.ObjectType = source.GetObjectType()
	o.PropertyType = source.GetPropertyType()
	o.Name = source.GetName()
	o.Uuid = source.GetUuid()
	o.Uri = source.GetUri()
	o.ObjectId = source.GetObjectId()
	if source.GetUpdateInterval() != nil {
		if o.UpdateInterval == nil {
			o.UpdateInterval = new(duration.Duration)
		}
		proto.Merge(o.UpdateInterval, source.GetUpdateInterval())
	}
	o.ProtoUrl = source.GetProtoUrl()
	if source.GetEnumValues() != nil {
		if o.EnumValues == nil {
			o.EnumValues = make(map[string]string, len(source.GetEnumValues()))
		}
		for key, sourceValue := range source.GetEnumValues() {
			o.EnumValues[key] = sourceValue
		}
	}
	if source.GetStringTags() != nil {
		if o.StringTags == nil {
			o.StringTags = make(map[string]string, len(source.GetStringTags()))
		}
		for key, sourceValue := range source.GetStringTags() {
			o.StringTags[key] = sourceValue
		}
	}
	for _, sourceValue := range source.GetMarkerTags() {
		exists := false
		for _, currentValue := range o.MarkerTags {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.MarkerTags = append(o.MarkerTags, newDstElement)
		}
	}

}

func (o *BACNetEntity) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*BACNetEntity))
}
