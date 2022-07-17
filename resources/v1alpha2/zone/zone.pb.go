// Code generated by protoc-gen-goten-go
// File: workplace/proto/v1alpha2/zone.proto
// DO NOT EDIT!!!

package zone

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &area.Area{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
	_ = &floor.Floor{}
	_ = &site.Site{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Area type
type Zone_Type int32

const (
	Zone_TYPE_UNSPECIFIED Zone_Type = 0
	Zone_COMMON_AREA      Zone_Type = 1
	Zone_BREAK_AREA       Zone_Type = 16
	Zone_BATHROOM         Zone_Type = 3
	Zone_SERVER_ROOM      Zone_Type = 4
	Zone_MECHANICAL_ROOM  Zone_Type = 10
	Zone_ELECTRICAL_ROOM  Zone_Type = 11
	Zone_STORAGE_ROOM     Zone_Type = 13
	Zone_OFFICE           Zone_Type = 2
	Zone_OPEN_OFFICE      Zone_Type = 5
	Zone_CONFERENCE_ROOM  Zone_Type = 6
	Zone_HUDDLE_SPACE     Zone_Type = 17
	Zone_RECEPTION_AREA   Zone_Type = 9
	Zone_WELLNESS_ROOM    Zone_Type = 14
	Zone_FITNESS_ROOM     Zone_Type = 15
	Zone_DESK             Zone_Type = 7
	Zone_TRAFFIC_LINE     Zone_Type = 8
)

var (
	Zone_Type_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "COMMON_AREA",
		16: "BREAK_AREA",
		3:  "BATHROOM",
		4:  "SERVER_ROOM",
		10: "MECHANICAL_ROOM",
		11: "ELECTRICAL_ROOM",
		13: "STORAGE_ROOM",
		2:  "OFFICE",
		5:  "OPEN_OFFICE",
		6:  "CONFERENCE_ROOM",
		17: "HUDDLE_SPACE",
		9:  "RECEPTION_AREA",
		14: "WELLNESS_ROOM",
		15: "FITNESS_ROOM",
		7:  "DESK",
		8:  "TRAFFIC_LINE",
	}

	Zone_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"COMMON_AREA":      1,
		"BREAK_AREA":       16,
		"BATHROOM":         3,
		"SERVER_ROOM":      4,
		"MECHANICAL_ROOM":  10,
		"ELECTRICAL_ROOM":  11,
		"STORAGE_ROOM":     13,
		"OFFICE":           2,
		"OPEN_OFFICE":      5,
		"CONFERENCE_ROOM":  6,
		"HUDDLE_SPACE":     17,
		"RECEPTION_AREA":   9,
		"WELLNESS_ROOM":    14,
		"FITNESS_ROOM":     15,
		"DESK":             7,
		"TRAFFIC_LINE":     8,
	}
)

func (x Zone_Type) Enum() *Zone_Type {
	p := new(Zone_Type)
	*p = x
	return p
}

func (x Zone_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (Zone_Type) Descriptor() preflect.EnumDescriptor {
	return workplace_proto_v1alpha2_zone_proto_enumTypes[0].Descriptor()
}

func (Zone_Type) Type() preflect.EnumType {
	return &workplace_proto_v1alpha2_zone_proto_enumTypes[0]
}

func (x Zone_Type) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use Zone_Type.ProtoReflect.Descriptor instead.
func (Zone_Type) EnumDescriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_proto_rawDescGZIP(), []int{0, 0}
}

