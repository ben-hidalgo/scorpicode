// import React from 'react';
import React, { Component } from 'react';
import './App.css';
import { observer }  from 'mobx-react';


// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">

//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//       </header>
//     </div>
//   );
// }

// export default App;

class App extends Component {

  onIncrement = () => {
    this.counter++;
  }
  onDecrement = () => {
    this.counter--;
  }
  render() {
    console.log(this.props.stores)
    console.log(this.props.stores.commonStore)
    console.log(this.props.stores.commonStore.counter)
    return (
      <div>
        <span>Counter: {this.props.counter}</span>
        
        <button onClick={this.onIncrement} type="button">Increment</button>
        <button onClick={this.onDecrement} type="button">Decrement</button>
      </div>
    );
  }
}
export default observer(App);
