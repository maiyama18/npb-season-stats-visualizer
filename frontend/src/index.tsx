import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { App } from './components/App';
import 'semantic-ui-css/semantic.min.css';
import { Provider } from 'react-redux';
import { store } from './store/store';

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
