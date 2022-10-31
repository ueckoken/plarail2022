import { BlockSystem } from '../src/BlockSystem';
import { BlockStateList, BlockState, Blocks } from '../spec/block_pb.d';

const state: BlockState[] = new BlockStateList()
const newState = new BlockState();
newState.setState(BlockState.State.OPEN)
newState.setBlock((): Blocks => { block: Blocks = return new Blocks(); Blocks.setBlockid(Blocks.BlockId.SHINJUKU_B1); return Blocks });


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

export function setId(arg0: string) {
  throw new Error('Function not implemented.');
}

export function setId(arg0: string) {
  throw new Error('Function not implemented.');
}

