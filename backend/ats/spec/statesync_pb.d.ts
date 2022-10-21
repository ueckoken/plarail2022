// package: 
// file: statesync.proto

import * as jspb from "google-protobuf";

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
    MOTOYAWATA_S1: 1;
    MOTOYAWATA_S2: 2;
    IWAMOTOCHO_S1: 11;
    IWAMOTOCHO_S2: 12;
    IWAMOTOCHO_S4: 13;
    IWAMOTOCHO_B1: 14;
    IWAMOTOCHO_B2: 15;
    IWAMOTOCHO_B3: 16;
    IWAMOTOCHO_B4: 17;
    KUDANSHITA_S5: 21;
    KUDANSHITA_S6: 22;
    SASAZUKA_B1: 31;
    SASAZUKA_B2: 32;
    SASAZUKA_S1: 33;
    SASAZUKA_S2: 34;
    SASAZUKA_S3: 35;
    SASAZUKA_S4: 36;
    SASAZUKA_S5: 37;
    MEIDAIMAE_S1: 41;
    MEIDAIMAE_S2: 42;
    CHOFU_S1: 51;
    CHOFU_S2: 52;
    CHOFU_S3: 53;
    CHOFU_S4: 54;
    CHOFU_S5: 55;
    CHOFU_S6: 56;
    CHOFU_B1: 57;
    CHOFU_B2: 58;
    CHOFU_B3: 59;
    CHOFU_B4: 60;
    CHOFU_B5: 61;
    KITANO_B1: 71;
    KITANO_B2: 72;
    KITANO_B3: 73;
    KITANO_S1: 74;
    KITANO_S2: 75;
    KITANO_S3: 76;
    KITANO_S4: 77;
    KITANO_S5: 78;
    KITANO_S6: 79;
    KITANO_S7: 80;
    TAKAO_S1: 91;
    TAKAO_S2: 92;
  }

  export const StationId: StationIdMap;
}

