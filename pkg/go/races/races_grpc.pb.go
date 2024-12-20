// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: races/races.proto

package races

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RaceService_GetRaceState_FullMethodName = "/races.RaceService/GetRaceState"
)

// RaceServiceClient is the client API for RaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaceServiceClient interface {
	GetRaceState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RaceState, error)
}

type raceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaceServiceClient(cc grpc.ClientConnInterface) RaceServiceClient {
	return &raceServiceClient{cc}
}

func (c *raceServiceClient) GetRaceState(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RaceState, error) {
	out := new(RaceState)
	err := c.cc.Invoke(ctx, RaceService_GetRaceState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaceServiceServer is the server API for RaceService service.
// All implementations should embed UnimplementedRaceServiceServer
// for forward compatibility
type RaceServiceServer interface {
	GetRaceState(context.Context, *emptypb.Empty) (*RaceState, error)
}

// UnimplementedRaceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRaceServiceServer struct {
}

func (UnimplementedRaceServiceServer) GetRaceState(context.Context, *emptypb.Empty) (*RaceState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRaceState not implemented")
}

// UnsafeRaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaceServiceServer will
// result in compilation errors.
type UnsafeRaceServiceServer interface {
	mustEmbedUnimplementedRaceServiceServer()
}

func RegisterRaceServiceServer(s grpc.ServiceRegistrar, srv RaceServiceServer) {
	s.RegisterService(&RaceService_ServiceDesc, srv)
}

func _RaceService_GetRaceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaceServiceServer).GetRaceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaceService_GetRaceState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaceServiceServer).GetRaceState(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RaceService_ServiceDesc is the grpc.ServiceDesc for RaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "races.RaceService",
	HandlerType: (*RaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRaceState",
			Handler:    _RaceService_GetRaceState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "races/races.proto",
}
