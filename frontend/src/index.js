import React from 'react'
import ReactDOM from 'react-dom'
import { Route } from 'react-router'
import { HashRouter } from 'react-router-dom'
import App from './containers/App'
import Helmet from './components/Helmet'
import NavBar from './containers/NavBar'
import HatList from './containers/HatList'
import HatEdit from './containers/HatEdit'
import Footer from './containers/Footer'
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
  <HashRouter>
    <Helmet />
    <NavBar stores={stores} />
    <Route path='/' render={() => <App stores={stores}/>}/>
    <Route path='/hats' render={() => <HatList stores={stores}/>}/>
    <Route path='/hatsnew' render={() => <HatEdit stores={stores}/>}/>
    <br/>
    <Footer stores={stores} />
  </HashRouter>,
  document.getElementById("root")
)
