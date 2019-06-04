import React from 'react';
import { Form } from 'semantic-ui-react';
import { bindActionCreators, Dispatch } from 'redux';
import { connect } from 'react-redux';
import { StatType } from '../../store/types';
import { changeGraphStat } from '../../store/modules/players';
import { AppState } from '../../store/reducers';

const actions = { changeGraphStat };

interface StateProps {
  selectedStatOption: StatOption;
  statOptions: StatOption[];
}

type DispatchProps = typeof actions;

interface StatOption {
  key: StatType;
  text: string;
  value: StatType;
}
const pitcherStatOptions: StatOption[] = [
  { key: 'game', value: 'game', text: '登板' },
  { key: 'era', value: 'era', text: '防御率' },
  { key: 'gameStart', value: 'gameStart', text: '先発' },
  { key: 'complete', value: 'complete', text: '完投' },
  { key: 'shutOut', value: 'shutOut', text: '完封' },
  { key: 'qualityStart', value: 'qualityStart', text: 'QS' },
  { key: 'win', value: 'win', text: '勝利' },
  { key: 'lose', value: 'lose', text: '敗戦' },
  { key: 'hold', value: 'hold', text: 'ホールド' },
  { key: 'holdPoint', value: 'holdPoint', text: 'HP' },
  { key: 'save', value: 'save', text: 'セーブ' },
  { key: 'winPercent', value: 'winPercent', text: '勝率' },
  { key: 'inning', value: 'inning', text: '投球回' },
  { key: 'hit', value: 'hit', text: '被安打' },
  { key: 'homeRun', value: 'homeRun', text: '被本塁打' },
  { key: 'strikeOut', value: 'strikeOut', text: '奪三振' },
  { key: 'strikeOutPercent', value: 'strikeOutPercent', text: '奪三振率' },
  { key: 'walk', value: 'walk', text: '与四球' },
  { key: 'hitByPitch', value: 'hitByPitch', text: '与死球' },
  { key: 'wildPitch', value: 'wildPitch', text: '暴投' },
  { key: 'balk', value: 'balk', text: 'ボーク' },
  { key: 'run', value: 'run', text: '失点' },
  { key: 'earnedRun', value: 'earnedRun', text: '自責点' },
  { key: 'average', value: 'average', text: '被打率' },
  { key: 'kbb', value: 'kbb', text: 'K/BB' },
  { key: 'whip', value: 'whip', text: 'WHIP' },
];

const batterStatOptions: StatOption[] = [
  { key: 'game', value: 'game', text: '試合' },
  { key: 'average', value: 'average', text: '打率' },
  { key: 'plateAppearance', value: 'plateAppearance', text: '打席' },
  { key: 'atBat', value: 'atBat', text: '打数' },
  { key: 'hit', value: 'hit', text: '安打' },
  { key: 'double', value: 'double', text: '二塁打' },
  { key: 'triple', value: 'triple', text: '三塁打' },
  { key: 'homeRun', value: 'homeRun', text: '本塁打' },
  { key: 'totalBase', value: 'totalBase', text: '塁打' },
  { key: 'runBattedIn', value: 'runBattedIn', text: '打点' },
  { key: 'run', value: 'run', text: '得点' },
  { key: 'strikeOut', value: 'strikeOut', text: '三振' },
  { key: 'walk', value: 'walk', text: '四球' },
  { key: 'hitByPitch', value: 'hitByPitch', text: '死球' },
  { key: 'sacrifice', value: 'sacrifice', text: '犠打' },
  { key: 'sacrificeFly', value: 'sacrificeFly', text: '犠飛' },
  { key: 'stolenBase', value: 'stolenBase', text: '盗塁' },
  { key: 'caughtStealing', value: 'caughtStealing', text: '盗塁死' },
  { key: 'doublePlay', value: 'doublePlay', text: '併殺打' },
  { key: 'onBasePercent', value: 'onBasePercent', text: '出塁率' },
  { key: 'sluggingPercent', value: 'sluggingPercent', text: '長打率' },
  { key: 'ops', value: 'ops', text: 'OPS' },
  { key: 'averageWithScoringPosition', value: 'averageWithScoringPosition', text: '得点圏打率' },
  { key: 'error', value: 'error', text: '失策' },
];

const GraphControl = (props: StateProps & DispatchProps) => {
  return (
    <Form>
      <Form.Group>
        <Form.Field>
          <label>成績</label>
          <Form.Select
            options={props.statOptions}
            value={props.selectedStatOption.value}
            onChange={(_e, data: any) => props.changeGraphStat(data.value)}
            style={{ minWidth: '9em' }}
          />
        </Form.Field>
      </Form.Group>
    </Form>
  );
};

const mapStateToProps = (state: AppState): StateProps => {
  const { graphStat, playersType } = state.players;

  const statOptions = playersType === 'pitchers' ? pitcherStatOptions : batterStatOptions;
  const selectedStatOption = statOptions.find(o => o.value === graphStat) || statOptions[0];

  return {
    statOptions,
    selectedStatOption,
  };
};
const mapDispatchToProps = (dispatch: Dispatch): DispatchProps => ({
  ...bindActionCreators(actions, dispatch),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(GraphControl);
