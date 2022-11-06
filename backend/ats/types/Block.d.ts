import { BlockState } from '../spec/block_pb'

export interface Block {
  id: BlockState.BlockIdMap;
  state: BlockState.StateMap;
}

export interface BlockMap {
  [key: string]: Block;
}