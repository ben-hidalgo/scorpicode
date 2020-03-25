import React, { Component } from 'react'
import {
  useParams
} from "react-router-dom";
import { observer }  from 'mobx-react'


class Hat extends Component {


  render() {

    return (
    <div className="container is-white">
      Hello
      <HatF />
    </div>
    )
  }

}

function HatF() {

  let { id } = useParams()
  return (
    <h3>Requested ID: {id}</h3>
    )

}

export default observer(Hat)
