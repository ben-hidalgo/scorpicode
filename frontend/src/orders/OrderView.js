import React from 'react'
import { useParams } from 'react-router-dom'
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext'

const OrderView = () => {

  let {
    orderStore,
    hatStore,
  } = React.useContext(StoreContext)

  let { id } = useParams()

  let order = orderStore.fetchOrder(id)

  if (!order) {
    return <NotFound id={id} />
  }


  return (
    <div className="container">
      <div className="card">
        <header className="card-header">
          <p className="card-header-title">
            Order: {order.id} Processing {hatStore.list.length}
          </p>
        </header>
        <div className="card-content">
          <div className="content">
            <div className="columns">

              <div className="column">
              <div className="field">
                  <label className="label">Size:</label>
                  {order.size}
                </div>
                <div className="field">
                  <label className="label">Quantity:</label>
                  {order.quantity}
                </div>
                <div className="field">
                  <label className="label">Notes:</label>
                  {order.notes}
                </div>
              </div>

              <div className="column">
                <div className="field">
                  <label className="label">Color:</label>
                  {order.color}
                </div>
                <div className="field">
                  <label className="label">Style:</label>
                  {order.style}
                </div>
                <div className="field">
                  <label className="label">Version:</label>
                  {order.version}
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

const NotFound = (props) => {
  return (
    <div className="container">
      <div className="card">
        <header className="card-header is-danger">
          <p className="card-header-title">
            Order: {props.id} not found
          </p>
        </header>
      </div>
    </div>
  )
} // NotFound

export default observer(OrderView)
