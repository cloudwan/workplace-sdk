// Code generated by protoc-gen-goten-go
// File: workplace/proto/v1alpha2/common.proto
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
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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
	_ = &latlng.LatLng{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vendor int32

const (
	Vendor_VENDOR_UNSPECIFIED Vendor = 0
	Vendor_EDGELQ_WORKPLACE   Vendor = 1
	Vendor_POINT_GRAB         Vendor = 2
	Vendor_RIPTIDE            Vendor = 3
	// used for floor plans. ids are tileset urls
	Vendor_MAPBOX Vendor = 4
)

var (
	Vendor_name = map[int32]string{
		0: "VENDOR_UNSPECIFIED",
		1: "EDGELQ_WORKPLACE",
		2: "POINT_GRAB",
		3: "RIPTIDE",
		4: "MAPBOX",
	}

	Vendor_value = map[string]int32{
		"VENDOR_UNSPECIFIED": 0,
		"EDGELQ_WORKPLACE":   1,
		"POINT_GRAB":         2,
		"RIPTIDE":            3,
		"MAPBOX":             4,
	}
)

func (x Vendor) Enum() *Vendor {
	p := new(Vendor)
	*p = x
	return p
}

func (x Vendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (Vendor) Descriptor() preflect.EnumDescriptor {
	return workplace_proto_v1alpha2_common_proto_enumTypes[0].Descriptor()
}

func (Vendor) Type() preflect.EnumType {
	return &workplace_proto_v1alpha2_common_proto_enumTypes[0]
}

func (x Vendor) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use Vendor.ProtoReflect.Descriptor instead.
func (Vendor) EnumDescriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{0}
}

// BBox descrbes rectangular geographical area by using two corners
type BBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// SouthWest-most corner
	SouthWest *latlng.LatLng `protobuf:"bytes,1,opt,name=south_west,json=southWest,proto3" json:"south_west,omitempty" firestore:"southWest"`
	// NorthEast-most corner
	NorthEast *latlng.LatLng `protobuf:"bytes,2,opt,name=north_east,json=northEast,proto3" json:"north_east,omitempty" firestore:"northEast"`
}

func (m *BBox) Reset() {
	*m = BBox{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BBox) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BBox) ProtoMessage() {}

func (m *BBox) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BBox) GotenMessage() {}

// Deprecated, Use BBox.ProtoReflect.Descriptor instead.
func (*BBox) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{0}
}

func (m *BBox) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BBox) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BBox) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BBox) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BBox) GetSouthWest() *latlng.LatLng {
	if m != nil {
		return m.SouthWest
	}
	return nil
}

func (m *BBox) GetNorthEast() *latlng.LatLng {
	if m != nil {
		return m.NorthEast
	}
	return nil
}

func (m *BBox) SetSouthWest(fv *latlng.LatLng) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SouthWest", "BBox"))
	}
	m.SouthWest = fv
}

func (m *BBox) SetNorthEast(fv *latlng.LatLng) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NorthEast", "BBox"))
	}
	m.NorthEast = fv
}

// Polygon describes polygon shape with using geo-coordinates
type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// List of coordinates. Each polygon vertice connects 2 consecutive points and
	// last point connects with first.
	Coordinates []*latlng.LatLng `protobuf:"bytes,1,rep,name=coordinates,proto3" json:"coordinates,omitempty" firestore:"coordinates"`
}

func (m *Polygon) Reset() {
	*m = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Polygon) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Polygon) ProtoMessage() {}

func (m *Polygon) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Polygon) GotenMessage() {}

// Deprecated, Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{1}
}

func (m *Polygon) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Polygon) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Polygon) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Polygon) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Polygon) GetCoordinates() []*latlng.LatLng {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Polygon) SetCoordinates(fv []*latlng.LatLng) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Coordinates", "Polygon"))
	}
	m.Coordinates = fv
}

