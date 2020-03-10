import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import './HeaderContainer.css'
import { observer }  from 'mobx-react'


class HeaderContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      username: 'me@user.com',
    })
  }

  render() {

    return (
      <div className="HeaderContainer">
        <a href="/#" >
          <img src="scorpicode210.png" alt="scorpicode logo" />
        </a>
      </div>
    )
  }

}

export default observer(HeaderContainer)
