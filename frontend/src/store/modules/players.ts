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
  | PayloadedAction<'CHANGE_PLAYERS_TYPE', { playersType: PlayersType }>
  | PayloadedAction<'START_SELECT', { id: number }>
  | PayloadedAction<'FINISH_SELECT_SUCCESS', { player: Player }>
  | Action<'FINISH_SELECT_FAILURE'>
  | PayloadedAction<'UNSELECT_PLAYER', { id: number }>;

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
const startSelect = (id: number): PlayersAction => ({
  type: 'START_SELECT',
  payload: { id },
});
const finishSelectSuccess = (player: Player): PlayersAction => ({
  type: 'FINISH_SELECT_SUCCESS',
  payload: {
    player,
  },
});
const finishSelectFailure = (): PlayersAction => ({
  type: 'FINISH_SELECT_FAILURE',
});
export const unselectPlayer = (id: number): PlayersAction => ({
  type: 'UNSELECT_PLAYER',
  payload: {
    id,
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
        query: '',
        candidates: [],
        selectedPlayers: [],
      };
    case 'START_SELECT':
      return {
        ...state,
        selecting: true,
      };
    case 'FINISH_SELECT_SUCCESS':
      return {
        ...state,
        query: '',
        candidates: [],
        selectedPlayers: [...state.selectedPlayers, action.payload.player],
        selecting: false,
      };
    case 'FINISH_SELECT_FAILURE':
      return {
        ...state,
        selecting: false,
      };
    case 'UNSELECT_PLAYER':
      return {
        ...state,
        selectedPlayers: state.selectedPlayers.filter(p => p.id != action.payload.id),
      };
    default:
      return state;
  }
};

const axios = axiosBase.create({
  baseURL: 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
  responseType: 'json',
});

export const changeQueryThunk = (query: string) => async (dispatch: Dispatch, getState: () => AppState) => {
  if (query === '') {
    dispatch(startSearch(query));
    dispatch(finishSearchSuccess([]));
    return;
  }

  dispatch(startSearch(query));
  try {
    const { playersType } = getState().players;
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

export const selectPlayerThunk = (id: number) => async (dispatch: Dispatch, getState: () => AppState) => {
  dispatch(startSelect(id));
  try {
    const { playersType } = getState().players;
    const playerResp: AxiosResponse = await axios.get(`/api/stats/${playersType}/${id}`);
    const player: Player = {
      id: playerResp.data.player.id,
      name: playerResp.data.player.name,
      stats: playerResp.data.stats,
    };
    dispatch(finishSelectSuccess(player));
  } catch (err) {
    console.error(err);
    dispatch(finishSelectFailure());
  }
};
