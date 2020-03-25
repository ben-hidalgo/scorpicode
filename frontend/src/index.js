import React from 'react'
import ReactDOM from 'react-dom'
import { Route } from 'react-router'
import { HashRouter } from 'react-router-dom'
import Hero from './containers/Hero'
import Helmet from './components/Helmet'
import NavBar from './containers/NavBar'
import Hat from './containers/Hat'
import HatList from './containers/HatList'
import HatEdit from './containers/HatEdit'
import Footer from './containers/Footer'

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
let jsx = <span>You have been logged out.  Please login to continue.</span>

if (authStore.isLoggedIn()) {

  jsx = 
  <HashRouter>
    <Helmet />
    <NavBar stores={stores} />
    
    <Route path='/' render={() => <Hero stores={stores}/>}/>
    <br/>
    <Route path='/hats' render={() => <HatList stores={stores}/>}/>
    <Route path='/hatsview/:id' render={() => <Hat stores={stores}/>}/>
    <Route path='/hatsnew' render={() => <HatEdit stores={stores}/>}/>
    <br/>
    <Footer stores={stores} />
  </HashRouter>  

}

ReactDOM.render(jsx , document.getElementById("root"))
