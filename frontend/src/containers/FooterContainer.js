import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import './FooterContainer.css'
import { observer }  from 'mobx-react'


class FooterContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      username: 'me@user.com',
    })
  }

  render() {

    return (
      <div className="FooterContainer">
        <h1>Footer</h1>
      </div>
    )
  }

}

export default observer(FooterContainer)
