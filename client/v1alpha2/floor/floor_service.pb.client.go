// Code generated by protoc-gen-goten-client
// API: FloorService
// DO NOT EDIT!!!

package floor_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	floor "github.com/cloudwan/workplace-sdk/resources/v1alpha2/floor"
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
	_ = &floor.Floor{}
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

// FloorServiceClient is the client API for FloorService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FloorServiceClient interface {
	GetFloor(ctx context.Context, in *GetFloorRequest, opts ...grpc.CallOption) (*floor.Floor, error)
	BatchGetFloors(ctx context.Context, in *BatchGetFloorsRequest, opts ...grpc.CallOption) (*BatchGetFloorsResponse, error)
	ListFloors(ctx context.Context, in *ListFloorsRequest, opts ...grpc.CallOption) (*ListFloorsResponse, error)
	WatchFloor(ctx context.Context, in *WatchFloorRequest, opts ...grpc.CallOption) (WatchFloorClientStream, error)
	WatchFloors(ctx context.Context, in *WatchFloorsRequest, opts ...grpc.CallOption) (WatchFloorsClientStream, error)
	CreateFloor(ctx context.Context, in *CreateFloorRequest, opts ...grpc.CallOption) (*floor.Floor, error)
	UpdateFloor(ctx context.Context, in *UpdateFloorRequest, opts ...grpc.CallOption) (*floor.Floor, error)
	DeleteFloor(ctx context.Context, in *DeleteFloorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewFloorServiceClient(cc grpc.ClientConnInterface) FloorServiceClient {
	return &client{cc}
}

func (c *client) GetFloor(ctx context.Context, in *GetFloorRequest, opts ...grpc.CallOption) (*floor.Floor, error) {
	out := new(floor.Floor)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.FloorService/GetFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetFloors(ctx context.Context, in *BatchGetFloorsRequest, opts ...grpc.CallOption) (*BatchGetFloorsResponse, error) {
	out := new(BatchGetFloorsResponse)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.FloorService/BatchGetFloors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListFloors(ctx context.Context, in *ListFloorsRequest, opts ...grpc.CallOption) (*ListFloorsResponse, error) {
	out := new(ListFloorsResponse)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.FloorService/ListFloors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchFloor(ctx context.Context, in *WatchFloorRequest, opts ...grpc.CallOption) (WatchFloorClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchFloor",
			ServerStreams: true,
		},
		"/ntt.workplace.v1alpha2.FloorService/WatchFloor", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchFloorWatchFloorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchFloorClientStream interface {
	Recv() (*WatchFloorResponse, error)
	grpc.ClientStream
}

type watchFloorWatchFloorClient struct {
	grpc.ClientStream
}

func (x *watchFloorWatchFloorClient) Recv() (*WatchFloorResponse, error) {
	m := new(WatchFloorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchFloors(ctx context.Context, in *WatchFloorsRequest, opts ...grpc.CallOption) (WatchFloorsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchFloors",
			ServerStreams: true,
		},
		"/ntt.workplace.v1alpha2.FloorService/WatchFloors", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchFloorsWatchFloorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchFloorsClientStream interface {
	Recv() (*WatchFloorsResponse, error)
	grpc.ClientStream
}

type watchFloorsWatchFloorsClient struct {
	grpc.ClientStream
}

func (x *watchFloorsWatchFloorsClient) Recv() (*WatchFloorsResponse, error) {
	m := new(WatchFloorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateFloor(ctx context.Context, in *CreateFloorRequest, opts ...grpc.CallOption) (*floor.Floor, error) {
	out := new(floor.Floor)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.FloorService/CreateFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateFloor(ctx context.Context, in *UpdateFloorRequest, opts ...grpc.CallOption) (*floor.Floor, error) {
	out := new(floor.Floor)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.FloorService/UpdateFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteFloor(ctx context.Context, in *DeleteFloorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha2.FloorService/DeleteFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
