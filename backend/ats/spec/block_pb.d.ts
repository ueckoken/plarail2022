// package: 
// file: block.proto

import * as jspb from "google-protobuf";

export class NotifyStateRequest extends jspb.Message {
  getState(): NotifyStateRequest.StateMap[keyof NotifyStateRequest.StateMap];
  setState(value: NotifyStateRequest.StateMap[keyof NotifyStateRequest.StateMap]): void;

  hasBlock(): boolean;
  clearBlock(): void;
  getBlock(): Blocks | undefined;
  setBlock(value?: Blocks): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyStateRequest): NotifyStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyStateRequest;
  static deserializeBinaryFromReader(message: NotifyStateRequest, reader: jspb.BinaryReader): NotifyStateRequest;
}

export namespace NotifyStateRequest {
  export type AsObject = {
    state: NotifyStateRequest.StateMap[keyof NotifyStateRequest.StateMap],
    block?: Blocks.AsObject,
  }

  export interface StateMap {
    UNKNOWN: 0;
    OPEN: 1;
    CLOSE: 2;
  }

  export const State: StateMap;
}

export class NotifyStateResponse extends jspb.Message {
  getResponse(): NotifyStateResponse.ResponseMap[keyof NotifyStateResponse.ResponseMap];
  setResponse(value: NotifyStateResponse.ResponseMap[keyof NotifyStateResponse.ResponseMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyStateResponse): NotifyStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyStateResponse;
  static deserializeBinaryFromReader(message: NotifyStateResponse, reader: jspb.BinaryReader): NotifyStateResponse;
}

export namespace NotifyStateResponse {
  export type AsObject = {
    response: NotifyStateResponse.ResponseMap[keyof NotifyStateResponse.ResponseMap],
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

