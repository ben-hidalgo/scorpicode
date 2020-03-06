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

    return (
      <div className="TopNavContainer">
        <a href="./" className="LeftLink">Link</a>
        <a href="./" className="LeftLink">Link</a>
        <a href="./" className="LeftLink">Link</a>
        <a href="./" className="RightLink">Link</a>
    </div>
    )
  }

}

export default observer(TopNavContainer)
