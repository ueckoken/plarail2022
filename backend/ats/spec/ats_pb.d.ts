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

