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
        <div className="copyright">&copy; Scorpicode 2020</div>
        <a href="https://www.instagram.com/scorpicode" >
          <img src="instagram504.png" alt="instagram logo" />
        </a>
        <a href="https://twitter.com/BenHidalgo8" >
          <img src="twitter400.png" alt="twitter logo" />
        </a>
      </div>
    )
  }

}

export default observer(FooterContainer)
