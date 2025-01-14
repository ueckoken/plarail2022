# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import ats_pb2 as ats__pb2


class AtsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendStatus = channel.unary_unary(
                '/Ats/SendStatus',
                request_serializer=ats__pb2.SendStatusRequest.SerializeToString,
                response_deserializer=ats__pb2.SendStatusResponse.FromString,
                )


class AtsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AtsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendStatus': grpc.unary_unary_rpc_method_handler(
                    servicer.SendStatus,
                    request_deserializer=ats__pb2.SendStatusRequest.FromString,
                    response_serializer=ats__pb2.SendStatusResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Ats', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Ats(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendStatus(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Ats/SendStatus',
            ats__pb2.SendStatusRequest.SerializeToString,
            ats__pb2.SendStatusResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
