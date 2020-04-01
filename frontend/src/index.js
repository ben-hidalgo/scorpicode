import React from 'react'
import ReactDOM from 'react-dom'
import { Route } from 'react-router'
import { HashRouter } from 'react-router-dom'

import StoreContext from './storeContext';

import Hero from './globals/Hero'
// import Helmet from './globals/Helmet'
import NavBar from './globals/NavBar'
import Footer from './globals/Footer'
import Hat from './hats/Hat'
import HatList from './hats/HatList'
import HatCreate from './hats/HatCreate'

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
  <StoreContext.Provider value={stores}>
    <HashRouter>
      {/*<Helmet />  disabled due to componentWillMount usage warning */}
      <NavBar stores={stores} />
      
      <Route path='/' render={() => <Hero stores={stores}/>}/>
      <br/>
      <Route path='/hats' render={() => <HatList stores={stores}/>}/>
      <Route path='/hatsview/:id' render={() => <Hat stores={stores}/>}/>
      <Route path='/hatsnew' render={() => <HatCreate stores={stores}/>}/>
      <br/>
      <Footer stores={stores} />
    </HashRouter>  
  </StoreContext.Provider>

}

ReactDOM.render(jsx , document.getElementById("root"))
