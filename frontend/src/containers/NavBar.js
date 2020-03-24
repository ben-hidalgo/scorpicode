import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { extendObservable } from 'mobx'
import { observer }  from 'mobx-react'
import logo from '../images/scorpicode210.png';

class NavBar extends Component {

  constructor() {
    super()
    extendObservable(this, {
      username: 'me@user.com',
    })
  }

  render() {

    const {
      authStore,
    } = this.props.stores

    return (
      <nav className="navbar is-light" role="navigation" aria-label="main navigation">
        <div className="navbar-brand">
          <a href="." className="navbar-item" >
            <img src={logo} alt="scorpicode logo" />
          </a>
        </div>

        <div className="navbar-menu">
          <div className="navbar-start">
            <a href="." className="navbar-item">
              Home
            </a>

            <div className="navbar-item has-dropdown is-hoverable">
              <a href="." className="navbar-link">
                Hats
              </a>

              <div className="navbar-dropdown">
                <Link to="/hatsnew" className="navbar-item">Create A Hat</Link>
                <Link to="/hats" className="navbar-item">List My Hats</Link>
              </div>
            </div>
          </div>

          <div className="navbar-end">
          <div className="navbar-item">
              <figure className="image is-48x48px">
                <img src={authStore.picture} className="is-rounded" style={{width: 'auto'}} alt="profile" />
              </figure>
            </div>
            <div className="navbar-item">
              <span>Welcome, {authStore.payload.given_name}</span>
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

}

export default observer(NavBar)
