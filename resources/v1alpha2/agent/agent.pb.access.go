// Code generated by protoc-gen-goten-resource
// Resource: Agent
// DO NOT EDIT!!!

package agent

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
)

type AgentAccess interface {
	GetAgent(context.Context, *GetQuery) (*Agent, error)
	BatchGetAgents(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryAgents(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchAgent(context.Context, *GetQuery, func(*AgentChange) error) error
	WatchAgents(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveAgent(context.Context, *Agent, ...gotenresource.SaveOption) error
	DeleteAgent(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	AgentAccess
}

func AsAnyCastAccess(access AgentAccess) gotenresource.Access {
	return &anyCastAccess{AgentAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asAgentQuery, ok := q.(*GetQuery); ok {
		return a.GetAgent(ctx, asAgentQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Agent, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asAgentQuery, ok := q.(*ListQuery); ok {
		return a.QueryAgents(ctx, asAgentQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Agent, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Agent")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asAgentQuery, ok := q.(*GetQuery); ok {
		return a.WatchAgent(ctx, asAgentQuery, func(change *AgentChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Agent, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asAgentQuery, ok := q.(*WatchQuery); ok {
		return a.WatchAgents(ctx, asAgentQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Agent, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asAgentRes, ok := res.(*Agent); ok {
		return a.SaveAgent(ctx, asAgentRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Agent, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asAgentRef, ok := ref.(*Reference); ok {
		return a.DeleteAgent(ctx, asAgentRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Agent, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	agentRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asAgentRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Agent, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			agentRefs = append(agentRefs, asAgentRef)
		}
	}
	return a.BatchGetAgents(ctx, agentRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
