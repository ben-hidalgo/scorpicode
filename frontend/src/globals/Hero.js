import React from 'react'
import { observer }  from 'mobx-react'
import '../style.scss'

const Hero = () => {

  return (
    <section className="hero is-primary">
      <div className="hero-body">
        <div className="container">
          <h1 className="title">
            Welcome to Scorpicode
          </h1>
          <h2 className="subtitle">
            Please select an option from the menu above
          </h2>
        </div>
      </div>
    </section>
  )
}

export default observer(Hero)
