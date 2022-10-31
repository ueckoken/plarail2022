import { BlockIdMap, StateMap } from './../spec/block_pb.d';
import { BlockRules } from './../types/BlockRule.d';
import { load } from 'js-yaml'
import { readFileSync } from 'fs'
import path from 'path'
import { Block } from '../types/Block';

const blockRules = load(readFileSync(path.join(__dirname, 'config/blockRule.yaml'), 'utf8')) as BlockRules;

export class BlockSystem {

  BlockState: Block[] = [];

  constructor(BlockState: Block[]) {
    this.BlockState = BlockState;
  }

  getBeUpdatedBlockByBlockId(block_id: string) {
    const blockRule = blockRules.sensors.find((blockRule) => blockRule.name === block_id);
    if (blockRule === undefined) {
      throw new Error(`Block ${block_id} is not found in config`);
    }
    return blockRule;
  }

  getBlockState(block_id: BlockIdMap) {
    return this.BlockState.find((block) => block.id === block_id);
  }

  updateBlockState(block_id: BlockIdMap, state: StateMap) {
    const block = this.getBlockState(block_id);
    if (block === undefined) {
      throw new Error(`Block ${block_id} is not found in BlockState`);
    }
    block.state = state;
  }
}