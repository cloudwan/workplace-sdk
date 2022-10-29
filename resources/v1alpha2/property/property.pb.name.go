// Code generated by protoc-gen-goten-resource
// Resource: Property
// DO NOT EDIT!!!

package property

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
	_ = codes.NotFound
	_ = new(fmt.Stringer)
	_ = new(proto.Message)
	_ = status.Status{}
	_ = url.URL{}
	_ = strings.Builder{}

	_ = new(goten.GotenMessage)
	_ = new(gotenresource.ListQuery)
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

var property_RegexpId = regexp.MustCompile("^(?P<property_id>[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building_Floor = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building_Floor_Area = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})/areas/(?P<area_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Zone = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/zones/(?P<zone_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building_Zone = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/zones/(?P<zone_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building_Floor_Zone = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})/zones/(?P<zone_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Building_Floor_Area_Zone = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})/areas/(?P<area_id>-|[\\w][\\w.-]{0,127})/zones/(?P<zone_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")
var regexPath_Project_Region_Site_Device = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/devices/(?P<device_id>-|[\\w][\\w.-]{0,127})/properties/(?P<property_id>-|[\\w][\\w.-]{0,127})$")

func (r *Property) MaybePopulateDefaults() error {
	propertyInterface := interface{}(r)
	if defaulter, ok := propertyInterface.(goten.Defaulter); ok {
		return defaulter.PopulateDefaults()
	}
	return nil
}

func (r *Property) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *Property) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

type Name struct {
	ParentName
	PropertyId string `firestore:"propertyId"`
}

func ParseName(name string) (*Name, error) {
	var matches []string
	if matches = regexPath_Project_Region_Site.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetId(matches[4]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Building.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetId(matches[5]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Building_Floor.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetFloorId(matches[5]).
			SetId(matches[6]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Building_Floor_Area.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetFloorId(matches[5]).
			SetAreaId(matches[6]).
			SetId(matches[7]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Zone.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetZoneId(matches[4]).
			SetId(matches[5]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Building_Zone.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetZoneId(matches[5]).
			SetId(matches[6]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Building_Floor_Zone.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetFloorId(matches[5]).
			SetZoneId(matches[6]).
			SetId(matches[7]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Building_Floor_Area_Zone.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetFloorId(matches[5]).
			SetAreaId(matches[6]).
			SetZoneId(matches[7]).
			SetId(matches[8]).
			Name(), nil
	}
	if matches = regexPath_Project_Region_Site_Device.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetDeviceId(matches[4]).
			SetId(matches[5]).
			Name(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as Property name", name)
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
	if property_RegexpId.MatchString(nameOrId) {
		return &Name{PropertyId: nameOrId}, nil
	} else {
		return nil, fmt.Errorf("unable to parse '%s' as Property name or id", nameOrId)
	}
}

func (name *Name) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 0 {
		return status.Errorf(codes.InvalidArgument, "No segments given for Property name")
	}
	if err := name.ParentName.SetFromSegments(segments[:len(segments)-1]); err != nil {
		return err
	}
	if segments[len(segments)-1].CollectionLowerJson != "properties" {
		return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form Property name", segments)
	}
	name.PropertyId = segments[len(segments)-1].Id
	return nil
}

func (name *Name) GetSiteName() *site.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetSiteName()
}
func (name *Name) GetBuildingName() *building.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetBuildingName()
}
func (name *Name) GetFloorName() *floor.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetFloorName()
}
func (name *Name) GetAreaName() *area.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetAreaName()
}
func (name *Name) GetZoneName() *zone.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetZoneName()
}
func (name *Name) GetDeviceName() *device.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetDeviceName()
}

func (name *Name) IsSpecified() bool {
	if name == nil || name.Pattern == "" || name.PropertyId == "" {
		return false
	}
	return name.ParentName.IsSpecified()
}

