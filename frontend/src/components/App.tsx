import * as React from 'react';
import { Grid, Container, Divider } from 'semantic-ui-react';
import { Header } from './Header';
import SearchForm from './select/SearchForm';
import SearchResults from './select/SearchResults';
import SelectedPlayers from './select/SelectedPlayers';
import GraphControl from './graph/GraphControl';
import StatsGraph from './graph/StatsGraph';

export const App = () => {
  return (
    <div style={{ overflowY: 'scroll', overflowX: 'hidden' }}>
      <Header />
      <Container>
        <Grid divided>
          <Grid.Row>
            <Grid.Column width={4}>
              <div style={{ height: '90px' }}>
                <SearchForm />
              </div>
              <Divider />
              <div style={{ height: '260px', overflowY: 'scroll' }}>
                <SearchResults />
              </div>
              <Divider />
              <div style={{ height: '200px', overflowY: 'scroll' }}>
                <SelectedPlayers />
              </div>
            </Grid.Column>
            <Grid.Column width={12}>
              <div style={{ height: '65px' }}>
                <GraphControl />
              </div>
              <div style={{ height: '545px', background: 'grey' }}>
                <StatsGraph />
              </div>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Container>
    </div>
  );
};
