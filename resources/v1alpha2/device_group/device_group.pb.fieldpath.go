// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/device_group.proto
// DO NOT EDIT!!!

package device_group

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = new(fmt.Stringer)
	_ = reflect.DeepEqual
	_ = strings.Builder{}
	_ = time.Second

	_ = strcase.ToLowerCamel
	_ = codes.NotFound
	_ = status.Status{}
	_ = protojson.UnmarshalOptions{}
	_ = new(proto.Message)
	_ = protoregistry.GlobalTypes
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type DeviceGroup_FieldPath interface {
	gotenobject.FieldPath
	Selector() DeviceGroup_FieldPathSelector
	Get(source *DeviceGroup) []interface{}
	GetSingle(source *DeviceGroup) (interface{}, bool)
	ClearValue(item *DeviceGroup)

	// Those methods build corresponding DeviceGroup_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) DeviceGroup_FieldPathValue
	WithIArrayOfValues(values interface{}) DeviceGroup_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) DeviceGroup_FieldPathArrayItemValue
}

type DeviceGroup_FieldPathSelector int32

const (
	DeviceGroup_FieldPathSelectorName        DeviceGroup_FieldPathSelector = 0
	DeviceGroup_FieldPathSelectorDisplayName DeviceGroup_FieldPathSelector = 1
	DeviceGroup_FieldPathSelectorMetadata    DeviceGroup_FieldPathSelector = 2
)

func (s DeviceGroup_FieldPathSelector) String() string {
	switch s {
	case DeviceGroup_FieldPathSelectorName:
		return "name"
	case DeviceGroup_FieldPathSelectorDisplayName:
		return "display_name"
	case DeviceGroup_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", s))
	}
}

func BuildDeviceGroup_FieldPath(fp gotenobject.RawFieldPath) (DeviceGroup_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object DeviceGroup")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorName}, nil
		case "display_name", "displayName", "display-name":
			return &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorDisplayName}, nil
		case "metadata":
			return &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := ntt_meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &DeviceGroup_FieldSubPath{selector: DeviceGroup_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object DeviceGroup", fp)
}

func ParseDeviceGroup_FieldPath(rawField string) (DeviceGroup_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildDeviceGroup_FieldPath(fp)
}

func MustParseDeviceGroup_FieldPath(rawField string) DeviceGroup_FieldPath {
	fp, err := ParseDeviceGroup_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type DeviceGroup_FieldTerminalPath struct {
	selector DeviceGroup_FieldPathSelector
}

var _ DeviceGroup_FieldPath = (*DeviceGroup_FieldTerminalPath)(nil)

func (fp *DeviceGroup_FieldTerminalPath) Selector() DeviceGroup_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *DeviceGroup_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *DeviceGroup_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source DeviceGroup
func (fp *DeviceGroup_FieldTerminalPath) Get(source *DeviceGroup) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case DeviceGroup_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case DeviceGroup_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case DeviceGroup_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
		}
	}
	return
}

func (fp *DeviceGroup_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*DeviceGroup))
}

// GetSingle returns value pointed by specific field of from source DeviceGroup
func (fp *DeviceGroup_FieldTerminalPath) GetSingle(source *DeviceGroup) (interface{}, bool) {
	switch fp.selector {
	case DeviceGroup_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case DeviceGroup_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case DeviceGroup_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
	}
}

func (fp *DeviceGroup_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*DeviceGroup))
}

// GetDefault returns a default value of the field type
func (fp *DeviceGroup_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case DeviceGroup_FieldPathSelectorName:
		return (*Name)(nil)
	case DeviceGroup_FieldPathSelectorDisplayName:
		return ""
	case DeviceGroup_FieldPathSelectorMetadata:
		return (*ntt_meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
	}
}

func (fp *DeviceGroup_FieldTerminalPath) ClearValue(item *DeviceGroup) {
	if item != nil {
		switch fp.selector {
		case DeviceGroup_FieldPathSelectorName:
			item.Name = nil
		case DeviceGroup_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case DeviceGroup_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
		}
	}
}

