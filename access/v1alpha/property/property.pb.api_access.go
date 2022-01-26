// Code generated by protoc-gen-goten-access
// Resource: Property
// DO NOT EDIT!!!

package property_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	property_client "github.com/cloudwan/workplace-sdk/client/v1alpha/property"
	property "github.com/cloudwan/workplace-sdk/resources/v1alpha/property"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = codes.NotFound
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiPropertyAccess struct {
	client property_client.PropertyServiceClient
}

func NewApiPropertyAccess(client property_client.PropertyServiceClient) property.PropertyAccess {
	return &apiPropertyAccess{client: client}
}

func (a *apiPropertyAccess) GetProperty(ctx context.Context, query *property.GetQuery) (*property.Property, error) {
	request := &property_client.GetPropertyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetProperty(ctx, request)
}

func (a *apiPropertyAccess) BatchGetProperties(ctx context.Context, refs []*property.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &property_client.BatchGetPropertiesRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetProperties(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[property.Name]*property.Property, len(refs))
	for _, resolvedRes := range resp.GetProperties() {
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

func (a *apiPropertyAccess) QueryProperties(ctx context.Context, query *property.ListQuery) (*property.QueryResultSnapshot, error) {
	request := &property_client.ListPropertiesRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListProperties(ctx, request)
	if err != nil {
		return nil, err
	}
	return &property.QueryResultSnapshot{
		Properties:     resp.Properties,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiPropertyAccess) WatchProperty(ctx context.Context, query *property.GetQuery, observerCb func(*property.PropertyChange) error) error {
	request := &property_client.WatchPropertyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchProperty(ctx, request)
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

func (a *apiPropertyAccess) WatchProperties(ctx context.Context, query *property.WatchQuery, observerCb func(*property.QueryResultChange) error) error {
	request := &property_client.WatchPropertiesRequest{
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
	changesStream, initErr := a.client.WatchProperties(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &property.QueryResultChange{
			Changes:      respChange.PropertyChanges,
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

func (a *apiPropertyAccess) SaveProperty(ctx context.Context, res *property.Property, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil {
		var err error
		previousRes, err = a.GetProperty(ctx, &property.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if previousRes != nil {
		updateRequest := &property_client.UpdatePropertyRequest{
			Property: res,
		}
		_, err := a.client.UpdateProperty(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &property_client.CreatePropertyRequest{
			Property: res,
		}
		_, err := a.client.CreateProperty(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiPropertyAccess) DeleteProperty(ctx context.Context, ref *property.Reference, opts ...gotenresource.DeleteOption) error {
	request := &property_client.DeletePropertyRequest{
		Name: ref,
	}
	_, err := a.client.DeleteProperty(ctx, request)
	return err
}