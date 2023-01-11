// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: servidor_2/servicios.proto

package gen

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

// ProveedorClient is the client API for Proveedor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProveedorClient interface {
	SuministrarProductos(ctx context.Context, in *Producto, opts ...grpc.CallOption) (*Respuesta, error)
}

type proveedorClient struct {
	cc grpc.ClientConnInterface
}

func NewProveedorClient(cc grpc.ClientConnInterface) ProveedorClient {
	return &proveedorClient{cc}
}

func (c *proveedorClient) SuministrarProductos(ctx context.Context, in *Producto, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/test.Proveedor/SuministrarProductos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProveedorServer is the server API for Proveedor service.
// All implementations must embed UnimplementedProveedorServer
// for forward compatibility
type ProveedorServer interface {
	SuministrarProductos(context.Context, *Producto) (*Respuesta, error)
	mustEmbedUnimplementedProveedorServer()
}

// UnimplementedProveedorServer must be embedded to have forward compatible implementations.
type UnimplementedProveedorServer struct {
}

func (UnimplementedProveedorServer) SuministrarProductos(context.Context, *Producto) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuministrarProductos not implemented")
}
func (UnimplementedProveedorServer) mustEmbedUnimplementedProveedorServer() {}

// UnsafeProveedorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProveedorServer will
// result in compilation errors.
type UnsafeProveedorServer interface {
	mustEmbedUnimplementedProveedorServer()
}

func RegisterProveedorServer(s grpc.ServiceRegistrar, srv ProveedorServer) {
	s.RegisterService(&Proveedor_ServiceDesc, srv)
}

func _Proveedor_SuministrarProductos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Producto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProveedorServer).SuministrarProductos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Proveedor/SuministrarProductos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProveedorServer).SuministrarProductos(ctx, req.(*Producto))
	}
	return interceptor(ctx, in, info, handler)
}

// Proveedor_ServiceDesc is the grpc.ServiceDesc for Proveedor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proveedor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.Proveedor",
	HandlerType: (*ProveedorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuministrarProductos",
			Handler:    _Proveedor_SuministrarProductos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servidor_2/servicios.proto",
}