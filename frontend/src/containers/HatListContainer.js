import React, { Component } from 'react'
import './HatListContainer.css'
import { observer }  from 'mobx-react'
import HatList from '../components/HatList'


class HatListContainer extends Component {

  render() {
    
    return (
      <div>
        <button onClick={this.props.stores.hatStore.listHats} type="button">List Hats</button>
        <HatList hats={this.props.stores.hatStore.hats}/>
      </div>
    )
  }

}
export default observer(HatListContainer)
