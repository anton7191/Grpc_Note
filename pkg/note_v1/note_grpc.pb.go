// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: note.proto

package note_v1

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

// NoteV1Client is the client API for NoteV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteV1Client interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*GetNoteResponse, error)
	GetListNote(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetListNoteResponse, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error)
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error)
}

type noteV1Client struct {
	cc grpc.ClientConnInterface
}

func NewNoteV1Client(cc grpc.ClientConnInterface) NoteV1Client {
	return &noteV1Client{cc}
}

func (c *noteV1Client) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, "/api.note_v1.NoteV1/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteV1Client) GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*GetNoteResponse, error) {
	out := new(GetNoteResponse)
	err := c.cc.Invoke(ctx, "/api.note_v1.NoteV1/GetNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteV1Client) GetListNote(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetListNoteResponse, error) {
	out := new(GetListNoteResponse)
	err := c.cc.Invoke(ctx, "/api.note_v1.NoteV1/GetListNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteV1Client) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error) {
	out := new(UpdateNoteResponse)
	err := c.cc.Invoke(ctx, "/api.note_v1.NoteV1/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteV1Client) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error) {
	out := new(DeleteNoteResponse)
	err := c.cc.Invoke(ctx, "/api.note_v1.NoteV1/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteV1Server is the server API for NoteV1 service.
// All implementations must embed UnimplementedNoteV1Server
// for forward compatibility
type NoteV1Server interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error)
	GetListNote(context.Context, *Empty) (*GetListNoteResponse, error)
	UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error)
	DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error)
	mustEmbedUnimplementedNoteV1Server()
}

// UnimplementedNoteV1Server must be embedded to have forward compatible implementations.
type UnimplementedNoteV1Server struct {
}

func (UnimplementedNoteV1Server) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNoteV1Server) GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNote not implemented")
}
func (UnimplementedNoteV1Server) GetListNote(context.Context, *Empty) (*GetListNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListNote not implemented")
}
func (UnimplementedNoteV1Server) UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedNoteV1Server) DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedNoteV1Server) mustEmbedUnimplementedNoteV1Server() {}

// UnsafeNoteV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteV1Server will
// result in compilation errors.
type UnsafeNoteV1Server interface {
	mustEmbedUnimplementedNoteV1Server()
}

func RegisterNoteV1Server(s grpc.ServiceRegistrar, srv NoteV1Server) {
	s.RegisterService(&NoteV1_ServiceDesc, srv)
}

func _NoteV1_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.note_v1.NoteV1/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteV1_GetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).GetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.note_v1.NoteV1/GetNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).GetNote(ctx, req.(*GetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteV1_GetListNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).GetListNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.note_v1.NoteV1/GetListNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).GetListNote(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteV1_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.note_v1.NoteV1/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteV1_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.note_v1.NoteV1/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).DeleteNote(ctx, req.(*DeleteNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteV1_ServiceDesc is the grpc.ServiceDesc for NoteV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.note_v1.NoteV1",
	HandlerType: (*NoteV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _NoteV1_CreateNote_Handler,
		},
		{
			MethodName: "GetNote",
			Handler:    _NoteV1_GetNote_Handler,
		},
		{
			MethodName: "GetListNote",
			Handler:    _NoteV1_GetListNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _NoteV1_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _NoteV1_DeleteNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note.proto",
}
