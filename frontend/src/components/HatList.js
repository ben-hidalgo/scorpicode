import React from 'react'
import PropTypes from 'prop-types'
import './HatList.css';

const HatList = ({ onClick, completed, text }) => (
  <p className="HatList-paragraph">HatList</p>
)

HatList.propTypes = {
  // onClick: PropTypes.func.isRequired,
  // completed: PropTypes.bool.isRequired,
  // text: PropTypes.string.isRequired
}

export default HatList
