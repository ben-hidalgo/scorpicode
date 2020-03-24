import React, { Component } from 'react'
import { observer }  from 'mobx-react'
import '../style.scss';

class App extends Component {

  render() {

    return (
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Welcome to Scorpicode
          </h1>
          <h2 class="subtitle">
            Please select an option from the menu above
          </h2>
        </div>
      </div>
    </section>)
  }
}
export default observer(App)
