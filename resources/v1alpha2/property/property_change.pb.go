// Code generated by protoc-gen-goten-go
// File: workplace/proto/v1alpha2/property_change.proto
// DO NOT EDIT!!!

package property

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
	device "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device"
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	zone "github.com/cloudwan/workplace-sdk/resources/v1alpha2/zone"
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
	_ = &device.Device{}
	_ = &floor.Floor{}
	_ = &site.Site{}
	_ = &zone.Zone{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PropertyChange is used by Watch notifications Responses to describe change of
// single Property One of Added, Modified, Removed
type PropertyChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Property change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*PropertyChange_Added_
	//	*PropertyChange_Modified_
	//	*PropertyChange_Current_
	//	*PropertyChange_Removed_
	ChangeType isPropertyChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *PropertyChange) Reset() {
	*m = PropertyChange{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PropertyChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PropertyChange) ProtoMessage() {}

func (m *PropertyChange) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PropertyChange) GotenMessage() {}

// Deprecated, Use PropertyChange.ProtoReflect.Descriptor instead.
func (*PropertyChange) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_property_change_proto_rawDescGZIP(), []int{0}
}

func (m *PropertyChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PropertyChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PropertyChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PropertyChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isPropertyChange_ChangeType interface {
	isPropertyChange_ChangeType()
}

type PropertyChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *PropertyChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type PropertyChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *PropertyChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type PropertyChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *PropertyChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type PropertyChange_Removed_ struct {
	// Removed is returned when Property is deleted or leaves Query view
	Removed *PropertyChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*PropertyChange_Added_) isPropertyChange_ChangeType()    {}
func (*PropertyChange_Modified_) isPropertyChange_ChangeType() {}
func (*PropertyChange_Current_) isPropertyChange_ChangeType()  {}
func (*PropertyChange_Removed_) isPropertyChange_ChangeType()  {}
func (m *PropertyChange) GetChangeType() isPropertyChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *PropertyChange) GetAdded() *PropertyChange_Added {
	if x, ok := m.GetChangeType().(*PropertyChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *PropertyChange) GetModified() *PropertyChange_Modified {
	if x, ok := m.GetChangeType().(*PropertyChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *PropertyChange) GetCurrent() *PropertyChange_Current {
	if x, ok := m.GetChangeType().(*PropertyChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *PropertyChange) GetRemoved() *PropertyChange_Removed {
	if x, ok := m.GetChangeType().(*PropertyChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *PropertyChange) SetChangeType(ofv isPropertyChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isPropertyChange_ChangeType", "PropertyChange"))
	}
	m.ChangeType = ofv
}
func (m *PropertyChange) SetAdded(fv *PropertyChange_Added) {
	m.SetChangeType(&PropertyChange_Added_{Added: fv})
}
func (m *PropertyChange) SetModified(fv *PropertyChange_Modified) {
	m.SetChangeType(&PropertyChange_Modified_{Modified: fv})
}
func (m *PropertyChange) SetCurrent(fv *PropertyChange_Current) {
	m.SetChangeType(&PropertyChange_Current_{Current: fv})
}
func (m *PropertyChange) SetRemoved(fv *PropertyChange_Removed) {
	m.SetChangeType(&PropertyChange_Removed_{Removed: fv})
}

// Property has been added to query view
type PropertyChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Property      *Property `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty" firestore:"property"`
	// Integer describing index of added Property in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PropertyChange_Added) Reset() {
	*m = PropertyChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PropertyChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PropertyChange_Added) ProtoMessage() {}

func (m *PropertyChange_Added) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PropertyChange_Added) GotenMessage() {}

// Deprecated, Use PropertyChange_Added.ProtoReflect.Descriptor instead.
func (*PropertyChange_Added) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_property_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *PropertyChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PropertyChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PropertyChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PropertyChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PropertyChange_Added) GetProperty() *Property {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *PropertyChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PropertyChange_Added) SetProperty(fv *Property) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Property", "PropertyChange_Added"))
	}
	m.Property = fv
}

