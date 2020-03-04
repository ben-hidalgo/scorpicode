import React from 'react'
import Hat from './Hat'
import './HatList.css'

function HatList(props) {
    return (
    <ul className="HatList">
      {
        props.hats.map(hat => {
          return (
              <Hat hat={hat} key={hat.id} />
          )
        })
      }
    </ul>
    )
}

export default HatList
