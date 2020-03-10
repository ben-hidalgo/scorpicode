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
      <button onClick={() => {deleteHat(id, version)}} type="button">Delete</button>
      <label>ID: {id}</label><br/>
      <label>Inches: {inches}</label><br/>
      <label>Color: {color}</label><br/>
      <label>Style: {style}</label><br/>
      <label>Version: {version}</label><br/>
    </li>
    <br/>
  </div>
}

export default Hat
