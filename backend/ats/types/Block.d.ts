export interface Block{
  id:string;
  state: BlockState;
}

type BlockState = 'error' | 'open' | 'close';