func (name *Name) IsFullyQualified() bool {
	if name == nil {
		return false
	}
	if name.ParentName.IsFullyQualified() == false {
		return false
	}
	if name.PropertyId == "" || name.PropertyId == gotenresource.WildcardId {
		return false
	}
	return true
}

func (name *Name) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for Property is not fully qualified")
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
			"propertyId": name.PropertyId,
			"buildingId": name.BuildingId,
			"floorId":    name.FloorId,
			"areaId":     name.AreaId,
			"zoneId":     name.ZoneId,
			"deviceId":   name.DeviceId,
		}
	}
	return map[string]string{
		"projectId":  "",
		"regionId":   "",
		"siteId":     "",
		"propertyId": "",
		"buildingId": "",
		"floorId":    "",
		"areaId":     "",
		"zoneId":     "",
		"deviceId":   "",
	}
}

func (name *Name) GetSegments() gotenresource.NameSegments {
	if name == nil || name.Pattern == "" {
		return nil
	}
	segments := name.ParentName.GetSegments()
	return append(segments, gotenresource.NameSegment{
		CollectionLowerJson: "properties",
		Id:                  name.PropertyId,
	})
}

func (name *Name) GetIParentName() gotenresource.Name {
	if name == nil {
		return (*ParentName)(nil)
	}
	return &name.ParentName
}

func (name *Name) GetIUnderlyingParentName() gotenresource.Name {
	if parentName := name.GetSiteName(); parentName != nil {
		return parentName
	}
	if parentName := name.GetBuildingName(); parentName != nil {
		return parentName
	}
	if parentName := name.GetFloorName(); parentName != nil {
		return parentName
	}
	if parentName := name.GetAreaName(); parentName != nil {
		return parentName
	}
	if parentName := name.GetZoneName(); parentName != nil {
		return parentName
	}
	if parentName := name.GetDeviceName(); parentName != nil {
		return parentName
	}
	return nil
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
	result += "properties/" + name.PropertyId
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
	if name.PropertyId != other1.PropertyId {
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
	if name.PropertyId != other1.PropertyId {
		return name.PropertyId == gotenresource.WildcardId
	}

	return true
}

// Google CEL integration Type
var celName = types.NewTypeValue("Name", traits.ReceiverType)

func (name *Name) TypeName() string {
	return ".ntt.workplace.v1alpha2.Property.Name"
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
		return types.String(name.PropertyId)
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
	property *Property
}

func MakeReference(name *Name, property *Property) (*Reference, error) {
	return &Reference{
		Name:     *name,
		property: property,
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

func (ref *Reference) Resolve(resolved *Property) {
	ref.property = resolved
}

func (ref *Reference) ResolveRaw(res gotenresource.Resource) error {
	if typedRes, ok := res.(*Property); ok && typedRes != nil {
		ref.Resolve(typedRes)
		return nil
	}
	return status.Errorf(codes.Internal, "Invalid resource type for Property: %s", reflect.TypeOf(res).Elem().Name())
}

func (ref *Reference) Resolved() bool {
	return ref.property != nil
}

func (ref *Reference) ClearCached() {
	ref.property = nil
}

func (ref *Reference) GetProperty() *Property {
	return ref.property
}

func (ref *Reference) GetRawResource() gotenresource.Resource {
	return ref.property
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
		return "", status.Errorf(codes.InvalidArgument, "Name for Property is not fully qualified")
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
		"propertyId": "",
		"buildingId": "",
		"floorId":    "",
		"areaId":     "",
		"zoneId":     "",
		"deviceId":   "",
	}
}

func (ref *Reference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.Name.GetSegments()
	}
	return nil
}

func (ref *Reference) GetIParentName() gotenresource.Name {
	if ref == nil {
		return (*ParentName)(nil)
	}
	return ref.Name.GetIParentName()
}

func (ref *Reference) GetIUnderlyingParentName() gotenresource.Name {
	if ref != nil {
		return ref.Name.GetIUnderlyingParentName()
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
