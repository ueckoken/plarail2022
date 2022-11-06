import { BlockSystem } from '../src/BlockSystem';
import { BlockState, BlockStateList } from '../spec/block_pb';

let state: BlockStateList;
state.addBlockStates(new BlockState());
console.log(state);

describe('BlockSystem', () => {
  const blockSystem = new BlockSystem(state);

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