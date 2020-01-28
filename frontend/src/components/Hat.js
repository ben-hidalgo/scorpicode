import React from 'react'
import PropTypes from 'prop-types'
import './Hat.css';

const Hat = ({ onClick, completed, text }) => (
  <p className="Hat-paragraph">hello</p>
)

Hat.propTypes = {
  // onClick: PropTypes.func.isRequired,
  // completed: PropTypes.bool.isRequired,
  // text: PropTypes.string.isRequired
}

export default Hat
