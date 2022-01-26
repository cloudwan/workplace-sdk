// Code generated by protoc-gen-goten-resource
// Resource: Floor
// DO NOT EDIT!!!

package floor

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/goten"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
)

// ensure the imports are used
var (
	_ = codes.NotFound
	_ = fmt.Stringer(nil)
	_ = proto.Message(nil)
	_ = status.Status{}
	_ = url.URL{}
	_ = strings.Builder{}

	_ = goten.GotenMessage(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
)

var floor_RegexpId = regexp.MustCompile("^(?P<floor_id>[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})$")

func (r *Floor) MaybePopulateDefaults() error {
	floorInterface := interface{}(r)
	if defaulter, ok := floorInterface.(goten.Defaulter); ok {
		return defaulter.PopulateDefaults()
	}
	return nil
}

func (r *Floor) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *Floor) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

type Name struct {
	ParentName
	FloorId string `firestore:"floorId"`
}

func ParseName(name string) (*Name, error) {
	var matches []string
	if matches = regexPath_Project_Region_Site_Building.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetId(matches[5]).
			Name(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as Floor name", name)
}

func MustParseName(name string) *Name {
	result, err := ParseName(name)
	if err != nil {
		panic(err)
	}
	return result
}

func ParseNameOrId(nameOrId string) (*Name, error) {
	name, err := ParseName(nameOrId)
	if err == nil {
		return name, err
	}
	if floor_RegexpId.MatchString(nameOrId) {
		return NewNameBuilder().SetId(nameOrId).Name(), nil
	} else {
		return nil, fmt.Errorf("unable to parse '%s' as Floor name or id", name)
	}
}

func (name *Name) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 0 {
		return status.Errorf(codes.InvalidArgument, "No segments given for Floor name")
	}
	if err := name.ParentName.SetFromSegments(segments[:len(segments)-1]); err != nil {
		return err
	}
	if segments[len(segments)-1].CollectionLowerJson != "floors" {
		return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form Floor name", segments)
	}
	name.FloorId = segments[len(segments)-1].Id
	return nil
}

func (name *Name) GetBuildingName() *building.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetBuildingName()
}

func (name *Name) IsSpecified() bool {
	return name != nil && (name.Pattern == NamePattern_Project_Region_Site_Building)
}

func (name *Name) IsFullyQualified() bool {
	if name == nil {
		return false
	}
	if name.ParentName.IsFullyQualified() == false {
		return false
	}
	if name.FloorId == "" || name.FloorId == gotenresource.WildcardId {
		return false
	}
	return true
}

func (name *Name) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for Floor is not fully qualified")
	}
	return fmt.Sprintf("//workplace.edgelq.com/%s", name.String()), nil
}

func (name *Name) String() string {
	if name == nil {
		return "<nil>"
	}
	if valueStr, err := name.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

func (name *Name) AsReference() *Reference {
	return &Reference{Name: *name}
}

func (name *Name) AsRawReference() gotenresource.Reference {
	return name.AsReference()
}

func (name *Name) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (name *Name) GetPattern() gotenresource.NamePattern {
	if name == nil {
		return ""
	}
	return name.Pattern
}

func (name *Name) GetIdParts() map[string]string {
	if name != nil {
		return map[string]string{
			"projectId":  name.ProjectId,
			"regionId":   name.RegionId,
			"siteId":     name.SiteId,
			"buildingId": name.BuildingId,
			"floorId":    name.FloorId,
		}
	}
	return map[string]string{
		"projectId":  "",
		"regionId":   "",
		"siteId":     "",
		"buildingId": "",
		"floorId":    "",
	}
}

func (name *Name) GetSegments() gotenresource.NameSegments {
	if name == nil || name.Pattern == "" {
		return nil
	}
	segments := name.ParentName.GetSegments()
	return append(segments, gotenresource.NameSegment{
		CollectionLowerJson: "floors",
		Id:                  name.FloorId,
	})
}

// implement methods required by protobuf-go library for string-struct conversion

func (name *Name) ProtoString() (string, error) {
	if name == nil {
		return "", nil
	}
	result := ""
	parentPrefix, err := name.ParentName.ProtoString()
	if err != nil {
		return "", err
	}
	if parentPrefix != "" {
		result += parentPrefix + "/"
	}
	result += "floors/" + name.FloorId
	return result, nil
}

func (name *Name) ParseProtoString(data string) error {
	parsed, err := ParseName(data)
	if err != nil {
		return err
	}
	*name = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (name *Name) GotenEqual(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*Name)
	if !ok {
		other2, ok := other.(Name)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	if name.ParentName.GotenEqual(other1.ParentName) == false {
		return false
	}
	if name.FloorId != other1.FloorId {
		return false
	}

	return true
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *Name) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*Name)
	if !ok {
		other2, ok := other.(Name)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	if name.ParentName.Matches(other1.ParentName) == false {
		return false
	}
	if name.FloorId != other1.FloorId {
		return name.FloorId == gotenresource.WildcardId
	}

	return true
}

// Google CEL integration Type
var celName = types.NewTypeValue("Name", traits.ReceiverType)

func (name *Name) TypeName() string {
	return ".ntt.workplace.v1alpha2.Floor.Name"
}

func (name *Name) HasTrait(trait int) bool {
	return trait == traits.ReceiverType
}

func (name *Name) Equal(other ref.Val) ref.Val {
	return types.Bool(false)
}

func (name *Name) Value() interface{} {
	return name
}

func (name *Name) Match(pattern ref.Val) ref.Val {
	return types.Bool(true)
}

func (name *Name) Receive(function string, overload string, args []ref.Val) ref.Val {
	switch function {
	case "getId":
		return types.String(name.FloorId)
	case "satisfies":
		rhsName, err := ParseName(string(args[0].(types.String)))
		if err != nil {
			return types.ValOrErr(celName, "function %s name parse error: %s", function, err)
		}
		return types.Bool(rhsName.Matches(name))
	default:
		return types.ValOrErr(celName, "no such function - %s", function)
	}
}

func (name *Name) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	panic("not required")
}

