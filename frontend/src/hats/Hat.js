import React from 'react'
import {
  useParams
} from "react-router-dom";
import { observer }  from 'mobx-react'


const Hat = (props) => {

  let {
    hatStore,
  } = props.stores

  let { id } = useParams()

  let hat = hatStore.fetchHat(id)

  if (!hat) {
    // TODO
    return <div>not found</div>
  }

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

export default observer(Hat)
