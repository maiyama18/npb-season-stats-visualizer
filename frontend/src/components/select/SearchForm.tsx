import React from 'react';
import { Form } from 'semantic-ui-react';
import { AppState } from '../../store/reducers';
import { bindActionCreators, Dispatch } from 'redux';
import { changePlayersType, changeQueryThunk } from '../../store/modules/players';
import { connect } from 'react-redux';
import { PlayersType } from '../../store/types';

const actions = { changePlayersType, changeQueryThunk };

interface StateProps {
  query: string;
  playersType: PlayersType;
}

type DispatchProps = typeof actions;

interface PlayerTypeOption {
  key: PlayersType;
  text: string;
  value: PlayersType;
}
const playersTypeOptions: PlayerTypeOption[] = [
  { key: 'batters', text: '打撃成績', value: 'batters' },
  { key: 'pitchers', text: '投球成績', value: 'pitchers' },
];

const SearchForm = (props: StateProps & DispatchProps) => (
  <Form>
    <Form.Field>
      <Form.Select
        options={playersTypeOptions}
        value={props.playersType}
        onChange={(_e, data: any) => {
          props.changePlayersType(data.value);
        }}
      />
    </Form.Field>
    <Form.Field>
      <input
        type="text"
        placeholder="選手名で検索"
        value={props.query}
        onChange={e => props.changeQueryThunk(e.target.value)}
      />
    </Form.Field>
  </Form>
);

const mapStateToProps = (state: AppState): StateProps => {
  return {
    query: state.players.query,
    playersType: state.players.playersType,
  };
};

const mapDispatchToProps = (dispatch: Dispatch): DispatchProps => ({
  ...bindActionCreators(actions, dispatch),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SearchForm);
