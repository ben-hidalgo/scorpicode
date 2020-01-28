import React from 'react';
import logo from '../logo.svg';
import './App.css';
import Hat from './Hat';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="/"
        >
          Sign Out
        </a>
        <Hat />
      </header>
    </div>
  );
}

export default App;
