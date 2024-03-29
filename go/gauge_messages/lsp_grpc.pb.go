//----------------------------------------------------------------
//  Copyright (c) ThoughtWorks, Inc.
//  Licensed under the Apache License, Version 2.0
//  See LICENSE in the project root for license information.
//----------------------------------------------------------------

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: lsp.proto

package gauge_messages

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

const (
	LspService_GetStepNames_FullMethodName           = "/gauge.messages.lspService/GetStepNames"
	LspService_CacheFile_FullMethodName              = "/gauge.messages.lspService/CacheFile"
	LspService_GetStepPositions_FullMethodName       = "/gauge.messages.lspService/GetStepPositions"
	LspService_GetImplementationFiles_FullMethodName = "/gauge.messages.lspService/GetImplementationFiles"
	LspService_ImplementStub_FullMethodName          = "/gauge.messages.lspService/ImplementStub"
	LspService_ValidateStep_FullMethodName           = "/gauge.messages.lspService/ValidateStep"
	LspService_Refactor_FullMethodName               = "/gauge.messages.lspService/Refactor"
	LspService_GetStepName_FullMethodName            = "/gauge.messages.lspService/GetStepName"
	LspService_GetGlobPatterns_FullMethodName        = "/gauge.messages.lspService/GetGlobPatterns"
	LspService_KillProcess_FullMethodName            = "/gauge.messages.lspService/KillProcess"
)

// LspServiceClient is the client API for LspService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LspServiceClient interface {
	// Deprecated: Do not use.
	GetStepNames(ctx context.Context, in *StepNamesRequest, opts ...grpc.CallOption) (*StepNamesResponse, error)
	// Deprecated: Do not use.
	CacheFile(ctx context.Context, in *CacheFileRequest, opts ...grpc.CallOption) (*Empty, error)
	// Deprecated: Do not use.
	GetStepPositions(ctx context.Context, in *StepPositionsRequest, opts ...grpc.CallOption) (*StepPositionsResponse, error)
	// Deprecated: Do not use.
	GetImplementationFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ImplementationFileListResponse, error)
	// Deprecated: Do not use.
	ImplementStub(ctx context.Context, in *StubImplementationCodeRequest, opts ...grpc.CallOption) (*FileDiff, error)
	// Deprecated: Do not use.
	ValidateStep(ctx context.Context, in *StepValidateRequest, opts ...grpc.CallOption) (*StepValidateResponse, error)
	// Deprecated: Do not use.
	Refactor(ctx context.Context, in *RefactorRequest, opts ...grpc.CallOption) (*RefactorResponse, error)
	// Deprecated: Do not use.
	GetStepName(ctx context.Context, in *StepNameRequest, opts ...grpc.CallOption) (*StepNameResponse, error)
	// Deprecated: Do not use.
	GetGlobPatterns(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ImplementationFileGlobPatternResponse, error)
	// Deprecated: Do not use.
	KillProcess(ctx context.Context, in *KillProcessRequest, opts ...grpc.CallOption) (*Empty, error)
}

type lspServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLspServiceClient(cc grpc.ClientConnInterface) LspServiceClient {
	return &lspServiceClient{cc}
}

