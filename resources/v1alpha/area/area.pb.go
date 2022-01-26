// Code generated by protoc-gen-goten-go
// File: workplace/proto/v1alpha/area.proto
// DO NOT EDIT!!!

package area

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
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha/common"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha/floor"
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
	_ = &workplace_common.BBox{}
	_ = &floor.Floor{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Area Resource
type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Area
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display name, e.g. "Room #203"
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Geometry of this area
	Geometry *workplace_common.Geometry `protobuf:"bytes,4,opt,name=geometry,proto3" json:"geometry,omitempty" firestore:"geometry"`
	// VendorSpec
	VendorSpec *Area_VendorSpec `protobuf:"bytes,6,opt,name=vendor_spec,json=vendorSpec,proto3" json:"vendor_spec,omitempty" firestore:"vendorSpec"`
	Metadata   *ntt_meta.Meta   `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *Area) Reset() {
	*m = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha_area_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Area) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Area) ProtoMessage() {}

func (m *Area) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha_area_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Area) GotenMessage() {}

// Deprecated, Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha_area_proto_rawDescGZIP(), []int{0}
}

func (m *Area) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Area) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Area) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Area) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Area) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Area) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Area) GetGeometry() *workplace_common.Geometry {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *Area) GetVendorSpec() *Area_VendorSpec {
	if m != nil {
		return m.VendorSpec
	}
	return nil
}

func (m *Area) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Area) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Area"))
	}
	m.Name = fv
}

func (m *Area) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "Area"))
	}
	m.DisplayName = fv
}

func (m *Area) SetGeometry(fv *workplace_common.Geometry) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Geometry", "Area"))
	}
	m.Geometry = fv
}

func (m *Area) SetVendorSpec(fv *Area_VendorSpec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "VendorSpec", "Area"))
	}
	m.VendorSpec = fv
}

func (m *Area) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Area"))
	}
	m.Metadata = fv
}

// VendorSpec
type Area_VendorSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// PointGrab Specification
	PointGrab *Area_VendorSpec_PointGrab `protobuf:"bytes,1,opt,name=point_grab,json=pointGrab,proto3" json:"point_grab,omitempty" firestore:"pointGrab"`
}

func (m *Area_VendorSpec) Reset() {
	*m = Area_VendorSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha_area_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Area_VendorSpec) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Area_VendorSpec) ProtoMessage() {}

func (m *Area_VendorSpec) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha_area_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Area_VendorSpec) GotenMessage() {}

// Deprecated, Use Area_VendorSpec.ProtoReflect.Descriptor instead.
func (*Area_VendorSpec) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha_area_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Area_VendorSpec) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Area_VendorSpec) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Area_VendorSpec) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Area_VendorSpec) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Area_VendorSpec) GetPointGrab() *Area_VendorSpec_PointGrab {
	if m != nil {
		return m.PointGrab
	}
	return nil
}

func (m *Area_VendorSpec) SetPointGrab(fv *Area_VendorSpec_PointGrab) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PointGrab", "Area_VendorSpec"))
	}
	m.PointGrab = fv
}

type Area_VendorSpec_PointGrab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// building id in PointGrab
	BuildingId string `protobuf:"bytes,1,opt,name=building_id,json=buildingId,proto3" json:"building_id,omitempty" firestore:"buildingId"`
}

func (m *Area_VendorSpec_PointGrab) Reset() {
	*m = Area_VendorSpec_PointGrab{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha_area_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Area_VendorSpec_PointGrab) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Area_VendorSpec_PointGrab) ProtoMessage() {}

func (m *Area_VendorSpec_PointGrab) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha_area_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Area_VendorSpec_PointGrab) GotenMessage() {}

// Deprecated, Use Area_VendorSpec_PointGrab.ProtoReflect.Descriptor instead.
func (*Area_VendorSpec_PointGrab) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha_area_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *Area_VendorSpec_PointGrab) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Area_VendorSpec_PointGrab) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Area_VendorSpec_PointGrab) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Area_VendorSpec_PointGrab) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Area_VendorSpec_PointGrab) GetBuildingId() string {
	if m != nil {
		return m.BuildingId
	}
	return ""
}

func (m *Area_VendorSpec_PointGrab) SetBuildingId(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "BuildingId", "Area_VendorSpec_PointGrab"))
	}
	m.BuildingId = fv
}

var workplace_proto_v1alpha_area_proto preflect.FileDescriptor

var workplace_proto_v1alpha_area_proto_rawDesc = []byte{
	0x0a, 0x22, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x05, 0x0a, 0x04, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x41, 0x72, 0x65, 0x61, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2b, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x8b, 0x01, 0x0a, 0x0a, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4f, 0x0a, 0x0a, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x67, 0x72, 0x61, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x62, 0x52, 0x09,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x62, 0x1a, 0x2c, 0x0a, 0x09, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x47, 0x72, 0x61, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x3a, 0x90, 0x02, 0xea, 0x41, 0x6d, 0x0a, 0x19, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x65, 0x61, 0x12, 0x50, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x69, 0x74,
	0x65, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74, 0x65, 0x7d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x7d, 0x2f, 0x61, 0x72,
	0x65, 0x61, 0x73, 0x2f, 0x7b, 0x61, 0x72, 0x65, 0x61, 0x7d, 0x92, 0xd9, 0x21, 0x4d, 0x0a, 0x05,
	0x61, 0x72, 0x65, 0x61, 0x73, 0x12, 0x05, 0x61, 0x72, 0x65, 0x61, 0x73, 0x1a, 0x05, 0x46, 0x6c,
	0x6f, 0x6f, 0x72, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x2e, 0x08, 0x02, 0x12, 0x0c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x02, 0xfa, 0xde, 0x21, 0x06, 0x0a,
	0x04, 0x41, 0x72, 0x65, 0x61, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0xc2, 0x85, 0x2c, 0x35, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x22, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x88, 0x03, 0xe8, 0xde, 0x21,
	0x01, 0xd2, 0xff, 0xd0, 0x02, 0x3e, 0x0a, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x61, 0x72, 0x65, 0x61, 0x92, 0x8c, 0xd1, 0x02, 0x48, 0x0a, 0x0f, 0x61, 0x72, 0x65, 0x61, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x72, 0x65,
	0x61, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x09, 0x41, 0x72, 0x65, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x72,
	0x65, 0x61, 0x3b, 0x61, 0x72, 0x65, 0x61, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xa2, 0x80, 0xd1, 0x02, 0x40, 0x0a, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x61, 0x72, 0x65, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	workplace_proto_v1alpha_area_proto_rawDescOnce sync.Once
	workplace_proto_v1alpha_area_proto_rawDescData = workplace_proto_v1alpha_area_proto_rawDesc
)

func workplace_proto_v1alpha_area_proto_rawDescGZIP() []byte {
	workplace_proto_v1alpha_area_proto_rawDescOnce.Do(func() {
		workplace_proto_v1alpha_area_proto_rawDescData = protoimpl.X.CompressGZIP(workplace_proto_v1alpha_area_proto_rawDescData)
	})
	return workplace_proto_v1alpha_area_proto_rawDescData
}

var workplace_proto_v1alpha_area_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var workplace_proto_v1alpha_area_proto_goTypes = []interface{}{
	(*Area)(nil),                      // 0: ntt.workplace.v1alpha.Area
	(*Area_VendorSpec)(nil),           // 1: ntt.workplace.v1alpha.Area.VendorSpec
	(*Area_VendorSpec_PointGrab)(nil), // 2: ntt.workplace.v1alpha.Area.VendorSpec.PointGrab
	(*workplace_common.Geometry)(nil), // 3: ntt.workplace.v1alpha.Geometry
	(*ntt_meta.Meta)(nil),             // 4: ntt.types.Meta
}
var workplace_proto_v1alpha_area_proto_depIdxs = []int32{
	3, // 0: ntt.workplace.v1alpha.Area.geometry:type_name -> ntt.workplace.v1alpha.Geometry
	1, // 1: ntt.workplace.v1alpha.Area.vendor_spec:type_name -> ntt.workplace.v1alpha.Area.VendorSpec
	4, // 2: ntt.workplace.v1alpha.Area.metadata:type_name -> ntt.types.Meta
	2, // 3: ntt.workplace.v1alpha.Area.VendorSpec.point_grab:type_name -> ntt.workplace.v1alpha.Area.VendorSpec.PointGrab
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { workplace_proto_v1alpha_area_proto_init() }
func workplace_proto_v1alpha_area_proto_init() {
	if workplace_proto_v1alpha_area_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		workplace_proto_v1alpha_area_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
		workplace_proto_v1alpha_area_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area_VendorSpec); i {
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
		workplace_proto_v1alpha_area_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area_VendorSpec_PointGrab); i {
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
			RawDescriptor: workplace_proto_v1alpha_area_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           workplace_proto_v1alpha_area_proto_goTypes,
		DependencyIndexes: workplace_proto_v1alpha_area_proto_depIdxs,
		MessageInfos:      workplace_proto_v1alpha_area_proto_msgTypes,
	}.Build()
	workplace_proto_v1alpha_area_proto = out.File
	workplace_proto_v1alpha_area_proto_rawDesc = nil
	workplace_proto_v1alpha_area_proto_goTypes = nil
	workplace_proto_v1alpha_area_proto_depIdxs = nil
}