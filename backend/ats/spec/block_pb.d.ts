// package: 
// file: block.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class BlockStateList extends jspb.Message {
  clearBlockStatesList(): void;
  getBlockStatesList(): Array<BlockState>;
  setBlockStatesList(value: Array<BlockState>): void;
  addBlockStates(value?: BlockState, index?: number): BlockState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BlockStateList.AsObject;
  static toObject(includeInstance: boolean, msg: BlockStateList): BlockStateList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BlockStateList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BlockStateList;
  static deserializeBinaryFromReader(message: BlockStateList, reader: jspb.BinaryReader): BlockStateList;
}

export namespace BlockStateList {
  export type AsObject = {
    blockStatesList: Array<BlockState.AsObject>,
  }
}

export class BlockState extends jspb.Message {
  hasBlock(): boolean;
  clearBlock(): void;
  getBlock(): Blocks | undefined;
  setBlock(value?: Blocks): void;

  getState(): BlockState.StateMap[keyof BlockState.StateMap];
  setState(value: BlockState.StateMap[keyof BlockState.StateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BlockState.AsObject;
  static toObject(includeInstance: boolean, msg: BlockState): BlockState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BlockState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BlockState;
  static deserializeBinaryFromReader(message: BlockState, reader: jspb.BinaryReader): BlockState;
}

export namespace BlockState {
  export type AsObject = {
    block?: Blocks.AsObject,
    state: BlockState.StateMap[keyof BlockState.StateMap],
  }

  export interface StateMap {
    UNKNOWN: 0;
    OPEN: 1;
    CLOSE: 2;
  }

  export const State: StateMap;
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

