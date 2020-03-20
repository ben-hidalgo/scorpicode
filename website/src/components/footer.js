import React from 'react'
import { FaTwitter, FaInstagram } from 'react-icons/fa'
import './style.scss'

const Footer = () => (
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
    </footer>
)

export default Footer
