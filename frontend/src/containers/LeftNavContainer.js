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
      </div>
    )
  }

}

export default observer(LeftNavContainer)
