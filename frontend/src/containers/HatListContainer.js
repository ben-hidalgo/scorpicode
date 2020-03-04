import React, { Component } from 'react'
import './HatListContainer.css'
import { observer }  from 'mobx-react'
import HatList from '../components/HatList'


class HatListContainer extends Component {

  render() {

    const {
      hats,
      listHats,
    } = this.props.stores.hatStore
    
    return (
      <div>
        <button onClick={listHats} type="button">List Hats</button>
        <HatList hats={hats}/>
      </div>
    )
  }

}
export default observer(HatListContainer)