// Zone Resource
type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Zone
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display name, e.g. "Room #203"
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Type of this area
	Type Zone_Type `protobuf:"varint,3,opt,name=type,proto3,enum=ntt.workplace.v1alpha2.Zone_Type" json:"type,omitempty" firestore:"type"`
	// Geometry of this area
	Geometry *workplace_common.Geometry `protobuf:"bytes,4,opt,name=geometry,proto3" json:"geometry,omitempty" firestore:"geometry"`
	// VendorSpec
	VendorSpec *Zone_VendorSpec `protobuf:"bytes,6,opt,name=vendor_spec,json=vendorSpec,proto3" json:"vendor_spec,omitempty" firestore:"vendorSpec"`
	State      *Zone_State      `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty" firestore:"state"`
	Metadata   *ntt_meta.Meta   `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *Zone) Reset() {
	*m = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Zone) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Zone) ProtoMessage() {}

func (m *Zone) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Zone) GotenMessage() {}

// Deprecated, Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_proto_rawDescGZIP(), []int{0}
}

func (m *Zone) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Zone) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Zone) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Zone) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Zone) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Zone) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Zone) GetType() Zone_Type {
	if m != nil {
		return m.Type
	}
	return Zone_TYPE_UNSPECIFIED
}

func (m *Zone) GetGeometry() *workplace_common.Geometry {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *Zone) GetVendorSpec() *Zone_VendorSpec {
	if m != nil {
		return m.VendorSpec
	}
	return nil
}

func (m *Zone) GetState() *Zone_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Zone) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Zone) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Zone"))
	}
	m.Name = fv
}

func (m *Zone) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "Zone"))
	}
	m.DisplayName = fv
}

func (m *Zone) SetType(fv Zone_Type) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Type", "Zone"))
	}
	m.Type = fv
}

func (m *Zone) SetGeometry(fv *workplace_common.Geometry) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Geometry", "Zone"))
	}
	m.Geometry = fv
}

func (m *Zone) SetVendorSpec(fv *Zone_VendorSpec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "VendorSpec", "Zone"))
	}
	m.VendorSpec = fv
}

func (m *Zone) SetState(fv *Zone_State) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "State", "Zone"))
	}
	m.State = fv
}

func (m *Zone) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Zone"))
	}
	m.Metadata = fv
}

// VendorSpec
type Zone_VendorSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// PointGrab Specification
	PointGrab *Zone_VendorSpec_PointGrab `protobuf:"bytes,1,opt,name=point_grab,json=pointGrab,proto3" json:"point_grab,omitempty" firestore:"pointGrab"`
}

func (m *Zone_VendorSpec) Reset() {
	*m = Zone_VendorSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Zone_VendorSpec) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Zone_VendorSpec) ProtoMessage() {}

func (m *Zone_VendorSpec) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Zone_VendorSpec) GotenMessage() {}

// Deprecated, Use Zone_VendorSpec.ProtoReflect.Descriptor instead.
func (*Zone_VendorSpec) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Zone_VendorSpec) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Zone_VendorSpec) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Zone_VendorSpec) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Zone_VendorSpec) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Zone_VendorSpec) GetPointGrab() *Zone_VendorSpec_PointGrab {
	if m != nil {
		return m.PointGrab
	}
	return nil
}

func (m *Zone_VendorSpec) SetPointGrab(fv *Zone_VendorSpec_PointGrab) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PointGrab", "Zone_VendorSpec"))
	}
	m.PointGrab = fv
}

type Zone_State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Occupancy     *Zone_State_Occupancy `protobuf:"bytes,1,opt,name=occupancy,proto3" json:"occupancy,omitempty" firestore:"occupancy"`
}

func (m *Zone_State) Reset() {
	*m = Zone_State{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Zone_State) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Zone_State) ProtoMessage() {}

func (m *Zone_State) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Zone_State) GotenMessage() {}

// Deprecated, Use Zone_State.ProtoReflect.Descriptor instead.
func (*Zone_State) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Zone_State) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Zone_State) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Zone_State) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Zone_State) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Zone_State) GetOccupancy() *Zone_State_Occupancy {
	if m != nil {
		return m.Occupancy
	}
	return nil
}

func (m *Zone_State) SetOccupancy(fv *Zone_State_Occupancy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Occupancy", "Zone_State"))
	}
	m.Occupancy = fv
}

type Zone_VendorSpec_PointGrab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// area id in PointGrab
	AreaId string `protobuf:"bytes,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty" firestore:"areaId"`
	// geo positions
	Polygon *workplace_common.Polygon `protobuf:"bytes,2,opt,name=polygon,proto3" json:"polygon,omitempty" firestore:"polygon"`
}

func (m *Zone_VendorSpec_PointGrab) Reset() {
	*m = Zone_VendorSpec_PointGrab{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Zone_VendorSpec_PointGrab) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Zone_VendorSpec_PointGrab) ProtoMessage() {}

func (m *Zone_VendorSpec_PointGrab) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Zone_VendorSpec_PointGrab) GotenMessage() {}

// Deprecated, Use Zone_VendorSpec_PointGrab.ProtoReflect.Descriptor instead.
func (*Zone_VendorSpec_PointGrab) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *Zone_VendorSpec_PointGrab) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Zone_VendorSpec_PointGrab) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Zone_VendorSpec_PointGrab) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Zone_VendorSpec_PointGrab) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Zone_VendorSpec_PointGrab) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *Zone_VendorSpec_PointGrab) GetPolygon() *workplace_common.Polygon {
	if m != nil {
		return m.Polygon
	}
	return nil
}

func (m *Zone_VendorSpec_PointGrab) SetAreaId(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AreaId", "Zone_VendorSpec_PointGrab"))
	}
	m.AreaId = fv
}

func (m *Zone_VendorSpec_PointGrab) SetPolygon(fv *workplace_common.Polygon) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Polygon", "Zone_VendorSpec_PointGrab"))
	}
	m.Polygon = fv
}

type Zone_State_Occupancy struct {
	state            protoimpl.MessageState
	sizeCache        protoimpl.SizeCache
	unknownFields    protoimpl.UnknownFields
	IsOccupied       bool                 `protobuf:"varint,1,opt,name=is_occupied,json=isOccupied,proto3" json:"is_occupied,omitempty" firestore:"isOccupied"`
	LastOccupiedTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_occupied_time,json=lastOccupiedTime,proto3" json:"last_occupied_time,omitempty" firestore:"lastOccupiedTime"`
}

