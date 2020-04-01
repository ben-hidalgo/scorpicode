import React from 'react'
import { observer }  from 'mobx-react'
import { FaTwitter, FaInstagram } from 'react-icons/fa'

const Footer = () => {

  return (
    <div>
      <div className="content has-text-centered">
        <p>
          <strong>Scorpicode</strong> by <a href="https://github.com/ben-hidalgo">Ben Hidalgo</a>
          <a href="https://www.instagram.com/scorpicode" >
            <FaInstagram size="29px" />
          </a>
          <a href="https://twitter.com/BenHidalgo8" >
              <FaTwitter size="29px" />
          </a>
        </p>
      </div>
    </div>
  )
}

export default observer(Footer)