// Deprecated: Do not use.
func (c *lspServiceClient) GetStepNames(ctx context.Context, in *StepNamesRequest, opts ...grpc.CallOption) (*StepNamesResponse, error) {
	out := new(StepNamesResponse)
	err := c.cc.Invoke(ctx, LspService_GetStepNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) CacheFile(ctx context.Context, in *CacheFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, LspService_CacheFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) GetStepPositions(ctx context.Context, in *StepPositionsRequest, opts ...grpc.CallOption) (*StepPositionsResponse, error) {
	out := new(StepPositionsResponse)
	err := c.cc.Invoke(ctx, LspService_GetStepPositions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) GetImplementationFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ImplementationFileListResponse, error) {
	out := new(ImplementationFileListResponse)
	err := c.cc.Invoke(ctx, LspService_GetImplementationFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) ImplementStub(ctx context.Context, in *StubImplementationCodeRequest, opts ...grpc.CallOption) (*FileDiff, error) {
	out := new(FileDiff)
	err := c.cc.Invoke(ctx, LspService_ImplementStub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) ValidateStep(ctx context.Context, in *StepValidateRequest, opts ...grpc.CallOption) (*StepValidateResponse, error) {
	out := new(StepValidateResponse)
	err := c.cc.Invoke(ctx, LspService_ValidateStep_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) Refactor(ctx context.Context, in *RefactorRequest, opts ...grpc.CallOption) (*RefactorResponse, error) {
	out := new(RefactorResponse)
	err := c.cc.Invoke(ctx, LspService_Refactor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) GetStepName(ctx context.Context, in *StepNameRequest, opts ...grpc.CallOption) (*StepNameResponse, error) {
	out := new(StepNameResponse)
	err := c.cc.Invoke(ctx, LspService_GetStepName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) GetGlobPatterns(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ImplementationFileGlobPatternResponse, error) {
	out := new(ImplementationFileGlobPatternResponse)
	err := c.cc.Invoke(ctx, LspService_GetGlobPatterns_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *lspServiceClient) KillProcess(ctx context.Context, in *KillProcessRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, LspService_KillProcess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LspServiceServer is the server API for LspService service.
// All implementations must embed UnimplementedLspServiceServer
// for forward compatibility
type LspServiceServer interface {
	// Deprecated: Do not use.
	GetStepNames(context.Context, *StepNamesRequest) (*StepNamesResponse, error)
	// Deprecated: Do not use.
	CacheFile(context.Context, *CacheFileRequest) (*Empty, error)
	// Deprecated: Do not use.
	GetStepPositions(context.Context, *StepPositionsRequest) (*StepPositionsResponse, error)
	// Deprecated: Do not use.
	GetImplementationFiles(context.Context, *Empty) (*ImplementationFileListResponse, error)
	// Deprecated: Do not use.
	ImplementStub(context.Context, *StubImplementationCodeRequest) (*FileDiff, error)
	// Deprecated: Do not use.
	ValidateStep(context.Context, *StepValidateRequest) (*StepValidateResponse, error)
	// Deprecated: Do not use.
	Refactor(context.Context, *RefactorRequest) (*RefactorResponse, error)
	// Deprecated: Do not use.
	GetStepName(context.Context, *StepNameRequest) (*StepNameResponse, error)
	// Deprecated: Do not use.
	GetGlobPatterns(context.Context, *Empty) (*ImplementationFileGlobPatternResponse, error)
	// Deprecated: Do not use.
	KillProcess(context.Context, *KillProcessRequest) (*Empty, error)
	mustEmbedUnimplementedLspServiceServer()
}

// UnimplementedLspServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLspServiceServer struct {
}

func (UnimplementedLspServiceServer) GetStepNames(context.Context, *StepNamesRequest) (*StepNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStepNames not implemented")
}
func (UnimplementedLspServiceServer) CacheFile(context.Context, *CacheFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheFile not implemented")
}
func (UnimplementedLspServiceServer) GetStepPositions(context.Context, *StepPositionsRequest) (*StepPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStepPositions not implemented")
}
func (UnimplementedLspServiceServer) GetImplementationFiles(context.Context, *Empty) (*ImplementationFileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImplementationFiles not implemented")
}
func (UnimplementedLspServiceServer) ImplementStub(context.Context, *StubImplementationCodeRequest) (*FileDiff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImplementStub not implemented")
}
func (UnimplementedLspServiceServer) ValidateStep(context.Context, *StepValidateRequest) (*StepValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateStep not implemented")
}
func (UnimplementedLspServiceServer) Refactor(context.Context, *RefactorRequest) (*RefactorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refactor not implemented")
}
func (UnimplementedLspServiceServer) GetStepName(context.Context, *StepNameRequest) (*StepNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStepName not implemented")
}
func (UnimplementedLspServiceServer) GetGlobPatterns(context.Context, *Empty) (*ImplementationFileGlobPatternResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobPatterns not implemented")
}
func (UnimplementedLspServiceServer) KillProcess(context.Context, *KillProcessRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillProcess not implemented")
}
func (UnimplementedLspServiceServer) mustEmbedUnimplementedLspServiceServer() {}

// UnsafeLspServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LspServiceServer will
// result in compilation errors.
type UnsafeLspServiceServer interface {
	mustEmbedUnimplementedLspServiceServer()
}

func RegisterLspServiceServer(s grpc.ServiceRegistrar, srv LspServiceServer) {
	s.RegisterService(&LspService_ServiceDesc, srv)
}

func _LspService_GetStepNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetStepNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_GetStepNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetStepNames(ctx, req.(*StepNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_CacheFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).CacheFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_CacheFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).CacheFile(ctx, req.(*CacheFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_GetStepPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetStepPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_GetStepPositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetStepPositions(ctx, req.(*StepPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_GetImplementationFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetImplementationFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_GetImplementationFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetImplementationFiles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_ImplementStub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StubImplementationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).ImplementStub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_ImplementStub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).ImplementStub(ctx, req.(*StubImplementationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_ValidateStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).ValidateStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_ValidateStep_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).ValidateStep(ctx, req.(*StepValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_Refactor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefactorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).Refactor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_Refactor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).Refactor(ctx, req.(*RefactorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_GetStepName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetStepName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_GetStepName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetStepName(ctx, req.(*StepNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_GetGlobPatterns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).GetGlobPatterns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_GetGlobPatterns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).GetGlobPatterns(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LspService_KillProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LspServiceServer).KillProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LspService_KillProcess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LspServiceServer).KillProcess(ctx, req.(*KillProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LspService_ServiceDesc is the grpc.ServiceDesc for LspService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LspService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gauge.messages.lspService",
	HandlerType: (*LspServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStepNames",
			Handler:    _LspService_GetStepNames_Handler,
		},
		{
			MethodName: "CacheFile",
			Handler:    _LspService_CacheFile_Handler,
		},
		{
			MethodName: "GetStepPositions",
			Handler:    _LspService_GetStepPositions_Handler,
		},
		{
			MethodName: "GetImplementationFiles",
			Handler:    _LspService_GetImplementationFiles_Handler,
		},
		{
			MethodName: "ImplementStub",
			Handler:    _LspService_ImplementStub_Handler,
		},
		{
			MethodName: "ValidateStep",
			Handler:    _LspService_ValidateStep_Handler,
		},
		{
			MethodName: "Refactor",
			Handler:    _LspService_Refactor_Handler,
		},
		{
			MethodName: "GetStepName",
			Handler:    _LspService_GetStepName_Handler,
		},
		{
			MethodName: "GetGlobPatterns",
			Handler:    _LspService_GetGlobPatterns_Handler,
		},
		{
			MethodName: "KillProcess",
			Handler:    _LspService_KillProcess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lsp.proto",
}
