// Code generated by protoc-gen-goten-go
// File: workplace/proto/v1alpha2/bacnet.proto
// DO NOT EDIT!!!

package workplace_common

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	duration "github.com/golang/protobuf/ptypes/duration"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var (
	_ = &duration.Duration{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Entity populated from BACNet system.
//
// In order to retrieve property value (present or historical),
// query [monitoring][ntt.monitoring.v3.Monitoring] service with resource type
// `workplace.edgelq.com/Property` and metric type depending on [property
// type][ntt.workplace.v1alpha2.BACNetEntity.property_type] and presence of
// [enum values][ntt.workplace.v1alpha2.BACNetEntity.enum_values], either metric
// type:
// * `workplace.edgelq.com/Property/value/enum`
// * `workplace.edgelq.com/Property/value/analog`
// * `workplace.edgelq.com/Property/value/binary`
type BACNetEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Object type, e.g.: DNET Analog Input
	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty" firestore:"objectType"`
	// BACnet property_type, e.g. AI, AO, BI, BO, AV
	PropertyType string `protobuf:"bytes,2,opt,name=property_type,json=propertyType,proto3" json:"property_type,omitempty" firestore:"propertyType"`
	// Name, e.g. S1 - Heat Exchanger Liquid Temperature
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// UUID of this object
	Uuid string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty" firestore:"uuid"`
	// Full URI (riptide format)
	Uri string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty" firestore:"uri"`
	// BACNet object id
	ObjectId int64 `protobuf:"varint,6,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty" firestore:"objectId"`
	// Property update interval
	UpdateInterval *duration.Duration `protobuf:"bytes,7,opt,name=update_interval,json=updateInterval,proto3" json:"update_interval,omitempty" firestore:"updateInterval"`
	// Proto URL provides details about backend location
	ProtoUrl string `protobuf:"bytes,9,opt,name=proto_url,json=protoUrl,proto3" json:"proto_url,omitempty" firestore:"protoUrl"`
	// Enum value mappings: decimal integer of value to readable string value
	EnumValues map[string]string `protobuf:"bytes,11,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"enumValues"`
	// String tags (or key-value tags)
	StringTags map[string]string `protobuf:"bytes,12,rep,name=string_tags,json=stringTags,proto3" json:"string_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"stringTags"`
	// Marker tags (denoting presence)
	MarkerTags []string `protobuf:"bytes,13,rep,name=marker_tags,json=markerTags,proto3" json:"marker_tags,omitempty" firestore:"markerTags"`
}

func (m *BACNetEntity) Reset() {
	*m = BACNetEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_bacnet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BACNetEntity) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BACNetEntity) ProtoMessage() {}

func (m *BACNetEntity) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_bacnet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BACNetEntity) GotenMessage() {}

// Deprecated, Use BACNetEntity.ProtoReflect.Descriptor instead.
func (*BACNetEntity) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_bacnet_proto_rawDescGZIP(), []int{0}
}

func (m *BACNetEntity) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BACNetEntity) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BACNetEntity) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BACNetEntity) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BACNetEntity) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *BACNetEntity) GetPropertyType() string {
	if m != nil {
		return m.PropertyType
	}
	return ""
}

func (m *BACNetEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BACNetEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BACNetEntity) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *BACNetEntity) GetObjectId() int64 {
	if m != nil {
		return m.ObjectId
	}
	return int64(0)
}

func (m *BACNetEntity) GetUpdateInterval() *duration.Duration {
	if m != nil {
		return m.UpdateInterval
	}
	return nil
}

func (m *BACNetEntity) GetProtoUrl() string {
	if m != nil {
		return m.ProtoUrl
	}
	return ""
}

func (m *BACNetEntity) GetEnumValues() map[string]string {
	if m != nil {
		return m.EnumValues
	}
	return nil
}

func (m *BACNetEntity) GetStringTags() map[string]string {
	if m != nil {
		return m.StringTags
	}
	return nil
}

func (m *BACNetEntity) GetMarkerTags() []string {
	if m != nil {
		return m.MarkerTags
	}
	return nil
}

func (m *BACNetEntity) SetObjectType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ObjectType", "BACNetEntity"))
	}
	m.ObjectType = fv
}

func (m *BACNetEntity) SetPropertyType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PropertyType", "BACNetEntity"))
	}
	m.PropertyType = fv
}

func (m *BACNetEntity) SetName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "BACNetEntity"))
	}
	m.Name = fv
}