func (m *Zone_State_Occupancy) Reset() {
	*m = Zone_State_Occupancy{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Zone_State_Occupancy) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Zone_State_Occupancy) ProtoMessage() {}

func (m *Zone_State_Occupancy) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Zone_State_Occupancy) GotenMessage() {}

// Deprecated, Use Zone_State_Occupancy.ProtoReflect.Descriptor instead.
func (*Zone_State_Occupancy) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (m *Zone_State_Occupancy) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Zone_State_Occupancy) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Zone_State_Occupancy) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Zone_State_Occupancy) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Zone_State_Occupancy) GetIsOccupied() bool {
	if m != nil {
		return m.IsOccupied
	}
	return false
}

func (m *Zone_State_Occupancy) GetLastOccupiedTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastOccupiedTime
	}
	return nil
}

func (m *Zone_State_Occupancy) SetIsOccupied(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IsOccupied", "Zone_State_Occupancy"))
	}
	m.IsOccupied = fv
}

func (m *Zone_State_Occupancy) SetLastOccupiedTime(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LastOccupiedTime", "Zone_State_Occupancy"))
	}
	m.LastOccupiedTime = fv
}

var workplace_proto_v1alpha2_zone_proto preflect.FileDescriptor

var workplace_proto_v1alpha2_zone_proto_rawDesc = []byte{
	0x0a, 0x23, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc1, 0x0e, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a,
	0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x48, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x38, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0xc5, 0x01, 0x0a, 0x0a, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x50, 0x0a, 0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x61,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x62, 0x52, 0x09, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x47, 0x72, 0x61, 0x62, 0x1a, 0x5f, 0x0a, 0x09, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72,
	0x61, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x07, 0x70,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x07, 0x70,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x1a, 0xcb, 0x01, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61,
	0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x1a, 0x76, 0x0a, 0x09, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64,
	0x12, 0x48, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb1, 0x02, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4d,
	0x4d, 0x4f, 0x4e, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x52,
	0x45, 0x41, 0x4b, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41,
	0x54, 0x48, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x43,
	0x48, 0x41, 0x4e, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x0a, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x4f,
	0x4d, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x52,
	0x4f, 0x4f, 0x4d, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45,
	0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x55, 0x44, 0x44, 0x4c,
	0x45, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x43,
	0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x10, 0x09, 0x12, 0x11, 0x0a,
	0x0d, 0x57, 0x45, 0x4c, 0x4c, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x0e,
	0x12, 0x10, 0x0a, 0x0c, 0x46, 0x49, 0x54, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x4f, 0x4d,
	0x10, 0x0f, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x4b, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x08, 0x3a, 0x83,
	0x06, 0xea, 0x41, 0x81, 0x03, 0x0a, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x6f, 0x6e, 0x65,
	0x12, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74,
	0x65, 0x7d, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x7d, 0x12,
	0x52, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74, 0x65,
	0x7d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f,
	0x6e, 0x65, 0x7d, 0x12, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b,
	0x73, 0x69, 0x74, 0x65, 0x7d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x7b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x66, 0x6c, 0x6f, 0x6f, 0x72,
	0x73, 0x2f, 0x7b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x7d, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f,
	0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x7d, 0x12, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x69, 0x74, 0x65,
	0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74, 0x65, 0x7d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x7b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x66, 0x6c,
	0x6f, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x7d, 0x2f, 0x61, 0x72, 0x65,
	0x61, 0x73, 0x2f, 0x7b, 0x61, 0x72, 0x65, 0x61, 0x7d, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f,
	0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x7d, 0x92, 0xd9, 0x21, 0x8c, 0x02, 0x0a, 0x05, 0x7a, 0x6f, 0x6e,
	0x65, 0x73, 0x12, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x1a, 0x04, 0x53, 0x69, 0x74, 0x65, 0x1a,
	0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72,
	0x1a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0xd8, 0x01, 0x08,
	0x02, 0x12, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x1e, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x67, 0x72, 0x61, 0x62, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x20, 0x47, 0x72, 0x61, 0x62, 0x20, 0x41, 0x72, 0x65, 0x61, 0x20, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x6f, 0x63, 0x63, 0x75, 0x70,
	0x61, 0x6e, 0x63, 0x79, 0x2e, 0x69, 0x73, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64,
	0x12, 0x0b, 0x49, 0x73, 0x20, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x12, 0x38, 0x0a,
	0x22, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79,
	0x2e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x20, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x69,
	0x65, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65, 0xda, 0x94, 0x23, 0x2a, 0x12, 0x06, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x1e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x62, 0x2e, 0x61, 0x72,
	0x65, 0x61, 0x5f, 0x69, 0x64, 0xc2, 0x85, 0x2c, 0x3b, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x0b, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x8c, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x3f,
	0x0a, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x09,
	0x5a, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x3b, 0x7a, 0x6f, 0x6e, 0x65, 0xd2, 0x84, 0xd1, 0x02, 0x45, 0x0a, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0xf2, 0x85, 0xd1, 0x02, 0x47, 0x0a, 0x0e, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x62, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0xa2, 0x80, 0xd1,
	0x02, 0x41, 0x0a, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a,
	0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	workplace_proto_v1alpha2_zone_proto_rawDescOnce sync.Once
	workplace_proto_v1alpha2_zone_proto_rawDescData = workplace_proto_v1alpha2_zone_proto_rawDesc
)

