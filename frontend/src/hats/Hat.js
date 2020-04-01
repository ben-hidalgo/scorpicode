import React from 'react'
import { useParams } from "react-router-dom";
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext';

// TODO: rename HatView
const Hat = () => {

  let {
    hatStore,
  } = React.useContext(StoreContext)

  let { id } = useParams()

  let hat = hatStore.fetchHat(id)

  if (!hat) {
    return <NotFound id={id} />
  }

  // TODO: add edit feature (link here)

  return (
    <div className="container">
      <div className="card">
        <header className="card-header is-danger">
          <p className="card-header-title">
            Hat: {hat.id}
          </p>
        </header>
        <div className="card-content">
          <div className="content">
            <div className="columns">

              <div className="column">              
                <div className="field">
                  <label className="label">Inches:</label>
                  {hat.inches}
                </div>
                <div className="field">
                  <label className="label">Version:</label>
                  {hat.version}
                </div>
              </div>

              <div className="column">
                <div className="field">
                  <label className="label">Color:</label>
                  {hat.color}
                </div>
                <div className="field">
                  <label className="label">Style:</label>
                  {hat.style}
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

const NotFound = (props) => {
  return (
    <div className="container">
      <div className="card">
        <header className="card-header is-danger">
          <p className="card-header-title">
            Hat: {props.id} not found
          </p>
        </header>
      </div>
    </div>
  )
} // NotFound

export default observer(Hat)
