// Code generated by protoc-gen-goten-object
// File: workplace/proto/v1alpha2/agent.proto
// DO NOT EDIT!!!

package agent

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &meta_service.Service{}
	_ = &timestamp.Timestamp{}
)

type AgentFieldPathBuilder struct{}

func NewAgentFieldPathBuilder() AgentFieldPathBuilder {
	return AgentFieldPathBuilder{}
}
func (AgentFieldPathBuilder) Name() AgentPathSelectorName {
	return AgentPathSelectorName{}
}
func (AgentFieldPathBuilder) Metadata() AgentPathSelectorMetadata {
	return AgentPathSelectorMetadata{}
}

type AgentPathSelectorName struct{}

func (AgentPathSelectorName) FieldPath() *Agent_FieldTerminalPath {
	return &Agent_FieldTerminalPath{selector: Agent_FieldPathSelectorName}
}

func (s AgentPathSelectorName) WithValue(value *Name) *Agent_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldTerminalPathValue)
}

func (s AgentPathSelectorName) WithArrayOfValues(values []*Name) *Agent_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldTerminalPathArrayOfValues)
}

type AgentPathSelectorMetadata struct{}

func (AgentPathSelectorMetadata) FieldPath() *Agent_FieldTerminalPath {
	return &Agent_FieldTerminalPath{selector: Agent_FieldPathSelectorMetadata}
}

func (s AgentPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *Agent_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldTerminalPathValue)
}

func (s AgentPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *Agent_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldTerminalPathArrayOfValues)
}

func (AgentPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *Agent_FieldSubPath {
	return &Agent_FieldSubPath{selector: Agent_FieldPathSelectorMetadata, subPath: subPath}
}

func (s AgentPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *Agent_FieldSubPathValue {
	return &Agent_FieldSubPathValue{Agent_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s AgentPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *Agent_FieldSubPathArrayOfValues {
	return &Agent_FieldSubPathArrayOfValues{Agent_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s AgentPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *Agent_FieldSubPathArrayItemValue {
	return &Agent_FieldSubPathArrayItemValue{Agent_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (AgentPathSelectorMetadata) CreateTime() AgentPathSelectorMetadataCreateTime {
	return AgentPathSelectorMetadataCreateTime{}
}

func (AgentPathSelectorMetadata) UpdateTime() AgentPathSelectorMetadataUpdateTime {
	return AgentPathSelectorMetadataUpdateTime{}
}

func (AgentPathSelectorMetadata) DeleteTime() AgentPathSelectorMetadataDeleteTime {
	return AgentPathSelectorMetadataDeleteTime{}
}

func (AgentPathSelectorMetadata) Uuid() AgentPathSelectorMetadataUuid {
	return AgentPathSelectorMetadataUuid{}
}

func (AgentPathSelectorMetadata) Tags() AgentPathSelectorMetadataTags {
	return AgentPathSelectorMetadataTags{}
}

func (AgentPathSelectorMetadata) Labels() AgentPathSelectorMetadataLabels {
	return AgentPathSelectorMetadataLabels{}
}

func (AgentPathSelectorMetadata) Annotations() AgentPathSelectorMetadataAnnotations {
	return AgentPathSelectorMetadataAnnotations{}
}

func (AgentPathSelectorMetadata) Generation() AgentPathSelectorMetadataGeneration {
	return AgentPathSelectorMetadataGeneration{}
}

func (AgentPathSelectorMetadata) ResourceVersion() AgentPathSelectorMetadataResourceVersion {
	return AgentPathSelectorMetadataResourceVersion{}
}

func (AgentPathSelectorMetadata) OwnerReferences() AgentPathSelectorMetadataOwnerReferences {
	return AgentPathSelectorMetadataOwnerReferences{}
}

func (AgentPathSelectorMetadata) Shards() AgentPathSelectorMetadataShards {
	return AgentPathSelectorMetadataShards{}
}

func (AgentPathSelectorMetadata) Syncing() AgentPathSelectorMetadataSyncing {
	return AgentPathSelectorMetadataSyncing{}
}

func (AgentPathSelectorMetadata) Lifecycle() AgentPathSelectorMetadataLifecycle {
	return AgentPathSelectorMetadataLifecycle{}
}

type AgentPathSelectorMetadataCreateTime struct{}

func (AgentPathSelectorMetadataCreateTime) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataUpdateTime struct{}

func (AgentPathSelectorMetadataUpdateTime) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataDeleteTime struct{}

func (AgentPathSelectorMetadataDeleteTime) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataDeleteTime) WithValue(value *timestamp.Timestamp) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamp.Timestamp) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataUuid struct{}

func (AgentPathSelectorMetadataUuid) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataUuid) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataTags struct{}

func (AgentPathSelectorMetadataTags) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataTags) WithValue(value []string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (s AgentPathSelectorMetadataTags) WithItemValue(value string) *Agent_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Agent_FieldSubPathArrayItemValue)
}

type AgentPathSelectorMetadataLabels struct{}

func (AgentPathSelectorMetadataLabels) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataLabels) WithValue(value map[string]string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (AgentPathSelectorMetadataLabels) WithKey(key string) AgentMapPathSelectorMetadataLabels {
	return AgentMapPathSelectorMetadataLabels{key: key}
}

type AgentMapPathSelectorMetadataLabels struct {
	key string
}

func (s AgentMapPathSelectorMetadataLabels) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s AgentMapPathSelectorMetadataLabels) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataAnnotations struct{}

func (AgentPathSelectorMetadataAnnotations) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (AgentPathSelectorMetadataAnnotations) WithKey(key string) AgentMapPathSelectorMetadataAnnotations {
	return AgentMapPathSelectorMetadataAnnotations{key: key}
}

type AgentMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s AgentMapPathSelectorMetadataAnnotations) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s AgentMapPathSelectorMetadataAnnotations) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataGeneration struct{}

