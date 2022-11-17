import * as t from "io-ts"

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/proto/statesync.proto#L28-L85
// 無限 user-defined type guard イヤイヤ期なので io-ts に頼る
export const unknownId = t.literal("unknown")
export const stopRailId = t.union([
  t.literal("motoyawata_s1"),
  t.literal("motoyawata_s2"),
  t.literal("iwamotocho_s1"),
  t.literal("iwamotocho_s2"),
  t.literal("iwamotocho_s4"),
  t.literal("kudanshita_s5"),
  t.literal("kudanshita_s6"),
  t.literal("sasazuka_s1"),
  t.literal("sasazuka_s2"),
  t.literal("sasazuka_s3"),
  t.literal("sasazuka_s4"),
  t.literal("sasazuka_s5"),
  t.literal("meidaimae_s1"),
  t.literal("meidaimae_s2"),
  t.literal("chofu_s1"),
  t.literal("chofu_s2"),
  t.literal("chofu_s3"),
  t.literal("chofu_s4"),
  t.literal("chofu_s5"),
  t.literal("chofu_s6"),
  t.literal("kitano_s1"),
  t.literal("kitano_s2"),
  t.literal("kitano_s3"),
  t.literal("kitano_s4"),
  t.literal("kitano_s5"),
  t.literal("kitano_s6"),
  t.literal("kitano_s7"),
  t.literal("takao_s1"),
  t.literal("takao_s2"),
])
export type StopRailId = t.TypeOf<typeof stopRailId>

export const bunkiRailId = t.union([
  t.literal("iwamotocho_b1"),
  t.literal("iwamotocho_b2"),
  t.literal("iwamotocho_b3"),
  t.literal("iwamotocho_b4"),
  t.literal("sasazuka_b1"),
  t.literal("sasazuka_b2"),
  t.literal("chofu_b1"),
  t.literal("chofu_b2"),
  t.literal("chofu_b3"),
  t.literal("chofu_b4"),
  t.literal("chofu_b5"),
  t.literal("kitano_b1"),
  t.literal("kitano_b2"),
  t.literal("kitano_b3"),
])
export type BunkiRailId = t.TypeOf<typeof bunkiRailId>

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
