# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import servicios_pb2 as servicios__pb2


class ProveedorStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SuministrarProductos = channel.unary_unary(
                '/test.Proveedor/SuministrarProductos',
                request_serializer=servicios__pb2.Producto.SerializeToString,
                response_deserializer=servicios__pb2.Respuesta.FromString,
                )


class ProveedorServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SuministrarProductos(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProveedorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SuministrarProductos': grpc.unary_unary_rpc_method_handler(
                    servicer.SuministrarProductos,
                    request_deserializer=servicios__pb2.Producto.FromString,
                    response_serializer=servicios__pb2.Respuesta.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'test.Proveedor', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Proveedor(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SuministrarProductos(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/test.Proveedor/SuministrarProductos',
            servicios__pb2.Producto.SerializeToString,
            servicios__pb2.Respuesta.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
