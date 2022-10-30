import { BlockRules } from './../types/BlockRule.d';
import{load}from 'js-yaml'
import{readFileSync}from'fs'
import path from 'path'

const blockRules = load(readFileSync(path.join(__dirname,'config/blockRule.yaml'), 'utf8')) as BlockRules;

export class BlockSystem{
  constructor(){

  }
  getBeUpdatedBlock(block_id: string){
    const blockRule = blockRules.sensors.find((blockRule) => blockRule.name === block_id);
    if(blockRule===undefined){
      throw new Error(`Block ${block_id} is not found in block-rules.json`);
    }
    return blockRule;
  }
}