func (m *BACNetEntity) SetUuid(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Uuid", "BACNetEntity"))
	}
	m.Uuid = fv
}

func (m *BACNetEntity) SetUri(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Uri", "BACNetEntity"))
	}
	m.Uri = fv
}

func (m *BACNetEntity) SetObjectId(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ObjectId", "BACNetEntity"))
	}
	m.ObjectId = fv
}

func (m *BACNetEntity) SetUpdateInterval(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "UpdateInterval", "BACNetEntity"))
	}
	m.UpdateInterval = fv
}

func (m *BACNetEntity) SetProtoUrl(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProtoUrl", "BACNetEntity"))
	}
	m.ProtoUrl = fv
}

func (m *BACNetEntity) SetEnumValues(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EnumValues", "BACNetEntity"))
	}
	m.EnumValues = fv
}

func (m *BACNetEntity) SetStringTags(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "StringTags", "BACNetEntity"))
	}
	m.StringTags = fv
}

func (m *BACNetEntity) SetMarkerTags(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MarkerTags", "BACNetEntity"))
	}
	m.MarkerTags = fv
}

var workplace_proto_v1alpha2_bacnet_proto preflect.FileDescriptor

var workplace_proto_v1alpha2_bacnet_proto_rawDesc = []byte{
	0x0a, 0x25, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x62, 0x61, 0x63, 0x6e, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x04, 0x0a, 0x0c, 0x42,
	0x41, 0x43, 0x4e, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x55, 0x0a, 0x0b, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x41, 0x43, 0x4e, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x55, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42,
	0x41, 0x43, 0x4e, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x72, 0x54, 0x61, 0x67, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x45, 0x6e, 0x75, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x7c, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0b, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	workplace_proto_v1alpha2_bacnet_proto_rawDescOnce sync.Once
	workplace_proto_v1alpha2_bacnet_proto_rawDescData = workplace_proto_v1alpha2_bacnet_proto_rawDesc
)

func workplace_proto_v1alpha2_bacnet_proto_rawDescGZIP() []byte {
	workplace_proto_v1alpha2_bacnet_proto_rawDescOnce.Do(func() {
		workplace_proto_v1alpha2_bacnet_proto_rawDescData = protoimpl.X.CompressGZIP(workplace_proto_v1alpha2_bacnet_proto_rawDescData)
	})
	return workplace_proto_v1alpha2_bacnet_proto_rawDescData
}

var workplace_proto_v1alpha2_bacnet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var workplace_proto_v1alpha2_bacnet_proto_goTypes = []interface{}{
	(*BACNetEntity)(nil),      // 0: ntt.workplace.v1alpha2.BACNetEntity
	nil,                       // 1: ntt.workplace.v1alpha2.BACNetEntity.EnumValuesEntry
	nil,                       // 2: ntt.workplace.v1alpha2.BACNetEntity.StringTagsEntry
	(*duration.Duration)(nil), // 3: google.protobuf.Duration
}
var workplace_proto_v1alpha2_bacnet_proto_depIdxs = []int32{
	3, // 0: ntt.workplace.v1alpha2.BACNetEntity.update_interval:type_name -> google.protobuf.Duration
	1, // 1: ntt.workplace.v1alpha2.BACNetEntity.enum_values:type_name -> ntt.workplace.v1alpha2.BACNetEntity.EnumValuesEntry
	2, // 2: ntt.workplace.v1alpha2.BACNetEntity.string_tags:type_name -> ntt.workplace.v1alpha2.BACNetEntity.StringTagsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { workplace_proto_v1alpha2_bacnet_proto_init() }
func workplace_proto_v1alpha2_bacnet_proto_init() {
	if workplace_proto_v1alpha2_bacnet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		workplace_proto_v1alpha2_bacnet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BACNetEntity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: workplace_proto_v1alpha2_bacnet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           workplace_proto_v1alpha2_bacnet_proto_goTypes,
		DependencyIndexes: workplace_proto_v1alpha2_bacnet_proto_depIdxs,
		MessageInfos:      workplace_proto_v1alpha2_bacnet_proto_msgTypes,
	}.Build()
	workplace_proto_v1alpha2_bacnet_proto = out.File
	workplace_proto_v1alpha2_bacnet_proto_rawDesc = nil
	workplace_proto_v1alpha2_bacnet_proto_goTypes = nil
	workplace_proto_v1alpha2_bacnet_proto_depIdxs = nil
}
