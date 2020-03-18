import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import './TopNavContainer.css'
import { observer }  from 'mobx-react'


class TopNavContainer extends Component {

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
      <div className="TopNavContainer">
        <a href="./" className="LeftLink">Hats</a>
        <a href="./" className="LeftLink">Accounts</a>
        <a href="./" className="LeftLink">Comms</a>
        <a href="./" className="RightLink"><img src={authStore.picture} alt="profile"></img></a>
        <button onClick={() => {authStore.logout()}} type="button">Logout</button>
    </div>
    )
  }

}

export default observer(TopNavContainer)
