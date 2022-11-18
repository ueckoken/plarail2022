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
export type BlocklId = t.TypeOf<typeof blockId>

export const stationId = t.union([unknownId, stopRailId, bunkiRailId])
export type StationId = t.TypeOf<typeof stationId>

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/proto/statesync.proto#L10-L14
export type StationState = "UNKNOWN" | "ON" | "OFF"

export type StationMessage = {
  station_name: StationId
  state: StationState
}

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/external/pkg/clientHandler/clientHandler.go#L28-L31
export type Message = StationMessage