func (m *PropertyChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PropertyChange_Added"))
	}
	m.ViewIndex = fv
}

// Property changed some of it's fields - contains either full document or
// masked change
type PropertyChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Property
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Property or masked difference, depending on mask_changes
	// instrumentation of issued [WatchPropertyRequest] or
	// [WatchPropertiesRequest]
	Property *Property `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty" firestore:"property"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Property_FieldMask `protobuf:"bytes,3,opt,customtype=Property_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Property.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Property new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PropertyChange_Modified) Reset() {
	*m = PropertyChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PropertyChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PropertyChange_Modified) ProtoMessage() {}

func (m *PropertyChange_Modified) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PropertyChange_Modified) GotenMessage() {}

// Deprecated, Use PropertyChange_Modified.ProtoReflect.Descriptor instead.
func (*PropertyChange_Modified) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_property_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *PropertyChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PropertyChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PropertyChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PropertyChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PropertyChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PropertyChange_Modified) GetProperty() *Property {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *PropertyChange_Modified) GetFieldMask() *Property_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *PropertyChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *PropertyChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PropertyChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PropertyChange_Modified"))
	}
	m.Name = fv
}

func (m *PropertyChange_Modified) SetProperty(fv *Property) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Property", "PropertyChange_Modified"))
	}
	m.Property = fv
}

func (m *PropertyChange_Modified) SetFieldMask(fv *Property_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "PropertyChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *PropertyChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "PropertyChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *PropertyChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PropertyChange_Modified"))
	}
	m.ViewIndex = fv
}

// Property has been added or modified in a query view. Version used for
// stateless watching
type PropertyChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Property      *Property `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty" firestore:"property"`
}

func (m *PropertyChange_Current) Reset() {
	*m = PropertyChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PropertyChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PropertyChange_Current) ProtoMessage() {}

func (m *PropertyChange_Current) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PropertyChange_Current) GotenMessage() {}

// Deprecated, Use PropertyChange_Current.ProtoReflect.Descriptor instead.
func (*PropertyChange_Current) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_property_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *PropertyChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PropertyChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PropertyChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PropertyChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PropertyChange_Current) GetProperty() *Property {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *PropertyChange_Current) SetProperty(fv *Property) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Property", "PropertyChange_Current"))
	}
	m.Property = fv
}

// Removed is returned when Property is deleted or leaves Query view
type PropertyChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Property index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PropertyChange_Removed) Reset() {
	*m = PropertyChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PropertyChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PropertyChange_Removed) ProtoMessage() {}

func (m *PropertyChange_Removed) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_property_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PropertyChange_Removed) GotenMessage() {}

// Deprecated, Use PropertyChange_Removed.ProtoReflect.Descriptor instead.
func (*PropertyChange_Removed) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_property_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *PropertyChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PropertyChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PropertyChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PropertyChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PropertyChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PropertyChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PropertyChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PropertyChange_Removed"))
	}
	m.Name = fv
}

func (m *PropertyChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PropertyChange_Removed"))
	}
	m.ViewIndex = fv
}

var workplace_proto_v1alpha2_property_change_proto preflect.FileDescriptor

var workplace_proto_v1alpha2_property_change_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x06, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a,
	0x64, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x8a, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x4b, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x32, 0x0a, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x1a, 0x47, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x4e, 0x0a, 0x07, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0e, 0x9a, 0xd9, 0x21,
	0x0a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x7e, 0xe8, 0xde, 0x21, 0x00,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x13, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	workplace_proto_v1alpha2_property_change_proto_rawDescOnce sync.Once
	workplace_proto_v1alpha2_property_change_proto_rawDescData = workplace_proto_v1alpha2_property_change_proto_rawDesc
)