// Geometry contains geo-information about place.
// It may contain just center, bounding box or polygon
type Geometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Center point of place
	Center *latlng.LatLng `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty" firestore:"center"`
	// Bounding Box of place (should contain polygon if present)
	Bbox *BBox `protobuf:"bytes,2,opt,name=bbox,proto3" json:"bbox,omitempty" firestore:"bbox"`
	// Polygon contains detailed shape of place
	Polygon *Polygon `protobuf:"bytes,3,opt,name=polygon,proto3" json:"polygon,omitempty" firestore:"polygon"`
	// Angle panning in degrees when displaying a centered map view.
	// Panning with value 0 results in default behavior (usually North).
	Panning float64 `protobuf:"fixed64,4,opt,name=panning,proto3" json:"panning,omitempty" firestore:"panning"`
}

func (m *Geometry) Reset() {
	*m = Geometry{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Geometry) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Geometry) ProtoMessage() {}

func (m *Geometry) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Geometry) GotenMessage() {}

// Deprecated, Use Geometry.ProtoReflect.Descriptor instead.
func (*Geometry) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{2}
}

func (m *Geometry) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Geometry) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Geometry) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Geometry) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Geometry) GetCenter() *latlng.LatLng {
	if m != nil {
		return m.Center
	}
	return nil
}

func (m *Geometry) GetBbox() *BBox {
	if m != nil {
		return m.Bbox
	}
	return nil
}

func (m *Geometry) GetPolygon() *Polygon {
	if m != nil {
		return m.Polygon
	}
	return nil
}

func (m *Geometry) GetPanning() float64 {
	if m != nil {
		return m.Panning
	}
	return float64(0)
}

func (m *Geometry) SetCenter(fv *latlng.LatLng) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Center", "Geometry"))
	}
	m.Center = fv
}

func (m *Geometry) SetBbox(fv *BBox) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Bbox", "Geometry"))
	}
	m.Bbox = fv
}

func (m *Geometry) SetPolygon(fv *Polygon) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Polygon", "Geometry"))
	}
	m.Polygon = fv
}

func (m *Geometry) SetPanning(fv float64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Panning", "Geometry"))
	}
	m.Panning = fv
}

// StreetLocation contains street address and resolved coordinates
type StreetLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Street address in english
	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty" firestore:"streetAddress"`
	// Resolved street address coordinates
	StreetCoordinates *latlng.LatLng `protobuf:"bytes,2,opt,name=street_coordinates,json=streetCoordinates,proto3" json:"street_coordinates,omitempty" firestore:"streetCoordinates"`
}

func (m *StreetLocation) Reset() {
	*m = StreetLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *StreetLocation) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*StreetLocation) ProtoMessage() {}

func (m *StreetLocation) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*StreetLocation) GotenMessage() {}

// Deprecated, Use StreetLocation.ProtoReflect.Descriptor instead.
func (*StreetLocation) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{3}
}

func (m *StreetLocation) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *StreetLocation) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *StreetLocation) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *StreetLocation) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *StreetLocation) GetStreetAddress() string {
	if m != nil {
		return m.StreetAddress
	}
	return ""
}

func (m *StreetLocation) GetStreetCoordinates() *latlng.LatLng {
	if m != nil {
		return m.StreetCoordinates
	}
	return nil
}

func (m *StreetLocation) SetStreetAddress(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "StreetAddress", "StreetLocation"))
	}
	m.StreetAddress = fv
}

func (m *StreetLocation) SetStreetCoordinates(fv *latlng.LatLng) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "StreetCoordinates", "StreetLocation"))
	}
	m.StreetCoordinates = fv
}

// VendorMapping contains entity id registered in another system
type VendorMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// vendor type
	Vendor Vendor `protobuf:"varint,1,opt,name=vendor,proto3,enum=ntt.workplace.v1alpha2.Vendor" json:"vendor,omitempty" firestore:"vendor"`
	// corresponding entity id in vendor system
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" firestore:"id"`
}

func (m *VendorMapping) Reset() {
	*m = VendorMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *VendorMapping) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*VendorMapping) ProtoMessage() {}

func (m *VendorMapping) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*VendorMapping) GotenMessage() {}

// Deprecated, Use VendorMapping.ProtoReflect.Descriptor instead.
func (*VendorMapping) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{4}
}