func (fp *DeviceGroup_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*DeviceGroup))
}

// IsLeaf - whether field path is holds simple value
func (fp *DeviceGroup_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == DeviceGroup_FieldPathSelectorName ||
		fp.selector == DeviceGroup_FieldPathSelectorDisplayName
}

func (fp *DeviceGroup_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *DeviceGroup_FieldTerminalPath) WithIValue(value interface{}) DeviceGroup_FieldPathValue {
	switch fp.selector {
	case DeviceGroup_FieldPathSelectorName:
		return &DeviceGroup_FieldTerminalPathValue{DeviceGroup_FieldTerminalPath: *fp, value: value.(*Name)}
	case DeviceGroup_FieldPathSelectorDisplayName:
		return &DeviceGroup_FieldTerminalPathValue{DeviceGroup_FieldTerminalPath: *fp, value: value.(string)}
	case DeviceGroup_FieldPathSelectorMetadata:
		return &DeviceGroup_FieldTerminalPathValue{DeviceGroup_FieldTerminalPath: *fp, value: value.(*ntt_meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
	}
}

func (fp *DeviceGroup_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *DeviceGroup_FieldTerminalPath) WithIArrayOfValues(values interface{}) DeviceGroup_FieldPathArrayOfValues {
	fpaov := &DeviceGroup_FieldTerminalPathArrayOfValues{DeviceGroup_FieldTerminalPath: *fp}
	switch fp.selector {
	case DeviceGroup_FieldPathSelectorName:
		return &DeviceGroup_FieldTerminalPathArrayOfValues{DeviceGroup_FieldTerminalPath: *fp, values: values.([]*Name)}
	case DeviceGroup_FieldPathSelectorDisplayName:
		return &DeviceGroup_FieldTerminalPathArrayOfValues{DeviceGroup_FieldTerminalPath: *fp, values: values.([]string)}
	case DeviceGroup_FieldPathSelectorMetadata:
		return &DeviceGroup_FieldTerminalPathArrayOfValues{DeviceGroup_FieldTerminalPath: *fp, values: values.([]*ntt_meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
	}
	return fpaov
}

func (fp *DeviceGroup_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *DeviceGroup_FieldTerminalPath) WithIArrayItemValue(value interface{}) DeviceGroup_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fp.selector))
	}
}

func (fp *DeviceGroup_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type DeviceGroup_FieldSubPath struct {
	selector DeviceGroup_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ DeviceGroup_FieldPath = (*DeviceGroup_FieldSubPath)(nil)

func (fps *DeviceGroup_FieldSubPath) Selector() DeviceGroup_FieldPathSelector {
	return fps.selector
}
func (fps *DeviceGroup_FieldSubPath) AsMetadataSubPath() (ntt_meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(ntt_meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *DeviceGroup_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *DeviceGroup_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source DeviceGroup
func (fps *DeviceGroup_FieldSubPath) Get(source *DeviceGroup) (values []interface{}) {
	if asMetaFieldPath, ok := fps.AsMetadataSubPath(); ok {
		values = append(values, asMetaFieldPath.Get(source.GetMetadata())...)
	} else {
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fps.selector))
	}
	return
}

func (fps *DeviceGroup_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*DeviceGroup))
}

// GetSingle returns value of selected field from source DeviceGroup
func (fps *DeviceGroup_FieldSubPath) GetSingle(source *DeviceGroup) (interface{}, bool) {
	switch fps.selector {
	case DeviceGroup_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fps.selector))
	}
}

func (fps *DeviceGroup_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*DeviceGroup))
}

// GetDefault returns a default value of the field type
func (fps *DeviceGroup_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *DeviceGroup_FieldSubPath) ClearValue(item *DeviceGroup) {
	if item != nil {
		switch fps.selector {
		case DeviceGroup_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fps.selector))
		}
	}
}

func (fps *DeviceGroup_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*DeviceGroup))
}

