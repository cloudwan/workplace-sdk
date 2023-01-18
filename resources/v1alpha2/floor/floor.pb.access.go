// Code generated by protoc-gen-goten-resource
// Resource: Floor
// DO NOT EDIT!!!

package floor

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
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha2/building"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
)

// ensure the imports are used
var (
	_ = new(context.Context)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &building.Building{}
	_ = &workplace_common.BBox{}
)

type FloorAccess interface {
	GetFloor(context.Context, *GetQuery) (*Floor, error)
	BatchGetFloors(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryFloors(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchFloor(context.Context, *GetQuery, func(*FloorChange) error) error
	WatchFloors(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveFloor(context.Context, *Floor, ...gotenresource.SaveOption) error
	DeleteFloor(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	FloorAccess
}

func AsAnyCastAccess(access FloorAccess) gotenresource.Access {
	return &anyCastAccess{FloorAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asFloorQuery, ok := q.(*GetQuery); ok {
		return a.GetFloor(ctx, asFloorQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Floor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asFloorQuery, ok := q.(*ListQuery); ok {
		return a.QueryFloors(ctx, asFloorQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Floor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Floor")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asFloorQuery, ok := q.(*GetQuery); ok {
		return a.WatchFloor(ctx, asFloorQuery, func(change *FloorChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Floor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asFloorQuery, ok := q.(*WatchQuery); ok {
		return a.WatchFloors(ctx, asFloorQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Floor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asFloorRes, ok := res.(*Floor); ok {
		return a.SaveFloor(ctx, asFloorRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Floor, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asFloorRef, ok := ref.(*Reference); ok {
		return a.DeleteFloor(ctx, asFloorRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Floor, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	floorRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asFloorRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Floor, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			floorRefs = append(floorRefs, asFloorRef)
		}
	}
	return a.BatchGetFloors(ctx, floorRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
