// Code generated by protoc-gen-goten-resource
// Resource: Zone
// DO NOT EDIT!!!

package zone

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

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
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &area.Area{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
	_ = &floor.Floor{}
	_ = &site.Site{}
)

var parentRegexPath_Project_Region_Site = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})$")
var parentRegexPath_Project_Region_Site_Building = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})$")
var parentRegexPath_Project_Region_Site_Building_Floor = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})$")
var parentRegexPath_Project_Region_Site_Building_Floor_Area = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/sites/(?P<site_id>-|[\\w][\\w.-]{0,127})/buildings/(?P<building_id>-|[\\w][\\w.-]{0,127})/floors/(?P<floor_id>-|[\\w][\\w.-]{0,127})/areas/(?P<area_id>-|[\\w][\\w.-]{0,127})$")

type ParentName struct {
	NamePattern
	ProjectId  string `firestore:"projectId"`
	RegionId   string `firestore:"regionId"`
	SiteId     string `firestore:"siteId"`
	BuildingId string `firestore:"buildingId"`
	FloorId    string `firestore:"floorId"`
	AreaId     string `firestore:"areaId"`
}

func ParseParentName(name string) (*ParentName, error) {
	var matches []string
	if matches = parentRegexPath_Project_Region_Site.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			Parent(), nil
	}
	if matches = parentRegexPath_Project_Region_Site_Building.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			Parent(), nil
	}
	if matches = parentRegexPath_Project_Region_Site_Building_Floor.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetFloorId(matches[5]).
			Parent(), nil
	}
	if matches = parentRegexPath_Project_Region_Site_Building_Floor_Area.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetSiteId(matches[3]).
			SetBuildingId(matches[4]).
			SetFloorId(matches[5]).
			SetAreaId(matches[6]).
			Parent(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as Zone parent name", name)
}

func MustParseParentName(name string) *ParentName {
	result, err := ParseParentName(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (name *ParentName) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 3 && segments[0].CollectionLowerJson == "projects" && segments[1].CollectionLowerJson == "regions" && segments[2].CollectionLowerJson == "sites" {
		name.Pattern = NamePattern_Project_Region_Site
		name.ProjectId = segments[0].Id
		name.RegionId = segments[1].Id
		name.SiteId = segments[2].Id
		return nil
	} else if len(segments) == 4 && segments[0].CollectionLowerJson == "projects" && segments[1].CollectionLowerJson == "regions" && segments[2].CollectionLowerJson == "sites" && segments[3].CollectionLowerJson == "buildings" {
		name.Pattern = NamePattern_Project_Region_Site_Building
		name.ProjectId = segments[0].Id
		name.RegionId = segments[1].Id
		name.SiteId = segments[2].Id
		name.BuildingId = segments[3].Id
		return nil
	} else if len(segments) == 5 && segments[0].CollectionLowerJson == "projects" && segments[1].CollectionLowerJson == "regions" && segments[2].CollectionLowerJson == "sites" && segments[3].CollectionLowerJson == "buildings" && segments[4].CollectionLowerJson == "floors" {
		name.Pattern = NamePattern_Project_Region_Site_Building_Floor
		name.ProjectId = segments[0].Id
		name.RegionId = segments[1].Id
		name.SiteId = segments[2].Id
		name.BuildingId = segments[3].Id
		name.FloorId = segments[4].Id
		return nil
	} else if len(segments) == 6 && segments[0].CollectionLowerJson == "projects" && segments[1].CollectionLowerJson == "regions" && segments[2].CollectionLowerJson == "sites" && segments[3].CollectionLowerJson == "buildings" && segments[4].CollectionLowerJson == "floors" && segments[5].CollectionLowerJson == "areas" {
		name.Pattern = NamePattern_Project_Region_Site_Building_Floor_Area
		name.ProjectId = segments[0].Id
		name.RegionId = segments[1].Id
		name.SiteId = segments[2].Id
		name.BuildingId = segments[3].Id
		name.FloorId = segments[4].Id
		name.AreaId = segments[5].Id
		return nil
	}
	return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form Zone parent name", segments)
}

func (name *ParentName) GetSiteName() *site.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		return site.NewNameBuilder().
			SetId(name.SiteId).
			SetProjectId(name.ProjectId).
			SetRegionId(name.RegionId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) GetBuildingName() *building.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site_Building:
		return building.NewNameBuilder().
			SetId(name.BuildingId).
			SetProjectId(name.ProjectId).
			SetRegionId(name.RegionId).
			SetSiteId(name.SiteId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) GetFloorName() *floor.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site_Building_Floor:
		return floor.NewNameBuilder().
			SetId(name.FloorId).
			SetProjectId(name.ProjectId).
			SetRegionId(name.RegionId).
			SetSiteId(name.SiteId).
			SetBuildingId(name.BuildingId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) GetAreaName() *area.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return area.NewNameBuilder().
			SetId(name.AreaId).
			SetProjectId(name.ProjectId).
			SetRegionId(name.RegionId).
			SetSiteId(name.SiteId).
			SetBuildingId(name.BuildingId).
			SetFloorId(name.FloorId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) IsSpecified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}
	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		return name.ProjectId != "" && name.RegionId != "" && name.SiteId != ""
	case NamePattern_Project_Region_Site_Building:
		return name.ProjectId != "" && name.RegionId != "" && name.SiteId != "" && name.BuildingId != ""
	case NamePattern_Project_Region_Site_Building_Floor:
		return name.ProjectId != "" && name.RegionId != "" && name.SiteId != "" && name.BuildingId != "" && name.FloorId != ""
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return name.ProjectId != "" && name.RegionId != "" && name.SiteId != "" && name.BuildingId != "" && name.FloorId != "" && name.AreaId != ""
	}
	return false
}

