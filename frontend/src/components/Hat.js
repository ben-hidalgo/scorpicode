import React from 'react'

function Hat(props) {
  const {
    id, 
    inches, 
    color, 
    style, 
    version,
  } = props.hat

  return <li>
    <div>ID: {id}</div>
    <div>Inches: {inches}</div>
    <div>Color: {color}</div>
    <div>Style: {style}</div>
    <div>Version: {version}</div>
  </li>
}

export default Hat
