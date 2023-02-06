// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: custommqttmessage/custommqttmessage.proto

package custommqttmessage

import (
	context "context"
	drivercommon "github.com/winc-link/edge-driver-proto/drivercommon"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcCustomMqttMessageClient is the client API for RpcCustomMqttMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcCustomMqttMessageClient interface {
	// 平台自定义消息推Public
	PlatformCustomPublish(ctx context.Context, in *PlatformCustomPublishRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error)
	// 平台自定义消息推Subscribe
	PlatformCustomSubscribe(ctx context.Context, in *PlatformCustomSubscribeRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error)
	// 平台自定义消息UnSubscribe
	PlatformCustomUnSubscribe(ctx context.Context, in *PlatformCustomUnSubscribeRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error)
}

type rpcCustomMqttMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcCustomMqttMessageClient(cc grpc.ClientConnInterface) RpcCustomMqttMessageClient {
	return &rpcCustomMqttMessageClient{cc}
}

func (c *rpcCustomMqttMessageClient) PlatformCustomPublish(ctx context.Context, in *PlatformCustomPublishRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error) {
	out := new(drivercommon.CommonResponse)
	err := c.cc.Invoke(ctx, "/custommqttmessage.RpcCustomMqttMessage/PlatformCustomPublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcCustomMqttMessageClient) PlatformCustomSubscribe(ctx context.Context, in *PlatformCustomSubscribeRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error) {
	out := new(drivercommon.CommonResponse)
	err := c.cc.Invoke(ctx, "/custommqttmessage.RpcCustomMqttMessage/PlatformCustomSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcCustomMqttMessageClient) PlatformCustomUnSubscribe(ctx context.Context, in *PlatformCustomUnSubscribeRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error) {
	out := new(drivercommon.CommonResponse)
	err := c.cc.Invoke(ctx, "/custommqttmessage.RpcCustomMqttMessage/PlatformCustomUnSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcCustomMqttMessageServer is the server API for RpcCustomMqttMessage service.
// All implementations must embed UnimplementedRpcCustomMqttMessageServer
// for forward compatibility
type RpcCustomMqttMessageServer interface {
	// 平台自定义消息推Public
	PlatformCustomPublish(context.Context, *PlatformCustomPublishRequest) (*drivercommon.CommonResponse, error)
	// 平台自定义消息推Subscribe
	PlatformCustomSubscribe(context.Context, *PlatformCustomSubscribeRequest) (*drivercommon.CommonResponse, error)
	// 平台自定义消息UnSubscribe
	PlatformCustomUnSubscribe(context.Context, *PlatformCustomUnSubscribeRequest) (*drivercommon.CommonResponse, error)
	mustEmbedUnimplementedRpcCustomMqttMessageServer()
}

// UnimplementedRpcCustomMqttMessageServer must be embedded to have forward compatible implementations.
type UnimplementedRpcCustomMqttMessageServer struct {
}

func (UnimplementedRpcCustomMqttMessageServer) PlatformCustomPublish(context.Context, *PlatformCustomPublishRequest) (*drivercommon.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlatformCustomPublish not implemented")
}
func (UnimplementedRpcCustomMqttMessageServer) PlatformCustomSubscribe(context.Context, *PlatformCustomSubscribeRequest) (*drivercommon.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlatformCustomSubscribe not implemented")
}
func (UnimplementedRpcCustomMqttMessageServer) PlatformCustomUnSubscribe(context.Context, *PlatformCustomUnSubscribeRequest) (*drivercommon.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlatformCustomUnSubscribe not implemented")
}
func (UnimplementedRpcCustomMqttMessageServer) mustEmbedUnimplementedRpcCustomMqttMessageServer() {}

// UnsafeRpcCustomMqttMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcCustomMqttMessageServer will
// result in compilation errors.
type UnsafeRpcCustomMqttMessageServer interface {
	mustEmbedUnimplementedRpcCustomMqttMessageServer()
}

func RegisterRpcCustomMqttMessageServer(s grpc.ServiceRegistrar, srv RpcCustomMqttMessageServer) {
	s.RegisterService(&RpcCustomMqttMessage_ServiceDesc, srv)
}

func _RpcCustomMqttMessage_PlatformCustomPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformCustomPublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcCustomMqttMessageServer).PlatformCustomPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/custommqttmessage.RpcCustomMqttMessage/PlatformCustomPublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcCustomMqttMessageServer).PlatformCustomPublish(ctx, req.(*PlatformCustomPublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcCustomMqttMessage_PlatformCustomSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformCustomSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcCustomMqttMessageServer).PlatformCustomSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/custommqttmessage.RpcCustomMqttMessage/PlatformCustomSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcCustomMqttMessageServer).PlatformCustomSubscribe(ctx, req.(*PlatformCustomSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcCustomMqttMessage_PlatformCustomUnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformCustomUnSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcCustomMqttMessageServer).PlatformCustomUnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/custommqttmessage.RpcCustomMqttMessage/PlatformCustomUnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcCustomMqttMessageServer).PlatformCustomUnSubscribe(ctx, req.(*PlatformCustomUnSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcCustomMqttMessage_ServiceDesc is the grpc.ServiceDesc for RpcCustomMqttMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcCustomMqttMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "custommqttmessage.RpcCustomMqttMessage",
	HandlerType: (*RpcCustomMqttMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlatformCustomPublish",
			Handler:    _RpcCustomMqttMessage_PlatformCustomPublish_Handler,
		},
		{
			MethodName: "PlatformCustomSubscribe",
			Handler:    _RpcCustomMqttMessage_PlatformCustomSubscribe_Handler,
		},
		{
			MethodName: "PlatformCustomUnSubscribe",
			Handler:    _RpcCustomMqttMessage_PlatformCustomUnSubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "custommqttmessage/custommqttmessage.proto",
}