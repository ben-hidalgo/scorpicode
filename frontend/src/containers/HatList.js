import React, { Component } from 'react'
import { observer }  from 'mobx-react'
import Hat from '../components/Hat'


class HatList extends Component {

  render() {

    const {
      hats,
      listHats,
    } = this.props.stores.hatStore
    
    return (
      <div>
        <button onClick={listHats} type="button">List Hats</button>
        <ul>
          {
            hats.map(h => {
              return (
                <Hat stores={this.props.stores} hat={h} key={h.id} />
              )
            })
          }
        </ul>
      </div>
    )
  }

}
export default observer(HatList)
