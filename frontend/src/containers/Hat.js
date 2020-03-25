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
    <fieldset className="container">
      <div className="columns">

        <div className="column">
          <div className="field">
            <label className="label">ID:</label>
              {hat.id}
          </div>
          
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
    </fieldset>    
    )
}

export default observer(Hat)
