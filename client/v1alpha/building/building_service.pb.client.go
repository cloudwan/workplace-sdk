// Code generated by protoc-gen-goten-client
// API: BuildingService
// DO NOT EDIT!!!

package building_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	building "github.com/cloudwan/workplace-sdk/resources/v1alpha/building"
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
	_ = &building.Building{}
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

// BuildingServiceClient is the client API for BuildingService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildingServiceClient interface {
	GetBuilding(ctx context.Context, in *GetBuildingRequest, opts ...grpc.CallOption) (*building.Building, error)
	BatchGetBuildings(ctx context.Context, in *BatchGetBuildingsRequest, opts ...grpc.CallOption) (*BatchGetBuildingsResponse, error)
	ListBuildings(ctx context.Context, in *ListBuildingsRequest, opts ...grpc.CallOption) (*ListBuildingsResponse, error)
	WatchBuilding(ctx context.Context, in *WatchBuildingRequest, opts ...grpc.CallOption) (WatchBuildingClientStream, error)
	WatchBuildings(ctx context.Context, in *WatchBuildingsRequest, opts ...grpc.CallOption) (WatchBuildingsClientStream, error)
	CreateBuilding(ctx context.Context, in *CreateBuildingRequest, opts ...grpc.CallOption) (*building.Building, error)
	UpdateBuilding(ctx context.Context, in *UpdateBuildingRequest, opts ...grpc.CallOption) (*building.Building, error)
	DeleteBuilding(ctx context.Context, in *DeleteBuildingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewBuildingServiceClient(cc grpc.ClientConnInterface) BuildingServiceClient {
	return &client{cc}
}

func (c *client) GetBuilding(ctx context.Context, in *GetBuildingRequest, opts ...grpc.CallOption) (*building.Building, error) {
	out := new(building.Building)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha.BuildingService/GetBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetBuildings(ctx context.Context, in *BatchGetBuildingsRequest, opts ...grpc.CallOption) (*BatchGetBuildingsResponse, error) {
	out := new(BatchGetBuildingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha.BuildingService/BatchGetBuildings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListBuildings(ctx context.Context, in *ListBuildingsRequest, opts ...grpc.CallOption) (*ListBuildingsResponse, error) {
	out := new(ListBuildingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha.BuildingService/ListBuildings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchBuilding(ctx context.Context, in *WatchBuildingRequest, opts ...grpc.CallOption) (WatchBuildingClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchBuilding",
			ServerStreams: true,
		},
		"/ntt.workplace.v1alpha.BuildingService/WatchBuilding", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchBuildingWatchBuildingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchBuildingClientStream interface {
	Recv() (*WatchBuildingResponse, error)
	grpc.ClientStream
}

type watchBuildingWatchBuildingClient struct {
	grpc.ClientStream
}

func (x *watchBuildingWatchBuildingClient) Recv() (*WatchBuildingResponse, error) {
	m := new(WatchBuildingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchBuildings(ctx context.Context, in *WatchBuildingsRequest, opts ...grpc.CallOption) (WatchBuildingsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchBuildings",
			ServerStreams: true,
		},
		"/ntt.workplace.v1alpha.BuildingService/WatchBuildings", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchBuildingsWatchBuildingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchBuildingsClientStream interface {
	Recv() (*WatchBuildingsResponse, error)
	grpc.ClientStream
}

type watchBuildingsWatchBuildingsClient struct {
	grpc.ClientStream
}

func (x *watchBuildingsWatchBuildingsClient) Recv() (*WatchBuildingsResponse, error) {
	m := new(WatchBuildingsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateBuilding(ctx context.Context, in *CreateBuildingRequest, opts ...grpc.CallOption) (*building.Building, error) {
	out := new(building.Building)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha.BuildingService/CreateBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateBuilding(ctx context.Context, in *UpdateBuildingRequest, opts ...grpc.CallOption) (*building.Building, error) {
	out := new(building.Building)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha.BuildingService/UpdateBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteBuilding(ctx context.Context, in *DeleteBuildingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.workplace.v1alpha.BuildingService/DeleteBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}