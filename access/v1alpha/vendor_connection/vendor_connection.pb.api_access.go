// Code generated by protoc-gen-goten-access
// Resource: VendorConnection
// DO NOT EDIT!!!

package vendor_connection_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	vendor_connection_client "github.com/cloudwan/workplace-sdk/client/v1alpha/vendor_connection"
	vendor_connection "github.com/cloudwan/workplace-sdk/resources/v1alpha/vendor_connection"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = codes.NotFound
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiVendorConnectionAccess struct {
	client vendor_connection_client.VendorConnectionServiceClient
}

func NewApiVendorConnectionAccess(client vendor_connection_client.VendorConnectionServiceClient) vendor_connection.VendorConnectionAccess {
	return &apiVendorConnectionAccess{client: client}
}

func (a *apiVendorConnectionAccess) GetVendorConnection(ctx context.Context, query *vendor_connection.GetQuery) (*vendor_connection.VendorConnection, error) {
	request := &vendor_connection_client.GetVendorConnectionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetVendorConnection(ctx, request)
}

func (a *apiVendorConnectionAccess) BatchGetVendorConnections(ctx context.Context, refs []*vendor_connection.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &vendor_connection_client.BatchGetVendorConnectionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetVendorConnections(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[vendor_connection.Name]*vendor_connection.VendorConnection, len(refs))
	for _, resolvedRes := range resp.GetVendorConnections() {
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

func (a *apiVendorConnectionAccess) QueryVendorConnections(ctx context.Context, query *vendor_connection.ListQuery) (*vendor_connection.QueryResultSnapshot, error) {
	request := &vendor_connection_client.ListVendorConnectionsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListVendorConnections(ctx, request)
	if err != nil {
		return nil, err
	}
	return &vendor_connection.QueryResultSnapshot{
		VendorConnections: resp.VendorConnections,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
	}, nil
}

func (a *apiVendorConnectionAccess) WatchVendorConnection(ctx context.Context, query *vendor_connection.GetQuery, observerCb func(*vendor_connection.VendorConnectionChange) error) error {
	request := &vendor_connection_client.WatchVendorConnectionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchVendorConnection(ctx, request)
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

func (a *apiVendorConnectionAccess) WatchVendorConnections(ctx context.Context, query *vendor_connection.WatchQuery, observerCb func(*vendor_connection.QueryResultChange) error) error {
	request := &vendor_connection_client.WatchVendorConnectionsRequest{
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
	changesStream, initErr := a.client.WatchVendorConnections(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &vendor_connection.QueryResultChange{
			Changes:      respChange.VendorConnectionChanges,
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

func (a *apiVendorConnectionAccess) SaveVendorConnection(ctx context.Context, res *vendor_connection.VendorConnection, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil {
		var err error
		previousRes, err = a.GetVendorConnection(ctx, &vendor_connection.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if previousRes != nil {
		updateRequest := &vendor_connection_client.UpdateVendorConnectionRequest{
			VendorConnection: res,
		}
		_, err := a.client.UpdateVendorConnection(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &vendor_connection_client.CreateVendorConnectionRequest{
			VendorConnection: res,
		}
		_, err := a.client.CreateVendorConnection(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiVendorConnectionAccess) DeleteVendorConnection(ctx context.Context, ref *vendor_connection.Reference, opts ...gotenresource.DeleteOption) error {
	request := &vendor_connection_client.DeleteVendorConnectionRequest{
		Name: ref,
	}
	_, err := a.client.DeleteVendorConnection(ctx, request)
	return err
}