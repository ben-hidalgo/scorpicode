import React, { Component } from 'react';
import './HatListContainer.css';
import { observer }  from 'mobx-react';

class HatListContainer extends Component {

  render() {
    
    const {
      counter,
      onIncrement,
      onDecrement,
    } = this.props.stores.commonStore

    return (
      <div>
        <span>Counter: {counter}</span>
        
        <button onClick={onIncrement} type="button">Increment</button>
        <button onClick={onDecrement} type="button">Decrement</button>
      </div>
    );
  }
}
export default observer(HatListContainer);
