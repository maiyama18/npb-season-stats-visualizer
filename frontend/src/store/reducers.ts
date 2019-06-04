import { initialPlayersState, playersReducer, PlayersState } from './modules/players';
import { combineReducers, Reducer } from 'redux';
import { AnyAction } from './types';

export interface AppState {
  players: PlayersState;
}

export const initialState: AppState = {
  players: initialPlayersState,
};

export const rootReducer: Reducer<AppState, AnyAction> = combineReducers({
  players: playersReducer,
});
