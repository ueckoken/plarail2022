import { BlockState, Blocks } from '../spec/block_pb'

export interface Block {
  id: Blocks.BlockIdMap;
  state: BlockState.StateMap;
}