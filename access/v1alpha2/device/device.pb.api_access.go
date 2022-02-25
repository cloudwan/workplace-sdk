// Code generated by protoc-gen-goten-access
// Resource: Device
// DO NOT EDIT!!!

package device_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	device_client "github.com/cloudwan/workplace-sdk/client/v1alpha2/device"
	device "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = grpc.ClientConnInterface(nil)
	_ = codes.NotFound
	_ = status.Status{}

	_ = gotenaccess.Watcher(nil)
	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiDeviceAccess struct {
	client device_client.DeviceServiceClient
}

func NewApiDeviceAccess(client device_client.DeviceServiceClient) device.DeviceAccess {
	return &apiDeviceAccess{client: client}
}

func (a *apiDeviceAccess) GetDevice(ctx context.Context, query *device.GetQuery) (*device.Device, error) {
	request := &device_client.GetDeviceRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetDevice(ctx, request)
}

func (a *apiDeviceAccess) BatchGetDevices(ctx context.Context, refs []*device.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &device_client.BatchGetDevicesRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetDevices(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[device.Name]*device.Device, len(refs))
	for _, resolvedRes := range resp.GetDevices() {
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

func (a *apiDeviceAccess) QueryDevices(ctx context.Context, query *device.ListQuery) (*device.QueryResultSnapshot, error) {
	request := &device_client.ListDevicesRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListDevices(ctx, request)
	if err != nil {
		return nil, err
	}
	return &device.QueryResultSnapshot{
		Devices:        resp.Devices,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiDeviceAccess) WatchDevice(ctx context.Context, query *device.GetQuery, observerCb func(*device.DeviceChange) error) error {
	request := &device_client.WatchDeviceRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchDevice(ctx, request)
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

func (a *apiDeviceAccess) WatchDevices(ctx context.Context, query *device.WatchQuery, observerCb func(*device.QueryResultChange) error) error {
	request := &device_client.WatchDevicesRequest{
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
	changesStream, initErr := a.client.WatchDevices(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &device.QueryResultChange{
			Changes:      respChange.DeviceChanges,
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

func (a *apiDeviceAccess) SaveDevice(ctx context.Context, res *device.Device, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetDevice(ctx, &device.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &device_client.UpdateDeviceRequest{
			Device: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*device.Device_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &device_client.UpdateDeviceRequest_CAS{
				ConditionalState: conditionalState.(*device.Device),
				FieldMask:        mask.(*device.Device_FieldMask),
			}
		}
		_, err := a.client.UpdateDevice(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &device_client.CreateDeviceRequest{
			Device: res,
		}
		_, err := a.client.CreateDevice(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiDeviceAccess) DeleteDevice(ctx context.Context, ref *device.Reference, opts ...gotenresource.DeleteOption) error {
	request := &device_client.DeleteDeviceRequest{
		Name: ref,
	}
	_, err := a.client.DeleteDevice(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(device.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return device.AsAnyCastAccess(NewApiDeviceAccess(device_client.NewDeviceServiceClient(cc)))
	})
}
