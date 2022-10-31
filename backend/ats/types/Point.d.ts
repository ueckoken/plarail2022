export interface Block{
  id:string;
  state: BlockState;
}

type BlockState = 'UNKNOWN' | 'ON' | 'OFF';