// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: mashupsdk/mashupsdk.proto

package mashupsdk

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MashupServer_CollaborateInit_FullMethodName      = "/mashupsdk.MashupServer/CollaborateInit"
	MashupServer_OnDisplayChange_FullMethodName      = "/mashupsdk.MashupServer/OnDisplayChange"
	MashupServer_GetElements_FullMethodName          = "/mashupsdk.MashupServer/GetElements"
	MashupServer_UpsertElements_FullMethodName       = "/mashupsdk.MashupServer/UpsertElements"
	MashupServer_TweakStates_FullMethodName          = "/mashupsdk.MashupServer/TweakStates"
	MashupServer_ResetStates_FullMethodName          = "/mashupsdk.MashupServer/ResetStates"
	MashupServer_TweakStatesByMotiv_FullMethodName   = "/mashupsdk.MashupServer/TweakStatesByMotiv"
	MashupServer_CollaborateBootstrap_FullMethodName = "/mashupsdk.MashupServer/CollaborateBootstrap"
	MashupServer_Shutdown_FullMethodName             = "/mashupsdk.MashupServer/Shutdown"
)

// MashupServerClient is the client API for MashupServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The mashup service definition.
type MashupServerClient interface {
	// Indicates to mashup it is time to shutdown.
	CollaborateInit(ctx context.Context, in *MashupConnectionConfigs, opts ...grpc.CallOption) (*MashupConnectionConfigs, error)
	OnDisplayChange(ctx context.Context, in *MashupDisplayBundle, opts ...grpc.CallOption) (*MashupDisplayHint, error)
	GetElements(ctx context.Context, in *MashupEmpty, opts ...grpc.CallOption) (*MashupDetailedElementBundle, error)
	UpsertElements(ctx context.Context, in *MashupDetailedElementBundle, opts ...grpc.CallOption) (*MashupDetailedElementBundle, error)
	TweakStates(ctx context.Context, in *MashupElementStateBundle, opts ...grpc.CallOption) (*MashupElementStateBundle, error)
	ResetStates(ctx context.Context, in *MashupEmpty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TweakStatesByMotiv(ctx context.Context, in *Motiv, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CollaborateBootstrap(ctx context.Context, in *MashupConnectionConfigs, opts ...grpc.CallOption) (*MashupEmpty, error)
	Shutdown(ctx context.Context, in *MashupEmpty, opts ...grpc.CallOption) (*MashupEmpty, error)
}

type mashupServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMashupServerClient(cc grpc.ClientConnInterface) MashupServerClient {
	return &mashupServerClient{cc}
}

func (c *mashupServerClient) CollaborateInit(ctx context.Context, in *MashupConnectionConfigs, opts ...grpc.CallOption) (*MashupConnectionConfigs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupConnectionConfigs)
	err := c.cc.Invoke(ctx, MashupServer_CollaborateInit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) OnDisplayChange(ctx context.Context, in *MashupDisplayBundle, opts ...grpc.CallOption) (*MashupDisplayHint, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupDisplayHint)
	err := c.cc.Invoke(ctx, MashupServer_OnDisplayChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) GetElements(ctx context.Context, in *MashupEmpty, opts ...grpc.CallOption) (*MashupDetailedElementBundle, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupDetailedElementBundle)
	err := c.cc.Invoke(ctx, MashupServer_GetElements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) UpsertElements(ctx context.Context, in *MashupDetailedElementBundle, opts ...grpc.CallOption) (*MashupDetailedElementBundle, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupDetailedElementBundle)
	err := c.cc.Invoke(ctx, MashupServer_UpsertElements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) TweakStates(ctx context.Context, in *MashupElementStateBundle, opts ...grpc.CallOption) (*MashupElementStateBundle, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupElementStateBundle)
	err := c.cc.Invoke(ctx, MashupServer_TweakStates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) ResetStates(ctx context.Context, in *MashupEmpty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MashupServer_ResetStates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) TweakStatesByMotiv(ctx context.Context, in *Motiv, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MashupServer_TweakStatesByMotiv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) CollaborateBootstrap(ctx context.Context, in *MashupConnectionConfigs, opts ...grpc.CallOption) (*MashupEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupEmpty)
	err := c.cc.Invoke(ctx, MashupServer_CollaborateBootstrap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mashupServerClient) Shutdown(ctx context.Context, in *MashupEmpty, opts ...grpc.CallOption) (*MashupEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MashupEmpty)
	err := c.cc.Invoke(ctx, MashupServer_Shutdown_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MashupServerServer is the server API for MashupServer service.
// All implementations must embed UnimplementedMashupServerServer
// for forward compatibility.
//
// The mashup service definition.
type MashupServerServer interface {
	// Indicates to mashup it is time to shutdown.
	CollaborateInit(context.Context, *MashupConnectionConfigs) (*MashupConnectionConfigs, error)
	OnDisplayChange(context.Context, *MashupDisplayBundle) (*MashupDisplayHint, error)
	GetElements(context.Context, *MashupEmpty) (*MashupDetailedElementBundle, error)
	UpsertElements(context.Context, *MashupDetailedElementBundle) (*MashupDetailedElementBundle, error)
	TweakStates(context.Context, *MashupElementStateBundle) (*MashupElementStateBundle, error)
	ResetStates(context.Context, *MashupEmpty) (*emptypb.Empty, error)
	TweakStatesByMotiv(context.Context, *Motiv) (*emptypb.Empty, error)
	CollaborateBootstrap(context.Context, *MashupConnectionConfigs) (*MashupEmpty, error)
	Shutdown(context.Context, *MashupEmpty) (*MashupEmpty, error)
	mustEmbedUnimplementedMashupServerServer()
}

// UnimplementedMashupServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMashupServerServer struct{}

func (UnimplementedMashupServerServer) CollaborateInit(context.Context, *MashupConnectionConfigs) (*MashupConnectionConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollaborateInit not implemented")
}
func (UnimplementedMashupServerServer) OnDisplayChange(context.Context, *MashupDisplayBundle) (*MashupDisplayHint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnDisplayChange not implemented")
}
func (UnimplementedMashupServerServer) GetElements(context.Context, *MashupEmpty) (*MashupDetailedElementBundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElements not implemented")
}
func (UnimplementedMashupServerServer) UpsertElements(context.Context, *MashupDetailedElementBundle) (*MashupDetailedElementBundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertElements not implemented")
}
func (UnimplementedMashupServerServer) TweakStates(context.Context, *MashupElementStateBundle) (*MashupElementStateBundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TweakStates not implemented")
}
func (UnimplementedMashupServerServer) ResetStates(context.Context, *MashupEmpty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetStates not implemented")
}
func (UnimplementedMashupServerServer) TweakStatesByMotiv(context.Context, *Motiv) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TweakStatesByMotiv not implemented")
}
func (UnimplementedMashupServerServer) CollaborateBootstrap(context.Context, *MashupConnectionConfigs) (*MashupEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollaborateBootstrap not implemented")
}
func (UnimplementedMashupServerServer) Shutdown(context.Context, *MashupEmpty) (*MashupEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedMashupServerServer) mustEmbedUnimplementedMashupServerServer() {}
func (UnimplementedMashupServerServer) testEmbeddedByValue()                      {}

// UnsafeMashupServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MashupServerServer will
// result in compilation errors.
type UnsafeMashupServerServer interface {
	mustEmbedUnimplementedMashupServerServer()
}

func RegisterMashupServerServer(s grpc.ServiceRegistrar, srv MashupServerServer) {
	// If the following call pancis, it indicates UnimplementedMashupServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MashupServer_ServiceDesc, srv)
}

func _MashupServer_CollaborateInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupConnectionConfigs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).CollaborateInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_CollaborateInit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).CollaborateInit(ctx, req.(*MashupConnectionConfigs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_OnDisplayChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupDisplayBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).OnDisplayChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_OnDisplayChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).OnDisplayChange(ctx, req.(*MashupDisplayBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_GetElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).GetElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_GetElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).GetElements(ctx, req.(*MashupEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_UpsertElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupDetailedElementBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).UpsertElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_UpsertElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).UpsertElements(ctx, req.(*MashupDetailedElementBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_TweakStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupElementStateBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).TweakStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_TweakStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).TweakStates(ctx, req.(*MashupElementStateBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_ResetStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).ResetStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_ResetStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).ResetStates(ctx, req.(*MashupEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_TweakStatesByMotiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Motiv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).TweakStatesByMotiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_TweakStatesByMotiv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).TweakStatesByMotiv(ctx, req.(*Motiv))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_CollaborateBootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupConnectionConfigs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).CollaborateBootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_CollaborateBootstrap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).CollaborateBootstrap(ctx, req.(*MashupConnectionConfigs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MashupServer_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MashupEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MashupServerServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MashupServer_Shutdown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MashupServerServer).Shutdown(ctx, req.(*MashupEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

// MashupServer_ServiceDesc is the grpc.ServiceDesc for MashupServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MashupServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mashupsdk.MashupServer",
	HandlerType: (*MashupServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollaborateInit",
			Handler:    _MashupServer_CollaborateInit_Handler,
		},
		{
			MethodName: "OnDisplayChange",
			Handler:    _MashupServer_OnDisplayChange_Handler,
		},
		{
			MethodName: "GetElements",
			Handler:    _MashupServer_GetElements_Handler,
		},
		{
			MethodName: "UpsertElements",
			Handler:    _MashupServer_UpsertElements_Handler,
		},
		{
			MethodName: "TweakStates",
			Handler:    _MashupServer_TweakStates_Handler,
		},
		{
			MethodName: "ResetStates",
			Handler:    _MashupServer_ResetStates_Handler,
		},
		{
			MethodName: "TweakStatesByMotiv",
			Handler:    _MashupServer_TweakStatesByMotiv_Handler,
		},
		{
			MethodName: "CollaborateBootstrap",
			Handler:    _MashupServer_CollaborateBootstrap_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _MashupServer_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mashupsdk/mashupsdk.proto",
}