// IsLeaf - whether field path is holds simple value
func (fps *DeviceGroup_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *DeviceGroup_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&DeviceGroup_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *DeviceGroup_FieldSubPath) WithIValue(value interface{}) DeviceGroup_FieldPathValue {
	return &DeviceGroup_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *DeviceGroup_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *DeviceGroup_FieldSubPath) WithIArrayOfValues(values interface{}) DeviceGroup_FieldPathArrayOfValues {
	return &DeviceGroup_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *DeviceGroup_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *DeviceGroup_FieldSubPath) WithIArrayItemValue(value interface{}) DeviceGroup_FieldPathArrayItemValue {
	return &DeviceGroup_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *DeviceGroup_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// DeviceGroup_FieldPathValue allows storing values for DeviceGroup fields according to their type
type DeviceGroup_FieldPathValue interface {
	DeviceGroup_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **DeviceGroup)
	CompareWith(*DeviceGroup) (cmp int, comparable bool)
}

func ParseDeviceGroup_FieldPathValue(pathStr, valueStr string) (DeviceGroup_FieldPathValue, error) {
	fp, err := ParseDeviceGroup_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing DeviceGroup field path value from %s: %v", valueStr, err)
	}
	return fpv.(DeviceGroup_FieldPathValue), nil
}

func MustParseDeviceGroup_FieldPathValue(pathStr, valueStr string) DeviceGroup_FieldPathValue {
	fpv, err := ParseDeviceGroup_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type DeviceGroup_FieldTerminalPathValue struct {
	DeviceGroup_FieldTerminalPath
	value interface{}
}

var _ DeviceGroup_FieldPathValue = (*DeviceGroup_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'DeviceGroup' as interface{}
func (fpv *DeviceGroup_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *DeviceGroup_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *DeviceGroup_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *DeviceGroup_FieldTerminalPathValue) AsMetadataValue() (*ntt_meta.Meta, bool) {
	res, ok := fpv.value.(*ntt_meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object DeviceGroup
func (fpv *DeviceGroup_FieldTerminalPathValue) SetTo(target **DeviceGroup) {
	if *target == nil {
		*target = new(DeviceGroup)
	}
	switch fpv.selector {
	case DeviceGroup_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case DeviceGroup_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case DeviceGroup_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*ntt_meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fpv.selector))
	}
}

func (fpv *DeviceGroup_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*DeviceGroup)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'DeviceGroup_FieldTerminalPathValue' with the value under path in 'DeviceGroup'.
func (fpv *DeviceGroup_FieldTerminalPathValue) CompareWith(source *DeviceGroup) (int, bool) {
	switch fpv.selector {
	case DeviceGroup_FieldPathSelectorName:
		leftValue := fpv.value.(*Name)
		rightValue := source.GetName()
		if leftValue == nil {
			if rightValue != nil {
				return -1, true
			}
			return 0, true
		}
		if rightValue == nil {
			return 1, true
		}
		if leftValue.String() == rightValue.String() {
			return 0, true
		} else if leftValue.String() < rightValue.String() {
			return -1, true
		} else {
			return 1, true
		}
	case DeviceGroup_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case DeviceGroup_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fpv.selector))
	}
}

func (fpv *DeviceGroup_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*DeviceGroup))
}

type DeviceGroup_FieldSubPathValue struct {
	DeviceGroup_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ DeviceGroup_FieldPathValue = (*DeviceGroup_FieldSubPathValue)(nil)

func (fpvs *DeviceGroup_FieldSubPathValue) AsMetadataPathValue() (ntt_meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *DeviceGroup_FieldSubPathValue) SetTo(target **DeviceGroup) {
	if *target == nil {
		*target = new(DeviceGroup)
	}
	switch fpvs.Selector() {
	case DeviceGroup_FieldPathSelectorMetadata:
		fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fpvs.Selector()))
	}
}

func (fpvs *DeviceGroup_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*DeviceGroup)
	fpvs.SetTo(&typedObject)
}

func (fpvs *DeviceGroup_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *DeviceGroup_FieldSubPathValue) CompareWith(source *DeviceGroup) (int, bool) {
	switch fpvs.Selector() {
	case DeviceGroup_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fpvs.Selector()))
	}
}

