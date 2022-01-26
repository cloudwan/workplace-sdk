// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/agent.proto
// DO NOT EDIT!!!

package agent

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
	_ = json.Marshaler(nil)
	_ = fmt.Stringer(nil)
	_ = reflect.DeepEqual
	_ = strings.Builder{}
	_ = time.Second

	_ = strcase.ToLowerCamel
	_ = codes.NotFound
	_ = status.Status{}
	_ = protojson.UnmarshalOptions{}
	_ = proto.Message(nil)
	_ = protoregistry.GlobalTypes
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Agent_FieldPath interface {
	gotenobject.FieldPath
	Selector() Agent_FieldPathSelector
	Get(source *Agent) []interface{}
	GetSingle(source *Agent) (interface{}, bool)
	ClearValue(item *Agent)

	// Those methods build corresponding Agent_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Agent_FieldPathValue
	WithIArrayOfValues(values interface{}) Agent_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Agent_FieldPathArrayItemValue
}

type Agent_FieldPathSelector int32

const (
	Agent_FieldPathSelectorName     Agent_FieldPathSelector = 0
	Agent_FieldPathSelectorMetadata Agent_FieldPathSelector = 1
)

func (s Agent_FieldPathSelector) String() string {
	switch s {
	case Agent_FieldPathSelectorName:
		return "name"
	case Agent_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", s))
	}
}

func BuildAgent_FieldPath(fp gotenobject.RawFieldPath) (Agent_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Agent")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Agent_FieldTerminalPath{selector: Agent_FieldPathSelectorName}, nil
		case "metadata":
			return &Agent_FieldTerminalPath{selector: Agent_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := ntt_meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Agent_FieldSubPath{selector: Agent_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Agent", fp)
}

func ParseAgent_FieldPath(rawField string) (Agent_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildAgent_FieldPath(fp)
}

func MustParseAgent_FieldPath(rawField string) Agent_FieldPath {
	fp, err := ParseAgent_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Agent_FieldTerminalPath struct {
	selector Agent_FieldPathSelector
}

var _ Agent_FieldPath = (*Agent_FieldTerminalPath)(nil)

func (fp *Agent_FieldTerminalPath) Selector() Agent_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Agent_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Agent_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Agent
func (fp *Agent_FieldTerminalPath) Get(source *Agent) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Agent_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Agent_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
		}
	}
	return
}

func (fp *Agent_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Agent))
}

// GetSingle returns value pointed by specific field of from source Agent
func (fp *Agent_FieldTerminalPath) GetSingle(source *Agent) (interface{}, bool) {
	switch fp.selector {
	case Agent_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Agent_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
	}
}

func (fp *Agent_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Agent))
}

// GetDefault returns a default value of the field type
func (fp *Agent_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Agent_FieldPathSelectorName:
		return (*Name)(nil)
	case Agent_FieldPathSelectorMetadata:
		return (*ntt_meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
	}
}

func (fp *Agent_FieldTerminalPath) ClearValue(item *Agent) {
	if item != nil {
		switch fp.selector {
		case Agent_FieldPathSelectorName:
			item.Name = nil
		case Agent_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
		}
	}
}

func (fp *Agent_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Agent))
}

// IsLeaf - whether field path is holds simple value
func (fp *Agent_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Agent_FieldPathSelectorName
}

func (fp *Agent_FieldTerminalPath) WithIValue(value interface{}) Agent_FieldPathValue {
	switch fp.selector {
	case Agent_FieldPathSelectorName:
		return &Agent_FieldTerminalPathValue{Agent_FieldTerminalPath: *fp, value: value.(*Name)}
	case Agent_FieldPathSelectorMetadata:
		return &Agent_FieldTerminalPathValue{Agent_FieldTerminalPath: *fp, value: value.(*ntt_meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
	}
}

func (fp *Agent_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Agent_FieldTerminalPath) WithIArrayOfValues(values interface{}) Agent_FieldPathArrayOfValues {
	fpaov := &Agent_FieldTerminalPathArrayOfValues{Agent_FieldTerminalPath: *fp}
	switch fp.selector {
	case Agent_FieldPathSelectorName:
		return &Agent_FieldTerminalPathArrayOfValues{Agent_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Agent_FieldPathSelectorMetadata:
		return &Agent_FieldTerminalPathArrayOfValues{Agent_FieldTerminalPath: *fp, values: values.([]*ntt_meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
	}
	return fpaov
}

func (fp *Agent_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Agent_FieldTerminalPath) WithIArrayItemValue(value interface{}) Agent_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fp.selector))
	}
}

func (fp *Agent_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Agent_FieldSubPath struct {
	selector Agent_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Agent_FieldPath = (*Agent_FieldSubPath)(nil)

func (fps *Agent_FieldSubPath) Selector() Agent_FieldPathSelector {
	return fps.selector
}
func (fps *Agent_FieldSubPath) AsMetadataSubPath() (ntt_meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(ntt_meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Agent_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Agent_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Agent
func (fps *Agent_FieldSubPath) Get(source *Agent) (values []interface{}) {
	if asMetaFieldPath, ok := fps.AsMetadataSubPath(); ok {
		values = append(values, asMetaFieldPath.Get(source.GetMetadata())...)
	} else {
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fps.selector))
	}
	return
}

func (fps *Agent_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Agent))
}

// GetSingle returns value of selected field from source Agent
func (fps *Agent_FieldSubPath) GetSingle(source *Agent) (interface{}, bool) {
	switch fps.selector {
	case Agent_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fps.selector))
	}
}

func (fps *Agent_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Agent))
}

// GetDefault returns a default value of the field type
func (fps *Agent_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Agent_FieldSubPath) ClearValue(item *Agent) {
	if item != nil {
		switch fps.selector {
		case Agent_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Agent: %d", fps.selector))
		}
	}
}

