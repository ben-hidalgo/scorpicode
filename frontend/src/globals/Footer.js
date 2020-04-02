import React from 'react'
import { observer }  from 'mobx-react'
import { FaTwitter, FaInstagram } from 'react-icons/fa'

const Footer = () => {

  return (
    <footer>
      <div className="content has-text-centered">
        <p>
          <a href="/">Scorpicode&nbsp;</a>
          ~&nbsp;
          <a href="https://github.com/ben-hidalgo">Ben Hidalgo&nbsp;</a>
        </p>
        <div>
          <a href="https://www.instagram.com/scorpicode" >
            <FaInstagram size="29px" />
          </a>
          <a href="https://twitter.com/BenHidalgo8" >
              <FaTwitter size="29px" />
          </a>
        </div>
      </div>
    </footer>
  )
}

export default observer(Footer)
