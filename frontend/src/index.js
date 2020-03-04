import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import commonStore from './stores/commonStore';
// import { Provider } from 'mobx-react';

// TODO: setup mobx stores
const stores = {
  commonStore,
}
  
// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()

ReactDOM.render(
  <App stores={stores}/>,
  document.getElementById("root")
)
