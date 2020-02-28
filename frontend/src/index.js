import ReactDOM from "react-dom";
import promiseFinally from "promise.prototype.finally";
import React from "react";
import { HashRouter } from "react-router-dom";
import { useStrict } from "mobx";
import { Provider } from "mobx-react";

import App from "./App";

import articlesStore from "./stores/articlesStore";
import hatStore from "./stores/hatStore";
import commentsStore from "./stores/commentsStore";
import authStore from "./stores/authStore";
import commonStore from "./stores/commonStore";
import editorStore from "./stores/editorStore";
import userStore from "./stores/userStore";
import profileStore from "./stores/profileStore";

const stores = {
  articlesStore,
  hatStore,
  commentsStore,
  authStore,
  commonStore,
  editorStore,
  userStore,
  profileStore
};


// TODO: remove make and load from here
// hatStore.makeHat(10);
// hatStore.makeHat(12);

hatStore.loadHats()

// For easier debugging
window._____APP_STATE_____ = stores;

promiseFinally.shim();

useStrict(true);

ReactDOM.render(
  <Provider {...stores}>
    <HashRouter>
      <App />
    </HashRouter>
  </Provider>,
  document.getElementById("root")
);
