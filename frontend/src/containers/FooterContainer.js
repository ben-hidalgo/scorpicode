import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import { observer }  from 'mobx-react'
import { FaTwitter, FaInstagram } from 'react-icons/fa'

class FooterContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      username: 'me@user.com',
    })
  }

  render() {

    return (
      <footer className="footer columns is-vcentered is-centered">
      <span>
          &copy; Scorpicode 2020
      </span>
      <span>&nbsp;&nbsp;&nbsp;</span>
      <a href="https://www.instagram.com/scorpicode" >
          <FaInstagram className="no-underline" size="29px" />
      </a>
      <a href="https://twitter.com/BenHidalgo8" >
          <FaTwitter size="29px" />
      </a>
  </footer>)
}

}

export default observer(FooterContainer)
