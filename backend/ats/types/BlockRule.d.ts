import { BlockState } from "./Block"

export interface BlockRule{
  name: string
  block: string
  state: BlockState
}

export interface BlockRules{
  sensors: BlockRule[]
}