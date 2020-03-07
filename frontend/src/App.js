import React, { Component } from 'react'
import './App.css'
import { observer }  from 'mobx-react'
import HeaderContainer from './containers/HeaderContainer'
import TopNavContainer from './containers/TopNavContainer'
import LeftNavContainer from './containers/LeftNavContainer'
import HatListContainer from './containers/HatListContainer'
import HatEditContainer from './containers/HatEditContainer'
import FooterContainer from './containers/FooterContainer'

class App extends Component {

  render() {

    const {
      stores,
    } = this.props

    return (
      <div className="App">
        <HeaderContainer stores={stores} />
        <TopNavContainer stores={stores} />
        <LeftNavContainer stores={stores} />
        <div className="RightColumn" >
          <HatEditContainer stores={stores} />
          <HatListContainer stores={stores} />
        </div>
        <FooterContainer stores={stores} />
      </div>
    )
  }
}
export default observer(App)
