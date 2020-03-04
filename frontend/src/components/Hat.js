import React from 'react'
import './Hat.css'

function Hat(props) {
  const {
    id, 
    inches, 
    color, 
    style, 
    version,
  } = props.hat

  return <li className="Hat">
    <div>ID: {id}</div>
    <div>Inches: {inches}</div>
    <div>Color: {color}</div>
    <div>Style: {style}</div>
    <div>Version: {version}</div>
  </li>
}

export default Hat
