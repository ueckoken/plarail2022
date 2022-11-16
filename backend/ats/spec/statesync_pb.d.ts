// package: 
// file: statesync.proto

import * as jspb from "google-protobuf";

export class Station extends jspb.Message {
  getStationid(): StationIdMap[keyof StationIdMap];
  setStationid(value: StationIdMap[keyof StationIdMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Station.AsObject;
  static toObject(includeInstance: boolean, msg: Station): Station.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Station, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Station;
  static deserializeBinaryFromReader(message: Station, reader: jspb.BinaryReader): Station;
}

export namespace Station {
  export type AsObject = {
    stationid: StationIdMap[keyof StationIdMap],
  }
}

export class UpdatePointStateRequest extends jspb.Message {
  hasStation(): boolean;
  clearStation(): void;
  getStation(): Station | undefined;
  setStation(value?: Station): void;

  getState(): PointStateMap[keyof PointStateMap];
  setState(value: PointStateMap[keyof PointStateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePointStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePointStateRequest): UpdatePointStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePointStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePointStateRequest;
  static deserializeBinaryFromReader(message: UpdatePointStateRequest, reader: jspb.BinaryReader): UpdatePointStateRequest;
}

export namespace UpdatePointStateRequest {
  export type AsObject = {
    station?: Station.AsObject,
    state: PointStateMap[keyof PointStateMap],
  }
}

export class UpdatePointStateResponse extends jspb.Message {
  getResponse(): ResponseCodeMap[keyof ResponseCodeMap];
  setResponse(value: ResponseCodeMap[keyof ResponseCodeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePointStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePointStateResponse): UpdatePointStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePointStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePointStateResponse;
  static deserializeBinaryFromReader(message: UpdatePointStateResponse, reader: jspb.BinaryReader): UpdatePointStateResponse;
}

export namespace UpdatePointStateResponse {
  export type AsObject = {
    response: ResponseCodeMap[keyof ResponseCodeMap],
  }
}

export class NotifyPointStateRequest extends jspb.Message {
  hasStation(): boolean;
  clearStation(): void;
  getStation(): Station | undefined;
  setStation(value?: Station): void;

  getState(): PointStateMap[keyof PointStateMap];
  setState(value: PointStateMap[keyof PointStateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyPointStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyPointStateRequest): NotifyPointStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyPointStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyPointStateRequest;
  static deserializeBinaryFromReader(message: NotifyPointStateRequest, reader: jspb.BinaryReader): NotifyPointStateRequest;
}

export namespace NotifyPointStateRequest {
  export type AsObject = {
    station?: Station.AsObject,
    state: PointStateMap[keyof PointStateMap],
  }
}

export class NotifyPointStateResponse extends jspb.Message {
  getResponse(): ResponseCodeMap[keyof ResponseCodeMap];
  setResponse(value: ResponseCodeMap[keyof ResponseCodeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyPointStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyPointStateResponse): NotifyPointStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyPointStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyPointStateResponse;
  static deserializeBinaryFromReader(message: NotifyPointStateResponse, reader: jspb.BinaryReader): NotifyPointStateResponse;
}

export namespace NotifyPointStateResponse {
  export type AsObject = {
    response: ResponseCodeMap[keyof ResponseCodeMap],
  }
}

export interface StationIdMap {
  STATIONID_UNKNOWN: 0;
  SHINJUKU_S1: 2;
  SHINJUKU_S2: 3;
  SAKURAJOSUI_P1: 11;
  SAKURAJOSUI_P2: 12;
  SAKURAJOSUI_S0: 13;
  SAKURAJOSUI_S1: 14;
  SAKURAJOSUI_S2: 15;
  SAKURAJOSUI_S3: 16;
  SAKURAJOSUI_S4: 17;
  SAKURAJOSUI_S5: 18;
  CHOFU_P1: 21;
  CHOFU_S0: 22;
  CHOFU_S1: 23;
  CHOFU_S2: 24;
  CHOFU_S3: 25;
  CHOFU_S4: 26;
  HASHIMOTO_S1: 31;
  HASHIMOTO_S2: 32;
  HACHIOJI_S1: 41;
  HACHIOJI_S2: 42;
}

export const StationId: StationIdMap;

export interface PointStateMap {
  UNKNOWN: 0;
  ON: 1;
  OFF: 2;
}

export const PointState: PointStateMap;

export interface ResponseCodeMap {
  RESPONSECODE_UNKNOWN: 0;
  RESPONSECODE_SUCCESS: 1;
  RESPONSECODE_FAILED: 2;
}

export const ResponseCode: ResponseCodeMap;

