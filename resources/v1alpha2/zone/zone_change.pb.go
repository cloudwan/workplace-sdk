// Code generated by protoc-gen-goten-go
// File: workplace/proto/v1alpha2/zone_change.proto
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
	area "github.com/cloudwan/workplace-sdk/resources/v1alpha2/area"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &field_mask.FieldMask{}
	_ = &area.Area{}
	_ = &building.Building{}
	_ = &floor.Floor{}
	_ = &site.Site{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ZoneChange is used by Watch notifications Responses to describe change of
// single Zone One of Added, Modified, Removed
type ZoneChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Zone change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*ZoneChange_Added_
	//	*ZoneChange_Modified_
	//	*ZoneChange_Current_
	//	*ZoneChange_Removed_
	ChangeType isZoneChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *ZoneChange) Reset() {
	*m = ZoneChange{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ZoneChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ZoneChange) ProtoMessage() {}

func (m *ZoneChange) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ZoneChange) GotenMessage() {}

// Deprecated, Use ZoneChange.ProtoReflect.Descriptor instead.
func (*ZoneChange) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_change_proto_rawDescGZIP(), []int{0}
}

func (m *ZoneChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ZoneChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ZoneChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ZoneChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isZoneChange_ChangeType interface {
	isZoneChange_ChangeType()
}

type ZoneChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *ZoneChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type ZoneChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *ZoneChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type ZoneChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *ZoneChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type ZoneChange_Removed_ struct {
	// Removed is returned when Zone is deleted or leaves Query view
	Removed *ZoneChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*ZoneChange_Added_) isZoneChange_ChangeType()    {}
func (*ZoneChange_Modified_) isZoneChange_ChangeType() {}
func (*ZoneChange_Current_) isZoneChange_ChangeType()  {}
func (*ZoneChange_Removed_) isZoneChange_ChangeType()  {}
func (m *ZoneChange) GetChangeType() isZoneChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *ZoneChange) GetAdded() *ZoneChange_Added {
	if x, ok := m.GetChangeType().(*ZoneChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *ZoneChange) GetModified() *ZoneChange_Modified {
	if x, ok := m.GetChangeType().(*ZoneChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *ZoneChange) GetCurrent() *ZoneChange_Current {
	if x, ok := m.GetChangeType().(*ZoneChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *ZoneChange) GetRemoved() *ZoneChange_Removed {
	if x, ok := m.GetChangeType().(*ZoneChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *ZoneChange) SetChangeType(ofv isZoneChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isZoneChange_ChangeType", "ZoneChange"))
	}
	m.ChangeType = ofv
}
func (m *ZoneChange) SetAdded(fv *ZoneChange_Added) {
	m.SetChangeType(&ZoneChange_Added_{Added: fv})
}
func (m *ZoneChange) SetModified(fv *ZoneChange_Modified) {
	m.SetChangeType(&ZoneChange_Modified_{Modified: fv})
}
func (m *ZoneChange) SetCurrent(fv *ZoneChange_Current) {
	m.SetChangeType(&ZoneChange_Current_{Current: fv})
}
func (m *ZoneChange) SetRemoved(fv *ZoneChange_Removed) {
	m.SetChangeType(&ZoneChange_Removed_{Removed: fv})
}

// Zone has been added to query view
type ZoneChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Zone          *Zone `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty" firestore:"zone"`
	// Integer describing index of added Zone in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ZoneChange_Added) Reset() {
	*m = ZoneChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ZoneChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ZoneChange_Added) ProtoMessage() {}

func (m *ZoneChange_Added) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ZoneChange_Added) GotenMessage() {}

// Deprecated, Use ZoneChange_Added.ProtoReflect.Descriptor instead.
func (*ZoneChange_Added) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ZoneChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ZoneChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ZoneChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ZoneChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ZoneChange_Added) GetZone() *Zone {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *ZoneChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ZoneChange_Added) SetZone(fv *Zone) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Zone", "ZoneChange_Added"))
	}
	m.Zone = fv
}