func workplace_proto_v1alpha2_zone_proto_rawDescGZIP() []byte {
	workplace_proto_v1alpha2_zone_proto_rawDescOnce.Do(func() {
		workplace_proto_v1alpha2_zone_proto_rawDescData = protoimpl.X.CompressGZIP(workplace_proto_v1alpha2_zone_proto_rawDescData)
	})
	return workplace_proto_v1alpha2_zone_proto_rawDescData
}

var workplace_proto_v1alpha2_zone_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var workplace_proto_v1alpha2_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var workplace_proto_v1alpha2_zone_proto_goTypes = []interface{}{
	(Zone_Type)(0),                    // 0: ntt.workplace.v1alpha2.Zone_Type
	(*Zone)(nil),                      // 1: ntt.workplace.v1alpha2.Zone
	(*Zone_VendorSpec)(nil),           // 2: ntt.workplace.v1alpha2.Zone.VendorSpec
	(*Zone_State)(nil),                // 3: ntt.workplace.v1alpha2.Zone.State
	(*Zone_VendorSpec_PointGrab)(nil), // 4: ntt.workplace.v1alpha2.Zone.VendorSpec.PointGrab
	(*Zone_State_Occupancy)(nil),      // 5: ntt.workplace.v1alpha2.Zone.State.Occupancy
	(*workplace_common.Geometry)(nil), // 6: ntt.workplace.v1alpha2.Geometry
	(*ntt_meta.Meta)(nil),             // 7: ntt.types.Meta
	(*workplace_common.Polygon)(nil),  // 8: ntt.workplace.v1alpha2.Polygon
	(*timestamp.Timestamp)(nil),       // 9: google.protobuf.Timestamp
}
var workplace_proto_v1alpha2_zone_proto_depIdxs = []int32{
	0, // 0: ntt.workplace.v1alpha2.Zone.type:type_name -> ntt.workplace.v1alpha2.Zone_Type
	6, // 1: ntt.workplace.v1alpha2.Zone.geometry:type_name -> ntt.workplace.v1alpha2.Geometry
	2, // 2: ntt.workplace.v1alpha2.Zone.vendor_spec:type_name -> ntt.workplace.v1alpha2.Zone.VendorSpec
	3, // 3: ntt.workplace.v1alpha2.Zone.state:type_name -> ntt.workplace.v1alpha2.Zone.State
	7, // 4: ntt.workplace.v1alpha2.Zone.metadata:type_name -> ntt.types.Meta
	4, // 5: ntt.workplace.v1alpha2.Zone.VendorSpec.point_grab:type_name -> ntt.workplace.v1alpha2.Zone.VendorSpec.PointGrab
	5, // 6: ntt.workplace.v1alpha2.Zone.State.occupancy:type_name -> ntt.workplace.v1alpha2.Zone.State.Occupancy
	8, // 7: ntt.workplace.v1alpha2.Zone.VendorSpec.PointGrab.polygon:type_name -> ntt.workplace.v1alpha2.Polygon
	9, // 8: ntt.workplace.v1alpha2.Zone.State.Occupancy.last_occupied_time:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { workplace_proto_v1alpha2_zone_proto_init() }
func workplace_proto_v1alpha2_zone_proto_init() {
	if workplace_proto_v1alpha2_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		workplace_proto_v1alpha2_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		workplace_proto_v1alpha2_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone_VendorSpec); i {
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
		workplace_proto_v1alpha2_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone_State); i {
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
		workplace_proto_v1alpha2_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone_VendorSpec_PointGrab); i {
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
		workplace_proto_v1alpha2_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone_State_Occupancy); i {
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
			RawDescriptor: workplace_proto_v1alpha2_zone_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           workplace_proto_v1alpha2_zone_proto_goTypes,
		DependencyIndexes: workplace_proto_v1alpha2_zone_proto_depIdxs,
		EnumInfos:         workplace_proto_v1alpha2_zone_proto_enumTypes,
		MessageInfos:      workplace_proto_v1alpha2_zone_proto_msgTypes,
	}.Build()
	workplace_proto_v1alpha2_zone_proto = out.File
	workplace_proto_v1alpha2_zone_proto_rawDesc = nil
	workplace_proto_v1alpha2_zone_proto_goTypes = nil
	workplace_proto_v1alpha2_zone_proto_depIdxs = nil
}
