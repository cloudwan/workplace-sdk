// Code generated by protoc-gen-goten-client
// API: DeviceGroupService
// DO NOT EDIT!!!

package device_group_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	device_group "github.com/cloudwan/workplace-sdk/resources/v1alpha2/device_group"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = context.Context(nil)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &empty.Empty{}
	_ = &device_group.DeviceGroup{}
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceGroupServiceClient is the client API for DeviceGroupService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceGroupServiceClient interface {
	GetDeviceGroup(ctx context.Context, in *GetDeviceGroupRequest, opts ...grpc.CallOption) (*device_group.DeviceGroup, error)
	BatchGetDeviceGroups(ctx context.Context, in *BatchGetDeviceGroupsRequest, opts ...grpc.CallOption) (*BatchGetDeviceGroupsResponse, error)
	ListDeviceGroups(ctx context.Context, in *ListDeviceGroupsRequest, opts ...grpc.CallOption) (*ListDeviceGroupsResponse, error)
	WatchDeviceGroup(ctx context.Context, in *WatchDeviceGroupRequest, opts ...grpc.CallOption) (WatchDeviceGroupClientStream, error)
	WatchDeviceGroups(ctx context.Context, in *WatchDeviceGroupsRequest, opts ...grpc.CallOption) (WatchDeviceGroupsClientStream, error)
	CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*device_group.DeviceGroup, error)
	UpdateDeviceGroup(ctx context.Context, in *UpdateDeviceGroupRequest, opts ...grpc.CallOption) (*device_group.DeviceGroup, error)
	DeleteDeviceGroup(ctx context.Context, in *DeleteDeviceGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewDeviceGroupServiceClient(cc grpc.ClientConnInterface) DeviceGroupServiceClient {
	return &client{cc}
}

func (c *client) GetDeviceGroup(ctx context.Context, in *GetDeviceGroupRequest, opts ...grpc.CallOption) (*device_group.DeviceGroup, error) {
	out := new(device_group.DeviceGroup)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.DeviceGroupService/GetDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetDeviceGroups(ctx context.Context, in *BatchGetDeviceGroupsRequest, opts ...grpc.CallOption) (*BatchGetDeviceGroupsResponse, error) {
	out := new(BatchGetDeviceGroupsResponse)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.DeviceGroupService/BatchGetDeviceGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListDeviceGroups(ctx context.Context, in *ListDeviceGroupsRequest, opts ...grpc.CallOption) (*ListDeviceGroupsResponse, error) {
	out := new(ListDeviceGroupsResponse)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.DeviceGroupService/ListDeviceGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchDeviceGroup(ctx context.Context, in *WatchDeviceGroupRequest, opts ...grpc.CallOption) (WatchDeviceGroupClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeviceGroup",
			ServerStreams: true,
		},
		"/ntt.workplace.v1alpha2.DeviceGroupService/WatchDeviceGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeviceGroupWatchDeviceGroupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeviceGroupClientStream interface {
	Recv() (*WatchDeviceGroupResponse, error)
	grpc.ClientStream
}

type watchDeviceGroupWatchDeviceGroupClient struct {
	grpc.ClientStream
}

func (x *watchDeviceGroupWatchDeviceGroupClient) Recv() (*WatchDeviceGroupResponse, error) {
	m := new(WatchDeviceGroupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchDeviceGroups(ctx context.Context, in *WatchDeviceGroupsRequest, opts ...grpc.CallOption) (WatchDeviceGroupsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeviceGroups",
			ServerStreams: true,
		},
		"/ntt.workplace.v1alpha2.DeviceGroupService/WatchDeviceGroups", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeviceGroupsWatchDeviceGroupsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeviceGroupsClientStream interface {
	Recv() (*WatchDeviceGroupsResponse, error)
	grpc.ClientStream
}

type watchDeviceGroupsWatchDeviceGroupsClient struct {
	grpc.ClientStream
}

func (x *watchDeviceGroupsWatchDeviceGroupsClient) Recv() (*WatchDeviceGroupsResponse, error) {
	m := new(WatchDeviceGroupsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*device_group.DeviceGroup, error) {
	out := new(device_group.DeviceGroup)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.DeviceGroupService/CreateDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateDeviceGroup(ctx context.Context, in *UpdateDeviceGroupRequest, opts ...grpc.CallOption) (*device_group.DeviceGroup, error) {
	out := new(device_group.DeviceGroup)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.DeviceGroupService/UpdateDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteDeviceGroup(ctx context.Context, in *DeleteDeviceGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.DeviceGroupService/DeleteDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