func (m *VendorMapping) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *VendorMapping) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *VendorMapping) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *VendorMapping) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *VendorMapping) GetVendor() Vendor {
	if m != nil {
		return m.Vendor
	}
	return Vendor_VENDOR_UNSPECIFIED
}

func (m *VendorMapping) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VendorMapping) SetVendor(fv Vendor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Vendor", "VendorMapping"))
	}
	m.Vendor = fv
}

func (m *VendorMapping) SetId(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Id", "VendorMapping"))
	}
	m.Id = fv
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	X             float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty" firestore:"x"`
	Y             float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty" firestore:"y"`
}

func (m *Point) Reset() {
	*m = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Point) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Point) ProtoMessage() {}

func (m *Point) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Point) GotenMessage() {}

// Deprecated, Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{5}
}

func (m *Point) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Point) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Point) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Point) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Point) GetX() float64 {
	if m != nil {
		return m.X
	}
	return float64(0)
}

func (m *Point) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return float64(0)
}

func (m *Point) SetX(fv float64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "X", "Point"))
	}
	m.X = fv
}

func (m *Point) SetY(fv float64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Y", "Point"))
	}
	m.Y = fv
}

type ReferencePoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// relative point
	Point *Point `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty" firestore:"point"`
	// absolute geo-coordinates
	LatLng *latlng.LatLng `protobuf:"bytes,2,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty" firestore:"latLng"`
}

func (m *ReferencePoint) Reset() {
	*m = ReferencePoint{}
	if protoimpl.UnsafeEnabled {
		mi := &workplace_proto_v1alpha2_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ReferencePoint) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ReferencePoint) ProtoMessage() {}

func (m *ReferencePoint) ProtoReflect() preflect.Message {
	mi := &workplace_proto_v1alpha2_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ReferencePoint) GotenMessage() {}

// Deprecated, Use ReferencePoint.ProtoReflect.Descriptor instead.
func (*ReferencePoint) Descriptor() ([]byte, []int) {
	return workplace_proto_v1alpha2_common_proto_rawDescGZIP(), []int{6}
}

func (m *ReferencePoint) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ReferencePoint) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ReferencePoint) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ReferencePoint) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ReferencePoint) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *ReferencePoint) GetLatLng() *latlng.LatLng {
	if m != nil {
		return m.LatLng
	}
	return nil
}

func (m *ReferencePoint) SetPoint(fv *Point) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Point", "ReferencePoint"))
	}
	m.Point = fv
}

func (m *ReferencePoint) SetLatLng(fv *latlng.LatLng) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LatLng", "ReferencePoint"))
	}
	m.LatLng = fv
}

var workplace_proto_v1alpha2_common_proto preflect.FileDescriptor

var workplace_proto_v1alpha2_common_proto_rawDesc = []byte{
	0x0a, 0x25, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74,
	0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x04, 0x42, 0x42, 0x6f, 0x78, 0x12, 0x32,
	0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x77, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x57, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x5f, 0x65, 0x61, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x09, 0x6e, 0x6f, 0x72,
	0x74, 0x68, 0x45, 0x61, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x0b, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x04,
	0x62, 0x62, 0x6f, 0x78, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x70, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x7b, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x42, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74,
	0x4c, 0x6e, 0x67, 0x52, 0x11, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x0d, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x79, 0x22, 0x73, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x6c,
	0x61, 0x74, 0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e,
	0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x2a, 0x5f, 0x0a, 0x06, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x44, 0x47, 0x45, 0x4c, 0x51, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x47, 0x52, 0x41, 0x42, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x49, 0x50, 0x54, 0x49, 0x44, 0x45, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x41, 0x50, 0x42, 0x4f, 0x58, 0x10, 0x04, 0x42, 0x7c, 0xe8, 0xde, 0x21, 0x01,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	workplace_proto_v1alpha2_common_proto_rawDescOnce sync.Once
	workplace_proto_v1alpha2_common_proto_rawDescData = workplace_proto_v1alpha2_common_proto_rawDesc
)

