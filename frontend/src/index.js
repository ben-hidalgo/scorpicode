import React from 'react'
import ReactDOM from 'react-dom'
import { Route } from 'react-router'
import { HashRouter } from 'react-router-dom'

import StoreContext from './storeContext'

import Hero from './globals/Hero'
// import Helmet from './globals/Helmet'
import NavBar from './globals/NavBar'
import Footer from './globals/Footer'
import Socket from './globals/Socket'
import OrderView from './orders/OrderView'
import HatView from './hats/HatView'
import HatList from './hats/HatList'
import HatCreate from './hats/HatCreate'

import commonStore from './stores/commonStore'
import authStore from './stores/authStore'
import hatStore from './stores/hatStore'
import orderStore from './stores/orderStore'
import languageStore from './stores/languageStore'

const stores = {
  commonStore,
  hatStore,
  orderStore,
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
      <NavBar />
      <Socket />
      <Route path='/' render={() => <Hero />}/>
      <br/>
      <Route path='/hats' render={() => <HatList />}/>
      <Route path='/hatsview/:id' render={() => <HatView />}/>
      <Route path='/hatsnew' render={() => <HatCreate stores={stores}/>}/>
      <Route path='/ordersview/:id' render={() => <OrderView />}/>
      <br/>
      <Footer stores={stores} />
    </HashRouter>  
  </StoreContext.Provider>

}

ReactDOM.render(jsx , document.getElementById("root"))
