export type AnyAction = Action<string> | PayloadedAction<string, any>;

export interface Action<T> {
  type: T;
}
export interface PayloadedAction<T, P> {
  type: T;
  payload: P;
}

export type PlayersType = 'batters' | 'pitchers';

export interface Candidate {
  id: number;
  name: string;
}

export interface Player {
  id: number;
  name: string;
  stats: Stats;
}

export type Stats = BatterStats | PitcherStats;

interface PitcherStats {
  game: Stat;
  era: Stat;
  gameStart: Stat;
  complete: Stat;
  shutOut: Stat;
  qualityStart: Stat;
  win: Stat;
  lose: Stat;
  hold: Stat;
  holdPoint: Stat;
  save: Stat;
  winPercent: Stat;
  inning: Stat;
  hit: Stat;
  homeRun: Stat;
  strikeOut: Stat;
  strikeOutPercent: Stat;
  walk: Stat;
  hitByPitch: Stat;
  wildPitch: Stat;
  balk: Stat;
  run: Stat;
  earnedRun: Stat;
  average: Stat;
  kbb: Stat;
  whip: Stat;
}

interface BatterStats {
  game: Stat;
  average: Stat;
  plateAppearance: Stat;
  atBat: Stat;
  hit: Stat;
  double: Stat;
  triple: Stat;
  homeRun: Stat;
  totalBase: Stat;
  runBattedIn: Stat;
  run: Stat;
  strikeOut: Stat;
  walk: Stat;
  hitByPitch: Stat;
  sacrifice: Stat;
  sacrificeFly: Stat;
  stolenBase: Stat;
  caughtStealing: Stat;
  doublePlay: Stat;
  onBasePercent: Stat;
  sluggingPercent: Stat;
  ops: Stat;
  averageWithScoringPosition: Stat;
  error: Stat;
}

interface Stat {
  dates: string[];
  values: number[];
}

export type StatType = PitcherStatType | BatterStatType;

export type PitcherStatType =
  | 'game'
  | 'era'
  | 'gameStart'
  | 'complete'
  | 'shutOut'
  | 'qualityStart'
  | 'win'
  | 'lose'
  | 'hold'
  | 'holdPoint'
  | 'save'
  | 'winPercent'
  | 'inning'
  | 'hit'
  | 'homeRun'
  | 'strikeOut'
  | 'strikeOutPercent'
  | 'walk'
  | 'hitByPitch'
  | 'wildPitch'
  | 'balk'
  | 'run'
  | 'earnedRun'
  | 'average'
  | 'kbb'
  | 'whip';

export type BatterStatType =
  | 'game'
  | 'average'
  | 'plateAppearance'
  | 'atBat'
  | 'hit'
  | 'double'
  | 'triple'
  | 'homeRun'
  | 'totalBase'
  | 'runBattedIn'
  | 'run'
  | 'strikeOut'
  | 'walk'
  | 'hitByPitch'
  | 'sacrifice'
  | 'sacrificeFly'
  | 'stolenBase'
  | 'caughtStealing'
  | 'doublePlay'
  | 'onBasePercent'
  | 'sluggingPercent'
  | 'ops'
  | 'averageWithScoringPosition'
  | 'error';
