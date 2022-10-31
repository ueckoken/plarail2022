// package: 
// file: ats.proto

import * as jspb from "google-protobuf";

export class SendStatusRequest extends jspb.Message {
  getSensor(): SendStatusRequest.SensorNameMap[keyof SendStatusRequest.SensorNameMap];
  setSensor(value: SendStatusRequest.SensorNameMap[keyof SendStatusRequest.SensorNameMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SendStatusRequest): SendStatusRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SendStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendStatusRequest;
  static deserializeBinaryFromReader(message: SendStatusRequest, reader: jspb.BinaryReader): SendStatusRequest;
}

export namespace SendStatusRequest {
  export type AsObject = {
    sensor: SendStatusRequest.SensorNameMap[keyof SendStatusRequest.SensorNameMap],
  }

  export interface SensorNameMap {
    UNKNOWN: 0;
    SHINJUKU_D1: 1;
    SHINJUKU_D2: 2;
    SAKURAJOSUI_D1: 11;
    SAKURAJOSUI_D2: 12;
    SAKURAJOSUI_D3: 13;
    SAKURAJOSUI_D4: 14;
    CHOFU_D1: 21;
    CHOFU_D2: 22;
    CHOFU_D3: 23;
    CHOFU_D4: 24;
    CHOFU_D5: 25;
    CHOFU_D6: 26;
    HASHIMOTO_D1: 31;
    HASHIMOTO_D2: 32;
    HACHIOJI_D1: 41;
    HACHIOJI_D2: 42;
    WAKABADAI_D1: 51;
    WAKABADAI_D2: 52;
  }

  export const SensorName: SensorNameMap;
}

export class SendStatusResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendStatusResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SendStatusResponse): SendStatusResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SendStatusResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendStatusResponse;
  static deserializeBinaryFromReader(message: SendStatusResponse, reader: jspb.BinaryReader): SendStatusResponse;
}

export namespace SendStatusResponse {
  export type AsObject = {
  }
}

