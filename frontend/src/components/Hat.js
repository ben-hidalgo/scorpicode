import React from 'react'
import PropTypes from 'prop-types'
import './Hat.css';

const Hat = ({ color, name, size }) => (
  <div>
    <p className="Hat-paragraph">Color: {color}</p>
    <p className="Hat-paragraph">Name: {name}</p>
    <p className="Hat-paragraph">Size: {size}</p>
  </div>
)

Hat.propTypes = {
  color: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  size: PropTypes.number.isRequired,
  // onClick: PropTypes.func.isRequired,
  // completed: PropTypes.bool.isRequired,
  // text: PropTypes.string.isRequired
}

export default Hat
