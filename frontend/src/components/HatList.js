import Hat from './Hat';
import React from 'react';
import { inject, observer } from "mobx-react";

@inject("hatStore")
@observer
export default class HatList extends React.Component {

  render() {
    console.log(this.props)
    let hats = this.props.hats
    console.log(hats)
    if (!this.props.hats || this.props.hats.length === 0) {
      return (
        <div className="article-preview">
          No hats
        </div>
      )
    }


    return (
      <div>
        <ul>Hats:</ul>
        {
          this.props.hats.map(hat => {
            return (
              <Hat hat={hat} key={hat.id} />
            )
          })
        }
      </div>
    )
  }
}
