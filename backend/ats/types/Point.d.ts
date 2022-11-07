import { Stations, RequestSync } from '../spec/statesync_pb'

export interface Point {
  id: Stations.StationIdMap;
  state: RequestSync.StateMap;
}