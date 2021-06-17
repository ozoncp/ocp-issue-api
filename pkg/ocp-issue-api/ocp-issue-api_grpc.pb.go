// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_issue_api

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

// OcpIssueApiClient is the client API for OcpIssueApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpIssueApiClient interface {
	// Returns list of issues
	ListIssuesV1(ctx context.Context, in *ListIssuesV1Request, opts ...grpc.CallOption) (*ListIssuesV1Response, error)
	// Create new issue
	CreateIssueV1(ctx context.Context, in *CreateIssueV1Request, opts ...grpc.CallOption) (*CreateIssueV1Response, error)
	// Get the issue
	DescribeIssueV1(ctx context.Context, in *DescribeIssueV1Request, opts ...grpc.CallOption) (*DescribeIssueV1Response, error)
	// Update the issue
	UpdateIssueV1(ctx context.Context, in *UpdateIssueV1Request, opts ...grpc.CallOption) (*UpdateIssueV1Response, error)
	// Remove the issue
	RemoveIssueV1(ctx context.Context, in *RemoveIssueV1Request, opts ...grpc.CallOption) (*RemoveIssueV1Response, error)
	// Create some issues
	MultiCreateIssueV1(ctx context.Context, in *MultiCreateIssueV1Request, opts ...grpc.CallOption) (*MultiCreateIssueV1Response, error)
}

type ocpIssueApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpIssueApiClient(cc grpc.ClientConnInterface) OcpIssueApiClient {
	return &ocpIssueApiClient{cc}
}

