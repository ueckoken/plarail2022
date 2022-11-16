// package: 
// file: block.proto

import * as jspb from "google-protobuf";

export class BlockAndState extends jspb.Message {
  getBlockid(): BlockIdMap[keyof BlockIdMap];
  setBlockid(value: BlockIdMap[keyof BlockIdMap]): void;

  getState(): BlockStateEnumMap[keyof BlockStateEnumMap];
  setState(value: BlockStateEnumMap[keyof BlockStateEnumMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BlockAndState.AsObject;
  static toObject(includeInstance: boolean, msg: BlockAndState): BlockAndState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BlockAndState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BlockAndState;
  static deserializeBinaryFromReader(message: BlockAndState, reader: jspb.BinaryReader): BlockAndState;
}

export namespace BlockAndState {
  export type AsObject = {
    blockid: BlockIdMap[keyof BlockIdMap],
    state: BlockStateEnumMap[keyof BlockStateEnumMap],
  }
}

export class UpdateBlockStateRequest extends jspb.Message {
  hasState(): boolean;
  clearState(): void;
  getState(): BlockAndState | undefined;
  setState(value?: BlockAndState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateBlockStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateBlockStateRequest): UpdateBlockStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateBlockStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateBlockStateRequest;
  static deserializeBinaryFromReader(message: UpdateBlockStateRequest, reader: jspb.BinaryReader): UpdateBlockStateRequest;
}

export namespace UpdateBlockStateRequest {
  export type AsObject = {
    state?: BlockAndState.AsObject,
  }
}

export class UpdateBlockStateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateBlockStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateBlockStateResponse): UpdateBlockStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateBlockStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateBlockStateResponse;
  static deserializeBinaryFromReader(message: UpdateBlockStateResponse, reader: jspb.BinaryReader): UpdateBlockStateResponse;
}

export namespace UpdateBlockStateResponse {
  export type AsObject = {
  }
}

export class NotifyBlockStateRequest extends jspb.Message {
  hasState(): boolean;
  clearState(): void;
  getState(): BlockAndState | undefined;
  setState(value?: BlockAndState): void;

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
    state?: BlockAndState.AsObject,
  }
}

export class NotifyBlockStateResponse extends jspb.Message {
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
  }
}

export interface BlockStateEnumMap {
  BLOCKSTATE_UNKNOWN: 0;
  BLOCKSTATE_OPEN: 1;
  BLOCKSTATE_CLOSE: 2;
}

export const BlockStateEnum: BlockStateEnumMap;

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

