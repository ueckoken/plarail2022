# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import block_pb2 as block__pb2


class BlockStateManagerStub(object):
    """BlockStateManagerはblock状態を管理します。externalで動作します。
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UpdateBlockState = channel.unary_unary(
                '/BlockStateManager/UpdateBlockState',
                request_serializer=block__pb2.UpdateBlockStateRequest.SerializeToString,
                response_deserializer=block__pb2.UpdateBlockStateResponse.FromString,
                )


class BlockStateManagerServicer(object):
    """BlockStateManagerはblock状態を管理します。externalで動作します。
    """

    def UpdateBlockState(self, request, context):
        """UpdateBlockStateはサーバのBlockStateを変更します。
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BlockStateManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UpdateBlockState': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateBlockState,
                    request_deserializer=block__pb2.UpdateBlockStateRequest.FromString,
                    response_serializer=block__pb2.UpdateBlockStateResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'BlockStateManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class BlockStateManager(object):
    """BlockStateManagerはblock状態を管理します。externalで動作します。
    """

    @staticmethod
    def UpdateBlockState(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/BlockStateManager/UpdateBlockState',
            block__pb2.UpdateBlockStateRequest.SerializeToString,
            block__pb2.UpdateBlockStateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class BlockStateNotificationStub(object):
    """BlockStateNotificationはblock状態の通知を受けます。auto_operationで動作します。
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.NotifyBlockState = channel.unary_unary(
                '/BlockStateNotification/NotifyBlockState',
                request_serializer=block__pb2.NotifyBlockStateRequest.SerializeToString,
                response_deserializer=block__pb2.NotifyBlockStateResponse.FromString,
                )


class BlockStateNotificationServicer(object):
    """BlockStateNotificationはblock状態の通知を受けます。auto_operationで動作します。
    """

    def NotifyBlockState(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BlockStateNotificationServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'NotifyBlockState': grpc.unary_unary_rpc_method_handler(
                    servicer.NotifyBlockState,
                    request_deserializer=block__pb2.NotifyBlockStateRequest.FromString,
                    response_serializer=block__pb2.NotifyBlockStateResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'BlockStateNotification', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class BlockStateNotification(object):
    """BlockStateNotificationはblock状態の通知を受けます。auto_operationで動作します。
    """

    @staticmethod
    def NotifyBlockState(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/BlockStateNotification/NotifyBlockState',
            block__pb2.NotifyBlockStateRequest.SerializeToString,
            block__pb2.NotifyBlockStateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
