import { applyMiddleware, compose, createStore } from 'redux';
import { initialState, rootReducer } from './reducers';
import thunk from 'redux-thunk';

const enhancers = [applyMiddleware(thunk)];

const reduxDevtoolsExtension = (window as any).__REDUX_DEVTOOLS_EXTENSION__;
if (process.env.NODE_ENV === 'development' && typeof reduxDevtoolsExtension === 'function') {
  enhancers.push(reduxDevtoolsExtension());
}
export const store = createStore(rootReducer, initialState, compose(...enhancers));
