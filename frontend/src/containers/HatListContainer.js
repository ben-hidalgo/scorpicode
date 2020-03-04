import React, { Component } from 'react'
import './HatListContainer.css'
import { observer }  from 'mobx-react'
import HatList from '../components/HatList'

class HatListContainer extends Component {

  render() {
    
    return (
      <div>
        <HatList hats={this.props.stores.hatStore.hats}/>
      </div>
    )
  }
}
export default observer(HatListContainer)
