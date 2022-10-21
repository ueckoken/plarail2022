// package: 
// file: speedControl.proto

import * as jspb from "google-protobuf";

export class SendSpeed extends jspb.Message {
  getSpeed(): number;
  setSpeed(value: number): void;

  getTrain(): SendSpeed.TrainMap[keyof SendSpeed.TrainMap];
  setTrain(value: SendSpeed.TrainMap[keyof SendSpeed.TrainMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendSpeed.AsObject;
  static toObject(includeInstance: boolean, msg: SendSpeed): SendSpeed.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SendSpeed, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendSpeed;
  static deserializeBinaryFromReader(message: SendSpeed, reader: jspb.BinaryReader): SendSpeed;
}

export namespace SendSpeed {
  export type AsObject = {
    speed: number,
    train: SendSpeed.TrainMap[keyof SendSpeed.TrainMap],
  }

  export interface TrainMap {
    UNKNOWN: 0;
    TAKAO: 1;
    CHICHIBU: 2;
    HAKONE: 3;
    OKUTAMA: 4;
    NIKKO: 5;
    ENOSHIMA: 6;
    KAMAKURA: 7;
    YOKOSUKA: 8;
  }

  export const Train: TrainMap;
}

export class StatusCode extends jspb.Message {
  getCode(): StatusCode.CodeMap[keyof StatusCode.CodeMap];
  setCode(value: StatusCode.CodeMap[keyof StatusCode.CodeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatusCode.AsObject;
  static toObject(includeInstance: boolean, msg: StatusCode): StatusCode.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StatusCode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatusCode;
  static deserializeBinaryFromReader(message: StatusCode, reader: jspb.BinaryReader): StatusCode;
}

export namespace StatusCode {
  export type AsObject = {
    code: StatusCode.CodeMap[keyof StatusCode.CodeMap],
  }

  export interface CodeMap {
    UNKNOWN: 0;
    SUCCESS: 1;
    FAILED: 2;
  }

  export const Code: CodeMap;
}