func (AgentPathSelectorMetadataGeneration) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataGeneration) WithValue(value int64) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataResourceVersion struct{}

func (AgentPathSelectorMetadataResourceVersion) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataResourceVersion) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferences struct{}

func (AgentPathSelectorMetadataOwnerReferences) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (s AgentPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *Agent_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Agent_FieldSubPathArrayItemValue)
}

func (AgentPathSelectorMetadataOwnerReferences) Kind() AgentPathSelectorMetadataOwnerReferencesKind {
	return AgentPathSelectorMetadataOwnerReferencesKind{}
}

func (AgentPathSelectorMetadataOwnerReferences) Version() AgentPathSelectorMetadataOwnerReferencesVersion {
	return AgentPathSelectorMetadataOwnerReferencesVersion{}
}

func (AgentPathSelectorMetadataOwnerReferences) Name() AgentPathSelectorMetadataOwnerReferencesName {
	return AgentPathSelectorMetadataOwnerReferencesName{}
}

func (AgentPathSelectorMetadataOwnerReferences) Region() AgentPathSelectorMetadataOwnerReferencesRegion {
	return AgentPathSelectorMetadataOwnerReferencesRegion{}
}

func (AgentPathSelectorMetadataOwnerReferences) Controller() AgentPathSelectorMetadataOwnerReferencesController {
	return AgentPathSelectorMetadataOwnerReferencesController{}
}

func (AgentPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() AgentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return AgentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

func (AgentPathSelectorMetadataOwnerReferences) RequiresOwnerReference() AgentPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return AgentPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type AgentPathSelectorMetadataOwnerReferencesKind struct{}

func (AgentPathSelectorMetadataOwnerReferencesKind) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferencesVersion struct{}

func (AgentPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferencesName struct{}

func (AgentPathSelectorMetadataOwnerReferencesName) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferencesRegion struct{}

func (AgentPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferencesController struct{}

func (AgentPathSelectorMetadataOwnerReferencesController) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (AgentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (AgentPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataShards struct{}

func (AgentPathSelectorMetadataShards) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataShards) WithValue(value map[string]int64) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (AgentPathSelectorMetadataShards) WithKey(key string) AgentMapPathSelectorMetadataShards {
	return AgentMapPathSelectorMetadataShards{key: key}
}

type AgentMapPathSelectorMetadataShards struct {
	key string
}

func (s AgentMapPathSelectorMetadataShards) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s AgentMapPathSelectorMetadataShards) WithValue(value int64) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataSyncing struct{}

func (AgentPathSelectorMetadataSyncing) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (AgentPathSelectorMetadataSyncing) OwningRegion() AgentPathSelectorMetadataSyncingOwningRegion {
	return AgentPathSelectorMetadataSyncingOwningRegion{}
}

func (AgentPathSelectorMetadataSyncing) Regions() AgentPathSelectorMetadataSyncingRegions {
	return AgentPathSelectorMetadataSyncingRegions{}
}

type AgentPathSelectorMetadataSyncingOwningRegion struct{}

func (AgentPathSelectorMetadataSyncingOwningRegion) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataSyncingRegions struct{}

func (AgentPathSelectorMetadataSyncingRegions) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataSyncingRegions) WithValue(value []string) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (s AgentPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Agent_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Agent_FieldSubPathArrayItemValue)
}

type AgentPathSelectorMetadataLifecycle struct{}

func (AgentPathSelectorMetadataLifecycle) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataLifecycle) WithValue(value *ntt_meta.Lifecycle) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataLifecycle) WithArrayOfValues(values []*ntt_meta.Lifecycle) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

func (AgentPathSelectorMetadataLifecycle) State() AgentPathSelectorMetadataLifecycleState {
	return AgentPathSelectorMetadataLifecycleState{}
}

func (AgentPathSelectorMetadataLifecycle) BlockDeletion() AgentPathSelectorMetadataLifecycleBlockDeletion {
	return AgentPathSelectorMetadataLifecycleBlockDeletion{}
}

type AgentPathSelectorMetadataLifecycleState struct{}

func (AgentPathSelectorMetadataLifecycleState) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataLifecycleState) WithValue(value ntt_meta.Lifecycle_State) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataLifecycleState) WithArrayOfValues(values []ntt_meta.Lifecycle_State) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}

type AgentPathSelectorMetadataLifecycleBlockDeletion struct{}

func (AgentPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Agent_FieldSubPath {
	return &Agent_FieldSubPath{
		selector: Agent_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s AgentPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Agent_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Agent_FieldSubPathValue)
}

func (s AgentPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Agent_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Agent_FieldSubPathArrayOfValues)
}
