import React from 'react'
import PropTypes from 'prop-types'
import './HatList.css';
import Hat from './Hat';

const HatList = ({ hats }) => (
  <ul>
    {hats.map(hat =>
      <Hat color= {hat.color} name= {hat.name} size= {hat.size} />
    )}
  </ul>  
)

HatList.propTypes = {
  // onClick: PropTypes.func.isRequired,
  // completed: PropTypes.bool.isRequired,
  // text: PropTypes.string.isRequired
}

export default HatList
