// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/device_group.proto
// DO NOT EDIT!!!

package device_group

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
	syncing_meta "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/syncing_meta"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &policy.Policy{}
	_ = &syncing_meta.SyncingMeta{}
	_ = &timestamp.Timestamp{}
)

type DeviceGroupFieldPathBuilder struct{}

func NewDeviceGroupFieldPathBuilder() DeviceGroupFieldPathBuilder {
	return DeviceGroupFieldPathBuilder{}
}
func (DeviceGroupFieldPathBuilder) Name() DeviceGroupPathSelectorName {
	return DeviceGroupPathSelectorName{}
}
func (DeviceGroupFieldPathBuilder) DisplayName() DeviceGroupPathSelectorDisplayName {
	return DeviceGroupPathSelectorDisplayName{}
}
func (DeviceGroupFieldPathBuilder) Metadata() DeviceGroupPathSelectorMetadata {
	return DeviceGroupPathSelectorMetadata{}
}

type DeviceGroupPathSelectorName struct{}

func (DeviceGroupPathSelectorName) FieldPath() *DeviceGroup_FieldTerminalPath {
	return &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorName}
}

func (s DeviceGroupPathSelectorName) WithValue(value *Name) *DeviceGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldTerminalPathValue)
}

func (s DeviceGroupPathSelectorName) WithArrayOfValues(values []*Name) *DeviceGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldTerminalPathArrayOfValues)
}

type DeviceGroupPathSelectorDisplayName struct{}

func (DeviceGroupPathSelectorDisplayName) FieldPath() *DeviceGroup_FieldTerminalPath {
	return &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorDisplayName}
}

func (s DeviceGroupPathSelectorDisplayName) WithValue(value string) *DeviceGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldTerminalPathValue)
}

func (s DeviceGroupPathSelectorDisplayName) WithArrayOfValues(values []string) *DeviceGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldTerminalPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadata struct{}

func (DeviceGroupPathSelectorMetadata) FieldPath() *DeviceGroup_FieldTerminalPath {
	return &DeviceGroup_FieldTerminalPath{selector: DeviceGroup_FieldPathSelectorMetadata}
}

func (s DeviceGroupPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *DeviceGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldTerminalPathValue)
}

func (s DeviceGroupPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *DeviceGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldTerminalPathArrayOfValues)
}

func (DeviceGroupPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{selector: DeviceGroup_FieldPathSelectorMetadata, subPath: subPath}
}

func (s DeviceGroupPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *DeviceGroup_FieldSubPathValue {
	return &DeviceGroup_FieldSubPathValue{DeviceGroup_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s DeviceGroupPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *DeviceGroup_FieldSubPathArrayOfValues {
	return &DeviceGroup_FieldSubPathArrayOfValues{DeviceGroup_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s DeviceGroupPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *DeviceGroup_FieldSubPathArrayItemValue {
	return &DeviceGroup_FieldSubPathArrayItemValue{DeviceGroup_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (DeviceGroupPathSelectorMetadata) CreateTime() DeviceGroupPathSelectorMetadataCreateTime {
	return DeviceGroupPathSelectorMetadataCreateTime{}
}

func (DeviceGroupPathSelectorMetadata) UpdateTime() DeviceGroupPathSelectorMetadataUpdateTime {
	return DeviceGroupPathSelectorMetadataUpdateTime{}
}

func (DeviceGroupPathSelectorMetadata) Uuid() DeviceGroupPathSelectorMetadataUuid {
	return DeviceGroupPathSelectorMetadataUuid{}
}

func (DeviceGroupPathSelectorMetadata) Tags() DeviceGroupPathSelectorMetadataTags {
	return DeviceGroupPathSelectorMetadataTags{}
}

func (DeviceGroupPathSelectorMetadata) Labels() DeviceGroupPathSelectorMetadataLabels {
	return DeviceGroupPathSelectorMetadataLabels{}
}

func (DeviceGroupPathSelectorMetadata) Annotations() DeviceGroupPathSelectorMetadataAnnotations {
	return DeviceGroupPathSelectorMetadataAnnotations{}
}

func (DeviceGroupPathSelectorMetadata) Generation() DeviceGroupPathSelectorMetadataGeneration {
	return DeviceGroupPathSelectorMetadataGeneration{}
}

func (DeviceGroupPathSelectorMetadata) ResourceVersion() DeviceGroupPathSelectorMetadataResourceVersion {
	return DeviceGroupPathSelectorMetadataResourceVersion{}
}

func (DeviceGroupPathSelectorMetadata) OwnerReferences() DeviceGroupPathSelectorMetadataOwnerReferences {
	return DeviceGroupPathSelectorMetadataOwnerReferences{}
}

func (DeviceGroupPathSelectorMetadata) Shards() DeviceGroupPathSelectorMetadataShards {
	return DeviceGroupPathSelectorMetadataShards{}
}

func (DeviceGroupPathSelectorMetadata) Syncing() DeviceGroupPathSelectorMetadataSyncing {
	return DeviceGroupPathSelectorMetadataSyncing{}
}

type DeviceGroupPathSelectorMetadataCreateTime struct{}

func (DeviceGroupPathSelectorMetadataCreateTime) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataUpdateTime struct{}

func (DeviceGroupPathSelectorMetadataUpdateTime) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataUuid struct{}

func (DeviceGroupPathSelectorMetadataUuid) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataUuid) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataUuid) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataTags struct{}

func (DeviceGroupPathSelectorMetadataTags) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataTags) WithValue(value []string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (s DeviceGroupPathSelectorMetadataTags) WithItemValue(value string) *DeviceGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*DeviceGroup_FieldSubPathArrayItemValue)
}

type DeviceGroupPathSelectorMetadataLabels struct{}

func (DeviceGroupPathSelectorMetadataLabels) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataLabels) WithValue(value map[string]string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (DeviceGroupPathSelectorMetadataLabels) WithKey(key string) DeviceGroupMapPathSelectorMetadataLabels {
	return DeviceGroupMapPathSelectorMetadataLabels{key: key}
}

type DeviceGroupMapPathSelectorMetadataLabels struct {
	key string
}

func (s DeviceGroupMapPathSelectorMetadataLabels) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s DeviceGroupMapPathSelectorMetadataLabels) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataAnnotations struct{}

func (DeviceGroupPathSelectorMetadataAnnotations) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataAnnotations) WithValue(value map[string]string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (DeviceGroupPathSelectorMetadataAnnotations) WithKey(key string) DeviceGroupMapPathSelectorMetadataAnnotations {
	return DeviceGroupMapPathSelectorMetadataAnnotations{key: key}
}

type DeviceGroupMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s DeviceGroupMapPathSelectorMetadataAnnotations) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s DeviceGroupMapPathSelectorMetadataAnnotations) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataGeneration struct{}

