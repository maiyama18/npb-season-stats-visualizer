import React from 'react';
import { List, Icon, Header } from 'semantic-ui-react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch } from 'redux';
import { AppState } from '../../store/reducers';
import { Candidate } from '../../store/types';

interface StateProps {
  candidates: Candidate[];
}

const SearchResults = (props: StateProps) => (
  <div>
    <Header as="h6">検索結果</Header>
    <List divided relaxed>
      {props.candidates.map(c => (
        <List.Item key={c.id}>
          <List.Content floated="right">
            <a
              href="/"
              onClick={e => {
                e.preventDefault();
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
const mapDispatchToProps = (dispatch: Dispatch) => ({
  ...bindActionCreators({}, dispatch),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SearchResults);
