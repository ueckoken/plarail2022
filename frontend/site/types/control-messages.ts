import { BlockId } from "./../../../backend/ats/spec/block_pb.d"
import * as t from "io-ts"

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/proto/statesync.proto#L28-L85
// 無限 user-defined type guard イヤイヤ期なので io-ts に頼る

export const unknownId = t.literal("unknown")
export const stopRailId = t.union([
  t.literal("shinjuku_s1"),
  t.literal("shinjuku_s2"),
  t.literal("sakurajosui_s0"),
  t.literal("sakurajosui_s1"),
  t.literal("sakurajosui_s2"),
  t.literal("sakurajosui_s3"),
  t.literal("sakurajosui_s4"),
  t.literal("sakurajosui_s5"),
  t.literal("chofu_s0"),
  t.literal("chofu_s1"),
  t.literal("chofu_s2"),
  t.literal("chofu_s3"),
  t.literal("chofu_s4"),
  t.literal("hashimoto_s1"),
  t.literal("hashimoto_s2"),
  t.literal("hachioji_s1"),
  t.literal("hachioji_s2"),
])

export const pointIdMap = {
  0: "unknown",
  2: "shinjuku_s1",
  3: "shinjuku_s2",
  11: "sakurajosui_p1",
  12: "sakurajosui_p2",
  13: "sakurajosui_s0",
  14: "sakurajosui_s1",
  15: "sakurajosui_s2",
  16: "sakurajosui_s3",
  17: "sakurajosui_s4",
  18: "sakurajosui_s5",
  21: "chofu_p1",
  22: "chofu_s0",
  23: "chofu_s1",
  24: "chofu_s2",
  25: "chofu_s3",
  26: "chofu_s4",
  31: "hashimoto_s1",
  32: "hashimoto_s2",
  41: "hachioji_s1",
  42: "hachioji_s2",
}

export const pointIdMapReverse = {
  unknown: 0,
  shinjuku_s1: 2,
  shinjuku_s2: 3,
  sakurajosui_p1: 11,
  sakurajosui_p2: 12,
  sakurajosui_s0: 13,
  sakurajosui_s1: 14,
  sakurajosui_s2: 15,
  sakurajosui_s3: 16,
  sakurajosui_s4: 17,
  sakurajosui_s5: 18,
  chofu_p1: 21,
  chofu_s0: 22,
  chofu_s1: 23,
  chofu_s2: 24,
  chofu_s3: 25,
  chofu_s4: 26,
  hashimoto_s1: 31,
  hashimoto_s2: 32,
  hachioji_s1: 41,
  hachioji_s2: 42,
}

export const blockIdMap = {
  0: "unknown",
  1: "shinjuku_b1",
  2: "shinjuku_b2",
  11: "sakurajosui_b1",
  12: "sakurajosui_b2",
  13: "sakurajosui_b3",
  14: "sakurajosui_b4",
  15: "sakurajosui_b5",
  16: "sakurajosui_b6",
  21: "chofu_b1",
  22: "chofu_b2",
  23: "chofu_b3",
  24: "chofu_b4",
  25: "chofu_b5",
  31: "hashimoto_b1",
  32: "hashimoto_b2",
  41: "hachioji_b1",
  42: "hashioji_b2",
}

export const blockIdMapReverse = {
  unknown: 0,
  shinjuku_b1: 1,
  shinjuku_b2: 2,
  sakurajosui_b1: 11,
  sakurajosui_b2: 12,
  sakurajosui_b3: 13,
  sakurajosui_b4: 14,
  sakurajosui_b5: 15,
  sakurajosui_b6: 16,
  chofu_b1: 21,
  chofu_b2: 22,
  chofu_b3: 23,
  chofu_b4: 24,
  chofu_b5: 25,
  hashimoto_b1: 31,
  hashimoto_b2: 32,
  hachioji_b1: 41,
  hashioji_b2: 42,
}

export type StopRailId = t.TypeOf<typeof stopRailId>

export const bunkiRailId = t.union([
  t.literal("sakurajosui_p1"),
  t.literal("chofu_p1"),
])
export type BunkiRailId = t.TypeOf<typeof bunkiRailId>

export const blockId = t.union([
  t.literal("shinjuku_b1"),
  t.literal("shinjuku_b2"),
  t.literal("sakurajosui_b1"),
  t.literal("sakurajosui_b2"),
  t.literal("sakurajosui_b3"),
  t.literal("sakurajosui_b4"),
  t.literal("sakurajosui_b5"),
  t.literal("sakurajosui_b6"),
  t.literal("chofu_b1"),
  t.literal("chofu_b2"),
  t.literal("chofu_b3"),
  t.literal("chofu_b4"),
  t.literal("chofu_b5"),
  t.literal("hashimoto_b1"),
  t.literal("hashimoto_b2"),
  t.literal("hachioji_b1"),
  t.literal("hachioji_b2"),
])
export type BlockId = t.TypeOf<typeof blockIds>

export const stationId = t.union([unknownId, stopRailId, bunkiRailId])
export const blockIds = t.union([unknownId, blockId])
export type StationId = t.TypeOf<typeof stationId>

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/proto/statesync.proto#L10-L14
export type StationState = "UNKNOWN" | "ON" | "OFF"
export type BlockState = "UNKNOWN" | "OPEN" | "CLOSE"

export const BlockStateIdMap = {
  0: "UNKNOWN",
  1: "OPEN",
  2: "CLOSE",
}

export const BlockStateIdMapReverse = {
  UNKNOWN: 0,
  OPEN: 1,
  CLOSE: 2,
}

export const StationStateIdMap = {
  0: "UNKNOWN",
  1: "ON",
  2: "OFF",
}

export const StationStateIdMapReverse = {
  UNKNOWN: 0,
  ON: 1,
  OFF: 2,
}

export type StationMessage = {
  station: {
    stationId: keyof typeof pointIdMap
  }
  state: keyof typeof BlockStateIdMap
}

export type BlockMessage = {
  block_name: BlockId
  state: BlockState
}
