import * as React from 'react';
import { Container, Menu } from 'semantic-ui-react';

export const Header = () => {
  return (
    <Menu
      attached={true}
      inverted={true}
      borderless={true}
      size={'huge'}
      color={'orange'}
      style={{ marginBottom: '1rem' }}
    >
      <Container>
        <Menu.Item fitted={'horizontally'}>title</Menu.Item>
      </Container>
    </Menu>
  );
};
