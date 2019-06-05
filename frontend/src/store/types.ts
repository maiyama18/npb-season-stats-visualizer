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

export interface PitcherStats {
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

export interface BatterStats {
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

export interface Stat {
  dates: string[];
  values: number[];
}

export type StatType = PitcherStatType | BatterStatType;

export type PitcherStatType = keyof PitcherStats;
export type BatterStatType = keyof BatterStats;
