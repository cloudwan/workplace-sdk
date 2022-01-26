// Code generated by protoc-gen-goten-resource
// Resource: DeviceGroup
// DO NOT EDIT!!!

package device_group

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
)

// OrderByField is single item in order_by specification
// it's string format is composed of 2 white-space separated values:
// - fieldPath and direction, e.g. "state.capacity desc".
// if direction is not provided, it defaults to "asc" (ascending)
type OrderByField struct {
	FieldPath DeviceGroup_FieldPath
	Direction gotenresource.OrderDirection
}

func (orderByFld *OrderByField) GetDirection() gotenresource.OrderDirection {
	return orderByFld.Direction
}

func (orderByFld *OrderByField) GetFieldPath() gotenobject.FieldPath {
	return orderByFld.FieldPath
}

func (orderByFld *OrderByField) CompareWithDirection(left, right *DeviceGroup) int {
	leftValue, leftPresent := orderByFld.FieldPath.GetSingle(left)
	_, rightPresent := orderByFld.FieldPath.GetSingle(right)
	if leftPresent != rightPresent {
		if orderByFld.Direction == gotenresource.DirectionASC {
			if !leftPresent {
				return -1
			} else {
				return 1
			}
		} else {
			if leftPresent {
				return -1
			} else {
				return 1
			}
		}
	}
	if !leftPresent {
		return 0
	}
	leftFpv := orderByFld.FieldPath.WithIValue(leftValue)
	result, comparable := leftFpv.CompareWith(right)
	if comparable && result != 0 {
		if orderByFld.Direction == gotenresource.DirectionASC {
			return result
		} else if result > 0 {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

// OrderBy Is string encoded Custom Protobuf type, which handles "order_by" field
// order_by consists of coma delimited OrderBy specs, which denote ordering priority,
// e.g. "state.value asc, state.capacity desc"
type OrderBy struct {
	OrderByFields []OrderByField
}

func (orderBy *OrderBy) String() string {
	if orderBy == nil {
		return "<nil>"
	}
	if valueStr, err := orderBy.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

// implement methods required by protobuf-go library for string-struct conversion

func (orderBy *OrderBy) ProtoString() (string, error) {
	if orderBy == nil {
		return "", nil
	}
	var results []string
	for _, orderByField := range orderBy.OrderByFields {
		results = append(results, fmt.Sprintf("%s %s", orderByField.FieldPath, orderByField.Direction))
	}
	return strings.Join(results, ","), nil
}

func (orderBy *OrderBy) ParseProtoString(data string) error {
	groups := strings.Split(data, ",")
	orderBys := make([]OrderByField, 0, len(groups))
	for _, orderByStr := range groups {
		orderByField, err := orderBy.parseOrderByField(orderByStr)
		if err != nil {
			return err
		}
		orderBys = append(orderBys, orderByField)
	}
	orderBy.OrderByFields = orderBys
	return nil
}

func (orderBy *OrderBy) Sort(results DeviceGroupList) {
	if orderBy == nil {
		return
	}
	sort.Slice(results, func(i, j int) bool {
		left, right := results[i], results[j]
		for _, fld := range orderBy.OrderByFields {
			compareResult := fld.CompareWithDirection(left, right)
			if compareResult != 0 {
				return compareResult < 0
			}
		}
		return false
	})
}

func (orderBy *OrderBy) SortRaw(results gotenresource.ResourceList) {
	orderBy.Sort(results.(DeviceGroupList))
}

func (orderBy *OrderBy) InsertSorted(sorted DeviceGroupList, elem *DeviceGroup) (DeviceGroupList, int) {
	if orderBy == nil {
		return append(sorted, elem), len(sorted)
	}
	i := sort.Search(len(sorted), func(i int) bool {
		comparedTo := sorted[i]
		for _, fld := range orderBy.OrderByFields {
			compareResult := fld.CompareWithDirection(elem, comparedTo)
			if compareResult != 0 {
				return compareResult < 0
			}
		}
		return false
	})
	sorted = append(sorted, elem)
	copy(sorted[i+1:], sorted[i:])
	sorted[i] = elem
	return sorted, i
}

func (orderBy *OrderBy) InsertSortedRaw(sorted gotenresource.ResourceList, elem gotenresource.Resource) (gotenresource.ResourceList, int) {
	return orderBy.InsertSorted(sorted.(DeviceGroupList), elem.(*DeviceGroup))
}

func (orderBy *OrderBy) Compare(left, right *DeviceGroup) int {
	if orderBy == nil {
		return 0
	}
	for _, fld := range orderBy.OrderByFields {
		compareResult := fld.CompareWithDirection(left, right)
		if compareResult != 0 {
			return compareResult
		}
	}
	return 0
}

func (orderBy *OrderBy) CompareRaw(left, right gotenresource.Resource) int {
	return orderBy.Compare(left.(*DeviceGroup), right.(*DeviceGroup))
}

func (orderBy *OrderBy) parseOrderByField(orderByStr string) (OrderByField, error) {
	split := strings.Fields(orderByStr)
	if len(split) > 2 || len(split) == 0 {
		return OrderByField{}, status.Errorf(codes.InvalidArgument, "string '%s' doesn't parse as order by query specifier", orderByStr)
	}
	// parse direction
	direction := gotenresource.DirectionASC
	if len(split) > 1 {
		var err error
		if direction, err = gotenresource.OrderDirectionFromString(split[1]); err != nil {
			return OrderByField{}, err
		}
	}
	// parse field path
	fp, err := ParseDeviceGroup_FieldPath(split[0])
	if err != nil {
		return OrderByField{}, err
	}

	return OrderByField{
		FieldPath: fp,
		Direction: direction,
	}, nil
}

func (orderBy *OrderBy) SetFromCliFlag(raw string) error {
	field, err := orderBy.parseOrderByField(raw)
	if err != nil {
		return err
	}
	orderBy.OrderByFields = append(orderBy.OrderByFields, field)
	return nil
}

func (orderBy *OrderBy) GetOrderByFields() []gotenresource.OrderByField {
	if orderBy == nil {
		return nil
	}

	fields := make([]gotenresource.OrderByField, len(orderBy.OrderByFields))
	for idx := range orderBy.OrderByFields {
		fields[idx] = &orderBy.OrderByFields[idx]
	}
	return fields
}

func (orderBy *OrderBy) GetFieldMask() *DeviceGroup_FieldMask {
	fieldMask := &DeviceGroup_FieldMask{}
	if orderBy == nil {
		return fieldMask
	}
	for _, orderByField := range orderBy.OrderByFields {
		fieldMask.Paths = append(fieldMask.Paths, orderByField.FieldPath)
	}
	return fieldMask
}

func (orderBy *OrderBy) GetRawFieldMask() gotenobject.FieldMask {
	return orderBy.GetFieldMask()
}

// PagerCursor is protobuf Custom Type, which (de)serializes "string page_token" for API List processing
// Database adapter implementation must use this cursor when Paginating list views
// Token is composed of 3 values (dot separated in serialized form)
// - CursorValue: Backend-specific value of the cursor.
// - PageDirection: either l (left) or r (right), which hints DB Adapter whether Snapshot marks Start or End of result
// - Inclusion: either i (inclusive) or e (exclusive) - Whether cursor marks exact point or right before/after (depending on direction)
type PagerCursor struct {
	CursorValue   gotenresource.CursorValue
	Inclusion     gotenresource.CursorInclusion
	PageDirection gotenresource.PageDirection
}

func (cursor *PagerCursor) String() string {
	if cursor == nil {
		return "<nil>"
	}
	if valueStr, err := cursor.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

func (cursor *PagerCursor) IsEmpty() bool {
	return cursor == nil || cursor.CursorValue == nil || reflect.ValueOf(cursor.CursorValue).IsNil()
}

// implement methods required by protobuf-go library for string-struct conversion

func (cursor *PagerCursor) ProtoString() (string, error) {
	if cursor.IsEmpty() {
		return "", nil
	}
	return fmt.Sprintf("%s.%s.%s.%s",
		cursor.PageDirection, cursor.Inclusion, cursor.CursorValue.GetValueType(), cursor.CursorValue.String()), nil
}

func (cursor *PagerCursor) ParseProtoString(data string) (err error) {
	split := strings.Split(data, ".")
	if len(split) != 4 {
		return status.Errorf(codes.InvalidArgument, "invalid cursor format: '%s'", data)
	}
	if cursor.PageDirection, err = gotenresource.PageDirectionFromString(split[0]); err != nil {
		return err
	}
	if cursor.Inclusion, err = gotenresource.CursorInclusionFromString(split[1]); err != nil {
		return err
	}
	switch gotenresource.CursorValueType(split[2]) {
	case gotenresource.CustomCursorValueType:
		cursor.CursorValue, err = gotenresource.ParseCustomCursorValue(split[3])
		if err != nil {
			return err
		}
	case gotenresource.OffsetCursorValueType:
		cursor.CursorValue, err = gotenresource.ParseOffsetCursorValue(split[3])
		if err != nil {
			return err
		}
	case gotenresource.SnapshotCursorValueType:
		cursor.CursorValue, err = gotenresource.ParseSnapshotCursorValue(GetDescriptor(), split[3])
		if err != nil {
			return err
		}
	default:
		return status.Errorf(codes.InvalidArgument, "invalid cursor value type: '%s'", split[2])
	}
	return nil
}

func (cursor *PagerCursor) SetFromCliFlag(raw string) error {
	return cursor.ParseProtoString(raw)
}

func (cursor *PagerCursor) GetPageDirection() gotenresource.PageDirection {
	if cursor == nil {
		return ""
	}
	return cursor.PageDirection
}

func (cursor *PagerCursor) GetInclusion() gotenresource.CursorInclusion {
	if cursor == nil {
		return ""
	}
	return cursor.Inclusion
}

func (cursor *PagerCursor) GetValue() gotenresource.CursorValue {
	if cursor == nil {
		return nil
	}
	return cursor.CursorValue
}

func (cursor *PagerCursor) SetPageDirection(direction gotenresource.PageDirection) {
	cursor.PageDirection = direction
}

func (cursor *PagerCursor) SetInclusion(inclusion gotenresource.CursorInclusion) {
	cursor.Inclusion = inclusion
}

func (cursor *PagerCursor) SetCursorValue(value gotenresource.CursorValue) {
	cursor.CursorValue = value
}

// PagerQuery is main struct used for assisting server and database to perform Pagination
type PagerQuery struct {
	OrderBy     *OrderBy
	Cursor      *PagerCursor
	Limit       int
	PeekForward bool
}

// MakePagerQuery builds pager from API data and applies defaults
func MakePagerQuery(orderBy *OrderBy, cursor *PagerCursor, pageSize int32, peekForward bool) *PagerQuery {
	if pageSize == 0 {
		pageSize = 100
	}
	if orderBy == nil {
		orderBy = &OrderBy{}
	}
	hasName := false
	for _, orderByField := range orderBy.OrderByFields {
		if orderByField.FieldPath.Selector() == DeviceGroup_FieldPathSelectorName {
			hasName = true
		}
	}
	if !hasName {
		nameDirection := gotenresource.DirectionASC
		if len(orderBy.OrderByFields) > 0 {
			lastOrderBy := orderBy.OrderByFields[len(orderBy.OrderByFields)-1]
			nameDirection = lastOrderBy.Direction
		}
		orderBy.OrderByFields = append(orderBy.OrderByFields, OrderByField{
			FieldPath: &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorName},
			Direction: nameDirection,
		})
	}

	return &PagerQuery{
		OrderBy:     orderBy,
		Cursor:      cursor,
		Limit:       int(pageSize),
		PeekForward: peekForward,
	}
}

func (p *PagerQuery) GetOrderBy() gotenresource.OrderBy {
	if p == nil {
		return (*OrderBy)(nil)
	}
	return p.OrderBy
}

func (p *PagerQuery) GetCursor() gotenresource.Cursor {
	if p == nil {
		return (*PagerCursor)(nil)
	}
	return p.Cursor
}

func (p *PagerQuery) GetLimit() int {
	if p == nil {
		return 0
	}
	return p.Limit
}

func (p *PagerQuery) GetPeekForward() bool {
	if p == nil {
		return false
	}
	return p.PeekForward
}

func (p *PagerQuery) PageDirection() gotenresource.PageDirection {
	if p == nil || p.Cursor == nil {
		return gotenresource.PageRight
	} else {
		return p.Cursor.PageDirection
	}
}

func (p *PagerQuery) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}
