import React, { Component } from 'react'
import { observer }  from 'mobx-react'
import HeaderContainer from './containers/HeaderContainer'
import TopNavContainer from './containers/TopNavContainer'
import HatListContainer from './containers/HatListContainer'
import HatEditContainer from './containers/HatEditContainer'
import FooterContainer from './containers/FooterContainer'
import './components/style.scss';

class App extends Component {

  render() {

    const {
      stores,
    } = this.props

    return (
      <div>
        <HeaderContainer stores={stores} />
        <TopNavContainer stores={stores} />
        <div>
          <HatEditContainer stores={stores} />
          <hr/>
          <HatListContainer stores={stores} />
        </div>
        <FooterContainer stores={stores} />
      </div>
    )
  }
}
export default observer(App)
