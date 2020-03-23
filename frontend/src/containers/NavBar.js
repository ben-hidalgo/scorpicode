import React, { Component } from 'react'
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
                More
              </a>

              <div className="navbar-dropdown">
                <a href="." className="navbar-item">
                  About
                </a>
                <a href="." className="navbar-item">
                  Jobs
                </a>
                <a href="." className="navbar-item">
                  Contact
                </a>
                <hr className="navbar-divider" />
                <a href="." className="navbar-item">
                  Report an issue
                </a>
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