func (m *ZoneChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ZoneChange_Added"))
	}
	m.ViewIndex = fv
}

// Zone changed some of it's fields - contains either full document or masked
// change
type ZoneChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Zone
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Zone or masked difference, depending on mask_changes
	// instrumentation of issued [WatchZoneRequest] or [WatchZonesRequest]
	Zone *Zone `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty" firestore:"zone"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Zone_FieldMask `protobuf:"bytes,3,opt,customtype=Zone_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Zone.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Zone new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ZoneChange_Modified) Reset() {
	*m = ZoneChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ZoneChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ZoneChange_Modified) ProtoMessage() {}

func (m *ZoneChange_Modified) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ZoneChange_Modified) GotenMessage() {}

// Deprecated, Use ZoneChange_Modified.ProtoReflect.Descriptor instead.
func (*ZoneChange_Modified) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ZoneChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ZoneChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ZoneChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ZoneChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ZoneChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ZoneChange_Modified) GetZone() *Zone {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *ZoneChange_Modified) GetFieldMask() *Zone_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ZoneChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *ZoneChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ZoneChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ZoneChange_Modified"))
	}
	m.Name = fv
}

func (m *ZoneChange_Modified) SetZone(fv *Zone) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Zone", "ZoneChange_Modified"))
	}
	m.Zone = fv
}

func (m *ZoneChange_Modified) SetFieldMask(fv *Zone_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ZoneChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *ZoneChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "ZoneChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *ZoneChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ZoneChange_Modified"))
	}
	m.ViewIndex = fv
}

// Zone has been added or modified in a query view. Version used for stateless
// watching
type ZoneChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Zone          *Zone `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty" firestore:"zone"`
}

func (m *ZoneChange_Current) Reset() {
	*m = ZoneChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ZoneChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ZoneChange_Current) ProtoMessage() {}

func (m *ZoneChange_Current) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ZoneChange_Current) GotenMessage() {}

// Deprecated, Use ZoneChange_Current.ProtoReflect.Descriptor instead.
func (*ZoneChange_Current) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *ZoneChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ZoneChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ZoneChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ZoneChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ZoneChange_Current) GetZone() *Zone {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *ZoneChange_Current) SetZone(fv *Zone) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Zone", "ZoneChange_Current"))
	}
	m.Zone = fv
}

// Removed is returned when Zone is deleted or leaves Query view
type ZoneChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Zone index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ZoneChange_Removed) Reset() {
	*m = ZoneChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ZoneChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ZoneChange_Removed) ProtoMessage() {}

func (m *ZoneChange_Removed) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_zone_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ZoneChange_Removed) GotenMessage() {}

// Deprecated, Use ZoneChange_Removed.ProtoReflect.Descriptor instead.
func (*ZoneChange_Removed) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_zone_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *ZoneChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ZoneChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ZoneChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ZoneChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ZoneChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ZoneChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ZoneChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ZoneChange_Removed"))
	}
	m.Name = fv
}

func (m *ZoneChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ZoneChange_Removed"))
	}
	m.ViewIndex = fv
}

var workplace_proto_v1alpha2_zone_change_proto preflect.FileDescriptor

var workplace_proto_v1alpha2_zone_change_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x06, 0x0a, 0x0a, 0x5a, 0x6f, 0x6e, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x5a,
	0x6f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x1a, 0x58, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xf6, 0x01,
	0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06,
	0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x47,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0c,
	0xb2, 0xda, 0x21, 0x08, 0x32, 0x06, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69,
	0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x3b, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x1a, 0x4a, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x20,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda,
	0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a,
	0x0a, 0x9a, 0xd9, 0x21, 0x06, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x72, 0xe8, 0xde, 0x21, 0x00,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x0f, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x3b, 0x7a, 0x6f, 0x6e, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	workplace_proto_v1alpha2_zone_change_proto_rawDescOnce sync.Once
	workplace_proto_v1alpha2_zone_change_proto_rawDescData = workplace_proto_v1alpha2_zone_change_proto_rawDesc
)

