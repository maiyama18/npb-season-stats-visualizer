import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch } from 'redux';
import { List, Icon, Loader, Header } from 'semantic-ui-react';
import { unselectPlayer } from '../../store/modules/players';
import { Player } from '../../store/types';
import { AppState } from '../../store/reducers';

const actions = { unselectPlayer };

interface StateProps {
  selecting: boolean;
  selectedPlayers: Player[];
}

type DispatchProps = typeof actions;

const SelectedPlayers = (props: StateProps & DispatchProps) => (
  <div>
    <Header as="h5">表示中</Header>
    <List divided>
      {props.selectedPlayers.map(p => (
        <List.Item key={p.id}>
          <List.Content floated="right">
            <a
              href="/"
              onClick={e => {
                e.preventDefault();
                props.unselectPlayer(p.id);
              }}
            >
              <Icon name="minus square" />
            </a>
          </List.Content>
          <List.Content>{p.name}</List.Content>
        </List.Item>
      ))}
      <Loader active={props.selecting} inline="centered" />
    </List>
  </div>
);

const mapStateToProps = (state: AppState): StateProps => ({
  selecting: state.players.selecting,
  selectedPlayers: state.players.selectedPlayers,
});
const mapDispatchToProps = (dispatch: Dispatch): DispatchProps => ({
  ...bindActionCreators(actions, dispatch),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SelectedPlayers);
