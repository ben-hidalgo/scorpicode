import React, { Component } from 'react'
import { observer }  from 'mobx-react'
import Helmet from '../components/Helmet'
import NavBar from './NavBar'
import HatList from './HatList'
import HatEdit from './HatEdit'
import Footer from './Footer'
import '../style.scss';

class App extends Component {

  render() {

    const {
      stores,
    } = this.props

    return (
      <div>
        <Helmet />
        <NavBar stores={stores} />
        <HatEdit stores={stores} />
        <HatList stores={stores} />
        <Footer stores={stores} />
      </div>
    )
  }
}
export default observer(App)