func (c *ocpIssueApiClient) ListIssuesV1(ctx context.Context, in *ListIssuesV1Request, opts ...grpc.CallOption) (*ListIssuesV1Response, error) {
	out := new(ListIssuesV1Response)
	err := c.cc.Invoke(ctx, "/ocp.issue.api.OcpIssueApi/ListIssuesV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpIssueApiClient) CreateIssueV1(ctx context.Context, in *CreateIssueV1Request, opts ...grpc.CallOption) (*CreateIssueV1Response, error) {
	out := new(CreateIssueV1Response)
	err := c.cc.Invoke(ctx, "/ocp.issue.api.OcpIssueApi/CreateIssueV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpIssueApiClient) DescribeIssueV1(ctx context.Context, in *DescribeIssueV1Request, opts ...grpc.CallOption) (*DescribeIssueV1Response, error) {
	out := new(DescribeIssueV1Response)
	err := c.cc.Invoke(ctx, "/ocp.issue.api.OcpIssueApi/DescribeIssueV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpIssueApiClient) UpdateIssueV1(ctx context.Context, in *UpdateIssueV1Request, opts ...grpc.CallOption) (*UpdateIssueV1Response, error) {
	out := new(UpdateIssueV1Response)
	err := c.cc.Invoke(ctx, "/ocp.issue.api.OcpIssueApi/UpdateIssueV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpIssueApiClient) RemoveIssueV1(ctx context.Context, in *RemoveIssueV1Request, opts ...grpc.CallOption) (*RemoveIssueV1Response, error) {
	out := new(RemoveIssueV1Response)
	err := c.cc.Invoke(ctx, "/ocp.issue.api.OcpIssueApi/RemoveIssueV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpIssueApiClient) MultiCreateIssueV1(ctx context.Context, in *MultiCreateIssueV1Request, opts ...grpc.CallOption) (*MultiCreateIssueV1Response, error) {
	out := new(MultiCreateIssueV1Response)
	err := c.cc.Invoke(ctx, "/ocp.issue.api.OcpIssueApi/MultiCreateIssueV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpIssueApiServer is the server API for OcpIssueApi service.
// All implementations must embed UnimplementedOcpIssueApiServer
// for forward compatibility
type OcpIssueApiServer interface {
	// Returns list of issues
	ListIssuesV1(context.Context, *ListIssuesV1Request) (*ListIssuesV1Response, error)
	// Create new issue
	CreateIssueV1(context.Context, *CreateIssueV1Request) (*CreateIssueV1Response, error)
	// Get the issue
	DescribeIssueV1(context.Context, *DescribeIssueV1Request) (*DescribeIssueV1Response, error)
	// Update the issue
	UpdateIssueV1(context.Context, *UpdateIssueV1Request) (*UpdateIssueV1Response, error)
	// Remove the issue
	RemoveIssueV1(context.Context, *RemoveIssueV1Request) (*RemoveIssueV1Response, error)
	// Create some issues
	MultiCreateIssueV1(context.Context, *MultiCreateIssueV1Request) (*MultiCreateIssueV1Response, error)
	mustEmbedUnimplementedOcpIssueApiServer()
}

// UnimplementedOcpIssueApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpIssueApiServer struct {
}

func (UnimplementedOcpIssueApiServer) ListIssuesV1(context.Context, *ListIssuesV1Request) (*ListIssuesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssuesV1 not implemented")
}
func (UnimplementedOcpIssueApiServer) CreateIssueV1(context.Context, *CreateIssueV1Request) (*CreateIssueV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssueV1 not implemented")
}
func (UnimplementedOcpIssueApiServer) DescribeIssueV1(context.Context, *DescribeIssueV1Request) (*DescribeIssueV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeIssueV1 not implemented")
}
func (UnimplementedOcpIssueApiServer) UpdateIssueV1(context.Context, *UpdateIssueV1Request) (*UpdateIssueV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIssueV1 not implemented")
}
func (UnimplementedOcpIssueApiServer) RemoveIssueV1(context.Context, *RemoveIssueV1Request) (*RemoveIssueV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveIssueV1 not implemented")
}
func (UnimplementedOcpIssueApiServer) MultiCreateIssueV1(context.Context, *MultiCreateIssueV1Request) (*MultiCreateIssueV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateIssueV1 not implemented")
}
func (UnimplementedOcpIssueApiServer) mustEmbedUnimplementedOcpIssueApiServer() {}

// UnsafeOcpIssueApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpIssueApiServer will
// result in compilation errors.
type UnsafeOcpIssueApiServer interface {
	mustEmbedUnimplementedOcpIssueApiServer()
}

func RegisterOcpIssueApiServer(s grpc.ServiceRegistrar, srv OcpIssueApiServer) {
	s.RegisterService(&OcpIssueApi_ServiceDesc, srv)
}

func _OcpIssueApi_ListIssuesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssuesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpIssueApiServer).ListIssuesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.issue.api.OcpIssueApi/ListIssuesV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpIssueApiServer).ListIssuesV1(ctx, req.(*ListIssuesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpIssueApi_CreateIssueV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpIssueApiServer).CreateIssueV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.issue.api.OcpIssueApi/CreateIssueV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpIssueApiServer).CreateIssueV1(ctx, req.(*CreateIssueV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpIssueApi_DescribeIssueV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeIssueV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpIssueApiServer).DescribeIssueV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.issue.api.OcpIssueApi/DescribeIssueV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpIssueApiServer).DescribeIssueV1(ctx, req.(*DescribeIssueV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpIssueApi_UpdateIssueV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIssueV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpIssueApiServer).UpdateIssueV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.issue.api.OcpIssueApi/UpdateIssueV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpIssueApiServer).UpdateIssueV1(ctx, req.(*UpdateIssueV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpIssueApi_RemoveIssueV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveIssueV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpIssueApiServer).RemoveIssueV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.issue.api.OcpIssueApi/RemoveIssueV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpIssueApiServer).RemoveIssueV1(ctx, req.(*RemoveIssueV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpIssueApi_MultiCreateIssueV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateIssueV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpIssueApiServer).MultiCreateIssueV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.issue.api.OcpIssueApi/MultiCreateIssueV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpIssueApiServer).MultiCreateIssueV1(ctx, req.(*MultiCreateIssueV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpIssueApi_ServiceDesc is the grpc.ServiceDesc for OcpIssueApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpIssueApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.issue.api.OcpIssueApi",
	HandlerType: (*OcpIssueApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIssuesV1",
			Handler:    _OcpIssueApi_ListIssuesV1_Handler,
		},
		{
			MethodName: "CreateIssueV1",
			Handler:    _OcpIssueApi_CreateIssueV1_Handler,
		},
		{
			MethodName: "DescribeIssueV1",
			Handler:    _OcpIssueApi_DescribeIssueV1_Handler,
		},
		{
			MethodName: "UpdateIssueV1",
			Handler:    _OcpIssueApi_UpdateIssueV1_Handler,
		},
		{
			MethodName: "RemoveIssueV1",
			Handler:    _OcpIssueApi_RemoveIssueV1_Handler,
		},
		{
			MethodName: "MultiCreateIssueV1",
			Handler:    _OcpIssueApi_MultiCreateIssueV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-issue-api/ocp-issue-api.proto",
}
