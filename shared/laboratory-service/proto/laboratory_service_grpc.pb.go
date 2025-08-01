// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: laboratory_service.proto

package proto

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
	LaboratoryService_RequestLab_FullMethodName                 = "/laboratory_service.LaboratoryService/RequestLab"
	LaboratoryService_AddTestResult_FullMethodName              = "/laboratory_service.LaboratoryService/AddTestResult"
	LaboratoryService_ViewLabRequests_FullMethodName            = "/laboratory_service.LaboratoryService/ViewLabRequests"
	LaboratoryService_ViewTestResults_FullMethodName            = "/laboratory_service.LaboratoryService/ViewTestResults"
	LaboratoryService_ViewLabRequestsByPatientID_FullMethodName = "/laboratory_service.LaboratoryService/ViewLabRequestsByPatientID"
	LaboratoryService_UpdateLabRequest_FullMethodName           = "/laboratory_service.LaboratoryService/UpdateLabRequest"
	LaboratoryService_UpdateTestResult_FullMethodName           = "/laboratory_service.LaboratoryService/UpdateTestResult"
)

// LaboratoryServiceClient is the client API for LaboratoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LaboratoryServiceClient interface {
	RequestLab(ctx context.Context, in *RequestLabRequest, opts ...grpc.CallOption) (*LabResponse, error)
	AddTestResult(ctx context.Context, in *AddTestResultRequest, opts ...grpc.CallOption) (*LabResponse, error)
	ViewLabRequests(ctx context.Context, in *ViewLabRequestsRequest, opts ...grpc.CallOption) (*ViewLabRequestsResponse, error)
	ViewTestResults(ctx context.Context, in *ViewTestResultsRequest, opts ...grpc.CallOption) (*ViewTestResultsResponse, error)
	ViewLabRequestsByPatientID(ctx context.Context, in *ViewLabRequestsByPatientIDRequest, opts ...grpc.CallOption) (*ViewLabRequestsResponse, error)
	UpdateLabRequest(ctx context.Context, in *UpdateLabRequestRequest, opts ...grpc.CallOption) (*LabResponse, error)
	UpdateTestResult(ctx context.Context, in *UpdateTestResultRequest, opts ...grpc.CallOption) (*LabResponse, error)
}

type laboratoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaboratoryServiceClient(cc grpc.ClientConnInterface) LaboratoryServiceClient {
	return &laboratoryServiceClient{cc}
}

