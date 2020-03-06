import React, { Component } from 'react'
import './App.css'
import { observer }  from 'mobx-react'
import HatListContainer from './containers/HatListContainer'
import HatEditContainer from './containers/HatEditContainer'

class App extends Component {

  render() {

    const {
      counter,
      onIncrement,
      onDecrement,
    } = this.props.stores.commonStore

    return (
      <div className="App">
        <span>Counter: {counter}</span>
        
        <button onClick={onIncrement} type="button">Increment</button>
        <button onClick={onDecrement} type="button">Decrement</button>
        <br/><br/><br/><br/><br/><br/>
        <HatEditContainer stores={this.props.stores}/>
        <HatListContainer stores={this.props.stores}/>
      </div>
    )
  }
}
export default observer(App)
