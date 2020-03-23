import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import { observer }  from 'mobx-react'


class Header extends Component {

  constructor() {
    super()
    extendObservable(this, {
    })
  }

  render() {

    return (
      <div>
      </div>
    )
  }

}

export default observer(Header)