func (DeviceGroupPathSelectorMetadataGeneration) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataGeneration) WithValue(value int64) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataResourceVersion struct{}

func (DeviceGroupPathSelectorMetadataResourceVersion) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataResourceVersion) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataOwnerReferences struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferences) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *DeviceGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*DeviceGroup_FieldSubPathArrayItemValue)
}

func (DeviceGroupPathSelectorMetadataOwnerReferences) ApiVersion() DeviceGroupPathSelectorMetadataOwnerReferencesApiVersion {
	return DeviceGroupPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (DeviceGroupPathSelectorMetadataOwnerReferences) Kind() DeviceGroupPathSelectorMetadataOwnerReferencesKind {
	return DeviceGroupPathSelectorMetadataOwnerReferencesKind{}
}

func (DeviceGroupPathSelectorMetadataOwnerReferences) Name() DeviceGroupPathSelectorMetadataOwnerReferencesName {
	return DeviceGroupPathSelectorMetadataOwnerReferencesName{}
}

func (DeviceGroupPathSelectorMetadataOwnerReferences) Uid() DeviceGroupPathSelectorMetadataOwnerReferencesUid {
	return DeviceGroupPathSelectorMetadataOwnerReferencesUid{}
}

func (DeviceGroupPathSelectorMetadataOwnerReferences) Controller() DeviceGroupPathSelectorMetadataOwnerReferencesController {
	return DeviceGroupPathSelectorMetadataOwnerReferencesController{}
}

func (DeviceGroupPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() DeviceGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return DeviceGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type DeviceGroupPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataOwnerReferencesKind struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferencesKind) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataOwnerReferencesName struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferencesName) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesName) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataOwnerReferencesUid struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferencesUid) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataOwnerReferencesController struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferencesController) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (DeviceGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataShards struct{}

func (DeviceGroupPathSelectorMetadataShards) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataShards) WithValue(value map[string]int64) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (DeviceGroupPathSelectorMetadataShards) WithKey(key string) DeviceGroupMapPathSelectorMetadataShards {
	return DeviceGroupMapPathSelectorMetadataShards{key: key}
}

type DeviceGroupMapPathSelectorMetadataShards struct {
	key string
}

func (s DeviceGroupMapPathSelectorMetadataShards) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s DeviceGroupMapPathSelectorMetadataShards) WithValue(value int64) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataSyncing struct{}

func (DeviceGroupPathSelectorMetadataSyncing) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataSyncing) WithValue(value *syncing_meta.SyncingMeta) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataSyncing) WithArrayOfValues(values []*syncing_meta.SyncingMeta) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (DeviceGroupPathSelectorMetadataSyncing) OwningRegion() DeviceGroupPathSelectorMetadataSyncingOwningRegion {
	return DeviceGroupPathSelectorMetadataSyncingOwningRegion{}
}

func (DeviceGroupPathSelectorMetadataSyncing) Regions() DeviceGroupPathSelectorMetadataSyncingRegions {
	return DeviceGroupPathSelectorMetadataSyncingRegions{}
}

type DeviceGroupPathSelectorMetadataSyncingOwningRegion struct{}

func (DeviceGroupPathSelectorMetadataSyncingOwningRegion) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

type DeviceGroupPathSelectorMetadataSyncingRegions struct{}

func (DeviceGroupPathSelectorMetadataSyncingRegions) FieldPath() *DeviceGroup_FieldSubPath {
	return &DeviceGroup_FieldSubPath{
		selector: DeviceGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s DeviceGroupPathSelectorMetadataSyncingRegions) WithValue(value []string) *DeviceGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceGroup_FieldSubPathValue)
}

func (s DeviceGroupPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *DeviceGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceGroup_FieldSubPathArrayOfValues)
}

func (s DeviceGroupPathSelectorMetadataSyncingRegions) WithItemValue(value string) *DeviceGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*DeviceGroup_FieldSubPathArrayItemValue)
}
