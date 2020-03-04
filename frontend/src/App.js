import React, { Component } from 'react'
import './App.css'
import { observer }  from 'mobx-react'
import HatListContainer from './containers/HatListContainer'

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
        <HatListContainer stores={this.props.stores}/>
      </div>
    )
  }
}
export default observer(App)
