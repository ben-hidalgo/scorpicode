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
  hats: PropTypes.arrayOf(PropTypes.shape({
    size: PropTypes.number.isRequired,
    color: PropTypes.string.isRequired,
    name: PropTypes.string.isRequired
  }).isRequired).isRequired,
}

export default HatList