func (name *Name) ConvertToType(typeVal ref.Type) ref.Val {
	switch typeVal.TypeName() {
	case types.StringType.TypeName():
		return types.String(name.String())
	default:
		panic(fmt.Errorf("unable to convert %s to CEL type %s", "Name", typeVal.TypeName()))
	}
}

func (name *Name) Type() ref.Type {
	return celName
}

// implement CustomTypeCliValue method
func (name *Name) SetFromCliFlag(raw string) error {
	parsedName, err := ParseName(raw)
	if err != nil {
		return err
	}
	*name = *parsedName
	return nil
}

type Reference struct {
	Name
	floor *Floor
}

func MakeReference(name *Name, floor *Floor) (*Reference, error) {
	return &Reference{
		Name:  *name,
		floor: floor,
	}, nil
}

func ParseReference(name string) (*Reference, error) {
	parsedName, err := ParseName(name)
	if err != nil {
		return nil, err
	}
	return MakeReference(parsedName, nil)
}

func MustParseReference(name string) *Reference {
	result, err := ParseReference(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (ref *Reference) Resolve(resolved *Floor) {
	ref.floor = resolved
}

func (ref *Reference) ResolveRaw(res gotenresource.Resource) error {
	if typedRes, ok := res.(*Floor); ok && typedRes != nil {
		ref.Resolve(typedRes)
		return nil
	}
	return status.Errorf(codes.Internal, "Invalid resource type for Floor: %s", reflect.TypeOf(res).Elem().Name())
}

func (ref *Reference) Resolved() bool {
	return ref.floor != nil
}

func (ref *Reference) ClearCached() {
	ref.floor = nil
}

func (ref *Reference) GetFloor() *Floor {
	return ref.floor
}

func (ref *Reference) GetRawResource() gotenresource.Resource {
	return ref.floor
}

func (ref *Reference) IsFullyQualified() bool {
	if ref == nil {
		return false
	}
	return ref.Name.IsFullyQualified()
}

func (ref *Reference) IsSpecified() bool {
	if ref == nil {
		return false
	}
	return ref.Name.IsSpecified()
}

func (ref *Reference) FullyQualifiedName() (string, error) {
	if !ref.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for Floor is not fully qualified")
	}
	return fmt.Sprintf("//workplace.edgelq.com/%s", ref.String()), nil
}

func (ref *Reference) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (ref *Reference) GetPattern() gotenresource.NamePattern {
	if ref == nil {
		return ""
	}
	return ref.Pattern
}

func (ref *Reference) GetIdParts() map[string]string {
	if ref != nil {
		return ref.Name.GetIdParts()
	}
	return map[string]string{
		"projectId":  "",
		"regionId":   "",
		"siteId":     "",
		"buildingId": "",
		"floorId":    "",
	}
}

func (ref *Reference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.Name.GetSegments()
	}
	return nil
}

func (ref *Reference) String() string {
	if ref == nil {
		return "<nil>"
	}
	return ref.Name.String()
}

// implement methods required by protobuf-go library for string-struct conversion

func (ref *Reference) ProtoString() (string, error) {
	if ref == nil {
		return "", nil
	}
	return ref.Name.ProtoString()
}

func (ref *Reference) ParseProtoString(data string) error {
	parsed, err := ParseReference(data)
	if err != nil {
		return err
	}
	*ref = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (ref *Reference) GotenEqual(other interface{}) bool {
	if other == nil {
		return ref == nil
	}
	other1, ok := other.(*Reference)
	if !ok {
		other2, ok := other.(Reference)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return ref == nil
	} else if ref == nil {
		return false
	}

	return ref.Name.GotenEqual(other1.Name)
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *Reference) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*Reference)
	if !ok {
		other2, ok := other.(Reference)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	return name.Name.Matches(&other1.Name)
}

// implement CustomTypeCliValue method
func (ref *Reference) SetFromCliFlag(raw string) error {
	parsedRef, err := ParseReference(raw)
	if err != nil {
		return err
	}
	*ref = *parsedRef
	return nil
}
