// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: checker.proto

package checkerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CheckerServiceClient is the client API for CheckerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckerServiceClient interface {
	CheckPermission(ctx context.Context, in *CheckPermissionReq, opts ...grpc.CallOption) (*CheckResp, error)
}

type checkerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckerServiceClient(cc grpc.ClientConnInterface) CheckerServiceClient {
	return &checkerServiceClient{cc}
}

func (c *checkerServiceClient) CheckPermission(ctx context.Context, in *CheckPermissionReq, opts ...grpc.CallOption) (*CheckResp, error) {
	out := new(CheckResp)
	err := c.cc.Invoke(ctx, "/checkerpb.CheckerService/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServiceServer is the server API for CheckerService service.
// All implementations must embed UnimplementedCheckerServiceServer
// for forward compatibility
type CheckerServiceServer interface {
	CheckPermission(context.Context, *CheckPermissionReq) (*CheckResp, error)
	mustEmbedUnimplementedCheckerServiceServer()
}

// UnimplementedCheckerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckerServiceServer struct {
}

func (UnimplementedCheckerServiceServer) CheckPermission(context.Context, *CheckPermissionReq) (*CheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedCheckerServiceServer) mustEmbedUnimplementedCheckerServiceServer() {}

// UnsafeCheckerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckerServiceServer will
// result in compilation errors.
type UnsafeCheckerServiceServer interface {
	mustEmbedUnimplementedCheckerServiceServer()
}

func RegisterCheckerServiceServer(s grpc.ServiceRegistrar, srv CheckerServiceServer) {
	s.RegisterService(&CheckerService_ServiceDesc, srv)
}

func _CheckerService_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServiceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkerpb.CheckerService/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServiceServer).CheckPermission(ctx, req.(*CheckPermissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckerService_ServiceDesc is the grpc.ServiceDesc for CheckerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "checkerpb.CheckerService",
	HandlerType: (*CheckerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPermission",
			Handler:    _CheckerService_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checker.proto",
}
