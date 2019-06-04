import React from 'react';
import { List, Icon, Header } from 'semantic-ui-react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch } from 'redux';
import { AppState } from '../../store/reducers';
import { Candidate } from '../../store/types';
import { selectPlayerThunk } from '../../store/modules/players';

const actions = { selectPlayerThunk };

interface StateProps {
  candidates: Candidate[];
}

type DispatchProps = typeof actions;

const SearchResults = (props: StateProps & DispatchProps) => (
  <div>
    <Header as="h5">検索結果</Header>
    <List divided relaxed>
      {props.candidates.map(c => (
        <List.Item key={c.id}>
          <List.Content floated="right">
            <a
              href="/"
              onClick={e => {
                e.preventDefault();
                props.selectPlayerThunk(c.id);
              }}
            >
              <Icon name="plus square" />
            </a>
          </List.Content>
          <List.Content>{c.name}</List.Content>
        </List.Item>
      ))}
    </List>
  </div>
);

const mapStateToProps = (state: AppState): StateProps => ({
  candidates: state.players.candidates,
});
const mapDispatchToProps = (dispatch: Dispatch): DispatchProps => ({
  ...bindActionCreators(actions, dispatch),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SearchResults);
