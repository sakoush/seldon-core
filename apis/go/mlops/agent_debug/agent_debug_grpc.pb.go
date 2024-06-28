/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: mlops/agent_debug/agent_debug.proto

package agent_debug

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AgentDebugService_ReplicaStatus_FullMethodName = "/seldon.mlops.agent_debug.AgentDebugService/ReplicaStatus"
)

// AgentDebugServiceClient is the client API for AgentDebugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentDebugServiceClient interface {
	ReplicaStatus(ctx context.Context, in *ReplicaStatusRequest, opts ...grpc.CallOption) (*ReplicaStatusResponse, error)
}

type agentDebugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentDebugServiceClient(cc grpc.ClientConnInterface) AgentDebugServiceClient {
	return &agentDebugServiceClient{cc}
}

func (c *agentDebugServiceClient) ReplicaStatus(ctx context.Context, in *ReplicaStatusRequest, opts ...grpc.CallOption) (*ReplicaStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplicaStatusResponse)
	err := c.cc.Invoke(ctx, AgentDebugService_ReplicaStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentDebugServiceServer is the server API for AgentDebugService service.
// All implementations must embed UnimplementedAgentDebugServiceServer
// for forward compatibility
type AgentDebugServiceServer interface {
	ReplicaStatus(context.Context, *ReplicaStatusRequest) (*ReplicaStatusResponse, error)
	mustEmbedUnimplementedAgentDebugServiceServer()
}

// UnimplementedAgentDebugServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentDebugServiceServer struct {
}

func (UnimplementedAgentDebugServiceServer) ReplicaStatus(context.Context, *ReplicaStatusRequest) (*ReplicaStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicaStatus not implemented")
}
func (UnimplementedAgentDebugServiceServer) mustEmbedUnimplementedAgentDebugServiceServer() {}

// UnsafeAgentDebugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentDebugServiceServer will
// result in compilation errors.
type UnsafeAgentDebugServiceServer interface {
	mustEmbedUnimplementedAgentDebugServiceServer()
}

func RegisterAgentDebugServiceServer(s grpc.ServiceRegistrar, srv AgentDebugServiceServer) {
	s.RegisterService(&AgentDebugService_ServiceDesc, srv)
}

func _AgentDebugService_ReplicaStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicaStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentDebugServiceServer).ReplicaStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentDebugService_ReplicaStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentDebugServiceServer).ReplicaStatus(ctx, req.(*ReplicaStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentDebugService_ServiceDesc is the grpc.ServiceDesc for AgentDebugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentDebugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seldon.mlops.agent_debug.AgentDebugService",
	HandlerType: (*AgentDebugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicaStatus",
			Handler:    _AgentDebugService_ReplicaStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mlops/agent_debug/agent_debug.proto",
}
