import * as React from 'react';
import { Grid, Container, Divider } from 'semantic-ui-react';
import { Header } from './Header';
import SearchForm from './select/SearchForm';

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
              <div style={{ height: '260px', overflowY: 'scroll', background: 'red' }} />
              <Divider />
              <div style={{ height: '200px', overflowY: 'scroll', background: 'green' }} />
            </Grid.Column>
            <Grid.Column width={12}>
              <div style={{ height: '65px', background: 'purple' }} />
              <div style={{ height: '545px', background: 'grey' }} />
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Container>
    </div>
  );
};
