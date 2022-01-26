// Code generated by protoc-gen-goten-resource
// Resource: Site
// DO NOT EDIT!!!

package site

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	workplace_common "github.com/cloudwan/workplace-sdk/resources/v1alpha2/common"
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
	_ = &iam_project.Project{}
	_ = &workplace_common.BBox{}
)

type SiteAccess interface {
	GetSite(context.Context, *GetQuery) (*Site, error)
	BatchGetSites(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QuerySites(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchSite(context.Context, *GetQuery, func(*SiteChange) error) error
	WatchSites(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveSite(context.Context, *Site, ...gotenresource.SaveOption) error
	DeleteSite(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	SiteAccess
}

func AsAnyCastAccess(access SiteAccess) gotenresource.Access {
	return &anyCastAccess{SiteAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asSiteQuery, ok := q.(*GetQuery); ok {
		return a.GetSite(ctx, asSiteQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Site, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asSiteQuery, ok := q.(*ListQuery); ok {
		return a.QuerySites(ctx, asSiteQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Site, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Site")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asSiteQuery, ok := q.(*GetQuery); ok {
		return a.WatchSite(ctx, asSiteQuery, func(change *SiteChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Site, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asSiteQuery, ok := q.(*WatchQuery); ok {
		return a.WatchSites(ctx, asSiteQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Site, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asSiteRes, ok := res.(*Site); ok {
		return a.SaveSite(ctx, asSiteRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Site, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asSiteRef, ok := ref.(*Reference); ok {
		return a.DeleteSite(ctx, asSiteRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Site, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	siteRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asSiteRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Site, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			siteRefs = append(siteRefs, asSiteRef)
		}
	}
	return a.BatchGetSites(ctx, siteRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