func workplace_proto_v1alpha2_common_proto_rawDescGZIP() []byte {
	workplace_proto_v1alpha2_common_proto_rawDescOnce.Do(func() {
		workplace_proto_v1alpha2_common_proto_rawDescData = protoimpl.X.CompressGZIP(workplace_proto_v1alpha2_common_proto_rawDescData)
	})
	return workplace_proto_v1alpha2_common_proto_rawDescData
}

var workplace_proto_v1alpha2_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var workplace_proto_v1alpha2_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var workplace_proto_v1alpha2_common_proto_goTypes = []interface{}{
	(Vendor)(0),            // 0: ntt.workplace.v1alpha2.Vendor
	(*BBox)(nil),           // 1: ntt.workplace.v1alpha2.BBox
	(*Polygon)(nil),        // 2: ntt.workplace.v1alpha2.Polygon
	(*Geometry)(nil),       // 3: ntt.workplace.v1alpha2.Geometry
	(*StreetLocation)(nil), // 4: ntt.workplace.v1alpha2.StreetLocation
	(*VendorMapping)(nil),  // 5: ntt.workplace.v1alpha2.VendorMapping
	(*Point)(nil),          // 6: ntt.workplace.v1alpha2.Point
	(*ReferencePoint)(nil), // 7: ntt.workplace.v1alpha2.ReferencePoint
	(*latlng.LatLng)(nil),  // 8: google.type.LatLng
}
var workplace_proto_v1alpha2_common_proto_depIdxs = []int32{
	8,  // 0: ntt.workplace.v1alpha2.BBox.south_west:type_name -> google.type.LatLng
	8,  // 1: ntt.workplace.v1alpha2.BBox.north_east:type_name -> google.type.LatLng
	8,  // 2: ntt.workplace.v1alpha2.Polygon.coordinates:type_name -> google.type.LatLng
	8,  // 3: ntt.workplace.v1alpha2.Geometry.center:type_name -> google.type.LatLng
	1,  // 4: ntt.workplace.v1alpha2.Geometry.bbox:type_name -> ntt.workplace.v1alpha2.BBox
	2,  // 5: ntt.workplace.v1alpha2.Geometry.polygon:type_name -> ntt.workplace.v1alpha2.Polygon
	8,  // 6: ntt.workplace.v1alpha2.StreetLocation.street_coordinates:type_name -> google.type.LatLng
	0,  // 7: ntt.workplace.v1alpha2.VendorMapping.vendor:type_name -> ntt.workplace.v1alpha2.Vendor
	6,  // 8: ntt.workplace.v1alpha2.ReferencePoint.point:type_name -> ntt.workplace.v1alpha2.Point
	8,  // 9: ntt.workplace.v1alpha2.ReferencePoint.lat_lng:type_name -> google.type.LatLng
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { workplace_proto_v1alpha2_common_proto_init() }
func workplace_proto_v1alpha2_common_proto_init() {
	if workplace_proto_v1alpha2_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		workplace_proto_v1alpha2_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBox); i {
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
		workplace_proto_v1alpha2_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Polygon); i {
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
		workplace_proto_v1alpha2_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geometry); i {
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
		workplace_proto_v1alpha2_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreetLocation); i {
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
		workplace_proto_v1alpha2_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VendorMapping); i {
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
		workplace_proto_v1alpha2_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		workplace_proto_v1alpha2_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferencePoint); i {
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
			RawDescriptor: workplace_proto_v1alpha2_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           workplace_proto_v1alpha2_common_proto_goTypes,
		DependencyIndexes: workplace_proto_v1alpha2_common_proto_depIdxs,
		EnumInfos:         workplace_proto_v1alpha2_common_proto_enumTypes,
		MessageInfos:      workplace_proto_v1alpha2_common_proto_msgTypes,
	}.Build()
	workplace_proto_v1alpha2_common_proto = out.File
	workplace_proto_v1alpha2_common_proto_rawDesc = nil
	workplace_proto_v1alpha2_common_proto_goTypes = nil
	workplace_proto_v1alpha2_common_proto_depIdxs = nil
}