func workplace_proto_v1alpha2_zone_change_proto_rawDescGZIP() []byte {
	workplace_proto_v1alpha2_zone_change_proto_rawDescOnce.Do(func() {
		workplace_proto_v1alpha2_zone_change_proto_rawDescData = protoimpl.X.CompressGZIP(workplace_proto_v1alpha2_zone_change_proto_rawDescData)
	})
	return workplace_proto_v1alpha2_zone_change_proto_rawDescData
}

var workplace_proto_v1alpha2_zone_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var workplace_proto_v1alpha2_zone_change_proto_goTypes = []interface{}{
	(*ZoneChange)(nil),          // 0: ntt.workplace.v1alpha2.ZoneChange
	(*ZoneChange_Added)(nil),    // 1: ntt.workplace.v1alpha2.ZoneChange.Added
	(*ZoneChange_Modified)(nil), // 2: ntt.workplace.v1alpha2.ZoneChange.Modified
	(*ZoneChange_Current)(nil),  // 3: ntt.workplace.v1alpha2.ZoneChange.Current
	(*ZoneChange_Removed)(nil),  // 4: ntt.workplace.v1alpha2.ZoneChange.Removed
	(*Zone)(nil),                // 5: ntt.workplace.v1alpha2.Zone
	(*Zone_FieldMask)(nil),      // 6: ntt.workplace.v1alpha2.Zone_FieldMask
}
var workplace_proto_v1alpha2_zone_change_proto_depIdxs = []int32{
	1, // 0: ntt.workplace.v1alpha2.ZoneChange.added:type_name -> ntt.workplace.v1alpha2.ZoneChange.Added
	2, // 1: ntt.workplace.v1alpha2.ZoneChange.modified:type_name -> ntt.workplace.v1alpha2.ZoneChange.Modified
	3, // 2: ntt.workplace.v1alpha2.ZoneChange.current:type_name -> ntt.workplace.v1alpha2.ZoneChange.Current
	4, // 3: ntt.workplace.v1alpha2.ZoneChange.removed:type_name -> ntt.workplace.v1alpha2.ZoneChange.Removed
	5, // 4: ntt.workplace.v1alpha2.ZoneChange.Added.zone:type_name -> ntt.workplace.v1alpha2.Zone
	5, // 5: ntt.workplace.v1alpha2.ZoneChange.Modified.zone:type_name -> ntt.workplace.v1alpha2.Zone
	6, // 6: ntt.workplace.v1alpha2.ZoneChange.Modified.field_mask:type_name -> ntt.workplace.v1alpha2.Zone_FieldMask
	5, // 7: ntt.workplace.v1alpha2.ZoneChange.Current.zone:type_name -> ntt.workplace.v1alpha2.Zone
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { workplace_proto_v1alpha2_zone_change_proto_init() }
func workplace_proto_v1alpha2_zone_change_proto_init() {
	if workplace_proto_v1alpha2_zone_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		workplace_proto_v1alpha2_zone_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneChange); i {
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
		workplace_proto_v1alpha2_zone_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneChange_Added); i {
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
		workplace_proto_v1alpha2_zone_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneChange_Modified); i {
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
		workplace_proto_v1alpha2_zone_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneChange_Current); i {
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
		workplace_proto_v1alpha2_zone_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneChange_Removed); i {
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

	workplace_proto_v1alpha2_zone_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ZoneChange_Added_)(nil),
		(*ZoneChange_Modified_)(nil),
		(*ZoneChange_Current_)(nil),
		(*ZoneChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: workplace_proto_v1alpha2_zone_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           workplace_proto_v1alpha2_zone_change_proto_goTypes,
		DependencyIndexes: workplace_proto_v1alpha2_zone_change_proto_depIdxs,
		MessageInfos:      workplace_proto_v1alpha2_zone_change_proto_msgTypes,
	}.Build()
	workplace_proto_v1alpha2_zone_change_proto = out.File
	workplace_proto_v1alpha2_zone_change_proto_rawDesc = nil
	workplace_proto_v1alpha2_zone_change_proto_goTypes = nil
	workplace_proto_v1alpha2_zone_change_proto_depIdxs = nil
}
