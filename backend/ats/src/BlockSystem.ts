import { BlockStateList } from './../spec/block_pb.d';
import { BlockRules } from './../types/BlockRule.d';
import { load } from 'js-yaml'
import { readFileSync } from 'fs'
import path from 'path'

const blockRules = load(readFileSync(path.join(__dirname, 'config/blockRule.yaml'), 'utf8')) as BlockRules;

export class BlockSystem {

  BlockState: BlockStateList;

  constructor(BlockState: BlockStateList) {
    this.BlockState = BlockState;
  }

  getBeUpdatedBlockByBlockId(block_id: string) {
    const blockRule = blockRules.sensors.find((blockRule) => blockRule.name === block_id);
    if (blockRule === undefined) {
      throw new Error(`Block ${block_id} is not found in config`);
    }
    return blockRule;
  }

  getBlockState(block_id: string) {
    return this.BlockState.getBlockStatesList.call(this.BlockState).find((block) => block.getBlockid.toString() === block_id);
  }
}