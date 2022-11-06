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

  updateBlock(blockId, blockState) {

  }
}