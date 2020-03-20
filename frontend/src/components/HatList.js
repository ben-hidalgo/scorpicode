import React from 'react'
import Hat from './Hat'

function HatList(props) {
    return (
    <ul className="HatList">
      {
        props.hats.map(hat => {
          return (
            <Hat stores={props.stores} hat={hat} key={hat.id} />
          )
        })
      }
    </ul>
    )
}

export default HatList
