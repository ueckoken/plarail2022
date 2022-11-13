// package: 
// file: statesync.proto

import * as jspb from "google-protobuf";

export class Command2InternalRequest extends jspb.Message {
  hasStation(): boolean;
  clearStation(): void;
  getStation(): Stations | undefined;
  setStation(value?: Stations): void;

  getState(): Command2InternalRequest.StateMap[keyof Command2InternalRequest.StateMap];
  setState(value: Command2InternalRequest.StateMap[keyof Command2InternalRequest.StateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Command2InternalRequest.AsObject;
  static toObject(includeInstance: boolean, msg: Command2InternalRequest): Command2InternalRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Command2InternalRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Command2InternalRequest;
  static deserializeBinaryFromReader(message: Command2InternalRequest, reader: jspb.BinaryReader): Command2InternalRequest;
}

export namespace Command2InternalRequest {
  export type AsObject = {
    station?: Stations.AsObject,
    state: Command2InternalRequest.StateMap[keyof Command2InternalRequest.StateMap],
  }

  export interface StateMap {
    UNKNOWN: 0;
    ON: 1;
    OFF: 2;
  }

  export const State: StateMap;
}

export class RequestSync extends jspb.Message {
  hasStation(): boolean;
  clearStation(): void;
  getStation(): Stations | undefined;
  setStation(value?: Stations): void;

  getState(): RequestSync.StateMap[keyof RequestSync.StateMap];
  setState(value: RequestSync.StateMap[keyof RequestSync.StateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RequestSync.AsObject;
  static toObject(includeInstance: boolean, msg: RequestSync): RequestSync.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RequestSync, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RequestSync;
  static deserializeBinaryFromReader(message: RequestSync, reader: jspb.BinaryReader): RequestSync;
}

export namespace RequestSync {
  export type AsObject = {
    station?: Stations.AsObject,
    state: RequestSync.StateMap[keyof RequestSync.StateMap],
  }

  export interface StateMap {
    UNKNOWN: 0;
    ON: 1;
    OFF: 2;
  }

  export const State: StateMap;
}

export class ResponseSync extends jspb.Message {
  getResponse(): ResponseSync.ResponseMap[keyof ResponseSync.ResponseMap];
  setResponse(value: ResponseSync.ResponseMap[keyof ResponseSync.ResponseMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResponseSync.AsObject;
  static toObject(includeInstance: boolean, msg: ResponseSync): ResponseSync.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResponseSync, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResponseSync;
  static deserializeBinaryFromReader(message: ResponseSync, reader: jspb.BinaryReader): ResponseSync;
}

export namespace ResponseSync {
  export type AsObject = {
    response: ResponseSync.ResponseMap[keyof ResponseSync.ResponseMap],
  }

  export interface ResponseMap {
    UNKNOWN: 0;
    SUCCESS: 1;
    FAILED: 2;
  }

  export const Response: ResponseMap;
}

export class Command2InternalResponse extends jspb.Message {
  getResponse(): Command2InternalResponse.ResponseMap[keyof Command2InternalResponse.ResponseMap];
  setResponse(value: Command2InternalResponse.ResponseMap[keyof Command2InternalResponse.ResponseMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Command2InternalResponse.AsObject;
  static toObject(includeInstance: boolean, msg: Command2InternalResponse): Command2InternalResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Command2InternalResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Command2InternalResponse;
  static deserializeBinaryFromReader(message: Command2InternalResponse, reader: jspb.BinaryReader): Command2InternalResponse;
}

export namespace Command2InternalResponse {
  export type AsObject = {
    response: Command2InternalResponse.ResponseMap[keyof Command2InternalResponse.ResponseMap],
  }

  export interface ResponseMap {
    UNKNOWN: 0;
    SUCCESS: 1;
    FAILED: 2;
  }

  export const Response: ResponseMap;
}

export class Stations extends jspb.Message {
  getStationid(): Stations.StationIdMap[keyof Stations.StationIdMap];
  setStationid(value: Stations.StationIdMap[keyof Stations.StationIdMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Stations.AsObject;
  static toObject(includeInstance: boolean, msg: Stations): Stations.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Stations, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Stations;
  static deserializeBinaryFromReader(message: Stations, reader: jspb.BinaryReader): Stations;
}

export namespace Stations {
  export type AsObject = {
    stationid: Stations.StationIdMap[keyof Stations.StationIdMap],
  }

  export interface StationIdMap {
    UNKNOWN: 0;
    SHINJUKU_P1: 1;
    SHINJUKU_S1: 2;
    SHINJUKU_S2: 3;
    SAKURAJOSUI_P1: 11;
    SAKURAJOSUI_P2: 12;
    SAKURAJOSUI_S1: 13;
    SAKURAJOSUI_S2: 14;
    SAKURAJOSUI_S3: 15;
    SAKURAJOSUI_S4: 16;
    CHOFU_P1: 21;
    CHOFU_P2: 22;
    CHOFU_P3: 23;
    CHOFU_P4: 24;
    CHOFU_P5: 25;
    CHOFU_S1: 26;
    CHOFU_S2: 27;
    CHOFU_S3: 28;
    CHOFU_S4: 29;
    HASHIMOTO_S1: 31;
    HASHIMOTO_S2: 32;
    HACHIOJI_S1: 41;
    HACHIOJI_S2: 42;
    WAKABADAI_P1: 51;
    WAKABADAI_P2: 52;
    WAKABADAI_S1: 53;
    WAKABADAI_S2: 54;
  }

  export const StationId: StationIdMap;
}

