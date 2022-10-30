import { BlockSystem } from './../src/BlockSystem';

describe('BlockSystem', () => {
  it('BlockIDが指定されたときにBlockRuleを取得する', () => {
    const blockSystem = new BlockSystem();
    const blockRule = blockSystem.getBeUpdatedBlock('shinjuku_d1');
    expect(blockRule).toEqual({ name:'shinjuku_d1', block: 'shinjuku_sakurajosui_up', state: 'open' });
  });
})