func (c *laboratoryServiceClient) RequestLab(ctx context.Context, in *RequestLabRequest, opts ...grpc.CallOption) (*LabResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LabResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_RequestLab_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laboratoryServiceClient) AddTestResult(ctx context.Context, in *AddTestResultRequest, opts ...grpc.CallOption) (*LabResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LabResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_AddTestResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laboratoryServiceClient) ViewLabRequests(ctx context.Context, in *ViewLabRequestsRequest, opts ...grpc.CallOption) (*ViewLabRequestsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ViewLabRequestsResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_ViewLabRequests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laboratoryServiceClient) ViewTestResults(ctx context.Context, in *ViewTestResultsRequest, opts ...grpc.CallOption) (*ViewTestResultsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ViewTestResultsResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_ViewTestResults_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laboratoryServiceClient) ViewLabRequestsByPatientID(ctx context.Context, in *ViewLabRequestsByPatientIDRequest, opts ...grpc.CallOption) (*ViewLabRequestsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ViewLabRequestsResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_ViewLabRequestsByPatientID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laboratoryServiceClient) UpdateLabRequest(ctx context.Context, in *UpdateLabRequestRequest, opts ...grpc.CallOption) (*LabResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LabResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_UpdateLabRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laboratoryServiceClient) UpdateTestResult(ctx context.Context, in *UpdateTestResultRequest, opts ...grpc.CallOption) (*LabResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LabResponse)
	err := c.cc.Invoke(ctx, LaboratoryService_UpdateTestResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaboratoryServiceServer is the server API for LaboratoryService service.
// All implementations must embed UnimplementedLaboratoryServiceServer
// for forward compatibility.
type LaboratoryServiceServer interface {
	RequestLab(context.Context, *RequestLabRequest) (*LabResponse, error)
	AddTestResult(context.Context, *AddTestResultRequest) (*LabResponse, error)
	ViewLabRequests(context.Context, *ViewLabRequestsRequest) (*ViewLabRequestsResponse, error)
	ViewTestResults(context.Context, *ViewTestResultsRequest) (*ViewTestResultsResponse, error)
	ViewLabRequestsByPatientID(context.Context, *ViewLabRequestsByPatientIDRequest) (*ViewLabRequestsResponse, error)
	UpdateLabRequest(context.Context, *UpdateLabRequestRequest) (*LabResponse, error)
	UpdateTestResult(context.Context, *UpdateTestResultRequest) (*LabResponse, error)
	mustEmbedUnimplementedLaboratoryServiceServer()
}

// UnimplementedLaboratoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLaboratoryServiceServer struct{}

func (UnimplementedLaboratoryServiceServer) RequestLab(context.Context, *RequestLabRequest) (*LabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLab not implemented")
}
func (UnimplementedLaboratoryServiceServer) AddTestResult(context.Context, *AddTestResultRequest) (*LabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTestResult not implemented")
}
func (UnimplementedLaboratoryServiceServer) ViewLabRequests(context.Context, *ViewLabRequestsRequest) (*ViewLabRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewLabRequests not implemented")
}
func (UnimplementedLaboratoryServiceServer) ViewTestResults(context.Context, *ViewTestResultsRequest) (*ViewTestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewTestResults not implemented")
}
func (UnimplementedLaboratoryServiceServer) ViewLabRequestsByPatientID(context.Context, *ViewLabRequestsByPatientIDRequest) (*ViewLabRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewLabRequestsByPatientID not implemented")
}
func (UnimplementedLaboratoryServiceServer) UpdateLabRequest(context.Context, *UpdateLabRequestRequest) (*LabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLabRequest not implemented")
}
func (UnimplementedLaboratoryServiceServer) UpdateTestResult(context.Context, *UpdateTestResultRequest) (*LabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestResult not implemented")
}
func (UnimplementedLaboratoryServiceServer) mustEmbedUnimplementedLaboratoryServiceServer() {}
func (UnimplementedLaboratoryServiceServer) testEmbeddedByValue()                           {}

// UnsafeLaboratoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LaboratoryServiceServer will
// result in compilation errors.
type UnsafeLaboratoryServiceServer interface {
	mustEmbedUnimplementedLaboratoryServiceServer()
}

func RegisterLaboratoryServiceServer(s grpc.ServiceRegistrar, srv LaboratoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedLaboratoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LaboratoryService_ServiceDesc, srv)
}

func _LaboratoryService_RequestLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).RequestLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_RequestLab_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).RequestLab(ctx, req.(*RequestLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaboratoryService_AddTestResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTestResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).AddTestResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_AddTestResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).AddTestResult(ctx, req.(*AddTestResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaboratoryService_ViewLabRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewLabRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).ViewLabRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_ViewLabRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).ViewLabRequests(ctx, req.(*ViewLabRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaboratoryService_ViewTestResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewTestResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).ViewTestResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_ViewTestResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).ViewTestResults(ctx, req.(*ViewTestResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaboratoryService_ViewLabRequestsByPatientID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewLabRequestsByPatientIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).ViewLabRequestsByPatientID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_ViewLabRequestsByPatientID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).ViewLabRequestsByPatientID(ctx, req.(*ViewLabRequestsByPatientIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaboratoryService_UpdateLabRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).UpdateLabRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_UpdateLabRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).UpdateLabRequest(ctx, req.(*UpdateLabRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaboratoryService_UpdateTestResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaboratoryServiceServer).UpdateTestResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LaboratoryService_UpdateTestResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaboratoryServiceServer).UpdateTestResult(ctx, req.(*UpdateTestResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LaboratoryService_ServiceDesc is the grpc.ServiceDesc for LaboratoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LaboratoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "laboratory_service.LaboratoryService",
	HandlerType: (*LaboratoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestLab",
			Handler:    _LaboratoryService_RequestLab_Handler,
		},
		{
			MethodName: "AddTestResult",
			Handler:    _LaboratoryService_AddTestResult_Handler,
		},
		{
			MethodName: "ViewLabRequests",
			Handler:    _LaboratoryService_ViewLabRequests_Handler,
		},
		{
			MethodName: "ViewTestResults",
			Handler:    _LaboratoryService_ViewTestResults_Handler,
		},
		{
			MethodName: "ViewLabRequestsByPatientID",
			Handler:    _LaboratoryService_ViewLabRequestsByPatientID_Handler,
		},
		{
			MethodName: "UpdateLabRequest",
			Handler:    _LaboratoryService_UpdateLabRequest_Handler,
		},
		{
			MethodName: "UpdateTestResult",
			Handler:    _LaboratoryService_UpdateTestResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "laboratory_service.proto",
}
