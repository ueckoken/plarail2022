// package: 
// file: block.proto

import * as jspb from "google-protobuf";

export class NotifyBlockStateRequest extends jspb.Message {
  getState(): NotifyBlockStateRequest.StateMap[keyof NotifyBlockStateRequest.StateMap];
  setState(value: NotifyBlockStateRequest.StateMap[keyof NotifyBlockStateRequest.StateMap]): void;

  hasBlock(): boolean;
  clearBlock(): void;
  getBlock(): Blocks | undefined;
  setBlock(value?: Blocks): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyBlockStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyBlockStateRequest): NotifyBlockStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyBlockStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyBlockStateRequest;
  static deserializeBinaryFromReader(message: NotifyBlockStateRequest, reader: jspb.BinaryReader): NotifyBlockStateRequest;
}

export namespace NotifyBlockStateRequest {
  export type AsObject = {
    state: NotifyBlockStateRequest.StateMap[keyof NotifyBlockStateRequest.StateMap],
    block?: Blocks.AsObject,
  }

  export interface StateMap {
    UNKNOWN: 0;
    OPEN: 1;
    CLOSE: 2;
  }

  export const State: StateMap;
}

export class NotifyBlockStateResponse extends jspb.Message {
  getResponse(): NotifyBlockStateResponse.ResponseMap[keyof NotifyBlockStateResponse.ResponseMap];
  setResponse(value: NotifyBlockStateResponse.ResponseMap[keyof NotifyBlockStateResponse.ResponseMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyBlockStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyBlockStateResponse): NotifyBlockStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyBlockStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyBlockStateResponse;
  static deserializeBinaryFromReader(message: NotifyBlockStateResponse, reader: jspb.BinaryReader): NotifyBlockStateResponse;
}

export namespace NotifyBlockStateResponse {
  export type AsObject = {
    response: NotifyBlockStateResponse.ResponseMap[keyof NotifyBlockStateResponse.ResponseMap],
  }

  export interface ResponseMap {
    UNKNOWN: 0;
    SUCCESS: 1;
    FAILED: 2;
  }

  export const Response: ResponseMap;
}

export class Blocks extends jspb.Message {
  getBlockid(): Blocks.BlockIdMap[keyof Blocks.BlockIdMap];
  setBlockid(value: Blocks.BlockIdMap[keyof Blocks.BlockIdMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Blocks.AsObject;
  static toObject(includeInstance: boolean, msg: Blocks): Blocks.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Blocks, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Blocks;
  static deserializeBinaryFromReader(message: Blocks, reader: jspb.BinaryReader): Blocks;
}

export namespace Blocks {
  export type AsObject = {
    blockid: Blocks.BlockIdMap[keyof Blocks.BlockIdMap],
  }

  export interface BlockIdMap {
    UNKNOWN: 0;
    SHINJUKU_B1: 1;
    SHINJUKU_B2: 2;
    SAKURAJOSUI_B1: 11;
    SAKURAJOSUI_B2: 12;
    SAKURAJOSUI_B3: 13;
    SAKURAJOSUI_B4: 14;
    SAKURAJOSUI_B5: 15;
    SAKURAJOSUI_B6: 16;
    CHOFU_B1: 21;
    CHOFU_B2: 22;
    CHOFU_B3: 23;
    CHOFU_B4: 24;
    CHOFU_B5: 25;
    HASHIMOTO_B1: 31;
    HASHIMOTO_B2: 32;
    HACHIOJI_B1: 41;
    HASHIOJI_B2: 42;
  }

  export const BlockId: BlockIdMap;
}