func (fps *Agent_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Agent))
}

// IsLeaf - whether field path is holds simple value
func (fps *Agent_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Agent_FieldSubPath) WithIValue(value interface{}) Agent_FieldPathValue {
	return &Agent_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Agent_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Agent_FieldSubPath) WithIArrayOfValues(values interface{}) Agent_FieldPathArrayOfValues {
	return &Agent_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Agent_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Agent_FieldSubPath) WithIArrayItemValue(value interface{}) Agent_FieldPathArrayItemValue {
	return &Agent_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Agent_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Agent_FieldPathValue allows storing values for Agent fields according to their type
type Agent_FieldPathValue interface {
	Agent_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Agent)
	CompareWith(*Agent) (cmp int, comparable bool)
}

func ParseAgent_FieldPathValue(pathStr, valueStr string) (Agent_FieldPathValue, error) {
	fp, err := ParseAgent_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Agent field path value from %s: %v", valueStr, err)
	}
	return fpv.(Agent_FieldPathValue), nil
}

func MustParseAgent_FieldPathValue(pathStr, valueStr string) Agent_FieldPathValue {
	fpv, err := ParseAgent_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Agent_FieldTerminalPathValue struct {
	Agent_FieldTerminalPath
	value interface{}
}

var _ Agent_FieldPathValue = (*Agent_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Agent' as interface{}
func (fpv *Agent_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Agent_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Agent_FieldTerminalPathValue) AsMetadataValue() (*ntt_meta.Meta, bool) {
	res, ok := fpv.value.(*ntt_meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object Agent
func (fpv *Agent_FieldTerminalPathValue) SetTo(target **Agent) {
	if *target == nil {
		*target = new(Agent)
	}
	switch fpv.selector {
	case Agent_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Agent_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*ntt_meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fpv.selector))
	}
}

func (fpv *Agent_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Agent)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Agent_FieldTerminalPathValue' with the value under path in 'Agent'.
func (fpv *Agent_FieldTerminalPathValue) CompareWith(source *Agent) (int, bool) {
	switch fpv.selector {
	case Agent_FieldPathSelectorName:
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
	case Agent_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fpv.selector))
	}
}

func (fpv *Agent_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Agent))
}

type Agent_FieldSubPathValue struct {
	Agent_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Agent_FieldPathValue = (*Agent_FieldSubPathValue)(nil)

func (fpvs *Agent_FieldSubPathValue) AsMetadataPathValue() (ntt_meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Agent_FieldSubPathValue) SetTo(target **Agent) {
	if *target == nil {
		*target = new(Agent)
	}
	switch fpvs.Selector() {
	case Agent_FieldPathSelectorMetadata:
		fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fpvs.Selector()))
	}
}

func (fpvs *Agent_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Agent)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Agent_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Agent_FieldSubPathValue) CompareWith(source *Agent) (int, bool) {
	switch fpvs.Selector() {
	case Agent_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fpvs.Selector()))
	}
}

func (fpvs *Agent_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Agent))
}

// Agent_FieldPathArrayItemValue allows storing single item in Path-specific values for Agent according to their type
// Present only for array (repeated) types.
type Agent_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Agent_FieldPath
	ContainsValue(*Agent) bool
}

// ParseAgent_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseAgent_FieldPathArrayItemValue(pathStr, valueStr string) (Agent_FieldPathArrayItemValue, error) {
	fp, err := ParseAgent_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Agent field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Agent_FieldPathArrayItemValue), nil
}

func MustParseAgent_FieldPathArrayItemValue(pathStr, valueStr string) Agent_FieldPathArrayItemValue {
	fpaiv, err := ParseAgent_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Agent_FieldTerminalPathArrayItemValue struct {
	Agent_FieldTerminalPath
	value interface{}
}

var _ Agent_FieldPathArrayItemValue = (*Agent_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Agent as interface{}
func (fpaiv *Agent_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Agent_FieldTerminalPathArrayItemValue) GetSingle(source *Agent) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Agent_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Agent))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Agent'
func (fpaiv *Agent_FieldTerminalPathArrayItemValue) ContainsValue(source *Agent) bool {
	slice := fpaiv.Agent_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type Agent_FieldSubPathArrayItemValue struct {
	Agent_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Agent_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Agent_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (ntt_meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Agent'
func (fpaivs *Agent_FieldSubPathArrayItemValue) ContainsValue(source *Agent) bool {
	switch fpaivs.Selector() {
	case Agent_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Agent: %d", fpaivs.Selector()))
	}
}

// Agent_FieldPathArrayOfValues allows storing slice of values for Agent fields according to their type
type Agent_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Agent_FieldPath
}

func ParseAgent_FieldPathArrayOfValues(pathStr, valuesStr string) (Agent_FieldPathArrayOfValues, error) {
	fp, err := ParseAgent_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Agent field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Agent_FieldPathArrayOfValues), nil
}

func MustParseAgent_FieldPathArrayOfValues(pathStr, valuesStr string) Agent_FieldPathArrayOfValues {
	fpaov, err := ParseAgent_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Agent_FieldTerminalPathArrayOfValues struct {
	Agent_FieldTerminalPath
	values interface{}
}

var _ Agent_FieldPathArrayOfValues = (*Agent_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Agent_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Agent_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Agent_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*ntt_meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Agent_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Agent_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*ntt_meta.Meta, bool) {
	res, ok := fpaov.values.([]*ntt_meta.Meta)
	return res, ok
}

type Agent_FieldSubPathArrayOfValues struct {
	Agent_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Agent_FieldPathArrayOfValues = (*Agent_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Agent_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Agent_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (ntt_meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(ntt_meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
