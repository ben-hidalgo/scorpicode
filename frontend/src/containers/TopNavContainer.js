import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import { observer }  from 'mobx-react'
// import logo from '/public/scorpicode210.png';
import '../components/style.scss';

class TopNavContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      username: 'me@user.com',
    })
  }

  render() {

    const {
      authStore,
    } = this.props.stores

    return (
      <nav className="navbar">
			<div className="container">
				<div className="navbar-brand">
					<a
						className="navbar-item"
						href="."
					>
						<img src="/scorpicode210.png" size="29px" alt="scorpicode logo" />
					</a>
          <button onClick={() => {authStore.logout()}} type="button">Logout</button>
				</div>
			</div>
		</nav>
    )
  }

}

export default observer(TopNavContainer)
