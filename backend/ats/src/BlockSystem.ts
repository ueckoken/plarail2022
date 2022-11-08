import { BlockStateList } from './../spec/block_pb.d';
import { load } from 'js-yaml'
import { readFileSync } from 'fs'
import path from 'path'

export class BlockSystem {

  BlockState: BlockStateList;

  constructor(BlockState: BlockStateList) {
    this.BlockState = BlockState;
  }
}