func (name *ParentName) IsFullyQualified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId && name.RegionId != "" && name.RegionId != gotenresource.WildcardId && name.SiteId != "" && name.SiteId != gotenresource.WildcardId
	case NamePattern_Project_Region_Site_Building:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId && name.RegionId != "" && name.RegionId != gotenresource.WildcardId && name.SiteId != "" && name.SiteId != gotenresource.WildcardId && name.BuildingId != "" && name.BuildingId != gotenresource.WildcardId
	case NamePattern_Project_Region_Site_Building_Floor:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId && name.RegionId != "" && name.RegionId != gotenresource.WildcardId && name.SiteId != "" && name.SiteId != gotenresource.WildcardId && name.BuildingId != "" && name.BuildingId != gotenresource.WildcardId && name.FloorId != "" && name.FloorId != gotenresource.WildcardId
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId && name.RegionId != "" && name.RegionId != gotenresource.WildcardId && name.SiteId != "" && name.SiteId != gotenresource.WildcardId && name.BuildingId != "" && name.BuildingId != gotenresource.WildcardId && name.FloorId != "" && name.FloorId != gotenresource.WildcardId && name.AreaId != "" && name.AreaId != gotenresource.WildcardId
	}

	return false
}

func (name *ParentName) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Parent name for Zone is not fully qualified")
	}
	return fmt.Sprintf("//workplace.edgelq.com/%s", name.String()), nil
}

func (name *ParentName) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (name *ParentName) GetPattern() gotenresource.NamePattern {
	if name == nil {
		return ""
	}
	return name.Pattern
}

func (name *ParentName) GetIdParts() map[string]string {
	if name != nil {
		return map[string]string{
			"projectId":  name.ProjectId,
			"regionId":   name.RegionId,
			"siteId":     name.SiteId,
			"buildingId": name.BuildingId,
			"floorId":    name.FloorId,
			"areaId":     name.AreaId,
		}
	}
	return map[string]string{
		"projectId":  "",
		"regionId":   "",
		"siteId":     "",
		"buildingId": "",
		"floorId":    "",
		"areaId":     "",
	}
}

func (name *ParentName) GetSegments() gotenresource.NameSegments {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "regions",
				Id:                  name.RegionId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "sites",
				Id:                  name.SiteId,
			},
		}
	case NamePattern_Project_Region_Site_Building:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "regions",
				Id:                  name.RegionId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "sites",
				Id:                  name.SiteId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "buildings",
				Id:                  name.BuildingId,
			},
		}
	case NamePattern_Project_Region_Site_Building_Floor:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "regions",
				Id:                  name.RegionId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "sites",
				Id:                  name.SiteId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "buildings",
				Id:                  name.BuildingId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "floors",
				Id:                  name.FloorId,
			},
		}
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "regions",
				Id:                  name.RegionId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "sites",
				Id:                  name.SiteId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "buildings",
				Id:                  name.BuildingId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "floors",
				Id:                  name.FloorId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "areas",
				Id:                  name.AreaId,
			},
		}
	}
	return nil
}

