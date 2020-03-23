import React from 'react'
import ReactDOM from 'react-dom'
import App from './containers/App'
import * as serviceWorker from './serviceWorker'

import commonStore from './stores/commonStore'
import authStore from './stores/authStore'
import hatStore from './stores/hatStore'
import languageStore from './stores/languageStore'

const stores = {
  commonStore,
  hatStore,
  languageStore,
  authStore,
}

// prevent anonymous access
if (authStore.decoded == null) {
  // TODO: inject value
  // window.location.href = 'http://localhost:8080/'
}

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()

ReactDOM.render(
  <App stores={stores}/>,
  document.getElementById("root")
)
