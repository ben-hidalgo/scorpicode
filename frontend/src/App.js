import React, { Component } from 'react';
import './App.css';
import { observer }  from 'mobx-react';

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
      </div>
    );
  }
}
export default observer(App);