func (name *ParentName) String() string {
	if name == nil {
		return "<nil>"
	}

	if valueStr, err := name.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

func (name *ParentName) DescendsFrom(ancestor string) bool {
	if name == nil {
		return false
	}

	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		return ancestor == "projects" || ancestor == "regions" || ancestor == "sites"
	case NamePattern_Project_Region_Site_Building:
		return ancestor == "projects" || ancestor == "regions" || ancestor == "sites" || ancestor == "buildings"
	case NamePattern_Project_Region_Site_Building_Floor:
		return ancestor == "projects" || ancestor == "regions" || ancestor == "sites" || ancestor == "buildings" || ancestor == "floors"
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return ancestor == "projects" || ancestor == "regions" || ancestor == "sites" || ancestor == "buildings" || ancestor == "floors" || ancestor == "areas"
	}

	return false
}

func (name *ParentName) AsReference() *ParentReference {
	return &ParentReference{ParentName: *name}
}

func (name *ParentName) AsRawReference() gotenresource.Reference {
	return name.AsReference()
}

// implement methods required by protobuf-go library for string-struct conversion

func (name *ParentName) ProtoString() (string, error) {
	if name == nil {
		return "", nil
	}
	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		return "projects/" + name.ProjectId + "/regions/" + name.RegionId + "/sites/" + name.SiteId, nil
	case NamePattern_Project_Region_Site_Building:
		return "projects/" + name.ProjectId + "/regions/" + name.RegionId + "/sites/" + name.SiteId + "/buildings/" + name.BuildingId, nil
	case NamePattern_Project_Region_Site_Building_Floor:
		return "projects/" + name.ProjectId + "/regions/" + name.RegionId + "/sites/" + name.SiteId + "/buildings/" + name.BuildingId + "/floors/" + name.FloorId, nil
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return "projects/" + name.ProjectId + "/regions/" + name.RegionId + "/sites/" + name.SiteId + "/buildings/" + name.BuildingId + "/floors/" + name.FloorId + "/areas/" + name.AreaId, nil
	}
	return "", nil
}

func (name *ParentName) ParseProtoString(data string) error {
	parsed, err := ParseParentName(data)
	if err != nil {
		return err
	}
	*name = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (name *ParentName) GotenEqual(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentName)
	if !ok {
		other2, ok := other.(ParentName)
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
	if name.ProjectId != other1.ProjectId {
		return false
	}
	if name.RegionId != other1.RegionId {
		return false
	}
	if name.SiteId != other1.SiteId {
		return false
	}
	if name.BuildingId != other1.BuildingId {
		return false
	}
	if name.FloorId != other1.FloorId {
		return false
	}
	if name.AreaId != other1.AreaId {
		return false
	}
	if name.Pattern != other1.Pattern {
		return false
	}

	return true
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *ParentName) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentName)
	if !ok {
		other2, ok := other.(ParentName)
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

	if name.Pattern != other1.Pattern {
		return false
	}
	switch name.Pattern {
	case NamePattern_Project_Region_Site:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
		if name.RegionId != other1.RegionId &&
			name.RegionId != gotenresource.WildcardId {
			return false
		}
		if name.SiteId != other1.SiteId &&
			name.SiteId != gotenresource.WildcardId {
			return false
		}
	case NamePattern_Project_Region_Site_Building:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
		if name.RegionId != other1.RegionId &&
			name.RegionId != gotenresource.WildcardId {
			return false
		}
		if name.SiteId != other1.SiteId &&
			name.SiteId != gotenresource.WildcardId {
			return false
		}
		if name.BuildingId != other1.BuildingId &&
			name.BuildingId != gotenresource.WildcardId {
			return false
		}
	case NamePattern_Project_Region_Site_Building_Floor:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
		if name.RegionId != other1.RegionId &&
			name.RegionId != gotenresource.WildcardId {
			return false
		}
		if name.SiteId != other1.SiteId &&
			name.SiteId != gotenresource.WildcardId {
			return false
		}
		if name.BuildingId != other1.BuildingId &&
			name.BuildingId != gotenresource.WildcardId {
			return false
		}
		if name.FloorId != other1.FloorId &&
			name.FloorId != gotenresource.WildcardId {
			return false
		}
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
		if name.RegionId != other1.RegionId &&
			name.RegionId != gotenresource.WildcardId {
			return false
		}
		if name.SiteId != other1.SiteId &&
			name.SiteId != gotenresource.WildcardId {
			return false
		}
		if name.BuildingId != other1.BuildingId &&
			name.BuildingId != gotenresource.WildcardId {
			return false
		}
		if name.FloorId != other1.FloorId &&
			name.FloorId != gotenresource.WildcardId {
			return false
		}
		if name.AreaId != other1.AreaId &&
			name.AreaId != gotenresource.WildcardId {
			return false
		}
	}

	return true
}

