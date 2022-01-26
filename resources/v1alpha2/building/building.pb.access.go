// Code generated by protoc-gen-goten-resource
// Resource: Building
// DO NOT EDIT!!!

package building

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	site "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"
)

// ensure the imports are used
var (
	_ = context.Context(nil)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &workplace_common.BBox{}
	_ = &site.Site{}
)

type BuildingAccess interface {
	GetBuilding(context.Context, *GetQuery) (*Building, error)
	BatchGetBuildings(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryBuildings(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchBuilding(context.Context, *GetQuery, func(*BuildingChange) error) error
	WatchBuildings(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveBuilding(context.Context, *Building, ...gotenresource.SaveOption) error
	DeleteBuilding(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	BuildingAccess
}

func AsAnyCastAccess(access BuildingAccess) gotenresource.Access {
	return &anyCastAccess{BuildingAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asBuildingQuery, ok := q.(*GetQuery); ok {
		return a.GetBuilding(ctx, asBuildingQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Building, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asBuildingQuery, ok := q.(*ListQuery); ok {
		return a.QueryBuildings(ctx, asBuildingQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Building, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Building")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asBuildingQuery, ok := q.(*GetQuery); ok {
		return a.WatchBuilding(ctx, asBuildingQuery, func(change *BuildingChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Building, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asBuildingQuery, ok := q.(*WatchQuery); ok {
		return a.WatchBuildings(ctx, asBuildingQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Building, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asBuildingRes, ok := res.(*Building); ok {
		return a.SaveBuilding(ctx, asBuildingRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Building, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asBuildingRef, ok := ref.(*Reference); ok {
		return a.DeleteBuilding(ctx, asBuildingRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Building, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	buildingRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asBuildingRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Building, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			buildingRefs = append(buildingRefs, asBuildingRef)
		}
	}
	return a.BatchGetBuildings(ctx, buildingRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
