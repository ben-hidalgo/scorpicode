import ReactDOM from "react-dom";
import promiseFinally from "promise.prototype.finally";
import React from "react";
import { HashRouter } from "react-router-dom";
import { Provider } from "mobx-react";

import App from "./App";


import hatStore from "./stores/hatStore";

const stores = {
  hatStore,
};


// TODO: remove make and load from here
// hatStore.makeHat(10);
// hatStore.makeHat(12);

hatStore.loadHats()

// For easier debugging
window._____APP_STATE_____ = stores;

promiseFinally.shim();


ReactDOM.render(
  <Provider {...stores}>
    <HashRouter>
      <App />
    </HashRouter>
  </Provider>,
  document.getElementById("root")
);
