import React from 'react';
import './App.css';
import HatList from './HatList';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <a
          className="App-link"
          href="/"
        >
          Sign Out
        </a>
        <HatList hats= {[{color: "red", name: "bowler", size: 10}, {color: "blue", name: "cap", size: 12}]} />
      </header>
    </div>
  );
}

export default App;
