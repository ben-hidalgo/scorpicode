import React, { Component } from 'react'
import { observer }  from 'mobx-react'
import Header from './Header'
import NavBar from './NavBar'
import HatList from './HatList'
import HatEdit from './HatEdit'
import Footer from './Footer'
import '../components/style.scss';

class App extends Component {

  render() {

    const {
      stores,
    } = this.props

    return (
      <div>
        <Header stores={stores} />
        <NavBar stores={stores} />
        <div>
          <HatEdit stores={stores} />
          <hr/>
          <HatList stores={stores} />
        </div>
        <Footer stores={stores} />
      </div>
    )
  }
}
export default observer(App)
