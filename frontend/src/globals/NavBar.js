import React from 'react'
import { Link } from 'react-router-dom'
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext'
import logo from '../images/scorpicode210.png'

const NavBar = () => {

  const {
    authStore,
  } = React.useContext(StoreContext)

  return (
    <nav className="navbar is-light" role="navigation" aria-label="main navigation">
      <div className="navbar-brand">
        <a href="." className="navbar-item" >
          <img src={logo} alt="scorpicode logo" />
        </a>
      </div>

      <div className="navbar-menu is-active">
        <div className="navbar-start">
          <Link to="/hatsnew" className="navbar-item">Make Hats</Link>
          <Link to="/hats" className="navbar-item">View Hats</Link>
        </div>

        {/* TODO: include the hamburger and dynamic 'is-active' for mobile UX */}
        {/* TODO: https://bulma.io/documentation/components/navbar/           */}
        
        <div className="navbar-end">
        <div className="navbar-item">
            <figure className="image is-48x48px">
              <img src={authStore.picture} style={{width: 'auto'}} alt="profile" />
            </figure>
          </div>
          <div className="navbar-item">
            <span>Welcome, {authStore.givenName()}</span>
          </div>
          <div className="navbar-item">
            <div className="buttons">
              <button className="button" onClick={() => {authStore.logout()}} type="button">Logout</button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  )
}

export default observer(NavBar)
