import { BlockStateList } from '../spec/block_pb';
import { BlockState } from '../spec/block_pb';


class BlockStateClass {
  blockState: BlockStateList = new BlockStateList();

  constructor() {
    const stationsNames = Object.getOwnPropertyNames(BlockState.BlockId);
    stationsNames.forEach((stationName) => {
      const newState = new BlockState();
      newState.setBlockid(BlockState.BlockId[stationName as keyof typeof BlockState.BlockId]);
      newState.setState(BlockState.State.OPEN);
      this.blockState.addBlockStates(newState);
    })
  }

  getBlockState(): BlockStateList.AsObject {
    return this.blockState.toObject();
  }

  setBlockState(blockId: BlockState.BlockIdMap[keyof BlockState.BlockIdMap], state: BlockState.StateMap[keyof BlockState.StateMap]): void {
    const blockState = this.blockState.getBlockStatesList().find((blockState) => blockState.getBlockid() === blockId);
    if (blockState) {
      blockState.setState(state);
    }
  }

}

export default BlockStateClass;

const blockStateClass = new BlockStateClass();
blockStateClass.setBlockState(BlockState.BlockId.SAKURAJOSUI_B1, BlockState.State.CLOSE);
console.log(blockStateClass.getBlockState());