// implement CustomTypeCliValue method
func (name *ParentName) SetFromCliFlag(raw string) error {
	parsedName, err := ParseParentName(raw)
	if err != nil {
		return err
	}
	*name = *parsedName
	return nil
}

type ParentReference struct {
	ParentName
	site     *site.Site
	building *building.Building
	floor    *floor.Floor
	area     *area.Area
}

func MakeParentReference(name *ParentName) (*ParentReference, error) {
	return &ParentReference{
		ParentName: *name,
	}, nil
}

func ParseParentReference(name string) (*ParentReference, error) {
	parsedName, err := ParseParentName(name)
	if err != nil {
		return nil, err
	}
	return MakeParentReference(parsedName)
}

func MustParseParentReference(name string) *ParentReference {
	result, err := ParseParentReference(name)
	if err != nil {
		panic(err)
	}
	return result
}
func (ref *ParentReference) GetSiteReference() *site.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project_Region_Site:
		return site.NewNameBuilder().
			SetId(ref.SiteId).
			SetProjectId(ref.ProjectId).
			SetRegionId(ref.RegionId).
			Reference()
	default:
		return nil
	}
}
func (ref *ParentReference) GetBuildingReference() *building.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project_Region_Site_Building:
		return building.NewNameBuilder().
			SetId(ref.BuildingId).
			SetProjectId(ref.ProjectId).
			SetRegionId(ref.RegionId).
			SetSiteId(ref.SiteId).
			Reference()
	default:
		return nil
	}
}
func (ref *ParentReference) GetFloorReference() *floor.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project_Region_Site_Building_Floor:
		return floor.NewNameBuilder().
			SetId(ref.FloorId).
			SetProjectId(ref.ProjectId).
			SetRegionId(ref.RegionId).
			SetSiteId(ref.SiteId).
			SetBuildingId(ref.BuildingId).
			Reference()
	default:
		return nil
	}
}
func (ref *ParentReference) GetAreaReference() *area.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project_Region_Site_Building_Floor_Area:
		return area.NewNameBuilder().
			SetId(ref.AreaId).
			SetProjectId(ref.ProjectId).
			SetRegionId(ref.RegionId).
			SetSiteId(ref.SiteId).
			SetBuildingId(ref.BuildingId).
			SetFloorId(ref.FloorId).
			Reference()
	default:
		return nil
	}
}

func (ref *ParentReference) GetUnderlyingReference() gotenresource.Reference {
	if ref == nil {
		return nil
	}
	siteRef := ref.GetSiteReference()
	if siteRef != nil {
		return siteRef
	}
	buildingRef := ref.GetBuildingReference()
	if buildingRef != nil {
		return buildingRef
	}
	floorRef := ref.GetFloorReference()
	if floorRef != nil {
		return floorRef
	}
	areaRef := ref.GetAreaReference()
	if areaRef != nil {
		return areaRef
	}

	return nil
}

func (ref *ParentReference) ResolveRaw(res gotenresource.Resource) error {
	switch typedRes := res.(type) {
	case *site.Site:
		if name := ref.GetSiteName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Site as parent of Zone, because pattern does not match")
		}
		ref.site = typedRes
		return nil
	case *building.Building:
		if name := ref.GetBuildingName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Building as parent of Zone, because pattern does not match")
		}
		ref.building = typedRes
		return nil
	case *floor.Floor:
		if name := ref.GetFloorName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Floor as parent of Zone, because pattern does not match")
		}
		ref.floor = typedRes
		return nil
	case *area.Area:
		if name := ref.GetAreaName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Area as parent of Zone, because pattern does not match")
		}
		ref.area = typedRes
		return nil
	default:
		return status.Errorf(codes.Internal, "Invalid parent type for Zone, got %s", reflect.TypeOf(res).Elem().Name())
	}
}

