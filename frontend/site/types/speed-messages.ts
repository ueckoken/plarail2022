import * as t from "io-ts"

export const trainId = t.union([
  t.literal("TAKAO"),
  t.literal("CHICHIBU"),
  t.literal("HAKONE"),
  t.literal("OKUTAMA"),
  t.literal("NIKKO"),
  t.literal("ENOSHIMA"),
  t.literal("KAMAKURA"),
  t.literal("YOKOSUKA"),
])
export type TrainId = t.TypeOf<typeof trainId>

export type SpeedMessage = {
  train_name: TrainId
  speed: number
}