func (fpvs *DeviceGroup_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*DeviceGroup))
}

// DeviceGroup_FieldPathArrayItemValue allows storing single item in Path-specific values for DeviceGroup according to their type
// Present only for array (repeated) types.
type DeviceGroup_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	DeviceGroup_FieldPath
	ContainsValue(*DeviceGroup) bool
}

// ParseDeviceGroup_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseDeviceGroup_FieldPathArrayItemValue(pathStr, valueStr string) (DeviceGroup_FieldPathArrayItemValue, error) {
	fp, err := ParseDeviceGroup_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing DeviceGroup field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(DeviceGroup_FieldPathArrayItemValue), nil
}

func MustParseDeviceGroup_FieldPathArrayItemValue(pathStr, valueStr string) DeviceGroup_FieldPathArrayItemValue {
	fpaiv, err := ParseDeviceGroup_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type DeviceGroup_FieldTerminalPathArrayItemValue struct {
	DeviceGroup_FieldTerminalPath
	value interface{}
}

var _ DeviceGroup_FieldPathArrayItemValue = (*DeviceGroup_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object DeviceGroup as interface{}
func (fpaiv *DeviceGroup_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *DeviceGroup_FieldTerminalPathArrayItemValue) GetSingle(source *DeviceGroup) (interface{}, bool) {
	return nil, false
}

func (fpaiv *DeviceGroup_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*DeviceGroup))
}

// Contains returns a boolean indicating if value that is being held is present in given 'DeviceGroup'
func (fpaiv *DeviceGroup_FieldTerminalPathArrayItemValue) ContainsValue(source *DeviceGroup) bool {
	slice := fpaiv.DeviceGroup_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if asProtoMsg, ok := fpaiv.value.(proto.Message); ok {
			if proto.Equal(asProtoMsg, v.(proto.Message)) {
				return true
			}
		} else if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type DeviceGroup_FieldSubPathArrayItemValue struct {
	DeviceGroup_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *DeviceGroup_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *DeviceGroup_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (ntt_meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'DeviceGroup'
func (fpaivs *DeviceGroup_FieldSubPathArrayItemValue) ContainsValue(source *DeviceGroup) bool {
	switch fpaivs.Selector() {
	case DeviceGroup_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for DeviceGroup: %d", fpaivs.Selector()))
	}
}

// DeviceGroup_FieldPathArrayOfValues allows storing slice of values for DeviceGroup fields according to their type
type DeviceGroup_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	DeviceGroup_FieldPath
}

func ParseDeviceGroup_FieldPathArrayOfValues(pathStr, valuesStr string) (DeviceGroup_FieldPathArrayOfValues, error) {
	fp, err := ParseDeviceGroup_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing DeviceGroup field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(DeviceGroup_FieldPathArrayOfValues), nil
}

func MustParseDeviceGroup_FieldPathArrayOfValues(pathStr, valuesStr string) DeviceGroup_FieldPathArrayOfValues {
	fpaov, err := ParseDeviceGroup_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type DeviceGroup_FieldTerminalPathArrayOfValues struct {
	DeviceGroup_FieldTerminalPath
	values interface{}
}

var _ DeviceGroup_FieldPathArrayOfValues = (*DeviceGroup_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *DeviceGroup_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case DeviceGroup_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case DeviceGroup_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case DeviceGroup_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*ntt_meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *DeviceGroup_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *DeviceGroup_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *DeviceGroup_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*ntt_meta.Meta, bool) {
	res, ok := fpaov.values.([]*ntt_meta.Meta)
	return res, ok
}

type DeviceGroup_FieldSubPathArrayOfValues struct {
	DeviceGroup_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ DeviceGroup_FieldPathArrayOfValues = (*DeviceGroup_FieldSubPathArrayOfValues)(nil)

func (fpsaov *DeviceGroup_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *DeviceGroup_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (ntt_meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(ntt_meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
