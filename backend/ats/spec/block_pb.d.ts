// package: 
// file: block.proto

import * as jspb from "google-protobuf";

export class NotifyStateRequest extends jspb.Message {
  getState(): NotifyStateRequest.StateMap[keyof NotifyStateRequest.StateMap];
  setState(value: NotifyStateRequest.StateMap[keyof NotifyStateRequest.StateMap]): void;

  hasBlockid(): boolean;
  clearBlockid(): void;
  getBlockid(): Blocks | undefined;
  setBlockid(value?: Blocks): void;

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
    blockid?: Blocks.AsObject,
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
    SHINJUKU_SAKURAJOSUI_UP: 10;
    SHINJUKU_SAKURAJOSUI_DOWN: 11;
    SAKURAJOSUI_CHOFU_UP: 20;
    SAKURAJOSUI_CHOFU_DOWN: 21;
    CHOFU_HACHIOJI_UP: 30;
    CHOFU_HACHIOJI_DOWN: 31;
    CHOFU_HASHIMOTO_UP: 40;
    CHOFU_HASHIMOTO_DOWN: 41;
    SHINJUKU_B1: 100;
    SHINJUKU_B2: 101;
    SAKURAJOSUI_B1: 110;
    SAKURAJOSUI_B2: 111;
    SAKURAJOSUI_B3: 120;
    SAKURAJOSUI_B4: 121;
    CHOFU_B1: 130;
    CHOFU_B2: 131;
    CHOFU_B3: 132;
    CHOFU_B4: 133;
    HASHIMOTO_B1: 140;
    HASHIMOTO_B2: 141;
    HACHIOJI_B1: 150;
    HACHIOJI_B2: 151;
  }

  export const BlockId: BlockIdMap;
}

