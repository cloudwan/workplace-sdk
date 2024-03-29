// Code generated by protoc-gen-goten-access
// Resource: Agent
// DO NOT EDIT!!!

package agent_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	agent_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/agent"
	agent "github.com/cloudwan/workplace-sdk/resources/v1alpha2/agent"
)

var (
	_ = new(context.Context)
	_ = new(fmt.GoStringer)

	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
)

type apiAgentAccess struct {
	client agent_client.AgentServiceClient
}

func NewApiAgentAccess(client agent_client.AgentServiceClient) agent.AgentAccess {
	return &apiAgentAccess{client: client}
}

func (a *apiAgentAccess) GetAgent(ctx context.Context, query *agent.GetQuery) (*agent.Agent, error) {
	request := &agent_client.GetAgentRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAgent(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAgentAccess) BatchGetAgents(ctx context.Context, refs []*agent.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &agent_client.BatchGetAgentsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetAgents(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[agent.Name]*agent.Agent, len(refs))
	for _, resolvedRes := range resp.GetAgents() {
		resultMap[*resolvedRes.GetName()] = resolvedRes
	}
	for _, ref := range refs {
		resolvedRes := resultMap[ref.Name]
		if resolvedRes != nil {
			ref.Resolve(resolvedRes)
		}
	}
	if batchGetOpts.MustResolveAll() && len(resp.GetMissing()) > 0 {
		return status.Errorf(codes.NotFound, "Number of references not found: %d", len(resp.GetMissing()))
	}
	return nil
}

func (a *apiAgentAccess) QueryAgents(ctx context.Context, query *agent.ListQuery) (*agent.QueryResultSnapshot, error) {
	request := &agent_client.ListAgentsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListAgents(ctx, request)
	if err != nil {
		return nil, err
	}
	return &agent.QueryResultSnapshot{
		Agents:            resp.Agents,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiAgentAccess) WatchAgent(ctx context.Context, query *agent.GetQuery, observerCb func(*agent.AgentChange) error) error {
	request := &agent_client.WatchAgentRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchAgent(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		resp, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		change := resp.GetChange()
		if err := observerCb(change); err != nil {
			return err
		}
	}
}

func (a *apiAgentAccess) WatchAgents(ctx context.Context, query *agent.WatchQuery, observerCb func(*agent.QueryResultChange) error) error {
	request := &agent_client.WatchAgentsRequest{
		Filter:       query.Filter,
		FieldMask:    query.Mask,
		MaxChunkSize: int32(query.ChunkSize),
		Type:         query.WatchType,
		ResumeToken:  query.ResumeToken,
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
	changesStream, initErr := a.client.WatchAgents(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &agent.QueryResultChange{
			Changes:      respChange.AgentChanges,
			IsCurrent:    respChange.IsCurrent,
			IsHardReset:  respChange.IsHardReset,
			IsSoftReset:  respChange.IsSoftReset,
			ResumeToken:  respChange.ResumeToken,
			SnapshotSize: respChange.SnapshotSize,
		}
		if respChange.PageTokenChange != nil {
			changesWithPaging.PrevPageCursor = respChange.PageTokenChange.PrevPageToken
			changesWithPaging.NextPageCursor = respChange.PageTokenChange.NextPageToken
		}
		if err := observerCb(changesWithPaging); err != nil {
			return err
		}
	}
}

func (a *apiAgentAccess) SaveAgent(ctx context.Context, res *agent.Agent, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetAgent(ctx, &agent.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &agent_client.UpdateAgentRequest{
			Agent: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*agent.Agent_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &agent_client.UpdateAgentRequest_CAS{
				ConditionalState: conditionalState.(*agent.Agent),
				FieldMask:        mask.(*agent.Agent_FieldMask),
			}
		}
		_, err := a.client.UpdateAgent(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &agent_client.CreateAgentRequest{
			Agent: res,
		}
		_, err := a.client.CreateAgent(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiAgentAccess) DeleteAgent(ctx context.Context, ref *agent.Reference, opts ...gotenresource.DeleteOption) error {
	request := &agent_client.DeleteAgentRequest{
		Name: ref,
	}
	_, err := a.client.DeleteAgent(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(agent.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return agent.AsAnyCastAccess(NewApiAgentAccess(agent_client.NewAgentServiceClient(cc)))
	})
}
