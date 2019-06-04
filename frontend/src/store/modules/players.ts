import { Action, Candidate, PayloadedAction, Player, PlayersType } from '../types';
import { Dispatch, Reducer } from 'redux';
import axiosBase, { AxiosResponse } from 'axios';
import { AppState } from '../reducers';

export interface PlayersState {
  playersType: PlayersType;
  query: string;
  searching: boolean;
  candidates: Candidate[];
  selecting: boolean;
  selectedPlayers: Player[];
}

export type PlayersAction =
  | PayloadedAction<'START_SEARCH', { query: string }>
  | PayloadedAction<'FINISH_SEARCH_SUCCESS', { candidates: Candidate[] }>
  | Action<'FINISH_SEARCH_FAILURE'>
  | PayloadedAction<'CHANGE_PLAYERS_TYPE', { playersType: PlayersType }>;

export const initialPlayersState: PlayersState = {
  playersType: 'batters',
  query: '',
  searching: false,
  candidates: [],
  selecting: false,
  selectedPlayers: [],
};

const startSearch = (query: string): PlayersAction => ({
  type: 'START_SEARCH',
  payload: { query },
});
const finishSearchSuccess = (candidates: Candidate[]): PlayersAction => ({
  type: 'FINISH_SEARCH_SUCCESS',
  payload: {
    candidates,
  },
});
const finishSearchFailure = (): PlayersAction => ({
  type: 'FINISH_SEARCH_FAILURE',
});
export const changePlayersType = (playersType: PlayersType): PlayersAction => ({
  type: 'CHANGE_PLAYERS_TYPE',
  payload: {
    playersType,
  },
});

export const playersReducer: Reducer<PlayersState, PlayersAction> = (
  state: PlayersState = initialPlayersState,
  action: PlayersAction
): PlayersState => {
  switch (action.type) {
    case 'START_SEARCH':
      return {
        ...state,
        query: action.payload.query,
        searching: true,
      };
    case 'FINISH_SEARCH_SUCCESS':
      return {
        ...state,
        candidates: action.payload.candidates,
        searching: false,
      };
    case 'FINISH_SEARCH_FAILURE':
      return {
        ...state,
        searching: false,
      };
    case 'CHANGE_PLAYERS_TYPE':
      return {
        ...state,
        playersType: action.payload.playersType,
      };
    default:
      return state;
  }
};

export const changeQueryThunk = (query: string) => async (dispatch: Dispatch, getState: () => AppState) => {
  if (query === '') {
    dispatch(startSearch(query));
    dispatch(finishSearchSuccess([]));
    return;
  }

  dispatch(startSearch(query));
  try {
    const { playersType } = getState().players;
    const axios = axiosBase.create({
      baseURL: 'http://localhost:8080',
      headers: {
        'Content-Type': 'application/json',
      },
      responseType: 'json',
    });
    const candidatesResp: AxiosResponse = await axios.get(`/api/search/${playersType}`, {
      params: {
        query,
      },
    });
    const candidates: Candidate[] = candidatesResp.data.players;
    dispatch(finishSearchSuccess(candidates));
  } catch (err) {
    console.error(err);
    dispatch(finishSearchFailure());
  }
};
