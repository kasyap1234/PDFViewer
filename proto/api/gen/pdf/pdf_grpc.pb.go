// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pdf.proto

package pdf

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
	PDFProcessor_GetPDFInfo_FullMethodName  = "/pdf.PDFProcessor/GetPDFInfo"
	PDFProcessor_GetPDFPage_FullMethodName  = "/pdf.PDFProcessor/GetPDFPage"
	PDFProcessor_SearchPDF_FullMethodName   = "/pdf.PDFProcessor/SearchPDF"
	PDFProcessor_ExtractText_FullMethodName = "/pdf.PDFProcessor/ExtractText"
)

// PDFProcessorClient is the client API for PDFProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PDFProcessorClient interface {
	GetPDFInfo(ctx context.Context, in *PDFInfoRequest, opts ...grpc.CallOption) (*PDFInfoResponse, error)
	GetPDFPage(ctx context.Context, in *PDFPageRequest, opts ...grpc.CallOption) (*PDFPageResponse, error)
	SearchPDF(ctx context.Context, in *SearchPDFRequest, opts ...grpc.CallOption) (*SearchPDFResponse, error)
	ExtractText(ctx context.Context, in *TextExtractionRequest, opts ...grpc.CallOption) (*TextExtractionResponse, error)
}

type pDFProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewPDFProcessorClient(cc grpc.ClientConnInterface) PDFProcessorClient {
	return &pDFProcessorClient{cc}
}

func (c *pDFProcessorClient) GetPDFInfo(ctx context.Context, in *PDFInfoRequest, opts ...grpc.CallOption) (*PDFInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PDFInfoResponse)
	err := c.cc.Invoke(ctx, PDFProcessor_GetPDFInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pDFProcessorClient) GetPDFPage(ctx context.Context, in *PDFPageRequest, opts ...grpc.CallOption) (*PDFPageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PDFPageResponse)
	err := c.cc.Invoke(ctx, PDFProcessor_GetPDFPage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pDFProcessorClient) SearchPDF(ctx context.Context, in *SearchPDFRequest, opts ...grpc.CallOption) (*SearchPDFResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPDFResponse)
	err := c.cc.Invoke(ctx, PDFProcessor_SearchPDF_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pDFProcessorClient) ExtractText(ctx context.Context, in *TextExtractionRequest, opts ...grpc.CallOption) (*TextExtractionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextExtractionResponse)
	err := c.cc.Invoke(ctx, PDFProcessor_ExtractText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PDFProcessorServer is the server API for PDFProcessor service.
// All implementations must embed UnimplementedPDFProcessorServer
// for forward compatibility.
type PDFProcessorServer interface {
	GetPDFInfo(context.Context, *PDFInfoRequest) (*PDFInfoResponse, error)
	GetPDFPage(context.Context, *PDFPageRequest) (*PDFPageResponse, error)
	SearchPDF(context.Context, *SearchPDFRequest) (*SearchPDFResponse, error)
	ExtractText(context.Context, *TextExtractionRequest) (*TextExtractionResponse, error)
	mustEmbedUnimplementedPDFProcessorServer()
}

// UnimplementedPDFProcessorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPDFProcessorServer struct{}

func (UnimplementedPDFProcessorServer) GetPDFInfo(context.Context, *PDFInfoRequest) (*PDFInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPDFInfo not implemented")
}
func (UnimplementedPDFProcessorServer) GetPDFPage(context.Context, *PDFPageRequest) (*PDFPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPDFPage not implemented")
}
func (UnimplementedPDFProcessorServer) SearchPDF(context.Context, *SearchPDFRequest) (*SearchPDFResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPDF not implemented")
}
func (UnimplementedPDFProcessorServer) ExtractText(context.Context, *TextExtractionRequest) (*TextExtractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractText not implemented")
}
func (UnimplementedPDFProcessorServer) mustEmbedUnimplementedPDFProcessorServer() {}
func (UnimplementedPDFProcessorServer) testEmbeddedByValue()                      {}

// UnsafePDFProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PDFProcessorServer will
// result in compilation errors.
type UnsafePDFProcessorServer interface {
	mustEmbedUnimplementedPDFProcessorServer()
}

func RegisterPDFProcessorServer(s grpc.ServiceRegistrar, srv PDFProcessorServer) {
	// If the following call pancis, it indicates UnimplementedPDFProcessorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PDFProcessor_ServiceDesc, srv)
}

func _PDFProcessor_GetPDFInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PDFInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDFProcessorServer).GetPDFInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PDFProcessor_GetPDFInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDFProcessorServer).GetPDFInfo(ctx, req.(*PDFInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PDFProcessor_GetPDFPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PDFPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDFProcessorServer).GetPDFPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PDFProcessor_GetPDFPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDFProcessorServer).GetPDFPage(ctx, req.(*PDFPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PDFProcessor_SearchPDF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPDFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDFProcessorServer).SearchPDF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PDFProcessor_SearchPDF_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDFProcessorServer).SearchPDF(ctx, req.(*SearchPDFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PDFProcessor_ExtractText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextExtractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDFProcessorServer).ExtractText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PDFProcessor_ExtractText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDFProcessorServer).ExtractText(ctx, req.(*TextExtractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PDFProcessor_ServiceDesc is the grpc.ServiceDesc for PDFProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PDFProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdf.PDFProcessor",
	HandlerType: (*PDFProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPDFInfo",
			Handler:    _PDFProcessor_GetPDFInfo_Handler,
		},
		{
			MethodName: "GetPDFPage",
			Handler:    _PDFProcessor_GetPDFPage_Handler,
		},
		{
			MethodName: "SearchPDF",
			Handler:    _PDFProcessor_SearchPDF_Handler,
		},
		{
			MethodName: "ExtractText",
			Handler:    _PDFProcessor_ExtractText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pdf.proto",
}
