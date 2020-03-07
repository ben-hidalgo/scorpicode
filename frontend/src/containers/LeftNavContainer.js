import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import './LeftNavContainer.css'
import { observer }  from 'mobx-react'


class LeftNavContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      username: 'me@user.com',
    })
  }

  render() {

    return (
      <div className="LeftNavContainer">
        <a href="./" >Context</a>
        <a href="./" >Menu</a>
        <a href="./" >Actions</a>
        <a href="./" >Here</a>
    </div>
    )
  }

}

export default observer(LeftNavContainer)
