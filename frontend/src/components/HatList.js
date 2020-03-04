import React from 'react'
import Hat from './Hat'

function HatList(props) {
    return (
    <ul>
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
