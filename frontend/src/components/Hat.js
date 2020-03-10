import React from 'react'
import './Hat.css'
// import { action } from 'mobx'

function Hat(props) {

  const {
    id, 
    inches, 
    color, 
    style, 
    version,
  } = props.hat

  const {
    deleteHat,
  } = props.stores.hatStore

  return <div>
    <li className="Hat">
      <div>ID: {id}</div>
      <div>Inches: {inches}</div>
      <div>Color: {color}</div>
      <div>Style: {style}</div>
      <div>Version: {version}</div>
    </li>
    
    <button onClick={() => {deleteHat(id, version)}} type="button">Delete</button>
  </div>
}

export default Hat
