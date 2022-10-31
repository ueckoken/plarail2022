import { BlockSystem } from '../src/BlockSystem';

describe('BlockSystem', () => {
  const blockSystem = new BlockSystem();

  it('BlockIDが指定されたときにBlockRuleを取得する', () => {
    const blockRule =
      blockSystem.getBeUpdatedBlockByBlockId('shinjuku_d1');
    expect(blockRule).toEqual({
      name: 'shinjuku_d1',
      block: 'shinjuku_sakurajosui_up',
      state: 'OPEN'
    });
  });

  it('不正なBlockIDが指定されたときにエラーを返す', () => {
    expect(() =>
      blockSystem.getBeUpdatedBlockByBlockId('none')
    ).toThrow(
      new Error("Block none is not found in config")
    );
  });
})