func (ref *ParentReference) Resolved() bool {
	if name := ref.GetSiteName(); name != nil {
		return ref.site != nil
	}
	if name := ref.GetBuildingName(); name != nil {
		return ref.building != nil
	}
	if name := ref.GetFloorName(); name != nil {
		return ref.floor != nil
	}
	if name := ref.GetAreaName(); name != nil {
		return ref.area != nil
	}
	return true
}

func (ref *ParentReference) ClearCached() {
	ref.site = nil
	ref.building = nil
	ref.floor = nil
	ref.area = nil
}

func (ref *ParentReference) GetSite() *site.Site {
	if ref == nil {
		return nil
	}
	return ref.site
}
func (ref *ParentReference) GetBuilding() *building.Building {
	if ref == nil {
		return nil
	}
	return ref.building
}
func (ref *ParentReference) GetFloor() *floor.Floor {
	if ref == nil {
		return nil
	}
	return ref.floor
}
func (ref *ParentReference) GetArea() *area.Area {
	if ref == nil {
		return nil
	}
	return ref.area
}

func (ref *ParentReference) GetRawResource() gotenresource.Resource {
	if name := ref.ParentName.GetSiteName(); name != nil {
		return ref.site
	}
	if name := ref.ParentName.GetBuildingName(); name != nil {
		return ref.building
	}
	if name := ref.ParentName.GetFloorName(); name != nil {
		return ref.floor
	}
	if name := ref.ParentName.GetAreaName(); name != nil {
		return ref.area
	}
	return nil
}

func (ref *ParentReference) IsFullyQualified() bool {
	if ref == nil {
		return false
	}
	return ref.ParentName.IsFullyQualified()
}

func (ref *ParentReference) IsSpecified() bool {
	if ref == nil {
		return false
	}
	return ref.ParentName.IsSpecified()
}

func (ref *ParentReference) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (ref *ParentReference) GetPattern() gotenresource.NamePattern {
	if ref == nil {
		return ""
	}
	return ref.Pattern
}

func (ref *ParentReference) GetIdParts() map[string]string {
	if ref != nil {
		return ref.ParentName.GetIdParts()
	}
	return map[string]string{
		"projectId":  "",
		"regionId":   "",
		"siteId":     "",
		"buildingId": "",
		"floorId":    "",
		"areaId":     "",
	}
}

func (ref *ParentReference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.ParentName.GetSegments()
	}
	return nil
}

func (ref *ParentReference) String() string {
	if ref == nil {
		return "<nil>"
	}
	return ref.ParentName.String()
}

// implement methods required by protobuf-go library for string-struct conversion

func (ref *ParentReference) ProtoString() (string, error) {
	if ref == nil {
		return "", nil
	}
	return ref.ParentName.ProtoString()
}

func (ref *ParentReference) ParseProtoString(data string) error {
	parsed, err := ParseParentReference(data)
	if err != nil {
		return err
	}
	*ref = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (ref *ParentReference) GotenEqual(other interface{}) bool {
	if other == nil {
		return ref == nil
	}
	other1, ok := other.(*ParentReference)
	if !ok {
		other2, ok := other.(ParentReference)
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
	if ref.site != other1.site {
		return false
	}
	if ref.building != other1.building {
		return false
	}
	if ref.floor != other1.floor {
		return false
	}
	if ref.area != other1.area {
		return false
	}

	return ref.ParentName.GotenEqual(other1.ParentName)
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *ParentReference) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentReference)
	if !ok {
		other2, ok := other.(ParentReference)
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
	return name.ParentName.Matches(&other1.ParentName)
}

// implement CustomTypeCliValue method
func (ref *ParentReference) SetFromCliFlag(raw string) error {
	parsedRef, err := ParseParentReference(raw)
	if err != nil {
		return err
	}
	*ref = *parsedRef
	return nil
}