func workplace_proto_v1alpha2_property_change_proto_rawDescGZIP() []byte {
	workplace_proto_v1alpha2_property_change_proto_rawDescOnce.Do(func() {
		workplace_proto_v1alpha2_property_change_proto_rawDescData = protoimpl.X.CompressGZIP(workplace_proto_v1alpha2_property_change_proto_rawDescData)
	})
	return workplace_proto_v1alpha2_property_change_proto_rawDescData
}

var workplace_proto_v1alpha2_property_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var workplace_proto_v1alpha2_property_change_proto_goTypes = []interface{}{
	(*PropertyChange)(nil),          // 0: ntt.workplace.v1alpha2.PropertyChange
	(*PropertyChange_Added)(nil),    // 1: ntt.workplace.v1alpha2.PropertyChange.Added
	(*PropertyChange_Modified)(nil), // 2: ntt.workplace.v1alpha2.PropertyChange.Modified
	(*PropertyChange_Current)(nil),  // 3: ntt.workplace.v1alpha2.PropertyChange.Current
	(*PropertyChange_Removed)(nil),  // 4: ntt.workplace.v1alpha2.PropertyChange.Removed
	(*Property)(nil),                // 5: ntt.workplace.v1alpha2.Property
	(*Property_FieldMask)(nil),      // 6: ntt.workplace.v1alpha2.Property_FieldMask
}
var workplace_proto_v1alpha2_property_change_proto_depIdxs = []int32{
	1, // 0: ntt.workplace.v1alpha2.PropertyChange.added:type_name -> ntt.workplace.v1alpha2.PropertyChange.Added
	2, // 1: ntt.workplace.v1alpha2.PropertyChange.modified:type_name -> ntt.workplace.v1alpha2.PropertyChange.Modified
	3, // 2: ntt.workplace.v1alpha2.PropertyChange.current:type_name -> ntt.workplace.v1alpha2.PropertyChange.Current
	4, // 3: ntt.workplace.v1alpha2.PropertyChange.removed:type_name -> ntt.workplace.v1alpha2.PropertyChange.Removed
	5, // 4: ntt.workplace.v1alpha2.PropertyChange.Added.property:type_name -> ntt.workplace.v1alpha2.Property
	5, // 5: ntt.workplace.v1alpha2.PropertyChange.Modified.property:type_name -> ntt.workplace.v1alpha2.Property
	6, // 6: ntt.workplace.v1alpha2.PropertyChange.Modified.field_mask:type_name -> ntt.workplace.v1alpha2.Property_FieldMask
	5, // 7: ntt.workplace.v1alpha2.PropertyChange.Current.property:type_name -> ntt.workplace.v1alpha2.Property
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { workplace_proto_v1alpha2_property_change_proto_init() }
func workplace_proto_v1alpha2_property_change_proto_init() {
	if workplace_proto_v1alpha2_property_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		workplace_proto_v1alpha2_property_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyChange); i {
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
		workplace_proto_v1alpha2_property_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyChange_Added); i {
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
		workplace_proto_v1alpha2_property_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyChange_Modified); i {
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
		workplace_proto_v1alpha2_property_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyChange_Current); i {
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
		workplace_proto_v1alpha2_property_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyChange_Removed); i {
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

	workplace_proto_v1alpha2_property_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PropertyChange_Added_)(nil),
		(*PropertyChange_Modified_)(nil),
		(*PropertyChange_Current_)(nil),
		(*PropertyChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: workplace_proto_v1alpha2_property_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           workplace_proto_v1alpha2_property_change_proto_goTypes,
		DependencyIndexes: workplace_proto_v1alpha2_property_change_proto_depIdxs,
		MessageInfos:      workplace_proto_v1alpha2_property_change_proto_msgTypes,
	}.Build()
	workplace_proto_v1alpha2_property_change_proto = out.File
	workplace_proto_v1alpha2_property_change_proto_rawDesc = nil
	workplace_proto_v1alpha2_property_change_proto_goTypes = nil
	workplace_proto_v1alpha2_property_change_proto_depIdxs = nil
}