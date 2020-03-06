import React, { Component } from 'react'
import './App.css'
import { observer }  from 'mobx-react'
import HatListContainer from './containers/HatListContainer'
import HatEditContainer from './containers/HatEditContainer'

class App extends Component {

  render() {

    const {
      stores,
    } = this.props

    return (
      <div className="App">
        <HatEditContainer stores={stores}/>
        <HatListContainer stores={stores}/>
      </div>
    )
  }
}
export default observer(App)
