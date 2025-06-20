// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: admin.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdminService_CreateTranscriptRequest_FullMethodName  = "/admin.AdminService/CreateTranscriptRequest"
	AdminService_ViewQueue_FullMethodName                = "/admin.AdminService/ViewQueue"
	AdminService_JoinQueue_FullMethodName                = "/admin.AdminService/JoinQueue"
	AdminService_RegisterRetake_FullMethodName           = "/admin.AdminService/RegisterRetake"
	AdminService_GetSchedule_FullMethodName              = "/admin.AdminService/GetSchedule"
	AdminService_UpdateSchedule_FullMethodName           = "/admin.AdminService/UpdateSchedule"
	AdminService_SubmitCertificateRequest_FullMethodName = "/admin.AdminService/SubmitCertificateRequest"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	CreateTranscriptRequest(ctx context.Context, in *TranscriptRequest, opts ...grpc.CallOption) (*Empty, error)
	ViewQueue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QueueList, error)
	JoinQueue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*Empty, error)
	RegisterRetake(ctx context.Context, in *RetakeRequest, opts ...grpc.CallOption) (*Empty, error)
	GetSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*Empty, error)
	SubmitCertificateRequest(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*Empty, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateTranscriptRequest(ctx context.Context, in *TranscriptRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_CreateTranscriptRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ViewQueue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QueueList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueueList)
	err := c.cc.Invoke(ctx, AdminService_ViewQueue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) JoinQueue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_JoinQueue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RegisterRetake(ctx context.Context, in *RetakeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_RegisterRetake_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, AdminService_GetSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_UpdateSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) SubmitCertificateRequest(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdminService_SubmitCertificateRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility.
type AdminServiceServer interface {
	CreateTranscriptRequest(context.Context, *TranscriptRequest) (*Empty, error)
	ViewQueue(context.Context, *Empty) (*QueueList, error)
	JoinQueue(context.Context, *QueueRequest) (*Empty, error)
	RegisterRetake(context.Context, *RetakeRequest) (*Empty, error)
	GetSchedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	UpdateSchedule(context.Context, *UpdateScheduleRequest) (*Empty, error)
	SubmitCertificateRequest(context.Context, *CertificateRequest) (*Empty, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminServiceServer struct{}

func (UnimplementedAdminServiceServer) CreateTranscriptRequest(context.Context, *TranscriptRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTranscriptRequest not implemented")
}
func (UnimplementedAdminServiceServer) ViewQueue(context.Context, *Empty) (*QueueList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewQueue not implemented")
}
func (UnimplementedAdminServiceServer) JoinQueue(context.Context, *QueueRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinQueue not implemented")
}
func (UnimplementedAdminServiceServer) RegisterRetake(context.Context, *RetakeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRetake not implemented")
}
func (UnimplementedAdminServiceServer) GetSchedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedAdminServiceServer) UpdateSchedule(context.Context, *UpdateScheduleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchedule not implemented")
}
func (UnimplementedAdminServiceServer) SubmitCertificateRequest(context.Context, *CertificateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCertificateRequest not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}
func (UnimplementedAdminServiceServer) testEmbeddedByValue()                      {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_CreateTranscriptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateTranscriptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateTranscriptRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateTranscriptRequest(ctx, req.(*TranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ViewQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ViewQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_ViewQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ViewQueue(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_JoinQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).JoinQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_JoinQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).JoinQueue(ctx, req.(*QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RegisterRetake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RegisterRetake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RegisterRetake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RegisterRetake(ctx, req.(*RetakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSchedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UpdateSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateSchedule(ctx, req.(*UpdateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_SubmitCertificateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).SubmitCertificateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_SubmitCertificateRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).SubmitCertificateRequest(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTranscriptRequest",
			Handler:    _AdminService_CreateTranscriptRequest_Handler,
		},
		{
			MethodName: "ViewQueue",
			Handler:    _AdminService_ViewQueue_Handler,
		},
		{
			MethodName: "JoinQueue",
			Handler:    _AdminService_JoinQueue_Handler,
		},
		{
			MethodName: "RegisterRetake",
			Handler:    _AdminService_RegisterRetake_Handler,
		},
		{
			MethodName: "GetSchedule",
			Handler:    _AdminService_GetSchedule_Handler,
		},
		{
			MethodName: "UpdateSchedule",
			Handler:    _AdminService_UpdateSchedule_Handler,
		},
		{
			MethodName: "SubmitCertificateRequest",
			Handler:    _AdminService_SubmitCertificateRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
